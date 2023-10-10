package source

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"unsafe"

	"github.com/bytedance/sonic"
	"github.com/spf13/cast"
	"github.com/uopensail/ulib/prome"
	"github.com/uopensail/ulib/sample"
	"github.com/uopensail/ulib/uno"
	"github.com/uopensail/ulib/zlog"
	"go.uber.org/zap"
)

type Collection []int

type Features struct {
	features sample.ImmutableFeatures
	id       int
	key      string
}

func (f *Features) Id() int {
	return f.id
}

func (f *Features) Key() string {
	return f.key
}

func (f *Features) Get(key string) sample.Feature {
	return f.features.Get(key)
}

type Index struct {
	data map[string]Collection
}

func (index *Index) Get(key interface{}) Collection {
	if array, ok := index.data[cast.ToString(key)]; ok {
		return array
	}
	return nil
}

type Source struct {
	area        *sample.Arena
	array       []Features
	dict        map[string]*Features
	collections map[string]Collection
	conditions  map[string]*Condition
	indeces     map[string]*Index
}

func NewSource(filepath string, keyField string) (*Source, error) {
	stat := prome.NewStat("NewSource")
	defer stat.End()
	file, err := os.Open(filepath)
	if err != nil {
		zlog.LOG.Error("failed to open file", zap.Error(err))
		stat.MarkErr()
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	source := &Source{
		area:  sample.NewArena(),
		array: make([]Features, 0, 1024),
		dict:  make(map[string]*Features),
	}

	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		feas := sample.NewImmutableFeatures(source.area)
		err = sonic.UnmarshalString(line, feas)
		if err != nil {
			zlog.LOG.Error("unmarshal immutableFeatures error", zap.String("data", line), zap.Error(err))
			continue
		}

		keyFea := feas.Get(keyField)
		if keyFea.Type() != sample.StringType {
			zlog.LOG.Error("key type not string", zap.Any("type", keyFea.Type()))
			continue
		}
		if err != nil {
			zlog.LOG.Error("get key error", zap.String("data", line), zap.Error(err))
			continue
		}
		key, _ := keyFea.GetString()
		w := Features{features: *feas, id: index, key: key}
		source.array = append(source.array, w)
		source.dict[key] = &w
		index++
	}

	if err := scanner.Err(); err != nil {
		zlog.LOG.Error("error while scanning file", zap.Error(err))
		stat.MarkErr()
		return nil, err
	}

	stat.SetCounter(index)
	source.collections = make(map[string]Collection)
	source.conditions = make(map[string]*Condition)
	source.indeces = make(map[string]*Index)
	return source, nil
}

func (s *Source) GetByKey(key string) *Features {
	stat := prome.NewStat("Source.Get")
	defer stat.End()
	if f, ok := s.dict[key]; ok {
		return f
	}
	stat.MarkMiss()
	return nil
}

func (s *Source) GetById(id int) *Features {
	stat := prome.NewStat("Source.GetById")
	defer stat.End()
	if id < 0 || id >= len(s.array) {
		stat.MarkMiss()
		return nil
	}
	return &s.array[id]
}

func (s *Source) Len() int {
	return len(s.array)
}

func (s *Source) BuildCollection(name string, condition string) {
	stat := prome.NewStat("Source.BuildCollection")
	defer stat.End()
	evaluator, err := uno.NewEvaluator(condition)
	if err != nil {
		zlog.LOG.Error("build collection condition error", zap.Error(err))
		stat.MarkErr()
		return
	}
	defer evaluator.Release()

	var status int32
	ret := make([]int, 0, 1024)
	for i := 0; i < len(s.array); i++ {
		slice := evaluator.Allocate()
		evaluator.Fill(&s.array[i].features, slice)
		status = evaluator.Eval(slice)
		if status == 1 {
			ret = append(ret, s.array[i].id)
		}
	}
	stat.SetCounter(len(ret))
	s.collections[name] = ret
	return
}

func (s *Source) GetCollection(name string) Collection {
	stat := prome.NewStat("Source.GetCollection")
	defer stat.End()
	if collection, ok := s.collections[name]; ok {
		return collection
	}
	stat.MarkMiss()
	return nil
}

