---
title: LFU
---

# LFU

Least Frequently Used。淘汰访问频次最低的条目。

## 工作原理

维护一个哈希表和一组按频次分层的双向链表。每个条目记录访问频次，`Get` 命中时频次 +1，条目从旧频次层移到新频次层。淘汰时从最低频次层的尾部移除。

```
频次层：
  freq=1:  [X] ←→ [Y]       ← 淘汰从这取
  freq=2:  [A] ←→ [B]
  freq=3:  [M]
```

同频次内按 LRU 顺序排列，频次相同时优先淘汰最久未访问的。

**时间复杂度**：Get、Put、Remove 均为 O(1)。

**线程安全**：使用 `sync.RWMutex`（`Contains`/`Peek`/`Len` 只需读锁）。

## 构造

```go
import "github.com/lazygophers/utils/cache/lfu"

cache, err := lfu.New[string, int](1000)

// 带淘汰回调
cache, err := lfu.NewWithEvict[string, int](1000, func(key string, value int) {
    log.Printf("evicted: %s = %d", key, value)
})
```

## 算法特有 API

```go
// 查看某个 key 的访问频次（不更新状态）
freq := cache.GetFreq("user:123")

// 统计信息
stats := cache.Stats()
// Stats struct:
//   Size             int         // 当前条目数
//   Capacity         int         // 容量上限
//   MinFreq          int         // 当前最低频次
//   FreqDistribution map[int]int // 频次 → 该频次的条目数
```

## 什么时候用

- 访问模式稳定，存在明确的热点数据
- 频次分布有长尾（少数 key 被大量访问）

## 什么时候别用

- 访问模式会随时间变化（曾经的热点不再访问却占着位置，考虑 [ALFU](/modules/cache/alfu)）
- 负载中有大量一次性扫描（新条目频次为 1，会立即被淘汰）

共享接口语义和选型建议见 [缓存策略总览](/modules/cache/)。
