---
title: W-TinyLFU 缓存
---

# W-TinyLFU (Windowed TinyLFU) 缓存

W-TinyLFU 是带滑动窗口的窗口化 TinyLFU。

## 概述

W-TinyLFU（Windowed TinyLFU）在 TinyLFU 的基础上增加了滑动窗口机制。它使用一个时间窗口来跟踪访问频率，能够更好地处理基于时间的访问模式和周期性数据访问。

## 特性

- **命中率**: 90%
- **内存占用**: 中等
- **并发性能**: 高
- **实现复杂度**: 复杂

## 使用场景

- 基于时间的访问模式
- 周期性数据访问
- 滑动窗口需求
- 需要时间窗口统计的场景

## 快速开始

### 安装

```bash
go get github.com/lazygophers/utils/cache/wtinylfu
```

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/wtinylfu"
)

func main() {
    // 创建容量为 1000，窗口大小为 100 的缓存
    cache := wtinylfu.New(1000, 100)

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

W-TinyLFU 使用滑动窗口：

```
时间窗口
┌────────────────────────────┐
│  [访问1, 访问2, ...]    │  ← 窗口内
│  [访问1, 访问2, ...]    │  ← 窗口外
└────────────────────────────┘
    ↑ 窗口滑动
```

窗口内的访问用于计算频率，窗口外的访问逐渐衰减。

## API 参考

### 构造函数

```go
// 创建新的 W-TinyLFU 缓存
func New(capacity int, windowSize int) *WTinyLFU

// 创建带选项的 W-TinyLFU 缓存
func NewWithOpts(opts Options) *WTinyLFU
```

### 主要方法

```go
// 设置键值对
func (c *WTinyLFU) Set(key string, value interface{})

// 获取值
func (c *WTinyLFU) Get(key string) (interface{}, bool)

// 删除键
func (c *WTinyLFU) Delete(key string)

// 清空缓存
func (c *WTinyLFU) Clear()

// 获取统计信息
func (c *WTinyLFU) Stats() Stats
```

## 性能特点

- **时间复杂度**:
  - Set: O(1) 平均
  - Get: O(1) 平均
  - Delete: O(1)
- **空间复杂度**: O(n)，其中 n 是缓存容量

## 最佳实践

1. **选择合适的窗口大小**: 根据访问模式调整窗口大小
2. **周期性访问**: W-TinyLFU 特别适合周期性访问模式
3. **时间窗口统计**: 适合需要基于时间窗口统计的场景

## 相关文档

- [缓存概览](./index.md)
- [TinyLFU 缓存](./tinylfu.md)
- [缓存选择指南](./index.md#快速选择指南)
