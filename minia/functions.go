package minia

import (
	"fmt"

	"github.com/uopensail/ulib/sample"
)

// ========== Base Function Interface ==========

/**
 * @brief Base interface for all function wrappers
 */
type FunctionWrapper interface {
	Call(args ...sample.Feature) (sample.Feature, error)
	GetSignature() string
}

/**
 * @brief Base function wrapper with common functionality
 */
type BaseFunctionWrapper struct {
	Name      string
	Signature string
}

/**
 * @brief Get the signature of the function
 * @return Function signature string
 */
func (b *BaseFunctionWrapper) GetSignature() string {
	return b.Signature
}

/**
 * @brief Validate argument count
 * @param args Input arguments
 * @param expected Expected argument count
 * @return Error if count mismatch, nil otherwise
 */
func (b *BaseFunctionWrapper) ValidateArgCount(args []sample.Feature, expected int) error {
	if len(args) != expected {
		return fmt.Errorf("function %s expects %d arguments, got %d", b.Name, expected, len(args))
	}
	return nil
}

/**
 * @brief Safely extract value from Feature with type checking
 * @tparam T Expected type
 * @param feature Input feature
 * @param argIndex Argument index for error reporting
 * @return Extracted value and error
 */
func ExtractValue[T any](feature sample.Feature, argIndex int) (T, error) {
	var zero T
	if feature == nil {
		return zero, fmt.Errorf("argument %d is nil", argIndex)
	}

	value, ok := feature.Get().(T)
	if !ok {
		return zero, fmt.Errorf("argument %d type mismatch: expected %T, got %T",
			argIndex, zero, feature.Get())
	}
	return value, nil
}

// ========== Runtime Function (0 parameters) ==========

/**
 * @brief Wrapper for runtime functions with no parameters
 * @tparam R Return type
 */
type RuntimeFunction[R any] struct {
	BaseFunctionWrapper
	f func() (R, error)
}

/**
 * @brief Create new runtime function wrapper
 * @tparam R Return type
 * @param name Function name
 * @param f Function to wrap
 * @return Function wrapper instance
 */
func NewRuntimeFunction[R any](name string, f func() (R, error)) *RuntimeFunction[R] {
	wrapper := &RuntimeFunction[R]{
		BaseFunctionWrapper: BaseFunctionWrapper{
			Name:      name,
			Signature: fmt.Sprintf("%s:0=[]", name),
		},
		f: f,
	}
	builtins[name] = wrapper
	return wrapper
}

/**
 * @brief Execute the runtime function
 * @param args Input arguments (should be empty)
 * @return Function result as Feature
 */
func (rf *RuntimeFunction[R]) Call(args ...sample.Feature) (sample.Feature, error) {
	if err := rf.ValidateArgCount(args, 0); err != nil {
		return nil, err
	}

	value, err := rf.f()
	if err != nil {
		return nil, fmt.Errorf("runtime function %s failed: %w", rf.Name, err)
	}

	return NewFeature(value)
}

// ========== Unary Function (1 parameter) ==========

/**
 * @brief Wrapper for unary functions
 * @tparam T Input type
 * @tparam R Return type
 */
type UnaryFunction[T, R any] struct {
	BaseFunctionWrapper
	f func(T) (R, error)
}

/**
 * @brief Create new unary function wrapper
 * @tparam T Input type
 * @tparam R Return type
 * @param name Function name
 * @param f Function to wrap
 * @return Function wrapper instance
 */
func NewUnaryFunction[T, R any](name string, f func(T) (R, error)) *UnaryFunction[T, R] {
	signature := generateSignature(name, f)
	wrapper := &UnaryFunction[T, R]{
		BaseFunctionWrapper: BaseFunctionWrapper{
			Name:      name,
			Signature: signature,
		},
		f: f,
	}
	builtins[signature] = wrapper
	return wrapper
}

/**
 * @brief Execute the unary function
 * @param args Input arguments (should contain 1 element)
 * @return Function result as Feature
 */
func (uf *UnaryFunction[T, R]) Call(args ...sample.Feature) (sample.Feature, error) {
	if err := uf.ValidateArgCount(args, 1); err != nil {
		return nil, err
	}

	arg0, err := ExtractValue[T](args[0], 0)
	if err != nil {
		return nil, fmt.Errorf("unary function %s: %w", uf.Name, err)
	}

	value, err := uf.f(arg0)
	if err != nil {
		return nil, fmt.Errorf("unary function %s failed: %w", uf.Name, err)
	}

	return NewFeature(value)
}

// ========== Unary Map Function (slice input) ==========

/**
 * @brief Wrapper for unary functions that operate on slices
 * @tparam T Input element type
 * @tparam R Return element type
 */
type UnaryMapFunction[T, R any] struct {
	BaseFunctionWrapper
	f func([]T) ([]R, error)
}

/**
 * @brief Create new unary map function wrapper
 * @tparam T Input element type
 * @tparam R Return element type
 * @param name Function name
 * @param f Original element-wise function
 * @return Function wrapper instance
 */
func NewUnaryMapFunction[T, R any](name string, f func(T) (R, error)) *UnaryMapFunction[T, R] {
	mapFunc := Map(f)
	signature := generateSignature(name, mapFunc)
	wrapper := &UnaryMapFunction[T, R]{
		BaseFunctionWrapper: BaseFunctionWrapper{
			Name:      name,
			Signature: signature,
		},
		f: mapFunc,
	}
	builtins[signature] = wrapper
	return wrapper
}

