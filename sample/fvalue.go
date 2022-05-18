package sample

import (
	"bytes"
	fmt "fmt"
	"strconv"
	"strings"

	"github.com/uopensail/ulib/utils"
)

type FValueType int8

const (
	StringListFValueType  = 1
	Float32ListFValueType = 2
	Int64ListFValueType   = 3
)

type FValue struct {
	FValueType
	StringList  []string
	Float32List []float32
	Int64List   []int64
}

func (fv *FValue) Length() int {
	switch fv.FValueType {
	case StringListFValueType:
		return len(fv.StringList)
	case Float32ListFValueType:
		return len(fv.Float32List)
	case Int64ListFValueType:
		return len(fv.Int64List)
	}
	return 0
}

func (fV *FValue) ToString(sep string) string {
	switch fV.FValueType {
	case StringListFValueType:
		return strings.Join(fV.StringList, sep)
	case Float32ListFValueType:
		buff := bytes.Buffer{}
		for i := 0; i < len(fV.Float32List); i++ {
			buff.WriteString(fmt.Sprintf("%.6f", fV.Float32List[i]))
			if i != 0 {
				buff.WriteString(sep)
			}
		}
		return buff.String()
	case Int64ListFValueType:
		buff := bytes.Buffer{}
		for i := 0; i < len(fV.Int64List); i++ {
			buff.WriteString(strconv.FormatInt(fV.Int64List[i], 10))
			if i != 0 {
				buff.WriteString(sep)
			}
		}
		return buff.String()
	}
	return ""
}

func (fV *FValue) GetInt64() int64 {
	if len(fV.Int64List) > 0 {
		return fV.Int64List[0]
	}
	return 0
}

func (fV *FValue) TryGetFloat32() float32 {
	switch fV.FValueType {
	case Float32ListFValueType:
		if len(fV.Float32List) > 0 {
			return fV.Float32List[0]
		}
	case Int64ListFValueType:
		if len(fV.Int64List) > 0 {
			return float32(fV.Int64List[0])
		}
	case StringListFValueType:
		if len(fV.StringList) > 0 {
			return utils.String2Float32(fV.StringList[0])
		}
	}
	return 0
}

func (fV *FValue) GetFloat32() float32 {
	if len(fV.Float32List) > 0 {
		return fV.Float32List[0]
	}
	return 0
}

func (fV *FValue) GetString() string {
	if len(fV.StringList) > 0 {
		return fV.StringList[0]
	}
	return ""
}

type FeaturesW struct {
	M map[string]FValue
	*Features
}

func (fs FeaturesW) GetFValue(field string) (FValue, bool) {
	if v, ok := fs.M[field]; ok {
		return v, ok
	}
	return FValue{}, false
}
func (fs FeaturesW) GetInt64ListByKey(field string) []int64 {
	if v, ok := fs.M[field]; ok {
		return v.Int64List
	}
	return nil
}

func (fs FeaturesW) GetStringListByKey(field string) []string {
	if v, ok := fs.M[field]; ok {
		return v.StringList
	}
	return nil
}

func (fs FeaturesW) GeFloat32ListByKey(field string) []float32 {
	if v, ok := fs.M[field]; ok {
		return v.Float32List
	}
	return nil
}

func (fs FeaturesW) GetInt64(filed string) int64 {
	if v, ok := fs.M[filed]; ok && len(v.Int64List) > 0 {
		return v.Int64List[0]
	}
	return 0
}

func (fs FeaturesW) TryGetFloat32(filed string) float32 {
	if v, ok := fs.M[filed]; ok {
		switch v.FValueType {
		case Float32ListFValueType:
			if len(v.Float32List) > 0 {
				return v.Float32List[0]
			}
		case Int64ListFValueType:
			if len(v.Int64List) > 0 {
				return float32(v.Int64List[0])
			}
		case StringListFValueType:
			if len(v.StringList) > 0 {
				return utils.String2Float32(v.StringList[0])
			}
		}

	}
	return 0
}

func (fs FeaturesW) GetFloat32(filed string) float32 {
	if v, ok := fs.M[filed]; ok && len(v.Float32List) > 0 {
		return v.Float32List[0]
	}
	return 0
}

func (fs FeaturesW) GetString(filed string) string {
	if v, ok := fs.M[filed]; ok && len(v.StringList) > 0 {
		return v.StringList[0]
	}
	return ""
}

func (fs FeaturesW) GetInt64List(filed string) []int64 {
	if v, ok := fs.M[filed]; ok {
		return v.Int64List
	}
	return nil
}

