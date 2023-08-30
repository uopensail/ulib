package sample

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unsafe"

	"github.com/bytedance/sonic"
)

const uintptrSize uintptr = unsafe.Sizeof(uintptr(0))
const intSize uintptr = unsafe.Sizeof(int(0))
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

func (f *ImmutableFeatures) GetInt64Array(key string) ([]int64, error) {
	if addr, ok := f.features[key]; ok {
		return getInt64Array(addr)
	}
	return nil, fmt.Errorf("key: %s not found", key)
}

func (f *ImmutableFeatures) GetFloat32Array(key string) ([]float32, error) {
	if addr, ok := f.features[key]; ok {
		return getFloat32Array(addr)
	}
	return nil, fmt.Errorf("key: %s not found", key)
}

func (f *ImmutableFeatures) GetStringArray(key string) ([]string, error) {
	if addr, ok := f.features[key]; ok {
		return getStringArray(addr)
	}
	return nil, fmt.Errorf("key: %s not found", key)
}

func (f *ImmutableFeatures) MarshalJSON() ([]byte, error) {
	feas := make(map[string]interface{})

	for key, addr := range f.features {
		dtype := *(*DataType)(unsafe.Pointer(addr))
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
		case Int64ArrayType:
			v, _ := getInt64Array(addr)
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value []int64  `json:"value"`
			}{
				Int64ArrayType,
				v,
			}
		case Float32ArrayType:
			v, _ := getFloat32Array(addr)
			feas[key] = struct {
				Type  DataType  `json:"type"`
				Value []float32 `json:"value"`
			}{
				Float32ArrayType,
				v,
			}
		case StringArrayType:
			v, _ := getStringArray(addr)
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value []string `json:"value"`
			}{
				StringArrayType,
				v,
			}
		}
	}
	return sonic.Marshal(feas)
}

func (f *ImmutableFeatures) Mutable() *MutableFeatures {
	ret := NewMutableFeatures()
	for key, addr := range f.features {
		dtype := *(*DataType)(unsafe.Pointer(addr))
		switch dtype {
		case Int64Type:
			v, _ := getInt64(addr)
			ret.Features[key] = &Int64{Value: v}
		case Float32Type:
			v, _ := getFloat32(addr)
			ret.Features[key] = &Float32{Value: v}
		case StringType:
			v, _ := getString(addr)
			ret.Features[key] = &String{Value: deepcopyOfString(v)}
		case Int64ArrayType:
			v, _ := getInt64Array(addr)
			nums := make([]int64, len(v))
			copy(nums, v)
			ret.Features[key] = &Int64Array{Value: nums}
		case Float32ArrayType:
			v, _ := getFloat32Array(addr)
			nums := make([]float32, len(v))
			copy(nums, v)
			ret.Features[key] = &Float32Array{Value: nums}
		case StringArrayType:
			v, _ := getStringArray(addr)
			ret.Features[key] = &StringArray{Value: deepcpyOfStringArray(v)}
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
		case Int64ArrayType:
			var nums []int64
			sonic.Unmarshal(value.Value, &nums)
			f.features[key] = putInt64Array(nums, f.arena)
		case Float32ArrayType:
			var nums []float32
			sonic.Unmarshal(value.Value, &nums)
			f.features[key] = putFloat32Array(nums, f.arena)
		case StringArrayType:
			var strs []string
			sonic.Unmarshal(value.Value, &strs)
			f.features[key] = putStringArray(strs, f.arena)
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
	if *(*byte)(unsafe.Pointer(addr)) != byte(Int64Type) {
		return 0, fmt.Errorf("type mismatch")
	}
	return *(*int64)(unsafe.Pointer(addr + 8)), nil
}

func putFloat32(value float32, arena *Arena) uintptr {
	data := arena.allocate(4 + float32Size)
	data[0] = byte(Float32Type)
	*(*float32)(unsafe.Pointer(&data[4])) = value
	return uintptr(unsafe.Pointer(&data[0]))
}

func getFloat32(addr uintptr) (float32, error) {
	if *(*byte)(unsafe.Pointer(addr)) != byte(Float32Type) {
		return 0.0, fmt.Errorf("type mismatch")
	}
	return *(*float32)(unsafe.Pointer(addr + 4)), nil
}

func putString(value string, arena *Arena) uintptr {
	size := 8 + stringHeaderSize + uintptr(len(value))
	if size&7 != 0 {
		size = size - size&7 + 8
	}
	data := arena.allocate(size)
	data[0] = byte(StringType)
	*(*uintptr)(unsafe.Pointer(&data[8])) = uintptr(unsafe.Pointer(&data[8+stringHeaderSize]))
	*(*int)(unsafe.Pointer(&data[8+uintptrSize])) = len(value)
	copy(data[8+stringHeaderSize:], *(*[]byte)(unsafe.Pointer(&value)))
	return uintptr(unsafe.Pointer(&data[0]))
}

func getString(addr uintptr) (string, error) {
	if *(*byte)(unsafe.Pointer(addr)) != byte(StringType) {
		return "", fmt.Errorf("type mismatch")
	}
	return *(*string)(unsafe.Pointer(addr + 8)), nil
}

func putInt64Array(arr []int64, arena *Arena) uintptr {
	data := arena.allocate(8 + sizeofInt64Array(arr))
	data[0] = byte(Int64ArrayType)
	copyInt64Array(arr, data[8:])
	return uintptr(unsafe.Pointer(&data[0]))
}

func getInt64Array(addr uintptr) ([]int64, error) {
	if *(*byte)(unsafe.Pointer(addr)) != byte(Int64ArrayType) {
		return nil, fmt.Errorf("type mismatch")
	}
	return toInt64Array(addr + 8), nil
}

func putFloat32Array(arr []float32, arena *Arena) uintptr {
	data := arena.allocate(8 + sizeofFloat32Array(arr))
	data[0] = byte(Float32ArrayType)
	copyFloat32Array(arr, data[8:])
	return uintptr(unsafe.Pointer(&data[0]))
}

func getFloat32Array(addr uintptr) ([]float32, error) {
	if *(*byte)(unsafe.Pointer(addr)) != byte(Float32ArrayType) {
		return nil, fmt.Errorf("type mismatch")
	}
	return toFloat32Array(addr + 8), nil
}

func putStringArray(arr []string, arena *Arena) uintptr {
	data := arena.allocate(8 + sizeofStringArray(arr))
	data[0] = byte(StringArrayType)
	copyStringArray(arr, data[8:])
	return uintptr(unsafe.Pointer(&data[0]))
}

func getStringArray(addr uintptr) ([]string, error) {
	if *(*byte)(unsafe.Pointer(addr)) != byte(StringArrayType) {
		return nil, fmt.Errorf("type mismatch")
	}
	return toStringArray(addr + 8), nil
}

func copyInt64Array(arr []int64, data []byte) {
	p := uintptr(unsafe.Pointer(&data[sliceHeaderSize]))
	*(*uintptr)(unsafe.Pointer(&data[0])) = p
	*(*int)(unsafe.Pointer(&data[uintptrSize])) = len(arr)
	*(*int)(unsafe.Pointer(&data[uintptrSize+intSize])) = len(arr)

	var intArr []int64
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&intArr))
	sliceHeader.Len = len(arr)
	sliceHeader.Cap = len(arr)
	sliceHeader.Data = uintptr(unsafe.Pointer(&data[sliceHeaderSize]))
	copy(intArr, arr)
}

