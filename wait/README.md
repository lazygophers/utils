# wait - Timeout, Retry, and Rate Limiting Utilities

The `wait` package provides utilities for controlling concurrent operations, implementing timeouts, retry mechanisms, and rate limiting. It includes semaphore pools, wait groups, and asynchronous operation management.

## Features

- **Semaphore Pools**: Control concurrent operations with named pools
- **Wait Groups**: Enhanced wait group operations with timeout support
- **Async Operations**: Asynchronous task execution with result handling
- **Rate Limiting**: Built-in rate limiting capabilities
- **Timeout Management**: Timeout support for all operations
- **Thread Safety**: All operations are goroutine-safe

## Installation

```bash
go get github.com/lazygophers/utils/wait
```

## Usage Examples

### Semaphore Pool Operations

```go
package main

import (
    "fmt"
    "sync"
    "time"
    "github.com/lazygophers/utils/wait"
)

func main() {
    // Create a semaphore pool with max 3 concurrent operations
    poolName := "api_requests"
    maxConcurrent := 3

    // Start concurrent workers
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()

            // Acquire semaphore
            wait.Lock(poolName, maxConcurrent)
            defer wait.Unlock(poolName)

            fmt.Printf("Worker %d started\n", id)
            time.Sleep(time.Second) // Simulate work
            fmt.Printf("Worker %d finished\n", id)
        }(i)
    }

    wg.Wait()
}
```

### Wait Group with Timeout

```go
// Create a wait group with timeout
group := wait.NewGroup()

// Add tasks
for i := 0; i < 5; i++ {
    group.Add(1)
    go func(id int) {
        defer group.Done()

        // Simulate work
        time.Sleep(time.Duration(id) * time.Second)
        fmt.Printf("Task %d completed\n", id)
    }(i)
}

// Wait with timeout
timeout := 3 * time.Second
if group.WaitTimeout(timeout) {
    fmt.Println("All tasks completed within timeout")
} else {
    fmt.Println("Timeout reached, some tasks may still be running")
}
```

### Asynchronous Operations

```go
// Define an async operation
operation := func() (interface{}, error) {
    time.Sleep(2 * time.Second)
    return "Operation completed", nil
}

// Execute asynchronously
asyncResult := wait.Async(operation)

// Do other work while operation runs
fmt.Println("Doing other work...")
time.Sleep(1 * time.Second)

// Get result with timeout
result, err := asyncResult.GetWithTimeout(5 * time.Second)
if err != nil {
    fmt.Printf("Error: %v\n", err)
} else {
    fmt.Printf("Result: %v\n", result)
}
```

### Pool Management

```go
// Check pool status
poolName := "database_connections"
maxConnections := 10

// Create pool
wait.NewPool(poolName, maxConnections)

// Check current usage
depth := wait.GetPoolDepth(poolName)
fmt.Printf("Current connections: %d/%d\n", depth, maxConnections)

// Lock multiple resources
wait.LockMultiple([]string{"db_pool", "cache_pool"}, []int{5, 3})
defer wait.UnlockMultiple([]string{"db_pool", "cache_pool"})

// Perform operations requiring both resources
fmt.Println("Using database and cache...")
```

## API Reference

### Semaphore Pool Functions

- `Lock(key string, max int)` - Acquire semaphore (creates pool if needed)
- `Unlock(key string)` - Release semaphore
- `TryLock(key string, max int) bool` - Try to acquire semaphore without blocking
- `LockWithTimeout(key string, max int, timeout time.Duration) bool` - Acquire with timeout

### Pool Management

- `NewPool(key string, max int)` - Create a new semaphore pool
- `GetPool(key string) *Pool` - Get existing pool
- `GetPoolDepth(key string) int` - Get current pool usage
- `DestroyPool(key string)` - Remove pool and release resources

### Multiple Pool Operations

- `LockMultiple(keys []string, maxes []int)` - Lock multiple pools atomically
- `UnlockMultiple(keys []string)` - Unlock multiple pools
- `TryLockMultiple(keys []string, maxes []int) bool` - Try lock multiple pools

### Wait Group Operations

```go
type Group struct {
    // Internal implementation
}

// Methods
func NewGroup() *Group
func (g *Group) Add(delta int)
func (g *Group) Done()
func (g *Group) Wait()
func (g *Group) WaitTimeout(timeout time.Duration) bool
func (g *Group) WaitContext(ctx context.Context) error
```

### Async Operations

```go
type AsyncResult struct {
    // Internal implementation
}

// Functions
func Async(fn func() (interface{}, error)) *AsyncResult
func AsyncWithContext(ctx context.Context, fn func() (interface{}, error)) *AsyncResult

// Methods
func (ar *AsyncResult) Get() (interface{}, error)
func (ar *AsyncResult) GetWithTimeout(timeout time.Duration) (interface{}, error)
func (ar *AsyncResult) IsReady() bool
func (ar *AsyncResult) Cancel()
```

### Pool Type

