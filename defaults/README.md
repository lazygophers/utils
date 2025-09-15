# Defaults - Struct Default Value Handling

A powerful Go package that provides comprehensive struct field default value assignment through reflection. The `defaults` package automatically populates struct fields with default values based on struct tags, custom functions, and type-specific logic.

## Features

- **Comprehensive Type Support**: Handles all Go types including primitives, pointers, structs, slices, arrays, maps, channels, and functions
- **Flexible Configuration**: Multiple error handling modes and customization options
- **Struct Tag Support**: Use `default:"value"` tags to specify field defaults
- **Custom Default Functions**: Register custom default value generators for specific types
- **Time Handling**: Advanced time parsing with multiple format support and "now" keyword
- **Nested Structure Support**: Recursively processes nested structs and pointers
- **JSON Integration**: Parse complex defaults from JSON strings
- **Thread Safe**: Safe for concurrent use with proper synchronization
- **Zero Dependencies**: Uses only Go standard library

## Installation

```bash
go get github.com/lazygophers/utils/defaults
```

## Quick Start

```go
package main

import (
    "fmt"
    "time"

    "github.com/lazygophers/utils/defaults"
)

type Config struct {
    Name        string        `default:"MyApp"`
    Port        int           `default:"8080"`
    Debug       bool          `default:"true"`
    Timeout     time.Duration `default:"30s"`
    CreatedAt   time.Time     `default:"now"`
    Tags        []string      `default:"[\"api\", \"web\"]"`
}

func main() {
    var config Config

    // Set defaults using struct tags
    defaults.SetDefaults(&config)

    fmt.Printf("Config: %+v\n", config)
    // Output: Config: {Name:MyApp Port:8080 Debug:true Timeout:30s CreatedAt:2024-01-01 12:00:00 Tags:[api web]}
}
```

## API Reference

### Core Functions

#### `SetDefaults(value interface{})`

Sets default values using struct tags with panic on error (default behavior).

**Parameters:**
- `value interface{}`: Pointer to the struct to populate

**Example:**
```go
type User struct {
    Name string `default:"Anonymous"`
    Age  int    `default:"18"`
}

var user User
defaults.SetDefaults(&user)
```

#### `SetDefaultsWithOptions(value interface{}, opts *Options) error`

Sets default values with custom configuration options.

**Parameters:**
- `value interface{}`: Pointer to the struct to populate
- `opts *Options`: Configuration options

**Returns:**
- `error`: Error if any occurred (depends on error mode)

### Configuration Options

#### `Options` struct

```go
type Options struct {
    ErrorMode        ErrorMode                // Error handling strategy
    CustomDefaults   map[string]DefaultFunc   // Custom default functions
    ValidateDefaults bool                     // Whether to validate defaults
    AllowOverwrite   bool                     // Allow overwriting non-zero values
}
```

#### `ErrorMode` constants

- `ErrorModePanic`: Panic on errors (default)
- `ErrorModeIgnore`: Ignore errors and continue
- `ErrorModeReturn`: Return errors without panicking

#### Custom Default Functions

#### `RegisterCustomDefault(typeName string, fn DefaultFunc)`

Registers a custom default value function for a specific type.

**Parameters:**
- `typeName string`: Type identifier ("string", "int", "float", "bool", "uint", "func")
- `fn DefaultFunc`: Function that returns default value

**Example:**
```go
// Register custom string default
defaults.RegisterCustomDefault("string", func() interface{} {
    return "custom-default-" + time.Now().Format("20060102")
})
```

#### `ClearCustomDefaults()`

Clears all registered custom default functions.

## Supported Types and Tags

### Primitive Types

#### String
```go
type Example struct {
    Name     string `default:"John Doe"`
    Empty    string `default:""`
    Optional string // No default, remains empty
}
```

#### Integer Types
```go
type Example struct {
    Age     int   `default:"25"`
    Count   int64 `default:"1000"`
    Retries uint  `default:"3"`
}
```

#### Float Types
```go
type Example struct {
    Price  float64 `default:"99.99"`
    Rating float32 `default:"4.5"`
}
```

#### Boolean
```go
type Example struct {
    Enabled  bool `default:"true"`
    Disabled bool `default:"false"`
}
```

### Complex Types

#### Time
```go
type Example struct {
    CreatedAt time.Time `default:"now"`
    UpdatedAt time.Time `default:"2024-01-01 15:04:05"`
    Birthday  time.Time `default:"1990-01-01"`
}
```

Supported time formats:
- `"now"` - Current time
- RFC3339: `"2006-01-02T15:04:05Z07:00"`
- RFC3339Nano: `"2006-01-02T15:04:05.999999999Z07:00"`
- Date time: `"2006-01-02 15:04:05"`
- Date only: `"2006-01-02"`
- Time only: `"15:04:05"`

