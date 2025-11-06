package sample

import (
	"errors"
)

var (
	ErrNotImplemented = errors.New("method not implemented for this type")
	ErrKeyNotFound    = errors.New("key not found")
	ErrTypeMismatch   = errors.New("type mismatch")
	ErrInvalidData    = errors.New("invalid data")
)

/**
 * @brief DataType represents the type of data stored in an ImmutableFeature
 *
 * Uses uint64 to ensure 8-byte alignment for optimal memory access.
 * Each stored value begins with a DataType field to enable type checking.
 */
type DataType uint64

const (
	Int64Type    DataType = iota // 64-bit signed integer
	Float32Type                  // 32-bit floating point
	StringType                   // UTF-8 string
	Int64sType                   // Slice of 64-bit signed integers
	Float32sType                 // Slice of 32-bit floating points
	StringsType                  // Slice of UTF-8 strings
	InvalidType  DataType = 127  // Invalid or uninitialized type
)

/**
 * @brief String returns the string representation of DataType
 *
 * @return String name of the data type
 */
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

/**
 * @brief IsValid checks if the DataType is a valid type
 *
 * @return True if the type is valid, false otherwise
 */
func (dt DataType) IsValid() bool {
	return dt >= Int64Type && dt <= StringsType
}

/**
 * @brief IsSliceType checks if the DataType represents a slice type
 *
 * @return True if the type is a slice type, false otherwise
 */
func (dt DataType) IsSliceType() bool {
	return dt == Int64sType || dt == Float32sType || dt == StringsType
}

/**
 * @brief Feature interface defines methods for accessing typed data values
 *
 * All feature implementations must provide type information and typed getter methods.
 * Type mismatches return errors to ensure type safety at runtime.
 *
 * The interface follows a consistent pattern where each getter method returns
 * the requested type and an error. Only the method corresponding to the feature's
 * actual type will succeed; all others return ErrNotImplemented.
 *
 * Safety Guidelines:
 * - Use safe methods (e.g., GetInt64) for type checking and error handling
 * - Use unsafe methods (e.g., GetInt64Unsafe) only when type is already verified
 * - Unsafe methods may panic or return garbage data if type doesn't match
 */
type Feature interface {
	/**
	 * @brief Returns the data type of this feature
	 *
	 * @return DataType enum value indicating the stored data type
	 */
	Type() DataType

	/**
	 * @brief Retrieves the stored value as int64
	 *
	 * @return The int64 value and nil error if type matches, otherwise ErrNotImplemented
	 */
	GetInt64() (int64, error)

	/**
	 * @brief Retrieves the stored value as int64 without type checking
	 *
	 * Warning: This method does not perform type checking. Use only when type is verified.
	 * @return The int64 value (undefined behavior if type doesn't match)
	 */
	GetInt64Unsafe() int64

	/**
	 * @brief Retrieves the stored value as float32
	 *
	 * @return The float32 value and nil error if type matches, otherwise ErrNotImplemented
	 */
	GetFloat32() (float32, error)

	/**
	 * @brief Retrieves the stored value as float32 without type checking
	 *
	 * Warning: This method does not perform type checking. Use only when type is verified.
	 * @return The float32 value (undefined behavior if type doesn't match)
	 */
	GetFloat32Unsafe() float32

	/**
	 * @brief Retrieves the stored value as string
	 *
	 * @return The string value and nil error if type matches, otherwise ErrNotImplemented
	 */
	GetString() (string, error)

	/**
	 * @brief Retrieves the stored value as string without type checking
	 *
	 * Warning: This method does not perform type checking. Use only when type is verified.
	 * @return The string value (undefined behavior if type doesn't match)
	 */
	GetStringUnsafe() string

	/**
	 * @brief Retrieves the stored value as int64 slice
	 *
	 * @return The int64 slice and nil error if type matches, otherwise ErrNotImplemented
	 */
	GetInt64s() ([]int64, error)

	/**
	 * @brief Retrieves the stored value as int64 slice without type checking
	 *
	 * Warning: This method does not perform type checking. Use only when type is verified.
	 * @return The int64 slice (undefined behavior if type doesn't match)
	 */
	GetInt64sUnsafe() []int64

	/**
	 * @brief Retrieves the stored value as float32 slice
	 *
	 * @return The float32 slice and nil error if type matches, otherwise ErrNotImplemented
	 */
	GetFloat32s() ([]float32, error)

	/**
	 * @brief Retrieves the stored value as float32 slice without type checking
	 *
	 * Warning: This method does not perform type checking. Use only when type is verified.
	 * @return The float32 slice (undefined behavior if type doesn't match)
	 */
	GetFloat32sUnsafe() []float32

	/**
	 * @brief Retrieves the stored value as string slice
	 *
	 * @return The string slice and nil error if type matches, otherwise ErrNotImplemented
	 */
	GetStrings() ([]string, error)

	/**
	 * @brief Retrieves the stored value as string slice without type checking
	 *
	 * Warning: This method does not perform type checking. Use only when type is verified.
	 * @return The string slice (undefined behavior if type doesn't match)
	 */
	GetStringsUnsafe() []string

	/**
	 * @brief Retrieves the stored value as any type
	 *
	 * This method returns the underlying value in its native Go type.
	 * It's useful for generic processing or when the exact type is not known.
	 *
	 * @return The underlying value as any type
	 */
	Get() any
}