func (s *Source) BuildCondition(name string, condition string) {
	stat := prome.NewStat("Source.BuildCondition")
	defer stat.End()
	evaluator, err := uno.NewEvaluator(condition)
	if err != nil {
		zlog.LOG.Error("build condition error", zap.Error(err))
		stat.MarkErr()
		return
	}
	c := &Condition{
		evaluator: evaluator,
		slices:    make([][]unsafe.Pointer, len(s.array)),
	}

	for i := 0; i < len(s.array); i++ {
		slice := evaluator.Allocate()
		evaluator.Fill(&s.array[i].features, slice)
		evaluator.PreEval(slice)
		c.slices[i] = slice
	}
	s.conditions[name] = c
}

func (s *Source) GetCondition(name string) *Condition {
	stat := prome.NewStat("Source.GetCondition")
	defer stat.End()
	if condition, ok := s.conditions[name]; ok {
		return condition
	}
	stat.MarkMiss()
	return nil
}

func (s *Source) Sort(collection Collection, key string, desc bool) Collection {
	stat := prome.NewStat("Source.Sort")
	defer stat.End()
	if len(key) == 0 {
		stat.MarkErr()
		zlog.LOG.Error("sort key is nil")
		return collection
	}

	ret := make(Collection, len(collection))
	copy(ret, collection)
	sort.Slice(ret, func(i, j int) bool {
		return less(&s.array[i].features, &s.array[j].features, key, desc)
	})

	return ret
}

func (s *Source) GetIndex(name string) *Index {
	stat := prome.NewStat("Source.GetIndex")
	defer stat.End()
	if index, ok := s.indeces[name]; ok {
		return index
	}
	stat.MarkMiss()
	return nil
}

func (s *Source) BuildIndex(name string, column string) {
	stat := prome.NewStat("Source.BuildIndex")
	defer stat.End()

	// generate func map
	funcs := make(map[sample.DataType]func(feature sample.Feature) []string)

	funcs[sample.Float32Type] = func(feature sample.Feature) []string {
		val, _ := feature.GetFloat32()
		return []string{cast.ToString(val)}
	}

	funcs[sample.Float32sType] = func(feature sample.Feature) []string {
		vals, _ := feature.GetFloat32s()
		ret := make([]string, len(vals))
		for i := 0; i < len(vals); i++ {
			ret[i] = cast.ToString(vals[i])
		}
		return ret
	}

	funcs[sample.Int64Type] = func(feature sample.Feature) []string {
		val, _ := feature.GetInt64()
		return []string{cast.ToString(val)}
	}

	funcs[sample.Int64sType] = func(feature sample.Feature) []string {
		vals, _ := feature.GetInt64s()
		ret := make([]string, len(vals))
		for i := 0; i < len(vals); i++ {
			ret[i] = cast.ToString(vals[i])
		}
		return ret
	}

	funcs[sample.StringType] = func(feature sample.Feature) []string {
		val, _ := feature.GetString()
		return []string{val}
	}

	funcs[sample.StringsType] = func(feature sample.Feature) []string {
		vals, _ := feature.GetStrings()
		return vals
	}

	foo := funcs[s.array[0].Get(column).Type()]
	dict := make(map[string]Collection)
	for i := 0; i < s.Len(); i++ {
		strs := foo(s.array[i].Get(column))
		for j := 0; j < len(strs); i++ {
			if list, ok := dict[strs[j]]; ok {
				list = append(list, i)
				dict[strs[j]] = list
			} else {
				dict[strs[j]] = []int{i}
			}
		}
	}

	s.indeces[name] = &Index{
		data: dict,
	}
}

func (s *Source) Release() {
	stat := prome.NewStat("Source.Release")
	defer stat.End()
	for _, condition := range s.conditions {
		condition.Release()
	}
}

func less(feasA, feasB sample.Features, key string, desc bool) bool {
	left := feasA.Get(key)
	right := feasB.Get(key)

	if left == nil || right == nil {
		return false
	}

	dtype := left.Type()
	switch dtype {
	case sample.Float32Type:
		lv, err1 := left.GetFloat32()
		rv, err2 := right.GetFloat32()
		if err1 != nil || err2 != nil {
			return false
		}
		return !((lv < rv) && desc)
	case sample.Int64Type:
		lv, err1 := left.GetInt64()
		rv, err2 := right.GetInt64()
		if err1 != nil || err2 != nil {
			return false
		}
		return !((lv < rv) && desc)
	case sample.StringType:
		lv, err1 := left.GetString()
		rv, err2 := right.GetString()
		if err1 != nil || err2 != nil {
			return false
		}
		return !((lv < rv) && desc)
	default:
		panic(fmt.Sprintf("data type: %d not support", dtype))
	}

}
