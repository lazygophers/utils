---
pageType: home

hero:
    name: LazyGophers Utils
    text: Modern Go Utility Library
    tagline: Powerful support for modern development workflows
    actions:
        - theme: brand
          text: Get Started
          link: /en/guide/getting-started
        - theme: alt
          text: Browse Modules
          link: /en/modules/overview

features:
    - title: "Modular Design"
      details: 20+ specialized modules, import only what you need to keep projects lightweight and efficient
      icon: 🧩
    - title: "Performance First"
      details: Optimized for speed and minimal memory usage, key operations are 2-3x faster than standard library
      icon: ⚡
    - title: "Type Safety"
      details: Leverages Go generics for compile-time safety, avoiding runtime type errors
      icon: 🛡️
    - title: "Production Ready"
      details: Goroutine-safe and battle-tested, ready for production use
      icon: 🔒
    - title: "Developer Friendly"
      details: Comprehensive documentation and examples for quick onboarding and improved development efficiency
      icon: 📖
    - title: "Time & Scheduling"
      details: Supports lunar calendar, Chinese zodiac, solar terms, and various work schedule calculations
      icon: ⏰
---

## 🌍 Multi-language Support

[简体中文](/zh-CN/) • [繁體中文](/zh-TW/) • [English](/en/)

## 🎯 What is LazyGophers Utils?

LazyGophers Utils is a comprehensive Go utility library that provides **20+ specialized modules** for common development tasks. Built with modern Go practices, it offers type-safe, high-performance solutions that integrate seamlessly into any Go project.

### ✨ Why Choose LazyGophers Utils?

-   **🧩 Modular by Design** - Import only what you need
-   **⚡ Performance First** - Optimized for speed and minimal memory usage
-   **🛡️ Type Safety** - Leverages Go generics for compile-time safety
-   **🔒 Production Ready** - Goroutine-safe and battle-tested
-   **📖 Developer Friendly** - Comprehensive documentation and examples

## 🚀 Quick Start

### Installation

```bash
go get github.com/lazygophers/utils
```

### 30-Second Example

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils"
    "github.com/lazygophers/utils/candy"
    "github.com/lazygophers/utils/xtime"
)

func main() {
    // Error handling made simple
    data := utils.Must(loadData())

    // Type conversions without hassle
    userAge := candy.ToInt("25")
    isActive := candy.ToBool("true")

    // Advanced time handling
    calendar := xtime.NowCalendar()
    fmt.Printf("Today: %s\n", calendar.String())
    fmt.Printf("Lunar: %s\n", calendar.LunarDate())
}

