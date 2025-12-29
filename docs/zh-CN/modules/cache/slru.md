---
title: SLRU 缓存
---

# SLRU (Segmented LRU) 缓存

SLRU 具有多个分段的分段 LRU 缓存，适合高并发环境。

## 概述

SLRU（Segmented LRU）将缓存分为多个段（segment），每个段是一个独立的 LRU 缓存。这种设计可以减少锁竞争，提高并发性能。

## 特性

- **命中率**: 90%
- **内存占用**: 高
- **并发性能**: 高
- **实现复杂度**: 中等

## 使用场景

- 减少锁竞争
- 高并发环境
- 大缓存大小
- 需要高并发性能的场景

## 快速开始

### 安装

```bash
go get github.com/lazygophers/utils/cache/slru
```

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/slru"
)

func main() {
    // 创建容量为 1000，分为 4 个段的缓存
    cache := slru.New(1000, 4)

    // 设置值
    cache.Set("key1", "value1")
    cache.Set("key2", "value2")

    // 获取值
    if value, ok := cache.Get("key1"); ok {
        fmt.Println("Found:", value)
    }

    // 删除值
    cache.Delete("key1")

    // 清空缓存
    cache.Clear()
}
```

### 高级使用

```go
// 带过期时间的缓存
cache.SetWithTTL("key", "value", time.Minute*5)

// 获取缓存统计
stats := cache.Stats()
fmt.Printf("Size: %d\n", stats.Size)
fmt.Printf("Hits: %d\n", stats.Hits)
fmt.Printf("Misses: %d\n", stats.Misses)
fmt.Printf("Hit Rate: %.2f%%\n", stats.HitRate())
```

## 工作原理

SLRU 将缓存分为多个段：

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

每个段有自己的锁，多个 goroutine 可以同时访问不同的段，减少锁竞争。

## API 参考

### 构造函数

```go
// 创建新的 SLRU 缓存
func New(capacity int, segments int) *SLRU

// 创建带选项的 SLRU 缓存
func NewWithOpts(opts Options) *SLRU
```

### 主要方法

```go
// 设置键值对
func (c *SLRU) Set(key string, value interface{})

// 设置键值对，带过期时间
func (c *SLRU) SetWithTTL(key string, value interface{}, ttl time.Duration)

// 获取值
func (c *SLRU) Get(key string) (interface{}, bool)

// 删除键
func (c *SLRU) Delete(key string)

// 清空缓存
func (c *SLRU) Clear()

// 获取统计信息
func (c *SLRU) Stats() Stats
```

## 性能特点

- **时间复杂度**:
  - Set: O(1)
  - Get: O(1)
  - Delete: O(1)
- **空间复杂度**: O(n)，其中 n 是缓存容量
- **并发性能**: 比标准 LRU 高 2-3 倍

## 最佳实践

1. **选择合适的段数**: 通常 4-8 个段效果最好
2. **高并发环境首选**: SLRU 在高并发下性能最佳
3. **监控段分布**: 定期检查各段的使用情况，确保负载均衡

## 相关文档

- [缓存概览](./index.md)
- [LRU 缓存](./lru.md)
- [缓存选择指南](./index.md#快速选择指南)
