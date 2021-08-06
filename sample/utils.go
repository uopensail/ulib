package sample

import (
	"github.com/uopensail/ulib/utils"
)

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

func MakeInt64List(vs []int64) *Feature {
	return &Feature{Kind: &Feature_Int64List{Int64List: &Int64List{Value: vs}}}
}

func MakeStringList(vs []string) *Feature {
	byteVs := make([][]byte, len(vs))
	for i := 0; i < len(vs); i++ {
		byteVs[i] = []byte(vs[i])
	}
	return &Feature{Kind: &Feature_BytesList{BytesList: &BytesList{Value: byteVs}}}
}

func MakeFloat32List(vs []float32) *Feature {
	return &Feature{Kind: &Feature_FloatList{FloatList: &FloatList{Value: vs}}}
}

func MakeFloat32ListWithStrs(ss []string) *Feature {
	vs := make([]float32, len(ss))
	for i := 0; i < len(ss); i++ {
		vs[i] = utils.String2Float32(ss[i])
	}
	return &Feature{Kind: &Feature_FloatList{FloatList: &FloatList{Value: vs}}}
}

func MakeByteListWithStrs(ss []string) *Feature {
	vs := make([][]byte, len(ss))
	for i := 0; i < len(ss); i++ {
		vs[i] = []byte(ss[i])
	}
	return &Feature{Kind: &Feature_BytesList{BytesList: &BytesList{Value: vs}}}
}

func MakeInt64ListWithStrs(ss []string) *Feature {
	vs := make([]int64, len(ss))
	for i := 0; i < len(ss); i++ {
		vs[i] = utils.String2Int64(ss[i])
	}
	return &Feature{Kind: &Feature_Int64List{Int64List: &Int64List{Value: vs}}}
}
