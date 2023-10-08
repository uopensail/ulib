package source

import (
	"fmt"

	"github.com/uopensail/ulib/sample"
)

type Collection []int

type wrappers struct {
	slice []*wrapper
	key   string
	desc  bool
}

func (w *wrappers) Less(i, j int) bool {
	left := w.slice[i].features.Get(w.key)
	right := w.slice[j].features.Get(w.key)

	if left == nil || right == nil {
		return false
	}

	dtype := left.Type()
	switch dtype {
	case sample.Float32Type:
		lv, err1 := left.GetFloat32()
		rv, err2 := right.GetFloat32()
		if err1 != nil || err2 != nil {
			return false
		}
		return !((lv < rv) && w.desc)
	case sample.Int64Type:
		lv, err1 := left.GetInt64()
		rv, err2 := right.GetInt64()
		if err1 != nil || err2 != nil {
			return false
		}
		return !((lv < rv) && w.desc)
	case sample.StringType:
		lv, err1 := left.GetString()
		rv, err2 := right.GetString()
		if err1 != nil || err2 != nil {
			return false
		}
		return !((lv < rv) && w.desc)
	default:
		panic(fmt.Sprintf("data type: %d not support", dtype))
	}

}

func (w *wrappers) Len() int {
	return len(w.slice)
}

func (w *wrappers) Swap(i, j int) {
	w.slice[i], w.slice[j] = w.slice[j], w.slice[i]
}
