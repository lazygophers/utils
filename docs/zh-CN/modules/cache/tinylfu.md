---
title: TinyLFU
---

# TinyLFU

TinyLFU 用 Count-Min Sketch 做频率估计，将近期性与频次结合，在准入阶段就决定新条目是否值得进入主缓存。

## 工作原理

缓存分为三层：

```
新条目 → [Window (1%)] ──准入评估──→ [Probation] ──再次访问──→ [Protected]
              ↑ 新条目先在这             ↑ 被准入的条目      ↑ 热点数据
              ↑ 满时评估是否值得          ↑ 可被更热门的挤出   ↑ 满时降级到 Probation
                 进入主缓存
```

准入策略：Window 满时，条目需要和 Probation 尾部（最可能被淘汰的条目）比较频率。只有频率更高才能进入主缓存，否则直接丢弃。

频率估计用 Count-Min Sketch（4 个哈希函数，每个计数器上限 15），配合 Doorkeeper（布隆过滤器替代）过滤首次访问。每隔 `capacity × 10` 次 Put 后，Sketch 计数器减半（老化），避免历史频次永久占优。

**时间复杂度**：Get、Put、Remove 均为 O(1)。

**线程安全**：使用 `sync.RWMutex`。

## 构造

```go
import "github.com/lazygophers/utils/cache/tinylfu"

cache, err := tinylfu.New[string, int](1000)

// 带淘汰回调
cache, err := tinylfu.NewWithEvict[string, int](1000, func(key string, value int) {
    log.Printf("evicted: %s = %d", key, value)
})
```

## 算法特有 API

```go
stats := cache.Stats()
// Stats struct:
//   Size           int  // 当前总条目数
//   Capacity       int  // 容量上限
//   WindowSize     int  // Window 当前条目数
//   ProbationSize  int  // Probation 当前条目数
//   ProtectedSize  int  // Protected 当前条目数
//   WindowCapacity int  // Window 容量
//   MainCapacity   int  // 主缓存容量（Probation + Protected）
//   SketchSize     int  // Sketch 累计采样数
//   DoorkeeperSize int  // Doorkeeper 条目数
```

## 什么时候用

- 混合负载（无明显单一模式），想用稳妥的默认选择
- 读多写少，需要频次信息但不想承受精确 LFU 的内存开销

## 什么时候别用

- 必须完全可控每一步淘汰逻辑（TinyLFU 内部状态较多）
- 缓存容量非常小（Window 1% 只有 1 个条目，准入机制几乎不生效）

共享接口语义和选型建议见 [缓存策略总览](/modules/cache/)。
