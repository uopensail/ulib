package sample

import (
	"encoding/json"
	"fmt"
	"iter"
	"maps"
	"slices"
	"strings"
	"unsafe"

	"github.com/bytedance/sonic"
)

// ImmutableFeature holds a single typed value stored inside an Arena page.
// The ptr field points to a contiguous block whose first 8 bytes encode the
// DataType, followed immediately by the value data.
//
// Memory layout per type (all fields 8-byte aligned):
//
//	Int64:    [DataType:8][Value:8]                         = 16 B
//	Float32:  [DataType:8][Value:4][Pad:4]                 = 16 B
//	String:   [DataType:8][Len:8][Data:align(len)]         = 16+align(len) B
//	Int64s:   [DataType:8][Len:8][Data:len*8]              = 16+len*8 B
//	Float32s: [DataType:8][Len:8][Data:align(len*4)]       = 16+align(len*4) B
//	Strings:  [DataType:8][Len:8][Headers:len*16][Data:…]
//
// All slice and string getters return a view directly into the arena — no
// copying occurs.
type ImmutableFeature struct {
	ptr unsafe.Pointer
}

// Type returns the DataType of the stored value, or InvalidType when ptr is nil.
func (f *ImmutableFeature) Type() DataType {
	if f.ptr == nil {
		return InvalidType
	}
	return *(*DataType)(f.ptr)
}

// Get returns the stored value using its native Go type.
func (f *ImmutableFeature) Get() any {
	switch f.Type() {
	case Int64Type:
		return f.GetInt64Unsafe()
	case Float32Type:
		return f.GetFloat32Unsafe()
	case StringType:
		return f.GetStringUnsafe()
	case Int64sType:
		return f.GetInt64sUnsafe()
	case Float32sType:
		return f.GetFloat32sUnsafe()
	case StringsType:
		return f.GetStringsUnsafe()
	}
	return nil
}

func (f *ImmutableFeature) GetInt64() (int64, error)    { return getInt64(f.ptr) }
func (f *ImmutableFeature) GetInt64Unsafe() int64       { return getInt64Unsafe(f.ptr) }
func (f *ImmutableFeature) GetFloat32() (float32, error) { return getFloat32(f.ptr) }
func (f *ImmutableFeature) GetFloat32Unsafe() float32   { return getFloat32Unsafe(f.ptr) }
func (f *ImmutableFeature) GetString() (string, error)  { return getString(f.ptr) }
func (f *ImmutableFeature) GetStringUnsafe() string     { return getStringUnsafe(f.ptr) }
func (f *ImmutableFeature) GetInt64s() ([]int64, error)  { return getInt64s(f.ptr) }
func (f *ImmutableFeature) GetInt64sUnsafe() []int64    { return getInt64sUnsafe(f.ptr) }
func (f *ImmutableFeature) GetFloat32s() ([]float32, error) { return getFloat32s(f.ptr) }
func (f *ImmutableFeature) GetFloat32sUnsafe() []float32 { return getFloat32sUnsafe(f.ptr) }
func (f *ImmutableFeature) GetStrings() ([]string, error) { return getStrings(f.ptr) }
func (f *ImmutableFeature) GetStringsUnsafe() []string  { return getStringsUnsafe(f.ptr) }

// ImmutableFeatures is a read-only, arena-backed collection of typed features.
// After construction the map cannot be modified, making it safe for concurrent
// reads without additional locking.
type ImmutableFeatures struct {
	arena    *Arena
	features map[string]ImmutableFeature
}

// NewImmutableFeatures returns an empty ImmutableFeatures backed by arena.
func NewImmutableFeatures(arena *Arena) *ImmutableFeatures {
	return &ImmutableFeatures{
		arena:    arena,
		features: make(map[string]ImmutableFeature),
	}
}

