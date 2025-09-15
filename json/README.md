# json - Enhanced JSON Processing

The `json` package provides enhanced JSON processing utilities with better error handling, type safety, and convenience functions. It extends Go's standard `encoding/json` package with additional functionality for common JSON operations.

## Features

- **Enhanced Error Handling**: Better error messages and error context
- **Type-Safe Operations**: Generic functions for type-safe JSON operations
- **Convenience Functions**: Simplified Marshal/Unmarshal operations
- **Pretty Printing**: Formatted JSON output with customizable indentation
- **Path Operations**: JSON path-based value extraction and modification
- **Validation**: JSON schema validation and structure checking
- **Streaming Support**: Efficient streaming JSON processing

## Installation

```bash
go get github.com/lazygophers/utils/json
```

## Usage Examples

### Basic JSON Operations

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/json"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    user := User{ID: 1, Name: "Alice", Age: 30}

    // Marshal to JSON string
    jsonStr, err := json.MarshalString(user)
    if err != nil {
        panic(err)
    }
    fmt.Println(jsonStr) // {"id":1,"name":"Alice","age":30}

    // Unmarshal from JSON string
    var newUser User
    err = json.UnmarshalString(jsonStr, &newUser)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%+v\n", newUser)
}
```

### Pretty Printing

```go
user := User{ID: 1, Name: "Alice", Age: 30}

// Pretty print with default indentation
pretty, err := json.MarshalPretty(user)
if err != nil {
    panic(err)
}
fmt.Println(string(pretty))
// Output:
// {
//   "id": 1,
//   "name": "Alice",
//   "age": 30
// }

// Pretty print with custom indentation
pretty, err = json.MarshalIndent(user, "", "    ")
if err != nil {
    panic(err)
}
fmt.Println(string(pretty))
```

### Type-Safe Operations

```go
// Generic marshal function
data := map[string]int{"apple": 5, "banana": 3}
result, err := json.Marshal[map[string]int](data)
if err != nil {
    panic(err)
}

// Generic unmarshal function
var restored map[string]int
err = json.Unmarshal[map[string]int](result, &restored)
if err != nil {
    panic(err)
}
fmt.Println(restored) // map[apple:5 banana:3]
```

### JSON Path Operations

```go
jsonData := `{
    "user": {
        "id": 123,
        "profile": {
            "name": "Alice",
            "age": 30
        }
    },
    "settings": {
        "theme": "dark"
    }
}`

// Extract value by path
name, err := json.GetValueByPath(jsonData, "user.profile.name")
if err != nil {
    panic(err)
}
fmt.Println(name) // "Alice"

// Set value by path
modified, err := json.SetValueByPath(jsonData, "user.profile.age", 31)
if err != nil {
    panic(err)
}
fmt.Println(modified)
```

### Validation and Structure Checking

```go
// Validate JSON structure
valid := json.IsValidJSON(`{"name": "Alice", "age": 30}`)
fmt.Println(valid) // true

// Check if JSON has required fields
hasFields := json.HasFields(jsonData, []string{"user.id", "user.profile.name"})
fmt.Println(hasFields) // true

// Validate against schema
schema := `{
    "type": "object",
    "properties": {
        "name": {"type": "string"},
        "age": {"type": "number"}
    },
    "required": ["name", "age"]
}`

valid, err = json.ValidateSchema(jsonData, schema)
if err != nil {
    panic(err)
}
fmt.Println(valid)
```

### Streaming Operations

```go
// Streaming encoder
var buf bytes.Buffer
encoder := json.NewStreamEncoder(&buf)

users := []User{
    {ID: 1, Name: "Alice", Age: 30},
    {ID: 2, Name: "Bob", Age: 25},
}

for _, user := range users {
    err := encoder.Encode(user)
    if err != nil {
        panic(err)
    }
}

