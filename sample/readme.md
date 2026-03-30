# sample — High-Performance Feature Storage

A Go package providing type-safe, arena-backed storage for heterogeneous data
features. Two complementary types cover the full read/write spectrum:

- **`ImmutableFeatures`** — zero-copy, arena-allocated; safe for concurrent reads.
- **`MutableFeatures`** — heap-allocated map; supports add / update / delete.

Requires **Go 1.26** or later.

## Supported types

| `DataType` constant | Value | Go type |
|---|---|---|
| `Int64Type` | 0 | `int64` |
| `Float32Type` | 1 | `float32` |
| `StringType` | 2 | `string` |
| `Int64sType` | 3 | `[]int64` |
| `Float32sType` | 4 | `[]float32` |
| `StringsType` | 5 | `[]string` |

## Core interfaces

```go
// Feature holds one typed value.
type Feature interface {
    Type() DataType

    Get() any

    GetInt64() (int64, error)
    GetInt64Unsafe() int64

    GetFloat32() (float32, error)
    GetFloat32Unsafe() float32

    GetString() (string, error)
    GetStringUnsafe() string

    GetInt64s() ([]int64, error)
    GetInt64sUnsafe() []int64

    GetFloat32s() ([]float32, error)
    GetFloat32sUnsafe() []float32

    GetStrings() ([]string, error)
    GetStringsUnsafe() []string
}

// Features is the common interface for both collection types.
type Features interface {
    GetType(key string) DataType
    Keys() []string
    Get(key string) Feature
    Len() int
    Has(key string) bool
    ForEach(fn IteratorFunc) error
    All() iter.Seq2[string, Feature]   // Go 1.23+ range-over-function
    MapAny() (map[string]any, error)
    MarshalJSON() ([]byte, error)
    UnmarshalJSON(data []byte) error
}
```

The `Unsafe` getters skip the type check and must only be called after
confirming `feature.Type()`. The safe variants (`GetInt64`, etc.) return
`ErrTypeMismatch` on a mismatch.

## ImmutableFeatures

### Construction

```go
arena := sample.NewArena()

// From a plain map
data := map[string]any{
    "user_id": int64(12345),
    "score":   float32(98.5),
    "name":    "Alice",
    "tags":    []string{"premium", "active"},
}
features, err := sample.NewImmutableFeaturesFromMap(data, arena)
if err != nil {
    log.Fatal(err)
}

// From MutableFeatures
mut := sample.NewMutableFeatures()
mut.SetValue("x", int64(1))
imm, err := mut.Immutable(arena)
```

### Reading

```go
// Single key
if f := features.Get("user_id"); f != nil {
    id, _ := f.GetInt64()
    fmt.Println(id)
}

// Type-switch style (no allocation)
if features.GetType("score") == sample.Float32Type {
    v := features.Get("score").GetFloat32Unsafe()
    fmt.Println(v)
}

// Callback iteration
features.ForEach(func(key string, f sample.Feature) error {
    fmt.Printf("%s: %s\n", key, f.Type())
    return nil
})

// Range-over-function (Go 1.23+)
for key, f := range features.All() {
    fmt.Printf("%s: %v\n", key, f.Get())
}
```

### JSON round-trip

```go
data, _ := features.MarshalJSON()

var restored sample.ImmutableFeatures
_ = restored.UnmarshalJSON(data)
```

### Memory layout

All values are stored in contiguous arena pages with 8-byte alignment.
No heap allocation occurs on read.

```
Int64:    [DataType:8][Value:8]                          = 16 B
Float32:  [DataType:8][Value:4][Pad:4]                  = 16 B
String:   [DataType:8][Len:8][Data:align(len)]           = 16 + align(len) B
Int64s:   [DataType:8][Len:8][Data:len×8]               = 16 + len×8 B
Float32s: [DataType:8][Len:8][Data:align(len×4)]         = 16 + align(len×4) B
Strings:  [DataType:8][Len:8][Headers:align(len×16)][Data:…]
```

## MutableFeatures

```go
f := sample.NewMutableFeatures()

// Set values
f.SetValue("user_id", int64(42))
f.SetValue("score", float32(9.5))
f.Set("tags", &sample.Strings{Value: []string{"a", "b"}})

// Get / check
feat := f.Get("score")           // Feature or nil
ok   := f.Has("missing")         // false
dt   := f.GetType("user_id")     // sample.Int64Type

// Mutate
f.SetValue("score", float32(10.0))
f.Delete("old_key")

// Deep copy
clone := f.Clone()

// Promote to immutable
arena := sample.NewArena()
imm, err := f.Immutable(arena)
```

## Arena

```go
arena := sample.NewArena()

fmt.Printf("pages: %d, bytes: %d\n", arena.PageCount(), arena.Size())
```

- Requests ≤ 4 KiB are served from the current page; a fresh page is appended
  when it fills up.
- Larger requests receive a dedicated page inserted just before the current one.
- All allocations are 8-byte aligned.
- Under Go 1.26's Green Tea GC the Arena remains valuable primarily for
  **zero-copy access**: strings and slices returned by `ImmutableFeature`
  point directly into arena pages rather than owning separate heap blocks.

## JSON format

```json
{
  "user_id": {"type": 0, "value": 12345},
  "score":   {"type": 1, "value": 98.5},
  "name":    {"type": 2, "value": "Alice"},
  "tags":    {"type": 5, "value": ["premium", "active"]}
}
```

Type integers correspond to the `DataType` iota constants (see table above).

## When to use which

| Scenario | Recommendation |
|---|---|
| Read-heavy, large dataset, concurrent access | `ImmutableFeatures` |
| Frequent mutations, small dataset | `MutableFeatures` |
| Build phase → serve phase | Build with `MutableFeatures`, promote with `Immutable()` |

## Dependencies

- [`github.com/bytedance/sonic`](https://github.com/bytedance/sonic) — JSON encoding/decoding
