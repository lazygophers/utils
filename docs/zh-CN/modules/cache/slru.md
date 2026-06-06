---
title: SLRU
---

# SLRU

Segmented LRU。将缓存分为两段：试用段（Probationary）和保护段（Protected），隔离一次性流量对热点数据的冲击。

## 工作原理

```
新条目 → [试用段 (20%)] ──二次访问──→ [保护段 (80%)]
              ↑ 淘汰优先从这里                ↑ 满时降级回试用段
```

- 新条目进入试用段
- 试用段中的条目被再次访问（`Get` 命中），提升到保护段
- 保护段满时，最久未访问的条目降级回试用段（不是直接淘汰）
- 试用段满时，最久未访问的条目才被淘汰

默认比例 20% 试用 / 80% 保护，可通过 `NewWithRatio` 自定义。

**时间复杂度**：Get、Put、Remove 均为 O(1)。

**线程安全**：使用 `sync.RWMutex`。

## 构造

```go
import "github.com/lazygophers/utils/cache/slru"

// 默认 20%/80% 比例
cache, err := slru.New[string, int](1000)

// 自定义试用段比例（0.0 ~ 1.0）
cache, err := slru.NewWithRatio[string, int](1000, 0.3) // 30% 试用，70% 保护

// 带淘汰回调
cache, err := slru.NewWithEvict[string, int](1000, func(key string, value int) {
    log.Printf("evicted: %s = %d", key, value)
})
```

## 算法特有 API

```go
stats := cache.Stats()
// Stats struct:
//   Size                 int  // 当前总条目数
//   Capacity             int  // 容量上限
//   ProbationarySize     int  // 试用段当前条目数
//   ProtectedSize        int  // 保护段当前条目数
//   ProbationaryCapacity int  // 试用段容量
//   ProtectedCapacity    int  // 保护段容量
```

`Keys()` 返回顺序：保护段在前，试用段在后。

## 什么时候用

- 需要在近期性内部做分层，保护已验证的热点
- 负载中混杂大量一次性请求（爬虫、批量导入）

## 什么时候别用

- 只需要最简单的通用缓存（直接用 [LRU](/modules/cache/lru)）
- 需要自适应调节（考虑 [ARC](/modules/cache/arc)）

共享接口语义和选型建议见 [缓存策略总览](/modules/cache/)。
