---
title: cryptox
---

# cryptox

`cryptox` 关注的是**通用加密与编码辅助**。如果你需要处理摘要、对称加密、签名辅助或安全相关的基础操作，可以先从这里判断是否匹配。

## 适合什么场景

- 项目需要统一封装部分加密/解密流程。
- 想减少直接操作底层安全原语时的样板代码。
- 需要与 URL、配置或网络请求中的安全字段处理配合使用。

## 使用建议

- 安全相关代码首先看算法、密钥管理和使用边界，其次才是便利性。
- 任何安全封装都不应绕过审计、轮换和最小暴露原则。
- 如果你的场景是 OpenPGP 密钥与消息处理，请直接看 [pgp](/modules/network/pgp)。

## 相关文档

- [pgp](/modules/network/pgp)
- [urlx](/modules/network/urlx)
- [network](/modules/network/network)
