---
title: FBR 缓存
---

# FBR (Frequency-Based Replacement) 缓存

FBR 是基于访问频率淘汰的 FBR 缓存。

## 概述

FBR（Frequency-Based Replacement）是一种基于访问频率的缓存淘汰策略。它跟踪每个键的访问频率，当缓存满时，淘汰访问频率最低的键。

## 特性

- **命中率**: 78%
- **内存占用**: 中等
- **并发性能**: 中等
- **实现复杂度**: 中等

## 使用场景

- 基于频率的访问
- 热数据保留
- 冷数据淘汰
- 需要明确区分热冷数据的场景

## 快速开始

### 安装

```bash
go get github.com/lazygophers/utils/cache/fbr
```

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/fbr"
)

func main() {
    // 创建容量为 1000 的缓存
    cache := fbr.New(1000)

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

FBR 维护每个键的访问频率：

```
Key1: 访问次数 = 5  (热数据)
Key2: 访问次数 = 1  (冷数据)
Key3: 访问次数 = 0  (未访问)
```

淘汰时优先淘汰访问频率最低的键。

## API 参考

### 构造函数

```go
// 创建新的 FBR 缓存
func New(capacity int) *FBR

// 创建带选项的 FBR 缓存
func NewWithOpts(opts Options) *FBR
```

### 主要方法

```go
// 设置键值对
func (c *FBR) Set(key string, value interface{})

// 获取值
func (c *FBR) Get(key string) (interface{}, bool)

// 删除键
func (c *FBR) Delete(key string)

// 清空缓存
func (c *FBR) Clear()

// 获取统计信息
func (c *FBR) Stats() Stats
```

## 性能特点

- **时间复杂度**:
  - Set: O(log n)
  - Get: O(log n)
  - Delete: O(log n)
- **空间复杂度**: O(n)，其中 n 是缓存容量

## 最佳实践

1. **热数据保留**: FBR 能够有效保留热数据
2. **冷数据淘汰**: 适合需要快速淘汰冷数据的场景
3. **频率差异大**: 适合访问频率差异大的场景

## 相关文档

- [缓存概览](./index.md)
- [LFU 缓存](./lfu.md)
- [ALFU 缓存](./alfu.md)
