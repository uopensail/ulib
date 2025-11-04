package minia

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/uopensail/ulib/sample"
)

/**
 * @brief Number interface constraint for numeric types
 */
type Number interface {
	int64 | float32
}

// Error definitions
var (
	ErrEmptyVector     = errors.New("cannot operate on empty vector")
	ErrDivisionByZero  = errors.New("division by zero")
	ErrInvalidArgument = errors.New("invalid argument")
	ErrDomainError     = errors.New("domain error")
	ErrTypeMismatch    = errors.New("type mismatch") // Fixed typo: "miss match" -> "mismatch"
)

/**
 * @brief Generic function interface for operations on features
 */
type Function interface {
	Call(...sample.Feature) (sample.Feature, error)
}

// Type code mapping for reflection
var types = map[reflect.Type]string{
	reflect.TypeOf(int64(0)):    fmt.Sprintf("%d", sample.Int64Type),
	reflect.TypeOf(float32(0)):  fmt.Sprintf("%d", sample.Float32Type),
	reflect.TypeOf(""):          fmt.Sprintf("%d", sample.StringType),
	reflect.TypeOf([]int64{}):   fmt.Sprintf("%d", sample.Int64sType),
	reflect.TypeOf([]float32{}): fmt.Sprintf("%d", sample.Float32sType),
	reflect.TypeOf([]string{}):  fmt.Sprintf("%d", sample.StringsType),
}

/**
 * @brief Get type code by reflection type
 * @param t The reflection type to lookup
 * @return The corresponding DataType, or InvalidType if not found
 */
func getTypeCodeByType(t reflect.Type) string {
	if code, exists := types[t]; exists {
		return code
	}
	return "0"
}

/**
 * @brief Identity function that returns the input value unchanged
 * @tparam T Any type
 * @param a Input value
 * @return The same input value and nil error
 */
func Identity[T any](a T) (T, error) {
	return a, nil
}

/**
 * @brief Add two numeric values
 * @tparam T Numeric type (int64 or float32)
 * @param a First operand
 * @param b Second operand
 * @return Sum of a and b, nil error
 */
func Add[T Number](a, b T) (T, error) {
	return a + b, nil
}

/**
 * @brief Subtract two numeric values
 * @tparam T Numeric type (int64 or float32)
 * @param a Minuend
 * @param b Subtrahend
 * @return Difference of a and b, nil error
 */
func Sub[T Number](a, b T) (T, error) {
	return a - b, nil
}

/**
 * @brief Multiply two numeric values
 * @tparam T Numeric type (int64 or float32)
 * @param a First factor
 * @param b Second factor
 * @return Product of a and b, nil error
 */
func Mul[T Number](a, b T) (T, error) {
	return a * b, nil
}

/**
 * @brief Divide two numeric values
 * @tparam T Numeric type (int64 or float32)
 * @param a Dividend
 * @param b Divisor
 * @return Quotient of a and b, or error if b is zero
 */
func Div[T Number](a, b T) (T, error) {
	if b == 0 {
		var zero T
		return zero, ErrDivisionByZero
	}
	return a / b, nil
}

/**
 * @brief Create a mixed-type function wrapper for float32 operations
 * @tparam T0 First numeric type
 * @tparam T1 Second numeric type
 * @param f Function that operates on float32 values
 * @return Wrapped function that accepts mixed numeric types
 */
func Mix[T0, T1 Number](f func(float32, float32) (float32, error)) func(T0, T1) (float32, error) {
	return func(arg0 T0, arg1 T1) (float32, error) {
		return f(float32(arg0), float32(arg1))
	}
}

/**
 * @brief Generate a unary math function wrapper
 * @tparam T Numeric type
 * @param f Math function that operates on float64
 * @return Wrapped function that accepts numeric type and returns float32
 */
