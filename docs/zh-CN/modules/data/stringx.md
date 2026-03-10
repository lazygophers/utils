---
title: stringx
---

# stringx

`stringx` 关注的是**字符串与字节切片处理、命名规整和常见字符串辅助**。当你需要统一 URL、字段名、缓存键或日志片段时，它通常是数据处理链路里的基础工具。

## 适合什么场景

- 在 `string` 与 `[]byte` 之间转换。
- 统一命名风格，例如驼峰与下划线之间转换。
- 做拆分、截断、拼接、反转或其他轻量字符串整理。

## 你会在这里关注什么

- `ToString` / `ToBytes` 这类底层转换。
- 命名格式转换函数。
- 是否涉及 `unsafe`、零拷贝以及由此带来的使用边界。

## 使用建议

- 涉及零拷贝转换时，先确认调用方是否会继续修改底层数据。
- 对外暴露稳定 API 时，优先保证语义清晰，再考虑微观性能差异。
- 如果目标是生成稳定 URL 或缓存键，可继续看 [urlx](/modules/network/urlx)。

## 相关文档

- [candy](/modules/data/candy)
- [json](/modules/data/json)
- [urlx](/modules/network/urlx)
