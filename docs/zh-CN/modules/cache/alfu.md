---
title: ALFU 缓存
---

# ALFU (Adaptive LFU) 缓存

ALFU 是根据访问模式调整的自适应 LFU 缓存。

## 概述

ALFU（Adaptive LFU）是一种自适应的缓存策略，能够根据访问模式动态调整淘汰策略。它在 LFU 的基础上增加了自适应机制，可以适应变化的访问模式。

## 特性

- **命中率**: 82%
- **内存占用**: 中等
- **并发性能**: 中等
- **实现复杂度**: 复杂

## 使用场景

- 未知访问模式
- 自适应需求
- 学习环境
- 访问模式变化的场景

## 快速开始

### 安装

```bash
go get github.com/lazygophers/utils/cache/alfu
```

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/alfu"
)

func main() {
    // 创建容量为 1000 的缓存
    cache := alfu.New(1000)

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

ALFU 使用自适应机制：

1. **频率跟踪**: 跟踪每个键的访问频率
2. **动态调整**: 根据访问模式动态调整淘汰策略
3. **学习模式**: 能够学习访问模式的变化

## API 参考

### 构造函数

```go
// 创建新的 ALFU 缓存
func New(capacity int) *ALFU

// 创建带选项的 ALFU 缓存
func NewWithOpts(opts Options) *ALFU
```

### 主要方法

```go
// 设置键值对
func (c *ALFU) Set(key string, value interface{})

// 获取值
func (c *ALFU) Get(key string) (interface{}, bool)

// 删除键
func (c *ALFU) Delete(key string)

// 清空缓存
func (c *ALFU) Clear()

// 获取统计信息
func (c *ALFU) Stats() Stats
```

## 性能特点

- **时间复杂度**:
  - Set: O(log n)
  - Get: O(log n)
  - Delete: O(log n)
- **空间复杂度**: O(n)，其中 n 是缓存容量

## 最佳实践

1. **未知访问模式**: ALFU 适合访问模式未知或变化的场景
2. **自适应需求**: 当需要缓存自动适应访问模式时使用
3. **监控学习过程**: 观察缓存的学习过程，确认自适应效果

## 相关文档

- [缓存概览](./index.md)
- [LFU 缓存](./lfu.md)
- [ARC 缓存](./arc.md)
