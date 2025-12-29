---
title: FBR 緩存
---

# FBR (基於頻率的替換) 緩存

FBR 是一種基於頻率的替換緩存。

## 概覽

FBR (基於頻率的替換) 是一種基於訪問頻率的緩存淘汰策略。它跟蹤每個鍵的訪問頻率，當緩存已滿時，淘汰訪問頻率最低的鍵。

## 特點

- **命中率**: 78%
- **內存佔用**: 中
- **並發性能**: 中
- **實現複雜度**: 中

## 使用場景

- 基於頻率的訪問
- 熱數據保留
- 冷數據淘汰
- 需要明確區分熱數據和冷數據的場景

## 快速開始

### 安裝

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
    // 創建容量為 1000 的緩存
    cache := fbr.New(1000)

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

FBR 維護每個鍵的訪問頻率：

```
Key1: 訪問次數 = 5  (熱數據)
Key2: 訪問次數 = 1  (冷數據)
Key3: 訪問次數 = 0  (從未訪問)
```

淘汰時，優先淘汰訪問頻率最低的鍵。

## API 參考

### 構造函數

```go
// 創建新的 FBR 緩存
func New(capacity int) *FBR

// 使用選項創建新的 FBR 緩存
func NewWithOpts(opts Options) *FBR
```

### 主要方法

```go
// 設置鍵值對
func (c *FBR) Set(key string, value interface{})

// 獲取值
func (c *FBR) Get(key string) (interface{}, bool)

// 刪除鍵
func (c *FBR) Delete(key string)

// 清空緩存
func (c *FBR) Clear()

// 獲取統計信息
func (c *FBR) Stats() Stats
```

## 性能特徵

- **時間複雜度**:
  - Set: O(log n)
  - Get: O(log n)
  - Delete: O(log n)
- **空間複雜度**: O(n)，其中 n 是緩存容量

## 最佳實踐

1. **熱數據保留**: FBR 可以有效地保留熱數據
2. **冷數據淘汰**: 適用於需要快速淘汰冷數據的場景
3. **大頻率差異**: 適用於訪問頻率差異大的場景

## 相關文檔

- [緩存概覽](./index.md)
- [LFU 緩存](./lfu.md)
- [ALFU 緩存](./alfu.md)