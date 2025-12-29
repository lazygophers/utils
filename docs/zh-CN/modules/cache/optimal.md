---
title: Optimal 缓存
---

# Optimal 缓存

Optimal 是用于理论性能的最优缓存。

## 概述

Optimal 缓存是一种理论上的最优缓存策略，假设知道未来的访问模式。它淘汰最久不会被访问的项目，提供理论上最高的命中率。主要用于性能基准测试和离线分析。

## 特性

- **命中率**: 95% (理论最优)
- **内存占用**: 高
- **并发性能**: 低
- **实现复杂度**: 复杂

## 使用场景

- 最大命中率
- 可预测的访问模式
- 离线分析
- 性能基准测试

## 快速开始

### 安装

```bash
go get github.com/lazygophers/utils/cache/optimal
```

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/optimal"
)

func main() {
    // 创建容量为 1000 的缓存，需要提供未来访问序列
    futureAccess := []string{"key1", "key2", "key3"}
    cache := optimal.New(1000, futureAccess)

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

Optimal 缓存需要预先知道访问序列：

```
访问序列: [key1, key2, key1, key3, key2, ...]

淘汰决策:
- key1: 下次访问 = 3
- key2: 下次访问 = 5
- key3: 下次访问 = ∞ (不再访问)

淘汰 key3，因为它的下次访问最远
```

## API 参考

### 构造函数

```go
// 创建新的 Optimal 缓存，需要提供未来访问序列
func New(capacity int, futureAccess []string) *Optimal

// 创建带选项的 Optimal 缓存
func NewWithOpts(opts Options) *Optimal
```

### 主要方法

```go
// 设置键值对
func (c *Optimal) Set(key string, value interface{})

// 获取值
func (c *Optimal) Get(key string) (interface{}, bool)

// 删除键
func (c *Optimal) Delete(key string)

// 清空缓存
func (c *Optimal) Clear()

// 获取统计信息
func (c *Optimal) Stats() Stats
```

## 性能特点

- **时间复杂度**:
  - Set: O(n)
  - Get: O(1)
  - Delete: O(n)
- **空间复杂度**: O(n)，其中 n 是缓存容量

## 最佳实践

1. **理论基准**: Optimal 主要用于理论性能基准测试
2. **离线分析**: 适合离线分析和性能评估
3. **不用于生产**: 由于需要预先知道访问模式，不适合生产环境

## 局限性

- **需要未来访问序列**: 必须预先知道访问模式
- **不适合在线场景**: 无法用于实时在线系统
- **主要用于研究**: 主要用于研究和基准测试

## 相关文档

- [缓存概览](./index.md)
- [LRU 缓存](./lru.md)
- [TinyLFU 缓存](./tinylfu.md)