```go
type Pool struct {
    // Internal channel-based semaphore
}

// Methods
func (p *Pool) Lock()
func (p *Pool) Unlock()
func (p *Pool) TryLock() bool
func (p *Pool) LockWithTimeout(timeout time.Duration) bool
func (p *Pool) Depth() int
func (p *Pool) Cap() int
```

## Advanced Usage Examples

### Rate Limiting HTTP Requests

```go
func makeAPIRequests(urls []string) {
    poolName := "api_rate_limit"
    maxConcurrent := 5 // Max 5 concurrent requests

    var wg sync.WaitGroup
    for _, url := range urls {
        wg.Add(1)
        go func(u string) {
            defer wg.Done()

            // Rate limit requests
            wait.Lock(poolName, maxConcurrent)
            defer wait.Unlock(poolName)

            resp, err := http.Get(u)
            if err != nil {
                fmt.Printf("Error fetching %s: %v\n", u, err)
                return
            }
            defer resp.Body.Close()

            fmt.Printf("Fetched %s: %d\n", u, resp.StatusCode)
        }(url)
    }

    wg.Wait()
}
```

### Database Connection Pool

```go
type DatabaseManager struct {
    poolName string
    maxConns int
}

func NewDatabaseManager(maxConnections int) *DatabaseManager {
    poolName := "database_connections"
    wait.NewPool(poolName, maxConnections)

    return &DatabaseManager{
        poolName: poolName,
        maxConns: maxConnections,
    }
}

func (dm *DatabaseManager) ExecuteQuery(query string) error {
    // Acquire connection from pool
    if !wait.LockWithTimeout(dm.poolName, dm.maxConns, 5*time.Second) {
        return fmt.Errorf("timeout acquiring database connection")
    }
    defer wait.Unlock(dm.poolName)

    // Execute query
    fmt.Printf("Executing query: %s\n", query)
    time.Sleep(100 * time.Millisecond) // Simulate query execution

    return nil
}

func (dm *DatabaseManager) GetStats() (current, max int) {
    return wait.GetPoolDepth(dm.poolName), dm.maxConns
}
```

### Batch Processing with Timeout

```go
func processBatchWithTimeout(items []string, timeout time.Duration) []string {
    results := make([]string, 0, len(items))
    resultsChan := make(chan string, len(items))

    group := wait.NewGroup()

    // Start workers
    for _, item := range items {
        group.Add(1)
        go func(data string) {
            defer group.Done()

            // Simulate processing
            processed := fmt.Sprintf("processed_%s", data)
            resultsChan <- processed
        }(item)
    }

    // Wait for completion or timeout
    done := make(chan bool)
    go func() {
        group.Wait()
        close(done)
    }()

    select {
    case <-done:
        // All completed
        close(resultsChan)
        for result := range resultsChan {
            results = append(results, result)
        }
    case <-time.After(timeout):
        // Timeout occurred
        fmt.Println("Batch processing timed out")
        close(resultsChan)
        for result := range resultsChan {
            results = append(results, result)
        }
    }

    return results
}
```

### Resource Management

```go
type ResourceManager struct {
    pools map[string]int
}

func NewResourceManager() *ResourceManager {
    return &ResourceManager{
        pools: map[string]int{
            "cpu_intensive": 2,
            "memory_intensive": 3,
            "io_operations": 10,
        },
    }
}

func (rm *ResourceManager) ExecuteTask(taskType string, task func()) error {
    maxConcurrent, exists := rm.pools[taskType]
    if !exists {
        return fmt.Errorf("unknown task type: %s", taskType)
    }

    // Try to acquire resource with timeout
    if !wait.LockWithTimeout(taskType, maxConcurrent, 30*time.Second) {
        return fmt.Errorf("timeout acquiring resource for %s", taskType)
    }
    defer wait.Unlock(taskType)

    // Execute task
    task()
    return nil
}
```

## Best Practices

1. **Pool Naming**: Use descriptive pool names that indicate the resource type
2. **Timeout Management**: Always use timeouts for long-running operations
3. **Resource Cleanup**: Ensure `Unlock()` is called in defer statements
4. **Proper Sizing**: Size pools based on actual resource constraints
5. **Monitoring**: Monitor pool usage to identify bottlenecks

## Performance Considerations

- **Channel-Based Semaphores**: Uses buffered channels for efficient semaphore implementation
- **Lock-Free Operations**: Pool depth checking is lock-free where possible
- **Memory Efficiency**: Minimal memory overhead per pool
- **Scalability**: Supports thousands of concurrent operations

## Error Handling

The package uses panic-free design patterns:

```go
// Safe operations that won't panic
success := wait.TryLock("mypool", 10)
if !success {
    fmt.Println("Could not acquire lock")
}

// Timeout-based operations
if wait.LockWithTimeout("mypool", 10, 5*time.Second) {
    defer wait.Unlock("mypool")
    // Perform work
} else {
    fmt.Println("Timeout acquiring lock")
}
```

## Related Packages

- `routine` - Goroutine management and task scheduling
- `event` - Event-driven programming utilities
- `hystrix` - Circuit breaker pattern implementation