package minia

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/uopensail/ulib/sample"
)

// ========== Benchmark Test Cases ==========

/**
 * @brief Benchmark complex mathematical expressions with nested operations
 * @param b Benchmark testing context
 * @details Tests arithmetic operations, function calls, and nested computations
 */
func BenchmarkMiniaComplexMathExpressions(b *testing.B) {
	expressions := []string{
		// Arithmetic operation combinations
		"result1 = ((a * b) + (c / d)) - (e / f)",
		"result2 = ((a + b) * (c - d)) / (e + 0.001)",
		"result3 = ((a * a) + (b * b)) + (c * c)",

		// Function calls
		"sqrt_result = sqrt((a * a) + (b * b))",
		"trig_result = sin(a) + (cos(b) * tan(c))",
		"log_result = log(abs(a) + 1) + exp(b / 10)",
		"power_result = pow(a, 2) + pow(b, 3)",

		// Nested function calls
		"nested1 = sqrt(pow(a, 2) + pow(b, 2))",
		"nested2 = log(exp(a) + exp(b))",
		"nested3 = abs(sin(a) - cos(b))",

		// Complex combinations
		"complex1 = (sqrt_result * trig_result) + log_result",
		"complex2 = (nested1 + nested2) / (nested3 + 0.1)",
		"final = ((complex1 * complex2) - result1) + result2",
	}

	m := NewMinia(expressions)
	ret := sample.NewMutableFeatures()
	ret.SetValue("a", 3.14159)
	ret.SetValue("b", 2.71828)
	ret.SetValue("c", 1.41421)
	ret.SetValue("d", 1.73205)
	ret.SetValue("e", 0.57721)
	ret.SetValue("f", 2.23606)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Eval(ret)
	}
}

/**
 * @brief Benchmark logical expressions with comparison and boolean operations
 * @param b Benchmark testing context
 * @details Tests comparison operators, logical operations, and mixed expressions
 */
func BenchmarkMiniaLogicalExpressions(b *testing.B) {
	expressions := []string{
		// Comparison operations
		"gt_result = (a > b)",
		"lt_result = (c < d)",
		"gte_result = (a >= c)",
		"lte_result = (b <= d)",
		"eq_result = (a == b)",
		"neq_result = (c != d)",

		// Logical operations
		"and_result = gt_result & lt_result",
		"or_result = eq_result | neq_result",
		"not_result = !and_result",

		// Complex logical combinations
		"complex_logic1 = ((a > b) & (c < d)) | (a == c)",
		"complex_logic2 = !((a > b) & (c < d)) | ((a >= c) & (b <= d))",
		"complex_logic3 = (((a > b) | (c > d)) & !(((a == b) | (c == d))))",

		// Mixed numerical and logical operations
		"mixed1 = (a * 1.0) + (b * 2.0)",
		"mixed2 = (a + b) * 2.0",
	}

	m := NewMinia(expressions)
	ret := sample.NewMutableFeatures()
	ret.SetValue("a", 5.5)
	ret.SetValue("b", 3.2)
	ret.SetValue("c", 7.8)
	ret.SetValue("d", 4.1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fea := m.Eval(ret)
		v, _ := fea.MarshalJSON()
		fmt.Println(string(v))
	}
}

/**
 * @brief Benchmark list operations including string, integer, and decimal arrays
 * @param b Benchmark testing context
 * @details Tests array operations, aggregation functions, and list manipulations
 */
func BenchmarkMiniaListOperations(b *testing.B) {
	expressions := []string{
		// String list operations
		`str_list = ["apple", "banana", "cherry"]`,
		`str_result = contains(str_list, "banana")`,
		`str_length = len(str_list)`,

		// Integer list operations
		`int_list = [1, 2, 3, 4, 5]`,
		`int_sum = sum(int_list)`,
		`int_avg = average(int_list)`,
		`int_max = max(int_list)`,
		`int_min = min(int_list)`,

		// Decimal list operations
		`decimal_list = [1.1, 2.2, 3.3, 4.4, 5.5]`,
		`decimal_sum = sum(decimal_list)`,
		`decimal_avg = average(decimal_list)`,

		// Mixed list and variable operations
		`mixed_calc = sum(int_list) * average(decimal_list)`,
		`list_comparison = (len(str_list) > len(int_list))`,
	}

	m := NewMinia(expressions)
	ret := sample.NewMutableFeatures()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Eval(ret)
	}
}

