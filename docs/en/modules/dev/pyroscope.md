---
title: pyroscope - Performance Profiling
---

# pyroscope - Performance Profiling

## Overview

The pyroscope module provides integration with Pyroscope for production monitoring and performance profiling.

## Functions

### load()

Load Pyroscope server address and start profiling.

```go
func load(address string)
```

**Parameters:**
- `address` - Pyroscope server address

**Behavior:**
- Starts Pyroscope client
- Configures profiling for the application

**Example:**
```go
pyroscope.load("http://localhost:4040")
```

---

## Usage Patterns

### Application Integration

```go
func main() {
    // Start profiling
    pyroscope.load("http://localhost:4040")
    
    // Application code
    runApplication()
}
```

### Production Monitoring

```go
func setupMonitoring() {
    address := os.Getenv("PYROSCOPE_ADDRESS")
    if address == "" {
        address = "http://localhost:4040"
    }
    
    pyroscope.load(address)
    
    log.Info("Pyroscope profiling enabled")
}
```

---

## Best Practices

### Configuration

```go
// Good: Configure Pyroscope address
address := os.Getenv("PYROSCOPE_ADDRESS")
if address == "" {
    address = "http://localhost:4040"
}

// Good: Handle connection errors
pyroscope.load(address)
// Connection errors are logged
```

---

## Related Documentation

- [runtime](/en/modules/runtime) - Runtime information
- [app](/en/modules/app) - Application framework
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
