# API Reference

<!-- Language selector -->
[🇺🇸 English](#english) | [🇨🇳 简体中文](#简体中文) | [🇭🇰 繁體中文](#繁體中文) | [🇷🇺 Русский](#русский) | [🇫🇷 Français](#français) | [🇸🇦 العربية](#العربية) | [🇪🇸 Español](#español)

---

## English

### Overview
This document provides a comprehensive API reference for all public functions, types, and interfaces in the LazyGophers Utils library, organized by package with detailed signatures, parameters, and usage examples.

### API Documentation Structure

#### Package Organization
```
github.com/lazygophers/utils/
├── core utilities (Must, Scan, Value, Validate)
├── candy/          # Type conversion utilities
├── json/           # JSON processing
├── anyx/           # Any type manipulation
├── routine/        # Goroutine management
├── wait/           # Synchronization utilities
├── singledo/       # Singleton execution
├── xtime/          # Enhanced time operations
├── cryptox/        # Cryptographic utilities
├── network/        # Network utilities
├── app/            # Application metadata
├── config/         # Configuration management
└── ...            # Additional packages
```

### Core Utilities API

#### Must Functions
```go
// Must returns the first argument if the error is nil, otherwise panics
func Must[T any](value T, err error) T

// MustSuccess panics if any of the provided errors is non-nil
func MustSuccess(errs ...error)

// MustOk panics if ok is false, returns the value otherwise
func MustOk[T any](value T, ok bool) T
```

**Examples:**
```go
// File reading with Must
content := utils.Must(os.ReadFile("config.json"))

// Multiple error checking
utils.MustSuccess(
    initDatabase(),
    startServer(),
    setupLogging(),
)

// Conditional value extraction
user := utils.MustOk(findUser(id))
```

#### Database Utilities
```go
// Scan deserializes JSON data into a struct with default values
func Scan(data []byte, v interface{}) error

// Value serializes a struct to JSON for database storage
func Value(v interface{}) ([]byte, error)
```

**Examples:**
```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age" default:"18"`
}

// Scanning from database
var user User
err := utils.Scan(jsonData, &user)

// Preparing for database storage
jsonData, err := utils.Value(user)
```

#### Validation
```go
// Validate performs struct validation using validator tags
func Validate(v interface{}) error
```

**Example:**
```go
type Config struct {
    Email string `validate:"required,email"`
    Port  int    `validate:"min=1,max=65535"`
}

config := Config{Email: "user@example.com", Port: 8080}
err := utils.Validate(&config)
```

### Candy Package API

#### Type Conversion Functions

```go
// String conversions
func ToString(v interface{}) string
func String[T constraints.Ordered](s T) string

// Integer conversions
func ToInt(v interface{}) int
func ToInt8(v interface{}) int8
func ToInt16(v interface{}) int16
func ToInt32(v interface{}) int32
func ToInt64(v interface{}) int64
func ToUint(v interface{}) uint
func ToUint8(v interface{}) uint8
func ToUint16(v interface{}) uint16
func ToUint32(v interface{}) uint32
func ToUint64(v interface{}) uint64

// Float conversions
func ToFloat32(v interface{}) float32
func ToFloat64(v interface{}) float64

// Boolean conversion
func ToBool(v interface{}) bool

// Byte slice conversion
func ToBytes(v interface{}) []byte
```

#### Collection Functions

```go
// Slice operations
func Contains[T comparable](slice []T, item T) bool
func Unique[T comparable](slice []T) []T
func Reverse[T any](slice []T) []T
func Chunk[T any](slice []T, size int) [][]T
func Filter[T any](slice []T, predicate func(T) bool) []T
func Map[T, U any](slice []T, mapper func(T) U) []U
func Reduce[T, U any](slice []T, initial U, reducer func(U, T) U) U

// Statistical functions
func Sum[T constraints.Ordered](slice []T) T
func Average[T constraints.Float](slice []T) T
func Min[T constraints.Ordered](slice []T) T
func Max[T constraints.Ordered](slice []T) T

// Mathematical functions
func Abs[T constraints.Signed | constraints.Float](x T) T
func Sqrt[T constraints.Float](x T) T
func Pow[T constraints.Integer | constraints.Float](base, exponent T) T
```

**Usage Examples:**
```go
// Type conversions
age := candy.ToInt("25")          // 25
price := candy.ToFloat64("99.99") // 99.99
active := candy.ToBool("true")    // true

