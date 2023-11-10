package uno

import (
	"encoding/json"
	"fmt"
	"unsafe"

	"github.com/uopensail/ulib/sample"
)

// define arithmetic expression interface
type ArithmeticExpression interface {
	Expression
	GetDataType() sample.DataType
	GetValue() unsafe.Pointer
	Simplify() ArithmeticExpression
}

type Int64 struct {
	BaseExpression
	value int64
}

func (i *Int64) GetDataType() sample.DataType {
	return sample.Int64Type
}

func (i *Int64) GetType() NodeType {
	return kInt64Node
}

func (i *Int64) Simplify() ArithmeticExpression {
	return i
}

func (i *Int64) GetValue() unsafe.Pointer {
	return unsafe.Pointer(&i.value)
}

func (i *Int64) Trivial() bool {
	return true
}

func (i *Int64) MarshalJSON() ([]byte, error) {
	type jsonNode struct {
		Id    int32    `json:"id"`
		Ntype NodeType `json:"ntype"`
		Value int64    `json:"value"`
	}

	node := &jsonNode{
		Id:    i.id,
		Ntype: i.GetType(),
		Value: i.value,
	}
	return json.Marshal(node)
}

func (i *Int64) ToList() []Expression {
	return []Expression{i}
}

type Int64s struct {
	BaseExpression
	value []int64
}

func (i *Int64s) GetDataType() sample.DataType {
	return sample.Int64sType
}

func (i *Int64s) Trivial() bool {
	return true
}

func (i *Int64s) GetValue() unsafe.Pointer {
	return unsafe.Pointer(&i.value)
}

func (i *Int64s) GetType() NodeType {
	return kInt64SliceNode
}

func (i *Int64s) Simplify() ArithmeticExpression {
	return i
}

func (i *Int64s) MarshalJSON() ([]byte, error) {
	type jsonNode struct {
		Id    int32    `json:"id"`
		Ntype NodeType `json:"ntype"`
		Value []int64  `json:"value"`
	}

	node := &jsonNode{
		Id:    i.id,
		Ntype: i.GetType(),
		Value: i.value,
	}
	return json.Marshal(node)
}

func (i *Int64s) ToList() []Expression {
	return []Expression{i}
}

type Float32 struct {
	BaseExpression
	value float32
}

func (f *Float32) GetDataType() sample.DataType {
	return sample.Float32Type
}

func (f *Float32) Trivial() bool {
	return true
}

func (f *Float32) GetType() NodeType {
	return kFloat32Node
}

func (f *Float32) Simplify() ArithmeticExpression {
	return f
}

func (f *Float32) GetValue() unsafe.Pointer {
	return unsafe.Pointer(&f.value)
}

func (f *Float32) MarshalJSON() ([]byte, error) {
	type jsonNode struct {
		Id    int32    `json:"id"`
		Ntype NodeType `json:"ntype"`
		Value float32  `json:"value"`
	}

	node := &jsonNode{
		Id:    f.id,
		Ntype: f.GetType(),
		Value: f.value,
	}
	return json.Marshal(node)
}

func (f *Float32) ToList() []Expression {
	return []Expression{f}
}

type Float32s struct {
	BaseExpression
	value []float32
}

func (f *Float32s) GetDataType() sample.DataType {
	return sample.Float32sType
}

func (i *Float32s) Trivial() bool {
	return true
}

func (f *Float32s) GetType() NodeType {
	return kFloat32SliceNode
}

func (f *Float32s) Simplify() ArithmeticExpression {
	return f
}

func (f *Float32s) GetValue() unsafe.Pointer {
	return unsafe.Pointer(&f.value)
}

func (f *Float32s) MarshalJSON() ([]byte, error) {
	type jsonNode struct {
		Id    int32     `json:"id"`
		Ntype NodeType  `json:"ntype"`
		Value []float32 `json:"value"`
	}

	node := &jsonNode{
		Id:    f.id,
		Ntype: f.GetType(),
		Value: f.value,
	}
	return json.Marshal(node)
}

func (f *Float32s) ToList() []Expression {
	return []Expression{f}
}

type String struct {
	BaseExpression
	value string
}

func (s *String) GetDataType() sample.DataType {
	return sample.StringType
}

func (s *String) GetType() NodeType {
	return kStringNode
}

func (s *String) Trivial() bool {
	return true
}

func (f *String) Simplify() ArithmeticExpression {
	return f
}

func (f *String) GetValue() unsafe.Pointer {
	return unsafe.Pointer(&f.value)
}

func (s *String) MarshalJSON() ([]byte, error) {
	type jsonNode struct {
		Id    int32    `json:"id"`
		Ntype NodeType `json:"ntype"`
		Value string   `json:"value"`
	}

	node := &jsonNode{
		Id:    s.id,
		Ntype: s.GetType(),
		Value: s.value,
	}
	return json.Marshal(node)
}

