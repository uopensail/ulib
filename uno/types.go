package uno

type NodeType int32
type CmpType int32
type DataType int32
type ArithmeticType string
type BooleanType string

const (
	kVarNode NodeType = iota
	kInt64Node
	kInt64SliceNode
	kFloat32Node
	kFloat32SliceNode
	kStringNode
	kStringSliceNode
	kFunctionNode
	kAndNode
	kOrNode
	kCmpNode
	kInNode
	kNotInNode
	kLiteralNode
	kErrorNode NodeType = 127
)

// const (
// 	Function   ArithmeticType = "function"
// 	Column     ArithmeticType = "column"
// 	Int64      ArithmeticType = "int64"
// 	Int64s     ArithmeticType = "int64s"
// 	Float32    ArithmeticType = "float32"
// 	Float32s   ArithmeticType = "float32s"
// 	String     ArithmeticType = "string"
// 	Strings    ArithmeticType = "strings"
// 	ErrorArith ArithmeticType = "error"
// )

// const (
// 	Literal   BooleanType = "literal"
// 	And       BooleanType = "and"
// 	Or        BooleanType = "or"
// 	Cmp       BooleanType = "cmp"
// 	ErrorBool BooleanType = "error"
// )

type Expression interface {
	GetId() int32
	SetId(int32)
	GetType() NodeType
	Trivial() bool
	Marshal() ([]byte, error)
	ToList() []Expression
}

type BaseExpression struct {
	id int32
}

func (e *BaseExpression) GetId() int32 {
	return e.id
}

func (e *BaseExpression) SetId(id int32) {
	e.id = id
}

const (
	kInt64 DataType = iota
	kFloat32
	kString
	kInt64s
	kFloat32s
	kStrings
	kErrorDType DataType = 127
)

const (
	kNil CmpType = iota
	kEqual
	kNotEqual
	kGreaterThan
	kGreaterThanEqual
	kLessThan
	kLessThanEqual
	kErrorCType CmpType = 127
)

type Stack[T any] struct {
	array []T
}

func (s *Stack[T]) Push(val T) {
	s.array = append(s.array, val)
}

func (s *Stack[T]) Pop() T {
	if len(s.array) == 0 {
		panic("stack empty")
	}

	val := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return val
}

var FUNCTION_IO_TYPES = map[string]map[string][]DataType{
	"mini": {
		"in":  {kInt64, kInt64},
		"out": {kInt64},
	},
	"minf": {
		"in":  {kFloat32, kFloat32},
		"out": {kFloat32},
	},
	"min": {
		"in":  {kFloat32, kFloat32},
		"out": {kFloat32},
	},
	"maxi": {
		"in":  {kInt64, kInt64},
		"out": {kInt64},
	},
	"maxf": {
		"in":  {kFloat32, kFloat32},
		"out": {kFloat32},
	},
	"max": {
		"in":  {kFloat32, kFloat32},
		"out": {kFloat32},
	},
	"addi": {
		"in":  {kInt64, kInt64},
		"out": {kInt64},
	},
	"addf": {
		"in":  {kFloat32, kFloat32},
		"out": {kFloat32},
	},
	"subi": {
		"in":  {kInt64, kInt64},
		"out": {kInt64},
	},
	"subf": {
		"in":  {kFloat32, kFloat32},
		"out": {kFloat32},
	},
	"muli": {
		"in":  {kInt64, kInt64},
		"out": {kInt64},
	},
	"mulf": {
		"in":  {kFloat32, kFloat32},
		"out": {kFloat32},
	},
	"divi": {
		"in":  {kInt64, kInt64},
		"out": {kInt64},
	},
	"divf": {
		"in":  {kFloat32, kFloat32},
		"out": {kFloat32},
	},
	"mod": {
		"in":  {kInt64, kInt64},
		"out": {kInt64},
	},
	"pow": {
		"in":  {kFloat32, kFloat32},
		"out": {kFloat32},
	},
	"round": {
		"in":  {kFloat32},
		"out": {kInt64},
	},
	"floor": {
		"in":  {kFloat32},
		"out": {kInt64},
	},
	"ceil": {
		"in":  {kFloat32},
		"out": {kInt64},
	},
	"log": {
		"in":  {kFloat32},
		"out": {kFloat32},
	},
	"exp": {
		"in":  {kFloat32},
		"out": {kFloat32},
	},
	"log10": {
		"in":  {kFloat32},
		"out": {kFloat32},
	},
	"log2": {
		"in":  {kFloat32},
		"out": {kFloat32},
	},
	"sqrt": {
		"in":  {kFloat32},
		"out": {kFloat32},
	},
	"abs": {
		"in":  {kFloat32},
		"out": {kFloat32},
	},
	"absf": {
		"in":  {kFloat32},
		"out": {kFloat32},
	},
	"absi": {
		"in":  {kInt64},
		"out": {kInt64},
	},
	"sin": {
		"in":  {kFloat32},
		"out": {kFloat32},
	},
	"asin": {
		"in":  {kFloat32},
		"out": {kFloat32},
	},
	"sinh": {
		"in":  {kFloat32},
		"out": {kFloat32},
	},
	"asinh": {
		"in":  {kFloat32},
		"out": {kFloat32},
	},
	"cos": {
		"in":  {kFloat32},
		"out": {kFloat32},
	},
	"acos": {
		"in":  {kFloat32},
		"out": {kFloat32},
	},
	"cosh": {
		"in":  {kFloat32},
		"out": {kFloat32},
	},
	"acosh": {
		"in":  {kFloat32},
		"out": {kFloat32},
	},
	"tanh": {
		"in":  {kFloat32},
		"out": {kFloat32},
	},
	"atan": {
		"in":  {kFloat32},
		"out": {kFloat32},
	},
	"atanh": {
		"in":  {kFloat32},
		"out": {kFloat32},
	},
	"sigmoid": {
		"in":  {kFloat32},
		"out": {kFloat32},
	},

	"year": {
		"in":  {},
		"out": {kString},
	},
	"month": {
		"in":  {},
		"out": {kString},
	},
	"day": {
		"in":  {},
		"out": {kString},
	},
	"date": {
		"in":  {},
		"out": {kString},
	},
	"hour": {
		"in":  {},
		"out": {kString},
	},
	"minute": {
		"in":  {},
		"out": {kString},
	},
	"second": {
		"in":  {},
		"out": {kString},
	},
	"now": {
		"in":  {},
		"out": {kInt64},
	},

	"from_unixtime": {
		"in":  {kInt64, kString},
		"out": {kString},
	},

	"unix_timestamp": {
		"in":  {kString, kString},
		"out": {kInt64},
	},
	"date_add": {
		"in":  {kString, kInt64},
		"out": {kString},
	},
	"date_sub": {
		"in":  {kString, kInt64},
		"out": {kString},
	},
	"date_diff": {
		"in":  {kString, kString},
		"out": {kInt64},
	},
	"reverse": {
		"in":  {kString},
		"out": {kString},
	},
	"upper": {
		"in":  {kString},
		"out": {kString},
	},
	"lower": {
		"in":  {kString},
		"out": {kString},
	},
	"substr": {
		"in":  {kString, kInt64, kInt64},
		"out": {kString},
	},
	"concat": {
		"in":  {kString, kString},
		"out": {kString},
	},
	"casti2f": {
		"in":  {kInt64},
		"out": {kFloat32},
	},
	"castf2i": {
		"in":  {kFloat32},
		"out": {kInt64},
	},
}
