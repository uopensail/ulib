package minia

import (
	"fmt"
	"strings"
	"sync"

	"github.com/antlr4-go/antlr/v4"
	"github.com/uopensail/ulib/sample"
)

// ========== Global Registry ==========

/**
 * @brief Global function registry with thread-safe initialization
 */
var (
	builtins     = make(map[string]FunctionWrapper)
	builtinsOnce sync.Once
)

// ========== Column Handler ==========

/**
 * @brief Column reference handler for data extraction
 */
type Column struct {
	Index      int    // Index in the evaluation array
	ColumnName string // Column name to lookup
}

/**
 * @brief Create new column handler
 * @param index Array index
 * @param columnName Column name
 * @return New column handler
 */
func NewColumn(index int, columnName string) *Column {
	return &Column{
		Index:      index,
		ColumnName: columnName,
	}
}

/**
 * @brief Extract column value from features and store in array
 * @param arr Evaluation array
 * @param features Variable number of feature sets
 */
func (c *Column) Extract(arr []sample.Feature, features ...sample.Features) {
	if c.Index < 0 || c.Index >= len(arr) {
		panic(fmt.Sprintf("column index out of bounds: %d", c.Index))
	}

	for _, fs := range features {
		if fs == nil {
			continue
		}

		feature := fs.Get(c.ColumnName)
		if feature != nil {
			arr[c.Index] = feature
			return // Use first non-nil feature found
		}
	}

	// If no feature found, this might be an error depending on use case
	// For now, we leave the array slot unchanged (could be a literal default)
}

/**
 * @brief String representation for debugging
 * @return String representation
 */
func (c *Column) String() string {
	return fmt.Sprintf("Column{%s -> [%d]}", c.ColumnName, c.Index)
}

// ========== Function Caller ==========

/**
 * @brief Function caller for expression evaluation
 */
type Caller struct {
	Index        int    // Index where result should be stored
	FunctionName string // Name of function to call
	Arguments    []int  // Indices of arguments in evaluation array
}

/**
 * @brief Create new function caller
 * @param index Result index
 * @param funcName Function name
 * @param args Argument indices
 * @return New caller
 */
func NewCaller(index int, funcName string, args []int) *Caller {
	return &Caller{
		Index:        index,
		FunctionName: funcName,
		Arguments:    args,
	}
}

/**
 * @brief Execute function call with arguments from array
 * @param arr Evaluation array
 * @throws panic if function not found or execution fails
 */
func (c *Caller) Execute(arr []sample.Feature) {
	if c.Index < 0 || c.Index >= len(arr) {
		panic(fmt.Sprintf("caller index out of bounds: %d", c.Index))
	}

	// Validate argument indices
	for i, argIdx := range c.Arguments {
		if argIdx < 0 || argIdx >= len(arr) {
			panic(fmt.Sprintf("argument %d index out of bounds: %d", i, argIdx))
		}
		if arr[argIdx] == nil {
			panic(fmt.Sprintf("argument %d at index %d is nil", i, argIdx))
		}
	}

	// Build function signature if not cached
	signature := c.buildSignature(arr)
	// Get function from registry
	function := GetFunction(signature)
	if function == nil {
		panic(fmt.Errorf("function not found: %s", signature))
	}

	// Prepare arguments
	inputs := make([]sample.Feature, len(c.Arguments))
	for i, argIdx := range c.Arguments {
		inputs[i] = arr[argIdx]
	}

	// Execute function
	result, err := function.Call(inputs...)
	if err != nil {
		panic(fmt.Errorf("function %s execution failed: %w", signature, err))
	}

	// Store result
	arr[c.Index] = result
}

/**
 * @brief Build and cache function signature
 * @param arr Evaluation array for type information
 */
func (c *Caller) buildSignature(arr []sample.Feature) string {
	types := make([]string, len(c.Arguments))
	for i, argIdx := range c.Arguments {
		types[i] = fmt.Sprintf("%d", arr[argIdx].Type())
	}
	return fmt.Sprintf("%s:%d=[%s]", c.FunctionName, len(c.Arguments), strings.Join(types, ","))
}

/**
 * @brief String representation for debugging
 * @return String representation
 */
func (c *Caller) String() string {
	return fmt.Sprintf("Caller{%s(%v) -> [%d]}", c.FunctionName, c.Arguments, c.Index)
}

// ========== Main Engine ==========

/**
 * @brief High-performance expression evaluation engine
 */
type Minia struct {
	// Output configuration
	outputNames   []string // Names of output features
	outputIndices []int    // Indices of output values in evaluation array

	// Execution plan
	columns []*Column // Column extractors (executed first)
	callers []*Caller // Function callers (executed in dependency order)

	// Evaluation state
	literals  []sample.Feature // Pre-computed literal values
	arraySize int              // Size of evaluation array

	pool chan []sample.Feature // Pool of reusable arrays
}

