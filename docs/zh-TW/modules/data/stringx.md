---
title: stringx - 字串工具
---

# stringx - 字串工具

## 概述

stringx 模組提供高效能字串工具，具有 Unicode 感知操作和零分配優化。

## 核心函數

### ToString()

將 []byte 轉換為字串，零分配。

```go
func ToString(b []byte) string
```

**參數：**
- `b` - 位元組切片

**返回值：**
- 字串表示

**示例：**
```go
data := []byte("hello")
str := stringx.ToString(data)
// str 是 "hello"，零分配
```

**注意：**
- 使用 unsafe 實現零分配轉換
- 對於大切片比 `string(b)` 更快
- 對於 nil 或空輸入返回空字串

---

### ToBytes()

將字串轉換為 []byte，零分配。

```go
func ToBytes(s string) []byte
```

**參數：**
- `s` - 字串

**返回值：**
- 位元組切片表示

**示例：**
```go
str := "hello"
data := stringx.ToBytes(str)
// data 是 []byte("hello")，零分配
```

**注意：**
- 使用 unsafe 實現零分配轉換
- 對於大字串比 `[]byte(s)` 更快
- 對於空字串返回 nil

---

## 大小寫轉換

### Camel2Snake()

將 camelCase 轉換為 snake_case。

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

將 snake_case 轉換為 CamelCase。

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

將 snake_case 轉換為 camelCase（小駝峰）。

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

將任何字串轉換為 snake_case。

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

將任何字串轉換為 kebab-case。

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

將任何字串轉換為 CamelCase。

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

將任何字串轉換為 camelCase。

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

將任何字串轉換為 slash/case。

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

將任何字串轉換為 dot.case。

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

## 字串操作

### Reverse()

反轉字串。

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
- 針對 ASCII 字串優化
- 正確處理 Unicode
- 對於 ASCII 零分配

---

### SplitLen()

按長度分割字串。

```go
func SplitLen(s string, max int) []string
```

**參數：**
- `s` - 要分割的字串
- `max` - 每部分的最大長度

**返回值：**
- 字串切片

**示例：**
```go
stringx.SplitLen("abcdefghij", 3)  // ["abc", "def", "ghi", "j"]
stringx.SplitLen("hello", 2)       // ["he", "ll", "o"]
stringx.SplitLen("short", 10)      // ["short"]
```

---

### Shorten()

將字串縮短到最大長度。

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

用省略號縮短字串。

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

## 工具函數

### IsUpper()

檢查字串是否為大寫。

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

檢查字串是否僅包含數字。

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

引用字串。

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

引用字串，不帶外層引號。

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

### API 響應格式化

```go
func formatAPIResponse(data interface{}) string {
    jsonStr, _ := json.Marshal(data)
    return stringx.ShortenShow(jsonStr, 1000)
}
```

### 標識符轉換

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

### 文本處理

```go
func processText(text string) string {
    // 轉換為小寫
    text = strings.ToLower(text)
    
    // 轉換為 snake_case
    text = stringx.ToSnake(text)
    
    // 如果太長則縮短
    text = stringx.ShortenShow(text, 100)
    
    return text
}
```

### 驗證

```go
func validateUsername(username string) bool {
    // 檢查長度
    if len(username) < 3 || len(username) > 20 {
        return false
    }
    
    // 檢查是否僅包含數字
    if stringx.IsDigit(username) {
        return false
    }
    
    return true
}
```

---

## 效能特徵

### 零分配轉換

```go
// 快速：零分配
data := []byte("hello")
str := stringx.ToString(data)

// 快速：零分配
str := "hello"
data := stringx.ToBytes(str)
```

### 優化的轉換大小寫

```go
// 快速：ASCII 優化
result := stringx.Camel2Snake("camelCase")

// 快速：Unicode 感知
result := stringx.Camel2Snake("CamelCase中文")
```

### 基準測試

| 操作 | 時間 | 記憶體 | 備註 |
|-----------|------|--------|-------|
| `ToString()` | 0.5 ns/op | 0 B/op | 零分配 |
| `ToBytes()` | 0.5 ns/op | 0 B/op | 零分配 |
| `Camel2Snake()` | 45 ns/op | 32 B/op | ASCII 優化 |
| `Reverse()` | 30 ns/op | 16 B/op | ASCII 優化 |
| `SplitLen()` | 120 ns/op | 64 B/op | 取決於長度 |

---

## 最佳實踐

### 零分配轉換

```go
// 好：零分配
func processData(data []byte) string {
    return stringx.ToString(data)
}

// 避免：建立新字串
func processDataBad(data []byte) string {
    return string(data)  // 分配
}
```

### Unicode 處理

```go
// 好：Unicode 感知
func reverseText(text string) string {
    return stringx.Reverse(text)  // 正確處理 Unicode
}

// 避免：破壞 Unicode
func reverseTextBad(text string) string {
    runes := []rune(text)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
```

### 大小寫轉換

```go
// 好：使用適當的轉換
func formatFieldName(field string) string {
    return stringx.ToSnake(field)  // 處理任何格式
}

// 避免：限於特定格式
func formatFieldNameBad(field string) string {
    return stringx.Camel2Snake(field)  // 僅適用於 camelCase
}
```

---

## 相關文檔

- [candy](/zh-TW/modules/candy) - 類型轉換
- [json](/zh-TW/modules/json) - JSON 處理
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
