---
pageType: home

hero:
    name: LazyGophers Utils
    text: 现代化 Go 工具库
    tagline: 为现代开发工作流程提供强大支持
    actions:
        - theme: brand
          text: 快速开始
          link: /zh-CN/guide/getting-started
        - theme: alt
          text: 浏览模块
          link: /zh-CN/modules/overview

features:
    - title: "模块化设计"
      details: 20+ 个专业模块，只导入您需要的内容，保持项目轻量高效
      icon: 🧩
    - title: "性能优先"
      details: 为速度和最小内存使用进行优化，关键操作比标准库快 2-3 倍
      icon: ⚡
    - title: "类型安全"
      details: 利用 Go 泛型实现编译时安全，避免运行时类型错误
      icon: 🛡️
    - title: "生产就绪"
      details: Goroutine 安全且经过实战检验，可直接用于生产环境
      icon: 🔒
    - title: "开发者友好"
      details: 全面的文档和示例，快速上手，提升开发效率
      icon: 📖
    - title: "时间与调度"
      details: 支持农历、中国生肖、节气，以及多种工作时间计算
      icon: ⏰
---

## 🌍 多语言支持

[简体中文](/zh-CN/) • [繁體中文](/zh-TW/) • [English](/en/)

## 🎯 什么是 LazyGophers Utils？

LazyGophers Utils 是一个全面的 Go 工具库，为常见开发任务提供 **20+ 个专业模块**。采用现代 Go 实践构建，它提供类型安全、高性能的解决方案，可以无缝集成到任何 Go 项目中。

### ✨ 为什么选择 LazyGophers Utils？

-   **🧩 模块化设计** - 只导入您需要的内容
-   **⚡ 性能优先** - 为速度和最小内存使用进行优化
-   **🛡️ 类型安全** - 利用 Go 泛型实现编译时安全
-   **🔒 生产就绪** - Goroutine 安全且经过实战检验
-   **📖 开发者友好** - 全面的文档和示例

## 🚀 快速开始

### 安装

```bash
go get github.com/lazygophers/utils
```

### 30 秒示例

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils"
    "github.com/lazygophers/utils/candy"
    "github.com/lazygophers/utils/xtime"
)

func main() {
    // 简化的错误处理
    data := utils.Must(loadData())

    // 无需麻烦的类型转换
    userAge := candy.ToInt("25")
    isActive := candy.ToBool("true")

    // 高级时间处理
    calendar := xtime.NowCalendar()
    fmt.Printf("Today: %s\n", calendar.String())
    fmt.Printf("Lunar: %s\n", calendar.LunarDate())
}

