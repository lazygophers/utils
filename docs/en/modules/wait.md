---
title: wait - Flow Control
---

# wait - Flow Control

## Overview

The wait module provides flow control utilities including semaphore pools, synchronization, and timeout management.

## Functions

### Lock()

Acquire lock for specified key.

```go
func Lock(key string)
```

**Parameters:**
- `key` - Lock identifier

**Behavior:**
- Blocks until lock is available
- Panics if key does not exist

**Example:**
```go
wait.Ready("my-resource", 10)
wait.Lock("my-resource")
defer wait.Unlock("my-resource")

// Critical section
processResource()
```

---

### Unlock()

Release lock for specified key.

```go
func Unlock(key string)
```

**Parameters:**
- `key` - Lock identifier

**Behavior:**
- Releases lock
- Panics if key does not exist

**Example:**
```go
wait.Lock("my-resource")
defer wait.Unlock("my-resource")
```

---

### Depth()

Get current depth for specified key.

```go
func Depth(key string) int
```

**Parameters:**
- `key` - Lock identifier

**Returns:**
- Current depth (number of acquired locks)

**Example:**
```go
depth := wait.Depth("my-resource")
log.Infof("Current depth: %d", depth)
```

---

### Sync()

Execute logic function with lock.

```go
func Sync(key string, logic func() error) error
```

**Parameters:**
- `key` - Lock identifier
- `logic` - Function to execute

**Returns:**
- Error from logic function

**Example:**
```go
err := wait.Sync("database", func() error {
    return updateDatabase()
})
if err != nil {
    log.Errorf("Database update failed: %v", err)
}
```

---

### Ready()

Initialize semaphore for specified key.

```go
func Ready(key string, max int)
```

**Parameters:**
- `key` - Lock identifier
- `max` - Maximum concurrency

**Example:**
```go
wait.Ready("api-requests", 10)
```

---

## Usage Patterns

### Rate Limiting

```go
func init() {
    wait.Ready("api-calls", 100)  // Max 100 concurrent calls
}

func makeAPICall() error {
    wait.Lock("api-calls")
    defer wait.Unlock("api-calls")
    
    return callAPI()
}
```

### Resource Pooling

```go
func init() {
    wait.Ready("database-connections", 10)
}

func queryDatabase(query string) (*Result, error) {
    wait.Lock("database-connections")
    defer wait.Unlock("database-connections")
    
    conn := getDatabaseConnection()
    defer releaseDatabaseConnection(conn)
    
    return conn.Query(query)
}
```

### Critical Sections

```go
func updateSharedResource() error {
    return wait.Sync("shared-resource", func() error {
        return performUpdate()
    })
}
```

### Concurrency Control

```go
func processItems(items []Item) error {
    wait.Ready("processing", 10)
    
    var wg sync.WaitGroup
    errors := make(chan error, len(items))
    
    for _, item := range items {
        wg.Add(1)
        go func(item Item) {
            defer wg.Done()
            
            wait.Lock("processing")
            defer wait.Unlock("processing")
            
            if err := processItem(item); err != nil {
                errors <- err
            }
        }(item)
    }
    
    wg.Wait()
    close(errors)
    
    for err := range errors {
        if err != nil {
            return err
        }
    }
    
    return nil
}
```

---

## Best Practices

### Lock Management

```go
// Good: Always use defer for unlock
func safeOperation() error {
    wait.Lock("resource")
    defer wait.Unlock("resource")
    
    return performOperation()
}

// Good: Check lock depth
func checkConcurrency() int {
    return wait.Depth("resource")
}
```

### Initialization

```go
// Good: Initialize semaphores during startup
func init() {
    wait.Ready("api", 100)
    wait.Ready("database", 10)
    wait.Ready("cache", 50)
}
```

---

## Related Documentation

- [routine](/en/modules/routine) - Goroutine management
- [hystrix](/en/modules/hystrix) - Circuit breaker
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
