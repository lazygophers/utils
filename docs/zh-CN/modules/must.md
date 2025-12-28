---
title: must.go - 错误断言
---

# must.go - 错误断言

## 概述

`must.go` 提供错误断言工具，通过基于 panic 的错误检查简化错误处理流程。这些函数专为初始化和关键操作设计，当失败时应终止程序执行。

## 函数

### MustOk()

断言第二个返回值为 true。

```go
func MustOk[T any](value T, ok bool) T
```

**参数：**
- `value` - 返回值
- `ok` - 成功标志

**返回值：**
- 如果 `ok` 为 true，返回 `value`
- 如果 `ok` 为 false，panic 并显示消息 "is not ok"

**示例：**
```go
value, ok := getValue()
result := utils.MustOk(value, ok)
```

**注意：**
- 当第二个返回值表示成功/失败时使用此函数
- 如果断言失败，panic 会终止程序执行
- 泛型类型 `T` 允许任何返回类型
- 常用于 map 查找和类型断言

---

### MustSuccess()

断言错误为 nil。

```go
func MustSuccess(err error)
```

**参数：**
- `err` - 要检查的错误

**行为：**
- 如果 `err` 为 nil，不执行任何操作
- 如果 `err` 不为 nil，panic 并显示格式化的错误消息

**示例：**
```go
utils.MustSuccess(config.Load(&cfg, "config.json"))
utils.MustSuccess(os.MkdirAll("data", 0755))
utils.MustSuccess(db.Ping())
```

**注意：**
- 常用于初始化和设置操作
- Panic 消息包含错误详情以便调试
- 用于程序必须成功才能继续的操作

---

### Must()

组合验证函数，检查错误状态并返回值。

```go
func Must[T any](value T, err error) T
```

**参数：**
- `value` - 返回值
- `err` - 错误

**返回值：**
- 如果 `err` 为 nil，返回 `value`
- 如果 `err` 不为 nil，panic

**示例：**
```go
data := utils.Must(loadData())
file := utils.Must(os.Open("file.txt"))
result := utils.Must(http.Get(url))
conn := utils.Must(net.Listen("tcp", ":8080"))
```

**注意：**
- must 模块中最常用的函数
- 组合错误检查和值提取
- 泛型类型 `T` 允许任何返回类型
- 适用于返回 `(T, error)` 的函数

---

### Ignore()

强制忽略任何参数。

```go
func Ignore[T any](value T, _ any) T
```

**参数：**
- `value` - 要返回的值
- `_` - 被忽略的参数

**返回值：**
- 返回 `value`

**示例：**
```go
result := utils.Ignore(data, err)
```

**注意：**
- 当需要抑制 linter 关于忽略值的警告时使用
- 第二个参数被显式忽略
- 有助于保持代码整洁而不出现 linter 警告
- 不实际处理错误，只是抑制警告

---

## 使用模式

### 初始化链

```go
func initApp() {
    cfg := utils.Must(loadConfig())
    db := utils.Must(connectDB(cfg.DatabaseURL))
    server := utils.Must(createServer(cfg.Port))
    
    utils.MustSuccess(server.Start())
}
```

### 文件操作

```go
func readFile(path string) []byte {
    file := utils.Must(os.Open(path))
    defer file.Close()
    
    data := utils.Must(io.ReadAll(file))
    return data
}
```

### Map 操作

```go
func getValue(m map[string]int, key string) int {
    value, ok := m[key]
    return utils.MustOk(value, ok)
}
```

### 配置加载

```go
type Config struct {
    Host string `json:"host"`
    Port int    `json:"port"`
}

func loadConfig(path string) *Config {
    data := utils.Must(os.ReadFile(path))
    
    var cfg Config
    utils.MustSuccess(json.Unmarshal(data, &cfg))
    
    return &cfg
}
```

## 最佳实践

### 何时使用 Must 函数

**使用 `Must()` 当：**
- 操作对程序启动至关重要
- 失败应终止程序执行
- 错误恢复不可能或没有意义
- 你在初始化代码中（main、init）

**避免 `Must()` 当：**
- 处理用户输入
- 可能失败的网络请求
- 可能不存在的文件操作
- 任何可能合理失败的操作

### 错误处理 vs Panic

```go
// 好：使用 Must 进行初始化
func init() {
    config = utils.Must(loadConfig())
}

// 好：为用户操作处理错误
func handleUserRequest() error {
    data, err := loadData()
    if err != nil {
        return err
    }
    // 处理数据
    return nil
}
```

## 相关文档

- [orm.go](/zh-CN/modules/orm) - 数据库操作
- [validator](/zh-CN/modules/validator) - 数据验证
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
