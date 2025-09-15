# AtExit - 优雅关闭处理

一个跨平台的 Go 包，通过信号拦截和回调注册提供优雅的应用程序关闭处理。`atexit` 包确保您的应用程序能够在终止前执行清理操作。

## 特性

- **跨平台支持**: 为 Linux、macOS、Windows 和通用 Unix 系统提供优化实现
- **信号处理**: 自动拦截常见的终止信号（SIGINT、SIGTERM、SIGHUP、SIGQUIT）
- **回调注册**: 注册多个清理函数在关闭时执行
- **恐慌恢复**: 内置恐慌恢复，防止一个回调影响其他回调
- **线程安全**: 支持并发注册和执行回调
- **零依赖**: 除了 Go 标准库外无外部依赖

## 安装

```bash
go get github.com/lazygophers/utils/atexit
```

## 快速开始

```go
package main

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/lazygophers/utils/atexit"
)

func main() {
    // 注册清理函数
    atexit.Register(func() {
        fmt.Println("正在关闭数据库连接...")
        // db.Close()
    })

    atexit.Register(func() {
        fmt.Println("正在保存应用程序状态...")
        // saveState()
    })

    // 模拟应用程序工作
    fmt.Println("应用程序运行中...")
    time.Sleep(30 * time.Second)
    fmt.Println("应用程序结束")
}
```

## API 参考

### 函数

#### `Register(callback func())`

注册一个回调函数，当应用程序收到终止信号时执行。

**参数:**
- `callback func()`: 在关闭时执行的函数。如果为 nil，则忽略此调用。

**示例:**
```go
atexit.Register(func() {
    log.Println("清理完成")
})
```

**行为:**
- 回调按注册顺序执行
- 每个回调在自己的受保护 goroutine 中运行，包含恐慌恢复
- 信号处理在首次调用 Register() 时初始化
- 支持并发注册的线程安全

## 平台特定行为

### Linux (`atexit_linux.go`)
- 处理信号: `SIGINT`、`SIGTERM`、`SIGQUIT`、`SIGHUP`
- 为 Linux 信号处理进行优化
- 使用 Linux 特定的信号处理优化

### macOS (`atexit_darwin.go`)
- 处理信号: `SIGINT`、`SIGTERM`、`SIGQUIT`、`SIGHUP`
- 支持额外的 Unix 信号
- 可与系统日志集成

### Windows (`atexit_windows.go`)
- 处理 Windows 特定的终止事件
- 控制台控制事件（Ctrl+C、Ctrl+Break）
- 系统关闭事件
- 服务停止请求

### 通用 Unix (`atexit.go`)
- 处理信号: `SIGINT`、`SIGTERM`
- 其他 Unix 系统的回退实现
- 基本信号处理包含恐慌恢复

## 使用示例

### 数据库清理

```go
package main

import (
    "database/sql"
    "log"

    "github.com/lazygophers/utils/atexit"
    _ "github.com/lib/pq"
)

func main() {
    db, err := sql.Open("postgres", "connection_string")
    if err != nil {
        log.Fatal(err)
    }

    // 注册数据库清理
    atexit.Register(func() {
        log.Println("正在关闭数据库连接...")
        if err := db.Close(); err != nil {
            log.Printf("关闭数据库时出错: %v", err)
        }
    })

    // 您的应用程序逻辑
    runApplication(db)
}
```

### HTTP 服务器优雅关闭

```go
package main

import (
    "context"
    "log"
    "net/http"
    "time"

    "github.com/lazygophers/utils/atexit"
)

func main() {
    server := &http.Server{
        Addr:    ":8080",
        Handler: http.DefaultServeMux,
    }

    // 注册服务器关闭
    atexit.Register(func() {
        log.Println("正在关闭 HTTP 服务器...")
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()

        if err := server.Shutdown(ctx); err != nil {
            log.Printf("服务器关闭错误: %v", err)
        }
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World!"))
    })

    log.Println("服务器在 :8080 启动")
    log.Fatal(server.ListenAndServe())
}
```

