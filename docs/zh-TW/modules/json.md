---
title: json - JSON 處理
---

# json - JSON 處理

## 概述

json 模組提供增強的 JSON 處理，具有更好的錯誤訊息和平台特定優化。它包裝了標準庫 JSON 編碼器/解碼器，提供改進的功能。

## 函數

### Marshal()

將值編碼為 JSON。

```go
func Marshal(v any) ([]byte, error)
```

**參數：**
- `v` - 要編碼的值

**返回值：**
- JSON 編碼的位元組
- 如果編碼失敗，返回錯誤

**示例：**
```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

user := User{Name: "John", Email: "john@example.com"}
data, err := json.Marshal(user)
if err != nil {
    log.Errorf("編碼失敗: %v", err)
}
// data 是 []byte(`{"name":"John","email":"john@example.com"}`)
```

---

### Unmarshal()

將 JSON 資料解碼為值。

```go
func Unmarshal(data []byte, v any) error
```

**參數：**
- `data` - JSON 編碼的位元組
- `v` - 目標值（必須是指針）

**返回值：**
- 如果解碼失敗，返回錯誤

**示例：**
```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

data := []byte(`{"name":"John","email":"john@example.com"}`)
var user User
if err := json.Unmarshal(data, &user); err != nil {
    log.Errorf("解碼失敗: %v", err)
}
```

---

### MarshalString()

將值編碼為 JSON 字串。

```go
func MarshalString(v any) (string, error)
```

**參數：**
- `v` - 要編碼的值

**返回值：**
- JSON 編碼的字串
- 如果編碼失敗，返回錯誤

**示例：**
```go
user := User{Name: "John", Email: "john@example.com"}
str, err := json.MarshalString(user)
if err != nil {
    log.Errorf("編碼失敗: %v", err)
}
// str 是 `{"name":"John","email":"john@example.com"}`
```

---

### UnmarshalString()

將 JSON 字串解碼為值。

```go
func UnmarshalString(data string, v any) error
```

**參數：**
- `data` - JSON 編碼的字串
- `v` - 目標值（必須是指針）

**返回值：**
- 如果解碼失敗，返回錯誤

**示例：**
```go
data := `{"name":"John","email":"john@example.com"}`
var user User
if err := json.UnmarshalString(data, &user); err != nil {
    log.Errorf("解碼失敗: %v", err)
}
```

---

### NewEncoder()

建立新的 JSON 編碼器。

```go
func NewEncoder(w io.Writer) *json.Encoder
```

**參數：**
- `w` - 要編碼到的寫入器

**返回值：**
- JSON 編碼器

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

建立新的 JSON 解碼器。

```go
func NewDecoder(r io.Reader) *json.Decoder
```

**參數：**
- `r` - 要解碼的讀取器

**返回值：**
- JSON 解碼器

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
        log.Errorf("解碼失敗: %v", err)
        break
    }
    // 處理使用者
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

### 檔案 I/O

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

### 美化列印

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

## 高級功能

### 自定義 Marshal/Unmarshal

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

### 串流處理

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
        fmt.Printf("處理中: %s\n", user.Name)
        return nil
    })
}
```

### 錯誤處理

```go
func safeUnmarshal(data []byte, v interface{}) error {
    if err := json.Unmarshal(data, v); err != nil {
        // 記錄詳細的錯誤訊息
        log.Errorf("JSON 解碼失敗: %v", err)
        log.Errorf("資料: %s", string(data))
        
        // 返回包裝的錯誤
        return fmt.Errorf("解碼 JSON 失敗: %w", err)
    }
    return nil
}

func validateJSON(data []byte) error {
    var v interface{}
    if err := json.Unmarshal(data, &v); err != nil {
        return fmt.Errorf("無效的 JSON: %w", err)
    }
    return nil
}
```

---

## 最佳實踐

### 錯誤處理

```go
// 好：正確處理錯誤
data, err := json.Marshal(user)
if err != nil {
    log.Errorf("編碼使用者失敗: %v", err)
    return err
}

// 好：解碼前驗證 JSON
if !json.Valid(data) {
    return fmt.Errorf("無效的 JSON 資料")
}

var user User
if err := json.Unmarshal(data, &user); err != nil {
    log.Errorf("解碼使用者失敗: %v", err)
    return err
}
```

### 記憶體效率

```go
// 好：對大檔案使用串流處理
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
        // 處理項目
    }
    return nil
}

// 避免：將整個檔案載入到記憶體
func processLargeFileBad(path string) error {
    data, err := os.ReadFile(path)  // 載入整個檔案
    if err != nil {
        return err
    }
    
    var items []interface{}
    if err := json.Unmarshal(data, &items); err != nil {
        return err
    }
    // 處理項目
    return nil
}
```

### 效能

```go
// 好：重用編碼器/解碼器
var encoder = json.NewEncoder(os.Stdout)
var decoder = json.NewDecoder(os.Stdin)

func process(data []byte) error {
    var v interface{}
    return decoder.Decode(&v)
}

// 避免：每次都建立新編碼器/解碼器
func processBad(data []byte) error {
    decoder := json.NewDecoder(bytes.NewReader(data))  // 昂貴
    var v interface{}
    return decoder.Decode(&v)
}
```

---

## 平台特定優化

json 模組根據平台自動選擇最佳實現：

### Linux AMD64

使用 `sonic` JSON 庫以獲得最大效能：
- 比標準庫 **快 3.5 倍**
- 零分配優化
- SIMD 加速解析

### 其他平台

使用具有增強錯誤訊息的標準庫 JSON：
- 與所有平台相容
- 更好的錯誤訊息用於除錯
- 標準庫效能

---

## 相關文檔

- [candy](/zh-TW/modules/candy) - 類型轉換
- [stringx](/zh-TW/modules/stringx) - 字串工具
- [anyx](/zh-TW/modules/anyx) - Interface{} 助手
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
