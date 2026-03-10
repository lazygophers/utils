---
title: anyx
---

# anyx

`anyx` 面向的是**弱类型值访问与动态数据读取**。当你拿到的是 `map[string]any`、JSON 解码后的动态结构，或者不方便先定义完整结构体时，可以先用它做过渡处理。

## 适合什么场景

- 从动态映射里安全地读取字段。
- 处理 JSON / YAML 解析后的弱类型数据。
- 想在“完全动态”和“强类型结构体”之间保留一层过渡。

## 你会在这里关注什么

- 动态 Map 封装。
- 类型读取与转换辅助。
- 与 JSON、配置、校验之间的衔接方式。

## 使用建议

- `anyx` 更像边界层工具；核心业务模型仍建议尽快落回结构体。
- 如果数据结构已经稳定，优先直接绑定到结构体，再配合 [validator](/modules/core/validator)。
- 需要统一 JSON 入口时，可结合 [json](/modules/data/json) 一起使用。

## 相关文档

- [candy](/modules/data/candy)
- [json](/modules/data/json)
- [validator](/modules/core/validator)