/**
 * @brief Benchmark large-scale nested dependencies with multi-layer computations
 * @param b Benchmark testing context
 * @details Tests performance with complex dependency graphs and deep nesting
 */
func BenchmarkMiniaLargeScaleNested(b *testing.B) {
	var expressions []string

	// First layer: basic computations
	for i := range 20 {
		expressions = append(expressions,
			fmt.Sprintf("base_%d = (input_%d * 2.5) + (input_%d / 1.7)", i, i%5, (i+1)%5))
	}

	// Second layer: computations based on first layer
	for i := range 15 {
		expressions = append(expressions,
			fmt.Sprintf("level2_%d = (base_%d + base_%d) - (base_%d * 0.3)",
				i, i%20, (i+5)%20, (i+10)%20))
	}

	// Third layer: more complex combinations
	for i := range 10 {
		expressions = append(expressions,
			fmt.Sprintf("level3_%d = sqrt(pow(level2_%d, 2) + pow(level2_%d, 2))",
				i, i%15, (i+7)%15))
	}

	// Final layer: aggregated results
	expressions = append(expressions,
		"final_sum = (((level3_0 + level3_1) + level3_2) + level3_3) + level3_4")
	expressions = append(expressions,
		"final_avg = final_sum / 5.0")
	expressions = append(expressions,
		"final_result = final_avg * sqrt(final_sum)")

	m := NewMinia(expressions)
	ret := sample.NewMutableFeatures()

	// Set input variables
	for i := 0; i < 5; i++ {
		ret.SetValue(fmt.Sprintf("input_%d", i), float64(i+1)*1.5)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Eval(ret)
	}
}

/**
 * @brief Benchmark string operations including concatenation and manipulation
 * @param b Benchmark testing context
 * @details Tests string functions, case conversion, and string-numeric mixing
 */
func BenchmarkMiniaStringOperations(b *testing.B) {
	expressions := []string{
		// Basic string operations
		`name = "John Doe"`,
		`greeting = concat("Hello, ", name)`,
		`upper_name = upper(name)`,
		`lower_greeting = lower(greeting)`,
		`name_length = len(name)`,

		// String functions
		`contains_john = contains(name, "John")`,
		`starts_hello = startswith(greeting, "Hello")`,
		`ends_doe = endswith(name, "Doe")`,

		// String and numeric mixing
		`name_score = len(name) * 10.5`,
		`is_long_name = (len(name) > 5)`,
		`formatted_score = tostring(name_score)`,
	}

	m := NewMinia(expressions)
	ret := sample.NewMutableFeatures()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Eval(ret)
	}
}

/**
 * @brief Benchmark real-world business scenarios (risk scoring system)
 * @param b Benchmark testing context
 * @details Tests practical business logic with scoring, normalization, and decision making
 */
func BenchmarkMiniaBusinessScenarios(b *testing.B) {
	// Risk scoring scenario
	expressions := []string{
		// Basic feature normalization
		"age_score = (age - 18) / 47.0",                 // Normalize age (18-65)
		"income_score = log(income + 1) / log(1000000)", // Income log normalization
		"credit_score = credit_rating / 850.0",          // Credit score normalization

		// Risk factors (simplified version without ternary operators)
		"age_risk_check = (age_score > 0.8)",
		"income_risk_check = (income_score < 0.3)",
		"credit_risk_check = (credit_score < 0.6)",

		// Combined scoring
		"base_score = ((age_score * 0.2) + (income_score * 0.4)) + (credit_score * 0.4)",
		"risk_count = (age_risk_check + income_risk_check) + credit_risk_check",
		"risk_penalty = risk_count * 0.1",
		"adjusted_score = base_score - risk_penalty",

		// Final decision making
		"approval_score = clamp(adjusted_score, 0.0, 1.0)",
		"is_approved = (approval_score > 0.6)",
		"base_loan = income * 5",
		"max_loan = clamp(base_loan, 0, 500000)",
	}

	m := NewMinia(expressions)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := sample.NewMutableFeatures()
		// Generate random customer data
		ret.SetValue("age", float64(rand.Intn(47)+18))             // Age 18-65
		ret.SetValue("income", float64(rand.Intn(200000)+30000))   // Income 30k-230k
		ret.SetValue("credit_rating", float64(rand.Intn(350)+500)) // Credit 500-850

		m.Eval(ret)
	}
}