// NewImmutableFeaturesFromMap constructs an ImmutableFeatures from a plain
// map, copying each value into arena. Supported value types: int, int64,
// float32, float64, string, []int, []int64, []float32, []float64, []string.
func NewImmutableFeaturesFromMap(data map[string]any, arena *Arena) (*ImmutableFeatures, error) {
	f := &ImmutableFeatures{
		arena:    arena,
		features: make(map[string]ImmutableFeature, len(data)),
	}
	for key, value := range data {
		ptr, err := putAnyValue(value, arena)
		if err != nil {
			return nil, fmt.Errorf("key %s: %w", key, err)
		}
		f.features[key] = ImmutableFeature{ptr: ptr}
	}
	return f, nil
}

// GetType returns the DataType for key, or InvalidType when absent.
func (f *ImmutableFeatures) GetType(key string) DataType {
	if fea, ok := f.features[key]; ok {
		return fea.Type()
	}
	return InvalidType
}

// Get returns the Feature for key, or nil when absent.
func (f *ImmutableFeatures) Get(key string) Feature {
	if fea, ok := f.features[key]; ok {
		return &fea
	}
	return nil
}

// Has reports whether key exists.
func (f *ImmutableFeatures) Has(key string) bool {
	_, ok := f.features[key]
	return ok
}

// Len returns the number of features.
func (f *ImmutableFeatures) Len() int { return len(f.features) }

// Keys returns all feature keys in unspecified order.
func (f *ImmutableFeatures) Keys() []string {
	return slices.Collect(maps.Keys(f.features))
}

// ForEach calls fn for every key-feature pair, stopping on the first non-nil error.
func (f *ImmutableFeatures) ForEach(fn IteratorFunc) error {
	for key, fea := range f.features {
		if err := fn(key, &fea); err != nil {
			return err
		}
	}
	return nil
}

// All returns an iterator over all key-feature pairs (Go 1.23+ range syntax).
//
//	for key, f := range features.All() {
//	    fmt.Println(key, f.Type())
//	}
func (f *ImmutableFeatures) All() iter.Seq2[string, Feature] {
	return func(yield func(string, Feature) bool) {
		for key, fea := range f.features {
			fea := fea // copy loop variable for safe address-taking
			if !yield(key, &fea) {
				return
			}
		}
	}
}

// MapAny converts every feature to a {type, value} struct for serialization.
func (f *ImmutableFeatures) MapAny() (map[string]any, error) {
	out := make(map[string]any, len(f.features))
	for key, fea := range f.features {
		v, err := featureToAny(&fea)
		if err != nil {
			return nil, fmt.Errorf("key %s: %w", key, err)
		}
		out[key] = v
	}
	return out, nil
}

// MarshalJSON encodes the collection as JSON.
func (f *ImmutableFeatures) MarshalJSON() ([]byte, error) {
	m, err := f.MapAny()
	if err != nil {
		return nil, err
	}
	return sonic.Marshal(m)
}

// UnmarshalJSON decodes JSON produced by MarshalJSON.
func (f *ImmutableFeatures) UnmarshalJSON(data []byte) error {
	type wire struct {
		Type  DataType        `json:"type"`
		Value json.RawMessage `json:"value"`
	}
	var raw map[string]wire
	if err := sonic.Unmarshal(data, &raw); err != nil {
		return err
	}

	if f.arena == nil {
		f.arena = NewArena()
	}
	if f.features == nil {
		f.features = make(map[string]ImmutableFeature, len(raw))
	}

	for key, w := range raw {
		ptr, err := unmarshalValue(w.Type, w.Value, f.arena)
		if err != nil {
			return fmt.Errorf("key %s: %w", key, err)
		}
		f.features[key] = ImmutableFeature{ptr: ptr}
	}
	return nil
}

