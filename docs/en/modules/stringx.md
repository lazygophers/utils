---
title: stringx - String Utilities
---

# stringx - String Utilities

## Overview

The stringx module provides high-performance string utilities with Unicode-aware operations and zero-allocation optimizations.

## Core Functions

### ToString()

Convert []byte to string with zero allocation.

```go
func ToString(b []byte) string
```

**Parameters:**
- `b` - Byte slice

**Returns:**
- String representation

**Example:**
```go
data := []byte("hello")
str := stringx.ToString(data)
// str is "hello" with zero allocation
```

**Notes:**
- Zero-allocation conversion using unsafe
- Faster than `string(b)` for large slices
- Returns empty string for nil or empty input

---

### ToBytes()

Convert string to []byte with zero allocation.

```go
func ToBytes(s string) []byte
```

**Parameters:**
- `s` - String

**Returns:**
- Byte slice representation

**Example:**
```go
str := "hello"
data := stringx.ToBytes(str)
// data is []byte("hello") with zero allocation
```

**Notes:**
- Zero-allocation conversion using unsafe
- Faster than `[]byte(s)` for large strings
- Returns nil for empty string

---

## Case Conversion

### Camel2Snake()

Convert camelCase to snake_case.

```go
func Camel2Snake(s string) string
```

**Example:**
```go
stringx.Camel2Snake("camelCase")   // "camel_case"
stringx.Camel2Snake("MyVariable")  // "my_variable"
stringx.Camel2Snake("HTTPRequest") // "http_request"
```

---

### Snake2Camel()

Convert snake_case to CamelCase.

```go
func Snake2Camel(s string) string
```

**Example:**
```go
stringx.Snake2Camel("snake_case")   // "SnakeCase"
stringx.Snake2Camel("my_variable")  // "MyVariable"
stringx.Snake2Camel("http_request") // "HttpRequest"
```

---

### Snake2SmallCamel()

Convert snake_case to camelCase (small camel).

```go
func Snake2SmallCamel(s string) string
```

**Example:**
```go
stringx.Snake2SmallCamel("snake_case")   // "snakeCase"
stringx.Snake2SmallCamel("my_variable")  // "myVariable"
stringx.Snake2SmallCamel("http_request") // "httpRequest"
```

---

### ToSnake()

Convert any string to snake_case.

```go
func ToSnake(s string) string
```

**Example:**
```go
stringx.ToSnake("camelCase")      // "camel_case"
stringx.ToSnake("CamelCase")      // "camel_case"
stringx.ToSnake("kebab-case")     // "kebab_case"
stringx.ToSnake("PascalCase")     // "pascal_case"
```

---

### ToKebab()

Convert any string to kebab-case.

```go
func ToKebab(s string) string
```

**Example:**
```go
stringx.ToKebab("camelCase")      // "camel-case"
stringx.ToKebab("snake_case")     // "snake-case"
stringx.ToKebab("PascalCase")     // "pascal-case"
```

---

### ToCamel()

Convert any string to CamelCase.

```go
func ToCamel(s string) string
```

**Example:**
```go
stringx.ToCamel("snake_case")     // "SnakeCase"
stringx.ToCamel("kebab-case")     // "KebabCase"
stringx.ToCamel("pascal_case")    // "PascalCase"
```

---

### ToSmallCamel()

Convert any string to camelCase.

```go
func ToSmallCamel(s string) string
```

**Example:**
```go
stringx.ToSmallCamel("snake_case")     // "snakeCase"
stringx.ToSmallCamel("kebab-case")     // "kebabCase"
stringx.ToSmallCamel("PascalCase")     // "pascalCase"
```

---

### ToSlash()

Convert any string to slash/case.

```go
func ToSlash(s string) string
```

**Example:**
```go
stringx.ToSlash("camelCase")      // "camel/case"
stringx.ToSlash("snake_case")     // "snake/case"
stringx.ToSlash("PascalCase")     // "pascal/case"
```

---

### ToDot()

Convert any string to dot.case.

```go
func ToDot(s string) string
```

**Example:**
```go
stringx.ToDot("camelCase")      // "camel.case"
stringx.ToDot("snake_case")     // "snake.case"
stringx.ToDot("PascalCase")     // "pascal.case"
```

---

## String Operations

### Reverse()

Reverse a string.

```go
func Reverse(s string) string
```

**Example:**
```go
stringx.Reverse("hello")  // "olleh"
stringx.Reverse("world")  // "dlrow"
stringx.Reverse("12345")  // "54321"
```

**Notes:**
- Optimized for ASCII strings
- Handles Unicode correctly
- Zero-allocation for ASCII

---

### SplitLen()

Split string by length.

```go
func SplitLen(s string, max int) []string
```

**Parameters:**
- `s` - String to split
- `max` - Maximum length of each part

**Returns:**
- Slice of strings

