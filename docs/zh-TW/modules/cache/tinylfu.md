---
title: TinyLFU 緩存
---

# TinyLFU 緩存

TinyLFU 結合 LRU 和 LFU 的高性能緩存，提供最佳的命中率。

## 概述

TinyLFU 是一種混合緩存策略，結合了 LRU（最近使用）和 LFU（最少使用）的優點。它使用兩個小緩存：一個跟蹤最近訪問的項目（LFU），另一個跟蹤訪問頻率（LFU），在淘汰時綜合考慮兩者。

## 特性

- **命中率**: 92% (最高)
- **內存佔用**: 中等
- **並發性能**: 高
- **實現複雜度**: 複雜

## 使用場景

- 高性能要求
- 混合訪問模式
- 大數據集
- 需要最佳命中率的場景

## 快速開始

### 安裝

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
    // 創建容量為 1000 的緩存
    cache := tinylfu.New(1000)

    // 設置值
    cache.Set("key1", "value1")
    cache.Set("key2", "value2")

    // 獲取值
    if value, ok := cache.Get("key1"); ok {
        fmt.Println("Found:", value)
    }

    // 刪除值
    cache.Delete("key1")

    // 清空緩存
    cache.Clear()
}
```

### 高級使用

```go
// 帶過期時間的緩存
cache.SetWithTTL("key", "value", time.Minute*5)

// 獲取緩存統計
stats := cache.Stats()
fmt.Printf("Size: %d\n", stats.Size)
fmt.Printf("Hits: %d\n", stats.Hits)
fmt.Printf("Misses: %d\n", stats.Misses)
fmt.Printf("Hit Rate: %.2f%%\n", stats.HitRate())
```

## 工作原理

TinyLFU 使用兩個小緩存：

1. **Window Cache (LRU)**: 跟蹤最近訪問的項目
2. **Main Cache (LFU)**: 跟蹤訪問頻率

當需要淘汰時：
- 優先淘汰 Window Cache 中的項目
- 如果 Window Cache 為空，淘汰 Main Cache 中頻率最低的項目

## API 參考

### 構造函數

```go
// 創建新的 TinyLFU 緩存
func New(capacity int) *TinyLFU

// 創建帶選項的 TinyLFU 緩存
func NewWithOpts(opts Options) *TinyLFU
```

### 主要方法

```go
// 設置鍵值對
func (c *TinyLFU) Set(key string, value interface{})

// 設置鍵值對，帶過期時間
func (c *TinyLFU) SetWithTTL(key string, value interface{}, ttl time.Duration)

// 獲取值
func (c *TinyLFU) Get(key string) (interface{}, bool)

// 刪除鍵
func (c *TinyLFU) Delete(key string)

// 清空緩存
func (c *TinyLFU) Clear()

// 獲取統計信息
func (c *TinyLFU) Stats() Stats
```

## 性能特點

- **時間複雜度**:
  - Set: O(1) 平均
  - Get: O(1) 平均
  - Delete: O(1)
- **空間複雜度**: O(n)，其中 n 是緩存容量

## 最佳實踐

1. **高性能場景首選**: TinyLFU 提供最高的命中率
2. **內存充足**: 需要額外的內存來維護兩個小緩存
3. **混合訪問模式**: 既有時間局部性又有頻率局部性的場景最佳

## 相關文檔

- [緩存概覽](./index.md)
- [LRU 緩存](./lru.md)
- [LFU 緩存](./lfu.md)
