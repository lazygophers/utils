---
title: LRU 緩存
---

# LRU (Least Recently Used) 緩存

LRU 是最常用的緩存淘汰策略，當達到容量時淘汰最近最少使用的項目。

## 概述

LRU（Least Recently Used）基於時間局部性原理，假設最近訪問的數據很可能再次被訪問。當緩存滿時，淘汰最久未使用的數據。

## 特性

- **命中率**: 85%
- **內存佔用**: 低
- **並發性能**: 中等
- **實現複雜度**: 簡單

## 使用場景

- 通用緩存
- 頻繁訪問的數據
- 可預測的訪問模式
- Web 應用緩存
- 數據庫查詢緩存

## 快速開始

### 安裝

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
    // 創建容量為 1000 的緩存
    cache := lru.New(1000)

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

// 批量操作
cache.SetMany(map[string]interface{}{
    "key1": "value1",
    "key2": "value2",
})

values := cache.GetMany([]string{"key1", "key2"})

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
// 創建新的 LRU 緩存
func New(capacity int) *LRU

// 創建帶選項的 LRU 緩存
func NewWithOpts(opts Options) *LRU
```

### 主要方法

```go
// 設置鍵值對
func (c *LRU) Set(key string, value interface{})

// 設置鍵值對，帶過期時間
func (c *LRU) SetWithTTL(key string, value interface{}, ttl time.Duration)

// 獲取值
func (c *LRU) Get(key string) (interface{}, bool)

// 刪除鍵
func (c *LRU) Delete(key string)

// 清空緩存
func (c *LRU) Clear()

// 獲取統計信息
func (c *LRU) Stats() Stats
```

## 性能特點

- **時間複雜度**:
  - Set: O(1)
  - Get: O(1)
  - Delete: O(1)
- **空間複雜度**: O(n)，其中 n 是緩存容量

## 最佳實踐

1. **選擇合適的緩存大小**: 根據可用內存和訪問模式調整
2. **監控命中率**: 定期檢查命中率，低於 50% 時考慮調整
3. **使用過期時間**: 對於時效性數據，使用 SetWithTTL
4. **批量操作**: 對於大量數據，使用 SetMany 和 GetMany 提高性能

## 相關文檔

- [緩存概覽](./index.md)
- [LFU 緩存](./lfu.md)
- [TinyLFU 緩存](./tinylfu.md)