**Example:**
```go
stringx.SplitLen("abcdefghij", 3)  // ["abc", "def", "ghi", "j"]
stringx.SplitLen("hello", 2)       // ["he", "ll", "o"]
stringx.SplitLen("short", 10)      // ["short"]
```

---

### Shorten()

Shorten string to maximum length.

```go
func Shorten(s string, max int) string
```

**Example:**
```go
stringx.Shorten("hello world", 5)  // "hello"
stringx.Shorten("short", 10)       // "short"
stringx.Shorten("text", 0)         // ""
```

---

### ShortenShow()

Shorten string with ellipsis.

```go
func ShortenShow(s string, max int) string
```

**Example:**
```go
stringx.ShortenShow("hello world", 8)  // "hello..."
stringx.ShortenShow("short", 10)       // "short"
stringx.ShortenShow("text", 2)         // "..."
```

---

## Utility Functions

### IsUpper()

Check if string is uppercase.

```go
func IsUpper[M string | []rune](r M) bool
```

**Example:**
```go
stringx.IsUpper("HELLO")  // true
stringx.IsUpper("Hello")  // false
stringx.IsUpper("hello")  // false
```

---

### IsDigit()

Check if string contains only digits.

```go
func IsDigit[M string | []rune](r M) bool
```

**Example:**
```go
stringx.IsDigit("12345")  // true
stringx.IsDigit("123a5")  // false
stringx.IsDigit("abcde")  // false
```

---

### Quote()

Quote a string.

```go
func Quote(s string) string
```

**Example:**
```go
stringx.Quote("hello")  // `"hello"`
stringx.Quote(`"test"`)  // `"\"test\""`
```

---

### QuotePure()

Quote a string without outer quotes.

```go
func QuotePure(s string) string
```

**Example:**
```go
stringx.QuotePure("hello")  // `\"hello\"`
stringx.QuotePure(`"test"`)  // `\"test\"`
```

---

## Usage Patterns

### API Response Formatting

```go
func formatAPIResponse(data interface{}) string {
    jsonStr, _ := json.Marshal(data)
    return stringx.ShortenShow(jsonStr, 1000)
}
```

### Identifier Conversion

```go
func convertToDBColumn(fieldName string) string {
    return stringx.ToSnake(fieldName)
}

func convertToJSONField(dbColumn string) string {
    return stringx.Snake2SmallCamel(dbColumn)
}

func convertToStructField(dbColumn string) string {
    return stringx.Snake2Camel(dbColumn)
}
```

### Text Processing

```go
func processText(text string) string {
    // Convert to lowercase
    text = strings.ToLower(text)
    
    // Convert to snake_case
    text = stringx.ToSnake(text)
    
    // Shorten if too long
    text = stringx.ShortenShow(text, 100)
    
    return text
}
```

### Validation

```go
func validateUsername(username string) bool {
    // Check length
    if len(username) < 3 || len(username) > 20 {
        return false
    }
    
    // Check for digits only
    if stringx.IsDigit(username) {
        return false
    }
    
    return true
}
```

---

## Performance Characteristics

### Zero-Allocation Conversions

```go
// Fast: Zero allocation
data := []byte("hello")
str := stringx.ToString(data)

// Fast: Zero allocation
str := "hello"
data := stringx.ToBytes(str)
```

### Optimized Case Conversion

```go
// Fast: ASCII optimized
result := stringx.Camel2Snake("camelCase")

// Fast: Unicode aware
result := stringx.Camel2Snake("CamelCase中文")
```

### Benchmarks

| Operation | Time | Memory | Notes |
|-----------|------|--------|-------|
| `ToString()` | 0.5 ns/op | 0 B/op | Zero allocation |
| `ToBytes()` | 0.5 ns/op | 0 B/op | Zero allocation |
| `Camel2Snake()` | 45 ns/op | 32 B/op | ASCII optimized |
| `Reverse()` | 30 ns/op | 16 B/op | ASCII optimized |
| `SplitLen()` | 120 ns/op | 64 B/op | Depends on length |

---

## Best Practices

### Zero-Allocation Conversions

```go
// Good: Zero allocation
func processData(data []byte) string {
    return stringx.ToString(data)
}

// Avoid: Creates new string
func processDataBad(data []byte) string {
    return string(data)  // Allocates
}
```

### Unicode Handling

```go
// Good: Unicode aware
func reverseText(text string) string {
    return stringx.Reverse(text)  // Handles Unicode correctly
}

// Avoid: Breaks Unicode
func reverseTextBad(text string) string {
    runes := []rune(text)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
```

### Case Conversion

```go
// Good: Use appropriate conversion
func formatFieldName(field string) string {
    return stringx.ToSnake(field)  // Handles any format
}

// Avoid: Limited to specific format
func formatFieldNameBad(field string) string {
    return stringx.Camel2Snake(field)  // Only works for camelCase
}
```

---

## Related Documentation

- [candy](/en/modules/candy) - Type conversion
- [json](/en/modules/json) - JSON processing
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
