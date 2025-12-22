package sample

import (
	"encoding/json"
	"fmt"
	"strings"
	"unsafe"

	"github.com/bytedance/sonic"
)

/**
 * @brief MutableFeatures is a collection of mutable typed features
 *
 * This collection allows modification of features after creation, unlike ImmutableFeatures.
 * It provides type-safe operations and supports conversion to immutable features
 * for performance-critical read operations.
 *
 * Key characteristics:
 * - Mutable: Features can be added, modified, or removed after creation
 * - Type-safe: Runtime type checking prevents type mismatches
 * - JSON serialization: Full support for marshaling/unmarshaling
 * - Conversion: Can be converted to ImmutableFeatures for zero-copy access
 *
 * Note: This collection is NOT thread-safe. External synchronization is required
 * for concurrent access.
 */
type MutableFeatures struct {
	features map[string]Feature // Map of feature name to feature implementation
}

/**
 * @brief Creates a new empty MutableFeatures collection
 *
 * @return Pointer to new MutableFeatures collection
 */
func NewMutableFeatures() *MutableFeatures {
	return &MutableFeatures{
		features: make(map[string]Feature),
	}
}

/**
 * @brief Creates MutableFeatures from a map of values
 *
 * @param data Map of feature names to values
 * @return Pointer to new MutableFeatures and error if conversion fails
 *
 * Supported value types:
 * - int, int64: stored as Int64 feature
 * - float32, float64: stored as Float32 feature (float64 converted to float32)
 * - string: stored as String feature
 * - []int, []int64: stored as Int64s feature
 * - []float32, []float64: stored as Float32s feature (float64 converted to float32)
 * - []string: stored as Strings feature
 */
func NewMutableFeaturesFromMap(data map[string]any) (*MutableFeatures, error) {
	features := NewMutableFeatures()

	for key, value := range data {
		feature := createFeatureFromValue(value)
		if feature == nil {
			return nil, fmt.Errorf("failed to create feature for key %s: %v", key, value)
		}
		features.features[key] = feature
	}

	return features, nil
}

/**
 * @brief Gets the data type of a feature by key
 *
 * @param key Feature name
 * @return DataType of the feature, InvalidType if not found
 */
func (f *MutableFeatures) GetType(key string) DataType {
	if feature, ok := f.features[key]; ok {
		return feature.Type()
	}
	return InvalidType
}

/**
 * @brief Returns all feature names
 *
 * @return Slice of feature names in no particular order
 */
func (f *MutableFeatures) Keys() []string {
	ret := make([]string, 0, len(f.features))
	for key := range f.features {
		ret = append(ret, key)
	}
	return ret
}

/**
 * @brief Returns the number of features in the collection
 *
 * @return Number of features
 */
func (f *MutableFeatures) Len() int {
	return len(f.features)
}

/**
 * @brief Retrieves a feature by name
 *
 * @param key Feature name
 * @return Feature interface or nil if not found
 */
func (f *MutableFeatures) Get(key string) Feature {
	if value, ok := f.features[key]; ok {
		return value
	}
	return nil
}

/**
 * @brief Checks if a feature exists
 *
 * @param key Feature name
 * @return True if feature exists, false otherwise
 */
func (f *MutableFeatures) Has(key string) bool {
	_, ok := f.features[key]
	return ok
}

/**
 * @brief Sets or updates a feature
 *
 * @param key Feature name
 * @param value Feature implementation to store
 */
func (f *MutableFeatures) Set(key string, value Feature) {
	f.features[key] = value
}

/**
 * @brief Sets a feature from any supported value type
 *
 * @param key Feature name
 * @param value Value to store (must be supported type)
 * @return Error if value type is unsupported
 */
func (f *MutableFeatures) SetValue(key string, value any) error {
	feature := createFeatureFromValue(value)
	if feature == nil {
		return fmt.Errorf("failed to create feature for key %s: %v", key, value)
	}
	f.features[key] = feature
	return nil
}

/**
 * @brief Removes a feature by key
 *
 * @param key Feature name to remove
 * @return True if feature was removed, false if not found
 */
