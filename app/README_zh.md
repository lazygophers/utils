# app - 应用程序生命周期管理

`app` 包提供了管理应用程序生命周期、环境检测和构建信息的工具。它帮助应用程序了解其运行时环境和构建上下文。

## 功能特性

- **构建信息**: 访问编译时构建信息（提交、分支、标签、构建日期）
- **环境检测**: 确定应用程序发布类型（调试、测试、alpha、beta、发布）
- **运行时上下文**: 访问 Go 版本和目标平台信息
- **应用程序元数据**: 管理应用程序名称、版本和组织

## 安装

```bash
go get github.com/lazygophers/utils/app
```

## 使用示例

### 基本应用程序信息

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/app"
)

func main() {
    app.Name = "MyApp"
    app.Version = "1.0.0"
    app.PackageType = app.Release

    fmt.Printf("应用程序: %s v%s\n", app.Name, app.Version)
    fmt.Printf("发布类型: %s\n", app.PackageType)
    fmt.Printf("组织: %s\n", app.Organization)
}
```

### 环境检测

```go
// 检查是否在开发环境中运行
if app.IsDebug() {
    fmt.Println("在调试模式下运行")
}

// 检查是否在生产环境中运行
if app.IsRelease() {
    fmt.Println("在生产环境中运行")
}

// 检查是否在运行测试
if app.IsTest() {
    fmt.Println("正在运行测试")
}

// 检查是否为 alpha/beta 版本
if app.IsAlpha() {
    fmt.Println("Alpha 版本")
}

if app.IsBeta() {
    fmt.Println("Beta 版本")
}
```

### 构建信息

```go
// 访问构建时信息
fmt.Printf("Git 提交: %s\n", app.Commit)
fmt.Printf("短提交: %s\n", app.ShortCommit)
fmt.Printf("分支: %s\n", app.Branch)
fmt.Printf("标签: %s\n", app.Tag)
fmt.Printf("构建日期: %s\n", app.BuildDate)
fmt.Printf("Go 版本: %s\n", app.GoVersion)
fmt.Printf("目标操作系统: %s\n", app.GoOS)
fmt.Printf("目标架构: %s\n", app.Goarch)
```

### 发布类型管理

```go
// 设置不同的发布类型
app.PackageType = app.Debug     // 开发
app.PackageType = app.Test      // 测试
app.PackageType = app.Alpha     // Alpha 版本
app.PackageType = app.Beta      // Beta 版本
app.PackageType = app.Release   // 生产版本

// 获取字符串表示
fmt.Println(app.PackageType.String()) // "release", "beta", 等
```

## API 参考

### 变量

#### 应用程序信息
- `Name string` - 应用程序名称
- `Version string` - 应用程序版本
- `Organization string` - 组织名称（默认："lazygophers"）
- `PackageType ReleaseType` - 当前发布类型

#### 构建信息
- `Commit string` - 完整 Git 提交哈希
- `ShortCommit string` - 短 Git 提交哈希
- `Branch string` - Git 分支名称
- `Tag string` - Git 标签名称
- `BuildDate string` - 构建时间戳
- `Description string` - 应用程序描述

#### 运行时信息
- `GoVersion string` - 用于构建的 Go 版本
- `GoOS string` - 目标操作系统
- `Goarch string` - 目标架构
- `Goarm string` - ARM 变体（如果适用）
- `Goamd64 string` - AMD64 变体（如果适用）
- `Gomips string` - MIPS 变体（如果适用）

### 类型

#### ReleaseType

```go
type ReleaseType uint8

const (
    Debug   ReleaseType = iota // 开发/调试
    Test                       // 测试环境
    Alpha                      // Alpha 版本
    Beta                       // Beta 版本
    Release                    // 生产版本
)
```

#### 方法

- `String() string` - 获取发布类型的字符串表示
- `Debug() string` - 获取调试字符串表示

### 函数

#### 环境检测
- `IsDebug() bool` - 检查是否在调试模式下运行
- `IsTest() bool` - 检查是否在测试模式下运行
- `IsAlpha() bool` - 检查是否运行 alpha 版本
- `IsBeta() bool` - 检查是否运行 beta 版本
- `IsRelease() bool` - 检查是否运行生产版本

## 构建集成

要在编译时填充构建信息，使用 Go 的 `-ldflags`：

```bash
go build -ldflags "
    -X github.com/lazygophers/utils/app.Name=MyApp
    -X github.com/lazygophers/utils/app.Version=1.0.0
    -X github.com/lazygophers/utils/app.Commit=$(git rev-parse HEAD)
    -X github.com/lazygophers/utils/app.ShortCommit=$(git rev-parse --short HEAD)
    -X github.com/lazygophers/utils/app.Branch=$(git branch --show-current)
    -X github.com/lazygophers/utils/app.BuildDate=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
"
```

## 使用模式

### 基于环境的配置

```go
func setupLogging() {
    switch app.PackageType {
    case app.Debug:
        // 启用调试日志
        log.SetLevel(log.DebugLevel)
    case app.Release:
        // 生产日志
        log.SetLevel(log.InfoLevel)
    default:
        // 默认日志
        log.SetLevel(log.WarnLevel)
    }
}
```

### 功能标志

```go
func enableExperimentalFeatures() bool {
    return app.IsDebug() || app.IsAlpha()
}

func enableMetrics() bool {
    return app.IsRelease() || app.IsBeta()
}
```

## 最佳实践

1. **早期设置应用程序信息**: 在应用程序的 init 函数中设置 `Name`、`Version` 和 `PackageType`
2. **使用环境检测**: 使用环境检测函数进行条件行为
3. **构建集成**: 将构建信息注入集成到 CI/CD 流水线中
4. **版本管理**: 对 `Version` 字段使用语义版本控制

## 相关包

- `config` - 配置管理
- `runtime` - 运行时系统信息
- `atexit` - 优雅关闭处理