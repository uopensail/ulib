package utils

import (
	"strconv"
	"strings"
)

func String2Float64(s string) float64 {
	v, _ := strconv.ParseFloat(s, 32)
	return v
}

func String2Float32(s string) float32 {
	v, _ := strconv.ParseFloat(s, 32)
	return float32(v)
}

func String2Int32(s string) int32 {
	v, _ := strconv.Atoi(s)
	return int32(v)
}

func String2Int8(s string) int8 {
	v, _ := strconv.Atoi(s)
	return int8(v)
}

func String2Int(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func String2UInt(s string) uint {
	v, _ := strconv.ParseUint(s, 10, 64)
	return uint(v)
}

func String2Int64(s string) int64 {
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}

func String2UInt32(s string) uint32 {
	v, _ := strconv.ParseUint(s, 10, 64)
	return uint32(v)
}

func String2UInt64(s string) uint64 {
	v, _ := strconv.ParseUint(s, 10, 64)
	return v
}

func UInt642String(v uint64) string {
	s := strconv.FormatUint(v, 10)
	return s
}

func Int642String(v int64) string {
	s := strconv.FormatInt(v, 10)
	return s
}

func Int322String(v int32) string {
	s := strconv.FormatInt(int64(v), 10)
	return s
}

func UInt322String(v uint32) string {
	s := strconv.FormatUint(uint64(v), 10)
	return s
}

func Int2String(v int) string {
	s := strconv.FormatInt(int64(v), 10)
	return s
}

func StringSplit(s string, sep string) []string {
	if len(s) <= 0 {
		return []string{}
	}
	return strings.Split(s, sep)
}

func String2Int32List(s string, split string) []int32 {
	ss := StringSplit(s, split)
	ret := make([]int32, len(ss))
	for i := 0; i < len(ss); i++ {
		v, _ := strconv.ParseInt(ss[i], 10, 64)
		ret[i] = int32(v)
	}

	return ret
}

func String2UInt32List(s string, split string) []uint32 {
	ss := StringSplit(s, split)
	ret := make([]uint32, len(ss))
	for i := 0; i < len(ss); i++ {
		v, _ := strconv.ParseUint(ss[i], 10, 64)
		ret[i] = uint32(v)
	}

	return ret
}

func String2Int64List(s string, split string) []int64 {
	ss := StringSplit(s, split)
	ret := make([]int64, len(ss))
	for i := 0; i < len(ss); i++ {
		v, _ := strconv.ParseInt(ss[i], 10, 64)
		ret[i] = int64(v)
	}

	return ret
}

func String2UInt64List(s string, split string) []uint64 {
	ss := StringSplit(s, split)
	ret := make([]uint64, len(ss))
	for i := 0; i < len(ss); i++ {
		v, _ := strconv.ParseUint(ss[i], 10, 64)
		ret[i] = v
	}

	return ret
}

func String2IntList(s string, split string) []int {
	ss := StringSplit(s, split)
	ret := make([]int, len(ss))
	for i := 0; i < len(ss); i++ {
		v, _ := strconv.ParseInt(ss[i], 10, 64)
		ret[i] = int(v)
	}

	return ret
}

func String2Float64List(s string, split string) []float64 {
	ss := StringSplit(s, split)
	ret := make([]float64, len(ss))
	for i := 0; i < len(ss); i++ {
		v, _ := strconv.ParseFloat(ss[i], 64)
		ret[i] = float64(v)
	}

	return ret
}

func String2Float32List(s string, split string) []float32 {
	ss := StringSplit(s, split)
	ret := make([]float32, len(ss))
	for i := 0; i < len(ss); i++ {
		v, _ := strconv.ParseFloat(ss[i], 32)
		ret[i] = float32(v)
	}

	return ret
}