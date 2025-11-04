package sample

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"unsafe"

	"github.com/bytedance/sonic"
)

var (
	ErrTypeMismatch = errors.New("type mismatch")
	ErrInvalidData  = errors.New("invalid data")
)

/**
 * @brief DataType represents the type of data stored in an ImmutableFeature
 *
 * Uses uint64 to ensure 8-byte alignment for optimal memory access.
 * Each stored value begins with a DataType field to enable type checking.
 */
type DataType uint64

const (
	InvalidType  DataType = iota // Invalid or uninitialized type
	Int64Type                    // 64-bit signed integer
	Float32Type                  // 32-bit floating point
	StringType                   // UTF-8 string
	Int64sType                   // Slice of 64-bit signed integers
	Float32sType                 // Slice of 32-bit floating points
	StringsType                  // Slice of UTF-8 strings
)

/**
 * @brief ImmutableFeature represents a read-only typed value stored in arena memory
 *
 * Memory Layout (all 8-byte aligned):
 * - Int64:     [DataType:8] + [Value:8] = 16 bytes
 * - Float32:   [DataType:8] + [Value:4] + [Padding:4] = 16 bytes
 * - String:    [DataType:8] + [Len:8] + [Data:aligned] = 16 + aligned(len)
 * - Int64s:    [DataType:8] + [Len:8] + [Data:len*8] = 16 + len*8 bytes
 * - Float32s:  [DataType:8] + [Len:8] + [Data:aligned(len*4)] = 16 + aligned(len*4) bytes
 * - Strings:   [DataType:8] + [Len:8] + [StringHeaders:len*16] + [StringData:aligned]
 *              = 16 + len*16 + aligned(total_string_length) bytes
 *
 * Key characteristics:
 * - Zero-copy access: All get operations return direct pointers to arena memory
 * - Type safety: DataType field enables runtime type checking
 * - Memory efficiency: Compact layout with proper alignment
 * - Immutable: Data cannot be modified after storage
 *
 * The addr field points to the beginning of the allocated memory block.
 */
type ImmutableFeature struct {
	addr uintptr // Pointer to the start of allocated memory
}

/**
 * @brief Returns the data type of the stored value
 *
 * @return DataType of the stored value, InvalidType if addr is null
 */
func (f *ImmutableFeature) Type() DataType {
	if f.addr == 0 {
		return InvalidType
	}
	return *(*DataType)(unsafe.Pointer(f.addr))
}

/**
 * @brief Retrieves the stored value as any type
 *
 * @return The underlying value as any type
 */
func (f *ImmutableFeature) Get() any {
	switch f.Type() {
	case Int64Type:
		value, _ := f.GetInt64()
		return value
	case Float32Type:
		value, _ := f.GetFloat32()
		return value
	case StringType:
		value, _ := f.GetString()
		return value
	case Int64sType:
		value, _ := f.GetInt64s()
		return value
	case Float32sType:
		value, _ := f.GetFloat32s()
		return value
	case StringsType:
		value, _ := f.GetStrings()
		return value
	}
	return nil
}

/**
 * @brief Retrieves the stored value as int64
 *
 * @return The int64 value and error if type mismatch
 */
func (f *ImmutableFeature) GetInt64() (int64, error) {
	return getInt64(f.addr)
}

/**
 * @brief Retrieves the stored value as float32
 *
 * @return The float32 value and error if type mismatch
 */
func (f *ImmutableFeature) GetFloat32() (float32, error) {
	return getFloat32(f.addr)
}

/**
 * @brief Retrieves the stored value as string
 *
 * @return The string value and error if type mismatch
 */
func (f *ImmutableFeature) GetString() (string, error) {
	return getString(f.addr)
}

/**
 * @brief Retrieves the stored value as []int64
 *
 * @return The []int64 slice and error if type mismatch
 */
func (f *ImmutableFeature) GetInt64s() ([]int64, error) {
	return getInt64s(f.addr)
}

/**
 * @brief Retrieves the stored value as []float32
 *
 * @return The []float32 slice and error if type mismatch
 */
func (f *ImmutableFeature) GetFloat32s() ([]float32, error) {
	return getFloat32s(f.addr)
}