// Collection operations
numbers := []int{1, 2, 3, 4, 5}
doubled := candy.Map(numbers, func(n int) int { return n * 2 })
evens := candy.Filter(numbers, func(n int) bool { return n%2 == 0 })
sum := candy.Sum(numbers)         // 15
avg := candy.Average([]float64{1.0, 2.0, 3.0}) // 2.0

// Mathematical operations
absolute := candy.Abs(-42)        // 42
square := candy.Pow(2, 3)        // 8
```

### JSON Package API

#### Core JSON Functions
```go
// Marshal/Unmarshal (platform-optimized)
func Marshal(v any) ([]byte, error)
func Unmarshal(data []byte, v any) error

// String operations
func MarshalString(v any) (string, error)
func UnmarshalString(data string, v any) error

// Must versions (panic on error)
func MustMarshal(v any) []byte
func MustMarshalString(v any) string

// File operations
func MarshalToFile(filename string, v any) error
func UnmarshalFromFile(filename string, v any) error
func MustMarshalToFile(filename string, v any)
func MustUnmarshalFromFile(filename string, v any)

// Stream operations
func NewEncoder(w io.Writer) Encoder
func NewDecoder(r io.Reader) Decoder
```

**Usage Examples:**
```go
// Basic marshal/unmarshal
data := map[string]interface{}{"name": "John", "age": 30}
jsonBytes, err := json.Marshal(data)
jsonString, err := json.MarshalString(data)

var person Person
err = json.Unmarshal(jsonBytes, &person)
err = json.UnmarshalString(jsonString, &person)

// File operations
err = json.MarshalToFile("config.json", config)
err = json.UnmarshalFromFile("config.json", &config)

// Must versions (for initialization)
configData := json.MustMarshal(defaultConfig)

// Streaming
encoder := json.NewEncoder(os.Stdout)
encoder.Encode(data)
```

### Anyx Package API

#### MapAny Type
```go
type MapAny struct { /* private fields */ }

// Constructor functions
func NewMap(m map[string]interface{}) *MapAny
func NewMapWithJson(s []byte) (*MapAny, error)
func NewMapWithYaml(s []byte) (*MapAny, error)
func NewMapWithAny(s interface{}) (*MapAny, error)

// Configuration methods
func (p *MapAny) EnableCut(seq string) *MapAny
func (p *MapAny) DisableCut() *MapAny

// Data access methods
func (p *MapAny) Get(key string) (interface{}, error)
func (p *MapAny) Set(key string, value interface{})
func (p *MapAny) Exists(key string) bool

// Type-specific getters
func (p *MapAny) GetString(key string) string
func (p *MapAny) GetInt(key string) int
func (p *MapAny) GetInt32(key string) int32
func (p *MapAny) GetInt64(key string) int64
func (p *MapAny) GetUint16(key string) uint16
func (p *MapAny) GetUint32(key string) uint32
func (p *MapAny) GetUint64(key string) uint64
func (p *MapAny) GetFloat64(key string) float64
func (p *MapAny) GetBool(key string) bool
func (p *MapAny) GetBytes(key string) []byte

// Collection getters
func (p *MapAny) GetSlice(key string) []interface{}
func (p *MapAny) GetStringSlice(key string) []string
func (p *MapAny) GetUint64Slice(key string) []uint64
func (p *MapAny) GetInt64Slice(key string) []int64
func (p *MapAny) GetUint32Slice(key string) []uint32

// Nested map access
func (p *MapAny) GetMap(key string) *MapAny

// Utility methods
func (p *MapAny) ToMap() map[string]interface{}
func (p *MapAny) ToSyncMap() *sync.Map
func (p *MapAny) Clone() *MapAny
func (p *MapAny) Range(f func(key, value interface{}) bool)
```

**Usage Examples:**
```go
// Creation
data := map[string]interface{}{
    "user": map[string]interface{}{
        "name": "John",
        "profile": map[string]interface{}{
            "age": 30,
        },
    },
}
m := anyx.NewMap(data)

// Nested access with dot notation
m.EnableCut(".")
name := m.GetString("user.name")           // "John"
age := m.GetInt("user.profile.age")        // 30

// Type-safe access
exists := m.Exists("user.email")           // false
userMap := m.GetMap("user")                // *MapAny for nested access

