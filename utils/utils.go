package utils

func GetFloat32Param(param map[string]string, key string, defaultVal float32) float32 {
	if v, ok := param[key]; ok {
		return String2Float32(v)
	}
	return defaultVal
}

func GetFloat64Param(param map[string]string, key string, defaultVal float64) float64 {
	if v, ok := param[key]; ok {
		return String2Float64(v)
	}
	return defaultVal
}

func GetIntParam(param map[string]string, key string, defaultVal int) int {
	if v, ok := param[key]; ok {
		return String2Int(v)
	}
	return defaultVal
}

func GetInt32Param(param map[string]string, key string, defaultVal int32) int32 {
	if v, ok := param[key]; ok {
		return String2Int32(v)
	}
	return defaultVal
}

func GetUInt32Param(param map[string]string, key string, defaultVal uint32) uint32 {
	if v, ok := param[key]; ok {
		return String2UInt32(v)
	}
	return defaultVal
}

func GetInt8Param(param map[string]string, key string, defaultVal int8) int8 {
	if v, ok := param[key]; ok {
		return String2Int8(v)
	}
	return defaultVal
}

func GetInt64Param(param map[string]string, key string, defaultVal int64) int64 {
	if v, ok := param[key]; ok {
		return String2Int64(v)
	}
	return defaultVal
}

func GetUInt64Param(param map[string]string, key string, defaultVal uint64) uint64 {
	if v, ok := param[key]; ok {
		return String2UInt64(v)
	}
	return defaultVal
}