/**
 * @brief Execute the unary map function
 * @param args Input arguments (should contain 1 slice)
 * @return Function result as Feature
 */
func (umf *UnaryMapFunction[T, R]) Call(args ...sample.Feature) (sample.Feature, error) {
	if err := umf.ValidateArgCount(args, 1); err != nil {
		return nil, err
	}

	arg0, err := ExtractValue[[]T](args[0], 0)
	if err != nil {
		return nil, fmt.Errorf("unary map function %s: %w", umf.Name, err)
	}

	value, err := umf.f(arg0)
	if err != nil {
		return nil, fmt.Errorf("unary map function %s failed: %w", umf.Name, err)
	}

	return NewFeature(value)
}

// ========== Binary Function (2 parameters) ==========

/**
 * @brief Wrapper for binary functions
 * @tparam T0 First input type
 * @tparam T1 Second input type
 * @tparam R Return type
 */
type BinaryFunction[T0, T1, R any] struct {
	BaseFunctionWrapper
	f func(T0, T1) (R, error)
}

/**
 * @brief Create new binary function wrapper
 * @tparam T0 First input type
 * @tparam T1 Second input type
 * @tparam R Return type
 * @param name Function name
 * @param f Function to wrap
 * @return Function wrapper instance
 */
func NewBinaryFunction[T0, T1, R any](name string, f func(T0, T1) (R, error)) *BinaryFunction[T0, T1, R] {
	signature := generateSignature(name, f)
	wrapper := &BinaryFunction[T0, T1, R]{
		BaseFunctionWrapper: BaseFunctionWrapper{
			Name:      name,
			Signature: signature,
		},
		f: f,
	}
	builtins[signature] = wrapper
	return wrapper
}

/**
 * @brief Execute the binary function
 * @param args Input arguments (should contain 2 elements)
 * @return Function result as Feature
 */
func (bf *BinaryFunction[T0, T1, R]) Call(args ...sample.Feature) (sample.Feature, error) {
	if err := bf.ValidateArgCount(args, 2); err != nil {
		return nil, err
	}

	arg0, err := ExtractValue[T0](args[0], 0)
	if err != nil {
		return nil, fmt.Errorf("binary function %s: %w", bf.Name, err)
	}

	arg1, err := ExtractValue[T1](args[1], 1)
	if err != nil {
		return nil, fmt.Errorf("binary function %s: %w", bf.Name, err)
	}

	value, err := bf.f(arg0, arg1)
	if err != nil {
		return nil, fmt.Errorf("binary function %s failed: %w", bf.Name, err)
	}

	return NewFeature(value)
}

// ========== Binary Map Functions ==========

/**
 * @brief Wrapper for binary functions where first parameter is a slice
 * @tparam T0 First input element type (slice)
 * @tparam T1 Second input type (scalar)
 * @tparam R Return element type
 */
type BinaryFirstMapFunction[T0, T1, R any] struct {
	BaseFunctionWrapper
	f func([]T0, T1) ([]R, error)
}

/**
 * @brief Create new binary first map function wrapper
 * @tparam T0 First input element type
 * @tparam T1 Second input type
 * @tparam R Return element type
 * @param name Function name
 * @param f Original binary function
 * @return Function wrapper instance
 */
func NewBinaryFirstMapFunction[T0, T1, R any](name string, f func(T0, T1) (R, error)) *BinaryFirstMapFunction[T0, T1, R] {
	mapFunc := Map2First(f)
	signature := generateSignature(name, mapFunc)
	wrapper := &BinaryFirstMapFunction[T0, T1, R]{
		BaseFunctionWrapper: BaseFunctionWrapper{
			Name:      name,
			Signature: signature,
		},
		f: mapFunc,
	}
	builtins[signature] = wrapper
	return wrapper
}

/**
 * @brief Execute the binary first map function
 * @param args Input arguments (slice, scalar)
 * @return Function result as Feature
 */
func (bfmf *BinaryFirstMapFunction[T0, T1, R]) Call(args ...sample.Feature) (sample.Feature, error) {
	if err := bfmf.ValidateArgCount(args, 2); err != nil {
		return nil, err
	}

	arg0, err := ExtractValue[[]T0](args[0], 0)
	if err != nil {
		return nil, fmt.Errorf("binary first map function %s: %w", bfmf.Name, err)
	}

	arg1, err := ExtractValue[T1](args[1], 1)
	if err != nil {
		return nil, fmt.Errorf("binary first map function %s: %w", bfmf.Name, err)
	}

	value, err := bfmf.f(arg0, arg1)
	if err != nil {
		return nil, fmt.Errorf("binary first map function %s failed: %w", bfmf.Name, err)
	}

	return NewFeature(value)
}

/**
 * @brief Wrapper for binary functions where second parameter is a slice
 * @tparam T0 First input type (scalar)
 * @tparam T1 Second input element type (slice)
 * @tparam R Return element type
 */
type BinarySecondMapFunction[T0, T1, R any] struct {
	BaseFunctionWrapper
	f func(T0, []T1) ([]R, error)
}

/**
 * @brief Create new binary second map function wrapper
 * @tparam T0 First input type
 * @tparam T1 Second input element type
 * @tparam R Return element type
 * @param name Function name
 * @param f Original binary function
 * @return Function wrapper instance
 */
