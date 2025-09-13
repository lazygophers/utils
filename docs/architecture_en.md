# LazyGophers Utils - Architecture Documentation

## 🏗️ Overview

LazyGophers Utils is a comprehensive Go utility library designed with modularity, performance, and developer experience in mind. The library follows modern Go practices including extensive use of generics, atomic operations, and zero-copy optimizations.

## 📊 Project Statistics

- **Total Packages**: 25 independent modules
- **Lines of Code**: 56,847
- **Go Files**: 323
- **Test Coverage**: 85.8%
- **Go Version**: 1.24.0+

## 🎯 Design Principles

### 1. Modular Architecture
Each package is designed as an independent module that can be imported and used separately, minimizing dependencies and binary size.

### 2. Performance-First Approach
- Extensive use of atomic operations for thread-safe operations
- Lock-free algorithms where possible
- Memory alignment optimizations
- Zero-copy string/byte conversions using unsafe operations
- Object pooling for high-frequency operations

### 3. Type Safety
Heavy use of Go 1.18+ generics to provide type-safe APIs while maintaining performance.

### 4. Consistent Error Handling
All packages follow a consistent error handling pattern: log errors using `github.com/lazygophers/log` before returning them.

## 🏛️ Package Architecture

### Core Packages
These packages form the foundation of the library:

#### `utils` (Root Package)
- **Purpose**: Fundamental utilities for error handling, database operations, and validation
- **Key Features**:
  - `Must[T any](value T, err error) T` - Panic-on-error wrapper
  - `Scan()` and `Value()` - Database integration utilities
  - `Validate()` - Struct validation using go-playground/validator
- **Dependencies**: Minimal external dependencies

#### `candy`
- **Purpose**: Comprehensive type conversion and slice manipulation
- **Features**: 143 files, 15,963 lines of code
- **Key Functions**:
  - Type conversion: `ToBool()`, `ToString()`, `ToInt*()`, `ToFloat*()`
  - Functional programming: `All()`, `Any()`, `Filter()`, `Map()`
  - Slice utilities: `Unique()`, `Sort()`, `Shuffle()`, `Chunk()`
- **Performance**: Optimized with generics for type safety

#### `json`
- **Purpose**: Enhanced JSON operations with performance optimization
- **Features**:
  - Platform-specific optimization using sonic on supported platforms
  - Consistent API wrapper over different JSON implementations
- **Performance Critical**: Yes

### Infrastructure Packages

#### `runtime`
- **Purpose**: Runtime utilities and system information
- **Key Features**:
  - Comprehensive panic handling and recovery
  - System directory utilities (`ExecDir()`, `UserHomeDir()`, etc.)
  - Platform detection utilities

#### `routine`
- **Purpose**: Enhanced goroutine management
- **Features**:
  - Safe goroutine execution with panic recovery
  - Goroutine tracing using `github.com/petermattis/goid`
  - Automatic cleanup and error handling

#### `app`
- **Purpose**: Application lifecycle and build information
- **Features**:
  - Build metadata (commit, branch, tag, build date)
  - Environment detection
  - Version information

### Utility Packages

#### `stringx`
- **Purpose**: High-performance string manipulation
- **Features**: 11 files, 4,385 lines of code
- **Performance Optimizations**:
  - Zero-copy string/byte conversion using unsafe operations
  - ASCII fast-path optimizations for case conversion
  - Memory-efficient string operations
- **Key Functions**:
  - `ToString()` / `ToBytes()` - Zero-copy conversions
  - `Camel2Snake()` - Optimized case conversion
  - Unicode utilities with performance optimizations

#### `anyx`
- **Purpose**: Type-agnostic map operations and value extraction
- **Features**: 4 files, 3,999 lines of code
- **Key Capabilities**:
  - Thread-safe map operations with type conversion
  - Nested key support with dot notation
  - Comprehensive type extraction utilities
- **Dependencies**: Uses `candy` and `json` packages

#### `wait`
- **Purpose**: Advanced concurrency and synchronization utilities
- **Features**: 6 files, 1,323 lines of code
- **Key Components**:
  - `Async()` - Goroutine pool with work distribution
  - `AsyncUnique()` - Task deduplication in concurrent processing
  - Enhanced WaitGroup with additional features
  - Object pooling for performance optimization

### Specialized Packages

#### `cryptox`
- **Purpose**: Comprehensive cryptographic operations
- **Features**: 40 files, 11,254 lines of code
- **Capabilities**:
  - Symmetric encryption: AES, DES, Triple DES, ChaCha20
  - Asymmetric cryptography: RSA, ECDSA, ECDH
  - Hashing: SHA family, MD5, HMAC, BLAKE2, SHA3
  - Key derivation: PBKDF2, Scrypt, Argon2
- **Security**: Production-ready implementations following best practices

