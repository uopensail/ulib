package uno

import (
	"fmt"
	"strconv"
	"strings"
)

type Listener struct {
	BaseunoListener
	arithmetics Stack[ArithmeticExpression]
	booleans    Stack[BooleanExpression]
}

// ExitCmpBooleanExpression is called when production CmpBooleanExpression is exited.
func (s *Listener) ExitCmpBooleanExpression(ctx *CmpBooleanExpressionContext) {
	right := s.arithmetics.Pop()
	left := s.arithmetics.Pop()
	if left.GetDataType() != right.GetDataType() {
		panic("DataType Mismatch")
	}

	dtype := left.GetDataType()
	mark := ctx.T_COMPARE().GetText()
	op := kErrorCType
	if mark == "==" || mark == "=" {
		op = kEqual
	} else if mark == "!=" || mark == "<>" {
		op = kNotEqual
	} else if mark == "<" {
		op = kLessThan
	} else if mark == "<=" {
		op = kLessThanEqual
	} else if mark == ">=" {
		op = kGreaterThanEqual
	} else if mark == ">" {
		op = kGreaterThan
	}
	tmp := &Cmp{left: left, right: right, op: op, dtype: dtype}
	s.booleans.Push(tmp)
}

// ExitNotBooleanExpression is called when production NotBooleanExpression is exited.
func (s *Listener) ExitNotBooleanExpression(ctx *NotBooleanExpressionContext) {
	node := s.booleans.Pop()
	node = node.Negation()
	s.booleans.Push(node)
}

// ExitOrBooleanExpression is called when production OrBooleanExpression is exited.
func (s *Listener) ExitOrBooleanExpression(ctx *OrBooleanExpressionContext) {
	right := s.booleans.Pop()
	left := s.booleans.Pop()
	tmp := &Or{left: left, right: right}
	s.booleans.Push(tmp)
}

// ExitTrueBooleanExpression is called when production TrueBooleanExpression is exited.
func (s *Listener) ExitTrueBooleanExpression(ctx *TrueBooleanExpressionContext) {
	s.booleans.Push(&Literal{value: true})
}

// ExitAndBooleanExpression is called when production AndBooleanExpression is exited.
func (s *Listener) ExitAndBooleanExpression(ctx *AndBooleanExpressionContext) {
	right := s.booleans.Pop()
	left := s.booleans.Pop()
	tmp := &And{left: left, right: right}
	s.booleans.Push(tmp)
}

// ExitNotInBooleanExpression is called when production NotInBooleanExpression is exited.
func (s *Listener) ExitNotInBooleanExpression(ctx *NotInBooleanExpressionContext) {
	expr := s.arithmetics.Pop()

	if ctx.DECIMAL_LIST() != nil {
		list := s.parseDecimalList(ctx.DECIMAL_LIST().GetText())
		array := &Float32s{value: list}
		if expr.GetDataType() != kFloat32 {
			panic("DataType Mismatch")
		}
		s.booleans.Push(&NotIn{left: expr, right: array, dtype: kFloat32})
	} else if ctx.INTEGER_LIST() != nil {
		if expr.GetDataType() != kInt64 {
			panic("DataType Mismatch")
		}
		list := s.parseIntegerList(ctx.INTEGER_LIST().GetText())
		array := &Int64s{value: list}
		s.booleans.Push(&NotIn{left: expr, right: array, dtype: kInt64})
	} else {
		if expr.GetDataType() != kString {
			panic("DataType MisMatch")
		}
		list := s.parseStringList(ctx.STRING_LIST().GetText())
		array := &Strings{value: list}
		s.booleans.Push(&NotIn{left: expr, right: array, dtype: kString})
	}
}

// ExitFalseBooleanExpression is called when production FalseBooleanExpression is exited.
func (s *Listener) ExitFalseBooleanExpression(ctx *FalseBooleanExpressionContext) {
	s.booleans.Push(&Literal{value: false})
}

// ExitInBooleanExpression is called when production InBooleanExpression is exited.
func (s *Listener) ExitInBooleanExpression(ctx *InBooleanExpressionContext) {
	expr := s.arithmetics.Pop()

	if ctx.DECIMAL_LIST() != nil {
		if expr.GetDataType() != kFloat32 {
			panic("DataType Mismatch")
		}
		list := s.parseDecimalList(ctx.DECIMAL_LIST().GetText())
		array := &Float32s{value: list}
		s.booleans.Push(&In{left: expr, right: array, dtype: kFloat32})
	} else if ctx.INTEGER_LIST() != nil {
		if expr.GetDataType() != kInt64 {
			panic("DataType Mismatch")
		}
		list := s.parseIntegerList(ctx.INTEGER_LIST().GetText())
		array := &Int64s{value: list}
		s.booleans.Push(&In{left: expr, right: array, dtype: kInt64})
	} else {
		if expr.GetDataType() != kString {
			panic("DataType Mismatch")
		}
		list := s.parseStringList(ctx.STRING_LIST().GetText())
		array := &Strings{value: list}
		s.booleans.Push(&In{left: expr, right: array, dtype: kString})
	}
}

