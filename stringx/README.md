# stringx - Advanced String Manipulation Utilities

The `stringx` package provides high-performance string manipulation utilities with zero-copy optimizations, Unicode support, and advanced string processing functions. It extends Go's standard `strings` package with additional functionality for common string operations.

## Features

- **Zero-Copy Conversions**: Efficient string/byte slice conversions using unsafe operations
- **Case Conversions**: Camel case, snake case, kebab case, and other case conversions
- **String Generation**: Random string generation with customizable character sets
- **Unicode Support**: Full Unicode support for text processing
- **Performance Optimized**: Memory-efficient operations with minimal allocations
- **Validation Functions**: String validation and checking utilities
- **Text Processing**: Advanced text manipulation and formatting functions

## Installation

```bash
go get github.com/lazygophers/utils/stringx
```

## Usage Examples

### Zero-Copy String/Byte Conversions

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/stringx"
)

func main() {
    // Zero-copy string to bytes conversion
    str := "hello world"
    bytes := stringx.ToBytes(str)
    fmt.Printf("String: %s, Bytes: %v\n", str, bytes)

    // Zero-copy bytes to string conversion
    data := []byte("hello world")
    result := stringx.ToString(data)
    fmt.Printf("Bytes: %v, String: %s\n", data, result)
}
```

### Case Conversions

```go
// Camel case to snake case
camelCase := "getUserProfile"
snakeCase := stringx.Camel2Snake(camelCase)
fmt.Println(snakeCase) // get_user_profile

// Snake case to camel case
snakeStr := "user_profile_data"
camelStr := stringx.Snake2Camel(snakeStr)
fmt.Println(camelStr) // userProfileData

// Kebab case conversions
kebabCase := stringx.Camel2Kebab("getUserProfile")
fmt.Println(kebabCase) // get-user-profile

// Pascal case conversions
pascalCase := stringx.Snake2Pascal("user_profile")
fmt.Println(pascalCase) // UserProfile
```

### Random String Generation

```go
// Generate random string with default charset (alphanumeric)
randomStr := stringx.RandString(10)
fmt.Println(randomStr) // e.g., "Kj8mNq2Lp9"

