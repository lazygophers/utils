---
title: anyx - Interface{} Helpers
---

# anyx - Interface{} Helpers

## Overview

The anyx module provides type-safe operations on `interface{}` values with convenient methods for accessing and converting data. It wraps `sync.Map` for thread-safe operations.

## Core Type

### MapAny

Thread-safe map with type-safe accessors.

```go
type MapAny struct {
    data *sync.Map
    cut  *atomic.Bool
    seq  *atomic.String
}
```

---

## Constructors

### NewMap()

Create a new MapAny from a map.

```go
func NewMap(m map[string]interface{}) *MapAny
```

**Example:**
```go
m := anyx.NewMap(map[string]interface{}{
    "name":  "John",
    "age":   30,
    "email": "john@example.com",
})
```

---

### NewMapWithJson()

Create a new MapAny from JSON bytes.

```go
func NewMapWithJson(s []byte) (*MapAny, error)
```

**Example:**
```go
data := []byte(`{"name":"John","age":30}`)
m, err := anyx.NewMapWithJson(data)
if err != nil {
    log.Errorf("Failed to parse JSON: %v", err)
}
```

---

### NewMapWithYaml()

Create a new MapAny from YAML bytes.

```go
func NewMapWithYaml(s []byte) (*MapAny, error)
```

**Example:**
```go
data := []byte("name: John\nage: 30")
m, err := anyx.NewMapWithYaml(data)
if err != nil {
    log.Errorf("Failed to parse YAML: %v", err)
}
```

---

### NewMapWithAny()

Create a new MapAny from any value.

```go
func NewMapWithAny(s interface{}) (*MapAny, error)
```

**Example:**
```go
m, err := anyx.NewMapWithAny(struct {
    Name string
    Age  int
}{Name: "John", Age: 30})
```

---

## Configuration Methods

### EnableCut()

Enable nested key access with separator.

```go
func (p *MapAny) EnableCut(seq string) *MapAny
```

**Parameters:**
- `seq` - Separator string (e.g., ".", "/")

**Returns:**
- MapAny instance for chaining

**Example:**
```go
m := anyx.NewMap(map[string]interface{}{
    "user": map[string]interface{}{
        "name": "John",
        "email": "john@example.com",
    },
})

m.EnableCut(".")
name := m.GetString("user.name")  // "John"
```

---

### DisableCut()

Disable nested key access.

```go
func (p *MapAny) DisableCut() *MapAny
```

**Example:**
```go
m.EnableCut(".").DisableCut()
```

---

## Basic Operations

### Set()

Set a key-value pair.

```go
func (p *MapAny) Set(key string, value interface{})
```

**Example:**
```go
m := anyx.NewMap(nil)
m.Set("name", "John")
m.Set("age", 30)
m.Set("active", true)
```

---

### Get()

Get a value by key.

```go
func (p *MapAny) Get(key string) (interface{}, error)
```

**Returns:**
- Value if key exists
- `ErrNotFound` if key does not exist

**Example:**
```go
value, err := m.Get("name")
if err != nil {
    log.Errorf("Key not found: %v", err)
} else {
    fmt.Printf("Name: %v\n", value)
}
```

---

### Exists()

Check if a key exists.

```go
func (p *MapAny) Exists(key string) bool
```

**Example:**
```go
if m.Exists("name") {
    fmt.Println("Name exists")
}
```

---

## Type-Safe Getters

### GetBool()

Get value as boolean.

```go
func (p *MapAny) GetBool(key string) bool
```

**Example:**
```go
m.Set("active", true)
active := m.GetBool("active")  // true
```

---

### GetInt()

Get value as int.

```go
func (p *MapAny) GetInt(key string) int
```

**Example:**
```go
m.Set("age", 30)
age := m.GetInt("age")  // 30
```

---

### GetInt32(), GetInt64()

Get value as int32 or int64.

```go
func (p *MapAny) GetInt32(key string) int32
func (p *MapAny) GetInt64(key string) int64
```

---

### GetUint16(), GetUint32(), GetUint64()

Get value as unsigned integer.

```go
func (p *MapAny) GetUint16(key string) uint16
func (p *MapAny) GetUint32(key string) uint32
func (p *MapAny) GetUint64(key string) uint64
```

---

### GetFloat64()

Get value as float64.

```go
func (p *MapAny) GetFloat64(key string) float64
```

**Example:**
```go
m.Set("price", 19.99)
price := m.GetFloat64("price")  // 19.99
```

---

### GetString()

Get value as string.

```go
func (p *MapAny) GetString(key string) string
```

**Example:**
```go
m.Set("name", "John")
name := m.GetString("name")  // "John"
```

---

### GetBytes()

Get value as []byte.

```go
func (p *MapAny) GetBytes(key string) []byte
```

---

### GetMap()

Get value as nested MapAny.

```go
func (p *MapAny) GetMap(key string) *MapAny
```

**Example:**
```go
m.Set("user", map[string]interface{}{
    "name": "John",
    "email": "john@example.com",
})

user := m.GetMap("user")
name := user.GetString("name")  // "John"
```