func (f *MutableFeatures) Delete(key string) bool {
	if _, exists := f.features[key]; exists {
		delete(f.features, key)
		return true
	}
	return false
}

/**
 * @brief Iterates over all features with a callback function
 *
 * @param fn Callback function called for each feature
 * @return Error if callback returns error
 */
func (f *MutableFeatures) ForEach(fn IteratorFunc) error {
	for key, feature := range f.features {
		if err := fn(key, feature); err != nil {
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
func (f *MutableFeatures) MapAny() (map[string]any, error) {
	feas := make(map[string]any, len(f.features))

	for key, value := range f.features {
		switch value.Type() {
		case Int64Type:
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value int64    `json:"value"`
			}{Int64Type, value.GetInt64Unsafe()}
		case Float32Type:
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value float32  `json:"value"`
			}{Float32Type, value.GetFloat32Unsafe()}

		case StringType:
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value string   `json:"value"`
			}{StringType, value.GetStringUnsafe()}

		case Int64sType:
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value []int64  `json:"value"`
			}{Int64sType, value.GetInt64sUnsafe()}

		case Float32sType:
			feas[key] = struct {
				Type  DataType  `json:"type"`
				Value []float32 `json:"value"`
			}{Float32sType, value.GetFloat32sUnsafe()}

		case StringsType:
			feas[key] = struct {
				Type  DataType `json:"type"`
				Value []string `json:"value"`
			}{StringsType, value.GetStringsUnsafe()}
		default:
			return nil, fmt.Errorf("unknown data type %v for key %s", value.Type(), key)
		}
	}
	return feas, nil
}

/**
 * @brief Marshals features to JSON
 *
 * @return JSON bytes and error if marshaling fails
 */
func (f *MutableFeatures) MarshalJSON() ([]byte, error) {
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

	// Initialize features map if not already initialized
	if f.features == nil {
		f.features = make(map[string]Feature)
	}

	for key, value := range fea {
		switch value.Type {
		case Int64Type:
			var num int64
			if err := sonic.Unmarshal(value.Value, &num); err != nil {
				return fmt.Errorf("failed to unmarshal int64 for key %s: %w", key, err)
			}
			f.features[key] = &Int64{Value: num}

		case Float32Type:
			var num float32
			if err := sonic.Unmarshal(value.Value, &num); err != nil {
				return fmt.Errorf("failed to unmarshal float32 for key %s: %w", key, err)
			}
			f.features[key] = &Float32{Value: num}

		case StringType:
			var str string
			if err := sonic.Unmarshal(value.Value, &str); err != nil {
				return fmt.Errorf("failed to unmarshal string for key %s: %w", key, err)
			}
			f.features[key] = &String{Value: str}

		case Int64sType:
			var nums []int64
			if err := sonic.Unmarshal(value.Value, &nums); err != nil {
				return fmt.Errorf("failed to unmarshal int64s for key %s: %w", key, err)
			}
			f.features[key] = &Int64s{Value: nums}

		case Float32sType:
			var nums []float32
			if err := sonic.Unmarshal(value.Value, &nums); err != nil {
				return fmt.Errorf("failed to unmarshal float32s for key %s: %w", key, err)
			}
			f.features[key] = &Float32s{Value: nums}

		case StringsType:
			var strs []string
			if err := sonic.Unmarshal(value.Value, &strs); err != nil {
				return fmt.Errorf("failed to unmarshal strings for key %s: %w", key, err)
			}
			f.features[key] = &Strings{Value: strs}

		default:
			return fmt.Errorf("unknown data type %v for key %s", value.Type, key)
		}
	}
	return nil
}

/**
 * @brief Converts to ImmutableFeatures for zero-copy read access
 *
 * @param arena Memory arena for storage
 * @return Pointer to new ImmutableFeatures and error if conversion fails
 *
 * All values are deep-copied into arena memory to ensure independence
 * from the original mutable features.
 */
