---
title: ARC Cache
---

# ARC (Adaptive Replacement Cache) Cache

ARC is an adaptive cache that adjusts between LRU and LFU.

## Overview

ARC (Adaptive Replacement Cache) is an adaptive cache strategy that dynamically adjusts between LRU (Least Recently Used) and LFU (Least Frequently Used). It maintains two lists: T1 (recently used) and T2 (frequently used), and dynamically adjusts the size of both lists based on access patterns.

## Features

- **Hit Rate**: 86%
- **Memory Usage**: Medium
- **Concurrency**: High
- **Implementation Complexity**: Medium

## Use Cases

- Mixed access patterns
- Adaptive requirements
- Balanced performance
- Scenarios needing balance between LRU and LFU

## Quick Start

### Installation

```bash
go get github.com/lazygophers/utils/cache/arc
```

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/arc"
)

func main() {
    // Create cache with capacity of 1000
    cache := arc.New(1000)

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

ARC maintains four lists:

1. **T1**: Items accessed once recently
2. **T2**: Items accessed twice or more recently
3. **B1**: Items recently evicted from T1
4. **B2**: Items recently evicted from T2

Dynamically adjusts the size of T1 and T2 based on access patterns.

## API Reference

### Constructors

```go
// Create new ARC cache
func New(capacity int) *ARC

// Create new ARC cache with options
func NewWithOpts(opts Options) *ARC
```

### Main Methods

```go
// Set key-value pair
func (c *ARC) Set(key string, value interface{})

// Get value
func (c *ARC) Get(key string) (interface{}, bool)

// Delete key
func (c *ARC) Delete(key string)

// Clear cache
func (c *ARC) Clear()

// Get statistics
func (c *ARC) Stats() Stats
```

## Performance Characteristics

- **Time Complexity**:
  - Set: O(1)
  - Get: O(1)
  - Delete: O(1)
- **Space Complexity**: O(n), where n is cache capacity

## Best Practices

1. **Mixed access patterns**: ARC is suitable for scenarios with both temporal and frequency locality
2. **Adaptive requirements**: Use when cache needs to automatically adjust between LRU and LFU
3. **Balanced performance**: ARC provides stable performance across various access patterns

## Related Documentation

- [Cache Overview](./index.md)
- [LRU Cache](./lru.md)
- [LFU Cache](./lfu.md)