---

### GetSlice()

Get value as []interface{}.

```go
func (p *MapAny) GetSlice(key string) []interface{}
```

**Example:**
```go
m.Set("tags", []interface{}{"go", "utils", "library"})
tags := m.GetSlice("tags")  // []interface{}{"go", "utils", "library"}
```

---

### GetStringSlice()

Get value as []string.

```go
func (p *MapAny) GetStringSlice(key string) []string
```

**Example:**
```go
m.Set("tags", []interface{}{"go", "utils"})
tags := m.GetStringSlice("tags")  // []string{"go", "utils"}
```

---

### GetInt64Slice()

Get value as []int64.

```go
func (p *MapAny) GetInt64Slice(key string) []int64
```

---

### GetUint32Slice()

Get value as []uint32.

```go
func (p *MapAny) GetUint32Slice(key string) []uint32
```

---

### GetUint64Slice()

Get value as []uint64.

```go
func (p *MapAny) GetUint64Slice(key string) []uint64
```

---

## Conversion Methods

### ToMap()

Convert to standard map[string]interface{}.

```go
func (p *MapAny) ToMap() map[string]interface{}
```

**Example:**
```go
m := anyx.NewMap(map[string]interface{}{
    "name": "John",
    "age":  30,
})

stdMap := m.ToMap()
// stdMap is map[string]interface{}{"name": "John", "age": 30}
```

---

### ToSyncMap()

Convert to *sync.Map.

```go
func (p *MapAny) ToSyncMap() *sync.Map
```

---

### Clone()

Create a copy of MapAny.

```go
func (p *MapAny) Clone() *MapAny
```

**Example:**
```go
original := anyx.NewMap(map[string]interface{}{
    "name": "John",
})

copy := original.Clone()
copy.Set("name", "Jane")
// original.GetString("name") == "John"
// copy.GetString("name") == "Jane"
```

---

## Iteration

### Range()

Iterate over all key-value pairs.

```go
func (p *MapAny) Range(f func(key, value interface{}) bool)
```

**Example:**
```go
m.Range(func(key, value interface{}) bool {
    fmt.Printf("%s: %v\n", key, value)
    return true  // Continue iteration
})
```

---

## Usage Patterns

### Configuration Management

```go
func loadConfig() *anyx.MapAny {
    data, _ := os.ReadFile("config.json")
    cfg, _ := anyx.NewMapWithJson(data)
    return cfg
}

func getConfig(cfg *anyx.MapAny, key string) string {
    return cfg.GetString(key)
}

func setConfig(cfg *anyx.MapAny, key string, value string) {
    cfg.Set(key, value)
}
```

### Nested Data Access

```go
m := anyx.NewMap(map[string]interface{}{
    "database": map[string]interface{}{
        "host": "localhost",
        "port": 5432,
        "credentials": map[string]interface{}{
            "username": "admin",
            "password": "secret",
        },
    },
})

m.EnableCut(".")
host := m.GetString("database.host")  // "localhost"
port := m.GetInt("database.port")  // 5432
username := m.GetString("database.credentials.username")  // "admin"
```

### Type Conversion

```go
func processData(data *anyx.MapAny) {
    name := data.GetString("name")
    age := data.GetInt("age")
    active := data.GetBool("active")
    price := data.GetFloat64("price")
    tags := data.GetStringSlice("tags")
    
    fmt.Printf("Name: %s, Age: %d, Active: %v\n", name, age, active)
    fmt.Printf("Price: %.2f, Tags: %v\n", price, tags)
}
```

### Thread-Safe Operations

```go
var cache = anyx.NewMap(nil)

func getFromCache(key string) (interface{}, bool) {
    val, err := cache.Get(key)
    if err == anyx.ErrNotFound {
        return nil, false
    }
    return val, true
}

func setCache(key string, value interface{}) {
    cache.Set(key, value)
}
```

---

## Best Practices

### Error Handling

```go
// Good: Handle not found errors
value, err := m.Get("key")
if err != nil {
    if err == anyx.ErrNotFound {
        fmt.Println("Key not found")
    } else {
        fmt.Printf("Error: %v\n", err)
    }
}

// Good: Use Exists() before Get
if m.Exists("key") {
    value, _ := m.Get("key")
    fmt.Printf("Value: %v\n", value)
}
```

### Type Safety

```go
// Good: Use type-safe getters
age := m.GetInt("age")
name := m.GetString("name")
active := m.GetBool("active")

// Avoid: Manual type assertions
value, _ := m.Get("age")
age, ok := value.(int)
if !ok {
    // Handle type mismatch
}
```

### Nested Access

```go
// Good: Enable cut for nested access
m.EnableCut(".")
value := m.GetString("user.profile.name")

// Alternative: Chain GetMap()
user := m.GetMap("user")
profile := user.GetMap("profile")
name := profile.GetString("name")
```

---

## Related Documentation

- [candy](/en/modules/candy) - Type conversion
- [json](/en/modules/json) - JSON processing
- [validator](/en/modules/validator) - Data validation
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
