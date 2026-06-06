---
title: LRU
---

# LRU

Least Recently Used。淘汰最近最少使用的条目。

## 工作原理

内部维护一个双向链表（`container/list`）加一个哈希表。每次 `Get` 命中将条目移到链表头部，`Put` 新条目也插入头部。容量满时淘汰链表尾部（最久未访问）。

```
链表头部（最近访问）          链表尾部（最久未访问，优先淘汰）
  [A] ←→ [B] ←→ [C] ←→ [D] ←→ [E]
   ↑ Put/Get 命中会移到这              ↑ 满时从这淘汰
```

**时间复杂度**：Get、Put、Remove 均为 O(1)。

**线程安全**：使用 `sync.Mutex`（不是 RWMutex，因为 `Get` 需要移动节点）。基准测试显示 Mutex 比 RWMutex 快约 25%。

## 构造

```go
import "github.com/lazygophers/utils/cache/lru"

// 基本构造
cache, err := lru.New[string, int](1000)

// 带淘汰回调
cache, err := lru.NewWithEvict[string, int](1000, func(key string, value int) {
    log.Printf("evicted: %s = %d", key, value)
})
```

`capacity` 必须大于 0，否则返回错误。

## 算法特有 API

```go
// 统计信息
stats := cache.Stats()
// Stats struct:
//   Size     int  // 当前条目数
//   Capacity int  // 容量上限
```

`Keys()` 返回顺序：从最近访问到最久访问。

## 什么时候用

- 通用场景的默认选择
- 访问模式明显呈时间局部性（刚访问的很可能再访问）

## 什么时候别用

- 大规模顺序扫描（一次扫描会把所有老条目挤掉）
- 已知频次比近期性更重要（考虑 LFU / ALFU）

共享接口语义和选型建议见 [缓存策略总览](/modules/cache/)。
