# routine - Goroutine Management and Task Scheduling

The `routine` package provides enhanced goroutine management with automatic error handling, panic recovery, trace context propagation, and lifecycle hooks. It simplifies concurrent programming while providing better observability and error handling.

## Features

- **Enhanced Goroutine Launch**: Safe goroutine creation with automatic error handling
- **Panic Recovery**: Automatic panic recovery with stack trace logging
- **Trace Context Propagation**: Automatic trace ID propagation across goroutines
- **Lifecycle Hooks**: Before/after hooks for goroutine execution
- **Error Handling**: Structured error handling and logging
- **Goroutine Groups**: Manage groups of related goroutines
- **Resource Management**: Automatic cleanup and resource management

## Installation

```bash
go get github.com/lazygophers/utils/routine
```

## Usage Examples

### Basic Goroutine Management

```go
package main

import (
    "fmt"
    "time"
    "github.com/lazygophers/utils/routine"
)

func main() {
    // Launch a simple goroutine with automatic error handling
    routine.Go(func() error {
        fmt.Println("Hello from goroutine!")
        time.Sleep(1 * time.Second)
        return nil
    })

    // Launch a goroutine with potential error
    routine.Go(func() error {
        fmt.Println("Processing data...")
        return fmt.Errorf("something went wrong")
    })

    time.Sleep(2 * time.Second)
}
```

### Panic Recovery

```go
// Launch goroutine with automatic panic recovery
routine.GoWithRecover(func() error {
    fmt.Println("This might panic...")

    // This will be caught and logged
    panic("unexpected error")

    return nil
})

// The panic is caught, logged, and doesn't crash the program
time.Sleep(1 * time.Second)
```

### Goroutine Groups

```go
// Create a goroutine group
group := routine.NewGroup()

// Add multiple tasks to the group
for i := 0; i < 5; i++ {
    taskID := i
    group.Go(func() error {
        fmt.Printf("Task %d started\n", taskID)
        time.Sleep(time.Duration(taskID) * time.Second)
        fmt.Printf("Task %d completed\n", taskID)
        return nil
    })
}

// Wait for all goroutines to complete
err := group.Wait()
if err != nil {
    fmt.Printf("Group execution failed: %v\n", err)
}
```

### Custom Lifecycle Hooks

```go
// Add custom before hook
routine.AddBeforeRoutine(func(baseGid, currentGid int64) {
    fmt.Printf("Starting goroutine %d from parent %d\n", currentGid, baseGid)
})

// Add custom after hook
routine.AddAfterRoutine(func(currentGid int64) {
    fmt.Printf("Finishing goroutine %d\n", currentGid)
})

// Launch goroutine - hooks will be called automatically
routine.Go(func() error {
    fmt.Println("Working...")
    return nil
})
```

### Background Tasks

```go
// Start background workers
routine.StartBackgroundWorker("data-processor", func() error {
    // Process data continuously
    for {
        err := processData()
        if err != nil {
            return err
        }
        time.Sleep(10 * time.Second)
    }
})

routine.StartBackgroundWorker("health-checker", func() error {
    // Health check loop
    for {
        err := performHealthCheck()
        if err != nil {
            return err
        }
        time.Sleep(30 * time.Second)
    }
})

// Stop all background workers gracefully
routine.StopAllBackgroundWorkers()
```

### Resource Management

```go
// Create a routine with automatic resource cleanup
routine.GoWithCleanup(
    func() error {
        // Main work
        file, err := os.Open("data.txt")
        if err != nil {
            return err
        }

        // Process file
        return processFile(file)
    },
    func() {
        // Cleanup function - always called
        fmt.Println("Cleaning up resources...")
    },
)
```

## API Reference

### Core Functions

- `Go(f func() error)` - Launch goroutine with error handling
- `GoWithRecover(f func() error)` - Launch goroutine with panic recovery
- `GoWithCleanup(work func() error, cleanup func())` - Launch with cleanup function
- `GoWithTimeout(f func() error, timeout time.Duration) error` - Launch with timeout

### Lifecycle Hooks

```go
type BeforeRoutine func(baseGid, currentGid int64)
type AfterRoutine func(currentGid int64)

// Functions
func AddBeforeRoutine(f BeforeRoutine)
func AddAfterRoutine(f AfterRoutine)
func RemoveBeforeRoutine(f BeforeRoutine)
func RemoveAfterRoutine(f AfterRoutine)
```

### Goroutine Groups

```go
type Group struct {
    // Internal implementation
}

// Functions
func NewGroup() *Group
func NewGroupWithLimit(limit int) *Group

// Methods
func (g *Group) Go(f func() error)
func (g *Group) GoWithRecover(f func() error)
func (g *Group) Wait() error
func (g *Group) WaitTimeout(timeout time.Duration) error
func (g *Group) Cancel()
func (g *Group) Size() int
```

### Background Workers

- `StartBackgroundWorker(name string, f func() error)` - Start named background worker
- `StopBackgroundWorker(name string)` - Stop specific background worker
- `StopAllBackgroundWorkers()` - Stop all background workers
- `GetBackgroundWorkerStatus(name string) WorkerStatus` - Get worker status

### Cache and Context

