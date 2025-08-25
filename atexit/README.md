# atexit

优雅的退出处理模块，用于在程序退出时执行清理操作。

## 特性

- 注册退出时的回调函数
- 支持多个退出处理器
- 确保资源正确释放

## 安装

```bash
go get github.com/lazygophers/utils/atexit
```

## 快速开始

### 基本用法

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/atexit"
)

func main() {
    // 注册退出处理器
    atexit.Register(func() {
        fmt.Println("清理资源...")
        // 执行清理操作
    })
    
    // 注册多个处理器
    atexit.Register(func() {
        fmt.Println("关闭数据库连接...")
        // 关闭数据库连接
    })
    
    // 程序退出时，所有注册的处理器会按顺序执行
}
```

## 文档

详细的 API 文档和更多示例，请参考 [GoDoc](https://pkg.go.dev/github.com/lazygophers/utils/atexit)。

## 许可证

本项目采用 AGPL-3.0 许可证。详情请参阅 [LICENSE](../LICENSE) 文件。