func MathUnaryFuncGenerator[T Number](f func(float64) float64) func(T) (float32, error) {
	return func(arg0 T) (float32, error) {
		return float32(f(float64(arg0))), nil
	}
}

/**
 * @brief Generate a binary math function wrapper
 * @tparam T0 First numeric type
 * @tparam T1 Second numeric type
 * @param f Math function that operates on float64 values
 * @return Wrapped function that accepts numeric types and returns float32
 */
func MathBinaryFuncGenerator[T0, T1 Number](f func(float64, float64) float64) func(T0, T1) (float32, error) {
	return func(arg0 T0, arg1 T1) (float32, error) {
		return float32(f(float64(arg0), float64(arg1))), nil
	}
}

/**
 * @brief Compute sigmoid activation function
 * @tparam T Numeric type
 * @param x Input value
 * @return Sigmoid of x (between 0 and 1), or error if result is not finite
 */
func Sigmoid[T Number](x T) (float32, error) {
	result := 1.0 / (1.0 + math.Exp(float64(-x)))
	if math.IsNaN(result) || math.IsInf(result, 0) {
		return 0, fmt.Errorf("sigmoid result is not finite")
	}
	return float32(result), nil
}

/**
 * @brief Compute softmax function for a vector
 * @tparam T Numeric type
 * @param values Input vector
 * @return Softmax probabilities, or error if input is empty
 */
func Softmax[T Number](values []T) ([]float32, error) {
	if len(values) == 0 {
		return nil, ErrEmptyVector
	}

	// Find max for numerical stability
	maxVal := values[0]
	for _, v := range values[1:] {
		if v > maxVal {
			maxVal = v
		}
	}

	// Compute exponentials and sum
	expValues := make([]float32, len(values))
	var sum float32
	for i, v := range values {
		expVal := float32(math.Exp(float64(v - maxVal)))
		expValues[i] = expVal
		sum += expVal
	}

	// Normalize
	for i := range expValues {
		expValues[i] /= sum
	}

	return expValues, nil
}

/**
 * @brief Compute log-softmax function for a vector
 * @tparam T Numeric type
 * @param values Input vector
 * @return Log-softmax values, or error if input is empty
 */
func LogSoftmax[T Number](values []T) ([]float32, error) {
	if len(values) == 0 {
		return nil, ErrEmptyVector
	}

	// Find max for numerical stability
	maxVal := values[0]
	for _, v := range values[1:] {
		if v > maxVal {
			maxVal = v
		}
	}

	// Compute log-sum-exp
	var sumExp float64
	for _, v := range values {
		sumExp += math.Exp(float64(v - maxVal))
	}
	logSumExp := float64(maxVal) + math.Log(sumExp)

	// Compute log-softmax
	result := make([]float32, len(values))
	for i, v := range values {
		result[i] = float32(float64(v) - logSumExp)
	}

	return result, nil
}

/**
 * @brief Calculate median of a numeric vector
 * @tparam T Numeric type
 * @param src Input vector
 * @return Median value, or error if input is empty
 */
func Median[T Number](src []T) (float32, error) {
	if len(src) == 0 {
		return 0, ErrEmptyVector
	}

	// Create a copy and sort it
	sorted := make([]T, len(src))
	copy(sorted, src)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	n := len(sorted)
	if n%2 == 0 {
		// Even number of elements - average of middle two
		return (float32(sorted[n/2-1]) + float32(sorted[n/2])) / 2, nil
	}
	// Odd number of elements - middle element
	return float32(sorted[n/2]), nil
}

/**
 * @brief Cast int64 to float32
 * @param v Input int64 value
 * @return Converted float32 value
 */
func CastToFloat(v int64) (float32, error) {
	return float32(v), nil
}

/**
 * @brief Clamp a value between min and max bounds
 * @tparam T Numeric type
 * @param value Input value to clamp
 * @param min Minimum bound
 * @param max Maximum bound
 * @return Clamped value, or error if min > max
 */
