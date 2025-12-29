---
title: LRU-K 缓存
---

# LRU-K (Least Recently Used with K) 缓存

LRU-K 跟踪访问频率的 LRU-K 缓存。

## 概述

LRU-K 是 LRU 的改进版本，跟踪每个键的访问频率（K 次）。当缓存满时，优先淘汰访问频率较低的键，而不是简单地淘汰最久未使用的键。

## 特性

- **命中率**: 88%
- **内存占用**: 中等
- **并发性能**: 中等
- **实现复杂度**: 中等

## 使用场景

- 平衡近期性和频率
- 混合访问模式
- 需要考虑访问频率的场景

## 快速开始

### 安装

```bash
go get github.com/lazygophers/utils/cache/lruk
```

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/lruk"
)

func main() {
    // 创建容量为 1000，K=2 的缓存
    cache := lruk.New(1000, 2)

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

## 工作原理

LRU-K 维护每个键的访问历史（最近 K 次访问）：

```
Key1: [访问1, 访问2, ..., 访问K]
Key2: [访问1, 访问2, ..., 访问K]
```

淘汰时综合考虑：
1. 最近访问时间
2. 访问频率（基于 K 次访问）

## API 参考

### 构造函数

```go
// 创建新的 LRU-K 缓存
func New(capacity int, k int) *LRUK

// 创建带选项的 LRU-K 缓存
func NewWithOpts(opts Options) *LRUK
```

### 主要方法

```go
// 设置键值对
func (c *LRUK) Set(key string, value interface{})

// 获取值
func (c *LRUK) Get(key string) (interface{}, bool)

// 删除键
func (c *LRUK) Delete(key string)

// 清空缓存
func (c *LRUK) Clear()

// 获取统计信息
func (c *LRUK) Stats() Stats
```

## 性能特点

- **时间复杂度**:
  - Set: O(1)
  - Get: O(1)
  - Delete: O(1)
- **空间复杂度**: O(n * k)，其中 n 是缓存容量，k 是跟踪的访问次数

## 最佳实践

1. **选择合适的 K 值**: 通常 K=2 或 K=3 效果较好
2. **平衡近期性和频率**: LRU-K 在两者之间提供平衡
3. **混合访问模式**: 适合既有时间局部性又有频率局部性的场景

## 相关文档

- [缓存概览](./index.md)
- [LRU 缓存](./lru.md)
- [TinyLFU 缓存](./tinylfu.md)
