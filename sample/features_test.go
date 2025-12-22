package sample

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/bytedance/sonic"
)

/**
 * @brief TestFeatures tests the basic functionality of both ImmutableFeatures and MutableFeatures
 *
 * This comprehensive test covers:
 * - JSON unmarshaling for both feature types
 * - Feature retrieval and type-specific value access
 * - Conversion between mutable and immutable features
 * - JSON marshaling
 * - Performance comparison between mutable and immutable access
 *
 * @param t Testing framework instance
 */
func TestFeatures(t *testing.T) {
	// Create test data with all supported types
	data := `{
		"A":{"type":0, "value":1}, 
		"B":{"type":1, "value":1.5},
		"C":{"type":2, "value":"hello world"},
		"D":{"type":3, "value":[5, 5, 6]},
		"E":{"type":4, "value":[3.4, 5.7]},
		"F":{"type":5, "value":[${data}]}
	}`

	// Generate large string array for testing
	strs := make([]string, 0, 100)
	for i := range 100 {
		strs = append(strs, fmt.Sprintf("\"%d\"", i))
	}
	s := strings.Join(strs, ",")
	data = strings.ReplaceAll(data, "${data}", s)

	// Initialize both feature types
	arena := NewArena()
	immutableFeas := NewImmutableFeatures(arena)
	mutableFeas := NewMutableFeatures()

	// Test JSON unmarshaling
	err := sonic.Unmarshal([]byte(data), immutableFeas)
	if err != nil {
		t.Fatalf("Failed to unmarshal immutable features: %v", err)
	}
	err = sonic.Unmarshal([]byte(data), mutableFeas)
	if err != nil {
		t.Fatalf("Failed to unmarshal mutable features: %v", err)
	}

	// Test immutable feature access
	t.Log("Testing ImmutableFeatures access")
	testFeatureAccess(t, immutableFeas, "ImmutableFeatures")

	// Test mutable feature access
	t.Log("Testing MutableFeatures access")
	testFeatureAccess(t, mutableFeas, "MutableFeatures")

	// Test conversion from immutable to mutable
	t.Log("Testing immutable to mutable conversion")
	convertedMutable := immutableFeas.Mutable()
	testFeatureAccess(t, convertedMutable, "Converted MutableFeatures")

	// Test JSON marshaling
	t.Log("Testing JSON marshaling")
	testJSONMarshaling(t, immutableFeas, "ImmutableFeatures")
	testJSONMarshaling(t, mutableFeas, "MutableFeatures")

	// Performance comparison
	t.Log("Running performance comparison")
	performanceComparison(t, mutableFeas, immutableFeas)
}

/**
 * @brief testFeatureAccess tests feature access for all supported data types
 *
 * @param t Testing framework instance
 * @param features Features collection to test
 * @param name Name of the feature type for logging
 */
