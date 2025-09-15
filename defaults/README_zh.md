# Defaults - 结构体默认值处理

一个强大的 Go 包，通过反射提供全面的结构体字段默认值分配。`defaults` 包基于结构体标签、自定义函数和类型特定逻辑自动填充结构体字段的默认值。

## 特性

- **全面类型支持**: 处理所有 Go 类型，包括原始类型、指针、结构体、切片、数组、映射、通道和函数
- **灵活配置**: 多种错误处理模式和自定义选项
- **结构体标签支持**: 使用 `default:"value"` 标签指定字段默认值
- **自定义默认函数**: 为特定类型注册自定义默认值生成器
- **时间处理**: 高级时间解析，支持多种格式和 "now" 关键字
- **嵌套结构支持**: 递归处理嵌套结构体和指针
- **JSON 集成**: 从 JSON 字符串解析复杂默认值
- **线程安全**: 通过适当同步实现并发安全
- **零依赖**: 仅使用 Go 标准库

## 安装

```bash
go get github.com/lazygophers/utils/defaults
```

## 快速开始

```go
package main

import (
    "fmt"
    "time"

    "github.com/lazygophers/utils/defaults"
)

type Config struct {
    Name        string        `default:"MyApp"`
    Port        int           `default:"8080"`
    Debug       bool          `default:"true"`
    Timeout     time.Duration `default:"30s"`
    CreatedAt   time.Time     `default:"now"`
    Tags        []string      `default:"[\"api\", \"web\"]"`
}

func main() {
    var config Config

    // 使用结构体标签设置默认值
    defaults.SetDefaults(&config)

    fmt.Printf("Config: %+v\n", config)
    // 输出: Config: {Name:MyApp Port:8080 Debug:true Timeout:30s CreatedAt:2024-01-01 12:00:00 Tags:[api web]}
}
```

## API 参考

### 核心函数

#### `SetDefaults(value interface{})`

使用结构体标签设置默认值，错误时恐慌（默认行为）。

**参数:**
- `value interface{}`: 指向要填充的结构体的指针

**示例:**
```go
type User struct {
    Name string `default:"Anonymous"`
    Age  int    `default:"18"`
}

var user User
defaults.SetDefaults(&user)
```

#### `SetDefaultsWithOptions(value interface{}, opts *Options) error`

使用自定义配置选项设置默认值。

**参数:**
- `value interface{}`: 指向要填充的结构体的指针
- `opts *Options`: 配置选项

**返回值:**
- `error`: 如果发生错误（取决于错误模式）

### 配置选项

#### `Options` 结构体

```go
type Options struct {
    ErrorMode        ErrorMode                // 错误处理策略
    CustomDefaults   map[string]DefaultFunc   // 自定义默认函数
    ValidateDefaults bool                     // 是否验证默认值
    AllowOverwrite   bool                     // 允许覆盖非零值
}
```

#### `ErrorMode` 常量

- `ErrorModePanic`: 错误时恐慌（默认）
- `ErrorModeIgnore`: 忽略错误并继续
- `ErrorModeReturn`: 返回错误而不恐慌

#### 自定义默认函数

#### `RegisterCustomDefault(typeName string, fn DefaultFunc)`

为特定类型注册自定义默认值函数。

**参数:**
- `typeName string`: 类型标识符（"string", "int", "float", "bool", "uint", "func"）
- `fn DefaultFunc`: 返回默认值的函数

**示例:**
```go
// 注册自定义字符串默认值
defaults.RegisterCustomDefault("string", func() interface{} {
    return "custom-default-" + time.Now().Format("20060102")
})
```

#### `ClearCustomDefaults()`

清除所有已注册的自定义默认函数。

## 支持的类型和标签

### 原始类型

#### 字符串
```go
type Example struct {
    Name     string `default:"John Doe"`
    Empty    string `default:""`
    Optional string // 无默认值，保持空
}
```

#### 整数类型
```go
type Example struct {
    Age     int   `default:"25"`
    Count   int64 `default:"1000"`
    Retries uint  `default:"3"`
}
```

#### 浮点类型
```go
type Example struct {
    Price  float64 `default:"99.99"`
    Rating float32 `default:"4.5"`
}
```

#### 布尔值
```go
type Example struct {
    Enabled  bool `default:"true"`
    Disabled bool `default:"false"`
}
```

### 复杂类型

#### 时间
```go
type Example struct {
    CreatedAt time.Time `default:"now"`
    UpdatedAt time.Time `default:"2024-01-01 15:04:05"`
    Birthday  time.Time `default:"1990-01-01"`
}
```

支持的时间格式:
- `"now"` - 当前时间
- RFC3339: `"2006-01-02T15:04:05Z07:00"`
- RFC3339Nano: `"2006-01-02T15:04:05.999999999Z07:00"`
- 日期时间: `"2006-01-02 15:04:05"`
- 仅日期: `"2006-01-02"`
- 仅时间: `"15:04:05"`

