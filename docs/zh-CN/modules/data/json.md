---
title: json
---

# json

`json` 包是围绕标准库 JSON 能力的工程化封装，目标不是发明新格式，而是让编码、解码和项目内的一致性更容易维护。

## 适合什么场景

- 想在项目里统一 JSON 编解码入口。
- 需要和仓库里其他工具一起使用，例如数据库字段映射、配置解析或弱类型处理。
- 希望在代码里明确区分“标准库原语”和“项目内统一封装”。

## 常见入口

- `json.Marshal`
- `json.Unmarshal`

## 使用建议

- 如果你已经在项目里约定统一 JSON 包，尽量保持全链路一致，避免混用多个实现。
- 这个包解决的是“编码/解码流程管理”，不是字段合法性校验；规则验证仍应交给 [validator](/modules/core/validator)。
- 文档不再保留脱离实际负载上下文的固定性能结论。

## 相关文档

- [candy](/modules/data/candy)
- [anyx](/modules/data/anyx)
- [orm](/modules/core/orm)
