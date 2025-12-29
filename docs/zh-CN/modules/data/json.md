---
title: json - JSON 处理
---

# json - JSON 处理

## 概述

json 模块提供增强的 JSON 处理，具有更好的错误消息和平台特定优化。它包装了标准库 JSON 编码器/解码器，提供改进的功能。

## 函数

### Marshal()

将值编码为 JSON。

```go
func Marshal(v any) ([]byte, error)
```

**参数：**
- `v` - 要编码的值

**返回值：**
- JSON 编码的字节
- 如果编码失败，返回错误

**示例：**
```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

user := User{Name: "John", Email: "john@example.com"}
data, err := json.Marshal(user)
if err != nil {
    log.Errorf("编码失败: %v", err)
}
// data 是 []byte(`{"name":"John","email":"john@example.com"}`)
```

---

### Unmarshal()

将 JSON 数据解码为值。

```go
func Unmarshal(data []byte, v any) error
```

**参数：**
- `data` - JSON 编码的字节
- `v` - 目标值（必须是指针）

**返回值：**
- 如果解码失败，返回错误

**示例：**
```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

data := []byte(`{"name":"John","email":"john@example.com"}`)
var user User
if err := json.Unmarshal(data, &user); err != nil {
    log.Errorf("解码失败: %v", err)
}
```

---

### MarshalString()

将值编码为 JSON 字符串。

```go
func MarshalString(v any) (string, error)
```

**参数：**
- `v` - 要编码的值

**返回值：**
- JSON 编码的字符串
- 如果编码失败，返回错误

**示例：**
```go
user := User{Name: "John", Email: "john@example.com"}
str, err := json.MarshalString(user)
if err != nil {
    log.Errorf("编码失败: %v", err)
}
// str 是 `{"name":"John","email":"john@example.com"}`
```

---

### UnmarshalString()

将 JSON 字符串解码为值。

```go
func UnmarshalString(data string, v any) error
```

**参数：**
- `data` - JSON 编码的字符串
- `v` - 目标值（必须是指针）

**返回值：**
- 如果解码失败，返回错误

**示例：**
```go
data := `{"name":"John","email":"john@example.com"}`
var user User
if err := json.UnmarshalString(data, &user); err != nil {
    log.Errorf("解码失败: %v", err)
}
```

---

### NewEncoder()

创建新的 JSON 编码器。

```go
func NewEncoder(w io.Writer) *json.Encoder
```

**参数：**
- `w` - 要编码到的写入器

**返回值：**
- JSON 编码器

**示例：**
```go
file, err := os.Create("users.json")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

encoder := json.NewEncoder(file)
encoder.Encode(User{Name: "John", Email: "john@example.com"})
encoder.Encode(User{Name: "Jane", Email: "jane@example.com"})
```

---

### NewDecoder()

创建新的 JSON 解码器。

```go
func NewDecoder(r io.Reader) *json.Decoder
```

**参数：**
- `r` - 要解码的读取器

**返回值：**
- JSON 解码器

**示例：**
```go
file, err := os.Open("users.json")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

decoder := json.NewDecoder(file)
for decoder.More() {
    var user User
    if err := decoder.Decode(&user); err != nil {
        log.Errorf("解码失败: %v", err)
        break
    }
    // 处理用户
}
```

---

## 使用模式

### HTTP API

```go
func createUserHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    if err := db.Create(&user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
```

### 文件 I/O

```go
func saveUsers(users []User, path string) error {
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()
    
    encoder := json.NewEncoder(file)
    for _, user := range users {
        if err := encoder.Encode(user); err != nil {
            return err
        }
    }
    return nil
}

func loadUsers(path string) ([]User, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    var users []User
    decoder := json.NewDecoder(file)
    for decoder.More() {
        var user User
        if err := decoder.Decode(&user); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}
```

### 配置

```go
type Config struct {
    Host     string `json:"host"`
    Port     int    `json:"port"`
    Debug    bool   `json:"debug"`
    Database struct {
        Name     string `json:"name"`
        User     string `json:"user"`
        Password string `json:"password"`
    } `json:"database"`
}

func loadConfig(path string) (*Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }
    
    var cfg Config
    if err := json.Unmarshal(data, &cfg); err != nil {
        return nil, err
    }
    
    return &cfg, nil
}

func saveConfig(cfg *Config, path string) error {
    data, err := json.Marshal(cfg)
    if err != nil {
        return err
    }
    
    return os.WriteFile(path, data, 0644)
}
```