// Generate random string with custom charset
customStr := stringx.RandStringWithCharset(8, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
fmt.Println(customStr) // e.g., "MXKQWERT"

// Generate random alphanumeric string
alphaNum := stringx.RandAlphaNumeric(12)
fmt.Println(alphaNum) // e.g., "Abc123Def456"

// Generate random alphabetic string
alpha := stringx.RandAlphabetic(6)
fmt.Println(alpha) // e.g., "XyZabc"

// Generate random numeric string
numeric := stringx.RandNumeric(4)
fmt.Println(numeric) // e.g., "1234"
```

### String Validation and Checking

```go
// Check if string is alphanumeric
isAlphaNum := stringx.IsAlphaNumeric("Abc123")
fmt.Println(isAlphaNum) // true

// Check if string is alphabetic only
isAlpha := stringx.IsAlphabetic("HelloWorld")
fmt.Println(isAlpha) // true

// Check if string is numeric only
isNumeric := stringx.IsNumeric("12345")
fmt.Println(isNumeric) // true

// Check if string is ASCII
isASCII := stringx.IsASCII("Hello World")
fmt.Println(isASCII) // true

// Check if string is empty or whitespace
isEmpty := stringx.IsBlank("   ")
fmt.Println(isEmpty) // true
```

### Text Processing

```go
// Reverse string
reversed := stringx.Reverse("hello")
fmt.Println(reversed) // "olleh"

// Capitalize first letter
capitalized := stringx.Capitalize("hello world")
fmt.Println(capitalized) // "Hello world"

// Title case
title := stringx.ToTitle("hello world")
fmt.Println(title) // "Hello World"

// Center string with padding
centered := stringx.Center("hello", 10, " ")
fmt.Println(centered) // "  hello   "

// Truncate string with ellipsis
truncated := stringx.Truncate("This is a long string", 10)
fmt.Println(truncated) // "This is..."

// Remove duplicates from string
noDupes := stringx.RemoveDuplicates("hello")
fmt.Println(noDupes) // "helo"
```

### Advanced String Operations

```go
// Split string and trim spaces
parts := stringx.SplitAndTrim("apple, banana, cherry", ",")
fmt.Println(parts) // ["apple", "banana", "cherry"]

// Join non-empty strings
joined := stringx.JoinNonEmpty([]string{"hello", "", "world", ""}, " ")
fmt.Println(joined) // "hello world"

// Extract words from string
words := stringx.ExtractWords("Hello, World! How are you?")
fmt.Println(words) // ["Hello", "World", "How", "are", "you"]

// Count occurrences of substring
count := stringx.CountOccurrences("hello world hello", "hello")
fmt.Println(count) // 2

// Replace multiple substrings
replacements := map[string]string{"hello": "hi", "world": "earth"}
replaced := stringx.ReplaceMultiple("hello world", replacements)
fmt.Println(replaced) // "hi earth"
```

## API Reference

### Zero-Copy Conversions

- `ToString(b []byte) string` - Convert byte slice to string (zero-copy)
- `ToBytes(s string) []byte` - Convert string to byte slice (zero-copy)

### Case Conversions

- `Camel2Snake(s string) string` - Convert camelCase to snake_case
- `Snake2Camel(s string) string` - Convert snake_case to camelCase
- `Camel2Kebab(s string) string` - Convert camelCase to kebab-case
- `Snake2Pascal(s string) string` - Convert snake_case to PascalCase
- `Pascal2Snake(s string) string` - Convert PascalCase to snake_case
- `Kebab2Camel(s string) string` - Convert kebab-case to camelCase

### Random String Generation

- `RandString(length int) string` - Generate random alphanumeric string
- `RandStringWithCharset(length int, charset string) string` - Generate with custom charset
- `RandAlphaNumeric(length int) string` - Generate alphanumeric string
- `RandAlphabetic(length int) string` - Generate alphabetic string
- `RandNumeric(length int) string` - Generate numeric string
- `RandHex(length int) string` - Generate hexadecimal string
- `RandBase64(length int) string` - Generate base64 string

### String Validation

- `IsAlphaNumeric(s string) bool` - Check if string is alphanumeric
- `IsAlphabetic(s string) bool` - Check if string is alphabetic
- `IsNumeric(s string) bool` - Check if string is numeric
- `IsASCII(s string) bool` - Check if string contains only ASCII characters
- `IsBlank(s string) bool` - Check if string is empty or whitespace
- `IsEmpty(s string) bool` - Check if string is empty
- `IsUpper(s string) bool` - Check if string is uppercase
- `IsLower(s string) bool` - Check if string is lowercase

### Text Processing

- `Reverse(s string) string` - Reverse string
- `Capitalize(s string) string` - Capitalize first letter
- `ToTitle(s string) string` - Convert to title case
- `Center(s string, width int, fillChar string) string` - Center with padding
- `PadLeft(s string, width int, padChar string) string` - Left pad
- `PadRight(s string, width int, padChar string) string` - Right pad
- `Truncate(s string, length int) string` - Truncate with ellipsis
- `TruncateWords(s string, wordCount int) string` - Truncate by word count

### Advanced Operations

- `SplitAndTrim(s, sep string) []string` - Split and trim whitespace
- `JoinNonEmpty(strs []string, sep string) string` - Join non-empty strings
- `ExtractWords(s string) []string` - Extract words from text
- `CountOccurrences(s, substr string) int` - Count substring occurrences
- `ReplaceMultiple(s string, replacements map[string]string) string` - Multiple replacements
- `RemoveDuplicates(s string) string` - Remove duplicate characters
- `Similarity(s1, s2 string) float64` - Calculate string similarity

### Unicode Support

- `ContainsUnicode(s string) bool` - Check if string contains Unicode
- `UnicodeLength(s string) int` - Get Unicode character count
- `UnicodeSubstring(s string, start, length int) string` - Unicode-aware substring
- `NormalizeSpaces(s string) string` - Normalize whitespace characters

## Performance Characteristics

The package is optimized for high performance:

### Zero-Copy Operations
```go
// These operations don't allocate new memory
bytes := stringx.ToBytes("hello")    // O(1), no allocation
str := stringx.ToString([]byte{...}) // O(1), no allocation
```

### Optimized Case Conversions
- ASCII-only strings use optimized fast path
- Unicode strings use efficient buffering
- Pre-calculated capacity to minimize reallocations

### Memory-Efficient Random Generation
- Reuses character sets and buffers
- Minimal allocations for repeated calls

## Best Practices

1. **Use Zero-Copy Functions**: Use `ToString()` and `ToBytes()` for efficient conversions
2. **Choose Appropriate Functions**: Use specific validation functions instead of regex
3. **Pre-allocate for Bulk Operations**: When processing many strings, consider pre-allocation
4. **Handle Unicode Properly**: Use Unicode-aware functions for international text
5. **Validate Input**: Always validate string input in public APIs

## Examples

### Configuration Key Conversion

```go
// Convert configuration keys between formats
configKeys := []string{
    "database_host",
    "database_port",
    "api_timeout",
    "log_level",
}

// Convert to environment variable format
for _, key := range configKeys {
    envKey := strings.ToUpper(key)
    fmt.Printf("%s -> %s\n", key, envKey)
}

// Convert to camelCase for JSON
for _, key := range configKeys {
    jsonKey := stringx.Snake2Camel(key)
    fmt.Printf("%s -> %s\n", key, jsonKey)
}
```

### Text Processing Pipeline

```go
// Process user input text
userInput := "  Hello, WORLD!  How are YOU today?  "

// Clean and normalize
cleaned := strings.TrimSpace(userInput)
normalized := stringx.NormalizeSpaces(cleaned)
title := stringx.ToTitle(normalized)

fmt.Println(title) // "Hello, World! How Are You Today?"

// Extract and analyze words
words := stringx.ExtractWords(normalized)
fmt.Printf("Word count: %d\n", len(words))
```

### ID and Token Generation

```go
// Generate various types of identifiers
sessionID := stringx.RandHex(32)
apiKey := stringx.RandBase64(24)
userToken := stringx.RandAlphaNumeric(16)

fmt.Printf("Session ID: %s\n", sessionID)
fmt.Printf("API Key: %s\n", apiKey)
fmt.Printf("User Token: %s\n", userToken)
```

## Related Packages

- `candy` - Type conversion utilities
- `validator` - String validation utilities
- `json` - JSON string processing