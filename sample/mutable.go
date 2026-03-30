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

// MutableFeatures is a modifiable collection of typed features. Unlike
// ImmutableFeatures it does not use an arena; each value is a separate heap
// allocation. Use Immutable() to convert to a zero-copy arena representation
// when the set of features is stable.
//
// MutableFeatures is NOT thread-safe.
type MutableFeatures struct {
	features map[string]Feature
}

// NewMutableFeatures returns an empty MutableFeatures.
func NewMutableFeatures() *MutableFeatures {
	return &MutableFeatures{features: make(map[string]Feature)}
}

// NewMutableFeaturesFromMap builds a MutableFeatures from a plain map.
// Supported value types: int, int64, float32, float64, string,
// []int, []int64, []float32, []float64, []string.
func NewMutableFeaturesFromMap(data map[string]any) (*MutableFeatures, error) {
	f := NewMutableFeatures()
	for key, value := range data {
		feat := createFeatureFromValue(value)
		if feat == nil {
			return nil, fmt.Errorf("unsupported value type for key %s: %T", key, value)
		}
		f.features[key] = feat
	}
	return f, nil
}

// GetType returns the DataType for key, or InvalidType when absent.
func (f *MutableFeatures) GetType(key string) DataType {
	if feat, ok := f.features[key]; ok {
		return feat.Type()
	}
	return InvalidType
}

// Keys returns all feature keys in unspecified order.
func (f *MutableFeatures) Keys() []string {
	return slices.Collect(maps.Keys(f.features))
}

// Len returns the number of features.
func (f *MutableFeatures) Len() int { return len(f.features) }

// Get returns the Feature for key, or nil when absent.
func (f *MutableFeatures) Get(key string) Feature {
	return f.features[key] // nil when absent — map returns zero value
}

// Has reports whether key exists.
func (f *MutableFeatures) Has(key string) bool {
	_, ok := f.features[key]
	return ok
}

// Set stores or replaces a feature.
func (f *MutableFeatures) Set(key string, value Feature) {
	f.features[key] = value
}

// SetValue stores a value of any supported type, wrapping it in the
// appropriate Feature implementation.
func (f *MutableFeatures) SetValue(key string, value any) error {
	feat := createFeatureFromValue(value)
	if feat == nil {
		return fmt.Errorf("unsupported value type for key %s: %T", key, value)
	}
	f.features[key] = feat
	return nil
}

// Delete removes key from the collection. It reports whether the key existed.
func (f *MutableFeatures) Delete(key string) bool {
	if _, ok := f.features[key]; !ok {
		return false
	}
	delete(f.features, key)
	return true
}

