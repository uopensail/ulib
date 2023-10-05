package uno

import (
	"encoding/json"
	"fmt"
)

type BooleanExpression interface {
	Expression
	Negation() BooleanExpression
	Simplify() BooleanExpression
}

type Literal struct {
	BaseExpression
	value bool
}

func (l *Literal) GetType() NodeType {
	return kLiteralNode
}

func (l *Literal) Negation() BooleanExpression {
	return &Literal{value: !l.value}
}

func (l *Literal) Simplify() BooleanExpression {
	return l
}

func (l *Literal) Trivial() bool {
	return true
}

func (l *Literal) Marshal() ([]byte, error) {
	type jsonNode struct {
		Id    int32    `json:"id"`
		Ntype NodeType `json:"ntype"`
		Value bool     `json:"value"`
	}

	node := &jsonNode{
		Id:    l.id,
		Ntype: l.GetType(),
		Value: l.value,
	}
	return json.Marshal(node)

}

func (l *Literal) ToList() []Expression {
	return []Expression{l}
}

type And struct {
	BaseExpression
	left  BooleanExpression
	right BooleanExpression
}

func (a *And) GetType() NodeType {
	return kAndNode
}

func (a *And) Negation() BooleanExpression {
	left := a.left.Negation()
	right := a.right.Negation()
	return &Or{left: left, right: right}
}

func (a *And) Simplify() BooleanExpression {
	left := a.left.Simplify()
	right := a.right.Simplify()

	if left.Trivial() {
		if left.(*Literal).value {
			return right
		}
		return &Literal{value: false}
	}
	if right.Trivial() {
		if right.(*Literal).value {
			return left
		}
		return &Literal{value: false}
	}
	a.left = left
	a.right = right
	return a
}

func (a *And) Trivial() bool {
	tmp := a.Simplify()
	if tmp == a {
		return false
	}
	return tmp.Trivial()
}

func (a *And) Marshal() ([]byte, error) {
	type jsonNode struct {
		Id    int32    `json:"id"`
		Ntype NodeType `json:"ntype"`
		Left  int32    `json:"left"`
		Right int32    `json:"right"`
	}

	node := &jsonNode{
		Id:    a.id,
		Ntype: a.GetType(),
		Left:  a.left.GetId(),
		Right: a.right.GetId(),
	}
	return json.Marshal(node)

}

func (a *And) ToList() []Expression {
	exprs := make([]Expression, 0, 3)
	left := a.left.ToList()
	right := a.right.ToList()
	exprs = append(exprs, left...)
	exprs = append(exprs, right...)
	exprs = append(exprs, a)
	return exprs
}

type Or struct {
	BaseExpression
	left  BooleanExpression
	right BooleanExpression
}

func (o *Or) GetType() NodeType {
	return kOrNode
}

func (o *Or) Negation() BooleanExpression {
	left := o.left.Negation()
	right := o.right.Negation()
	return &And{left: left, right: right}
}

func (o *Or) Simplify() BooleanExpression {
	left := o.left.Simplify()
	right := o.right.Simplify()

	if left.Trivial() {
		if !left.(*Literal).value {
			return right
		}
		return &Literal{value: true}
	}
	if right.Trivial() {
		if !right.(*Literal).value {
			return left
		}
		return &Literal{value: true}
	}
	o.left = left
	o.right = right
	return o
}

func (o *Or) Trivial() bool {
	tmp := o.Simplify()
	if tmp == o {
		return false
	}
	return tmp.Trivial()
}

func (o *Or) Marshal() ([]byte, error) {
	type jsonNode struct {
		Id    int32    `json:"id"`
		Ntype NodeType `json:"ntype"`
		Left  int32    `json:"left"`
		Right int32    `json:"right"`
	}

	node := &jsonNode{
		Id:    o.id,
		Ntype: o.GetType(),
		Left:  o.left.GetId(),
		Right: o.right.GetId(),
	}
	return json.Marshal(node)

}

func (o *Or) ToList() []Expression {
	exprs := make([]Expression, 0, 3)
	left := o.left.ToList()
	right := o.right.ToList()
	exprs = append(exprs, left...)
	exprs = append(exprs, right...)
	exprs = append(exprs, o)
	return exprs
}

type Cmp struct {
	BaseExpression
	left  ArithmeticExpression
	right ArithmeticExpression
	op    CmpType
	dtype DataType
}

func (c *Cmp) GetType() NodeType {
	return kCmpNode
}

func (c *Cmp) Negation() BooleanExpression {
	if c.op == kLessThan {
		return &Cmp{left: c.left, right: c.right, op: kGreaterThanEqual}
	} else if c.op == kLessThanEqual {
		return &Cmp{left: c.left, right: c.right, op: kGreaterThan}
	} else if c.op == kEqual {
		return &Cmp{left: c.left, right: c.right, op: kNotEqual}
	} else if c.op == kNotEqual {
		return &Cmp{left: c.left, right: c.right, op: kEqual}
	} else if c.op == kGreaterThanEqual {
		return &Cmp{left: c.left, right: c.right, op: kLessThan}
	} else if c.op == kGreaterThan {
		return &Cmp{left: c.left, right: c.right, op: kLessThanEqual}
	}
	panic(fmt.Sprintf("cmp op:%d not supported", c.op))
}

func (c *Cmp) Simplify() BooleanExpression {
	c.left = c.left.Simplify()
	c.right = c.right.Simplify()

	if !c.left.Trivial() || !c.right.Trivial() {
		return c
	}

	if c.dtype == kInt64 {
		status := compare[int64](c.left.(*Int64).value, c.right.(*Int64).value, c.op)
		return &Literal{value: status}
	} else if c.dtype == kFloat32 {
		status := compare[float32](c.left.(*Float32).value, c.right.(*Float32).value, c.op)
		return &Literal{value: status}
	} else if c.dtype == kString {
		status := compare[string](c.left.(*String).value, c.right.(*String).value, c.op)
		return &Literal{value: status}
	}
	panic(fmt.Sprintf("data type:%d not supported", c.dtype))
}

