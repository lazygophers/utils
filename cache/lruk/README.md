# LRU-K Cache

A thread-safe LRU-K cache implementation that tracks the K most recent access times for each item, providing better replacement decisions than traditional LRU for workloads with more complex temporal patterns.

## Features

- **K-distance tracking** for superior replacement decisions
- **Dual-tier architecture** with history and main cache
- **Configurable K value** for different workload optimization
- **O(1)** operations for all cache methods
- **Thread-safe** with optimized locking
- **Generic support** for any comparable key and any value type

## How LRU-K Works

LRU-K maintains two data structures:

1. **History Buffer**: Tracks items with fewer than K accesses
2. **Main Cache**: Contains items with K or more accesses

The algorithm considers the **K-th most recent access time** (K-distance) rather than just the most recent access when making eviction decisions. This provides better resistance to temporal locality violations.

## Use Cases

- **Database buffer pools** where LRU-2 is proven superior
- **File system caches** with mixed access patterns  
- **Application caches** with both sequential and random access
- **Storage systems** where K=2 reduces cache pollution
- **Analytics workloads** with scan + lookup patterns

## API

```go
// Create LRU-K cache with capacity and K value
cache := lruk.New[string, int](capacity, k)

// Common configurations
cache := lruk.New[string, int](1000, 2)  // LRU-2 (most common)
cache := lruk.New[string, int](1000, 3)  // LRU-3
cache := lruk.New[string, int](1000, 1)  // LRU-1 (equivalent to LRU)

// With eviction callback
cache := lruk.NewWithEvict[string, int](capacity, k, func(key string, value int) {
    fmt.Printf("Evicted: %s = %d\n", key, value)
})

// Basic operations
cache.Put("key", 42)           // Add item (counts as 1st access)
value, ok := cache.Get("key")  // Retrieve and record access
cache.Remove("key")            // Remove specific item
cache.Contains("key")          // Check if in main cache
value, ok := cache.Peek("key") // Get without recording access

// LRU-K specific
k := cache.GetK()             // Get the K value
stats := cache.Stats()        // Cache and history statistics

// Management
cache.Clear()                 // Remove all items
cache.Resize(newSize)        // Change capacity
cache.Len()                  // Items in main cache only
cache.Cap()                  // Maximum capacity
```

## Example

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/lruk"
)

