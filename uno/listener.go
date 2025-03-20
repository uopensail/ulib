package uno

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr"
	"github.com/uopensail/ulib/sample"
)

type Listener struct {
	*BaseunoListener
	*antlr.DefaultErrorListener
	arithmetics *Stack[ArithmeticExpression]
	booleans    *Stack[BooleanExpression]
	types       map[string]sample.DataType
}

func NewListener(types map[string]sample.DataType) *Listener {
	return &Listener{
		arithmetics: NewStack[ArithmeticExpression](),
		booleans:    NewStack[BooleanExpression](),
		types:       types,
	}
}

func (s *Listener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	panic(fmt.Sprintf("line: %d column: %d: %s", line, column, msg))
}

// ExitCmpBooleanExpression is called when production CmpBooleanExpression is exited.
func (s *Listener) ExitCmpBooleanExpression(ctx *CmpBooleanExpressionContext) {
	right := s.arithmetics.Pop()
	left := s.arithmetics.Pop()

	// support integer cast to float
	var leftExpr ArithmeticExpression
	var rightExpr ArithmeticExpression

	leftType := left.GetDataType()
	rightType := right.GetDataType()
	dtype := sample.ErrorType
	if leftType == sample.Float32Type && rightType == sample.Float32Type {
		leftExpr, rightExpr, dtype = left, right, sample.Float32Type
	} else if leftType == sample.Float32Type && rightType == sample.Int64Type {
		leftExpr, rightExpr, dtype = left, &Function{
			function: "casti2f",
			args:     []ArithmeticExpression{right},
		}, sample.Float32Type
	} else if leftType == sample.Int64Type && rightType == sample.Int64Type {
		leftExpr, rightExpr = left, right
	} else if leftType == sample.Int64Type && rightType == sample.Float32Type {
		leftExpr, rightExpr, dtype = &Function{
			function: "casti2f",
			args:     []ArithmeticExpression{left},
		}, right, sample.Float32Type
	} else if leftType == sample.StringType && rightType == sample.StringType {
		leftExpr, rightExpr, dtype = left, right, sample.StringType
	} else {
		panic("DataType Mismatch")
	}

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
	} else {
		panic(fmt.Sprintf("%s not a valid operation", mark))
	}
	tmp := &Cmp{left: leftExpr, right: rightExpr, op: op, dtype: dtype}
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
	exprType := expr.GetDataType()
	if ctx.INTEGER_LIST() != nil {
		if exprType != sample.Int64Type && exprType != sample.Int64sType {
			panic("DataType Mismatch")
		}
		list := s.parseIntegerList(ctx.INTEGER_LIST().GetText())
		array := &Int64s{value: list}
		s.booleans.Push(&NotIn{left: expr, right: array, dtype: sample.Int64Type})
	} else {
		if exprType != sample.StringType && exprType != sample.StringsType {
			panic("DataType Mismatch")
		}
		list := s.parseStringList(ctx.STRING_LIST().GetText())
		array := &Strings{value: list}
		s.booleans.Push(&NotIn{left: expr, right: array, dtype: sample.StringType})
	}
}

// ExitFalseBooleanExpression is called when production FalseBooleanExpression is exited.
func (s *Listener) ExitFalseBooleanExpression(ctx *FalseBooleanExpressionContext) {
	s.booleans.Push(&Literal{value: false})
}

// ExitInBooleanExpression is called when production InBooleanExpression is exited.
func (s *Listener) ExitInBooleanExpression(ctx *InBooleanExpressionContext) {
	expr := s.arithmetics.Pop()
	exprType := expr.GetDataType()
	if ctx.INTEGER_LIST() != nil {
		if exprType != sample.Int64Type && exprType != sample.Int64sType {
			panic("DataType Mismatch")
		}
		list := s.parseIntegerList(ctx.INTEGER_LIST().GetText())
		array := &Int64s{value: list}
		s.booleans.Push(&In{left: expr, right: array, dtype: sample.Int64Type})
	} else {
		if exprType != sample.StringType && exprType != sample.StringsType {
			panic("DataType Mismatch")
		}
		list := s.parseStringList(ctx.STRING_LIST().GetText())
		array := &Strings{value: list}
		s.booleans.Push(&In{left: expr, right: array, dtype: sample.StringType})
	}
}

