---
title: MRU Cache
---

# MRU (Most Recently Used) Cache

MRU evicts the most recently used items, suitable for scenarios with strong temporal locality.

## Overview

MRU (Most Recently Used) is based on the principle of temporal locality, assuming that recently accessed data will not be accessed again. When the cache is full, it evicts the most recently used data.

## Features

- **Hit Rate**: 80%
- **Memory Usage**: Low
- **Concurrency**: Medium
- **Implementation Complexity**: Simple

## Use Cases

- Temporal locality
- Sequential access patterns
- Cache warmup
- Scan-style access

## Quick Start

### Installation

```bash
go get github.com/lazygophers/utils/cache/mru
```

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/mru"
)

func main() {
    // Create cache with capacity of 1000
    cache := mru.New(1000)

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

## API Reference

### Constructors

```go
// Create new MRU cache
func New(capacity int) *MRU

// Create new MRU cache with options
func NewWithOpts(opts Options) *MRU
```

### Main Methods

```go
// Set key-value pair
func (c *MRU) Set(key string, value interface{})

// Get value
func (c *MRU) Get(key string) (interface{}, bool)

// Delete key
func (c *MRU) Delete(key string)

// Clear cache
func (c *MRU) Clear()

// Get statistics
func (c *MRU) Stats() Stats
```

## Performance Characteristics

- **Time Complexity**:
  - Set: O(1)
  - Get: O(1)
  - Delete: O(1)
- **Space Complexity**: O(n), where n is cache capacity

## Best Practices

1. **Sequential access scenarios**: MRU is suitable for scan-style access
2. **Cache warmup**: During warmup phase, MRU can quickly evict warmup data
3. **Avoid frequent access**: Not suitable for scenarios where the same data is accessed frequently

## Related Documentation

- [Cache Overview](./index.md)
- [LRU Cache](./lru.md)
- [Cache Selection Guide](./index.md#quick-selection-guide)