func Clamp[T0, T1, T2 Number](value T0, min T1, max T2) (float32, error) {
	if float32(min) > float32(max) {
		return float32(value), fmt.Errorf("min cannot be greater than max")
	}

	if float32(value) < float32(min) {
		return float32(min), nil
	}
	if float32(value) > float32(max) {
		return float32(max), nil
	}
	return float32(value), nil
}

/**
 * @brief Modulus operation with Euclidean behavior
 * @param a Dividend
 * @param b Divisor
 * @return Euclidean modulus result, or error if b is zero or negative
 */
func Mod(a, b int64) (int64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	if b < 0 {
		return 0, fmt.Errorf("negative divisor not supported")
	}

	result := a % b
	if result < 0 {
		result += b
	}
	return result, nil
}

/**
 * @brief Find minimum value in a vector
 * @tparam T Numeric type
 * @param src Input vector
 * @return Minimum value, or error if vector is empty
 */
func Min[T Number](src []T) (T, error) {
	if len(src) == 0 {
		var zero T
		return zero, ErrEmptyVector
	}

	ret := src[0]
	for _, v := range src[1:] {
		if v < ret {
			ret = v
		}
	}
	return ret, nil
}

/**
 * @brief Sum of values in a vector
 * @tparam T Numeric type
 * @param src Input vector
 * @return Sum value, or error if vector is empty
 */
func Sum[T Number](src []T) (T, error) {
	if len(src) == 0 {
		var zero T
		return zero, ErrEmptyVector
	}

	ret := src[0]
	for _, v := range src[1:] {
		ret += v
	}
	return ret, nil
}

/**
 * @brief Find maximum value in a vector
 * @tparam T Numeric type
 * @param src Input vector
 * @return Maximum value, or error if vector is empty
 */
func Max[T Number](src []T) (T, error) {
	if len(src) == 0 {
		var zero T
		return zero, ErrEmptyVector
	}

	ret := src[0]
	for _, v := range src[1:] {
		if v > ret {
			ret = v
		}
	}
	return ret, nil
}

/**
 * @brief Calculate average of a vector
 * @tparam T Numeric type
 * @param src Input vector
 * @return Average value, or error if vector is empty
 */
func Average[T Number](src []T) (float32, error) {
	if len(src) == 0 {
		return 0, ErrEmptyVector
	}

	var sum T
	for _, v := range src {
		sum += v
	}
	return float32(sum) / float32(len(src)), nil
}

/**
 * @brief Calculate variance of a vector
 * @tparam T Numeric type
 * @param src Input vector
 * @return Population variance, or error if calculation fails
 */
func Variance[T Number](src []T) (float32, error) {
	if len(src) <= 1 {
		return 0.0, nil
	}

	var sum, sqSum T
	for _, val := range src {
		sum += val
		sqSum += val * val
	}

	n := float64(len(src))
	mean := float64(sum) / n
	result := (float64(sqSum) - mean*mean*n) / n

	if math.IsNaN(result) || math.IsInf(result, 0) {
		return 0, fmt.Errorf("variance calculation resulted in non-finite value")
	}

	return float32(result), nil
}

/**
 * @brief Calculate standard deviation of a vector
 * @tparam T Numeric type
 * @param src Input vector
 * @return Standard deviation, or error if calculation fails
 */
func Stddev[T Number](src []T) (float32, error) {
	variance, err := Variance(src)
	if err != nil {
		return 0, err
	}

	if variance > 0 {
		result := math.Sqrt(float64(variance))
		if math.IsNaN(result) || math.IsInf(result, 0) {
			return 0, fmt.Errorf("standard deviation calculation resulted in non-finite value")
		}
		return float32(result), nil
	}
	return 0.0, nil
}

/**
 * @brief Normalize vector using L-p norm
 * @tparam T Numeric type
 * @param src Input vector
 * @param norm Norm degree (must be >= 1)
 * @return Normalized vector, or error if norm < 1 or vector is zero
 */