#### Pointers
```go
type Example struct {
    Name *string `default:"John"`
    Age  *int    `default:"30"`
}
```

#### Slices
```go
type Example struct {
    Tags     []string `default:"[\"tag1\", \"tag2\"]"`
    Numbers  []int    `default:"1,2,3,4,5"`
    Empty    []string // Initialized as empty slice
}
```

#### Arrays
```go
type Example struct {
    Colors [3]string `default:"red,green,blue"`
    Matrix [2]int    `default:"10,20"`
}
```

#### Maps
```go
type Example struct {
    Config   map[string]string `default:"{\"key1\":\"value1\", \"key2\":\"value2\"}"`
    Settings map[string]int    // Initialized as empty map
}
```

#### Channels
```go
type Example struct {
    Messages chan string `default:"10"` // Buffer size
    Events   chan int    `default:"0"`  // Unbuffered
}
```

#### Nested Structs
```go
type Address struct {
    Street string `default:"123 Main St"`
    City   string `default:"Springfield"`
}

type Person struct {
    Name    string  `default:"John"`
    Address Address // Automatically processed
}
```

## Usage Examples

### Basic Configuration

```go
package main

import (
    "fmt"
    "time"

    "github.com/lazygophers/utils/defaults"
)

type DatabaseConfig struct {
    Host         string        `default:"localhost"`
    Port         int           `default:"5432"`
    Username     string        `default:"admin"`
    Password     string        // No default for security
    MaxConns     int           `default:"10"`
    Timeout      time.Duration `default:"30s"`
    SSL          bool          `default:"true"`
    RetryAttempts uint          `default:"3"`
}

func main() {
    var dbConfig DatabaseConfig
    defaults.SetDefaults(&dbConfig)

    fmt.Printf("Database Config:\n")
    fmt.Printf("  Host: %s\n", dbConfig.Host)
    fmt.Printf("  Port: %d\n", dbConfig.Port)
    fmt.Printf("  SSL: %t\n", dbConfig.SSL)
    fmt.Printf("  Timeout: %v\n", dbConfig.Timeout)
}
```

### Error Handling Options

```go
package main

import (
    "fmt"
    "log"

    "github.com/lazygophers/utils/defaults"
)

type Config struct {
    Port    int    `default:"invalid"`  // Invalid integer
    Timeout string `default:"30s"`
}

func main() {
    var config Config

    // Option 1: Return errors instead of panicking
    opts := &defaults.Options{
        ErrorMode: defaults.ErrorModeReturn,
    }

    if err := defaults.SetDefaultsWithOptions(&config, opts); err != nil {
        log.Printf("Error setting defaults: %v", err)
    }

    // Option 2: Ignore errors and continue
    opts.ErrorMode = defaults.ErrorModeIgnore
    defaults.SetDefaultsWithOptions(&config, opts)

    fmt.Printf("Config: %+v\n", config)
}
```

### Custom Default Functions

```go
package main

import (
    "fmt"
    "os"
    "time"

    "github.com/lazygophers/utils/defaults"
)

type AppConfig struct {
    AppName     string `default:""`  // Will use custom default
    Environment string `default:""`  // Will use custom default
    Version     string `default:"1.0.0"`
}

func main() {
    // Register custom defaults
    defaults.RegisterCustomDefault("string", func() interface{} {
        if appName := os.Getenv("APP_NAME"); appName != "" {
            return appName
        }
        return "MyApplication"
    })

    var config AppConfig
    defaults.SetDefaults(&config)

    fmt.Printf("App Config: %+v\n", config)
}
```

### Complex Nested Structures

```go
package main

import (
    "fmt"
    "time"

    "github.com/lazygophers/utils/defaults"
)

type Server struct {
    Host string `default:"0.0.0.0"`
    Port int    `default:"8080"`
}

type Database struct {
    Host     string `default:"localhost"`
    Port     int    `default:"5432"`
    Username string `default:"admin"`
    Pool     *PoolConfig
}

type PoolConfig struct {
    MaxConnections int           `default:"10"`
    IdleTimeout    time.Duration `default:"5m"`
}

type ApplicationConfig struct {
    Name      string    `default:"MyApp"`
    Debug     bool      `default:"false"`
    CreatedAt time.Time `default:"now"`
    Server    Server
    Database  Database
    Features  []string `default:"[\"auth\", \"api\", \"web\"]"`
    Metadata  map[string]interface{} `default:"{\"version\":\"1.0\"}"`
}

func main() {
    var config ApplicationConfig
    defaults.SetDefaults(&config)

    fmt.Printf("Application: %s\n", config.Name)
    fmt.Printf("Server: %s:%d\n", config.Server.Host, config.Server.Port)
    fmt.Printf("Database: %s:%d\n", config.Database.Host, config.Database.Port)
    fmt.Printf("Pool Max Connections: %d\n", config.Database.Pool.MaxConnections)
    fmt.Printf("Features: %v\n", config.Features)
}
```

