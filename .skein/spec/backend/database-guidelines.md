# Database Helpers

> ORM 辅助工具（非 ORM 框架）。本项目是工具库，不包含数据库层。

---

## Overview

根包 `utils` 提供 `sql.Scanner` / `driver.Valuer` 辅助函数，方便调用方在 SQL 数据库中存储/读取 JSON 序列化的 Go struct。

## Core Functions

```go
// Value — 将任意 Go 值序列化为 JSON driver.Value
data, err := utils.Value(user)

// Scan — 从 SQL 结果扫描 JSON driver.Value 并反序列化到 Go 值
err := utils.Scan(data, &user)
```

## Usage Pattern

```go
// 数据库模型中嵌入 JSON 存储
type User struct {
    ID   int64
    Data string // JSON-encoded
}

// 写入
value, err := utils.Value(user)
db.Exec("INSERT INTO users (data) VALUES (?)", value)

// 读取
var data string
row.Scan(&data)
var user User
utils.Scan(data, &user)
```

## Round-trip 保证

`Value()` → `Scan()` 往返保证类型完整。测试覆盖：`TestScanValue_RoundTrip`, `TestScanValue_ConcurrentAccess`。
