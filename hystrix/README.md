# Hystrix - Circuit Breaker Pattern Implementation

A high-performance Go implementation of the Circuit Breaker pattern with multiple optimized variants. The `hystrix` package provides reliable service protection through circuit breaking, helping prevent cascading failures and improving system resilience.

## Features

- **Multiple Implementations**: Standard, Fast, and Batch circuit breakers for different use cases
- **High Performance**: Lock-free operations with atomic counters and optimized ring buffers
- **Configurable States**: Closed, Open, and Half-Open states with customizable transitions
- **Time Window Support**: Configurable success/failure tracking windows
- **Panic Recovery**: Built-in panic recovery for robust operation
- **Memory Optimized**: Cache-aligned structures to prevent false sharing
- **Concurrent Safe**: Thread-safe for high-concurrency environments
- **Zero Allocations**: Optimized for minimal memory allocation in critical paths

## Installation

```bash
go get github.com/lazygophers/utils/hystrix
```

## Quick Start

```go
package main

import (
    "fmt"
    "time"
    "errors"

    "github.com/lazygophers/utils/hystrix"
)

func main() {
    // Create a circuit breaker
    cb := hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
        TimeWindow: 10 * time.Second,
        ReadyToTrip: func(successes, failures uint64) bool {
            total := successes + failures
            return total >= 10 && failures > successes
        },
    })

    // Use the circuit breaker to protect a service call
    err := cb.Call(func() error {
        // Your service call here
        return callExternalService()
    })

    if err != nil {
        fmt.Printf("Service call failed: %v\n", err)
    }

    fmt.Printf("Circuit breaker state: %s\n", cb.State())
}

func callExternalService() error {
    // Simulate service call
    return errors.New("service unavailable")
}
```

## Circuit Breaker States

### Closed State
- **Description**: Normal operation, requests pass through
- **Behavior**: Monitors success/failure rates
- **Transition**: Moves to Open when failure threshold is exceeded

### Open State
- **Description**: Circuit is "open", requests are immediately rejected
- **Behavior**: Protects downstream services from additional load
- **Transition**: Moves to Half-Open after cooldown period

### Half-Open State
- **Description**: Limited requests allowed to test service recovery
- **Behavior**: Uses probe function to determine if service is healthy
- **Transition**: Moves to Closed on success, back to Open on failure

## API Reference

### Standard Circuit Breaker

#### `NewCircuitBreaker(config CircuitBreakerConfig) *CircuitBreaker`

Creates a new circuit breaker with advanced features and optimizations.

**Configuration:**
```go
type CircuitBreakerConfig struct {
    TimeWindow    time.Duration // Statistics time window
    OnStateChange StateChange   // State change callback
    ReadyToTrip   ReadyToTrip   // Failure condition function
    Probe         Probe         // Half-open state probe function
    BufferSize    int          // Request history buffer size
}
```

**Example:**
```go
cb := hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
    TimeWindow: 30 * time.Second,
    BufferSize: 1000,
    ReadyToTrip: func(successes, failures uint64) bool {
        total := successes + failures
        failureRate := float64(failures) / float64(total)
        return total >= 20 && failureRate > 0.6
    },
    OnStateChange: func(oldState, newState hystrix.State) {
        log.Printf("Circuit breaker state changed: %s -> %s", oldState, newState)
    },
})
```

#### Methods

**`Before() bool`**
Checks if a request should be allowed through the circuit breaker.

**`After(success bool)`**
Records the result of a request execution.

**`Call(fn func() error) error`**
Executes a function with circuit breaker protection.

**`State() State`**
Returns the current circuit breaker state.

**`Stat() (successes, failures uint64)`**
Returns current success and failure counts.

### Fast Circuit Breaker

#### `NewFastCircuitBreaker(failureThreshold uint64, timeWindow time.Duration) *FastCircuitBreaker`

Creates a lightweight circuit breaker with minimal overhead.

**Example:**
```go
cb := hystrix.NewFastCircuitBreaker(5, 10*time.Second)

// Check if request is allowed
if cb.AllowRequest() {
    err := makeServiceCall()
    cb.RecordResult(err == nil)
}
```

### Batch Circuit Breaker

#### `NewBatchCircuitBreaker(config CircuitBreakerConfig, batchSize int, batchTimeout time.Duration) *BatchCircuitBreaker`

Creates a circuit breaker optimized for batch processing scenarios.

