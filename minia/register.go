package minia

import (
	"math"
)

// ========== Simplified Registration Functions ==========

/**
 * @brief Register unary function with all mapping variants
 * @tparam T Input type
 * @tparam R Return type
 * @param name Function name
 * @param f Unary function to register
 */
func RegisterUnaryFunc[T, R any](name string, f func(T) (R, error)) {
	NewUnaryFunction(name, f)
	NewUnaryMapFunction(name, f)
}

/**
 * @brief Register binary function with all mapping variants
 * @tparam T0 First input type
 * @tparam T1 Second input type
 * @tparam R Return type
 * @param name Function name
 * @param f Binary function to register
 */
func RegisterBinaryFunc[T0, T1, R any](name string, f func(T0, T1) (R, error)) {
	NewBinaryFunction(name, f)
	NewBinaryFirstMapFunction(name, f)
	NewBinarySecondMapFunction(name, f)
	NewBinaryBothMapFunction(name, f)
	NewBinaryProductMapFunction(name, f) // Added missing product mapping
}

/**
 * @brief Register ternary function with all mapping variants
 * @tparam T0 First input type
 * @tparam T1 Second input type
 * @tparam T2 Third input type
 * @tparam R Return type
 * @param name Function name
 * @param f Ternary function to register
 */
func RegisterTernaryFunc[T0, T1, T2, R any](name string, f func(T0, T1, T2) (R, error)) {
	NewTernaryFunction(name, f)
	NewTernaryFirstMapFunction(name, f)
	NewTernarySecondMapFunction(name, f)
	NewTernaryThirdMapFunction(name, f)
	NewTernaryFirstSecondMapFunction(name, f)
	NewTernaryFirstThirdMapFunction(name, f)
	NewTernarySecondThirdMapFunction(name, f)
	NewTernaryAllMapFunction(name, f)
}

/**
 * @brief Register runtime function (no parameters)
 * @tparam R Return type
 * @param name Function name
 * @param f Runtime function to register
 */
func RegisterRuntimeFunc[R any](name string, f func() (R, error)) {
	NewRuntimeFunction(name, f)
}

// ========== Batch Registration Helpers ==========

/**
 * @brief Register math unary functions for both int64 and float32
 * @param name Function name
 * @param mathFunc Math function from math package
 */
func registerMathUnary(name string, mathFunc func(float64) float64) {
	RegisterUnaryFunc(name, MathUnaryFuncGenerator[int64](mathFunc))
	RegisterUnaryFunc(name, MathUnaryFuncGenerator[float32](mathFunc))
}

/**
 * @brief Register math binary functions for all numeric type combinations
 * @param name Function name
 * @param mathFunc Math function from math package
 */
func registerMathBinary(name string, mathFunc func(float64, float64) float64) {
	RegisterBinaryFunc(name, MathBinaryFuncGenerator[int64, int64](mathFunc))
	RegisterBinaryFunc(name, MathBinaryFuncGenerator[float32, float32](mathFunc))
	RegisterBinaryFunc(name, MathBinaryFuncGenerator[int64, float32](mathFunc))
	RegisterBinaryFunc(name, MathBinaryFuncGenerator[float32, int64](mathFunc))
}

/**
 * @brief Register aggregate functions for numeric types
 * @param name Function name
 * @param intFunc Function for int64 slices
 * @param floatFunc Function for float32 slices
 */
func registerAggregate[R any](name string, intFunc func([]int64) (R, error), floatFunc func([]float32) (R, error)) {
	RegisterUnaryFunc(name, intFunc)
	RegisterUnaryFunc(name, floatFunc)
}

// ========== Main Registration Function ==========

/**
 * @brief Register all built-in functions
 * @details This function registers all mathematical, utility, and string functions
 *          with their appropriate type variants and mapping functions
 */