/**
 * @brief Retrieves the stored value as []string
 *
 * @return The []string slice and error if type mismatch
 */
func (f *ImmutableFeature) GetStrings() ([]string, error) {
	return getStrings(f.addr)
}

/**
 * @brief ImmutableFeatures is a read-only collection of typed features
 *
 * This collection stores features in arena-allocated memory for efficient
 * access and reduced garbage collection pressure. Once created, the features
 * cannot be modified, ensuring data integrity in concurrent environments.
 *
 * Key characteristics:
 * - Immutable after creation
 * - Memory-efficient storage using arena allocation
 * - Type-safe access to stored values
 * - JSON serialization support
 * - Conversion to mutable features
 */
type ImmutableFeatures struct {
	arena    *Arena                      // Memory arena for storage
	features map[string]ImmutableFeature // Map of feature name to feature
}

/**
 * @brief Creates a new empty ImmutableFeatures collection
 *
 * @param arena Memory arena for storage (creates new if nil)
 * @return Pointer to new ImmutableFeatures collection
 */
func NewImmutableFeatures(arena *Arena) *ImmutableFeatures {
	return &ImmutableFeatures{
		arena:    arena,
		features: make(map[string]ImmutableFeature),
	}
}

/**
 * @brief Creates ImmutableFeatures from a map of values
 *
 * @param data Map of feature names to values
 * @param arena Memory arena for storage (creates new if nil)
 * @return Pointer to new ImmutableFeatures and error if conversion fails
 *
 * Supported value types:
 * - int, int64: stored as Int64Type
 * - float32, float64: stored as Float32Type (float64 converted to float32)
 * - string: stored as StringType
 * - []int, []int64: stored as Int64sType
 * - []float32, []float64: stored as Float32sType
 * - []string: stored as StringsType
 */
func NewImmutableFeaturesFromMap(data map[string]any, arena *Arena) (*ImmutableFeatures, error) {
	if arena == nil {
		arena = NewArena()
	}

	features := &ImmutableFeatures{
		arena:    arena,
		features: make(map[string]ImmutableFeature, len(data)),
	}

	for key, value := range data {
		addr, err := putAnyValue(value, arena)
		if err != nil {
			return nil, fmt.Errorf("failed to store value for key %s: %w", key, err)
		}
		features.features[key] = ImmutableFeature{addr: addr}
	}

	return features, nil
}

/**
 * @brief Retrieves a feature by name
 *
 * @param key Feature name
 * @return Feature interface or nil if not found
 */
func (f *ImmutableFeatures) Get(key string) Feature {
	if fea, ok := f.features[key]; ok {
		return &fea
	}
	return nil
}

/**
 * @brief Checks if a feature exists
 *
 * @param key Feature name
 * @return True if feature exists, false otherwise
 */
func (f *ImmutableFeatures) Has(key string) bool {
	_, ok := f.features[key]
	return ok
}

/**
 * @brief Returns the number of features in the collection
 *
 * @return Number of features
 */
func (f *ImmutableFeatures) Len() int {
	return len(f.features)
}

/**
 * @brief Returns all feature names
 *
 * @return Slice of feature names
 */
func (f *ImmutableFeatures) Keys() []string {
	ret := make([]string, 0, len(f.features))
	for key := range f.features {
		ret = append(ret, key)
	}
	return ret
}

/**
 * @brief Iterates over all features with a callback function
 *
 * @param fn Callback function called for each feature
 * @return Error if callback returns error
 */
func (f *ImmutableFeatures) ForEach(fn func(key string, feature Feature) error) error {
	for key, fea := range f.features {
		if err := fn(key, &fea); err != nil {
			return err
		}
	}
	return nil
}

/**
 * @brief Converts features to a map for serialization
 *
 * @return Map of feature names to structured values and error if conversion fails
 *
 * Each value in the returned map contains:
 * - Type: DataType of the feature
 * - Value: The actual value
 */
