---
title: W-TinyLFU 緩存
---

# W-TinyLFU (窗口式 TinyLFU) 緩存

W-TinyLFU 是具有滑動窗口的窗口式 TinyLFU。

## 概覽

W-TinyLFU (窗口式 TinyLFU) 在 TinyLFU 中添加了滑動窗口機制。它使用時間窗口來跟蹤訪問頻率，更好地處理基於時間的訪問模式和周期性數據訪問。

## 特點

- **命中率**: 90%
- **內存佔用**: 中
- **並發性能**: 高
- **實現複雜度**: 複雜

## 使用場景

- 基於時間的訪問模式
- 周期性數據訪問
- 滑動窗口需求
- 需要時間窗口統計的場景

## 快速開始

### 安裝

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
    // 創建容量為 1000，窗口大小為 100 的緩存
    cache := wtinylfu.New(1000, 100)

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

W-TinyLFU 使用滑動窗口：

```
時間窗口
┌────────────────────────────┐
│  [access1, access2, ...]    │  ← 窗口內
│  [access1, access2, ...]    │  ← 窗口外
└────────────────────────────┘
    ↑ 窗口滑動
```

窗口內的訪問用於計算頻率，窗口外的訪問逐漸衰減。

## API 參考

### 構造函數

```go
// 創建新的 W-TinyLFU 緩存
func New(capacity int, windowSize int) *WTinyLFU

// 使用選項創建新的 W-TinyLFU 緩存
func NewWithOpts(opts Options) *WTinyLFU
```

### 主要方法

```go
// 設置鍵值對
func (c *WTinyLFU) Set(key string, value interface{})

// 獲取值
func (c *WTinyLFU) Get(key string) (interface{}, bool)

// 刪除鍵
func (c *WTinyLFU) Delete(key string)

// 清空緩存
func (c *WTinyLFU) Clear()

// 獲取統計信息
func (c *WTinyLFU) Stats() Stats
```

## 性能特徵

- **時間複雜度**:
  - Set: O(1) 平均
  - Get: O(1) 平均
  - Delete: O(1)
- **空間複雜度**: O(n)，其中 n 是緩存容量

## 最佳實踐

1. **選擇適當的窗口大小**: 根據訪問模式調整窗口大小
2. **周期性訪問**: W-TinyLFU 特別適用於周期性訪問模式
3. **時間窗口統計**: 適用於需要時間窗口統計的場景

## 相關文檔

- [緩存概覽](./index.md)
- [TinyLFU 緩存](./tinylfu.md)
- [緩存選擇指南](./index.md#快速選擇指南)