---
title: LFU 缓存
---

# LFU (Least Frequently Used) 缓存

LFU 淘汰最少使用的项目，适合不常访问的数据场景。

## 概述

LFU（Least Frequently Used）基于频率局部性原理，假设频繁访问的数据应该保留。当缓存满时，淘汰访问频率最低的数据。

## 特性

- **命中率**: 75%
- **内存占用**: 低
- **并发性能**: 中等
- **实现复杂度**: 中等

## 使用场景

- 不常访问的数据
- 大数据集
- 内存受限环境
- 冷热数据分离场景

## 快速开始

### 安装

```bash
go get github.com/lazygophers/utils/cache/lfu
```

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/lfu"
)

func main() {
    // 创建容量为 1000 的缓存
    cache := lfu.New(1000)

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

## API 参考

### 构造函数

```go
// 创建新的 LFU 缓存
func New(capacity int) *LFU

// 创建带选项的 LFU 缓存
func NewWithOpts(opts Options) *LFU
```

### 主要方法

```go
// 设置键值对
func (c *LFU) Set(key string, value interface{})

// 设置键值对，带过期时间
func (c *LFU) SetWithTTL(key string, value interface{}, ttl time.Duration)

// 获取值
func (c *LFU) Get(key string) (interface{}, bool)

// 删除键
func (c *LFU) Delete(key string)

// 清空缓存
func (c *LFU) Clear()

// 获取统计信息
func (c *LFU) Stats() Stats
```

## 性能特点

- **时间复杂度**:
  - Set: O(log n)
  - Get: O(log n)
  - Delete: O(log n)
- **空间复杂度**: O(n)，其中 n 是缓存容量

## 最佳实践

1. **适用于冷数据**: LFU 适合访问频率差异大的场景
2. **避免缓存污染**: 对于突发访问，考虑使用其他策略
3. **监控频率分布**: 定期检查访问频率分布，确认 LFU 是否合适

## 相关文档

- [缓存概览](./index.md)
- [LRU 缓存](./lru.md)
- [ALFU 缓存](./alfu.md)
