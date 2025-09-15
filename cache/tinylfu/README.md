# TinyLFU Cache

A memory-efficient Least Frequently Used (LFU) cache implementation using a Count-Min Sketch for frequency estimation and periodic aging to handle changing access patterns.

## Features

- **Memory-efficient** frequency tracking using Count-Min Sketch
- **Probabilistic** frequency estimation with high accuracy
- **Automatic aging** to adapt to changing access patterns
- **O(1)** operations for all cache methods
- **Thread-safe** with optimized locking
- **Generic support** for any comparable key and any value type

## How TinyLFU Works

TinyLFU uses several key components:

1. **Count-Min Sketch**: Probabilistic data structure for frequency estimation
2. **Aging mechanism**: Periodic halving of all frequencies to handle pattern changes
3. **LRU ordering**: Items with same frequency are ordered by recency
4. **Compact storage**: Much smaller memory footprint than traditional LFU

## Use Cases

- **Large-scale caches** where memory efficiency is critical
- **Web caches** with millions of items and limited memory
- **CDN edge caches** that need to track popularity efficiently
- **Application caches** where frequency matters more than exact counts
- **Database query caches** with varying query patterns

## API

```go
// Create TinyLFU cache with default parameters
cache := tinylfu.New[string, int](capacity)

// With custom aging threshold (items accessed before aging)
cache := tinylfu.NewWithAging[string, int](capacity, agingThreshold)

// With eviction callback
cache := tinylfu.NewWithEvict[string, int](capacity, func(key string, value int) {
    fmt.Printf("Evicted: %s = %d\n", key, value)
})

// Basic operations
cache.Put("key", 42)           // Add/update item
value, ok := cache.Get("key")  // Retrieve and increment frequency
cache.Remove("key")            // Remove specific item
cache.Contains("key")          // Check existence
value, ok := cache.Peek("key") // Get without affecting frequency

// Frequency and aging
cache.ForceAging()            // Manually trigger aging
freq := cache.Frequency("key") // Get estimated frequency

// Management
cache.Clear()                 // Remove all items
cache.Resize(newSize)        // Change capacity
cache.Len()                  // Current item count
cache.Cap()                  // Maximum capacity
```

## Example

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/tinylfu"
)

func main() {
    // Create TinyLFU cache with capacity 1000
    cache := tinylfu.New[string, int](1000)
    
    // Add some items
    cache.Put("popular", 1)
    cache.Put("normal", 2)
    cache.Put("rare", 3)
    
    // Simulate different access frequencies
    for i := 0; i < 100; i++ {
        cache.Get("popular")  // High frequency
    }
    
    for i := 0; i < 10; i++ {
        cache.Get("normal")   // Medium frequency  
    }
    
    cache.Get("rare")        // Low frequency
    
    // Check estimated frequencies
    fmt.Printf("Popular frequency: %d\n", cache.Frequency("popular"))
    fmt.Printf("Normal frequency: %d\n", cache.Frequency("normal"))
    fmt.Printf("Rare frequency: %d\n", cache.Frequency("rare"))
    
    // Fill cache to capacity to trigger evictions
    for i := 0; i < 1000; i++ {
        cache.Put(fmt.Sprintf("item_%d", i), i)
    }
    
    // Popular item should still be present
    if cache.Contains("popular") {
        fmt.Println("Popular item survived eviction")
    }
}
```

## Configuration

### Aging Threshold

The aging threshold determines how many items are accessed before frequencies are halved:

```go
// Conservative aging (less frequent, more stable)
cache := tinylfu.NewWithAging[string, int](10000, 50000)

// Aggressive aging (more frequent, adapts faster)
cache := tinylfu.NewWithAging[string, int](10000, 5000)

// Default aging (balanced)
cache := tinylfu.New[string, int](10000) // threshold = capacity * 10
```

### Memory vs Accuracy Trade-offs

| Sketch Size | Memory Usage | Accuracy | Best For |
|-------------|--------------|----------|-----------|
| Small | Minimal | Good | Memory-constrained |
| Medium | Moderate | Very Good | Balanced workloads |
| Large | Higher | Excellent | Accuracy-critical |

## Performance

- **Get**: O(1) - Hash lookup + sketch update
- **Put**: O(1) - Hash insert + frequency-based eviction
- **Remove**: O(1) - Hash delete + list removal
- **Memory**: O(n + s) where n=items, s=sketch size (s << n)
- **Frequency estimation**: O(1) with high probability

## Benchmarks

```
BenchmarkTinyLFUPut-8     6000000    220 ns/op    48 B/op    1 allocs/op
BenchmarkTinyLFUGet-8    12000000    130 ns/op     0 B/op    0 allocs/op
BenchmarkTinyLFUMixed-8   8000000    175 ns/op    24 B/op    0 allocs/op
```

## Count-Min Sketch Details

The Count-Min Sketch provides probabilistic frequency estimation:

- **Width**: Proportional to cache capacity
- **Depth**: Fixed at 4 rows for good accuracy
- **Hash functions**: Independent hash functions per row
- **Error bound**: With high probability, estimates are close to true frequency
- **Space complexity**: O(width × depth) = O(capacity)

## Aging Mechanism

Aging prevents the cache from being biased toward historical access patterns:

1. **Trigger**: After every `agingThreshold` accesses
2. **Process**: All frequencies in the sketch are halved
3. **Effect**: Recent accesses become more influential
4. **Benefit**: Adapts to changing workload patterns

## When to Use TinyLFU

### ✅ Excellent for:
- **Large caches** (>10K items) where memory efficiency matters
- **Web applications** with varying content popularity
- **CDN systems** tracking millions of files
- **Database caches** with changing query patterns
- **Microservices** with limited memory budgets

### ✅ Better than LFU when:
- **Memory is constrained** and exact counts aren't needed
- **Access patterns change** over time
- **Cache size is large** relative to available memory
- **Approximate frequency** is sufficient for decisions

### ❌ Consider alternatives for:
- **Small caches** (<1K items) where exact LFU overhead is acceptable
- **Critical applications** where frequency accuracy is paramount
- **Static workloads** where patterns don't change
- **Debug scenarios** where exact frequencies are needed

## Comparison with Other Algorithms

| Algorithm | Memory Efficiency | Frequency Accuracy | Adaptability | Complexity |
|-----------|------------------|-------------------|--------------|------------|
| **LFU** | Poor | Perfect | Poor | Medium |
| **TinyLFU** | Excellent | Very Good | Excellent | Medium |
| **LRU** | Excellent | N/A | Excellent | Simple |
| **Adaptive LFU** | Poor | Perfect | Excellent | High |

TinyLFU provides the best balance of memory efficiency and frequency tracking for large-scale caching scenarios where traditional LFU would consume too much memory.