# config - Multi-Format Configuration File Loading

The `config` package provides utilities for loading and managing configuration files in multiple formats. It supports JSON, YAML, TOML, INI, and HCL formats with automatic format detection and environment variable substitution.

## Features

- **Multiple Formats**: Support for JSON, YAML, TOML, INI, and HCL configuration files
- **Auto-Detection**: Automatic format detection based on file extension
- **Environment Variables**: Environment variable substitution in configuration values
- **Flexible Loading**: Load from files, strings, or readers
- **Type Safety**: Strong typing with struct mapping
- **Validation**: Built-in validation support

## Installation

```bash
go get github.com/lazygophers/utils/config
```

## Usage Examples

### Basic Configuration Loading

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/config"
)

type AppConfig struct {
    Server struct {
        Host string `json:"host" yaml:"host"`
        Port int    `json:"port" yaml:"port"`
    } `json:"server" yaml:"server"`
    Database struct {
        URL      string `json:"url" yaml:"url"`
        MaxConns int    `json:"max_conns" yaml:"max_conns"`
    } `json:"database" yaml:"database"`
}

func main() {
    var cfg AppConfig

    // Load from YAML file
    err := config.LoadFile("config.yaml", &cfg)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Server: %s:%d\n", cfg.Server.Host, cfg.Server.Port)
}
```

### Multi-Format Support

```go
// Load JSON configuration
err := config.LoadFile("config.json", &cfg)

// Load YAML configuration
err := config.LoadFile("config.yaml", &cfg)

// Load TOML configuration
err := config.LoadFile("config.toml", &cfg)

// Load INI configuration
err := config.LoadFile("config.ini", &cfg)

// Load HCL configuration
err := config.LoadFile("config.hcl", &cfg)
```

### Environment Variable Substitution

```go
// config.yaml with environment variables
// server:
//   host: ${HOST:localhost}
//   port: ${PORT:8080}

var cfg AppConfig
err := config.LoadFile("config.yaml", &cfg)
// Automatically substitutes HOST and PORT environment variables
// Uses default values if environment variables are not set
```

### Loading from String

```go
yamlContent := `
server:
  host: localhost
  port: 8080
database:
  url: postgres://localhost/mydb
`

var cfg AppConfig
err := config.LoadString(yamlContent, &cfg, config.FormatYAML)
if err != nil {
    panic(err)
}
```

### Loading with Validation

```go
type Config struct {
    Server struct {
        Host string `yaml:"host" validate:"required"`
        Port int    `yaml:"port" validate:"required,min=1,max=65535"`
    } `yaml:"server" validate:"required"`
}

var cfg Config
err := config.LoadAndValidate("config.yaml", &cfg)
if err != nil {
    // Handle validation errors
    panic(err)
}
```

## API Reference

### Types

#### Format

```go
type Format int

const (
    FormatAuto Format = iota // Auto-detect based on file extension
    FormatJSON               // JSON format
    FormatYAML               // YAML format
    FormatTOML               // TOML format
    FormatINI                // INI format
    FormatHCL                // HCL format
)
```

### Functions

#### File Loading
- `LoadFile(filename string, v interface{}) error` - Load configuration from file with auto-detection
- `LoadFileWithFormat(filename string, v interface{}, format Format) error` - Load with specific format
- `LoadAndValidate(filename string, v interface{}) error` - Load and validate configuration

#### String Loading
- `LoadString(content string, v interface{}, format Format) error` - Load from string content
- `LoadReader(r io.Reader, v interface{}, format Format) error` - Load from reader

#### Environment Processing
- `ProcessEnvVars(content string) string` - Process environment variable substitutions
- `SetEnvPrefix(prefix string)` - Set prefix for environment variable lookup

#### Validation
- `Validate(v interface{}) error` - Validate struct using validator tags
- `ValidateWithRules(v interface{}, rules map[string]string) error` - Validate with custom rules

### Configuration Options

```go
type Options struct {
    EnvPrefix       string            // Prefix for environment variables
    EnvSubstitution bool              // Enable environment variable substitution
    Validation      bool              // Enable validation
    CaseSensitive   bool              // Case-sensitive field matching
    DefaultValues   map[string]interface{} // Default values
}

// LoadWithOptions provides full control over loading behavior
func LoadWithOptions(filename string, v interface{}, opts Options) error
```

## Environment Variable Substitution

The package supports flexible environment variable substitution:

### Syntax
- `${VAR}` - Required environment variable (error if not set)
- `${VAR:default}` - Optional with default value
- `${PREFIX_VAR}` - With prefix (when EnvPrefix is set)

### Examples

```yaml
# config.yaml
database:
  host: ${DB_HOST:localhost}
  port: ${DB_PORT:5432}
  user: ${DB_USER}
  password: ${DB_PASSWORD}

server:
  debug: ${DEBUG:false}
  workers: ${WORKERS:4}
```

```go
// Set environment variables
os.Setenv("DB_HOST", "prod-db.example.com")
os.Setenv("DB_USER", "myapp")
os.Setenv("DB_PASSWORD", "secret")
os.Setenv("DEBUG", "true")

// Load configuration
var cfg Config
config.LoadFile("config.yaml", &cfg)
// Results in:
// cfg.Database.Host = "prod-db.example.com"
// cfg.Database.Port = 5432 (default)
// cfg.Database.User = "myapp"
// cfg.Database.Password = "secret"
// cfg.Server.Debug = true
// cfg.Server.Workers = 4 (default)
```

## File Format Examples

### JSON
```json
{
  "server": {
    "host": "localhost",
    "port": 8080
  },
  "database": {
    "url": "postgres://localhost/mydb"
  }
}
```

### YAML
```yaml
server:
  host: localhost
  port: 8080
database:
  url: postgres://localhost/mydb
```

### TOML
```toml
[server]
host = "localhost"
port = 8080

[database]
url = "postgres://localhost/mydb"
```

### INI
```ini
[server]
host = localhost
port = 8080

[database]
url = postgres://localhost/mydb
```

### HCL
```hcl
server {
  host = "localhost"
  port = 8080
}

database {
  url = "postgres://localhost/mydb"
}
```

## Best Practices

1. **Use Environment Variables**: Use environment variable substitution for secrets and environment-specific values
2. **Validate Configuration**: Always validate configuration after loading
3. **Provide Defaults**: Use default values for optional configuration
4. **Structure Configuration**: Organize configuration into logical sections
5. **Document Format**: Choose the format that best fits your team's preferences

## Error Handling

The package provides detailed error information:

```go
err := config.LoadFile("config.yaml", &cfg)
if err != nil {
    switch {
    case config.IsValidationError(err):
        // Handle validation errors
        fmt.Printf("Validation failed: %v\n", err)
    case config.IsFormatError(err):
        // Handle format/parsing errors
        fmt.Printf("Format error: %v\n", err)
    case config.IsFileError(err):
        // Handle file access errors
        fmt.Printf("File error: %v\n", err)
    default:
        fmt.Printf("Unknown error: %v\n", err)
    }
}
```

## Related Packages

- `validator` - Struct validation utilities
- `app` - Application lifecycle management
- `defaults` - Struct default value handling