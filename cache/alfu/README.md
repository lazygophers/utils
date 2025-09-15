# Adaptive LFU Cache

A thread-safe Adaptive Least Frequently Used cache that automatically adjusts frequencies over time using configurable decay mechanisms. This cache adapts to changing access patterns while maintaining the benefits of frequency-based replacement.

## Features

- **Time-based frequency decay** for pattern adaptation
- **Configurable decay parameters** for workload optimization  
- **Combined frequency and time factors** in replacement decisions
- **Automatic aging** to handle changing workloads
- **O(1)** operations for most cache methods
- **Thread-safe** with optimized locking
- **Generic support** for any comparable key and any value type

## How Adaptive LFU Works

Adaptive LFU extends traditional LFU with several key features:

1. **Frequency Tracking**: Like LFU, tracks access frequency for each item
2. **Time-based Decay**: Frequencies gradually decrease over time
3. **Access Time Recording**: Tracks when each item was last accessed
4. **Adaptive Eviction**: Considers both frequency and time since last access
5. **Periodic Aging**: Automatically reduces all frequencies to adapt to new patterns

## Use Cases

- **Dynamic web applications** with changing user behavior
- **Content streaming** where popularity shifts over time
- **Database query caches** with evolving query patterns
- **API gateways** adapting to changing endpoint popularity
- **Analytics systems** with seasonal or trending data patterns

## API

```go
// Create Adaptive LFU cache with default parameters
cache := alfu.New[string, int](capacity)

// With custom decay parameters
cache := alfu.NewWithConfig[string, int](
    capacity,
    0.8,                    // decay factor (80% retention)
    10 * time.Minute,      // decay interval
)

// With eviction callback
cache := alfu.NewWithEvict[string, int](capacity, func(key string, value int) {
    fmt.Printf("Evicted: %s = %d\n", key, value)
})

// Basic operations
cache.Put("key", 42)           // Add/update and increment frequency
value, ok := cache.Get("key")  // Retrieve and increment frequency  
cache.Remove("key")            // Remove specific item
cache.Contains("key")          // Check existence
value, ok := cache.Peek("key") // Get without affecting frequency

// Adaptive features
cache.ForceDecay()            // Manually trigger decay
stats := cache.Stats()        // Detailed statistics including decay info

// Management
cache.Clear()                 // Remove all items and reset decay
cache.Resize(newSize)        // Change capacity
cache.Len()                  // Current item count
cache.Cap()                  // Maximum capacity
```

## Example

```go
package main

import (
    "fmt"
    "time"
    "github.com/lazygophers/utils/cache/alfu"
)

func main() {
    // Create adaptive cache with aggressive decay for demonstration
    cache := alfu.NewWithConfig[string, int](
        100,                   // capacity
        0.5,                   // 50% decay factor (aggressive)
        1 * time.Second,       // 1 second decay interval
    )
    
    // Phase 1: Old popular content
    cache.Put("old_video", 1)
    for i := 0; i < 50; i++ {
        cache.Get("old_video")  // Very popular initially
    }
    
    cache.Put("old_article", 2)
    for i := 0; i < 20; i++ {
        cache.Get("old_article") // Moderately popular
    }
    
    fmt.Printf("Phase 1 - Old content frequencies:\n")
    stats := cache.Stats()
    fmt.Printf("  old_video freq level: high\n")
    fmt.Printf("  old_article freq level: medium\n")
    
    // Wait for decay to occur
    time.Sleep(2 * time.Second)
    
    // Phase 2: New trending content
    cache.Put("trending_video", 3)
    for i := 0; i < 30; i++ {
        cache.Get("trending_video") // New popular content
    }
    
    // Force decay to see adaptation
    cache.ForceDecay()
    
    // Add more content to trigger evictions
    for i := 0; i < 90; i++ {
        cache.Put(fmt.Sprintf("content_%d", i), i+100)
    }
    
    fmt.Printf("\nPhase 2 - After decay and new content:\n")
    fmt.Printf("  old_video present: %v\n", cache.Contains("old_video"))
    fmt.Printf("  trending_video present: %v\n", cache.Contains("trending_video"))
    
    // Check final statistics
    finalStats := cache.Stats()
    fmt.Printf("\nFinal statistics:\n")
    fmt.Printf("  Size: %d/%d\n", finalStats.Size, finalStats.Capacity)
    fmt.Printf("  Min frequency: %d\n", finalStats.MinFrequency)
    fmt.Printf("  Max frequency: %d\n", finalStats.MaxFrequency)
    fmt.Printf("  Last decay: %v ago\n", time.Since(finalStats.LastDecay))
}
```

## Configuration Parameters

### Decay Factor

Controls how much frequencies are reduced during decay:

