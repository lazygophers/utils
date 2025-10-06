# atexit - Graceful Shutdown Handler

[English](#english) | [简体中文](#简体中文)

---

## English

### Overview

The `atexit` package provides a cross-platform mechanism to register callback functions that are executed when a program exits. It ensures graceful shutdown handling across all Go-supported operating systems and architectures.

### Platform Support

This package supports **all platforms** listed by `go tool dist list`:

#### Signal-Based Platforms
- **Linux** (including Android) - Uses `gomonkey` to hook `os.Exit`
- **Darwin/macOS** (including iOS) - Handles SIGINT, SIGTERM, SIGHUP, SIGQUIT
- **Windows** - Handles SIGINT, SIGTERM, os.Interrupt
- **BSD Family** (FreeBSD, OpenBSD, NetBSD, DragonFly BSD) - Extended Unix signal handling
- **Solaris** (including illumos) - Standard Unix signals
- **AIX** - Standard Unix signals
- **Other Unix-like systems** - Fallback with SIGINT, SIGTERM

#### Non-Signal Platforms
- **Plan9** - Requires explicit `Exit()` call
- **js/wasm** - Requires explicit `Exit()` call
- **wasip1/wasm** - Requires explicit `Exit()` call

### Features

- ✅ Cross-platform compatibility
- ✅ Thread-safe callback registration
- ✅ Panic recovery in callbacks
- ✅ Callbacks executed in registration order
- ✅ Graceful exit with code 0 after signal handling
- ✅ Zero dependencies (except gomonkey on Linux/Android)
- ✅ Support for all Go architectures (amd64, arm64, 386, arm, ppc64, riscv64, s390x, mips, loong64, etc.)

### Installation

```bash
go get github.com/lazygophers/utils/atexit
```

### Usage

#### Basic Example

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/atexit"
)

func main() {
    // Register cleanup callback
    atexit.Register(func() {
        fmt.Println("Cleaning up resources...")
    })

    // Your application code
    fmt.Println("Application running...")

    // Program exits normally or via signal
}
```

#### Multiple Callbacks

```go
atexit.Register(func() {
    fmt.Println("Closing database connection...")
})

atexit.Register(func() {
    fmt.Println("Flushing logs...")
})

atexit.Register(func() {
    fmt.Println("Sending final metrics...")
})
```

#### For Plan9, js/wasm, wasip1/wasm

On platforms without signal support, use `atexit.Exit()`:

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/atexit"
)

func main() {
    atexit.Register(func() {
        fmt.Println("Cleanup executed")
    })

    // Use atexit.Exit instead of os.Exit
    atexit.Exit(0)  // Executes callbacks before exiting
}
```

### API

#### `Register(callback func())`

Registers a callback function to be called when the program exits.

- **Parameters**: `callback` - Function to execute on exit
- **Returns**: None
- **Notes**:
  - Nil callbacks are ignored
  - Callbacks are executed in registration order
  - Panics in callbacks are recovered

#### `Exit(code int)` (Plan9, js/wasm, wasip1/wasm only)

Executes all registered callbacks and then exits with the given code.

- **Parameters**: `code` - Exit code (0 for success, non-zero for error)
- **Returns**: Does not return
- **Notes**: Only available on non-signal platforms

### Platform-Specific Behavior

| Platform | Signal Handling | Exit Hook | Notes |
|----------|----------------|-----------|-------|
| Linux/Android | ✅ | ✅ gomonkey | Hooks `os.Exit` for comprehensive coverage |
| Darwin/macOS/iOS | ✅ | ❌ | Handles SIGINT, SIGTERM, SIGHUP, SIGQUIT |
| Windows | ✅ | ❌ | Handles SIGINT, SIGTERM, os.Interrupt |
| BSD (FreeBSD, OpenBSD, NetBSD, DragonFly) | ✅ | ❌ | Extended signal support |
| Solaris/illumos | ✅ | ❌ | Standard Unix signals |
| AIX | ✅ | ❌ | Standard Unix signals |
| Plan9 | ❌ | ❌ | Use `atexit.Exit()` |
| js/wasm | ❌ | ❌ | Use `atexit.Exit()` |
| wasip1/wasm | ❌ | ❌ | Use `atexit.Exit()` |

