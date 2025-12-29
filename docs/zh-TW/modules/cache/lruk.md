---
title: LRU-K 緩存
---

# LRU-K (基於 K 次最近使用) 緩存

LRU-K 是一種跟蹤訪問頻率的 LRU 緩存。

## 概覽

LRU-K 是 LRU 的改進版本，跟蹤每個鍵的訪問頻率（K 次）。當緩存已滿時，它優先淘汰訪問頻率較低的鍵，而不是簡單地淘汰最近最少使用的鍵。

## 特點

- **命中率**: 88%
- **內存佔用**: 中
- **並發性能**: 中
- **實現複雜度**: 中

## 使用場景

- 平衡近期性和頻率
- 混合訪問模式
- 需要考慮訪問頻率的場景

## 快速開始

### 安裝

```bash
go get github.com/lazygophers/utils/cache/lruk
```

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/lruk"
)

func main() {
    // 創建容量為 1000，K=2 的緩存
    cache := lruk.New(1000, 2)

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

## 工作原理

LRU-K 維護每個鍵的訪問歷史（最近 K 次訪問）：

```
Key1: [access1, access2, ..., accessK]
Key2: [access1, access2, ..., accessK]
```

淘汰時，考慮兩個因素：
1. 最近訪問時間
2. 訪問頻率（基於最近 K 次訪問）

## API 參考

### 構造函數

```go
// 創建新的 LRU-K 緩存
func New(capacity int, k int) *LRUK

// 使用選項創建新的 LRU-K 緩存
func NewWithOpts(opts Options) *LRUK
```

### 主要方法

```go
// 設置鍵值對
func (c *LRUK) Set(key string, value interface{})

// 獲取值
func (c *LRUK) Get(key string) (interface{}, bool)

// 刪除鍵
func (c *LRUK) Delete(key string)

// 清空緩存
func (c *LRUK) Clear()

// 獲取統計信息
func (c *LRUK) Stats() Stats
```

## 性能特徵

- **時間複雜度**:
  - Set: O(1)
  - Get: O(1)
  - Delete: O(1)
- **空間複雜度**: O(n * k)，其中 n 是緩存容量，k 是跟蹤的訪問次數

## 最佳實踐

1. **選擇適當的 K 值**: 通常 K=2 或 K=3 效果良好
2. **平衡近期性和頻率**: LRU-K 提供了兩者的平衡
3. **混合訪問模式**: 適用於既有時間局部性又有頻率局部性的場景

## 相關文檔

- [緩存概覽](./index.md)
- [LRU 緩存](./lru.md)
- [TinyLFU 緩存](./tinylfu.md)