# 用户指南

欢迎使用 LazyGophers Utils 用户指南！这里包含了帮助您快速上手和深入使用整个工具库的完整文档。

## 📚 指南概览

LazyGophers Utils 是一个功能全面的 Go 工具库，包含 20+ 个专业模块，覆盖日常开发的各种需求。

### 🎯 快速导航

| 模块类型 | 模块数量 | 主要功能 | 文档状态 |
|---------|----------|----------|----------|
| **基础工具** | 5+ | 错误处理、验证、数据库操作 | ✅ 完整 |
| **数据处理** | 6+ | 类型转换、字符串处理、JSON | ✅ 完整 |
| **系统工具** | 4+ | 时间处理、网络、运行时 | ✅ 完整 |
| **开发工具** | 3+ | 配置管理、测试、调试 | ✅ 完整 |
| **高级功能** | 3+ | 加密、事件、并发控制 | ✅ 完整 |

## 🚀 快速开始

### 安装

```bash
go get github.com/lazygophers/utils
```

### 基础使用

```go
package main

import (
    "github.com/lazygophers/utils"
    "github.com/lazygophers/utils/candy"
    "github.com/lazygophers/utils/xtime"
)

func main() {
    // 错误处理工具
    result := utils.Must(someFunction())
    
    // 类型转换工具
    str := candy.ToString(123)
    
    // 时间处理工具
    cal := xtime.NowCalendar()
    fmt.Println(cal.String())
}
```

## 📖 模块指南

### 🔧 核心工具模块

#### must.go - 错误处理工具
```go
// 简化错误处理
value := utils.Must(getValue())         // 出错时panic
utils.MustSuccess(doSomething())        // 验证操作成功
result := utils.MustOk(checkStatus())   // 验证状态正确
```

#### validate.go - 数据验证
```go
type User struct {
    Name  string `validate:"required"`
    Email string `validate:"required,email"`
}

// 快速验证
err := utils.Validate(&user)
```

#### orm.go - 数据库操作
```go
// 数据库扫描和序列化
err := utils.Scan(dbData, &result)
value, err := utils.Value(&data)
```

### 🍭 Candy - 数据处理工具

**主要功能**：
- 类型转换：支持所有基础类型的相互转换
- 切片操作：过滤、映射、去重、排序等
- 数组处理：合并、分割、查找等

```go
import "github.com/lazygophers/utils/candy"

// 类型转换
str := candy.ToString(123)           // "123"
num := candy.ToInt("456")            // 456
slice := candy.ToSlice(data)         // []interface{}

// 切片操作
filtered := candy.Filter(slice, func(v interface{}) bool {
    return v.(int) > 10
})
mapped := candy.Map(slice, func(v interface{}) interface{} {
    return v.(int) * 2
})
```

### 🕰️ XTime - 时间处理工具

**核心特性**：
- 中国农历支持
- 24节气计算
- 生肖天干地支
- 传统节日检测

```go
import "github.com/lazygophers/utils/xtime"

// 创建日历对象
cal := xtime.NowCalendar()

// 获取各种信息
fmt.Println(cal.String())              // 2023年08月15日 六月廿九 兔年 处暑
fmt.Println(cal.LunarDate())           // 农历二零二三年六月廿九
fmt.Println(cal.Animal())              // 兔
fmt.Println(cal.CurrentSolarTerm())    // 处暑

// 详细信息获取
data := cal.ToMap() // 完整的JSON格式数据
```

### 🔐 Cryptox - 加密工具

**支持算法**：
- 对称加密：AES, DES, Blowfish, ChaCha20
- 非对称加密：RSA, ECC
- 哈希算法：MD5, SHA系列, Blake2
- 消息认证：HMAC

```go
import "github.com/lazygophers/utils/cryptox"

// AES加密
encrypted, err := cryptox.AESEncrypt(data, key)
decrypted, err := cryptox.AESDecrypt(encrypted, key)

// RSA加密
publicKey, privateKey, err := cryptox.GenerateRSAKeyPair(2048)
encrypted, err := cryptox.RSAEncrypt(data, publicKey)
```

### 🌐 Network - 网络工具

**主要功能**：
- IP地址操作和验证
- 网络接口信息获取
- Fiber框架增强

```go
import "github.com/lazygophers/utils/network"

// IP操作
isValid := network.IsValidIP("192.168.1.1")
isPrivate := network.IsPrivateIP("10.0.0.1")

// 获取网络接口信息
interfaces, err := network.GetNetworkInterfaces()
```

### ⚙️ Config - 配置管理

**支持格式**：
- JSON, YAML, TOML
- 环境变量
- 命令行参数

```go
import "github.com/lazygophers/utils/config"

type AppConfig struct {
    Port     int    `json:"port"`
    Database string `json:"database"`
}

var cfg AppConfig
err := config.Load("config.json", &cfg)
```

### 🔄 Routine - 并发控制

**功能特性**：
- Goroutine池管理
- 任务队列
- 缓存机制