// ExitAddArithmeticExpression is called when production AddArithmeticExpression is exited.
func (s *Listener) ExitAddArithmeticExpression(ctx *AddArithmeticExpressionContext) {
	second := s.arithmetics.Pop()
	first := s.arithmetics.Pop()

	if first.GetDataType() != second.GetDataType() {
		panic("DataType mismatch")
	}

	function := "addf"
	if second.Trivial() {
		second = second.Simplify()
		if second.GetDataType() == kFloat32 {
			function = "divf"
		} else if second.GetDataType() == kInt64 {
			function = "addi"
		}
	}

	tmp := &Function{function: function, args: []ArithmeticExpression{first, second}}
	s.arithmetics.Push(tmp)
}

// ExitStringArithmeticExpression is called when production StringArithmeticExpression is exited.
func (s *Listener) ExitStringArithmeticExpression(ctx *StringArithmeticExpressionContext) {
	val := ctx.STRING().GetText()
	s.arithmetics.Push(&String{value: val[1 : len(val)-1]})
}

// ExitIntegerArithmeticExpression is called when production IntegerArithmeticExpression is exited.
func (s *Listener) ExitIntegerArithmeticExpression(ctx *IntegerArithmeticExpressionContext) {
	val, err := strconv.ParseInt(ctx.INTEGER().GetText(), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("parse int error:%v", err))
	}
	s.arithmetics.Push(&Int64{value: val})
}

// ExitDecimalArithmeticExpression is called when production DecimalArithmeticExpression is exited.
func (s *Listener) ExitDecimalArithmeticExpression(ctx *DecimalArithmeticExpressionContext) {
	val, err := strconv.ParseFloat(ctx.DECIMAL().GetText(), 32)
	if err != nil {
		panic(fmt.Sprintf("parse float error:%v", err))
	}
	s.arithmetics.Push(&Float32{value: float32(val)})
}

// ExitFuncArithmeticExpression is called when production FuncArithmeticExpression is exited.
func (s *Listener) ExitFuncArithmeticExpression(ctx *FuncArithmeticExpressionContext) {
	function := strings.ToLower(ctx.IDENTIFIER().GetText())
	n := len(ctx.AllArithmetic_expression())
	args := make([]ArithmeticExpression, n)
	for i := 0; i < n; i++ {
		args[n-1-i] = s.arithmetics.Pop()
	}

	tmp := &Function{function: function, args: args}
	s.arithmetics.Push(tmp)
}

// ExitColumnArithmeticExpression is called when production ColumnArithmeticExpression is exited.
func (s *Listener) ExitColumnArithmeticExpression(ctx *ColumnArithmeticExpressionContext) {
	mark := ctx.Type_marker().GetText()
	mark = strings.ToLower(mark)
	column := ctx.IDENTIFIER().GetText()
	dtype := kFloat32
	if mark == "[int64]" {
		dtype = kInt64
	} else if mark == "[int64s]" {
		dtype = kInt64s
	} else if mark == "[float32]" {
		dtype = kFloat32
	} else if mark == "[float32s]" {
		dtype = kFloat32s
	} else if mark == "[string]" {
		dtype = kString
	} else if mark == "[strings]" {
		dtype = kStrings
	}
	tmp := &Variable{value: column, dtype: dtype}
	s.arithmetics.Push(tmp)
}

// ExitDivArithmeticExpression is called when production DivArithmeticExpression is exited.
func (s *Listener) ExitDivArithmeticExpression(ctx *DivArithmeticExpressionContext) {
	second := s.arithmetics.Pop()
	first := s.arithmetics.Pop()

	if first.GetDataType() != second.GetDataType() {
		panic("DataType Mismatch")
	}

	function := "divf"
	if second.Trivial() {
		second = second.Simplify()
		if second.GetDataType() == kFloat32 {
			if second.(*Float32).value == 0.0 {
				panic("Devide By Zero")
			}
		} else if second.GetDataType() == kInt64 {
			if second.(*Int64).value == 0 {
				panic("Devide By Zero")
			}
			function = "divi"
		}
	}

	tmp := &Function{function: function, args: []ArithmeticExpression{first, second}}
	s.arithmetics.Push(tmp)
}