func NewBinarySecondMapFunction[T0, T1, R any](name string, f func(T0, T1) (R, error)) *BinarySecondMapFunction[T0, T1, R] {
	mapFunc := Map2Second(f)
	signature := generateSignature(name, mapFunc)
	wrapper := &BinarySecondMapFunction[T0, T1, R]{
		BaseFunctionWrapper: BaseFunctionWrapper{
			Name:      name,
			Signature: signature,
		},
		f: mapFunc,
	}
	builtins[signature] = wrapper
	return wrapper
}

/**
 * @brief Execute the binary second map function
 * @param args Input arguments (scalar, slice)
 * @return Function result as Feature
 */
func (bsmf *BinarySecondMapFunction[T0, T1, R]) Call(args ...sample.Feature) (sample.Feature, error) {
	if err := bsmf.ValidateArgCount(args, 2); err != nil {
		return nil, err
	}

	arg0, err := ExtractValue[T0](args[0], 0)
	if err != nil {
		return nil, fmt.Errorf("binary second map function %s: %w", bsmf.Name, err)
	}

	arg1, err := ExtractValue[[]T1](args[1], 1)
	if err != nil {
		return nil, fmt.Errorf("binary second map function %s: %w", bsmf.Name, err)
	}

	value, err := bsmf.f(arg0, arg1)
	if err != nil {
		return nil, fmt.Errorf("binary second map function %s failed: %w", bsmf.Name, err)
	}

	return NewFeature(value)
}

/**
 * @brief Wrapper for binary functions where both parameters are slices
 * @tparam T0 First input element type (slice)
 * @tparam T1 Second input element type (slice)
 * @tparam R Return element type
 */
type BinaryBothMapFunction[T0, T1, R any] struct {
	BaseFunctionWrapper
	f func([]T0, []T1) ([]R, error)
}

/**
 * @brief Create new binary both map function wrapper
 * @tparam T0 First input element type
 * @tparam T1 Second input element type
 * @tparam R Return element type
 * @param name Function name
 * @param f Original binary function
 * @return Function wrapper instance
 */
func NewBinaryBothMapFunction[T0, T1, R any](name string, f func(T0, T1) (R, error)) *BinaryBothMapFunction[T0, T1, R] {
	mapFunc := Map2Both(f)
	signature := generateSignature(name, mapFunc)
	wrapper := &BinaryBothMapFunction[T0, T1, R]{
		BaseFunctionWrapper: BaseFunctionWrapper{
			Name:      name,
			Signature: signature,
		},
		f: mapFunc,
	}
	builtins[signature] = wrapper
	return wrapper
}

/**
 * @brief Execute the binary both map function
 * @param args Input arguments (slice, slice)
 * @return Function result as Feature
 */
func (bbmf *BinaryBothMapFunction[T0, T1, R]) Call(args ...sample.Feature) (sample.Feature, error) {
	if err := bbmf.ValidateArgCount(args, 2); err != nil {
		return nil, err
	}

	arg0, err := ExtractValue[[]T0](args[0], 0)
	if err != nil {
		return nil, fmt.Errorf("binary both map function %s: %w", bbmf.Name, err)
	}

	arg1, err := ExtractValue[[]T1](args[1], 1)
	if err != nil {
		return nil, fmt.Errorf("binary both map function %s: %w", bbmf.Name, err)
	}

	value, err := bbmf.f(arg0, arg1)
	if err != nil {
		return nil, fmt.Errorf("binary both map function %s failed: %w", bbmf.Name, err)
	}

	return NewFeature(value)
}

/**
 * @brief Wrapper for binary functions with Cartesian product mapping
 * @tparam T0 First input element type (slice)
 * @tparam T1 Second input element type (slice)
 * @tparam R Return element type
 */
type BinaryProductMapFunction[T0, T1, R any] struct {
	BaseFunctionWrapper
	f func([]T0, []T1) ([]R, error)
}

/**
 * @brief Create new binary product map function wrapper
 * @tparam T0 First input element type
 * @tparam T1 Second input element type
 * @tparam R Return element type
 * @param name Function name
 * @param f Original binary function
 * @return Function wrapper instance
 */
func NewBinaryProductMapFunction[T0, T1, R any](name string, f func(T0, T1) (R, error)) *BinaryProductMapFunction[T0, T1, R] {
	mapFunc := Product(f)
	signature := generateSignature(name, mapFunc)
	wrapper := &BinaryProductMapFunction[T0, T1, R]{
		BaseFunctionWrapper: BaseFunctionWrapper{
			Name:      name,
			Signature: signature,
		},
		f: mapFunc,
	}
	builtins[signature] = wrapper
	return wrapper
}

/**
 * @brief Execute the binary product map function
 * @param args Input arguments (slice, slice)
 * @return Function result as Feature (Cartesian product)
 */
func (bpmf *BinaryProductMapFunction[T0, T1, R]) Call(args ...sample.Feature) (sample.Feature, error) {
	if err := bpmf.ValidateArgCount(args, 2); err != nil {
		return nil, err
	}

	arg0, err := ExtractValue[[]T0](args[0], 0)
	if err != nil {
		return nil, fmt.Errorf("binary product map function %s: %w", bpmf.Name, err)
	}

	arg1, err := ExtractValue[[]T1](args[1], 1)
	if err != nil {
		return nil, fmt.Errorf("binary product map function %s: %w", bpmf.Name, err)
	}

	value, err := bpmf.f(arg0, arg1)
	if err != nil {
		return nil, fmt.Errorf("binary product map function %s failed: %w", bpmf.Name, err)
	}

	return NewFeature(value)
}

