---
title: app - 應用程式框架
---

# app - 應用程式框架

## 概述

app 模組提供應用程式框架工具，包括組織名稱、版本和發布類型管理。

## 變量

### Organization

```go
var Organization = "lazygophers"
```

**描述：**
- 應用程式的組織名稱
- 用於配置和緩存目錄

---

### Name

```go
var Name string
```

**描述：**
- 應用程式名稱
- 在構建時設置

---

### Version

```go
var Version string
```

**描述：**
- 應用程式版本
- 在構建時設置

---

### 構建資訊

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

## 發布類型

### 類型定義

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

獲取發布類型字符串。

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

獲取發布類型調試字符串。

```go
func (p ReleaseType) Debug() string
```

**返回值：**
- 發布類型字符串

---

## 環境配置

### 包類型

```go
var PackageType ReleaseType
```

**描述：**
- 當前包類型
- 基於 `APP_ENV` 環境變量設置

**環境變量：**
- `APP_ENV=dev` 或 `development` → Debug
- `APP_ENV=test` 或 `canary` → Test
- `APP_ENV=prod` 或 `release` 或 `production` → Release
- `APP_ENV=alpha` → Alpha
- `APP_ENV=beta` → Beta

**示例：**
```bash
# 設置環境
export APP_ENV=prod
# 在 Go 代碼中
if app.PackageType == app.Release {
    fmt.Println("在生產模式下運行")
}
```

---

## 使用模式

### 應用程式資訊

```go
func printAppInfo() {
    fmt.Printf("應用程式: %s\n", app.Name)
    fmt.Printf("版本: %s\n", app.Version)
    fmt.Printf("組織: %s\n", app.Organization)
    fmt.Printf("發布類型: %s\n", app.PackageType.String())
    fmt.Printf("Go 版本: %s\n", app.GoVersion)
    fmt.Printf("構建日期: %s\n", app.BuildDate)
}
```

### 基於環境的行為

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

### 功能標誌

```go
func isFeatureEnabled(feature string) bool {
    switch app.PackageType {
    case app.Debug, app.Test:
        return true  // 調試/測試中啟用所有功能
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

### 配置路徑

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

## 最佳實踐

### 構建資訊

```go
// 好：使用構建時變量
func getVersion() string {
    if app.Version != "" {
        return app.Version
    }
    return "dev"
}

// 好：顯示構建資訊
func showBuildInfo() {
    fmt.Printf("版本: %s\n", getVersion())
    fmt.Printf("提交: %s\n", app.Commit)
    fmt.Printf("構建日期: %s\n", app.BuildDate)
}
```

### 環境處理

```go
// 好：檢查包類型
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

## 相關文檔

- [runtime](/zh-TW/modules/runtime) - 運行時資訊
- [config](/zh-TW/modules/config) - 配置管理
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
