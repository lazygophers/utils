---
title: routine - Goroutine Management
---

# routine - Goroutine Management

## Overview

The routine module provides goroutine management utilities including worker pools, task scheduling, and panic recovery.

## Functions

### Go()

Execute function in a goroutine with panic recovery.

```go
func Go(f func() (err error))
```

**Parameters:**
- `f` - Function to execute

**Behavior:**
- Executes function in a goroutine
- Logs errors if function returns error
- Automatically manages trace IDs

**Example:**
```go
routine.Go(func() error {
    if err := processData(); err != nil {
        return err
    }
    return nil
})
```

---

### GoWithRecover()

Execute function in a goroutine with full panic recovery.

```go
func GoWithRecover(f func() (err error))
```

**Parameters:**
- `f` - Function to execute

**Behavior:**
- Executes function in a goroutine
- Catches panics and logs stack trace
- Logs errors if function returns error

**Example:**
```go
routine.GoWithRecover(func() error {
    // This will be caught and logged
    panic("Something went wrong")
    return nil
})
```

---

### GoWithMustSuccess()

Execute function in a goroutine with panic on error.

```go
func GoWithMustSuccess(f func() (err error))
```

**Parameters:**
- `f` - Function to execute

**Behavior:**
- Executes function in a goroutine
- Exits process if function returns error

**Example:**
```go
routine.GoWithMustSuccess(func() error {
    if err := criticalOperation(); err != nil {
        return err
    }
    return nil
})
// Process will exit if criticalOperation fails
```

---

### AddBeforeRoutine()

Add callback to execute before goroutine starts.

```go
func AddBeforeRoutine(f BeforeRoutine)
```

**Parameters:**
- `f` - Callback function

**Example:**
```go
routine.AddBeforeRoutine(func(baseGid, currentGid int64) {
    log.Infof("Starting goroutine: %d -> %d", baseGid, currentGid)
})
```

---

### AddAfterRoutine()

Add callback to execute after goroutine completes.

```go
func AddAfterRoutine(f AfterRoutine)
```

**Parameters:**
- `f` - Callback function

**Example:**
```go
routine.AddAfterRoutine(func(currentGid int64) {
    log.Infof("Completed goroutine: %d", currentGid)
})
```

---

## Usage Patterns

### Background Tasks

```go
func startBackgroundTasks() {
    routine.Go(func() error {
        ticker := time.NewTicker(time.Minute)
        defer ticker.Stop()
        
        for range ticker.C {
            if err := performMaintenance(); err != nil {
                log.Errorf("Maintenance failed: %v", err)
            }
        }
        return nil
    })
}
```

### Error Handling

```go
func safeAsyncOperation() {
    routine.GoWithRecover(func() error {
        // This panic will be caught
        if someCondition {
            panic("Unexpected error")
        }
        return nil
    })
}
```

### Task Scheduling

```go
func scheduleTask(delay time.Duration, task func()) {
    routine.Go(func() error {
        time.Sleep(delay)
        task()
        return nil
    })
}
```

---

## Best Practices

### Error Recovery

```go
// Good: Use GoWithRecover for critical goroutines
routine.GoWithRecover(func() error {
    criticalOperation()
    return nil
})

// Good: Use Go for simple tasks
routine.Go(func() error {
    simpleOperation()
    return nil
})
```

---

## Related Documentation

- [wait](/en/modules/wait) - Flow control
- [hystrix](/en/modules/hystrix) - Circuit breaker
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
