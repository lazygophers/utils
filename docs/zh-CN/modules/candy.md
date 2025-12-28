---
title: candy - 类型转换
---

# candy - 类型转换

## 概述

candy 模块提供高性能类型转换工具，具有零分配优化。它支持各种 Go 类型之间的转换，包括字符串、数字、布尔值、切片和映射。

## 整数转换函数

### ToInt()

将任何类型转换为 int。

```go
func ToInt(val interface{}) int
```

**支持的类型：**
- bool: true → 1, false → 0
- 所有整数类型（int、int8、int16、int32、int64）
- 所有无符号整数类型（uint、uint8、uint16、uint32、uint64）
- 浮点类型（float32、float64）
- string、[]byte: 解析为整数
- 其他类型: 返回 0

**示例：**
```go
candy.ToInt("123")        // 123
candy.ToInt(45.67)       // 45
candy.ToInt(true)         // 1
candy.ToInt(false)        // 0
candy.ToInt([]byte("99")) // 99
```

---

### ToInt8(), ToInt16(), ToInt32(), ToInt64()

转换为特定的整数类型。

```go
func ToInt8(val interface{}) int8
func ToInt16(val interface{}) int16
func ToInt32(val interface{}) int32
func ToInt64(val interface{}) int64
```

**示例：**
```go
candy.ToInt8("127")      // 127
candy.ToInt16("32767")    // 32767
candy.ToInt32("2147483647") // 2147483647
candy.ToInt64("9223372036854775807") // 9223372036854775807
```

---

### ToUint(), ToUint8(), ToUint16(), ToUint32(), ToUint64()

转换为无符号整数类型。

```go
func ToUint(val interface{}) uint
func ToUint8(val interface{}) uint8
func ToUint16(val interface{}) uint16
func ToUint32(val interface{}) uint32
func ToUint64(val interface{}) uint64
```

**示例：**
```go
candy.ToUint("255")       // 255
candy.ToUint8("255")      // 255
candy.ToUint16("65535")   // 65535
candy.ToUint32("4294967295") // 4294967295
candy.ToUint64("18446744073709551615") // 18446744073709551615
```

---

## 浮点转换函数

### ToFloat(), ToFloat32(), ToFloat64()

转换为浮点类型。

```go
func ToFloat(val interface{}) float64
func ToFloat32(val interface{}) float32
func ToFloat64(val interface{}) float64
```

**支持的类型：**
- bool: true → 1.0, false → 0.0
- 所有整数类型
- 浮点类型
- string、[]byte: 解析为浮点数
- 其他类型: 返回 0.0

**示例：**
```go
candy.ToFloat("123.45")    // 123.45
candy.ToFloat(123)         // 123.0
candy.ToFloat(true)        // 1.0
candy.ToFloat32("3.14")   // 3.14
candy.ToFloat64("2.71828") // 2.71828
```

---

## 布尔转换

### ToBool()

将任何类型转换为布尔值。

```go
func ToBool(val interface{}) bool
```

**支持的类型：**
- bool: 原样返回
- 数字: 非零 → true, 零 → false
- string、[]byte: "1"、"true"、"yes"、"on" → true
- 其他类型: 返回 false

**示例：**
```go
candy.ToBool(true)        // true
candy.ToBool(false)       // false
candy.ToBool(1)          // true
candy.ToBool(0)          // false
candy.ToBool("true")     // true
candy.ToBool("false")    // false
candy.ToBool("yes")      // true
```

---

## 字符串转换

### ToString()

将任何类型转换为字符串。

```go
func ToString(val interface{}) string
```

**支持的类型：**
- bool: true → "1", false → "0"
- 所有整数类型: 转换为十进制字符串
- 浮点类型: 使用适当的精度转换
- time.Duration: 格式化为持续时间字符串
- string: 原样返回
- []byte: 转换为字符串
- nil: 返回 ""
- error: 返回错误消息
- 其他类型: JSON 序列化

**示例：**
```go
candy.ToString(123)           // "123"
candy.ToString(45.67)         // "45.670000"
candy.ToString(true)          // "1"
candy.ToString([]byte("abc"))  // "abc"
candy.ToString(errors.New("error")) // "error"
candy.ToString(time.Hour)      // "1h0m0s"
```

---

## 切片转换

### ToSlice()

将任何类型转换为切片。

```go
func ToSlice(val interface{}) []interface{}
```

**支持的类型：**
- 任何类型的数组和切片
- 映射: 转换为值的切片
- 其他类型: 包装在单元素切片中

**示例：**
```go
candy.ToSlice([]int{1, 2, 3})        // []interface{}{1, 2, 3}
candy.ToSlice([3]string{"a", "b", "c"}) // []interface{}{"a", "b", "c"}
candy.ToSlice(map[string]int{"a": 1, "b": 2}) // []interface{}{1, 2}
candy.ToSlice(123)                      // []interface{}{123}
```

---

### ToIntSlice(), ToStringSlice(), 等

转换为类型化切片。