// From JSON
jsonData := []byte(`{"config": {"debug": true}}`)
m, err := anyx.NewMapWithJson(jsonData)
debug := m.GetBool("config.debug")         // true
```

### Routine Package API

#### Goroutine Management
```go
// Enhanced goroutine functions
func Go(f func() error)
func GoWithRecover(f func() error)
func GoWithMustSuccess(f func() error)

// Lifecycle hooks
type BeforeRoutine func(baseGid, currentGid int64)
type AfterRoutine func(currentGid int64)

func AddBeforeRoutine(f BeforeRoutine)
func AddAfterRoutine(f AfterRoutine)
```

**Usage Examples:**
```go
// Basic enhanced goroutine
routine.Go(func() error {
    return processData()
})

// With panic recovery
routine.GoWithRecover(func() error {
    return riskyOperation()
})

// Critical operation (exits on error)
routine.GoWithMustSuccess(func() error {
    return initializeCriticalService()
})

// Custom lifecycle hooks
routine.AddBeforeRoutine(func(baseGid, currentGid int64) {
    log.Printf("Starting goroutine %d from parent %d", currentGid, baseGid)
})

routine.AddAfterRoutine(func(currentGid int64) {
    log.Printf("Goroutine %d completed", currentGid)
})
```

### Wait Package API

#### Synchronization Utilities
```go
type Group struct { /* private fields */ }

// Constructor
func NewGroup() *Group

// Core methods
func (g *Group) Go(f func() error) 
func (g *Group) Wait() []error
func (g *Group) WaitFirst() error

// Timeout and context support
func (g *Group) SetTimeout(timeout time.Duration) *Group
func (g *Group) WithContext(ctx context.Context) *Group
```

**Usage Examples:**
```go
// Basic wait group
wg := wait.NewGroup()
wg.Go(func() error { return task1() })
wg.Go(func() error { return task2() })
wg.Go(func() error { return task3() })

errors := wg.Wait()  // Wait for all tasks
if len(errors) > 0 {
    // Handle errors
}

// With timeout
wg := wait.NewGroup().SetTimeout(30 * time.Second)
wg.Go(func() error { return longRunningTask() })
err := wg.WaitFirst()  // Returns first error or timeout
```

### SingleDo Package API

#### Singleton Execution
```go
type Group struct { /* private fields */ }

// Constructor
func NewGroup() *Group

// Core methods
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error)
func (g *Group) DoChan(key string, fn func() (interface{}, error)) <-chan Result

type Result struct {
    Val interface{}
    Err error
}
```

**Usage Examples:**
```go
// Singleton execution
group := singledo.NewGroup()

// Multiple calls with same key will execute fn only once
result1, err1 := group.Do("cache-key", func() (interface{}, error) {
    return expensiveOperation(), nil
})

result2, err2 := group.Do("cache-key", func() (interface{}, error) {
    return expensiveOperation(), nil  // This won't execute
})
// result1 == result2

// Channel-based execution
ch := group.DoChan("async-key", func() (interface{}, error) {
    return asyncOperation(), nil
})
result := <-ch  // Result{Val: ..., Err: ...}
```

### App Package API

#### Application Information
```go
// Global variables
var (
    Organization = "lazygophers"
    Name         string
    Version      string
    Description  string
    PackageType  ReleaseType
)

// Build information
var (
    Commit      string
    ShortCommit string
    Branch      string
    Tag         string
    BuildDate   string
    GoVersion   string
    GoOS        string
    Goarch      string
)

// Release type enumeration
type ReleaseType uint8
const (
    Debug ReleaseType = iota
    Test
    Alpha
    Beta
    Release
)

func (p ReleaseType) String() string
```

**Usage Examples:**
```go
// Set application metadata
app.Name = "MyApplication"
app.Version = "1.0.0"
app.Description = "A sample application"

// Access build information
fmt.Printf("Built from commit: %s\n", app.ShortCommit)
fmt.Printf("Build date: %s\n", app.BuildDate)
fmt.Printf("Release type: %s\n", app.PackageType.String())

// Conditional behavior based on release type
switch app.PackageType {
case app.Debug:
    enableDebugMode()
case app.Release:
    enableProductionMode()
}
```

### Error Handling Patterns

#### Common Error Types
```go
var (
    ErrNotFound = errors.New("not found")  // anyx package
    // Additional package-specific errors
)
```

#### Error Handling Examples
```go
// Graceful error handling
result, err := json.Marshal(data)
if err != nil {
    log.Printf("JSON marshal failed: %v", err)
    return err
}

