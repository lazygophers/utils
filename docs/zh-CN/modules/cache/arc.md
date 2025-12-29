---
title: ARC 缓存
---

# ARC (Adaptive Replacement Cache) 缓存

ARC 是在 LRU 和 LFU 之间自适应调整的缓存。

## 概述

ARC（Adaptive Replacement Cache）是一种自适应缓存策略，动态地在 LRU（最近使用）和 LFU（最少使用）之间调整。它维护两个列表：T1（最近使用）和 T2（频繁使用），并根据访问模式动态调整两个列表的大小。

## 特性

- **命中率**: 86%
- **内存占用**: 中等
- **并发性能**: 高
- **实现复杂度**: 中等

## 使用场景

- 混合访问模式
- 自适应需求
- 平衡性能
- 需要在 LRU 和 LFU 之间平衡的场景

## 快速开始

### 安装

```bash
go get github.com/lazygophers/utils/cache/arc
```

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/arc"
)

func main() {
    // 创建容量为 1000 的缓存
    cache := arc.New(1000)

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

ARC 维护四个列表：

1. **T1**: 最近使用一次的项目
2. **T2**: 最近使用两次或以上的项目
3. **B1**: 最近从 T1 淘汰的项目
4. **B2**: 最近从 T2 淘汰的项目

根据访问模式动态调整 T1 和 T2 的大小。

## API 参考

### 构造函数

```go
// 创建新的 ARC 缓存
func New(capacity int) *ARC

// 创建带选项的 ARC 缓存
func NewWithOpts(opts Options) *ARC
```

### 主要方法

```go
// 设置键值对
func (c *ARC) Set(key string, value interface{})

// 获取值
func (c *ARC) Get(key string) (interface{}, bool)

// 删除键
func (c *ARC) Delete(key string)

// 清空缓存
func (c *ARC) Clear()

// 获取统计信息
func (c *ARC) Stats() Stats
```

## 性能特点

- **时间复杂度**:
  - Set: O(1)
  - Get: O(1)
  - Delete: O(1)
- **空间复杂度**: O(n)，其中 n 是缓存容量

## 最佳实践

1. **混合访问模式**: ARC 适合既有时间局部性又有频率局部性的场景
2. **自适应需求**: 当需要缓存自动在 LRU 和 LFU 之间调整时使用
3. **平衡性能**: ARC 在各种访问模式下都能提供稳定的性能

## 相关文档

- [缓存概览](./index.md)
- [LRU 缓存](./lru.md)
- [LFU 缓存](./lfu.md)
