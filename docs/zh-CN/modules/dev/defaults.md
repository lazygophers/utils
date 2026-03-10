---
title: defaults
---

# defaults

`defaults` 用于**给结构体或对象补默认值**。它适合把“没有显式传值时该怎么办”从业务分支里抽离出来。

## 适合什么场景

- 配置对象有稳定的默认值策略。
- 请求或内部对象在进入主流程前需要补齐默认值。
- 想让默认值规则和验证规则分开管理。

## 常见入口

- `defaults.SetDefaults`

## 使用建议

- 默认值是“缺省策略”，不是“容错万能药”；关键字段仍应继续做 [validator](/modules/core/validator)。
- 默认值与配置加载通常一起出现，但职责应分开：先加载，再补默认，再校验。
- 对嵌套结构或零值语义不清晰的字段，要先统一团队约定。

## 相关文档

- [validator](/modules/core/validator)
- [config](/modules/system/config)
- [fake](/modules/dev/fake)
