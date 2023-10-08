package source

import (
	"bufio"
	"os"
	"sort"
	"unsafe"

	"github.com/bytedance/sonic"
	"github.com/uopensail/ulib/prome"
	"github.com/uopensail/ulib/sample"
	"github.com/uopensail/ulib/uno"
	"github.com/uopensail/ulib/zlog"
	"go.uber.org/zap"
)

type wrapper struct {
	features sample.ImmutableFeatures
	id       int
}

type Source struct {
	area        *sample.Arena
	array       []wrapper
	dict        map[string]*wrapper
	collections map[string]Collection
	conditions  map[string]*Condition
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
		array: make([]wrapper, 0, 1024),
		dict:  make(map[string]*wrapper),
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
		w := wrapper{features: *feas, id: index}
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
	return source, nil
}

// if not found, return -1
func (s *Source) GetId(key string) int {
	stat := prome.NewStat("Source.GetId")
	defer stat.End()
	if w, ok := s.dict[key]; ok {
		return w.id
	}
	stat.MarkMiss()
	return -1
}

func (s *Source) GetByKey(key string) *sample.ImmutableFeatures {
	stat := prome.NewStat("Source.GetByKey")
	defer stat.End()
	if w, ok := s.dict[key]; ok {
		return &w.features
	}
	stat.MarkMiss()
	return nil
}

func (s *Source) GetById(id int) *sample.ImmutableFeatures {
	stat := prome.NewStat("Source.GetById")
	defer stat.End()
	if id < 0 || id >= len(s.array) {
		stat.MarkMiss()
		return nil
	}
	return &s.array[id].features
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

func (s *Source) Sort(collection Collection, key string, desc bool) []int {
	stat := prome.NewStat("Source.Sort")
	defer stat.End()
	if len(key) == 0 {
		stat.MarkErr()
		zlog.LOG.Error("sort key is nil")
		return collection
	}
	slice := make([]*wrapper, len(collection))
	for i := 0; i < len(collection); i++ {
		slice[i] = &s.array[collection[i]]
	}
	w := &wrappers{
		slice: slice,
		key:   key,
		desc:  desc,
	}
	if len(key) > 0 {
		sort.Stable(w)
	}

	for i := 0; i < w.Len(); i++ {
		collection[i] = w.slice[i].id
	}

	return collection
}

func (s *Source) Release() {
	stat := prome.NewStat("Source.Release")
	defer stat.End()
	for _, condition := range s.conditions {
		condition.Release()
	}
}
