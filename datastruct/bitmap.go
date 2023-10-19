package datastruct

type BitMapStatus int8

const (
	NoMarkBitMapStatus BitMapStatus = 0
	TrueBitMapStatus   BitMapStatus = 1
	FalseBitMapStatus  BitMapStatus = 2
	ErrorBitMapStatus  BitMapStatus = 3
)

type DoubleBitMap []byte

func CreateDoubleBitMap(size int) DoubleBitMap {
	bitmap := make(DoubleBitMap, ((size>>3)+1)<<1)
	return bitmap
}

func (bitmap DoubleBitMap) Mark(index int, bv bool) {
	byteIndex := index >> 2
	offset := (index & 3) << 1
	if bv {
		bitmap[byteIndex] |= (1 << offset)
	} else {
		bitmap[byteIndex] |= (2 << offset)
	}
}

func (bitmap DoubleBitMap) Check(index int) BitMapStatus {
	byteIndex := index >> 2
	offset := (index & 3) << 1
	return BitMapStatus((bitmap[byteIndex] >> offset) & 3)
}

type BitMap []byte

func CreateBitMap(size int) BitMap {
	bitmap := make(BitMap, (size>>3)+1)
	return bitmap
}

func (bitmap BitMap) MarkTrue(index int) {
	byteIndex := index >> 3
	offset := index & 7
	bitmap[byteIndex] |= (1 << offset)
}

func (bitmap BitMap) Check(index int) bool {
	byteIndex := index >> 3
	offset := index & 7
	return !(bitmap[byteIndex]&(1<<offset) == 0)
}

func (bitmap BitMap) Clear() {
	if len(bitmap) > 0 {
		bitmap[0] = 0
		for bp := 1; bp < len(bitmap); bp *= 2 {
			copy(bitmap[bp:], bitmap[:bp])
		}
	}
}
