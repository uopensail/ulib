package uno

/*
#cgo CXXFLAGS: -std=c++17
#include <stdlib.h>
#include "uno.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type Column struct {
	Addr   int32
	Column string
	Type   DataType
}

type Columns struct {
	addrs []int32
	cols  []string
	types []DataType
}

type Expression struct {
	columns  Columns
	varSlice []unsafe.Pointer
}

type Evaluator struct {
	expression unsafe.Pointer
	columns    map[string]Column
}

func NewEvaluator(code string) *Evaluator {
	bytes := *(*[]byte)(unsafe.Pointer(&code))
	expr := C.uno_create_expression(unsafe.Pointer(&bytes[0]), C.int(len(code)))
	columns := make(map[string]Column)
	cols := (*Expression)(expr).columns
	for i := 0; i < len(cols.addrs); i++ {
		columns[cols.cols[i]] = Column{
			Addr:   cols.addrs[i],
			Column: cols.cols[i],
			Type:   cols.types[i],
		}
	}
	return &Evaluator{expression: expr, columns: columns}
}

func (e *Evaluator) FillInt64(col string, data int64, value []unsafe.Pointer) error {
	column, ok := e.columns[col]
	if !ok {
		return fmt.Errorf("column %s not found", col)
	}

	if column.Type == Int64 {
		value[column.Addr] = unsafe.Pointer(&data)
		return nil
	}
	return fmt.Errorf("column type check error: %d not found", column.Type)
}

func (e *Evaluator) FillInt64s(col string, data []int64, value []unsafe.Pointer) error {
	column, ok := e.columns[col]
	if !ok {
		return fmt.Errorf("column %s not found", col)
	}

	if column.Type == Int64s {
		value[column.Addr] = unsafe.Pointer(&data)
		return nil
	}
	return fmt.Errorf("column type check error: %d not found", column.Type)
}

func (e *Evaluator) FillFloat32(col string, data float32, value []unsafe.Pointer) error {
	column, ok := e.columns[col]
	if !ok {
		return fmt.Errorf("column %s not found", col)
	}

	if column.Type == Float32 {
		value[column.Addr] = unsafe.Pointer(&data)
		return nil
	}
	return fmt.Errorf("column type check error: %d not found", column.Type)
}

func (e *Evaluator) FillFloat32s(col string, data []float32, value []unsafe.Pointer) error {
	column, ok := e.columns[col]
	if !ok {
		return fmt.Errorf("column %s not found", col)
	}

	if column.Type == Float32s {
		value[column.Addr] = unsafe.Pointer(&data)
		return nil
	}
	return fmt.Errorf("column type check error: %d not found", column.Type)
}

func (e *Evaluator) FillString(col string, data string, value []unsafe.Pointer) error {
	column, ok := e.columns[col]
	if !ok {
		return fmt.Errorf("column %s not found", col)
	}

	if column.Type == String {
		value[column.Addr] = unsafe.Pointer(&data)
		return nil
	}
	return fmt.Errorf("column type check error: %d not found", column.Type)
}

func (e *Evaluator) FillStrings(col string, data []string, value []unsafe.Pointer) error {
	column, ok := e.columns[col]
	if !ok {
		return fmt.Errorf("column %s not found", col)
	}

	if column.Type == Strings {
		value[column.Addr] = unsafe.Pointer(&data)
		return nil
	}
	return fmt.Errorf("column type check error: %d not found", column.Type)
}

func (e *Evaluator) Allocate() []unsafe.Pointer {
	ret := make([]unsafe.Pointer, len((*Expression)(e.expression).varSlice))
	copy(ret, (*Expression)(e.expression).varSlice)
	return ret
}

func (e *Evaluator) Eval(slice []unsafe.Pointer) int32 {
	return int32(C.uno_eval(e.expression, (*C.char)(unsafe.Pointer(&slice))))
}

func (e *Evaluator) PreEval(slice []unsafe.Pointer) {
	C.uno_preeval(e.expression, (*C.char)(unsafe.Pointer(&slice)))
}

func (e *Evaluator) BatchEval(slices [][]unsafe.Pointer) []int32 {
	ret := make([]int32, len(slices))
	C.uno_batch_eval(e.expression, (*C.char)(unsafe.Pointer(&slices)), (*C.char)(unsafe.Pointer(&ret)))
	return ret
}

func (e *Evaluator) Clean(slice []unsafe.Pointer) {
	C.uno_clean_varslice(e.expression, (*C.char)(unsafe.Pointer(&slice)))
}

func (e *Evaluator) Release() {
	C.uno_release_expression(unsafe.Pointer(e.expression))
}
