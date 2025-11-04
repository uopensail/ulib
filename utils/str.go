package utils

import (
	"strconv"
	"strings"
)

// String2Float64 converts a string to float64, returns 0 if conversion fails
//
// @param s: The string to convert
// @return: The converted float64 value or 0 on error
func String2Float64(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

// String2Float32 converts a string to float32, returns 0 if conversion fails
//
// @param s: The string to convert
// @return: The converted float32 value or 0 on error
func String2Float32(s string) float32 {
	v, _ := strconv.ParseFloat(s, 32)
	return float32(v)
}

// String2Int32 converts a string to int32, returns 0 if conversion fails
//
// @param s: The string to convert
// @return: The converted int32 value or 0 on error
func String2Int32(s string) int32 {
	v, _ := strconv.ParseInt(s, 10, 32)
	return int32(v)
}

// String2Int8 converts a string to int8, returns 0 if conversion fails
//
// @param s: The string to convert
// @return: The converted int8 value or 0 on error
func String2Int8(s string) int8 {
	v, _ := strconv.ParseInt(s, 10, 8)
	return int8(v)
}

// String2Int converts a string to int, returns 0 if conversion fails
//
// @param s: The string to convert
// @return: The converted int value or 0 on error
func String2Int(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

// String2UInt converts a string to uint, returns 0 if conversion fails
//
// @param s: The string to convert
// @return: The converted uint value or 0 on error
func String2UInt(s string) uint {
	v, _ := strconv.ParseUint(s, 10, 64)
	return uint(v)
}

// String2Int64 converts a string to int64, returns 0 if conversion fails
//
// @param s: The string to convert
// @return: The converted int64 value or 0 on error
func String2Int64(s string) int64 {
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}

// String2UInt32 converts a string to uint32, returns 0 if conversion fails
//
// @param s: The string to convert
// @return: The converted uint32 value or 0 on error
func String2UInt32(s string) uint32 {
	v, _ := strconv.ParseUint(s, 10, 32)
	return uint32(v)
}

// String2UInt64 converts a string to uint64, returns 0 if conversion fails
//
// @param s: The string to convert
// @return: The converted uint64 value or 0 on error
func String2UInt64(s string) uint64 {
	v, _ := strconv.ParseUint(s, 10, 64)
	return v
}

// UInt642String converts a uint64 to string
//
// @param v: The uint64 value to convert
// @return: The string representation
func UInt642String(v uint64) string {
	return strconv.FormatUint(v, 10)
}

// Int642String converts an int64 to string
//
// @param v: The int64 value to convert
// @return: The string representation
func Int642String(v int64) string {
	return strconv.FormatInt(v, 10)
}

// Int322String converts an int32 to string
//
// @param v: The int32 value to convert
// @return: The string representation
func Int322String(v int32) string {
	return strconv.FormatInt(int64(v), 10)
}

// UInt322String converts a uint32 to string
//
// @param v: The uint32 value to convert
// @return: The string representation
func UInt322String(v uint32) string {
	return strconv.FormatUint(uint64(v), 10)
}

// Int2String converts an int to string
//
// @param v: The int value to convert
// @return: The string representation
func Int2String(v int) string {
	return strconv.Itoa(v)
}

// StringSplit splits a string by separator and returns a slice
// Returns empty slice if input string is empty
//
// @param s: The string to split
// @param sep: The separator string
// @return: Slice of split strings
func StringSplit(s string, sep string) []string {
	if len(s) == 0 {
		return []string{}
	}
	return strings.Split(s, sep)
}

// String2Int32List converts a separated string to []int32
//
// @param s: The string to convert
// @param split: The separator string
// @return: Slice of int32 values
func String2Int32List(s string, split string) []int32 {
	ss := StringSplit(s, split)
	ret := make([]int32, len(ss))
	for i := 0; i < len(ss); i++ {
		v, _ := strconv.ParseInt(ss[i], 10, 32)
		ret[i] = int32(v)
	}
	return ret
}

// String2UInt32List converts a separated string to []uint32
//
// @param s: The string to convert
// @param split: The separator string
// @return: Slice of uint32 values
func String2UInt32List(s string, split string) []uint32 {
	ss := StringSplit(s, split)
	ret := make([]uint32, len(ss))
	for i := 0; i < len(ss); i++ {
		v, _ := strconv.ParseUint(ss[i], 10, 32)
		ret[i] = uint32(v)
	}
	return ret
}

// String2Int64List converts a separated string to []int64
//
// @param s: The string to convert
// @param split: The separator string
// @return: Slice of int64 values
func String2Int64List(s string, split string) []int64 {
	ss := StringSplit(s, split)
	ret := make([]int64, len(ss))
	for i := 0; i < len(ss); i++ {
		v, _ := strconv.ParseInt(ss[i], 10, 64)
		ret[i] = v
	}
	return ret
}

// String2UInt64List converts a separated string to []uint64
//
// @param s: The string to convert
// @param split: The separator string
// @return: Slice of uint64 values
func String2UInt64List(s string, split string) []uint64 {
	ss := StringSplit(s, split)
	ret := make([]uint64, len(ss))
	for i := 0; i < len(ss); i++ {
		v, _ := strconv.ParseUint(ss[i], 10, 64)
		ret[i] = v
	}
	return ret
}

// String2IntList converts a separated string to []int
//
// @param s: The string to convert
// @param split: The separator string
// @return: Slice of int values
func String2IntList(s string, split string) []int {
	ss := StringSplit(s, split)
	ret := make([]int, len(ss))
	for i := 0; i < len(ss); i++ {
		v, _ := strconv.Atoi(ss[i])
		ret[i] = v
	}
	return ret
}

// String2Float64List converts a separated string to []float64
//
// @param s: The string to convert
// @param split: The separator string
// @return: Slice of float64 values
func String2Float64List(s string, split string) []float64 {
	ss := StringSplit(s, split)
	ret := make([]float64, len(ss))
	for i := 0; i < len(ss); i++ {
		v, _ := strconv.ParseFloat(ss[i], 64)
		ret[i] = v
	}
	return ret
}

// String2Float32List converts a separated string to []float32
//
// @param s: The string to convert
// @param split: The separator string
// @return: Slice of float32 values
func String2Float32List(s string, split string) []float32 {
	ss := StringSplit(s, split)
	ret := make([]float32, len(ss))
	for i := 0; i < len(ss); i++ {
		v, _ := strconv.ParseFloat(ss[i], 32)
		ret[i] = float32(v)
	}
	return ret
}

// String2Bool converts a string to bool, returns false if conversion fails
// Supports "true", "false", "1", "0", "t", "f" etc.
//
// @param s: The string to convert
// @return: The converted boolean value
func String2Bool(s string) bool {
	v, _ := strconv.ParseBool(s)
	return v
}

// Bool2String converts a bool to string
//
// @param v: The boolean value to convert
// @return: "true" or "false"
func Bool2String(v bool) string {
	return strconv.FormatBool(v)
}

// String2BoolList converts a separated string to []bool
//
// @param s: The string to convert
// @param split: The separator string
// @return: Slice of boolean values
func String2BoolList(s string, split string) []bool {
	ss := StringSplit(s, split)
	ret := make([]bool, len(ss))
	for i := 0; i < len(ss); i++ {
		v, _ := strconv.ParseBool(ss[i])
		ret[i] = v
	}
	return ret
}
