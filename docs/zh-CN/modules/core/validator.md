---
title: validator
---

# validator

`validator` 包负责**结构体验证**。如果你已经把配置、请求体或持久化对象映射成 Go 结构体，下一步通常就是在这里做规则检查。

## 适合什么场景

- 启动时校验配置是否完整。
- 请求参数已经绑定到结构体，需要统一执行字段规则验证。
- 想把“字段名 + 规则 + 本地化错误信息”集中到一套机制里管理。

## 常用入口

- `validator.Struct(v)`：对结构体执行规则校验。

## 使用建议

- 把验证放在“数据进入系统后的第一站”，不要等到业务深处再补救。
- 如果你文档里看到旧的 `utils.Validate`，请以当前源码为准：这里应使用 `validator.Struct`。
- 本包带有多语言与字段名策略能力，涉及错误文案时先确认项目实际展示语言。

## 相关文档

- [must](/modules/core/must)
- [config](/modules/system/config)
- [API 概览](/api/overview)
