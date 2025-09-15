# SingleDo - Singleton Execution Pattern Implementation

The `singledo` module provides thread-safe singleton execution pattern implementation with intelligent caching and deduplication capabilities. It ensures that expensive operations are executed only once within a specified time window, making it ideal for caching expensive computations, API calls, and resource-intensive operations.

## Features

- **Thread-Safe Execution**: Guarantees only one execution per key at a time
- **Time-Based Caching**: Results are cached for a configurable duration
- **Generic Type Support**: Fully type-safe with Go generics
- **Deduplication**: Multiple concurrent calls for the same key share the same result
- **Group Management**: Organize operations by keys for independent caching
- **Zero Memory Allocation**: Efficient implementation with minimal overhead
- **Panic Recovery**: Graceful handling of panics in executed functions

## Installation

```bash
go get github.com/lazygophers/utils
```

## Usage

### Basic Single Execution

```go
package main

import (
    "fmt"
    "time"
    "github.com/lazygophers/utils/singledo"
)

func main() {
    // Create a Single instance with 5-minute cache duration
    single := singledo.NewSingle[string](5 * time.Minute)

    // Expensive operation that will be cached
    expensiveOperation := func() (string, error) {
        fmt.Println("Executing expensive operation...")
        time.Sleep(2 * time.Second) // Simulate expensive work
        return "computed result", nil
    }

    // First call - executes the function
    result1, err := single.Do(expensiveOperation)
    fmt.Printf("Result 1: %s, Error: %v\n", result1, err)

    // Second call within cache window - returns cached result
    result2, err := single.Do(expensiveOperation)
    fmt.Printf("Result 2: %s, Error: %v\n", result2, err)

    // Reset cache manually if needed
    single.Reset()
}
```

### Concurrent Deduplication

```go
package main

import (
    "fmt"
    "sync"
    "time"
    "github.com/lazygophers/utils/singledo"
)

func main() {
    single := singledo.NewSingle[int](1 * time.Minute)

    expensiveComputation := func() (int, error) {
        fmt.Println("Computing...")
        time.Sleep(3 * time.Second)
        return 42, nil
    }

    var wg sync.WaitGroup

    // Start multiple goroutines
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            result, err := single.Do(expensiveComputation)
            fmt.Printf("Goroutine %d got result: %d, error: %v\n", id, result, err)
        }(i)
    }

    wg.Wait()
    // Only one "Computing..." will be printed, all goroutines get the same result
}
```

### Group-Based Key Management

```go
package main

import (
    "fmt"
    "time"
    "github.com/lazygophers/utils/singledo"
)

func main() {
    // Create a group for managing multiple cached operations
    group := singledo.NewSingleGroup[string](2 * time.Minute)

    // Different operations for different keys
    fetchUserData := func() (string, error) {
        fmt.Println("Fetching user data...")
        time.Sleep(1 * time.Second)
        return "user data", nil
    }

    fetchConfigData := func() (string, error) {
        fmt.Println("Fetching config data...")
        time.Sleep(1 * time.Second)
        return "config data", nil
    }

    // Execute operations with different keys
    userData, _ := group.Do("user:123", fetchUserData)
    configData, _ := group.Do("config:app", fetchConfigData)

    fmt.Printf("User: %s, Config: %s\n", userData, configData)

    // Subsequent calls within cache window return cached results
    userData2, _ := group.Do("user:123", fetchUserData)
    fmt.Printf("Cached user data: %s\n", userData2)
}
```

### API Response Caching

```go
package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"
    "github.com/lazygophers/utils/singledo"
)

type APIResponse struct {
    Data    map[string]interface{} `json:"data"`
    Status  string                 `json:"status"`
}

func main() {
    // Cache API responses for 10 minutes
    apiCache := singledo.NewSingleGroup[*APIResponse](10 * time.Minute)

    fetchFromAPI := func(endpoint string) func() (*APIResponse, error) {
        return func() (*APIResponse, error) {
            fmt.Printf("Making API call to %s...\n", endpoint)

            resp, err := http.Get("https://api.example.com/" + endpoint)
            if err != nil {
                return nil, err
            }
            defer resp.Body.Close()

            body, err := io.ReadAll(resp.Body)
            if err != nil {
                return nil, err
            }

            var apiResp APIResponse
            if err := json.Unmarshal(body, &apiResp); err != nil {
                return nil, err
            }

            return &apiResp, nil
        }
    }

    // Multiple calls to the same endpoint will be deduplicated
    result1, err := apiCache.Do("users", fetchFromAPI("users"))
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    result2, err := apiCache.Do("users", fetchFromAPI("users"))
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    fmt.Printf("Same instance: %v\n", result1 == result2) // true - same cached instance
}
```

