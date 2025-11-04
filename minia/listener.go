package minia

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/uopensail/ulib/sample"
)

// ========== Expression Type System ==========

/**
 * @brief Expression type enumeration
 */
type ExprType int

const (
	ExprTypeLiteral  ExprType = iota // Literal value
	ExprTypeVariable                 // Function call or computed value
	ExprTypeColumn                   // Column reference
)

// ========== Expression Interface ==========

/**
 * @brief Base expression interface
 */
type Expr interface {
	GetType() ExprType
	GetName() string
	SetName(name string)
	String() string
	Clone() Expr
}

// ========== Literal Expression ==========

/**
 * @brief Literal expression containing constant values
 */
type Literal struct {
	Name  string
	Value sample.Feature
	Type  ExprType
}

/**
 * @brief Get expression type
 * @return ExprTypeLiteral
 */
func (l *Literal) GetType() ExprType {
	return ExprTypeLiteral
}

/**
 * @brief Get expression name
 * @return Expression name
 */
func (l *Literal) GetName() string {
	return l.Name
}

/**
 * @brief Set expression name
 * @param name New name
 */
func (l *Literal) SetName(name string) {
	l.Name = name
}

/**
 * @brief String representation
 * @return String representation of literal
 */
func (l *Literal) String() string {
	if l.Value == nil {
		return fmt.Sprintf("Literal{%s: <nil>}", l.Name)
	}
	return fmt.Sprintf("Literal{%s: %v}", l.Name, l.Value.Get())
}

/**
 * @brief Clone the literal expression
 * @return Cloned expression
 */
func (l *Literal) Clone() Expr {
	return &Literal{
		Name:  l.Name,
		Value: l.Value, // Note: shallow copy of Feature
		Type:  l.Type,
	}
}

/**
 * @brief Create new literal expression
 * @param name Expression name
 * @param feature Feature value
 * @return New literal expression
 */
func NewLiteral(name string, feature sample.Feature) *Literal {
	return &Literal{
		Name:  name,
		Value: feature,
		Type:  ExprTypeLiteral,
	}
}

// ========== Variable Expression ==========

/**
 * @brief Variable expression representing function calls or computed values
 */
type Variable struct {
	Name string
	Func string
	Args []Expr
	Type ExprType
}

/**
 * @brief Get expression type
 * @return ExprTypeVariable
 */
func (v *Variable) GetType() ExprType {
	return ExprTypeVariable
}

/**
 * @brief Get expression name
 * @return Expression name
 */
func (v *Variable) GetName() string {
	return v.Name
}

/**
 * @brief Set expression name
 * @param name New name
 */
func (v *Variable) SetName(name string) {
	v.Name = name
}

/**
 * @brief String representation
 * @return String representation of variable
 */
func (v *Variable) String() string {
	argStrs := make([]string, len(v.Args))
	for i, arg := range v.Args {
		argStrs[i] = arg.String()
	}
	return fmt.Sprintf("Variable{%s: %s(%s)}", v.Name, v.Func, strings.Join(argStrs, ", "))
}

/**
 * @brief Clone the variable expression
 * @return Cloned expression
 */
func (v *Variable) Clone() Expr {
	clonedArgs := make([]Expr, len(v.Args))
	for i, arg := range v.Args {
		clonedArgs[i] = arg.Clone()
	}
	return &Variable{
		Name: v.Name,
		Func: v.Func,
		Args: clonedArgs,
		Type: v.Type,
	}
}

/**
 * @brief Create new variable expression
 * @param name Expression name
 * @param funcName Function name
 * @param args Function arguments
 * @return New variable expression
 */
func NewVariable(name, funcName string, args []Expr) *Variable {
	return &Variable{
		Name: name,
		Func: funcName,
		Args: args,
		Type: ExprTypeVariable,
	}
}

// ========== Column Expression ==========

/**
 * @brief Column expression representing data column references
 */
type ColumnNode struct {
	Name string
	Type ExprType
}

/**
 * @brief Get expression type
 * @return ExprTypeColumn
 */
func (c *ColumnNode) GetType() ExprType {
	return ExprTypeColumn
}

/**
 * @brief Get expression name
 * @return Expression name
 */
func (c *ColumnNode) GetName() string {
	return c.Name
}

/**
 * @brief Set expression name
 * @param name New name
 */
