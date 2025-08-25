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

### 多配置源合并

```go
// 创建配置管理器
cfg := config.New()

// 设置默认值
cfg.SetDefault("app.name", "MyApp")
cfg.SetDefault("app.version", "1.0.0")

// 从文件加载
cfg.LoadFile("config.json")

// 从环境变量加载（前缀为 APP_）
cfg.LoadEnv("APP")

// 监听配置变化
cfg.Watch("database.url", func(oldValue, newValue string) {
    fmt.Printf("Database URL changed from %s to %s\n", oldValue, newValue)
})
```

## 文档

详细的 API 文档和更多示例，请参考 [GoDoc](https://pkg.go.dev/github.com/lazygophers/utils/config)。

## 许可证

本项目采用 AGPL-3.0 许可证。详情请参阅 [LICENSE](../LICENSE) 文件。