func Normalize[T Number](src []T, norm float32) ([]float32, error) {
	if norm < 1 {
		return nil, fmt.Errorf("norm degree must be >= 1")
	}

	// Calculate norm value
	var sum float64
	for _, v := range src {
		sum += math.Pow(math.Abs(float64(v)), float64(norm))
	}
	normValue := math.Pow(sum, 1.0/float64(norm))

	if normValue == 0 {
		return nil, fmt.Errorf("cannot normalize zero vector")
	}

	if math.IsNaN(normValue) || math.IsInf(normValue, 0) {
		return nil, fmt.Errorf("norm value is not finite")
	}

	// Normalize
	dst := make([]float32, len(src))
	for i, v := range src {
		dst[i] = float32(float64(v) / normValue)
	}
	return dst, nil
}

/**
 * @brief Get top K elements from a slice
 * @tparam T Any type
 * @param src Input slice (should be pre-sorted if order matters)
 * @param k Number of elements to return
 * @return First k elements, or error if k is negative
 */
func TopK[T any](src []T, k int64) ([]T, error) {
	if k < 0 {
		return nil, fmt.Errorf("k must be non-negative")
	}
	if k > int64(len(src)) {
		k = int64(len(src))
	}

	result := make([]T, k)
	copy(result, src[:k])
	return result, nil
}

/**
 * @brief Count occurrences of a value in a slice
 * @tparam T Comparable type
 * @param list Input slice
 * @param v Value to count
 * @return Number of occurrences
 */
func Count[T comparable](list []T, v T) (int64, error) {
	var count int64
	for _, item := range list {
		if item == v {
			count++
		}
	}
	return count, nil
}

/**
 * @brief Get length of a slice
 * @tparam T Any type
 * @param list Input slice
 * @return Length of the slice
 */
func Len[T any](list []T) (int64, error) {
	return int64(len(list)), nil
}

/**
 * @brief Min-max normalization
 * @tparam T Numeric type
 * @param v Value to normalize
 * @param min Minimum value in the range
 * @param max Maximum value in the range
 * @return Normalized value between 0 and 1, or error if min == max
 */
func MinMax[T Number](v, min, max T) (float32, error) {
	if min == max {
		return 0, fmt.Errorf("min and max cannot be equal")
	}
	return float32(v-min) / float32(max-min), nil
}

/**
 * @brief Binarize a value based on threshold
 * @tparam T Numeric type
 * @param v Input value
 * @param threshold Threshold for binarization
 * @return 1 if v >= threshold, 0 otherwise
 */
func Binarize[T Number](v, threshold T) (int64, error) {
	if v < threshold {
		return 0, nil
	}
	return 1, nil
}

/**
 * @brief Bucketize a value based on boundaries
 * @tparam T Numeric type
 * @param v Input value
 * @param boundaries Sorted boundary values
 * @return Bucket index (0-based)
 */
func Bucketize[T Number](v T, boundaries []T) (int64, error) {
	// Find the first boundary greater than v
	for i, boundary := range boundaries {
		if v <= boundary {
			return int64(i), nil
		}
	}
	return int64(len(boundaries)), nil
}

/**
 * @brief Z-score normalization
 * @param value Input value
 * @param mean Mean of the distribution
 * @param stdDev Standard deviation of the distribution
 * @return Z-score, or error if stdDev <= 0 or result is not finite
 */
func ZScore(value, mean, stdDev float32) (float32, error) {
	if stdDev <= 0 {
		return 0, fmt.Errorf("standard deviation must be positive")
	}

	result := float64((value - mean) / stdDev)
	if math.IsNaN(result) || math.IsInf(result, 0) {
		return 0, fmt.Errorf("z-score calculation resulted in non-finite value")
	}

	return float32(result), nil
}

