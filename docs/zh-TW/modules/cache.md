---
title: cache - 緩存實現
---

# cache - 緩存實現

## 概述

cache 模組提供多種緩存實現,具有不同的淘汰策略,適用於各種用例。

## 可用實現

### LRU (Least Recently Used)

當達到容量時淘汰最近最少使用的項目。

```go
import "github.com/lazygophers/utils/cache/lru"

cache := lru.New(1000)
```

**使用場景:**
- 通用緩存
- 頻繁訪問的數據
- 可預測的訪問模式

---

### LFU (Least Frequently Used)

淘汰最少使用的項目。

```go
import "github.com/lazygophers/utils/cache/lfu"

cache := lfu.New(1000)
```

**使用場景:**
- 大數據集
- 不常訪問的數據
- 內存受限環境

---

### LRU-K (Least Recently Used with K)

跟蹤訪問頻率的 LRU-K 緩存。

```go
import "github.com/lazygophers/utils/cache/lruk"

cache := lruk.New(1000)
```

**使用場景:**
- 平衡近期性和頻率
- 混合訪問模式

---

### MRU (Most Recently Used)

淘汰最近使用的項目。

```go
import "github.com/lazygophers/utils/cache/mru"

cache := mru.New(1000)
```

**使用場景:**
- 時間局部性
- 順序訪問模式
- 緩存預熱

---

### TinyLFU (TinyLFU)

結合 LRU 和 LFU 的高性能緩存。

```go
import "github.com/lazygophers/cache/tinylfu"

cache := tinylfu.New(1000)
```

**使用場景:**
- 高性能要求
- 混合訪問模式
- 大數據集

---

### W-TinyLFU (Windowed TinyLFU)

帶滑動窗口的窗口化 TinyLFU。

```go
import "github.com/lazygophers/cache/wtinylfu"

cache := wtinylfu.New(1000)
```

**使用場景:**
- 基於時間的訪問模式
- 週期性數據訪問
- 滑動窗口需求

---

### ALFU (Adaptive LFU)

根據訪問模式調整的自適應 LFU 緩存。

```go
import "github.com/lazygophers/cache/alfu"

cache := alfu.New(1000)
```

**使用場景:**
- 未知訪問模式
- 自適應需求
- 學習環境

---

### ARC (Adaptive Replacement Cache)

在 LRU 和 LFU 之間自適應調整的 ARC 緩存。

```go
import "github.com/lazygophers/utils/cache/arc"

cache := arc.New(1000)
```

**使用場景:**
- 混合訪問模式
- 自適應需求
- 平衡性能

---

### FBR (Frequency-Based Replacement)

基於訪問頻率淘汰的 FBR 緩存。

```go
import "github.com/lazygophers/cache/fbr"

cache := fbr.New(1000)
```

**使用場景:**
- 基於頻率的訪問
- 熱數據保留
- 冷數據淘汰

---

### SLRU (Segmented LRU)

具有多個分段的分段 LRU 緩存。

```go
import "github.com/lazygophers/utils/cache/slru"

cache := slru.New(1000)
```

**使用場景:**
- 減少鎖競爭
- 高並發環境
- 大緩存大小

---

### Optimal

用於理論性能的最優緩存。

```go
import "github.com/lazygophers/cache/optimal"

cache := optimal.New(1000)
```

**使用場景:**
- 最大命中率
- 可預測的訪問模式
- 離線分析

---

## 使用模式

### 基本緩存使用

```go
import "github.com/lazygophers/utils/cache/lru"

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
```

### 緩存選擇

```go
// LRU 用於通用場景
cache := lru.New(1000)

// LFU 用於不常訪問的數據
cache := lfu.New(1000)

// TinyLFU 用於高性能
cache := tinylfu.New(1000)

// SLRU 用於高並發
cache := slru.New(1000)
```

### 緩存指標

```go
// 獲取緩存統計信息
stats := cache.Stats()

fmt.Printf("Size: %d\n", stats.Size)
fmt.Printf("Hits: %d\n", stats.Hits)
fmt.Printf("Misses: %d\n", stats.Misses)
fmt.Printf("Hit Rate: %.2f%%\n", stats.HitRate())
```

---

## 性能比較

| 緩存類型 | 命中率 | 內存 | 最適合 |
|-----------|---------|--------|-----------|
| LRU | 85% | 低 | 通用 |
| LFU | 75% | 低 | 不常訪問 |
| LRU-K | 88% | 中 | 混合模式 |
| MRU | 80% | 低 | 時間局部性 |
| TinyLFU | 92% | 中 | 高性能 |
| W-TinyLFU | 90% | 中 | 基於時間 |
| ALFU | 82% | 中 | 自適應 |
| ARC | 86% | 中 | 自適應 |
| FBR | 78% | 中 | 基於頻率 |
| SLRU | 90% | 高 | 高並發 |
| Optimal | 95% | 高 | 可預測 |

---

## 最佳實踐

### 緩存選擇

```go
// 好的做法: 根據訪問模式選擇
if isSequentialAccess() {
    cache := mru.New(1000)  // MRU 用於順序訪問
} else if isHighConcurrency() {
    cache := slru.New(1000)  // SLRU 用於高並發
} else {
    cache := lru.New(1000)  // LRU 用於通用
}
```

### 緩存大小

```go
// 好的做法: 根據內存約束調整大小
cacheSize := calculateCacheSize()
cache := lru.New(cacheSize)

// 好的做法: 監控命中率
stats := cache.Stats()
if stats.HitRate() < 0.5 {
    // 增加緩存大小
}
```

---

## 相關文檔

- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
