package datastruct

import "unsafe"

// CombineSliceBuilder appends slices of various element types into per-type
// contiguous backing arrays and returns sub-slices that share that backing
// memory. This avoids separate heap allocations for each small slice.
//
// Strings are stored by copying their bytes into the byte backing array and
// constructing a string header that points into that array, yielding
// zero-allocation string storage.
//
// NOT thread-safe — external synchronization is required for concurrent use.
type CombineSliceBuilder struct {
	combineInt32   []int32
	combineInt64   []int64
	combineString  []string
	combineBytes   []byte
	combineFloat32 []float32
}

// CombineInt64s appends vs to the int64 backing array and returns the
// sub-slice that corresponds to vs within that array.
func (cs *CombineSliceBuilder) CombineInt64s(vs []int64) []int64 {
	cur := len(cs.combineInt64)
	cs.combineInt64 = append(cs.combineInt64, vs...)
	return cs.combineInt64[cur:]
}

// CombineFloat32s appends vs to the float32 backing array and returns the
// corresponding sub-slice.
func (cs *CombineSliceBuilder) CombineFloat32s(vs []float32) []float32 {
	cur := len(cs.combineFloat32)
	cs.combineFloat32 = append(cs.combineFloat32, vs...)
	return cs.combineFloat32[cur:]
}

// CombineBytes appends vs to the byte backing array and returns the
// corresponding sub-slice.
func (cs *CombineSliceBuilder) CombineBytes(vs []byte) []byte {
	cur := len(cs.combineBytes)
	cs.combineBytes = append(cs.combineBytes, vs...)
	return cs.combineBytes[cur:]
}

// CombineInt32s appends vs to the int32 backing array and returns the
// corresponding sub-slice.
func (cs *CombineSliceBuilder) CombineInt32s(vs []int32) []int32 {
	cur := len(cs.combineInt32)
	cs.combineInt32 = append(cs.combineInt32, vs...)
	return cs.combineInt32[cur:]
}

// CombineStrings appends vs to the string backing array and returns the
// corresponding sub-slice.
func (cs *CombineSliceBuilder) CombineStrings(vs []string) []string {
	cur := len(cs.combineString)
	cs.combineString = append(cs.combineString, vs...)
	return cs.combineString[cur:]
}

// CombineStringByte copies vs into the byte backing array and returns a
// string whose data pointer points directly into that array — no heap
// allocation beyond the backing array growth itself.
//
// The returned string is valid only as long as the CombineSliceBuilder is
// alive and has not been garbage-collected. Do NOT use the result after the
// builder is discarded.
func (cs *CombineSliceBuilder) CombineStringByte(vs []byte) string {
	cur := len(cs.combineBytes)
	cs.combineBytes = append(cs.combineBytes, vs...)
	bs := cs.combineBytes[cur:]
	// unsafe.Pointer cast avoids copying the bytes a second time.
	return *(*string)(unsafe.Pointer(&bs))
}

// CombineStringBytes converts each element of vvs into a string backed by
// the shared byte array (see CombineStringByte) and returns the resulting
// string slice.
func (cs *CombineSliceBuilder) CombineStringBytes(vvs [][]byte) []string {
	ret := make([]string, len(vvs))
	for i, v := range vvs {
		ret[i] = cs.CombineStringByte(v)
	}
	return ret
}