/**
 * @brief Get weekday from date string
 * @param dateStr Date string in "2006-01-02" format
 * @return Weekday name, or error if parsing fails
 */
func Weekday(dateStr string) (string, error) {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse date: %w", err)
	}
	return t.Weekday().String(), nil
}

/**
 * @brief Get number of days in a month
 * @param year Year value
 * @param month Month value (1-12)
 * @return Number of days in the month, or error if month is invalid
 */
func DaysOfMonth(year, month int64) (int64, error) {
	if month < 1 || month > 12 {
		return 0, fmt.Errorf("month must be between 1 and 12")
	}

	t := time.Date(int(year), time.Month(month+1), 0, 0, 0, 0, 0, time.UTC)
	return int64(t.Day()), nil
}

/**
 * @brief Convert numeric value to string
 * @tparam T Numeric type
 * @param v Input value
 * @return String representation of the value
 */
func ToString[T Number](v T) (string, error) {
	return fmt.Sprintf("%v", v), nil
}

/**
 * @brief Split string by separator
 * @param s Input string
 * @param separator Separator string
 * @return Array of split strings
 */
func SplitString(s, separator string) ([]string, error) {
	return strings.Split(s, separator), nil
}

/**
 * @brief Logical AND operation
 * @param a First operand (0 is false, non-zero is true)
 * @param b Second operand (0 is false, non-zero is true)
 * @return 1 if both are true, 0 otherwise
 */
func And(a, b int64) (int64, error) {
	if a != 0 && b != 0 {
		return 1, nil
	}
	return 0, nil
}

/**
 * @brief Get the value of π (pi)
 * @return Value of π as float32
 */
func Pi() (float32, error) {
	return float32(math.Pi), nil
}

/**
 * @brief Get the value of e (Euler's number)
 * @return Value of e as float32
 */
func E() (float32, error) {
	return float32(math.E), nil
}

/**
 * @brief Get the value of φ (golden ratio)
 * @return Value of φ as float32
 */
func Phi() (float32, error) {
	return float32((1.0 + math.Sqrt(5.0)) / 2.0), nil
}

/**
 * @brief Get the value of √2
 * @return Value of √2 as float32
 */
func Sqrt2() (float32, error) {
	return float32(math.Sqrt2), nil
}

/**
 * @brief Get the value of ln(2)
 * @return Value of ln(2) as float32
 */
func Ln2() (float32, error) {
	return float32(math.Ln2), nil
}

/**
 * @brief Get the value of ln(10)
 * @return Value of ln(10) as float32
 */
func Ln10() (float32, error) {
	return float32(math.Ln10), nil
}

/**
 * @brief Get sign of a numeric value
 * @tparam T Numeric type
 * @param value Input value
 * @return -1 if negative, 1 if positive, 0 if zero
 */
func Sign[T Number](value T) (int64, error) {
	if value < 0 {
		return -1, nil
	}
	if value > 0 {
		return 1, nil
	}
	return 0, nil
}

// StringLength returns the length of a string in characters (Unicode-safe).
//
// @param s Input string
// @return Length in characters
func StringLength(s string) (int64, error) {
	return int64(len([]rune(s))), nil
}

// StringContains checks if a string contains a substring.
//
// @param s Input string
// @param substr Substring to search for
// @return 1 if contains, 0 otherwise
func StringContains(s, substr string) (int64, error) {
	if strings.Contains(s, substr) {
		return 1, nil
	}
	return 0, nil
}

/**
 * @brief Check if string starts with prefix
 * @param s Input string
 * @param prefix Prefix to check
 * @return 1 if starts with prefix, 0 otherwise
 */
func StartsWith(s, prefix string) (int64, error) {
	if strings.HasPrefix(s, prefix) {
		return 1, nil
	}
	return 0, nil
}

/**
 * @brief Check if string ends with suffix
 * @param s Input string
 * @param suffix Suffix to check
 * @return 1 if ends with suffix, 0 otherwise
 */