### Database Query Caching

```go
package main

import (
    "database/sql"
    "fmt"
    "time"
    "github.com/lazygophers/utils/singledo"
)

type User struct {
    ID   int    `db:"id"`
    Name string `db:"name"`
    Email string `db:"email"`
}

func main() {
    // Cache database queries for 5 minutes
    queryCache := singledo.NewSingleGroup[*User](5 * time.Minute)

    // Mock database connection
    var db *sql.DB // Initialize your database connection

    fetchUser := func(userID int) func() (*User, error) {
        return func() (*User, error) {
            fmt.Printf("Querying database for user ID: %d\n", userID)

            query := "SELECT id, name, email FROM users WHERE id = ?"
            row := db.QueryRow(query, userID)

            user := &User{}
            err := row.Scan(&user.ID, &user.Name, &user.Email)
            if err != nil {
                return nil, err
            }

            return user, nil
        }
    }

    // Cache key includes user ID
    userKey := fmt.Sprintf("user:%d", 123)

    // First call hits database
    user1, err := queryCache.Do(userKey, fetchUser(123))
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    // Second call returns cached result
    user2, err := queryCache.Do(userKey, fetchUser(123))
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    fmt.Printf("User 1: %+v\n", user1)
    fmt.Printf("Same cached user: %v\n", user1 == user2)
}
```

### Complex Data Processing

```go
package main

import (
    "crypto/md5"
    "fmt"
    "time"
    "github.com/lazygophers/utils/singledo"
)

type ProcessingResult struct {
    Hash      string
    Size      int
    Processed time.Time
}

func main() {
    // Cache processing results for 30 minutes
    processor := singledo.NewSingleGroup[*ProcessingResult](30 * time.Minute)

    processData := func(data []byte) func() (*ProcessingResult, error) {
        return func() (*ProcessingResult, error) {
            fmt.Printf("Processing %d bytes of data...\n", len(data))

            // Simulate expensive processing
            time.Sleep(2 * time.Second)

            hash := fmt.Sprintf("%x", md5.Sum(data))

            return &ProcessingResult{
                Hash:      hash,
                Size:      len(data),
                Processed: time.Now(),
            }, nil
        }
    }

    data1 := []byte("Hello, World!")
    data2 := []byte("Hello, World!") // Same content
    data3 := []byte("Different data")

    // Use content hash as cache key
    key1 := fmt.Sprintf("data:%x", md5.Sum(data1))
    key2 := fmt.Sprintf("data:%x", md5.Sum(data2))
    key3 := fmt.Sprintf("data:%x", md5.Sum(data3))

    result1, _ := processor.Do(key1, processData(data1))
    result2, _ := processor.Do(key2, processData(data2)) // Same key, cached result
    result3, _ := processor.Do(key3, processData(data3)) // Different key, new processing

    fmt.Printf("Result 1: %+v\n", result1)
    fmt.Printf("Result 2: %+v\n", result2)
    fmt.Printf("Result 3: %+v\n", result3)
    fmt.Printf("Results 1 and 2 are same instance: %v\n", result1 == result2)
}
```

## API Reference

### Single Type

#### `NewSingle[T any](wait time.Duration) *Single[T]`
Creates a new Single instance for type T with the specified cache duration.

**Parameters:**
- `wait`: Duration to cache successful results

**Returns:**
- `*Single[T]`: New Single instance

#### `(s *Single[T]) Do(fn func() (T, error)) (T, error)`
Executes the function if not already cached or in progress, returns cached result if available.

**Parameters:**
- `fn`: Function to execute (will be called at most once per cache window)

**Returns:**
- `T`: Result of the function or cached value
- `error`: Error from the function or nil

**Behavior:**
- If result is cached and not expired, returns cached value immediately
- If function is currently executing, waits for completion and returns the result
- If no cache and no execution in progress, executes the function
- Only successful results (error == nil) are cached

