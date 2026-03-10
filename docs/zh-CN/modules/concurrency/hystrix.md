---
title: hystrix
---

# hystrix

`hystrix` 解决的是 **熔断、隔离与故障保护**。当下游服务不稳定、失败会放大连锁影响时，它比普通重试更值得优先考虑。

## 适合什么场景

- 外部依赖不稳定，需要避免故障扩散。
- 想给调用链加上熔断、降级或隔离保护。
- 需要把“失败如何被限制”明确成基础设施规则。

## 使用建议

- 熔断不是简单的“失败就重试”，而是要结合阈值、窗口、恢复策略一起设计。
- 接入前先明确降级路径；没有降级目标的熔断价值会大幅下降。
- 如果你首先缺的是任务组织能力，可先看 [routine](/modules/concurrency/routine)。

## 相关文档

- [routine](/modules/concurrency/routine)
- [wait](/modules/concurrency/wait)
- [singledo](/modules/concurrency/singledo)