func registerFunctions() {
	// ========== Runtime Functions ==========
	RegisterRuntimeFunc("pi", Pi)
	RegisterRuntimeFunc("e", E)
	RegisterRuntimeFunc("phi", Phi)
	RegisterRuntimeFunc("sqrt2", Sqrt2)
	RegisterRuntimeFunc("ln2", Ln2)
	RegisterRuntimeFunc("ln10", Ln10)
	RegisterRuntimeFunc("year", Year)
	RegisterRuntimeFunc("month", Month)
	RegisterRuntimeFunc("day", Day)
	RegisterRuntimeFunc("curdate", CurDate)
	RegisterRuntimeFunc("unixtimestamp", UnixTimestamp)

	// ========== Math Unary Functions ==========
	registerMathUnary("abs", math.Abs)
	registerMathUnary("sin", math.Sin)
	registerMathUnary("cos", math.Cos)
	registerMathUnary("tan", math.Tan)
	registerMathUnary("asin", math.Asin)
	registerMathUnary("acos", math.Acos)
	registerMathUnary("atan", math.Atan)
	registerMathUnary("sinh", math.Sinh)
	registerMathUnary("cosh", math.Cosh)
	registerMathUnary("tanh", math.Tanh)
	registerMathUnary("exp", math.Exp)
	registerMathUnary("exp2", math.Exp2)
	registerMathUnary("sqrt", math.Sqrt)
	registerMathUnary("cbrt", math.Cbrt)
	registerMathUnary("log", math.Log)
	registerMathUnary("log2", math.Log2)
	registerMathUnary("log10", math.Log10)
	registerMathUnary("log1p", math.Log1p)

	// Float-only math functions
	RegisterUnaryFunc("ceil", MathUnaryFuncGenerator[float32](math.Ceil))
	RegisterUnaryFunc("floor", MathUnaryFuncGenerator[float32](math.Floor))
	RegisterUnaryFunc("round", MathUnaryFuncGenerator[float32](math.Round))
	RegisterUnaryFunc("trunc", MathUnaryFuncGenerator[float32](math.Trunc))

	// ========== Activation Functions ==========
	RegisterUnaryFunc("sigmoid", Sigmoid[int64])
	RegisterUnaryFunc("sigmoid", Sigmoid[float32])

	// ========== Type Conversion ==========
	RegisterUnaryFunc("cast", CastToFloat)
	RegisterUnaryFunc("sign", Sign[int64])
	RegisterUnaryFunc("sign", Sign[float32])

	// ========== Aggregate Functions ==========
	RegisterUnaryFunc("min", Min[int64])
	RegisterUnaryFunc("min", Min[float32])

	RegisterUnaryFunc("max", Max[int64])
	RegisterUnaryFunc("max", Max[float32])

	RegisterUnaryFunc("sum", Sum[int64])
	RegisterUnaryFunc("sum", Sum[float32])

	registerAggregate("average", Average[int64], Average[float32])
	registerAggregate("variance", Variance[int64], Variance[float32])
	registerAggregate("stddev", Stddev[int64], Stddev[float32])
	registerAggregate("median", Median[int64], Median[float32])

	// ========== String Conversion ==========
	RegisterUnaryFunc("tostring", ToString[int64])
	RegisterUnaryFunc("tostring", ToString[float32])

	// ========== Logical Functions ==========
	RegisterUnaryFunc("not", Not)
	RegisterBinaryFunc("and", And)
	RegisterBinaryFunc("or", Or)

	// ========== Collection Functions ==========
	RegisterUnaryFunc("len", Len[int64])
	RegisterUnaryFunc("len", Len[float32])
	RegisterUnaryFunc("len", Len[string])

	RegisterBinaryFunc("contains", Contains[int64])
	RegisterBinaryFunc("contains", Contains[float32])
	RegisterBinaryFunc("contains", Contains[string])

	RegisterBinaryFunc("count", Count[int64])
	RegisterBinaryFunc("count", Count[float32])
	RegisterBinaryFunc("count", Count[string])

	// ========== String Functions ==========
	RegisterUnaryFunc("reverse", Reverse)
	RegisterUnaryFunc("upper", Upper)
	RegisterUnaryFunc("len", StringLength)
	RegisterUnaryFunc("lower", Lower)
	RegisterBinaryFunc("concat", Concat)
	RegisterBinaryFunc("split", SplitString)
	RegisterBinaryFunc("startswith", StartsWith)
	RegisterBinaryFunc("endswith", EndsWith)
	RegisterBinaryFunc("indexof", IndexOf)
	RegisterBinaryFunc("contains", StringContains)

	// ========== Comparison Functions ==========
	// Equality
	RegisterBinaryFunc("eq", EqualSameType[int64])
	RegisterBinaryFunc("eq", EqualSameType[float32])
	RegisterBinaryFunc("eq", EqualSameType[string])
	RegisterBinaryFunc("eq", Equal[int64, float32])
	RegisterBinaryFunc("eq", Equal[int64, string])
	RegisterBinaryFunc("eq", Equal[float32, int64])
	RegisterBinaryFunc("eq", Equal[float32, string])
	RegisterBinaryFunc("eq", Equal[string, int64])
	RegisterBinaryFunc("eq", Equal[string, float32])

	RegisterBinaryFunc("neq", NotEqualSameType[int64])
	RegisterBinaryFunc("neq", NotEqualSameType[float32])
	RegisterBinaryFunc("neq", NotEqualSameType[string])
	RegisterBinaryFunc("neq", NotEqual[int64, float32])
	RegisterBinaryFunc("neq", NotEqual[int64, string])
	RegisterBinaryFunc("neq", NotEqual[float32, int64])
	RegisterBinaryFunc("neq", NotEqual[float32, string])
	RegisterBinaryFunc("neq", NotEqual[string, int64])
	RegisterBinaryFunc("neq", NotEqual[string, float32])

	// Numeric comparisons
	RegisterBinaryFunc("lt", LessThanSameType[int64])
	RegisterBinaryFunc("lt", LessThanSameType[float32])
	RegisterBinaryFunc("lt", LessThan[int64, float32])
	RegisterBinaryFunc("lt", LessThan[float32, int64])

	RegisterBinaryFunc("lte", LessThanEqualSameType[int64])
	RegisterBinaryFunc("lte", LessThanEqualSameType[float32])
	RegisterBinaryFunc("lte", LessThanEqual[int64, float32])
	RegisterBinaryFunc("lte", LessThanEqual[float32, int64])

	RegisterBinaryFunc("gt", GreaterThanSameType[int64])
	RegisterBinaryFunc("gt", GreaterThanSameType[float32])
	RegisterBinaryFunc("gt", GreaterThan[int64, float32])
	RegisterBinaryFunc("gt", GreaterThan[float32, int64])

	RegisterBinaryFunc("gte", GreaterThanEqualSameType[int64])
	RegisterBinaryFunc("gte", GreaterThanEqualSameType[float32])
	RegisterBinaryFunc("gte", GreaterThanEqual[int64, float32])
	RegisterBinaryFunc("gte", GreaterThanEqual[float32, int64])

	// ========== Arithmetic Functions ==========
	RegisterBinaryFunc("add", Add[int64])
	RegisterBinaryFunc("add", Add[float32])

	RegisterBinaryFunc("add", Mix[int64, float32](Add[float32]))
	RegisterBinaryFunc("add", Mix[float32, int64](Add[float32]))

	RegisterBinaryFunc("sub", Sub[int64])
	RegisterBinaryFunc("sub", Sub[float32])
	RegisterBinaryFunc("sub", Mix[int64, float32](Sub[float32]))
	RegisterBinaryFunc("sub", Mix[float32, int64](Sub[float32]))

	RegisterBinaryFunc("mul", Mul[int64])
	RegisterBinaryFunc("mul", Mul[float32])
	RegisterBinaryFunc("mul", Mix[int64, float32](Mul[float32]))
	RegisterBinaryFunc("mul", Mix[float32, int64](Mul[float32]))

	RegisterBinaryFunc("div", Div[int64])
	RegisterBinaryFunc("div", Div[float32])
	RegisterBinaryFunc("div", Mix[int64, float32](Div[float32]))
	RegisterBinaryFunc("div", Mix[float32, int64](Div[float32]))

	RegisterBinaryFunc("mod", Mod)

	// ========== Advanced Math Functions ==========
	registerMathBinary("pow", math.Pow)
	registerMathBinary("atan2", math.Atan2)
	registerMathBinary("hypot", math.Hypot)
	registerMathBinary("remainder", math.Remainder)

	// ========== Normalization Functions ==========
	RegisterBinaryFunc("normalize", Normalize[int64])
	RegisterBinaryFunc("normalize", Normalize[float32])

	RegisterBinaryFunc("binarize", Binarize[int64])
	RegisterBinaryFunc("binarize", Binarize[float32])

	RegisterBinaryFunc("bucketize", Bucketize[int64])
	RegisterBinaryFunc("bucketize", Bucketize[float32])

	// ========== Date/Time Functions ==========
	RegisterUnaryFunc("weekday", Weekday)
	RegisterBinaryFunc("fromunixtime", FromUnixTime)
	RegisterBinaryFunc("datediff", DateDiff)
	RegisterBinaryFunc("daysofmonth", DaysOfMonth)

	// ========== Ternary Functions ==========
	RegisterTernaryFunc("clamp", Clamp[int64, int64, int64])
	RegisterTernaryFunc("clamp", Clamp[int64, int64, float32])
	RegisterTernaryFunc("clamp", Clamp[int64, float32, int64])
	RegisterTernaryFunc("clamp", Clamp[int64, float32, float32])
	RegisterTernaryFunc("clamp", Clamp[float32, int64, int64])
	RegisterTernaryFunc("clamp", Clamp[float32, int64, float32])
	RegisterTernaryFunc("clamp", Clamp[float32, float32, int64])
	RegisterTernaryFunc("clamp", Clamp[float32, float32, float32])

	RegisterTernaryFunc("minmax", MinMax[int64])
	RegisterTernaryFunc("minmax", MinMax[float32])

	RegisterTernaryFunc("zscore", ZScore)

	RegisterTernaryFunc("date_add", DateAdd)
	RegisterTernaryFunc("date_sub", DateSub)

	RegisterTernaryFunc("substr", Substr)

	// ========== Advanced Collection Functions ==========
	RegisterBinaryFunc("topk", TopK[int64])
	RegisterBinaryFunc("topk", TopK[float32])
	RegisterBinaryFunc("topk", TopK[string])

	// ========== Softmax Functions ==========
	RegisterUnaryFunc("softmax", Softmax[int64])
	RegisterUnaryFunc("softmax", Softmax[float32])
	RegisterUnaryFunc("logsoftmax", LogSoftmax[int64])
	RegisterUnaryFunc("logsoftmax", LogSoftmax[float32])
}
