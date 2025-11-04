package sample

/**
 * @brief Feature interface defines methods for accessing typed data values
 *
 * All feature implementations must provide type information and typed getter methods.
 * Type mismatches return errors to ensure type safety at runtime.
 *
 * The interface follows a consistent pattern where each getter method returns
 * the requested type and an error. Only the method corresponding to the feature's
 * actual type will succeed; all others return ErrNotImplemented.
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
	 * @return The int64 value and nil error if type matches, otherwise error
	 */
	GetInt64() (int64, error)

	/**
	 * @brief Retrieves the stored value as float32
	 *
	 * @return The float32 value and nil error if type matches, otherwise error
	 */
	GetFloat32() (float32, error)

	/**
	 * @brief Retrieves the stored value as string
	 *
	 * @return The string value and nil error if type matches, otherwise error
	 */
	GetString() (string, error)

	/**
	 * @brief Retrieves the stored value as int64 slice
	 *
	 * @return The int64 slice and nil error if type matches, otherwise error
	 */
	GetInt64s() ([]int64, error)

	/**
	 * @brief Retrieves the stored value as float32 slice
	 *
	 * @return The float32 slice and nil error if type matches, otherwise error
	 */
	GetFloat32s() ([]float32, error)

	/**
	 * @brief Retrieves the stored value as string slice
	 *
	 * @return The string slice and nil error if type matches, otherwise error
	 */
	GetStrings() ([]string, error)

	/**
	 * @brief Retrieves the stored value as any type
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
 */
type Features interface {
	/**
	 * @brief Returns all feature keys in the collection
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
	Get(string) Feature

	/**
	 * @brief Returns the number of features in the collection
	 *
	 * @return Count of features currently stored
	 */
	Len() int

	/**
	 * @brief Checks if a feature key exists in the collection
	 *
	 * @param key The feature key to check
	 * @return True if key exists, false otherwise
	 */
	Has(string) bool

	/**
	 * @brief Converts features to a map for serialization or inspection
	 *
	 * Each feature is converted to a structured format containing both
	 * type information and the actual value. This method is useful for
	 * debugging, logging, or custom serialization needs.
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
	 * @param data JSON byte array to unmarshal
	 * @return Error if unmarshaling fails or data is invalid
	 */
	UnmarshalJSON([]byte) error
}