/**
 * @brief Benchmark deeply nested parenthetical expressions
 * @param b Benchmark testing context
 * @details Tests parser and evaluator performance with complex nested structures
 */
func BenchmarkMiniaDeepNesting(b *testing.B) {
	expressions := []string{
		// Deeply nested mathematical expressions
		"deep1 = (((a + b) * (c - d)) / ((e + f) * (g - h)))",
		"deep2 = sqrt(((a * a) + (b * b)) + ((c * c) + (d * d)))",
		"deep3 = log(exp((a + b) / 2) + exp((c + d) / 2))",

		// Complex logical nesting
		"logic1 = (((a > b) & (c > d)) | ((e > f) & (g > h)))",
		"logic2 = !(((a == b) | (c == d)) & ((e == f) | (g == h)))",

		// Function nesting
		"nested_func1 = sin(cos(tan(a)))",
		"nested_func2 = sqrt(abs(log(exp(b))))",
		"nested_func3 = pow(sqrt(a), log(b))",

		// Final composite calculation
		"final_complex = (((deep1 + deep2) * deep3)) / ((nested_func1 + nested_func2) + (nested_func3 + 0.001))",
	}

	m := NewMinia(expressions)
	ret := sample.NewMutableFeatures()
	ret.SetValue("a", 2.5)
	ret.SetValue("b", 3.7)
	ret.SetValue("c", 1.8)
	ret.SetValue("d", 4.2)
	ret.SetValue("e", 0.9)
	ret.SetValue("f", 2.1)
	ret.SetValue("g", 3.3)
	ret.SetValue("h", 1.6)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Eval(ret)
	}
}

/**
 * @brief Benchmark memory allocation patterns with different expression types
 * @param b Benchmark testing context
 * @details Tests memory efficiency across various expression patterns
 */
func BenchmarkMiniaMemoryPatterns(b *testing.B) {
	expressions := []string{
		// Literal-heavy expressions (should have fewer allocations)
		"literal1 = 42 + 3.14159",
		"literal2 = (100 * 2) / 4",
		"literal3 = sqrt(16) + pow(2, 3)",

		// Variable-heavy expressions
		"var1 = (a * b) + (c * d)",
		"var2 = (sqrt(a) + log(b)) + exp(c)",
		"var3 = (sin(a) * cos(b)) * tan(c)",

		// Mixed expressions
		"mixed1 = (a + 10) * (b - 5.5)",
		"mixed2 = sqrt(a * a + 25) / 3.0",
		"mixed3 = log(a + 1) + 2.71828",

		// List operations (higher allocation)
		"list1 = [1, 2, 3, 4, 5]",
		"list2 = average(list1) * 2.0",
		"list3 = max(list1) - min(list1)",
	}

	m := NewMinia(expressions)
	ret := sample.NewMutableFeatures()
	ret.SetValue("a", 3.0)
	ret.SetValue("b", 4.0)
	ret.SetValue("c", 5.0)
	ret.SetValue("d", 6.0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Eval(ret)
	}
}

// ========== Validation Test Cases ==========

/**
 * @brief Test case structure for validation tests
 */
type ValidationTestCase struct {
	Name        string                 // Test case name
	Expressions []string               // Expression strings to evaluate
	Inputs      map[string]interface{} // Input variable values
	Expected    map[string]float64     // Expected output values
	Tolerance   float64                // Acceptable error tolerance
}

/**
 * @brief Comprehensive performance test with result validation
 * @param t Testing context
 * @details Tests various scenarios and validates correctness of results
 */