### 美化打印

```go
func prettyPrint(v interface{}) error {
    data, err := json.MarshalIndent(v, "", "  ")
    if err != nil {
        return err
    }
    fmt.Println(string(data))
    return nil
}

func prettyPrintToFile(v interface{}, path string) error {
    file, err := os.Create(path)
    if err != nil {
        return err
    }
    defer file.Close()
    
    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    return encoder.Encode(v)
}
```

---

## 高级功能

### 自定义 Marshal/Unmarshal

```go
type Time struct {
    time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
    return []byte(fmt.Sprintf(`"%s"`, t.Time.Format("2006-01-02"))), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
    str := string(data)
    if str == "null" {
        return nil
    }
    
    str = strings.Trim(str, `"`)
    parsed, err := time.Parse("2006-01-02", str)
    if err != nil {
        return err
    }
    
    t.Time = parsed
    return nil
}

type Event struct {
    Name string `json:"name"`
    Date Time  `json:"date"`
}
```

### 流式处理

```go
func processLargeJSON(reader io.Reader, processor func(User) error) error {
    decoder := json.NewDecoder(reader)
    
    for decoder.More() {
        var user User
        if err := decoder.Decode(&user); err != nil {
            return err
        }
        
        if err := processor(user); err != nil {
            return err
        }
    }
    
    return nil
}

func processUsersFile(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()
    
    return processLargeJSON(file, func(user User) error {
        fmt.Printf("处理中: %s\n", user.Name)
        return nil
    })
}
```

### 错误处理

```go
func safeUnmarshal(data []byte, v interface{}) error {
    if err := json.Unmarshal(data, v); err != nil {
        // 记录详细的错误信息
        log.Errorf("JSON 解码失败: %v", err)
        log.Errorf("数据: %s", string(data))
        
        // 返回包装的错误
        return fmt.Errorf("解码 JSON 失败: %w", err)
    }
    return nil
}

func validateJSON(data []byte) error {
    var v interface{}
    if err := json.Unmarshal(data, &v); err != nil {
        return fmt.Errorf("无效的 JSON: %w", err)
    }
    return nil
}
```

---

## 最佳实践

### 错误处理

```go
// 好：正确处理错误
data, err := json.Marshal(user)
if err != nil {
    log.Errorf("编码用户失败: %v", err)
    return err
}

// 好：解码前验证 JSON
if !json.Valid(data) {
    return fmt.Errorf("无效的 JSON 数据")
}

var user User
if err := json.Unmarshal(data, &user); err != nil {
    log.Errorf("解码用户失败: %v", err)
    return err
}
```

### 内存效率

```go
// 好：对大文件使用流式处理
func processLargeFile(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()
    
    decoder := json.NewDecoder(file)
    for decoder.More() {
        var item interface{}
        if err := decoder.Decode(&item); err != nil {
            return err
        }
        // 处理项目
    }
    return nil
}

// 避免：将整个文件加载到内存
func processLargeFileBad(path string) error {
    data, err := os.ReadFile(path)  // 加载整个文件
    if err != nil {
        return err
    }
    
    var items []interface{}
    if err := json.Unmarshal(data, &items); err != nil {
        return err
    }
    // 处理项目
    return nil
}
```

### 性能

```go
// 好：重用编码器/解码器
var encoder = json.NewEncoder(os.Stdout)
var decoder = json.NewDecoder(os.Stdin)

func process(data []byte) error {
    var v interface{}
    return decoder.Decode(&v)
}

// 避免：每次都创建新编码器/解码器
func processBad(data []byte) error {
    decoder := json.NewDecoder(bytes.NewReader(data))  // 昂贵
    var v interface{}
    return decoder.Decode(&v)
}
```

---

## 平台特定优化

json 模块根据平台自动选择最佳实现：

### Linux AMD64

使用 `sonic` JSON 库以获得最大性能：
- 比标准库 **快 3.5 倍**
- 零分配优化
- SIMD 加速解析

### 其他平台

使用具有增强错误消息的标准库 JSON：
- 与所有平台兼容
- 更好的错误消息用于调试
- 标准库性能

---

## 相关文档

- [candy](/zh-CN/modules/candy) - 类型转换
- [stringx](/zh-CN/modules/stringx) - 字符串工具
- [anyx](/zh-CN/modules/anyx) - Interface{} 助手
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
