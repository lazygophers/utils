---
title: hystrix - Circuit Breaker
---

# hystrix - Circuit Breaker

## Overview

The hystrix module provides circuit breaker functionality for fault tolerance and graceful degradation. It includes optimized implementations for different use cases.

## Types

### State

Circuit breaker state.

```go
type State string

const (
    Closed   State = "closed"   // Service available
    Open     State = "open"      // Service unavailable
    HalfOpen State = "half-open" // Probing state
)
```

---

### CircuitBreakerConfig

Configuration for circuit breaker.

```go
type CircuitBreakerConfig struct {
    TimeWindow    time.Duration // Statistics time window
    OnStateChange StateChange   // State change callback
    ReadyToTrip   ReadyToTrip   // Trip condition function
    Probe         Probe         // Half-open probe function
    BufferSize    int           // Request result cache size
}
```

---

## Circuit Breaker Types

### NewCircuitBreaker()

Create optimized circuit breaker.

```go
func NewCircuitBreaker(c CircuitBreakerConfig) *CircuitBreaker
```

**Parameters:**
- `c` - Configuration options

**Returns:**
- Circuit breaker instance

**Example:**
```go
cb := hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
    TimeWindow: time.Minute,
    OnStateChange: func(oldState, newState hystrix.State) {
        log.Infof("State changed: %s -> %s", oldState, newState)
    },
    ReadyToTrip: func(successes, failures uint64) bool {
        total := successes + failures
        return total >= 10 && failures > successes
    },
})
```

---

### NewFastCircuitBreaker()

Create ultra-lightweight circuit breaker.

```go
func NewFastCircuitBreaker(failureThreshold uint64, timeWindow time.Duration) *FastCircuitBreaker
```

**Parameters:**
- `failureThreshold` - Failure threshold
- `timeWindow` - Time window

**Returns:**
- Fast circuit breaker instance

**Example:**
```go
cb := hystrix.NewFastCircuitBreaker(5, time.Minute)
```

---

### NewBatchCircuitBreaker()

Create batch processing circuit breaker.

```go
func NewBatchCircuitBreaker(config CircuitBreakerConfig, batchSize int, batchTimeout time.Duration) *BatchCircuitBreaker
```

**Parameters:**
- `config` - Configuration options
- `batchSize` - Batch size
- `batchTimeout` - Batch timeout

**Returns:**
- Batch circuit breaker instance

**Example:**
```go
cb := hystrix.NewBatchCircuitBreaker(hystrix.CircuitBreakerConfig{
    TimeWindow: time.Minute,
}, 100, time.Second*10)
```

---

## Circuit Breaker Methods

### Before()

Check if request is allowed.

```go
func (p *CircuitBreaker) Before() bool
```

**Returns:**
- true if request is allowed
- false if circuit is open

**Example:**
```go
if !cb.Before() {
    return errors.New("circuit breaker is open")
}

// Execute request
return executeRequest()
```

---

### After()

Record request result.

```go
func (p *CircuitBreaker) After(success bool)
```

**Parameters:**
- `success` - Whether request succeeded

**Example:**
```go
err := executeRequest()
cb.After(err == nil)
```

---

### Call()

Execute function with circuit breaker protection.

```go
func (p *CircuitBreaker) Call(fn func() error) error
```

**Parameters:**
- `fn` - Function to execute

**Returns:**
- Error from function
- Error if circuit is open

**Example:**
```go
err := cb.Call(func() error {
    return callExternalService()
})
if err != nil {
    log.Errorf("Service call failed: %v", err)
}
```

---

### State()

Get current circuit breaker state.

```go
func (p *CircuitBreaker) State() State
```

**Returns:**
- Current state

**Example:**
```go
state := cb.State()
log.Infof("Circuit breaker state: %s", state)
```

---

## Usage Patterns

### Service Call Protection

```go
func callExternalService() (string, error) {
    err := cb.Call(func() error {
        return makeHTTPRequest()
    })
    
    if err != nil {
        return "", err
    }
    
    return "success", nil
}
```

### Batch Processing

```go
func processBatch(items []Item) error {
    for _, item := range items {
        err := cb.Call(func() error {
            return processItem(item)
        })
        
        cb.After(err == nil)
    }
    
    return nil
}
```

### State Monitoring

```go
func monitorCircuitBreaker() {
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()
    
    for range ticker.C {
        state := cb.State()
        successes, failures := cb.Stat()
        total := cb.Total()
        
        log.Infof("State: %s, Success: %d, Failures: %d, Total: %d",
            state, successes, failures, total)
    }
}
```

---

## Best Practices

### Circuit Breaker Configuration

```go
// Good: Configure appropriate thresholds
cb := hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
    TimeWindow: time.Minute,
    ReadyToTrip: func(successes, failures uint64) bool {
        // Trip if failure rate > 50%
        total := successes + failures
        return total >= 10 && failures > total/2
    },
})
```

### Error Handling

```go
// Good: Handle circuit breaker errors
func safeServiceCall() (string, error) {
    err := cb.Call(func() error {
        return callService()
    })
    
    if err != nil {
        if err.Error() == "circuit breaker is open" {
            return "", nil  // Return default value
        }
        return "", err
    }
    
    return "success", nil
}
```

---

## Related Documentation

- [routine](/en/modules/routine) - Goroutine management
- [wait](/en/modules/wait) - Flow control
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