func (s *String) ToList() []Expression {
	return []Expression{s}
}

type Strings struct {
	BaseExpression
	value []string
}

func (s *Strings) GetDataType() sample.DataType {
	return sample.StringsType
}

func (s *Strings) Trivial() bool {
	return true
}

func (s *Strings) GetType() NodeType {
	return kStringSliceNode
}

func (s *Strings) Simplify() ArithmeticExpression {
	return s
}

func (f *Strings) GetValue() unsafe.Pointer {
	return unsafe.Pointer(&f.value)
}

func (s *Strings) MarshalJSON() ([]byte, error) {
	type jsonNode struct {
		Id    int32    `json:"id"`
		Ntype NodeType `json:"ntype"`
		Value []string `json:"value"`
	}

	node := &jsonNode{
		Id:    s.id,
		Ntype: s.GetType(),
		Value: s.value,
	}
	return json.Marshal(node)
}

func (s *Strings) ToList() []Expression {
	return []Expression{s}
}

type Variable struct {
	BaseExpression
	value string
	dtype sample.DataType
}

func (v *Variable) GetDataType() sample.DataType {
	return v.dtype
}

func (v *Variable) Trivial() bool {
	return false
}

func (v *Variable) GetType() NodeType {
	return kVarNode
}

func (v *Variable) Simplify() ArithmeticExpression {
	return v
}

func (v *Variable) GetValue() unsafe.Pointer {
	panic("not implemented")
}

func (c *Variable) MarshalJSON() ([]byte, error) {
	type jsonNode struct {
		Id    int32           `json:"id"`
		Ntype NodeType        `json:"ntype"`
		Value string          `json:"value"`
		Dtype sample.DataType `json:"dtype"`
	}

	node := &jsonNode{
		Id:    c.id,
		Ntype: c.GetType(),
		Value: c.value,
		Dtype: c.dtype,
	}
	return json.Marshal(node)
}

func (v *Variable) ToList() []Expression {
	return []Expression{v}
}

type Function struct {
	BaseExpression
	function string
	args     []ArithmeticExpression
}

func (f *Function) check() {
	inputs := FUNCTION_IO_TYPES[f.function]["in"]
	if len(inputs) != len(f.args) {
		panic("function argument check failed")
	}
	for i := 0; i < len(f.args); i++ {
		if f.args[i].GetDataType() != inputs[i] {
			panic(fmt.Sprintf("function argument:%d type check failed", i))
		}
	}
}

func (f *Function) GetDataType() sample.DataType {
	return FUNCTION_IO_TYPES[f.function]["out"][0]
}

func (f *Function) GetType() NodeType {
	return kFunctionNode
}

func (f *Function) Trivial() bool {
	if len(f.args) == 0 {
		return false
	}
	for i := 0; i < len(f.args); i++ {
		if !f.args[i].Trivial() {
			return false
		}
	}
	return true
}

func (f *Function) Simplify() ArithmeticExpression {
	for i := 0; i < len(f.args); i++ {
		f.args[i] = f.args[i].Simplify()
	}

	if !f.Trivial() {
		return f
	}

	dtype := f.GetDataType()
	args := make([]unsafe.Pointer, len(f.args)+1)
	for i := 0; i < len(f.args); i++ {
		args[i] = f.args[i].GetValue()
	}
	if dtype == sample.Int64Type {
		ret := call_for_int64(f.function, args)
		return &Int64{value: ret}
	} else if dtype == sample.Float32Type {
		ret := call_for_float32(f.function, args)
		return &Float32{value: ret}
	} else if dtype == sample.StringType {
		ret := call_for_string(f.function, args)
		return &String{value: ret}
	}
	panic(fmt.Sprintf("datatype: %d not supported", dtype))
}

func (f *Function) GetValue() unsafe.Pointer {
	panic("not implemented")
}

func (f *Function) MarshalJSON() ([]byte, error) {
	type jsonNode struct {
		ID    int32    `json:"id"`
		Ntype NodeType `json:"ntype"`
		Func  string   `json:"func"`
		Args  []int    `json:"args"`
	}

	node := &jsonNode{
		ID:    f.id,
		Ntype: f.GetType(),
		Func:  f.function,
		Args:  make([]int, len(f.args)),
	}
	for i := 0; i < len(f.args); i++ {
		node.Args[i] = int(f.args[i].GetId())
	}
	return json.Marshal(node)
}

func (f *Function) ToList() []Expression {
	exprs := make([]Expression, 0, len(f.args)+1)
	for i := 0; i < len(f.args); i++ {
		tmp := f.args[i].ToList()
		exprs = append(exprs, tmp...)
	}
	exprs = append(exprs, f)
	return exprs
}
