---
title: Module Overview
---

# Module Overview

LazyGophers Utils provides 20+ specialized modules covering all aspects of Go development.

## 🔧 Core Utilities

### must.go

Error assertion tool that simplifies error handling flow.

**Key Features:**
- `Must()` - Assert operation succeeds, panic on failure
- `MustSuccess()` - Assert error is nil
- `MustOk()` - Assert second return value is true

**Example:**
```go
data := utils.Must(loadData())
utils.MustSuccess(config.Load(&cfg, "config.json"))
value := utils.MustOk(getValue())
```

### orm.go

Database operation tool providing convenient data conversion methods.

**Key Features:**
- `Scan()` - Scan database results into structs
- `Value()` - Convert structs to database values

### validator

Data validation module supporting struct validation.

**Key Features:**
- `Validate()` - Validate struct data
- Built-in validation rules: `required`, `email`, `min`, `max`, etc.

## 🍭 Data Processing

### candy

Type conversion tool providing zero-allocation type conversions.

**Key Features:**
- `ToInt()` - String to integer
- `ToFloat()` - String to float
- `ToBool()` - String to boolean
- `ToString()` - Any type to string
- `ToSlice()` - Any type to slice
- `ToMap()` - Any type to map

**Performance:** Zero allocation, 3.2x faster than standard library

### json

Enhanced JSON handling with better error messages.

**Key Features:**
- `Marshal()` - JSON encoding
- `Unmarshal()` - JSON decoding
- More friendly error messages

### stringx

String utilities with Unicode-aware operations.

**Key Features:**
- `Rand()` - Generate random strings
- `Reverse()` - Reverse strings
- `Truncate()` - Truncate strings
- Unicode-aware string operations

### anyx

interface{} helpers providing type-safe any operations.

**Key Features:**
- Type-safe any operations
- Generic support

## ⏰ Time & Scheduling

### xtime

Advanced time processing with lunar calendar, zodiac, and solar terms support.

**Key Features:**
- `NowCalendar()` - Get current calendar
- `LunarDate()` - Get lunar date
- `Animal()` - Get zodiac animal
- `CurrentSolarTerm()` - Get current solar term
- `Format()` - Format dates

**Special Features:**
- 🌙 Lunar calendar support
- 🐲 Zodiac calculation
- 🌾 Solar term calculation

### xtime996

996 work schedule calculation.

**Key Features:**
- `IsWorkTime()` - Check if it's work time
- Work time calculations

### xtime955

955 work schedule calculation.

**Key Features:**
- `IsWorkTime()` - Check if it's work time
- Balanced schedule support

### xtime007

24/7 operations time utilities.

**Key Features:**
- Always-on time utilities
- 24/7 time handling

## 🔧 System & Configuration

### config

Configuration management supporting multiple configuration formats.

**Key Features:**
- `Load()` - Load configuration files
- Supported formats: JSON, YAML, TOML, INI, HCL

### runtime

Runtime information providing system detection and diagnostics.

**Key Features:**
- System information detection
- Runtime diagnostics

### osx

OS operations providing file and process management.

**Key Features:**
- File operations
- Process management

### app

Application framework providing lifecycle management.

**Key Features:**
- Application lifecycle management
- Startup and shutdown handling

### atexit

Graceful shutdown providing clean exit handling.

**Key Features:**
- Register exit handlers
- Graceful shutdown

## 🌐 Network & Security

### network

HTTP utilities providing connection pooling and retry logic.

**Key Features:**
- Connection pool management
- Retry logic
- HTTP client utilities

### cryptox

Cryptographic functions providing hashing, encryption, and secure random.

**Key Features:**
- Hash functions: MD5, SHA1, SHA256, SHA512
- Encryption: AES, DES, RSA, ECDSA, ECDH
- Secure random number generation
- UUID generation

### pgp

PGP operations providing email encryption and file signing.

**Key Features:**
- PGP encryption
- File signing
- Email encryption

### urlx

URL manipulation providing parsing, building, and validation.

**Key Features:**
- URL parsing
- URL building
- URL validation
- Query parameter handling

## 🚀 Concurrency & Control Flow

### routine

Goroutine management providing worker pools and task scheduling.

**Key Features:**
- `NewPool()` - Create worker pool
- `Submit()` - Submit tasks
- `Close()` - Close worker pool

### wait

Flow control providing timeout, retry, and rate limiting.

**Key Features:**
- `For()` - Wait for condition
- `Retry()` - Retry operations
- Timeout control
- Rate limiting

### hystrix

Circuit breaker providing fault tolerance and graceful degradation.

**Key Features:**
- `Do()` - Execute operations with circuit breaker protection
- Circuit breaker configuration
- Fault tolerance

### singledo

Singleton execution preventing duplicate operations.

**Key Features:**
- Ensure operation runs only once
- Prevent duplicate calculations

### event

Event system implementing pub/sub pattern.

**Key Features:**
- Event publishing
- Event subscription
- Event handling

## 🧪 Development & Testing

### fake

Test data generation supporting unit testing and integration testing.

**Key Features:**
- Generate random names
- Generate random addresses
- Generate random companies
- Generate random text
- Support for multiple languages

### randx

Random utilities providing cryptographically secure random.

**Key Features:**
- Cryptographically secure random numbers
- Random booleans
- Random numbers
- Random times

### defaults

Default values providing struct initialization.

**Key Features:**
- Set struct default values
- Read default values from tags

### pyroscope

Performance profiling providing production monitoring.

**Key Features:**
- Pyroscope integration
- Performance profiling
- Production monitoring

## Performance Comparison

| Operation | Time | Memory | vs Standard Library |
|-----------|------|--------|-------------------|
| `candy.ToInt()` | 12.3 ns/op | 0 B/op | **3.2x faster** |
| `json.Marshal()` | 156 ns/op | 64 B/op | **1.8x faster** |
| `xtime.Now()` | 45.2 ns/op | 0 B/op | **2.1x faster** |
| `utils.Must()` | 2.1 ns/op | 0 B/op | **Zero overhead** |

## Next Steps

- Check out [API Documentation](/en/api/overview) for detailed API information
- Check out [Getting Started](/en/guide/getting-started) to start using
- Visit [GitHub Repository](https://github.com/lazygophers/utils) for more examples
