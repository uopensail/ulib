package sample

import (
	"encoding/json"
	"fmt"

	"github.com/bytedance/sonic"
)

type Feature interface {
	Type() DataType
}

type MutableFeatures struct {
	features map[string]Feature
}

func NewMutableFeatures() *MutableFeatures {
	return &MutableFeatures{
		features: make(map[string]Feature),
	}
}

func (f *MutableFeatures) GetType(key string) DataType {
	if feature, ok := f.features[key]; ok {
		return feature.Type()
	}
	return ErrorType
}

func (f *MutableFeatures) Keys() []string {
	ret := make([]string, 0, len(f.features))
	for key := range f.features {
		ret = append(ret, key)
	}
	return ret
}

func (f *MutableFeatures) GetInt64(key string) (int64, error) {
	if value, ok := f.features[key]; ok {
		if value.Type() != Int64Type {
			return 0, fmt.Errorf("type mismatch")
		}
		return value.(*Int64).Value, nil
	}
	return 0, fmt.Errorf("key: %s not found", key)
}

func (f *MutableFeatures) GetFloat32(key string) (float32, error) {
	if value, ok := f.features[key]; ok {
		if value.Type() != Float32Type {
			return 0.0, fmt.Errorf("type mismatch")
		}
		return value.(*Float32).Value, nil
	}
	return 0.0, fmt.Errorf("key: %s not found", key)
}

func (f *MutableFeatures) GetString(key string) (string, error) {
	if value, ok := f.features[key]; ok {
		if value.Type() != StringType {
			return "", fmt.Errorf("type mismatch")
		}

		return value.(*String).Value, nil
	}
	return "", fmt.Errorf("key: %s not found", key)
}

func (f *MutableFeatures) GetInt64s(key string) ([]int64, error) {
	if value, ok := f.features[key]; ok {
		if value.Type() != Int64sType {
			return nil, fmt.Errorf("type mismatch")
		}
		return value.(*Int64s).Value, nil
	}
	return nil, fmt.Errorf("key: %s not found", key)
}

func (f *MutableFeatures) GetFloat32s(key string) ([]float32, error) {
	if value, ok := f.features[key]; ok {
		if value.Type() != Float32sType {
			return nil, fmt.Errorf("type mismatch")
		}
		return value.(*Float32s).Value, nil
	}
	return nil, fmt.Errorf("key: %s not found", key)
}

func (f *MutableFeatures) GetStrings(key string) ([]string, error) {
	if value, ok := f.features[key]; ok {
		if value.Type() != StringsType {
			return nil, fmt.Errorf("type mismatch")
		}
		return value.(*Strings).Value, nil
	}
	return nil, fmt.Errorf("key: %s not found", key)
}

func (f *MutableFeatures) MarshalJSON() ([]byte, error) {
	feas := make(map[string]interface{})
	for key, value := range f.features {
		switch value.Type() {
		case Int64Type:
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value int64    `json:"value"`
			}{
				Int64Type,
				value.(*Int64).Value,
			}
		case Float32Type:
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value float32  `json:"value"`
			}{
				Float32Type,
				value.(*Float32).Value,
			}
		case StringType:
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value string   `json:"value"`
			}{
				StringType,
				value.(*String).Value,
			}
		case Int64sType:
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value []int64  `json:"value"`
			}{
				Int64sType,
				value.(*Int64s).Value,
			}
		case Float32sType:
			feas[key] = struct {
				Type  DataType  `json:"type"`
				Value []float32 `json:"value"`
			}{
				Float32sType,
				value.(*Float32s).Value,
			}
		case StringsType:
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value []string `json:"value"`
			}{
				StringsType,
				value.(*Strings).Value,
			}
		}
	}
	return sonic.Marshal(feas)
}

func (f *MutableFeatures) UnmarshalJSON(data []byte) error {
	type Fea struct {
		Type  DataType        `json:"type"`
		Value json.RawMessage `json:"value"`
	}

	var fea map[string]Fea
	err := sonic.Unmarshal(data, &fea)
	if err != nil {
		return err
	}

	for key, value := range fea {
		switch value.Type {
		case Int64Type:
			var num int64
			sonic.Unmarshal(value.Value, &num)
			f.features[key] = &Int64{Value: num}
		case Float32Type:
			var num float32
			sonic.Unmarshal(value.Value, &num)
			f.features[key] = &Float32{Value: num}
		case StringType:
			var str string
			sonic.Unmarshal(value.Value, &str)
			f.features[key] = &String{Value: str}
		case Int64sType:
			var nums []int64
			sonic.Unmarshal(value.Value, &nums)
			f.features[key] = &Int64s{Value: nums}
		case Float32sType:
			var nums []float32
			sonic.Unmarshal(value.Value, &nums)
			f.features[key] = &Float32s{Value: nums}
		case StringsType:
			var strs []string
			sonic.Unmarshal(value.Value, &strs)
			f.features[key] = &Strings{Value: strs}
		}
	}
	return nil
}

type Int64 struct {
	Value int64
}

func (f *Int64) Type() DataType {
	return Int64Type
}

type Int64s struct {
	Value []int64
}

func (f *Int64s) Type() DataType {
	return Int64sType
}

type Float32 struct {
	Value float32
}

func (f *Float32) Type() DataType {
	return Float32Type
}

type Float32s struct {
	Value []float32
}

func (f *Float32s) Type() DataType {
	return Float32sType
}

type String struct {
	Value string
}

func (f *String) Type() DataType {
	return StringType
}

type Strings struct {
	Value []string
}

func (f *Strings) Type() DataType {
	return StringsType
}