func testFeatureAccess(t *testing.T, features interface{}, name string) {
	// Type assertion to get the common interface
	var getter interface {
		Get(string) Feature
	}

	switch f := features.(type) {
	case *ImmutableFeatures:
		getter = f
	case *MutableFeatures:
		getter = f
	default:
		t.Fatalf("Unsupported features type: %T", features)
	}

	t.Logf("Testing %s feature access", name)

	// Test Int64 access
	if feature := getter.Get("A"); feature != nil {
		if v, err := feature.GetInt64(); err != nil {
			t.Errorf("Failed to get int64 value: %v", err)
		} else {
			t.Logf("A (int64): %v", v)
			if v != 1 {
				t.Errorf("Expected int64 value 1, got %v", v)
			}
		}
	} else {
		t.Error("Feature A not found")
	}

	// Test Float32 access
	if feature := getter.Get("B"); feature != nil {
		if v, err := feature.GetFloat32(); err != nil {
			t.Errorf("Failed to get float32 value: %v", err)
		} else {
			t.Logf("B (float32): %v", v)
			if v != 1.5 {
				t.Errorf("Expected float32 value 1.5, got %v", v)
			}
		}
	} else {
		t.Error("Feature B not found")
	}

	// Test String access
	if feature := getter.Get("C"); feature != nil {
		if v, err := feature.GetString(); err != nil {
			t.Errorf("Failed to get string value: %v", err)
		} else {
			t.Logf("C (string): %v", v)
			if v != "hello world" {
				t.Errorf("Expected string value 'hello world', got %v", v)
			}
		}
	} else {
		t.Error("Feature C not found")
	}

	// Test Int64s access
	if feature := getter.Get("D"); feature != nil {
		if v, err := feature.GetInt64s(); err != nil {
			t.Errorf("Failed to get int64s value: %v", err)
		} else {
			t.Logf("D ([]int64): %v", v)
			expected := []int64{5, 5, 6}
			if len(v) != len(expected) {
				t.Errorf("Expected int64s length %d, got %d", len(expected), len(v))
			}
		}
	} else {
		t.Error("Feature D not found")
	}

	// Test Float32s access
	if feature := getter.Get("E"); feature != nil {
		if v, err := feature.GetFloat32s(); err != nil {
			t.Errorf("Failed to get float32s value: %v", err)
		} else {
			t.Logf("E ([]float32): %v", v)
			expected := []float32{3.4, 5.7}
			if len(v) != len(expected) {
				t.Errorf("Expected float32s length %d, got %d", len(expected), len(v))
			}
		}
	} else {
		t.Error("Feature E not found")
	}

	// Test Strings access
	if feature := getter.Get("F"); feature != nil {
		if v, err := feature.GetStrings(); err != nil {
			t.Errorf("Failed to get strings value: %v", err)
		} else {
			t.Logf("F ([]string): length=%d, first_few=%v", len(v), v[:min(5, len(v))])
			if len(v) != 100 {
				t.Errorf("Expected strings length 100, got %d", len(v))
			}
		}
	} else {
		t.Error("Feature F not found")
	}
}

/**
 * @brief testJSONMarshaling tests JSON marshaling functionality
 *
 * @param t Testing framework instance
 * @param features Features collection to test
 * @param name Name of the feature type for logging
 */
func testJSONMarshaling(t *testing.T, features interface{}, name string) {
	var marshaler interface {
		MarshalJSON() ([]byte, error)
	}

	switch f := features.(type) {
	case *ImmutableFeatures:
		marshaler = f
	case *MutableFeatures:
		marshaler = f
	default:
		t.Fatalf("Unsupported features type for marshaling: %T", features)
	}

	msg, err := marshaler.MarshalJSON()
	if err != nil {
		t.Errorf("Failed to marshal %s: %v", name, err)
		return
	}

	t.Logf("%s JSON length: %d bytes", name, len(msg))

	// Verify JSON is valid by attempting to unmarshal
	var testData map[string]interface{}
	if err := sonic.Unmarshal(msg, &testData); err != nil {
		t.Errorf("Generated JSON is invalid for %s: %v", name, err)
	}
}

/**
 * @brief performanceComparison compares access performance between mutable and immutable features
 *
 * @param t Testing framework instance
 * @param mutableFeas Mutable features instance
 * @param immutableFeas Immutable features instance
 */
func performanceComparison(t *testing.T, mutableFeas *MutableFeatures, immutableFeas *ImmutableFeatures) {
	const iterations = 1000000

	// Benchmark mutable features access
	start := time.Now()
	for range iterations {
		if feature := mutableFeas.Get("F"); feature != nil {
			feature.GetStrings()
		}
	}
	mutableDuration := time.Since(start)

	// Benchmark immutable features access
	start = time.Now()
	for range iterations {
		if feature := immutableFeas.Get("F"); feature != nil {
			feature.GetStrings()
		}
	}
	immutableDuration := time.Since(start)

	t.Logf("Performance comparison (%d iterations):", iterations)
	t.Logf("  MutableFeatures:   %v (%v per op)", mutableDuration, mutableDuration/iterations)
	t.Logf("  ImmutableFeatures: %v (%v per op)", immutableDuration, immutableDuration/iterations)

	if immutableDuration < mutableDuration {
		speedup := float64(mutableDuration) / float64(immutableDuration)
		t.Logf("  ImmutableFeatures is %.2fx faster", speedup)
	}
}

/**
 * @brief TestImmutableFeaturesMemoryUsage tests memory allocation patterns for ImmutableFeatures
 *
 * This test measures the number of heap objects allocated when creating ImmutableFeatures
 * from JSON data. It demonstrates the memory efficiency of arena allocation.
 *
 * @param t Testing framework instance
 */