func (f *ImmutableFeatures) MapAny() (map[string]any, error) {
	feas := make(map[string]interface{}, len(f.features))

	for key, fea := range f.features {
		dtype := fea.Type()
		switch dtype {
		case Int64Type:
			v, err := fea.GetInt64()
			if err != nil {
				return nil, fmt.Errorf("failed to get int64 for key %s: %w", key, err)
			}
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value int64    `json:"value"`
			}{Int64Type, v}

		case Float32Type:
			v, err := fea.GetFloat32()
			if err != nil {
				return nil, fmt.Errorf("failed to get float32 for key %s: %w", key, err)
			}
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value float32  `json:"value"`
			}{Float32Type, v}

		case StringType:
			v, err := fea.GetString()
			if err != nil {
				return nil, fmt.Errorf("failed to get string for key %s: %w", key, err)
			}
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value string   `json:"value"`
			}{StringType, v}

		case Int64sType:
			v, err := fea.GetInt64s()
			if err != nil {
				return nil, fmt.Errorf("failed to get int64s for key %s: %w", key, err)
			}
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value []int64  `json:"value"`
			}{Int64sType, v}

		case Float32sType:
			v, err := fea.GetFloat32s()
			if err != nil {
				return nil, fmt.Errorf("failed to get float32s for key %s: %w", key, err)
			}
			feas[key] = struct {
				Type  DataType  `json:"type"`
				Value []float32 `json:"value"`
			}{Float32sType, v}

		case StringsType:
			v, err := fea.GetStrings()
			if err != nil {
				return nil, fmt.Errorf("failed to get strings for key %s: %w", key, err)
			}
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value []string `json:"value"`
			}{StringsType, v}

		default:
			return nil, fmt.Errorf("unknown data type %v for key %s", dtype, key)
		}
	}
	return feas, nil
}

/**
 * @brief Marshals features to JSON
 *
 * @return JSON bytes and error if marshaling fails
 */
func (f *ImmutableFeatures) MarshalJSON() ([]byte, error) {
	feas, err := f.MapAny()
	if err != nil {
		return nil, err
	}
	return sonic.Marshal(feas)
}

/**
 * @brief Unmarshals features from JSON
 *
 * @param data JSON bytes to unmarshal
 * @return Error if unmarshaling fails
 */
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
		var addr uintptr
		var unmarshalErr error

		switch value.Type {
		case Int64Type:
			var num int64
			if err := sonic.Unmarshal(value.Value, &num); err != nil {
				return fmt.Errorf("failed to unmarshal int64 for key %s: %w", key, err)
			}
			addr, unmarshalErr = putInt64(num, f.arena)

		case Float32Type:
			var num float32
			if err := sonic.Unmarshal(value.Value, &num); err != nil {
				return fmt.Errorf("failed to unmarshal float32 for key %s: %w", key, err)
			}
			addr, unmarshalErr = putFloat32(num, f.arena)

		case StringType:
			var str string
			if err := sonic.Unmarshal(value.Value, &str); err != nil {
				return fmt.Errorf("failed to unmarshal string for key %s: %w", key, err)
			}
			addr, unmarshalErr = putString(str, f.arena)

		case Int64sType:
			var nums []int64
			if err := sonic.Unmarshal(value.Value, &nums); err != nil {
				return fmt.Errorf("failed to unmarshal int64s for key %s: %w", key, err)
			}
			addr, unmarshalErr = putInt64s(nums, f.arena)

		case Float32sType:
			var nums []float32
			if err := sonic.Unmarshal(value.Value, &nums); err != nil {
				return fmt.Errorf("failed to unmarshal float32s for key %s: %w", key, err)
			}
			addr, unmarshalErr = putFloat32s(nums, f.arena)

		case StringsType:
			var strs []string
			if err := sonic.Unmarshal(value.Value, &strs); err != nil {
				return fmt.Errorf("failed to unmarshal strings for key %s: %w", key, err)
			}
			addr, unmarshalErr = putStrings(strs, f.arena)

		default:
			return fmt.Errorf("unknown data type %v for key %s", value.Type, key)
		}

		if unmarshalErr != nil {
			return fmt.Errorf("failed to store value for key %s: %w", key, unmarshalErr)
		}

		f.features[key] = ImmutableFeature{addr: addr}
	}
	return nil
}

