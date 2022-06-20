package luban

import (
	"github.com/uopensail/ulib/utils"
)

func GetFloat32ByKey(features *Features, key string) float32 {
	feature, ok := features.Feature[key]
	if !ok {
		return 0.0
	}
	return feature.GetFloat().GetValue()
}

func GetStringByKey(features *Features, key string) string {
	feature, ok := features.Feature[key]
	if !ok {
		return ""
	}
	values := feature.GetBytes().GetValue()
	if values == nil {
		return ""
	}
	return string(values)
}

func GetInt64ByKey(features *Features, key string) int64 {
	feature, ok := features.Feature[key]
	if !ok {
		return 0
	}
	return feature.GetInt64().GetValue()
}

func GetUint64ByKey(features *Features, key string) uint64 {
	feature, ok := features.Feature[key]
	if !ok {
		return 0
	}
	return feature.GetUint64().GetValue()
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

func GetUint64ListByKey(features *Features, key string) []uint64 {
	feature, ok := features.Feature[key]
	if !ok {
		return nil
	}
	return feature.GetUint64List().GetValue()
}

func MakeInt64List(vs []int64) *Feature {
	return &Feature{Kind: &Feature_Int64List{Int64List: &Int64List{Value: vs}}}
}

func MakeInt64(v int64) *Feature {
	return &Feature{Kind: &Feature_Int64{Int64: &Int64{Value: v}}}
}

func MakeUint64List(vs []uint64) *Feature {
	return &Feature{Kind: &Feature_Uint64List{Uint64List: &Uint64List{Value: vs}}}
}

func MakeUint64(v uint64) *Feature {
	return &Feature{Kind: &Feature_Uint64{Uint64: &Uint64{Value: v}}}
}

func MakeStringList(vs []string) *Feature {
	byteVs := make([][]byte, len(vs))
	for i := 0; i < len(vs); i++ {
		byteVs[i] = []byte(vs[i])
	}
	return &Feature{Kind: &Feature_BytesList{BytesList: &BytesList{Value: byteVs}}}
}

func MakeString(v string) *Feature {
	return &Feature{Kind: &Feature_Bytes{Bytes: &Bytes{Value: []byte(v)}}}
}

func MakeFloat32List(vs []float32) *Feature {
	return &Feature{Kind: &Feature_FloatList{FloatList: &FloatList{Value: vs}}}
}

func MakeFloat32(v float32) *Feature {
	return &Feature{Kind: &Feature_Float{Float: &Float{Value: v}}}
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

func MakeUint64ListWithStrs(ss []string) *Feature {
	vs := make([]uint64, len(ss))
	for i := 0; i < len(ss); i++ {
		vs[i] = utils.String2UInt64(ss[i])
	}
	return &Feature{Kind: &Feature_Uint64List{Uint64List: &Uint64List{Value: vs}}}
}