func TestMiniaComplexScenariosWithValidation(t *testing.T) {
	testCases := []ValidationTestCase{
		{
			Name: "Mathematical Operations Validation",
			Expressions: []string{
				"a_plus_b = (a + b)",
				"a_times_b = (a * b)",
				"complex = (a + b) * (a - b)", // Should equal a²-b²
				"sqrt_sum = sqrt((a * a) + (b * b))",
			},
			Inputs: map[string]interface{}{
				"a": 3.0,
				"b": 4.0,
			},
			Expected: map[string]float64{
				"a_plus_b":  7.0,
				"a_times_b": 12.0,
				"complex":   -7.0, // 3²-4² = 9-16 = -7
				"sqrt_sum":  5.0,  // √(9+16) = √25 = 5
			},
			Tolerance: 0.0001,
		},
		{
			Name: "Logical Operations Validation",
			Expressions: []string{
				"is_greater = (a > b)",
				"is_equal = (a == b)",
				"logical_and = (a > 0) & (b > 0)",
				"logical_or = (a > 10) | (b > 10)",
			},
			Inputs: map[string]interface{}{
				"a": 5.0,
				"b": 3.0,
			},
			Expected: map[string]float64{
				"is_greater":  1.0, // true
				"is_equal":    0.0, // false
				"logical_and": 1.0, // true
				"logical_or":  0.0, // false
			},
			Tolerance: 0.0001,
		},
		{
			Name: "Parentheses Priority Validation",
			Expressions: []string{
				"no_paren = a + b * c",     // Should be a + (b * c)
				"with_paren = (a + b) * c", // Should be (a + b) * c
				"complex_paren = ((a + b) * c) - (d / e)",
			},
			Inputs: map[string]interface{}{
				"a": 2.0,
				"b": 3.0,
				"c": 4.0,
				"d": 10.0,
				"e": 2.0,
			},
			Expected: map[string]float64{
				"no_paren":      14.0, // 2 + (3 * 4) = 2 + 12 = 14
				"with_paren":    20.0, // (2 + 3) * 4 = 5 * 4 = 20
				"complex_paren": 15.0, // ((2 + 3) * 4) - (10 / 2) = 20 - 5 = 15
			},
			Tolerance: 0.0001,
		},
		{
			Name: "Trigonometric Functions Validation",
			Expressions: []string{
				"sin_zero = sin(0)",
				"cos_zero = cos(0)",
				"sin_pi_half = sin(1.5707963)",                        // π/2
				"cos_pi = cos(3.1415926)",                             // π
				"pythagorean = (sin(a) * sin(a)) + (cos(a) * cos(a))", // Should be 1
			},
			Inputs: map[string]interface{}{
				"a": 0.5, // Any angle for Pythagorean identity
			},
			Expected: map[string]float64{
				"sin_zero":    0.0,
				"cos_zero":    1.0,
				"sin_pi_half": 1.0,
				"cos_pi":      -1.0,
				"pythagorean": 1.0, // sin²(x) + cos²(x) = 1
			},
			Tolerance: 0.001, // Slightly higher tolerance for trig functions
		},
		{
			Name: "Logarithmic and Exponential Functions Validation",
			Expressions: []string{
				"log_e = log(2.71828)",           // ln(e) ≈ 1
				"exp_zero = exp(0)",              // e⁰ = 1
				"log_exp_identity = log(exp(a))", // Should equal a
				"exp_log_identity = exp(log(b))", // Should equal b
			},
			Inputs: map[string]interface{}{
				"a": 2.5,
				"b": 7.3,
			},
			Expected: map[string]float64{
				"log_e":            1.0,
				"exp_zero":         1.0,
				"log_exp_identity": 2.5, // log(exp(2.5)) = 2.5
				"exp_log_identity": 7.3, // exp(log(7.3)) = 7.3
			},
			Tolerance: 0.001,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			start := time.Now()

			m := NewMinia(tc.Expressions)
			ret := sample.NewMutableFeatures()

			for key, value := range tc.Inputs {
				ret.SetValue(key, value)
			}

			fs := m.Eval(ret)
			evalTime := time.Since(start)

			t.Logf("\n=== %s ===", tc.Name)
			t.Logf("Execution time: %v", evalTime)

			// Validate results
			for varName, expectedVal := range tc.Expected {
				if actualVal, err := fs.Get(varName).GetFloat32(); err == nil {
					t.Logf("%s: %.6f (expected: %.6f)", varName, actualVal, expectedVal)
					if abs(float64(actualVal)-expectedVal) > tc.Tolerance {
						t.Errorf("%s: expected %.6f, got %.6f (tolerance: %.6f)",
							varName, expectedVal, actualVal, tc.Tolerance)
					}
				} else {
					if actualVal, err := fs.Get(varName).GetInt64(); err == nil {
						t.Logf("%s: %d (expected: %.6f)", varName, actualVal, expectedVal)
						if abs(float64(actualVal)-expectedVal) > tc.Tolerance {
							t.Errorf("%s: expected %.6f, got %d (tolerance: %.6f)",
								varName, expectedVal, actualVal, tc.Tolerance)
						}
					} else {
						t.Errorf("Failed to get value for variable %s: %v", varName, err)
					}
				}
			}
		})
	}
}

