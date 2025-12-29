---
title: defaults - 默認值
---

# defaults - 默認值

## 概述

defaults 模組為 Go 結構體提供默認值設置,支持各種數據類型和自定義默認函數。

## 函數

### SetDefaults()

為結構體設置默認值。

```go
func SetDefaults(value interface{}, opts ...Options) error
```

**參數:**
- `value` - 要設置默認值的結構體指針
- `opts` - 默認行為選項

**選項:**
- `ErrorMode` - 錯誤處理模式(Panic, Ignore, Return)
- `ValidateDefaults` - 是否驗證默認值
- `AllowOverwrite` - 是否允許覆蓋非零值

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

為類型註冊自定義默認函數。

```go
func RegisterCustomDefault(typeName string, fn DefaultFunc)
```

**參數:**
- `typeName` - 註冊的類型名稱
- `fn` - 默認函數

**示例:**
```go
defaults.RegisterCustomDefault("time.Time", func() interface{} {
    return time.Now()
})
```

---

## 使用模式

### 結構體初始化

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

### 自定義默認值

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

### 錯誤處理

```go
// 好的做法: 錯誤時 panic
defaults.SetDefaults(&cfg, defaults.Options{
    ErrorMode: defaults.ErrorModePanic,
})

// 好的做法: 忽略錯誤
defaults.SetDefaults(&cfg, defaults.Options{
    ErrorMode: defaults.ErrorModeIgnore,
})

// 好的做法: 返回錯誤
defaults.SetDefaults(&cfg, defaults.Options{
    ErrorMode: defaults.ErrorModeReturn,
})
```

---

## 最佳實踐

### 默認標籤

```go
// 好的做法: 使用適當的默認值
type Config struct {
    Host     string `default:"localhost"`  // 好的默認值
    Port     int    `default:"8080"`   // 好的默認值
    Debug    bool   `default:"false"`   // 好的默認值
    Timeout  int    `default:"30"`     // 好的默認值
}

// 避免: 不必要的默認值
type Config struct {
    Host     string `default:""`         // 不好的默認值
    Port     int    `default:"0"`         // 不好的默認值
    Debug    bool   `default:"false"`   // 好的默認值
}
```

---

## 相關文檔

- [validator](/zh-TW/modules/validator) - 數據驗證
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
