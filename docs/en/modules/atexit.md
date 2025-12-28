---
title: atexit - Graceful Shutdown
---

# atexit - Graceful Shutdown

## Overview

The atexit module provides graceful shutdown functionality by registering exit handlers that are called when the application terminates.

## Functions

### Register()

Register a callback function to be called on exit.

```go
func Register(callback func())
```

**Parameters:**
- `callback` - Function to call on exit

**Behavior:**
- Registers callback for execution on exit
- Initializes signal handler on first registration
- Callbacks are executed in registration order

**Example:**
```go
func main() {
    atexit.Register(cleanupResources)
    atexit.Register(closeConnections)
    atexit.Register(saveState)
    
    // Application code
    runApplication()
    
    // Exit handlers will be called automatically
}

func cleanupResources() {
    log.Info("Cleaning up resources...")
}

func closeConnections() {
    log.Info("Closing connections...")
}

func saveState() {
    log.Info("Saving state...")
}
```

---

## Usage Patterns

### Resource Cleanup

```go
func setupDatabase() *sql.DB {
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    
    atexit.Register(func() {
        log.Info("Closing database connection")
        db.Close()
    })
    
    return db
}

func setupHTTPServer() *http.Server {
    server := &http.Server{
        Addr:    ":8080",
        Handler: router,
    }
    
    atexit.Register(func() {
        log.Info("Shutting down HTTP server")
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()
        server.Shutdown(ctx)
    })
    
    go server.ListenAndServe()
    return server
}
```

### Signal Handling

The atexit module automatically handles common termination signals:

- **Unix-like systems**: SIGINT, SIGTERM
- **Windows**: Console events

```go
func main() {
    atexit.Register(func() {
        log.Info("Received termination signal")
        gracefulShutdown()
    })
    
    // Application will exit gracefully on SIGINT/SIGTERM
    select {}
}
```

### Multiple Handlers

```go
func main() {
    // Register multiple cleanup handlers
    atexit.Register(cleanupDatabase)
    atexit.Register(closeFiles)
    atexit.Register(flushLogs)
    atexit.Register(notifyMonitoring)
    
    // Application code
    runApplication()
}

func cleanupDatabase() {
    log.Info("Cleaning up database...")
}

func closeFiles() {
    log.Info("Closing open files...")
}

func flushLogs() {
    log.Info("Flushing logs...")
}

func notifyMonitoring() {
    log.Info("Notifying monitoring system...")
}
```

---

## Best Practices

### Handler Registration

```go
// Good: Register handlers during initialization
func init() {
    atexit.Register(cleanupResources)
}

// Good: Register handlers with error recovery
func registerHandler() {
    atexit.Register(func() {
        defer func() {
            if r := recover(); r != nil {
                log.Errorf("Panic in exit handler: %v", r)
            }
        }()
        
        cleanup()
    })
}
```

### Resource Management

```go
// Good: Use defer for immediate cleanup
func processFile(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()
    
    // Process file
    return nil
}

// Good: Use atexit for application-level cleanup
func main() {
    db := setupDatabase()
    server := setupHTTPServer()
    
    atexit.Register(func() {
        db.Close()
        server.Shutdown(context.Background())
    })
    
    // Application code
}
```

---

## Related Documentation

- [runtime](/en/modules/runtime) - Runtime information
- [app](/en/modules/app) - Application framework
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
