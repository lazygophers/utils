---
title: config - Configuration Management
---

# config - Configuration Management

## Overview

The config module provides configuration file loading with support for multiple formats including JSON, YAML, TOML, INI, HCL, and more.

## Supported Formats

- **JSON** - `.json`, `.json5`
- **YAML** - `.yaml`, `.yml`
- **TOML** - `.toml`
- **INI** - `.ini`
- **HCL** - `.hcl`
- **XML** - `.xml`
- **Properties** - `.properties`
- **ENV** - `.env`

---

## Core Functions

### LoadConfig()

Load configuration file with validation.

```go
func LoadConfig(c any, paths ...string) error
```

**Parameters:**
- `c` - Config struct pointer
- `paths` - Optional list of file paths to search

**Returns:**
- Error if loading or validation fails

**Search Order:**
1. Explicit paths provided
2. Environment variable `LAZYGOPHERS_CONFIG`
3. Current directory (`conf.*` or `config.*`)
4. Executable directory (`conf.*` or `config.*`)

**Example:**
```go
type Config struct {
    Host     string `json:"host" toml:"host" yaml:"host" ini:"host"`
    Port     int    `json:"port" toml:"port" yaml:"port" ini:"port"`
    Debug    bool   `json:"debug" toml:"debug" yaml:"debug" ini:"debug"`
    Database struct {
        Name     string `json:"name" toml:"name" yaml:"name" ini:"name"`
        User     string `json:"user" toml:"user" yaml:"user" ini:"user"`
        Password string `json:"password" toml:"password" yaml:"password" ini:"password"`
    } `json:"database" toml:"database" yaml:"database" ini:"database"`
}

var cfg Config
if err := config.LoadConfig(&cfg, "config.json"); err != nil {
    log.Fatalf("Failed to load config: %v", err)
}
```

---

### LoadConfigSkipValidate()

Load configuration file without validation.

```go
func LoadConfigSkipValidate(c any, paths ...string) error
```

**Parameters:**
- `c` - Config struct pointer
- `paths` - Optional list of file paths to search

**Returns:**
- Error if loading fails

**Example:**
```go
var cfg Config
if err := config.LoadConfigSkipValidate(&cfg, "config.yaml"); err != nil {
    log.Fatalf("Failed to load config: %v", err)
}
```

---

### SetConfig()

Save configuration to file.

```go
func SetConfig(c any) error
```

**Parameters:**
- `c` - Config struct to save

**Returns:**
- Error if saving fails

**Example:**
```go
cfg := Config{
    Host: "localhost",
    Port: 8080,
    Debug: true,
}

if err := config.SetConfig(&cfg); err != nil {
    log.Fatalf("Failed to save config: %v", err)
}
```

---

### RegisterParser()

Register custom parser for file extension.

```go
func RegisterParser(ext string, m Marshaler, u Unmarshaler)
```

**Parameters:**
- `ext` - File extension (e.g., ".custom")
- `m` - Marshaler function
- `u` - Unmarshaler function

**Example:**
```go
config.RegisterParser(".custom",
    func(writer io.Writer, v interface{}) error {
        // Custom marshal logic
        return nil
    },
    func(reader io.Reader, v interface{}) error {
        // Custom unmarshal logic
        return nil
    },
)
```

---

## Usage Patterns

### Multi-Format Support

```go
type Config struct {
    Host string `json:"host" toml:"host" yaml:"host" ini:"host"`
    Port int    `json:"port" toml:"port" yaml:"port" ini:"port"`
}

// Works with any supported format
var cfg Config
config.LoadConfig(&cfg, "config.json")   // JSON
config.LoadConfig(&cfg, "config.yaml")   // YAML
config.LoadConfig(&cfg, "config.toml")   // TOML
config.LoadConfig(&cfg, "config.ini")    // INI
```

### Environment-Based Loading

```go
func loadConfig() *Config {
    var cfg Config
    
    // Try multiple paths
    paths := []string{
        "/etc/myapp/config.json",
        os.Getenv("HOME") + "/.myapp/config.json",
        "./config.json",
    }
    
    if err := config.LoadConfig(&cfg, paths...); err != nil {
        log.Warnf("Using default config: %v", err)
        return &Config{
            Host: "localhost",
            Port: 8080,
        }
    }
    
    return &cfg
}
```

### Nested Configuration

```go
type Config struct {
    Server struct {
        Host string `json:"host"`
        Port int    `json:"port"`
    } `json:"server"`
    
    Database struct {
        Host     string `json:"host"`
        Port     int    `json:"port"`
        Name     string `json:"name"`
        User     string `json:"user"`
        Password string `json:"password"`
    } `json:"database"`
}

var cfg Config
config.LoadConfig(&cfg, "config.json")
```

### Configuration Validation

```go
type Config struct {
    Host     string `json:"host" validate:"required"`
    Port     int    `json:"port" validate:"required,min=1,max=65535"`
    Database string `json:"database" validate:"required"`
}

var cfg Config
if err := config.LoadConfig(&cfg, "config.json"); err != nil {
    log.Fatalf("Configuration validation failed: %v", err)
}
```

---

## Best Practices

### Default Values

```go
type Config struct {
    Host     string `json:"host" default:"localhost"`
    Port     int    `json:"port" default:"8080"`
    Debug    bool   `json:"debug" default:"false"`
}

var cfg Config
config.LoadConfigSkipValidate(&cfg, "config.json")
// Default values are applied if not in config file
```

### Error Handling

```go
func loadConfigWithFallback() *Config {
    var cfg Config
    
    if err := config.LoadConfig(&cfg, "config.json"); err != nil {
        log.Warnf("Failed to load config, using defaults: %v", err)
        
        return &Config{
            Host: "localhost",
            Port: 8080,
            Debug: false,
        }
    }
    
    return &cfg
}
```

### Configuration Hot Reload

```go
func watchConfig(path string, cfg *Config) {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatalf("Failed to create watcher: %v", err)
    }
    
    watcher.Add(path)
    
    for {
        select {
        case event := <-watcher.Events:
            if event.Op&fsnotify.Write == fsnotify.Write {
                log.Info("Configuration changed, reloading...")
                
                if err := config.LoadConfig(cfg, path); err != nil {
                    log.Errorf("Failed to reload config: %v", err)
                }
            }
        case err := <-watcher.Errors:
            log.Errorf("Watcher error: %v", err)
        }
    }
}
```

---

## Related Documentation

- [validator](/en/modules/validator) - Data validation
- [defaults](/en/modules/defaults) - Default values
- [json](/en/modules/json) - JSON processing
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
