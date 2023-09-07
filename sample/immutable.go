package sample

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unsafe"

	"github.com/bytedance/sonic"
)

const int64Size uintptr = unsafe.Sizeof(int64(0))
const float32Size uintptr = unsafe.Sizeof(float32(0))
const stringHeaderSize uintptr = unsafe.Sizeof(reflect.StringHeader{Data: 0, Len: 0})
const sliceHeaderSize uintptr = unsafe.Sizeof(reflect.SliceHeader{Data: 0, Len: 0, Cap: 0})

type ImmutableFeatures struct {
	arena    *Arena
	features map[string]uintptr
}

func NewImmutableFeatures(arena *Arena) *ImmutableFeatures {
	return &ImmutableFeatures{
		arena:    arena,
		features: make(map[string]uintptr),
	}
}

func (f *ImmutableFeatures) GetType(key string) DataType {
	if addr, ok := f.features[key]; ok {
		return *(*DataType)(uintptr2Pointer(addr))
	}
	return ErrorType
}

func (f *ImmutableFeatures) Keys() []string {
	ret := make([]string, 0, len(f.features))
	for key := range f.features {
		ret = append(ret, key)
	}
	return ret
}

func (f *ImmutableFeatures) GetInt64(key string) (int64, error) {
	if addr, ok := f.features[key]; ok {
		return getInt64(addr)
	}
	return 0, fmt.Errorf("key: %s not found", key)
}

func (f *ImmutableFeatures) GetFloat32(key string) (float32, error) {
	if addr, ok := f.features[key]; ok {
		return getFloat32(addr)
	}
	return 0.0, fmt.Errorf("key: %s not found", key)
}

func (f *ImmutableFeatures) GetString(key string) (string, error) {
	if addr, ok := f.features[key]; ok {
		return getString(addr)
	}
	return "", fmt.Errorf("key: %s not found", key)
}

func (f *ImmutableFeatures) GetInt64s(key string) ([]int64, error) {
	if addr, ok := f.features[key]; ok {
		return getInt64s(addr)
	}
	return nil, fmt.Errorf("key: %s not found", key)
}

func (f *ImmutableFeatures) GetFloat32s(key string) ([]float32, error) {
	if addr, ok := f.features[key]; ok {
		return getFloat32s(addr)
	}
	return nil, fmt.Errorf("key: %s not found", key)
}

func (f *ImmutableFeatures) GetStrings(key string) ([]string, error) {
	if addr, ok := f.features[key]; ok {
		return getStrings(addr)
	}
	return nil, fmt.Errorf("key: %s not found", key)
}

func (f *ImmutableFeatures) MarshalJSON() ([]byte, error) {
	feas := make(map[string]interface{})

	for key, addr := range f.features {
		dtype := *(*DataType)(uintptr2Pointer(addr))
		switch dtype {
		case Int64Type:
			v, _ := getInt64(addr)
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value int64    `json:"value"`
			}{
				Int64Type,
				v,
			}
		case Float32Type:
			v, _ := getFloat32(addr)
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value float32  `json:"value"`
			}{
				Float32Type,
				v,
			}
		case StringType:
			v, _ := getString(addr)
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value string   `json:"value"`
			}{
				StringType,
				v,
			}
		case Int64sType:
			v, _ := getInt64s(addr)
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value []int64  `json:"value"`
			}{
				Int64sType,
				v,
			}
		case Float32sType:
			v, _ := getFloat32s(addr)
			feas[key] = struct {
				Type  DataType  `json:"type"`
				Value []float32 `json:"value"`
			}{
				Float32sType,
				v,
			}
		case StringsType:
			v, _ := getStrings(addr)
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value []string `json:"value"`
			}{
				StringsType,
				v,
			}
		}
	}
	return sonic.Marshal(feas)
}

