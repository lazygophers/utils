---
title: routine
---

# routine

`routine` 关注的是 **goroutine 组织与运行期辅助**。当你需要更清晰地管理并发任务，而不是零散地直接起协程时，可以先看这里。

## 适合什么场景

- 想统一启动、跟踪或封装 goroutine。
- 需要把并发执行逻辑和业务代码解耦。
- 想在异常恢复、日志或任务组织上保持一致写法。

## 使用建议

- 并发问题的核心仍是生命周期、取消、错误传播和共享状态控制。
- 如果你的问题本质是等待条件满足，而不是起协程，请看 [wait](/modules/concurrency/wait)。
- 如果并发路径需要失败熔断，可继续看 [hystrix](/modules/concurrency/hystrix)。

## 相关文档

- [wait](/modules/concurrency/wait)
- [hystrix](/modules/concurrency/hystrix)
- [event](/modules/concurrency/event)
