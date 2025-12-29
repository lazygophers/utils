---
title: anyx - Interface{} 助手
---

# anyx - Interface{} 助手

## 概述

anyx 模組提供對 `interface{}` 值的類型安全操作，具有訪問和轉換資料的便捷方法。它包裝 `sync.Map` 以實現線程安全操作。

## 核心類型

### MapAny

具有類型安全訪問器的線程安全映射。

```go
type MapAny struct {
    data *sync.Map
    cut  *atomic.Bool
    seq  *atomic.String
}
```

---

## 建構函數

### NewMap()

從映射建立新的 MapAny。

```go
func NewMap(m map[string]interface{}) *MapAny
```

**示例：**
```go
m := anyx.NewMap(map[string]interface{}{
    "name":  "John",
    "age":   30,
    "email": "john@example.com",
})
```

---

### NewMapWithJson()

從 JSON 位元組建立新的 MapAny。

```go
func NewMapWithJson(s []byte) (*MapAny, error)
```

**示例：**
```go
data := []byte(`{"name":"John","age":30}`)
m, err := anyx.NewMapWithJson(data)
if err != nil {
    log.Errorf("解析 JSON 失敗: %v", err)
}
```

---

### NewMapWithYaml()

從 YAML 位元組建立新的 MapAny。

```go
func NewMapWithYaml(s []byte) (*MapAny, error)
```

**示例：**
```go
data := []byte("name: John\nage: 30")
m, err := anyx.NewMapWithYaml(data)
if err != nil {
    log.Errorf("解析 YAML 失敗: %v", err)
}
```

---

### NewMapWithAny()

從任何值建立新的 MapAny。

```go
func NewMapWithAny(s interface{}) (*MapAny, error)
```

**示例：**
```go
m, err := anyx.NewMapWithAny(struct {
    Name string
    Age  int
}{Name: "John", Age: 30})
```

---

## 配置方法

### EnableCut()

啟用帶分隔符的嵌套鍵訪問。

```go
func (p *MapAny) EnableCut(seq string) *MapAny
```

**參數：**
- `seq` - 分隔符字串（例如 "."、"/"）

**返回值：**
- MapAny 實例用於鏈式調用

**示例：**
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

禁用嵌套鍵訪問。

```go
func (p *MapAny) DisableCut() *MapAny
```

**示例：**
```go
m.EnableCut(".").DisableCut()
```

---

## 基本操作

### Set()

設置鍵值對。

```go
func (p *MapAny) Set(key string, value interface{})
```

**示例：**
```go
m := anyx.NewMap(nil)
m.Set("name", "John")
m.Set("age", 30)
m.Set("active", true)
```

---

### Get()

按鍵獲取值。

```go
func (p *MapAny) Get(key string) (interface{}, error)
```

**返回值：**
- 如果鍵存在，返回值
- 如果鍵不存在，返回 `ErrNotFound`

**示例：**
```go
value, err := m.Get("name")
if err != nil {
    log.Errorf("鍵未找到: %v", err)
} else {
    fmt.Printf("名稱: %v\n", value)
}
```

---

### Exists()

檢查鍵是否存在。

```go
func (p *MapAny) Exists(key string) bool
```

**示例：**
```go
if m.Exists("name") {
    fmt.Println("名稱存在")
}
```

---

## 類型安全獲取器

### GetBool()

獲取布林值。

```go
func (p *MapAny) GetBool(key string) bool
```

**示例：**
```go
m.Set("active", true)
active := m.GetBool("active")  // true
```

---

### GetInt()

獲取 int 值。

```go
func (p *MapAny) GetInt(key string) int
```

**示例：**
```go
m.Set("age", 30)
age := m.GetInt("age")  // 30
```

---

### GetInt32(), GetInt64()

獲取 int32 或 int64 值。

```go
func (p *MapAny) GetInt32(key string) int32
func (p *MapAny) GetInt64(key string) int64
```

---

### GetUint16(), GetUint32(), GetUint64()

獲取無符號整數值。

```go
func (p *MapAny) GetUint16(key string) uint16
func (p *MapAny) GetUint32(key string) uint32
func (p *MapAny) GetUint64(key string) uint64
```

---

### GetFloat64()

獲取 float64 值。

```go
func (p *MapAny) GetFloat64(key string) float64
```

