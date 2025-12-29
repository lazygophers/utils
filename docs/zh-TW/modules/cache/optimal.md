---
title: Optimal 緩存
---

# Optimal 緩存

Optimal 是理論上性能最佳的緩存。

## 概覽

Optimal 緩存是一種理論上最佳的緩存策略，假設知道未來的訪問模式。當緩存已滿時，它會淘汰在未來最久之後才會被訪問的項目，提供理論上最高的命中率。主要用於性能基準測試和離線分析。

## 特點

- **命中率**: 95%（理論上最佳）
- **內存佔用**: 高
- **並發性能**: 低
- **實現複雜度**: 複雜

## 使用場景

- 最大命中率需求
- 可預測的訪問模式
- 離線分析
- 性能基準測試

## 快速開始

### 安裝

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
    // 創建容量為 1000 的緩存，需要提供未來訪問序列
    futureAccess := []string{"key1", "key2", "key3"}
    cache := optimal.New(1000, futureAccess)

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

Optimal 緩存需要事先知道訪問序列：

```
訪問序列: [key1, key2, key1, key3, key2, ...]

淘汰決策:
- key1: 下次訪問 = 3
- key2: 下次訪問 = 5
- key3: 下次訪問 = ∞ (從不再次訪問)

淘汰 key3，因為它的下次訪問最久
```

## API 參考

### 構造函數

```go
// 創建新的 Optimal 緩存，需要提供未來訪問序列
func New(capacity int, futureAccess []string) *Optimal

// 使用選項創建新的 Optimal 緩存
func NewWithOpts(opts Options) *Optimal
```

### 主要方法

```go
// 設置鍵值對
func (c *Optimal) Set(key string, value interface{})

// 獲取值
func (c *Optimal) Get(key string) (interface{}, bool)

// 刪除鍵
func (c *Optimal) Delete(key string)

// 清空緩存
func (c *Optimal) Clear()

// 獲取統計信息
func (c *Optimal) Stats() Stats
```

## 性能特徵

- **時間複雜度**:
  - Set: O(n)
  - Get: O(1)
  - Delete: O(n)
- **空間複雜度**: O(n)，其中 n 是緩存容量

## 最佳實踐

1. **理論基準測試**: Optimal 主要用於理論性能基準測試
2. **離線分析**: 適用於離線分析和性能評估
3. **不適用於生產環境**: 由於需要事先知道訪問模式，不適用於生產環境

## 限制

- **需要未來訪問序列**: 必須事先知道訪問模式
- **不適用於線上場景**: 不能用於實時線上系統
- **主要用於研究**: 主要用於研究和基準測試

## 相關文檔

- [緩存概覽](./index.md)
- [LRU 緩存](./lru.md)
- [TinyLFU 緩存](./tinylfu.md)