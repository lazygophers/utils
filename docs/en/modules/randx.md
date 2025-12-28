---
title: randx - Random Utilities
---

# randx - Random Utilities

## Overview

The randx module provides high-performance random number generation with cryptographically secure random support.

## Functions

### Int()

Generate random integer.

```go
func Int() int
```

**Returns:**
- Random integer

**Example:**
```go
num := randx.Int()
```

---

### Intn()

Generate random integer in range [0, n).

```go
func Intn(n int) int
```

**Parameters:**
- `n` - Maximum value (exclusive)

**Returns:**
- Random integer in [0, n)

**Example:**
```go
num := randx.Intn(10)  // 0-9
```

---

### IntnRange()

Generate random integer in range [min, max].

```go
func IntnRange(min, max int) int
```

**Parameters:**
- `min` - Minimum value (inclusive)
- `max` - Maximum value (inclusive)

**Returns:**
- Random integer in [min, max]

**Example:**
```go
num := randx.IntnRange(1, 10)  // 1-10
```

---

### Int64()

Generate random int64.

```go
func Int64() int64
```

**Returns:**
- Random int64

**Example:**
```go
num := randx.Int64()
```

---

### Float64()

Generate random float64.

```go
func Float64() float64
```

**Returns:**
- Random float64 in [0.0, 1.0)

**Example:**
go
num := randx.Float64()
```

---

### Float64Range()

Generate random float64 in range [min, max].

```go
func Float64Range(min, max float64) float64
```

**Parameters:**
- `min` - Minimum value
- `max` - Maximum value

**Returns:**
- Random float64 in [min, max]

**Example:**
```go
num := randx.Float64Range(0.0, 100.0)
```

---

### Bool()

Generate random boolean.

```go
func Bool() bool
```

**Returns:**
- Random boolean

**Example:**
```go
flag := randx.Bool()
```

---

## Usage Patterns

### Random Numbers

```go
// Random age
age := randx.IntnRange(18, 65)

// Random price
price := randx.Float64Range(10.0, 1000.0)

// Random quantity
quantity := randx.IntnRange(1, 100)
```

### Random Selection

```go
items := []string{"apple", "banana", "cherry", "date", "elderberry"}
index := randx.Intn(len(items))
selected := items[index]
```

### Random Data

```go
type User struct {
    ID       int64
    Name     string
    Age      int
    Active   bool
    Price    float64
}

user := User{
    ID:     randx.Int64(),
    Name:    "User",
    Age:     randx.IntnRange(18, 65),
    Active:  randx.Bool(),
    Price:   randx.Float64Range(0.0, 1000.0),
}
```

---

## Performance

The randx module uses high-performance random number generation with optimized implementations for better performance.

---

## Related Documentation

- [fake](/en/modules/fake) - Test data generation
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
