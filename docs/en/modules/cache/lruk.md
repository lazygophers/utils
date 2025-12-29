---
title: LRU-K Cache
---

# LRU-K (Least Recently Used with K) Cache

LRU-K is an LRU cache that tracks access frequency.

## Overview

LRU-K is an improved version of LRU that tracks the access frequency for each key (K times). When the cache is full, it prioritizes evicting keys with lower access frequencies, rather than simply evicting the least recently used key.

## Features

- **Hit Rate**: 88%
- **Memory Usage**: Medium
- **Concurrency**: Medium
- **Implementation Complexity**: Medium

## Use Cases

- Balance recency and frequency
- Mixed access patterns
- Scenarios needing to consider access frequency

## Quick Start

### Installation

```bash
go get github.com/lazygophers/utils/cache/lruk
```

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/lruk"
)

func main() {
    // Create cache with capacity of 1000, K=2
    cache := lruk.New(1000, 2)

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

## How It Works

LRU-K maintains access history for each key (most recent K accesses):

```
Key1: [access1, access2, ..., accessK]
Key2: [access1, access2, ..., accessK]
```

When evicting, considers both:
1. Recent access time
2. Access frequency (based on K recent accesses)

## API Reference

### Constructors

```go
// Create new LRU-K cache
func New(capacity int, k int) *LRUK

// Create new LRU-K cache with options
func NewWithOpts(opts Options) *LRUK
```

### Main Methods

```go
// Set key-value pair
func (c *LRUK) Set(key string, value interface{})

// Get value
func (c *LRUK) Get(key string) (interface{}, bool)

// Delete key
func (c *LRUK) Delete(key string)

// Clear cache
func (c *LRUK) Clear()

// Get statistics
func (c *LRUK) Stats() Stats
```

## Performance Characteristics

- **Time Complexity**:
  - Set: O(1)
  - Get: O(1)
  - Delete: O(1)
- **Space Complexity**: O(n * k), where n is cache capacity, k is number of tracked accesses

## Best Practices

1. **Choose appropriate K value**: Usually K=2 or K=3 works well
2. **Balance recency and frequency**: LRU-K provides balance between both
3. **Mixed access patterns**: Suitable for scenarios with both temporal and frequency locality

## Related Documentation

- [Cache Overview](./index.md)
- [LRU Cache](./lru.md)
- [TinyLFU Cache](./tinylfu.md)
