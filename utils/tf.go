package utils

import "github.com/uopensail/ulib/sample"

func MakeFeatureFloat32List(vs []float32) *sample.Feature {

	return &sample.Feature{
		Kind: &sample.Feature_FloatList{
			FloatList: &sample.FloatList{Value: vs},
		},
	}
}

func MakeFeatureInt64List(vs []int64) *sample.Feature {
	return &sample.Feature{
		Kind: &sample.Feature_Int64List{
			Int64List: &sample.Int64List{Value: vs},
		},
	}
}

func MakeFloat32ListWithStrs(ss []string) *sample.Feature {
	vs := make([]float32, len(ss))
	for i := 0; i < len(ss); i++ {
		vs[i] = String2Float32(ss[i])
	}
	return &sample.Feature{
		Kind: &sample.Feature_FloatList{
			FloatList: &sample.FloatList{Value: vs},
		},
	}
}

func MakeByteListWithStrs(ss []string) *sample.Feature {
	vs := make([][]byte, len(ss))
	for i := 0; i < len(ss); i++ {
		vs[i] = []byte(ss[i])
	}
	return &sample.Feature{
		Kind: &sample.Feature_BytesList{
			BytesList: &sample.BytesList{Value: vs},
		},
	}
}

func MakeInt64ListWithStrs(ss []string) *sample.Feature {
	vs := make([]int64, len(ss))
	for i := 0; i < len(ss); i++ {
		vs[i] = String2Int64(ss[i])
	}
	return &sample.Feature{
		Kind: &sample.Feature_Int64List{
			Int64List: &sample.Int64List{Value: vs},
		},
	}
}
