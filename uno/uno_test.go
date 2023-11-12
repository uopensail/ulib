package uno

import (
	"fmt"
	"testing"
)

func TestIn(t *testing.T) {
	{
		condition := `cat1[int64] in (12,1,3)`
		instance, err := NewEvaluator(condition)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		slice := instance.Allocate()

		err = instance.FillInt64("", "cat1", 1, slice)
		if instance.Eval(slice) != 1 {
			t.FailNow()
		}

		slice = instance.Allocate()
		err = instance.FillInt64("", "cat1", 2, slice)
		if instance.Eval(slice) != 0 {
			t.FailNow()
		}
	}
	{
		condition := `d_s_language[string] IN ("en","fr","jp")`
		instance, err := NewEvaluator(condition)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		slice := instance.Allocate()

		err = instance.FillString("", "d_s_language", "xx", slice)
		if instance.Eval(slice) != 0 {
			t.FailNow()
		}

		slice = instance.Allocate()
		err = instance.FillString("", "d_s_language", "fr", slice)
		if instance.Eval(slice) != 1 {
			t.FailNow()
		}
	}

}

func TestEval(t *testing.T) {
	condition := `v[float32] > 1000.0+1.5+2.9`
	instance, err := NewEvaluator(condition)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	slice := instance.Allocate()
	defer instance.Clean(slice)
	fmt.Printf("%v\n", slice)
	err = instance.FillFloat32("", "v", float32(1001), slice)
	fmt.Printf("%v\n", err)

	fmt.Printf("%v\n", slice)
	if instance.Eval(slice) != 0 {
		t.FailNow()
	}
	slice = instance.Allocate()
	defer instance.Clean(slice)
	err = instance.FillFloat32("", "v", float32(2001), slice)

	if instance.Eval(slice) != 1 {
		t.FailNow()
	}
}