func TestImmutableFeaturesMemoryUsage(t *testing.T) {
	// Create test data with large string array
	data := createTestData(10000)

	arena := NewArena()
	immutableFeas := NewImmutableFeatures(arena)

	// Force garbage collection and measure initial state
	runtime.GC()
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)
	t.Logf("Initial state: Mallocs=%d Frees=%d HeapObjects=%d",
		m1.Mallocs, m1.Frees, m1.HeapObjects)

	// Unmarshal data and measure memory usage
	err := sonic.Unmarshal([]byte(data), immutableFeas)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	runtime.KeepAlive(immutableFeas)
	runtime.GC()

	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	objectsAllocated := int(m2.HeapObjects) - int(m1.HeapObjects)
	t.Logf("ImmutableFeatures allocated %d heap objects", objectsAllocated)
	t.Logf("Final state: Mallocs=%d Frees=%d HeapObjects=%d",
		m2.Mallocs, m2.Frees, m2.HeapObjects)

	// Verify arena usage
	t.Logf("Arena memory usage: %d bytes across %d pages",
		arena.Size(), arena.PageCount())
}

/**
 * @brief TestMutableFeaturesMemoryUsage tests memory allocation patterns for MutableFeatures
 *
 * This test measures the number of heap objects allocated when creating MutableFeatures
 * from JSON data. It provides a comparison baseline against ImmutableFeatures.
 *
 * @param t Testing framework instance
 */
func TestMutableFeaturesMemoryUsage(t *testing.T) {
	// Create test data with large string array
	data := createTestData(10000)

	mutableFeas := NewMutableFeatures()

	// Force garbage collection and measure initial state
	runtime.GC()
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)
	t.Logf("Initial state: Mallocs=%d Frees=%d HeapObjects=%d",
		m1.Mallocs, m1.Frees, m1.HeapObjects)

	// Unmarshal data and measure memory usage
	err := sonic.Unmarshal([]byte(data), mutableFeas)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	runtime.KeepAlive(mutableFeas)
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	runtime.GC()

	objectsAllocated := int(m2.HeapObjects) - int(m1.HeapObjects)
	t.Logf("MutableFeatures allocated %d heap objects", objectsAllocated)
	t.Logf("Final state: Mallocs=%d Frees=%d HeapObjects=%d",
		m2.Mallocs, m2.Frees, m2.HeapObjects)
}

/**
 * @brief TestFeatureConversion tests bidirectional conversion between mutable and immutable features
 *
 * @param t Testing framework instance
 */
func TestFeatureConversion(t *testing.T) {
	// Create initial mutable features
	mutable := NewMutableFeatures()
	mutable.SetValue("int_val", int64(42))
	mutable.SetValue("float_val", float32(3.14))
	mutable.SetValue("string_val", "test string")
	mutable.SetValue("int_array", []int64{1, 2, 3})
	mutable.SetValue("float_array", []float32{1.1, 2.2})
	mutable.SetValue("string_array", []string{"a", "b", "c"})

	t.Logf("Original mutable features has %d items", mutable.Len())

	// Convert to immutable
	immutable, err := mutable.Immutable(NewArena())
	if err != nil {
		t.Fatalf("Failed to convert to immutable: %v", err)
	}

	t.Logf("Converted immutable features has %d items", immutable.Len())

	// Verify all values are preserved
	verifyFeatureValues(t, immutable, "converted immutable")

	// Convert back to mutable
	mutable2 := immutable.Mutable()
	t.Logf("Converted back mutable features has %d items", mutable2.Len())

	// Verify all values are still preserved
	verifyFeatureValues(t, mutable2, "converted back mutable")

	// Test modification of converted mutable
	mutable2.SetValue("new_val", "added after conversion")
	if !mutable2.Has("new_val") {
		t.Error("Failed to add new value to converted mutable features")
	}
}

/**
 * @brief createTestData creates JSON test data with specified number of strings
 *
 * @param stringCount Number of strings to include in the test data
 * @return JSON string containing test data
 */
