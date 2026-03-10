---
title: wait
---

# wait

`wait` 面向的是 **条件等待、轮询等待与时机控制**。当你需要“等某件事发生”而不是立即执行下一步时，这类工具更合适。

## 适合什么场景

- 等待服务启动、状态就绪、文件生成或条件满足。
- 想把轮询等待、超时控制和重试节奏收口。
- 需要让测试、初始化或后台流程的等待逻辑更可读。

## 使用建议

- 等待工具要明确超时与退出条件，避免形成隐式死等。
- 如果能通过事件通知替代轮询，优先考虑事件驱动方案。
- 需要把等待放进更大的并发组织里时，可结合 [routine](/modules/concurrency/routine)。

## 相关文档

- [routine](/modules/concurrency/routine)
- [hystrix](/modules/concurrency/hystrix)
- [event](/modules/concurrency/event)