// unmarshalValue decodes a single JSON value of the given DataType into arena.
func unmarshalValue(dt DataType, raw json.RawMessage, arena *Arena) (unsafe.Pointer, error) {
	switch dt {
	case Int64Type:
		var v int64
		if err := sonic.Unmarshal(raw, &v); err != nil {
			return nil, err
		}
		return putInt64(v, arena)
	case Float32Type:
		var v float32
		if err := sonic.Unmarshal(raw, &v); err != nil {
			return nil, err
		}
		return putFloat32(v, arena)
	case StringType:
		var v string
		if err := sonic.Unmarshal(raw, &v); err != nil {
			return nil, err
		}
		return putString(v, arena)
	case Int64sType:
		var v []int64
		if err := sonic.Unmarshal(raw, &v); err != nil {
			return nil, err
		}
		return putInt64s(v, arena)
	case Float32sType:
		var v []float32
		if err := sonic.Unmarshal(raw, &v); err != nil {
			return nil, err
		}
		return putFloat32s(v, arena)
	case StringsType:
		var v []string
		if err := sonic.Unmarshal(raw, &v); err != nil {
			return nil, err
		}
		return putStrings(v, arena)
	default:
		return nil, fmt.Errorf("unknown data type %v", dt)
	}
}

// Mutable returns a deep-copied MutableFeatures, independent of the arena.
func (f *ImmutableFeatures) Mutable() *MutableFeatures {
	ret := NewMutableFeatures()
	for key, fea := range f.features {
		switch fea.Type() {
		case Int64Type:
			ret.features[key] = &Int64{Value: fea.GetInt64Unsafe()}
		case Float32Type:
			ret.features[key] = &Float32{Value: fea.GetFloat32Unsafe()}
		case StringType:
			ret.features[key] = &String{Value: strings.Clone(fea.GetStringUnsafe())}
		case Int64sType:
			src := fea.GetInt64sUnsafe()
			dst := make([]int64, len(src))
			copy(dst, src)
			ret.features[key] = &Int64s{Value: dst}
		case Float32sType:
			src := fea.GetFloat32sUnsafe()
			dst := make([]float32, len(src))
			copy(dst, src)
			ret.features[key] = &Float32s{Value: dst}
		case StringsType:
			ret.features[key] = &Strings{Value: deepcopyStrings(fea.GetStringsUnsafe())}
		}
	}
	return ret
}

// ---------------------------------------------------------------------------
// Arena storage helpers
// ---------------------------------------------------------------------------

// putInt64 writes [DataType:8][Value:8] into a 16-byte arena block.
func putInt64(value int64, arena *Arena) (unsafe.Pointer, error) {
	data, err := arena.allocate(16)
	if err != nil {
		return nil, err
	}
	*(*DataType)(unsafe.Pointer(&data[0])) = Int64Type
	*(*int64)(unsafe.Pointer(&data[8])) = value
	return unsafe.Pointer(&data[0]), nil
}

func getInt64(ptr unsafe.Pointer) (int64, error) {
	if getDataType(ptr) != Int64Type {
		return 0, fmt.Errorf("%w: expected %v, got %v", ErrTypeMismatch, Int64Type, getDataType(ptr))
	}
	return getInt64Unsafe(ptr), nil
}

func getInt64Unsafe(ptr unsafe.Pointer) int64 {
	return *(*int64)(unsafe.Add(ptr, 8))
}

// putFloat32 writes [DataType:8][Value:4][Pad:4] into a 16-byte arena block.
func putFloat32(value float32, arena *Arena) (unsafe.Pointer, error) {
	data, err := arena.allocate(16)
	if err != nil {
		return nil, err
	}
	*(*DataType)(unsafe.Pointer(&data[0])) = Float32Type
	*(*float32)(unsafe.Pointer(&data[8])) = value
	return unsafe.Pointer(&data[0]), nil
}

func getFloat32(ptr unsafe.Pointer) (float32, error) {
	if getDataType(ptr) != Float32Type {
		return 0, fmt.Errorf("%w: expected %v, got %v", ErrTypeMismatch, Float32Type, getDataType(ptr))
	}
	return getFloat32Unsafe(ptr), nil
}

