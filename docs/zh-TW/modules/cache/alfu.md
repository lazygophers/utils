---
title: ALFU 緩存
---

# ALFU (自適應 LFU) 緩存

ALFU 是一種自適應 LFU 緩存，根據訪問模式進行調整。

## 概覽

ALFU (自適應 LFU) 是一種自適應緩存策略，根據訪問模式動態調整淘汰策略。它基於 LFU (最少使用) 策略，添加了自適應機制以適應不斷變化的訪問模式。

## 特點

- **命中率**: 82%
- **內存佔用**: 中
- **並發性能**: 中
- **實現複雜度**: 複雜

## 使用場景

- 未知訪問模式
- 自適應需求
- 學習環境
- 訪問模式變化的場景

## 快速開始

### 安裝

```bash
go get github.com/lazygophers/utils/cache/alfu
```

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/cache/alfu"
)

func main() {
    // 創建容量為 1000 的緩存
    cache := alfu.New(1000)

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

ALFU 使用自適應機制：

1. **頻率跟蹤**: 跟蹤每個鍵的訪問頻率
2. **動態調整**: 根據訪問模式動態調整淘汰策略
3. **模式學習**: 可以學習訪問模式的變化

## API 參考

### 構造函數

```go
// 創建新的 ALFU 緩存
func New(capacity int) *ALFU

// 使用選項創建新的 ALFU 緩存
func NewWithOpts(opts Options) *ALFU
```

### 主要方法

```go
// 設置鍵值對
func (c *ALFU) Set(key string, value interface{})

// 獲取值
func (c *ALFU) Get(key string) (interface{}, bool)

// 刪除鍵
func (c *ALFU) Delete(key string)

// 清空緩存
func (c *ALFU) Clear()

// 獲取統計信息
func (c *ALFU) Stats() Stats
```

## 性能特徵

- **時間複雜度**:
  - Set: O(log n)
  - Get: O(log n)
  - Delete: O(log n)
- **空間複雜度**: O(n)，其中 n 是緩存容量

## 最佳實踐

1. **未知訪問模式**: ALFU 適用於訪問模式未知或不斷變化的場景
2. **自適應需求**: 當緩存需要自動適應訪問模式時使用
3. **監控學習過程**: 觀察緩存的學習過程，確認自適應效果

## 相關文檔

- [緩存概覽](./index.md)
- [LFU 緩存](./lfu.md)
- [ARC 緩存](./arc.md)