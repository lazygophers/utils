---
    title: cryptox
    ---

    # cryptox

    仓库里的密码学工具集合，覆盖哈希、对称加密、非对称加密、签名、UUID/ULID 等能力。

    ## 适用场景

    - 需要快速生成标识符。
- 需要常见 hash / HMAC / ECDSA / ECDH / DES 等能力。
- 需要用一致的工具包管理安全相关辅助函数。

    ## 你会接触到什么

    - 标识符：`UUID`、`ULID`、`ULIDWithTimestamp`。
- 哈希：`Hash32`、`Hash64` 等 FNV 辅助。
- 密钥与签名：ECDH / ECDSA 相关生成、签名、验证与 PEM 转换能力。

    ## 使用建议

    - 密码学能力面很广，进入本包后应按算法族继续细读，而不是把它当作一个单函数包。
- 如果是面向业务安全边界，建议再封装一层自己的领域语义。

    ## 相关文档

    - [模块总览](/modules/overview)
    - [API 概览](/api/overview)
