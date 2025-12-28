---
title: defaults - Default Values
---

# defaults - Default Values

## Overview

The defaults module provides default value setting for Go structs with support for various data types and custom default functions.

## Functions

### SetDefaults()

Set default values for a struct.

```go
func SetDefaults(value interface{}, opts ...Options) error
```

**Parameters:**
- `value` - Struct pointer to set defaults for
- `opts` - Options for default behavior

**Options:**
- `ErrorMode` - Error handling mode (Panic, Ignore, Return)
- `ValidateDefaults` - Whether to validate default values
- `AllowOverwrite` - Whether to allow overwriting non-zero values

**Example:**
```go
type Config struct {
    Host     string `default:"localhost"`
    Port     int    `default:"8080"`
    Debug    bool   `default:"false"`
    Timeout  int    `default:"30"`
}

var cfg Config
defaults.SetDefaults(&cfg)
// cfg.Host == "localhost"
// cfg.Port == 8080
// cfg.Debug == false
// cfg.Timeout == 30
```

---

### RegisterCustomDefault()

Register custom default function for a type.

```go
func RegisterCustomDefault(typeName string, fn DefaultFunc)
```

**Parameters:**
- `typeName` - Type name for registration
- `fn` - Default function

**Example:**
```go
defaults.RegisterCustomDefault("time.Time", func() interface{} {
    return time.Now()
})
```

---

## Usage Patterns

### Struct Initialization

```go
type User struct {
    Name     string `default:"Anonymous"`
    Email     string `default:""`
    Age       int    `default:"0"`
    Active    bool   `default:"true"`
    CreatedAt time.Time `default:"now"`
}

var user User
defaults.SetDefaults(&user)
// user.Name == "Anonymous"
// user.Email == ""
// user.Age == 0
// user.Active == true
// user.CreatedAt == time.Now()
```

### Custom Defaults

```go
type Product struct {
    Name     string
    Price     float64
    Stock     int
    Category  string
}

defaults.RegisterCustomDefault("Product.Price", func() interface{} {
    return 0.0
})

defaults.RegisterCustomDefault("Product.Stock", func() interface{} {
    return 100
})

var product Product
defaults.SetDefaults(&product)
// product.Price == 0.0
// product.Stock == 100
```

### Error Handling

```go
// Good: Panic on error
defaults.SetDefaults(&cfg, defaults.Options{
    ErrorMode: defaults.ErrorModePanic,
})

// Good: Ignore errors
defaults.SetDefaults(&cfg, defaults.Options{
    ErrorMode: defaults.ErrorModeIgnore,
})

// Good: Return errors
defaults.SetDefaults(&cfg, defaults.Options{
    ErrorMode: defaults.ErrorModeReturn,
})
```

---

## Best Practices

### Default Tags

```go
// Good: Use appropriate default values
type Config struct {
    Host     string `default:"localhost"`  // Good default
    Port     int    `default:"8080"`   // Good default
    Debug    bool   `default:"false"`   // Good default
    Timeout  int    `default:"30"`     // Good default
}

// Avoid: Unnecessary defaults
type Config struct {
    Host     string `default:""`         // Bad default
    Port     int    `default:"0"`         // Bad default
    Debug    bool   `default:"false"`   // Good default
}
```

---

## Related Documentation

- [validator](/en/modules/validator) - Data validation
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
