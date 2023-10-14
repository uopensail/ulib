package items

import (
	"bufio"
	"os"

	"github.com/bytedance/sonic"
	"github.com/uopensail/ulib/datastruct"
	"github.com/uopensail/ulib/prome"
	"github.com/uopensail/ulib/sample"
	"github.com/uopensail/ulib/zlog"
	"go.uber.org/zap"
)

type Features struct {
	Feats sample.ImmutableFeatures
	ID    int
	key   string
}

func (f *Features) Id() int {
	return f.ID
}

func (f *Features) Key() string {
	return f.key
}

func (f *Features) Get(key string) sample.Feature {
	return f.Feats.Get(key)
}

type ItemPool struct {
	area            *sample.Arena
	Array           []Features
	ids             []datastruct.Tuple[string, float32]
	WholeCollection []int
	dict            map[string]*Features
}

func NewItemPool(filepath string, keyField string) (*ItemPool, error) {
	stat := prome.NewStat("NewItemPool")
	defer stat.End()
	file, err := os.Open(filepath)
	if err != nil {
		zlog.LOG.Error("failed to open file", zap.Error(err))
		stat.MarkErr()
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pool := &ItemPool{
		area:  sample.NewArena(),
		Array: make([]Features, 0, 1024),
		dict:  make(map[string]*Features),
		ids:   make([]datastruct.Tuple[string, float32], 0, 1024),
	}

	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		feas := sample.NewImmutableFeatures(pool.area)
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
		w := Features{Feats: *feas, ID: index, key: key}
		pool.Array = append(pool.Array, w)
		pool.ids = append(pool.ids, datastruct.Tuple[string, float32]{First: key, Second: 0.0})
		pool.WholeCollection = append(pool.WholeCollection, w.ID)
		pool.dict[key] = &w
		index++
	}

	if err := scanner.Err(); err != nil {
		zlog.LOG.Error("error while scanning file", zap.Error(err))
		stat.MarkErr()
		return nil, err
	}

	stat.SetCounter(index)

	return pool, nil
}

func (s *ItemPool) GetByKey(key string) *Features {
	stat := prome.NewStat("pool.GetByKey")
	defer stat.End()
	if f, ok := s.dict[key]; ok {
		return f
	}
	stat.MarkMiss()
	return nil
}

func (s *ItemPool) GetById(id int) *Features {
	stat := prome.NewStat("pool.GetById")
	defer stat.End()
	if id < 0 || id >= len(s.Array) {
		stat.MarkMiss()
		return nil
	}
	return &s.Array[id]
}

func (s *ItemPool) Len() int {
	return len(s.Array)
}

func (s *ItemPool) List() []datastruct.Tuple[string, float32] {
	return s.ids
}
