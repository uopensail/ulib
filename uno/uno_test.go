package uno

import (
	"fmt"
	"testing"
)

func TestEval(t *testing.T) {
	condition := `v[float32] > 1000.0+1.5+2.9`
	instance, err := NewEvaluator(condition)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	slice := instance.Allocate()
	fmt.Printf("%v\n", slice)
	err = instance.FillFloat32("v", float32(1001), slice)
	fmt.Printf("%v\n", err)

	fmt.Printf("%v\n", slice)
	fmt.Println(instance.Eval(slice))
}
