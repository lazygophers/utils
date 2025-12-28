---
title: orm.go - 数据库操作
---

# orm.go - 数据库操作

## 概述

`orm.go` 提供数据库操作工具，用于在 Go 结构体和数据库值之间进行便捷的数据转换。它实现了 `driver.Valuer` 和 `sql.Scanner` 接口以实现无缝数据库集成。

## 函数

### Scan()

将数据库结果扫描到结构体中。

```go
func Scan(src interface{}, dst interface{}) error
```

**参数：**
- `src` - 源数据（[]byte 或 string）
- `dst` - 目标结构体指针

**返回值：**
- 如果扫描失败，返回错误

**支持的格式：**
- JSON 对象：`{...}`
- JSON 数组：`[...]`
- 空值：应用结构体标签中的默认值

**示例：**
```go
type User struct {
    Name  string `json:"name" default:"Anonymous"`
    Email string `json:"email"`
    Age   int    `json:"age" default:"0"`
}

var user User
err := utils.Scan(dbResult, &user)
if err != nil {
    log.Errorf("扫描用户失败: %v", err)
}
```

**注意：**
- 自动从源检测 JSON 格式
- 当源为空时，应用结构体标签中的默认值
- 在扫描过程中记录错误
- 支持 []byte 和 string 输入

---

### Value()

将结构体转换为数据库值。

```go
func Value(m interface{}) (value driver.Value, err error)
```

**参数：**
- `m` - 结构体或结构体指针

**返回值：**
- `value` - 数据库值（[]byte）
- `err` - 如果转换失败则返回错误

**行为：**
- 如果输入为 nil，返回 `[]byte("null")`
- 在转换前为非 nil 结构体应用默认值
- 将结构体转换为 JSON 格式

**示例：**
```go
type User struct {
    Name  string `json:"name" default:"John"`
    Email string `json:"email" default:"john@example.com"`
}

user := User{Name: "Jane", Email: "jane@example.com"}
value, err := utils.Value(&user)
if err != nil {
    log.Errorf("转换用户失败: %v", err)
}
// value 是 []byte(`{"name":"Jane","email":"jane@example.com"}`)
```

**注意：**
- 实现 `driver.Valuer` 接口以实现数据库兼容性
- 在转换前自动应用默认值
- 返回 JSON 编码的字节
- 优雅地处理 nil 值

---

## 使用模式

### 数据库模型定义

```go
type User struct {
    ID        int64     `json:"id" default:"0"`
    Name      string    `json:"name" default:""`
    Email     string    `json:"email" default:""`
    Age       int       `json:"age" default:"0"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) Scan(value interface{}) error {
    return utils.Scan(value, u)
}

func (u *User) Value() (driver.Value, error) {
    return utils.Value(u)
}
```

### CRUD 操作

```go
// 创建
func CreateUser(db *sql.DB, user *User) error {
    value := utils.Must(utils.Value(user))
    _, err := db.Exec(
        "INSERT INTO users (data) VALUES (?)",
        value,
    )
    return err
}

// 读取
func GetUser(db *sql.DB, id int64) (*User, error) {
    var data []byte
    err := db.QueryRow(
        "SELECT data FROM users WHERE id = ?",
        id,
    ).Scan(&data)
    if err != nil {
        return nil, err
    }
    
    var user User
    err = utils.Scan(data, &user)
    return &user, err
}

// 更新
func UpdateUser(db *sql.DB, user *User) error {
    value := utils.Must(utils.Value(user))
    _, err := db.Exec(
        "UPDATE users SET data = ? WHERE id = ?",
        value,
        user.ID,
    )
    return err
}

// 删除
func DeleteUser(db *sql.DB, id int64) error {
    _, err := db.Exec(
        "DELETE FROM users WHERE id = ?",
        id,
    )
    return err
}
```

### 与 GORM 集成

```go
type User struct {
    gorm.Model
    Profile Profile `gorm:"type:json"`
}

type Profile struct {
    Name  string `json:"name" default:""`
    Email string `json:"email" default:""`
}

// Scan 实现 sql.Scanner 接口
func (p *Profile) Scan(value interface{}) error {
    return utils.Scan(value, p)
}

// Value 实现 driver.Valuer 接口
func (p Profile) Value() (driver.Value, error) {
    return utils.Value(p)
}
```

### 批量操作

```go
func BatchInsertUsers(db *sql.DB, users []*User) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()
    
    stmt, err := tx.Prepare(
        "INSERT INTO users (data) VALUES (?)",
    )
    if err != nil {
        return err
    }
    defer stmt.Close()
    
    for _, user := range users {
        value := utils.Must(utils.Value(user))
        _, err := stmt.Exec(value)
        if err != nil {
            return err
        }
    }
    
    return tx.Commit()
}
```

## 最佳实践

### 默认值

```go
type Config struct {
    Host     string `json:"host" default:"localhost"`
    Port     int    `json:"port" default:"8080"`
    Debug    bool   `json:"debug" default:"false"`
    Timeout  int    `json:"timeout" default:"30"`
}

// 扫描空数据时，应用默认值
var cfg Config
utils.Scan([]byte{}, &cfg)
// cfg.Host == "localhost"
// cfg.Port == 8080
// cfg.Debug == false
// cfg.Timeout == 30
```

### 错误处理

```go
func SafeScan(data []byte, dst interface{}) error {
    if err := utils.Scan(data, dst); err != nil {
        // 记录错误以便调试
        log.Errorf("扫描失败: %v, 数据: %s", err, string(data))
        
        // 返回错误或优雅处理
        return fmt.Errorf("扫描数据失败: %w", err)
    }
    return nil
}
```

### 性能考虑

```go
// 好：重用结构体进行扫描
var user User
rows, _ := db.Query("SELECT data FROM users")
for rows.Next() {
    // 重用同一个结构体
    rows.Scan(&data)
    utils.Scan(data, &user)
    // 处理用户
}

// 避免：在循环中创建新结构体
for rows.Next() {
    var user User  // 每次迭代创建新结构体
    rows.Scan(&data)
    utils.Scan(data, &user)
}
```

## 集成示例

### 与 SQL 数据库

```go
db, err := sql.Open("postgres", connStr)
if err != nil {
    log.Fatal(err)
}
defer db.Close()

// 创建表
_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        data JSONB NOT NULL
    )
`)

// 插入用户
user := &User{Name: "John", Email: "john@example.com"}
value := utils.Must(utils.Value(user))
_, err = db.Exec("INSERT INTO users (data) VALUES ($1)", value)
```

### 与 Redis

```go
func SaveUserToRedis(client *redis.Client, key string, user *User) error {
    value := utils.Must(utils.Value(user))
    return client.Set(ctx, key, value, 0).Err()
}

func LoadUserFromRedis(client *redis.Client, key string) (*User, error) {
    data, err := client.Get(ctx, key).Bytes()
    if err != nil {
        return nil, err
    }
    
    var user User
    err = utils.Scan(data, &user)
    return &user, err
}
```

### 与 HTTP API

```go
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    value := utils.Must(utils.Value(user))
    if _, err := db.Exec("INSERT INTO users (data) VALUES (?)", value); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusCreated)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    
    var data []byte
    if err := db.QueryRow("SELECT data FROM users WHERE id = ?", id).Scan(&data); err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    
    var user User
    if err := utils.Scan(data, &user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    json.NewEncoder(w).Encode(user)
}
```

## 相关文档

- [must.go](/zh-CN/modules/must) - 错误断言工具
- [validator](/zh-CN/modules/validator) - 数据验证
- [defaults](/zh-CN/modules/defaults) - 默认值处理
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