func (c *ColumnNode) SetName(name string) {
	c.Name = name
}

/**
 * @brief String representation
 * @return String representation of column
 */
func (c *ColumnNode) String() string {
	return fmt.Sprintf("Column{%s}", c.Name)
}

/**
 * @brief Clone the column expression
 * @return Cloned expression
 */
func (c *ColumnNode) Clone() Expr {
	return &ColumnNode{
		Name: c.Name,
		Type: c.Type,
	}
}

/**
 * @brief Create new column expression
 * @param name Column name
 * @return New column expression
 */
func NewColumnNode(name string) *ColumnNode {
	return &ColumnNode{
		Name: name,
		Type: ExprTypeColumn,
	}
}

// ========== Expression Listener ==========

/**
 * @brief ANTLR listener for expression parsing and optimization
 */
type Listener struct {
	*BaseminiaListener
	*antlr.DefaultErrorListener

	exprs     *Stack[Expr]
	nodes     []Expr
	features  []string
	nodeCount int
}

/**
 * @brief Create new listener instance
 * @return New listener instance
 */
func NewListener() *Listener {
	return &Listener{
		BaseminiaListener:    &BaseminiaListener{},
		DefaultErrorListener: &antlr.DefaultErrorListener{},
		exprs:                NewStack[Expr](),
		nodes:                make([]Expr, 0),
		features:             make([]string, 0),
	}
}

// ========== ANTLR Exit Methods ==========

/**
 * @brief Handle start expression exit
 * @param ctx Start context
 */
func (l *Listener) ExitStart(ctx *StartContext) {
	if l.exprs.Empty() {
		panic("operation stack is empty at exitStart")
	}

	identifier := ctx.IDENTIFIER()
	if identifier == nil {
		panic("missing identifier in start expression")
	}

	str := identifier.GetText()
	expr := l.exprs.Pop()
	expr.SetName(str)
	l.nodes = append(l.nodes, expr)
	l.features = append(l.features, str)
}

/**
 * @brief Handle binary arithmetic operations
 * @param op Operation name
 */
func (l *Listener) handleBinaryOp(op string) {
	if l.exprs.Size() < 2 {
		panic(fmt.Sprintf("insufficient operands for %s operation", op))
	}

	right := l.exprs.Pop()
	left := l.exprs.Pop()

	ptr := NewVariable(
		fmt.Sprintf("node:%d", l.nodeCount),
		op,
		[]Expr{left, right},
	)
	l.nodeCount++

	expr := l.optimizeExpression(ptr)
	l.exprs.Push(expr)
	l.nodes = append(l.nodes, expr)
}

/**
 * @brief Handle unary operations
 * @param op Operation name
 */
func (l *Listener) handleUnaryOp(op string) {
	if l.exprs.Size() < 1 {
		panic(fmt.Sprintf("insufficient operands for %s operation", op))
	}

	operand := l.exprs.Pop()

	ptr := NewVariable(
		fmt.Sprintf("node:%d", l.nodeCount),
		op,
		[]Expr{operand},
	)
	l.nodeCount++

	expr := l.optimizeExpression(ptr)
	l.exprs.Push(expr)
	l.nodes = append(l.nodes, expr)
}

// Arithmetic operations
func (l *Listener) ExitMulExpr(ctx *MulExprContext) { l.handleBinaryOp("mul") }
func (l *Listener) ExitSubExpr(ctx *SubExprContext) { l.handleBinaryOp("sub") }
func (l *Listener) ExitAddExpr(ctx *AddExprContext) { l.handleBinaryOp("add") }
func (l *Listener) ExitDivExpr(ctx *DivExprContext) { l.handleBinaryOp("div") }
func (l *Listener) ExitModExpr(ctx *ModExprContext) { l.handleBinaryOp("mod") }

// Logical operations
func (l *Listener) ExitAndExpr(ctx *AndExprContext) { l.handleBinaryOp("and") }
func (l *Listener) ExitOrExpr(ctx *OrExprContext)   { l.handleBinaryOp("or") }
func (l *Listener) ExitNotExpr(ctx *NotExprContext) { l.handleUnaryOp("not") }