### 多资源清理

```go
package main

import (
    "log"
    "os"

    "github.com/lazygophers/utils/atexit"
)

func main() {
    // 打开文件
    logFile, err := os.Create("app.log")
    if err != nil {
        log.Fatal(err)
    }

    configFile, err := os.Open("config.json")
    if err != nil {
        log.Fatal(err)
    }

    // 为每个资源注册清理
    atexit.Register(func() {
        log.Println("正在关闭日志文件...")
        logFile.Close()
    })

    atexit.Register(func() {
        log.Println("正在关闭配置文件...")
        configFile.Close()
    })

    atexit.Register(func() {
        log.Println("执行最终清理...")
        os.Remove("temp.lock")
    })

    // 应用程序逻辑
    runApplication()
}
```

## 最佳实践

### 1. 尽早注册
在应用程序生命周期中尽早注册清理回调:

```go
func main() {
    // 在创建资源后立即注册清理
    db := setupDatabase()
    atexit.Register(func() { db.Close() })

    cache := setupCache()
    atexit.Register(func() { cache.Shutdown() })

    // 继续应用程序逻辑
}
```

### 2. 优雅处理错误
清理函数应该优雅地处理错误而不是恐慌:

```go
atexit.Register(func() {
    if err := resource.Close(); err != nil {
        log.Printf("警告: 关闭资源失败: %v", err)
        // 不要恐慌 - 其他回调需要运行
    }
})
```

### 3. 为长时间运行的操作设置超时
为可能长时间运行的清理操作设置超时:

```go
atexit.Register(func() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := server.Shutdown(ctx); err != nil {
        log.Printf("服务器关闭超时: %v", err)
    }
})
```

### 4. 依赖顺序
按反向依赖顺序注册回调（最后的依赖先注册）:

```go
func main() {
    cache := setupCache()
    db := setupDatabase()
    server := setupServer(db, cache)

    // 按依赖的反向顺序注册
    atexit.Register(func() { server.Shutdown() })  // 依赖 db 和 cache
    atexit.Register(func() { cache.Close() })      // 独立
    atexit.Register(func() { db.Close() })         // 独立
}
```

## 信号处理详情

### 支持的信号

| 平台     | SIGINT | SIGTERM | SIGQUIT | SIGHUP | Windows 事件 |
|----------|--------|---------|---------|--------|-------------|
| Linux    | ✓      | ✓       | ✓       | ✓      | -           |
| macOS    | ✓      | ✓       | ✓       | ✓      | -           |
| Windows  | -      | -       | -       | -      | ✓           |
| 通用     | ✓      | ✓       | -       | -      | -           |

### 信号来源

- **SIGINT**: 键盘中断（Ctrl+C）
- **SIGTERM**: 终止请求
- **SIGQUIT**: 键盘退出（Ctrl+\）
- **SIGHUP**: 检测到控制终端断开
- **Windows**: 控制台控制事件，系统关闭

## 性能考虑

- **低开销**: 信号处理仅初始化一次
- **并发安全**: 使用 RWMutex 实现线程安全的回调管理
- **恐慌恢复**: 每个回调在受保护的环境中运行
- **内存高效**: 信号处理的最小内存占用

## 线程安全

atexit 包完全线程安全:

- **注册**: 多个 goroutine 可以安全地并发注册回调
- **执行**: 回调按顺序执行，但每个都在自己的受保护作用域中
- **信号处理**: 信号处理器使用 `sync.Once` 仅初始化一次

## 限制

1. **一次性执行**: 回调仅在每次应用程序关闭时执行一次
2. **无法取消**: 一旦注册，回调无法取消注册
3. **顺序执行**: 回调顺序运行，不并行
4. **平台差异**: 信号处理在不同操作系统之间有所不同

## 贡献

欢迎贡献！请确保:

1. 跨平台兼容性
2. 线程安全
3. 全面测试
4. 文档更新

## 许可证

此包是 LazyGophers Utils 库的一部分，遵循相同的许可条款。