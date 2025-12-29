---
title: config - 配置管理
---

# config - 配置管理

## 概述

config 模块提供配置文件加载，支持多种格式包括 JSON、YAML、TOML、INI、HCL 等。

## 支持的格式

- **JSON** - `.json`, `.json5`
- **YAML** - `.yaml`, `.yml`
- **TOML** - `.toml`
- **INI** - `.ini`
- **HCL** - `.hcl`
- **XML** - `.xml`
- **Properties** - `.properties`
- **ENV** - `.env`

---

## 核心函数

### LoadConfig()

加载配置文件并验证。

```go
func LoadConfig(c any, paths ...string) error
```

**参数：**
- `c` - 配置结构体指针
- `paths` - 可选的文件路径列表

**返回值：**
- 如果加载或验证失败，返回错误

**搜索顺序：**
1. 提供的显式路径
2. 环境变量 `LAZYGOPHERS_CONFIG`
3. 当前目录（`conf.*` 或 `config.*`）
4. 可执行文件目录（`conf.*` 或 `config.*`）

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
    log.Fatalf("加载配置失败: %v", err)
}
```

---

### LoadConfigSkipValidate()

加载配置文件但不验证。

```go
func LoadConfigSkipValidate(c any, paths ...string) error
```

**参数：**
- `c` - 配置结构体指针
- `paths` - 可选的文件路径列表

**返回值：**
- 如果加载失败，返回错误

**示例：**
```go
var cfg Config
if err := config.LoadConfigSkipValidate(&cfg, "config.yaml"); err != nil {
    log.Fatalf("加载配置失败: %v", err)
}
```

---

### SetConfig()

保存配置到文件。

```go
func SetConfig(c any) error
```

**参数：**
- `c` - 要保存的配置结构体

**返回值：**
- 如果保存失败，返回错误

**示例：**
```go
cfg := Config{
    Host: "localhost",
    Port: 8080,
    Debug: true,
}

if err := config.SetConfig(&cfg); err != nil {
    log.Fatalf("保存配置失败: %v", err)
}
```

---

### RegisterParser()

为文件扩展名注册自定义解析器。

```go
func RegisterParser(ext string, m Marshaler, u Unmarshaler)
```

**参数：**
- `ext` - 文件扩展名（例如 ".custom"）
- `m` - 编码器函数
- `u` - 解码器函数

**示例：**
```go
config.RegisterParser(".custom",
    func(writer io.Writer, v interface{}) error {
        // 自定义编码逻辑
        return nil
    },
    func(reader io.Reader, v interface{}) error {
        // 自定义解码逻辑
        return nil
    },
)
```

---

## 使用模式

### 多格式支持

```go
type Config struct {
    Host string `json:"host" toml:"host" yaml:"host" ini:"host"`
    Port int    `json:"port" toml:"port" yaml:"port" ini:"port"`
}

// 支持任何格式
var cfg Config
config.LoadConfig(&cfg, "config.json")   // JSON
config.LoadConfig(&cfg, "config.yaml")   // YAML
config.LoadConfig(&cfg, "config.toml")   // TOML
config.LoadConfig(&cfg, "config.ini")    // INI
```

### 基于环境的加载

```go
func loadConfig() *Config {
    var cfg Config
    
    // 尝试多个路径
    paths := []string{
        "/etc/myapp/config.json",
        os.Getenv("HOME") + "/.myapp/config.json",
        "./config.json",
    }
    
    if err := config.LoadConfig(&cfg, paths...); err != nil {
        log.Warnf("使用默认配置: %v", err)
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

### 配置验证

```go
type Config struct {
    Host     string `json:"host" validate:"required"`
    Port     int    `json:"port" validate:"required,min=1,max=65535"`
    Database string `json:"database" validate:"required"`
}

var cfg Config
if err := config.LoadConfig(&cfg, "config.json"); err != nil {
    log.Fatalf("配置验证失败: %v", err)
}
```

---

## 最佳实践

### 默认值

```go
type Config struct {
    Host     string `json:"host" default:"localhost"`
    Port     int    `json:"port" default:"8080"`
    Debug    bool   `json:"debug" default:"false"`
}

var cfg Config
config.LoadConfigSkipValidate(&cfg, "config.json")
// 如果配置文件中没有，则应用默认值
```

### 错误处理

```go
func loadConfigWithFallback() *Config {
    var cfg Config
    
    if err := config.LoadConfig(&cfg, "config.json"); err != nil {
        log.Warnf("加载配置失败，使用默认值: %v", err)
        
        return &Config{
            Host: "localhost",
            Port: 8080,
            Debug: false,
        }
    }
    
    return &cfg
}
```

### 配置热重载

```go
func watchConfig(path string, cfg *Config) {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatalf("创建监视器失败: %v", err)
    }
    
    watcher.Add(path)
    
    for {
        select {
        case event := <-watcher.Events:
            if event.Op&fsnotify.Write == fsnotify.Write {
                log.Info("配置已更改，正在重新加载...")
                
                if err := config.LoadConfig(cfg, path); err != nil {
                    log.Errorf("重新加载配置失败: %v", err)
                }
            }
        case err := <-watcher.Errors:
            log.Errorf("监视器错误: %v", err)
        }
    }
}
```

---

## 相关文档

- [validator](/zh-CN/modules/validator) - 数据验证
- [defaults](/zh-CN/modules/defaults) - 默认值
- [json](/zh-CN/modules/json) - JSON 处理
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