func EndsWith(s, suffix string) (int64, error) {
	if strings.HasSuffix(s, suffix) {
		return 1, nil
	}
	return 0, nil
}

/**
 * @brief Find index of substring in string
 * @param s Input string
 * @param substr Substring to find
 * @return Index of first occurrence, or -1 if not found
 */
func IndexOf(s, substr string) (int64, error) {
	index := strings.Index(s, substr)
	return int64(index), nil
}

/**
 * @brief Logical OR operation
 * @param a First operand (0 is false, non-zero is true)
 * @param b Second operand (0 is false, non-zero is true)
 * @return 1 if either is true, 0 otherwise
 */
func Or(a, b int64) (int64, error) {
	if a != 0 || b != 0 {
		return 1, nil
	}
	return 0, nil
}

/**
 * @brief Logical NOT operation
 * @param v Input value (0 is false, non-zero is true)
 * @return 1 if input is false, 0 if input is true
 */
func Not(v int64) (int64, error) {
	if v == 0 {
		return 1, nil
	}
	return 0, nil
}

/**
 * @brief Check if slice contains a value
 * @tparam T Comparable type
 * @param list Input slice
 * @param v Value to search for
 * @return 1 if found, 0 otherwise
 */
func Contains[T comparable](list []T, v T) (int64, error) {
	if slices.Contains(list, v) {
		return 1, nil
	}
	return 0, nil
}

/**
 * @brief Check equality of two values
 * @tparam T Comparable type
 * @param a First value
 * @param b Second value
 * @return 1 if equal, 0 otherwise
 */
func EqualSameType[T comparable](a, b T) (int64, error) {
	if a == b {
		return 1, nil
	}
	return 0, nil
}

func Equal[T0, T1 comparable](a T0, b T1) (int64, error) {
	if reflect.DeepEqual(a, b) {
		return 1, nil
	} else {
		return 0, nil
	}
}

/**
 * @brief Check inequality of two values
 * @tparam T Comparable type
 * @param a First value
 * @param b Second value
 * @return 1 if not equal, 0 otherwise
 */
func NotEqualSameType[T comparable](a, b T) (int64, error) {
	if a != b {
		return 1, nil
	}
	return 0, nil
}

func NotEqual[T0, T1 comparable](a T0, b T1) (int64, error) {
	if reflect.DeepEqual(a, b) {
		return 0, nil
	} else {
		return 1, nil
	}
}

/**
 * @brief Check if first value is less than second
 * @tparam T Numeric type
 * @param a First value
 * @param b Second value
 * @return 1 if a < b, 0 otherwise
 */
func LessThanSameType[T Number](a, b T) (int64, error) {
	if a < b {
		return 1, nil
	}
	return 0, nil
}

func LessThan[T0, T1 Number](a T0, b T1) (int64, error) {
	if float32(a) < float32(b) {
		return 1, nil
	}
	return 0, nil
}

/**
 * @brief Check if first value is less than or equal to second
 * @tparam T Numeric type
 * @param a First value
 * @param b Second value
 * @return 1 if a <= b, 0 otherwise
 */
func LessThanEqualSameType[T Number](a, b T) (int64, error) {
	if a < b {
		return 1, nil
	}
	return 0, nil
}

func LessThanEqual[T0, T1 Number](a T0, b T1) (int64, error) {
	if float32(a) <= float32(b) {
		return 1, nil
	}
	return 0, nil
}

/**
 * @brief Check if first value is greater than second
 * @tparam T Numeric type
 * @param a First value
 * @param b Second value
 * @return 1 if a > b, 0 otherwise
 */
func GreaterThanSameType[T Number](a, b T) (int64, error) {
	if a > b {
		return 1, nil
	}
	return 0, nil
}

func GreaterThan[T0, T1 Number](a T0, b T1) (int64, error) {
	if float32(a) > float32(b) {
		return 1, nil
	}
	return 0, nil
}

