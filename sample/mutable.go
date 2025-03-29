package sample

import (
	"encoding/json"
	"fmt"

	"github.com/bytedance/sonic"
)

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

func (f *MutableFeatures) Len() int {
	return len(f.features)
}

func (f *MutableFeatures) Get(key string) Feature {
	if value, ok := f.features[key]; ok {
		return value
	}
	return nil
}

func (f *MutableFeatures) Set(key string, value Feature) {
	f.features[key] = value
}

func (f *MutableFeatures) MapAny() map[string]any {
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
	return feas
}
func (f *MutableFeatures) MarshalJSON() ([]byte, error) {
	feas := f.MapAny()
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

type ErrorFeature struct{}

func (f *ErrorFeature) Type() DataType {
	return ErrorType
}

func (f *ErrorFeature) GetInt64() (int64, error) {
	return 0, fmt.Errorf("not implemented")
}

func (f *ErrorFeature) GetFloat32() (float32, error) {
	return 0.0, fmt.Errorf("not implemented")
}

func (f *ErrorFeature) GetString() (string, error) {
	return "", fmt.Errorf("not implemented")
}

func (f *ErrorFeature) GetInt64s() ([]int64, error) {
	return nil, fmt.Errorf("not implemented")
}

func (f *ErrorFeature) GetFloat32s() ([]float32, error) {
	return nil, fmt.Errorf("not implemented")
}

func (f *ErrorFeature) GetStrings() ([]string, error) {
	return nil, fmt.Errorf("not implemented")
}

type Int64 struct {
	ErrorFeature
	Value int64
}

func (f *Int64) Type() DataType {
	return Int64Type
}

func (f *Int64) GetInt64() (int64, error) {
	return f.Value, nil
}

type Int64s struct {
	ErrorFeature
	Value []int64
}

func (f *Int64s) Type() DataType {
	return Int64sType
}

func (f *Int64s) GetInt64s() ([]int64, error) {
	return f.Value, nil
}

type Float32 struct {
	ErrorFeature
	Value float32
}

func (f *Float32) Type() DataType {
	return Float32Type
}

func (f *Float32) GetFloat32() (float32, error) {
	return f.Value, nil
}

type Float32s struct {
	ErrorFeature
	Value []float32
}

func (f *Float32s) Type() DataType {
	return Float32sType
}

func (f *Float32s) GetFloat32s() ([]float32, error) {
	return f.Value, nil
}

type String struct {
	ErrorFeature
	Value string
}

func (f *String) Type() DataType {
	return StringType
}

func (f *String) GetString() (string, error) {
	return f.Value, nil
}

type Strings struct {
	ErrorFeature
	Value []string
}

func (f *Strings) Type() DataType {
	return StringsType
}

func (f *Strings) GetStrings() ([]string, error) {
	return f.Value, nil
}
