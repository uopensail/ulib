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

func uno_create_expression(code string) unsafe.Pointer {
	bytes := *(*[]byte)(unsafe.Pointer(&code))
	return C.uno_create_expression(unsafe.Pointer(&bytes[0]), C.int(len(code)))
}

func uno_eval(expression unsafe.Pointer, slice []unsafe.Pointer) int32 {
	return int32(C.uno_eval(expression, (*C.char)(unsafe.Pointer(&slice))))
}

func uno_preeval(expression unsafe.Pointer, slice []unsafe.Pointer) {
	C.uno_preeval(expression, (*C.char)(unsafe.Pointer(&slice)))
}

func uno_batch_eval(expression unsafe.Pointer, slices [][]unsafe.Pointer) []int32 {
	ret := make([]int32, len(slices))
	C.uno_batch_eval(expression, (*C.char)(unsafe.Pointer(&slices)), (*C.char)(unsafe.Pointer(&ret)))
	return ret
}

func uno_clean_varslice(expression unsafe.Pointer, slice []unsafe.Pointer) {
	C.uno_clean_varslice(expression, (*C.char)(unsafe.Pointer(&slice)))
}

func uno_release_expression(expression unsafe.Pointer) {
	C.uno_release_expression(unsafe.Pointer(expression))
}

func call(function string, args []unsafe.Pointer) {
	bytes := *(*[]byte)(unsafe.Pointer(&function))
	C.uno_call_unsafe(unsafe.Pointer(&bytes[0]), C.int(len(function)), (*C.char)(unsafe.Pointer(&args)))
}

func call_for_int64(function string, args []unsafe.Pointer) int64 {
	call(function, args)
	ret := *(*int64)(args[len(args)-1])
	C.free(args[len(args)-1])
	args[len(args)-1] = nil
	return ret
}

func call_for_float32(function string, args []unsafe.Pointer) float32 {
	call(function, args)
	ret := *(*float32)(args[len(args)-1])
	C.free(args[len(args)-1])
	args[len(args)-1] = nil
	return ret
}

func call_for_string(function string, args []unsafe.Pointer) string {
	call(function, args)
	s := *(*string)(args[len(args)-1])
	data := make([]byte, len(s))
	copy(data, *(*[]byte)(unsafe.Pointer(&s)))
	C.free(args[len(args)-1])
	args[len(args)-1] = nil
	return string(data)
}
