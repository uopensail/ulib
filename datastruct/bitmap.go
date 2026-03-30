// Package datastruct provides generic data-structure primitives.
package datastruct

// BitMapStatus is the three-valued result of a DoubleBitMap check.
type BitMapStatus int8

const (
	// NoMarkBitMapStatus means the index has never been marked.
	NoMarkBitMapStatus BitMapStatus = 0
	// TrueBitMapStatus means the index was marked true.
	TrueBitMapStatus BitMapStatus = 1
	// FalseBitMapStatus means the index was marked false.
	FalseBitMapStatus BitMapStatus = 2
	// ErrorBitMapStatus is reserved for error states.
	ErrorBitMapStatus BitMapStatus = 3
)

// DoubleBitMap is a compact bitmap that stores two bits per logical index,
// allowing each position to carry one of four states (NoMark, True, False,
// Error). It is backed by a plain []byte and is NOT thread-safe.
//
// Storage layout: index i occupies bits (i%4)*2 and (i%4)*2+1 of byte i/4.
type DoubleBitMap []byte

// CreateDoubleBitMap allocates a DoubleBitMap large enough to hold size
// entries. All entries are initialised to NoMarkBitMapStatus (zero).
func CreateDoubleBitMap(size int) DoubleBitMap {
	// 2 bits per entry → 4 entries per byte; add one byte of headroom.
	return make(DoubleBitMap, (size>>2)+1)
}

// Mark sets the two-bit slot at index to true (bv=true → TrueBitMapStatus)
// or false (bv=false → FalseBitMapStatus). Calling Mark twice on the same
// index ORs the bits together, which may produce ErrorBitMapStatus; reset
// the bitmap before reuse if that is undesirable.
func (bitmap DoubleBitMap) Mark(index int, bv bool) {
	byteIndex := index >> 2
	offset := uint((index & 3) << 1)
	if bv {
		bitmap[byteIndex] |= 1 << offset
	} else {
		bitmap[byteIndex] |= 2 << offset
	}
}

// Check returns the BitMapStatus stored at index.
func (bitmap DoubleBitMap) Check(index int) BitMapStatus {
	byteIndex := index >> 2
	offset := uint((index & 3) << 1)
	return BitMapStatus((bitmap[byteIndex] >> offset) & 3)
}

// BitMap is a compact single-bit-per-index bitmap backed by a []byte.
// It is NOT thread-safe.
type BitMap []byte

// CreateBitMap allocates a BitMap large enough to hold size entries.
// All entries are initialised to false (zero).
func CreateBitMap(size int) BitMap {
	return make(BitMap, (size>>3)+1)
}

// MarkTrue sets the bit at index to 1.
func (bitmap BitMap) MarkTrue(index int) {
	bitmap[index>>3] |= 1 << uint(index&7)
}

// Check returns true if the bit at index is set.
func (bitmap BitMap) Check(index int) bool {
	return bitmap[index>>3]&(1<<uint(index&7)) != 0
}

// Clear resets all bits to zero using an exponential-copy trick that avoids
// a function call per byte when the slice is large.
func (bitmap BitMap) Clear() {
	if len(bitmap) == 0 {
		return
	}
	bitmap[0] = 0
	for bp := 1; bp < len(bitmap); bp <<= 1 {
		copy(bitmap[bp:], bitmap[:bp])
	}
}

// And performs a bitwise AND of bitmap with b, storing the result in bitmap.
// Only the bytes present in both slices are modified; surplus bytes in bitmap
// are left unchanged.
func (bitmap BitMap) And(b BitMap) {
	n := len(b)
	if len(bitmap) < n {
		n = len(bitmap)
	}
	for i := 0; i < n; i++ {
		bitmap[i] &= b[i]
	}
}

// Or performs a bitwise OR of bitmap with b, storing the result in bitmap.
func (bitmap BitMap) Or(b BitMap) {
	n := len(b)
	if len(bitmap) < n {
		n = len(bitmap)
	}
	for i := 0; i < n; i++ {
		bitmap[i] |= b[i]
	}
}