/**
 * @brief Creates a mutable copy of the features
 *
 * @return Pointer to new MutableFeatures with deep-copied values
 *
 * All values are deep-copied to ensure the mutable copy is independent
 * of the original immutable features.
 */
func (f *ImmutableFeatures) Mutable() *MutableFeatures {
	ret := NewMutableFeatures()
	for key, fea := range f.features {
		dtype := fea.Type()
		switch dtype {
		case Int64Type:
			if v, err := fea.GetInt64(); err == nil {
				ret.features[key] = &Int64{Value: v}
			}
		case Float32Type:
			if v, err := fea.GetFloat32(); err == nil {
				ret.features[key] = &Float32{Value: v}
			}
		case StringType:
			if v, err := fea.GetString(); err == nil {
				ret.features[key] = &String{Value: strings.Clone(v)}
			}
		case Int64sType:
			if v, err := fea.GetInt64s(); err == nil {
				nums := make([]int64, len(v))
				copy(nums, v)
				ret.features[key] = &Int64s{Value: nums}
			}
		case Float32sType:
			if v, err := fea.GetFloat32s(); err == nil {
				nums := make([]float32, len(v))
				copy(nums, v)
				ret.features[key] = &Float32s{Value: nums}
			}
		case StringsType:
			if v, err := fea.GetStrings(); err == nil {
				ret.features[key] = &Strings{Value: deepcopyStrings(v)}
			}
		}
	}
	return ret
}

// Storage functions with 8-byte aligned memory layout

/**
 * @brief Stores an int64 value in arena memory
 *
 * @param value The int64 value to store
 * @param arena Memory arena for allocation
 * @return Memory address of stored value and error if allocation fails
 *
 * Memory Layout: [DataType:8] + [Value:8] = 16 bytes total
 */
func putInt64(value int64, arena *Arena) (uintptr, error) {
	data, err := arena.allocate(16)
	if err != nil {
		return 0, err
	}

	*(*DataType)(unsafe.Pointer(&data[0])) = Int64Type
	*(*int64)(unsafe.Pointer(&data[8])) = value

	return uintptr(unsafe.Pointer(&data[0])), nil
}

/**
 * @brief Retrieves an int64 value from arena memory
 *
 * @param addr Memory address of the stored value
 * @return The int64 value and error if type mismatch
 */
func getInt64(addr uintptr) (int64, error) {
	if *(*DataType)(unsafe.Pointer(addr)) != Int64Type {
		return 0, fmt.Errorf("%w: expected %v, got %v", ErrTypeMismatch, Int64Type, *(*DataType)(unsafe.Pointer(addr)))
	}
	return *(*int64)(unsafe.Pointer(addr + 8)), nil
}

/**
 * @brief Stores a float32 value in arena memory
 *
 * @param value The float32 value to store
 * @param arena Memory arena for allocation
 * @return Memory address of stored value and error if allocation fails
 *
 * Memory Layout: [DataType:8] + [Value:4] + [Padding:4] = 16 bytes total
 */
func putFloat32(value float32, arena *Arena) (uintptr, error) {
	data, err := arena.allocate(16)
	if err != nil {
		return 0, err
	}

	*(*DataType)(unsafe.Pointer(&data[0])) = Float32Type
	*(*float32)(unsafe.Pointer(&data[8])) = value
	// Bytes 12-16 are padding for alignment

	return uintptr(unsafe.Pointer(&data[0])), nil
}

/**
 * @brief Retrieves a float32 value from arena memory
 *
 * @param addr Memory address of the stored value
 * @return The float32 value and error if type mismatch
 */
func getFloat32(addr uintptr) (float32, error) {
	if *(*DataType)(unsafe.Pointer(addr)) != Float32Type {
		return 0, fmt.Errorf("%w: expected %v, got %v", ErrTypeMismatch, Float32Type, *(*DataType)(unsafe.Pointer(addr)))
	}
	return *(*float32)(unsafe.Pointer(addr + 8)), nil
}

/**
 * @brief Stores a string value in arena memory
 *
 * @param value The string value to store
 * @param arena Memory arena for allocation
 * @return Memory address of stored value and error if allocation fails
 *
 * Memory Layout: [DataType:8] + [Len:8] + [Data:aligned] = 16 + aligned(len) bytes
 */
