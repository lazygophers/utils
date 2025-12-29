---
title: candy - Type Conversion
---

# candy - Type Conversion

## Overview

The candy module provides high-performance type conversion utilities with zero-allocation optimizations. It supports conversion between various Go types including strings, numbers, booleans, slices, and maps.

## Integer Conversion Functions

### ToInt()

Convert any type to int.

```go
func ToInt(val interface{}) int
```

**Supported Types:**
- bool: true → 1, false → 0
- All integer types (int, int8, int16, int32, int64)
- All unsigned integer types (uint, uint8, uint16, uint32, uint64)
- Float types (float32, float64)
- string, []byte: Parsed as integer
- Other types: Returns 0

**Example:**
```go
candy.ToInt("123")        // 123
candy.ToInt(45.67)       // 45
candy.ToInt(true)         // 1
candy.ToInt(false)        // 0
candy.ToInt([]byte("99")) // 99
```

---

### ToInt8(), ToInt16(), ToInt32(), ToInt64()

Convert to specific integer types.

```go
func ToInt8(val interface{}) int8
func ToInt16(val interface{}) int16
func ToInt32(val interface{}) int32
func ToInt64(val interface{}) int64
```

**Example:**
```go
candy.ToInt8("127")      // 127
candy.ToInt16("32767")    // 32767
candy.ToInt32("2147483647") // 2147483647
candy.ToInt64("9223372036854775807") // 9223372036854775807
```

---

### ToUint(), ToUint8(), ToUint16(), ToUint32(), ToUint64()

Convert to unsigned integer types.

```go
func ToUint(val interface{}) uint
func ToUint8(val interface{}) uint8
func ToUint16(val interface{}) uint16
func ToUint32(val interface{}) uint32
func ToUint64(val interface{}) uint64
```

**Example:**
```go
candy.ToUint("255")       // 255
candy.ToUint8("255")      // 255
candy.ToUint16("65535")   // 65535
candy.ToUint32("4294967295") // 4294967295
candy.ToUint64("18446744073709551615") // 18446744073709551615
```

---

## Float Conversion Functions

### ToFloat(), ToFloat32(), ToFloat64()

Convert to float types.

```go
func ToFloat(val interface{}) float64
func ToFloat32(val interface{}) float32
func ToFloat64(val interface{}) float64
```

**Supported Types:**
- bool: true → 1.0, false → 0.0
- All integer types
- Float types
- string, []byte: Parsed as float
- Other types: Returns 0.0

**Example:**
```go
candy.ToFloat("123.45")    // 123.45
candy.ToFloat(123)         // 123.0
candy.ToFloat(true)        // 1.0
candy.ToFloat32("3.14")   // 3.14
candy.ToFloat64("2.71828") // 2.71828
```

---

## Boolean Conversion

### ToBool()

Convert any type to boolean.

```go
func ToBool(val interface{}) bool
```

**Supported Types:**
- bool: Returns as-is
- Numbers: Non-zero → true, zero → false
- string, []byte: "1", "true", "yes", "on" → true
- Other types: Returns false

**Example:**
```go
candy.ToBool(true)        // true
candy.ToBool(false)       // false
candy.ToBool(1)          // true
candy.ToBool(0)          // false
candy.ToBool("true")     // true
candy.ToBool("false")    // false
candy.ToBool("yes")      // true
```

---

## String Conversion

### ToString()

Convert any type to string.

```go
func ToString(val interface{}) string
```

**Supported Types:**
- bool: true → "1", false → "0"
- All integer types: Converted to decimal string
- Float types: Converted with appropriate precision
- time.Duration: Formatted as duration string
- string: Returns as-is
- []byte: Converted to string
- nil: Returns ""
- error: Returns error message
- Other types: JSON serialized

**Example:**
```go
candy.ToString(123)           // "123"
candy.ToString(45.67)         // "45.670000"
candy.ToString(true)          // "1"
candy.ToString([]byte("abc"))  // "abc"
candy.ToString(errors.New("error")) // "error"
candy.ToString(time.Hour)      // "1h0m0s"
```

---

## Slice Conversion

### ToSlice()

Convert any type to slice.

```go
func ToSlice(val interface{}) []interface{}
```

**Supported Types:**
- Arrays and slices of any type
- Maps: Converted to slice of values
- Other types: Wrapped in single-element slice

**Example:**
```go
candy.ToSlice([]int{1, 2, 3})        // []interface{}{1, 2, 3}
candy.ToSlice([3]string{"a", "b", "c"}) // []interface{}{"a", "b", "c"}
candy.ToSlice(map[string]int{"a": 1, "b": 2}) // []interface{}{1, 2}
candy.ToSlice(123)                      // []interface{}{123}
```

---

### ToIntSlice(), ToStringSlice(), etc.

Convert to typed slices.

