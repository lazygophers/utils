---
title: config - 配置管理
---

# config - 配置管理

## 概述

config 模組提供配置文件載入，支援多種格式包括 JSON、YAML、TOML、INI、HCL 等。

## 支援的格式

- **JSON** - `.json`, `.json5`
- **YAML** - `.yaml`, `.yml`
- **TOML** - `.toml`
- **INI** - `.ini`
- **HCL** - `.hcl`
- **XML** - `.xml`
- **Properties** - `.properties`
- **ENV** - `.env`

---

## 核心函數

### LoadConfig()

載入配置文件並驗證。

```go
func LoadConfig(c any, paths ...string) error
```

**參數：**
- `c` - 配置結構體指針
- `paths` - 可選的文件路徑列表

**返回值：**
- 如果載入或驗證失敗，返回錯誤

**搜索順序：**
1. 提供的顯式路徑
2. 環境變量 `LAZYGOPHERS_CONFIG`
3. 當前目錄（`conf.*` 或 `config.*`）
4. 可執行文件目錄（`conf.*` 或 `config.*`）

**示例：**
```go
type Config struct {
    Host     string `json:"host" toml:"host" yaml:"host" ini:"host"`
    Port     int    `json:"port" toml:"port" yaml:"port" ini:"port"`
    Debug    bool   `json:"debug" toml:"debug" yaml:"debug" ini:"debug"`
    Database struct {
        Name     string `json:"name" toml:"name" yaml:"name" ini:"name"`
        User     string `json:"user" toml:"user" yaml:"user" ini:"user"`
        Password string `json:"password" toml:"password" yaml:"password" ini:"password"`
    } `json:"database" toml:"database" yaml:"database" ini:"database"`
}

var cfg Config
if err := config.LoadConfig(&cfg, "config.json"); err != nil {
    log.Fatalf("載入配置失敗: %v", err)
}
```

---

### LoadConfigSkipValidate()

載入配置文件但不驗證。

```go
func LoadConfigSkipValidate(c any, paths ...string) error
```

**參數：**
- `c` - 配置結構體指針
- `paths` - 可選的文件路徑列表

**返回值：**
- 如果載入失敗，返回錯誤

**示例：**
```go
var cfg Config
if err := config.LoadConfigSkipValidate(&cfg, "config.yaml"); err != nil {
    log.Fatalf("載入配置失敗: %v", err)
}
```

---

### SetConfig()

保存配置到文件。

```go
func SetConfig(c any) error
```

**參數：**
- `c` - 要保存的配置結構體

**返回值：**
- 如果保存失敗，返回錯誤

**示例：**
```go
cfg := Config{
    Host: "localhost",
    Port: 8080,
    Debug: true,
}

if err := config.SetConfig(&cfg); err != nil {
    log.Fatalf("保存配置失敗: %v", err)
}
```

---

### RegisterParser()

為文件擴展名註冊自定義解析器。

```go
func RegisterParser(ext string, m Marshaler, u Unmarshaler)
```

**參數：**
- `ext` - 文件擴展名（例如 ".custom"）
- `m` - 編碼器函數
- `u` - 解碼器函數

**示例：**
```go
config.RegisterParser(".custom",
    func(writer io.Writer, v interface{}) error {
        // 自定義編碼邏輯
        return nil
    },
    func(reader io.Reader, v interface{}) error {
        // 自定義解碼邏輯
        return nil
    },
)
```

---

## 使用模式

### 多格式支援

```go
type Config struct {
    Host string `json:"host" toml:"host" yaml:"host" ini:"host"`
    Port int    `json:"port" toml:"port" yaml:"port" ini:"port"`
}

// 支援任何格式
var cfg Config
config.LoadConfig(&cfg, "config.json")   // JSON
config.LoadConfig(&cfg, "config.yaml")   // YAML
config.LoadConfig(&cfg, "config.toml")   // TOML
config.LoadConfig(&cfg, "config.ini")    // INI
```

### 基於環境的載入

```go
func loadConfig() *Config {
    var cfg Config
    
    // 嘗試多個路徑
    paths := []string{
        "/etc/myapp/config.json",
        os.Getenv("HOME") + "/.myapp/config.json",
        "./config.json",
    }
    
    if err := config.LoadConfig(&cfg, paths...); err != nil {
        log.Warnf("使用默認配置: %v", err)
        return &Config{
            Host: "localhost",
            Port: 8080,
        }
    }
    
    return &cfg
}
```

### 嵌套配置

```go
type Config struct {
    Server struct {
        Host string `json:"host"`
        Port int    `json:"port"`
    } `json:"server"`
    
    Database struct {
        Host     string `json:"host"`
        Port     int    `json:"port"`
        Name     string `json:"name"`
        User     string `json:"user"`
        Password string `json:"password"`
    } `json:"database"`
}

var cfg Config
config.LoadConfig(&cfg, "config.json")
```

### 配置驗證

```go
type Config struct {
    Host     string `json:"host" validate:"required"`
    Port     int    `json:"port" validate:"required,min=1,max=65535"`
    Database string `json:"database" validate:"required"`
}

var cfg Config
if err := config.LoadConfig(&cfg, "config.json"); err != nil {
    log.Fatalf("配置驗證失敗: %v", err)
}
```

---

## 最佳實踐

### 默認值

```go
type Config struct {
    Host     string `json:"host" default:"localhost"`
    Port     int    `json:"port" default:"8080"`
    Debug    bool   `json:"debug" default:"false"`
}

var cfg Config
config.LoadConfigSkipValidate(&cfg, "config.json")
// 如果配置文件中沒有，則應用默認值
```

### 錯誤處理

```go
func loadConfigWithFallback() *Config {
    var cfg Config
    
    if err := config.LoadConfig(&cfg, "config.json"); err != nil {
        log.Warnf("載入配置失敗，使用默認值: %v", err)
        
        return &Config{
            Host: "localhost",
            Port: 8080,
            Debug: false,
        }
    }
    
    return &cfg
}
```

### 配置熱重載

```go
func watchConfig(path string, cfg *Config) {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatalf("建立監視器失敗: %v", err)
    }
    
    watcher.Add(path)
    
    for {
        select {
        case event := <-watcher.Events:
            if event.Op&fsnotify.Write == fsnotify.Write {
                log.Info("配置已更改，正在重新載入...")
                
                if err := config.LoadConfig(cfg, path); err != nil {
                    log.Errorf("重新載入配置失敗: %v", err)
                }
            }
        case err := <-watcher.Errors:
            log.Errorf("監視器錯誤: %v", err)
        }
    }
}
```

---

## 相關文檔

- [validator](/zh-TW/modules/validator) - 資料驗證
- [defaults](/zh-TW/modules/defaults) - 默認值
- [json](/zh-TW/modules/json) - JSON 處理
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
