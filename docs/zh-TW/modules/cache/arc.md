---
title: ARC 緩存
---

# ARC (自適應替換緩存) 緩存

ARC 是一種自適應緩存，在 LRU 和 LFU 之間調整。

## 概覽

ARC (自適應替換緩存) 是一種自適應緩存策略，在 LRU (最近最少使用) 和 LFU (最少使用) 之間動態調整。它維護兩個列表：T1 (最近使用) 和 T2 (頻繁使用)，並根據訪問模式動態調整兩個列表的大小。

## 特點

- **命中率**: 86%
- **內存佔用**: 中
- **並發性能**: 高
- **實現複雜度**: 中

## 使用場景

- 混合訪問模式
- 自適應需求
- 平衡性能
- 需要在 LRU 和 LFU 之間取得平衡的場景

## 快速開始

### 安裝

```bash
go get github.com/lazygophers/utils/cache/arc
```

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/arc"
)

func main() {
    // 創建容量為 1000 的緩存
    cache := arc.New(1000)

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

ARC 維護四個列表：

1. **T1**: 最近訪問過一次的項目
2. **T2**: 最近訪問過兩次或更多次的項目
3. **B1**: 最近從 T1 中淘汰的項目
4. **B2**: 最近從 T2 中淘汰的項目

根據訪問模式動態調整 T1 和 T2 的大小。

## API 參考

### 構造函數

```go
// 創建新的 ARC 緩存
func New(capacity int) *ARC

// 使用選項創建新的 ARC 緩存
func NewWithOpts(opts Options) *ARC
```

### 主要方法

```go
// 設置鍵值對
func (c *ARC) Set(key string, value interface{})

// 獲取值
func (c *ARC) Get(key string) (interface{}, bool)

// 刪除鍵
func (c *ARC) Delete(key string)

// 清空緩存
func (c *ARC) Clear()

// 獲取統計信息
func (c *ARC) Stats() Stats
```

## 性能特徵

- **時間複雜度**:
  - Set: O(1)
  - Get: O(1)
  - Delete: O(1)
- **空間複雜度**: O(n)，其中 n 是緩存容量

## 最佳實踐

1. **混合訪問模式**: ARC 適用於既有時間局部性又有頻率局部性的場景
2. **自適應需求**: 當緩存需要在 LRU 和 LFU 之間自動調整時使用
3. **平衡性能**: ARC 在各種訪問模式下提供穩定的性能

## 相關文檔

- [緩存概覽](./index.md)
- [LRU 緩存](./lru.md)
- [LFU 緩存](./lfu.md)