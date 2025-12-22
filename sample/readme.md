# Sample - High-Performance Feature Storage System

A Go library providing efficient, type-safe storage and access for heterogeneous data features with both mutable and immutable implementations.

## Overview

This library offers two complementary approaches to feature storage:
- **ImmutableFeatures**: Zero-copy, arena-allocated storage for high-performance read operations
- **MutableFeatures**: Traditional map-based storage for flexible data manipulation

## Key Features

- 🚀 **Zero-copy access** for immutable features
- 🔒 **Type safety** with runtime type checking
- 🧠 **Memory efficient** with custom arena allocation
- 🔄 **Bidirectional conversion** between mutable and immutable forms
- 📦 **JSON serialization** support
- ⚡ **High performance** with 8-byte memory alignment
- 🛡️ **Go 1.20+ optimized** using latest unsafe operations

## Supported Data Types

| Type | Description | Storage |
|------|-------------|---------|
| `Int64Type` | 64-bit signed integer | `int64` |
| `Float32Type` | 32-bit floating point | `float32` |
| `StringType` | UTF-8 string | `string` |
| `Int64sType` | Integer array | `[]int64` |
| `Float32sType` | Float array | `[]float32` |
| `StringsType` | String array | `[]string` |

## Architecture

### Feature Interface

All feature types implement a common interface:

```go
type Feature interface {
    Type() DataType                    // Returns the data type of this feature
    Get() any                          // Retrieves value
    GetInt64() (int64, error)          // Retrieves int64 value
    GetFloat32() (float32, error)      // Retrieves float32 value  
    GetString() (string, error)        // Retrieves string value
    GetInt64s() ([]int64, error)       // Retrieves int64 slice
    GetFloat32s() ([]float32, error)   // Retrieves float32 slice
    GetStrings() ([]string, error)     // Retrieves string slice
}
```

### Features Collection Interface

Both mutable and immutable collections provide:

```go
type Features interface {
    Keys() []string                    // Returns all feature keys
    Get(string) Feature                // Retrieves feature by key
    Len() int                         // Returns number of features
    Has(string) bool                  // Checks if key exists
    MarshalJSON() ([]byte, error)     // JSON serialization
    UnmarshalJSON([]byte) error       // JSON deserialization
}
```

## ImmutableFeatures

Optimized for read-heavy workloads with zero-copy access and minimal memory overhead.

### Benefits
- **Zero-copy reads**: Direct pointer access to arena memory
- **Reduced GC pressure**: Arena allocation minimizes garbage collection
- **Memory efficiency**: Compact, aligned memory layout
- **Thread-safe reads**: Immutable data allows concurrent access

### Memory Layout

All data structures use 8-byte alignment for optimal CPU performance:

#### Scalar Types
```
Int64:    [DataType:8] + [Value:8] = 16 bytes
Float32:  [DataType:8] + [Value:4] + [Padding:4] = 16 bytes
String:   [DataType:8] + [Len:8] + [Data:aligned] = 16 + aligned(len)
```

#### Array Types
```
Int64s:   [DataType:8] + [Len:8] + [Data:len*8] = 16 + len*8 bytes
Float32s: [DataType:8] + [Len:8] + [Data:aligned(len*4)] = 16 + aligned(len*4) bytes
Strings:  [DataType:8] + [Len:8] + [StringHeaders:len*16] + [StringData:aligned]
```

### Usage Example

```go
// Create from map
data := map[string]any{
    "user_id": int64(12345),
    "score": float32(98.5),
    "name": "John Doe",
    "tags": []string{"premium", "active"},
}

features, err := NewImmutableFeaturesFromMap(data, nil)
if err != nil {
    log.Fatal(err)
}

// Zero-copy access
if feature := features.Get("user_id"); feature != nil {
    userID, _ := feature.GetInt64()
    fmt.Printf("User ID: %d\n", userID)
}

// Iterate over all features
features.ForEach(func(key string, feature Feature) error {
    fmt.Printf("Key: %s, Type: %s\n", key, feature.Type())
    return nil
})

// JSON serialization
jsonData, _ := features.MarshalJSON()
fmt.Printf("JSON: %s\n", jsonData)
```

## MutableFeatures

Traditional map-based storage for scenarios requiring frequent modifications.

### Benefits
- **Full mutability**: Add, update, and remove features dynamically
- **Flexible API**: Multiple ways to set and access data
- **Easy conversion**: Convert to ImmutableFeatures for performance
- **Standard semantics**: Familiar map-like operations

### Usage Example

```go
// Create and populate
features := NewMutableFeatures()
features.SetValue("user_id", int64(12345))
features.SetValue("score", float32(98.5))

// Or create from map
data := map[string]any{
    "name": "John Doe",
    "tags": []string{"premium", "active"},
}
features, _ := NewMutableFeaturesFromMap(data)

// Modify features
features.SetValue("score", float32(99.0))
features.Delete("old_key")

// Convert to immutable for performance
immutable, _ := features.Immutable(nil)

// Clone for independent copy
clone := features.Clone()
```

## Arena Memory Management

The Arena allocator provides efficient memory management for ImmutableFeatures:

```go
// Create custom arena
arena := NewArena()

// Check memory usage
fmt.Printf("Total memory: %d bytes\n", arena.Size())
fmt.Printf("Page count: %d\n", arena.PageCount())

// Use with features
features, _ := NewImmutableFeaturesFromMap(data, arena)
```

### Arena Benefits
- **Reduced allocations**: Large chunks allocated upfront
- **Memory locality**: Related data stored together
- **Lower GC pressure**: Fewer objects for garbage collector
- **Alignment**: All allocations are 8-byte aligned

## Performance Characteristics

### ImmutableFeatures
- **Read access**: O(1) with zero-copy
- **Memory overhead**: Minimal, only type headers
- **GC impact**: Very low, arena-allocated
- **Thread safety**: Read-safe without locks

### MutableFeatures
- **Read/Write access**: O(1) map operations
- **Memory overhead**: Standard Go map overhead
- **GC impact**: Standard Go object lifecycle
- **Thread safety**: Not thread-safe (external sync required)

## JSON Format

Both feature types use a consistent JSON format:

```json
{
  "user_id": {
    "type": 1,
    "value": 12345
  },
  "score": {
    "type": 2,
    "value": 98.5
  },
  "name": {
    "type": 3,
    "value": "John Doe"
  },
  "tags": {
    "type": 6,
    "value": ["premium", "active"]
  }
}
```

## Best Practices

### When to Use ImmutableFeatures
- Read-heavy workloads
- Large datasets with memory constraints
- High-performance requirements
- Concurrent read access needed

### When to Use MutableFeatures
- Frequent modifications required
- Small to medium datasets
- Prototyping and development
- Dynamic feature sets

### Memory Optimization
- Reuse Arena instances when possible
- Convert MutableFeatures to ImmutableFeatures for read-heavy phases
- Monitor memory usage with `Arena.Size()`

## Requirements

- Go 1.20 or later (for unsafe.Slice and unsafe.String)
- github.com/bytedance/sonic (for JSON operations)
