---
title: LFU 緩存
---

# LFU (Least Frequently Used) 緩存

LFU 淘汰最少使用的項目，適合不常訪問的數據場景。

## 概述

LFU（Least Frequently Used）基於頻率局部性原理，假設頻繁訪問的數據應該保留。當緩存滿時，淘汰訪問頻率最低的數據。

## 特性

- **命中率**: 75%
- **內存佔用**: 低
- **並發性能**: 中等
- **實現複雜度**: 中等

## 使用場景

- 不常訪問的數據
- 大數據集
- 內存受限環境
- 冷熱數據分離場景

## 快速開始

### 安裝

```bash
go get github.com/lazygophers/utils/cache/lfu
```

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/lfu"
)

func main() {
    // 創建容量為 1000 的緩存
    cache := lfu.New(1000)

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

## API 參考

### 構造函數

```go
// 創建新的 LFU 緩存
func New(capacity int) *LFU

// 創建帶選項的 LFU 緩存
func NewWithOpts(opts Options) *LFU
```

### 主要方法

```go
// 設置鍵值對
func (c *LFU) Set(key string, value interface{})

// 設置鍵值對，帶過期時間
func (c *LFU) SetWithTTL(key string, value interface{}, ttl time.Duration)

// 獲取值
func (c *LFU) Get(key string) (interface{}, bool)

// 刪除鍵
func (c *LFU) Delete(key string)

// 清空緩存
func (c *LFU) Clear()

// 獲取統計信息
func (c *LFU) Stats() Stats
```

## 性能特點

- **時間複雜度**:
  - Set: O(log n)
  - Get: O(log n)
  - Delete: O(log n)
- **空間複雜度**: O(n)，其中 n 是緩存容量

## 最佳實踐

1. **適用於冷數據**: LFU 適合訪問頻率差異大的場景
2. **避免緩存污染**: 對於突發訪問，考慮使用其他策略
3. **監控頻率分佈**: 定期檢查訪問頻率分佈，確認 LFU 是否合適

## 相關文檔

- [緩存概覽](./index.md)
- [LRU 緩存](./lru.md)
- [ALFU 緩存](./alfu.md)
