---
title: orm
---

# orm

`utils.Scan` 与 `utils.Value` 解决的是**结构体字段和数据库 JSON / 文本字段之间的转换**问题，适合配合 `sql.Scanner`、`driver.Valuer` 使用。

## 适合什么场景

- 表字段里存的是 JSON，但业务里希望直接使用结构体或切片。
- 想把数据库读写时的序列化 / 反序列化逻辑收口到统一辅助函数。
- 需要在数据库边界层减少重复的 `json.Unmarshal` / `json.Marshal` 样板。

## 常用入口

- `utils.Scan(src, dst)`：把数据库读出的值扫描到目标对象。
- `utils.Value(v)`：把结构体或集合转换为数据库可写入值。

## 使用建议

- 让 `Scan` / `Value` 停留在数据访问层最合适，不要把数据库边界细节散到业务层。
- 对外部不可信数据仍要做好字段校验；能完成反序列化，不代表业务语义一定合法。
- 如果字段结构会频繁演进，建议搭配版本字段或兼容逻辑一起设计。

## 相关文档

- [must](/modules/core/must)
- [validator](/modules/core/validator)
- [config](/modules/system/config)