// ========== Ternary Function (3 parameters) ==========

/**
 * @brief Wrapper for ternary functions
 * @tparam T0 First input type
 * @tparam T1 Second input type
 * @tparam T2 Third input type
 * @tparam R Return type
 */
type TernaryFunction[T0, T1, T2, R any] struct {
	BaseFunctionWrapper
	f func(T0, T1, T2) (R, error)
}

/**
 * @brief Create new ternary function wrapper
 * @tparam T0 First input type
 * @tparam T1 Second input type
 * @tparam T2 Third input type
 * @tparam R Return type
 * @param name Function name
 * @param f Function to wrap
 * @return Function wrapper instance
 */
func NewTernaryFunction[T0, T1, T2, R any](name string, f func(T0, T1, T2) (R, error)) *TernaryFunction[T0, T1, T2, R] {
	signature := generateSignature(name, f)
	wrapper := &TernaryFunction[T0, T1, T2, R]{
		BaseFunctionWrapper: BaseFunctionWrapper{
			Name:      name,
			Signature: signature,
		},
		f: f,
	}
	builtins[signature] = wrapper
	return wrapper
}

/**
 * @brief Execute the ternary function
 * @param args Input arguments (should contain 3 elements)
 * @return Function result as Feature
 */
func (tf *TernaryFunction[T0, T1, T2, R]) Call(args ...sample.Feature) (sample.Feature, error) {
	if err := tf.ValidateArgCount(args, 3); err != nil {
		return nil, err
	}

	arg0, err := ExtractValue[T0](args[0], 0)
	if err != nil {
		return nil, fmt.Errorf("ternary function %s: %w", tf.Name, err)
	}

	arg1, err := ExtractValue[T1](args[1], 1)
	if err != nil {
		return nil, fmt.Errorf("ternary function %s: %w", tf.Name, err)
	}

	arg2, err := ExtractValue[T2](args[2], 2)
	if err != nil {
		return nil, fmt.Errorf("ternary function %s: %w", tf.Name, err)
	}

	value, err := tf.f(arg0, arg1, arg2)
	if err != nil {
		return nil, fmt.Errorf("ternary function %s failed: %w", tf.Name, err)
	}

	return NewFeature(value)
}

// ========== Ternary Map Functions ==========

/**
 * @brief Wrapper for ternary functions where first parameter is a slice
 * @tparam T0 First input element type (slice)
 * @tparam T1 Second input type (scalar)
 * @tparam T2 Third input type (scalar)
 * @tparam R Return element type
 */
type TernaryFirstMapFunction[T0, T1, T2, R any] struct {
	BaseFunctionWrapper
	f func([]T0, T1, T2) ([]R, error)
}

/**
 * @brief Create new ternary first map function wrapper
 */
func NewTernaryFirstMapFunction[T0, T1, T2, R any](name string, f func(T0, T1, T2) (R, error)) *TernaryFirstMapFunction[T0, T1, T2, R] {
	mapFunc := Map3First(f)
	signature := generateSignature(name, mapFunc)
	wrapper := &TernaryFirstMapFunction[T0, T1, T2, R]{
		BaseFunctionWrapper: BaseFunctionWrapper{
			Name:      name,
			Signature: signature,
		},
		f: mapFunc,
	}
	builtins[signature] = wrapper
	return wrapper
}

/**
 * @brief Execute the ternary first map function
 */
func (tfmf *TernaryFirstMapFunction[T0, T1, T2, R]) Call(args ...sample.Feature) (sample.Feature, error) {
	if err := tfmf.ValidateArgCount(args, 3); err != nil {
		return nil, err
	}

	arg0, err := ExtractValue[[]T0](args[0], 0)
	if err != nil {
		return nil, fmt.Errorf("ternary first map function %s: %w", tfmf.Name, err)
	}

	arg1, err := ExtractValue[T1](args[1], 1)
	if err != nil {
		return nil, fmt.Errorf("ternary first map function %s: %w", tfmf.Name, err)
	}

	arg2, err := ExtractValue[T2](args[2], 2)
	if err != nil {
		return nil, fmt.Errorf("ternary first map function %s: %w", tfmf.Name, err)
	}

	value, err := tfmf.f(arg0, arg1, arg2)
	if err != nil {
		return nil, fmt.Errorf("ternary first map function %s failed: %w", tfmf.Name, err)
	}

	return NewFeature(value)
}

/**
 * @brief Wrapper for ternary functions where second parameter is a slice
 */
type TernarySecondMapFunction[T0, T1, T2, R any] struct {
	BaseFunctionWrapper
	f func(T0, []T1, T2) ([]R, error)
}

/**
 * @brief Create new ternary second map function wrapper
 */
func NewTernarySecondMapFunction[T0, T1, T2, R any](name string, f func(T0, T1, T2) (R, error)) *TernarySecondMapFunction[T0, T1, T2, R] {
	mapFunc := Map3Second(f)
	signature := generateSignature(name, mapFunc)
	wrapper := &TernarySecondMapFunction[T0, T1, T2, R]{
		BaseFunctionWrapper: BaseFunctionWrapper{
			Name:      name,
			Signature: signature,
		},
		f: mapFunc,
	}
	builtins[signature] = wrapper
	return wrapper
}

/**
 * @brief Execute the ternary second map function
 */
