# app

应用启动器模块，用于创建和配置应用程序实例。

## 特性

-   提供统一的应用程序创建接口
-   支持自定义配置和中间件
-   简化应用初始化流程

## 安装

```bash
go get github.com/lazygophers/utils/app
```

## 快速开始

### 基本用法

```go
package main

import (
    "fmt"

    "github.com/lazygophers/utils/app"
)

func main() {
    fmt.Println(app.Version)
}
```

## 文档

详细的 API 文档和更多示例，请参考 [GoDoc](https://pkg.go.dev/github.com/lazygophers/utils/app)。

## 许可证

本项目采用 AGPL-3.0 许可证。详情请参阅 [LICENSE](../LICENSE) 文件。