/**
 * @brief Create new Minia engine from expression strings
 * @param exprs Array of expression strings
 * @return New Minia engine instance
 */
func NewMinia(exprs []string) *Minia {
	if len(exprs) == 0 {
		panic("no expressions provided")
	}

	// Parse expressions
	input := strings.Join(exprs, ";")
	stream := antlr.NewInputStream(input)
	lexer := NewminiaLexer(stream)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewminiaParser(tokens)

	// Add error handling
	parser.RemoveErrorListeners()
	errorListener := &MiniaErrorListener{}
	parser.AddErrorListener(errorListener)

	listener := NewListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Prog())

	return buildMiniaEngine(listener)
}

/**
 * @brief Custom error listener for better error reporting
 */
type MiniaErrorListener struct {
	*antlr.DefaultErrorListener
}

/**
 * @brief Handle syntax errors
 */
func (mel *MiniaErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	panic(fmt.Sprintf("syntax error at line %d:%d - %s", line, column, msg))
}

/**
 * @brief Build Minia engine from parsed expressions
 * @param listener Expression listener with parsed nodes
 * @return Configured Minia engine
 */
func buildMiniaEngine(listener *Listener) *Minia {
	nodes := listener.nodes
	features := listener.features

	if len(nodes) == 0 {
		panic("no nodes generated from expressions")
	}

	// Initialize collections
	literals := make([]sample.Feature, len(nodes))
	columns := make([]*Column, 0)
	callers := make([]*Caller, 0)
	nameToIndex := make(map[string]int, len(nodes))

	// Process nodes in order
	for i, node := range nodes {
		switch node.GetType() {
		case ExprTypeLiteral:
			literal := node.(*Literal)
			literals[i] = literal.Value
			nameToIndex[node.GetName()] = i

		case ExprTypeColumn:
			column := node.(*ColumnNode)
			if _, exists := nameToIndex[node.GetName()]; !exists {
				columns = append(columns, NewColumn(i, column.GetName()))
				nameToIndex[node.GetName()] = i
			}

		case ExprTypeVariable:
			variable := node.(*Variable)

			// Resolve argument indices
			args := make([]int, len(variable.Args))
			for j, arg := range variable.Args {
				if idx, exists := nameToIndex[arg.GetName()]; exists {
					args[j] = idx
				} else {
					panic(fmt.Sprintf("undefined reference: %s", arg.GetName()))
				}
			}

			callers = append(callers, NewCaller(i, variable.Func, args))
			nameToIndex[node.GetName()] = i

		default:
			panic(fmt.Sprintf("unknown expression type: %v", node.GetType()))
		}
	}

	// Build output mapping
	outputIndices := make([]int, len(features))
	for i, featureName := range features {
		if idx, exists := nameToIndex[featureName]; exists {
			outputIndices[i] = idx
		} else {
			panic(fmt.Sprintf("output feature not found: %s", featureName))
		}
	}

	// Create engine
	engine := &Minia{
		outputNames:   features,
		outputIndices: outputIndices,
		columns:       columns,
		callers:       callers,
		literals:      literals,
		arraySize:     len(nodes),
		pool:          make(chan []sample.Feature, 16),
	}

	return engine
}

func (m *Minia) get() []sample.Feature {
	select {
	case arr := <-m.pool:
		return arr
	default:
		arr := make([]sample.Feature, m.arraySize)
		copy(arr, m.literals)
		return arr
	}
}

func (m *Minia) put(arr []sample.Feature) {
	copy(arr, m.literals)
	select {
	case m.pool <- arr:
	default:
	}
}

/**
 * @brief Evaluate expressions with given feature sets
 * @param features Variable number of feature sets to evaluate against
 * @return Computed features
 */
func (m *Minia) Eval(features ...sample.Features) sample.Features {
	if len(features) == 0 {
		panic("no feature sets provided for evaluation")
	}

	// Get evaluation array
	evalArray := m.get()
	defer m.put(evalArray)

	// Extract column values
	for _, column := range m.columns {
		column.Extract(evalArray, features...)
	}

	// Execute function calls in dependency order
	for _, caller := range m.callers {
		caller.Execute(evalArray)
	}

	// Build result
	result := sample.NewMutableFeatures()
	for i, name := range m.outputNames {
		idx := m.outputIndices[i]
		if idx >= 0 && idx < len(evalArray) && evalArray[idx] != nil {
			result.Set(name, evalArray[idx])
		}
	}

	return result
}

func init() {
	builtinsOnce.Do(func() {
		builtins = make(map[string]FunctionWrapper)
		registerFunctions()
	})
}
