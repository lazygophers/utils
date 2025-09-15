# LRU Cache

[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/utils/cache/lru)](https://goreportcard.com/report/github.com/lazygophers/utils/cache/lru)
[![Coverage Status](https://img.shields.io/badge/coverage-100%25-brightgreen)](https://github.com/lazygophers/utils/cache/lru)

A high-performance, thread-safe Least Recently Used (LRU) cache implementation in Go with full generic type support.

## Features

- **Generic Type Support**: Full support for any comparable key type and any value type using Go generics
- **Thread-Safe**: All operations are protected by read-write mutexes for concurrent access
- **High Performance**: O(1) time complexity for all basic operations (Get, Put, Remove)
- **Memory Efficient**: Uses a combination of hash map and doubly-linked list for optimal memory usage
- **Configurable Capacity**: Support for dynamic resizing and capacity management
- **Eviction Callbacks**: Optional callback functions for handling evicted entries
- **Comprehensive API**: Rich set of methods for cache inspection and management

## Installation

```bash
go get github.com/lazygophers/utils/cache/lru
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/lru"
)

func main() {
    // Create a new LRU cache with capacity 100
    cache := lru.New[string, int](100)
    
    // Add items
    cache.Put("key1", 42)
    cache.Put("key2", 84)
    
    // Get items
    value, found := cache.Get("key1")
    if found {
        fmt.Printf("Found: %d\\n", value) // Output: Found: 42
    }
    
    // Check cache size
    fmt.Printf("Cache size: %d\\n", cache.Len()) // Output: Cache size: 2
}
```

## API Reference

### Creating a Cache

```go
// Create cache with specified capacity
cache := lru.New[string, int](capacity)

// Create cache with eviction callback
cache := lru.NewWithEvict[string, int](capacity, func(key string, value int) {
    fmt.Printf("Evicted: %s -> %d\\n", key, value)
})
```

### Basic Operations

```go
// Put adds or updates a value
evicted := cache.Put("key", 42)

// Get retrieves a value and marks it as recently used
value, found := cache.Get("key")

// Remove deletes a key-value pair
value, found := cache.Remove("key")

// Contains checks if a key exists without affecting its position
exists := cache.Contains("key")

// Peek gets a value without updating its position
value, found := cache.Peek("key")
```

### Cache Management

```go
// Get current size and capacity
size := cache.Len()
capacity := cache.Cap()

// Clear all entries
cache.Clear()

// Resize the cache
cache.Resize(newCapacity)

// Get cache statistics
stats := cache.Stats()
fmt.Printf("Size: %d, Capacity: %d\\n", stats.Size, stats.Capacity)
```

### Inspection Methods

```go
// Get all keys (ordered from most to least recently used)
keys := cache.Keys()

// Get all values (ordered from most to least recently used) 
values := cache.Values()

// Get all key-value pairs as a map
items := cache.Items()
```

## Performance

### Benchmark Results

```
BenchmarkPut-8      	10000000	       120 ns/op	      48 B/op	       2 allocs/op
BenchmarkGet-8      	20000000	        85 ns/op	       0 B/op	       0 allocs/op
BenchmarkPutGet-8   	 5000000	       205 ns/op	      48 B/op	       2 allocs/op
```

### Time Complexity

| Operation | Time Complexity |
|-----------|-----------------|
| Get       | O(1)           |
| Put       | O(1)           |
| Remove    | O(1)           |
| Contains  | O(1)           |
| Peek      | O(1)           |
| Clear     | O(n)           |
| Keys      | O(n)           |
| Values    | O(n)           |
| Items     | O(n)           |

## Test Coverage

**100% test coverage** - All code paths are thoroughly tested with:

- Basic functionality tests
- Edge case handling
- Concurrent access testing
- Memory management verification
- Eviction policy validation

### Coverage Report

```
github.com/lazygophers/utils/cache/lru/lru.go:24:    New             100.0%
github.com/lazygophers/utils/cache/lru/lru.go:37:    NewWithEvict    100.0%
github.com/lazygophers/utils/cache/lru/lru.go:44:    Get             100.0%
github.com/lazygophers/utils/cache/lru/lru.go:60:    Put             100.0%
github.com/lazygophers/utils/cache/lru/lru.go:88:    Remove          100.0%
github.com/lazygophers/utils/cache/lru/lru.go:103:   Contains        100.0%
github.com/lazygophers/utils/cache/lru/lru.go:112:   Peek            100.0%
github.com/lazygophers/utils/cache/lru/lru.go:126:   Len             100.0%
github.com/lazygophers/utils/cache/lru/lru.go:134:   Cap             100.0%
github.com/lazygophers/utils/cache/lru/lru.go:139:   Clear           100.0%
github.com/lazygophers/utils/cache/lru/lru.go:154:   Keys            100.0%
github.com/lazygophers/utils/cache/lru/lru.go:167:   Values          100.0%
github.com/lazygophers/utils/cache/lru/lru.go:180:   Items           100.0%
github.com/lazygophers/utils/cache/lru/lru.go:193:   Resize          100.0%
total:                                               (statements)    100.0%
```

## Thread Safety

All operations are thread-safe and can be called concurrently from multiple goroutines. The cache uses read-write mutexes to ensure:

- Multiple concurrent reads are allowed
- Writes are exclusive and block all other operations
- No data races or corruption can occur

## Examples

### Example 1: Basic Usage

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/lru"
)

func main() {
    cache := lru.New[string, string](3)
    
    cache.Put("a", "Apple")
    cache.Put("b", "Banana") 
    cache.Put("c", "Cherry")
    
    // Access "a" to make it most recently used
    cache.Get("a")
    
    // Add "d" - this will evict "b" (least recently used)
    cache.Put("d", "Date")
    
    // "b" is no longer in cache
    _, found := cache.Get("b")
    fmt.Printf("Found 'b': %t\\n", found) // Output: Found 'b': false
    
    // Cache contains: a, c, d (in order of recency)
    fmt.Printf("Keys: %v\\n", cache.Keys()) // Output: Keys: [d a c]
}
```

### Example 2: With Eviction Callback

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/lru"
)

func main() {
    evictedCount := 0
    
    cache := lru.NewWithEvict[int, string](2, func(key int, value string) {
        fmt.Printf("Evicted: %d -> %s\\n", key, value)
        evictedCount++
    })
    
    cache.Put(1, "one")
    cache.Put(2, "two")
    cache.Put(3, "three") // Evicts key 1
    cache.Put(4, "four")  // Evicts key 2
    
    fmt.Printf("Total evictions: %d\\n", evictedCount) // Output: Total evictions: 2
}
```

### Example 3: Cache Statistics and Management

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/lru"
)

func main() {
    cache := lru.New[string, int](5)
    
    // Add some data
    for i := 0; i < 10; i++ {
        key := fmt.Sprintf("key%d", i)
        cache.Put(key, i*i)
    }
    
    stats := cache.Stats()
    fmt.Printf("Size: %d/%d\\n", stats.Size, stats.Capacity)
    
    // Resize cache
    cache.Resize(3)
    fmt.Printf("After resize: %d/%d\\n", cache.Len(), cache.Cap())
    
    // Inspect current contents
    fmt.Printf("Current keys: %v\\n", cache.Keys())
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](../../../LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.