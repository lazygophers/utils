---
title: urlx
---

# urlx

`urlx` 负责 **URL 与查询参数规范化**。如果你想生成稳定缓存键、签名串或统一 URL 表达，这个包通常比手写字符串拼接更稳。

## 适合什么场景

- 需要稳定地排序查询参数。
- 想把 URL 规范化逻辑收口到统一辅助函数。
- 需要生成便于缓存、比较或签名的 URL 字符串。

## 常用入口

- `urlx.SortQuery`：统一查询参数顺序。

## 使用建议

- URL 规范化前先明确保留哪些语义：参数顺序、编码方式、空值策略都可能影响结果。
- 如果规范化结果会参与签名，请和调用方保持完全一致的编码约定。
- 字符串层面的处理可配合 [stringx](/modules/data/stringx) 一起使用。

## 相关文档

- [stringx](/modules/data/stringx)
- [network](/modules/network/network)
- [cryptox](/modules/network/cryptox)
