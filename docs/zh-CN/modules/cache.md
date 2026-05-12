---
title: cache - 缓存实现
---

# cache - 缓存实现

## 概览

cache 模块提供多种缓存实现，具有不同的淘汰策略，适用于各种使用场景。

## 可用实现

### LRU (Least Recently Used - 最近最少使用)

当容量达到上限时，淘汰最近最少使用的项。

```go
import "github.com/lazygophers/utils/cache/lru"

cache := lru.New(1000)
```

**使用场景：**
- 通用缓存
- 频繁访问的数据
- 可预测的访问模式

---

### LFU (Least Frequently Used - 最少使用频率)

淘汰使用频率最低的项。

```go
import "github.com/lazygophers/utils/cache/lfu"

cache := lfu.New(1000)
```

**使用场景：**
- 大型数据集
- 不常访问的数据
- 内存受限环境

---

### LRU-K (Least Recently Used with K - 带K值的LRU)

跟踪访问频率的 LRU-K 缓存。

```go
import "github.com/lazygophers/utils/cache/lruk"

cache := lruk.New(1000)
```

**使用场景：**
- 平衡近期性和频率
- 混合访问模式

---

### MRU (Most Recently Used - 最近最多使用)

淘汰最近最多使用的项。

```go
import "github.com/lazygophers/utils/cache/mru"

cache := mru.New(1000)
```

**使用场景：**
- 时间局部性
- 顺序访问模式
- 缓存预热

---

### TinyLFU

结合 LRU 和 LFU 的高性能缓存。

```go
import "github.com/lazygophers/cache/tinylfu"

cache := tinylfu.New(1000)
```

**使用场景：**
- 高性能要求
- 混合访问模式
- 大型数据集

---

### W-TinyLFU (Windowed TinyLFU - 窗口 TinyLFU)

带滑动窗口的 TinyLFU。

```go
import "github.com/lazygophers/cache/wtinylfu"

cache := wtinylfu.New(1000)
```

**使用场景：**
- 突发流量
- 时间局部性变化
- 需要快速适应的缓存

---

### ARC (Adaptive Replacement Cache - 自适应替换缓存)

自动在 LRU 和 LFU 之间调整。

```go
import "github.com/lazygophers/utils/cache/arc"

cache := arc.New(1000)
```

**使用场景：**
- 访问模式变化
- 自适应性能
- 通用缓存

---

### ALFU (Adaptive LFU - 自适应 LFU)

具有自适应性的 LFU 缓存。

```go
import "github.com/lazygophers/cache/alfu"

cache := alfu.New(1000)
```

**使用场景：**
- 动态访问模式
- 需要频率自适应
- 复杂访问模式

---

### FBR (Frequency-Based Replacement - 基于频率的替换)

基于访问频率的替换策略。

```go
import "github.com/lazygophers/cache/fbr"

cache := fbr.New(1000)
```

**使用场景：**
- 频率驱动的淘汰
- 需要频率跟踪
- 大型数据集

---

### Optimal (最优 - 理论基线)

理论最优缓存策略（用于对比）。

```go
import "github.com/lazygophers/cache/optimal"

cache := optimal.New(1000)
```

**使用场景：**
- 性能测试
- 算法研究
- 理论对比

## 选择建议

参见 [缓存策略](./cache/index.md) 了解如何选择合适的缓存实现。
