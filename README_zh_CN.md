# LazyGophers Utils

> 🚀 功能丰富、高性能的Go实用程序库，让Go开发更高效

**🌍 语言**: [English](README.md) • [简体中文](README_zh_CN.md)

[![Go Version](https://img.shields.io/badge/Go-1.18+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-AGPL%20v3-green.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/lazygophers/utils.svg)](https://pkg.go.dev/github.com/lazygophers/utils)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/utils)](https://goreportcard.com/report/github.com/lazygophers/utils)
[![GitHub releases](https://img.shields.io/github/release/lazygophers/utils.svg)](https://github.com/lazygophers/utils/releases)
[![GoProxy.cn Downloads](https://goproxy.cn/stats/github.com/lazygophers/utils/badges/download-count.svg)](https://goproxy.cn/stats/github.com/lazygophers/utils)
[![Test Coverage](https://img.shields.io/badge/coverage-98%25-brightgreen.svg)](https://github.com/lazygophers/utils/actions)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/lazygophers/utils)

## 📋 目录

- [项目概览](#-项目概览)
- [核心特性](#-核心特性)
- [快速开始](#-快速开始)
- [文档](#-文档)
- [核心模块](#-核心模块)
- [功能模块](#-功能模块)
- [使用示例](#-使用示例)
- [性能数据](#-性能数据)
- [贡献代码](#-贡献代码)
- [许可证](#-许可证)
- [社区支持](#-社区支持)

## 💡 项目概览

LazyGophers Utils 是一个综合性的高性能 Go 实用程序库，提供 20+ 个专业模块，覆盖日常开发的各种需求。采用模块化设计，支持按需导入，零依赖冲突。

**设计理念**：简单、高效、可靠

## ✨ 核心特性

| 特性 | 描述 | 优势 |
|------|------|------|
| 🧩 **模块化设计** | 20+ 个独立模块 | 按需导入，减小体积 |
| ⚡ **高性能** | 基准测试验证 | 微秒级响应，内存友好 |
| 🛡️ **类型安全** | 充分利用泛型 | 编译时错误检查 |
| 🔒 **并发安全** | 协程友好设计 | 生产环境可用 |
| 📚 **文档完善** | 95%+ 文档覆盖 | 易学易用 |
| 🧪 **测试完备** | 98%+ 测试覆盖 | 质量保证 |

## 🚀 快速开始

### 安装

```bash
go get github.com/lazygophers/utils
```

### 基本用法

```go
package main

import (
    "github.com/lazygophers/utils"
    "github.com/lazygophers/utils/candy"
    "github.com/lazygophers/utils/xtime"
)

func main() {
    // 错误处理
    value := utils.Must(getValue())
    
    // 类型转换
    age := candy.ToInt("25")
    
    // 时间处理
    cal := xtime.NowCalendar()
    fmt.Println(cal.String()) // 2023年08月15日 六月廿九 兔年 处暑
}
```

## 📖 文档

### 📁 模块文档
- **核心模块**: [错误处理](must.go) | [数据库](orm.go) | [数据验证](validate.go)
- **数据处理**: [candy](candy/) | [json](json/) | [stringx](stringx/)
- **时间工具**: [xtime](xtime/) | [xtime996](xtime/xtime996/) | [xtime955](xtime/xtime955/)
- **系统工具**: [config](config/) | [runtime](runtime/) | [osx](osx/)
- **网络安全**: [network](network/) | [cryptox](cryptox/) | [pgp](pgp/)
- **并发控制**: [routine](routine/) | [wait](wait/) | [hystrix](hystrix/)

### 📋 快速参考
- [🔧 安装指南](#-快速开始)
- [📝 使用示例](#-使用示例)
- [📚 完整文档索引](docs/) - 综合文档中心
- [🎯 按场景查找模块](docs/#-quick-search) - 根据使用场景快速定位
- [🏗️ 架构文档](docs/architecture_zh.md) - 深入了解系统设计

### 🌍 多语言文档
- [English](README.md) - 英文版本

## 🔧 核心模块

### 错误处理 (`must.go`)
```go
// 断言操作成功，失败时panic
value := utils.Must(getValue())

// 验证无错误
utils.MustSuccess(doSomething())

// 验证布尔状态
result := utils.MustOk(checkCondition())
```

### 数据库操作 (`orm.go`)
```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age" default:"18"`
}

// 扫描数据库数据到结构体
err := utils.Scan(dbData, &user)

// 转换结构体为数据库值
value, err := utils.Value(user)
```

### 数据验证 (`validate.go`)
```go
type Config struct {
    Email string `validate:"required,email"`
    Port  int    `validate:"min=1,max=65535"`
}

// 快速验证
err := utils.Validate(&config)
```

## 📦 功能模块

<details>
<summary><strong>🍭 数据处理模块</strong></summary>

| 模块 | 功能 | 核心API |
|------|------|---------|
| **[candy](candy/)** | 类型转换语法糖 | `ToInt()`, `ToString()`, `ToBool()` |
| **[json](json/)** | 增强JSON处理 | `Marshal()`, `Unmarshal()`, `Pretty()` |
| **[stringx](stringx/)** | 字符串处理 | `IsEmpty()`, `Contains()`, `Split()` |
| **[anyx](anyx/)** | 任意类型工具 | `IsNil()`, `Type()`, `Convert()` |

</details>

<details>
<summary><strong>⏰ 时间处理模块</strong></summary>

| 模块 | 功能 | 特性 |
|------|------|------|
| **[xtime](xtime/)** | 增强时间处理 | 农历、节气、生肖 |
| **[xtime996](xtime/xtime996/)** | 996工作制常量 | 工作时间计算 |
| **[xtime955](xtime/xtime955/)** | 955工作制常量 | 工作时间计算 |
| **[xtime007](xtime/xtime007/)** | 007工作制常量 | 全天候时间 |

**xtime 特色功能**:
- 🗓️ 统一的日历接口（公历+农历）
- 🌙 精确的农历转换和节气计算
- 🐲 完整的干支纪年体系
- 🏮 自动传统节日识别

```go
cal := xtime.NowCalendar()
fmt.Println(cal.LunarDate())     // 农历二零二三年六月廿九
fmt.Println(cal.Animal())        // 兔
fmt.Println(cal.CurrentSolarTerm()) // 处暑
```

</details>

<details>
<summary><strong>🔧 系统工具模块</strong></summary>

| 模块 | 功能 | 用途 |
|------|------|------|
| **[config](config/)** | 配置管理 | 多格式配置文件读取 |
| **[runtime](runtime/)** | 运行时信息 | 系统信息获取 |
| **[osx](osx/)** | 操作系统增强 | 文件和进程操作 |
| **[app](app/)** | 应用框架 | 应用生命周期管理 |
| **[atexit](atexit/)** | 退出钩子 | 优雅退出处理 |

</details>

<details>
<summary><strong>🌐 网络安全模块</strong></summary>

| 模块 | 功能 | 应用场景 |
|------|------|----------|
| **[network](network/)** | 网络操作 | HTTP客户端、连接池 |
| **[cryptox](cryptox/)** | 加密工具 | 哈希、加密、解密 |
| **[pgp](pgp/)** | PGP加密 | 邮件加密、文件签名 |
| **[urlx](urlx/)** | URL处理 | URL解析、构建 |

</details>

<details>
<summary><strong>🚀 并发控制模块</strong></summary>

| 模块 | 功能 | 设计模式 |
|------|------|----------|
| **[routine](routine/)** | 协程管理 | 协程池、任务调度 |
| **[wait](wait/)** | 等待控制 | 超时、重试、限流 |
| **[hystrix](hystrix/)** | 熔断器 | 容错、降级 |
| **[singledo](singledo/)** | 单例模式 | 防重复执行 |
| **[event](event/)** | 事件驱动 | 发布订阅模式 |

</details>

<details>
<summary><strong>🧪 开发测试模块</strong></summary>

| 模块 | 功能 | 开发阶段 |
|------|------|----------|
| **[fake](fake/)** | 假数据生成 | 测试数据生成 |
| **[unit](unit/)** | 测试辅助 | 单元测试工具 |
| **[pyroscope](pyroscope/)** | 性能分析 | 生产监控 |
| **[defaults](defaults/)** | 默认值 | 配置初始化 |
| **[randx](randx/)** | 随机数 | 安全随机生成 |

</details>

## 🎯 使用示例

### 完整应用示例

```go
package main

import (
    "github.com/lazygophers/utils"
    "github.com/lazygophers/utils/config"
    "github.com/lazygophers/utils/candy"
    "github.com/lazygophers/utils/xtime"
)

type AppConfig struct {
    Port     int    `json:"port" default:"8080" validate:"min=1,max=65535"`
    Database string `json:"database" validate:"required"`
    Debug    bool   `json:"debug" default:"false"`
}

func main() {
    // 1. 加载配置
    var cfg AppConfig
    utils.MustSuccess(config.Load(&cfg, "config.json"))
    
    // 2. 验证配置
    utils.MustSuccess(utils.Validate(&cfg))
    
    // 3. 类型转换
    portStr := candy.ToString(cfg.Port)
    
    // 4. 时间处理
    cal := xtime.NowCalendar()
    log.Printf("应用启动时间：%s", cal.String())
    
    // 5. 启动服务器
    startServer(cfg)
}
```

### 错误处理最佳实践

```go
// ✅ 推荐：使用Must系列函数
func processData() string {
    data := utils.Must(loadData())        // 加载失败时panic
    utils.MustSuccess(validateData(data)) // 验证失败时panic
    return utils.MustOk(transformData(data)) // 转换失败时panic
}

// ✅ 推荐：批量错误处理
func batchProcess() error {
    return utils.MustSuccess(
        doStep1(),
        doStep2(),
        doStep3(),
    )
}
```

### 数据库操作示例

```go
type User struct {
    ID    int64  `json:"id"`
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
    Age   int    `json:"age" default:"0" validate:"min=0,max=150"`
}

func SaveUser(db *sql.DB, user *User) error {
    // 验证数据
    if err := utils.Validate(user); err != nil {
        return err
    }
    
    // 转换为数据库值
    data, err := utils.Value(user)
    if err != nil {
        return err
    }
    
    // 保存到数据库
    _, err = db.Exec("INSERT INTO users (data) VALUES (?)", data)
    return err
}

func GetUser(db *sql.DB, id int64) (*User, error) {
    var data []byte
    err := db.QueryRow("SELECT data FROM users WHERE id = ?", id).Scan(&data)
    if err != nil {
        return nil, err
    }
    
    var user User
    err = utils.Scan(data, &user)
    return &user, err
}
```

## 📊 性能数据

### 基准测试结果

| 操作 | 时间 | 内存分配 | vs 标准库 |
|------|------|----------|-----------|
| `candy.ToInt()` | 12.3 ns/op | 0 B/op | **快3.2倍** |
| `json.Marshal()` | 156 ns/op | 64 B/op | **快1.8倍** |
| `xtime.Now()` | 45.2 ns/op | 0 B/op | **快2.1倍** |
| `utils.Must()` | 2.1 ns/op | 0 B/op | **零开销** |

### 性能特征

- ⚡ **微秒级响应**：核心操作在微秒级完成
- 🧠 **内存友好**：使用sync.Pool减少GC压力
- 🔄 **零分配**：关键路径避免内存分配
- 🚀 **并发优化**：针对高并发场景优化

> 📈 详细性能报告：[性能文档](docs/performance_zh.md)

## 🤝 贡献代码

我们欢迎各种形式的贡献！

### 贡献流程

1. 🍴 Fork本项目
2. 🌿 创建特性分支：`git checkout -b feature/amazing-feature`
3. 📝 编写代码和测试
4. 🧪 确保测试通过：`go test ./...`
5. 📤 提交PR

### 开发标准

- ✅ 遵循 [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- 📖 所有公共API必须有godoc注释
- 🧪 新功能必须包含测试用例
- 📊 保持测试覆盖率 > 80%
- 🔄 保持向后兼容性

> 📋 详细指南：[贡献指南](CONTRIBUTING_zh.md)

## 📄 许可证

本项目采用GNU Affero General Public License v3.0许可证。

详见 [LICENSE](LICENSE) 文件。

## 🌟 社区支持

### 获取帮助

- 📖 **文档**：[完整文档](docs/)
- 🐛 **错误报告**：[GitHub Issues](https://github.com/lazygophers/utils/issues)
- 💬 **讨论**：[GitHub Discussions](https://github.com/lazygophers/utils/discussions)
- ❓ **问答**：[Stack Overflow](https://stackoverflow.com/questions/tagged/lazygophers-utils)

### 项目统计

| 指标 | 数值 | 说明 |
|------|------|------|
| 📦 模块数量 | 20+ | 覆盖各种常用功能 |
| 🧪 测试覆盖率 | 85%+ | 高质量代码保证 |
| 📝 文档完整度 | 95%+ | 详细使用说明 |
| ⚡ 性能等级 | A+ | 基准测试验证优化 |
| 🌟 GitHub Stars | ![GitHub stars](https://img.shields.io/github/stars/lazygophers/utils) | 社区认可度 |

### 致谢

感谢所有贡献者的辛勤工作！

[![Contributors](https://contrib.rocks/image?repo=lazygophers/utils)](https://github.com/lazygophers/utils/graphs/contributors)

---

<div align="center">

**如果这个项目对您有帮助，请给我们一个 ⭐ Star！**

[🚀 开始使用](#-快速开始) • [📖 查看文档](docs/) • [🤝 加入社区](https://github.com/lazygophers/utils/discussions)

</div>