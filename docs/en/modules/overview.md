---
title: Module Overview
---

# Module Overview

LazyGophers Utils provides 20+ specialized modules covering all aspects of Go development. All modules are organized by category for easy navigation.

## 🔧 Core Utilities

Basic utility modules providing error handling, database operations, and data validation.

- [must](./core/must.md) - Error assertion utilities
- [orm](./core/orm.md) - Database operation utilities
- [validator](./core/validator.md) - Data validation module

## 🍭 Data Processing

Provides type conversion, JSON processing, string operations, and other data processing utilities.

- [candy](./data/candy.md) - Type conversion sugar (zero allocation, 3.2x faster)
- [json](./data/json.md) - Enhanced JSON handling
- [stringx](./data/stringx.md) - String utilities
- [anyx](./data/anyx.md) - Interface{} helpers

## ⏰ Time & Scheduling

Advanced time processing, work schedule calculations, and scheduling utilities.

- [xtime](./time/xtime.md) - Advanced time processing (🌙 Lunar, 🐲 Zodiac, 🌾 Solar Terms)
- [xtime996](./time/xtime996.md) - 996 work schedule
- [xtime955](./time/xtime955.md) - 955 work schedule
- [xtime007](./time/xtime007.md) - 24/7 operations

## 🔧 System & Configuration

Configuration management, runtime information, OS operations, and other system-level utilities.

- [config](./system/config.md) - Configuration management (JSON, YAML, TOML, INI, HCL)
- [runtime](./system/runtime.md) - Runtime information
- [osx](./system/osx.md) - OS operations
- [app](./system/app.md) - Application framework
- [atexit](./system/atexit.md) - Graceful shutdown

## 🌐 Network & Security

HTTP utilities, cryptographic functions, PGP operations, and other network and security tools.

- [network](./network/network.md) - HTTP utilities (connection pooling, retry logic)
- [cryptox](./network/cryptox.md) - Cryptographic functions (hashing, encryption, secure random)
- [pgp](./network/pgp.md) - PGP operations
- [urlx](./network/urlx.md) - URL manipulation

## 🚀 Concurrency & Control Flow

Goroutine management, flow control, circuit breakers, and other concurrency and flow control utilities.

- [routine](./concurrency/routine.md) - Goroutine management (worker pools, task scheduling)
- [wait](./concurrency/wait.md) - Flow control (timeout, retry, rate limiting)
- [hystrix](./concurrency/hystrix.md) - Circuit breaker (fault tolerance, graceful degradation)
- [singledo](./concurrency/singledo.md) - Singleton execution
- [event](./concurrency/event.md) - Event system (pub/sub)

## 🧪 Development & Testing

Test data generation, random utilities, default values, and other development and testing utilities.

- [fake](./dev/fake.md) - Test data generation
- [randx](./dev/randx.md) - Random utilities (cryptographically secure)
- [defaults](./dev/defaults.md) - Default values
- [pyroscope](./dev/pyroscope.md) - Performance profiling

## 📊 Performance Highlights

| Operation | Time | Memory | vs Standard Library |
|-----------|------|--------|-------------------|
| `candy.ToInt()` | 12.3 ns/op | 0 B/op | **3.2x faster** |
| `json.Marshal()` | 156 ns/op | 64 B/op | **1.8x faster** |
| `xtime.Now()` | 45.2 ns/op | 0 B/op | **2.1x faster** |
| `utils.Must()` | 2.1 ns/op | 0 B/op | **Zero overhead** |

## Next Steps

- View [API Documentation](/en/api/overview) for detailed API information
- View [Getting Started](/en/guide/getting-started) to start using
- Visit [GitHub Repository](https://github.com/lazygophers/utils) for more examples