### Architecture Support

Tested and supported on all Go architectures:
- **x86**: 386, amd64
- **ARM**: arm, arm64
- **PowerPC**: ppc64, ppc64le
- **RISC-V**: riscv64
- **MIPS**: mips, mipsle, mips64, mips64le
- **LoongArch**: loong64
- **s390x**: IBM Z architecture
- **wasm**: WebAssembly

### Exit Behavior

**Signal-based platforms** (Linux, macOS, Windows, BSD, etc.):
- When a termination signal is received (SIGINT, SIGTERM, SIGHUP, etc.), all registered callbacks are executed
- After callbacks complete, the program exits with **code 0** (graceful shutdown)
- This ensures clean shutdown is considered successful

**Non-signal platforms** (Plan9, js/wasm, wasip1/wasm):
- Use `atexit.Exit(code)` to execute callbacks before exiting
- The provided exit code is used as-is

### Best Practices

1. **Register Early**: Register callbacks early in `main()` to ensure they're set up before potential exits
2. **Keep Callbacks Short**: Exit callbacks should complete quickly
3. **Handle Errors Gracefully**: Callbacks should handle their own errors
4. **Avoid Blocking**: Don't use blocking operations in callbacks
5. **Platform Awareness**: Use `atexit.Exit()` on Plan9 and WASM platforms
6. **Exit Code**: Signal-triggered shutdowns exit with code 0, indicating graceful termination

### Testing

```bash
# Run tests on current platform
go test -v

# Cross-compile for specific platform
GOOS=linux GOARCH=amd64 go build
GOOS=windows GOARCH=amd64 go build
GOOS=js GOARCH=wasm go build
```

---

## 简体中文

### 概述

`atexit` 包提供了一个跨平台机制，用于注册程序退出时执行的回调函数。它确保在所有 Go 支持的操作系统和架构上实现优雅的关闭处理。

### 平台支持

本包支持 `go tool dist list` 列出的**所有平台**：

#### 基于信号的平台
- **Linux**（包括 Android）- 使用 `gomonkey` 钩住 `os.Exit`
- **Darwin/macOS**（包括 iOS）- 处理 SIGINT、SIGTERM、SIGHUP、SIGQUIT
- **Windows** - 处理 SIGINT、SIGTERM、os.Interrupt
- **BSD 家族**（FreeBSD、OpenBSD、NetBSD、DragonFly BSD）- 扩展 Unix 信号处理
- **Solaris**（包括 illumos）- 标准 Unix 信号
- **AIX** - 标准 Unix 信号
- **其他类 Unix 系统** - 使用 SIGINT、SIGTERM 的后备实现

#### 非信号平台
- **Plan9** - 需要显式调用 `Exit()`
- **js/wasm** - 需要显式调用 `Exit()`
- **wasip1/wasm** - 需要显式调用 `Exit()`

### 特性

- ✅ 跨平台兼容性
- ✅ 线程安全的回调注册
- ✅ 回调中的 panic 恢复
- ✅ 按注册顺序执行回调
- ✅ 信号处理后以退出码 0 优雅退出
- ✅ 零依赖（Linux/Android 上除 gomonkey 外）
- ✅ 支持所有 Go 架构（amd64、arm64、386、arm、ppc64、riscv64、s390x、mips、loong64 等）

### 安装

```bash
go get github.com/lazygophers/utils/atexit
```

### 使用方法

#### 基本示例

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/atexit"
)

func main() {
    // 注册清理回调
    atexit.Register(func() {
        fmt.Println("正在清理资源...")
    })

    // 应用程序代码
    fmt.Println("应用程序正在运行...")

    // 程序正常退出或通过信号退出
}
```

#### 多个回调

```go
atexit.Register(func() {
    fmt.Println("正在关闭数据库连接...")
})

