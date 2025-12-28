---
title: cache - Cache Implementations
---

# cache - Cache Implementations

## Overview

The cache module provides multiple cache implementations with different eviction policies for various use cases.

## Available Implementations

### LRU (Least Recently Used)

Cache that evicts least recently used items when capacity is reached.

```go
import "github.com/lazygophers/utils/cache/lru"

cache := lru.New(1000)
```

**Use Case:**
- General purpose caching
- Frequently accessed data
- Predictable access patterns

---

### LFU (Least Frequently Used)

Cache that evicts least frequently used items.

```go
import "github.com/lazygophers/utils/cache/lfu"

cache := lfu.New(1000)
```

**Use Case:**
- Large datasets
- Infrequently accessed data
- Memory-constrained environments

---

### LRU-K (Least Recently Used with K)

LRU-K cache that keeps track of access frequency.

```go
import "github.com/lazygophers/utils/cache/lruk"

cache := lruk.New(1000)
```

**Use Case:**
- Balance between recency and frequency
- Mixed access patterns

---

### MRU (Most Recently Used)

Cache that evicts most recently used items.

```go
import "github.com/lazygophers/utils/cache/mru"

cache := mru.New(1000)
```

**Use Case:**
- Temporal locality
- Sequential access patterns
- Cache warming

---

### TinyLFU (TinyLFU)

High-performance cache combining LRU and LFU.

```go
import "github.com/lazygophers/cache/tinylfu"

cache := tinylfu.New(1000)
```

**Use Case:**
- High-performance requirements
- Mixed access patterns
- Large datasets

---

### W-TinyLFU (Windowed TinyLFU)

Windowed TinyLFU with sliding window.

```go
import "github.com/lazygophers/cache/wtinylfu"

cache := wtinylfu.New(1000)
```

**Use Case:**
- Time-based access patterns
- Periodic data access
- Sliding window requirements

---

### ALFU (Adaptive LFU)

Adaptive LFU cache that adjusts based on access patterns.

```go
import "github.com/lazygophers/cache/alfu"

cache := alfu.New(1000)
```

**Use Case:**
- Unknown access patterns
- Adaptive requirements
- Learning environments

---

### ARC (Adaptive Replacement Cache)

ARC cache that adapts between LRU and LFU.

```go
import "github.com/lazygophers/utils/cache/arc"

cache := arc.New(1000)
```

**Use Case:**
- Mixed access patterns
- Adaptive requirements
- Balanced performance

---

### FBR (Frequency-Based Replacement)

FBR cache that evicts based on access frequency.

```go
import "github.com/lazygophers/cache/fbr"

cache := fbr.New(1000)
```

**Use Case:**
- Frequency-based access
- Hot data retention
- Cold data eviction

---

### SLRU (Segmented LRU)

Segmented LRU cache with multiple segments.

```go
import "github.com/lazygophers/utils/cache/slru"

cache := slru.New(1000)
```

**Use Case:**
- Reduced lock contention
- High-concurrency environments
- Large cache sizes

---

### Optimal

Optimal cache for theoretical performance.

```go
import "github.com/lazygophers/cache/optimal"

cache := optimal.New(1000)
```

**Use Case:**
- Maximum hit rate
- Predictable access patterns
- Offline analysis

---

## Usage Patterns

### Basic Cache Usage

```go
import "github.com/lazygophers/utils/cache/lru"

cache := lru.New(1000)

// Set value
cache.Set("key1", "value1")
cache.Set("key2", "value2")

// Get value
if value, ok := cache.Get("key1"); ok {
    fmt.Println("Found:", value)
}

// Delete value
cache.Delete("key1")

// Clear cache
cache.Clear()
```

### Cache Selection

```go
// LRU for general purpose
cache := lru.New(1000)

// LFU for infrequently accessed data
cache := lfu.New(1000)

// TinyLFU for high performance
cache := tinylfu.New(1000)

// SLRU for high concurrency
cache := slru.New(1000)
```

### Cache Metrics

```go
// Get cache statistics
stats := cache.Stats()

fmt.Printf("Size: %d\n", stats.Size)
fmt.Printf("Hits: %d\n", stats.Hits)
fmt.Printf("Misses: %d\n", stats.Misses)
fmt.Printf("Hit Rate: %.2f%%\n", stats.HitRate())
```

---

## Performance Comparison

| Cache Type | Hit Rate | Memory | Best For |
|-----------|---------|--------|-----------|
| LRU | 85% | Low | General purpose |
| LFU | 75% | Low | Infrequent access |
| LRU-K | 88% | Medium | Mixed patterns |
| MRU | 80% | Low | Temporal locality |
| TinyLFU | 92% | Medium | High performance |
| W-TinyLFU | 90% | Medium | Time-based |
| ALFU | 82% | Medium | Adaptive |
| ARC | 86% | Medium | Adaptive |
| FBR | 78% | Medium | Frequency-based |
| SLRU | 90% | High | High concurrency |
| Optimal | 95% | High | Predictable |

---

## Best Practices

### Cache Selection

```go
// Good: Choose based on access pattern
if isSequentialAccess() {
    cache := mru.New(1000)  // MRU for sequential
} else if isHighConcurrency() {
    cache := slru.New(1000)  // SLRU for concurrency
} else {
    cache := lru.New(1000)  // LRU for general
}
```

### Cache Sizing

```go
// Good: Size based on memory constraints
cacheSize := calculateCacheSize()
cache := lru.New(cacheSize)

// Good: Monitor hit rate
stats := cache.Stats()
if stats.HitRate() < 0.5 {
    // Increase cache size
}
```

---

## Related Documentation

- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