func (f *ImmutableFeatures) Mutable() *MutableFeatures {
	ret := NewMutableFeatures()
	for key, addr := range f.features {
		dtype := *(*DataType)(uintptr2Pointer(addr))
		switch dtype {
		case Int64Type:
			v, _ := getInt64(addr)
			ret.features[key] = &Int64{Value: v}
		case Float32Type:
			v, _ := getFloat32(addr)
			ret.features[key] = &Float32{Value: v}
		case StringType:
			v, _ := getString(addr)
			ret.features[key] = &String{Value: deepcopyOfString(v)}
		case Int64sType:
			v, _ := getInt64s(addr)
			nums := make([]int64, len(v))
			copy(nums, v)
			ret.features[key] = &Int64s{Value: nums}
		case Float32sType:
			v, _ := getFloat32s(addr)
			nums := make([]float32, len(v))
			copy(nums, v)
			ret.features[key] = &Float32s{Value: nums}
		case StringsType:
			v, _ := getStrings(addr)
			ret.features[key] = &Strings{Value: deepcpyOfStrings(v)}
		}
	}
	return ret
}

func (f *ImmutableFeatures) UnmarshalJSON(data []byte) error {
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
			f.features[key] = putInt64(num, f.arena)
		case Float32Type:
			var num float32
			sonic.Unmarshal(value.Value, &num)
			f.features[key] = putFloat32(num, f.arena)
		case StringType:
			var str string
			sonic.Unmarshal(value.Value, &str)
			f.features[key] = putString(str, f.arena)
		case Int64sType:
			var nums []int64
			sonic.Unmarshal(value.Value, &nums)
			f.features[key] = putInt64s(nums, f.arena)
		case Float32sType:
			var nums []float32
			sonic.Unmarshal(value.Value, &nums)
			f.features[key] = putFloat32s(nums, f.arena)
		case StringsType:
			var strs []string
			sonic.Unmarshal(value.Value, &strs)
			f.features[key] = putStrings(strs, f.arena)
		}
	}
	return nil
}

func putInt64(value int64, arena *Arena) uintptr {
	data := arena.allocate(8 + int64Size)
	data[0] = byte(Int64Type)
	*(*int64)(unsafe.Pointer(&data[8])) = value
	return uintptr(unsafe.Pointer(&data[0]))
}

func getInt64(addr uintptr) (int64, error) {
	if *(*DataType)(uintptr2Pointer(addr)) != Int64Type {
		return 0, fmt.Errorf("type mismatch")
	}
	return *(*int64)(uintptr2Pointer(addr + 8)), nil
}

func putFloat32(value float32, arena *Arena) uintptr {
	data := arena.allocate(4 + float32Size)
	data[0] = byte(Float32Type)
	*(*float32)(unsafe.Pointer(&data[4])) = value
	return uintptr(unsafe.Pointer(&data[0]))
}

func getFloat32(addr uintptr) (float32, error) {
	if *(*DataType)(uintptr2Pointer(addr)) != Float32Type {
		return 0.0, fmt.Errorf("type mismatch")
	}
	return *(*float32)(uintptr2Pointer(addr + 4)), nil
}

func putString(value string, arena *Arena) uintptr {
	size := 8 + stringHeaderSize + uintptr(len(value))
	size = ((size + 7) >> 3) << 3
	data := arena.allocate(size)
	data[0] = byte(StringType)
	header := (*reflect.StringHeader)(unsafe.Pointer(&data[8]))
	header.Data = uintptr(unsafe.Pointer(&data[8+stringHeaderSize]))
	header.Len = len(value)
	copy(data[8+stringHeaderSize:], *(*[]byte)(unsafe.Pointer(&value)))
	return uintptr(unsafe.Pointer(&data[0]))
}

func getString(addr uintptr) (string, error) {
	if *(*DataType)(uintptr2Pointer(addr)) != StringType {
		return "", fmt.Errorf("type mismatch")
	}
	return *(*string)(uintptr2Pointer(addr + 8)), nil
}

func putInt64s(arr []int64, arena *Arena) uintptr {
	data := arena.allocate(8 + sizeofInt64s(arr))
	data[0] = byte(Int64sType)
	packInts(arr, data[8:])
	return uintptr(unsafe.Pointer(&data[0]))
}

func getInt64s(addr uintptr) ([]int64, error) {
	if *(*DataType)(uintptr2Pointer(addr)) != Int64sType {
		return nil, fmt.Errorf("type mismatch")
	}
	return unpackInts(addr + 8), nil
}