func (f *MutableFeatures) Immutable(arena *Arena) (*ImmutableFeatures, error) {
	immutable := NewImmutableFeatures(arena)

	for key, feature := range f.features {
		var ptr unsafe.Pointer
		var err error

		switch feature.Type() {
		case Int64Type:
			ptr, err = putInt64(feature.GetInt64Unsafe(), arena)
		case Float32Type:
			ptr, err = putFloat32(feature.GetFloat32Unsafe(), arena)
		case StringType:
			ptr, err = putString(feature.GetStringUnsafe(), arena)
		case Int64sType:
			ptr, err = putInt64s(feature.GetInt64sUnsafe(), arena)
		case Float32sType:
			ptr, err = putFloat32s(feature.GetFloat32sUnsafe(), arena)
		case StringsType:
			ptr, err = putStrings(feature.GetStringsUnsafe(), arena)
		default:
			err = fmt.Errorf("unsupported feature type %v for key %s", feature.Type(), key)
		}

		if err != nil {
			return nil, fmt.Errorf("failed to convert feature %s: %w", key, err)
		}

		immutable.features[key] = ImmutableFeature{ptr: ptr}
	}

	return immutable, nil
}

/**
 * @brief Creates a deep copy of the MutableFeatures
 *
 * @return Pointer to new MutableFeatures with copied values
 *
 * All values are deep-copied to ensure the copy is independent.
 * This includes deep copying of slices and strings using strings.Clone.
 */
func (f *MutableFeatures) Clone() *MutableFeatures {
	clone := NewMutableFeatures()

	for key, feature := range f.features {
		switch feature.Type() {
		case Int64Type:
			if v, err := feature.GetInt64(); err == nil {
				clone.features[key] = &Int64{Value: v}
			}
		case Float32Type:
			if v, err := feature.GetFloat32(); err == nil {
				clone.features[key] = &Float32{Value: v}
			}
		case StringType:
			if v, err := feature.GetString(); err == nil {
				clone.features[key] = &String{Value: strings.Clone(v)}
			}
		case Int64sType:
			if v, err := feature.GetInt64s(); err == nil {
				nums := make([]int64, len(v))
				copy(nums, v)
				clone.features[key] = &Int64s{Value: nums}
			}
		case Float32sType:
			if v, err := feature.GetFloat32s(); err == nil {
				nums := make([]float32, len(v))
				copy(nums, v)
				clone.features[key] = &Float32s{Value: nums}
			}
		case StringsType:
			if v, err := feature.GetStrings(); err == nil {
				strs := make([]string, len(v))
				for i, s := range v {
					strs[i] = strings.Clone(s)
				}
				clone.features[key] = &Strings{Value: strs}
			}
		}
	}

	return clone
}

// Feature implementations with embedded error handling

/**
 * @brief ErrorFeature provides default error implementations for all Feature methods
 *
 * This struct is embedded in concrete feature types to provide default error
 * responses for methods that don't apply to the specific type. This follows
 * the composition pattern to avoid code duplication.
 */
type ErrorFeature struct{}

/**
 * @brief Returns InvalidType for error features
 *
 * @return InvalidType
 */
func (f *ErrorFeature) Type() DataType {
	return InvalidType
}

/**
 * @brief Retrieves the stored value as any type
 *
 * @return nil for error features
 */
func (f *ErrorFeature) Get() any {
	return nil
}

/**
 * @brief Default error implementation for GetInt64
 *
 * @return Zero value and ErrNotImplemented error
 */
func (f *ErrorFeature) GetInt64() (int64, error) {
	return 0, ErrNotImplemented
}

/**
 * @brief Default unsafe implementation for GetInt64Unsafe
 *
 * @return Zero value
 */
func (f *ErrorFeature) GetInt64Unsafe() int64 {
	return 0
}

/**
 * @brief Default error implementation for GetFloat32
 *
 * @return Zero value and ErrNotImplemented error
 */
func (f *ErrorFeature) GetFloat32() (float32, error) {
	return 0.0, ErrNotImplemented
}

/**
 * @brief Default unsafe implementation for GetFloat32Unsafe
 *
 * @return Zero value
 */
func (f *ErrorFeature) GetFloat32Unsafe() float32 {
	return 0.0
}

/**
 * @brief Default error implementation for GetString
 *
 * @return Empty string and ErrNotImplemented error
 */
func (f *ErrorFeature) GetString() (string, error) {
	return "", ErrNotImplemented
}

