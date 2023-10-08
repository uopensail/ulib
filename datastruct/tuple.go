package datastruct

import "fmt"

type Tuple[T1 int | int32 | int64 | uint64 | float32 | float64 | string,
	T2 int | int32 | int64 | uint64 | float32 | float64 | string] struct {
	First  T1
	Second T2
}

type TupleList[T1 int | int32 | int64 | uint64 | float32 | float64 | string,
	T2 int | int32 | int64 | uint64 | float32 | float64 | string] []Tuple[T1, T2]

func (t TupleList[T1, T2]) Less(i, j int) bool {
	return t[i].Second < t[j].Second
}

func (t TupleList[T1, T2]) Len() int {
	return len(t)
}

func (t TupleList[T1, T2]) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t TupleList[T1, T2]) Print() {
	fmt.Print("[")
	for i := 0; i < len(t); i++ {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print("(", t[i].First, ",", t[i].Second, ")")
	}
	fmt.Println(")")
}
