---
title: LFU Cache
---

# LFU (Least Frequently Used) Cache

LFU evicts the least frequently used items, suitable for infrequently accessed data scenarios.

## Overview

LFU (Least Frequently Used) is based on the principle of frequency locality, assuming that frequently accessed data should be retained. When the cache is full, it evicts the least frequently accessed data.

## Features

- **Hit Rate**: 75%
- **Memory Usage**: Low
- **Concurrency**: Medium
- **Implementation Complexity**: Medium

## Use Cases

- Infrequently accessed data
- Large datasets
- Memory constrained environments
- Hot/cold data separation scenarios

## Quick Start

### Installation

```bash
go get github.com/lazygophers/utils/cache/lfu
```

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/lfu"
)

func main() {
    // Create cache with capacity of 1000
    cache := lfu.New(1000)

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
// Create new LFU cache
func New(capacity int) *LFU

// Create new LFU cache with options
func NewWithOpts(opts Options) *LFU
```

### Main Methods

```go
// Set key-value pair
func (c *LFU) Set(key string, value interface{})

// Set key-value pair with TTL
func (c *LFU) SetWithTTL(key string, value interface{}, ttl time.Duration)

// Get value
func (c *LFU) Get(key string) (interface{}, bool)

// Delete key
func (c *LFU) Delete(key string)

// Clear cache
func (c *LFU) Clear()

// Get statistics
func (c *LFU) Stats() Stats
```

## Performance Characteristics

- **Time Complexity**:
  - Set: O(log n)
  - Get: O(log n)
  - Delete: O(log n)
- **Space Complexity**: O(n), where n is cache capacity

## Best Practices

1. **Suitable for cold data**: LFU is suitable for scenarios with large access frequency differences
2. **Avoid cache pollution**: For burst access, consider using other strategies
3. **Monitor frequency distribution**: Periodically check access frequency distribution to confirm LFU is appropriate

## Related Documentation

- [Cache Overview](./index.md)
- [LRU Cache](./lru.md)
- [ALFU Cache](./alfu.md)