func getFloat32Unsafe(ptr unsafe.Pointer) float32 {
	return *(*float32)(unsafe.Add(ptr, 8))
}

// putString writes [DataType:8][Len:8][Data:align(len)] into the arena.
func putString(value string, arena *Arena) (unsafe.Pointer, error) {
	strLen := len(value)
	total := 16 + alignSize(uintptr(strLen)) // 8 (type) + 8 (len) + data
	data, err := arena.allocate(total)
	if err != nil {
		return nil, err
	}
	*(*DataType)(unsafe.Pointer(&data[0])) = StringType
	*(*uint64)(unsafe.Pointer(&data[8])) = uint64(strLen)
	if strLen > 0 {
		copy(data[16:], value)
	}
	return unsafe.Pointer(&data[0]), nil
}

func getString(ptr unsafe.Pointer) (string, error) {
	if getDataType(ptr) != StringType {
		return "", fmt.Errorf("%w: expected %v, got %v", ErrTypeMismatch, StringType, getDataType(ptr))
	}
	return getStringUnsafe(ptr), nil
}

func getStringUnsafe(ptr unsafe.Pointer) string {
	n := *(*uint64)(unsafe.Add(ptr, 8))
	if n == 0 {
		return ""
	}
	return unsafe.String((*byte)(unsafe.Add(ptr, 16)), int(n))
}

// putInt64s writes [DataType:8][Len:8][Data:len*8] into the arena.
// The slice returned by getInt64sUnsafe points directly into the arena block.
func putInt64s(arr []int64, arena *Arena) (unsafe.Pointer, error) {
	n := len(arr)
	total := uintptr(16 + n*8)
	data, err := arena.allocate(total)
	if err != nil {
		return nil, err
	}
	*(*DataType)(unsafe.Pointer(&data[0])) = Int64sType
	*(*uint64)(unsafe.Pointer(&data[8])) = uint64(n)
	if n > 0 {
		copy(unsafe.Slice((*int64)(unsafe.Pointer(&data[16])), n), arr)
	}
	return unsafe.Pointer(&data[0]), nil
}

func getInt64s(ptr unsafe.Pointer) ([]int64, error) {
	if getDataType(ptr) != Int64sType {
		return nil, fmt.Errorf("%w: expected %v, got %v", ErrTypeMismatch, Int64sType, getDataType(ptr))
	}
	return getInt64sUnsafe(ptr), nil
}

func getInt64sUnsafe(ptr unsafe.Pointer) []int64 {
	n := *(*uint64)(unsafe.Add(ptr, 8))
	if n == 0 {
		return nil
	}
	return unsafe.Slice((*int64)(unsafe.Add(ptr, 16)), int(n))
}

// putFloat32s writes [DataType:8][Len:8][Data:align(len*4)] into the arena.
func putFloat32s(arr []float32, arena *Arena) (unsafe.Pointer, error) {
	n := len(arr)
	total := 16 + alignSize(uintptr(n)*4)
	data, err := arena.allocate(total)
	if err != nil {
		return nil, err
	}
	*(*DataType)(unsafe.Pointer(&data[0])) = Float32sType
	*(*uint64)(unsafe.Pointer(&data[8])) = uint64(n)
	if n > 0 {
		copy(unsafe.Slice((*float32)(unsafe.Pointer(&data[16])), n), arr)
	}
	return unsafe.Pointer(&data[0]), nil
}

func getFloat32s(ptr unsafe.Pointer) ([]float32, error) {
	if getDataType(ptr) != Float32sType {
		return nil, fmt.Errorf("%w: expected %v, got %v", ErrTypeMismatch, Float32sType, getDataType(ptr))
	}
	return getFloat32sUnsafe(ptr), nil
}

func getFloat32sUnsafe(ptr unsafe.Pointer) []float32 {
	n := *(*uint64)(unsafe.Add(ptr, 8))
	if n == 0 {
		return nil
	}
	return unsafe.Slice((*float32)(unsafe.Add(ptr, 16)), int(n))
}

