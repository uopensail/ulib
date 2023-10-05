package uno

/*
#cgo CXXFLAGS: -std=c++17
#include <stdlib.h>
#include "uno.h"
*/
import "C"

import (
	"unsafe"
)

func eval(expression unsafe.Pointer, slice []unsafe.Pointer) int32 {
	return int32(C.uno_eval(expression, (*C.char)(unsafe.Pointer(&slice))))
}

func preEval(expression unsafe.Pointer, slice []unsafe.Pointer) {
	C.uno_preeval(expression, (*C.char)(unsafe.Pointer(&slice)))
}

func batchEval(expression unsafe.Pointer, slices [][]unsafe.Pointer) []int32 {
	ret := make([]int32, len(slices))
	C.uno_batch_eval(expression, (*C.char)(unsafe.Pointer(&slices)), (*C.char)(unsafe.Pointer(&ret)))
	return ret
}

func clean(expression unsafe.Pointer, slice []unsafe.Pointer) {
	C.uno_clean_varslice(expression, (*C.char)(unsafe.Pointer(&slice)))
}

func release(expression unsafe.Pointer) {
	C.uno_release_expression(unsafe.Pointer(expression))
}

func call(function string, args []unsafe.Pointer) {
	f := C.CString(function)
	defer C.free(unsafe.Pointer(f))
	C.uno_call_unsafe(unsafe.Pointer(f), C.int(len(function)), (*C.char)(unsafe.Pointer(&args)))
}

func callForInt64(function string, args []unsafe.Pointer) int64 {
	call(function, args)
	var ret int64
	ret = *(*int64)(args[len(args)-1])
	C.free(args[len(args)-1])
	args[len(args)-1] = nil
	return ret
}

func callForFloat32(function string, args []unsafe.Pointer) float32 {
	call(function, args)
	var ret float32
	ret = *(*float32)(args[len(args)-1])
	C.free(args[len(args)-1])
	args[len(args)-1] = nil
	return ret
}

func callForString(function string, args []unsafe.Pointer) string {
	call(function, args)
	s := *(*string)(args[len(args)-1])
	data := make([]byte, len(s))
	copy(data, *(*[]byte)(unsafe.Pointer(&s)))
	C.free(args[len(args)-1])
	args[len(args)-1] = nil
	return string(data)
}