/**
 * @brief Check if first value is greater than or equal to second
 * @tparam T Numeric type
 * @param a First value
 * @param b Second value
 * @return 1 if a >= b, 0 otherwise
 */
func GreaterThanEqualSameType[T Number](a, b T) (int64, error) {
	if a > b {
		return 1, nil
	}
	return 0, nil
}

func GreaterThanEqual[T0, T1 Number](a T0, b T1) (int64, error) {
	if float32(a) >= float32(b) {
		return 1, nil
	}
	return 0, nil
}

/**
 * @brief Format time using specified format string
 * @param t Time to format
 * @param format Format string (C-style, will be converted to Go format)
 * @return Formatted time string
 */
func FormatDate(t time.Time, format string) (string, error) {
	if format == "" {
		format = "2006-01-02"
	}
	// Convert C-style format to Go format
	goFormat := convertDateFormat(format)
	return t.Format(goFormat), nil
}

/**
 * @brief Convert C-style date format to Go format
 * @param cFormat C-style format string
 * @return Go-style format string
 */
func convertDateFormat(cFormat string) string {
	replacements := map[string]string{
		"%Y": "2006",
		"%m": "01",
		"%d": "02",
		"%H": "15",
		"%M": "04",
		"%S": "05",
	}

	result := cFormat
	for old, new := range replacements {
		result = strings.ReplaceAll(result, old, new)
	}
	return result
}

/**
 * @brief Get current year as string
 * @return Current year in "2006" format
 */
func Year() (string, error) {
	return time.Now().Format("2006"), nil
}

/**
 * @brief Get current month as string
 * @return Current month in "01" format
 */
func Month() (string, error) {
	return time.Now().Format("01"), nil
}

/**
 * @brief Get current day as string
 * @return Current day in "02" format
 */
func Day() (string, error) {
	return time.Now().Format("02"), nil
}

/**
 * @brief Get current date as string
 * @return Current date in "2006-01-02" format
 */
func CurDate() (string, error) {
	return time.Now().Format("2006-01-02"), nil
}

/**
 * @brief Get current Unix timestamp
 * @return Unix timestamp (seconds since epoch)
 */
func UnixTimestamp() (int64, error) {
	return time.Now().Unix(), nil
}

/**
 * @brief Format Unix timestamp to date string
 * @param timestamp Unix timestamp
 * @param format Format string (C-style, will be converted to Go format)
 * @return Formatted date string
 */
func FromUnixTime(timestamp int64, format string) (string, error) {
	t := time.Unix(timestamp, 0)
	goFormat := convertDateFormat(format)
	return t.Format(goFormat), nil
}

/**
 * @brief Add time interval to a date
 * @param dateStr Date string in "2006-01-02" format
 * @param interval Number of units to add
 * @param unit Time unit ("day", "days", "month", "months", "year", "years")
 * @return New date string, or error if parsing fails or unit is unsupported
 */
func DateAdd(dateStr string, interval int64, unit string) (string, error) {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse date: %w", err)
	}

	switch strings.ToLower(unit) {
	case "day", "days":
		t = t.AddDate(0, 0, int(interval))
	case "month", "months":
		t = t.AddDate(0, int(interval), 0)
	case "year", "years":
		t = t.AddDate(int(interval), 0, 0)
	default:
		return "", fmt.Errorf("unsupported unit: %s", unit)
	}

	return t.Format("2006-01-02"), nil
}

/**
 * @brief Subtract time interval from a date
 * @param dateStr Date string in "2006-01-02" format
 * @param interval Number of units to subtract
 * @param unit Time unit ("day", "days", "month", "months", "year", "years")
 * @return New date string, or error if parsing fails or unit is unsupported
 */
func DateSub(dateStr string, interval int64, unit string) (string, error) {
	return DateAdd(dateStr, -interval, unit)
}

