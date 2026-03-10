---
title: pgp
---

# pgp

`pgp` 聚焦 **OpenPGP** 相关能力，适合处理密钥、加密文本、签名消息等明确属于 PGP 语义的场景。

## 适合什么场景

- 需要生成、读取或使用 PGP 密钥。
- 想对文本、文件或消息做 PGP 加密/解密。
- 需要在系统中接入一条与通用对称加密不同的安全链路。

## 使用建议

- 先区分你的需求是不是“必须是 PGP”；如果不是，通用加密方案可能更简单。
- 密钥存储、权限控制和轮换策略比示例代码本身更关键。
- 与 `cryptox` 的关系是“主题不同、边界不同”，不建议混成一个统一黑盒。

## 相关文档

- [cryptox](/modules/network/cryptox)
- [network](/modules/network/network)
- [API 概览](/api/overview)
