---
title: Cache Implementations
---

# Cache Implementations

The cache module provides multiple cache implementations with different eviction strategies for various use cases.

## How to Choose the Right Cache

When selecting a cache implementation, consider the following factors:

### 1. Access Pattern
- **Temporal Locality**: Recently accessed data likely to be accessed again → Choose MRU
- **Frequency Locality**: Frequently accessed data should be retained → Choose LFU
- **Mixed Pattern**: Both temporal and frequency locality → Choose TinyLFU or LRU-K

### 2. Concurrency Requirements
- **Low Concurrency**: Standard LRU is sufficient
- **High Concurrency**: Choose SLRU (Segmented LRU) to reduce lock contention

### 3. Performance Requirements
- **General Performance**: LRU (85% hit rate)
- **High Performance**: TinyLFU (92% hit rate) or Optimal (95% hit rate)
- **Memory Sensitive**: LFU (lowest memory usage)

### 4. Adaptive Needs
- **Known Access Pattern**: Choose corresponding strategy based on pattern
- **Unknown Access Pattern**: Choose ALFU (Adaptive LFU) or ARC (Adaptive Replacement)

## Cache Implementation Comparison

| Cache Type | Hit Rate | Memory | Concurrency | Best For | Recommendation |
|-----------|---------|---------|------------|---------|
| **[LRU](./lru.md)** | 85% | Low | Medium | General cache, frequently accessed data | ⭐⭐⭐⭐ |
| **[LFU](./lfu.md)** | 75% | Low | Medium | Infrequently accessed data, memory constrained | ⭐⭐⭐ |
| **[LRU-K](./lruk.md)** | 88% | Medium | Medium | Balance recency and frequency, mixed pattern | ⭐⭐⭐⭐ |
| **[MRU](./mru.md)** | 80% | Low | Medium | Temporal locality, sequential access | ⭐⭐⭐ |
| **[TinyLFU](./tinylfu.md)** | 92% | Medium | High | High performance requirements, mixed pattern | ⭐⭐⭐⭐⭐ |
| **[W-TinyLFU](./wtinylfu.md)** | 90% | Medium | High | Time-based access pattern, periodic data | ⭐⭐⭐ |
| **[ALFU](./alfu.md)** | 82% | Medium | Medium | Unknown access pattern, adaptive needs | ⭐⭐⭐ |
| **[ARC](./arc.md)** | 86% | Medium | High | Mixed access pattern, adaptive | ⭐⭐⭐⭐ |
| **[FBR](./fbr.md)** | 78% | Medium | Medium | Frequency-based access, hot data retention | ⭐⭐⭐ |
| **[SLRU](./slru.md)** | 90% | High | High | High concurrency, large cache size | ⭐⭐⭐⭐ |
| **[Optimal](./optimal.md)** | 95% | High | Low | Predictable access pattern, offline analysis | ⭐⭐⭐ |

## Quick Selection Guide

### Choose by Scenario

```go
// Scenario 1: General web application cache
import "github.com/lazygophers/utils/cache/lru"
cache := lru.New(1000)  // LRU - Most general

// Scenario 2: High concurrency API cache
import "github.com/lazygophers/utils/cache/slru"
cache := slru.New(1000)  // SLRU - Reduce lock contention

// Scenario 3: High performance requirements
import "github.com/lazygophers/utils/cache/tinylfu"
cache := tinylfu.New(1000)  // TinyLFU - Highest hit rate

// Scenario 4: Unknown access pattern
import "github.com/lazygophers/utils/cache/alfu"
cache := alfu.New(1000)  // ALFU - Adaptive

// Scenario 5: Sequential access data
import "github.com/lazygophers/utils/cache/mru"
cache := mru.New(1000)  // MRU - Temporal locality
```

### Decision Tree

```
Known access pattern?
├─ Yes
│  ├─ High concurrency? → SLRU
│  ├─ Sequential access? → MRU
│  ├─ High performance? → TinyLFU
│  └─ General scenario? → LRU
└─ No
   ├─ Need adaptation? → ALFU or ARC
   └─ Predictable pattern? → Optimal
```

## Basic Usage Example

### Creating Cache

```go
import "github.com/lazygophers/utils/cache/lru"

// Create cache with capacity of 1000
cache := lru.New(1000)
```

### Basic Operations

```go
// Set values
cache.Set("key1", "value1")
cache.Set("key2", "value2")

// Get values
if value, ok := cache.Get("key1"); ok {
    fmt.Println("Found:", value)
}

// Delete values
cache.Delete("key1")

// Clear cache
cache.Clear()
```

### Cache Statistics

```go
// Get cache statistics
stats := cache.Stats()

fmt.Printf("Size: %d\n", stats.Size)
fmt.Printf("Hits: %d\n", stats.Hits)
fmt.Printf("Misses: %d\n", stats.Misses)
fmt.Printf("Hit Rate: %.2f%%\n", stats.HitRate())
```

## Best Practices

### 1. Cache Size Selection

```go
// Adjust cache size based on available memory
func calculateCacheSize() int {
    // Good practice: Based on memory constraints
    availableMem := getAvailableMemory()
    return availableMem / 1024  // ~1KB per entry
}

cache := lru.New(calculateCacheSize())
```

### 2. Monitor Hit Rate

```go
// Periodically check hit rate
func monitorCache(cache Cache) {
    stats := cache.Stats()
    if stats.HitRate() < 0.5 {
        // Hit rate too low, consider:
        // 1. Increase cache size
        // 2. Change cache strategy
        // 3. Check access pattern
    }
}
```

### 3. Choose Appropriate Cache Type

```go
// Select based on actual scenario
func createCache() Cache {
    if isSequentialAccess() {
        return mru.New(1000)  // MRU for sequential access
    } else if isHighConcurrency() {
        return slru.New(1000)  // SLRU for high concurrency
    } else if isHighPerformance() {
        return tinylfu.New(1000)  // TinyLFU for high performance
    } else {
        return lru.New(1000)  // LRU for general
    }
}
```

## Related Documentation

- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
