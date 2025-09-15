# AtExit - Graceful Shutdown Handling

A cross-platform Go package that provides graceful application shutdown handling through signal interception and callback registration. The `atexit` package ensures your application can perform cleanup operations before termination.

## Features

- **Cross-Platform**: Optimized implementations for Linux, macOS, Windows, and generic Unix systems
- **Signal Handling**: Automatic interception of common termination signals (SIGINT, SIGTERM, SIGHUP, SIGQUIT)
- **Callback Registration**: Register multiple cleanup functions to execute on shutdown
- **Panic Recovery**: Built-in panic recovery to prevent one callback from affecting others
- **Thread-Safe**: Concurrent registration and execution of callbacks
- **Zero Dependencies**: No external dependencies beyond Go standard library

## Installation

```bash
go get github.com/lazygophers/utils/atexit
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/lazygophers/utils/atexit"
)

func main() {
    // Register cleanup functions
    atexit.Register(func() {
        fmt.Println("Closing database connections...")
        // db.Close()
    })

    atexit.Register(func() {
        fmt.Println("Saving application state...")
        // saveState()
    })

    // Simulate application work
    fmt.Println("Application running...")
    time.Sleep(30 * time.Second)
    fmt.Println("Application finished")
}
```

## API Reference

### Functions

#### `Register(callback func())`

Registers a callback function to be executed when the application receives a termination signal.

**Parameters:**
- `callback func()`: The function to execute on shutdown. If nil, the call is ignored.

**Example:**
```go
atexit.Register(func() {
    log.Println("Cleanup completed")
})
```

**Behavior:**
- Callbacks are executed in the order they were registered
- Each callback runs in its own protected goroutine with panic recovery
- Signal handling is initialized on the first call to Register()
- Thread-safe for concurrent registration

## Platform-Specific Behavior

### Linux (`atexit_linux.go`)
- Handles: `SIGINT`, `SIGTERM`, `SIGQUIT`, `SIGHUP`
- Optimized for Linux signal handling
- Uses Linux-specific signal handling optimizations

### macOS (`atexit_darwin.go`)
- Handles: `SIGINT`, `SIGTERM`, `SIGQUIT`, `SIGHUP`
- Supports additional Unix signals
- Can integrate with system logging

### Windows (`atexit_windows.go`)
- Handles Windows-specific termination events
- Console control events (Ctrl+C, Ctrl+Break)
- System shutdown events
- Service stop requests

### Generic Unix (`atexit.go`)
- Handles: `SIGINT`, `SIGTERM`
- Fallback implementation for other Unix systems
- Basic signal handling with panic recovery

## Usage Examples

### Database Cleanup

```go
package main

import (
    "database/sql"
    "log"

    "github.com/lazygophers/utils/atexit"
    _ "github.com/lib/pq"
)

func main() {
    db, err := sql.Open("postgres", "connection_string")
    if err != nil {
        log.Fatal(err)
    }

    // Register database cleanup
    atexit.Register(func() {
        log.Println("Closing database connection...")
        if err := db.Close(); err != nil {
            log.Printf("Error closing database: %v", err)
        }
    })

    // Your application logic here
    runApplication(db)
}
```

### HTTP Server Graceful Shutdown

```go
package main

import (
    "context"
    "log"
    "net/http"
    "time"

    "github.com/lazygophers/utils/atexit"
)

func main() {
    server := &http.Server{
        Addr:    ":8080",
        Handler: http.DefaultServeMux,
    }

    // Register server shutdown
    atexit.Register(func() {
        log.Println("Shutting down HTTP server...")
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        if err := server.Shutdown(ctx); err != nil {
            log.Printf("Server shutdown error: %v", err)
        }
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })

    log.Println("Server starting on :8080")
    log.Fatal(server.ListenAndServe())
}
```

### Multiple Resource Cleanup

