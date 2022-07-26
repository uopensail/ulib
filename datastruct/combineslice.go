package datastruct

import "unsafe"

//Thread unsafe
type CombineSliceBuilder struct {
	combineInt32  []int32
	combineInt64  []int64
	combineString []string
	combineBytes  []byte
	combineFloat  []float32
}

func (cs *CombineSliceBuilder) CombineInt64s(vs []int64) []int64 {
	cur := len(cs.combineInt64)
	cs.combineInt64 = append(cs.combineInt64, vs...)
	return cs.combineInt64[cur:]
}
func (cs *CombineSliceBuilder) CombineBytes(vs []byte) []byte {
	cur := len(cs.combineBytes)
	cs.combineBytes = append(cs.combineBytes, vs...)
	return cs.combineBytes[cur:]
}

func (cs *CombineSliceBuilder) CombineInt32s(vs []int32) []int32 {
	cur := len(cs.combineInt32)
	cs.combineInt32 = append(cs.combineInt32, vs...)
	return cs.combineInt32[cur:]
}

func (cs *CombineSliceBuilder) CombineStrings(vs []string) []string {
	cur := len(cs.combineString)
	cs.combineString = append(cs.combineString, vs...)
	return cs.combineString[cur:]
}
func (cs *CombineSliceBuilder) CombineStringByte(vs []byte) string {
	cur := len(cs.combineBytes)
	cs.combineBytes = append(cs.combineBytes, vs...)
	bs := cs.combineBytes[cur:]
	return *(*string)(unsafe.Pointer(&bs))
}

func (cs *CombineSliceBuilder) CombineStringBytes(vvs [][]byte) []string {
	ret := make([]string, len(vvs))
	for i := 0; i < len(vvs); i++ {
		ret[i] = cs.CombineStringByte(vvs[i])
	}
	return ret
}
