---
title: randx - 隨機工具
---

# randx - 隨機工具

## 概述

randx 模組提供高性能隨機數生成,支持加密安全的隨機數。

## 函數

### Int()

生成隨機整數。

```go
func Int() int
```

**返回:**
- 隨機整數

**示例:**
```go
num := randx.Int()
```

---

### Intn()

生成範圍 [0, n) 內的隨機整數。

```go
func Intn(n int) int
```

**參數:**
- `n` - 最大值(不包含)

**返回:**
- [0, n) 內的隨機整數

**示例:**
```go
num := randx.Intn(10)  // 0-9
```

---

### IntnRange()

生成範圍 [min, max] 內的隨機整數。

```go
func IntnRange(min, max int) int
```

**參數:**
- `min` - 最小值(包含)
- `max` - 最大值(包含)

**返回:**
- [min, max] 內的隨機整數

**示例:**
```go
num := randx.IntnRange(1, 10)  // 1-10
```

---

### Int64()

生成隨機 int64。

```go
func Int64() int64
```

**返回:**
- 隨機 int64

**示例:**
```go
num := randx.Int64()
```

---

### Float64()

生成隨機 float64。

```go
func Float64() float64
```

**返回:**
- [0.0, 1.0) 內的隨機 float64

**示例:**
```go
num := randx.Float64()
```

---

### Float64Range()

生成範圍 [min, max] 內的隨機 float64。

```go
func Float64Range(min, max float64) float64
```

**參數:**
- `min` - 最小值
- `max` - 最大值

**返回:**
- [min, max] 內的隨機 float64

**示例:**
```go
num := randx.Float64Range(0.0, 100.0)
```

---

### Bool()

生成隨機布爾值。

```go
func Bool() bool
```

**返回:**
- 隨機布爾值

**示例:**
```go
flag := randx.Bool()
```

---

## 使用模式

### 隨機數字

```go
// 隨機年齡
age := randx.IntnRange(18, 65)

// 隨機價格
price := randx.Float64Range(10.0, 1000.0)

// 隨機數量
quantity := randx.IntnRange(1, 100)
```

### 隨機選擇

```go
items := []string{"apple", "banana", "cherry", "date", "elderberry"}
index := randx.Intn(len(items))
selected := items[index]
```

### 隨機數據

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

randx 模組使用高性能隨機數生成,具有優化實現以獲得更好的性能。

---

## 相關文檔

- [fake](/zh-TW/modules/fake) - 測試數據生成
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