**示例：**
```go
m.Set("price", 19.99)
price := m.GetFloat64("price")  // 19.99
```

---

### GetString()

獲取字串值。

```go
func (p *MapAny) GetString(key string) string
```

**示例：**
```go
m.Set("name", "John")
name := m.GetString("name")  // "John"
```

---

### GetBytes()

獲取 []byte 值。

```go
func (p *MapAny) GetBytes(key string) []byte
```

---

### GetMap()

獲取嵌套的 MapAny。

```go
func (p *MapAny) GetMap(key string) *MapAny
```

**示例：**
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

獲取 []interface{} 值。

```go
func (p *MapAny) GetSlice(key string) []interface{}
```

**示例：**
```go
m.Set("tags", []interface{}{"go", "utils", "library"})
tags := m.GetSlice("tags")  // []interface{}{"go", "utils", "library"}
```

---

### GetStringSlice()

獲取 []string 值。

```go
func (p *MapAny) GetStringSlice(key string) []string
```

**示例：**
```go
m.Set("tags", []interface{}{"go", "utils"})
tags := m.GetStringSlice("tags")  // []string{"go", "utils"}
```

---

### GetInt64Slice()

獲取 []int64 值。

```go
func (p *MapAny) GetInt64Slice(key string) []int64
```

---

### GetUint32Slice()

獲取 []uint32 值。

```go
func (p *MapAny) GetUint32Slice(key string) []uint32
```

---

### GetUint64Slice()

獲取 []uint64 值。

```go
func (p *MapAny) GetUint64Slice(key string) []uint64
```

---

## 轉換方法

### ToMap()

轉換為標準 map[string]interface{}。

```go
func (p *MapAny) ToMap() map[string]interface{}
```

**示例：**
```go
m := anyx.NewMap(map[string]interface{}{
    "name": "John",
    "age":  30,
})

stdMap := m.ToMap()
// stdMap 是 map[string]interface{}{"name": "John", "age": 30}
```

---

### ToSyncMap()

轉換為 *sync.Map。

```go
func (p *MapAny) ToSyncMap() *sync.Map
```

---

### Clone()

建立 MapAny 的副本。

```go
func (p *MapAny) Clone() *MapAny
```

**示例：**
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

## 迭代

### Range()

遍歷所有鍵值對。

```go
func (p *MapAny) Range(f func(key, value interface{}) bool)
```

**示例：**
```go
m.Range(func(key, value interface{}) bool {
    fmt.Printf("%s: %v\n", key, value)
    return true  // 繼續迭代
})
```

---

## 使用模式

### 配置管理

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

### 嵌套資料訪問

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

### 類型轉換

```go
func processData(data *anyx.MapAny) {
    name := data.GetString("name")
    age := data.GetInt("age")
    active := data.GetBool("active")
    price := data.GetFloat64("price")
    tags := data.GetStringSlice("tags")
    
    fmt.Printf("名稱: %s, 年齡: %d, 活躍: %v\n", name, age, active)
    fmt.Printf("價格: %.2f, 標籤: %v\n", price, tags)
}
```

### 線程安全操作

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

## 最佳實踐

### 錯誤處理

```go
// 好：處理未找到錯誤
value, err := m.Get("key")
if err != nil {
    if err == anyx.ErrNotFound {
        fmt.Println("鍵未找到")
    } else {
        fmt.Printf("錯誤: %v\n", err)
    }
}

// 好：在 Get 之前使用 Exists()
if m.Exists("key") {
    value, _ := m.Get("key")
    fmt.Printf("值: %v\n", value)
}
```

### 類型安全

```go
// 好：使用類型安全獲取器
age := m.GetInt("age")
name := m.GetString("name")
active := m.GetBool("active")

// 避免：手動類型斷言
value, _ := m.Get("age")
age, ok := value.(int)
if !ok {
    // 處理類型不匹配
}
```

### 嵌套訪問

```go
// 好：啟用 cut 進行嵌套訪問
m.EnableCut(".")
value := m.GetString("user.profile.name")

// 替代方案：鏈式調用 GetMap()
user := m.GetMap("user")
profile := user.GetMap("profile")
name := profile.GetString("name")
```

---

## 相關文檔

- [candy](/zh-TW/modules/candy) - 類型轉換
- [json](/zh-TW/modules/json) - JSON 處理
- [validator](/zh-TW/modules/validator) - 資料驗證
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
