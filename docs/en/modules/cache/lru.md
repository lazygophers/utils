---
title: LRU Cache
---

# LRU (Least Recently Used) Cache

LRU is the most commonly used cache eviction strategy, evicting the least recently used items when capacity is reached.

## Overview

LRU (Least Recently Used) is based on the principle of temporal locality, assuming that recently accessed data is likely to be accessed again. When the cache is full, it evicts the least recently used data.

## Features

- **Hit Rate**: 85%
- **Memory Usage**: Low
- **Concurrency**: Medium
- **Implementation Complexity**: Simple

## Use Cases

- General purpose caching
- Frequently accessed data
- Predictable access patterns
- Web application caching
- Database query caching

## Quick Start

### Installation

```bash
go get github.com/lazygophers/utils/cache/lru
```

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/lru"
)

func main() {
    // Create cache with capacity of 1000
    cache := lru.New(1000)

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
}
```

### Advanced Usage

```go
// Set with TTL
cache.SetWithTTL("key", "value", time.Minute*5)

// Batch operations
cache.SetMany(map[string]interface{}{
    "key1": "value1",
    "key2": "value2",
})

values := cache.GetMany([]string{"key1", "key2"})

// Get cache statistics
stats := cache.Stats()
fmt.Printf("Size: %d\n", stats.Size)
fmt.Printf("Hits: %d\n", stats.Hits)
fmt.Printf("Misses: %d\n", stats.Misses)
fmt.Printf("Hit Rate: %.2f%%\n", stats.HitRate())
```

## API Reference

### Constructors

```go
// Create new LRU cache
func New(capacity int) *LRU

// Create new LRU cache with options
func NewWithOpts(opts Options) *LRU
```

### Main Methods

```go
// Set key-value pair
func (c *LRU) Set(key string, value interface{})

// Set key-value pair with TTL
func (c *LRU) SetWithTTL(key string, value interface{}, ttl time.Duration)

// Get value
func (c *LRU) Get(key string) (interface{}, bool)

// Delete key
func (c *LRU) Delete(key string)

// Clear cache
func (c *LRU) Clear()

// Get statistics
func (c *LRU) Stats() Stats
```

## Performance Characteristics

- **Time Complexity**:
  - Set: O(1)
  - Get: O(1)
  - Delete: O(1)
- **Space Complexity**: O(n), where n is cache capacity

## Best Practices

1. **Choose appropriate cache size**: Adjust based on available memory and access patterns
2. **Monitor hit rate**: Periodically check hit rate, adjust if below 50%
3. **Use TTL**: For time-sensitive data, use SetWithTTL
4. **Batch operations**: For large data, use SetMany and GetMany for better performance

## Related Documentation

- [Cache Overview](./index.md)
- [LFU Cache](./lfu.md)
- [TinyLFU Cache](./tinylfu.md)