- `GetCache(key string) interface{}` - Get goroutine-local cache value
- `SetCache(key string, value interface{})` - Set goroutine-local cache value
- `ClearCache()` - Clear goroutine-local cache
- `GetGoroutineID() int64` - Get current goroutine ID

## Advanced Usage Examples

### Task Processing Pipeline

```go
// Create a processing pipeline
pipeline := routine.NewPipeline()

// Add processing stages
pipeline.AddStage("validate", func(data interface{}) (interface{}, error) {
    // Validate input data
    return validateData(data), nil
})

pipeline.AddStage("transform", func(data interface{}) (interface{}, error) {
    // Transform data
    return transformData(data), nil
})

pipeline.AddStage("save", func(data interface{}) (interface{}, error) {
    // Save to database
    return saveData(data), nil
})

// Process data through pipeline
data := []interface{}{item1, item2, item3}
results, err := pipeline.Process(data)
if err != nil {
    fmt.Printf("Pipeline error: %v\n", err)
}
```

### Worker Pool

```go
// Create a worker pool
pool := routine.NewWorkerPool(10) // 10 workers

// Submit jobs
for i := 0; i < 100; i++ {
    jobID := i
    pool.Submit(func() error {
        fmt.Printf("Processing job %d\n", jobID)
        time.Sleep(100 * time.Millisecond)
        return nil
    })
}

// Wait for all jobs to complete
pool.Wait()

// Shutdown the pool
pool.Shutdown()
```

### Scheduled Tasks

```go
// Create a task scheduler
scheduler := routine.NewScheduler()

// Schedule periodic task
scheduler.SchedulePeriodic("backup", 1*time.Hour, func() error {
    return performBackup()
})

// Schedule one-time task
scheduler.ScheduleOnce("cleanup", 5*time.Minute, func() error {
    return performCleanup()
})

// Schedule cron-style task
scheduler.ScheduleCron("report", "0 0 * * *", func() error {
    return generateDailyReport()
})

// Start the scheduler
scheduler.Start()

// Stop the scheduler
defer scheduler.Stop()
```

### Error Handling and Monitoring

```go
// Set up error handlers
routine.SetErrorHandler(func(err error, gid int64) {
    fmt.Printf("Goroutine %d error: %v\n", gid, err)
    // Send to monitoring system
    sendToMonitoring(err, gid)
})

routine.SetPanicHandler(func(panicValue interface{}, stack []byte, gid int64) {
    fmt.Printf("Goroutine %d panic: %v\n", gid, panicValue)
    fmt.Printf("Stack trace:\n%s\n", stack)
    // Send alert
    sendAlert(panicValue, stack, gid)
})

// Monitor goroutine statistics
stats := routine.GetStats()
fmt.Printf("Active goroutines: %d\n", stats.ActiveGoroutines)
fmt.Printf("Total launched: %d\n", stats.TotalLaunched)
fmt.Printf("Errors: %d\n", stats.Errors)
fmt.Printf("Panics: %d\n", stats.Panics)
```

## Best Practices

1. **Use Error Handling**: Always return errors from goroutine functions instead of panicking
2. **Resource Cleanup**: Use cleanup functions or defer statements for resource management
3. **Avoid Blocking**: Don't block indefinitely in goroutines without timeout
4. **Monitor Resources**: Monitor goroutine count and resource usage
5. **Use Groups**: Use goroutine groups for related tasks that should complete together

## Performance Considerations

- **Goroutine Overhead**: Each goroutine has ~8KB initial stack size
- **Context Switching**: Too many goroutines can cause excessive context switching
- **Memory Usage**: Monitor memory usage with large numbers of goroutines
- **Error Handling**: Error handling adds minimal overhead compared to panics

## Error Handling Patterns

```go
// Graceful error handling
routine.Go(func() error {
    if err := doSomething(); err != nil {
        return fmt.Errorf("doSomething failed: %w", err)
    }
    return nil
})

// With retry logic
routine.GoWithRetry(func() error {
    return doSomethingThatMightFail()
}, 3, time.Second) // 3 retries with 1 second delay

// With circuit breaker
routine.GoWithCircuitBreaker(func() error {
    return callExternalService()
}, "external-service")
```

## Integration Examples

### HTTP Server

```go
// In HTTP handler
func handleRequest(w http.ResponseWriter, r *http.Request) {
    // Process request asynchronously
    routine.Go(func() error {
        return processRequestAsync(r)
    })

    w.WriteHeader(http.StatusAccepted)
}
```

### Message Queue Consumer

```go
// Message queue consumer
routine.StartBackgroundWorker("message-consumer", func() error {
    for message := range messageChannel {
        routine.Go(func() error {
            return processMessage(message)
        })
    }
    return nil
})
```

### Database Operations

```go
// Batch database operations
group := routine.NewGroup()

for _, record := range records {
    group.Go(func() error {
        return db.Insert(record)
    })
}

if err := group.Wait(); err != nil {
    // Handle batch operation failure
    return fmt.Errorf("batch insert failed: %w", err)
}
```

## Related Packages

- `wait` - Timeout, retry, and rate limiting utilities
- `event` - Event-driven programming utilities
- `hystrix` - Circuit breaker pattern implementation