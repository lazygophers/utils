---
title: cache - 缓存实现
---

# cache - 缓存实现

## 概述

cache 模块提供多种缓存实现,具有不同的淘汰策略,适用于各种用例。

## 可用实现

### LRU (Least Recently Used)

当达到容量时淘汰最近最少使用的项目。

```go
import "github.com/lazygophers/utils/cache/lru"

cache := lru.New(1000)
```

**使用场景:**
- 通用缓存
- 频繁访问的数据
- 可预测的访问模式

---

### LFU (Least Frequently Used)

淘汰最少使用的项目。

```go
import "github.com/lazygophers/utils/cache/lfu"

cache := lfu.New(1000)
```

**使用场景:**
- 大数据集
- 不常访问的数据
- 内存受限环境

---

### LRU-K (Least Recently Used with K)

跟踪访问频率的 LRU-K 缓存。

```go
import "github.com/lazygophers/utils/cache/lruk"

cache := lruk.New(1000)
```

**使用场景:**
- 平衡近期性和频率
- 混合访问模式

---

### MRU (Most Recently Used)

淘汰最近使用的项目。

```go
import "github.com/lazygophers/utils/cache/mru"

cache := mru.New(1000)
```

**使用场景:**
- 时间局部性
- 顺序访问模式
- 缓存预热

---

### TinyLFU (TinyLFU)

结合 LRU 和 LFU 的高性能缓存。

```go
import "github.com/lazygophers/cache/tinylfu"

cache := tinylfu.New(1000)
```

**使用场景:**
- 高性能要求
- 混合访问模式
- 大数据集

---

### W-TinyLFU (Windowed TinyLFU)

带滑动窗口的窗口化 TinyLFU。

```go
import "github.com/lazygophers/cache/wtinylfu"

cache := wtinylfu.New(1000)
```

**使用场景:**
- 基于时间的访问模式
- 周期性数据访问
- 滑动窗口需求

---

### ALFU (Adaptive LFU)

根据访问模式调整的自适应 LFU 缓存。

```go
import "github.com/lazygophers/cache/alfu"

cache := alfu.New(1000)
```

**使用场景:**
- 未知访问模式
- 自适应需求
- 学习环境

---

### ARC (Adaptive Replacement Cache)

在 LRU 和 LFU 之间自适应调整的 ARC 缓存。

```go
import "github.com/lazygophers/utils/cache/arc"

cache := arc.New(1000)
```

**使用场景:**
- 混合访问模式
- 自适应需求
- 平衡性能

---

### FBR (Frequency-Based Replacement)

基于访问频率淘汰的 FBR 缓存。

```go
import "github.com/lazygophers/cache/fbr"

cache := fbr.New(1000)
```

**使用场景:**
- 基于频率的访问
- 热数据保留
- 冷数据淘汰

---

### SLRU (Segmented LRU)

具有多个分段的分段 LRU 缓存。

```go
import "github.com/lazygophers/utils/cache/slru"

cache := slru.New(1000)
```

**使用场景:**
- 减少锁竞争
- 高并发环境
- 大缓存大小

---

### Optimal

用于理论性能的最优缓存。

```go
import "github.com/lazygophers/cache/optimal"

cache := optimal.New(1000)
```

**使用场景:**
- 最大命中率
- 可预测的访问模式
- 离线分析

---

## 使用模式

### 基本缓存使用

```go
import "github.com/lazygophers/utils/cache/lru"

cache := lru.New(1000)

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

### 缓存选择

```go
// LRU 用于通用场景
cache := lru.New(1000)

// LFU 用于不常访问的数据
cache := lfu.New(1000)

// TinyLFU 用于高性能
cache := tinylfu.New(1000)

// SLRU 用于高并发
cache := slru.New(1000)
```

### 缓存指标

```go
// 获取缓存统计信息
stats := cache.Stats()

fmt.Printf("Size: %d\n", stats.Size)
fmt.Printf("Hits: %d\n", stats.Hits)
fmt.Printf("Misses: %d\n", stats.Misses)
fmt.Printf("Hit Rate: %.2f%%\n", stats.HitRate())
```

---

## 性能比较

| 缓存类型 | 命中率 | 内存 | 最适合 |
|-----------|---------|--------|-----------|
| LRU | 85% | 低 | 通用 |
| LFU | 75% | 低 | 不常访问 |
| LRU-K | 88% | 中 | 混合模式 |
| MRU | 80% | 低 | 时间局部性 |
| TinyLFU | 92% | 中 | 高性能 |
| W-TinyLFU | 90% | 中 | 基于时间 |
| ALFU | 82% | 中 | 自适应 |
| ARC | 86% | 中 | 自适应 |
| FBR | 78% | 中 | 基于频率 |
| SLRU | 90% | 高 | 高并发 |
| Optimal | 95% | 高 | 可预测 |

---

## 最佳实践

### 缓存选择

```go
// 好的做法: 根据访问模式选择
if isSequentialAccess() {
    cache := mru.New(1000)  // MRU 用于顺序访问
} else if isHighConcurrency() {
    cache := slru.New(1000)  // SLRU 用于高并发
} else {
    cache := lru.New(1000)  // LRU 用于通用
}
```

### 缓存大小

```go
// 好的做法: 根据内存约束调整大小
cacheSize := calculateCacheSize()
cache := lru.New(cacheSize)

// 好的做法: 监控命中率
stats := cache.Stats()
if stats.HitRate() < 0.5 {
    // 增加缓存大小
}
```

---

## 相关文档

- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
