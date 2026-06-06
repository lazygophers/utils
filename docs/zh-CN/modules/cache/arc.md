---
title: ARC
---

# ARC

Adaptive Replacement Cache。在近期性（LRU）和频次（LFU）之间自适应切换，不需要手动调参。

## 工作原理

ARC 维护四个列表：

```
T1（近期）    T2（频次）
 最近访问的    被多次访问的
    ↑ 新条目进这    ↑ 命中后移到这

B1（幽灵-T1）  B2（幽灵-T2）
 T1 淘汰记录    T2 淘汰记录
```

- **T1**：存放只被访问过一次的条目（近期性）
- **T2**：存放被多次访问的条目（频次）
- **B1/B2**：幽灵列表，不存数据，只记录被淘汰条目的 key

自适应机制：当 B1 中的幽灵条目被再次访问（`Put`），说明 T1 空间不够，增大目标参数 `p`（倾向 T1）。B2 命中则相反（倾向 T2）。`p` 值决定了淘汰时优先从 T1 还是 T2 取。

```
Put 命中 B1 幽灵 → p 增大 → 更多空间给 T1（近期）
Put 命中 B2 幽灵 → p 减小 → 更多空间给 T2（频次）
```

**时间复杂度**：Get、Put、Remove 均为 O(1)。

**线程安全**：使用 `sync.RWMutex`。

## 构造

```go
import "github.com/lazygophers/utils/cache/arc"

cache, err := arc.New[string, int](1000)

// 带淘汰回调
cache, err := arc.NewWithEvict[string, int](1000, func(key string, value int) {
    log.Printf("evicted: %s = %d", key, value)
})
```

## 算法特有 API

```go
stats := cache.Stats()
// Stats struct:
//   Size     int  // 实际缓存条目数（T1 + T2）
//   Capacity int  // 容量上限
//   T1Size   int  // T1（近期）条目数
//   T2Size   int  // T2（频次）条目数
//   B1Size   int  // B1 幽灵条目数
//   B2Size   int  // B2 幽灵条目数
//   P        int  // 自适应参数（T1 目标大小）
```

`P` 值的变化趋势可以帮你理解当前负载是偏近期性还是偏频次。

## 什么时候用

- 负载模式会波动（有时偏近期，有时偏频次）
- 不想手动选 LRU 还是 LFU
- 扫描型流量和热点流量共存

## 什么时候别用

- 需要简单可预测的策略（ARC 内部状态较复杂）
- 内存非常紧张（幽灵列表会额外占用空间，虽然不存 value）

共享接口语义和选型建议见 [缓存策略总览](/modules/cache/)。
