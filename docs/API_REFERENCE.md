# lazygophers/utils API Reference

A comprehensive Go utility library with modular design, providing essential functionality for common development tasks. This library emphasizes type safety, performance optimization, and follows Go 1.24+ standards.

## Installation

```go
go get github.com/lazygophers/utils
```

## Table of Contents

- [Core Package (utils)](#core-package-utils)
- [Type Conversion (candy)](#type-conversion-candy)
- [String Manipulation (stringx)](#string-manipulation-stringx)
- [Map Operations (anyx)](#map-operations-anyx)
- [Concurrency Control (wait)](#concurrency-control-wait)
- [Circuit Breaker (hystrix)](#circuit-breaker-hystrix)
- [Configuration Management (config)](#configuration-management-config)
- [Cryptographic Operations (cryptox)](#cryptographic-operations-cryptox)
- [Time Extensions (xtime)](#time-extensions-xtime)
- [Random Number Generation (randx)](#random-number-generation-randx)
- [Network Utilities (network)](#network-utilities-network)

---

## Core Package (utils)

The root package provides fundamental utilities including error handling, database operations, and validation.

### Must Functions

Error handling utilities that panic on error - useful for initialization and critical operations.

```go
// Must combines error checking with value return
func Must[T any](value T, err error) T

// MustOk panics if ok is false
func MustOk[T any](value T, ok bool) T

// MustSuccess panics if error is not nil  
func MustSuccess(err error)

// Ignore explicitly ignores return values
func Ignore[T any](value T, _ any) T
```

**Usage Example:**
```go
import "github.com/lazygophers/utils"

// Parse configuration that must succeed
config := utils.Must(loadConfig("app.json"))

// Get map value that must exist
value := utils.MustOk(m["key"])

// Ensure operation succeeds
utils.MustSuccess(file.Close())
```

### Database Operations

JSON-based database field scanning and value conversion with automatic default population.

```go
// Scan deserializes database field to struct
func Scan(src interface{}, dst interface{}) error

// Value serializes struct for database storage
func Value(m interface{}) (driver.Value, error)
```

**Usage Example:**
```go
type User struct {
    Name string `json:"name" default:"anonymous"`
    Age  int    `json:"age" default:"18"`
}

var user User
// Scan from database JSON field
err := utils.Scan(jsonBytes, &user)

// Convert to database value
value, err := utils.Value(user)
```

### Validation

Struct validation using go-playground/validator with automatic error logging.

```go
// Validate validates struct fields using tags
func Validate(m interface{}) error
```

**Usage Example:**
```go
type Request struct {
    Email string `validate:"required,email"`
    Age   int    `validate:"min=18,max=120"`
}

req := Request{Email: "user@example.com", Age: 25}
err := utils.Validate(req) // Returns nil if valid
```

---

## Type Conversion (candy)

High-performance type conversion utilities with comprehensive support for Go's type system.

### Basic Type Conversions

```go
// Convert any type to int64
func ToInt64(val interface{}) int64

// Convert any type to string  
func String[T constraints.Ordered](s T) string

// Convert to pointer
func ToPtr[T any](val T) *T

// Convert to boolean
func ToBool(val interface{}) bool

// Convert to bytes
func ToBytes(val interface{}) []byte
```

**Supported Input Types:**
- All numeric types (int, int8-64, uint, uint8-64, float32/64)
- Boolean values (true=1, false=0)
- Strings and []byte (parsed as numbers)
- time.Duration
- Returns 0 for unsupported types

**Usage Example:**
```go
import "github.com/lazygophers/utils/candy"

// Type conversions
num := candy.ToInt64("123")        // 123
str := candy.String(42)            // "42"
ptr := candy.ToPtr("hello")        // *string pointing to "hello"
flag := candy.ToBool("true")       // true
bytes := candy.ToBytes("test")     // []byte("test")
```

### Collection Operations

```go
// Get first element or default
func First[T any](slice []T) T
func FirstOr[T any](slice []T, defaultValue T) T

// Get last element or default  
func Last[T any](slice []T) T
func LastOr[T any](slice []T, defaultValue T) T

// Check if slice contains element
func Contains[T comparable](slice []T, item T) bool

// Filter slice elements
func Filter[T any](slice []T, predicate func(T) bool) []T

// Transform slice elements
func Map[T, R any](slice []T, mapper func(T) R) []R

// Remove duplicates
func Unique[T comparable](slice []T) []T

// Reverse slice
func Reverse[T any](slice []T) []T

// Chunk slice into smaller slices
func Chunk[T any](slice []T, size int) [][]T
```

**Performance Notes:**
- `Contains` uses optimized loops for primitive types
- `Filter` and `Map` pre-allocate result slices for better performance
- `Unique` uses map-based deduplication

**Usage Example:**
```go
slice := []int{1, 2, 3, 2, 4, 1}

first := candy.First(slice)                    // 1
last := candy.LastOr(slice, 0)                // 1
hasTwo := candy.Contains(slice, 2)             // true
evens := candy.Filter(slice, func(x int) bool { return x%2 == 0 }) // [2, 2, 4]
doubled := candy.Map(slice, func(x int) int { return x * 2 })      // [2, 4, 6, 4, 8, 2]
unique := candy.Unique(slice)                  // [1, 2, 3, 4]
reversed := candy.Reverse(slice)               // [1, 4, 2, 3, 2, 1]
chunks := candy.Chunk(slice, 2)                // [[1, 2], [3, 2], [4, 1]]
```

### Mathematical Operations

```go
// Sum numeric slice
func Sum[T Number](slice []T) T

// Find minimum/maximum
func Min[T constraints.Ordered](slice []T) T
func Max[T constraints.Ordered](slice []T) T

// Calculate average
func Average[T Number](slice []T) float64

// Mathematical functions
func Abs[T Number](x T) T
func Pow[T Number](base, exp T) T
func Sqrt[T Number](x T) float64
func Cbrt[T Number](x T) float64
```

**Usage Example:**
```go
numbers := []int{1, 2, 3, 4, 5}

total := candy.Sum(numbers)        // 15
min := candy.Min(numbers)          // 1  
max := candy.Max(numbers)          // 5
avg := candy.Average(numbers)      // 3.0

absVal := candy.Abs(-42)           // 42
power := candy.Pow(2, 3)           // 8
squareRoot := candy.Sqrt(16)       // 4.0
cubeRoot := candy.Cbrt(27)         // 3.0
```

---

## String Manipulation (stringx)

High-performance string manipulation with memory and CPU optimizations.

### Case Conversion

Memory-optimized string case conversions with ASCII fast-path optimization.

```go
// Convert camelCase to snake_case (optimized)
func Camel2Snake(s string) string

// Convert snake_case to CamelCase
func Snake2Camel(s string) string

// Convert snake_case to camelCase  
func Snake2SmallCamel(s string) string

// Convert to snake_case (general purpose)
func ToSnake(s string) string

// Convert to kebab-case
func ToKebab(s string) string

// Convert to PascalCase
func ToCamel(s string) string

// Convert to camelCase
func ToSmallCamel(s string) string

// Convert to slash/separated
func ToSlash(s string) string

// Convert to dot.separated
func ToDot(s string) string
```

**Performance Features:**
- ASCII fast-path for maximum performance
- Memory pre-allocation to avoid repeated allocations
- Zero-copy string conversion using unsafe operations
- Specialized optimizations for pure ASCII strings

**Usage Example:**
```go
import "github.com/lazygophers/utils/stringx"

input := "UserProfileData"

snake := stringx.Camel2Snake(input)      // "user_profile_data"
kebab := stringx.ToKebab(input)          // "user-profile-data"
camel := stringx.ToCamel("user_name")    // "UserName"
smallCamel := stringx.ToSmallCamel("user_name") // "userName"
slash := stringx.ToSlash(input)          // "user/profile/data"
dot := stringx.ToDot(input)              // "user.profile.data"
```

### String Utilities

```go
// Zero-copy bytes to string conversion (unsafe)
func ToString(b []byte) string

// Zero-copy string to bytes conversion (unsafe)  
func ToBytes(s string) []byte

// Split string by length
func SplitLen(s string, max int) []string

// Truncate string
func Shorten(s string, max int) string

// Truncate with ellipsis
func ShortenShow(s string, max int) string

// Reverse string (ASCII optimized)
func Reverse(s string) string

// Quote string
func Quote(s string) string
func QuotePure(s string) string
```

**Safety Notes:**
- `ToString` and `ToBytes` use unsafe operations for zero-copy conversion
- Only use when you control the lifecycle of the underlying data
- For safety-critical code, use standard library conversions

**Usage Example:**
```go
// Zero-copy conversions (use with caution)
bytes := []byte("hello")
str := stringx.ToString(bytes)           // "hello" (zero-copy)
backToBytes := stringx.ToBytes(str)      // []byte("hello") (zero-copy)

// String manipulation
text := "Hello, World!"
parts := stringx.SplitLen(text, 5)       // ["Hello", ", Wor", "ld!"]
short := stringx.Shorten(text, 5)        // "Hello"
shortShow := stringx.ShortenShow(text, 8) // "Hello..."
reversed := stringx.Reverse(text)        // "!dlroW ,olleH"
quoted := stringx.Quote(text)            // "\"Hello, World!\""
```

### Type Checking

```go
// Check if string/runes are uppercase
func IsUpper[M string | []rune](r M) bool

// Check if string/runes are digits
func IsDigit[M string | []rune](r M) bool
```

**Usage Example:**
```go
stringx.IsUpper("HELLO")     // true
stringx.IsUpper("Hello")     // false
stringx.IsDigit("12345")     // true
stringx.IsDigit("123a5")     // false
```

---

## Map Operations (anyx)

Comprehensive map and slice manipulation utilities with reflection-based operations.

### Map Key/Value Extraction

```go
// Extract map keys by type
func MapKeysString(m interface{}) []string
func MapKeysInt(m interface{}) []int
func MapKeysInt32(m interface{}) []int32
func MapKeysInt64(m interface{}) []int64
func MapKeysUint32(m interface{}) []uint32
func MapKeysUint64(m interface{}) []uint64
func MapKeysFloat32(m interface{}) []float32
func MapKeysFloat64(m interface{}) []float64

// Extract map values  
func MapValues[K constraints.Ordered, V any](m map[K]V) []V
func MapValuesAny(m interface{}) []interface{}
func MapValuesString(m interface{}) []string
func MapValuesInt(m interface{}) []int
func MapValuesFloat64(m interface{}) []float64
```

**Usage Example:**
```go
import "github.com/lazygophers/utils/anyx"

userMap := map[string]int{"alice": 25, "bob": 30}

keys := anyx.MapKeysString(userMap)        // ["alice", "bob"]
values := anyx.MapValues(userMap)          // [25, 30]

// Works with interface{} maps too
genericMap := map[interface{}]interface{}{
    "key1": "value1",
    42: "value2",
}
allKeys := anyx.MapKeysAny(genericMap)     // ["key1", 42]
```

### Map Merging and Transformation

```go
// Merge two maps (source + target)
func MergeMap[K constraints.Ordered, V any](source, target map[K]V) map[K]V

// Convert slice to map[item]bool
func Slice2Map[M constraints.Ordered](list []M) map[M]bool
```

**Usage Example:**
```go
map1 := map[string]int{"a": 1, "b": 2}
map2 := map[string]int{"b": 3, "c": 4}

merged := anyx.MergeMap(map1, map2)        // {"a": 1, "b": 3, "c": 4}

slice := []string{"apple", "banana", "apple"}
setMap := anyx.Slice2Map(slice)            // {"apple": true, "banana": true}
```

### Struct to Map Conversion

```go
// Convert slice to map by field name
func KeyBy(list interface{}, fieldName string) interface{}
func KeyByString[M any](list []*M, fieldName string) map[string]*M
func KeyByInt64[M any](list []*M, fieldName string) map[int64]*M
func KeyByUint64[M any](list []*M, fieldName string) map[uint64]*M
func KeyByInt32[M any](list []*M, fieldName string) map[int32]*M
```

**Usage Example:**
```go
type User struct {
    ID   int64  `json:"id"`
    Name string `json:"name"`
}

users := []*User{
    {ID: 1, Name: "Alice"},
    {ID: 2, Name: "Bob"},
}

// Key by ID field
usersByID := anyx.KeyByInt64(users, "ID")
// Result: map[int64]*User{1: &User{1, "Alice"}, 2: &User{2, "Bob"}}

usersByName := anyx.KeyByString(users, "Name")  
// Result: map[string]*User{"Alice": &User{1, "Alice"}, "Bob": &User{2, "Bob"}}
```

### Value Type Detection

```go
type ValueType int

const (
    ValueUnknown ValueType = iota
    ValueNumber
    ValueString  
    ValueBool
)

// Detect value type
func CheckValueType(val interface{}) ValueType
```

**Usage Example:**
```go
anyx.CheckValueType(42)        // ValueNumber
anyx.CheckValueType("hello")   // ValueString
anyx.CheckValueType(true)      // ValueBool
anyx.CheckValueType(struct{}{}) // ValueUnknown
```

---

## Concurrency Control (wait)

Advanced concurrency utilities including semaphore pools and async task processing.

### Semaphore Pool Management

Named semaphore pools for controlling concurrent access to resources.

```go
// Initialize semaphore pool
func Ready(key string, max int)

// Acquire/release semaphore
func Lock(key string)
func Unlock(key string)

// Get current semaphore usage
func Depth(key string) int

// Execute function with automatic semaphore management
func Sync(key string, logic func() error) error
```

**Usage Example:**
```go
import "github.com/lazygophers/utils/wait"

// Setup semaphore pool for database connections
wait.Ready("db_pool", 10)

// Manual semaphore management
wait.Lock("db_pool")
// ... database operation
wait.Unlock("db_pool")

// Automatic semaphore management
err := wait.Sync("db_pool", func() error {
    // Database operation here
    return db.Query("SELECT * FROM users")
})
```

### Async Task Processing

High-performance async task processing with goroutine pools and object reuse.

```go
// Process tasks with goroutine pool
func Async[M any](process int, push func(chan M), logic func(M))

// Process tasks with persistent goroutines  
func AsyncAlwaysWithChan[M any](process int, c chan M, logic func(M))

// Task interface for uniqueness checking
type UniqueTask interface {
    UniqueKey() string
}

// Process unique tasks (prevents duplicate execution)
func AsyncUnique[M UniqueTask](process int, push func(chan M), logic func(M))
func AsyncAlwaysUnique[M UniqueTask](process int, logic func(M)) chan M
```

**Performance Features:**
- WaitGroup object pooling to reduce GC pressure
- Buffered channels sized for optimal throughput
- Unique task processing prevents duplicate work
- Automatic goroutine lifecycle management

**Usage Example:**
```go
// Basic async processing
tasks := []int{1, 2, 3, 4, 5}

wait.Async(3, func(ch chan int) {
    for _, task := range tasks {
        ch <- task
    }
}, func(task int) {
    // Process task
    fmt.Printf("Processing task %d\n", task)
})

// Unique task processing
type ProcessTask struct {
    ID   string
    Data string
}

func (t ProcessTask) UniqueKey() string {
    return t.ID
}

taskChan := wait.AsyncAlwaysUnique(2, func(task ProcessTask) {
    // Process unique task
    fmt.Printf("Processing %s: %s\n", task.ID, task.Data)
})

// Send tasks (duplicates will be ignored)
taskChan <- ProcessTask{ID: "task1", Data: "data1"}
taskChan <- ProcessTask{ID: "task1", Data: "data2"} // Ignored
```

### Advanced Pool Management

```go
// Pool represents a semaphore pool
type Pool struct {
    // Internal implementation
}

// Pool methods
func (p *Pool) Lock()
func (p *Pool) Unlock()  
func (p *Pool) Depth() int
```

**Usage Example:**
```go
// Direct pool usage (advanced)
wait.Ready("api_calls", 5)
pool := wait.getPool("api_calls") // Internal function

// Custom pool management
pool.Lock()
defer pool.Unlock()

fmt.Printf("Current depth: %d\n", pool.Depth())
```

---

## Circuit Breaker (hystrix)

High-performance circuit breaker implementation with multiple optimization levels.

### Standard Circuit Breaker

Feature-rich circuit breaker with configurable behavior and state management.

```go
type CircuitBreakerConfig struct {
    TimeWindow    time.Duration // Statistics time window
    OnStateChange StateChange   // State change callback
    ReadyToTrip   ReadyToTrip   // Circuit break condition
    Probe         Probe         // Half-open state probe
    BufferSize    int          // Request result buffer size (default: 1000)
}

// Create circuit breaker
func NewCircuitBreaker(config CircuitBreakerConfig) *CircuitBreaker

// Circuit breaker states
type State string
const (
    Closed   State = "closed"    // Service available
    Open     State = "open"      // Service unavailable  
    HalfOpen State = "half-open" // Testing service recovery
)
```

**Core Methods:**
```go
// Check if request should be allowed
func (cb *CircuitBreaker) Before() bool

// Record request result
func (cb *CircuitBreaker) After(success bool)

// Execute function with circuit breaker protection
func (cb *CircuitBreaker) Call(fn func() error) error

// Get current state and statistics
func (cb *CircuitBreaker) State() State
func (cb *CircuitBreaker) Stat() (successes, failures uint64)
```

**Performance Features:**
- Lock-free atomic operations for counters
- Memory-aligned structures to prevent false sharing
- Optimized ring buffer with packed data storage
- Fast state transitions with CAS operations

**Usage Example:**
```go
import "github.com/lazygophers/utils/hystrix"

// Configure circuit breaker
config := hystrix.CircuitBreakerConfig{
    TimeWindow: 10 * time.Second,
    OnStateChange: func(oldState, newState hystrix.State) {
        log.Printf("Circuit breaker state: %s -> %s", oldState, newState)
    },
    ReadyToTrip: func(successes, failures uint64) bool {
        total := successes + failures
        return total >= 10 && failures > successes
    },
}

cb := hystrix.NewCircuitBreaker(config)

// Use circuit breaker
err := cb.Call(func() error {
    // Call external service
    return externalAPI.Call()
})

if err != nil {
    log.Printf("Service call failed: %v", err)
}

// Manual usage
if cb.Before() {
    err := externalService()
    cb.After(err == nil)
}
```

### Fast Circuit Breaker

Ultra-lightweight circuit breaker for high-performance scenarios.

```go
// Create fast circuit breaker
func NewFastCircuitBreaker(failureThreshold uint64, timeWindow time.Duration) *FastCircuitBreaker

// Fast circuit breaker methods
func (cb *FastCircuitBreaker) AllowRequest() bool
func (cb *FastCircuitBreaker) RecordResult(success bool)
func (cb *FastCircuitBreaker) CallFast(fn func() error) error
```

**Usage Example:**
```go
// Ultra-fast circuit breaker
fastCB := hystrix.NewFastCircuitBreaker(5, 30*time.Second)

// Fastest possible usage
if fastCB.AllowRequest() {
    err := doWork()
    fastCB.RecordResult(err == nil)
}

// Or use CallFast wrapper
err := fastCB.CallFast(func() error {
    return doWork()
})
```

### Batch Circuit Breaker

Optimized for high-throughput scenarios with batch result processing.

```go
// Create batch circuit breaker
func NewBatchCircuitBreaker(config CircuitBreakerConfig, batchSize int, batchTimeout time.Duration) *BatchCircuitBreaker

// Batch methods
func (cb *BatchCircuitBreaker) AfterBatch(success bool)
```

**Usage Example:**
```go
// High-throughput batch processing
batchCB := hystrix.NewBatchCircuitBreaker(config, 100, 100*time.Millisecond)

// Process many results efficiently
for i := 0; i < 1000; i++ {
    success := processItem(i)
    batchCB.AfterBatch(success) // Batched internally
}
```

### Helper Functions

```go
// Probe functions for half-open state
func ProbeWithChance(percentage int) Probe

// Common ready-to-trip functions
func DefaultReadyToTrip(successes, failures uint64) bool
```

---

## Configuration Management (config)

Multi-format configuration loading with automatic file discovery and validation.

### Supported Formats

The library supports multiple configuration formats:
- **JSON/JSON5**: Standard and extended JSON
- **YAML/YML**: YAML configuration files  
- **TOML**: Tom's Obvious Minimal Language
- **INI**: Traditional INI files
- **XML**: XML configuration
- **Properties**: Java-style properties files
- **ENV**: Environment variable files
- **HCL**: HashiCorp Configuration Language

### Core Functions

```go
// Load and validate configuration
func LoadConfig(c any, paths ...string) error

// Load without validation
func LoadConfigSkipValidate(c any, paths ...string) error

// Save configuration to file
func SetConfig(c any) error

// Register custom format parser
func RegisterParser(ext string, m Marshaler, u Unmarshaler) error
```

### Automatic Discovery

Configuration files are automatically discovered in this order:
1. Explicitly provided paths
2. Environment variable `LAZYGOPHERS_CONFIG`
3. Current working directory (`conf.*` or `config.*`)
4. Executable directory (`conf.*` or `config.*`)

**Usage Example:**
```go
import "github.com/lazygophers/utils/config"

type AppConfig struct {
    Server struct {
        Host string `json:"host" yaml:"host" validate:"required"`
        Port int    `json:"port" yaml:"port" validate:"min=1,max=65535"`
    } `json:"server" yaml:"server"`
    
    Database struct {
        URL      string `json:"url" yaml:"url" validate:"required"`
        MaxConns int    `json:"max_conns" yaml:"max_conns" default:"10"`
    } `json:"database" yaml:"database"`
}

var cfg AppConfig

// Load with automatic discovery and validation
err := config.LoadConfig(&cfg)
if err != nil {
    log.Fatal(err)
}

// Load specific file without validation  
err = config.LoadConfigSkipValidate(&cfg, "app.yaml", "app.json")

// Save configuration
err = config.SetConfig(&cfg)
```

### Format-Specific Features

**JSON5 Support:**
- Comments in JSON files
- Trailing commas
- Unquoted keys
- Multi-line strings

**Properties/ENV Support:**
- Nested structure support with dot notation
- Environment variable substitution
- Comment lines (# or !)
- Automatic quote handling

**HCL Support:**
- Terraform-compatible syntax
- Nested blocks
- String interpolation
- Type inference

**Custom Format Registration:**
```go
// Register custom parser
config.RegisterParser(".custom", 
    func(writer io.Writer, v interface{}) error {
        // Marshal implementation
        return nil
    },
    func(reader io.Reader, v interface{}) error {
        // Unmarshal implementation  
        return nil
    },
)
```

### Field Tag Priority

Field mapping uses tags in this priority order:
1. `properties` tag
2. `env` tag  
3. `json` tag
4. `yaml` tag
5. `toml` tag
6. `ini` tag
7. Lowercase field name

**Example Configuration:**
```yaml
# config.yaml
server:
  host: "localhost"
  port: 8080
  
database:
  url: "postgres://user:pass@localhost/db"
  max_conns: 20
```

```json
{
  "server": {
    "host": "localhost", 
    "port": 8080
  },
  "database": {
    "url": "postgres://user:pass@localhost/db",
    "max_conns": 20
  }
}
```

---

## Cryptographic Operations (cryptox)

Comprehensive cryptographic utilities with multiple encryption algorithms and security features.

### AES Encryption

AES-256 encryption with multiple operation modes.

```go
// AES-GCM (Recommended - provides authentication)
func Encrypt(key, plaintext []byte) ([]byte, error)
func Decrypt(key, ciphertext []byte) ([]byte, error)

// AES-CBC (Cipher Block Chaining)
func EncryptCBC(key, plaintext []byte) ([]byte, error)  
func DecryptCBC(key, ciphertext []byte) ([]byte, error)

// AES-CFB (Cipher Feedback)
func EncryptCFB(key, plaintext []byte) ([]byte, error)
func DecryptCFB(key, ciphertext []byte) ([]byte, error)

// AES-CTR (Counter Mode)
func EncryptCTR(key, plaintext []byte) ([]byte, error)
func DecryptCTR(key, ciphertext []byte) ([]byte, error)

// AES-OFB (Output Feedback)
func EncryptOFB(key, plaintext []byte) ([]byte, error)
func DecryptOFB(key, ciphertext []byte) ([]byte, error)

// AES-ECB (NOT RECOMMENDED - included for compatibility)
func EncryptECB(key, plaintext []byte) ([]byte, error)
func DecryptECB(key, ciphertext []byte) ([]byte, error)
```

**Security Notes:**
- All functions require 32-byte (256-bit) keys
- GCM mode provides both encryption and authentication
- ECB mode is cryptographically unsafe for most uses
- Random IVs/nonces are automatically generated

**Usage Example:**
```go
import "github.com/lazygophers/utils/cryptox"

// Generate 32-byte key (in practice, use secure key derivation)
key := make([]byte, 32)
rand.Read(key)

plaintext := []byte("Hello, World!")

// AES-GCM (recommended)
ciphertext, err := cryptox.Encrypt(key, plaintext)
if err != nil {
    log.Fatal(err)
}

decrypted, err := cryptox.Decrypt(key, ciphertext)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Decrypted: %s\n", decrypted) // "Hello, World!"

// AES-CBC
cbcCiphertext, err := cryptox.EncryptCBC(key, plaintext)
cbcDecrypted, err := cryptox.DecryptCBC(key, cbcCiphertext)
```

### Hash Functions

Comprehensive hash function implementations.

```go
// Basic hash functions
func MD5(data []byte) []byte
func SHA1(data []byte) []byte
func SHA256(data []byte) []byte
func SHA512(data []byte) []byte

// SHA-3 family
func SHA3_224(data []byte) []byte
func SHA3_256(data []byte) []byte
func SHA3_384(data []byte) []byte  
func SHA3_512(data []byte) []byte

// BLAKE2 family
func BLAKE2b_256(data []byte) []byte
func BLAKE2b_512(data []byte) []byte
func BLAKE2s_256(data []byte) []byte

// CRC functions
func CRC32(data []byte) uint32
func CRC64(data []byte) uint64

// FNV hash functions
func FNV32(data []byte) uint32
func FNV32a(data []byte) uint32
func FNV64(data []byte) uint64
func FNV64a(data []byte) uint64
```

**Usage Example:**
```go
data := []byte("Hello, World!")

// Basic hashes
md5Hash := cryptox.MD5(data)
sha256Hash := cryptox.SHA256(data)
sha512Hash := cryptox.SHA512(data)

// Modern hashes (recommended)
sha3Hash := cryptox.SHA3_256(data)
blakeHash := cryptox.BLAKE2b_256(data)

// Fast non-cryptographic hashes
crc32Val := cryptox.CRC32(data)
fnvHash := cryptox.FNV64a(data)

fmt.Printf("SHA256: %x\n", sha256Hash)
```

### HMAC Functions

Message authentication codes for verifying data integrity and authenticity.

```go
// HMAC with various hash functions
func HMAC_MD5(key, data []byte) []byte
func HMAC_SHA1(key, data []byte) []byte
func HMAC_SHA256(key, data []byte) []byte
func HMAC_SHA512(key, data []byte) []byte
func HMAC_SHA3_256(key, data []byte) []byte
func HMAC_SHA3_512(key, data []byte) []byte
```

**Usage Example:**
```go
key := []byte("secret-key")
message := []byte("important message")

// Generate HMAC
mac := cryptox.HMAC_SHA256(key, message)

// Verify HMAC (in practice, use constant-time comparison)
expectedMAC := cryptox.HMAC_SHA256(key, message)
valid := bytes.Equal(mac, expectedMAC)
```

### UUID Generation

```go
// Generate UUIDs
func UUID() string           // UUID v4 (random)
func UUIDShort() string     // Short UUID
func UUIDTimeBased() string // UUID v1 (time-based)
```

**Usage Example:**
```go
uuid := cryptox.UUID()              // "f47ac10b-58cc-4372-a567-0e02b2c3d479"
shortUUID := cryptox.UUIDShort()    // "f47ac10b58cc4372"
timeUUID := cryptox.UUIDTimeBased() // Time-based UUID
```

### Advanced Features

**Key Derivation:**
```go
// PBKDF2 key derivation
func PBKDF2(password, salt []byte, iter, keyLen int) []byte

// scrypt key derivation  
func Scrypt(password, salt []byte, N, r, p, keyLen int) ([]byte, error)

// Argon2 key derivation
func Argon2i(password, salt []byte, time, memory uint32, threads uint8, keyLen uint32) []byte
func Argon2id(password, salt []byte, time, memory uint32, threads uint8, keyLen uint32) []byte
```

**RSA Operations:**
```go
// Generate RSA key pair
func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error)

// RSA encryption/decryption
func RSAEncrypt(publicKey *rsa.PublicKey, plaintext []byte) ([]byte, error)
func RSADecrypt(privateKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error)

// RSA signing/verification
func RSASign(privateKey *rsa.PrivateKey, message []byte) ([]byte, error)
func RSAVerify(publicKey *rsa.PublicKey, message, signature []byte) error
```

---

## Time Extensions (xtime)

Extended time handling with business time calculations and specialized time constants.

### Time Constants

Business-oriented time duration constants for common scenarios.

```go
const (
    // Standard time units
    Nanosecond  = time.Nanosecond
    Microsecond = time.Microsecond  
    Millisecond = time.Millisecond
    Second      = time.Second
    Minute      = time.Minute
    Hour        = time.Hour
    
    // Extended time units
    HalfHour = time.Minute * 30
    HalfDay  = time.Hour * 12
    Day      = time.Hour * 24
    
    // Business time units
    WorkDayWeek  = Day * 5           // Monday-Friday
    ResetDayWeek = Day * 2           // Saturday-Sunday  
    Week         = Day * 7
    
    WorkDayMonth  = Day*21 + HalfDay // ~21.5 working days
    ResetDayMonth = Day*8 + HalfDay  // ~8.5 weekend days
    Month         = Day * 30
    
    QUARTER = Day * 91               // ~3 months
    Year    = Day * 365
    Decade  = Year*10 + Day*2        // Accounts for leap years
    Century = Year*100 + Day*25      // Accounts for leap years
)
```

**Usage Example:**
```go
import "github.com/lazygophers/utils/xtime"

// Business time calculations
workingHours := 8 * xtime.Hour
lunchBreak := xtime.HalfHour
fullWorkDay := workingHours + lunchBreak

// Project timelines
projectDuration := 2 * xtime.WorkDayWeek  // 2 working weeks
vacationTime := xtime.ResetDayWeek        // Weekend

// Long-term planning
quarterDeadline := time.Now().Add(xtime.QUARTER)
yearlyReview := time.Now().Add(xtime.Year)

fmt.Printf("Project will take: %v\n", projectDuration)
fmt.Printf("Quarter ends: %v\n", quarterDeadline)
```

### Time Utilities

```go
// Current time helpers
func Now() time.Time
func NowUnix() int64
func NowUnixMilli() int64
func NowUnixNano() int64

// Time formatting
func FormatTime(t time.Time, layout string) string
func ParseTime(value, layout string) (time.Time, error)

// Business day calculations  
func IsWorkday(t time.Time) bool
func NextWorkday(t time.Time) time.Time
func PrevWorkday(t time.Time) time.Time
func WorkdaysBetween(start, end time.Time) int
```

**Usage Example:**
```go
now := xtime.Now()
unixTime := xtime.NowUnix()
milliTime := xtime.NowUnixMilli()

// Business day logic
if xtime.IsWorkday(now) {
    fmt.Println("It's a working day")
}

nextWork := xtime.NextWorkday(now)
workdays := xtime.WorkdaysBetween(now, nextWork.Add(xtime.Week))

fmt.Printf("Next workday: %v\n", nextWork)
fmt.Printf("Workdays in next week: %d\n", workdays)
```

### Lunar Calendar Support

Traditional Chinese lunar calendar with solar terms and festivals.

```go
// Lunar calendar conversion
func SolarToLunar(year, month, day int) (lunarYear, lunarMonth, lunarDay int, isLeap bool)
func LunarToSolar(lunarYear, lunarMonth, lunarDay int, isLeap bool) (year, month, day int)

// Solar terms (24 traditional Chinese seasons)
func GetSolarTerm(year, index int) time.Time
func GetSolarTermName(index int) string
func GetCurrentSolarTerm(t time.Time) (index int, name string)

// Traditional festivals
func GetTraditionalFestivals(year int) map[string]time.Time
```

**Usage Example:**
```go
// Convert solar date to lunar
lunarYear, lunarMonth, lunarDay, isLeap := xtime.SolarToLunar(2024, 2, 10)
fmt.Printf("Lunar date: %d/%d/%d (leap: %v)\n", lunarYear, lunarMonth, lunarDay, isLeap)

// Get solar terms
springStart := xtime.GetSolarTerm(2024, 3) // 立春 (Start of Spring)
termName := xtime.GetSolarTermName(3)
fmt.Printf("%s starts at: %v\n", termName, springStart)

// Current solar term
now := time.Now()
termIndex, termName := xtime.GetCurrentSolarTerm(now)
fmt.Printf("Current solar term: %s (index %d)\n", termName, termIndex)

// Traditional festivals
festivals := xtime.GetTraditionalFestivals(2024)
for name, date := range festivals {
    fmt.Printf("%s: %v\n", name, date)
}
```

---

## Random Number Generation (randx)

High-performance random number generation with thread-safe pools and batch operations.

### Performance Features

- **Thread-safe pool**: Avoids lock contention with sync.Pool
- **Batch operations**: Generate multiple random numbers efficiently  
- **Fast variants**: Global locked generator for single-threaded use
- **Memory optimization**: Object pooling for random generators

### Basic Random Numbers

```go
// Integer random numbers
func Intn(n int) int                    // [0, n)
func IntnRange(min, max int) int        // [min, max]
func Int() int                          // Full range
func Int64() int64                      // Full range
func Int64n(n int64) int64             // [0, n)
func Int64nRange(min, max int64) int64  // [min, max]

// Unsigned integers
func Uint32() uint32
func Uint32Range(min, max uint32) uint32
func Uint64() uint64  
func Uint64Range(min, max uint64) uint64

// Floating point
func Float32() float32                           // [0.0, 1.0)
func Float32Range(min, max float32) float32     // [min, max]
func Float64() float64                          // [0.0, 1.0)
func Float64Range(min, max float64) float64     // [min, max]
```

**Usage Example:**
```go
import "github.com/lazygophers/utils/randx"

// Basic random numbers
dice := randx.IntnRange(1, 6)           // Dice roll [1, 6]
percent := randx.Intn(100)              // Percentage [0, 99]
id := randx.Int64()                     // Random ID

// Range examples
temperature := randx.Float64Range(-10.0, 35.0)  // Temperature
delay := randx.IntnRange(100, 500)               // Delay in ms

fmt.Printf("Dice: %d, Temp: %.1f°C, Delay: %dms\n", dice, temperature, delay)
```

### High-Performance Variants

```go
// Fast variants (global locked generator)
func FastIntn(n int) int
func FastInt() int
func FastFloat64() float64
```

**Usage Example:**
```go
// Use fast variants for single-threaded scenarios
for i := 0; i < 1000; i++ {
    value := randx.FastIntn(100)  // Slightly faster for single thread
    process(value)
}
```

### Batch Operations

Generate multiple random numbers efficiently by reusing the same generator.

```go
// Batch generation functions
func BatchIntn(n int, count int) []int
func BatchInt64n(n int64, count int) []int64  
func BatchFloat64(count int) []float64
```

**Performance Benefits:**
- Single generator acquisition from pool
- Reduced sync.Pool overhead for large batches
- Memory-friendly slice pre-allocation

**Usage Example:**
```go
// Generate 1000 random numbers efficiently
randomNumbers := randx.BatchIntn(100, 1000)
float64Numbers := randx.BatchFloat64(500)
int64Numbers := randx.BatchInt64n(1000000, 250)

// Process batch
for _, num := range randomNumbers {
    process(num)
}
```

### Random Utilities

```go
// Boolean random
func Bool() bool
func BoolWithChance(percentage int) bool

// Random choice from slice
func Choice[T any](slice []T) T
func ChoiceMultiple[T any](slice []T, count int) []T

// Random time
func RandomTime(start, end time.Time) time.Time
func RandomDuration(min, max time.Duration) time.Duration
```

**Usage Example:**
```go
// Boolean random
coinFlip := randx.Bool()                    // 50% chance
rare := randx.BoolWithChance(5)             // 5% chance

// Random selection
colors := []string{"red", "green", "blue", "yellow"}
randomColor := randx.Choice(colors)
threeColors := randx.ChoiceMultiple(colors, 3)

// Random time
start := time.Now()
end := start.Add(24 * time.Hour)
randomMoment := randx.RandomTime(start, end)

randomDelay := randx.RandomDuration(100*time.Millisecond, 2*time.Second)

fmt.Printf("Color: %s, Time: %v, Delay: %v\n", randomColor, randomMoment, randomDelay)
```

### Thread Safety

All functions are thread-safe and use efficient pool-based generators:

```go
// Safe for concurrent use
go func() {
    for i := 0; i < 1000; i++ {
        value := randx.Intn(100)
        // Process value
    }
}()

go func() {
    batch := randx.BatchFloat64(500)
    // Process batch
}()
```

### Performance Recommendations

1. **Use batch operations** for generating many random numbers
2. **Use Fast variants** for single-threaded scenarios
3. **Use regular functions** for concurrent access
4. **Avoid frequent range calls** - cache ranges when possible

```go
// Good: Batch generation
numbers := randx.BatchIntn(100, 1000)

// Good: Single-threaded fast path  
if singleThreaded {
    value := randx.FastIntn(100)
}

// Good: Concurrent safe
value := randx.Intn(100)

// Avoid: Many individual calls in tight loop
for i := 0; i < 1000; i++ {
    value := randx.IntnRange(1, 100) // Pool overhead per call
}
```

---

## Network Utilities (network)

Network-related utilities for IP address validation and network interface operations.

### IP Address Utilities

```go
// Check if IP address is local/private
func IsLocalIp(ip string) bool
```

**Local IP Detection:**
- Private IP ranges (RFC 1918)
- Loopback addresses (127.0.0.0/8, ::1)
- Link-local addresses (169.254.0.0/16, fe80::/10)

**Usage Example:**
```go
import "github.com/lazygophers/utils/network"

ips := []string{
    "192.168.1.1",    // Private
    "10.0.0.1",       // Private
    "172.16.0.1",     // Private
    "127.0.0.1",      // Loopback
    "8.8.8.8",        // Public
    "::1",            // IPv6 loopback
}

for _, ip := range ips {
    isLocal := network.IsLocalIp(ip)
    fmt.Printf("%s is local: %v\n", ip, isLocal)
}
// Output:
// 192.168.1.1 is local: true
// 10.0.0.1 is local: true  
// 172.16.0.1 is local: true
// 127.0.0.1 is local: true
// 8.8.8.8 is local: false
// ::1 is local: true
```

### Network Interface Operations

```go
// Get network interface information
func GetNetworkInterfaces() ([]NetworkInterface, error)
func GetActiveInterface() (*NetworkInterface, error)
func GetInterfaceByName(name string) (*NetworkInterface, error)

type NetworkInterface struct {
    Name      string
    IPs       []string
    MAC       string
    IsUp      bool
    IsLoopback bool
}
```

**Usage Example:**
```go
// Get all network interfaces
interfaces, err := network.GetNetworkInterfaces()
if err != nil {
    log.Fatal(err)
}

for _, iface := range interfaces {
    fmt.Printf("Interface: %s\n", iface.Name)
    fmt.Printf("  MAC: %s\n", iface.MAC)
    fmt.Printf("  IPs: %v\n", iface.IPs)
    fmt.Printf("  Up: %v, Loopback: %v\n", iface.IsUp, iface.IsLoopback)
}

// Get active interface (typically default route interface)
active, err := network.GetActiveInterface()
if err == nil {
    fmt.Printf("Active interface: %s (%v)\n", active.Name, active.IPs)
}
```

### HTTP Client Utilities

```go
// Create HTTP client with custom configuration
func NewHTTPClient(timeout time.Duration, options ...ClientOption) *http.Client

// Client configuration options
type ClientOption func(*http.Client)

func WithTimeout(timeout time.Duration) ClientOption
func WithTransport(transport *http.Transport) ClientOption  
func WithRetry(maxRetries int) ClientOption
```

**Usage Example:**
```go
// Create configured HTTP client
client := network.NewHTTPClient(
    30*time.Second,
    network.WithRetry(3),
    network.WithTransport(&http.Transport{
        MaxIdleConns:       10,
        IdleConnTimeout:    90 * time.Second,
        DisableCompression: true,
    }),
)

// Use client for requests
resp, err := client.Get("https://api.example.com/data")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()
```

---

## Best Practices and Performance Tips

### Memory Management

1. **Use batch operations** when generating multiple items:
   ```go
   // Good: Single allocation
   numbers := randx.BatchIntn(100, 1000)
   
   // Avoid: Multiple allocations
   for i := 0; i < 1000; i++ {
       randx.Intn(100)
   }
   ```

2. **Leverage object pools** for high-frequency operations:
   ```go
   // WaitGroup pooling is automatic in wait.Async
   wait.Async(10, pushFunc, processFunc)
   ```

3. **Use zero-copy conversions** carefully:
   ```go
   // Fast but unsafe - ensure data lifecycle control
   str := stringx.ToString(bytes)
   ```

### Error Handling

1. **Use Must functions** for initialization:
   ```go
   config := utils.Must(loadConfig("app.json"))
   ```

2. **Consistent error logging** is built-in:
   ```go
   // Errors are automatically logged before return
   err := utils.Validate(request)
   ```

### Concurrency

1. **Use semaphore pools** for resource control:
   ```go
   wait.Ready("db_pool", 10)
   wait.Sync("db_pool", func() error {
       return db.Query()
   })
   ```

2. **Circuit breakers** for external services:
   ```go
   cb := hystrix.NewCircuitBreaker(config)
   err := cb.Call(externalAPICall)
   ```

### Type Safety

1. **Leverage generics** for type-safe operations:
   ```go
   result := candy.Filter(numbers, func(x int) bool {
       return x > 0
   })
   ```

2. **Use typed map operations**:
   ```go
   userMap := anyx.KeyByString(users, "Name")
   ```

---

## Migration Guide

### From Standard Library

```go
// Standard library
strconv.Atoi(str)
// utils/candy
candy.ToInt64(str)

// Standard library  
json.Marshal(data)
// utils (with automatic defaults)
utils.Value(data)

// Standard library
rand.Intn(100)
// utils/randx (thread-safe, pooled)
randx.Intn(100)
```

### From Other Libraries

The utilities are designed to be drop-in replacements or enhancements for common operations while providing better performance and safety.

---

## Examples and Recipes

### Complete Application Setup

```go
package main

import (
    "log"
    "time"
    
    "github.com/lazygophers/utils"
    "github.com/lazygophers/utils/config"
    "github.com/lazygophers/utils/wait"
    "github.com/lazygophers/utils/hystrix"
)

type Config struct {
    Server struct {
        Host string `json:"host" validate:"required"`
        Port int    `json:"port" validate:"min=1,max=65535"`
    } `json:"server"`
    Database struct {
        MaxConns int `json:"max_conns" default:"10"`
    } `json:"database"`
}

func main() {
    // Load configuration with validation
    var cfg Config
    utils.MustSuccess(config.LoadConfig(&cfg))
    
    // Setup concurrency control
    wait.Ready("db_pool", cfg.Database.MaxConns)
    
    // Setup circuit breaker
    cb := hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
        TimeWindow: 10 * time.Second,
    })
    
    // Application logic here
    log.Printf("Server starting on %s:%d", cfg.Server.Host, cfg.Server.Port)
}
```

### High-Performance Data Processing

```go
func processLargeDataset(data []Item) error {
    // Use async processing with unique task filtering
    wait.AsyncUnique(runtime.NumCPU(), 
        func(ch chan Item) {
            for _, item := range data {
                ch <- item
            }
        },
        func(item Item) {
            // Process item with circuit breaker protection
            cb.Call(func() error {
                return processItem(item)
            })
        },
    )
    return nil
}
```

This completes the comprehensive API reference for the lazygophers/utils library. Each package provides production-ready utilities with a focus on performance, safety, and ease of use.