/**
 * @brief Calculate difference between two dates in days
 * @param date1 First date string in "2006-01-02" format
 * @param date2 Second date string in "2006-01-02" format
 * @return Difference in days (date1 - date2), or error if parsing fails
 */
func DateDiff(date1, date2 string) (int64, error) {
	t1, err := time.Parse("2006-01-02", date1)
	if err != nil {
		return 0, fmt.Errorf("failed to parse date1: %w", err)
	}

	t2, err := time.Parse("2006-01-02", date2)
	if err != nil {
		return 0, fmt.Errorf("failed to parse date2: %w", err)
	}

	diff := t1.Sub(t2)
	return int64(diff.Hours() / 24), nil
}

/**
 * @brief Reverse a string (Unicode-safe)
 * @param s Input string
 * @return Reversed string
 */
func Reverse(s string) (string, error) {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes), nil
}

/**
 * @brief Convert string to uppercase
 * @param s Input string
 * @return Uppercase string
 */
func Upper(s string) (string, error) {
	return strings.ToUpper(s), nil
}

/**
 * @brief Convert string to lowercase
 * @param s Input string
 * @return Lowercase string
 */
func Lower(s string) (string, error) {
	return strings.ToLower(s), nil
}

/**
 * @brief Extract substring from string (Unicode-safe)
 * @param s Input string
 * @param start Starting index (0-based)
 * @param length Number of characters to extract
 * @return Substring, or error if parameters are invalid
 */
func Substr(s string, start, length int64) (string, error) {
	if start < 0 {
		return "", fmt.Errorf("start index cannot be negative")
	}
	if length < 0 {
		return "", fmt.Errorf("length cannot be negative")
	}

	runes := []rune(s)

	if start >= int64(len(runes)) {
		return "", nil
	}

	end := start + length
	if end > int64(len(runes)) {
		end = int64(len(runes))
	}

	return string(runes[start:end]), nil
}

/**
 * @brief Concatenate two strings
 * @param a First string
 * @param b Second string
 * @return Concatenated string
 */
func Concat(a, b string) (string, error) {
	return a + b, nil
}

/**
 * @brief Create a new Feature from a value
 * @param value Input value of supported types
 * @return Feature wrapper, or error if type is not supported
 */
func NewFeature(value any) (sample.Feature, error) {
	switch val := value.(type) {
	case int64:
		return &sample.Int64{Value: val}, nil
	case float32:
		return &sample.Float32{Value: val}, nil
	case string:
		return &sample.String{Value: val}, nil
	case []int64:
		return &sample.Int64s{Value: val}, nil
	case []float32:
		return &sample.Float32s{Value: val}, nil
	case []string:
		return &sample.Strings{Value: val}, nil
	default:
		return nil, ErrTypeMismatch
	}
}

/**
 * @brief Generate function signature string for reflection-based function registration
 * @details Creates a signature in format "name:argCount=[type1,type2,...]" where types
 *          are represented by their corresponding DataType codes
 * @param name Function name identifier
 * @param fn Function interface to analyze via reflection
 * @return Signature string, empty string if fn is not a function
 * @example generateSignature("Add", Add[int64]) returns "Add:2=[1,1]"
 */
func generateSignature(name string, fn any) string {
	// Validate input parameters
	if name == "" || fn == nil {
		return ""
	}

	fnType := reflect.TypeOf(fn)
	if fnType.Kind() != reflect.Func {
		return ""
	}

	argCount := fnType.NumIn()

	// Pre-allocate slice for better performance
	typeCodes := make([]string, 0, argCount)

	// Collect type codes
	for i := range argCount {
		argType := fnType.In(i)
		typeCode := getTypeCodeByType(argType)
		typeCodes = append(typeCodes, typeCode)
	}
	return fmt.Sprintf("%s:%d=[%s]", name, argCount, strings.Join(typeCodes, ","))
}
