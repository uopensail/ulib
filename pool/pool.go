package pool

import (
	"bufio"
	"os"
	"strings"

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

type Pool struct {
	area            *sample.Arena
	Array           []Features
	ids             []datastruct.Tuple[string, float32]
	WholeCollection []int
	dict            map[string]*Features
}

func NewPool(filepath string) (*Pool, error) {
	stat := prome.NewStat("NewPool")
	defer stat.End()
	file, err := os.Open(filepath)
	if err != nil {
		zlog.LOG.Error("failed to open file", zap.Error(err))
		stat.MarkErr()
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pool := &Pool{
		area:  sample.NewArena(),
		Array: make([]Features, 0, 1024),
		dict:  make(map[string]*Features),
		ids:   make([]datastruct.Tuple[string, float32], 0, 1024),
	}

	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		ss := strings.Split(line, "\t")
		if len(ss) != 2 {
			zlog.LOG.Warn("ingore line", zap.String("line", line))
			continue
		}
		feas := sample.NewImmutableFeatures(pool.area)
		err = sonic.UnmarshalString(ss[1], feas)
		if err != nil {
			zlog.LOG.Error("unmarshal immutableFeatures error", zap.String("data", line), zap.Error(err))
			continue
		}

		key := ss[0]

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

func (s *Pool) GetByKey(key string) *Features {
	stat := prome.NewStat("pool.GetByKey")
	defer stat.End()
	if f, ok := s.dict[key]; ok {
		return f
	}
	stat.MarkMiss()
	return nil
}

func (s *Pool) GetById(id int) *Features {
	stat := prome.NewStat("pool.GetById")
	defer stat.End()
	if id < 0 || id >= len(s.Array) {
		stat.MarkMiss()
		return nil
	}
	return &s.Array[id]
}

func (s *Pool) Len() int {
	return len(s.Array)
}

func (s *Pool) List() []datastruct.Tuple[string, float32] {
	return s.ids
}