func putFloat32s(arr []float32, arena *Arena) uintptr {
	data := arena.allocate(8 + sizeofFloat32s(arr))
	data[0] = byte(Float32sType)
	packFloats(arr, data[8:])
	return uintptr(unsafe.Pointer(&data[0]))
}

func getFloat32s(addr uintptr) ([]float32, error) {
	if *(*DataType)(uintptr2Pointer(addr)) != Float32sType {
		return nil, fmt.Errorf("type mismatch")
	}
	return unpackFloats(addr + 8), nil
}

func putStrings(arr []string, arena *Arena) uintptr {
	data := arena.allocate(8 + sizeofStrings(arr))
	data[0] = byte(StringsType)
	packStrs(arr, data[8:])
	return uintptr(unsafe.Pointer(&data[0]))
}

func getStrings(addr uintptr) ([]string, error) {
	if *(*DataType)(uintptr2Pointer(addr)) != StringsType {
		return nil, fmt.Errorf("type mismatch")
	}
	return unpackStrs(addr + 8), nil
}

func packInts(arr []int64, data []byte) {
	p := uintptr(unsafe.Pointer(&data[sliceHeaderSize]))
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&data[0]))
	slice.Data = p
	slice.Cap = len(arr)
	slice.Len = len(arr)
	copy(*(*[]int64)(unsafe.Pointer(&data[0])), arr)
}

func unpackInts(addr uintptr) []int64 {
	return *(*[]int64)(uintptr2Pointer(addr))
}

func sizeofInt64s(arr []int64) uintptr {
	return sliceHeaderSize + int64Size*uintptr(len(arr))
}

func packFloats(arr []float32, data []byte) {
	p := uintptr(unsafe.Pointer(&data[sliceHeaderSize]))
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&data[0]))
	slice.Data = p
	slice.Cap = len(arr)
	slice.Len = len(arr)
	copy(*(*[]float32)(unsafe.Pointer(&data[0])), arr)
}

func unpackFloats(addr uintptr) []float32 {
	return *(*[]float32)(uintptr2Pointer(addr))
}

func sizeofFloat32s(arr []float32) uintptr {
	return ((sliceHeaderSize + float32Size*uintptr(len(arr)) + 7) >> 3) << 3
}

func packStrs(arr []string, data []byte) {
	p := uintptr(unsafe.Pointer(&data[sliceHeaderSize]))
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&data[0]))
	slice.Data = p
	slice.Cap = len(arr)
	slice.Len = len(arr)
	headerOffset := sliceHeaderSize
	dataOffset := headerOffset + stringHeaderSize*uintptr(len(arr))
	var str *reflect.StringHeader
	var size uintptr
	for i := 0; i < len(arr); i++ {
		size = uintptr(len(arr[i]))
		str = (*reflect.StringHeader)(unsafe.Pointer(&data[headerOffset]))
		str.Data = uintptr(unsafe.Pointer(&data[dataOffset]))
		str.Len = len(arr[i])
		copy(data[dataOffset:dataOffset+size], *(*[]byte)(unsafe.Pointer(&arr[i])))
		dataOffset += size
		headerOffset += stringHeaderSize
	}
}

func unpackStrs(addr uintptr) []string {
	return *(*[]string)(uintptr2Pointer(addr))
}

func sizeofStrings(arr []string) uintptr {
	size := sliceHeaderSize + stringHeaderSize*uintptr(len(arr))
	for i := 0; i < len(arr); i++ {
		size += uintptr(len(arr[i]))
	}
	return ((size + 7) >> 3) << 3
}

func deepcpyOfStrings(arr []string) []string {
	ret := make([]string, len(arr))
	for i := 0; i < len(arr); i++ {
		ret[i] = deepcopyOfString(arr[i])
	}
	return ret
}

func deepcopyOfString(s string) (str string) {
	data := make([]byte, len(s))
	copy(data, *(*[]byte)(unsafe.Pointer(&s)))
	return string(data)
}

func uintptr2Pointer(addr uintptr) unsafe.Pointer {
	return unsafe.Pointer(addr)
}
