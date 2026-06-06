---
title: cache - 缓存实现
---

# cache - 缓存实现

cache 不是单一实现，而是按淘汰策略拆分的独立子包。选型建议和共享接口语义见 [缓存策略总览](/modules/cache/)。

## 可用实现

| 策略 | 导入路径 | 核心信号 |
| --- | --- | --- |
| LRU | `github.com/lazygophers/utils/cache/lru` | 最近访问时间 |
| LFU | `github.com/lazygophers/utils/cache/lfu` | 访问频次 |
| LRU-K | `github.com/lazygophers/utils/cache/lruk` | 第 K 次访问 |
| MRU | `github.com/lazygophers/utils/cache/mru` | 反向淘汰最近访问 |
| TinyLFU | `github.com/lazygophers/utils/cache/tinylfu` | 近期性 + 频次 |
| W-TinyLFU | `github.com/lazygophers/utils/cache/wtinylfu` | 窗口 + TinyLFU |
| ARC | `github.com/lazygophers/utils/cache/arc` | 自适应近期性/频次 |
| ALFU | `github.com/lazygophers/utils/cache/alfu` | 自适应频次 |
| SLRU | `github.com/lazygophers/utils/cache/slru` | 分段近期性 |
| FBR | `github.com/lazygophers/utils/cache/fbr` | 频次分区替换 |
| Optimal | `github.com/lazygophers/utils/cache/optimal` | 理论基线 |

```go
import "github.com/lazygophers/utils/cache/lru"

cache := lru.New(1000)
```

选型指引见 [缓存策略](/modules/cache/)。
