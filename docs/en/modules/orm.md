---
title: orm.go - Database Operations
---

# orm.go - Database Operations

## Overview

`orm.go` provides database operation utilities for convenient data conversion between Go structs and database values. It implements `driver.Valuer` and `sql.Scanner` interfaces for seamless database integration.

## Functions

### Scan()

Scan database results into structs.

```go
func Scan(src interface{}, dst interface{}) error
```

**Parameters:**
- `src` - Source data ([]byte or string)
- `dst` - Destination struct pointer

**Returns:**
- Error if scanning fails

**Supported Formats:**
- JSON objects: `{...}`
- JSON arrays: `[...]`
- Empty values: Applies default values from struct tags

**Example:**
```go
type User struct {
    Name  string `json:"name" default:"Anonymous"`
    Email string `json:"email"`
    Age   int    `json:"age" default:"0"`
}

var user User
err := utils.Scan(dbResult, &user)
if err != nil {
    log.Errorf("Failed to scan user: %v", err)
}
```

**Notes:**
- Automatically detects JSON format from source
- Applies default values from struct tags when source is empty
- Logs errors during scanning process
- Supports both []byte and string inputs

---

### Value()

Convert struct to database value.

```go
func Value(m interface{}) (value driver.Value, err error)
```

**Parameters:**
- `m` - Struct or struct pointer

**Returns:**
- `value` - Database value ([]byte)
- `err` - Error if conversion fails

**Behavior:**
- Returns `[]byte("null")` if input is nil
- Applies default values for non-nil structs before conversion
- Converts struct to JSON format

**Example:**
```go
type User struct {
    Name  string `json:"name" default:"John"`
    Email string `json:"email" default:"john@example.com"`
}

user := User{Name: "Jane", Email: "jane@example.com"}
value, err := utils.Value(&user)
if err != nil {
    log.Errorf("Failed to convert user: %v", err)
}
// value is []byte(`{"name":"Jane","email":"jane@example.com"}`)
```

**Notes:**
- Implements `driver.Valuer` interface for database compatibility
- Automatically applies default values before conversion
- Returns JSON-encoded bytes
- Handles nil values gracefully

---

## Usage Patterns

### Database Model Definition

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

### CRUD Operations

```go
// Create
func CreateUser(db *sql.DB, user *User) error {
    value := utils.Must(utils.Value(user))
    _, err := db.Exec(
        "INSERT INTO users (data) VALUES (?)",
        value,
    )
    return err
}

// Read
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

// Update
func UpdateUser(db *sql.DB, user *User) error {
    value := utils.Must(utils.Value(user))
    _, err := db.Exec(
        "UPDATE users SET data = ? WHERE id = ?",
        value,
        user.ID,
    )
    return err
}

// Delete
func DeleteUser(db *sql.DB, id int64) error {
    _, err := db.Exec(
        "DELETE FROM users WHERE id = ?",
        id,
    )
    return err
}
```

### With GORM Integration

```go
type User struct {
    gorm.Model
    Profile Profile `gorm:"type:json"`
}

type Profile struct {
    Name  string `json:"name" default:""`
    Email string `json:"email" default:""`
}

// Scan implements sql.Scanner interface
func (p *Profile) Scan(value interface{}) error {
    return utils.Scan(value, p)
}

// Value implements driver.Valuer interface
func (p Profile) Value() (driver.Value, error) {
    return utils.Value(p)
}
```

### Batch Operations

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

## Best Practices

### Default Values

```go
type Config struct {
    Host     string `json:"host" default:"localhost"`
    Port     int    `json:"port" default:"8080"`
    Debug    bool   `json:"debug" default:"false"`
    Timeout  int    `json:"timeout" default:"30"`
}

// When scanning empty data, defaults are applied
var cfg Config
utils.Scan([]byte{}, &cfg)
// cfg.Host == "localhost"
// cfg.Port == 8080
// cfg.Debug == false
// cfg.Timeout == 30
```

### Error Handling

```go
func SafeScan(data []byte, dst interface{}) error {
    if err := utils.Scan(data, dst); err != nil {
        // Log error for debugging
        log.Errorf("Scan failed: %v, data: %s", err, string(data))
        
        // Return error or handle gracefully
        return fmt.Errorf("failed to scan data: %w", err)
    }
    return nil
}
```

### Performance Considerations

```go
// Good: Reuse structs for scanning
var user User
rows, _ := db.Query("SELECT data FROM users")
for rows.Next() {
    // Reuse the same struct
    rows.Scan(&data)
    utils.Scan(data, &user)
    // process user
}

// Avoid: Creating new structs in loop
for rows.Next() {
    var user User  // Creates new struct each iteration
    rows.Scan(&data)
    utils.Scan(data, &user)
}
```

## Integration Examples

### With SQL Database

```go
db, err := sql.Open("postgres", connStr)
if err != nil {
    log.Fatal(err)
}
defer db.Close()

// Create table
_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        data JSONB NOT NULL
    )
`)

// Insert user
user := &User{Name: "John", Email: "john@example.com"}
value := utils.Must(utils.Value(user))
_, err = db.Exec("INSERT INTO users (data) VALUES ($1)", value)
```

### With Redis

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

### With HTTP API

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

## Related Documentation

- [must.go](/en/modules/must) - Error assertion utilities
- [validator](/en/modules/validator) - Data validation
- [defaults](/en/modules/defaults) - Default value handling
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
