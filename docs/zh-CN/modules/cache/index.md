---
title: 缓存实现
---

# 缓存实现

cache 模块提供多种缓存实现，具有不同的淘汰策略，适用于各种使用场景。

## 如何选择合适的缓存

选择缓存实现时，需要考虑以下因素：

### 1. 访问模式
- **时间局部性**：最近访问的数据很可能再次被访问 → 选择 MRU
- **频率局部性**：频繁访问的数据应该保留 → 选择 LFU
- **混合模式**：既有时间局部性又有频率局部性 → 选择 TinyLFU 或 LRU-K

### 2. 并发需求
- **低并发**：标准 LRU 即可
- **高并发**：选择 SLRU（分段 LRU）以减少锁竞争

### 3. 性能要求
- **一般性能**：LRU（85% 命中率）
- **高性能**：TinyLFU（92% 命中率）或 Optimal（95% 命中率）
- **内存敏感**：LFU（内存占用最低）

### 4. 自适应需求
- **已知访问模式**：根据模式选择对应策略
- **未知访问模式**：选择 ALFU（自适应 LFU）或 ARC（自适应替换）

## 缓存实现对比

| 缓存类型 | 命中率 | 内存占用 | 并发性能 | 最适合场景 | 推荐度 |
|-----------|---------|---------|-----------|------------|---------|
| **[LRU](./lru.md)** | 85% | 低 | 中 | 通用缓存、频繁访问的数据 | ⭐⭐⭐⭐⭐ |
| **[LFU](./lfu.md)** | 75% | 低 | 中 | 不常访问的数据、内存受限 | ⭐⭐⭐ |
| **[LRU-K](./lruk.md)** | 88% | 中 | 中 | 平衡近期性和频率、混合模式 | ⭐⭐⭐⭐⭐ |
| **[MRU](./mru.md)** | 80% | 低 | 中 | 时间局部性、顺序访问 | ⭐⭐⭐ |
| **[TinyLFU](./tinylfu.md)** | 92% | 中 | 高 | 高性能要求、混合模式 | ⭐⭐⭐⭐⭐⭐ |
| **[W-TinyLFU](./wtinylfu.md)** | 90% | 中 | 高 | 基于时间的访问模式、周期性数据 | ⭐⭐⭐⭐ |
| **[ALFU](./alfu.md)** | 82% | 中 | 中 | 未知访问模式、自适应需求 | ⭐⭐⭐ |
| **[ARC](./arc.md)** | 86% | 中 | 高 | 混合访问模式、自适应 | ⭐⭐⭐⭐⭐ |
| **[FBR](./fbr.md)** | 78% | 中 | 中 | 基于频率的访问、热数据保留 | ⭐⭐⭐ |
| **[SLRU](./slru.md)** | 90% | 高 | 高 | 高并发环境、大缓存大小 | ⭐⭐⭐⭐⭐ |
| **[Optimal](./optimal.md)** | 95% | 高 | 低 | 可预测的访问模式、离线分析 | ⭐⭐⭐ |

## 快速选择指南

### 根据场景选择

```go
// 场景 1: 通用 Web 应用缓存
import "github.com/lazygophers/utils/cache/lru"
cache := lru.New(1000)  // LRU - 最通用

// 场景 2: 高并发 API 缓存
import "github.com/lazygophers/utils/cache/slru"
cache := slru.New(1000)  // SLRU - 减少锁竞争

// 场景 3: 高性能要求
import "github.com/lazygophers/utils/cache/tinylfu"
cache := tinylfu.New(1000)  // TinyLFU - 最高命中率

// 场景 4: 未知访问模式
import "github.com/lazygophers/utils/cache/alfu"
cache := alfu.New(1000)  // ALFU - 自适应

// 场景 5: 顺序访问数据
import "github.com/lazygophers/utils/cache/mru"
cache := mru.New(1000)  // MRU - 时间局部性
```

### 决策树

```
是否已知访问模式？
├─ 是
│  ├─ 高并发？ → SLRU
│  ├─ 顺序访问？ → MRU
│  ├─ 高性能要求？ → TinyLFU
│  └─ 通用场景？ → LRU
└─ 否
   ├─ 需要自适应？ → ALFU 或 ARC
   └─ 可预测模式？ → Optimal
```

## 基本使用示例

### 创建缓存

```go
import "github.com/lazygophers/utils/cache/lru"

// 创建容量为 1000 的缓存
cache := lru.New(1000)
```

### 基本操作

```go
// 设置值
cache.Set("key1", "value1")
cache.Set("key2", "value2")

// 获取值
if value, ok := cache.Get("key1"); ok {
    fmt.Println("Found:", value)
}

// 删除值
cache.Delete("key1")

// 清空缓存
cache.Clear()
```

### 缓存统计

```go
// 获取缓存统计信息
stats := cache.Stats()

fmt.Printf("Size: %d\n", stats.Size)
fmt.Printf("Hits: %d\n", stats.Hits)
fmt.Printf("Misses: %d\n", stats.Misses)
fmt.Printf("Hit Rate: %.2f%%\n", stats.HitRate())
```

## 最佳实践

### 1. 缓存大小选择

```go
// 根据可用内存调整缓存大小
func calculateCacheSize() int {
    // 好的做法：基于内存约束
    availableMem := getAvailableMemory()
    return availableMem / 1024  // 每个条目约 1KB
}

cache := lru.New(calculateCacheSize())
```

### 2. 监控命中率

```go
// 定期检查命中率
func monitorCache(cache Cache) {
    stats := cache.Stats()
    if stats.HitRate() < 0.5 {
        // 命中率过低，考虑：
        // 1. 增加缓存大小
        // 2. 更换缓存策略
        // 3. 检查访问模式
    }
}
```

### 3. 选择合适的缓存类型

```go
// 根据实际场景选择
func createCache() Cache {
    if isSequentialAccess() {
        return mru.New(1000)  // MRU 用于顺序访问
    } else if isHighConcurrency() {
        return slru.New(1000)  // SLRU 用于高并发
    } else if isHighPerformance() {
        return tinylfu.New(1000)  // TinyLFU 用于高性能
    } else {
        return lru.New(1000)  // LRU 用于通用
    }
}
```

## 相关文档

- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
