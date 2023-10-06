package uno

import (
	"fmt"
	"testing"
)

func TestEval(t *testing.T) {
	condition := `a[int64] > (5+100)+100`
	instance, err := NewEvaluator(condition)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	slice := instance.Allocate()
	fmt.Printf("%v\n", slice)
	err = instance.FillInt64("a", int64(1001), slice)
	fmt.Printf("%v\n", err)

	fmt.Printf("%v\n", slice)
	fmt.Println(instance.Eval(slice))
}