// Comparison operations
func (l *Listener) ExitGreaterThanEqualExpr(ctx *GreaterThanEqualExprContext) {
	l.handleBinaryOp("gte")
}
func (l *Listener) ExitLessThanEqualExpr(ctx *LessThanEqualExprContext) { l.handleBinaryOp("lte") }
func (l *Listener) ExitLessThanExpr(ctx *LessThanExprContext)           { l.handleBinaryOp("lt") }
func (l *Listener) ExitGreaterThanExpr(ctx *GreaterThanExprContext)     { l.handleBinaryOp("gt") }
func (l *Listener) ExitNotEqualExpr(ctx *NotEqualExprContext)           { l.handleBinaryOp("neq") }
func (l *Listener) ExitEqualExpr(ctx *EqualExprContext)                 { l.handleBinaryOp("eq") }

/**
 * @brief Handle negation expression
 * @param ctx Negation context
 */
func (l *Listener) ExitNegExpr(ctx *NegExprContext) {
	if l.exprs.Size() < 1 {
		panic("insufficient operands for negation")
	}

	expr := l.exprs.Pop()

	// Create zero literal
	zeroFeature := &sample.Int64{Value: 0}
	zeroLiteral := NewLiteral(fmt.Sprintf("node:%d", l.nodeCount), zeroFeature)
	l.nodeCount++

	ptr := NewVariable(
		fmt.Sprintf("node:%d", l.nodeCount),
		"sub",
		[]Expr{zeroLiteral, expr},
	)
	l.nodeCount++

	optimizedExpr := l.optimizeExpression(ptr)
	l.exprs.Push(optimizedExpr)
	l.nodes = append(l.nodes, optimizedExpr)
}

/**
 * @brief Handle function call expression
 * @param ctx Function call context
 */
func (l *Listener) ExitFuncCall(ctx *FuncCallContext) {
	identifier := ctx.IDENTIFIER()
	if identifier == nil {
		panic("missing function name in function call")
	}

	funcName := identifier.GetText()

	var argCount int
	if ctx.ExprList() != nil {
		if exprList, ok := ctx.ExprList().(*ExprListContext); ok {
			argCount = len(exprList.AllExpr())
		}
	}

	if l.exprs.Size() < argCount {
		panic(fmt.Sprintf("insufficient arguments for function call %s: expected %d, got %d",
			funcName, argCount, l.exprs.Size()))
	}

	args := make([]Expr, argCount)
	for i := argCount - 1; i >= 0; i-- {
		args[i] = l.exprs.Pop()
	}

	ptr := NewVariable(
		fmt.Sprintf("node:%d", l.nodeCount),
		funcName,
		args,
	)
	l.nodeCount++

	expr := l.optimizeExpression(ptr)
	l.exprs.Push(expr)
	l.nodes = append(l.nodes, expr)
}

// ========== Literal Value Handlers ==========

/**
 * @brief Handle string literal
 * @param ctx String context
 */
func (l *Listener) ExitStringExpr(ctx *StringExprContext) {
	str := ctx.STRING().GetText()
	if len(str) < 2 {
		panic("invalid string format")
	}

	// Remove quotes
	value := str[1 : len(str)-1]
	f := &sample.String{Value: value}

	ptr := NewLiteral(fmt.Sprintf("node:%d", l.nodeCount), f)
	l.exprs.Push(ptr)
	l.nodeCount++
	l.nodes = append(l.nodes, ptr)
}

/**
 * @brief Handle integer literal
 * @param ctx Integer context
 */
func (l *Listener) ExitIntegerExpr(ctx *IntegerExprContext) {
	str := ctx.INTEGER().GetText()

	value, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("invalid integer format: %s", str))
	}

	f := &sample.Int64{Value: value}
	ptr := NewLiteral(fmt.Sprintf("node:%d", l.nodeCount), f)
	l.exprs.Push(ptr)
	l.nodes = append(l.nodes, ptr)
	l.nodeCount++
}

/**
 * @brief Handle decimal literal
 * @param ctx Decimal context
 */
func (l *Listener) ExitDecimalExpr(ctx *DecimalExprContext) {
	str := ctx.GetText()

	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(fmt.Sprintf("invalid decimal format: %s", str))
	}

	f := &sample.Float32{Value: float32(value)}
	ptr := NewLiteral(fmt.Sprintf("node:%d", l.nodeCount), f)
	l.exprs.Push(ptr)
	l.nodes = append(l.nodes, ptr)
	l.nodeCount++
}

/**
 * @brief Handle boolean literals
 */
