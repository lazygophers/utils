# config - 多格式配置文件加载

`config` 包提供了加载和管理多种格式配置文件的工具。它支持 JSON、YAML、TOML、INI 和 HCL 格式，具有自动格式检测和环境变量替换功能。

## 功能特性

- **多种格式**: 支持 JSON、YAML、TOML、INI 和 HCL 配置文件
- **自动检测**: 基于文件扩展名的自动格式检测
- **环境变量**: 配置值中的环境变量替换
- **灵活加载**: 从文件、字符串或读取器加载
- **类型安全**: 具有结构体映射的强类型
- **验证**: 内置验证支持

## 安装

```bash
go get github.com/lazygophers/utils/config
```

## 使用示例

### 基本配置加载

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/config"
)

type AppConfig struct {
    Server struct {
        Host string `json:"host" yaml:"host"`
        Port int    `json:"port" yaml:"port"`
    } `json:"server" yaml:"server"`
    Database struct {
        URL      string `json:"url" yaml:"url"`
        MaxConns int    `json:"max_conns" yaml:"max_conns"`
    } `json:"database" yaml:"database"`
}

func main() {
    var cfg AppConfig

    // 从 YAML 文件加载
    err := config.LoadFile("config.yaml", &cfg)
    if err != nil {
        panic(err)
    }

    fmt.Printf("服务器: %s:%d\n", cfg.Server.Host, cfg.Server.Port)
}
```

### 多格式支持

```go
// 加载 JSON 配置
err := config.LoadFile("config.json", &cfg)

// 加载 YAML 配置
err := config.LoadFile("config.yaml", &cfg)

// 加载 TOML 配置
err := config.LoadFile("config.toml", &cfg)

// 加载 INI 配置
err := config.LoadFile("config.ini", &cfg)

// 加载 HCL 配置
err := config.LoadFile("config.hcl", &cfg)
```

### 环境变量替换

```go
// 带有环境变量的 config.yaml
// server:
//   host: ${HOST:localhost}
//   port: ${PORT:8080}

var cfg AppConfig
err := config.LoadFile("config.yaml", &cfg)
// 自动替换 HOST 和 PORT 环境变量
// 如果环境变量未设置，则使用默认值
```

### 从字符串加载

```go
yamlContent := `
server:
  host: localhost
  port: 8080
database:
  url: postgres://localhost/mydb
`

var cfg AppConfig
err := config.LoadString(yamlContent, &cfg, config.FormatYAML)
if err != nil {
    panic(err)
}
```

### 带验证的加载

```go
type Config struct {
    Server struct {
        Host string `yaml:"host" validate:"required"`
        Port int    `yaml:"port" validate:"required,min=1,max=65535"`
    } `yaml:"server" validate:"required"`
}

var cfg Config
err := config.LoadAndValidate("config.yaml", &cfg)
if err != nil {
    // 处理验证错误
    panic(err)
}
```

## API 参考

### 类型

#### Format

```go
type Format int

const (
    FormatAuto Format = iota // 基于文件扩展名自动检测
    FormatJSON               // JSON 格式
    FormatYAML               // YAML 格式
    FormatTOML               // TOML 格式
    FormatINI                // INI 格式
    FormatHCL                // HCL 格式
)
```

### 函数

#### 文件加载
- `LoadFile(filename string, v interface{}) error` - 从文件加载配置并自动检测格式
- `LoadFileWithFormat(filename string, v interface{}, format Format) error` - 使用特定格式加载
- `LoadAndValidate(filename string, v interface{}) error` - 加载并验证配置

#### 字符串加载
- `LoadString(content string, v interface{}, format Format) error` - 从字符串内容加载
- `LoadReader(r io.Reader, v interface{}, format Format) error` - 从读取器加载

#### 环境处理
- `ProcessEnvVars(content string) string` - 处理环境变量替换
- `SetEnvPrefix(prefix string)` - 设置环境变量查找前缀

#### 验证
- `Validate(v interface{}) error` - 使用验证器标签验证结构体
- `ValidateWithRules(v interface{}, rules map[string]string) error` - 使用自定义规则验证

### 配置选项

```go
type Options struct {
    EnvPrefix       string            // 环境变量前缀
    EnvSubstitution bool              // 启用环境变量替换
    Validation      bool              // 启用验证
    CaseSensitive   bool              // 大小写敏感字段匹配
    DefaultValues   map[string]interface{} // 默认值
}

// LoadWithOptions 提供对加载行为的完全控制
func LoadWithOptions(filename string, v interface{}, opts Options) error
```

## 环境变量替换

该包支持灵活的环境变量替换：

### 语法
- `${VAR}` - 必需的环境变量（如果未设置则错误）
- `${VAR:default}` - 可选的，带默认值
- `${PREFIX_VAR}` - 带前缀（当设置了 EnvPrefix 时）

### 示例

```yaml
# config.yaml
database:
  host: ${DB_HOST:localhost}
  port: ${DB_PORT:5432}
  user: ${DB_USER}
  password: ${DB_PASSWORD}

server:
  debug: ${DEBUG:false}
  workers: ${WORKERS:4}
```

```go
// 设置环境变量
os.Setenv("DB_HOST", "prod-db.example.com")
os.Setenv("DB_USER", "myapp")
os.Setenv("DB_PASSWORD", "secret")
os.Setenv("DEBUG", "true")

// 加载配置
var cfg Config
config.LoadFile("config.yaml", &cfg)
// 结果：
// cfg.Database.Host = "prod-db.example.com"
// cfg.Database.Port = 5432 (默认值)
// cfg.Database.User = "myapp"
// cfg.Database.Password = "secret"
// cfg.Server.Debug = true
// cfg.Server.Workers = 4 (默认值)
```

## 文件格式示例

### JSON
```json
{
  "server": {
    "host": "localhost",
    "port": 8080
  },
  "database": {
    "url": "postgres://localhost/mydb"
  }
}
```

### YAML
```yaml
server:
  host: localhost
  port: 8080
database:
  url: postgres://localhost/mydb
```

### TOML
```toml
[server]
host = "localhost"
port = 8080

[database]
url = "postgres://localhost/mydb"
```

### INI
```ini
[server]
host = localhost
port = 8080

[database]
url = postgres://localhost/mydb
```

### HCL
```hcl
server {
  host = "localhost"
  port = 8080
}

database {
  url = "postgres://localhost/mydb"
}
```

## 最佳实践

1. **使用环境变量**: 对于机密信息和环境特定值使用环境变量替换
2. **验证配置**: 加载后始终验证配置
3. **提供默认值**: 对可选配置使用默认值
4. **结构化配置**: 将配置组织成逻辑部分
5. **记录格式**: 选择最适合团队偏好的格式

## 错误处理

该包提供详细的错误信息：

```go
err := config.LoadFile("config.yaml", &cfg)
if err != nil {
    switch {
    case config.IsValidationError(err):
        // 处理验证错误
        fmt.Printf("验证失败: %v\n", err)
    case config.IsFormatError(err):
        // 处理格式/解析错误
        fmt.Printf("格式错误: %v\n", err)
    case config.IsFileError(err):
        // 处理文件访问错误
        fmt.Printf("文件错误: %v\n", err)
    default:
        fmt.Printf("未知错误: %v\n", err)
    }
}
```

## 相关包

- `validator` - 结构体验证工具
- `app` - 应用程序生命周期管理
- `defaults` - 结构体默认值处理