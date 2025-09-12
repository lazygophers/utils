# config

灵活的配置管理模块，支持多种配置源和动态配置更新。

## 特性

- 多源配置支持（文件、环境变量、命令行参数）
- 支持多种配置格式（JSON、YAML、TOML、INI）
- 配置热重载
- 配置验证和默认值
- 配置项监听和变更通知

## 安装

```bash
go get github.com/lazygophers/utils/config
```

## 快速开始

### 基本用法

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/config"
)

func main() {
    // 创建配置管理器
    cfg := config.New()
    
    // 从文件加载配置
    err := cfg.LoadFile("config.yaml")
    if err != nil {
        panic(err)
    }
    
    // 获取配置值
    serverPort := cfg.GetInt("server.port", 8080)
    debugMode := cfg.GetBool("debug", false)
    
    fmt.Printf("Server will run on port %d\n", serverPort)
    fmt.Printf("Debug mode: %v\n", debugMode)
}
```

### 多格式配置文件支持

#### JSON 格式
```json
{
    "app": {
        "name": "MyApp",
        "port": 8080,
        "debug": true
    },
    "database": {
        "host": "localhost",
        "username": "user"
    }
}
```

#### YAML 格式
```yaml
app:
  name: MyApp
  port: 8080
  debug: true
database:
  host: localhost
  username: user
```

#### TOML 格式
```toml
[app]
name = "MyApp"
port = 8080
debug = true

[database]
host = "localhost"
username = "user"
```

#### INI 格式
```ini
[app]
name = MyApp
port = 8080
debug = true

[database]
host = localhost
username = user
```

### 加载配置示例

```go
type Config struct {
    App struct {
        Name  string `json:"name" yaml:"name" toml:"name" ini:"name"`
        Port  int    `json:"port" yaml:"port" toml:"port" ini:"port"`
        Debug bool   `json:"debug" yaml:"debug" toml:"debug" ini:"debug"`
    } `json:"app" yaml:"app" toml:"app" ini:"app"`
    Database struct {
        Host     string `json:"host" yaml:"host" toml:"host" ini:"host"`
        Username string `json:"username" yaml:"username" toml:"username" ini:"username"`
    } `json:"database" yaml:"database" toml:"database" ini:"database"`
}

func main() {
    var cfg Config
    
    // 支持多种格式：.json, .yaml, .yml, .toml, .ini
    err := config.LoadConfig(&cfg, "config.ini")
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("App: %s running on port %d\n", cfg.App.Name, cfg.App.Port)
}
```

## 文档

详细的 API 文档和更多示例，请参考 [GoDoc](https://pkg.go.dev/github.com/lazygophers/utils/config)。

## 许可证

本项目采用 AGPL-3.0 许可证。详情请参阅 [LICENSE](../LICENSE) 文件。