func (tsmf *TernarySecondMapFunction[T0, T1, T2, R]) Call(args ...sample.Feature) (sample.Feature, error) {
	if err := tsmf.ValidateArgCount(args, 3); err != nil {
		return nil, err
	}

	arg0, err := ExtractValue[T0](args[0], 0)
	if err != nil {
		return nil, fmt.Errorf("ternary second map function %s: %w", tsmf.Name, err)
	}

	arg1, err := ExtractValue[[]T1](args[1], 1)
	if err != nil {
		return nil, fmt.Errorf("ternary second map function %s: %w", tsmf.Name, err)
	}

	arg2, err := ExtractValue[T2](args[2], 2)
	if err != nil {
		return nil, fmt.Errorf("ternary second map function %s: %w", tsmf.Name, err)
	}

	value, err := tsmf.f(arg0, arg1, arg2)
	if err != nil {
		return nil, fmt.Errorf("ternary second map function %s failed: %w", tsmf.Name, err)
	}

	return NewFeature(value)
}

/**
 * @brief Wrapper for ternary functions where third parameter is a slice
 */
type TernaryThirdMapFunction[T0, T1, T2, R any] struct {
	BaseFunctionWrapper
	f func(T0, T1, []T2) ([]R, error)
}

/**
 * @brief Create new ternary third map function wrapper
 */
func NewTernaryThirdMapFunction[T0, T1, T2, R any](name string, f func(T0, T1, T2) (R, error)) *TernaryThirdMapFunction[T0, T1, T2, R] {
	mapFunc := Map3Third(f)
	signature := generateSignature(name, mapFunc)
	wrapper := &TernaryThirdMapFunction[T0, T1, T2, R]{
		BaseFunctionWrapper: BaseFunctionWrapper{
			Name:      name,
			Signature: signature,
		},
		f: mapFunc,
	}
	builtins[signature] = wrapper
	return wrapper
}

/**
 * @brief Execute the ternary third map function
 */
func (ttmf *TernaryThirdMapFunction[T0, T1, T2, R]) Call(args ...sample.Feature) (sample.Feature, error) {
	if err := ttmf.ValidateArgCount(args, 3); err != nil {
		return nil, err
	}

	arg0, err := ExtractValue[T0](args[0], 0)
	if err != nil {
		return nil, fmt.Errorf("ternary third map function %s: %w", ttmf.Name, err)
	}

	arg1, err := ExtractValue[T1](args[1], 1)
	if err != nil {
		return nil, fmt.Errorf("ternary third map function %s: %w", ttmf.Name, err)
	}

	arg2, err := ExtractValue[[]T2](args[2], 2)
	if err != nil {
		return nil, fmt.Errorf("ternary third map function %s: %w", ttmf.Name, err)
	}

	value, err := ttmf.f(arg0, arg1, arg2)
	if err != nil {
		return nil, fmt.Errorf("ternary third map function %s failed: %w", ttmf.Name, err)
	}

	return NewFeature(value)
}

/**
 * @brief Wrapper for ternary functions where first and second parameters are slices
 */
type TernaryFirstSecondMapFunction[T0, T1, T2, R any] struct {
	BaseFunctionWrapper
	f func([]T0, []T1, T2) ([]R, error)
}

/**
 * @brief Create new ternary first-second map function wrapper
 */
func NewTernaryFirstSecondMapFunction[T0, T1, T2, R any](name string, f func(T0, T1, T2) (R, error)) *TernaryFirstSecondMapFunction[T0, T1, T2, R] {
	mapFunc := Map3FirstSecond(f)
	signature := generateSignature(name, mapFunc)
	wrapper := &TernaryFirstSecondMapFunction[T0, T1, T2, R]{
		BaseFunctionWrapper: BaseFunctionWrapper{
			Name:      name,
			Signature: signature,
		},
		f: mapFunc,
	}
	builtins[signature] = wrapper
	return wrapper
}

/**
 * @brief Execute the ternary first-second map function
 */
func (tfsmf *TernaryFirstSecondMapFunction[T0, T1, T2, R]) Call(args ...sample.Feature) (sample.Feature, error) {
	if err := tfsmf.ValidateArgCount(args, 3); err != nil {
		return nil, err
	}

	arg0, err := ExtractValue[[]T0](args[0], 0)
	if err != nil {
		return nil, fmt.Errorf("ternary first-second map function %s: %w", tfsmf.Name, err)
	}

	arg1, err := ExtractValue[[]T1](args[1], 1)
	if err != nil {
		return nil, fmt.Errorf("ternary first-second map function %s: %w", tfsmf.Name, err)
	}

	arg2, err := ExtractValue[T2](args[2], 2)
	if err != nil {
		return nil, fmt.Errorf("ternary first-second map function %s: %w", tfsmf.Name, err)
	}

	value, err := tfsmf.f(arg0, arg1, arg2)
	if err != nil {
		return nil, fmt.Errorf("ternary first-second map function %s failed: %w", tfsmf.Name, err)
	}

	return NewFeature(value)
}

/**
 * @brief Wrapper for ternary functions where first and third parameters are slices
 */
type TernaryFirstThirdMapFunction[T0, T1, T2, R any] struct {
	BaseFunctionWrapper
	f func([]T0, T1, []T2) ([]R, error)
}

/**
 * @brief Create new ternary first-third map function wrapper
 */
