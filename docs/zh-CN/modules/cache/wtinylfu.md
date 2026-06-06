---
title: W-TinyLFU
---

# W-TinyLFU

Window-TinyLFU。在 TinyLFU 基础上增大窗口比例（10%），给新条目更多观察时间后再做准入决策。是混合负载中综合表现最好的策略之一。

## 工作原理

```
新条目 → [Window (10%)] → [Probation (20% of main)] → [Protected (80% of main)]
              ↑ 比TinyLFU窗口更大       ↑ 准入/淘汰竞争区         ↑ 已验证热点
```

三段晋升路径：
1. **Window**：所有新条目先进窗口，按 LRU 管理
2. **Probation**：Window 溢出时，条目与 Probation 尾部竞争（比较频率），高频者留下
3. **Protected**：Probation 中的条目被 `Get` 命中后晋升。Protected 满时降级到 Probation

频率估计用 Count-Min Sketch（4 层，带采样），晋升/准入决策基于 Sketch 估计的频率对比。

与 TinyLFU 的区别：
- Window 占 10%（TinyLFU 是 1%），新条目有更多缓冲
- Protected 占主缓存的 80%（更强的热点保护）
- Sketch 使用采样（大缓存每 10 次记录 1 次），开销更低

**时间复杂度**：Get、Put、Remove 均为 O(1)。

**线程安全**：使用 `sync.RWMutex`。

## 构造

```go
import "github.com/lazygophers/utils/cache/wtinylfu"

cache, err := wtinylfu.New[string, int](1000)

// 带淘汰回调
cache, err := wtinylfu.NewWithEvict[string, int](1000, func(key string, value int) {
    log.Printf("evicted: %s = %d", key, value)
})
```

小容量自动调整：容量为 1 时 Window 占全部；主缓存 ≤ 4 时 Probation 和 Protected 各半。

## 算法特有 API

```go
stats := cache.Stats()
// Stats struct:
//   Size           int  // 当前总条目数
//   Capacity       int  // 容量上限
//   WindowSize     int  // Window 当前条目数
//   WindowCapacity int  // Window 容量
//   ProbationSize  int  // Probation 当前条目数
//   ProbationCap   int  // Probation 容量
//   ProtectedSize  int  // Protected 当前条目数
//   ProtectedCap   int  // Protected 容量
```

## 什么时候用

- 不确定负载模式，需要"最安全的默认选择"
- 混合了热点、扫描、突发流量的场景
- 需要同时保留热点数据和给新热点机会

## 什么时候别用

- 需要简单的策略（W-TinyLFU 内部有三段 + Sketch，复杂度较高）
- 已确认负载纯近期性或纯频次（直接用 [LRU](/modules/cache/lru) 或 [LFU](/modules/cache/lfu) 更轻量）

共享接口语义和选型建议见 [缓存策略总览](/modules/cache/)。
