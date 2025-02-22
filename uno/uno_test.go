package uno

import (
	"fmt"
	"testing"

	"github.com/uopensail/ulib/sample"
)

func TestIn(t *testing.T) {
	types := map[string]sample.DataType{
		"cat1":         sample.Int64Type,
		"d_s_language": sample.StringType,
	}
	{
		condition := `cat1 in (12,1,3)`
		instance, err := NewEvaluator(condition, types)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		slice := instance.Allocate()

		err = instance.FillInt64("cat1", 1, slice)
		if instance.Eval(slice) != 1 {
			t.FailNow()
		}

		slice = instance.Allocate()
		err = instance.FillInt64("cat1", 2, slice)
		if instance.Eval(slice) != 0 {
			t.FailNow()
		}
	}
	{
		condition := `d_s_language IN ("en","fr","jp")`
		instance, err := NewEvaluator(condition, types)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		slice := instance.Allocate()

		err = instance.FillString("d_s_language", "xx", slice)
		if instance.Eval(slice) != 0 {
			t.FailNow()
		}

		slice = instance.Allocate()
		err = instance.FillString("d_s_language", "fr", slice)
		if instance.Eval(slice) != 1 {
			t.FailNow()
		}
	}

}

func Test_Fill(t *testing.T) {
	types := map[string]sample.DataType{
		"d_s_language": sample.StringType,
		"u_s_language": sample.StringType,
	}
	{
		condition := `d_s_language = u_s_language`
		instance, err := NewEvaluator(condition, types)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		dFeat := sample.NewMutableFeatures()
		dFeat.Set("d_s_language", &sample.String{Value: "en"})

		slice := instance.Allocate()
		instance.Fill(dFeat, slice)
		uFeat := sample.NewMutableFeatures()
		uFeat.Set("u_s_language", &sample.String{Value: "fr"})
		instance.Fill(uFeat, slice)
		if instance.Eval(slice) != 0 {
			t.FailNow()
		}

		uFeat.Set("u_s_language", &sample.String{Value: "en"})
		instance.Fill(uFeat, slice)
		if instance.Eval(slice) != 1 {
			t.FailNow()
		}
	}
}

func TestEval(t *testing.T) {
	types := map[string]sample.DataType{
		"v": sample.Float32Type,
	}
	condition := `v > 1000.0+1.5+2.9`
	instance, err := NewEvaluator(condition, types)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	slice := instance.Allocate()

	fmt.Printf("%v\n", slice)
	err = instance.FillFloat32("v", float32(1001), slice)
	fmt.Printf("%v\n", err)

	fmt.Printf("%v\n", slice)
	if instance.Eval(slice) != 0 {
		t.FailNow()
	}
	slice = instance.Allocate()

	err = instance.FillFloat32("v", float32(2001), slice)

	if instance.Eval(slice) != 1 {
		t.FailNow()
	}
}