/**
 * @brief Features interface defines methods for managing collections of typed features
 *
 * This interface is implemented by both MutableFeatures and ImmutableFeatures,
 * providing a consistent API for feature collection operations. It supports
 * basic collection operations, serialization, and type-safe feature access.
 *
 * Key characteristics:
 * - Uniform API across mutable and immutable implementations
 * - Type-safe feature retrieval
 * - JSON serialization/deserialization support
 * - Efficient key-based access patterns
 * - Thread-safety depends on implementation (ImmutableFeatures is thread-safe)
 */
type Features interface {

	/**
	 * @brief Gets the data type of a feature by key
	 *
	 * This is a convenience method that combines Get() and Type() calls.
	 *
	 * @param key Feature name
	 * @return DataType of the feature, InvalidType if not found
	 */
	GetType(key string) DataType
	/**
	 * @brief Returns all feature keys in the collection
	 *
	 * The returned slice contains all feature keys currently in the collection.
	 * The order of keys is not guaranteed and may vary between calls.
	 *
	 * @return Slice containing all feature keys (order not guaranteed)
	 */
	Keys() []string

	/**
	 * @brief Retrieves a feature by its key
	 *
	 * @param key The feature key to look up
	 * @return Feature interface if found, nil if key doesn't exist
	 */
	Get(key string) Feature

	/**
	 * @brief Returns the number of features in the collection
	 *
	 * @return Count of features currently stored
	 */
	Len() int

	/**
	 * @brief Checks if a feature key exists in the collection
	 *
	 * This method is more efficient than calling Get() and checking for nil
	 * when you only need to verify existence.
	 *
	 * @param key The feature key to check
	 * @return True if key exists, false otherwise
	 */
	Has(key string) bool

	/**
	 * @brief Converts features to a map for serialization or inspection
	 *
	 * Each feature is converted to a structured format containing both
	 * type information and the actual value. This method is useful for
	 * debugging, logging, or custom serialization needs.
	 *
	 * The returned map structure for each feature:
	 * {
	 *   "type": DataType,
	 *   "value": <actual_value>
	 * }
	 *
	 * @return Map of feature names to structured values, error if conversion fails
	 */
	MapAny() (map[string]any, error)

	/**
	 * @brief Marshals the feature collection to JSON format
	 *
	 * The JSON format includes both type and value information for each feature,
	 * enabling accurate reconstruction during unmarshaling. The format is
	 * consistent between mutable and immutable implementations.
	 *
	 * Example JSON structure:
	 * {
	 *   "feature1": {"type": 0, "value": 123},
	 *   "feature2": {"type": 2, "value": "hello"}
	 * }
	 *
	 * @return JSON byte array and error if marshaling fails
	 */
	MarshalJSON() ([]byte, error)

	/**
	 * @brief Unmarshals feature collection from JSON format
	 *
	 * Reconstructs the feature collection from JSON data created by MarshalJSON.
	 * The implementation handles type validation and creates appropriate
	 * feature instances based on the type information in the JSON.
	 *
	 * Note: This method may modify the existing collection. For mutable features,
	 * existing features may be overwritten. For immutable features, this typically
	 * replaces the entire collection.
	 *
	 * @param data JSON byte array to unmarshal
	 * @return Error if unmarshaling fails or data is invalid
	 */
	UnmarshalJSON(data []byte) error

	/**
	 * @brief Iterates over all features with a callback function
	 *
	 * The callback function is called for each feature in the collection.
	 * If the callback returns an error, iteration stops and the error is returned.
	 *
	 * @param fn Callback function called for each feature
	 * @return Error if callback returns error, nil if iteration completes
	 */
	ForEach(fn IteratorFunc) error
}

/**
 * @brief IteratorFunc defines the function signature for feature iteration
 *
 * @param key The feature key
 * @param feature The feature instance
 * @return Error to stop iteration, nil to continue
 */
type IteratorFunc func(key string, feature Feature) error
