# FBR (Frequency-Based Replacement) Cache

A thread-safe cache implementation that uses pure frequency-based replacement with LRU ordering within frequency groups. Items with the lowest access frequency are evicted first, and among items with the same frequency, the least recently used is chosen.

## Features

- **Pure frequency-based** eviction strategy
- **LRU ordering** within each frequency level
- **O(1)** operations for Get and Put
- **Thread-safe** with optimized locking
- **Dynamic frequency tracking** with efficient data structures
- **Generic support** for any comparable key and any value type

## How FBR Works

FBR maintains:

1. **Frequency groups**: Separate LRU lists for each frequency level
2. **Minimum frequency tracking**: Efficient identification of victims
3. **Dynamic frequency adjustment**: Items move between frequency groups
4. **LRU within frequency**: Recent items survive within same frequency

## Use Cases

- **Content delivery networks** where popularity is the primary factor
- **Database buffer pools** for frequently accessed pages
- **Application caches** where access frequency predicts future use
- **Media streaming** where popular content should be prioritized
- **Analytics systems** processing hot vs cold data

## API

```go
// Create FBR cache
cache := fbr.New[string, int](capacity)

// With eviction callback
cache := fbr.NewWithEvict[string, int](capacity, func(key string, value int) {
    fmt.Printf("Evicted: %s = %d (freq: %d)\n", key, value, freq)
})

// Basic operations
cache.Put("key", 42)           // Add/update and increment frequency
value, ok := cache.Get("key")  // Retrieve and increment frequency
cache.Remove("key")            // Remove specific item
cache.Contains("key")          // Check existence
value, ok := cache.Peek("key") // Get without affecting frequency

// Frequency inspection
freq := cache.Frequency("key") // Get current frequency
stats := cache.Stats()         // Frequency distribution statistics

// Management
cache.Clear()                  // Remove all items
cache.Resize(newSize)         // Change capacity
cache.Len()                   // Current item count
cache.Cap()                   // Maximum capacity
```

## Example

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/fbr"
)

func main() {
    // Create FBR cache with capacity 5
    cache := fbr.New[string, string](5)
    
    // Add items
    cache.Put("video1", "cat-video.mp4")
    cache.Put("video2", "dog-video.mp4") 
    cache.Put("video3", "bird-video.mp4")
    
    // Simulate different popularity levels
    // video1: very popular (freq 10)
    for i := 0; i < 9; i++ {
        cache.Get("video1")
    }
    
    // video2: moderately popular (freq 5)
    for i := 0; i < 4; i++ {
        cache.Get("video2")
    }
    
    // video3: unpopular (freq 1)
    // No additional accesses
    
    // Add more videos to trigger eviction
    cache.Put("video4", "fish-video.mp4")
    cache.Put("video5", "lion-video.mp4")
    cache.Put("video6", "tiger-video.mp4") // This will evict video3 (lowest freq)
    
    fmt.Printf("video1 freq: %d\n", cache.Frequency("video1")) // 10
    fmt.Printf("video2 freq: %d\n", cache.Frequency("video2")) // 5
    fmt.Printf("video3 present: %v\n", cache.Contains("video3")) // false - evicted
    
    // Check frequency distribution
    stats := cache.Stats()
    fmt.Printf("Frequency distribution: %+v\n", stats.FrequencyDistribution)
}
```

## Frequency Management

### Frequency Increment Strategy

Every successful Get() and Put() operation increments the item's frequency:

```go
cache.Put("key", "value")  // frequency = 1
cache.Get("key")          // frequency = 2  
cache.Get("key")          // frequency = 3
cache.Put("key", "new")   // frequency = 4 (update also increments)
```

### Eviction Order

Items are evicted in this priority order:
1. **Lowest frequency first**
2. **Within same frequency: LRU first**

```
Frequency 1: [C, B, A] (A most recent)
Frequency 2: [F, E, D] (D most recent) 
Frequency 3: [H, G]    (G most recent)

Eviction order: C → B → A → F → E → D → H → G
```

## Performance

- **Get**: O(1) - Hash lookup + list move + frequency update
- **Put**: O(1) - Hash insert + frequency management
- **Remove**: O(1) - Hash delete + list removal
- **Memory**: O(n + f) where n=items, f=unique frequencies
- **Frequency tracking**: O(1) with efficient min-frequency management

## Benchmarks

```
BenchmarkFBRPut-8       7000000    190 ns/op    40 B/op    1 allocs/op
BenchmarkFBRGet-8      12000000    140 ns/op     0 B/op    0 allocs/op
BenchmarkFBRMixed-8     9000000    165 ns/op    20 B/op    0 allocs/op
```

## Statistics

Monitor frequency distribution and cache behavior:

```go
stats := cache.Stats()
fmt.Printf("Cache statistics:\n")
fmt.Printf("  Size: %d/%d\n", stats.Size, stats.Capacity)
fmt.Printf("  Min frequency: %d\n", stats.MinFrequency)
fmt.Printf("  Max frequency: %d\n", stats.MaxFrequency)
fmt.Printf("  Frequency distribution:\n")

for freq, count := range stats.FrequencyDistribution {
    fmt.Printf("    Freq %d: %d items\n", freq, count)
}
```

## Configuration Best Practices

### Cache Size Guidelines

| Cache Size | Best For | Frequency Range |
|------------|----------|----------------|
| Small (<1K) | Hot data caching | High frequency differences |
| Medium (1K-100K) | Application caches | Mixed access patterns |
| Large (>100K) | CDN, database buffers | Wide frequency distribution |

### Workload Suitability

FBR works best when:
- **Access frequency predicts future access**
- **Popular items are much more popular than unpopular ones**
- **Frequency differences are significant**
- **Workload has stable popularity patterns**

## When to Use FBR

### ✅ Excellent for:
- **Content delivery** where popularity is key
- **Database systems** with hot pages
- **Web caches** serving popular content
- **Media streaming** with hit songs/videos
- **API caching** with popular endpoints

### ✅ Better than LRU when:
- **Frequency is more important than recency**
- **Popular items should never be evicted**
- **Access patterns have clear hot/cold distinction**
- **Long-term popularity matters**

### ❌ Consider alternatives for:
- **Temporal locality workloads** (use LRU)
- **Scan-heavy workloads** (use SLRU) 
- **Rapidly changing patterns** (use Adaptive LFU)
- **Memory-constrained environments** (use TinyLFU)

## Comparison with Other Algorithms

| Algorithm | Frequency Tracking | Memory Overhead | Adaptability | Complexity |
|-----------|-------------------|-----------------|--------------|------------|
| **FBR** | Perfect | Medium | Poor | Medium |
| **LFU** | Perfect | High | Poor | High |
| **TinyLFU** | Approximate | Low | Good | Medium |
| **LRU** | None | Low | Excellent | Simple |

FBR provides exact frequency tracking with better performance than traditional LFU implementations, making it ideal for workloads where frequency is the primary access predictor.