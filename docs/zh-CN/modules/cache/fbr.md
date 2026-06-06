---
title: FBR
---

# FBR

Frequency-Based Replacement。基于频次的分区淘汰策略，比 LFU 简单，没有衰减机制。

## 工作原理

与 LFU 结构类似：哈希表 + 按频次分层的链表。区别在于 FBR 没有衰减机制，频次只增不减。淘汰时始终从最低频次层的尾部移除。

```
频次层：
  freq=1:  [X] ←→ [Y]       ← 淘汰从这取（尾部最旧）
  freq=2:  [A] ←→ [B]
  freq=3:  [M] ←→ [N]
```

`Get` 命中时频次 +1，条目移到更高频次层。同频次内按 LRU 顺序排列。

**时间复杂度**：Get、Put、Remove 均为 O(1)。

**线程安全**：使用 `sync.RWMutex`。

## 构造

```go
import "github.com/lazygophers/utils/cache/fbr"

cache, err := fbr.New[string, int](1000)

// 带淘汰回调
cache, err := fbr.NewWithEvict[string, int](1000, func(key string, value int) {
    log.Printf("evicted: %s = %d", key, value)
})
```

## 算法特有 API

```go
stats := cache.Stats()
// Stats struct:
//   Size                  int         // 当前条目数
//   Capacity              int         // 容量上限
//   MinFrequency          int         // 当前最低频次
//   MaxFrequency          int         // 当前最高频次
//   FrequencyDistribution map[int]int // 频次 → 条目数
```

`Keys()` 和 `Values()` 按频次从高到低返回。

## 什么时候用

- 需要 LFU 的频次语义但不需要衰减
- 访问模式非常稳定，热点不会迁移

## 什么时候别用

- 访问模式会变化（没有衰减，旧热点会永久占据高位，考虑 [ALFU](/modules/cache/alfu)）
- 需要更复杂的准入策略（考虑 [TinyLFU](/modules/cache/tinylfu)）

共享接口语义和选型建议见 [缓存策略总览](/modules/cache/)。
