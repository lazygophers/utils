# anyx - Generic Type Utilities

The `anyx` package provides utilities for working with `interface{}` (any) types and generic map operations in Go. It offers type-safe operations for extracting keys and values from maps and converting data structures.

## Features

- **Value Type Detection**: Detect the type category of any value (number, string, bool, unknown)
- **Map Key Extraction**: Extract keys from maps with type-specific functions
- **Map Value Extraction**: Extract values from maps with type-specific functions
- **Type Conversion**: Convert between different data types safely
- **Slice to Map**: Convert slices to maps for fast lookups
- **KeyBy Operations**: Create indexed maps from struct slices

## Installation

```bash
go get github.com/lazygophers/utils/anyx
```

## Usage Examples

### Value Type Detection

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/anyx"
)

func main() {
    fmt.Println(anyx.CheckValueType(42))       // ValueNumber
    fmt.Println(anyx.CheckValueType("hello"))  // ValueString
    fmt.Println(anyx.CheckValueType(true))     // ValueBool
}
```

### Map Key Extraction

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
keys := anyx.MapKeysString(m)
fmt.Println(keys) // [a b c]

numMap := map[int]string{1: "one", 2: "two"}
numKeys := anyx.MapKeysInt(numMap)
fmt.Println(numKeys) // [1 2]
```

### Map Value Extraction

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
values := anyx.MapValues(m)
fmt.Println(values) // [1 2 3]

// For interface{} maps
anyMap := map[string]interface{}{"x": 1, "y": "hello"}
anyValues := anyx.MapValuesAny(anyMap)
fmt.Println(anyValues) // [1 hello]
```

### KeyBy Operations

```go
type User struct {
    ID   uint64
    Name string
}

users := []*User{
    {ID: 1, Name: "Alice"},
    {ID: 2, Name: "Bob"},
}

// Create map indexed by ID
userMap := anyx.KeyByUint64(users, "ID")
fmt.Println(userMap[1].Name) // Alice

// Create map indexed by Name
nameMap := anyx.KeyByString(users, "Name")
fmt.Println(nameMap["Bob"].ID) // 2
```

### Slice to Map Conversion

```go
slice := []string{"apple", "banana", "cherry"}
lookup := anyx.Slice2Map(slice)
fmt.Println(lookup["apple"])  // true
fmt.Println(lookup["grape"])  // false
```

## API Reference

### Value Type Functions

- `CheckValueType(val interface{}) ValueType` - Determine the type category of a value

### Map Key Extraction Functions

- `MapKeysString(m interface{}) []string` - Extract string keys
- `MapKeysInt(m interface{}) []int` - Extract int keys
- `MapKeysUint64(m interface{}) []uint64` - Extract uint64 keys
- `MapKeysFloat64(m interface{}) []float64` - Extract float64 keys
- `MapKeysAny(m interface{}) []interface{}` - Extract keys of any type

### Map Value Extraction Functions

- `MapValues[K, V any](m map[K]V) []V` - Extract values (generic)
- `MapValuesAny(m interface{}) []interface{}` - Extract values of any type
- `MapValuesString(m interface{}) []string` - Extract string values
- `MapValuesInt(m interface{}) []int` - Extract int values

### KeyBy Functions

- `KeyBy(list interface{}, fieldName string) interface{}` - Generic KeyBy operation
- `KeyByUint64[M any](list []*M, fieldName string) map[uint64]*M` - KeyBy with uint64 keys
- `KeyByString[M any](list []*M, fieldName string) map[string]*M` - KeyBy with string keys
- `KeyByInt64[M any](list []*M, fieldName string) map[int64]*M` - KeyBy with int64 keys

### Utility Functions

- `MergeMap[K, V any](source, target map[K]V) map[K]V` - Merge two maps
- `Slice2Map[M comparable](list []M) map[M]bool` - Convert slice to lookup map

## Performance

The package is optimized for performance with:
- Pre-allocated slices with known capacity
- Minimal memory allocations
- Type-specific functions to avoid reflection overhead where possible

## Notes

- Functions that work with `interface{}` use reflection and may panic if types don't match expectations
- Generic versions (using type parameters) are preferred for better type safety and performance
- All key extraction functions handle nil maps gracefully
- KeyBy functions expect struct elements and will panic if the field is not found

## Related Packages

- `candy` - Type conversion utilities
- `stringx` - String manipulation utilities
- `randx` - Random value generation