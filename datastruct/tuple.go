package datastruct

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

type Tuple[T1 any, T2 any] struct {
	First  T1
	Second T2
}

func (t *Tuple[T1, T2]) MarshalJSON() ([]byte, error) {
	return json.Marshal((*(struct {
		First  T1 `json:"first"`
		Second T2 `json:"second"`
	}))(unsafe.Pointer(t)))
}

func (t *Tuple[T1, T2]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*(struct {
		First  T1 `json:"first"`
		Second T2 `json:"second"`
	}))(unsafe.Pointer(t)))
}

func (t *Tuple[T1, T2]) Print() {
	fmt.Printf("(%v,%v)\n", t.First, t.Second)
}

type Ordered interface {
	Integer | Float | ~string
}

type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Float interface {
	~float32 | ~float64
}