// Streaming decoder
decoder := json.NewStreamDecoder(&buf)
for decoder.More() {
    var user User
    err := decoder.Decode(&user)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%+v\n", user)
}
```

## API Reference

### Basic Functions

- `Marshal(v interface{}) ([]byte, error)` - Marshal value to JSON bytes
- `MarshalString(v interface{}) (string, error)` - Marshal value to JSON string
- `Unmarshal(data []byte, v interface{}) error` - Unmarshal JSON bytes
- `UnmarshalString(s string, v interface{}) error` - Unmarshal JSON string

### Pretty Printing

- `MarshalPretty(v interface{}) ([]byte, error)` - Marshal with default pretty formatting
- `MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)` - Marshal with custom indentation
- `PrettyPrint(data []byte) ([]byte, error)` - Format existing JSON data

### Generic Functions

- `Marshal[T any](v T) ([]byte, error)` - Type-safe marshal
- `Unmarshal[T any](data []byte, v *T) error` - Type-safe unmarshal
- `MarshalString[T any](v T) (string, error)` - Type-safe marshal to string

### Path Operations

- `GetValueByPath(jsonStr, path string) (interface{}, error)` - Get value by JSON path
- `SetValueByPath(jsonStr, path string, value interface{}) (string, error)` - Set value by path
- `DeleteByPath(jsonStr, path string) (string, error)` - Delete value by path
- `PathExists(jsonStr, path string) bool` - Check if path exists

### Validation

- `IsValidJSON(s string) bool` - Check if string is valid JSON
- `ValidateSchema(jsonStr, schema string) (bool, error)` - Validate against JSON schema
- `HasFields(jsonStr string, fields []string) bool` - Check for required fields
- `GetType(jsonStr, path string) (string, error)` - Get type of value at path

### Streaming

- `NewStreamEncoder(w io.Writer) *StreamEncoder` - Create streaming encoder
- `NewStreamDecoder(r io.Reader) *StreamDecoder` - Create streaming decoder

### Utility Functions

- `Merge(json1, json2 string) (string, error)` - Merge two JSON objects
- `Filter(jsonStr string, keys []string) (string, error)` - Filter JSON by keys
- `Transform(jsonStr string, transformer func(key, value interface{}) interface{}) (string, error)` - Transform JSON values
- `Compare(json1, json2 string) (bool, error)` - Compare two JSON structures

## Error Handling

The package provides enhanced error information:

```go
type JSONError struct {
    Op   string // Operation that failed
    Path string // JSON path where error occurred
    Err  error  // Underlying error
}

func (e *JSONError) Error() string {
    return fmt.Sprintf("json: %s at path %s: %v", e.Op, e.Path, e.Err)
}
```

### Error Types

- `IsSyntaxError(err error) bool` - Check for JSON syntax errors
- `IsTypeError(err error) bool` - Check for type conversion errors
- `IsPathError(err error) bool` - Check for path-related errors

## Performance Optimizations

- **Zero-Copy Operations**: Minimize memory allocations where possible
- **Streaming Support**: Process large JSON files efficiently
- **Path Caching**: Cache frequently used JSON paths
- **Pool Reuse**: Reuse buffers and decoders for better performance

## Best Practices

1. **Use Type-Safe Functions**: Prefer generic functions for compile-time safety
2. **Handle Errors**: Always check and handle JSON errors appropriately
3. **Validate Input**: Validate JSON structure before processing
4. **Use Streaming**: Use streaming for large JSON data sets
5. **Cache Paths**: Cache frequently accessed JSON paths

## Examples

### Configuration Processing

```go
configJSON := `{
    "server": {
        "host": "localhost",
        "port": 8080,
        "ssl": {
            "enabled": true,
            "cert": "/path/to/cert.pem"
        }
    },
    "database": {
        "url": "postgres://localhost/mydb",
        "pool_size": 10
    }
}`

// Extract specific configuration values
host, _ := json.GetValueByPath(configJSON, "server.host")
port, _ := json.GetValueByPath(configJSON, "server.port")
sslEnabled, _ := json.GetValueByPath(configJSON, "server.ssl.enabled")

fmt.Printf("Server: %s:%d (SSL: %v)\n", host, port, sslEnabled)
```

### API Response Processing

```go
type APIResponse struct {
    Status  string      `json:"status"`
    Data    interface{} `json:"data"`
    Message string      `json:"message,omitempty"`
}

// Create and marshal response
response := APIResponse{
    Status: "success",
    Data:   map[string]interface{}{"user_id": 123, "name": "Alice"},
}

jsonResponse, err := json.MarshalPretty(response)
if err != nil {
    panic(err)
}

fmt.Println(string(jsonResponse))
```

### Data Transformation

```go
// Transform JSON data
originalJSON := `{"prices": [10, 20, 30], "currency": "USD"}`

transformed, err := json.Transform(originalJSON, func(key, value interface{}) interface{} {
    if key == "prices" && reflect.TypeOf(value).Kind() == reflect.Slice {
        // Convert prices from USD to EUR (example rate)
        prices := value.([]interface{})
        for i, p := range prices {
            if price, ok := p.(float64); ok {
                prices[i] = price * 0.85 // Example conversion rate
            }
        }
        return prices
    }
    if key == "currency" {
        return "EUR"
    }
    return value
})

if err != nil {
    panic(err)
}
fmt.Println(transformed)
```

## Related Packages

- `candy` - Type conversion utilities
- `validator` - Struct validation utilities
- `stringx` - String manipulation utilities