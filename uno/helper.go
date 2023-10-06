package uno

import "github.com/uopensail/ulib/sample"

type NodeType int32
type CmpType int32

const (
	kVarNode NodeType = iota
	kInt64Node
	kFloat32Node
	kStringNode
	kInt64SliceNode
	kFloat32SliceNode
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

type Expression interface {
	GetId() int32
	SetId(int32)
	GetType() NodeType
	Trivial() bool
	MarshalJSON() ([]byte, error)
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

type Stack[T any] struct {
	array []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		array: make([]T, 0),
	}
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

var FUNCTION_IO_TYPES = map[string]map[string][]sample.DataType{
	"mini": {
		"in":  {sample.Int64Type, sample.Int64Type},
		"out": {sample.Int64Type},
	},
	"minf": {
		"in":  {sample.Float32Type, sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"min": {
		"in":  {sample.Float32Type, sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"maxi": {
		"in":  {sample.Int64Type, sample.Int64Type},
		"out": {sample.Int64Type},
	},
	"maxf": {
		"in":  {sample.Float32Type, sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"max": {
		"in":  {sample.Float32Type, sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"addi": {
		"in":  {sample.Int64Type, sample.Int64Type},
		"out": {sample.Int64Type},
	},
	"addf": {
		"in":  {sample.Float32Type, sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"subi": {
		"in":  {sample.Int64Type, sample.Int64Type},
		"out": {sample.Int64Type},
	},
	"subf": {
		"in":  {sample.Float32Type, sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"muli": {
		"in":  {sample.Int64Type, sample.Int64Type},
		"out": {sample.Int64Type},
	},
	"mulf": {
		"in":  {sample.Float32Type, sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"divi": {
		"in":  {sample.Int64Type, sample.Int64Type},
		"out": {sample.Int64Type},
	},
	"divf": {
		"in":  {sample.Float32Type, sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"mod": {
		"in":  {sample.Int64Type, sample.Int64Type},
		"out": {sample.Int64Type},
	},
	"pow": {
		"in":  {sample.Float32Type, sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"round": {
		"in":  {sample.Float32Type},
		"out": {sample.Int64Type},
	},
	"floor": {
		"in":  {sample.Float32Type},
		"out": {sample.Int64Type},
	},
	"ceil": {
		"in":  {sample.Float32Type},
		"out": {sample.Int64Type},
	},
	"log": {
		"in":  {sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"exp": {
		"in":  {sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"log10": {
		"in":  {sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"log2": {
		"in":  {sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"sqrt": {
		"in":  {sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"abs": {
		"in":  {sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"absf": {
		"in":  {sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"absi": {
		"in":  {sample.Int64Type},
		"out": {sample.Int64Type},
	},
	"sin": {
		"in":  {sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"asin": {
		"in":  {sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"sinh": {
		"in":  {sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"asinh": {
		"in":  {sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"cos": {
		"in":  {sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"acos": {
		"in":  {sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"cosh": {
		"in":  {sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"acosh": {
		"in":  {sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"tanh": {
		"in":  {sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"atan": {
		"in":  {sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"atanh": {
		"in":  {sample.Float32Type},
		"out": {sample.Float32Type},
	},
	"sigmoid": {
		"in":  {sample.Float32Type},
		"out": {sample.Float32Type},
	},

	"year": {
		"in":  {},
		"out": {sample.StringType},
	},
	"month": {
		"in":  {},
		"out": {sample.StringType},
	},
	"day": {
		"in":  {},
		"out": {sample.StringType},
	},
	"date": {
		"in":  {},
		"out": {sample.StringType},
	},
	"hour": {
		"in":  {},
		"out": {sample.StringType},
	},
	"minute": {
		"in":  {},
		"out": {sample.StringType},
	},
	"second": {
		"in":  {},
		"out": {sample.StringType},
	},
	"now": {
		"in":  {},
		"out": {sample.Int64Type},
	},

	"from_unixtime": {
		"in":  {sample.Int64Type, sample.StringType},
		"out": {sample.StringType},
	},

	"unix_timestamp": {
		"in":  {sample.StringType, sample.StringType},
		"out": {sample.Int64Type},
	},
	"date_add": {
		"in":  {sample.StringType, sample.Int64Type},
		"out": {sample.StringType},
	},
	"date_sub": {
		"in":  {sample.StringType, sample.Int64Type},
		"out": {sample.StringType},
	},
	"date_diff": {
		"in":  {sample.StringType, sample.StringType},
		"out": {sample.Int64Type},
	},
	"reverse": {
		"in":  {sample.StringType},
		"out": {sample.StringType},
	},
	"upper": {
		"in":  {sample.StringType},
		"out": {sample.StringType},
	},
	"lower": {
		"in":  {sample.StringType},
		"out": {sample.StringType},
	},
	"substr": {
		"in":  {sample.StringType, sample.Int64Type, sample.Int64Type},
		"out": {sample.StringType},
	},
	"concat": {
		"in":  {sample.StringType, sample.StringType},
		"out": {sample.StringType},
	},
	"casti2f": {
		"in":  {sample.Int64Type},
		"out": {sample.Float32Type},
	},
	"castf2i": {
		"in":  {sample.Float32Type},
		"out": {sample.Int64Type},
	},
}