### Overwrite Existing Values

```go
package main

import (
    "fmt"

    "github.com/lazygophers/utils/defaults"
)

type Config struct {
    Name    string `default:"Default Name"`
    Port    int    `default:"8080"`
    Enabled bool   `default:"true"`
}

func main() {
    // Pre-populate some fields
    config := Config{
        Name: "Custom Name",
        Port: 3000,
    }

    fmt.Printf("Before: %+v\n", config)

    // Option 1: Don't overwrite existing values (default behavior)
    defaults.SetDefaults(&config)
    fmt.Printf("After (no overwrite): %+v\n", config)

    // Option 2: Allow overwriting existing values
    opts := &defaults.Options{
        AllowOverwrite: true,
    }
    defaults.SetDefaultsWithOptions(&config, opts)
    fmt.Printf("After (with overwrite): %+v\n", config)
}
```

### Validation and Custom Logic

```go
package main

import (
    "fmt"
    "strings"

    "github.com/lazygophers/utils/defaults"
)

type UserProfile struct {
    Username string `default:"guest"`
    Email    string `default:"user@example.com"`
    Role     string `default:"user"`
}

func main() {
    // Register custom validation/transformation
    defaults.RegisterCustomDefault("string", func() interface{} {
        return strings.ToLower("DEFAULT_VALUE")
    })

    opts := &defaults.Options{
        ValidateDefaults: true,
        ErrorMode:        defaults.ErrorModeReturn,
    }

    var profile UserProfile
    if err := defaults.SetDefaultsWithOptions(&profile, opts); err != nil {
        fmt.Printf("Validation error: %v\n", err)
    }

    fmt.Printf("Profile: %+v\n", profile)
}
```

## Advanced Features

### Working with Interfaces

```go
type Config struct {
    Data interface{} `default:"{\"key\": \"value\"}"`
    List interface{} `default:"[1, 2, 3]"`
}

var config Config
defaults.SetDefaults(&config)
// Data will be parsed as JSON object/array
```

### Channel Buffer Sizes

```go
type EventSystem struct {
    Events    chan Event `default:"100"`  // Buffered channel with size 100
    Errors    chan error `default:"0"`    // Unbuffered channel
    Shutdown  chan bool  `default:"1"`    // Buffered channel with size 1
}
```

### Function Type Defaults

```go
type Handlers struct {
    OnError   func(error)   // Custom default through RegisterCustomDefault
    OnSuccess func(string)
}

// Register custom function defaults
defaults.RegisterCustomDefault("func", func() interface{} {
    return func(err error) {
        log.Printf("Default error handler: %v", err)
    }
})
```

## Performance Considerations

### Memory Allocation
- Minimal allocations for primitive types
- Efficient handling of slices and maps
- Reuses reflection values where possible

### Execution Speed
- Caches reflection type information
- Optimized for repeated use on similar structs
- Concurrent-safe for multiple goroutines

### Best Practices for Performance

1. **Reuse Options Objects:**
```go
var opts = &defaults.Options{
    ErrorMode: defaults.ErrorModeIgnore,
}

// Reuse opts for multiple calls
defaults.SetDefaultsWithOptions(&config1, opts)
defaults.SetDefaultsWithOptions(&config2, opts)
```

2. **Minimize Custom Functions:**
Custom default functions add overhead - use them sparingly.

3. **Prefer Simple Tags:**
Simple string/number defaults are fastest.

## Error Handling Strategies

### Panic Mode (Default)
```go
defaults.SetDefaults(&config) // Panics on error
```

### Error Return Mode
```go
opts := &defaults.Options{ErrorMode: defaults.ErrorModeReturn}
if err := defaults.SetDefaultsWithOptions(&config, opts); err != nil {
    log.Printf("Error: %v", err)
}
```

### Ignore Mode
```go
opts := &defaults.Options{ErrorMode: defaults.ErrorModeIgnore}
defaults.SetDefaultsWithOptions(&config, opts) // Continues on error
```

## Thread Safety

The defaults package is thread-safe:

- **SetDefaults**: Safe for concurrent use
- **SetDefaultsWithOptions**: Safe for concurrent use
- **RegisterCustomDefault**: Safe for concurrent registration
- **ClearCustomDefaults**: Safe for concurrent access

## Contributing

Contributions are welcome! Please ensure:

1. Comprehensive test coverage
2. Proper error handling
3. Documentation updates
4. Performance benchmarks for new features

## License

This package is part of the LazyGophers Utils library and follows the same licensing terms.