#### 指针
```go
type Example struct {
    Name *string `default:"John"`
    Age  *int    `default:"30"`
}
```

#### 切片
```go
type Example struct {
    Tags     []string `default:"[\"tag1\", \"tag2\"]"`
    Numbers  []int    `default:"1,2,3,4,5"`
    Empty    []string // 初始化为空切片
}
```

#### 数组
```go
type Example struct {
    Colors [3]string `default:"red,green,blue"`
    Matrix [2]int    `default:"10,20"`
}
```

#### 映射
```go
type Example struct {
    Config   map[string]string `default:"{\"key1\":\"value1\", \"key2\":\"value2\"}"`
    Settings map[string]int    // 初始化为空映射
}
```

#### 通道
```go
type Example struct {
    Messages chan string `default:"10"` // 缓冲区大小
    Events   chan int    `default:"0"`  // 无缓冲
}
```

#### 嵌套结构体
```go
type Address struct {
    Street string `default:"123 Main St"`
    City   string `default:"Springfield"`
}

type Person struct {
    Name    string  `default:"John"`
    Address Address // 自动处理
}
```

## 使用示例

### 基本配置

```go
package main

import (
    "fmt"
    "time"

    "github.com/lazygophers/utils/defaults"
)

type DatabaseConfig struct {
    Host         string        `default:"localhost"`
    Port         int           `default:"5432"`
    Username     string        `default:"admin"`
    Password     string        // 出于安全考虑无默认值
    MaxConns     int           `default:"10"`
    Timeout      time.Duration `default:"30s"`
    SSL          bool          `default:"true"`
    RetryAttempts uint          `default:"3"`
}

func main() {
    var dbConfig DatabaseConfig
    defaults.SetDefaults(&dbConfig)

    fmt.Printf("数据库配置:\n")
    fmt.Printf("  Host: %s\n", dbConfig.Host)
    fmt.Printf("  Port: %d\n", dbConfig.Port)
    fmt.Printf("  SSL: %t\n", dbConfig.SSL)
    fmt.Printf("  Timeout: %v\n", dbConfig.Timeout)
}
```

### 错误处理选项

```go
package main

import (
    "fmt"
    "log"

    "github.com/lazygophers/utils/defaults"
)

type Config struct {
    Port    int    `default:"invalid"`  // 无效整数
    Timeout string `default:"30s"`
}

func main() {
    var config Config

    // 选项 1: 返回错误而不是恐慌
    opts := &defaults.Options{
        ErrorMode: defaults.ErrorModeReturn,
    }

    if err := defaults.SetDefaultsWithOptions(&config, opts); err != nil {
        log.Printf("设置默认值错误: %v", err)
    }

    // 选项 2: 忽略错误并继续
    opts.ErrorMode = defaults.ErrorModeIgnore
    defaults.SetDefaultsWithOptions(&config, opts)

    fmt.Printf("Config: %+v\n", config)
}
```

### 自定义默认函数

```go
package main

import (
    "fmt"
    "os"
    "time"

    "github.com/lazygophers/utils/defaults"
)

type AppConfig struct {
    AppName     string `default:""`  // 将使用自定义默认值
    Environment string `default:""`  // 将使用自定义默认值
    Version     string `default:"1.0.0"`
}

func main() {
    // 注册自定义默认值
    defaults.RegisterCustomDefault("string", func() interface{} {
        if appName := os.Getenv("APP_NAME"); appName != "" {
            return appName
        }
        return "MyApplication"
    })

    var config AppConfig
    defaults.SetDefaults(&config)

    fmt.Printf("应用配置: %+v\n", config)
}
```

### 复杂嵌套结构

```go
package main

import (
    "fmt"
    "time"

    "github.com/lazygophers/utils/defaults"
)

type Server struct {
    Host string `default:"0.0.0.0"`
    Port int    `default:"8080"`
}

type Database struct {
    Host     string `default:"localhost"`
    Port     int    `default:"5432"`
    Username string `default:"admin"`
    Pool     *PoolConfig
}

type PoolConfig struct {
    MaxConnections int           `default:"10"`
    IdleTimeout    time.Duration `default:"5m"`
}

type ApplicationConfig struct {
    Name      string    `default:"MyApp"`
    Debug     bool      `default:"false"`
    CreatedAt time.Time `default:"now"`
    Server    Server
    Database  Database
    Features  []string `default:"[\"auth\", \"api\", \"web\"]"`
    Metadata  map[string]interface{} `default:"{\"version\":\"1.0\"}"`
}

func main() {
    var config ApplicationConfig
    defaults.SetDefaults(&config)

    fmt.Printf("应用程序: %s\n", config.Name)
    fmt.Printf("服务器: %s:%d\n", config.Server.Host, config.Server.Port)
    fmt.Printf("数据库: %s:%d\n", config.Database.Host, config.Database.Port)
    fmt.Printf("连接池最大连接数: %d\n", config.Database.Pool.MaxConnections)
    fmt.Printf("功能: %v\n", config.Features)
}
```