// ExitAddArithmeticExpression is called when production AddArithmeticExpression is exited.
func (s *Listener) ExitAddArithmeticExpression(ctx *AddArithmeticExpressionContext) {
	second := s.arithmetics.Pop()
	first := s.arithmetics.Pop()
	function := "addf"
	firstType := first.GetDataType()
	secondType := second.GetDataType()

	var firstExpr ArithmeticExpression
	var secondExpr ArithmeticExpression
	if firstType == sample.Int64Type && secondType == sample.Int64Type {
		function, firstExpr, secondExpr = "addi", first, second
	} else if firstType == sample.Int64Type && secondType == sample.Float32Type {
		function, firstExpr, secondExpr = "addf", &Function{
			function: "casti2f",
			args:     []ArithmeticExpression{first},
		}, second
	} else if firstType == sample.Float32Type && secondType == sample.Int64Type {
		function, firstExpr, secondExpr = "addf", first, &Function{
			function: "casti2f",
			args:     []ArithmeticExpression{second},
		}
	} else if firstType == sample.Float32Type && secondType == sample.Float32Type {
		function, firstExpr, secondExpr = "addf", first, second
	} else {
		panic("DataType Not supported")
	}

	tmp := &Function{
		function: function,
		args:     []ArithmeticExpression{firstExpr, secondExpr},
	}
	tmp.check()
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
	tmp.check()
	s.arithmetics.Push(tmp)
}

// ExitColumnArithmeticExpression is called when production ColumnArithmeticExpression is exited.
func (s *Listener) ExitColumnArithmeticExpression(ctx *ColumnArithmeticExpressionContext) {
	column := ctx.IDENTIFIER().GetText()
	var dtype sample.DataType
	var ok bool
	if dtype, ok = s.types[column]; !ok {
		panic(fmt.Sprintf("column:%s not found", column))
	}

	tmp := &Variable{value: column, dtype: dtype}
	s.arithmetics.Push(tmp)
}

// ExitDivArithmeticExpression is called when production DivArithmeticExpression is exited.
func (s *Listener) ExitDivArithmeticExpression(ctx *DivArithmeticExpressionContext) {
	second := s.arithmetics.Pop()
	first := s.arithmetics.Pop()
	function := "divf"
	firstType := first.GetDataType()
	secondType := second.GetDataType()

	var firstExpr ArithmeticExpression
	var secondExpr ArithmeticExpression
	if firstType == sample.Int64Type && secondType == sample.Int64Type {
		function, firstExpr, secondExpr = "divi", first, second
	} else if firstType == sample.Int64Type && secondType == sample.Float32Type {
		function, firstExpr, secondExpr = "divf", &Function{
			function: "casti2f",
			args:     []ArithmeticExpression{first},
		}, second
	} else if firstType == sample.Float32Type && secondType == sample.Int64Type {
		function, firstExpr, secondExpr = "divf", first, &Function{
			function: "casti2f",
			args:     []ArithmeticExpression{second},
		}
	} else if firstType == sample.Float32Type && secondType == sample.Float32Type {
		function, firstExpr, secondExpr = "divf", first, second
	} else {
		panic("DataType Not supported")
	}

	tmp := &Function{
		function: function,
		args:     []ArithmeticExpression{firstExpr, secondExpr},
	}
	tmp.check()
	s.arithmetics.Push(tmp)
}

