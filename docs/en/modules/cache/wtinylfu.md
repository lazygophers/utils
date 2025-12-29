---
title: W-TinyLFU Cache
---

# W-TinyLFU (Windowed TinyLFU) Cache

W-TinyLFU is a windowed TinyLFU with sliding window.

## Overview

W-TinyLFU (Windowed TinyLFU) adds a sliding window mechanism to TinyLFU. It uses a time window to track access frequency, better handling time-based access patterns and periodic data access.

## Features

- **Hit Rate**: 90%
- **Memory Usage**: Medium
- **Concurrency**: High
- **Implementation Complexity**: Complex

## Use Cases

- Time-based access patterns
- Periodic data access
- Sliding window requirements
- Scenarios requiring time window statistics

## Quick Start

### Installation

```bash
go get github.com/lazygophers/utils/cache/wtinylfu
```

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/wtinylfu"
)

func main() {
    // Create cache with capacity of 1000, window size of 100
    cache := wtinylfu.New(1000, 100)

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

W-TinyLFU uses a sliding window:

```
Time Window
┌────────────────────────────┐
│  [access1, access2, ...]    │  ← Inside window
│  [access1, access2, ...]    │  ← Outside window
└────────────────────────────┘
    ↑ Window slides
```

Accesses within the window are used to calculate frequency, accesses outside the window gradually decay.

## API Reference

### Constructors

```go
// Create new W-TinyLFU cache
func New(capacity int, windowSize int) *WTinyLFU

// Create new W-TinyLFU cache with options
func NewWithOpts(opts Options) *WTinyLFU
```

### Main Methods

```go
// Set key-value pair
func (c *WTinyLFU) Set(key string, value interface{})

// Get value
func (c *WTinyLFU) Get(key string) (interface{}, bool)

// Delete key
func (c *WTinyLFU) Delete(key string)

// Clear cache
func (c *WTinyLFU) Clear()

// Get statistics
func (c *WTinyLFU) Stats() Stats
```

## Performance Characteristics

- **Time Complexity**:
  - Set: O(1) average
  - Get: O(1) average
  - Delete: O(1)
- **Space Complexity**: O(n), where n is cache capacity

## Best Practices

1. **Choose appropriate window size**: Adjust window size based on access patterns
2. **Periodic access**: W-TinyLFU is particularly suitable for periodic access patterns
3. **Time window statistics**: Suitable for scenarios requiring time window statistics

## Related Documentation

- [Cache Overview](./index.md)
- [TinyLFU Cache](./tinylfu.md)
- [Cache Selection Guide](./index.md#quick-selection-guide)