```go
import "github.com/lazygophers/utils/routine"

// 使用goroutine池
pool := routine.NewPool(10)
pool.Submit(func() {
    // 任务执行
})
```

## 📊 模块对比和选择

### 按使用场景选择模块

#### 🔄 数据处理场景
| 需求 | 推荐模块 | 主要功能 |
|------|----------|----------|
| 类型转换 | `candy` | 安全的类型转换 |
| 字符串处理 | `stringx` | 增强的字符串操作 |
| JSON操作 | `json` | 高性能JSON处理 |
| 数据验证 | `validate` | 结构体验证 |

#### 🕰️ 时间处理场景
| 需求 | 推荐模块 | 主要功能 |
|------|----------|----------|
| 基础时间操作 | `xtime` | 增强的时间处理 |
| 农历节气 | `xtime` | 中国传统历法 |
| 工作时间计算 | `xtime007/955/996` | 工作制时间常量 |

#### 🔐 安全处理场景
| 需求 | 推荐模块 | 主要功能 |
|------|----------|----------|
| 数据加密 | `cryptox` | 各种加密算法 |
| 哈希计算 | `cryptox` | 哈希和消息认证 |
| 随机数生成 | `randx` | 安全随机数 |

#### 🌐 网络处理场景
| 需求 | 推荐模块 | 主要功能 |
|------|----------|----------|
| IP地址处理 | `network` | IP操作和验证 |
| 网络接口 | `network` | 接口信息获取 |
| Web框架增强 | `network` | Fiber增强功能 |

## 🎯 使用最佳实践

### 1. 错误处理最佳实践

```go
// ✅ 推荐：使用Must系列函数简化错误处理
func processData() {
    // 对于不应该出错的操作
    data := utils.Must(getData())
    
    // 对于需要验证成功的操作
    utils.MustSuccess(saveData(data))
    
    // 对于需要验证状态的操作
    status := utils.MustOk(checkStatus())
}

// ✅ 推荐：结合日志记录
func safeProcess() {
    result, err := processData()
    if err != nil {
        log.Error("Processing failed", log.Error(err))
        return
    }
    // 继续处理
}
```

### 2. 类型转换最佳实践

```go
// ✅ 推荐：使用candy进行安全转换
import "github.com/lazygophers/utils/candy"

func convertData(input interface{}) {
    // 安全的类型转换，有默认值
    str := candy.ToString(input)
    num := candy.ToInt(input)
    
    // 检查转换是否成功
    if str == "" {
        log.Warn("Failed to convert to string")
    }
}
```

### 3. 配置管理最佳实践

```go
// ✅ 推荐：统一配置结构
type Config struct {
    Server   ServerConfig   `json:"server"`
    Database DatabaseConfig `json:"database"`
    Redis    RedisConfig    `json:"redis"`
}

func loadConfig() *Config {
    var cfg Config
    
    // 加载配置文件
    err := config.Load("config.json", &cfg)
    if err != nil {
        log.Fatal("Failed to load config", log.Error(err))
    }
    
    // 验证配置
    if err := utils.Validate(&cfg); err != nil {
        log.Fatal("Invalid config", log.Error(err))
    }
    
    return &cfg
}
```

### 4. 性能优化最佳实践

```go
// ✅ 推荐：重用对象减少分配
var calendarPool = sync.Pool{
    New: func() interface{} {
        return &xtime.Calendar{}
    },
}

func getCalendarInfo(t time.Time) string {
    cal := calendarPool.Get().(*xtime.Calendar)
    defer calendarPool.Put(cal)
    
    // 使用calendar对象
    return cal.String()
}
```

## 🔗 深入学习

### 📚 进阶指南
- **[模块详细文档](../modules/)** - 每个模块的详细使用指南
- **[API参考](../api/)** - 完整的API文档
- **[性能指南](../performance/)** - 性能优化和基准测试

### 🛠️ 开发相关
- **[贡献指南](../development/contributing.md)** - 如何参与项目开发
- **[测试文档](../testing/)** - 测试策略和质量保证

### 💡 实用资源
- **[最佳实践](../guides/best-practices.md)** - 生产环境使用建议
- **[故障排除](../guides/troubleshooting.md)** - 常见问题解答
- **[示例代码](../guides/examples.md)** - 丰富的实际示例

## 💬 获取帮助

### 常见问题
1. **Q: 如何选择合适的模块？**
   A: 查看上方的模块对比表，根据具体需求选择

2. **Q: 模块之间有依赖关系吗？**
   A: 大部分模块独立设计，可单独使用

3. **Q: 如何处理版本兼容性？**
   A: 项目遵循语义版本，向后兼容

### 获取支持
- 📖 **查看文档**: 优先查阅相关模块文档
- 🔍 **搜索Issue**: 在GitHub搜索已知问题
- 💬 **社区讨论**: GitHub Discussions
- 🐛 **报告问题**: 创建详细的Issue

---

*用户指南最后更新: 2025年09月13日*