/**
 * @brief Default unsafe implementation for GetStringUnsafe
 *
 * @return Empty string
 */
func (f *ErrorFeature) GetStringUnsafe() string {
	return ""
}

/**
 * @brief Default error implementation for GetInt64s
 *
 * @return Nil slice and ErrNotImplemented error
 */
func (f *ErrorFeature) GetInt64s() ([]int64, error) {
	return nil, ErrNotImplemented
}

/**
 * @brief Default unsafe implementation for GetInt64sUnsafe
 *
 * @return Nil slice
 */
func (f *ErrorFeature) GetInt64sUnsafe() []int64 {
	return nil
}

/**
 * @brief Default error implementation for GetFloat32s
 *
 * @return Nil slice and ErrNotImplemented error
 */
func (f *ErrorFeature) GetFloat32s() ([]float32, error) {
	return nil, ErrNotImplemented
}

/**
 * @brief Default unsafe implementation for GetFloat32sUnsafe
 *
 * @return Nil slice
 */
func (f *ErrorFeature) GetFloat32sUnsafe() []float32 {
	return nil
}

/**
 * @brief Default error implementation for GetStrings
 *
 * @return Nil slice and ErrNotImplemented error
 */
func (f *ErrorFeature) GetStrings() ([]string, error) {
	return nil, ErrNotImplemented
}

/**
 * @brief Default unsafe implementation for GetStringsUnsafe
 *
 * @return Nil slice
 */
func (f *ErrorFeature) GetStringsUnsafe() []string {
	return nil
}

/**
 * @brief Int64 feature stores a 64-bit signed integer value
 */
type Int64 struct {
	ErrorFeature
	Value int64 // The stored integer value
}

/**
 * @brief Returns Int64Type
 *
 * @return Int64Type
 */
func (f *Int64) Type() DataType {
	return Int64Type
}

/**
 * @brief Retrieves the stored value as any type
 *
 * @return The underlying int64 value as any type
 */
func (f *Int64) Get() any {
	return f.Value
}

/**
 * @brief Retrieves the stored int64 value
 *
 * @return The int64 value and nil error
 */
func (f *Int64) GetInt64() (int64, error) {
	return f.Value, nil
}

/**
 * @brief Retrieves the stored int64 value without error checking
 *
 * @return The int64 value
 */
func (f *Int64) GetInt64Unsafe() int64 {
	return f.Value
}

/**
 * @brief Int64s feature stores a slice of 64-bit signed integers
 */
type Int64s struct {
	ErrorFeature
	Value []int64 // The stored integer slice
}

/**
 * @brief Returns Int64sType
 *
 * @return Int64sType
 */
func (f *Int64s) Type() DataType {
	return Int64sType
}

/**
 * @brief Retrieves the stored value as any type
 *
 * @return The underlying []int64 slice as any type
 */
func (f *Int64s) Get() any {
	return f.Value
}

/**
 * @brief Retrieves the stored int64 slice
 *
 * @return The int64 slice and nil error
 */
func (f *Int64s) GetInt64s() ([]int64, error) {
	return f.Value, nil
}

/**
 * @brief Retrieves the stored int64 slice without error checking
 *
 * @return The int64 slice
 */
func (f *Int64s) GetInt64sUnsafe() []int64 {
	return f.Value
}

/**
 * @brief Float32 feature stores a 32-bit floating point value
 */
type Float32 struct {
	ErrorFeature
	Value float32 // The stored float value
}

/**
 * @brief Returns Float32Type
 *
 * @return Float32Type
 */
func (f *Float32) Type() DataType {
	return Float32Type
}

/**
 * @brief Retrieves the stored value as any type
 *
 * @return The underlying float32 value as any type
 */
func (f *Float32) Get() any {
	return f.Value
}

/**
 * @brief Retrieves the stored float32 value
 *
 * @return The float32 value and nil error
 */
func (f *Float32) GetFloat32() (float32, error) {
	return f.Value, nil
}

/**
 * @brief Retrieves the stored float32 value without error checking
 *
 * @return The float32 value
 */
func (f *Float32) GetFloat32Unsafe() float32 {
	return f.Value
}

/**
 * @brief Float32s feature stores a slice of 32-bit floating point values
 */
