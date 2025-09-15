# Cache Algorithms Collection

A comprehensive collection of high-performance, thread-safe cache implementations in Go using generics.

## Available Algorithms

| Algorithm | Path | Coverage | Description |
|-----------|------|----------|-------------|
| **LRU** | `lru/` | 99.3% | Least Recently Used - Classic cache with recency-based eviction |
| **LFU** | `lfu/` | 98.5% | Least Frequently Used - Frequency-based eviction with aging |
| **MRU** | `mru/` | 97.7% | Most Recently Used - Opposite of LRU, evicts most recent items |
| **SLRU** | `slru/` | 97.8% | Segmented LRU - Two-tier LRU with probationary and protected segments |
| **TinyLFU** | `tinylfu/` | 97.5% | Memory-efficient LFU using count-min sketch |
| **FBR** | `fbr/` | 99.2% | Frequency-Based Replacement with LRU ordering within frequencies |
| **LRU-K** | `lruk/` | 97.6% | LRU-K algorithm tracking K most recent access times |
| **Adaptive LFU** | `alfu/` | 96.0% | LFU with time-based frequency decay and adaptation |
| **Window-TinyLFU** | `wtinylfu/` | 82.7% | Combines LRU window with TinyLFU main cache |
| **Optimal** | `optimal/` | 94.8% | Belady's optimal algorithm for analysis and benchmarking |

## Features

- **Type Safety**: Full Go generics support for compile-time type safety
- **Thread Safety**: All implementations are goroutine-safe with optimized locking
- **High Performance**: Optimized data structures for microsecond operations
- **Consistent API**: Uniform interface across all cache implementations
- **Comprehensive Testing**: >95% average test coverage with extensive benchmarks
- **Memory Efficient**: Careful memory management and resource cleanup

## Common Interface

All cache implementations follow this interface:

```go
type Cache[K comparable, V any] interface {
    Get(key K) (value V, ok bool)
    Put(key K, value V) (evicted bool)
    Remove(key K) (value V, ok bool)
    Contains(key K) bool
    Peek(key K) (value V, ok bool)
    Clear()
    Len() int
    Cap() int
    Keys() []K
    Values() []V
    Items() map[K]V
    Resize(capacity int)
}
```

## Quick Start

```go
import "github.com/lazygophers/utils/cache/lru"

// Create an LRU cache with capacity 1000
cache := lru.New[string, int](1000)

// Basic operations
cache.Put("key1", 42)
value, ok := cache.Get("key1")
cache.Remove("key1")

// With eviction callback
cache := lru.NewWithEvict[string, int](1000, func(key string, value int) {
    fmt.Printf("Evicted: %s = %d\n", key, value)
})
```

## Algorithm Selection Guide

### For General Purpose
- **LRU**: Simple, predictable, good for most workloads
- **Window-TinyLFU**: Best overall hit rates for mixed workloads

### For Frequency-Sensitive Workloads
- **LFU**: When frequency matters more than recency
- **TinyLFU**: Memory-efficient frequency tracking
- **Adaptive LFU**: Dynamic workloads with changing patterns

### For Specialized Use Cases
- **SLRU**: Scan-resistant workloads
- **LRU-K**: Better than LRU for database buffer pools
- **MRU**: Scenarios where recent items are less likely to be reused
- **Optimal**: Analysis and benchmarking reference

### For Analysis
- **Optimal**: Theoretical maximum performance for comparison

## Performance Characteristics

| Algorithm | Time Complexity | Space Overhead | Best Use Case |
|-----------|----------------|----------------|---------------|
| LRU | O(1) | Low | General purpose |
| LFU | O(log n) | Medium | Frequency-sensitive |
| TinyLFU | O(1) | Low | Large caches |
| SLRU | O(1) | Low | Scan-resistant |
| LRU-K | O(1) | Medium | Database buffers |
| Window-TinyLFU | O(1) | Medium | Mixed workloads |

## Benchmarks

Run benchmarks for all implementations:

```bash
go test -bench=. -benchmem ./cache/...
```

## Testing

Run comprehensive tests with coverage:

```bash
go test -cover ./cache/...
```

Each implementation includes:
- Unit tests for all operations
- Concurrency safety tests
- Edge case handling
- Performance benchmarks
- Memory usage analysis