func (l *Listener) ExitTrueExpr(ctx *TrueExprContext) {
	f := &sample.Int64{Value: 1}
	ptr := NewLiteral(fmt.Sprintf("node:%d", l.nodeCount), f)
	l.exprs.Push(ptr)
	l.nodeCount++
	l.nodes = append(l.nodes, ptr)
}

func (l *Listener) ExitFalseExpr(ctx *FalseExprContext) {
	f := &sample.Int64{Value: 0}
	ptr := NewLiteral(fmt.Sprintf("node:%d", l.nodeCount), f)
	l.exprs.Push(ptr)
	l.nodeCount++
	l.nodes = append(l.nodes, ptr)
}

/**
 * @brief Handle column reference
 * @param ctx Column context
 */
func (l *Listener) ExitColumnExpr(ctx *ColumnExprContext) {
	identifier := ctx.IDENTIFIER()
	if identifier == nil {
		panic("missing identifier in column expression")
	}

	str := identifier.GetText()
	ptr := NewColumnNode(str)
	l.exprs.Push(ptr)
	l.nodes = append(l.nodes, ptr)
}

// ========== List Handlers ==========

/**
 * @brief Handle integer list literal
 * @param ctx Integer list context
 */
func (l *Listener) ExitIntegerListExpr(ctx *IntegerListExprContext) {
	str := ctx.GetText()
	result, err := l.parseIntList(str)
	if err != nil {
		panic(fmt.Sprintf("failed to parse integer list: %v", err))
	}

	f := &sample.Int64s{Value: result}
	ptr := NewLiteral(fmt.Sprintf("node:%d", l.nodeCount), f)
	l.exprs.Push(ptr)
	l.nodeCount++
	l.nodes = append(l.nodes, ptr)
}

/**
 * @brief Handle decimal list literal
 * @param ctx Decimal list context
 */
func (l *Listener) ExitDecimalListExpr(ctx *DecimalListExprContext) {
	str := ctx.GetText()
	result, err := l.parseFloatList(str)
	if err != nil {
		panic(fmt.Sprintf("failed to parse decimal list: %v", err))
	}

	values := make([]float32, len(result))
	for i, v := range result {
		values[i] = float32(v)
	}

	f := &sample.Float32s{Value: values}
	ptr := NewLiteral(fmt.Sprintf("node:%d", l.nodeCount), f)
	l.exprs.Push(ptr)
	l.nodeCount++
	l.nodes = append(l.nodes, ptr)
}

/**
 * @brief Handle string list literal
 * @param ctx String list context
 */
func (l *Listener) ExitStringListExpr(ctx *StringListExprContext) {
	str := ctx.GetText()
	result, err := l.parseStringList(str)
	if err != nil {
		panic(fmt.Sprintf("failed to parse string list: %v", err))
	}

	f := &sample.Strings{Value: result}
	ptr := NewLiteral(fmt.Sprintf("node:%d", l.nodeCount), f)
	l.exprs.Push(ptr)
	l.nodes = append(l.nodes, ptr)
	l.nodeCount++
}

// ========== Expression Optimization ==========

/**
 * @brief Optimize expression with various optimization techniques
 * @param expr Expression to optimize
 * @return Optimized expression
 */
func (l *Listener) optimizeExpression(expr Expr) Expr {
	variable, ok := expr.(*Variable)
	if !ok {
		return expr
	}

	// Apply optimization passes
	optimized := l.applyConstantFolding(variable)
	optimized = l.applyLogicalOptimization(optimized)

	return optimized
}

/**
 * @brief Apply constant folding optimization
 * @param expr Variable expression to optimize
 * @return Optimized expression
 */
func (l *Listener) applyConstantFolding(expr *Variable) Expr {
	// Check if all arguments are literals
	allLiteral := true
	for _, arg := range expr.Args {
		if arg.GetType() != ExprTypeLiteral {
			allLiteral = false
			break
		}
	}

	if !allLiteral {
		return expr
	}

	// Try to evaluate the function with literal arguments
	return l.evaluateBuiltinFunction(expr)
}

/**
 * @brief Apply logical optimization (short-circuiting)
 * @param expr Expression to optimize
 * @return Optimized expression
 */
