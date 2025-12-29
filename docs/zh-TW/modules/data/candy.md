---
title: candy - 類型轉換
---

# candy - 類型轉換

## 概述

candy 模組提供高效能類型轉換工具，具有零分配優化。它支援各種 Go 類型之間的轉換，包括字串、數字、布林值、切片和映射。

## 整數轉換函數

### ToInt()

將任何類型轉換為 int。

```go
func ToInt(val interface{}) int
```

**支援的類型：**
- bool: true → 1, false → 0
- 所有整數類型（int、int8、int16、int32、int64）
- 所無符號整數類型（uint、uint8、uint16、uint32、uint64）
- 浮點類型（float32、float64）
- string、[]byte: 解析為整數
- 其他類型: 返回 0

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

轉換為特定的整數類型。

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

轉換為無符號整數類型。

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

## 浮點轉換函數

### ToFloat(), ToFloat32(), ToFloat64()

轉換為浮點類型。

```go
func ToFloat(val interface{}) float64
func ToFloat32(val interface{}) float32
func ToFloat64(val interface{}) float64
```

**支援的類型：**
- bool: true → 1.0, false → 0.0
- 所有整數類型
- 浮點類型
- string、[]byte: 解析為浮點數
- 其他類型: 返回 0.0

**示例：**
```go
candy.ToFloat("123.45")    // 123.45
candy.ToFloat(123)         // 123.0
candy.ToFloat(true)        // 1.0
candy.ToFloat32("3.14")   // 3.14
candy.ToFloat64("2.71828") // 2.71828
```

---

## 布林轉換

### ToBool()

將任何類型轉換為布林值。

```go
func ToBool(val interface{}) bool
```

**支援的類型：**
- bool: 原樣返回
- 數字: 非零 → true, 零 → false
- string、[]byte: "1"、"true"、"yes"、"on" → true
- 其他類型: 返回 false

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

## 字串轉換

### ToString()

將任何類型轉換為字串。

```go
func ToString(val interface{}) string
```

**支援的類型：**
- bool: true → "1", false → "0"
- 所有整數類型: 轉換為十進制字串
- 浮點類型: 使用適當的精度轉換
- time.Duration: 格式化為持續時間字串
- string: 原樣返回
- []byte: 轉換為字串
- nil: 返回 ""
- error: 返回錯誤訊息
- 其他類型: JSON 序列化

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

## 切片轉換

### ToSlice()

將任何類型轉換為切片。

```go
func ToSlice(val interface{}) []interface{}
```

**支援的類型：**
- 任何類型的陣列和切片
- 映射: 轉換為值的切片
- 其他類型: 包裝在單元素切片中

**示例：**
```go
candy.ToSlice([]int{1, 2, 3})        // []interface{}{1, 2, 3}
candy.ToSlice([3]string{"a", "b", "c"}) // []interface{}{"a", "b", "c"}
candy.ToSlice(map[string]int{"a": 1, "b": 2}) // []interface{}{1, 2}
candy.ToSlice(123)                      // []interface{}{123}
```

---

### ToIntSlice(), ToStringSlice(), 等

轉換為類型化切片。

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

## 映射轉換

### ToMap()

將任何類型轉換為映射。

```go
func ToMap(val interface{}) map[string]interface{}
```

**支援的類型：**
- map[string]interface{}: 原樣返回
- map[interface{}]interface{}: 鍵轉換為字串
- 結構體: 欄位轉換為映射條目
- 其他類型: 返回空映射

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

## 指針轉換

### ToPtr()

將任何類型轉換為指針。

```go
func ToPtr[T any](val T) *T
```

**示例：**
```go
candy.ToPtr(123)           // *int 值為 123
candy.ToPtr("hello")       // *string 值為 "hello"
candy.ToPtr(true)          // *bool 值為 true
```

---

## 使用模式

### HTTP 請求解析

```go
func handleRequest(r *http.Request) {
    query := r.URL.Query()
    
    // 解析查詢參數
    page := candy.ToInt(query.Get("page"))
    limit := candy.ToInt(query.Get("limit"))
    active := candy.ToBool(query.Get("active"))
    
    // 使用解析的值
    fmt.Printf("頁碼: %d, 限制: %d, 活躍: %v\n", page, limit, active)
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

### 資料處理

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

### 資料庫操作

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

## 效能特徵

### 零分配轉換

candy 模組針對效能進行了優化，盡可能實現零分配轉換：

```go
// 快速：直接類型轉換
val := candy.ToInt("123")

// 快速：無堆分配
val := candy.ToInt64(123)

// 快速：直接轉換
val := candy.ToString(123)
```

### 基準測試

| 操作 | 時間 | 記憶體 | vs 標準庫 |
|-----------|------|--------|-------------------|
| `ToInt()` | 12.3 ns/op | 0 B/op | **快 3.2 倍** |
| `ToFloat64()` | 15.7 ns/op | 0 B/op | **快 2.8 倍** |
| `ToString()` | 8.9 ns/op | 0 B/op | **快 4.1 倍** |
| `ToBool()` | 5.2 ns/op | 0 B/op | **快 5.3 倍** |

---

## 最佳實踐

### 類型安全

```go
// 好：優雅地處理轉換錯誤
func safeToInt(val interface{}) (int, error) {
    result := candy.ToInt(val)
    if result == 0 && val != 0 {
        return 0, fmt.Errorf("轉換失敗")
    }
    return result, nil
}

// 更好：對關鍵轉換使用類型斷言
func criticalToInt(val interface{}) (int, error) {
    if i, ok := val.(int); ok {
        return i, nil
    }
    return candy.ToInt(val), nil
}
```

### 錯誤處理

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

## 相關文檔

- [stringx](/zh-TW/modules/stringx) - 字串工具
- [anyx](/zh-TW/modules/anyx) - Interface{} 助手
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