func toInt64Array(addr uintptr) []int64 {
	return *(*[]int64)(unsafe.Pointer(addr))
}

func sizeofInt64Array(arr []int64) uintptr {
	size := sliceHeaderSize + int64Size*uintptr(len(arr))
	return size
}

func copyFloat32Array(arr []float32, data []byte) {
	p := uintptr(unsafe.Pointer(&data[sliceHeaderSize]))
	*(*uintptr)(unsafe.Pointer(&data[0])) = p
	*(*int)(unsafe.Pointer(&data[uintptrSize])) = len(arr)
	*(*int)(unsafe.Pointer(&data[uintptrSize+intSize])) = len(arr)

	var float32Arr []float32
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&float32Arr))
	sliceHeader.Len = len(arr)
	sliceHeader.Cap = len(arr)
	sliceHeader.Data = uintptr(unsafe.Pointer(&data[sliceHeaderSize]))
	copy(float32Arr, arr)
}

func toFloat32Array(addr uintptr) []float32 {
	return *(*[]float32)(unsafe.Pointer(addr))
}

func sizeofFloat32Array(arr []float32) uintptr {
	size := sliceHeaderSize + float32Size*uintptr(len(arr))
	if size&7 != 0 {
		size = size - size&7 + 8
	}
	return size
}

func copyStringArray(arr []string, data []byte) {
	p := uintptr(unsafe.Pointer(&data[sliceHeaderSize]))
	*(*uintptr)(unsafe.Pointer(&data[0])) = p
	*(*int)(unsafe.Pointer(&data[uintptrSize])) = len(arr)
	*(*int)(unsafe.Pointer(&data[uintptrSize+intSize])) = len(arr)
	headerOffset := sliceHeaderSize
	dataOffset := headerOffset + stringHeaderSize*uintptr(len(arr))
	var size uintptr
	for i := 0; i < len(arr); i++ {
		size = uintptr(len(arr[i]))
		*(*uintptr)(unsafe.Pointer(&data[headerOffset])) = uintptr(unsafe.Pointer(&data[dataOffset]))
		*(*int)(unsafe.Pointer(&data[headerOffset+uintptrSize])) = len(arr[i])
		copy(data[dataOffset:dataOffset+size], *(*[]byte)(unsafe.Pointer(&arr[i])))
		dataOffset += size
		headerOffset += stringHeaderSize
	}
}

func toStringArray(addr uintptr) []string {
	return *(*[]string)(unsafe.Pointer(addr))
}

func sizeofStringArray(arr []string) uintptr {
	size := sliceHeaderSize + stringHeaderSize*uintptr(len(arr))
	for i := 0; i < len(arr); i++ {
		size += uintptr(len(arr[i]))
	}
	if size&7 != 0 {
		size = size - size&7 + 8
	}
	return size
}

func deepcpyOfStringArray(arr []string) []string {
	ret := make([]string, len(arr))
	for i := 0; i < len(arr); i++ {
		ret[i] = deepcopyOfString(arr[i])
	}
	return ret
}

func deepcopyOfString(s string) string {
	b := []byte(s)
	copyB := make([]byte, len(b))
	copy(copyB, b)
	return string(copyB)
}