func putString(value string, arena *Arena) (uintptr, error) {
	strLen := len(value)
	dataSize := alignSize(uintptr(strLen))
	totalSize := 16 + dataSize // 16 = 8(type) + 8(len)

	data, err := arena.allocate(totalSize)
	if err != nil {
		return 0, err
	}

	*(*DataType)(unsafe.Pointer(&data[0])) = StringType
	*(*uint64)(unsafe.Pointer(&data[8])) = uint64(strLen)

	if strLen > 0 {
		copy(data[16:16+strLen], value)
	}

	return uintptr(unsafe.Pointer(&data[0])), nil
}

/**
 * @brief Retrieves a string value from arena memory
 *
 * @param addr Memory address of the stored value
 * @return The string value and error if type mismatch
 */
func getString(addr uintptr) (string, error) {
	if *(*DataType)(unsafe.Pointer(addr)) != StringType {
		return "", fmt.Errorf("%w: expected %v, got %v", ErrTypeMismatch, StringType, *(*DataType)(unsafe.Pointer(addr)))
	}

	strLen := *(*uint64)(unsafe.Pointer(addr + 8))
	if strLen == 0 {
		return "", nil
	}

	strData := unsafe.Slice((*byte)(unsafe.Pointer(addr+16)), strLen)
	return unsafe.String(&strData[0], int(strLen)), nil
}

/**
 * @brief Stores an int64 slice in arena memory with zero-copy retrieval support
 *
 * @param arr The int64 slice to store
 * @param arena Memory arena for allocation
 * @return Memory address of stored slice and error if allocation fails
 *
 * Memory Layout (designed for zero-copy access):
 * [DataType:8] + [SliceLen:8] + [Data:arrLen*8]
 *
 * The layout directly stores the slice data after the header, allowing
 * unsafe.Slice to create a zero-copy view of the data.
 */
func putInt64s(arr []int64, arena *Arena) (uintptr, error) {
	arrLen := len(arr)
	if arrLen == 0 {
		// Empty slice: [DataType:8] + [SliceLen:8] = 16 bytes
		data, err := arena.allocate(16)
		if err != nil {
			return 0, err
		}

		*(*DataType)(unsafe.Pointer(&data[0])) = Int64sType
		*(*uint64)(unsafe.Pointer(&data[8])) = 0

		return uintptr(unsafe.Pointer(&data[0])), nil
	}

	// Calculate required space: header + data
	dataSize := uintptr(arrLen) * 8 // int64 is 8 bytes each
	totalSize := 16 + dataSize      // 16 bytes header + data

	data, err := arena.allocate(totalSize)
	if err != nil {
		return 0, err
	}

	*(*DataType)(unsafe.Pointer(&data[0])) = Int64sType
	*(*uint64)(unsafe.Pointer(&data[8])) = uint64(arrLen)

	// Copy slice data directly after header
	if arrLen > 0 {
		dataStart := unsafe.Pointer(&data[16])
		dataSlice := unsafe.Slice((*int64)(dataStart), arrLen)
		copy(dataSlice, arr)
	}

	return uintptr(unsafe.Pointer(&data[0])), nil
}

/**
 * @brief Retrieves an int64 slice from arena memory with zero-copy access
 *
 * @param addr Memory address of the stored slice
 * @return The int64 slice pointing directly to arena memory, error if type mismatch
 *
 * Returns a slice that directly references the arena memory without copying.
 */
func getInt64s(addr uintptr) ([]int64, error) {
	if *(*DataType)(unsafe.Pointer(addr)) != Int64sType {
		return nil, fmt.Errorf("%w: expected %v, got %v", ErrTypeMismatch, Int64sType, *(*DataType)(unsafe.Pointer(addr)))
	}

	arrLen := *(*uint64)(unsafe.Pointer(addr + 8))
	if arrLen == 0 {
		return nil, nil
	}

	// Create zero-copy slice pointing directly to arena data
	dataStart := unsafe.Pointer(addr + 16)
	return unsafe.Slice((*int64)(dataStart), int(arrLen)), nil
}

