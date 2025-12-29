---
title: anyx - Interface{} 助手
---

# anyx - Interface{} 助手

## 概述

anyx 模块提供对 `interface{}` 值的类型安全操作，具有访问和转换数据的便捷方法。它包装 `sync.Map` 以实现线程安全操作。

## 核心类型

### MapAny

具有类型安全访问器的线程安全映射。

```go
type MapAny struct {
    data *sync.Map
    cut  *atomic.Bool
    seq  *atomic.String
}
```

---

## 构造函数

### NewMap()

从映射创建新的 MapAny。

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

从 JSON 字节创建新的 MapAny。

```go
func NewMapWithJson(s []byte) (*MapAny, error)
```

**示例：**
```go
data := []byte(`{"name":"John","age":30}`)
m, err := anyx.NewMapWithJson(data)
if err != nil {
    log.Errorf("解析 JSON 失败: %v", err)
}
```

---

### NewMapWithYaml()

从 YAML 字节创建新的 MapAny。

```go
func NewMapWithYaml(s []byte) (*MapAny, error)
```

**示例：**
```go
data := []byte("name: John\nage: 30")
m, err := anyx.NewMapWithYaml(data)
if err != nil {
    log.Errorf("解析 YAML 失败: %v", err)
}
```

---

### NewMapWithAny()

从任何值创建新的 MapAny。

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

启用带分隔符的嵌套键访问。

```go
func (p *MapAny) EnableCut(seq string) *MapAny
```

**参数：**
- `seq` - 分隔符字符串（例如 "."、"/"）

**返回值：**
- MapAny 实例用于链式调用

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

禁用嵌套键访问。

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

设置键值对。

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

按键获取值。

```go
func (p *MapAny) Get(key string) (interface{}, error)
```

**返回值：**
- 如果键存在，返回值
- 如果键不存在，返回 `ErrNotFound`

**示例：**
```go
value, err := m.Get("name")
if err != nil {
    log.Errorf("键未找到: %v", err)
} else {
    fmt.Printf("名称: %v\n", value)
}
```

---

### Exists()

检查键是否存在。

```go
func (p *MapAny) Exists(key string) bool
```

**示例：**
```go
if m.Exists("name") {
    fmt.Println("名称存在")
}
```

---

## 类型安全获取器

### GetBool()

获取布尔值。

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

获取 int 值。

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

获取 int32 或 int64 值。

```go
func (p *MapAny) GetInt32(key string) int32
func (p *MapAny) GetInt64(key string) int64
```

---

### GetUint16(), GetUint32(), GetUint64()

获取无符号整数值。

```go
func (p *MapAny) GetUint16(key string) uint16
func (p *MapAny) GetUint32(key string) uint32
func (p *MapAny) GetUint64(key string) uint64
```

---

### GetFloat64()

获取 float64 值。

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

获取字符串值。

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

获取 []byte 值。

```go
func (p *MapAny) GetBytes(key string) []byte
```

---

### GetMap()

获取嵌套的 MapAny。

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

获取 []interface{} 值。

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

获取 []string 值。

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

获取 []int64 值。

```go
func (p *MapAny) GetInt64Slice(key string) []int64
```

---

### GetUint32Slice()

获取 []uint32 值。

```go
func (p *MapAny) GetUint32Slice(key string) []uint32
```

---

### GetUint64Slice()

获取 []uint64 值。

```go
func (p *MapAny) GetUint64Slice(key string) []uint64
```

---

## 转换方法

### ToMap()

转换为标准 map[string]interface{}。

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

转换为 *sync.Map。

```go
func (p *MapAny) ToSyncMap() *sync.Map
```

---

### Clone()

创建 MapAny 的副本。

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

遍历所有键值对。

```go
func (p *MapAny) Range(f func(key, value interface{}) bool)
```

**示例：**
```go
m.Range(func(key, value interface{}) bool {
    fmt.Printf("%s: %v\n", key, value)
    return true  // 继续迭代
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

### 嵌套数据访问

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

### 类型转换

```go
func processData(data *anyx.MapAny) {
    name := data.GetString("name")
    age := data.GetInt("age")
    active := data.GetBool("active")
    price := data.GetFloat64("price")
    tags := data.GetStringSlice("tags")
    
    fmt.Printf("名称: %s, 年龄: %d, 活跃: %v\n", name, age, active)
    fmt.Printf("价格: %.2f, 标签: %v\n", price, tags)
}
```

### 线程安全操作

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

## 最佳实践

### 错误处理

```go
// 好：处理未找到错误
value, err := m.Get("key")
if err != nil {
    if err == anyx.ErrNotFound {
        fmt.Println("键未找到")
    } else {
        fmt.Printf("错误: %v\n", err)
    }
}

// 好：在 Get 之前使用 Exists()
if m.Exists("key") {
    value, _ := m.Get("key")
    fmt.Printf("值: %v\n", value)
}
```

### 类型安全

```go
// 好：使用类型安全获取器
age := m.GetInt("age")
name := m.GetString("name")
active := m.GetBool("active")

// 避免：手动类型断言
value, _ := m.Get("age")
age, ok := value.(int)
if !ok {
    // 处理类型不匹配
}
```

### 嵌套访问

```go
// 好：启用 cut 进行嵌套访问
m.EnableCut(".")
value := m.GetString("user.profile.name")

// 替代方案：链式调用 GetMap()
user := m.GetMap("user")
profile := user.GetMap("profile")
name := profile.GetString("name")
```

---

## 相关文档

- [candy](/zh-CN/modules/candy) - 类型转换
- [json](/zh-CN/modules/json) - JSON 处理
- [validator](/zh-CN/modules/validator) - 数据验证
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