// ExitFieldColumnArithmeticExpression is called when production FieldColumnArithmeticExpression is exited.
func (s *Listener) ExitFieldColumnArithmeticExpression(ctx *FieldColumnArithmeticExpressionContext) {
	mark := ctx.Type_marker().GetText()
	mark = strings.ToLower(mark)
	column := fmt.Sprintf("%s.%s", ctx.IDENTIFIER(0).GetText(), ctx.IDENTIFIER(1).GetText())
	dtype := kFloat32
	if mark == "[int64]" {
		dtype = kInt64
	} else if mark == "[int64s]" {
		dtype = kInt64s
	} else if mark == "[float32]" {
		dtype = kFloat32
	} else if mark == "[float32s]" {
		dtype = kFloat32s
	} else if mark == "[string]" {
		dtype = kString
	} else if mark == "[strings]" {
		dtype = kStrings
	}
	tmp := &Variable{value: column, dtype: dtype}
	s.arithmetics.Push(tmp)
}

// ExitSubArithmeticExpression is called when production SubArithmeticExpression is exited.
func (s *Listener) ExitSubArithmeticExpression(ctx *SubArithmeticExpressionContext) {
	second := s.arithmetics.Pop()
	first := s.arithmetics.Pop()
	if first.GetDataType() != second.GetDataType() {
		panic("DataType Mismatch")
	}
	function := "subf"
	if first.GetDataType() == kInt64 {
		function = "subi"
	}
	tmp := &Function{function: function,
		args: []ArithmeticExpression{first, second}}
	s.arithmetics.Push(tmp)
}

// ExitModArithmeticExpression is called when production ModArithmeticExpression is exited.
func (s *Listener) ExitModArithmeticExpression(ctx *ModArithmeticExpressionContext) {
	second := s.arithmetics.Pop()
	first := s.arithmetics.Pop()
	if first.GetDataType() != second.GetDataType() {
		panic("DataType Mismatch")
	}

	if first.GetDataType() != kInt64 {
		panic("DataType Must Be Int64")
	}
	if second.Trivial() {
		second = second.Simplify()
		if second.(*Int64).value == 0 {
			panic("Mod By Zero")
		}
	}

	tmp := &Function{function: "modi", args: []ArithmeticExpression{first, second}}
	s.arithmetics.Push(tmp)
}

// ExitRuntTimeFuncArithmeticExpression is called when production RuntTimeFuncArithmeticExpression is exited.
func (s *Listener) ExitRuntTimeFuncArithmeticExpression(ctx *RuntTimeFuncArithmeticExpressionContext) {
	function := ctx.IDENTIFIER().GetText()
	tmp := &Function{
		function: strings.ToLower(function),
		args:     []ArithmeticExpression{},
	}
	s.arithmetics.Push(tmp)
}

// ExitMulArithmeticExpression is called when production MulArithmeticExpression is exited.
func (s *Listener) ExitMulArithmeticExpression(ctx *MulArithmeticExpressionContext) {
	second := s.arithmetics.Pop()
	first := s.arithmetics.Pop()
	if first.GetDataType() != second.GetDataType() {
		panic("DataType Mismatch")
	}
	function := "mulf"
	if first.GetDataType() == kInt64 {
		function = "muli"
	}
	tmp := &Function{
		function: function,
		args:     []ArithmeticExpression{first, second},
	}
	s.arithmetics.Push(tmp)
}

func (s *Listener) parseDecimalList(str string) []float32 {
	str = strings.TrimSpace(str)
	str = str[1 : len(str)-1]
	items := strings.Split(str, ",")
	array := make([]float32, 0, len(items))
	for i := 0; i < len(items); i++ {
		tmp := strings.TrimSpace(items[i])
		val, _ := strconv.ParseFloat(tmp, 32)
		array = append(array, float32(val))
	}
	return array
}

func (s *Listener) parseIntegerList(str string) []int64 {
	str = strings.TrimSpace(str)
	str = str[1 : len(str)-1]
	items := strings.Split(str, ",")
	array := make([]int64, 0, len(items))
	for i := 0; i < len(items); i++ {
		tmp := strings.TrimSpace(items[i])
		val, _ := strconv.ParseInt(tmp, 10, 64)
		array = append(array, val)
	}
	return array
}

func (s *Listener) parseStringList(str string) []string {
	str = strings.TrimSpace(str)
	str = str[1 : len(str)-1]
	items := strings.Split(str, ",")
	array := make([]string, 0, len(items))
	for i := 0; i < len(items); i++ {
		tmp := strings.TrimSpace(items[i])
		array = append(array, tmp[1:len(tmp)-1])
	}
	return array
}
