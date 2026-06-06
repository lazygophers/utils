---
title: LRU-K
---

# LRU-K

LRU-K 是 LRU 的改进版。普通 LRU 只看"最近一次访问"，LRU-K 要求条目被访问 K 次后才进入主缓存，过滤掉偶发的一次性访问。

## 工作原理

维护两个列表：**历史列表**（History）和**主缓存列表**（Main Cache）。

```
新条目 → [历史列表] ──第 K 次访问──→ [主缓存列表]
              ↑ 记录访问次数                ↑ 按 LRU 管理
              ↑ 未达 K 次不占主缓存         ↑ 满时从尾部淘汰
```

- 条目首次 `Put` 进入历史列表，开始计数
- 访问次数达到 K 后，提升到主缓存
- 主缓存内按标准 LRU 管理
- `Contains` 和 `Peek` 只对主缓存中的条目返回 true

访问时间用环形缓冲区记录，避免切片扩容开销。

**时间复杂度**：Get、Put、Remove 均为 O(1)。

**线程安全**：使用 `sync.RWMutex`。

## 构造

```go
import "github.com/lazygophers/utils/cache/lruk"

// capacity=1000, K=2（需要被访问 2 次才进入主缓存）
cache, err := lruk.New[string, int](1000, 2)

// K=3：更严格，需要 3 次访问
cache, err := lruk.New[string, int](1000, 3)

// 带淘汰回调
cache, err := lruk.NewWithEvict[string, int](1000, 2, func(key string, value int) {
    log.Printf("evicted: %s = %d", key, value)
})
```

`capacity` 和 `k` 都必须大于 0。

## 算法特有 API

```go
// 查看 K 值
k := cache.GetK() // 返回 2

// 统计信息
stats := cache.Stats()
// Stats struct:
//   Size         int  // 主缓存条目数（不含历史）
//   Capacity     int  // 主缓存容量上限
//   K            int  // K 值
//   HistorySize  int  // 历史列表条目数
//   TotalEntries int  // 总条目数（主缓存 + 历史）
```

注意：`Len()` 只返回主缓存的条目数，不包括历史列表中的条目。

## 什么时候用

- 需要过滤偶发访问（只有反复访问的才值得缓存）
- "持续被访问"比"一次最近访问"更重要的场景

## 什么时候别用

- 负载规模很小，普通 LRU 已足够
- K 值设置过高会导致新热点迟迟无法进入缓存

共享接口语义和选型建议见 [缓存策略总览](/modules/cache/)。
