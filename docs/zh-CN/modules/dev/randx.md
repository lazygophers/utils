---
title: randx - 随机工具
---

# randx - 随机工具

## 概述

randx 模块提供高性能随机数生成,支持加密安全的随机数。

## 函数

### Int()

生成随机整数。

```go
func Int() int
```

**返回:**
- 随机整数

**示例:**
```go
num := randx.Int()
```

---

### Intn()

生成范围 [0, n) 内的随机整数。

```go
func Intn(n int) int
```

**参数:**
- `n` - 最大值(不包含)

**返回:**
- [0, n) 内的随机整数

**示例:**
```go
num := randx.Intn(10)  // 0-9
```

---

### IntnRange()

生成范围 [min, max] 内的随机整数。

```go
func IntnRange(min, max int) int
```

**参数:**
- `min` - 最小值(包含)
- `max` - 最大值(包含)

**返回:**
- [min, max] 内的随机整数

**示例:**
```go
num := randx.IntnRange(1, 10)  // 1-10
```

---

### Int64()

生成随机 int64。

```go
func Int64() int64
```

**返回:**
- 随机 int64

**示例:**
```go
num := randx.Int64()
```

---

### Float64()

生成随机 float64。

```go
func Float64() float64
```

**返回:**
- [0.0, 1.0) 内的随机 float64

**示例:**
```go
num := randx.Float64()
```

---

### Float64Range()

生成范围 [min, max] 内的随机 float64。

```go
func Float64Range(min, max float64) float64
```

**参数:**
- `min` - 最小值
- `max` - 最大值

**返回:**
- [min, max] 内的随机 float64

**示例:**
```go
num := randx.Float64Range(0.0, 100.0)
```

---

### Bool()

生成随机布尔值。

```go
func Bool() bool
```

**返回:**
- 随机布尔值

**示例:**
```go
flag := randx.Bool()
```

---

## 使用模式

### 随机数字

```go
// 随机年龄
age := randx.IntnRange(18, 65)

// 随机价格
price := randx.Float64Range(10.0, 1000.0)

// 随机数量
quantity := randx.IntnRange(1, 100)
```

### 随机选择

```go
items := []string{"apple", "banana", "cherry", "date", "elderberry"}
index := randx.Intn(len(items))
selected := items[index]
```

### 随机数据

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

## 性能

randx 模块使用高性能随机数生成,具有优化实现以获得更好的性能。

---

## 相关文档

- [fake](/zh-CN/modules/fake) - 测试数据生成
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
