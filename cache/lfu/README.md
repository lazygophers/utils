# LFU Cache

[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/utils/cache/lfu)](https://goreportcard.com/report/github.com/lazygophers/utils/cache/lfu)
[![Coverage Status](https://img.shields.io/badge/coverage-96.9%25-brightgreen)](https://github.com/lazygophers/utils/cache/lfu)

A high-performance, thread-safe Least Frequently Used (LFU) cache implementation in Go with full generic type support.

## Features

- **Generic Type Support**: Full support for any comparable key type and any value type using Go generics
- **Thread-Safe**: All operations are protected by read-write mutexes for concurrent access
- **High Performance**: O(1) time complexity for most operations with frequency-based eviction
- **Memory Efficient**: Uses hash maps and linked lists for optimal memory usage
- **Frequency Tracking**: Accurate frequency counting for intelligent eviction decisions
- **Configurable Capacity**: Support for dynamic resizing and capacity management
- **Eviction Callbacks**: Optional callback functions for handling evicted entries
- **Comprehensive API**: Rich set of methods for cache inspection and frequency management

## Installation

```bash
go get github.com/lazygophers/utils/cache/lfu
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/lfu"
)

func main() {
    // Create a new LFU cache with capacity 100
    cache := lfu.New[string, int](100)
    
    // Add items
    cache.Put("key1", 42)
    cache.Put("key2", 84)
    
    // Access items to increase frequency
    cache.Get("key1") // key1 frequency: 2
    cache.Get("key1") // key1 frequency: 3
    cache.Get("key2") // key2 frequency: 2
    
    // Check frequency
    fmt.Printf("key1 frequency: %d\\n", cache.GetFreq("key1")) // Output: 3
    
    // When cache is full, least frequently used items are evicted first
    fmt.Printf("Cache size: %d\\n", cache.Len())
}
```

## How LFU Works

The Least Frequently Used (LFU) cache evicts items that have been accessed the least number of times. Key characteristics:

1. **Frequency Tracking**: Each cache entry maintains an access frequency counter
2. **Eviction Policy**: When the cache is full, the item with the lowest frequency is removed
3. **Tie Breaking**: Among items with the same frequency, the least recently used is evicted
4. **Frequency Updates**: Every `Get` and `Put` (for existing keys) increments the frequency

## API Reference

### Creating a Cache

```go
// Create cache with specified capacity
cache := lfu.New[string, int](capacity)

// Create cache with eviction callback
cache := lfu.NewWithEvict[string, int](capacity, func(key string, value int) {
    fmt.Printf("Evicted: %s -> %d (freq: %d)\\n", key, value, freq)
})
```

### Basic Operations

```go
// Put adds or updates a value (increments frequency for existing keys)
evicted := cache.Put("key", 42)

// Get retrieves a value and increments its frequency
value, found := cache.Get("key")

// Remove deletes a key-value pair
value, found := cache.Remove("key")

// Contains checks if a key exists without affecting its frequency
exists := cache.Contains("key")

// Peek gets a value without updating its frequency
value, found := cache.Peek("key")
```

### Frequency Management

```go
// Get the access frequency of a key
freq := cache.GetFreq("key")

// Get cache statistics including frequency distribution
stats := cache.Stats()
fmt.Printf("Min frequency: %d\\n", stats.MinFreq)
fmt.Printf("Frequency distribution: %v\\n", stats.FreqDistribution)
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
```

### Inspection Methods

```go
// Get all keys
keys := cache.Keys()

// Get all values
values := cache.Values()

// Get all key-value pairs as a map
items := cache.Items()
```

## Performance

### Benchmark Results

```
BenchmarkPut-8      	 8000000	       150 ns/op	      64 B/op	       3 allocs/op
BenchmarkGet-8      	15000000	       110 ns/op	       0 B/op	       0 allocs/op
BenchmarkPutGet-8   	 4000000	       260 ns/op	      64 B/op	       3 allocs/op
```

### Time Complexity

| Operation | Time Complexity | Notes |
|-----------|-----------------|-------|
| Get       | O(1)           | Increments frequency |
| Put       | O(1) amortized | May trigger eviction |
| Remove    | O(1)           | Updates frequency lists |
| Contains  | O(1)           | No frequency change |
| Peek      | O(1)           | No frequency change |
| GetFreq   | O(1)           | Direct frequency lookup |
| Clear     | O(n)           | Clears all structures |
| Keys      | O(n)           | Iterates all entries |
| Values    | O(n)           | Iterates all entries |
| Items     | O(n)           | Iterates all entries |

## Test Coverage

**96.9% test coverage** - Comprehensive testing including:

- Basic functionality tests
- Frequency tracking validation
- Edge case handling
- Concurrent access testing
- Memory management verification
- Eviction policy validation
- Complex frequency scenarios

### Coverage Report

```
github.com/lazygophers/utils/cache/lfu/lfu.go:27:     New             100.0%
github.com/lazygophers/utils/cache/lfu/lfu.go:41:     NewWithEvict    100.0%
github.com/lazygophers/utils/cache/lfu/lfu.go:48:     Get             100.0%
github.com/lazygophers/utils/cache/lfu/lfu.go:62:     Put             100.0%
github.com/lazygophers/utils/cache/lfu/lfu.go:99:     Remove          100.0%
github.com/lazygophers/utils/cache/lfu/lfu.go:113:    Contains        100.0%
github.com/lazygophers/utils/cache/lfu/lfu.go:122:    Peek            100.0%
github.com/lazygophers/utils/cache/lfu/lfu.go:135:    Len             100.0%
github.com/lazygophers/utils/cache/lfu/lfu.go:143:    Cap             100.0%
github.com/lazygophers/utils/cache/lfu/lfu.go:148:    Clear           100.0%
github.com/lazygophers/utils/cache/lfu/lfu.go:164:    Keys            100.0%
github.com/lazygophers/utils/cache/lfu/lfu.go:176:    Values          100.0%
github.com/lazygophers/utils/cache/lfu/lfu.go:188:    Items           100.0%
github.com/lazygophers/utils/cache/lfu/lfu.go:200:    Resize          100.0%
github.com/lazygophers/utils/cache/lfu/lfu.go:217:    GetFreq         100.0%
github.com/lazygophers/utils/cache/lfu/lfu.go:228:    incrementFreq   100.0%
github.com/lazygophers/utils/cache/lfu/lfu.go:250:    evictLFU        85.7%
github.com/lazygophers/utils/cache/lfu/lfu.go:266:    removeEntry     100.0%
github.com/lazygophers/utils/cache/lfu/lfu.go:286:    updateMinFreq   70.0%
github.com/lazygophers/utils/cache/lfu/lfu.go:308:    Stats           100.0%
total:                                                (statements)    96.9%
```

## Thread Safety

All operations are thread-safe and can be called concurrently from multiple goroutines. The cache uses read-write mutexes to ensure:

- Multiple concurrent reads are allowed
- Writes are exclusive and block all other operations
- Frequency updates are atomic and consistent
- No data races or corruption can occur

## Examples

### Example 1: Basic LFU Behavior

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/lfu"
)

func main() {
    cache := lfu.New[string, string](3)
    
    cache.Put("a", "Apple")   // freq: 1
    cache.Put("b", "Banana")  // freq: 1
    cache.Put("c", "Cherry")  // freq: 1
    
    // Make "a" and "b" more frequently used
    cache.Get("a") // freq: 2
    cache.Get("a") // freq: 3
    cache.Get("b") // freq: 2
    
    // Add "d" - this will evict "c" (least frequently used)
    cache.Put("d", "Date")
    
    // "c" is no longer in cache
    _, found := cache.Get("c")
    fmt.Printf("Found 'c': %t\\n", found) // Output: Found 'c': false
    
    // Check frequencies
    fmt.Printf("'a' frequency: %d\\n", cache.GetFreq("a")) // Output: 4 (3 + 1 from Get check)
    fmt.Printf("'b' frequency: %d\\n", cache.GetFreq("b")) // Output: 2
    fmt.Printf("'d' frequency: %d\\n", cache.GetFreq("d")) // Output: 1
}
```

### Example 2: Frequency Tracking

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/lfu"
)

func main() {
    cache := lfu.New[int, string](5)
    
    // Add items with different access patterns
    cache.Put(1, "one")
    cache.Put(2, "two")
    cache.Put(3, "three")
    
    // Create different frequency patterns
    for i := 0; i < 5; i++ {
        cache.Get(1) // Make 1 very frequent
    }
    
    for i := 0; i < 2; i++ {
        cache.Get(2) // Make 2 moderately frequent  
    }
    
    // 3 remains least frequent (freq: 1)
    
    stats := cache.Stats()
    fmt.Printf("Min frequency: %d\\n", stats.MinFreq)
    fmt.Printf("Frequency distribution: %v\\n", stats.FreqDistribution)
    
    // When cache fills up, item 3 will be evicted first
}
```

### Example 3: Using Peek vs Get

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/lfu"
)

func main() {
    cache := lfu.New[string, int](3)
    
    cache.Put("key", 42)
    fmt.Printf("Initial frequency: %d\\n", cache.GetFreq("key")) // Output: 1
    
    // Get increments frequency
    cache.Get("key")
    fmt.Printf("After Get: %d\\n", cache.GetFreq("key")) // Output: 2
    
    // Peek does not increment frequency
    cache.Peek("key")
    fmt.Printf("After Peek: %d\\n", cache.GetFreq("key")) // Output: 2 (unchanged)
}
```

### Example 4: Cache Statistics

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/lfu"
)

func main() {
    cache := lfu.New[string, int](10)
    
    // Add items with varying access patterns
    for i := 0; i < 5; i++ {
        key := fmt.Sprintf("key%d", i)
        cache.Put(key, i)
        
        // Create different frequencies
        for j := 0; j <= i; j++ {
            cache.Get(key)
        }
    }
    
    stats := cache.Stats()
    fmt.Printf("Cache size: %d/%d\\n", stats.Size, stats.Capacity)
    fmt.Printf("Minimum frequency: %d\\n", stats.MinFreq)
    fmt.Printf("Frequency distribution:\\n")
    
    for freq, count := range stats.FreqDistribution {
        fmt.Printf("  Frequency %d: %d items\\n", freq, count)
    }
}
```

## Comparison with LRU

| Aspect | LRU | LFU |
|--------|-----|-----|
| Eviction Policy | Least Recently Used | Least Frequently Used |
| Time Complexity | O(1) all operations | O(1) most operations |
| Memory Overhead | Lower | Higher (frequency tracking) |
| Use Case | Temporal locality | Frequency-based patterns |
| Access Pattern | Recent access matters | Total access count matters |

Choose LFU when:
- You have clear frequency-based access patterns
- Some items are accessed much more frequently than others
- You want to keep "hot" data in cache regardless of recent access

Choose LRU when:
- You have temporal locality in access patterns
- Recent access is more important than total access count
- You want simpler, faster cache operations

## License

This project is licensed under the MIT License - see the [LICENSE](../../../LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.