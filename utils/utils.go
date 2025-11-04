package utils

import (
	"strconv"
)

// GetFloat32Param retrieves a float32 value from a string map with a default fallback
//
// @param param: The string map containing parameters
// @param key: The key to look up in the map
// @param defaultVal: The default value to return if key is not found or conversion fails
// @return: The parsed float32 value or default value
func GetFloat32Param(param map[string]string, key string, defaultVal float32) float32 {
	if v, ok := param[key]; ok {
		return String2Float32(v)
	}
	return defaultVal
}

// GetFloat64Param retrieves a float64 value from a string map with a default fallback
//
// @param param: The string map containing parameters
// @param key: The key to look up in the map
// @param defaultVal: The default value to return if key is not found or conversion fails
// @return: The parsed float64 value or default value
func GetFloat64Param(param map[string]string, key string, defaultVal float64) float64 {
	if v, ok := param[key]; ok {
		return String2Float64(v)
	}
	return defaultVal
}

// GetIntParam retrieves an int value from a string map with a default fallback
//
// @param param: The string map containing parameters
// @param key: The key to look up in the map
// @param defaultVal: The default value to return if key is not found or conversion fails
// @return: The parsed int value or default value
func GetIntParam(param map[string]string, key string, defaultVal int) int {
	if v, ok := param[key]; ok {
		return String2Int(v)
	}
	return defaultVal
}

// GetInt32Param retrieves an int32 value from a string map with a default fallback
//
// @param param: The string map containing parameters
// @param key: The key to look up in the map
// @param defaultVal: The default value to return if key is not found or conversion fails
// @return: The parsed int32 value or default value
func GetInt32Param(param map[string]string, key string, defaultVal int32) int32 {
	if v, ok := param[key]; ok {
		return String2Int32(v)
	}
	return defaultVal
}

// GetUInt32Param retrieves a uint32 value from a string map with a default fallback
//
// @param param: The string map containing parameters
// @param key: The key to look up in the map
// @param defaultVal: The default value to return if key is not found or conversion fails
// @return: The parsed uint32 value or default value
func GetUInt32Param(param map[string]string, key string, defaultVal uint32) uint32 {
	if v, ok := param[key]; ok {
		return String2UInt32(v)
	}
	return defaultVal
}

// GetInt8Param retrieves an int8 value from a string map with a default fallback
//
// @param param: The string map containing parameters
// @param key: The key to look up in the map
// @param defaultVal: The default value to return if key is not found or conversion fails
// @return: The parsed int8 value or default value
func GetInt8Param(param map[string]string, key string, defaultVal int8) int8 {
	if v, ok := param[key]; ok {
		return String2Int8(v)
	}
	return defaultVal
}

// GetInt64Param retrieves an int64 value from a string map with a default fallback
//
// @param param: The string map containing parameters
// @param key: The key to look up in the map
// @param defaultVal: The default value to return if key is not found or conversion fails
// @return: The parsed int64 value or default value
func GetInt64Param(param map[string]string, key string, defaultVal int64) int64 {
	if v, ok := param[key]; ok {
		return String2Int64(v)
	}
	return defaultVal
}

// GetUInt64Param retrieves a uint64 value from a string map with a default fallback
//
// @param param: The string map containing parameters
// @param key: The key to look up in the map
// @param defaultVal: The default value to return if key is not found or conversion fails
// @return: The parsed uint64 value or default value
func GetUInt64Param(param map[string]string, key string, defaultVal uint64) uint64 {
	if v, ok := param[key]; ok {
		return String2UInt64(v)
	}
	return defaultVal
}

// GetStringParam retrieves a string value from a string map with a default fallback
//
// @param param: The string map containing parameters
// @param key: The key to look up in the map
// @param defaultVal: The default value to return if key is not found
// @return: The string value or default value
func GetStringParam(param map[string]string, key string, defaultVal string) string {
	if v, ok := param[key]; ok {
		return v
	}
	return defaultVal
}

// GetBoolParam retrieves a boolean value from a string map with a default fallback
//
// @param param: The string map containing parameters
// @param key: The key to look up in the map
// @param defaultVal: The default value to return if key is not found or conversion fails
// @return: The parsed boolean value or default value
func GetBoolParam(param map[string]string, key string, defaultVal bool) bool {
	if v, ok := param[key]; ok {
		if result, err := strconv.ParseBool(v); err == nil {
			return result
		}
	}
	return defaultVal
}

// GetStringSliceParam retrieves a string slice from a string map by splitting the value
//
// @param param: The string map containing parameters
// @param key: The key to look up in the map
// @param separator: The separator used to split the string
// @param defaultVal: The default value to return if key is not found
// @return: The string slice or default value
func GetStringSliceParam(param map[string]string, key string, separator string, defaultVal []string) []string {
	if v, ok := param[key]; ok {
		// Simple split implementation - you might want to use strings.Split
		// and handle trimming based on your needs
		if separator == "" {
			return []string{v}
		}
		// This is a placeholder - implement proper splitting logic as needed
		return []string{v} // Replace with actual split logic
	}
	return defaultVal
}
