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
	Features map[string]Feature
}

func NewMutableFeatures() *MutableFeatures {
	return &MutableFeatures{
		Features: make(map[string]Feature),
	}
}

func (f *MutableFeatures) GetInt64(key string) (int64, error) {
	if value, ok := f.Features[key]; ok {
		if value.Type() != Int64Type {
			return 0, fmt.Errorf("type mismatch")
		}
		return value.(*Int64).Value, nil
	}
	return 0, fmt.Errorf("key: %s not found", key)
}

func (f *MutableFeatures) GetFloat32(key string) (float32, error) {
	if value, ok := f.Features[key]; ok {
		if value.Type() != Float32Type {
			return 0.0, fmt.Errorf("type mismatch")
		}
		return value.(*Float32).Value, nil
	}
	return 0.0, fmt.Errorf("key: %s not found", key)
}

func (f *MutableFeatures) GetString(key string) (string, error) {
	if value, ok := f.Features[key]; ok {
		if value.Type() != StringType {
			return "", fmt.Errorf("type mismatch")
		}
		return value.(*String).Value, nil
	}
	return "", fmt.Errorf("key: %s not found", key)
}

func (f *MutableFeatures) GetInt64Array(key string) ([]int64, error) {
	if value, ok := f.Features[key]; ok {
		if value.Type() != Int64ArrayType {
			return nil, fmt.Errorf("type mismatch")
		}
		return value.(*Int64Array).Value, nil
	}
	return nil, fmt.Errorf("key: %s not found", key)
}

func (f *MutableFeatures) GetFloat32Array(key string) ([]float32, error) {
	if value, ok := f.Features[key]; ok {
		if value.Type() != Float32ArrayType {
			return nil, fmt.Errorf("type mismatch")
		}
		return value.(*Float32Array).Value, nil
	}
	return nil, fmt.Errorf("key: %s not found", key)
}

func (f *MutableFeatures) GetStringArray(key string) ([]string, error) {
	if value, ok := f.Features[key]; ok {
		if value.Type() != StringArrayType {
			return nil, fmt.Errorf("type mismatch")
		}
		return value.(*StringArray).Value, nil
	}
	return nil, fmt.Errorf("key: %s not found", key)
}

func (f *MutableFeatures) MarshalJSON() ([]byte, error) {
	feas := make(map[string]interface{})
	for key, value := range f.Features {
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
		case Int64ArrayType:
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value []int64  `json:"value"`
			}{
				Int64ArrayType,
				value.(*Int64Array).Value,
			}
		case Float32ArrayType:
			feas[key] = struct {
				Type  DataType  `json:"type"`
				Value []float32 `json:"value"`
			}{
				Float32ArrayType,
				value.(*Float32Array).Value,
			}
		case StringArrayType:
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value []string `json:"value"`
			}{
				StringArrayType,
				value.(*StringArray).Value,
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
			f.Features[key] = &Int64{Value: num}
		case Float32Type:
			var num float32
			sonic.Unmarshal(value.Value, &num)
			f.Features[key] = &Float32{Value: num}
		case StringType:
			var str string
			sonic.Unmarshal(value.Value, &str)
			f.Features[key] = &String{Value: str}
		case Int64ArrayType:
			var nums []int64
			sonic.Unmarshal(value.Value, &nums)
			f.Features[key] = &Int64Array{Value: nums}
		case Float32ArrayType:
			var nums []float32
			sonic.Unmarshal(value.Value, &nums)
			f.Features[key] = &Float32Array{Value: nums}
		case StringArrayType:
			var strs []string
			sonic.Unmarshal(value.Value, &strs)
			f.Features[key] = &StringArray{Value: strs}
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

type Int64Array struct {
	Value []int64
}

func (f *Int64Array) Type() DataType {
	return Int64ArrayType
}

type Float32 struct {
	Value float32
}

func (f *Float32) Type() DataType {
	return Float32Type
}

type Float32Array struct {
	Value []float32
}

func (f *Float32Array) Type() DataType {
	return Float32ArrayType
}

type String struct {
	Value string
}

func (f *String) Type() DataType {
	return StringType
}

type StringArray struct {
	Value []string
}

func (f *StringArray) Type() DataType {
	return StringArrayType
}