// putStrings writes a compact block into the arena:
//
//	[DataType:8][Len:8][Headers:align(len*16)][StringData:align(totalChars)]
//
// Each header is a Go string (pointer+length) pointing into the StringData
// region, so getStringsUnsafe can return a zero-copy []string view.
func putStrings(arr []string, arena *Arena) (unsafe.Pointer, error) {
	n := len(arr)
	headersSize := alignSize(uintptr(n) * 16)
	var totalLen uintptr
	for _, s := range arr {
		totalLen += uintptr(len(s))
	}
	total := 16 + headersSize + alignSize(totalLen)
	data, err := arena.allocate(total)
	if err != nil {
		return nil, err
	}

	*(*DataType)(unsafe.Pointer(&data[0])) = StringsType
	*(*uint64)(unsafe.Pointer(&data[8])) = uint64(n)

	headerBase := 16
	dataBase := headerBase + int(headersSize)
	var offset uintptr
	for i, s := range arr {
		hdrPtr := unsafe.Pointer(&data[headerBase+i*16])
		if l := len(s); l > 0 {
			copy(data[dataBase+int(offset):], s)
			str := unsafe.String((*byte)(unsafe.Pointer(&data[dataBase+int(offset)])), l)
			*(*string)(hdrPtr) = str
			offset += uintptr(l)
		} else {
			*(*string)(hdrPtr) = ""
		}
	}
	return unsafe.Pointer(&data[0]), nil
}

func getStrings(ptr unsafe.Pointer) ([]string, error) {
	if getDataType(ptr) != StringsType {
		return nil, fmt.Errorf("%w: expected %v, got %v", ErrTypeMismatch, StringsType, getDataType(ptr))
	}
	return getStringsUnsafe(ptr), nil
}

// getStringsUnsafe returns a zero-copy []string whose headers live inside the
// arena block.
func getStringsUnsafe(ptr unsafe.Pointer) []string {
	n := *(*uint64)(unsafe.Add(ptr, 8))
	if n == 0 {
		return nil
	}
	return unsafe.Slice((*string)(unsafe.Add(ptr, 16)), int(n))
}

// getDataType reads the first 8 bytes of a memory block as a DataType.
func getDataType(ptr unsafe.Pointer) DataType {
	if ptr == nil {
		return InvalidType
	}
	return *(*DataType)(ptr)
}

// putAnyValue stores value into arena, dispatching on its Go type.
// Supported: int, int64, float32, float64, string, []int, []int64, []float32, []float64, []string.
func putAnyValue(value any, arena *Arena) (unsafe.Pointer, error) {
	switch v := value.(type) {
	case int64:
		return putInt64(v, arena)
	case int:
		return putInt64(int64(v), arena)
	case float32:
		return putFloat32(v, arena)
	case float64:
		return putFloat32(float32(v), arena)
	case string:
		return putString(v, arena)
	case []int64:
		return putInt64s(v, arena)
	case []int:
		s := make([]int64, len(v))
		for i, x := range v {
			s[i] = int64(x)
		}
		return putInt64s(s, arena)
	case []float32:
		return putFloat32s(v, arena)
	case []float64:
		s := make([]float32, len(v))
		for i, x := range v {
			s[i] = float32(x)
		}
		return putFloat32s(s, arena)
	case []string:
		return putStrings(v, arena)
	default:
		return nil, fmt.Errorf("unsupported type: %T", value)
	}
}

// deepcopyStrings returns a fresh []string whose elements are independent
// copies of the originals (using strings.Clone for efficiency).
func deepcopyStrings(arr []string) []string {
	if len(arr) == 0 {
		return nil
	}
	out := make([]string, len(arr))
	for i, s := range arr {
		out[i] = strings.Clone(s)
	}
	return out
}
