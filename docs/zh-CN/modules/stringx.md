---
title: stringx - 字符串工具
---

# stringx - 字符串工具

## 概述

stringx 模块提供高性能字符串工具，具有 Unicode 感知操作和零分配优化。

## 核心函数

### ToString()

将 []byte 转换为字符串，零分配。

```go
func ToString(b []byte) string
```

**参数：**
- `b` - 字节切片

**返回值：**
- 字符串表示

**示例：**
```go
data := []byte("hello")
str := stringx.ToString(data)
// str 是 "hello"，零分配
```

**注意：**
- 使用 unsafe 实现零分配转换
- 对于大切片比 `string(b)` 更快
- 对于 nil 或空输入返回空字符串

---

### ToBytes()

将字符串转换为 []byte，零分配。

```go
func ToBytes(s string) []byte
```

**参数：**
- `s` - 字符串

**返回值：**
- 字节切片表示

**示例：**
```go
str := "hello"
data := stringx.ToBytes(str)
// data 是 []byte("hello")，零分配
```

**注意：**
- 使用 unsafe 实现零分配转换
- 对于大字符串比 `[]byte(s)` 更快
- 对于空字符串返回 nil

---

## 大小写转换

### Camel2Snake()

将 camelCase 转换为 snake_case。

```go
func Camel2Snake(s string) string
```

**示例：**
```go
stringx.Camel2Snake("camelCase")   // "camel_case"
stringx.Camel2Snake("MyVariable")  // "my_variable"
stringx.Camel2Snake("HTTPRequest") // "http_request"
```

---

### Snake2Camel()

将 snake_case 转换为 CamelCase。

```go
func Snake2Camel(s string) string
```

**示例：**
```go
stringx.Snake2Camel("snake_case")   // "SnakeCase"
stringx.Snake2Camel("my_variable")  // "MyVariable"
stringx.Snake2Camel("http_request") // "HttpRequest"
```

---

### Snake2SmallCamel()

将 snake_case 转换为 camelCase（小驼峰）。

```go
func Snake2SmallCamel(s string) string
```

**示例：**
```go
stringx.Snake2SmallCamel("snake_case")   // "snakeCase"
stringx.Snake2SmallCamel("my_variable")  // "myVariable"
stringx.Snake2SmallCamel("http_request") // "httpRequest"
```

---

### ToSnake()

将任何字符串转换为 snake_case。

```go
func ToSnake(s string) string
```

**示例：**
```go
stringx.ToSnake("camelCase")      // "camel_case"
stringx.ToSnake("CamelCase")      // "camel_case"
stringx.ToSnake("kebab-case")     // "kebab_case"
stringx.ToSnake("PascalCase")     // "pascal_case"
```

---

### ToKebab()

将任何字符串转换为 kebab-case。

```go
func ToKebab(s string) string
```

**示例：**
```go
stringx.ToKebab("camelCase")      // "camel-case"
stringx.ToKebab("snake_case")     // "snake-case"
stringx.ToKebab("PascalCase")     // "pascal-case"
```

---

### ToCamel()

将任何字符串转换为 CamelCase。

```go
func ToCamel(s string) string
```

**示例：**
```go
stringx.ToCamel("snake_case")     // "SnakeCase"
stringx.ToCamel("kebab-case")     // "KebabCase"
stringx.ToCamel("pascal_case")    // "PascalCase"
```

---

### ToSmallCamel()

将任何字符串转换为 camelCase。

```go
func ToSmallCamel(s string) string
```

**示例：**
```go
stringx.ToSmallCamel("snake_case")     // "snakeCase"
stringx.ToSmallCamel("kebab-case")     // "kebabCase"
stringx.ToSmallCamel("PascalCase")     // "pascalCase"
```

---

### ToSlash()

将任何字符串转换为 slash/case。

```go
func ToSlash(s string) string
```

**示例：**
```go
stringx.ToSlash("camelCase")      // "camel/case"
stringx.ToSlash("snake_case")     // "snake/case"
stringx.ToSlash("PascalCase")     // "pascal/case"
```

---

### ToDot()

将任何字符串转换为 dot.case。

```go
func ToDot(s string) string
```

**示例：**
```go
stringx.ToDot("camelCase")      // "camel.case"
stringx.ToDot("snake_case")     // "snake.case"
stringx.ToDot("PascalCase")     // "pascal.case"
```

---

## 字符串操作

### Reverse()

反转字符串。

```go
func Reverse(s string) string
```

**示例：**
```go
stringx.Reverse("hello")  // "olleh"
stringx.Reverse("world")  // "dlrow"
stringx.Reverse("12345")  // "54321"
```

**注意：**
- 针对 ASCII 字符串优化
- 正确处理 Unicode
- 对于 ASCII 零分配

---

### SplitLen()

按长度分割字符串。

```go
func SplitLen(s string, max int) []string
```

**参数：**
- `s` - 要分割的字符串
- `max` - 每部分的最大长度

**返回值：**
- 字符串切片

