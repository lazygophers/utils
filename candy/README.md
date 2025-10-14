# candy - Type Conversion Utilities

The `candy` package provides comprehensive type conversion utilities for Go. It offers safe, flexible conversions between different data types with extensive support for numeric, string, and boolean conversions.

## Features

- **Type Conversion**: Convert between different primitive types safely
- **Collection Operations**: Work with slices and arrays using functional programming patterns
- **Boolean Conversion**: Flexible string-to-boolean conversion with multiple formats
- **Numeric Conversion**: Safe numeric conversions with overflow protection
- **String Utilities**: String manipulation and conversion functions
- **Array Operations**: Functions for working with arrays and slices

## Installation

```bash
go get github.com/lazygophers/utils/candy
```

## Usage Examples

### Boolean Conversions

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/candy"
)

func main() {
    // Various boolean conversions
    fmt.Println(candy.ToBool("true"))   // true
    fmt.Println(candy.ToBool("yes"))    // true
    fmt.Println(candy.ToBool("1"))      // true
    fmt.Println(candy.ToBool("on"))     // true
    fmt.Println(candy.ToBool("false"))  // false
    fmt.Println(candy.ToBool("no"))     // false
    fmt.Println(candy.ToBool("0"))      // false
    fmt.Println(candy.ToBool("off"))    // false
    fmt.Println(candy.ToBool(42))       // true (non-zero)
    fmt.Println(candy.ToBool(0))        // false (zero)
    fmt.Println(candy.ToBool("hello"))  // true (non-empty string)
    fmt.Println(candy.ToBool(""))       // false (empty string)
}
```

### Slice Operations

```go
// Check if any element satisfies condition
numbers := []int{1, 2, 3, 4, 5}
hasEven := candy.Any(numbers, func(n int) bool { return n%2 == 0 })
fmt.Println(hasEven) // true

// Check if all elements satisfy condition
allPositive := candy.All(numbers, func(n int) bool { return n > 0 })
fmt.Println(allPositive) // true

// Calculate average
average := candy.Average(numbers)
fmt.Println(average) // 3.0

// Get absolute values
negatives := []int{-1, -2, 3, -4}
absolutes := candy.Abs(negatives)
fmt.Println(absolutes) // [1, 2, 3, 4]
```

### Array Conversions

```go
// Convert slice to string slice
items := []interface{}{"apple", 42, true}
stringSlice := candy.ToStringSlice(items)
fmt.Println(stringSlice) // ["apple", "42", "true"]
```

## API Reference

### Boolean Conversion

- `ToBool(val interface{}) bool` - Convert any value to boolean

Boolean conversion rules:
- **Numbers**: 0 = false, non-zero = true
- **Strings**: "true", "1", "t", "y", "yes", "on" = true; "false", "0", "f", "n", "no", "off" = false
- **Other strings**: non-empty = true, empty = false
- **nil**: false
- **Other types**: false

### Slice Operations

- `Any[T any](ss []T, f func(T) bool) bool` - Check if any element satisfies condition
- `All[T any](ss []T, f func(T) bool) bool` - Check if all elements satisfy condition
- `Average[T numeric](ss []T) float64` - Calculate average of numeric slice
- `Sum[T numeric](ss []T) T` - Calculate sum of numeric slice
- `Min[T comparable](ss []T) T` - Find minimum value
- `Max[T comparable](ss []T) T` - Find maximum value

### Array Utilities

- `Abs[T numeric](ss []T) []T` - Get absolute values of all elements
- `ToStringSlice(val interface{}) []string` - Convert any slice to string slice
- `ToArrayString(val interface{}) []string` - Alias for ToStringSlice (deprecated)
- `Bottom[T any](ss []T, n int) []T` - Get bottom N elements
- `Top[T any](ss []T, n int) []T` - Get top N elements
- `Unique[T comparable](ss []T) []T` - Remove duplicate elements
- `Reverse[T any](ss []T) []T` - Reverse slice order

### Type Conversions

- `ToString(val interface{}) string` - Convert any value to string
- `ToInt(val interface{}) (int, error)` - Convert to int with error handling
- `ToInt64(val interface{}) (int64, error)` - Convert to int64 with error handling
- `ToFloat64(val interface{}) (float64, error)` - Convert to float64 with error handling
- `ToBytes(val interface{}) []byte` - Convert to byte slice

### Math Operations

- `Abs[T numeric](val T) T` - Get absolute value
- `Max[T comparable](a, b T) T` - Get maximum of two values
- `Min[T comparable](a, b T) T` - Get minimum of two values
- `Clamp[T comparable](val, min, max T) T` - Clamp value between min and max

## Type Constraints

The package uses Go generics with these type constraints:

```go
type numeric interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
    ~float32 | ~float64
}

type comparable interface {
    comparable
}
```

## Performance

The package is optimized for performance with:
- Zero-allocation operations where possible
- Type-specific implementations to avoid reflection
- Efficient slice operations with pre-allocated capacity
- Generic implementations for type safety without runtime overhead

## Error Handling

Most conversion functions use the panic-free pattern, returning sensible defaults:
- Boolean conversions default to `false` for unrecognized types
- Numeric conversions return zero values for invalid inputs
- String conversions always succeed with string representation

For cases where you need error information, use the error-returning variants:
- `ToIntE()`, `ToInt64E()`, `ToFloat64E()` etc.

## Examples

### Data Processing Pipeline

```go
// Process a slice of mixed data
data := []interface{}{1, "2", 3.5, "4", true}

// Convert all to numbers, filter positives, get average
numbers := candy.Map(data, func(v interface{}) float64 {
    if f, err := candy.ToFloat64(v); err == nil {
        return f
    }
    return 0
})

positives := candy.Filter(numbers, func(n float64) bool {
    return n > 0
})

average := candy.Average(positives)
fmt.Printf("Average of positive numbers: %.2f\n", average)
```

### Configuration Parsing

```go
// Parse configuration values
config := map[string]interface{}{
    "debug":     "true",
    "port":      "8080",
    "timeout":   "30s",
    "enabled":   1,
}

debug := candy.ToBool(config["debug"])       // true
port := candy.ToString(config["port"])       // "8080"
enabled := candy.ToBool(config["enabled"])   // true
```

## Related Packages

- `anyx` - Generic type utilities for interface{} handling
- `stringx` - Advanced string manipulation utilities
- `validator` - Struct validation utilities