// Must functions for critical paths
config := utils.Must(loadConfig())

// Validation with detailed errors
if err := utils.Validate(&user); err != nil {
    return fmt.Errorf("user validation failed: %w", err)
}
```

---

## 简体中文

### 概述
本文档为 LazyGophers Utils 库中的所有公共函数、类型和接口提供全面的 API 参考，按包组织，包含详细的签名、参数和使用示例。

### 核心工具 API

#### Must 函数
```go
// Must 如果错误为 nil 返回第一个参数，否则恐慌
func Must[T any](value T, err error) T

// MustSuccess 如果任何提供的错误非 nil 则恐慌
func MustSuccess(errs ...error)

// MustOk 如果 ok 为 false 则恐慌，否则返回值
func MustOk[T any](value T, ok bool) T
```

**示例:**
```go
// 使用 Must 读取文件
content := utils.Must(os.ReadFile("config.json"))

// 多个错误检查
utils.MustSuccess(
    initDatabase(),
    startServer(),
    setupLogging(),
)
```

### Candy 包 API

#### 类型转换函数
```go
// 字符串转换
func ToString(v interface{}) string
func String[T constraints.Ordered](s T) string

// 整数转换
func ToInt(v interface{}) int
func ToInt32(v interface{}) int32
func ToInt64(v interface{}) int64

// 浮点数转换
func ToFloat32(v interface{}) float32
func ToFloat64(v interface{}) float64

// 布尔值转换
func ToBool(v interface{}) bool
```

**使用示例:**
```go
// 类型转换
age := candy.ToInt("25")          // 25
price := candy.ToFloat64("99.99") // 99.99
active := candy.ToBool("true")    // true
```

---

## 繁體中文

### 概述
本文件為 LazyGophers Utils 函式庫中的所有公開函數、型別和介面提供全面的 API 參考，按套件組織，包含詳細的簽名、參數和使用範例。

### 核心工具 API

#### Must 函數
```go
// Must 如果錯誤為 nil 回傳第一個參數，否則恐慌
func Must[T any](value T, err error) T
```

### Candy 套件 API
```go
// 型別轉換函數
func ToString(v interface{}) string
func ToInt(v interface{}) int
func ToBool(v interface{}) bool
```

---

## Русский

### Обзор
Этот документ предоставляет комплексный справочник API для всех публичных функций, типов и интерфейсов в библиотеке LazyGophers Utils.

### API основных утилит
```go
// Must возвращает первый аргумент если ошибка nil, иначе паникует
func Must[T any](value T, err error) T
```

### API пакета Candy
```go
// Функции преобразования типов
func ToString(v interface{}) string
func ToInt(v interface{}) int
func ToBool(v interface{}) bool
```

---

## Français

### Aperçu
Ce document fournit une référence API complète pour toutes les fonctions publiques, types et interfaces de la bibliothèque LazyGophers Utils.

### API des utilitaires principaux
```go
// Must retourne le premier argument si l'erreur est nil, sinon panique
func Must[T any](value T, err error) T
```

### API du package Candy
```go
// Fonctions de conversion de type
func ToString(v interface{}) string
func ToInt(v interface{}) int
func ToBool(v interface{}) bool
```

---

## العربية

### نظرة عامة
توفر هذه الوثيقة مرجعاً شاملاً لواجهة برمجة التطبيقات لجميع الوظائف العامة والأنواع والواجهات في مكتبة LazyGophers Utils.

### واجهة برمجة التطبيقات للأدوات الأساسية
```go
// Must ترجع المعطى الأول إذا كان الخطأ nil، وإلا فإنها تصدر panic
func Must[T any](value T, err error) T
```

### واجهة برمجة التطبيقات لحزمة Candy
```go
// وظائف تحويل النوع
func ToString(v interface{}) string
func ToInt(v interface{}) int
func ToBool(v interface{}) bool
```

---

## Español

### Descripción general
Este documento proporciona una referencia de API integral para todas las funciones públicas, tipos e interfaces en la biblioteca LazyGophers Utils.

### API de utilidades principales
```go
// Must devuelve el primer argumento si el error es nil, de lo contrario entra en pánico
func Must[T any](value T, err error) T
```

### API del paquete Candy
```go
// Funciones de conversión de tipo
func ToString(v interface{}) string
func ToInt(v interface{}) int
func ToBool(v interface{}) bool
```