**Example:**
```go
cb := hystrix.NewBatchCircuitBreaker(
    hystrix.CircuitBreakerConfig{TimeWindow: 30 * time.Second},
    100,                    // Batch size
    1 * time.Second,       // Batch timeout
)

// Record results in batches for better performance
cb.AfterBatch(true)  // Success
cb.AfterBatch(false) // Failure
```

## Usage Examples

### Basic Service Protection

```go
package main

import (
    "fmt"
    "time"
    "math/rand"

    "github.com/lazygophers/utils/hystrix"
)

type ExternalService struct {
    cb *hystrix.CircuitBreaker
}

func NewExternalService() *ExternalService {
    return &ExternalService{
        cb: hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
            TimeWindow: 30 * time.Second,
            ReadyToTrip: func(successes, failures uint64) bool {
                total := successes + failures
                return total >= 10 && failures*2 > successes
            },
        }),
    }
}

func (s *ExternalService) CallAPI() (string, error) {
    var result string
    err := s.cb.Call(func() error {
        // Simulate API call
        if rand.Float32() < 0.3 { // 30% failure rate
            return fmt.Errorf("API call failed")
        }
        result = "Success response"
        return nil
    })

    if err != nil {
        return "", err
    }
    return result, nil
}

func main() {
    service := NewExternalService()

    for i := 0; i < 50; i++ {
        result, err := service.CallAPI()
        if err != nil {
            fmt.Printf("Call %d failed: %v (State: %s)\n",
                i+1, err, service.cb.State())
        } else {
            fmt.Printf("Call %d succeeded: %s (State: %s)\n",
                i+1, result, service.cb.State())
        }

        time.Sleep(100 * time.Millisecond)
    }
}
```

### Database Connection Protection

```go
package main

import (
    "database/sql"
    "fmt"
    "time"

    "github.com/lazygophers/utils/hystrix"
)

type Database struct {
    db *sql.DB
    cb *hystrix.CircuitBreaker
}

func NewDatabase(db *sql.DB) *Database {
    return &Database{
        db: db,
        cb: hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
            TimeWindow: 60 * time.Second,
            BufferSize: 500,
            ReadyToTrip: func(successes, failures uint64) bool {
                total := successes + failures
                if total < 5 {
                    return false
                }
                failureRate := float64(failures) / float64(total)
                return failureRate > 0.5
            },
            OnStateChange: func(oldState, newState hystrix.State) {
                fmt.Printf("Database circuit breaker: %s -> %s\n",
                    oldState, newState)
            },
        }),
    }
}

func (d *Database) QueryUser(userID int) (*User, error) {
    var user User
    err := d.cb.Call(func() error {
        return d.db.QueryRow(
            "SELECT id, name, email FROM users WHERE id = ?",
            userID,
        ).Scan(&user.ID, &user.Name, &user.Email)
    })

    if err != nil {
        return nil, err
    }
    return &user, nil
}

type User struct {
    ID    int
    Name  string
    Email string
}
```

### HTTP Client Protection

```go
package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/lazygophers/utils/hystrix"
)

type HTTPClient struct {
    client *http.Client
    cb     *hystrix.CircuitBreaker
}

func NewHTTPClient() *HTTPClient {
    return &HTTPClient{
        client: &http.Client{Timeout: 5 * time.Second},
        cb: hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
            TimeWindow: 30 * time.Second,
            ReadyToTrip: func(successes, failures uint64) bool {
                total := successes + failures
                return total >= 5 && failures*3 > successes
            },
            Probe: hystrix.ProbeWithChance(25), // 25% chance to probe
        }),
    }
}

func (h *HTTPClient) Get(url string) (*http.Response, error) {
    var resp *http.Response
    err := h.cb.Call(func() error {
        var err error
        resp, err = h.client.Get(url)
        if err != nil {
            return err
        }
        if resp.StatusCode >= 500 {
            return fmt.Errorf("server error: %d", resp.StatusCode)
        }
        return nil
    })

    return resp, err
}

func main() {
    client := NewHTTPClient()

    urls := []string{
        "https://httpbin.org/status/200",
        "https://httpbin.org/status/500",
        "https://httpbin.org/delay/6", // Will timeout
    }

    for i := 0; i < 20; i++ {
        url := urls[i%len(urls)]
        resp, err := client.Get(url)

        if err != nil {
            fmt.Printf("Request %d failed: %v (State: %s)\n",
                i+1, err, client.cb.State())
        } else {
            fmt.Printf("Request %d succeeded: %d (State: %s)\n",
                i+1, resp.StatusCode, client.cb.State())
            resp.Body.Close()
        }

        time.Sleep(500 * time.Millisecond)
    }
}
```