/**
 * @brief Test expression parsing edge cases
 * @param t Testing context
 * @details Tests parser robustness with various edge cases and error conditions
 */
func TestMiniaParsingEdgeCases(t *testing.T) {
	edgeCases := []struct {
		name        string
		expressions []string
		shouldPanic bool
		description string
	}{
		{
			name:        "Empty Expression List",
			expressions: []string{},
			shouldPanic: true,
			description: "Should panic with empty expression list",
		},
		{
			name:        "Single Variable Assignment",
			expressions: []string{"x = 42"},
			shouldPanic: false,
			description: "Should handle simple variable assignment",
		},
		{
			name:        "Deeply Nested Parentheses",
			expressions: []string{"result = ((((a + b) * c) - d) / e)"},
			shouldPanic: false,
			description: "Should handle deeply nested parentheses",
		},
		{
			name:        "Multiple Dependencies",
			expressions: []string{"a = 1", "b = a + 2", "c = a + b", "result = (a + b) + c"},
			shouldPanic: false,
			description: "Should handle multiple variable dependencies",
		},
		{
			name:        "Complex Function Nesting",
			expressions: []string{"result = sqrt(abs(sin(cos(tan(3.14159)))))"},
			shouldPanic: false,
			description: "Should handle complex function nesting",
		},
	}

	for _, tc := range edgeCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tc.shouldPanic {
						t.Errorf("Unexpected panic: %v", r)
					} else {
						t.Logf("Expected panic occurred: %v", r)
					}
				} else if tc.shouldPanic {
					t.Errorf("Expected panic did not occur")
				}
			}()

			m := NewMinia(tc.expressions)
			if !tc.shouldPanic {
				ret := sample.NewMutableFeatures()
				ret.SetValue("a", 1.0)
				ret.SetValue("b", 2.0)
				ret.SetValue("c", 3.0)
				ret.SetValue("d", 4.0)
				ret.SetValue("e", 5.0)

				result := m.Eval(ret)
				t.Logf("Successfully evaluated: %v", result)
			}
		})
	}
}

// ========== Helper Functions ==========

/**
 * @brief Calculate absolute value of a float64
 * @param x Input value
 * @return Absolute value of x
 */
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

/**
 * @brief Generate random test data for benchmarking
 * @param size Number of data points to generate
 * @return Map of variable names to random values
 */
func generateRandomTestData(size int) map[string]float64 {
	data := make(map[string]float64, size)
	for i := range size {
		varName := fmt.Sprintf("var_%d", i)
		data[varName] = rand.Float64() * 100.0 // Random value 0-100
	}
	return data
}

/**
 * @brief Create sample features from test data
 * @param data Map of variable names to values
 * @return Sample features object
 */
func createSampleFeatures(data map[string]float64) sample.Features {
	ret := sample.NewMutableFeatures()
	for key, value := range data {
		ret.SetValue(key, value)
	}
	return ret
}

/**
 * @brief Benchmark memory allocation efficiency
 * @param b Benchmark testing context
 * @details Specifically tests allocation patterns and GC pressure
 */
func BenchmarkMiniaMemoryEfficiency(b *testing.B) {
	expressions := []string{
		"result1 = (a + b) * (c - d)",
		"result2 = sqrt(a * a + b * b)",
		"result3 = log(abs(a) + 1) + exp(b / 10)",
		"final = (result1 + result2) / (result3 + 0.001)",
	}

	m := NewMinia(expressions)
	testData := map[string]float64{
		"a": 2.0,
		"b": 3.0,
		"c": 4.0,
		"d": 10.0,
		"e": 2.0,
	}

	features := createSampleFeatures(testData)

	b.ReportAllocs() // Enable allocation reporting
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Eval(features)
	}
}
