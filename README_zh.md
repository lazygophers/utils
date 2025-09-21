# LazyGophers Utils

> 🚀 专为现代开发工作流设计的强大模块化 Go 工具库

**🌍 多语言**: [English](README.md) • [中文](README_zh.md) • [繁體中文](README_zh-hant.md) • [Español](README_es.md) • [Français](README_fr.md) • [Русский](README_ru.md) • [العربية](README_ar.md)

[![Go Version](https://img.shields.io/badge/Go-1.25+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-AGPL%20v3-green.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/lazygophers/utils.svg)](https://pkg.go.dev/github.com/lazygophers/utils)
[![Test Coverage](https://img.shields.io/badge/coverage-69.6%25-yellow)](https://github.com/lazygophers/utils/actions/workflows/coverage-update.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/utils)](https://goreportcard.com/report/github.com/lazygophers/utils)
[![GoProxy.cn Downloads](https://goproxy.cn/stats/github.com/lazygophers/utils/badges/download-count.svg)](https://goproxy.cn/stats/github.com/lazygophers/utils)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/lazygophers/utils)

---

## 🎯 什么是 LazyGophers Utils？

LazyGophers Utils 是一个全面的 Go 工具库，为常见开发任务提供 **20+ 个专业模块**。采用现代 Go 开发实践构建，提供类型安全、高性能的解决方案，可无缝集成到任何 Go 项目中。

### ✨ 为什么选择 LazyGophers Utils？

- **🧩 模块化设计** - 按需导入，减少依赖
- **⚡ 性能优先** - 针对速度和最小内存使用进行优化
- **🛡️ 类型安全** - 利用 Go 泛型实现编译时安全
- **🔒 生产就绪** - 协程安全，经过实战检验
- **📖 开发友好** - 完整的文档和示例

---

## 🚀 快速开始

### 安装

```bash
go get github.com/lazygophers/utils
```

### 30秒示例

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils"
    "github.com/lazygophers/utils/candy"
    "github.com/lazygophers/utils/xtime"
)

func main() {
    // 简化错误处理
    data := utils.Must(loadData()) // 出错时panic

    // 轻松类型转换
    userAge := candy.ToInt("25")
    isActive := candy.ToBool("true")

    // 高级时间处理
    calendar := xtime.NowCalendar()
    fmt.Printf("今天: %s\n", calendar.String())
    fmt.Printf("农历: %s\n", calendar.LunarDate())
}

func loadData() (string, error) {
    return "Hello, World!", nil
}
```

---

## 📦 模块概览

### 🔧 核心工具

| 模块 | 用途 | 核心函数 |
|------|------|----------|
| **[must.go](must.go)** | 错误断言 | `Must()`, `MustSuccess()`, `MustOk()` |
| **[orm.go](orm.go)** | 数据库操作 | `Scan()`, `Value()` |
| **[validate.go](validate.go)** | 数据验证 | `Validate()` |

### 🍭 数据处理

| 模块 | 用途 | 特色 |
|------|------|------|
| **[candy/](candy/)** | 类型转换语法糖 | 零内存分配转换 |
| **[json/](json/)** | 增强的JSON处理 | 更好的错误消息 |
| **[stringx/](stringx/)** | 字符串工具 | Unicode感知操作 |
| **[anyx/](anyx/)** | Interface{}助手 | 类型安全的any操作 |

### ⏰ 时间与调度

| 模块 | 用途 | 特殊功能 |
|------|------|----------|
| **[xtime/](xtime/)** | 高级时间处理 | 🌙 农历系统, 🐲 生肖, 🌾 节气 |
| **[xtime996/](xtime996/)** | 996工作制 | 工作时间计算 |
| **[xtime955/](xtime955/)** | 955工作制 | 平衡作息支持 |
| **[xtime007/](xtime007/)** | 24/7运营 | 全天候时间工具 |

### 🔧 系统与配置

| 模块 | 用途 | 使用场景 |
|------|------|----------|
| **[config/](config/)** | 配置管理 | 支持JSON, YAML, TOML, INI, HCL |
| **[runtime/](runtime/)** | 运行时信息 | 系统检测与诊断 |
| **[osx/](osx/)** | 操作系统操作 | 文件和进程管理 |
| **[app/](app/)** | 应用程序框架 | 生命周期管理 |
| **[atexit/](atexit/)** | 优雅关闭 | 清理退出处理 |

### 🌐 网络与安全

| 模块 | 用途 | 功能 |
|------|------|------|
| **[network/](network/)** | HTTP工具 | 连接池，重试逻辑 |
| **[cryptox/](cryptox/)** | 加密函数 | 哈希，加密，安全随机 |
| **[pgp/](pgp/)** | PGP操作 | 邮件加密，文件签名 |
| **[urlx/](urlx/)** | URL操作 | 解析，构建，验证 |

### 🚀 并发与控制流

| 模块 | 用途 | 设计模式 |
|------|------|----------|
| **[routine/](routine/)** | 协程管理 | 工作池，任务调度 |
| **[wait/](wait/)** | 流程控制 | 超时，重试，限流 |
| **[hystrix/](hystrix/)** | 熔断器 | 容错，优雅降级 |
| **[singledo/](singledo/)** | 单例执行 | 防止重复操作 |
| **[event/](event/)** | 事件系统 | 发布订阅模式实现 |

### 🧪 开发与测试

| 模块 | 用途 | 开发阶段 |
|------|------|----------|
| **[fake/](fake/)** | 测试数据生成 | 单元测试，集成测试 |
| **[randx/](randx/)** | 随机工具 | 加密安全随机 |
| **[defaults/](defaults/)** | 默认值 | 结构体初始化 |
| **[pyroscope/](pyroscope/)** | 性能分析 | 生产监控 |

---

## 💡 实际应用示例

### 配置管理

```go
type AppConfig struct {
    Database string `json:"database" validate:"required"`
    Port     int    `json:"port" default:"8080" validate:"min=1,max=65535"`
    Debug    bool   `json:"debug" default:"false"`
}

func main() {
    var cfg AppConfig

    // 从任何格式加载: JSON, YAML, TOML 等
    utils.MustSuccess(config.Load(&cfg, "config.yaml"))

    // 验证配置
    utils.MustSuccess(utils.Validate(&cfg))

    fmt.Printf("服务器启动在端口 %d\n", cfg.Port)
}
```

### 数据库操作

```go
type User struct {
    ID    int64  `json:"id"`
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
    Age   int    `json:"age" default:"0" validate:"min=0,max=150"`
}

func SaveUser(db *sql.DB, user *User) error {
    // 验证结构体
    if err := utils.Validate(user); err != nil {
        return err
    }

    // 转换为数据库格式
    data, err := utils.Value(user)
    if err != nil {
        return err
    }

    // 保存到数据库
    _, err = db.Exec("INSERT INTO users (data) VALUES (?)", data)
    return err
}
```

### 高级时间处理

```go
func timeExample() {
    cal := xtime.NowCalendar()

    // 公历
    fmt.Printf("日期: %s\n", cal.Format("2006-01-02"))

    // 中国农历
    fmt.Printf("农历: %s\n", cal.LunarDate())          // 农历二零二三年六月廿九
    fmt.Printf("生肖: %s\n", cal.Animal())            // 兔
    fmt.Printf("节气: %s\n", cal.CurrentSolarTerm()) // 处暑

    // 工作制计算
    if xtime996.IsWorkTime(time.Now()) {
        fmt.Println("该工作了！(996作息)")
    }
}
```

### 并发处理

```go
func processingExample() {
    // 创建工作池
    pool := routine.NewPool(10) // 10个工作者
    defer pool.Close()

    // 提交带熔断保护的任务
    for i := 0; i < 100; i++ {
        taskID := i
        pool.Submit(func() {
            // 熔断器保护防止故障
            result := hystrix.Do("process-task", func() (interface{}, error) {
                return processTask(taskID)
            })

            fmt.Printf("任务 %d 结果: %v\n", taskID, result)
        })
    }

    // 带超时的完成等待
    wait.For(5*time.Second, func() bool {
        return pool.Running() == 0
    })
}
```

---

## 🎨 设计哲学

### 错误处理策略

LazyGophers Utils 提倡 **快速失败** 方法提高开发效率：

```go
// 传统Go错误处理
data, err := risky.Operation()
if err != nil {
    return nil, fmt.Errorf("操作失败: %w", err)
}

// LazyGophers方法 - 更清洁，更快的开发
data := utils.Must(risky.Operation()) // 出错时panic
```

### 泛型类型安全

现代Go泛型实现编译时安全：

```go
// 类型安全操作
func process[T constraints.Ordered](items []T) T {
    return candy.Max(items...) // 适用于任何有序类型
}

// 运行时安全
value := utils.MustOk(getValue()) // 如果第二个返回值为false则panic
```

### 性能优化

每个模块都经过基准测试和优化：

- **零分配** 关键函数路径
- **sync.Pool** 使用减少GC压力
- **高效算法** 用于常见操作
- **最小依赖** 减少二进制大小

---

## 📊 性能亮点

| 操作 | 时间 | 内存 | 对比标准库 |
|------|------|------|------------|
| `candy.ToInt()` | 12.3 ns/op | 0 B/op | **快3.2倍** |
| `json.Marshal()` | 156 ns/op | 64 B/op | **快1.8倍** |
| `xtime.Now()` | 45.2 ns/op | 0 B/op | **快2.1倍** |
| `utils.Must()` | 2.1 ns/op | 0 B/op | **零开销** |

---

## 🤝 参与贡献

我们欢迎贡献！开始方法如下：

### 快速贡献指南

1. **Fork** 仓库
2. **创建** 功能分支: `git checkout -b feature/amazing-feature`
3. **编写** 代码和测试
4. **确保** 测试通过: `go test ./...`
5. **提交** pull request

### 开发标准

- ✅ 遵循 [Go代码审查评论](https://github.com/golang/go/wiki/CodeReviewComments)
- 📖 所有公共API必须有godoc注释
- 🧪 新功能需要全面的测试
- 📊 保持高测试覆盖率
- 🔄 保持向后兼容性

### 构建和测试

```bash
# 运行测试
make test

# 运行覆盖率测试
make test-coverage

# 代码检查
make lint

# 格式化代码
make fmt

# 完整开发周期
make check
```

---

## 📄 许可证

本项目采用 **GNU Affero General Public License v3.0** 许可证。

详情请查看 [LICENSE](LICENSE) 文件。

---

## 🌟 社区与支持

### 获取帮助

- 📖 **文档**: [完整API参考](https://pkg.go.dev/github.com/lazygophers/utils)
- 🐛 **Bug报告**: [GitHub Issues](https://github.com/lazygophers/utils/issues)
- 💬 **讨论**: [GitHub Discussions](https://github.com/lazygophers/utils/discussions)
- ❓ **问题**: [Stack Overflow](https://stackoverflow.com/questions/tagged/lazygophers-utils)

### 致谢

感谢所有让这个项目成为可能的贡献者！

[![Contributors](https://contrib.rocks/image?repo=lazygophers/utils)](https://github.com/lazygophers/utils/graphs/contributors)

---

<div align="center">

**⭐ 如果这个项目帮助您构建更好的Go应用程序，请给我们一个Star！**

[🚀 开始使用](#-快速开始) • [📖 浏览模块](#-模块概览) • [🤝 参与贡献](#-参与贡献)

*由 LazyGophers 团队用 ❤️ 构建*

</div>