func NewTernaryFirstThirdMapFunction[T0, T1, T2, R any](name string, f func(T0, T1, T2) (R, error)) *TernaryFirstThirdMapFunction[T0, T1, T2, R] {
	mapFunc := Map3FirstThird(f)
	signature := generateSignature(name, mapFunc)
	wrapper := &TernaryFirstThirdMapFunction[T0, T1, T2, R]{
		BaseFunctionWrapper: BaseFunctionWrapper{
			Name:      name,
			Signature: signature,
		},
		f: mapFunc,
	}
	builtins[signature] = wrapper
	return wrapper
}

/**
 * @brief Execute the ternary first-third map function
 */
func (tftmf *TernaryFirstThirdMapFunction[T0, T1, T2, R]) Call(args ...sample.Feature) (sample.Feature, error) {
	if err := tftmf.ValidateArgCount(args, 3); err != nil {
		return nil, err
	}

	arg0, err := ExtractValue[[]T0](args[0], 0)
	if err != nil {
		return nil, fmt.Errorf("ternary first-third map function %s: %w", tftmf.Name, err)
	}

	arg1, err := ExtractValue[T1](args[1], 1)
	if err != nil {
		return nil, fmt.Errorf("ternary first-third map function %s: %w", tftmf.Name, err)
	}

	arg2, err := ExtractValue[[]T2](args[2], 2)
	if err != nil {
		return nil, fmt.Errorf("ternary first-third map function %s: %w", tftmf.Name, err)
	}

	value, err := tftmf.f(arg0, arg1, arg2)
	if err != nil {
		return nil, fmt.Errorf("ternary first-third map function %s failed: %w", tftmf.Name, err)
	}

	return NewFeature(value)
}

/**
 * @brief Wrapper for ternary functions where second and third parameters are slices
 */
type TernarySecondThirdMapFunction[T0, T1, T2, R any] struct {
	BaseFunctionWrapper
	f func(T0, []T1, []T2) ([]R, error)
}

/**
 * @brief Create new ternary second-third map function wrapper
 */
func NewTernarySecondThirdMapFunction[T0, T1, T2, R any](name string, f func(T0, T1, T2) (R, error)) *TernarySecondThirdMapFunction[T0, T1, T2, R] {
	mapFunc := Map3SecondThird(f)
	signature := generateSignature(name, mapFunc)
	wrapper := &TernarySecondThirdMapFunction[T0, T1, T2, R]{
		BaseFunctionWrapper: BaseFunctionWrapper{
			Name:      name,
			Signature: signature,
		},
		f: mapFunc,
	}
	builtins[signature] = wrapper
	return wrapper
}

/**
 * @brief Execute the ternary second-third map function
 */
func (tstmf *TernarySecondThirdMapFunction[T0, T1, T2, R]) Call(args ...sample.Feature) (sample.Feature, error) {
	if err := tstmf.ValidateArgCount(args, 3); err != nil {
		return nil, err
	}

	arg0, err := ExtractValue[T0](args[0], 0)
	if err != nil {
		return nil, fmt.Errorf("ternary second-third map function %s: %w", tstmf.Name, err)
	}

	arg1, err := ExtractValue[[]T1](args[1], 1)
	if err != nil {
		return nil, fmt.Errorf("ternary second-third map function %s: %w", tstmf.Name, err)
	}

	arg2, err := ExtractValue[[]T2](args[2], 2)
	if err != nil {
		return nil, fmt.Errorf("ternary second-third map function %s: %w", tstmf.Name, err)
	}

	value, err := tstmf.f(arg0, arg1, arg2)
	if err != nil {
		return nil, fmt.Errorf("ternary second-third map function %s failed: %w", tstmf.Name, err)
	}

	return NewFeature(value)
}

/**
 * @brief Wrapper for ternary functions where all three parameters are slices
 */
type TernaryAllMapFunction[T0, T1, T2, R any] struct {
	BaseFunctionWrapper
	f func([]T0, []T1, []T2) ([]R, error)
}

/**
 * @brief Create new ternary all map function wrapper
 */
func NewTernaryAllMapFunction[T0, T1, T2, R any](name string, f func(T0, T1, T2) (R, error)) *TernaryAllMapFunction[T0, T1, T2, R] {
	mapFunc := Map3All(f)
	signature := generateSignature(name, mapFunc)
	wrapper := &TernaryAllMapFunction[T0, T1, T2, R]{
		BaseFunctionWrapper: BaseFunctionWrapper{
			Name:      name,
			Signature: signature,
		},
		f: mapFunc,
	}
	builtins[signature] = wrapper
	return wrapper
}

/**
 * @brief Execute the ternary all map function
 */
func (tamf *TernaryAllMapFunction[T0, T1, T2, R]) Call(args ...sample.Feature) (sample.Feature, error) {
	if err := tamf.ValidateArgCount(args, 3); err != nil {
		return nil, err
	}

	arg0, err := ExtractValue[[]T0](args[0], 0)
	if err != nil {
		return nil, fmt.Errorf("ternary all map function %s: %w", tamf.Name, err)
	}

	arg1, err := ExtractValue[[]T1](args[1], 1)
	if err != nil {
		return nil, fmt.Errorf("ternary all map function %s: %w", tamf.Name, err)
	}

	arg2, err := ExtractValue[[]T2](args[2], 2)
	if err != nil {
		return nil, fmt.Errorf("ternary all map function %s: %w", tamf.Name, err)
	}

	value, err := tamf.f(arg0, arg1, arg2)
	if err != nil {
		return nil, fmt.Errorf("ternary all map function %s failed: %w", tamf.Name, err)
	}

	return NewFeature(value)
}