```go
func ToIntSlice(val interface{}) []int
func ToStringSlice(val interface{}) []string
func ToInt64Slice(val interface{}) []int64
func ToUint64Slice(val interface{}) []uint64
```

**示例：**
```go
candy.ToIntSlice([]interface{}{1, 2, 3})          // []int{1, 2, 3}
candy.ToIntSlice([]string{"1", "2", "3"})          // []int{1, 2, 3}
candy.ToStringSlice([]interface{}{"a", "b", "c"})    // []string{"a", "b", "c"}
candy.ToInt64Slice([]int{1, 2, 3})               // []int64{1, 2, 3}
candy.ToUint64Slice([]int{1, 2, 3})              // []uint64{1, 2, 3}
```

---

## 映射转换

### ToMap()

将任何类型转换为映射。

```go
func ToMap(val interface{}) map[string]interface{}
```

**支持的类型：**
- map[string]interface{}: 原样返回
- map[interface{}]interface{}: 键转换为字符串
- 结构体: 字段转换为映射条目
- 其他类型: 返回空映射

**示例：**
```go
type User struct {
    Name  string
    Email string
}

user := User{Name: "John", Email: "john@example.com"}
candy.ToMap(user)
// map[string]interface{}{"Name": "John", "Email": "john@example.com"}

candy.ToMap(map[interface{}]interface{}{1: "a", 2: "b"})
// map[string]interface{}{"1": "a", "2": "b"}
```

---

## 指针转换

### ToPtr()

将任何类型转换为指针。

```go
func ToPtr[T any](val T) *T
```

**示例：**
```go
candy.ToPtr(123)           // *int 值为 123
candy.ToPtr("hello")       // *string 值为 "hello"
candy.ToPtr(true)          // *bool 值为 true
```

---

## 使用模式

### HTTP 请求解析

```go
func handleRequest(r *http.Request) {
    query := r.URL.Query()
    
    // 解析查询参数
    page := candy.ToInt(query.Get("page"))
    limit := candy.ToInt(query.Get("limit"))
    active := candy.ToBool(query.Get("active"))
    
    // 使用解析的值
    fmt.Printf("页码: %d, 限制: %d, 活跃: %v\n", page, limit, active)
}
```

### 配置解析

```go
type Config struct {
    Port     int
    Debug    bool
    Timeout  int
}

func loadConfigFromEnv() *Config {
    return &Config{
        Port:     candy.ToInt(os.Getenv("PORT")),
        Debug:    candy.ToBool(os.Getenv("DEBUG")),
        Timeout:  candy.ToInt(os.Getenv("TIMEOUT")),
    }
}
```

### 数据处理

```go
func processData(data []interface{}) []int {
    var result []int
    for _, item := range data {
        result = append(result, candy.ToInt(item))
    }
    return result
}

func processStrings(data []interface{}) []string {
    return candy.ToStringSlice(data)
}
```

### 数据库操作

```go
func convertRow(row []interface{}) map[string]interface{} {
    return candy.ToMap(row)
}

func convertRows(rows [][]interface{}) []map[string]interface{} {
    var result []map[string]interface{}
    for _, row := range rows {
        result = append(result, convertRow(row))
    }
    return result
}
```

---

## 性能特征

### 零分配转换

candy 模块针对性能进行了优化，尽可能实现零分配转换：

```go
// 快速：直接类型转换
val := candy.ToInt("123")

// 快速：无堆分配
val := candy.ToInt64(123)

// 快速：直接转换
val := candy.ToString(123)
```

### 基准测试

| 操作 | 时间 | 内存 | vs 标准库 |
|-----------|------|--------|-------------------|
| `ToInt()` | 12.3 ns/op | 0 B/op | **快 3.2 倍** |
| `ToFloat64()` | 15.7 ns/op | 0 B/op | **快 2.8 倍** |
| `ToString()` | 8.9 ns/op | 0 B/op | **快 4.1 倍** |
| `ToBool()` | 5.2 ns/op | 0 B/op | **快 5.3 倍** |

---

## 最佳实践

### 类型安全

```go
// 好：优雅地处理转换错误
func safeToInt(val interface{}) (int, error) {
    result := candy.ToInt(val)
    if result == 0 && val != 0 {
        return 0, fmt.Errorf("转换失败")
    }
    return result, nil
}

// 更好：对关键转换使用类型断言
func criticalToInt(val interface{}) (int, error) {
    if i, ok := val.(int); ok {
        return i, nil
    }
    return candy.ToInt(val), nil
}
```

### 错误处理

```go
func convertWithDefault(val interface{}, defaultVal int) int {
    result := candy.ToInt(val)
    if result == 0 && val != nil {
        return defaultVal
    }
    return result
}

func safeConvert(val interface{}) (int, bool) {
    result := candy.ToInt(val)
    success := result != 0 || val == 0
    return result, success
}
```

---

## 相关文档

- [stringx](/zh-CN/modules/stringx) - 字符串工具
- [anyx](/zh-CN/modules/anyx) - Interface{} 助手
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
