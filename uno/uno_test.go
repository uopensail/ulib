package uno

import (
	"fmt"
	"testing"
)

func TestEval(t *testing.T) {
	code := `{"nodes":[{"id":0,"ntype":0,"value":"a","dtype":1},{"id":1,"ntype":7,"args":[0],"func":"castf2i"},{"id":2,"ntype":0,"value":"b","dtype":0},{"id":3,"left":1,"right":2,"ntype":10,"dtype":0,"cmp":3}]}`
	instance := NewEvaluator(code)
	slice := instance.Allocate()
	fmt.Printf("%v\n", slice)
	err := instance.FillFloat32("a", float32(1001), slice)
	fmt.Printf("%v\n", err)
	err = instance.FillInt64("b", int64(1000), slice)
	fmt.Printf("%v\n", err)

	fmt.Printf("%v\n", slice)
	fmt.Println(instance.Eval(slice))
	//fmt.Printf("%d\n", *(*int64)(slice[0]))
}
