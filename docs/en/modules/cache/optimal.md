---
title: Optimal Cache
---

# Optimal Cache

Optimal is a theoretically optimal cache for performance.

## Overview

Optimal cache is a theoretically optimal cache strategy that assumes knowledge of future access patterns. It evicts the item that will be accessed furthest in the future, providing the theoretically highest hit rate. Mainly used for performance benchmarking and offline analysis.

## Features

- **Hit Rate**: 95% (theoretically optimal)
- **Memory Usage**: High
- **Concurrency**: Low
- **Implementation Complexity**: Complex

## Use Cases

- Maximum hit rate
- Predictable access patterns
- Offline analysis
- Performance benchmarking

## Quick Start

### Installation

```bash
go get github.com/lazygophers/utils/cache/optimal
```

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/optimal"
)

func main() {
    // Create cache with capacity of 1000, need to provide future access sequence
    futureAccess := []string{"key1", "key2", "key3"}
    cache := optimal.New(1000, futureAccess)

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

Optimal cache requires prior knowledge of access sequence:

```
Access sequence: [key1, key2, key1, key3, key2, ...]

Eviction decision:
- key1: next access = 3
- key2: next access = 5
- key3: next access = ∞ (never accessed again)

Evict key3, because its next access is furthest
```

## API Reference

### Constructors

```go
// Create new Optimal cache, need to provide future access sequence
func New(capacity int, futureAccess []string) *Optimal

// Create new Optimal cache with options
func NewWithOpts(opts Options) *Optimal
```

### Main Methods

```go
// Set key-value pair
func (c *Optimal) Set(key string, value interface{})

// Get value
func (c *Optimal) Get(key string) (interface{}, bool)

// Delete key
func (c *Optimal) Delete(key string)

// Clear cache
func (c *Optimal) Clear()

// Get statistics
func (c *Optimal) Stats() Stats
```

## Performance Characteristics

- **Time Complexity**:
  - Set: O(n)
  - Get: O(1)
  - Delete: O(n)
- **Space Complexity**: O(n), where n is cache capacity

## Best Practices

1. **Theoretical benchmark**: Optimal is mainly used for theoretical performance benchmarking
2. **Offline analysis**: Suitable for offline analysis and performance evaluation
3. **Not for production**: Due to requiring prior knowledge of access patterns, not suitable for production environments

## Limitations

- **Requires future access sequence**: Must know access patterns in advance
- **Not suitable for online scenarios**: Cannot be used for real-time online systems
- **Mainly for research**: Mainly used for research and benchmarking

## Related Documentation

- [Cache Overview](./index.md)
- [LRU Cache](./lru.md)
- [TinyLFU Cache](./tinylfu.md)
