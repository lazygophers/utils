---
title: defaults - 默认值
---

# defaults - 默认值

## 概述

defaults 模块为 Go 结构体提供默认值设置,支持各种数据类型和自定义默认函数。

## 函数

### SetDefaults()

为结构体设置默认值。

```go
func SetDefaults(value interface{}, opts ...Options) error
```

**参数:**
- `value` - 要设置默认值的结构体指针
- `opts` - 默认行为选项

**选项:**
- `ErrorMode` - 错误处理模式(Panic, Ignore, Return)
- `ValidateDefaults` - 是否验证默认值
- `AllowOverwrite` - 是否允许覆盖非零值

**示例:**
```go
type Config struct {
    Host     string `default:"localhost"`
    Port     int    `default:"8080"`
    Debug    bool   `default:"false"`
    Timeout  int    `default:"30"`
}

var cfg Config
defaults.SetDefaults(&cfg)
// cfg.Host == "localhost"
// cfg.Port == 8080
// cfg.Debug == false
// cfg.Timeout == 30
```

---

### RegisterCustomDefault()

为类型注册自定义默认函数。

```go
func RegisterCustomDefault(typeName string, fn DefaultFunc)
```

**参数:**
- `typeName` - 注册的类型名称
- `fn` - 默认函数

**示例:**
```go
defaults.RegisterCustomDefault("time.Time", func() interface{} {
    return time.Now()
})
```

---

## 使用模式

### 结构体初始化

```go
type User struct {
    Name     string `default:"Anonymous"`
    Email     string `default:""`
    Age       int    `default:"0"`
    Active    bool   `default:"true"`
    CreatedAt time.Time `default:"now"`
}

var user User
defaults.SetDefaults(&user)
// user.Name == "Anonymous"
// user.Email == ""
// user.Age == 0
// user.Active == true
// user.CreatedAt == time.Now()
```

### 自定义默认值

```go
type Product struct {
    Name     string
    Price     float64
    Stock     int
    Category  string
}

defaults.RegisterCustomDefault("Product.Price", func() interface{} {
    return 0.0
})

defaults.RegisterCustomDefault("Product.Stock", func() interface{} {
    return 100
})

var product Product
defaults.SetDefaults(&product)
// product.Price == 0.0
// product.Stock == 100
```

### 错误处理

```go
// 好的做法: 错误时 panic
defaults.SetDefaults(&cfg, defaults.Options{
    ErrorMode: defaults.ErrorModePanic,
})

// 好的做法: 忽略错误
defaults.SetDefaults(&cfg, defaults.Options{
    ErrorMode: defaults.ErrorModeIgnore,
})

// 好的做法: 返回错误
defaults.SetDefaults(&cfg, defaults.Options{
    ErrorMode: defaults.ErrorModeReturn,
})
```

---

## 最佳实践

### 默认标签

```go
// 好的做法: 使用适当的默认值
type Config struct {
    Host     string `default:"localhost"`  // 好的默认值
    Port     int    `default:"8080"`   // 好的默认值
    Debug    bool   `default:"false"`   // 好的默认值
    Timeout  int    `default:"30"`     // 好的默认值
}

// 避免: 不必要的默认值
type Config struct {
    Host     string `default:""`         // 不好的默认值
    Port     int    `default:"0"`         // 不好的默认值
    Debug    bool   `default:"false"`   // 好的默认值
}
```

---

## 相关文档

- [validator](/zh-CN/modules/validator) - 数据验证
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
