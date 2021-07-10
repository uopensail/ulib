package sample

func GetFloat32ByKey(features *Features, key string) float32 {
	feature, ok := features.Feature[key]
	if !ok {
		return 0.0
	}
	values := feature.GetFloatList().GetValue()
	if values == nil {
		return 0.0
	}
	return values[0]
}

func GetStringByKey(features *Features, key string) string {
	feature, ok := features.Feature[key]
	if !ok {
		return ""
	}
	values := feature.GetBytesList().GetValue()
	if values == nil {
		return ""
	}
	return string(values[0])
}

func GetInt64ByKey(features *Features, key string) int64 {
	feature, ok := features.Feature[key]
	if !ok {
		return 0
	}
	values := feature.GetInt64List().GetValue()
	if values == nil {
		return 0
	}
	return values[0]
}

func GetFloat32ListByKey(features *Features, key string) []float32 {
	feature, ok := features.Feature[key]
	if !ok {
		return nil
	}
	return feature.GetFloatList().GetValue()
}

func GetStringListByKey(features *Features, key string) []string {
	feature, ok := features.Feature[key]
	if !ok {
		return nil
	}
	values := feature.GetBytesList().GetValue()
	if values == nil {
		return nil
	}
	ret := make([]string, len(values))
	for i := 0; i < len(values); i++ {
		ret[i] = string(values[i])
	}
	return ret
}

func GetInt64ListByKey(features *Features, key string) []int64 {
	feature, ok := features.Feature[key]
	if !ok {
		return nil
	}
	return feature.GetInt64List().GetValue()
}