func (l *Listener) applyLogicalOptimization(expr Expr) Expr {
	variable, ok := expr.(*Variable)
	if !ok {
		return expr
	}

	switch variable.Func {
	case "and":
		return l.optimizeAndExpression(variable)
	case "or":
		return l.optimizeOrExpression(variable)
	case "not":
		return l.optimizeNotExpression(variable)
	default:
		return variable
	}
}

/**
 * @brief Optimize AND expression
 * @param expr AND expression
 * @return Optimized expression
 */
func (l *Listener) optimizeAndExpression(expr *Variable) Expr {
	if len(expr.Args) != 2 {
		return expr
	}

	left, leftIsLit := expr.Args[0].(*Literal)
	right, rightIsLit := expr.Args[1].(*Literal)

	// Short-circuit evaluation
	if leftIsLit {
		if l.getIntValue(left) == 0 {
			// false && anything = false
			return l.createIntLiteral(expr.Name, 0)
		} else {
			// true && right = right
			return expr.Args[1]
		}
	}

	if rightIsLit {
		if l.getIntValue(right) == 0 {
			// anything && false = false
			return l.createIntLiteral(expr.Name, 0)
		} else {
			// left && true = left
			return expr.Args[0]
		}
	}

	return expr
}

/**
 * @brief Optimize OR expression
 * @param expr OR expression
 * @return Optimized expression
 */
func (l *Listener) optimizeOrExpression(expr *Variable) Expr {
	if len(expr.Args) != 2 {
		return expr
	}

	left, leftIsLit := expr.Args[0].(*Literal)
	right, rightIsLit := expr.Args[1].(*Literal)

	// Short-circuit evaluation
	if leftIsLit {
		if l.getIntValue(left) != 0 {
			// true || anything = true
			return l.createIntLiteral(expr.Name, 1)
		} else {
			// false || right = right
			return expr.Args[1]
		}
	}

	if rightIsLit {
		if l.getIntValue(right) != 0 {
			// anything || true = true
			return l.createIntLiteral(expr.Name, 1)
		} else {
			// left || false = left
			return expr.Args[0]
		}
	}

	return expr
}

/**
 * @brief Optimize NOT expression
 * @param expr NOT expression
 * @return Optimized expression
 */
func (l *Listener) optimizeNotExpression(expr *Variable) Expr {
	if len(expr.Args) != 1 {
		return expr
	}

	if lit, ok := expr.Args[0].(*Literal); ok {
		value := int64(0)
		if l.getIntValue(lit) == 0 {
			value = 1
		}
		return l.createIntLiteral(expr.Name, value)
	}

	return expr
}

// ========== Helper Functions ==========

/**
 * @brief Evaluate builtin function with literal arguments
 * @param expr Variable expression with literal arguments
 * @return Evaluated literal or original expression if evaluation fails
 */
func (l *Listener) evaluateBuiltinFunction(expr *Variable) Expr {
	// Build function signature
	args := make([]sample.Feature, len(expr.Args))
	typeCodes := make([]string, len(expr.Args))

	for i, arg := range expr.Args {
		if lit, ok := arg.(*Literal); ok {
			if lit.Value == nil {
				// Cannot evaluate with nil value
				return expr
			}
			typeCodes = append(typeCodes, fmt.Sprintf("%d", lit.Value.Type()))
			args[i] = lit.Value
		} else {
			// Cannot evaluate with non-literal arguments
			return expr
		}
	}
	signature := fmt.Sprintf("%s:%d=[%s]", expr.Func, len(expr.Args), strings.Join(typeCodes, ","))

	// Try to find and call the function
	if f := GetFunction(signature); f != nil {
		result, err := f.Call(args...)
		if err != nil {
			// If evaluation fails, return original expression
			// Don't panic here to allow graceful degradation
			return expr
		}
		return NewLiteral(expr.Name, result)
	}

	// Function not found, return original expression
	return expr
}

/**
 * @brief Get integer value from literal
 * @param lit Literal expression
 * @return Integer value, 0 if conversion fails
 */
func (l *Listener) getIntValue(lit *Literal) int64 {
	if lit.Value == nil {
		return 0
	}

	if v, err := lit.Value.GetInt64(); err == nil {
		return v
	}

	// Try to get as float and convert
	if v, err := lit.Value.GetFloat32(); err == nil {
		return int64(v)
	}

	return 0
}

/**
 * @brief Create integer literal expression
 * @param name Expression name
 * @param value Integer value
 * @return New literal expression
 */