| Decay Factor | Retention | Behavior | Use Case |
|--------------|-----------|----------|-----------|
| **0.9** | 90% | Conservative | Stable workloads |
| **0.8** | 80% | Balanced | General purpose |
| **0.7** | 70% | Moderate | Changing patterns |
| **0.5** | 50% | Aggressive | Rapidly changing |

```go
// Conservative: slow adaptation, stable patterns
cache := alfu.NewWithConfig[K, V](capacity, 0.9, 10*time.Minute)

// Aggressive: fast adaptation, trending content
cache := alfu.NewWithConfig[K, V](capacity, 0.5, 1*time.Minute)
```

### Decay Interval

Determines how often decay occurs:

| Interval | Frequency | Best For |
|----------|-----------|----------|
| **1 minute** | Very frequent | Real-time trending |
| **5 minutes** | Frequent | Social media |
| **30 minutes** | Moderate | Web applications |
| **2 hours** | Infrequent | Stable applications |

## Performance

- **Get**: O(1) - Hash lookup + frequency increment + potential decay
- **Put**: O(1) - Hash insert + frequency-based eviction
- **Remove**: O(1) - Hash delete + frequency list management  
- **Decay**: O(n) - Applied to all entries, but infrequent
- **Memory**: O(n + f) where n=items, f=unique frequencies

## Benchmarks

```
BenchmarkALFUPut-8      5000000    280 ns/op    56 B/op    1 allocs/op
BenchmarkALFUGet-8      8000000    200 ns/op     8 B/op    0 allocs/op
BenchmarkALFUMixed-8    6000000    240 ns/op    32 B/op    0 allocs/op
```

## Decay Algorithm

The decay process combines frequency and time factors:

### Frequency Decay
```
new_frequency = old_frequency × decay_factor
```

### Time-based Decay
```
time_factor = exp(-time_since_access / 1_hour)
effective_frequency = frequency × time_factor
```

### Combined Effect
Items that are both:
- **Less frequently accessed** AND
- **Not accessed recently**

Are most likely to be evicted.

## Statistics

Monitor adaptation and cache behavior:

```go
stats := cache.Stats()
fmt.Printf("Adaptive LFU Statistics:\n")
fmt.Printf("  Size: %d/%d\n", stats.Size, stats.Capacity)
fmt.Printf("  Frequency range: %d - %d\n", stats.MinFrequency, stats.MaxFrequency)
fmt.Printf("  Decay factor: %.2f\n", stats.DecayFactor)
fmt.Printf("  Decay interval: %v\n", stats.DecayInterval)
fmt.Printf("  Last decay: %v ago\n", time.Since(stats.LastDecay))

fmt.Printf("  Frequency distribution:\n")
for freq, count := range stats.FrequencyDistribution {
    fmt.Printf("    Freq %d: %d items\n", freq, count)
}
```

## Adaptation Patterns

### Trending Content Scenario
```
Initial: item A (freq 100), item B (freq 50)
After decay (0.8): item A (freq 80), item B (freq 40)  
New trending: item C gets 60 accesses
Result: C (60) > A (80*time_decay) > B (40*time_decay)
```

### Seasonal Pattern Adaptation
```
Summer content: high frequency, then unused for months
Winter access: frequencies have decayed significantly
New winter content: can compete effectively with old summer content
```

## When to Use Adaptive LFU

### ✅ Excellent for:
- **Dynamic web applications** with changing user interests
- **Content platforms** with trending and viral content
- **News websites** where article popularity shifts rapidly
- **Social media caches** following engagement trends
- **E-commerce** with seasonal and promotional patterns

### ✅ Better than LFU when:
- **Access patterns change** significantly over time
- **Historical frequency** shouldn't dominate indefinitely  
- **Trending content** needs to compete with established items
- **Workload has temporal phases** (daily, weekly, seasonal)

### ❌ Consider alternatives for:
- **Stable access patterns** (use regular LFU)
- **Memory-constrained** environments (use TinyLFU)
- **Simple temporal locality** (use LRU)
- **Very small caches** (overhead not justified)

## Comparison with Other Adaptive Algorithms

| Algorithm | Adaptation Method | Complexity | Memory | Responsiveness |
|-----------|------------------|------------|---------|----------------|
| **Adaptive LFU** | Time-based decay | Medium | Medium | Good |
| **TinyLFU** | Sketch aging | Medium | Low | Excellent |
| **Window-TinyLFU** | Window rotation | High | Medium | Excellent |
| **LRU** | Implicit (recency) | Simple | Low | Excellent |

Adaptive LFU provides a middle ground between the stability of traditional LFU and the responsiveness of recency-based algorithms, making it ideal for workloads with evolving but predictable frequency patterns.