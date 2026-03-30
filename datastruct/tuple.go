package datastruct

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

// Tuple is a generic ordered pair. Both fields are exported so the value can
// be used directly without getters.
type Tuple[T1 any, T2 any] struct {
	First  T1
	Second T2
}

// MarshalJSON serialises the tuple as {"first":…,"second":…}.
// It reuses the same memory layout as an anonymous struct with JSON tags to
// avoid copying the field values.
func (t *Tuple[T1, T2]) MarshalJSON() ([]byte, error) {
	type wire struct {
		First  T1 `json:"first"`
		Second T2 `json:"second"`
	}
	return json.Marshal((*wire)(unsafe.Pointer(t)))
}

// UnmarshalJSON deserialises {"first":…,"second":…} into the tuple.
func (t *Tuple[T1, T2]) UnmarshalJSON(data []byte) error {
	type wire struct {
		First  T1 `json:"first"`
		Second T2 `json:"second"`
	}
	return json.Unmarshal(data, (*wire)(unsafe.Pointer(t)))
}

// String returns a human-readable representation of the tuple.
func (t *Tuple[T1, T2]) String() string {
	return fmt.Sprintf("(%v, %v)", t.First, t.Second)
}

// Print writes the tuple to stdout in (first, second) format.
func (t *Tuple[T1, T2]) Print() {
	fmt.Println(t.String())
}

// Ordered is a type constraint satisfied by all integer, float, and string
// types, making them directly comparable with <, >, <=, >=.
type Ordered interface {
	Integer | Float | ~string
}

// Integer is a type constraint for all signed and unsigned integer types.
type Integer interface {
	Signed | Unsigned
}

// Signed is a type constraint for all signed integer types.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a type constraint for all unsigned integer types.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Float is a type constraint for 32-bit and 64-bit floating-point types.
type Float interface {
	~float32 | ~float64
}