func (l *Listener) createIntLiteral(name string, value int64) *Literal {
	f := &sample.Int64{Value: value}
	return NewLiteral(name, f)
}

// ========== List Parsing Functions ==========

/**
 * @brief Parse integer list from string representation
 * @param str String representation like "[1,2,3]"
 * @return Parsed integer slice and error
 */
func (l *Listener) parseIntList(str string) ([]int64, error) {
	if len(str) < 2 || str[0] != '[' || str[len(str)-1] != ']' {
		return nil, fmt.Errorf("invalid list format: %s", str)
	}

	content := strings.TrimSpace(str[1 : len(str)-1])
	if content == "" {
		return []int64{}, nil
	}

	parts := strings.Split(content, ",")
	result := make([]int64, 0, len(parts))

	for i, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			continue
		}

		value, err := strconv.ParseInt(trimmed, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid integer at position %d: %s", i, trimmed)
		}
		result = append(result, value)
	}

	return result, nil
}

/**
 * @brief Parse float list from string representation
 * @param str String representation like "[1.0,2.5,3.14]"
 * @return Parsed float slice and error
 */
func (l *Listener) parseFloatList(str string) ([]float64, error) {
	if len(str) < 2 || str[0] != '[' || str[len(str)-1] != ']' {
		return nil, fmt.Errorf("invalid list format: %s", str)
	}

	content := strings.TrimSpace(str[1 : len(str)-1])
	if content == "" {
		return []float64{}, nil
	}

	parts := strings.Split(content, ",")
	result := make([]float64, 0, len(parts))

	for i, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			continue
		}

		value, err := strconv.ParseFloat(trimmed, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid float at position %d: %s", i, trimmed)
		}
		result = append(result, value)
	}

	return result, nil
}

/**
 * @brief Parse string list from string representation
 * @param str String representation like '["hello","world"]'
 * @return Parsed string slice and error
 */
func (l *Listener) parseStringList(str string) ([]string, error) {
	if len(str) < 2 || str[0] != '[' || str[len(str)-1] != ']' {
		return nil, fmt.Errorf("invalid list format: %s", str)
	}

	content := strings.TrimSpace(str[1 : len(str)-1])
	if content == "" {
		return []string{}, nil
	}

	result := make([]string, 0)
	current := ""
	inQuotes := false
	escaped := false

	for i, ch := range content {
		if escaped {
			current += string(ch)
			escaped = false
			continue
		}

		switch ch {
		case '\\':
			if inQuotes {
				escaped = true
			} else {
				current += string(ch)
			}
		case '"':
			inQuotes = !inQuotes
			current += string(ch)
		case ',':
			if !inQuotes {
				trimmed := strings.TrimSpace(current)
				if len(trimmed) >= 2 && trimmed[0] == '"' && trimmed[len(trimmed)-1] == '"' {
					// Remove quotes and handle escape sequences
					unquoted := trimmed[1 : len(trimmed)-1]
					unquoted = strings.ReplaceAll(unquoted, `\"`, `"`)
					unquoted = strings.ReplaceAll(unquoted, `\\`, `\`)
					result = append(result, unquoted)
				} else if trimmed != "" {
					return nil, fmt.Errorf("invalid string format at position %d: %s", i, trimmed)
				}
				current = ""
			} else {
				current += string(ch)
			}
		default:
			current += string(ch)
		}
	}

	// Handle the last element
	if current != "" {
		trimmed := strings.TrimSpace(current)
		if len(trimmed) >= 2 && trimmed[0] == '"' && trimmed[len(trimmed)-1] == '"' {
			unquoted := trimmed[1 : len(trimmed)-1]
			unquoted = strings.ReplaceAll(unquoted, `\"`, `"`)
			unquoted = strings.ReplaceAll(unquoted, `\\`, `\`)
			result = append(result, unquoted)
		} else if trimmed != "" {
			return nil, fmt.Errorf("invalid string format: %s", trimmed)
		}
	}

	return result, nil
}

// ========== Error Handling ==========

/**
 * @brief Handle syntax errors from ANTLR
 * @param recognizer ANTLR recognizer
 * @param offendingSymbol Offending symbol
 * @param line Line number
 * @param column Column number
 * @param msg Error message
 * @param e Exception
 */
func (l *Listener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	panic(fmt.Sprintf("syntax error at line %d:%d - %s", line, column, msg))
}
