---
title: must.go - 錯誤斷言
---

# must.go - 錯誤斷言

## 概述

`must.go` 提供錯誤斷言工具，透過基於 panic 的錯誤檢查簡化錯誤處理流程。這些函數專為初始化和關鍵操作設計，當失敗時應終止程式執行。

## 函數

### MustOk()

斷言第二個返回值為 true。

```go
func MustOk[T any](value T, ok bool) T
```

**參數：**
- `value` - 返回值
- `ok` - 成功標誌

**返回值：**
- 如果 `ok` 為 true，返回 `value`
- 如果 `ok` 為 false，panic 並顯示訊息 "is not ok"

**示例：**
```go
value, ok := getValue()
result := utils.MustOk(value, ok)
```

**注意：**
- 當第二個返回值表示成功/失敗時使用此函數
- 如果斷言失敗，panic 會終止程式執行
- 泛型類型 `T` 允許任何返回類型
- 常用於 map 查找和類型斷言

---

### MustSuccess()

斷言錯誤為 nil。

```go
func MustSuccess(err error)
```

**參數：**
- `err` - 要檢查的錯誤

**行為：**
- 如果 `err` 為 nil，不執行任何操作
- 如果 `err` 不為 nil，panic 並顯示格式化的錯誤訊息

**示例：**
```go
utils.MustSuccess(config.Load(&cfg, "config.json"))
utils.MustSuccess(os.MkdirAll("data", 0755))
utils.MustSuccess(db.Ping())
```

**注意：**
- 常用於初始化和設置操作
- Panic 訊息包含錯誤詳情以便除錯
- 用於程式必須成功才能繼續的操作

---

### Must()

組合驗證函數，檢查錯誤狀態並返回值。

```go
func Must[T any](value T, err error) T
```

**參數：**
- `value` - 返回值
- `err` - 錯誤

**返回值：**
- 如果 `err` 為 nil，返回 `value`
- 如果 `err` 不為 nil，panic

**示例：**
```go
data := utils.Must(loadData())
file := utils.Must(os.Open("file.txt"))
result := utils.Must(http.Get(url))
conn := utils.Must(net.Listen("tcp", ":8080"))
```

**注意：**
- must 模組中最常用的函數
- 組合錯誤檢查和值提取
- 泛型類型 `T` 允許任何返回類型
- 適用於返回 `(T, error)` 的函數

---

### Ignore()

強制忽略任何參數。

```go
func Ignore[T any](value T, _ any) T
```

**參數：**
- `value` - 要返回的值
- `_` - 被忽略的參數

**返回值：**
- 返回 `value`

**示例：**
```go
result := utils.Ignore(data, err)
```

**注意：**
- 當需要抑制 linter 關於忽略值的警告時使用
- 第二個參數被顯式忽略
- 有助於保持程式碼整潔而不出現 linter 警告
- 不實際處理錯誤，只是抑制警告

---

## 使用模式

### 初始化鏈

```go
func initApp() {
    cfg := utils.Must(loadConfig())
    db := utils.Must(connectDB(cfg.DatabaseURL))
    server := utils.Must(createServer(cfg.Port))
    
    utils.MustSuccess(server.Start())
}
```

### 檔案操作

```go
func readFile(path string) []byte {
    file := utils.Must(os.Open(path))
    defer file.Close()
    
    data := utils.Must(io.ReadAll(file))
    return data
}
```

### Map 操作

```go
func getValue(m map[string]int, key string) int {
    value, ok := m[key]
    return utils.MustOk(value, ok)
}
```

### 配置載入

```go
type Config struct {
    Host string `json:"host"`
    Port int    `json:"port"`
}

func loadConfig(path string) *Config {
    data := utils.Must(os.ReadFile(path))
    
    var cfg Config
    utils.MustSuccess(json.Unmarshal(data, &cfg))
    
    return &cfg
}
```

## 最佳實踐

### 何時使用 Must 函數

**使用 `Must()` 當：**
- 操作對程式啟動至關重要
- 失敗應終止程式執行
- 錯誤恢復不可能或沒有意義
- 你在初始化程式碼中（main、init）

**避免 `Must()` 當：**
- 處理使用者輸入
- 可能失敗的網路請求
- 可能不存在的檔案操作
- 任何可能合理失敗的操作

### 錯誤處理 vs Panic

```go
// 好：使用 Must 進行初始化
func init() {
    config = utils.Must(loadConfig())
}

// 好：為使用者操作處理錯誤
func handleUserRequest() error {
    data, err := loadData()
    if err != nil {
        return err
    }
    // 處理資料
    return nil
}
```

## 相關文檔

- [orm.go](/zh-TW/modules/orm) - 資料庫操作
- [validator](/zh-TW/modules/validator) - 資料驗證
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
