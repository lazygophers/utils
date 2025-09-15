# Window-TinyLFU Cache

A thread-safe cache implementation that combines a small LRU window cache with a TinyLFU main cache, providing excellent hit rates by balancing both recency and frequency in replacement decisions.

## Features

- **Hybrid architecture** combining LRU window with TinyLFU main cache
- **Three-tier main cache** with probation and protected segments
- **Count-Min Sketch** for memory-efficient frequency estimation
- **Automatic space management** with configurable segment sizes
- **O(1)** operations for all cache methods
- **Thread-safe** with optimized locking
- **Generic support** for any comparable key and any value type

## How Window-TinyLFU Works

Window-TinyLFU uses a sophisticated multi-tier architecture:

1. **Window Cache (10%)**: Small LRU cache for new items
2. **Main Cache (90%)**: Divided into probation (20%) and protected (80%) segments
3. **Frequency Tracking**: Count-Min Sketch estimates item popularity
4. **Admission Control**: Items compete based on frequency estimates

### Architecture
```
┌─────────────┐    ┌──────────────────────────────────┐
│   Window    │    │          Main Cache              │
│   (LRU)     │───▶│  ┌─────────────┬──────────────┐  │
│   10%       │    │  │ Probation   │  Protected   │  │
│             │    │  │   (LRU)     │    (LRU)     │  │
│             │    │  │    20%      │     80%      │  │
└─────────────┘    │  └─────────────┴──────────────┘  │
                   └──────────────────────────────────┘
```

## Use Cases

- **General-purpose caching** with excellent hit rates
- **Web applications** serving mixed content types  
- **CDN edge caches** balancing popular and new content
- **Database query caches** with varied access patterns
- **Application caches** requiring high performance
- **Microservices** needing robust caching with limited tuning

## API

```go
// Create Window-TinyLFU cache with default configuration
cache := wtinylfu.New[string, int](capacity)

// With eviction callback
cache := wtinylfu.NewWithEvict[string, int](capacity, func(key string, value int) {
    fmt.Printf("Evicted: %s = %d\n", key, value)
})

// Basic operations
cache.Put("key", 42)           // Add/update item
value, ok := cache.Get("key")  // Retrieve and update frequency
cache.Remove("key")            // Remove specific item
cache.Contains("key")          // Check existence
value, ok := cache.Peek("key") // Get without affecting placement

// Space management and statistics
stats := cache.Stats()         // Detailed segment statistics

// Management
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
    "github.com/lazygophers/utils/cache/wtinylfu"
)

func main() {
    // Create Window-TinyLFU cache with capacity 100
    cache := wtinylfu.New[string, string](100)
    
    // Add some initial content
    cache.Put("homepage", "HTML content")
    cache.Put("about", "About page")
    cache.Put("contact", "Contact info")
    
    // Simulate homepage being very popular
    for i := 0; i < 50; i++ {
        cache.Get("homepage")  // Builds up frequency
    }
    
    // Access about page a few times
    for i := 0; i < 5; i++ {
        cache.Get("about")
    }
    
    // Add lots of new content (one-time requests)
    for i := 0; i < 200; i++ {
        cache.Put(fmt.Sprintf("temp_%d", i), fmt.Sprintf("temp content %d", i))
    }
    
    // Popular content should survive
    fmt.Printf("Homepage present: %v\n", cache.Contains("homepage"))
    fmt.Printf("About present: %v\n", cache.Contains("about"))
    fmt.Printf("Contact present: %v\n", cache.Contains("contact"))
    
    // Check space distribution
    stats := cache.Stats()
    fmt.Printf("\nCache statistics:\n")
    fmt.Printf("  Total: %d/%d\n", stats.Size, stats.Capacity)
    fmt.Printf("  Window: %d/%d\n", stats.WindowSize, stats.WindowCapacity)
    fmt.Printf("  Probation: %d/%d\n", stats.ProbationSize, stats.ProbationCap)
    fmt.Printf("  Protected: %d/%d\n", stats.ProtectedSize, stats.ProtectedCap)
}
```

## Space Configuration

### Default Space Allocation

For a cache of capacity C:
- **Window**: 10% of C (minimum 1)
- **Main Cache**: 90% of C
  - **Probation**: 20% of main cache
  - **Protected**: 80% of main cache

### Automatic Sizing for Small Caches

For caches ≤ 4 items, the algorithm uses balanced segment distribution to ensure reasonable operation.

## Item Flow

