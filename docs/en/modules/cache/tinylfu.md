---
title: TinyLFU Cache
---

# TinyLFU Cache

TinyLFU is a high-performance cache combining LRU and LFU.

## Overview

TinyLFU is a hybrid cache strategy that combines the advantages of LRU (Least Recently Used) and LFU (Least Frequently Used). It uses two small caches: one tracking recently accessed items (LRU), another tracking access frequency (LFU), and considers both when evicting.

## Features

- **Hit Rate**: 92% (Highest)
- **Memory Usage**: Medium
- **Concurrency**: High
- **Implementation Complexity**: Complex

## Use Cases

- High performance requirements
- Mixed access patterns
- Large datasets
- Scenarios requiring best hit rate

## Quick Start

### Installation

```bash
go get github.com/lazygophers/utils/cache/tinylfu
```

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/tinylfu"
)

func main() {
    // Create cache with capacity of 1000
    cache := tinylfu.New(1000)

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

## How It Works

TinyLFU uses two small caches:

1. **Window Cache (LRU)**: Tracks recently accessed items
2. **Main Cache (LFU)**: Tracks access frequency

When eviction is needed:
- Prioritize evicting items from Window Cache
- If Window Cache is empty, evict the lowest frequency item from Main Cache

## API Reference

### Constructors

```go
// Create new TinyLFU cache
func New(capacity int) *TinyLFU

// Create new TinyLFU cache with options
func NewWithOpts(opts Options) *TinyLFU
```

### Main Methods

```go
// Set key-value pair
func (c *TinyLFU) Set(key string, value interface{})

// Set key-value pair with TTL
func (c *TinyLFU) SetWithTTL(key string, value interface{}, ttl time.Duration)

// Get value
func (c *TinyLFU) Get(key string) (interface{}, bool)

// Delete key
func (c *TinyLFU) Delete(key string)

// Clear cache
func (c *TinyLFU) Clear()

// Get statistics
func (c *TinyLFU) Stats() Stats
```

## Performance Characteristics

- **Time Complexity**:
  - Set: O(1) average
  - Get: O(1) average
  - Delete: O(1)
- **Space Complexity**: O(n), where n is cache capacity

## Best Practices

1. **First choice for high performance**: TinyLFU provides the highest hit rate
2. **Sufficient memory**: Requires additional memory to maintain two small caches
3. **Mixed access patterns**: Best for scenarios with both temporal and frequency locality

## Related Documentation

- [Cache Overview](./index.md)
- [LRU Cache](./lru.md)
- [LFU Cache](./lfu.md)
