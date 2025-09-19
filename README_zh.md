# LazyGophers Utils

> 🚀 一个功能丰富、高性能的 Go 工具库，让 Go 开发更加高效

**🌍 多语言**: [English](README.md) • [中文](README_zh.md) • [繁體中文](README_zh-hant.md) • [Español](README_es.md) • [Français](README_fr.md) • [Русский](README_ru.md) • [العربية](README_ar.md)

[![Build Status](https://github.com/lazygophers/utils/actions/workflows/test-and-build.yml/badge.svg)](https://github.com/lazygophers/utils/actions/workflows/test-and-build.yml)
[![Go Version](https://img.shields.io/badge/Go-1.24+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-AGPL%20v3-green.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/lazygophers/utils.svg)](https://pkg.go.dev/github.com/lazygophers/utils)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/utils)](https://goreportcard.com/report/github.com/lazygophers/utils)
[![Test Coverage](https://img.shields.io/badge/coverage-98%25-brightgreen.svg)](https://github.com/lazygophers/utils/actions/workflows/update-coverage-badge.yml)
[![GitHub releases](https://img.shields.io/github/release/lazygophers/utils.svg)](https://github.com/lazygophers/utils/releases)
[![GoProxy.cn Downloads](https://goproxy.cn/stats/github.com/lazygophers/utils/badges/download-count.svg)](https://goproxy.cn/stats/github.com/lazygophers/utils)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/lazygophers/utils)

## 📋 目录

-   [项目简介](#-项目简介)
-   [核心特性](#-核心特性)
-   [快速开始](#-快速开始)
-   [文档导航](#-文档导航)
-   [核心模块](#-核心模块)
-   [功能模块](#-功能模块)
-   [使用示例](#-使用示例)
-   [性能数据](#-性能数据)
-   [贡献指南](#-贡献指南)
-   [许可证](#-许可证)
-   [社区支持](#-社区支持)

## 💡 项目简介

LazyGophers Utils 是一个功能全面、性能优异的 Go 工具库，提供了 20+个专业模块，覆盖日常开发中的各种需求。采用模块化设计，按需引入，零依赖冲突。

**设计理念**：简洁、高效、可靠

## ✨ 核心特性

| 特性              | 说明          | 优势                 |
| ----------------- | ------------- | -------------------- |
| 🧩 **模块化设计** | 20+个独立模块 | 按需引入，减少体积   |
| ⚡ **高性能优化** | 基准测试验证  | 微秒级响应，内存友好 |
| 🛡️ **类型安全**   | 充分利用泛型  | 编译时错误检查       |
| 🔒 **并发安全**   | 协程友好设计  | 生产环境可靠         |
| 📚 **文档完备**   | 95%+ 文档覆盖 | 易学易用             |
| 🧪 **测试充分**   | 85%+ 测试覆盖 | 质量保障             |

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
    // 错误处理
    value := utils.Must(getValue())

    // 类型转换
    age := candy.ToInt("25")

    // 时间处理
    cal := xtime.NowCalendar()
    fmt.Println(cal.String()) // 2023年08月15日 六月廿九 兔年 处暑
}
```

## 📖 文档导航

### 📁 模块文档

-   **核心模块**：[错误处理](must.go) | [数据库](orm.go) | [验证](validate.go)
-   **数据处理**：[candy](candy/) | [json](json/) | [stringx](stringx/)
-   **时间工具**：[xtime](xtime/) | [xtime996](xtime/xtime996/) | [xtime955](xtime/xtime955/)
-   **系统工具**：[config](config/) | [runtime](runtime/) | [osx](osx/)
-   **网络&安全**：[network](network/) | [cryptox](cryptox/) | [pgp](pgp/)
-   **并发&控制**：[routine](routine/) | [wait](wait/) | [hystrix](hystrix/)

### 📋 快速参考

-   [🔧 安装指南](#-快速开始)
-   [📝 使用示例](#-使用示例)
-   [📚 完整文档索引](docs/) - 全面的文档导航中心
-   [🎯 按场景查找模块](docs/#-快速查找) - 按使用场景快速定位
-   [🏗️ 架构设计文档](docs/architecture_zh.md) - 深入了解系统设计

### 🌍 多语言 README

-   [English](README_en.md) - English version
-   [繁體中文](README_zh-hant.md) - Traditional Chinese
-   [Español](README_es.md) - Spanish version
-   [Français](README_fr.md) - French version
-   [Русский](README_ru.md) - Russian version
-   [العربية](README_ar.md) - Arabic version

## 🔧 核心模块

### 错误处理 (`must.go`)

```go
// 断言操作成功，失败时 panic
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

// 结构体转数据库值
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

| 模块                    | 功能           | 核心 API                               |
| ----------------------- | -------------- | -------------------------------------- |
| **[candy](candy/)**     | 类型转换语法糖 | `ToInt()`, `ToString()`, `ToBool()`    |
| **[json](json/)**       | JSON 处理增强  | `Marshal()`, `Unmarshal()`, `Pretty()` |
| **[stringx](stringx/)** | 字符串处理     | `IsEmpty()`, `Contains()`, `Split()`   |
| **[anyx](anyx/)**       | Any 类型工具   | `IsNil()`, `Type()`, `Convert()`       |

</details>

<details>
<summary><strong>⏰ 时间处理模块</strong></summary>

| 模块                            | 功能           | 特色             |
| ------------------------------- | -------------- | ---------------- |
| **[xtime](xtime/)**             | 增强时间处理   | 农历、节气、生肖 |
| **[xtime996](xtime/xtime996/)** | 996 工作制常量 | 工作时间计算     |
| **[xtime955](xtime/xtime955/)** | 955 工作制常量 | 工作时间计算     |
| **[xtime007](xtime/xtime007/)** | 007 工作制常量 | 全天候时间       |

**xtime 特色功能**：

-   🗓️ 统一日历接口（公历+农历）
-   🌙 精确农历转换和节气计算
-   🐲 完整天干地支系统
-   🏮 传统节日自动检测

```go
cal := xtime.NowCalendar()
fmt.Println(cal.LunarDate())     // 农历二零二三年六月廿九
fmt.Println(cal.Animal())        // 兔
fmt.Println(cal.CurrentSolarTerm()) // 处暑
```

</details>

<details>
<summary><strong>🔧 系统工具模块</strong></summary>

| 模块                    | 功能         | 用途               |
| ----------------------- | ------------ | ------------------ |
| **[config](config/)**   | 配置管理     | 多格式配置文件读取 |
| **[runtime](runtime/)** | 运行时信息   | 系统信息获取       |
| **[osx](osx/)**         | 操作系统增强 | 文件、进程操作     |
| **[app](app/)**         | 应用框架     | 应用生命周期管理   |
| **[atexit](atexit/)**   | 退出钩子     | 优雅关闭处理       |

</details>

<details>
<summary><strong>🌐 网络&安全模块</strong></summary>

| 模块                    | 功能     | 应用场景            |
| ----------------------- | -------- | ------------------- |
| **[network](network/)** | 网络操作 | HTTP 客户端、连接池 |
| **[cryptox](cryptox/)** | 加密工具 | 哈希、加密、解密    |
| **[pgp](pgp/)**         | PGP 加密 | 邮件加密、文件签名  |
| **[urlx](urlx/)**       | URL 处理 | URL 解析、构建      |

</details>

<details>
<summary><strong>🚀 并发&控制模块</strong></summary>

| 模块                      | 功能     | 设计模式         |
| ------------------------- | -------- | ---------------- |
| **[routine](routine/)**   | 协程管理 | 协程池、任务调度 |
| **[wait](wait/)**         | 等待控制 | 超时、重试、限流 |
| **[hystrix](hystrix/)**   | 熔断器   | 容错、降级       |
| **[singledo](singledo/)** | 单例模式 | 防重复执行       |
| **[event](event/)**       | 事件驱动 | 发布订阅模式     |

</details>

<details>
<summary><strong>🧪 开发&测试模块</strong></summary>

| 模块                        | 功能       | 开发阶段       |
| --------------------------- | ---------- | -------------- |
| **[fake](fake/)**           | 假数据生成 | 测试数据生成   |
| **[unit](unit/)**           | 测试辅助   | 单元测试工具   |
| **[pyroscope](pyroscope/)** | 性能分析   | 生产监控       |
| **[defaults](defaults/)**   | 默认值     | 配置初始化     |
| **[randx](randx/)**         | 随机数     | 安全随机数生成 |

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
    log.Printf("应用启动: %s", cal.String())

    // 5. 启动服务
    startServer(cfg)
}
```

### 错误处理最佳实践

```go
// ✅ 推荐：使用 Must 系列函数
func processData() string {
    data := utils.Must(loadData())        // 加载失败时 panic
    utils.MustSuccess(validateData(data)) // 验证失败时 panic
    return utils.MustOk(transformData(data)) // 转换失败时 panic
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

| 操作             | 耗时       | 内存分配 | 对比标准库        |
| ---------------- | ---------- | -------- | ----------------- |
| `candy.ToInt()`  | 12.3 ns/op | 0 B/op   | **3.2x faster**   |
| `json.Marshal()` | 156 ns/op  | 64 B/op  | **1.8x faster**   |
| `xtime.Now()`    | 45.2 ns/op | 0 B/op   | **2.1x faster**   |
| `utils.Must()`   | 2.1 ns/op  | 0 B/op   | **Zero overhead** |

### 性能特点

-   ⚡ **微秒级响应**：核心操作在微秒级完成
-   🧠 **内存友好**：使用 sync.Pool 减少 GC 压力
-   🔄 **零分配**：关键路径避免内存分配
-   🚀 **并发优化**：针对高并发场景优化

> 📈 详细性能报告：[性能测试文档](docs/performance.md)

## 🤝 贡献指南

我们欢迎任何形式的贡献！

### 贡献流程

1. 🍴 Fork 项目
2. 🌿 创建特性分支: `git checkout -b feature/amazing-feature`
3. 📝 编写代码和测试
4. 🧪 确保测试通过: `go test ./...`
5. 📤 提交 PR

### 开发规范

-   ✅ 遵循 [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
-   📖 所有公共 API 必须有 godoc 注释
-   🧪 新功能必须包含测试用例
-   📊 保持测试覆盖率 > 80%
-   🔄 保持向后兼容性

> 📋 详细规范：[贡献指南](CONTRIBUTING.md)

## 📄 许可证

本项目采用 GNU Affero General Public License v3.0 许可证。

查看 [LICENSE](LICENSE) 文件了解详情。

## 🌟 社区支持

### 获取帮助

-   📖 **文档**：[完整文档](docs/)
-   🐛 **Bug 报告**：[GitHub Issues](https://github.com/lazygophers/utils/issues)
-   💬 **讨论**：[GitHub Discussions](https://github.com/lazygophers/utils/discussions)
-   ❓ **问答**：[Stack Overflow](https://stackoverflow.com/questions/tagged/lazygophers-utils)

### 项目统计

| 指标            | 数值                                                                   | 说明             |
| --------------- | ---------------------------------------------------------------------- | ---------------- |
| 📦 模块数量     | 20+                                                                    | 涵盖各种常用功能 |
| 🧪 测试覆盖率   | 85%+                                                                   | 高质量代码保障   |
| 📝 文档完整度   | 95%+                                                                   | 详尽的使用说明   |
| ⚡ 性能等级     | A+                                                                     | 经过基准测试优化 |
| 🌟 GitHub Stars | ![GitHub stars](https://img.shields.io/github/stars/lazygophers/utils) | 社区认可度       |

### 致谢

感谢所有贡献者的辛勤付出！

[![Contributors](https://contrib.rocks/image?repo=lazygophers/utils)](https://github.com/lazygophers/utils/graphs/contributors)

---

<div align="center">

**如果这个项目对你有帮助，请给我们一个 ⭐ Star！**

[🚀 开始使用](#-快速开始) • [📖 查看文档](docs/) • [🤝 加入社区](https://github.com/lazygophers/utils/discussions)

</div>