// ForEach calls fn for every key-feature pair, stopping on the first non-nil error.
func (f *MutableFeatures) ForEach(fn IteratorFunc) error {
	for key, feat := range f.features {
		if err := fn(key, feat); err != nil {
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
func (f *MutableFeatures) All() iter.Seq2[string, Feature] {
	return func(yield func(string, Feature) bool) {
		for key, feat := range f.features {
			if !yield(key, feat) {
				return
			}
		}
	}
}

// MapAny converts every feature to a {type, value} struct for serialization.
func (f *MutableFeatures) MapAny() (map[string]any, error) {
	out := make(map[string]any, len(f.features))
	for key, feat := range f.features {
		v, err := featureToAny(feat)
		if err != nil {
			return nil, fmt.Errorf("key %s: %w", key, err)
		}
		out[key] = v
	}
	return out, nil
}

// MarshalJSON encodes the collection as JSON.
func (f *MutableFeatures) MarshalJSON() ([]byte, error) {
	m, err := f.MapAny()
	if err != nil {
		return nil, err
	}
	return sonic.Marshal(m)
}

// UnmarshalJSON decodes JSON produced by MarshalJSON.
func (f *MutableFeatures) UnmarshalJSON(data []byte) error {
	type wire struct {
		Type  DataType        `json:"type"`
		Value json.RawMessage `json:"value"`
	}
	var raw map[string]wire
	if err := sonic.Unmarshal(data, &raw); err != nil {
		return err
	}
	if f.features == nil {
		f.features = make(map[string]Feature, len(raw))
	}
	for key, w := range raw {
		feat, err := unmarshalMutableValue(w.Type, w.Value)
		if err != nil {
			return fmt.Errorf("key %s: %w", key, err)
		}
		f.features[key] = feat
	}
	return nil
}

// unmarshalMutableValue decodes a single JSON value into a heap-allocated Feature.
func unmarshalMutableValue(dt DataType, raw json.RawMessage) (Feature, error) {
	switch dt {
	case Int64Type:
		var v int64
		if err := sonic.Unmarshal(raw, &v); err != nil {
			return nil, err
		}
		return &Int64{Value: v}, nil
	case Float32Type:
		var v float32
		if err := sonic.Unmarshal(raw, &v); err != nil {
			return nil, err
		}
		return &Float32{Value: v}, nil
	case StringType:
		var v string
		if err := sonic.Unmarshal(raw, &v); err != nil {
			return nil, err
		}
		return &String{Value: v}, nil
	case Int64sType:
		var v []int64
		if err := sonic.Unmarshal(raw, &v); err != nil {
			return nil, err
		}
		return &Int64s{Value: v}, nil
	case Float32sType:
		var v []float32
		if err := sonic.Unmarshal(raw, &v); err != nil {
			return nil, err
		}
		return &Float32s{Value: v}, nil
	case StringsType:
		var v []string
		if err := sonic.Unmarshal(raw, &v); err != nil {
			return nil, err
		}
		return &Strings{Value: v}, nil
	default:
		return nil, fmt.Errorf("unknown data type %v", dt)
	}
}

// Immutable copies all features into arena-allocated memory and returns an
// ImmutableFeatures for zero-copy read access.
func (f *MutableFeatures) Immutable(arena *Arena) (*ImmutableFeatures, error) {
	imm := NewImmutableFeatures(arena)
	for key, feat := range f.features {
		ptr, err := putFeature(feat, arena)
		if err != nil {
			return nil, fmt.Errorf("key %s: %w", key, err)
		}
		imm.features[key] = ImmutableFeature{ptr: ptr}
	}
	return imm, nil
}

// putFeature serialises feat into arena, dispatching on feat.Type().
func putFeature(feat Feature, arena *Arena) (unsafe.Pointer, error) {
	switch feat.Type() {
	case Int64Type:
		return putInt64(feat.GetInt64Unsafe(), arena)
	case Float32Type:
		return putFloat32(feat.GetFloat32Unsafe(), arena)
	case StringType:
		return putString(feat.GetStringUnsafe(), arena)
	case Int64sType:
		return putInt64s(feat.GetInt64sUnsafe(), arena)
	case Float32sType:
		return putFloat32s(feat.GetFloat32sUnsafe(), arena)
	case StringsType:
		return putStrings(feat.GetStringsUnsafe(), arena)
	default:
		return nil, fmt.Errorf("unsupported feature type %v", feat.Type())
	}
}

// Clone returns a deep copy of the collection. All slice and string values are
// independently copied so mutations to the clone do not affect the original.
func (f *MutableFeatures) Clone() *MutableFeatures {
	clone := &MutableFeatures{features: make(map[string]Feature, len(f.features))}
	for key, feat := range f.features {
		// Type is already known from the switch — use Unsafe getters to avoid
		// a redundant type check on every field.
		switch feat.Type() {
		case Int64Type:
			clone.features[key] = &Int64{Value: feat.GetInt64Unsafe()}
		case Float32Type:
			clone.features[key] = &Float32{Value: feat.GetFloat32Unsafe()}
		case StringType:
			clone.features[key] = &String{Value: strings.Clone(feat.GetStringUnsafe())}
		case Int64sType:
			src := feat.GetInt64sUnsafe()
			dst := make([]int64, len(src))
			copy(dst, src)
			clone.features[key] = &Int64s{Value: dst}
		case Float32sType:
			src := feat.GetFloat32sUnsafe()
			dst := make([]float32, len(src))
			copy(dst, src)
			clone.features[key] = &Float32s{Value: dst}
		case StringsType:
			clone.features[key] = &Strings{Value: deepcopyStrings(feat.GetStringsUnsafe())}
		}
	}
	return clone
}

// ---------------------------------------------------------------------------
// Concrete Feature implementations
// ---------------------------------------------------------------------------

// ErrorFeature is embedded in every concrete Feature type to provide default
// "not implemented" responses for getter methods that don't apply to that type.
type ErrorFeature struct{}

func (f *ErrorFeature) Type() DataType            { return InvalidType }
func (f *ErrorFeature) Get() any                  { return nil }
func (f *ErrorFeature) GetInt64() (int64, error)  { return 0, ErrNotImplemented }
func (f *ErrorFeature) GetInt64Unsafe() int64     { return 0 }
func (f *ErrorFeature) GetFloat32() (float32, error) { return 0, ErrNotImplemented }
func (f *ErrorFeature) GetFloat32Unsafe() float32 { return 0 }
func (f *ErrorFeature) GetString() (string, error) { return "", ErrNotImplemented }
func (f *ErrorFeature) GetStringUnsafe() string   { return "" }
func (f *ErrorFeature) GetInt64s() ([]int64, error) { return nil, ErrNotImplemented }
func (f *ErrorFeature) GetInt64sUnsafe() []int64  { return nil }
func (f *ErrorFeature) GetFloat32s() ([]float32, error) { return nil, ErrNotImplemented }
func (f *ErrorFeature) GetFloat32sUnsafe() []float32 { return nil }
func (f *ErrorFeature) GetStrings() ([]string, error) { return nil, ErrNotImplemented }
func (f *ErrorFeature) GetStringsUnsafe() []string { return nil }

// Int64 stores a single int64 value.
type Int64 struct {
	ErrorFeature
	Value int64
}

func (f *Int64) Type() DataType           { return Int64Type }
func (f *Int64) Get() any                 { return f.Value }
func (f *Int64) GetInt64() (int64, error) { return f.Value, nil }
func (f *Int64) GetInt64Unsafe() int64    { return f.Value }

// Int64s stores a []int64 value.
type Int64s struct {
	ErrorFeature
	Value []int64
}

func (f *Int64s) Type() DataType             { return Int64sType }
func (f *Int64s) Get() any                   { return f.Value }
func (f *Int64s) GetInt64s() ([]int64, error) { return f.Value, nil }
func (f *Int64s) GetInt64sUnsafe() []int64   { return f.Value }

// Float32 stores a single float32 value.
type Float32 struct {
	ErrorFeature
	Value float32
}

func (f *Float32) Type() DataType              { return Float32Type }
func (f *Float32) Get() any                    { return f.Value }
func (f *Float32) GetFloat32() (float32, error) { return f.Value, nil }
func (f *Float32) GetFloat32Unsafe() float32   { return f.Value }

// Float32s stores a []float32 value.
type Float32s struct {
	ErrorFeature
	Value []float32
}

func (f *Float32s) Type() DataType               { return Float32sType }
func (f *Float32s) Get() any                     { return f.Value }
func (f *Float32s) GetFloat32s() ([]float32, error) { return f.Value, nil }
func (f *Float32s) GetFloat32sUnsafe() []float32 { return f.Value }

// String stores a single string value.
type String struct {
	ErrorFeature
	Value string
}

func (f *String) Type() DataType            { return StringType }
func (f *String) Get() any                  { return f.Value }
func (f *String) GetString() (string, error) { return f.Value, nil }
func (f *String) GetStringUnsafe() string   { return f.Value }

// Strings stores a []string value.
type Strings struct {
	ErrorFeature
	Value []string
}

func (f *Strings) Type() DataType              { return StringsType }
func (f *Strings) Get() any                    { return f.Value }
func (f *Strings) GetStrings() ([]string, error) { return f.Value, nil }
func (f *Strings) GetStringsUnsafe() []string  { return f.Value }

// createFeatureFromValue wraps value in the appropriate Feature implementation.
// Returns nil for unsupported types.
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
		dst := make([]int64, len(v))
		copy(dst, v)
		return &Int64s{Value: dst}
	case []int:
		dst := make([]int64, len(v))
		for i, x := range v {
			dst[i] = int64(x)
		}
		return &Int64s{Value: dst}
	case []float32:
		dst := make([]float32, len(v))
		copy(dst, v)
		return &Float32s{Value: dst}
	case []float64:
		dst := make([]float32, len(v))
		for i, x := range v {
			dst[i] = float32(x)
		}
		return &Float32s{Value: dst}
	case []string:
		return &Strings{Value: deepcopyStrings(v)}
	default:
		return nil
	}
}