// ExitSubArithmeticExpression is called when production SubArithmeticExpression is exited.
func (s *Listener) ExitSubArithmeticExpression(ctx *SubArithmeticExpressionContext) {
	second := s.arithmetics.Pop()
	first := s.arithmetics.Pop()
	function := "subf"
	firstType := first.GetDataType()
	secondType := second.GetDataType()

	var firstExpr ArithmeticExpression
	var secondExpr ArithmeticExpression
	if firstType == sample.Int64Type && secondType == sample.Int64Type {
		function, firstExpr, secondExpr = "subi", first, second
	} else if firstType == sample.Int64Type && secondType == sample.Float32Type {
		function, firstExpr, secondExpr = "subf", &Function{
			function: "casti2f",
			args:     []ArithmeticExpression{first},
		}, second
	} else if firstType == sample.Float32Type && secondType == sample.Int64Type {
		function, firstExpr, secondExpr = "subf", first, &Function{
			function: "casti2f",
			args:     []ArithmeticExpression{second},
		}
	} else if firstType == sample.Float32Type && secondType == sample.Float32Type {
		function, firstExpr, secondExpr = "subf", first, second
	} else {
		panic("DataType Not supported")
	}

	tmp := &Function{
		function: function,
		args:     []ArithmeticExpression{firstExpr, secondExpr},
	}
	tmp.check()
	s.arithmetics.Push(tmp)
}

// ExitModArithmeticExpression is called when production ModArithmeticExpression is exited.
func (s *Listener) ExitModArithmeticExpression(ctx *ModArithmeticExpressionContext) {
	second := s.arithmetics.Pop()
	first := s.arithmetics.Pop()
	if first.GetDataType() != second.GetDataType() {
		panic("DataType Mismatch")
	}

	if first.GetDataType() != sample.Int64Type {
		panic("DataType Must Be Int64")
	}
	if second.Trivial() {
		second = second.Simplify()
		if second.(*Int64).value == 0 {
			panic("Mod By Zero")
		}
	}

	tmp := &Function{function: "mod", args: []ArithmeticExpression{first, second}}
	tmp.check()
	s.arithmetics.Push(tmp)
}

// ExitRuntTimeFuncArithmeticExpression is called when production RuntTimeFuncArithmeticExpression is exited.
func (s *Listener) ExitRuntTimeFuncArithmeticExpression(ctx *RuntTimeFuncArithmeticExpressionContext) {
	function := ctx.IDENTIFIER().GetText()
	tmp := &Function{
		function: strings.ToLower(function),
		args:     []ArithmeticExpression{},
	}
	tmp.check()
	s.arithmetics.Push(tmp)
}

// ExitMulArithmeticExpression is called when production MulArithmeticExpression is exited.
func (s *Listener) ExitMulArithmeticExpression(ctx *MulArithmeticExpressionContext) {
	second := s.arithmetics.Pop()
	first := s.arithmetics.Pop()
	function := "mulf"
	firstType := first.GetDataType()
	secondType := second.GetDataType()

	var firstExpr ArithmeticExpression
	var secondExpr ArithmeticExpression
	if firstType == sample.Int64Type && secondType == sample.Int64Type {
		function, firstExpr, secondExpr = "muli", first, second
	} else if firstType == sample.Int64Type && secondType == sample.Float32Type {
		function, firstExpr, secondExpr = "mulf", &Function{
			function: "casti2f",
			args:     []ArithmeticExpression{first},
		}, second
	} else if firstType == sample.Float32Type && secondType == sample.Int64Type {
		function, firstExpr, secondExpr = "mulf", first, &Function{
			function: "casti2f",
			args:     []ArithmeticExpression{second},
		}
	} else if firstType == sample.Float32Type && secondType == sample.Float32Type {
		function, firstExpr, secondExpr = "mulf", first, second
	} else {
		panic("DataType Not supported")
	}

	tmp := &Function{
		function: function,
		args:     []ArithmeticExpression{firstExpr, secondExpr},
	}
	tmp.check()
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