func (c *Cmp) Trivial() bool {
	if c.left.Trivial() && c.right.Trivial() {
		return true
	}
	return false
}

func (c *Cmp) Marshal() ([]byte, error) {
	type jsonNode struct {
		Id    int32    `json:"id"`
		Ntype NodeType `json:"ntype"`
		Dtype DataType `json:"dtype"`
		Left  int32    `json:"left"`
		Right int32    `json:"right"`
		Cmp   CmpType  `json:"cmp"`
	}

	node := &jsonNode{
		Id:    c.id,
		Ntype: c.GetType(),
		Left:  c.left.GetId(),
		Right: c.right.GetId(),
		Cmp:   c.op,
		Dtype: c.dtype,
	}
	return json.Marshal(node)
}

func (c *Cmp) ToList() []Expression {
	exprs := make([]Expression, 0, 3)
	left := c.left.ToList()
	right := c.right.ToList()
	exprs = append(exprs, left...)
	exprs = append(exprs, right...)
	exprs = append(exprs, c)
	return exprs
}

type In struct {
	BaseExpression
	left  ArithmeticExpression
	right ArithmeticExpression
	dtype DataType
}

func (i *In) GetType() NodeType {
	return kInNode
}

func (i *In) Negation() BooleanExpression {
	return &NotIn{left: i.left, right: i.right}
}

func (i *In) Simplify() BooleanExpression {
	i.left = i.left.Simplify()

	if !i.left.Trivial() {
		return i
	}

	if i.dtype == kInt64 {

		status := inarray[int64](i.left.(*Int64).value, i.right.(*Int64s).value)
		return &Literal{value: status}
	} else if i.dtype == kFloat32 {
		status := inarray[float32](i.left.(*Float32).value, i.right.(*Float32s).value)
		return &Literal{value: status}
	} else if i.dtype == kString {
		status := inarray[string](i.left.(*String).value, i.right.(*Strings).value)
		return &Literal{value: status}
	}
	panic(fmt.Sprintf("data type:%d not supported", i.dtype))
}

func (i *In) Trivial() bool {
	return i.left.Trivial()
}

func (i *In) Marshal() ([]byte, error) {
	type jsonNode struct {
		Id    int32    `json:"id"`
		Ntype NodeType `json:"ntype"`
		Dtype DataType `json:"dtype"`
		Left  int32    `json:"left"`
		Right int32    `json:"right"`
	}

	node := &jsonNode{
		Id:    i.id,
		Ntype: i.GetType(),
		Left:  i.left.GetId(),
		Right: i.right.GetId(),
		Dtype: i.dtype,
	}
	return json.Marshal(node)
}

func (i *In) ToList() []Expression {
	exprs := make([]Expression, 0, 3)
	left := i.left.ToList()
	exprs = append(exprs, left...)
	exprs = append(exprs, i)
	return exprs
}

type NotIn struct {
	BaseExpression
	left  ArithmeticExpression
	right ArithmeticExpression
	dtype DataType
}

func (i *NotIn) GetId() int32 {
	return i.id
}

func (i *NotIn) GetType() NodeType {
	return kNotInNode
}

func (i *NotIn) Negation() BooleanExpression {
	return &In{left: i.left, right: i.right}
}

func (i *NotIn) Simplify() BooleanExpression {
	i.left = i.left.Simplify()

	if !i.left.Trivial() {
		return i
	}

	if i.dtype == kInt64 {
		status := inarray[int64](i.left.(*Int64).value, i.right.(*Int64s).value)
		return &Literal{value: !status}
	} else if i.dtype == kFloat32 {
		status := inarray[float32](i.left.(*Float32).value, i.right.(*Float32s).value)
		return &Literal{value: !status}
	} else if i.dtype == kString {
		status := inarray[string](i.left.(*String).value, i.right.(*Strings).value)
		return &Literal{value: !status}
	}
	panic(fmt.Sprintf("data type:%d not supported", i.dtype))
}

func (i *NotIn) Trivial() bool {
	return i.left.Trivial()
}

func (i *NotIn) Marshal() ([]byte, error) {
	type jsonNode struct {
		Id    int32    `json:"id"`
		Ntype NodeType `json:"ntype"`
		Dtype DataType `json:"dtype"`
		Left  int32    `json:"left"`
		Right int32    `json:"right"`
	}

	node := &jsonNode{
		Id:    i.id,
		Ntype: i.GetType(),
		Left:  i.left.GetId(),
		Right: i.right.GetId(),
		Dtype: i.dtype,
	}
	return json.Marshal(node)
}

func (i *NotIn) ToList() []Expression {
	exprs := make([]Expression, 0, 3)
	left := i.left.ToList()
	exprs = append(exprs, left...)
	exprs = append(exprs, i)
	return exprs
}

func inarray[T int64 | float32 | string](a T, array []T) bool {
	for i := 0; i < len(array); i++ {
		if a == array[i] {
			return true
		}
	}
	return false
}

func compare[T int64 | float32 | string](a, b T, cmp CmpType) bool {
	if cmp == kEqual {
		return a == b
	} else if cmp == kNotEqual {
		return a != b
	} else if cmp == kGreaterThan {
		return a > b
	} else if cmp == kGreaterThanEqual {
		return a >= b
	} else if cmp == kLessThan {
		return a < b
	} else if cmp == kLessThanEqual {
		return a <= b
	}
	panic(fmt.Sprintf("invalid cmp: %d", cmp))
}