// ========== Utility Functions ==========

/**
 * @brief Get function wrapper from registry by signature
 * @param signature Function signature string
 * @return Function wrapper if found, nil otherwise
 */
func GetFunction(signature string) FunctionWrapper {
	return builtins[signature]
}

/**
 * @brief Register a function wrapper in the global registry
 * @param signature Function signature string
 * @param wrapper Function wrapper to register
 */
func RegisterFunction(signature string, wrapper FunctionWrapper) {
	builtins[signature] = wrapper
}

/**
 * @brief List all registered function signatures
 * @return Slice of all registered function signatures
 */
func ListFunctions() []string {
	signatures := make([]string, 0, len(builtins))
	for sig := range builtins {
		signatures = append(signatures, sig)
	}
	return signatures
}

/**
 * @brief Clear all registered functions
 */
func ClearFunctions() {
	builtins = make(map[string]FunctionWrapper)
}

/**
 * @brief Check if a function is registered
 * @param signature Function signature string
 * @return True if function is registered, false otherwise
 */
func HasFunction(signature string) bool {
	_, exists := builtins[signature]
	return exists
}

/**
 * @brief Get the number of registered functions
 * @return Number of registered functions
 */
func FunctionCount() int {
	return len(builtins)
}

// ========== Helper Functions for Map Operations ==========

/**
 * @brief Map function for single parameter slice operations
 * @tparam T Input element type
 * @tparam R Return element type
 * @param f Function to apply to each element
 * @return Function that operates on slices
 */
func Map[T, R any](f func(T) (R, error)) func([]T) ([]R, error) {
	return func(slice []T) ([]R, error) {
		if len(slice) == 0 {
			return []R{}, nil
		}

		result := make([]R, len(slice))
		for i, item := range slice {
			val, err := f(item)
			if err != nil {
				return nil, fmt.Errorf("map operation failed at index %d: %w", i, err)
			}
			result[i] = val
		}
		return result, nil
	}
}

/**
 * @brief Map function for binary operations where first parameter is a slice
 * @tparam T0 First input element type
 * @tparam T1 Second input type
 * @tparam R Return element type
 * @param f Binary function to apply
 * @return Function that operates on slice and scalar
 */
func Map2First[T0, T1, R any](f func(T0, T1) (R, error)) func([]T0, T1) ([]R, error) {
	return func(slice []T0, scalar T1) ([]R, error) {
		if len(slice) == 0 {
			return []R{}, nil
		}

		result := make([]R, len(slice))
		for i, item := range slice {
			val, err := f(item, scalar)
			if err != nil {
				return nil, fmt.Errorf("map2first operation failed at index %d: %w", i, err)
			}
			result[i] = val
		}
		return result, nil
	}
}

/**
 * @brief Map function for binary operations where second parameter is a slice
 * @tparam T0 First input type
 * @tparam T1 Second input element type
 * @tparam R Return element type
 * @param f Binary function to apply
 * @return Function that operates on scalar and slice
 */
func Map2Second[T0, T1, R any](f func(T0, T1) (R, error)) func(T0, []T1) ([]R, error) {
	return func(scalar T0, slice []T1) ([]R, error) {
		if len(slice) == 0 {
			return []R{}, nil
		}

		result := make([]R, len(slice))
		for i, item := range slice {
			val, err := f(scalar, item)
			if err != nil {
				return nil, fmt.Errorf("map2second operation failed at index %d: %w", i, err)
			}
			result[i] = val
		}
		return result, nil
	}
}

/**
 * @brief Map function for binary operations where both parameters are slices
 * @tparam T0 First input element type
 * @tparam T1 Second input element type
 * @tparam R Return element type
 * @param f Binary function to apply
 * @return Function that operates on two slices element-wise
 */
func Map2Both[T0, T1, R any](f func(T0, T1) (R, error)) func([]T0, []T1) ([]R, error) {
	return func(slice1 []T0, slice2 []T1) ([]R, error) {
		if len(slice1) != len(slice2) {
			return nil, fmt.Errorf("slice lengths must be equal: got %d and %d", len(slice1), len(slice2))
		}

		if len(slice1) == 0 {
			return []R{}, nil
		}

		result := make([]R, len(slice1))
		for i := range slice1 {
			val, err := f(slice1[i], slice2[i])
			if err != nil {
				return nil, fmt.Errorf("map2both operation failed at index %d: %w", i, err)
			}
			result[i] = val
		}
		return result, nil
	}
}

/**
 * @brief Product function for Cartesian product operations
 * @tparam T0 First input element type
 * @tparam T1 Second input element type
 * @tparam R Return element type
 * @param f Binary function to apply
 * @return Function that computes Cartesian product
 */
func Product[T0, T1, R any](f func(T0, T1) (R, error)) func([]T0, []T1) ([]R, error) {
	return func(slice1 []T0, slice2 []T1) ([]R, error) {
		if len(slice1) == 0 || len(slice2) == 0 {
			return []R{}, nil
		}

		result := make([]R, 0, len(slice1)*len(slice2))
		for i, item1 := range slice1 {
			for j, item2 := range slice2 {
				val, err := f(item1, item2)
				if err != nil {
					return nil, fmt.Errorf("product operation failed at (%d,%d): %w", i, j, err)
				}
				result = append(result, val)
			}
		}
		return result, nil
	}
}

