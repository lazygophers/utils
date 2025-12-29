---
title: SLRU Cache
---

# SLRU (Segmented LRU) Cache

SLRU is a segmented LRU cache with multiple segments, suitable for high concurrency environments.

## Overview

SLRU (Segmented LRU) divides the cache into multiple segments, where each segment is an independent LRU cache. This design can reduce lock contention and improve concurrency performance.

## Features

- **Hit Rate**: 90%
- **Memory Usage**: High
- **Concurrency**: High
- **Implementation Complexity**: Medium

## Use Cases

- Reduce lock contention
- High concurrency environments
- Large cache sizes
- Scenarios requiring high concurrency performance

## Quick Start

### Installation

```bash
go get github.com/lazygophers/utils/cache/slru
```

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/slru"
)

func main() {
    // Create cache with capacity of 1000, divided into 4 segments
    cache := slru.New(1000, 4)

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

SLRU divides the cache into multiple segments:

```
┌─────────────┐
│  Segment 1 │
├─────────────┤
│  Segment 2 │
├─────────────┤
│  Segment 3 │
├─────────────┤
│  Segment 4 │
└─────────────┘
```

Each segment has its own lock, allowing multiple goroutines to access different segments simultaneously, reducing lock contention.

## API Reference

### Constructors

```go
// Create new SLRU cache
func New(capacity int, segments int) *SLRU

// Create new SLRU cache with options
func NewWithOpts(opts Options) *SLRU
```

### Main Methods

```go
// Set key-value pair
func (c *SLRU) Set(key string, value interface{})

// Set key-value pair with TTL
func (c *SLRU) SetWithTTL(key string, value interface{}, ttl time.Duration)

// Get value
func (c *SLRU) Get(key string) (interface{}, bool)

// Delete key
func (c *SLRU) Delete(key string)

// Clear cache
func (c *SLRU) Clear()

// Get statistics
func (c *SLRU) Stats() Stats
```

## Performance Characteristics

- **Time Complexity**:
  - Set: O(1)
  - Get: O(1)
  - Delete: O(1)
- **Space Complexity**: O(n), where n is cache capacity
- **Concurrency Performance**: 2-3x higher than standard LRU

## Best Practices

1. **Choose appropriate segment count**: Usually 4-8 segments work best
2. **First choice for high concurrency**: SLRU provides best performance under high concurrency
3. **Monitor segment distribution**: Periodically check usage of each segment to ensure load balancing

## Related Documentation

- [Cache Overview](./index.md)
- [LRU Cache](./lru.md)
- [Cache Selection Guide](./index.md#quick-selection-guide)
