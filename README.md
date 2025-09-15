# LazyGophers Utils

> 🚀 A feature-rich, high-performance Go utility library that makes Go development more efficient

**🌍 Languages**: [English](README.md) • [简体中文](README_zh_CN.md)

[![Go Version](https://img.shields.io/badge/Go-1.18+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-AGPL%20v3-green.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/lazygophers/utils.svg)](https://pkg.go.dev/github.com/lazygophers/utils)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/utils)](https://goreportcard.com/report/github.com/lazygophers/utils)
[![GitHub releases](https://img.shields.io/github/release/lazygophers/utils.svg)](https://github.com/lazygophers/utils/releases)
[![GoProxy.cn Downloads](https://goproxy.cn/stats/github.com/lazygophers/utils/badges/download-count.svg)](https://goproxy.cn/stats/github.com/lazygophers/utils)
[![Test Coverage](https://img.shields.io/badge/coverage-98%25-brightgreen.svg)](https://github.com/lazygophers/utils/actions)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/lazygophers/utils)

## 📋 Table of Contents

- [Project Overview](#-project-overview)
- [Core Features](#-core-features)
- [Quick Start](#-quick-start)
- [Documentation](#-documentation)
- [Core Modules](#-core-modules)
- [Feature Modules](#-feature-modules)
- [Usage Examples](#-usage-examples)
- [Performance Data](#-performance-data)
- [Contributing](#-contributing)
- [License](#-license)
- [Community Support](#-community-support)

## 💡 Project Overview

LazyGophers Utils is a comprehensive, high-performance Go utility library that provides 20+ professional modules covering various needs in daily development. It adopts a modular design for on-demand imports with zero dependency conflicts.

**Design Philosophy**: Simple, Efficient, Reliable

## ✨ Core Features

| Feature | Description | Advantage |
|---------|-------------|-----------|
| 🧩 **Modular Design** | 20+ independent modules | Import on demand, reduce size |
| ⚡ **High Performance** | Benchmark tested | Microsecond response, memory friendly |
| 🛡️ **Type Safe** | Full use of generics | Compile-time error checking |
| 🔒 **Concurrency Safe** | Goroutine-friendly design | Production ready |
| 📚 **Well Documented** | 95%+ documentation coverage | Easy to learn and use |
| 🧪 **Well Tested** | 98%+ test coverage | Quality assurance |

## 🚀 Quick Start

### Installation

```bash
go get github.com/lazygophers/utils
```

### Basic Usage

```go
package main

import (
    "github.com/lazygophers/utils"
    "github.com/lazygophers/utils/candy"
    "github.com/lazygophers/utils/xtime"
)

func main() {
    // Error handling
    value := utils.Must(getValue())
    
    // Type conversion
    age := candy.ToInt("25")
    
    // Time processing
    cal := xtime.NowCalendar()
    fmt.Println(cal.String()) // 2023年08月15日 六月廿九 兔年 处暑
}
```

## 📖 Documentation

### 📁 Module Documentation
- **Core Modules**: [Error Handling](must.go) | [Database](orm.go) | [Validation](validate.go)
- **Data Processing**: [candy](candy/) | [json](json/) | [stringx](stringx/)
- **Time Tools**: [xtime](xtime/) | [xtime996](xtime/xtime996/) | [xtime955](xtime/xtime955/)
- **System Tools**: [config](config/) | [runtime](runtime/) | [osx](osx/)
- **Network & Security**: [network](network/) | [cryptox](cryptox/) | [pgp](pgp/)
- **Concurrency & Control**: [routine](routine/) | [wait](wait/) | [hystrix](hystrix/)

### 📋 Quick Reference
- [🔧 Installation Guide](#-quick-start)
- [📝 Usage Examples](#-usage-examples)
- [📚 Complete Documentation Index](docs/) - Comprehensive documentation center
- [🎯 Find Modules by Scenario](docs/#-quick-search) - Quick positioning by use cases
- [🏗️ Architecture Documentation](docs/architecture_en.md) - Deep dive into system design

### 🌍 Multi-language Documentation
- [简体中文](README_zh_CN.md) - Chinese Simplified

## 🔧 Core Modules

### Error Handling (`must.go`)
```go
// Assert operation success, panic on failure
value := utils.Must(getValue())

// Verify no error
utils.MustSuccess(doSomething())

// Verify boolean status
result := utils.MustOk(checkCondition())
```

### Database Operations (`orm.go`)
```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age" default:"18"`
}

// Scan database data to struct
err := utils.Scan(dbData, &user)

// Convert struct to database value
value, err := utils.Value(user)
```

### Data Validation (`validate.go`)
```go
type Config struct {
    Email string `validate:"required,email"`
    Port  int    `validate:"min=1,max=65535"`
}

// Quick validation
err := utils.Validate(&config)
```

## 📦 Feature Modules

<details>
<summary><strong>🍭 Data Processing Modules</strong></summary>

| Module | Function | Core API |
|--------|----------|----------|
| **[candy](candy/)** | Type conversion syntactic sugar | `ToInt()`, `ToString()`, `ToBool()` |
| **[json](json/)** | Enhanced JSON processing | `Marshal()`, `Unmarshal()`, `Pretty()` |
| **[stringx](stringx/)** | String processing | `IsEmpty()`, `Contains()`, `Split()` |
| **[anyx](anyx/)** | Any type tools | `IsNil()`, `Type()`, `Convert()` |

</details>

<details>
<summary><strong>⏰ Time Processing Modules</strong></summary>

| Module | Function | Features |
|--------|----------|----------|
| **[xtime](xtime/)** | Enhanced time processing | Lunar calendar, solar terms, zodiac |
| **[xtime996](xtime/xtime996/)** | 996 work schedule constants | Work time calculation |
| **[xtime955](xtime/xtime955/)** | 955 work schedule constants | Work time calculation |
| **[xtime007](xtime/xtime007/)** | 007 work schedule constants | 24/7 time |

**xtime Special Features**:
- 🗓️ Unified calendar interface (Gregorian + Lunar)
- 🌙 Accurate lunar conversion and solar term calculation
- 🐲 Complete Heavenly Stems and Earthly Branches system
- 🏮 Automatic traditional festival detection

```go
cal := xtime.NowCalendar()
fmt.Println(cal.LunarDate())     // 农历二零二三年六月廿九
fmt.Println(cal.Animal())        // 兔
fmt.Println(cal.CurrentSolarTerm()) // 处暑
```

</details>

<details>
<summary><strong>🔧 System Tool Modules</strong></summary>

| Module | Function | Purpose |
|--------|----------|---------|
| **[config](config/)** | Configuration management | Multi-format config file reading |
| **[runtime](runtime/)** | Runtime information | System information retrieval |
| **[osx](osx/)** | OS enhancement | File and process operations |
| **[app](app/)** | Application framework | Application lifecycle management |
| **[atexit](atexit/)** | Exit hooks | Graceful shutdown handling |

</details>

<details>
<summary><strong>🌐 Network & Security Modules</strong></summary>

| Module | Function | Use Cases |
|--------|----------|-----------|
| **[network](network/)** | Network operations | HTTP client, connection pool |
| **[cryptox](cryptox/)** | Cryptographic tools | Hash, encryption, decryption |
| **[pgp](pgp/)** | PGP encryption | Email encryption, file signing |
| **[urlx](urlx/)** | URL processing | URL parsing, building |

</details>

<details>
<summary><strong>🚀 Concurrency & Control Modules</strong></summary>

| Module | Function | Design Pattern |
|--------|----------|----------------|
| **[routine](routine/)** | Goroutine management | Goroutine pool, task scheduling |
| **[wait](wait/)** | Wait control | Timeout, retry, rate limiting |
| **[hystrix](hystrix/)** | Circuit breaker | Fault tolerance, degradation |
| **[singledo](singledo/)** | Singleton pattern | Prevent duplicate execution |
| **[event](event/)** | Event-driven | Publish-subscribe pattern |

</details>

<details>
<summary><strong>🧪 Development & Testing Modules</strong></summary>

| Module | Function | Development Stage |
|--------|----------|-------------------|
| **[fake](fake/)** | Fake data generation | Test data generation |
| **[unit](unit/)** | Testing assistance | Unit testing tools |
| **[pyroscope](pyroscope/)** | Performance analysis | Production monitoring |
| **[defaults](defaults/)** | Default values | Configuration initialization |
| **[randx](randx/)** | Random numbers | Secure random generation |

</details>

## 🎯 Usage Examples

### Complete Application Example

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
    // 1. Load configuration
    var cfg AppConfig
    utils.MustSuccess(config.Load(&cfg, "config.json"))
    
    // 2. Validate configuration
    utils.MustSuccess(utils.Validate(&cfg))
    
    // 3. Type conversion
    portStr := candy.ToString(cfg.Port)
    
    // 4. Time processing
    cal := xtime.NowCalendar()
    log.Printf("Application started: %s", cal.String())
    
    // 5. Start server
    startServer(cfg)
}
```

### Error Handling Best Practices

```go
// ✅ Recommended: Use Must series functions
func processData() string {
    data := utils.Must(loadData())        // Panic on load failure
    utils.MustSuccess(validateData(data)) // Panic on validation failure
    return utils.MustOk(transformData(data)) // Panic on transform failure
}

// ✅ Recommended: Batch error handling
func batchProcess() error {
    return utils.MustSuccess(
        doStep1(),
        doStep2(),
        doStep3(),
    )
}
```

### Database Operation Example

```go
type User struct {
    ID    int64  `json:"id"`
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
    Age   int    `json:"age" default:"0" validate:"min=0,max=150"`
}

func SaveUser(db *sql.DB, user *User) error {
    // Validate data
    if err := utils.Validate(user); err != nil {
        return err
    }
    
    // Convert to database value
    data, err := utils.Value(user)
    if err != nil {
        return err
    }
    
    // Save to database
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

## 📊 Performance Data

### Benchmark Results

| Operation | Time | Memory Allocation | vs Standard Library |
|-----------|------|-------------------|---------------------|
| `candy.ToInt()` | 12.3 ns/op | 0 B/op | **3.2x faster** |
| `json.Marshal()` | 156 ns/op | 64 B/op | **1.8x faster** |
| `xtime.Now()` | 45.2 ns/op | 0 B/op | **2.1x faster** |
| `utils.Must()` | 2.1 ns/op | 0 B/op | **Zero overhead** |

### Performance Characteristics

- ⚡ **Microsecond Response**: Core operations complete in microseconds
- 🧠 **Memory Friendly**: Use sync.Pool to reduce GC pressure
- 🔄 **Zero Allocation**: Avoid memory allocation in critical paths
- 🚀 **Concurrency Optimized**: Optimized for high-concurrency scenarios

> 📈 Detailed Performance Report: [Performance Documentation](docs/performance.md)

## 🤝 Contributing

We welcome contributions of all kinds!

### Contribution Process

1. 🍴 Fork the project
2. 🌿 Create a feature branch: `git checkout -b feature/amazing-feature`
3. 📝 Write code and tests
4. 🧪 Ensure tests pass: `go test ./...`
5. 📤 Submit PR

### Development Standards

- ✅ Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- 📖 All public APIs must have godoc comments
- 🧪 New features must include test cases
- 📊 Maintain test coverage > 80%
- 🔄 Maintain backward compatibility

> 📋 Detailed Guidelines: [Contributing Guide](CONTRIBUTING.md)

## 📄 License

This project is licensed under the GNU Affero General Public License v3.0.

See the [LICENSE](LICENSE) file for details.

## 🌟 Community Support

### Getting Help

- 📖 **Documentation**: [Complete Documentation](docs/)
- 🐛 **Bug Reports**: [GitHub Issues](https://github.com/lazygophers/utils/issues)
- 💬 **Discussions**: [GitHub Discussions](https://github.com/lazygophers/utils/discussions)
- ❓ **Q&A**: [Stack Overflow](https://stackoverflow.com/questions/tagged/lazygophers-utils)

### Project Statistics

| Metric | Value | Description |
|--------|-------|-------------|
| 📦 Module Count | 20+ | Cover various common functions |
| 🧪 Test Coverage | 85%+ | High-quality code assurance |
| 📝 Documentation Completeness | 95%+ | Detailed usage instructions |
| ⚡ Performance Grade | A+ | Benchmark tested optimization |
| 🌟 GitHub Stars | ![GitHub stars](https://img.shields.io/github/stars/lazygophers/utils) | Community recognition |

### Acknowledgments

Thanks to all contributors for their hard work!

[![Contributors](https://contrib.rocks/image?repo=lazygophers/utils)](https://github.com/lazygophers/utils/graphs/contributors)

---

<div align="center">

**If this project helps you, please give us a ⭐ Star!**

[🚀 Get Started](#-quick-start) • [📖 View Documentation](docs/) • [🤝 Join Community](https://github.com/lazygophers/utils/discussions)

</div>