### 覆盖现有值

```go
package main

import (
    "fmt"

    "github.com/lazygophers/utils/defaults"
)

type Config struct {
    Name    string `default:"Default Name"`
    Port    int    `default:"8080"`
    Enabled bool   `default:"true"`
}

func main() {
    // 预填充一些字段
    config := Config{
        Name: "Custom Name",
        Port: 3000,
    }

    fmt.Printf("之前: %+v\n", config)

    // 选项 1: 不覆盖现有值（默认行为）
    defaults.SetDefaults(&config)
    fmt.Printf("之后（无覆盖）: %+v\n", config)

    // 选项 2: 允许覆盖现有值
    opts := &defaults.Options{
        AllowOverwrite: true,
    }
    defaults.SetDefaultsWithOptions(&config, opts)
    fmt.Printf("之后（覆盖）: %+v\n", config)
}
```

### 验证和自定义逻辑

```go
package main

import (
    "fmt"
    "strings"

    "github.com/lazygophers/utils/defaults"
)

type UserProfile struct {
    Username string `default:"guest"`
    Email    string `default:"user@example.com"`
    Role     string `default:"user"`
}

func main() {
    // 注册自定义验证/转换
    defaults.RegisterCustomDefault("string", func() interface{} {
        return strings.ToLower("DEFAULT_VALUE")
    })

    opts := &defaults.Options{
        ValidateDefaults: true,
        ErrorMode:        defaults.ErrorModeReturn,
    }

    var profile UserProfile
    if err := defaults.SetDefaultsWithOptions(&profile, opts); err != nil {
        fmt.Printf("验证错误: %v\n", err)
    }

    fmt.Printf("配置文件: %+v\n", profile)
}
```

## 高级功能

### 使用接口

```go
type Config struct {
    Data interface{} `default:"{\"key\": \"value\"}"`
    List interface{} `default:"[1, 2, 3]"`
}

var config Config
defaults.SetDefaults(&config)
// Data 将被解析为 JSON 对象/数组
```

### 通道缓冲区大小

```go
type EventSystem struct {
    Events    chan Event `default:"100"`  // 大小为 100 的缓冲通道
    Errors    chan error `default:"0"`    // 无缓冲通道
    Shutdown  chan bool  `default:"1"`    // 大小为 1 的缓冲通道
}
```

### 函数类型默认值

```go
type Handlers struct {
    OnError   func(error)   // 通过 RegisterCustomDefault 自定义默认值
    OnSuccess func(string)
}

// 注册自定义函数默认值
defaults.RegisterCustomDefault("func", func() interface{} {
    return func(err error) {
        log.Printf("默认错误处理器: %v", err)
    }
})
```

## 性能考虑

### 内存分配
- 原始类型的最小分配
- 切片和映射的高效处理
- 尽可能重用反射值

### 执行速度
- 缓存反射类型信息
- 优化在相似结构体上的重复使用
- 多个 goroutine 的并发安全

### 性能最佳实践

1. **重用选项对象:**
```go
var opts = &defaults.Options{
    ErrorMode: defaults.ErrorModeIgnore,
}

// 重用 opts 进行多次调用
defaults.SetDefaultsWithOptions(&config1, opts)
defaults.SetDefaultsWithOptions(&config2, opts)
```

2. **最小化自定义函数:**
自定义默认函数会增加开销 - 谨慎使用。

3. **优先使用简单标签:**
简单的字符串/数字默认值最快。

## 错误处理策略

### 恐慌模式（默认）
```go
defaults.SetDefaults(&config) // 错误时恐慌
```

### 错误返回模式
```go
opts := &defaults.Options{ErrorMode: defaults.ErrorModeReturn}
if err := defaults.SetDefaultsWithOptions(&config, opts); err != nil {
    log.Printf("错误: %v", err)
}
```

### 忽略模式
```go
opts := &defaults.Options{ErrorMode: defaults.ErrorModeIgnore}
defaults.SetDefaultsWithOptions(&config, opts) // 错误时继续
```

## 线程安全

defaults 包是线程安全的:

- **SetDefaults**: 并发使用安全
- **SetDefaultsWithOptions**: 并发使用安全
- **RegisterCustomDefault**: 并发注册安全
- **ClearCustomDefaults**: 并发访问安全

## 贡献

欢迎贡献！请确保:

1. 全面的测试覆盖
2. 正确的错误处理
3. 文档更新
4. 新功能的性能基准测试

## 许可证

此包是 LazyGophers Utils 库的一部分，遵循相同的许可条款。