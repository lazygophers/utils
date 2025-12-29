---
title: orm.go - 資料庫操作
---

# orm.go - 資料庫操作

## 概述

`orm.go` 提供資料庫操作工具，用於在 Go 結構體和資料庫值之間進行便捷的資料轉換。它實現了 `driver.Valuer` 和 `sql.Scanner` 介面以實現無縫資料庫集成。

## 函數

### Scan()

將資料庫結果掃描到結構體中。

```go
func Scan(src interface{}, dst interface{}) error
```

**參數：**
- `src` - 源資料（[]byte 或 string）
- `dst` - 目標結構體指針

**返回值：**
- 如果掃描失敗，返回錯誤

**支援的格式：**
- JSON 對象：`{...}`
- JSON 陣列：`[...]`
- 空值：應用結構體標籤中的預設值

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
    log.Errorf("掃描使用者失敗: %v", err)
}
```

**注意：**
- 自動從源檢測 JSON 格式
- 當源為空時，應用結構體標籤中的預設值
- 在掃描過程中記錄錯誤
- 支援 []byte 和 string 輸入

---

### Value()

將結構體轉換為資料庫值。

```go
func Value(m interface{}) (value driver.Value, err error)
```

**參數：**
- `m` - 結構體或結構體指針

**返回值：**
- `value` - 資料庫值（[]byte）
- `err` - 如果轉換失敗則返回錯誤

**行為：**
- 如果輸入為 nil，返回 `[]byte("null")`
- 在轉換前為非 nil 結構體應用預設值
- 將結構體轉換為 JSON 格式

**示例：**
```go
type User struct {
    Name  string `json:"name" default:"John"`
    Email string `json:"email" default:"john@example.com"`
}

user := User{Name: "Jane", Email: "jane@example.com"}
value, err := utils.Value(&user)
if err != nil {
    log.Errorf("轉換使用者失敗: %v", err)
}
// value 是 []byte(`{"name":"Jane","email":"jane@example.com"}`)
```

**注意：**
- 實現 `driver.Valuer` 介面以實現資料庫相容性
- 在轉換前自動應用預設值
- 返回 JSON 編碼的位元組
- 優雅地處理 nil 值

---

## 使用模式

### 資料庫模型定義

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
// 建立資料
func CreateUser(db *sql.DB, user *User) error {
    value := utils.Must(utils.Value(user))
    _, err := db.Exec(
        "INSERT INTO users (data) VALUES (?)",
        value,
    )
    return err
}

// 讀取資料
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

// 更新資料
func UpdateUser(db *sql.DB, user *User) error {
    value := utils.Must(utils.Value(user))
    _, err := db.Exec(
        "UPDATE users SET data = ? WHERE id = ?",
        value,
        user.ID,
    )
    return err
}

// 刪除資料
func DeleteUser(db *sql.DB, id int64) error {
    _, err := db.Exec(
        "DELETE FROM users WHERE id = ?",
        id,
    )
    return err
}
```

### 與 GORM 集成

```go
type User struct {
    gorm.Model
    Profile Profile `gorm:"type:json"`
}

type Profile struct {
    Name  string `json:"name" default:""`
    Email string `json:"email" default:""`
}

// Scan 實現 sql.Scanner 介面
func (p *Profile) Scan(value interface{}) error {
    return utils.Scan(value, p)
}

// Value 實現 driver.Valuer 介面
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

## 最佳實踐

### 預設值

```go
type Config struct {
    Host     string `json:"host" default:"localhost"`
    Port     int    `json:"port" default:"8080"`
    Debug    bool   `json:"debug" default:"false"`
    Timeout  int    `json:"timeout" default:"30"`
}

// 掃描空資料時，應用預設值
var cfg Config
utils.Scan([]byte{}, &cfg)
// cfg.Host == "localhost"
// cfg.Port == 8080
// cfg.Debug == false
// cfg.Timeout == 30
```

### 錯誤處理

```go
func SafeScan(data []byte, dst interface{}) error {
    if err := utils.Scan(data, dst); err != nil {
        // 記錄錯誤以便除錯
        log.Errorf("掃描失敗: %v, 資料: %s", err, string(data))
        
        // 返回錯誤或優雅處理
        return fmt.Errorf("掃描資料失敗: %w", err)
    }
    return nil
}
```

### 效能考慮

```go
// 好：重用結構體進行掃描
var user User
rows, _ := db.Query("SELECT data FROM users")
for rows.Next() {
    // 重用同一個結構體
    rows.Scan(&data)
    utils.Scan(data, &user)
    // 處理使用者
}

// 避免：在循環中建立新結構體
for rows.Next() {
    var user User  // 每次迭代建立新結構體
    rows.Scan(&data)
    utils.Scan(data, &user)
}
```

## 集成示例

### 與 SQL 資料庫

```go
db, err := sql.Open("postgres", connStr)
if err != nil {
    log.Fatal(err)
}
defer db.Close()

// 建立表
_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        data JSONB NOT NULL
    )
`)

// 插入使用者
user := &User{Name: "John", Email: "john@example.com"}
value := utils.Must(utils.Value(user))
_, err = db.Exec("INSERT INTO users (data) VALUES ($1)", value)
```

### 與 Redis

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

### 與 HTTP API

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

## 相關文檔

- [must.go](/zh-TW/modules/must) - 錯誤斷言工具
- [validator](/zh-TW/modules/validator) - 資料驗證
- [defaults](/zh-TW/modules/defaults) - 預設值處理
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
