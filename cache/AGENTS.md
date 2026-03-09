# Cache Package - Algorithm Selection Guide

**Package:** `github.com/lazygophers/utils/cache`

## OVERVIEW
11 cache algorithms for different workloads - LRU, LFU, ARC, TinyLFU, etc.

## STRUCTURE
```
cache/
├── lru/        # Least Recently Used
├── lfu/        # Least Frequently Used
├── arc/        # Adaptive Replacement Cache
├── fbr/        # Frequency-Based Replacement
├── alfu/       # Adaptive LFU
├── lruk/       # LRU-K (K-references)
├── slru/       # Segmented LRU
├── mru/        # Most Recently Used
├── optimal/    # Belady's optimal algorithm
├── tinylfu/    # TinyLFU (W-TinyLFU variant)
└── wtinylfu/   # Window-TinyLFU
```

## WHERE TO LOOK

| Use Case | Algorithm | Package | Notes |
|----------|-----------|---------|-------|
| **General purpose** | TinyLFU/W-TinyLFU | `tinylfu/`, `wtinylfu/` | Best hit rate for mixed workloads |
| **Scan resistance** | ARC | `arc/` | Adaptive, resists scan-heavy workloads |
| **Frequency bias** | LFU | `lfu/` | Hot data stays cached |
| **Recency bias** | LRU | `lru/` | Standard LRU, simple and fast |
| **Optimal benchmark** | Optimal | `optimal/` | Belady's algorithm (theoretical best) |
| **Write-heavy** | MRU | `mru/` | Evicts most recent, good for loops |

## CONVENTIONS

### Interface Pattern
All caches implement `CacheInterface`:
```go
type CacheInterface[K comparable, V any] interface {
    Get(key K) (V, bool)
    Set(key K, value V)
    Has(key K) bool
    Del(key K)
    Purge()
    Keys() []K
    Len() int
}
```

### Thread Safety
- **Non-thread-safe by default**: Use external synchronization
- **Exception**: Some implementations have `*Sync` variants

## ANTI-PATTERNS

### DO NOT
- **DO NOT** use Optimal cache in production (theoretical only, requires future knowledge)
- **DO NOT** use LRU for scan-heavy workloads (pollutes cache with single-use items)
- **DO NOT** ignore eviction policies (each algorithm has specific trade-offs)

### ALWAYS
- **ALWAYS** benchmark with your actual workload before choosing
- **ALWAYS** check memory limits (some algorithms track more metadata)
- **ALWAYS** consider read/write ratio when selecting algorithm

## PERFORMANCE

| Algorithm | Get | Set | Memory Overhead |
|-----------|-----|-----|-----------------|
| LRU | O(1) | O(1) | Low |
| LFU | O(1) | O(log n) | Medium |
| ARC | O(1) | O(1) | High (2x) |
| TinyLFU | O(1) | O(1) | Medium |

## SELECTION MATRIX

See [SELECTION_STRATEGY.md](SELECTION_STRATEGY.md) and [SELECTION_STRATEGY_ZH.md](SELECTION_STRATEGY_ZH.md) for detailed decision trees.
