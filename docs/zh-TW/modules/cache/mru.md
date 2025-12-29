---
title: MRU 緩存
---

# MRU (最近最多使用) 緩存

MRU 淘汰最近使用的項目，適用於具有強時間局部性的場景。

## 概覽

MRU (最近最多使用) 基於時間局部性原理，假設最近訪問的數據不會再次被訪問。當緩存已滿時，它會淘汰最近使用的數據。

## 特點

- **命中率**: 80%
- **內存佔用**: 低
- **並發性能**: 中
- **實現複雜度**: 簡單

## 使用場景

- 時間局部性
- 順序訪問模式
- 緩存預熱
- 掃描式訪問

## 快速開始

### 安裝

```bash
go get github.com/lazygophers/utils/cache/mru
```

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/mru"
)

func main() {
    // 創建容量為 1000 的緩存
    cache := mru.New(1000)

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

## API 參考

### 構造函數

```go
// 創建新的 MRU 緩存
func New(capacity int) *MRU

// 使用選項創建新的 MRU 緩存
func NewWithOpts(opts Options) *MRU
```

### 主要方法

```go
// 設置鍵值對
func (c *MRU) Set(key string, value interface{})

// 獲取值
func (c *MRU) Get(key string) (interface{}, bool)

// 刪除鍵
func (c *MRU) Delete(key string)

// 清空緩存
func (c *MRU) Clear()

// 獲取統計信息
func (c *MRU) Stats() Stats
```

## 性能特徵

- **時間複雜度**:
  - Set: O(1)
  - Get: O(1)
  - Delete: O(1)
- **空間複雜度**: O(n)，其中 n 是緩存容量

## 最佳實踐

1. **順序訪問場景**: MRU 適用於掃描式訪問
2. **緩存預熱**: 在預熱階段，MRU 可以快速淘汰預熱數據
3. **避免頻繁訪問**: 不適合相同數據被頻繁訪問的場景

## 相關文檔

- [緩存概覽](./index.md)
- [LRU 緩存](./lru.md)
- [緩存選擇指南](./index.md#快速選擇指南)