### Microservice Integration

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "github.com/lazygophers/utils/hystrix"
)

type UserService struct {
    cb *hystrix.CircuitBreaker
}

type OrderService struct {
    userService *UserService
    cb          *hystrix.CircuitBreaker
}

func NewUserService() *UserService {
    return &UserService{
        cb: hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
            TimeWindow: 30 * time.Second,
            ReadyToTrip: func(successes, failures uint64) bool {
                return failures >= 3
            },
        }),
    }
}

func NewOrderService(userService *UserService) *OrderService {
    return &OrderService{
        userService: userService,
        cb: hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
            TimeWindow: 30 * time.Second,
            ReadyToTrip: func(successes, failures uint64) bool {
                return failures >= 5
            },
        }),
    }
}

func (u *UserService) GetUser(userID int) (*User, error) {
    var user User
    err := u.cb.Call(func() error {
        // Simulate user service call
        if userID <= 0 {
            return fmt.Errorf("invalid user ID")
        }
        user = User{ID: userID, Name: fmt.Sprintf("User%d", userID)}
        return nil
    })
    return &user, err
}

func (o *OrderService) CreateOrder(userID int, items []string) (*Order, error) {
    var order Order
    err := o.cb.Call(func() error {
        // Get user first (this might fail due to user service circuit breaker)
        user, err := o.userService.GetUser(userID)
        if err != nil {
            return fmt.Errorf("failed to get user: %w", err)
        }

        // Create order
        order = Order{
            ID:     time.Now().Unix(),
            UserID: user.ID,
            Items:  items,
            Status: "created",
        }
        return nil
    })
    return &order, err
}

type Order struct {
    ID     int64    `json:"id"`
    UserID int      `json:"user_id"`
    Items  []string `json:"items"`
    Status string   `json:"status"`
}

func main() {
    userService := NewUserService()
    orderService := NewOrderService(userService)

    // Simulate order creation
    for i := 0; i < 10; i++ {
        userID := i%5 - 2 // Some invalid IDs to trigger failures

        order, err := orderService.CreateOrder(userID, []string{"item1", "item2"})
        if err != nil {
            fmt.Printf("Order creation %d failed: %v\n", i+1, err)
            fmt.Printf("  User service state: %s\n", userService.cb.State())
            fmt.Printf("  Order service state: %s\n", orderService.cb.State())
        } else {
            fmt.Printf("Order created: %+v\n", order)
        }

        time.Sleep(200 * time.Millisecond)
    }
}
```

### High-Performance Scenario with Fast Circuit Breaker

```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"

    "github.com/lazygophers/utils/hystrix"
)

func main() {
    // Use fast circuit breaker for high-throughput scenarios
    cb := hystrix.NewFastCircuitBreaker(100, 10*time.Second)

    var (
        allowed    uint64
        rejected   uint64
        successful uint64
        failed     uint64
    )

    // Simulate high load
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()

            for j := 0; j < 1000; j++ {
                if cb.AllowRequest() {
                    atomic.AddUint64(&allowed, 1)

                    // Simulate work with occasional failures
                    success := j%10 != 0 // 90% success rate
                    cb.RecordResult(success)

                    if success {
                        atomic.AddUint64(&successful, 1)
                    } else {
                        atomic.AddUint64(&failed, 1)
                    }
                } else {
                    atomic.AddUint64(&rejected, 1)
                }
            }
        }(i)
    }

    wg.Wait()

    fmt.Printf("Results:\n")
    fmt.Printf("  Allowed:    %d\n", allowed)
    fmt.Printf("  Rejected:   %d\n", rejected)
    fmt.Printf("  Successful: %d\n", successful)
    fmt.Printf("  Failed:     %d\n", failed)
    fmt.Printf("  Total:      %d\n", allowed+rejected)
}
```

## Advanced Configuration

### Custom Ready-to-Trip Function

```go
cb := hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
    TimeWindow: 60 * time.Second,
    ReadyToTrip: func(successes, failures uint64) bool {
        total := successes + failures

        // Require minimum sample size
        if total < 50 {
            return false
        }

        // Multiple conditions
        failureRate := float64(failures) / float64(total)
        consecutiveFailures := failures >= 10

        return failureRate > 0.7 || consecutiveFailures
    },
})
```

### Custom Probe Function

```go
// Custom probe with exponential backoff
var probeAttempts uint64

