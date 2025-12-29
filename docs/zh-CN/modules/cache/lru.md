---
title: LRU 缓存
---

# LRU (Least Recently Used) 缓存

LRU 是最常用的缓存淘汰策略，当达到容量时淘汰最近最少使用的项目。

## 概述

LRU（Least Recently Used）基于时间局部性原理，假设最近访问的数据很可能再次被访问。当缓存满时，淘汰最久未使用的数据。

## 特性

- **命中率**: 85%
- **内存占用**: 低
- **并发性能**: 中等
- **实现复杂度**: 简单

## 使用场景

- 通用缓存
- 频繁访问的数据
- 可预测的访问模式
- Web 应用缓存
- 数据库查询缓存

## 快速开始

### 安装

```bash
go get github.com/lazygophers/utils/cache/lru
```

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/lru"
)

func main() {
    // 创建容量为 1000 的缓存
    cache := lru.New(1000)

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

// 批量操作
cache.SetMany(map[string]interface{}{
    "key1": "value1",
    "key2": "value2",
})

values := cache.GetMany([]string{"key1", "key2"})

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
// 创建新的 LRU 缓存
func New(capacity int) *LRU

// 创建带选项的 LRU 缓存
func NewWithOpts(opts Options) *LRU
```

### 主要方法

```go
// 设置键值对
func (c *LRU) Set(key string, value interface{})

// 设置键值对，带过期时间
func (c *LRU) SetWithTTL(key string, value interface{}, ttl time.Duration)

// 获取值
func (c *LRU) Get(key string) (interface{}, bool)

// 删除键
func (c *LRU) Delete(key string)

// 清空缓存
func (c *LRU) Clear()

// 获取统计信息
func (c *LRU) Stats() Stats
```

## 性能特点

- **时间复杂度**:
  - Set: O(1)
  - Get: O(1)
  - Delete: O(1)
- **空间复杂度**: O(n)，其中 n 是缓存容量

## 最佳实践

1. **选择合适的缓存大小**: 根据可用内存和访问模式调整
2. **监控命中率**: 定期检查命中率，低于 50% 时考虑调整
3. **使用过期时间**: 对于时效性数据，使用 SetWithTTL
4. **批量操作**: 对于大量数据，使用 SetMany 和 GetMany 提高性能

## 相关文档

- [缓存概览](./index.md)
- [LFU 缓存](./lfu.md)
- [TinyLFU 缓存](./tinylfu.md)
