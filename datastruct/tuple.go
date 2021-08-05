package datastruct

type TupleSS struct {
	First  string
	Second string
}

type TupleSF struct {
	First  string
	Second float64
}

type TupleIF struct {
	First  int
	Second float64
}

type TupleII struct {
	First  int
	Second int
}

type TupleSFList []TupleSF

func (sort TupleSFList) Less(i, j int) bool {
	return sort[i].Second < sort[j].Second
}
func (sort TupleSFList) Len() int {
	return len(sort)
}
func (sort TupleSFList) Swap(i, j int) {
	sort[i], sort[j] = sort[j], sort[i]
}

type TupleIFList []TupleIF

func (sort TupleIFList) Less(i, j int) bool {
	return sort[i].Second < sort[j].Second
}
func (sort TupleIFList) Len() int {
	return len(sort)
}
func (sort TupleIFList) Swap(i, j int) {
	sort[i], sort[j] = sort[j], sort[i]
}

type TupleIFListP []*TupleIF

func (sort TupleIFListP) Less(i, j int) bool {
	return sort[i].Second < sort[j].Second
}
func (sort TupleIFListP) Len() int {
	return len(sort)
}
func (sort TupleIFListP) Swap(i, j int) {
	sort[i], sort[j] = sort[j], sort[i]
}

type TupleIIListP []*TupleII
