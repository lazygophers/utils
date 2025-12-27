---
title: API Documentation
---

# API Documentation

This document provides detailed API reference for LazyGophers Utils.

## Core Utilities

### utils.Must()

Assert operation succeeds, panic on failure.

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
```

### utils.MustSuccess()

Assert error is nil.

```go
func MustSuccess(err error)
```

**Parameters:**
- `err` - Error

**Example:**
```go
utils.MustSuccess(config.Load(&cfg, "config.json"))
```

### utils.MustOk()

Assert second return value is true.

```go
func MustOk[T any](value T, ok bool) T
```

**Parameters:**
- `value` - Return value
- `ok` - Success flag

**Returns:**
- Returns `value` if `ok` is true
- Panics if `ok` is false

**Example:**
```go
value := utils.MustOk(getValue())
```

### utils.Validate()

Validate struct data.

```go
func Validate(v any) error
```

**Parameters:**
- `v` - Struct to validate

**Returns:**
- Validation error if validation fails

**Example:**
```go
type User struct {
    Name  string `validate:"required"`
    Email string `validate:"required,email"`
    Age   int    `validate:"min=0,max=150"`
}

user := User{
    Name:  "John Doe",
    Email: "john@example.com",
    Age:   25,
}

if err := utils.Validate(&user); err != nil {
    fmt.Printf("Validation failed: %v\n", err)
}
```

## Data Processing

### candy.ToInt()

Convert string to integer.

```go
func ToInt(s string) int
```

**Parameters:**
- `s` - String

**Returns:**
- Integer value

**Example:**
```go
age := candy.ToInt("25")
```

### candy.ToFloat()

Convert string to float.

```go
func ToFloat(s string) float64
```

**Parameters:**
- `s` - String

**Returns:**
- Float value

**Example:**
```go
price := candy.ToFloat("99.99")
```

### candy.ToBool()

Convert string to boolean.

```go
func ToBool(s string) bool
```

**Parameters:**
- `s` - String

**Returns:**
- Boolean value

**Example:**
```go
active := candy.ToBool("true")
```

### candy.ToString()

Convert any type to string.

```go
func ToString(v any) string
```

**Parameters:**
- `v` - Any value

**Returns:**
- String

**Example:**
```go
str := candy.ToString(123)
```

## Time Processing

### xtime.NowCalendar()

Get current calendar.

```go
func NowCalendar() *Calendar
```

**Returns:**
- Current calendar object

**Example:**
```go
cal := xtime.NowCalendar()
fmt.Printf("Today: %s\n", cal.String())
```

### Calendar.LunarDate()

Get lunar date.

```go
func (c *Calendar) LunarDate() string
```

**Returns:**
- Lunar date string

**Example:**
```go
fmt.Printf("Lunar: %s\n", cal.LunarDate())
```

### Calendar.Animal()

Get zodiac animal.

```go
func (c *Calendar) Animal() string
```

**Returns:**
- Zodiac animal string

**Example:**
```go
fmt.Printf("Animal: %s\n", cal.Animal())
```

### Calendar.CurrentSolarTerm()

Get current solar term.

```go
func (c *Calendar) CurrentSolarTerm() string
```

**Returns:**
- Solar term string

**Example:**
```go
fmt.Printf("Solar Term: %s\n", cal.CurrentSolarTerm())
```

## Configuration Management

### config.Load()

Load configuration file.

```go
func Load(v any, filename string) error
```

**Parameters:**
- `v` - Config struct pointer
- `filename` - Configuration file name

**Returns:**
- Error if loading fails

**Supported formats:**
- JSON
- YAML
- TOML
- INI
- HCL

**Example:**
```go
type Config struct {
    Database string `json:"database"`
    Port     int    `json:"port"`
    Debug    bool   `json:"debug"`
}

var cfg Config
utils.MustSuccess(config.Load(&cfg, "config.json"))
```

## Concurrency Control

### routine.NewPool()

Create worker pool.

```go
func NewPool(size int) *Pool
```

**Parameters:**
- `size` - Pool size

**Returns:**
- Worker pool object

**Example:**
```go
pool := routine.NewPool(10)
defer pool.Close()
```

### Pool.Submit()

Submit task to worker pool.

```go
func (p *Pool) Submit(fn func())
```

**Parameters:**
- `fn` - Function to execute

**Example:**
```go
pool.Submit(func() {
    fmt.Println("Task executed")
})
```

### wait.For()

Wait for condition.

```go
func For(timeout time.Duration, condition func() bool) bool
```

**Parameters:**
- `timeout` - Timeout duration
- `condition` - Condition function

**Returns:**
- Whether condition was met before timeout

**Example:**
```go
success := wait.For(5*time.Second, func() bool {
    return pool.Running() == 0
})
```

## More APIs

For complete API documentation, visit:

- [pkg.go.dev](https://pkg.go.dev/github.com/lazygophers/utils)
- [GitHub Repository](https://github.com/lazygophers/utils)

## Related Documentation

- [Getting Started](/en/guide/getting-started)
- [Module Overview](/en/modules/overview)
