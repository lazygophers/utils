---
title: SLRU 緩存
---

# SLRU (Segmented LRU) 緩存

SLRU 具有多個分段的分段 LRU 緩存，適合高並發環境。

## 概述

SLRU（Segmented LRU）將緩存分為多個段（segment），每個段是一個獨立的 LRU 緩存。這種設計可以減少鎖競爭，提高並發性能。

## 特性

- **命中率**: 90%
- **內存佔用**: 高
- **並發性能**: 高
- **實現複雜度**: 中等

## 使用場景

- 減少鎖競爭
- 高並發環境
- 大緩存大小
- 需要高並發性能的場景

## 快速開始

### 安裝

```bash
go get github.com/lazygophers/utils/cache/slru
```

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/slru"
)

func main() {
    // 創建容量為 1000，分為 4 個段的緩存
    cache := slru.New(1000, 4)

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

SLRU 將緩存分為多個段：

```
┌─────────────┐
│  Segment 1 │
├─────────────┤
│  Segment 2 │
├─────────────┤
│  Segment 3 │
├─────────────┤
│  Segment 4 │
└─────────────┘
```

每個段有自己的鎖，多個 goroutine 可以同時訪問不同的段，減少鎖競爭。

## API 參考

### 構造函數

```go
// 創建新的 SLRU 緩存
func New(capacity int, segments int) *SLRU

// 創建帶選項的 SLRU 緩存
func NewWithOpts(opts Options) *SLRU
```

### 主要方法

```go
// 設置鍵值對
func (c *SLRU) Set(key string, value interface{})

// 設置鍵值對，帶過期時間
func (c *SLRU) SetWithTTL(key string, value interface{}, ttl time.Duration)

// 獲取值
func (c *SLRU) Get(key string) (interface{}, bool)

// 刪除鍵
func (c *SLRU) Delete(key string)

// 清空緩存
func (c *SLRU) Clear()

// 獲取統計信息
func (c *SLRU) Stats() Stats
```

## 性能特點

- **時間複雜度**:
  - Set: O(1)
  - Get: O(1)
  - Delete: O(1)
- **空間複雜度**: O(n)，其中 n 是緩存容量
- **並發性能**: 比標準 LRU 高 2-3 倍

## 最佳實踐

1. **選擇合適的段數**: 通常 4-8 個段效果最好
2. **高並發環境首選**: SLRU 在高並發下性能最佳
3. **監控段分佈**: 定期檢查各段的使用情況，確保負載均衡

## 相關文檔

- [緩存概覽](./index.md)
- [LRU 緩存](./lru.md)
- [緩存選擇指南](./index.md#快速選擇指南)