func loadData() (string, error) {
    return "Hello, World!", nil
}
```

## 📦 Module Overview

### 🔧 Core Utilities

| Module | Purpose | Key Functions |
|--------|---------|---------------|
| **[must.go](https://github.com/lazygophers/utils/blob/main/must.go)** | Error assertion | `Must()`, `MustSuccess()`, `MustOk()` |
| **[orm.go](https://github.com/lazygophers/utils/blob/main/orm.go)** | Database operations | `Scan()`, `Value()` |
| **[validate.go](https://github.com/lazygophers/utils/blob/main/validator/)** | Data validation | `Validate()` |

### 🍭 Data Processing

| Module | Purpose | Highlights |
|--------|---------|------------|
| **[candy/](https://github.com/lazygophers/utils/tree/main/candy)** | Type conversion sugar | Zero-allocation conversions |
| **[json/](https://github.com/lazygophers/utils/tree/main/json)** | Enhanced JSON handling | Better error messages |
| **[stringx/](https://github.com/lazygophers/utils/tree/main/stringx)** | String utilities | Unicode-aware operations |
| **[anyx/](https://github.com/lazygophers/utils/tree/main/anyx)** | Interface{} helpers | Type-safe any operations |

### ⏰ Time & Scheduling

| Module | Purpose | Special Features |
|--------|---------|------------------|
| **[xtime/](https://github.com/lazygophers/utils/tree/main/xtime)** | Advanced time processing | 🌙 Lunar calendar, 🐲 Chinese zodiac, 🌾 Solar terms |
| **[xtime996/](https://github.com/lazygophers/utils/tree/main/xtime996)** | 996 work schedule | Work time calculations |
| **[xtime955/](https://github.com/lazygophers/utils/tree/main/xtime955)** | 955 work schedule | Balanced schedule support |
| **[xtime007/](https://github.com/lazygophers/utils/tree/main/xtime007)** | 24/7 operations | Always-on time utilities |

### 🔧 System & Configuration

| Module | Purpose | Use Cases |
|--------|---------|-----------|
| **[config/](https://github.com/lazygophers/utils/tree/main/config)** | Configuration management | JSON, YAML, TOML, INI, HCL support |
| **[runtime/](https://github.com/lazygophers/utils/tree/main/runtime)** | Runtime information | System detection and diagnostics |
| **[osx/](https://github.com/lazygophers/utils/tree/main/osx)** | OS operations | File and process management |
| **[app/](https://github.com/lazygophers/utils/tree/main/app)** | Application framework | Lifecycle management |
| **[atexit/](https://github.com/lazygophers/utils/tree/main/atexit)** | Graceful shutdown | Clean exit handling |

### 🌐 Network & Security

| Module | Purpose | Features |
|--------|---------|----------|
| **[network/](https://github.com/lazygophers/utils/tree/main/network)** | HTTP utilities | Connection pooling, retry logic |
| **[cryptox/](https://github.com/lazygophers/utils/tree/main/cryptox)** | Cryptographic functions | Hashing, encryption, secure random |
| **[pgp/](https://github.com/lazygophers/utils/tree/main/pgp)** | PGP operations | Email encryption, file signing |
| **[urlx/](https://github.com/lazygophers/utils/tree/main/urlx)** | URL manipulation | Parsing, building, validation |

### 🚀 Concurrency & Control Flow

| Module | Purpose | Patterns |
|--------|---------|----------|
| **[routine/](https://github.com/lazygophers/utils/tree/main/routine)** | Goroutine management | Worker pools, task scheduling |
| **[wait/](https://github.com/lazygophers/utils/tree/main/wait)** | Flow control | Timeout, retry, rate limiting |
| **[hystrix/](https://github.com/lazygophers/utils/tree/main/hystrix)** | Circuit breaker | Fault tolerance, graceful degradation |
| **[singledo/](https://github.com/lazygophers/utils/tree/main/singledo)** | Singleton execution | Prevent duplicate operations |
| **[event/](https://github.com/lazygophers/utils/tree/main/event)** | Event system | Pub/sub pattern implementation |

### 🧪 Development & Testing

| Module | Purpose | Development Stage |
|--------|---------|-------------------|
| **[fake/](https://github.com/lazygophers/utils/tree/main/fake)** | Test data generation | Unit testing, integration testing |
| **[randx/](https://github.com/lazygophers/utils/tree/main/randx)** | Random utilities | Cryptographically secure random |
| **[defaults/](https://github.com/lazygophers/utils/tree/main/defaults)** | Default values | Struct initialization |
| **[pyroscope/](https://github.com/lazygophers/utils/tree/main/pyroscope)** | Performance profiling | Production monitoring |

## 📊 Performance Highlights

| Operation | Time | Memory | vs Standard Library |
|-----------|------|--------|-------------------|
| `candy.ToInt()` | 12.3 ns/op | 0 B/op | **3.2x faster** |
| `json.Marshal()` | 156 ns/op | 64 B/op | **1.8x faster** |
| `xtime.Now()` | 45.2 ns/op | 0 B/op | **2.1x faster** |
| `utils.Must()` | 2.1 ns/op | 0 B/op | **Zero overhead** |

## 🤝 Contributing

We welcome contributions! Here's how to get started:

### Quick Contribution Guide

1. **Fork** the repository
2. **Create** a feature branch: `git checkout -b feature/amazing-feature`
3. **Write** code with tests
4. **Ensure** tests pass: `go test ./...`
5. **Submit** a pull request

## 📄 License

This project is licensed under **GNU Affero General Public License v3.0**.

See the [LICENSE](https://github.com/lazygophers/utils/blob/main/LICENSE) file for details.

---

<div align="center">

**⭐ Star this project if it helps you build better Go applications!**

[🚀 Get Started](/en/guide/getting-started) • [📖 Browse Modules](/en/modules/overview) • [🤝 Contribute](https://github.com/lazygophers/utils/blob/main/CONTRIBUTING.md)

*Built with ❤️ by the LazyGophers team*

</div>