/**
 * @brief Stores a float32 slice in arena memory with zero-copy retrieval support
 *
 * @param arr The float32 slice to store
 * @param arena Memory arena for allocation
 * @return Memory address of stored slice and error if allocation fails
 *
 * Memory Layout (designed for zero-copy access):
 * [DataType:8] + [SliceLen:8] + [Data:aligned(arrLen*4)]
 *
 * Float32 data is aligned to 8-byte boundary for consistent alignment.
 */
func putFloat32s(arr []float32, arena *Arena) (uintptr, error) {
	arrLen := len(arr)
	if arrLen == 0 {
		// Empty slice: [DataType:8] + [SliceLen:8] = 16 bytes
		data, err := arena.allocate(16)
		if err != nil {
			return 0, err
		}

		*(*DataType)(unsafe.Pointer(&data[0])) = Float32sType
		*(*uint64)(unsafe.Pointer(&data[8])) = 0

		return uintptr(unsafe.Pointer(&data[0])), nil
	}

	// Calculate required space with alignment
	dataSize := alignSize(uintptr(arrLen) * 4) // float32 is 4 bytes each, aligned to 8
	totalSize := 16 + dataSize                 // 16 bytes header + aligned data

	data, err := arena.allocate(totalSize)
	if err != nil {
		return 0, err
	}

	*(*DataType)(unsafe.Pointer(&data[0])) = Float32sType
	*(*uint64)(unsafe.Pointer(&data[8])) = uint64(arrLen)

	// Copy slice data directly after header
	if arrLen > 0 {
		dataStart := unsafe.Pointer(&data[16])
		dataSlice := unsafe.Slice((*float32)(dataStart), arrLen)
		copy(dataSlice, arr)
	}

	return uintptr(unsafe.Pointer(&data[0])), nil
}

/**
 * @brief Retrieves a float32 slice from arena memory with zero-copy access
 *
 * @param addr Memory address of the stored slice
 * @return The float32 slice pointing directly to arena memory, error if type mismatch
 *
 * Returns a slice that directly references the arena memory without copying.
 */
func getFloat32s(addr uintptr) ([]float32, error) {
	if *(*DataType)(unsafe.Pointer(addr)) != Float32sType {
		return nil, fmt.Errorf("%w: expected %v, got %v", ErrTypeMismatch, Float32sType, *(*DataType)(unsafe.Pointer(addr)))
	}

	arrLen := *(*uint64)(unsafe.Pointer(addr + 8))
	if arrLen == 0 {
		return nil, nil
	}

	// Create zero-copy slice pointing directly to arena data
	dataStart := unsafe.Pointer(addr + 16)
	return unsafe.Slice((*float32)(dataStart), int(arrLen)), nil
}

/**
 * @brief Stores a string slice in arena memory with zero-copy retrieval support
 *
 * @param arr The string slice to store
 * @param arena Memory arena for allocation
 * @return Memory address of stored slice and error if allocation fails
 *
 * Memory Layout (designed for zero-copy access):
 * [DataType:8] + [SliceLen:8] + [StringHeaders:arrLen*16] + [StringData:aligned]
 *
 * Where:
 * - SliceLen: Number of strings in the slice
 * - StringHeaders: Array of string headers (pointer + length) for each string
 * - StringData: Concatenated string data referenced by string headers
 *
 * This layout allows unsafe.Slice to create a zero-copy []string view.
 */