#### `(s *Single[T]) Reset()`
Clears the cached result, forcing the next call to execute the function.

### Group Type

#### `NewSingleGroup[T any](wait time.Duration) *Group[T]`
Creates a new Group instance for managing multiple cached operations by key.

**Parameters:**
- `wait`: Duration to cache successful results for each key

**Returns:**
- `*Group[T]`: New Group instance

#### `(g *Group[T]) Do(key string, fn func() (T, error)) (T, error)`
Executes the function for the given key if not cached or in progress.

**Parameters:**
- `key`: Unique identifier for the operation
- `fn`: Function to execute for this key

**Returns:**
- `T`: Result of the function or cached value for the key
- `error`: Error from the function or nil

**Behavior:**
- Each key maintains independent cache and execution state
- Keys are never automatically cleaned up (design choice for simplicity)
- Suitable for bounded key sets or short-lived applications

## Best Practices

### 1. Choose Appropriate Cache Duration
```go
// Short-lived data
userSession := singledo.NewSingle[*Session](5 * time.Minute)

// Configuration data
appConfig := singledo.NewSingle[*Config](1 * time.Hour)

// Static reference data
currencies := singledo.NewSingle[[]Currency](24 * time.Hour)
```

### 2. Handle Errors Appropriately
```go
result, err := single.Do(func() (string, error) {
    // Only successful results are cached
    if someCondition {
        return "", errors.New("temporary failure") // Not cached
    }
    return "success", nil // This will be cached
})

if err != nil {
    // Handle error - result will be zero value
    log.Printf("Operation failed: %v", err)
    return
}
```

### 3. Use Meaningful Keys for Groups
```go
group := singledo.NewSingleGroup[*User](10 * time.Minute)

// Good - descriptive and unique
userKey := fmt.Sprintf("user:id:%d", userID)
profileKey := fmt.Sprintf("profile:user:%d:full", userID)

// Avoid - too generic or collision-prone
badKey := fmt.Sprintf("%d", userID)
```

### 4. Consider Memory Usage with Groups
```go
// Groups never clean up keys automatically
// For unbounded key sets, implement cleanup or use Single instances

// Good for bounded sets
userCache := singledo.NewSingleGroup[*User](time.Hour)

// Consider alternatives for unbounded sets
// Or implement your own cleanup mechanism
```

## Performance Considerations

- **Memory**: Groups retain references to all keys ever used
- **Concurrency**: Minimal lock contention with efficient RWMutex usage
- **CPU**: Near-zero overhead for cache hits
- **Goroutines**: No goroutines created, synchronous execution model

## Thread Safety

The singledo module is fully thread-safe:

- Multiple goroutines can safely call `Do()` concurrently
- Only one execution per key will occur at a time
- Concurrent calls to the same operation will wait and receive the same result
- `Reset()` is safe to call concurrently (though it may cause cache misses)

## Error Handling

- Only successful results (error == nil) are cached
- Errors are returned immediately and not cached
- Panics in executed functions are recovered and returned as errors
- Cache state is properly maintained even when errors occur

## Use Cases

1. **API Response Caching**: Reduce external API calls
2. **Database Query Optimization**: Cache expensive queries
3. **Computational Caching**: Cache results of heavy computations
4. **Resource Initialization**: Ensure resources are initialized only once
5. **Configuration Loading**: Cache configuration data
6. **File Processing**: Cache results of file parsing/processing
7. **Authentication**: Cache user authentication/authorization results

## Comparison with Other Patterns

### vs sync.Once
- **Advantage**: Time-based expiration, error handling, multiple executions over time
- **Use Case**: When you need periodic re-execution rather than one-time initialization

### vs Manual Caching
- **Advantage**: Thread-safe, deduplication of concurrent calls, built-in expiration
- **Use Case**: When you need more than simple map-based caching

### vs Channel-based Patterns
- **Advantage**: Simpler API, no goroutine management, immediate results for cache hits
- **Use Case**: When you don't need complex flow control or async execution

## Related Packages

- [`sync`](https://pkg.go.dev/sync): Standard synchronization primitives
- [`context`](https://pkg.go.dev/context): For timeout and cancellation (consider adding context support)
- [`time`](https://pkg.go.dev/time): Time-based functionality