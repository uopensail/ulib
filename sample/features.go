package sample

import (
	"errors"
	"fmt"
	"iter"
)

var (
	// ErrNotImplemented is returned by Feature getter methods that do not
	// apply to the feature's actual type.
	ErrNotImplemented = errors.New("method not implemented for this type")
	// ErrKeyNotFound is returned when a requested key is absent.
	ErrKeyNotFound = errors.New("key not found")
	// ErrTypeMismatch is returned when a type-safe getter is called on a
	// feature whose stored type differs from the requested type.
	ErrTypeMismatch = errors.New("type mismatch")
	// ErrInvalidData indicates that stored data is corrupt or unreadable.
	ErrInvalidData = errors.New("invalid data")
)

// DataType identifies the Go type held inside a Feature.
// It is stored as the first 8 bytes of every arena-allocated value so that
// the correct getter can be dispatched without a separate type field.
type DataType uint64

const (
	Int64Type    DataType = iota // int64
	Float32Type                  // float32
	StringType                   // string
	Int64sType                   // []int64
	Float32sType                 // []float32
	StringsType                  // []string
	InvalidType  DataType = 127  // uninitialized / unknown
)

// String returns the human-readable name of dt.
func (dt DataType) String() string {
	switch dt {
	case Int64Type:
		return "Int64"
	case Float32Type:
		return "Float32"
	case StringType:
		return "String"
	case Int64sType:
		return "Int64s"
	case Float32sType:
		return "Float32s"
	case StringsType:
		return "Strings"
	case InvalidType:
		return "Invalid"
	default:
		return "Unknown"
	}
}

// IsValid reports whether dt is a recognized scalar or slice type.
func (dt DataType) IsValid() bool {
	return dt >= Int64Type && dt <= StringsType
}

// IsSliceType reports whether dt represents a slice value.
func (dt DataType) IsSliceType() bool {
	return dt == Int64sType || dt == Float32sType || dt == StringsType
}

// Feature holds a single strongly-typed value. The safe getter methods
// (GetInt64, GetFloat32, …) return an error on type mismatch; the Unsafe
// variants skip that check and must only be called after verifying Type().
type Feature interface {
	// Type returns the DataType of the stored value.
	Type() DataType

	// Get returns the value as any, using the native Go type.
	Get() any

	GetInt64() (int64, error)
	GetInt64Unsafe() int64

	GetFloat32() (float32, error)
	GetFloat32Unsafe() float32

	GetString() (string, error)
	GetStringUnsafe() string

	GetInt64s() ([]int64, error)
	GetInt64sUnsafe() []int64

	GetFloat32s() ([]float32, error)
	GetFloat32sUnsafe() []float32

	GetStrings() ([]string, error)
	GetStringsUnsafe() []string
}

// IteratorFunc is the callback for Features.ForEach. Returning a non-nil
// error stops iteration and causes ForEach to return that error.
type IteratorFunc func(key string, feature Feature) error

// Features is the common interface for both MutableFeatures and
// ImmutableFeatures. Methods that iterate over the collection are available
// in two forms:
//
//   - ForEach — callback-based, compatible with all Go versions.
//   - All — returns an iter.Seq2 usable with range (Go 1.23+).
type Features interface {
	// GetType returns the DataType for key, or InvalidType if absent.
	GetType(key string) DataType
	// Keys returns all feature keys in unspecified order.
	Keys() []string
	// Get returns the Feature for key, or nil if absent.
	Get(key string) Feature
	// Len returns the number of features.
	Len() int
	// Has reports whether key exists.
	Has(key string) bool
	// MapAny converts each feature to a typed struct for serialization.
	MapAny() (map[string]any, error)
	// MarshalJSON encodes all features as JSON.
	MarshalJSON() ([]byte, error)
	// UnmarshalJSON decodes JSON produced by MarshalJSON.
	UnmarshalJSON(data []byte) error
	// ForEach calls fn for every key-feature pair. It stops on the first
	// non-nil error returned by fn.
	ForEach(fn IteratorFunc) error
	// All returns an iterator over all key-feature pairs.
	//
	//  for key, f := range features.All() {
	//      fmt.Println(key, f.Type())
	//  }
	All() iter.Seq2[string, Feature]
}

// featureToAny converts f to an anonymous struct with "type" and "value" JSON
// fields. Both MutableFeatures.MapAny and ImmutableFeatures.MapAny delegate
// here to avoid duplicating the type-switch.
func featureToAny(f Feature) (any, error) {
	switch f.Type() {
	case Int64Type:
		return struct {
			Type  DataType `json:"type"`
			Value int64    `json:"value"`
		}{Int64Type, f.GetInt64Unsafe()}, nil
	case Float32Type:
		return struct {
			Type  DataType `json:"type"`
			Value float32  `json:"value"`
		}{Float32Type, f.GetFloat32Unsafe()}, nil
	case StringType:
		return struct {
			Type  DataType `json:"type"`
			Value string   `json:"value"`
		}{StringType, f.GetStringUnsafe()}, nil
	case Int64sType:
		return struct {
			Type  DataType `json:"type"`
			Value []int64  `json:"value"`
		}{Int64sType, f.GetInt64sUnsafe()}, nil
	case Float32sType:
		return struct {
			Type  DataType  `json:"type"`
			Value []float32 `json:"value"`
		}{Float32sType, f.GetFloat32sUnsafe()}, nil
	case StringsType:
		return struct {
			Type  DataType `json:"type"`
			Value []string `json:"value"`
		}{StringsType, f.GetStringsUnsafe()}, nil
	default:
		return nil, fmt.Errorf("unknown data type %v", f.Type())
	}
}