func putStrings(arr []string, arena *Arena) (uintptr, error) {
	arrLen := len(arr)
	if arrLen == 0 {
		// Empty slice: [DataType:8] + [SliceLen:8] = 16 bytes
		data, err := arena.allocate(16)
		if err != nil {
			return 0, err
		}

		*(*DataType)(unsafe.Pointer(&data[0])) = StringsType
		*(*uint64)(unsafe.Pointer(&data[8])) = 0

		return uintptr(unsafe.Pointer(&data[0])), nil
	}

	// Calculate required space
	stringHeadersSize := alignSize(uintptr(arrLen) * 16) // Each string header: 16 bytes (ptr+len), aligned

	var totalStrLen uintptr
	for _, s := range arr {
		totalStrLen += uintptr(len(s))
	}
	stringDataSize := alignSize(totalStrLen)

	totalSize := 16 + stringHeadersSize + stringDataSize
	data, err := arena.allocate(totalSize)
	if err != nil {
		return 0, err
	}

	*(*DataType)(unsafe.Pointer(&data[0])) = StringsType
	*(*uint64)(unsafe.Pointer(&data[8])) = uint64(arrLen)

	// Setup string headers and data
	stringHeadersStart := 16
	stringDataStart := stringHeadersStart + int(stringHeadersSize)
	var currentOffset uintptr

	for i, s := range arr {
		strLen := len(s)

		// Calculate header position (each header is 16 bytes)
		headerPos := stringHeadersStart + i*16

		if strLen > 0 {
			// Copy string data
			copy(data[stringDataStart+int(currentOffset):], s)

			// Create string header: [Data:8] + [Len:8]
			// Use unsafe.String to create the string pointing to arena memory
			strPtr := unsafe.Pointer(&data[stringDataStart+int(currentOffset)])
			str := unsafe.String((*byte)(strPtr), strLen)

			// Store the string in the header area
			*(*string)(unsafe.Pointer(&data[headerPos])) = str

			currentOffset += uintptr(strLen)
		} else {
			// Empty string
			*(*string)(unsafe.Pointer(&data[headerPos])) = ""
		}
	}

	return uintptr(unsafe.Pointer(&data[0])), nil
}

/**
 * @brief Retrieves a string slice from arena memory with zero-copy access
 *
 * @param addr Memory address of the stored slice
 * @return The string slice pointing directly to arena memory, error if type mismatch
 *
 * This function creates a slice that directly references the string headers
 * stored in arena memory, achieving true zero-copy access.
 */
func getStrings(addr uintptr) ([]string, error) {
	if *(*DataType)(unsafe.Pointer(addr)) != StringsType {
		return nil, fmt.Errorf("%w: expected %v, got %v", ErrTypeMismatch, StringsType, *(*DataType)(unsafe.Pointer(addr)))
	}

	arrLen := *(*uint64)(unsafe.Pointer(addr + 8))
	if arrLen == 0 {
		return nil, nil
	}

	// Create zero-copy slice pointing directly to the string headers in arena memory
	stringHeadersStart := addr + 16

	// Each string is stored as a complete Go string (16 bytes on 64-bit)
	// Use unsafe.Slice to create a slice view of the stored strings
	return unsafe.Slice((*string)(unsafe.Pointer(stringHeadersStart)), int(arrLen)), nil
}

/**
 * @brief Stores any supported value type in arena memory
 *
 * @param value The value to store (must be supported type)
 * @param arena Memory arena for allocation
 * @return Memory address of stored value and error if type unsupported or allocation fails
 *
 * Supported Types:
 * - int, int64 -> Int64Type
 * - float32, float64 -> Float32Type (float64 converted to float32)
 * - string -> StringType
 * - []int, []int64 -> Int64sType
 * - []float32, []float64 -> Float32sType
 * - []string -> StringsType
 */
func putAnyValue(value any, arena *Arena) (uintptr, error) {
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
		int64s := make([]int64, len(v))
		for i, val := range v {
			int64s[i] = int64(val)
		}
		return putInt64s(int64s, arena)
	case []float32:
		return putFloat32s(v, arena)
	case []float64:
		float32s := make([]float32, len(v))
		for i, val := range v {
			float32s[i] = float32(val)
		}
		return putFloat32s(float32s, arena)
	case []string:
		return putStrings(v, arena)
	default:
		return 0, fmt.Errorf("unsupported type: %T", value)
	}
}

/**
 * @brief Creates deep copies of string slices using Go 1.20+ strings.Clone
 *
 * @param arr The string slice to copy
 * @return Deep copy of the string slice
 *
 * Uses strings.Clone (Go 1.20+) for efficient string copying that ensures
 * the copied strings are independent of the original memory.
 */
func deepcopyStrings(arr []string) []string {
	if len(arr) == 0 {
		return nil
	}

	ret := make([]string, len(arr))
	for i, s := range arr {
		ret[i] = strings.Clone(s)
	}
	return ret
}
