---
title: ALFU Cache
---

# ALFU (Adaptive LFU) Cache

ALFU is an adaptive LFU cache that adjusts based on access patterns.

## Overview

ALFU (Adaptive LFU) is an adaptive cache strategy that dynamically adjusts eviction strategies based on access patterns. It builds on LFU (Least Frequently Used) by adding adaptive mechanisms to adapt to changing access patterns.

## Features

- **Hit Rate**: 82%
- **Memory Usage**: Medium
- **Concurrency**: Medium
- **Implementation Complexity**: Complex

## Use Cases

- Unknown access patterns
- Adaptive requirements
- Learning environments
- Scenarios with changing access patterns

## Quick Start

### Installation

```bash
go get github.com/lazygophers/utils/cache/alfu
```

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/alfu"
)

func main() {
    // Create cache with capacity of 1000
    cache := alfu.New(1000)

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

ALFU uses adaptive mechanisms:

1. **Frequency Tracking**: Tracks access frequency for each key
2. **Dynamic Adjustment**: Dynamically adjusts eviction strategy based on access patterns
3. **Pattern Learning**: Can learn changes in access patterns

## API Reference

### Constructors

```go
// Create new ALFU cache
func New(capacity int) *ALFU

// Create new ALFU cache with options
func NewWithOpts(opts Options) *ALFU
```

### Main Methods

```go
// Set key-value pair
func (c *ALFU) Set(key string, value interface{})

// Get value
func (c *ALFU) Get(key string) (interface{}, bool)

// Delete key
func (c *ALFU) Delete(key string)

// Clear cache
func (c *ALFU) Clear()

// Get statistics
func (c *ALFU) Stats() Stats
```

## Performance Characteristics

- **Time Complexity**:
  - Set: O(log n)
  - Get: O(log n)
  - Delete: O(log n)
- **Space Complexity**: O(n), where n is cache capacity

## Best Practices

1. **Unknown access patterns**: ALFU is suitable for scenarios where access patterns are unknown or changing
2. **Adaptive requirements**: Use when cache needs to automatically adapt to access patterns
3. **Monitor learning process**: Observe the cache's learning process to confirm adaptive effectiveness

## Related Documentation

- [Cache Overview](./index.md)
- [LFU Cache](./lfu.md)
- [ARC Cache](./arc.md)