```go
package main

import (
    "log"
    "os"

    "github.com/lazygophers/utils/atexit"
)

func main() {
    // Open files
    logFile, err := os.Create("app.log")
    if err != nil {
        log.Fatal(err)
    }

    configFile, err := os.Open("config.json")
    if err != nil {
        log.Fatal(err)
    }

    // Register cleanup for each resource
    atexit.Register(func() {
        log.Println("Closing log file...")
        logFile.Close()
    })

    atexit.Register(func() {
        log.Println("Closing config file...")
        configFile.Close()
    })

    atexit.Register(func() {
        log.Println("Performing final cleanup...")
        os.Remove("temp.lock")
    })

    // Application logic
    runApplication()
}
```

## Best Practices

### 1. Register Early
Register your cleanup callbacks as early as possible in your application lifecycle:

```go
func main() {
    // Register cleanup immediately after resource creation
    db := setupDatabase()
    atexit.Register(func() { db.Close() })

    cache := setupCache()
    atexit.Register(func() { cache.Shutdown() })

    // Continue with application logic
}
```

### 2. Handle Errors Gracefully
Cleanup functions should handle errors without panicking:

```go
atexit.Register(func() {
    if err := resource.Close(); err != nil {
        log.Printf("Warning: Failed to close resource: %v", err)
        // Don't panic - other callbacks need to run
    }
})
```

### 3. Timeout Long-Running Operations
Set timeouts for potentially long-running cleanup operations:

```go
atexit.Register(func() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := server.Shutdown(ctx); err != nil {
        log.Printf("Server shutdown timeout: %v", err)
    }
})
```

### 4. Order Dependencies
Register callbacks in reverse dependency order (last dependency first):

```go
func main() {
    cache := setupCache()
    db := setupDatabase()
    server := setupServer(db, cache)

    // Register in reverse order of dependencies
    atexit.Register(func() { server.Shutdown() })  // Depends on db and cache
    atexit.Register(func() { cache.Close() })      // Independent
    atexit.Register(func() { db.Close() })         // Independent
}
```

## Signal Handling Details

### Supported Signals

| Platform | SIGINT | SIGTERM | SIGQUIT | SIGHUP | Windows Events |
|----------|--------|---------|---------|--------|----------------|
| Linux    | ✓      | ✓       | ✓       | ✓      | -              |
| macOS    | ✓      | ✓       | ✓       | ✓      | -              |
| Windows  | -      | -       | -       | -      | ✓              |
| Generic  | ✓      | ✓       | -       | -      | -              |

### Signal Sources

- **SIGINT**: Interrupt from keyboard (Ctrl+C)
- **SIGTERM**: Termination request
- **SIGQUIT**: Quit from keyboard (Ctrl+\)
- **SIGHUP**: Hangup detected on controlling terminal
- **Windows**: Console control events, system shutdown

## Performance Considerations

- **Low Overhead**: Signal handling is initialized only once
- **Concurrent Safe**: Uses RWMutex for thread-safe callback management
- **Panic Recovery**: Each callback runs in a protected environment
- **Memory Efficient**: Minimal memory footprint for signal handling

## Thread Safety

The atexit package is fully thread-safe:

- **Registration**: Multiple goroutines can safely register callbacks concurrently
- **Execution**: Callbacks are executed sequentially but each in its own protected scope
- **Signal Handling**: Signal handlers are initialized once using `sync.Once`

## Limitations

1. **One-Time Execution**: Callbacks are executed only once per application shutdown
2. **No Cancellation**: Once registered, callbacks cannot be unregistered
3. **Sequential Execution**: Callbacks run sequentially, not in parallel
4. **Platform Differences**: Signal handling varies between operating systems

## Contributing

Contributions are welcome! Please ensure:

1. Cross-platform compatibility
2. Thread safety
3. Comprehensive tests
4. Documentation updates

## License

This package is part of the LazyGophers Utils library and follows the same licensing terms.