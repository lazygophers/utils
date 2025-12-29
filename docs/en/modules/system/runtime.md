---
title: runtime - Runtime Information
---

# runtime - Runtime Information

## Overview

The runtime module provides system information, runtime diagnostics, and path utilities for Go applications.

## Functions

### CachePanic()

Cache panic and prevent stack overflow.

```go
func CachePanic()
```

**Behavior:**
- Catches panic and prevents stack overflow
- Writes panic information to stderr
- Dumps stack trace

---

### CachePanicWithHandle()

Cache panic with custom handler.

```go
func CachePanicWithHandle(handle func(err interface{}))
```

**Parameters:**
- `handle` - Custom panic handler function

**Example:**
```go
runtime.CachePanicWithHandle(func(err interface{}) {
    log.Errorf("Panic occurred: %v", err)
    // Custom error handling
})
```

---

### PrintStack()

Print current stack trace.

```go
func PrintStack()
```

**Example:**
```go
func debugFunction() {
    runtime.PrintStack()
}
```

---

### ExecDir()

Get executable directory.

```go
func ExecDir() string
```

**Returns:**
- Directory containing the executable
- Empty string if error occurs

**Example:**
```go
execDir := runtime.ExecDir()
configPath := filepath.Join(execDir, "config.json")
```

---

### ExecFile()

Get executable file path.

```go
func ExecFile() string
```

**Returns:**
- Full path to the executable
- Empty string if error occurs

**Example:**
```go
execFile := runtime.ExecFile()
log.Infof("Running from: %s", execFile)
```

---

### Pwd()

Get current working directory.

```go
func Pwd() string
```

**Returns:**
- Current working directory
- Empty string if error occurs

**Example:**
```go
cwd := runtime.Pwd()
log.Infof("Current directory: %s", cwd)
```

---

### UserHomeDir()

Get user home directory.

```go
func UserHomeDir() string
```

**Returns:**
- User home directory
- Empty string if error occurs

**Example:**
```go
homeDir := runtime.UserHomeDir()
configPath := filepath.Join(homeDir, ".myapp", "config.json")
```

---

### UserConfigDir()

Get user config directory.

```go
func UserConfigDir() string
```

**Returns:**
- Platform-specific user config directory
- Empty string if error occurs

**Example:**
```go
configDir := runtime.UserConfigDir()
appConfigDir := filepath.Join(configDir, "myapp")
```

---

### UserCacheDir()

Get user cache directory.

```go
func UserCacheDir() string
```

**Returns:**
- Platform-specific user cache directory
- Empty string if error occurs

**Example:**
```go
cacheDir := runtime.UserCacheDir()
appCacheDir := filepath.Join(cacheDir, "myapp")
```

---

### LazyConfigDir()

Get lazygophers config directory.

```go
func LazyConfigDir() string
```

**Returns:**
- User config directory with lazygophers organization

**Example:**
```go
lazyConfigDir := runtime.LazyConfigDir()
configPath := filepath.Join(lazyConfigDir, "config.json")
```

---

### LazyCacheDir()

Get lazygophers cache directory.

```go
func LazyCacheDir() string
```

**Returns:**
- User cache directory with lazygophers organization

**Example:**
```go
lazyCacheDir := runtime.LazyCacheDir()
cachePath := filepath.Join(lazyCacheDir, "cache.db")
```

---

## Usage Patterns

### Application Initialization

```go
func initApp() {
    // Get executable directory
    execDir := runtime.ExecDir()
    
    // Get config path
    configPath := filepath.Join(execDir, "config.json")
    
    // Load configuration
    var cfg Config
    if err := config.LoadConfig(&cfg, configPath); err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }
    
    // Get cache directory
    cacheDir := runtime.LazyCacheDir()
    os.MkdirAll(cacheDir, 0755)
    
    // Initialize application
    app.Init(&cfg, cacheDir)
}
```

### Panic Recovery

```go
func main() {
    defer runtime.CachePanic()
    
    // Application code
    if err := runApplication(); err != nil {
        log.Fatalf("Application error: %v", err)
    }
}

func runApplication() error {
    // Application logic
    return nil
}
```

### Debug Information

```go
func printDebugInfo() {
    log.Infof("Executable: %s", runtime.ExecFile())
    log.Infof("Directory: %s", runtime.ExecDir())
    log.Infof("Working: %s", runtime.Pwd())
    log.Infof("Home: %s", runtime.UserHomeDir())
    log.Infof("Config: %s", runtime.UserConfigDir())
    log.Infof("Cache: %s", runtime.UserCacheDir())
}
```

### Custom Panic Handler

```go
func setupPanicHandler() {
    runtime.CachePanicWithHandle(func(err interface{}) {
        log.Errorf("Panic occurred: %v", err)
        
        // Send alert
        sendAlert(fmt.Sprintf("Panic: %v", err))
        
        // Save stack trace
        saveStackTrace()
        
        // Graceful shutdown
        gracefulShutdown()
    })
}

func sendAlert(message string) {
    // Send alert to monitoring system
}

func saveStackTrace() {
    // Save stack trace to file
    runtime.PrintStack()
}

func gracefulShutdown() {
    // Cleanup resources
    log.Info("Performing graceful shutdown...")
}
```

---

## Platform-Specific Paths

### Linux/Unix

```go
UserHomeDir()    // /home/username
UserConfigDir()  // /home/username/.config
UserCacheDir()   // /home/username/.cache
```

### macOS

```go
UserHomeDir()    // /Users/username
UserConfigDir()  // /Users/username/Library/Application Support
UserCacheDir()   // /Users/username/Library/Caches
```

### Windows

```go
UserHomeDir()    // C:\Users\username
UserConfigDir()  // C:\Users\username\AppData\Roaming
UserCacheDir()   // C:\Users\username\AppData\Local
```

---

## Best Practices

### Panic Handling

```go
// Good: Use defer for panic recovery
func safeFunction() {
    defer runtime.CachePanic()
    
    // Code that might panic
}

// Avoid: Not handling panics
func unsafeFunction() {
    // Code that might panic
}
```

### Path Resolution

```go
// Good: Use runtime functions for cross-platform paths
func getConfigPath() string {
    execDir := runtime.ExecDir()
    return filepath.Join(execDir, "config.json")
}

// Avoid: Hardcoding paths
func getConfigPathBad() string {
    return "/usr/local/myapp/config.json"  // Not cross-platform
}
```

### Debug Information

```go
// Good: Print debug information on startup
func main() {
    printDebugInfo()
    
    if err := runApplication(); err != nil {
        log.Fatalf("Application error: %v", err)
    }
}

func printDebugInfo() {
    log.Infof("Executable: %s", runtime.ExecFile())
    log.Infof("Working: %s", runtime.Pwd())
}
```

---

## Related Documentation

- [osx](/en/modules/osx) - OS operations
- [app](/en/modules/app) - Application framework
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
