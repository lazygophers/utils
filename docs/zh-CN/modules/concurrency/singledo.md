---
title: singledo
---

# singledo

`singledo` 面向的是 **同类请求去重**。如果多个并发调用本质上都在做同一件事，这类工具可以让系统只做一次真正工作。

## 适合什么场景

- 避免相同 key 的重复加载、重复刷新或重复初始化。
- 减少缓存击穿时对下游的并发放大。
- 想让并发请求共享同一次计算结果。

## 使用建议

- 去重 key 的设计非常关键，过粗会误共享，过细则失去价值。
- `singledo` 解决的是“同一次工作不要重复做”，不负责缓存长期结果。
- 如果你要解决的是下游故障扩散而不是重复调用，请看 [hystrix](/modules/concurrency/hystrix)。

## 相关文档

- [routine](/modules/concurrency/routine)
- [wait](/modules/concurrency/wait)
- [缓存策略](/modules/cache/)
