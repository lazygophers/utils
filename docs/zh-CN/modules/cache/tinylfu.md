---
title: TinyLFU 缓存
---

# TinyLFU 缓存

TinyLFU 结合 LRU 和 LFU 的高性能缓存，提供最佳的命中率。

## 概述

TinyLFU 是一种混合缓存策略，结合了 LRU（最近使用）和 LFU（最少使用）的优点。它使用两个小缓存：一个跟踪最近访问的项目（LFU），另一个跟踪访问频率（LFU），在淘汰时综合考虑两者。

## 特性

- **命中率**: 92% (最高)
- **内存占用**: 中等
- **并发性能**: 高
- **实现复杂度**: 复杂

## 使用场景

- 高性能要求
- 混合访问模式
- 大数据集
- 需要最佳命中率的场景

## 快速开始

### 安装

```bash
go get github.com/lazygophers/utils/cache/tinylfu
```

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/tinylfu"
)

func main() {
    // 创建容量为 1000 的缓存
    cache := tinylfu.New(1000)

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

TinyLFU 使用两个小缓存：

1. **Window Cache (LRU)**: 跟踪最近访问的项目
2. **Main Cache (LFU)**: 跟踪访问频率

当需要淘汰时：
- 优先淘汰 Window Cache 中的项目
- 如果 Window Cache 为空，淘汰 Main Cache 中频率最低的项目

## API 参考

### 构造函数

```go
// 创建新的 TinyLFU 缓存
func New(capacity int) *TinyLFU

// 创建带选项的 TinyLFU 缓存
func NewWithOpts(opts Options) *TinyLFU
```

### 主要方法

```go
// 设置键值对
func (c *TinyLFU) Set(key string, value interface{})

// 设置键值对，带过期时间
func (c *TinyLFU) SetWithTTL(key string, value interface{}, ttl time.Duration)

// 获取值
func (c *TinyLFU) Get(key string) (interface{}, bool)

// 删除键
func (c *TinyLFU) Delete(key string)

// 清空缓存
func (c *TinyLFU) Clear()

// 获取统计信息
func (c *TinyLFU) Stats() Stats
```

## 性能特点

- **时间复杂度**:
  - Set: O(1) 平均
  - Get: O(1) 平均
  - Delete: O(1)
- **空间复杂度**: O(n)，其中 n 是缓存容量

## 最佳实践

1. **高性能场景首选**: TinyLFU 提供最高的命中率
2. **内存充足**: 需要额外的内存来维护两个小缓存
3. **混合访问模式**: 既有时间局部性又有频率局部性的场景最佳

## 相关文档

- [缓存概览](./index.md)
- [LRU 缓存](./lru.md)
- [LFU 缓存](./lfu.md)