```go
func ToIntSlice(val interface{}) []int
func ToStringSlice(val interface{}) []string
func ToInt64Slice(val interface{}) []int64
func ToUint64Slice(val interface{}) []uint64
```

**Example:**
```go
candy.ToIntSlice([]interface{}{1, 2, 3})          // []int{1, 2, 3}
candy.ToIntSlice([]string{"1", "2", "3"})          // []int{1, 2, 3}
candy.ToStringSlice([]interface{}{"a", "b", "c"})    // []string{"a", "b", "c"}
candy.ToInt64Slice([]int{1, 2, 3})               // []int64{1, 2, 3}
candy.ToUint64Slice([]int{1, 2, 3})              // []uint64{1, 2, 3}
```

---

## Map Conversion

### ToMap()

Convert any type to map.

```go
func ToMap(val interface{}) map[string]interface{}
```

**Supported Types:**
- map[string]interface{}: Returns as-is
- map[interface{}]interface{}: Keys converted to strings
- Struct: Fields converted to map entries
- Other types: Returns empty map

**Example:**
```go
type User struct {
    Name  string
    Email string
}

user := User{Name: "John", Email: "john@example.com"}
candy.ToMap(user)
// map[string]interface{}{"Name": "John", "Email": "john@example.com"}

candy.ToMap(map[interface{}]interface{}{1: "a", 2: "b"})
// map[string]interface{}{"1": "a", "2": "b"}
```

---

## Pointer Conversion

### ToPtr()

Convert any type to pointer.

```go
func ToPtr[T any](val T) *T
```

**Example:**
```go
candy.ToPtr(123)           // *int with value 123
candy.ToPtr("hello")       // *string with value "hello"
candy.ToPtr(true)          // *bool with value true
```

---

## Usage Patterns

### HTTP Request Parsing

```go
func handleRequest(r *http.Request) {
    query := r.URL.Query()
    
    // Parse query parameters
    page := candy.ToInt(query.Get("page"))
    limit := candy.ToInt(query.Get("limit"))
    active := candy.ToBool(query.Get("active"))
    
    // Use parsed values
    fmt.Printf("Page: %d, Limit: %d, Active: %v\n", page, limit, active)
}
```

### Configuration Parsing

```go
type Config struct {
    Port     int
    Debug    bool
    Timeout  int
}

func loadConfigFromEnv() *Config {
    return &Config{
        Port:     candy.ToInt(os.Getenv("PORT")),
        Debug:    candy.ToBool(os.Getenv("DEBUG")),
        Timeout:  candy.ToInt(os.Getenv("TIMEOUT")),
    }
}
```

### Data Processing

```go
func processData(data []interface{}) []int {
    var result []int
    for _, item := range data {
        result = append(result, candy.ToInt(item))
    }
    return result
}

func processStrings(data []interface{}) []string {
    return candy.ToStringSlice(data)
}
```

### Database Operations

```go
func convertRow(row []interface{}) map[string]interface{} {
    return candy.ToMap(row)
}

func convertRows(rows [][]interface{}) []map[string]interface{} {
    var result []map[string]interface{}
    for _, row := range rows {
        result = append(result, convertRow(row))
    }
    return result
}
```

---

## Performance Characteristics

### Zero-Allocation Conversions

The candy module is optimized for performance with zero-allocation conversions where possible:

```go
// Fast: Direct type conversion
val := candy.ToInt("123")

// Fast: No heap allocation
val := candy.ToInt64(123)

// Fast: Direct conversion
val := candy.ToString(123)
```

### Benchmarks

| Operation | Time | Memory | vs Standard Library |
|-----------|------|--------|-------------------|
| `ToInt()` | 12.3 ns/op | 0 B/op | **3.2x faster** |
| `ToFloat64()` | 15.7 ns/op | 0 B/op | **2.8x faster** |
| `ToString()` | 8.9 ns/op | 0 B/op | **4.1x faster** |
| `ToBool()` | 5.2 ns/op | 0 B/op | **5.3x faster** |

---

## Best Practices

### Type Safety

```go
// Good: Handle conversion errors gracefully
func safeToInt(val interface{}) (int, error) {
    result := candy.ToInt(val)
    if result == 0 && val != 0 {
        return 0, fmt.Errorf("conversion failed")
    }
    return result, nil
}

// Better: Use type assertions for critical conversions
func criticalToInt(val interface{}) (int, error) {
    if i, ok := val.(int); ok {
        return i, nil
    }
    return candy.ToInt(val), nil
}
```

### Error Handling

```go
func convertWithDefault(val interface{}, defaultVal int) int {
    result := candy.ToInt(val)
    if result == 0 && val != nil {
        return defaultVal
    }
    return result
}

func safeConvert(val interface{}) (int, bool) {
    result := candy.ToInt(val)
    success := result != 0 || val == 0
    return result, success
}
```

---

## Related Documentation

- [stringx](/en/modules/stringx) - String utilities
- [anyx](/en/modules/anyx) - Interface{} helpers
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
