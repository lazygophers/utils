---
title: must.go - Error Assertion
---

# must.go - Error Assertion

## Overview

`must.go` provides error assertion utilities that simplify error handling flow by providing panic-based error checking. These functions are designed for initialization and critical operations where failure should halt program execution.

## Functions

### MustOk()

Assert second return value is true.

```go
func MustOk[T any](value T, ok bool) T
```

**Parameters:**
- `value` - Return value
- `ok` - Success flag

**Returns:**
- Returns `value` if `ok` is true
- Panics with message "is not ok" if `ok` is false

**Example:**
```go
value, ok := getValue()
result := utils.MustOk(value, ok)
```

**Notes:**
- Use this function when second return value indicates success/failure
- Panic will halt program execution if assertion fails
- Generic type `T` allows any return type
- Commonly used with map lookups and type assertions

---

### MustSuccess()

Assert error is nil.

```go
func MustSuccess(err error)
```

**Parameters:**
- `err` - Error to check

**Behavior:**
- Does nothing if `err` is nil
- Panics with formatted error message if `err` is not nil

**Example:**
```go
utils.MustSuccess(config.Load(&cfg, "config.json"))
utils.MustSuccess(os.MkdirAll("data", 0755))
utils.MustSuccess(db.Ping())
```

**Notes:**
- Commonly used for initialization and setup operations
- Panic message includes error details for debugging
- Use for operations that must succeed for the program to continue

---

### Must()

Combine validation function that checks error status and returns value.

```go
func Must[T any](value T, err error) T
```

**Parameters:**
- `value` - Return value
- `err` - Error

**Returns:**
- Returns `value` if `err` is nil
- Panics if `err` is not nil

**Example:**
```go
data := utils.Must(loadData())
file := utils.Must(os.Open("file.txt"))
result := utils.Must(http.Get(url))
conn := utils.Must(net.Listen("tcp", ":8080"))
```

**Notes:**
- Most commonly used function in must module
- Combines error checking and value extraction
- Generic type `T` allows any return type
- Ideal for functions that return `(T, error)`

---

### Ignore()

Forcefully ignore any parameter.

```go
func Ignore[T any](value T, _ any) T
```

**Parameters:**
- `value` - Value to return
- `_` - Ignored parameter

**Returns:**
- Returns `value`

**Example:**
```go
result := utils.Ignore(data, err)
```

**Notes:**
- Used when you need to suppress linter warnings about ignored values
- The second parameter is explicitly ignored
- Useful for maintaining clean code without linter warnings
- Does not actually handle the error, just suppresses warnings

---

## Usage Patterns

### Initialization Chain

```go
func initApp() {
    cfg := utils.Must(loadConfig())
    db := utils.Must(connectDB(cfg.DatabaseURL))
    server := utils.Must(createServer(cfg.Port))
    
    utils.MustSuccess(server.Start())
}
```

### File Operations

```go
func readFile(path string) []byte {
    file := utils.Must(os.Open(path))
    defer file.Close()
    
    data := utils.Must(io.ReadAll(file))
    return data
}
```

### Map Operations

```go
func getValue(m map[string]int, key string) int {
    value, ok := m[key]
    return utils.MustOk(value, ok)
}
```

### Configuration Loading

```go
type Config struct {
    Host string `json:"host"`
    Port int    `json:"port"`
}

func loadConfig(path string) *Config {
    data := utils.Must(os.ReadFile(path))
    
    var cfg Config
    utils.MustSuccess(json.Unmarshal(data, &cfg))
    
    return &cfg
}
```

## Best Practices

### When to Use Must Functions

**Use `Must()` when:**
- The operation is critical to program startup
- Failure should halt program execution
- Error recovery is not possible or meaningful
- You're in initialization code (main, init)

**Avoid `Must()` when:**
- Handling user input
- Network requests that may fail
- File operations that may not exist
- Any operation that can reasonably fail

### Error Handling vs Panic

```go
// Good: Use Must for initialization
func init() {
    config = utils.Must(loadConfig())
}

// Good: Handle errors for user operations
func handleUserRequest() error {
    data, err := loadData()
    if err != nil {
        return err
    }
    // process data
    return nil
}
```

## Related Documentation

- [orm.go](/en/modules/orm) - Database operations
- [validator](/en/modules/validator) - Data validation
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
