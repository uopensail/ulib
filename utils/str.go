package utils

import (
	"strconv"
	"strings"
)

// String2Float64 converts s to float64, returning 0 on parse failure.
func String2Float64(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

// String2Float32 converts s to float32, returning 0 on parse failure.
func String2Float32(s string) float32 {
	v, _ := strconv.ParseFloat(s, 32)
	return float32(v)
}

// String2Int32 converts s to int32, returning 0 on parse failure.
func String2Int32(s string) int32 {
	v, _ := strconv.ParseInt(s, 10, 32)
	return int32(v)
}

// String2Int8 converts s to int8, returning 0 on parse failure.
func String2Int8(s string) int8 {
	v, _ := strconv.ParseInt(s, 10, 8)
	return int8(v)
}

// String2Int converts s to int, returning 0 on parse failure.
func String2Int(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

// String2UInt converts s to uint, returning 0 on parse failure.
func String2UInt(s string) uint {
	v, _ := strconv.ParseUint(s, 10, 64)
	return uint(v)
}

// String2Int64 converts s to int64, returning 0 on parse failure.
func String2Int64(s string) int64 {
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}

// String2UInt32 converts s to uint32, returning 0 on parse failure.
func String2UInt32(s string) uint32 {
	v, _ := strconv.ParseUint(s, 10, 32)
	return uint32(v)
}

// String2UInt64 converts s to uint64, returning 0 on parse failure.
func String2UInt64(s string) uint64 {
	v, _ := strconv.ParseUint(s, 10, 64)
	return v
}

// UInt642String formats v as a decimal string.
func UInt642String(v uint64) string {
	return strconv.FormatUint(v, 10)
}

// Int642String formats v as a decimal string.
func Int642String(v int64) string {
	return strconv.FormatInt(v, 10)
}

// Int322String formats v as a decimal string.
func Int322String(v int32) string {
	return strconv.FormatInt(int64(v), 10)
}

// UInt322String formats v as a decimal string.
func UInt322String(v uint32) string {
	return strconv.FormatUint(uint64(v), 10)
}

// Int2String formats v as a decimal string.
func Int2String(v int) string {
	return strconv.Itoa(v)
}

// StringSplit splits s by sep. Returns nil when s is empty, avoiding the
// single-element [""] slice that strings.Split would produce.
func StringSplit(s string, sep string) []string {
	if len(s) == 0 {
		return nil
	}
	return strings.Split(s, sep)
}

// String2Int32List splits s by split and converts each token to int32.
func String2Int32List(s string, split string) []int32 {
	ss := StringSplit(s, split)
	ret := make([]int32, len(ss))
	for i, tok := range ss {
		v, _ := strconv.ParseInt(tok, 10, 32)
		ret[i] = int32(v)
	}
	return ret
}

// String2UInt32List splits s by split and converts each token to uint32.
func String2UInt32List(s string, split string) []uint32 {
	ss := StringSplit(s, split)
	ret := make([]uint32, len(ss))
	for i, tok := range ss {
		v, _ := strconv.ParseUint(tok, 10, 32)
		ret[i] = uint32(v)
	}
	return ret
}

// String2Int64List splits s by split and converts each token to int64.
func String2Int64List(s string, split string) []int64 {
	ss := StringSplit(s, split)
	ret := make([]int64, len(ss))
	for i, tok := range ss {
		ret[i], _ = strconv.ParseInt(tok, 10, 64)
	}
	return ret
}

// String2UInt64List splits s by split and converts each token to uint64.
func String2UInt64List(s string, split string) []uint64 {
	ss := StringSplit(s, split)
	ret := make([]uint64, len(ss))
	for i, tok := range ss {
		ret[i], _ = strconv.ParseUint(tok, 10, 64)
	}
	return ret
}

// String2IntList splits s by split and converts each token to int.
func String2IntList(s string, split string) []int {
	ss := StringSplit(s, split)
	ret := make([]int, len(ss))
	for i, tok := range ss {
		ret[i], _ = strconv.Atoi(tok)
	}
	return ret
}

// String2Float64List splits s by split and converts each token to float64.
func String2Float64List(s string, split string) []float64 {
	ss := StringSplit(s, split)
	ret := make([]float64, len(ss))
	for i, tok := range ss {
		ret[i], _ = strconv.ParseFloat(tok, 64)
	}
	return ret
}

// String2Float32List splits s by split and converts each token to float32.
func String2Float32List(s string, split string) []float32 {
	ss := StringSplit(s, split)
	ret := make([]float32, len(ss))
	for i, tok := range ss {
		v, _ := strconv.ParseFloat(tok, 32)
		ret[i] = float32(v)
	}
	return ret
}

// String2Bool converts s to bool using strconv.ParseBool, returning false on failure.
// Recognised values: "1", "t", "T", "TRUE", "true", "True",
// "0", "f", "F", "FALSE", "false", "False".
func String2Bool(s string) bool {
	v, _ := strconv.ParseBool(s)
	return v
}

// Bool2String converts v to "true" or "false".
func Bool2String(v bool) string {
	return strconv.FormatBool(v)
}

// String2BoolList splits s by split and converts each token to bool.
func String2BoolList(s string, split string) []bool {
	ss := StringSplit(s, split)
	ret := make([]bool, len(ss))
	for i, tok := range ss {
		ret[i], _ = strconv.ParseBool(tok)
	}
	return ret
}
