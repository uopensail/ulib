package source

import (
	"unsafe"

	"github.com/uopensail/ulib/prome"
	"github.com/uopensail/ulib/sample"
	"github.com/uopensail/ulib/uno"
)

type Condition struct {
	evaluator *uno.Evaluator
	slices    [][]unsafe.Pointer
}

func (c *Condition) Check(features sample.Features, collection Collection) Collection {
	stat := prome.NewStat("Condition.Check")
	defer stat.End()
	var slice []unsafe.Pointer
	var oldSlice []unsafe.Pointer
	var newSlice []unsafe.Pointer

	slices := make([][]unsafe.Pointer, 0, len(collection))
	slice = c.evaluator.Allocate()
	c.evaluator.Fill(features, slice)
	address := make([]uintptr, len(slice))
	for i := 0; i < len(slice); i++ {
		address[i] = uintptr(slice[i])
	}

	for i := 0; i < len(collection); i++ {
		oldSlice = c.slices[collection[i]]
		newSlice = make([]unsafe.Pointer, len(oldSlice))
		for j := 0; j < len(slice); j++ {
			newSlice[j] = unsafe.Pointer(uintptr(oldSlice[j]) | address[j])
		}
		slices = append(slices, newSlice)
	}

	ret := make([]int, 0, len(collection))
	results := c.evaluator.BatchEval(slices)
	for i := 0; i < len(results); i++ {
		if results[i] == 1 {
			ret = append(ret, collection[i])
		}
	}
	stat.SetCounter(len(ret))
	return ret
}

func (c *Condition) CheckAll(features sample.Features) Collection {
	stat := prome.NewStat("Condition.CheckAll")
	defer stat.End()
	var slice []unsafe.Pointer
	var newSlice []unsafe.Pointer

	slices := make([][]unsafe.Pointer, 0, len(c.slices))
	slice = c.evaluator.Allocate()
	address := make([]uintptr, len(slice))
	for i := 0; i < len(slice); i++ {
		address[i] = uintptr(slice[i])
	}
	c.evaluator.Fill(features, slice)

	for i := 0; i < len(c.slices); i++ {
		newSlice = make([]unsafe.Pointer, len(c.slices[i]))
		for j := 0; j < len(slice); j++ {
			newSlice[j] = unsafe.Pointer(uintptr(c.slices[i][j]) | address[j])
		}
		slices = append(slices, newSlice)
	}

	ret := make([]int, 0, len(c.slices))
	results := c.evaluator.BatchEval(slices)
	for i := 0; i < len(results); i++ {
		if results[i] == 1 {
			ret = append(ret, i)
		}
	}
	stat.SetCounter(len(ret))
	return ret
}

func (c *Condition) Release() {
	for i := 0; i < len(c.slices); i++ {
		c.evaluator.Clean(c.slices[i])
	}
	c.evaluator.Release()
}
