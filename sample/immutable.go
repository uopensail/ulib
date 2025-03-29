package sample

import (
	"encoding/json"
	"fmt"
	"strings"
	"unsafe"

	"github.com/bytedance/sonic"
)

// stringHeader, sliceHeader are deprecated, so we need to defined it here
type stringHeader struct {
	Data uintptr
	Len  int
}

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

const int64Size uintptr = unsafe.Sizeof(int64(0))
const float32Size uintptr = unsafe.Sizeof(float32(0))
const stringHeaderSize uintptr = unsafe.Sizeof(stringHeader{Data: 0, Len: 0})
const sliceHeaderSize uintptr = unsafe.Sizeof(sliceHeader{Data: 0, Len: 0, Cap: 0})

type ImmutableFeature struct {
	addr uintptr
}

func (f *ImmutableFeature) Type() DataType {
	return *(*DataType)(uintptr2Pointer(f.addr))
}

func (f *ImmutableFeature) GetInt64() (int64, error) {
	return getInt64(f.addr)
}

func (f *ImmutableFeature) GetFloat32() (float32, error) {
	return getFloat32(f.addr)
}

func (f *ImmutableFeature) GetString() (string, error) {
	return getString(f.addr)
}

func (f *ImmutableFeature) GetInt64s() ([]int64, error) {
	return getInt64s(f.addr)
}

func (f *ImmutableFeature) GetFloat32s() ([]float32, error) {
	return getFloat32s(f.addr)
}

func (f *ImmutableFeature) GetStrings() ([]string, error) {
	return getStrings(f.addr)
}

type ImmutableFeatures struct {
	arena    *Arena
	features map[string]ImmutableFeature
}

func NewImmutableFeatures(arena *Arena) *ImmutableFeatures {
	return &ImmutableFeatures{
		arena:    arena,
		features: make(map[string]ImmutableFeature),
	}
}

func (f *ImmutableFeatures) Get(key string) Feature {
	if fea, ok := f.features[key]; ok {
		return &fea
	}
	return nil
}
func (f *ImmutableFeatures) Len() int {
	return len(f.features)
}
func (f *ImmutableFeatures) Keys() []string {
	ret := make([]string, 0, len(f.features))
	for key := range f.features {
		ret = append(ret, key)
	}
	return ret
}
func (f *ImmutableFeatures) MapAny() map[string]any {
	feas := make(map[string]interface{})

	for key, fea := range f.features {
		dtype := fea.Type()
		switch dtype {
		case Int64Type:
			v, _ := fea.GetInt64()
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value int64    `json:"value"`
			}{
				Int64Type,
				v,
			}
		case Float32Type:
			v, _ := fea.GetFloat32()
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value float32  `json:"value"`
			}{
				Float32Type,
				v,
			}
		case StringType:
			v, _ := fea.GetString()
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value string   `json:"value"`
			}{
				StringType,
				v,
			}
		case Int64sType:
			v, _ := fea.GetInt64s()
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value []int64  `json:"value"`
			}{
				Int64sType,
				v,
			}
		case Float32sType:
			v, _ := fea.GetFloat32s()
			feas[key] = struct {
				Type  DataType  `json:"type"`
				Value []float32 `json:"value"`
			}{
				Float32sType,
				v,
			}
		case StringsType:
			v, _ := fea.GetStrings()
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value []string `json:"value"`
			}{
				StringsType,
				v,
			}
		}
	}
	return feas
}

func (f *ImmutableFeatures) MarshalJSON() ([]byte, error) {
	feas := f.MapAny()
	return sonic.Marshal(feas)
}

func (f *ImmutableFeatures) Mutable() *MutableFeatures {
	ret := NewMutableFeatures()
	for key, fea := range f.features {
		dtype := fea.Type()
		switch dtype {
		case Int64Type:
			v, _ := fea.GetInt64()
			ret.features[key] = &Int64{Value: v}
		case Float32Type:
			v, _ := fea.GetFloat32()
			ret.features[key] = &Float32{Value: v}
		case StringType:
			v, _ := fea.GetString()
			ret.features[key] = &String{Value: deepcopyOfString(v)}
		case Int64sType:
			v, _ := fea.GetInt64s()
			nums := make([]int64, len(v))
			copy(nums, v)
			ret.features[key] = &Int64s{Value: nums}
		case Float32sType:
			v, _ := fea.GetFloat32s()
			nums := make([]float32, len(v))
			copy(nums, v)
			ret.features[key] = &Float32s{Value: nums}
		case StringsType:
			v, _ := fea.GetStrings()
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
			f.features[key] = ImmutableFeature{addr: putInt64(num, f.arena)}
		case Float32Type:
			var num float32
			sonic.Unmarshal(value.Value, &num)
			f.features[key] = ImmutableFeature{addr: putFloat32(num, f.arena)}
		case StringType:
			var str string
			sonic.Unmarshal(value.Value, &str)
			f.features[key] = ImmutableFeature{addr: putString(str, f.arena)}
		case Int64sType:
			var nums []int64
			sonic.Unmarshal(value.Value, &nums)
			f.features[key] = ImmutableFeature{addr: putInt64s(nums, f.arena)}
		case Float32sType:
			var nums []float32
			sonic.Unmarshal(value.Value, &nums)
			f.features[key] = ImmutableFeature{addr: putFloat32s(nums, f.arena)}
		case StringsType:
			var strs []string
			sonic.Unmarshal(value.Value, &strs)
			f.features[key] = ImmutableFeature{addr: putStrings(strs, f.arena)}
		}
	}
	return nil
}

func setDataType(data []byte, dataType DataType) {
	*(*DataType)(unsafe.Pointer(&data[0])) = dataType
}

func putInt64(value int64, arena *Arena) uintptr {
	data := arena.allocate(8 + int64Size)
	setDataType(data, Int64Type)
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
	setDataType(data, Float32Type)
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
	setDataType(data, StringType)
	header := (*stringHeader)(unsafe.Pointer(&data[8]))
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
	setDataType(data, Int64sType)
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
	setDataType(data, Float32sType)
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
	setDataType(data, StringsType)
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
	slice := (*sliceHeader)(unsafe.Pointer(&data[0]))
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
	slice := (*sliceHeader)(unsafe.Pointer(&data[0]))
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
	slice := (*sliceHeader)(unsafe.Pointer(&data[0]))
	slice.Data = p
	slice.Cap = len(arr)
	slice.Len = len(arr)
	headerOffset := sliceHeaderSize
	dataOffset := headerOffset + stringHeaderSize*uintptr(len(arr))
	var str *stringHeader
	var size uintptr
	for i := range arr {
		size = uintptr(len(arr[i]))
		str = (*stringHeader)(unsafe.Pointer(&data[headerOffset]))
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
	for i := range arr {
		size += uintptr(len(arr[i]))
	}
	return ((size + 7) >> 3) << 3
}

func deepcpyOfStrings(arr []string) []string {
	ret := make([]string, len(arr))
	for i := range arr {
		ret[i] = deepcopyOfString(arr[i])
	}
	return ret
}

func deepcopyOfString(s string) (str string) {
	var builder strings.Builder
	builder.WriteString(s)
	return builder.String()
}

func uintptr2Pointer(addr uintptr) unsafe.Pointer {
	return unsafe.Pointer(addr)
}
