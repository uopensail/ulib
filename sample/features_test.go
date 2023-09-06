package sample

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/bytedance/sonic"
)

func TestFeatures(t *testing.T) {
	data := `{
		"A":{"type":0, "value":1}, 
		"B":{"type":1, "value":1.5},
		"C":{"type":2, "value":"hello world"},
		"D":{"type":3, "value":[5, 5, 6]},
		"E":{"type":4, "value":[3.4, 5.7]},
		"F":{"type":5, "value":[${data}]}
	}`

	strs := make([]string, 0, 100)
	for i := 0; i < 100; i++ {
		strs = append(strs, fmt.Sprintf("\"%d\"", i))
	}
	s := strings.Join(strs, ",")
	data = strings.ReplaceAll(data, "${data}", s)

	arena := NewArena()
	immutableFeas := NewImmutableFeatures(arena)
	mutableFeas := NewMutableFeatures()
	err := sonic.Unmarshal([]byte(data), immutableFeas)
	if err != nil {
		panic(err)
	}
	err = sonic.Unmarshal([]byte(data), mutableFeas)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("ImmutableFeatures: %v\n", fea)

	//fmt.Printf("MutableFeatures: %v\n", mfea)

	fmt.Println("Get Feature from immutableFeas")
	v1, _ := immutableFeas.GetInt64("A")
	fmt.Printf("A: %v\n", v1)

	v2, _ := immutableFeas.GetFloat32("B")
	fmt.Printf("B: %v\n", v2)

	v3, _ := immutableFeas.GetString("C")
	fmt.Printf("C: %v\n", v3)

	v4, _ := immutableFeas.GetInt64s("D")
	fmt.Printf("D: %v\n", v4)

	v5, _ := immutableFeas.GetFloat32s("E")
	fmt.Printf("E: %v\n", v5)

	v6, _ := immutableFeas.GetStrings("F")
	fmt.Printf("F: %v\n", v6)

	fmt.Println("Get Feature from mutableFeas")
	v1, _ = mutableFeas.GetInt64("A")
	fmt.Printf("A: %v\n", v1)

	v2, _ = mutableFeas.GetFloat32("B")
	fmt.Printf("B: %v\n", v2)

	v3, _ = mutableFeas.GetString("C")
	fmt.Printf("C: %v\n", v3)

	v4, _ = mutableFeas.GetInt64s("D")
	fmt.Printf("D: %v\n", v4)

	v5, _ = mutableFeas.GetFloat32s("E")
	fmt.Printf("E: %v\n", v5)

	v6, _ = mutableFeas.GetStrings("F")
	fmt.Printf("F: %v\n", v6)

	fmt.Println("immutableFeas -> mutableFeas")

	mutableFeas = immutableFeas.Mutable()

	fmt.Println("Get Feature from mutableFeas")
	v1, _ = mutableFeas.GetInt64("A")
	fmt.Printf("A: %v\n", v1)

	v2, _ = mutableFeas.GetFloat32("B")
	fmt.Printf("B: %v\n", v2)

	v3, _ = mutableFeas.GetString("C")
	fmt.Printf("C: %v\n", v3)

	v4, _ = mutableFeas.GetInt64s("D")
	fmt.Printf("D: %v\n", v4)

	v5, _ = mutableFeas.GetFloat32s("E")
	fmt.Printf("E: %v\n", v5)

	v6, _ = mutableFeas.GetStrings("F")
	fmt.Printf("F: %v\n", v6)

	fmt.Println("immutableFeas Marshal")
	msg, _ := sonic.Marshal(immutableFeas)
	fmt.Println(string(msg))

	fmt.Println("mutableFeas Marshal")
	msg, _ = sonic.Marshal(mutableFeas)
	fmt.Println(string(msg))

	count := 1000000
	start := time.Now()
	for i := 0; i < count; i++ {
		mutableFeas.GetStrings("E")
	}
	elapsed := time.Since(start)
	fmt.Printf("mutableFeas time cost: %s\n", elapsed)

	start = time.Now()
	for i := 0; i < count; i++ {
		immutableFeas.GetStrings("E")
	}
	elapsed = time.Since(start)
	fmt.Printf("immutableFeas time cost: %s\n", elapsed)
}

func TestImmutableFeas(t *testing.T) {
	data := `{
		"A":{"type":0, "value":1}, 
		"B":{"type":1, "value":1.5},
		"C":{"type":2, "value":"hello world"},
		"D":{"type":3, "value":[5, 5, 6]},
		"E":{"type":4, "value":[3.4, 5.7]},
		"F":{"type":5, "value":[${data}]}
	}`
	strs := make([]string, 0, 100)
	for i := 0; i < 10000; i++ {
		strs = append(strs, fmt.Sprintf("\"%d\"", i))
	}
	s := strings.Join(strs, ",")
	data = strings.ReplaceAll(data, "${data}", s)

	arena := NewArena()
	immutableFeas := NewImmutableFeatures(arena)
	runtime.GC()

	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)
	fmt.Printf("At start: Mallocs=%d Frees=%d HeapObjects=%d\n",
		m1.Mallocs, m1.Frees, m1.HeapObjects)
	sonic.Unmarshal([]byte(data), immutableFeas)
	runtime.KeepAlive(immutableFeas)
	runtime.GC()

	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)
	fmt.Printf("Allocated %d objects\n", int(m2.HeapObjects)-int(m1.HeapObjects))

	fmt.Printf("At end: Mallocs=%d Frees=%d HeapObjects=%d\n",
		m2.Mallocs, m2.Frees, m2.HeapObjects)

}

func TestMutableFeas(t *testing.T) {
	data := `{
		"A":{"type":0, "value":1}, 
		"B":{"type":1, "value":1.5},
		"C":{"type":2, "value":"hello world"},
		"D":{"type":3, "value":[5, 5, 6]},
		"E":{"type":4, "value":[3.4, 5.7]},
		"F":{"type":5, "value":[${data}]}
	}`
	strs := make([]string, 0, 100)
	for i := 0; i < 10000; i++ {
		strs = append(strs, fmt.Sprintf("\"%d\"", i))
	}
	s := strings.Join(strs, ",")
	data = strings.ReplaceAll(data, "${data}", s)
	mutableFeas := NewMutableFeatures()
	runtime.GC()

	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)
	fmt.Printf("At start: Mallocs=%d Frees=%d HeapObjects=%d\n",
		m1.Mallocs, m1.Frees, m1.HeapObjects)

	sonic.Unmarshal([]byte(data), mutableFeas)
	runtime.KeepAlive(mutableFeas)
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	runtime.GC()

	fmt.Printf("Allocated %d objects\n", int(m2.HeapObjects)-int(m1.HeapObjects))

	fmt.Printf("At end: Mallocs=%d Frees=%d HeapObjects=%d\n",
		m2.Mallocs, m2.Frees, m2.HeapObjects)
}
