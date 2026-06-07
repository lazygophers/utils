---
title: Module Overview
---

# Module Overview

LazyGophers Utils provides 20+ specialized modules covering all aspects of Go development. All modules are organized by category for easy navigation.

## Eight Category Groups

| Category | When to start here | Modules |
| --- | --- | --- |
| [Core Utilities](/en/modules/core/) | Reduce boilerplate or handle database field mapping | [must](/en/modules/core/must), [orm](/en/modules/core/orm) |
| [Data Processing](/en/modules/data/) | Type conversion, collection ops, string normalization, JSON | [candy](/en/modules/data/candy), [json](/en/modules/data/json), [stringx](/en/modules/data/stringx), [anyx](/en/modules/data/anyx) |
| [Cache Strategies](/en/modules/cache/) | Choose a cache eviction policy that fits your workload | [Overview](/en/modules/cache/), [LRU](/en/modules/cache/lru), [LFU](/en/modules/cache/lfu), [TinyLFU](/en/modules/cache/tinylfu), [SLRU](/en/modules/cache/slru), [MRU](/en/modules/cache/mru), [ALFU](/en/modules/cache/alfu), [ARC](/en/modules/cache/arc), [LRU-K](/en/modules/cache/lruk), [W-TinyLFU](/en/modules/cache/wtinylfu), [FBR](/en/modules/cache/fbr), [Optimal](/en/modules/cache/optimal) |
| [Time & Scheduling](/en/modules/time/) | Lunar calendar, solar terms, or fixed shift rules | [xtime](/en/modules/time/xtime), [xtime996](/en/modules/time/xtime996), [xtime955](/en/modules/time/xtime955), [xtime007](/en/modules/time/xtime007) |
| [System & Configuration](/en/modules/system/) | Config loading, path resolution, app init, exit cleanup | [config](/en/modules/system/config), [runtime](/en/modules/system/runtime), [osx](/en/modules/system/osx), [app](/en/modules/system/app), [atexit](/en/modules/system/atexit) |
| [Network & Security](/en/modules/network/) | HTTP helpers, encryption, signing, URL normalization | [network](/en/modules/network/network), [cryptox](/en/modules/network/cryptox), [pgp](/en/modules/network/pgp), [urlx](/en/modules/network/urlx) |
| [Concurrency & Control Flow](/en/modules/concurrency/) | Task execution, waiting, circuit breaking, dedup | [routine](/en/modules/concurrency/routine), [wait](/en/modules/concurrency/wait), [hystrix](/en/modules/concurrency/hystrix), [singledo](/en/modules/concurrency/singledo), [event](/en/modules/concurrency/event) |
| [Development & Testing](/en/modules/dev/) | Default values, random/fake data, profiling | [randx](/en/modules/dev/randx), [fake](/en/modules/dev/fake), [defaults](/en/modules/dev/defaults), [pyroscope](/en/modules/dev/pyroscope) |

## Standalone Module

| Module | Description |
| --- | --- |
| [Validator](/en/validator/) | Data validation — 169 built-in validators, 100% coverage of go-playground/validator v10 rules |

## Choosing the Right Module

### Infrastructure vs. Business Helper

- Startup, config, validation, DB field mapping → **Core Utilities** & **System & Configuration**
- Collection handling, string normalization, JSON, random data → **Data Processing** & **Development & Testing**
- Caching, concurrency, retry, scheduling → **Cache Strategies**, **Concurrency & Control Flow**, **Time & Scheduling**

### Modules with Special Considerations

- **Cache**: Each strategy has different eviction logic and thread-safety semantics — read before choosing.
- **xtime**: Beyond time helpers, includes lunar calendar, solar terms, and shift rules.
- **atexit**: Exit behavior varies across platforms.

## Recommended Reading Path

1. New project: `must` → `config` → `validator` → relevant business module.
2. Existing project: browse by category, then read individual module pages.
3. Precise signatures: visit [pkg.go.dev](https://pkg.go.dev/github.com/lazygophers/utils).
