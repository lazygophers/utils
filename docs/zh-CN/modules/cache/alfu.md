---
title: ALFU
---

# ALFU

Adaptive LFU。在 LFU 基础上增加频率衰减（Decay）机制，解决"历史热点占着不走"的问题。

## 工作原理

数据结构与 LFU 类似：哈希表 + 按频次分层的链表。关键区别是引入了**衰减机制**：

```
每隔 decayInterval：
  所有条目的频次 × decayFactor（默认 0.9）
  频次低于 1 的重置为 1
```

这意味着长期不被访问的条目频次会逐渐降低，最终被淘汰。新热点可以更快地超过旧热点。

```
初始：条目 A freq=100（历史热点）
衰减后：条目 A freq=90 → 81 → 72 → ...
新条目 B freq=1 → 2 → 4 → 8 → ...
一段时间后 B 的频次可能超过 A
```

衰减检查在 `Get`/`Put` 时自动触发，无需外部定时器。

**时间复杂度**：Get、Put 均为 O(1)（衰减时 O(n)，但分摊到每次访问中）。

**线程安全**：使用 `sync.RWMutex`。

## 构造

```go
import "github.com/lazygophers/utils/cache/alfu"

// 默认配置：decayFactor=0.9, decayInterval=5 分钟
cache, err := alfu.New[string, int](1000)

// 自定义衰减参数
cache, err := alfu.NewWithConfig[string, int](1000, 0.8, 3*time.Minute)

// 带淘汰回调
cache, err := alfu.NewWithEvict[string, int](1000, func(key string, value int) {
    log.Printf("evicted: %s = %d", key, value)
})
```

`decayFactor` 范围 (0, 1]，值越小衰减越快。`decayInterval` 越短衰减越频繁。

## 算法特有 API

```go
// 强制立即执行一次衰减
cache.ForceDecay()

// 统计信息
stats := cache.Stats()
// Stats struct:
//   Size                  int           // 当前条目数
//   Capacity              int           // 容量上限
//   MinFrequency          int           // 当前最低频次
//   MaxFrequency          int           // 当前最高频次
//   DecayFactor           float64       // 衰减因子
//   DecayInterval         time.Duration // 衰减间隔
//   LastDecay             time.Time     // 上次衰减时间
//   FrequencyDistribution map[int]int   // 频次 → 条目数
```

`Keys()` 和 `Values()` 按频次从高到低返回。

## 什么时候用

- 访问模式不稳定，热点会随时间迁移
- 需要 LFU 的频次语义，但又不想旧数据永久占据高位

## 什么时候别用

- 访问模式非常稳定（直接用 [LFU](/modules/cache/lfu) 更简单）
- 对衰减的额外 CPU 开销敏感（每 `decayInterval` 一次全量扫描）

共享接口语义和选型建议见 [缓存策略总览](/modules/cache/)。
