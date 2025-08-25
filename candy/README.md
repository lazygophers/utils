# candy

优雅的配置管理模块，提供灵活且强大的配置处理能力。

## 特性

- 多格式配置文件支持（JSON、YAML、TOML 等）
- 环境变量集成
- 配置热重载
- 配置验证
- 类型安全的配置访问

## 安装

```bash
go get github.com/lazygophers/utils/candy
```

## 快速开始

### 基本用法

```go
package main

import (
    "github.com/lazygophers/utils/candy"
)

func main() {
    // 加载配置
    cfg := candy.New()
    
    // 从文件加载
    if err := cfg.LoadFromFile("config.yaml"); err != nil {
        panic(err)
    }
    
    // 获取配置值
    port := cfg.GetString("server.port", "8080")
    debug := cfg.GetBool("debug", false)
    
    // 使用配置
    // ...
}
```

### 环境变量支持

```go
// 自动从环境变量读取配置
cfg.SetEnvPrefix("APP") // 读取 APP_ 开头的环境变量
cfg.AutoEnv()           // 自动映射环境变量
```

## 文档

详细的 API 文档和更多示例，请参考 [GoDoc](https://pkg.go.dev/github.com/lazygophers/utils/candy)。

## 许可证

本项目采用 AGPL-3.0 许可证。详情请参阅 [LICENSE](../LICENSE) 文件。