**示例：**
```go
stringx.SplitLen("abcdefghij", 3)  // ["abc", "def", "ghi", "j"]
stringx.SplitLen("hello", 2)       // ["he", "ll", "o"]
stringx.SplitLen("short", 10)      // ["short"]
```

---

### Shorten()

将字符串缩短到最大长度。

```go
func Shorten(s string, max int) string
```

**示例：**
```go
stringx.Shorten("hello world", 5)  // "hello"
stringx.Shorten("short", 10)       // "short"
stringx.Shorten("text", 0)         // ""
```

---

### ShortenShow()

用省略号缩短字符串。

```go
func ShortenShow(s string, max int) string
```

**示例：**
```go
stringx.ShortenShow("hello world", 8)  // "hello..."
stringx.ShortenShow("short", 10)       // "short"
stringx.ShortenShow("text", 2)         // "..."
```

---

## 工具函数

### IsUpper()

检查字符串是否为大写。

```go
func IsUpper[M string | []rune](r M) bool
```

**示例：**
```go
stringx.IsUpper("HELLO")  // true
stringx.IsUpper("Hello")  // false
stringx.IsUpper("hello")  // false
```

---

### IsDigit()

检查字符串是否仅包含数字。

```go
func IsDigit[M string | []rune](r M) bool
```

**示例：**
```go
stringx.IsDigit("12345")  // true
stringx.IsDigit("123a5")  // false
stringx.IsDigit("abcde")  // false
```

---

### Quote()

引用字符串。

```go
func Quote(s string) string
```

**示例：**
```go
stringx.Quote("hello")  // `"hello"`
stringx.Quote(`"test"`)  // `"\"test\""`
```

---

### QuotePure()

引用字符串，不带外层引号。

```go
func QuotePure(s string) string
```

**示例：**
```go
stringx.QuotePure("hello")  // `\"hello\"`
stringx.QuotePure(`"test"`)  // `\"test\""`
```

---

## 使用模式

### API 响应格式化

```go
func formatAPIResponse(data interface{}) string {
    jsonStr, _ := json.Marshal(data)
    return stringx.ShortenShow(jsonStr, 1000)
}
```

### 标识符转换

```go
func convertToDBColumn(fieldName string) string {
    return stringx.ToSnake(fieldName)
}

func convertToJSONField(dbColumn string) string {
    return stringx.Snake2SmallCamel(dbColumn)
}

func convertToStructField(dbColumn string) string {
    return stringx.Snake2Camel(dbColumn)
}
```

### 文本处理

```go
func processText(text string) string {
    // 转换为小写
    text = strings.ToLower(text)
    
    // 转换为 snake_case
    text = stringx.ToSnake(text)
    
    // 如果太长则缩短
    text = stringx.ShortenShow(text, 100)
    
    return text
}
```

### 验证

```go
func validateUsername(username string) bool {
    // 检查长度
    if len(username) < 3 || len(username) > 20 {
        return false
    }
    
    // 检查是否仅包含数字
    if stringx.IsDigit(username) {
        return false
    }
    
    return true
}
```

---

## 性能特征

### 零分配转换

```go
// 快速：零分配
data := []byte("hello")
str := stringx.ToString(data)

// 快速：零分配
str := "hello"
data := stringx.ToBytes(str)
```

### 优化的转换大小写

```go
// 快速：ASCII 优化
result := stringx.Camel2Snake("camelCase")

// 快速：Unicode 感知
result := stringx.Camel2Snake("CamelCase中文")
```

### 基准测试

| 操作 | 时间 | 内存 | 备注 |
|-----------|------|--------|-------|
| `ToString()` | 0.5 ns/op | 0 B/op | 零分配 |
| `ToBytes()` | 0.5 ns/op | 0 B/op | 零分配 |
| `Camel2Snake()` | 45 ns/op | 32 B/op | ASCII 优化 |
| `Reverse()` | 30 ns/op | 16 B/op | ASCII 优化 |
| `SplitLen()` | 120 ns/op | 64 B/op | 取决于长度 |

---

## 最佳实践

### 零分配转换

```go
// 好：零分配
func processData(data []byte) string {
    return stringx.ToString(data)
}

// 避免：创建新字符串
func processDataBad(data []byte) string {
    return string(data)  // 分配
}
```

### Unicode 处理

```go
// 好：Unicode 感知
func reverseText(text string) string {
    return stringx.Reverse(text)  // 正确处理 Unicode
}

// 避免：破坏 Unicode
func reverseTextBad(text string) string {
    runes := []rune(text)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
```

### 大小写转换

```go
// 好：使用适当的转换
func formatFieldName(field string) string {
    return stringx.ToSnake(field)  // 处理任何格式
}

// 避免：限于特定格式
func formatFieldNameBad(field string) string {
    return stringx.Camel2Snake(field)  // 仅适用于 camelCase
}
```

---

## 相关文档

- [candy](/zh-CN/modules/candy) - 类型转换
- [json](/zh-CN/modules/json) - JSON 处理
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
