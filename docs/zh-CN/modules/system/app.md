---
title: app - 应用程序框架
---

# app - 应用程序框架

## 概述

app 模块提供应用程序框架工具，包括组织名称、版本和发布类型管理。

## 变量

### Organization

```go
var Organization = "lazygophers"
```

**描述：**
- 应用程序的组织名称
- 用于配置和缓存目录

---

### Name

```go
var Name string
```

**描述：**
- 应用程序名称
- 在构建时设置

---

### Version

```go
var Version string
```

**描述：**
- 应用程序版本
- 在构建时设置

---

### 构建信息

```go
var (
    Commit      string
    ShortCommit string
    Branch      string
    Tag         string
    BuildDate string
    GoVersion string
    GoOS    string
    Goarch  string
    Goarm   string
    Goamd64 string
    Gomips  string
    Description string
)
```

---

## 发布类型

### 类型定义

```go
type ReleaseType uint8
const (
    Debug ReleaseType = iota
    Test
    Alpha
    Beta
    Release
)
```

---

### String()

获取发布类型字符串。

```go
func (p ReleaseType) String() string
```

**返回值：**
- "debug"、"test"、"alpha"、"beta" 或 "release"

**示例：**
```go
fmt.Println(app.Debug.String())    // "debug"
fmt.Println(app.Release.String())  // "release"
```

---

### Debug()

获取发布类型调试字符串。

```go
func (p ReleaseType) Debug() string
```

**返回值：**
- 发布类型字符串

---

## 环境配置

### 包类型

```go
var PackageType ReleaseType
```

**描述：**
- 当前包类型
- 基于 `APP_ENV` 环境变量设置

**环境变量：**
- `APP_ENV=dev` 或 `development` → Debug
- `APP_ENV=test` 或 `canary` → Test
- `APP_ENV=prod` 或 `release` 或 `production` → Release
- `APP_ENV=alpha` → Alpha
- `APP_ENV=beta` → Beta

**示例：**
```bash
# 设置环境
export APP_ENV=prod
# 在 Go 代码中
if app.PackageType == app.Release {
    fmt.Println("在生产模式下运行")
}
```

---

## 使用模式

### 应用程序信息

```go
func printAppInfo() {
    fmt.Printf("应用程序: %s\n", app.Name)
    fmt.Printf("版本: %s\n", app.Version)
    fmt.Printf("组织: %s\n", app.Organization)
    fmt.Printf("发布类型: %s\n", app.PackageType.String())
    fmt.Printf("Go 版本: %s\n", app.GoVersion)
    fmt.Printf("构建日期: %s\n", app.BuildDate)
}
```

### 基于环境的行为

```go
func setupLogging() {
    switch app.PackageType {
    case app.Debug:
        log.SetLevel(log.DebugLevel)
    case app.Test:
        log.SetLevel(log.InfoLevel)
    case app.Release:
        log.SetLevel(log.WarnLevel)
    default:
        log.SetLevel(log.InfoLevel)
    }
}
```

### 功能标志

```go
func isFeatureEnabled(feature string) bool {
    switch app.PackageType {
    case app.Debug, app.Test:
        return true  // 调试/测试中启用所有功能
    case app.Alpha:
        return isAlphaFeature(feature)
    case app.Beta:
        return isBetaFeature(feature)
    case app.Release:
        return isReleaseFeature(feature)
    default:
        return false
    }
}
```

### 配置路径

```go
func getConfigPath() string {
    configDir := runtime.LazyConfigDir()
    return filepath.Join(configDir, app.Name, "config.json")
}

func getCachePath() string {
    cacheDir := runtime.LazyCacheDir()
    return filepath.Join(cacheDir, app.Name, "cache.db")
}
```

---

## 最佳实践

### 构建信息

```go
// 好：使用构建时变量
func getVersion() string {
    if app.Version != "" {
        return app.Version
    }
    return "dev"
}

// 好：显示构建信息
func showBuildInfo() {
    fmt.Printf("版本: %s\n", getVersion())
    fmt.Printf("提交: %s\n", app.Commit)
    fmt.Printf("构建日期: %s\n", app.BuildDate)
}
```

### 环境处理

```go
// 好：检查包类型
func initApplication() {
    switch app.PackageType {
    case app.Debug:
        enableDebugFeatures()
    case app.Release:
        enableProductionFeatures()
    default:
        enableStandardFeatures()
    }
}
```

---

## 相关文档

- [runtime](/zh-CN/modules/runtime) - 运行时信息
- [config](/zh-CN/modules/config) - 配置管理
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
