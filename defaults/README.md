# defaults

默认值管理模块，为 Go 结构体提供灵活的默认值设置功能。

## 特性

- 支持多种默认值设置方式
- 支持嵌套结构体的默认值
- 支持默认值标签
- 支持环境变量作为默认值
- 支持回调函数设置默认值
- 支持默认值验证

## 安装

```bash
go get github.com/lazygophers/utils/defaults
```

## 快速开始

### 基本用法

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/defaults"
)

type Config struct {
    Host     string `default:"localhost"`
    Port     int    `default:"8080"`
    Debug    bool   `default:"false"`
    MaxConn  int    `default:"100"`
    Timeout  int    `default:"30"`
}

func main() {
    // 创建配置实例
    config := &Config{}
    
    // 应用默认值
    defaults.SetDefaults(config)
    
    fmt.Printf("Host: %s\n", config.Host)
    fmt.Printf("Port: %d\n", config.Port)
    fmt.Printf("Debug: %v\n", config.Debug)
}
```

### 使用回调函数

```go
type Database struct {
    URL      string `default:""`
    Username string `default:""`
    Password string `default:"env:DB_PASSWORD"`
}

func (d *Database) SetDefaults() {
    if d.URL == "" {
        d.URL = "mysql://localhost:3306/mydb"
    }
    if d.Username == "" {
        d.Username = "root"
    }
}

func main() {
    db := &Database{}
    defaults.SetDefaults(db)
}
```

## 文档

详细的 API 文档和更多示例，请参考 [GoDoc](https://pkg.go.dev/github.com/lazygophers/utils/defaults)。

## 许可证

本项目采用 AGPL-3.0 许可证。详情请参阅 [LICENSE](../LICENSE) 文件。