func (fs FeaturesW) GetFloat32List(filed string) []float32 {
	if v, ok := fs.M[filed]; ok {
		return v.Float32List
	}
	return nil
}

func (fs FeaturesW) GetStringList(filed string) []string {
	if v, ok := fs.M[filed]; ok {
		return v.StringList
	}
	return nil
}

func StringListIntersectionIsNone(aList []string, bList []string) bool {
	for a := 0; a < len(aList); a++ {
		for b := 0; b < len(bList); b++ {
			if aList[a] == bList[b] {
				return false
			}
		}
	}
	return true
}

func Float32ListIntersectionIsNone(aList []float32, bList []float32) bool {
	for a := 0; a < len(aList); a++ {
		for b := 0; b < len(bList); b++ {
			if aList[a] == bList[b] {
				return false
			}
		}
	}
	return true
}

func Int64ListIntersectionIsNone(aList []int64, bList []int64) bool {
	for a := 0; a < len(aList); a++ {
		for b := 0; b < len(bList); b++ {
			if aList[a] == bList[b] {
				return false
			}
		}
	}
	return true
}

//判断交集是否为空
func (fs FeaturesW) IntersectionIsNone(filed string, featB *FValue) bool {
	featA, ok := fs.M[filed]
	if ok == false || featB == nil {
		return true
	}
	if featB.FValueType != featA.FValueType {
		return true
	}
	switch featA.FValueType {
	case StringListFValueType:

		return StringListIntersectionIsNone(featA.StringList, featB.StringList)
	case Float32ListFValueType:
		return Float32ListIntersectionIsNone(featA.Float32List, featB.Float32List)
	case Int64ListFValueType:
		return Int64ListIntersectionIsNone(featA.Int64List, featB.Int64List)

	}
	return true
}
func StringListIntersectionLen(aList []string, bList []string) int {
	k := 0
	for a := 0; a < len(aList); a++ {
		for b := 0; b < len(bList); b++ {
			if aList[a] == bList[b] {
				k++
			}
		}
	}
	return k
}

func Float32ListIntersectionLen(aList []float32, bList []float32) int {
	k := 0
	for a := 0; a < len(aList); a++ {
		for b := 0; b < len(bList); b++ {
			if aList[a] == bList[b] {
				k++
			}
		}
	}
	return k
}
func Int64ListIntersectionLen(aList []int64, bList []int64) int {
	k := 0
	for a := 0; a < len(aList); a++ {
		for b := 0; b < len(bList); b++ {
			if aList[a] == bList[b] {
				k++
			}
		}
	}
	return k
}

//判断交集的长度
func (fs FeaturesW) IntersectionLen(filed string, featB *FValue) int {
	featA, ok := fs.M[filed]
	if ok == false || featB == nil {
		return 0
	}
	if featB.FValueType != featA.FValueType {
		return 0
	}
	switch featA.FValueType {
	case StringListFValueType:
		return StringListIntersectionLen(featA.StringList, featB.StringList)
	case Float32ListFValueType:
		return Float32ListIntersectionLen(featA.Float32List, featB.Float32List)
	case Int64ListFValueType:
		return Int64ListIntersectionLen(featA.Int64List, featB.Int64List)

	}
	return 0
}

func MakeFeaturesM(feats map[string]*Feature) FeaturesW {
	ret := FeaturesW{
		M: make(map[string]FValue, len(feats)),
		Features: &Features{
			Feature: feats,
		},
	}

	for k, v := range feats {
		switch v.Kind.(type) {
		case *Feature_BytesList:
			vs := v.GetBytesList()
			if vs != nil && len(vs.Value) > 0 {
				strList := make([]string, len(vs.Value))
				for i := 0; i < len(vs.Value); i++ {
					strList[i] = string(vs.Value[i])
				}
				ret.M[k] = FValue{
					FValueType: StringListFValueType,
					StringList: strList,
				}
			}

		case *Feature_Int64List:
			vs := v.GetInt64List()
			if vs != nil && len(vs.Value) > 0 {
				ret.M[k] = FValue{
					FValueType: Int64ListFValueType,
					Int64List:  vs.Value,
				}
			}
		case *Feature_FloatList:
			vs := v.GetFloatList()
			if vs != nil && len(vs.Value) > 0 {
				ret.M[k] = FValue{
					FValueType:  Float32ListFValueType,
					Float32List: vs.Value,
				}
			}
		}
	}
	return ret
}

func MakeFeatures(tfs *Features) FeaturesW {
	return MakeFeaturesM(tfs.Feature)
}