func loadData() (string, error) {
    return "Hello, World!", nil
}
```

## 📦 模块概览

### 🔧 核心工具

| 模块                                                                         | 用途       | 主要函数                              |
| ---------------------------------------------------------------------------- | ---------- | ------------------------------------- |
| **[must.go](https://github.com/lazygophers/utils/blob/main/must.go)**        | 错误断言   | `Must()`, `MustSuccess()`, `MustOk()` |
| **[orm.go](https://github.com/lazygophers/utils/blob/main/orm.go)**          | 数据库操作 | `Scan()`, `Value()`                   |
| **[validate.go](https://github.com/lazygophers/utils/blob/main/validator/)** | 数据验证   | `Validate()`                          |

### 🍭 数据处理

| 模块                                                                   | 用途             | 亮点                |
| ---------------------------------------------------------------------- | ---------------- | ------------------- |
| **[candy/](https://github.com/lazygophers/utils/tree/main/candy)**     | 类型转换糖       | 零分配转换          |
| **[json/](https://github.com/lazygophers/utils/tree/main/json)**       | 增强的 JSON 处理 | 更好的错误消息      |
| **[stringx/](https://github.com/lazygophers/utils/tree/main/stringx)** | 字符串工具       | Unicode 感知操作    |
| **[anyx/](https://github.com/lazygophers/utils/tree/main/anyx)**       | Interface{} 辅助 | 类型安全的 any 操作 |

### ⏰ 时间与调度

| 模块                                                                     | 用途         | 特殊功能                      |
| ------------------------------------------------------------------------ | ------------ | ----------------------------- |
| **[xtime/](https://github.com/lazygophers/utils/tree/main/xtime)**       | 高级时间处理 | 🌙 农历, 🐲 中国生肖, 🌾 节气 |
| **[xtime996/](https://github.com/lazygophers/utils/tree/main/xtime996)** | 996 工作时间 | 工作时间计算                  |
| **[xtime955/](https://github.com/lazygophers/utils/tree/main/xtime955)** | 955 工作时间 | 平衡时间表支持                |
| **[xtime007/](https://github.com/lazygophers/utils/tree/main/xtime007)** | 24/7 操作    | 始终在线的时间工具            |

### 🔧 系统与配置

| 模块                                                                   | 用途         | 用例                            |
| ---------------------------------------------------------------------- | ------------ | ------------------------------- |
| **[config/](https://github.com/lazygophers/utils/tree/main/config)**   | 配置管理     | JSON, YAML, TOML, INI, HCL 支持 |
| **[runtime/](https://github.com/lazygophers/utils/tree/main/runtime)** | 运行时信息   | 系统检测和诊断                  |
| **[osx/](https://github.com/lazygophers/utils/tree/main/osx)**         | 操作系统操作 | 文件和进程管理                  |
| **[app/](https://github.com/lazygophers/utils/tree/main/app)**         | 应用框架     | 生命周期管理                    |
| **[atexit/](https://github.com/lazygophers/utils/tree/main/atexit)**   | 优雅关闭     | 干净的退出处理                  |

### 🌐 网络与安全

| 模块                                                                   | 用途      | 功能                 |
| ---------------------------------------------------------------------- | --------- | -------------------- |
| **[network/](https://github.com/lazygophers/utils/tree/main/network)** | HTTP 工具 | 连接池, 重试逻辑     |
| **[cryptox/](https://github.com/lazygophers/utils/tree/main/cryptox)** | 加密函数  | 哈希, 加密, 安全随机 |
| **[pgp/](https://github.com/lazygophers/utils/tree/main/pgp)**         | PGP 操作  | 邮件加密, 文件签名   |
| **[urlx/](https://github.com/lazygophers/utils/tree/main/urlx)**       | URL 操作  | 解析, 构建, 验证     |

### 🚀 并发与控制流

| 模块                                                                     | 用途           | 模式                 |
| ------------------------------------------------------------------------ | -------------- | -------------------- |
| **[routine/](https://github.com/lazygophers/utils/tree/main/routine)**   | Goroutine 管理 | 工作池, 任务调度     |
| **[wait/](https://github.com/lazygophers/utils/tree/main/wait)**         | 流程控制       | 超时, 重试, 速率限制 |
| **[hystrix/](https://github.com/lazygophers/utils/tree/main/hystrix)**   | 熔断器         | 容错, 优雅降级       |
| **[singledo/](https://github.com/lazygophers/utils/tree/main/singledo)** | 单例执行       | 防止重复操作         |
| **[event/](https://github.com/lazygophers/utils/tree/main/event)**       | 事件系统       | 发布/订阅模式实现    |

### 🧪 开发与测试

| 模块                                                                       | 用途         | 开发阶段           |
| -------------------------------------------------------------------------- | ------------ | ------------------ |
| **[fake/](https://github.com/lazygophers/utils/tree/main/fake)**           | 测试数据生成 | 单元测试, 集成测试 |
| **[randx/](https://github.com/lazygophers/utils/tree/main/randx)**         | 随机工具     | 加密安全的随机数   |
| **[defaults/](https://github.com/lazygophers/utils/tree/main/defaults)**   | 默认值       | 结构体初始化       |
| **[pyroscope/](https://github.com/lazygophers/utils/tree/main/pyroscope)** | 性能分析     | 生产监控           |

## 📊 性能亮点

| 操作             | 时间       | 内存    | vs 标准库     |
| ---------------- | ---------- | ------- | ------------- |
| `candy.ToInt()`  | 12.3 ns/op | 0 B/op  | **快 3.2 倍** |
| `json.Marshal()` | 156 ns/op  | 64 B/op | **快 1.8 倍** |
| `xtime.Now()`    | 45.2 ns/op | 0 B/op  | **快 2.1 倍** |
| `utils.Must()`   | 2.1 ns/op  | 0 B/op  | **零开销**    |

## 🤝 贡献

我们欢迎贡献！以下是入门方法：

### 快速贡献指南

1. **Fork** 仓库
2. **创建** 功能分支：`git checkout -b feature/amazing-feature`
3. **编写** 代码和测试
4. **确保** 测试通过：`go test ./...`
5. **提交** 拉取请求

## 📄 许可证

本项目采用 **GNU Affero General Public License v3.0** 许可。

详见 [LICENSE](https://github.com/lazygophers/utils/blob/main/LICENSE) 文件。

---

<div align="center">

**⭐ 如果这个项目帮助您构建更好的 Go 应用，请给它一个星标！**

[🚀 开始使用](/zh-CN/guide/getting-started) • [📖 浏览模块](/zh-CN/modules/overview) • [🤝 贡献](https://github.com/lazygophers/utils/blob/main/CONTRIBUTING.md)

_由 LazyGophers 团队用 ❤️ 构建_

</div>
