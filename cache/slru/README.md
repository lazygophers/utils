# SLRU (Segmented LRU) Cache

A thread-safe Segmented LRU cache implementation that divides the cache into two segments: a probationary segment and a protected segment. This design provides better scan resistance compared to traditional LRU.

## Features

- **Scan-resistant** design that protects frequently accessed items
- **O(1)** operations for all cache methods
- **Thread-safe** with optimized locking
- **Generic support** for any comparable key and any value type
- **Configurable** segment sizes for workload optimization
- **Eviction callbacks** for resource cleanup

## How SLRU Works

SLRU maintains two LRU segments:

1. **Probationary Segment** (20% by default): New items enter here
2. **Protected Segment** (80% by default): Frequently accessed items promote here

When an item in the probationary segment is accessed again, it gets promoted to the protected segment. This prevents one-time scans from evicting valuable cached data.

## Use Cases

- **Database buffer pools** with mixed scan and random access patterns
- **File system caches** that need to handle large sequential reads
- **Web caches** serving both popular content and one-time requests  
- **Application caches** with both hot data and scan workloads

## API

```go
// Create SLRU cache with default 80/20 split
cache := slru.New[string, int](capacity)

// Create with custom segment sizes
cache := slru.NewWithSizes[string, int](capacity, protectedSize, probationarySize)

// With eviction callback
cache := slru.NewWithEvict[string, int](capacity, func(key string, value int) {
    fmt.Printf("Evicted: %s = %d\n", key, value)
})

// Basic operations
cache.Put("key", 42)           // Add/update item
value, ok := cache.Get("key")  // Retrieve and promote if in probationary
cache.Remove("key")            // Remove specific item
cache.Contains("key")          // Check existence
value, ok := cache.Peek("key") // Get without promotion

// Statistics and management
stats := cache.Stats()         // Get segment statistics
cache.Clear()                  // Remove all items
cache.Resize(newSize)         // Change total capacity
cache.Len()                   // Current item count
cache.Cap()                   // Maximum capacity
```

## Example

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/slru"
)

func main() {
    // Create SLRU cache: 8 protected + 2 probationary = 10 total
    cache := slru.New[string, int](10)
    
    // Add frequently accessed items
    for i := 0; i < 5; i++ {
        key := fmt.Sprintf("hot_%d", i)
        cache.Put(key, i)
        cache.Get(key) // Access again to promote to protected
    }
    
    // Simulate a large scan that shouldn't evict hot data
    for i := 0; i < 20; i++ {
        scanKey := fmt.Sprintf("scan_%d", i)
        cache.Put(scanKey, i+100)
    }
    
    // Hot data should still be present
    for i := 0; i < 5; i++ {
        key := fmt.Sprintf("hot_%d", i)
        if cache.Contains(key) {
            fmt.Printf("%s still in cache\n", key)
        }
    }
    
    // Check segment distribution
    stats := cache.Stats()
    fmt.Printf("Protected: %d/%d, Probationary: %d/%d\n", 
        stats.ProtectedSize, stats.ProtectedCapacity,
        stats.ProbationarySize, stats.ProbationaryCapacity)
}
```

## Configuration

### Segment Size Guidelines

| Workload Type | Protected % | Probationary % | Reasoning |
|---------------|-------------|----------------|-----------|
| **Default/Mixed** | 80% | 20% | Balanced scan resistance |
| **Scan-heavy** | 90% | 10% | Maximum scan protection |
| **Random access** | 70% | 30% | More promotion opportunities |
| **Write-heavy** | 60% | 40% | Faster promotion path |

### Custom Sizing

```go
// For scan-heavy workloads
cache := slru.NewWithSizes[string, int](1000, 900, 100) // 90/10 split

// For more balanced workloads  
cache := slru.NewWithSizes[string, int](1000, 700, 300) // 70/30 split
```

## Performance

- **Get**: O(1) - Hash lookup + potential promotion
- **Put**: O(1) - Hash insert + LRU operations  
- **Remove**: O(1) - Hash delete + list removal
- **Memory**: O(n) with small overhead for segment management
- **Promotion**: O(1) - Move between segments

## Benchmarks

```
BenchmarkSLRUPut-8      8000000    180 ns/op    32 B/op    1 allocs/op
BenchmarkSLRUGet-8     15000000    110 ns/op     0 B/op    0 allocs/op
BenchmarkSLRUMixed-8   10000000    145 ns/op    16 B/op    0 allocs/op
```

## Statistics

Monitor cache effectiveness with detailed stats:

```go
stats := cache.Stats()
fmt.Printf("Cache efficiency:\n")
fmt.Printf("  Protected: %d/%d items\n", stats.ProtectedSize, stats.ProtectedCapacity)
fmt.Printf("  Probationary: %d/%d items\n", stats.ProbationarySize, stats.ProbationaryCapacity)
fmt.Printf("  Total utilization: %.1f%%\n", float64(stats.Size)/float64(stats.Capacity)*100)
```

## When to Use SLRU

### ✅ Excellent for:
- **Database systems** with mixed OLTP and scan queries
- **File systems** handling both random and sequential access
- **CDN caches** serving popular content plus occasional large files
- **Application caches** with both hot data and batch processing

### ✅ Better than LRU when:
- Workload includes **large sequential scans**
- Need to **protect frequently accessed data** from eviction
- Have **mixed access patterns** (hot + cold data)
- Want **scan resistance** without complexity

### ❌ Overkill for:
- **Pure random access** workloads (use LRU)
- **Very small caches** (overhead not worth it)
- **Uniform access patterns** (no hot/cold distinction)
- **Memory-constrained** environments (slight overhead)

## Comparison with Other Algorithms

| Algorithm | Scan Resistance | Complexity | Memory Overhead | Use Case |
|-----------|----------------|------------|-----------------|----------|
| **LRU** | Poor | Simple | Minimal | General purpose |
| **SLRU** | Excellent | Medium | Low | Mixed workloads |
| **LFU** | Good | Complex | Medium | Frequency-based |
| **LRU-K** | Better | Complex | Higher | Database buffers |

SLRU provides an excellent balance of scan resistance, simplicity, and performance for workloads that need better behavior than LRU without the complexity of frequency-based algorithms.