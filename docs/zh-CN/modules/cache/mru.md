---
title: MRU 缓存
---

# MRU (Most Recently Used) 缓存

MRU 淘汰最近使用的项目，适合时间局部性强的场景。

## 概述

MRU（Most Recently Used）基于时间局部性原理，假设最近访问的数据不会再次被访问。当缓存满时，淘汰最近使用的数据。

## 特性

- **命中率**: 80%
- **内存占用**: 低
- **并发性能**: 中等
- **实现复杂度**: 简单

## 使用场景

- 时间局部性
- 顺序访问模式
- 缓存预热
- 扫描式访问

## 快速开始

### 安装

```bash
go get github.com/lazygophers/utils/cache/mru
```

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/mru"
)

func main() {
    // 创建容量为 1000 的缓存
    cache := mru.New(1000)

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

## API 参考

### 构造函数

```go
// 创建新的 MRU 缓存
func New(capacity int) *MRU

// 创建带选项的 MRU 缓存
func NewWithOpts(opts Options) *MRU
```

### 主要方法

```go
// 设置键值对
func (c *MRU) Set(key string, value interface{})

// 获取值
func (c *MRU) Get(key string) (interface{}, bool)

// 删除键
func (c *MRU) Delete(key string)

// 清空缓存
func (c *MRU) Clear()

// 获取统计信息
func (c *MRU) Stats() Stats
```

## 性能特点

- **时间复杂度**:
  - Set: O(1)
  - Get: O(1)
  - Delete: O(1)
- **空间复杂度**: O(n)，其中 n 是缓存容量

## 最佳实践

1. **顺序访问场景**: MRU 适合扫描式访问
2. **缓存预热**: 在预热阶段使用 MRU 可以快速淘汰预热数据
3. **避免频繁访问**: 不适合频繁访问同一数据的场景

## 相关文档

- [缓存概览](./index.md)
- [LRU 缓存](./lru.md)
- [缓存选择指南](./index.md#快速选择指南)