func main() {
    // Create LRU-2 cache (K=2)
    cache := lruk.New[string, int](5, 2)
    
    // Add items (each Put counts as first access)
    cache.Put("A", 1)  // A: 1 access, in history
    cache.Put("B", 2)  // B: 1 access, in history  
    cache.Put("C", 3)  // C: 1 access, in history
    
    // Access A again - promotes to main cache (2 accesses)
    cache.Get("A")     // A: 2 accesses, promoted to main cache
    
    // Access B again - promotes to main cache
    cache.Get("B")     // B: 2 accesses, promoted to main cache
    
    // Add more items
    cache.Put("D", 4)  // D: 1 access, in history
    cache.Put("E", 5)  // E: 1 access, in history
    
    // Now add items to fill cache and trigger evictions
    cache.Put("F", 6)  // F: 1 access, in history
    cache.Put("G", 7)  // G: 1 access, in history
    
    // Access G again to try to promote it
    cache.Get("G")     // G: 2 accesses, but main cache might be full
    
    // Check what's in main cache vs history
    stats := cache.Stats()
    fmt.Printf("Main cache size: %d\n", stats.Size)
    fmt.Printf("History size: %d\n", stats.HistorySize)
    fmt.Printf("Total entries: %d\n", stats.TotalEntries)
    
    // Items in main cache have 2+ accesses
    fmt.Printf("A in main cache: %v\n", cache.Contains("A"))
    fmt.Printf("B in main cache: %v\n", cache.Contains("B"))
}
```

## K Value Selection

The choice of K significantly affects cache behavior:

| K Value | Behavior | Use Case | Memory Overhead |
|---------|----------|----------|-----------------|
| **K=1** | Same as LRU | Simple temporal locality | Minimal |
| **K=2** | Best general purpose | Database buffers, mixed workloads | Low |
| **K=3** | More conservative | Scan-resistant applications | Medium |
| **K=5+** | Very conservative | Highly irregular access patterns | Higher |

### LRU-2 (K=2) - Most Popular

LRU-2 is widely used because it:
- **Reduces cache pollution** from one-time accesses
- **Maintains good hit rates** for repeated accesses
- **Has low overhead** compared to higher K values
- **Is proven effective** in database buffer pool research

## Performance

- **Get**: O(1) - Hash lookup + access time tracking
- **Put**: O(1) - Hash insert + potential promotion
- **Remove**: O(1) - Hash delete + structure cleanup
- **Memory**: O(n × K) for access time tracking
- **Promotion**: O(1) when moving between history and cache

## Benchmarks

```
BenchmarkLRUKPut-8      6000000    250 ns/op    64 B/op    1 allocs/op
BenchmarkLRUKGet-8     10000000    180 ns/op     8 B/op    0 allocs/op
BenchmarkLRUKMixed-8    7000000    215 ns/op    36 B/op    0 allocs/op
```

## Architecture Details

### Two-Tier Design

```
┌─────────────────┐    ┌──────────────────┐
│   History       │    │   Main Cache     │
│   (< K access)  │───▶│   (≥ K access)   │
│                 │    │                  │
│ New items enter │    │ Promoted items   │
│ One-time scans  │    │ Frequently used  │
└─────────────────┘    └──────────────────┘
```

### Access Tracking

Each entry maintains up to K timestamps:

```go
type entry struct {
    accessTimes []time.Time  // Last K access times
    inCache     bool         // In main cache or history
    // ... other fields
}
```

### Promotion Logic

An item is promoted from history to main cache when:
1. It receives its K-th access
2. There's space in the main cache, OR
3. It can displace a less recently accessed item in main cache

## Statistics

Monitor cache efficiency and behavior:

```go
stats := cache.Stats()
fmt.Printf("LRU-K Cache Statistics:\n")
fmt.Printf("  Main cache: %d/%d items\n", stats.Size, stats.Capacity)
fmt.Printf("  History buffer: %d items\n", stats.HistorySize)
fmt.Printf("  Total tracked: %d items\n", stats.TotalEntries)
fmt.Printf("  K value: %d\n", stats.K)

// Calculate promotion rate
promotionRate := float64(stats.Size) / float64(stats.TotalEntries) * 100
fmt.Printf("  Promotion rate: %.1f%%\n", promotionRate)
```

## When to Use LRU-K

### ✅ Excellent for:
- **Database buffer pools** (LRU-2 is industry standard)
- **File system caches** with mixed access patterns
- **Application caches** with scan + lookup workloads
- **Storage systems** needing scan resistance
- **Analytics platforms** with batch + interactive queries

### ✅ Better than LRU when:
- **One-time scans** shouldn't evict valuable data
- **Access patterns are more complex** than simple recency
- **Cache pollution** from irregular accesses is a problem
- **Historical research** shows LRU-K performs better

### ❌ Consider alternatives for:
- **Pure LRU workloads** (unnecessary overhead)
- **Memory-constrained** environments (K × access time overhead)
- **Very dynamic** access patterns (consider Adaptive algorithms)
- **Small caches** where overhead isn't justified

## Comparison with Other Algorithms

| Algorithm | Scan Resistance | Memory Overhead | Complexity | Research Backing |
|-----------|----------------|-----------------|------------|------------------|
| **LRU** | Poor | Minimal | Simple | Basic |
| **LRU-K** | Good | Medium | Medium | Strong (databases) |
| **SLRU** | Good | Low | Medium | Moderate |
| **LFU** | Excellent | High | High | Strong |

## Research Background

LRU-K, particularly LRU-2, has strong theoretical and empirical support:

- **Database research** shows LRU-2 significantly outperforms LRU
- **Buffer pool studies** demonstrate reduced I/O with LRU-2
- **Theoretical analysis** proves better worst-case behavior
- **Industry adoption** in major database systems

LRU-K strikes an excellent balance between improved hit rates and manageable complexity, making it ideal for systems where traditional LRU isn't quite good enough but full LFU is overkill.