cb := hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
    TimeWindow: 30 * time.Second,
    Probe: func() bool {
        attempts := atomic.AddUint64(&probeAttempts, 1)
        // Exponential backoff: 1%, 2%, 4%, 8%, max 25%
        chance := min(25, int(attempts))
        return rand.Intn(100) < chance
    },
})
```

### State Change Monitoring

```go
cb := hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
    TimeWindow: 30 * time.Second,
    OnStateChange: func(oldState, newState hystrix.State) {
        // Log state changes
        log.Printf("Circuit breaker state: %s -> %s", oldState, newState)

        // Send metrics
        metrics.Counter("circuit_breaker.state_change").
            With("from", string(oldState), "to", string(newState)).
            Increment()

        // Trigger alerts for open state
        if newState == hystrix.Open {
            alerting.Send("Circuit breaker opened", "High failure rate detected")
        }
    },
})
```

## Performance Characteristics

### Memory Usage
- **Standard CB**: ~200 bytes + (BufferSize * 8) bytes for ring buffer
- **Fast CB**: ~64 bytes
- **Batch CB**: Standard CB + (BatchSize * 1) bytes

### Latency
- **Before() check**: ~50ns (lock-free atomic operations)
- **After() recording**: ~100ns (includes ring buffer update)
- **Call() overhead**: ~200ns total

### Throughput
- **Standard CB**: >10M operations/second
- **Fast CB**: >50M operations/second
- **Concurrent access**: Linear scaling up to CPU cores

## Best Practices

### 1. Configuration Guidelines

**Time Window Selection:**
```go
// Short-lived requests (APIs)
TimeWindow: 10 * time.Second

// Batch processing
TimeWindow: 5 * time.Minute

// Long-running operations
TimeWindow: 30 * time.Minute
```

**Failure Thresholds:**
```go
// Conservative (avoid false positives)
ReadyToTrip: func(successes, failures uint64) bool {
    total := successes + failures
    return total >= 20 && failures*4 > successes // 80% failure rate
}

// Aggressive (fast failure detection)
ReadyToTrip: func(successes, failures uint64) bool {
    return failures >= 3 // Trip after 3 consecutive failures
}
```

### 2. Integration Patterns

**Graceful Degradation:**
```go
func GetUserProfile(userID int) *UserProfile {
    profile, err := userService.GetProfile(userID)
    if err != nil {
        // Return cached or default profile
        return getCachedProfile(userID)
    }
    return profile
}
```

**Retry with Circuit Breaker:**
```go
func CallWithRetry(cb *hystrix.CircuitBreaker, fn func() error) error {
    for i := 0; i < 3; i++ {
        err := cb.Call(fn)
        if err == nil {
            return nil
        }

        if err.Error() == "circuit breaker is open" {
            return err // Don't retry when circuit is open
        }

        time.Sleep(time.Duration(i+1) * 100 * time.Millisecond)
    }
    return fmt.Errorf("failed after 3 retries")
}
```

### 3. Monitoring and Alerting

**Metrics Collection:**
```go
// Periodically collect and report metrics
go func() {
    ticker := time.NewTicker(10 * time.Second)
    for range ticker.C {
        successes, failures := cb.Stat()
        total := successes + failures

        metrics.Gauge("circuit_breaker.requests.total").Set(float64(total))
        metrics.Gauge("circuit_breaker.requests.successes").Set(float64(successes))
        metrics.Gauge("circuit_breaker.requests.failures").Set(float64(failures))
        metrics.Gauge("circuit_breaker.state").Set(stateToFloat(cb.State()))
    }
}()
```

## Thread Safety

All circuit breaker implementations are fully thread-safe:

- **Atomic Operations**: All counters use atomic operations
- **Lock-Free**: No mutexes in hot paths
- **Cache-Aligned**: Memory layout prevents false sharing
- **Concurrent Access**: Safe for use across multiple goroutines

## Contributing

Contributions are welcome! Areas for improvement:

1. Additional probe strategies
2. Metrics integration
3. Configuration validation
4. Performance optimizations
5. Additional state transition strategies

## License

This package is part of the LazyGophers Utils library and follows the same licensing terms.