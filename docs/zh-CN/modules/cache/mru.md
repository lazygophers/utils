---
title: MRU
---

# MRU

Most Recently Used。淘汰**最近刚访问过**的条目，和 LRU 相反。

## 工作原理

数据结构与 LRU 相同：双向链表 + 哈希表。区别在于容量满时淘汰链表**头部**（最近访问的）而非尾部。

```
链表头部（最近访问，优先淘汰）      链表尾部（最久未访问）
  [A] ←→ [B] ←→ [C] ←→ [D] ←→ [E]
   ↑ 满时从这淘汰                    ↑ 留着
```

这听起来反直觉，但在特定模式下是正确的：如果刚访问的数据短期内不会再被访问（比如顺序扫描），那最近访问的反而是最不可能再用的。

**时间复杂度**：Get、Put、Remove 均为 O(1)。

**线程安全**：使用 `sync.Mutex`（`Get` 需要移动节点）。

## 构造

```go
import "github.com/lazygophers/utils/cache/mru"

cache, err := mru.New[string, int](1000)
cache, err := mru.NewWithEvict[string, int](1000, func(key string, value int) {
    log.Printf("evicted: %s = %d", key, value)
})
```

## 算法特有 API

```go
stats := cache.Stats()
// Stats struct:
//   Size     int  // 当前条目数
//   Capacity int  // 容量上限
```

`Keys()` 返回顺序：从最近访问到最久访问。

## 什么时候用

- 顺序扫描型负载（遍历完不会回头）
- 已知最近访问的数据短期内不会再用

## 什么时候别用

- 典型的热点缓存（最近访问的就是最可能再用的）
- Web 请求缓存、数据库查询缓存

共享接口语义和选型建议见 [缓存策略总览](/modules/cache/)。