/**
 * @brief Map function for ternary operations where first parameter is a slice
 */
func Map3First[T0, T1, T2, R any](f func(T0, T1, T2) (R, error)) func([]T0, T1, T2) ([]R, error) {
	return func(slice []T0, scalar1 T1, scalar2 T2) ([]R, error) {
		if len(slice) == 0 {
			return []R{}, nil
		}

		result := make([]R, len(slice))
		for i, item := range slice {
			val, err := f(item, scalar1, scalar2)
			if err != nil {
				return nil, fmt.Errorf("map3first operation failed at index %d: %w", i, err)
			}
			result[i] = val
		}
		return result, nil
	}
}

/**
 * @brief Map function for ternary operations where second parameter is a slice
 */
func Map3Second[T0, T1, T2, R any](f func(T0, T1, T2) (R, error)) func(T0, []T1, T2) ([]R, error) {
	return func(scalar1 T0, slice []T1, scalar2 T2) ([]R, error) {
		if len(slice) == 0 {
			return []R{}, nil
		}

		result := make([]R, len(slice))
		for i, item := range slice {
			val, err := f(scalar1, item, scalar2)
			if err != nil {
				return nil, fmt.Errorf("map3second operation failed at index %d: %w", i, err)
			}
			result[i] = val
		}
		return result, nil
	}
}

/**
 * @brief Map function for ternary operations where third parameter is a slice
 */
func Map3Third[T0, T1, T2, R any](f func(T0, T1, T2) (R, error)) func(T0, T1, []T2) ([]R, error) {
	return func(scalar1 T0, scalar2 T1, slice []T2) ([]R, error) {
		if len(slice) == 0 {
			return []R{}, nil
		}

		result := make([]R, len(slice))
		for i, item := range slice {
			val, err := f(scalar1, scalar2, item)
			if err != nil {
				return nil, fmt.Errorf("map3third operation failed at index %d: %w", i, err)
			}
			result[i] = val
		}
		return result, nil
	}
}

/**
 * @brief Map function for ternary operations where first and second parameters are slices
 */
func Map3FirstSecond[T0, T1, T2, R any](f func(T0, T1, T2) (R, error)) func([]T0, []T1, T2) ([]R, error) {
	return func(slice1 []T0, slice2 []T1, scalar T2) ([]R, error) {
		if len(slice1) != len(slice2) {
			return nil, fmt.Errorf("slice lengths must be equal: got %d and %d", len(slice1), len(slice2))
		}

		if len(slice1) == 0 {
			return []R{}, nil
		}

		result := make([]R, len(slice1))
		for i := range slice1 {
			val, err := f(slice1[i], slice2[i], scalar)
			if err != nil {
				return nil, fmt.Errorf("map3firstsecond operation failed at index %d: %w", i, err)
			}
			result[i] = val
		}
		return result, nil
	}
}

/**
 * @brief Map function for ternary operations where first and third parameters are slices
 */
func Map3FirstThird[T0, T1, T2, R any](f func(T0, T1, T2) (R, error)) func([]T0, T1, []T2) ([]R, error) {
	return func(slice1 []T0, scalar T1, slice2 []T2) ([]R, error) {
		if len(slice1) != len(slice2) {
			return nil, fmt.Errorf("slice lengths must be equal: got %d and %d", len(slice1), len(slice2))
		}

		if len(slice1) == 0 {
			return []R{}, nil
		}

		result := make([]R, len(slice1))
		for i := range slice1 {
			val, err := f(slice1[i], scalar, slice2[i])
			if err != nil {
				return nil, fmt.Errorf("map3firstthird operation failed at index %d: %w", i, err)
			}
			result[i] = val
		}
		return result, nil
	}
}

/**
 * @brief Map function for ternary operations where second and third parameters are slices
 */
func Map3SecondThird[T0, T1, T2, R any](f func(T0, T1, T2) (R, error)) func(T0, []T1, []T2) ([]R, error) {
	return func(scalar T0, slice1 []T1, slice2 []T2) ([]R, error) {
		if len(slice1) != len(slice2) {
			return nil, fmt.Errorf("slice lengths must be equal: got %d and %d", len(slice1), len(slice2))
		}

		if len(slice1) == 0 {
			return []R{}, nil
		}

		result := make([]R, len(slice1))
		for i := range slice1 {
			val, err := f(scalar, slice1[i], slice2[i])
			if err != nil {
				return nil, fmt.Errorf("map3secondthird operation failed at index %d: %w", i, err)
			}
			result[i] = val
		}
		return result, nil
	}
}

/**
 * @brief Map function for ternary operations where all three parameters are slices
 */
func Map3All[T0, T1, T2, R any](f func(T0, T1, T2) (R, error)) func([]T0, []T1, []T2) ([]R, error) {
	return func(slice1 []T0, slice2 []T1, slice3 []T2) ([]R, error) {
		if len(slice1) != len(slice2) || len(slice1) != len(slice3) {
			return nil, fmt.Errorf("all slice lengths must be equal: got %d, %d, and %d",
				len(slice1), len(slice2), len(slice3))
		}

		if len(slice1) == 0 {
			return []R{}, nil
		}

		result := make([]R, len(slice1))
		for i := range slice1 {
			val, err := f(slice1[i], slice2[i], slice3[i])
			if err != nil {
				return nil, fmt.Errorf("map3all operation failed at index %d: %w", i, err)
			}
			result[i] = val
		}
		return result, nil
	}
}
