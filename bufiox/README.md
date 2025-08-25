# bufiox

增强的缓冲区 I/O 操作模块，提供高性能的数据读写功能。

## 特性

- 优化的缓冲区操作
- 高效的读写性能
- 减少系统调用次数
- 内存池管理

## 安装

```bash
go get github.com/lazygophers/utils/bufiox
```

## 快速开始

### 基本用法

```go
package main

import (
    "github.com/lazygophers/utils/bufiox"
)

func main() {
    // 创建缓冲区读取器
    reader := bufiox.NewReader(src)
    
    // 创建缓冲区写入器
    writer := bufiox.NewWriter(dest)
    
    // 使用缓冲区进行高效读写
    // ...
}
```

## 文档

详细的 API 文档和更多示例，请参考 [GoDoc](https://pkg.go.dev/github.com/lazygophers/utils/bufiox)。

## 许可证

本项目采用 AGPL-3.0 许可证。详情请参阅 [LICENSE](../LICENSE) 文件。