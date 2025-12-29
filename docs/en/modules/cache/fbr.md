---
title: FBR Cache
---

# FBR (Frequency-Based Replacement) Cache

FBR is a frequency-based replacement cache.

## Overview

FBR (Frequency-Based Replacement) is a cache eviction strategy based on access frequency. It tracks the access frequency for each key, and when the cache is full, it evicts the key with the lowest access frequency.

## Features

- **Hit Rate**: 78%
- **Memory Usage**: Medium
- **Concurrency**: Medium
- **Implementation Complexity**: Medium

## Use Cases

- Frequency-based access
- Hot data retention
- Cold data eviction
- Scenarios needing clear distinction between hot and cold data

## Quick Start

### Installation

```bash
go get github.com/lazygophers/utils/cache/fbr
```

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/fbr"
)

func main() {
    // Create cache with capacity of 1000
    cache := fbr.New(1000)

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

FBR maintains access frequency for each key:

```
Key1: access count = 5  (hot data)
Key2: access count = 1  (cold data)
Key3: access count = 0  (never accessed)
```

When evicting, prioritizes evicting keys with the lowest access frequency.

## API Reference

### Constructors

```go
// Create new FBR cache
func New(capacity int) *FBR

// Create new FBR cache with options
func NewWithOpts(opts Options) *FBR
```

### Main Methods

```go
// Set key-value pair
func (c *FBR) Set(key string, value interface{})

// Get value
func (c *FBR) Get(key string) (interface{}, bool)

// Delete key
func (c *FBR) Delete(key string)

// Clear cache
func (c *FBR) Clear()

// Get statistics
func (c *FBR) Stats() Stats
```

## Performance Characteristics

- **Time Complexity**:
  - Set: O(log n)
  - Get: O(log n)
  - Delete: O(log n)
- **Space Complexity**: O(n), where n is cache capacity

## Best Practices

1. **Hot data retention**: FBR can effectively retain hot data
2. **Cold data eviction**: Suitable for scenarios needing to quickly evict cold data
3. **Large frequency differences**: Suitable for scenarios with large access frequency differences

## Related Documentation

- [Cache Overview](./index.md)
- [LFU Cache](./lfu.md)
- [ALFU Cache](./alfu.md)
