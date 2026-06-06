# LazyGophers Utils

> 🚀 A powerful, modular Go utility library designed for modern development workflows

**🌍 Languages**: [English](README.md) • [中文](README_zh.md) • [繁體中文](README_zh-hant.md) • [Español](README_es.md) • [Français](README_fr.md) • [Русский](README_ru.md) • [العربية](README_ar.md)

📖 **Documentation**: [lazygophers.github.io/utils](https://lazygophers.github.io/utils/)

[![Go Version](https://img.shields.io/badge/Go-1.25+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-AGPL%20v3-green.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/lazygophers/utils.svg)](https://pkg.go.dev/github.com/lazygophers/utils)
[![Test Coverage](https://img.shields.io/badge/coverage-98.7%25-brightgreen)](https://github.com/lazygophers/utils/actions/workflows/coverage-update.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/utils)](https://goreportcard.com/report/github.com/lazygophers/utils)
[![GoProxy.cn Downloads](https://goproxy.cn/stats/github.com/lazygophers/utils/badges/download-count.svg)](https://goproxy.cn/stats/github.com/lazygophers/utils)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/lazygophers/utils)

---

## 🎯 What is LazyGophers Utils?

LazyGophers Utils is a comprehensive Go utility library that provides **20+ specialized modules** for common development tasks. Built with modern Go practices, it offers type-safe, high-performance solutions that integrate seamlessly into any Go project.

### ✨ Why Choose LazyGophers Utils?

- **🧩 Modular by Design** - Import only what you need
- **⚡ Performance First** - Optimized for speed and minimal memory usage
- **🛡️ Type Safety** - Leverages Go generics for compile-time safety
- **🔒 Production Ready** - Goroutine-safe and battle-tested
- **📖 Developer Friendly** - Comprehensive documentation and examples

---

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
    data := utils.Must(loadData()) // Panics on error

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

---

## 📦 Module Overview

### 🔧 Core Utilities

| Module | Purpose | Key Functions |
|--------|---------|---------------|
| **[must.go](must.go)** | Error assertion | `Must()`, `MustSuccess()`, `MustOk()` |
| **[orm.go](orm.go)** | Database operations | `Scan()`, `Value()` |
| **[validator/](validator/)** | Data validation | Email, URL, strong password, IPv4, UUID, MAC |

### 🍭 Data Processing

| Module | Purpose | Highlights |
|--------|---------|------------|
| **[candy/](candy/)** | Type conversion sugar | Zero-allocation conversions |
| **[json/](json/)** | Enhanced JSON handling | Better error messages |
| **[stringx/](stringx/)** | String utilities | Unicode-aware operations |
| **[anyx/](anyx/)** | Interface{} helpers | Type-safe any operations |

### ⏰ Time & Scheduling

| Module | Purpose | Special Features |
|--------|---------|------------------|
| **[xtime/](xtime/)** | Advanced time processing | 🌙 Lunar calendar, 🐲 Chinese zodiac, 🌾 Solar terms |
| **[xtime996/](xtime996/)** | 996 work schedule | Work time calculations |
| **[xtime955/](xtime955/)** | 955 work schedule | Balanced schedule support |
| **[xtime007/](xtime007/)** | 24/7 operations | Always-on time utilities |

### 🔧 System & Configuration

| Module | Purpose | Use Cases |
|--------|---------|-----------|
| **[config/](config/)** | Configuration management | JSON, YAML, TOML, INI, HCL support |
| **[runtime/](runtime/)** | Runtime information | System detection and diagnostics |
| **[osx/](osx/)** | OS operations | File and process management |
| **[app/](app/)** | Application framework | Lifecycle management |
| **[atexit/](atexit/)** | Graceful shutdown | Clean exit handling |

### 🌐 Network & Security

| Module | Purpose | Features |
|--------|---------|----------|
| **[network/](network/)** | HTTP utilities | Connection pooling, retry logic |
| **[cryptox/](cryptox/)** | Cryptographic functions | Hashing, encryption, secure random ([Security Guide](cryptox/SECURITY.md)) |
| **[pgp/](pgp/)** | PGP operations | Email encryption, file signing |
| **[urlx/](urlx/)** | URL manipulation | Parsing, building, validation |

### 🚀 Concurrency & Control Flow

| Module | Purpose | Patterns |
|--------|---------|----------|
| **[routine/](routine/)** | Goroutine management | Worker pools, task scheduling |
| **[wait/](wait/)** | Flow control | Timeout, retry, rate limiting |
| **[hystrix/](hystrix/)** | Circuit breaker | Fault tolerance, graceful degradation |
| **[singledo/](singledo/)** | Singleton execution | Prevent duplicate operations |
| **[event/](event/)** | Event system | Pub/sub pattern implementation |

### 🧪 Development & Testing

| Module | Purpose | Development Stage |
|--------|---------|-------------------|
| **[fake/](fake/)** | Test data generation | Unit testing, integration testing |
| **[randx/](randx/)** | Random utilities | Cryptographically secure random |
| **[defaults/](defaults/)** | Default values | Struct initialization |
| **[pyroscope/](pyroscope/)** | Performance profiling | Production monitoring |

---

## 💡 Real-World Examples

### Configuration Management

```go
type AppConfig struct {
    Database string `json:"database" validate:"required"`
    Port     int    `json:"port" default:"8080" validate:"min=1,max=65535"`
    Debug    bool   `json:"debug" default:"false"`
}

func main() {
    var cfg AppConfig

    // Load from any format: JSON, YAML, TOML, etc.
    utils.MustSuccess(config.Load(&cfg, "config.yaml"))

    // Validate configuration
    utils.MustSuccess(utils.Validate(&cfg))

    fmt.Printf("Server starting on port %d\n", cfg.Port)
}
```

### Database Operations

```go
type User struct {
    ID    int64  `json:"id"`
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
    Age   int    `json:"age" default:"0" validate:"min=0,max=150"`
}

func SaveUser(db *sql.DB, user *User) error {
    // Validate struct
    if err := utils.Validate(user); err != nil {
        return err
    }

    // Convert to database format
    data, err := utils.Value(user)
    if err != nil {
        return err
    }

    // Save to database
    _, err = db.Exec("INSERT INTO users (data) VALUES (?)", data)
    return err
}
```

### Advanced Time Processing

```go
func timeExample() {
    cal := xtime.NowCalendar()

    // Gregorian calendar
    fmt.Printf("Date: %s\n", cal.Format("2006-01-02"))

    // Chinese lunar calendar
    fmt.Printf("Lunar: %s\n", cal.LunarDate())          // 农历二零二三年六月廿九
    fmt.Printf("Animal: %s\n", cal.Animal())            // 兔 (Rabbit)
    fmt.Printf("Solar Term: %s\n", cal.CurrentSolarTerm()) // 处暑 (End of Heat)

    // Work schedule calculations
    if xtime996.IsWorkTime(time.Now()) {
        fmt.Println("Time to work! (996 schedule)")
    }
}
```

### Concurrent Processing

```go
func processingExample() {
    // Create a worker pool
    pool := routine.NewPool(10) // 10 workers
    defer pool.Close()

    // Submit tasks with circuit breaker protection
    for i := 0; i < 100; i++ {
        taskID := i
        pool.Submit(func() {
            // Circuit breaker protects against failures
            result := hystrix.Do("process-task", func() (interface{}, error) {
                return processTask(taskID)
            })

            fmt.Printf("Task %d result: %v\n", taskID, result)
        })
    }

    // Wait for completion with timeout
    wait.For(5*time.Second, func() bool {
        return pool.Running() == 0
    })
}
```

---

## 🎨 Design Philosophy

### Error Handling Strategy

LazyGophers Utils promotes a **fail-fast** approach for development efficiency:

```go
// Traditional Go error handling
data, err := risky.Operation()
if err != nil {
    return nil, fmt.Errorf("operation failed: %w", err)
}

// LazyGophers approach - cleaner, faster development
data := utils.Must(risky.Operation()) // Panics on error
```

### Type Safety with Generics

Modern Go generics enable compile-time safety:

```go
// Type-safe operations
func process[T constraints.Ordered](items []T) T {
    return candy.Max(items...) // Works with any ordered type
}

// Runtime safety
value := utils.MustOk(getValue()) // Panics if second return value is false
```

### Performance Optimization

Every module is benchmarked and optimized:

- **Zero-allocation** paths in critical functions
- **sync.Pool** usage to reduce GC pressure
- **Efficient algorithms** for common operations
- **Minimal dependencies** to reduce binary size

---

## 📊 Performance Highlights

| Operation | Time | Memory | vs Standard Library |
|-----------|------|--------|-------------------|
| `candy.ToInt()` | 12.3 ns/op | 0 B/op | **3.2x faster** |
| `json.Marshal()` | 156 ns/op | 64 B/op | **1.8x faster** |
| `xtime.Now()` | 45.2 ns/op | 0 B/op | **2.1x faster** |
| `utils.Must()` | 0.28 ns/op | 0 B/op | **Near-zero overhead** |

---

## 🤝 Contributing

We welcome contributions! Here's how to get started:

### Quick Contribution Guide

1. **Fork** the repository
2. **Create** a feature branch: `git checkout -b feature/amazing-feature`
3. **Write** code with tests
4. **Ensure** tests pass: `go test ./...`
5. **Submit** a pull request

### Development Standards

- ✅ Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- 📖 All public APIs must have godoc comments
- 🧪 New features require comprehensive tests
- 📊 Maintain high test coverage
- 🔄 Preserve backward compatibility

### Building and Testing

```bash
# Run tests
make test

# Run tests with coverage
make test-coverage

# Lint code
make lint

# Format code
make fmt

# Full development cycle
make check
```

---

## 📄 License

This project is licensed under the **GNU Affero General Public License v3.0**.

See the [LICENSE](LICENSE) file for details.

---

## 🌟 Community & Support

### Get Help

- 📖 **Documentation**: [Complete API Reference](https://pkg.go.dev/github.com/lazygophers/utils)
- 🐛 **Bug Reports**: [GitHub Issues](https://github.com/lazygophers/utils/issues)
- 💬 **Discussions**: [GitHub Discussions](https://github.com/lazygophers/utils/discussions)
- ❓ **Questions**: [Stack Overflow](https://stackoverflow.com/questions/tagged/lazygophers-utils)

### Acknowledgments

Thanks to all our contributors who make this project possible!

[![Contributors](https://contrib.rocks/image?repo=lazygophers/utils)](https://github.com/lazygophers/utils/graphs/contributors)

---

<div align="center">

**⭐ Star this project if it helps you build better Go applications!**

[🚀 Get Started](#-quick-start) • [📖 Browse Modules](#-module-overview) • [🤝 Contribute](#-contributing)

*Built with ❤️ by the LazyGophers team*

</div>