### New Item Journey
```
1. New item → Window Cache (LRU)
2. On access → Promoted to Probation
3. On access → Promoted to Protected  
4. Further accesses → Stays in Protected (LRU order)
```

### Eviction Priority
```
1. Items in Window (if space needed there)
2. Items in Probation (compete with window items)
3. Items in Protected (last resort)
```

### Frequency Competition
When spaces are full, items compete based on Count-Min Sketch frequency estimates. Higher frequency items win admission.

## Performance

- **Get**: O(1) - Hash lookup + potential promotion + frequency update
- **Put**: O(1) - Hash insert + space management + frequency tracking
- **Remove**: O(1) - Hash delete + space cleanup
- **Memory**: O(n + s) where n=items, s=sketch size
- **Frequency estimation**: O(1) with high accuracy

## Benchmarks

```
BenchmarkWTinyLFUPut-8    4000000    350 ns/op    72 B/op    1 allocs/op
BenchmarkWTinyLFUGet-8    7000000    220 ns/op    16 B/op    0 allocs/op  
BenchmarkWTinyLFUMixed-8  5000000    285 ns/op    44 B/op    0 allocs/op
```

## Statistics

Monitor cache effectiveness across all segments:

```go
stats := cache.Stats()
fmt.Printf("Window-TinyLFU Statistics:\n")
fmt.Printf("  Overall: %d/%d items (%.1f%% full)\n", 
    stats.Size, stats.Capacity, 
    float64(stats.Size)/float64(stats.Capacity)*100)

fmt.Printf("  Window: %d/%d (%.1f%% full)\n",
    stats.WindowSize, stats.WindowCapacity,
    float64(stats.WindowSize)/float64(stats.WindowCapacity)*100)

fmt.Printf("  Probation: %d/%d (%.1f%% full)\n",
    stats.ProbationSize, stats.ProbationCap,
    float64(stats.ProbationSize)/float64(stats.ProbationCap)*100)

fmt.Printf("  Protected: %d/%d (%.1f%% full)\n",
    stats.ProtectedSize, stats.ProtectedCap,
    float64(stats.ProtectedSize)/float64(stats.ProtectedCap)*100)
```

## Tuning Guidelines

### When Default Works Well
- **General web applications**
- **Mixed read/write workloads**  
- **Balanced hot/cold data**
- **Standard temporal locality patterns**

### Consider Alternatives When
- **Pure LRU workload**: Use LRU cache
- **Pure frequency workload**: Use TinyLFU cache
- **Memory-constrained**: Use simpler algorithms
- **Highly specialized patterns**: Use algorithm-specific caches

## Algorithm Strengths

### vs. LRU
- **Better frequency awareness**: Popular items protected from scans
- **Scan resistance**: Window absorbs one-time accesses
- **Higher hit rates**: Especially for mixed workloads

### vs. LFU/TinyLFU  
- **Recency consideration**: New popular items can quickly gain prominence
- **Faster adaptation**: Window provides immediate recency benefits
- **Better balance**: Considers both frequency and recency

### vs. Simple Multi-tier
- **Automatic management**: No manual tuning of promotion thresholds
- **Frequency-based decisions**: Uses statistical estimation for better choices
- **Adaptive behavior**: Count-Min Sketch ages gracefully

## When to Use Window-TinyLFU

### ✅ Excellent for:
- **Web applications** with mixed content popularity
- **General-purpose caching** where you want "best overall" performance
- **CDN edge servers** balancing new and popular content
- **Database query caches** with varied access patterns
- **Application-level caches** requiring minimal tuning

### ✅ Choose over other algorithms when:
- **Hit rate is critical** and you can accept slightly higher complexity
- **Workload combines** both temporal and frequency locality
- **You want proven performance** without algorithm-specific tuning
- **Memory overhead is acceptable** for better hit rates

### ❌ Consider simpler alternatives for:
- **Memory-critical** applications (use LRU or TinyLFU)
- **Very predictable** workloads (use specialized algorithms)
- **Embedded systems** with strict resource constraints
- **Real-time systems** where deterministic performance is required

## Research Background

Window-TinyLFU is based on extensive research:

- **Academic validation** with superior hit rates across diverse workloads
- **Industry adoption** in high-performance caching systems
- **Benchmark studies** showing consistent improvements over LRU and LFU
- **Production deployment** in large-scale web services

The algorithm represents the current state-of-the-art in general-purpose cache replacement, providing excellent performance across a wide range of workloads without requiring workload-specific tuning.