type Float32s struct {
	ErrorFeature
	Value []float32 // The stored float slice
}

/**
 * @brief Returns Float32sType
 *
 * @return Float32sType
 */
func (f *Float32s) Type() DataType {
	return Float32sType
}

/**
 * @brief Retrieves the stored value as any type
 *
 * @return The underlying []float32 slice as any type
 */
func (f *Float32s) Get() any {
	return f.Value
}

/**
 * @brief Retrieves the stored float32 slice
 *
 * @return The float32 slice and nil error
 */
func (f *Float32s) GetFloat32s() ([]float32, error) {
	return f.Value, nil
}

/**
 * @brief Retrieves the stored float32 slice without error checking
 *
 * @return The float32 slice
 */
func (f *Float32s) GetFloat32sUnsafe() []float32 {
	return f.Value
}

/**
 * @brief String feature stores a UTF-8 string value
 */
type String struct {
	ErrorFeature
	Value string // The stored string value
}

/**
 * @brief Returns StringType
 *
 * @return StringType
 */
func (f *String) Type() DataType {
	return StringType
}

/**
 * @brief Retrieves the stored value as any type
 *
 * @return The underlying string value as any type
 */
func (f *String) Get() any {
	return f.Value
}

/**
 * @brief Retrieves the stored string value
 *
 * @return The string value and nil error
 */
func (f *String) GetString() (string, error) {
	return f.Value, nil
}

/**
 * @brief Retrieves the stored string value without error checking
 *
 * @return The string value
 */
func (f *String) GetStringUnsafe() string {
	return f.Value
}

/**
 * @brief Strings feature stores a slice of UTF-8 strings
 */
type Strings struct {
	ErrorFeature
	Value []string // The stored string slice
}

/**
 * @brief Returns StringsType
 *
 * @return StringsType
 */
func (f *Strings) Type() DataType {
	return StringsType
}

/**
 * @brief Retrieves the stored value as any type
 *
 * @return The underlying []string slice as any type
 */
func (f *Strings) Get() any {
	return f.Value
}

/**
 * @brief Retrieves the stored string slice
 *
 * @return The string slice and nil error
 */
func (f *Strings) GetStrings() ([]string, error) {
	return f.Value, nil
}

/**
 * @brief Retrieves the stored string slice without error checking
 *
 * @return The string slice
 */
func (f *Strings) GetStringsUnsafe() []string {
	return f.Value
}

// Helper functions

/**
 * @brief Creates a Feature from any supported value type
 *
 * @param value The value to convert to a Feature
 * @return Feature implementation
 *
 * Supported types:
 * - int, int64: converted to Int64 feature
 * - float32, float64: converted to Float32 feature (float64 truncated to float32)
 * - string: converted to String feature
 * - []int, []int64: converted to Int64s feature
 * - []float32, []float64: converted to Float32s feature (float64 truncated to float32)
 * - []string: converted to Strings feature with deep-copied strings
 */
func createFeatureFromValue(value any) Feature {
	switch v := value.(type) {
	case int64:
		return &Int64{Value: v}
	case int:
		return &Int64{Value: int64(v)}
	case float32:
		return &Float32{Value: v}
	case float64:
		return &Float32{Value: float32(v)}
	case string:
		return &String{Value: v}
	case []int64:
		// Create a copy to ensure independence
		nums := make([]int64, len(v))
		copy(nums, v)
		return &Int64s{Value: nums}
	case []int:
		int64s := make([]int64, len(v))
		for i, val := range v {
			int64s[i] = int64(val)
		}
		return &Int64s{Value: int64s}
	case []float32:
		// Create a copy to ensure independence
		nums := make([]float32, len(v))
		copy(nums, v)
		return &Float32s{Value: nums}
	case []float64:
		float32s := make([]float32, len(v))
		for i, val := range v {
			float32s[i] = float32(val)
		}
		return &Float32s{Value: float32s}
	case []string:
		// Deep copy strings to ensure independence
		strs := make([]string, len(v))
		for i, s := range v {
			strs[i] = strings.Clone(s)
		}
		return &Strings{Value: strs}
	default:
		return nil
	}
}