#### `hystrix`
- **Purpose**: High-performance circuit breaker pattern implementation
- **Features**: 4 files, 1,367 lines of code
- **Performance Optimizations**:
  - Atomic operations for state management
  - Lock-free algorithms
  - Memory alignment for CPU cache efficiency
  - Three variants: standard, fast, and batch-optimized
- **Benchmark**: Register operations at ~46ns/op with zero allocations

#### `xtime`
- **Purpose**: Extended time operations with Chinese calendar support
- **Features**: 21 files, 10,744 lines of code
- **Unique Features**:
  - Chinese lunar calendar calculations
  - Solar terms computation
  - Business time constants for work schedules
  - Subpackages for different work patterns (007, 955, 996)

### Configuration and I/O Packages

#### `config`
- **Purpose**: Multi-format configuration loading
- **Supported Formats**: JSON, YAML, TOML, INI, HCL
- **Features**: Environment-aware configuration loading with validation

#### `bufiox`
- **Purpose**: Buffered I/O operations and utilities
- **Features**: Custom scanning utilities for performance optimization

#### `osx`
- **Purpose**: Cross-platform OS interface and file operations
- **Features**: 9 files, 2,554 lines of code
- **Capabilities**: File system utilities with cross-platform compatibility

### Network and Communication

#### `network`
- **Purpose**: Network utilities and helpers
- **Features**: IP address utilities, interface detection, real IP extraction

### Random and Testing Utilities

#### `randx`
- **Purpose**: Extended random number and data generation
- **Features**: 9 files, 2,014 lines of code
- **Capabilities**:
  - Various probability distributions
  - Time-based random utilities
  - Performance-optimized random generators

#### `fake`
- **Purpose**: Fake data generation for testing
- **Features**: User agent generation, test data utilities

## 🔗 Dependency Graph

```
Root Package (utils)
├── json (core JSON operations)
├── candy (type conversion)
└── ... (minimal external dependencies)

Infrastructure Layer
├── runtime → app
├── routine → runtime, log
└── osx (OS abstraction)

Utility Layer
├── stringx (string manipulation)
├── anyx → candy, json
├── wait → routine, runtime
└── xtime (time operations)

Specialized Layer
├── cryptox (cryptographic operations)
├── hystrix → randx (circuit breaker)
├── config → json, osx, runtime
└── network (networking utilities)
```

## 🚀 Performance Characteristics

### Benchmarks (Apple M3)

| Package | Operation | Performance | Memory |
|---------|-----------|-------------|--------|
| atexit | Register | 46.69 ns/op | 43 B/op, 0 allocs/op |
| atexit | RegisterConcurrent | 43.81 ns/op | 44 B/op, 0 allocs/op |
| atexit | ExecuteCallbacks | 545.9 ns/op | 896 B/op, 1 allocs/op |

### Performance Features

1. **Lock-Free Operations**: Critical paths use atomic operations instead of mutexes
2. **Memory Alignment**: Structs aligned for optimal CPU cache performance
3. **Zero-Copy Operations**: String/byte conversions without memory allocation
4. **Object Pooling**: Reduces GC pressure in high-frequency operations
5. **Generic Optimizations**: Type-safe operations without runtime reflection

## 🧪 Testing and Quality

### Test Coverage by Package
- **candy**: 99.3%
- **anyx**: 99.0%
- **atexit**: 100.0%
- **bufiox**: 100.0%
- **cryptox**: 100.0%
- **defaults**: 100.0%
- **stringx**: 96.4%
- **osx**: 97.7%
- **config**: 95.7%
- **network**: 89.1%

### Quality Assurance
- Comprehensive unit tests with edge case coverage
- Benchmark tests for performance-critical operations
- Race condition testing for concurrent operations
- Memory leak testing for long-running operations

## 🌏 Cultural Features

### Chinese Calendar Support (xtime package)
- **Lunar Calendar**: Full implementation of Chinese lunar calendar
- **Solar Terms**: 24 traditional Chinese solar terms calculation
- **Work Schedules**: Support for Chinese work patterns (007, 955, 996)
- **Traditional Holidays**: Chinese traditional holiday calculations

## 🔮 Future Architecture Considerations

1. **Plugin System**: Consider implementing a plugin architecture for extensibility
2. **Observability**: Enhanced metrics and tracing integration
3. **Configuration**: Hot-reload configuration capabilities
4. **Caching**: Distributed caching layer for high-performance scenarios
5. **Streaming**: Enhanced streaming utilities for large data processing

## 📈 Scalability

The architecture is designed to scale both vertically and horizontally:

- **Vertical Scaling**: Optimized memory usage and CPU performance
- **Horizontal Scaling**: Thread-safe operations support concurrent usage
- **Microservices**: Each package can be used independently in different services
- **Cloud Native**: Compatible with container environments and cloud platforms

This architecture provides a solid foundation for building high-performance Go applications while maintaining code clarity and developer productivity.