package source

import (
	"bufio"
	"os"

	"github.com/bytedance/sonic"
	"github.com/uopensail/ulib/sample"
	"github.com/uopensail/ulib/zlog"
	"go.uber.org/zap"
)

type Collection []*sample.ImmutableFeatures

type Source struct {
	area        *sample.Arena
	array       []sample.ImmutableFeatures
	dict        map[string]*sample.ImmutableFeatures
	collections map[string]Collection
}

func NewSource(filepath string, keyField string) (*Source, error) {
	file, err := os.Open(filepath)
	if err != nil {
		zlog.LOG.Error("failed to open file", zap.Error(err))
		return nil, err
	}
	defer file.Close()
	index := 0
	scanner := bufio.NewScanner(file)
	source := &Source{
		area:  sample.NewArena(),
		array: make([]sample.ImmutableFeatures, 1024),
		dict:  make(map[string]*sample.ImmutableFeatures),
	}

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
		}
		if err != nil {
			zlog.LOG.Error("get key error", zap.String("data", line), zap.Error(err))
			continue
		}
		key, _ := keyFea.GetString()
		source.array = append(source.array, *feas)
		source.dict[key] = feas
		index++
	}
	if err := scanner.Err(); err != nil {
		zlog.LOG.Error("error while scanning file", zap.Error(err))
		return nil, err
	}
	return source, nil
}

func (s *Source) GetByKey(key string) *sample.ImmutableFeatures {
	if feas, ok := s.dict[key]; ok {
		return feas
	}
	return nil
}

func (s *Source) GetByIndex(index int) *sample.ImmutableFeatures {
	if index < 0 || index >= len(s.array) {
		return nil
	}
	return &s.array[index]
}

func (s *Source) BuildCollection(name string, condition string) {

}
