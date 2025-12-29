---
title: 緩存實現
---

# 緩存實現

cache 模組提供多種緩存實現，具有不同的淘汰策略，適用於各種使用場景。

## 如何選擇合適的緩存

選擇緩存實現時，需要考慮以下因素：

### 1. 訪問模式
- **時間局部性**：最近訪問的數據很可能再次被訪問 → 選擇 MRU
- **頻率局部性**：頻繁訪問的數據應該保留 → 選擇 LFU
- **混合模式**：既有時間局部性又有頻率局部性 → 選擇 TinyLFU 或 LRU-K

### 2. 並發需求
- **低並發**：標準 LRU 即可
- **高並發**：選擇 SLRU（分段 LRU）以減少鎖競爭

### 3. 性能要求
- **一般性能**：LRU（85% 命中率）
- **高性能**：TinyLFU（92% 命中率）或 Optimal（95% 命中率）
- **內存敏感**：LFU（內存佔用最低）

### 4. 自適應需求
- **已知訪問模式**：根據模式選擇對應策略
- **未知訪問模式**：選擇 ALFU（自適應 LFU）或 ARC（自適應替換）

## 緩存實現對比

| 緩存類型 | 命中率 | 內存佔用 | 並發性能 | 最適合場景 | 推薦度 |
|-----------|---------|---------|-----------|------------|---------|
| **[LRU](./lru.md)** | 85% | 低 | 中 | 通用緩存、頻繁訪問的數據 | ⭐⭐⭐⭐ |
| **[LFU](./lfu.md)** | 75% | 低 | 中 | 不常訪問的數據、內存受限 | ⭐⭐⭐ |
| **[LRU-K](./lruk.md)** | 88% | 中 | 中 | 平衡近期性和頻率、混合模式 | ⭐⭐⭐⭐ |
| **[MRU](./mru.md)** | 80% | 低 | 中 | 時間局部性、順序訪問 | ⭐⭐⭐ |
| **[TinyLFU](./tinylfu.md)** | 92% | 中 | 高 | 高性能要求、混合模式 | ⭐⭐⭐⭐⭐ |
| **[W-TinyLFU](./wtinylfu.md)** | 90% | 中 | 高 | 基於時間的訪問模式、週期性數據 | ⭐⭐⭐⭐ |
| **[ALFU](./alfu.md)** | 82% | 中 | 中 | 未知訪問模式、自適應需求 | ⭐⭐⭐ |
| **[ARC](./arc.md)** | 86% | 中 | 高 | 混合訪問模式、自適應 | ⭐⭐⭐⭐ |
| **[FBR](./fbr.md)** | 78% | 中 | 中 | 基於頻率的訪問、熱數據保留 | ⭐⭐⭐ |
| **[SLRU](./slru.md)** | 90% | 高 | 高 | 高並發環境、大緩存大小 | ⭐⭐⭐⭐ |
| **[Optimal](./optimal.md)** | 95% | 高 | 低 | 可預測的訪問模式、離線分析 | ⭐⭐⭐ |

## 快速選擇指南

### 根據場景選擇

```go
// 場景 1: 通用 Web 應用緩存
import "github.com/lazygophers/utils/cache/lru"
cache := lru.New(1000)  // LRU - 最通用

// 場景 2: 高並發 API 緩存
import "github.com/lazygophers/utils/cache/slru"
cache := slru.New(1000)  // SLRU - 減少鎖競爭

// 場景 3: 高性能要求
import "github.com/lazygophers/utils/cache/tinylfu"
cache := tinylfu.New(1000)  // TinyLFU - 最高命中率

// 場景 4: 未知訪問模式
import "github.com/lazygophers/utils/cache/alfu"
cache := alfu.New(1000)  // ALFU - 自適應

// 場景 5: 順序訪問數據
import "github.com/lazygophers/utils/cache/mru"
cache := mru.New(1000)  // MRU - 時間局部性
```

### 決策樹

```
是否已知訪問模式？
├─ 是
│  ├─ 高並發？ → SLRU
│  ├─ 順序訪問？ → MRU
│  ├─ 高性能要求？ → TinyLFU
│  └─ 通用場景？ → LRU
└─ 否
   ├─ 需要自適應？ → ALFU 或 ARC
   └─ 可預測模式？ → Optimal
```

## 基本使用示例

### 創建緩存

```go
import "github.com/lazygophers/utils/cache/lru"

// 創建容量為 1000 的緩存
cache := lru.New(1000)
```

### 基本操作

```go
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

### 緩存統計

```go
// 獲取緩存統計信息
stats := cache.Stats()

fmt.Printf("Size: %d\n", stats.Size)
fmt.Printf("Hits: %d\n", stats.Hits)
fmt.Printf("Misses: %d\n", stats.Misses)
fmt.Printf("Hit Rate: %.2f%%\n", stats.HitRate())
```

## 最佳實踐

### 1. 緩存大小選擇

```go
// 根據可用內存調整緩存大小
func calculateCacheSize() int {
    // 好的做法：基於內存約束
    availableMem := getAvailableMemory()
    return availableMem / 1024  // 每個條目約 1KB
}

cache := lru.New(calculateCacheSize())
```

### 2. 監控命中率

```go
// 定期檢查命中率
func monitorCache(cache Cache) {
    stats := cache.Stats()
    if stats.HitRate() < 0.5 {
        // 命中率過低，考慮：
        // 1. 增加緩存大小
        // 2. 更換緩存策略
        // 3. 檢查訪問模式
    }
}
```

### 3. 選擇合適的緩存類型

```go
// 根據實際場景選擇
func createCache() Cache {
    if isSequentialAccess() {
        return mru.New(1000)  // MRU 用於順序訪問
    } else if isHighConcurrency() {
        return slru.New(1000)  // SLRU 用於高並發
    } else if isHighPerformance() {
        return tinylfu.New(1000)  // TinyLFU 用於高性能
    } else {
        return lru.New(1000)  // LRU 用於通用
    }
}
```

## 相關文檔

- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