atexit.Register(func() {
    fmt.Println("正在刷新日志...")
})

atexit.Register(func() {
    fmt.Println("正在发送最终指标...")
})
```

#### 对于 Plan9、js/wasm、wasip1/wasm

在不支持信号的平台上，使用 `atexit.Exit()`：

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/atexit"
)

func main() {
    atexit.Register(func() {
        fmt.Println("清理已执行")
    })

    // 使用 atexit.Exit 而不是 os.Exit
    atexit.Exit(0)  // 在退出前执行回调
}
```

### API

#### `Register(callback func())`

注册程序退出时要调用的回调函数。

- **参数**：`callback` - 退出时执行的函数
- **返回值**：无
- **注意事项**：
  - 忽略 nil 回调
  - 按注册顺序执行回调
  - 恢复回调中的 panic

#### `Exit(code int)`（仅 Plan9、js/wasm、wasip1/wasm）

执行所有注册的回调，然后以给定代码退出。

- **参数**：`code` - 退出代码（0 表示成功，非零表示错误）
- **返回值**：不返回
- **注意事项**：仅在非信号平台上可用

### 平台特定行为

| 平台 | 信号处理 | 退出钩子 | 备注 |
|------|---------|---------|------|
| Linux/Android | ✅ | ✅ gomonkey | 钩住 `os.Exit` 以实现全面覆盖 |
| Darwin/macOS/iOS | ✅ | ❌ | 处理 SIGINT、SIGTERM、SIGHUP、SIGQUIT |
| Windows | ✅ | ❌ | 处理 SIGINT、SIGTERM、os.Interrupt |
| BSD（FreeBSD、OpenBSD、NetBSD、DragonFly）| ✅ | ❌ | 扩展信号支持 |
| Solaris/illumos | ✅ | ❌ | 标准 Unix 信号 |
| AIX | ✅ | ❌ | 标准 Unix 信号 |
| Plan9 | ❌ | ❌ | 使用 `atexit.Exit()` |
| js/wasm | ❌ | ❌ | 使用 `atexit.Exit()` |
| wasip1/wasm | ❌ | ❌ | 使用 `atexit.Exit()` |

### 架构支持

在所有 Go 架构上测试并支持：
- **x86**：386、amd64
- **ARM**：arm、arm64
- **PowerPC**：ppc64、ppc64le
- **RISC-V**：riscv64
- **MIPS**：mips、mipsle、mips64、mips64le
- **LoongArch**：loong64
- **s390x**：IBM Z 架构
- **wasm**：WebAssembly

### 退出行为

**基于信号的平台**（Linux、macOS、Windows、BSD 等）：
- 当收到终止信号（SIGINT、SIGTERM、SIGHUP 等）时，执行所有注册的回调
- 回调完成后，程序以**退出码 0** 退出（优雅关闭）
- 这确保了干净的关闭被视为成功

**非信号平台**（Plan9、js/wasm、wasip1/wasm）：
- 使用 `atexit.Exit(code)` 在退出前执行回调
- 使用提供的退出码

### 最佳实践

1. **早期注册**：在 `main()` 中尽早注册回调，以确保在潜在退出之前设置好
2. **保持回调简短**：退出回调应快速完成
3. **优雅处理错误**：回调应处理自己的错误
4. **避免阻塞**：不要在回调中使用阻塞操作
5. **平台意识**：在 Plan9 和 WASM 平台上使用 `atexit.Exit()`
6. **退出码**：信号触发的关闭以退出码 0 退出，表示优雅终止

### 测试

```bash
# 在当前平台上运行测试
go test -v

# 为特定平台交叉编译
GOOS=linux GOARCH=amd64 go build
GOOS=windows GOARCH=amd64 go build
GOOS=js GOARCH=wasm go build
```

### 许可证

本项目是 LazyGophers Utils 库的一部分。
