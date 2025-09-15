# MRU (Most Recently Used) Cache

A thread-safe MRU cache implementation that evicts the most recently used items when the cache reaches capacity. This is the opposite of LRU and is useful in scenarios where recently accessed items are less likely to be needed again.

## Features

- **O(1)** Get, Put, and Remove operations
- **Thread-safe** with optimized read-write locking
- **Generic support** for any comparable key and any value type
- **Eviction callbacks** for cleanup when items are removed
- **Memory efficient** with minimal overhead per item

## Use Cases

- **Streaming data processing** where recent items are less valuable
- **One-time scan workloads** where items won't be reaccessed
- **LIFO-like cache behavior** for stack-based data patterns
- **Anti-caching scenarios** where you want to keep older, stable data

## API

```go
// Create a new MRU cache
cache := mru.New[string, int](capacity)

// With eviction callback
cache := mru.NewWithEvict[string, int](capacity, func(key string, value int) {
    fmt.Printf("Evicted: %s = %d\n", key, value)
})

// Basic operations
cache.Put("key", 42)           // Returns true if eviction occurred
value, ok := cache.Get("key")  // Returns (value, true) if found
cache.Remove("key")            // Returns (value, true) if found
cache.Contains("key")          // Returns true if key exists
value, ok := cache.Peek("key") // Get without affecting order

// Bulk operations
keys := cache.Keys()           // All keys (most to least recent)
values := cache.Values()       // All values 
items := cache.Items()         // All key-value pairs

// Management
cache.Clear()                  // Remove all items
cache.Resize(newSize)         // Change capacity
cache.Len()                   // Current number of items
cache.Cap()                   // Maximum capacity
```

## Example

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/mru"
)

func main() {
    // Create MRU cache with capacity 3
    cache := mru.New[string, int](3)
    
    // Add items
    cache.Put("A", 1)
    cache.Put("B", 2) 
    cache.Put("C", 3)
    
    // Access A (makes it most recent)
    cache.Get("A")
    
    // Add D - will evict A (most recent)
    cache.Put("D", 4)
    
    fmt.Println(cache.Contains("A")) // false - A was evicted
    fmt.Println(cache.Contains("B")) // true - B is older
    fmt.Println(cache.Contains("C")) // true - C is older  
    fmt.Println(cache.Contains("D")) // true - D was just added
}
```

## Performance

- **Get**: O(1) - Direct hash table lookup + list move
- **Put**: O(1) - Hash table insert + list operations
- **Remove**: O(1) - Hash table delete + list removal
- **Memory**: O(n) where n is the number of items
- **Concurrency**: Optimized read-write locks for high throughput

## Benchmarks

```
BenchmarkMRUPut-8     10000000    150 ns/op    24 B/op    1 allocs/op
BenchmarkMRUGet-8     20000000     85 ns/op     0 B/op    0 allocs/op
BenchmarkMRUMixed-8   15000000    120 ns/op    12 B/op    0 allocs/op
```

## When to Use MRU

### ✅ Good for:
- **Sequential scan workloads** where data won't be reaccessed
- **Stream processing** where recent data has lower future probability
- **One-time data import/export** operations
- **Cache warming scenarios** where you want to preserve initially loaded data

### ❌ Avoid for:
- **General-purpose caching** (use LRU instead)
- **Hot data workloads** where recent items are frequently reaccessed
- **Interactive applications** with locality of reference
- **Database query caching** (recent queries often repeat)

## Comparison with LRU

| Aspect | MRU | LRU |
|--------|-----|-----|
| Eviction | Most recent first | Least recent first |
| Use case | Anti-locality workloads | Temporal locality workloads |
| Performance | O(1) all operations | O(1) all operations |
| Memory | Same overhead | Same overhead |
| Common | Rare, specialized | Very common |

MRU is essentially LRU with inverted eviction logic, making it suitable for very specific workload patterns where temporal locality assumptions don't hold.