func createTestData(stringCount int) string {
	data := `{
		"A":{"type":0, "value":1}, 
		"B":{"type":1, "value":1.5},
		"C":{"type":2, "value":"hello world"},
		"D":{"type":3, "value":[5, 5, 6]},
		"E":{"type":4, "value":[3.4, 5.7]},
		"F":{"type":5, "value":[${data}]}
	}`

	strs := make([]string, 0, stringCount)
	for i := range stringCount {
		strs = append(strs, fmt.Sprintf("\"%d\"", i))
	}
	s := strings.Join(strs, ",")
	return strings.ReplaceAll(data, "${data}", s)
}

/**
 * @brief verifyFeatureValues verifies that all expected feature values are present and correct
 *
 * @param t Testing framework instance
 * @param features Features collection to verify
 * @param name Name for logging purposes
 */
func verifyFeatureValues(t *testing.T, features any, name string) {
	var getter interface {
		Get(string) Feature
		Has(string) bool
	}

	switch f := features.(type) {
	case *ImmutableFeatures:
		getter = f
	case *MutableFeatures:
		getter = f
	default:
		t.Fatalf("Unsupported features type: %T", features)
	}

	// Verify int64 value
	if !getter.Has("int_val") {
		t.Errorf("%s missing int_val", name)
	} else if feature := getter.Get("int_val"); feature != nil {
		if v, err := feature.GetInt64(); err != nil || v != 42 {
			t.Errorf("%s int_val incorrect: expected 42, got %v (err: %v)", name, v, err)
		}
	}

	// Verify float32 value
	if !getter.Has("float_val") {
		t.Errorf("%s missing float_val", name)
	} else if feature := getter.Get("float_val"); feature != nil {
		if v, err := feature.GetFloat32(); err != nil || v != 3.14 {
			t.Errorf("%s float_val incorrect: expected 3.14, got %v (err: %v)", name, v, err)
		}
	}

	// Verify string value
	if !getter.Has("string_val") {
		t.Errorf("%s missing string_val", name)
	} else if feature := getter.Get("string_val"); feature != nil {
		if v, err := feature.GetString(); err != nil || v != "test string" {
			t.Errorf("%s string_val incorrect: expected 'test string', got %v (err: %v)", name, v, err)
		}
	}

	// Verify array values
	expectedIntArray := []int64{1, 2, 3}
	if feature := getter.Get("int_array"); feature != nil {
		if v, err := feature.GetInt64s(); err != nil || len(v) != len(expectedIntArray) {
			t.Errorf("%s int_array incorrect length: expected %d, got %d (err: %v)",
				name, len(expectedIntArray), len(v), err)
		}
	}
}

/**
 * @brief min returns the minimum of two integers
 *
 * @param a First integer
 * @param b Second integer
 * @return Minimum value
 */
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * @brief BenchmarkImmutableFeatureAccess benchmarks immutable feature access performance
 *
 * @param b Benchmark framework instance
 */
func BenchmarkImmutableFeatureAccess(b *testing.B) {
	data := createTestData(1000)
	arena := NewArena()
	features := NewImmutableFeatures(arena)
	sonic.Unmarshal([]byte(data), features)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if feature := features.Get("F"); feature != nil {
				feature.GetStrings()
			}
		}
	})
}

/**
 * @brief BenchmarkMutableFeatureAccess benchmarks mutable feature access performance
 *
 * @param b Benchmark framework instance
 */
func BenchmarkMutableFeatureAccess(b *testing.B) {
	data := createTestData(1000)
	features := NewMutableFeatures()
	sonic.Unmarshal([]byte(data), features)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if feature := features.Get("F"); feature != nil {
				feature.GetStrings()
			}
		}
	})
}

/**
 * @brief BenchmarkJSONUnmarshalImmutable benchmarks JSON unmarshaling for immutable features
 *
 * @param b Benchmark framework instance
 */
func BenchmarkJSONUnmarshalImmutable(b *testing.B) {
	data := []byte(createTestData(100))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arena := NewArena()
		features := NewImmutableFeatures(arena)
		sonic.Unmarshal(data, features)
	}
}

/**
 * @brief BenchmarkJSONUnmarshalMutable benchmarks JSON unmarshaling for mutable features
 *
 * @param b Benchmark framework instance
 */
func BenchmarkJSONUnmarshalMutable(b *testing.B) {
	data := []byte(createTestData(100))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		features := NewMutableFeatures()
		sonic.Unmarshal(data, features)
	}
}
