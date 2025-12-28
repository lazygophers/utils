---
title: runtime - 運行時資訊
---

# runtime - 運行時資訊

## 概述

runtime 模組為 Go 應用程式提供系統資訊、運行時診斷和路徑工具。

## 函數

### CachePanic()

捕獲 panic 並防止堆棧溢出。

```go
func CachePanic()
```

**行為：**
- 捕獲 panic 並防止堆棧溢出
- 將 panic 資訊寫入 stderr
- 轉儲堆棧跟踪

---

### CachePanicWithHandle()

使用自定義處理器捕獲 panic。

```go
func CachePanicWithHandle(handle func(err interface{}))
```

**參數：**
- `handle` - 自定義 panic 處理器函數

**示例：**
```go
runtime.CachePanicWithHandle(func(err interface{}) {
    log.Errorf("發生 panic: %v", err)
    // 自定義錯誤處理
})
```

---

### PrintStack()

打印當前堆棧跟踪。

```go
func PrintStack()
```

**示例：**
```go
func debugFunction() {
    runtime.PrintStack()
}
```

---

### ExecDir()

獲取可執行文件目錄。

```go
func ExecDir() string
```

**返回值：**
- 包含可執行文件的目錄
- 如果發生錯誤，返回空字符串

**示例：**
```go
execDir := runtime.ExecDir()
configPath := filepath.Join(execDir, "config.json")
```

---

### ExecFile()

獲取可執行文件路徑。

```go
func ExecFile() string
```

**返回值：**
- 可執行文件的完整路徑
- 如果發生錯誤，返回空字符串

**示例：**
```go
execFile := runtime.ExecFile()
log.Infof("運行自: %s", execFile)
```

---

### Pwd()

獲取當前工作目錄。

```go
func Pwd() string
```

**返回值：**
- 當前工作目錄
- 如果發生錯誤，返回空字符串

**示例：**
```go
cwd := runtime.Pwd()
log.Infof("當前目錄: %s", cwd)
```

---

### UserHomeDir()

獲取用戶主目錄。

```go
func UserHomeDir() string
```

**返回值：**
- 用戶主目錄
- 如果發生錯誤，返回空字符串

**示例：**
```go
homeDir := runtime.UserHomeDir()
configPath := filepath.Join(homeDir, ".myapp", "config.json")
```

---

### UserConfigDir()

獲取用戶配置目錄。

```go
func UserConfigDir() string
```

**返回值：**
- 平台特定的用戶配置目錄
- 如果發生錯誤，返回空字符串

**示例：**
```go
configDir := runtime.UserConfigDir()
appConfigDir := filepath.Join(configDir, "myapp")
```

---

### UserCacheDir()

獲取用戶緩存目錄。

```go
func UserCacheDir() string
```

**返回值：**
- 平台特定的用戶緩存目錄
- 如果發生錯誤，返回空字符串

**示例：**
```go
cacheDir := runtime.UserCacheDir()
appCacheDir := filepath.Join(cacheDir, "myapp")
```

---

### LazyConfigDir()

獲取 lazygophers 配置目錄。

```go
func LazyConfigDir() string
```

**返回值：**
- 帶有 lazygophers 組織的用戶配置目錄

**示例：**
```go
lazyConfigDir := runtime.LazyConfigDir()
configPath := filepath.Join(lazyConfigDir, "config.json")
```

---

### LazyCacheDir()

獲取 lazygophers 緩存目錄。

```go
func LazyCacheDir() string
```

**返回值：**
- 帶有 lazygophers 組織的用戶緩存目錄

**示例：**
```go
lazyCacheDir := runtime.LazyCacheDir()
cachePath := filepath.Join(lazyCacheDir, "cache.db")
```

---

## 使用模式

### 應用程式初始化

```go
func initApp() {
    // 獲取可執行文件目錄
    execDir := runtime.ExecDir()
    
    // 獲取配置路徑
    configPath := filepath.Join(execDir, "config.json")
    
    // 載入配置
    var cfg Config
    if err := config.LoadConfig(&cfg, configPath); err != nil {
        log.Fatalf("載入配置失敗: %v", err)
    }
    
    // 獲取緩存目錄
    cacheDir := runtime.LazyCacheDir()
    os.MkdirAll(cacheDir, 0755)
    
    // 初始化應用程式
    app.Init(&cfg, cacheDir)
}
```

### Panic 恢復

```go
func main() {
    defer runtime.CachePanic()
    
    // 應用程式代碼
    if err := runApplication(); err != nil {
        log.Fatalf("應用程式錯誤: %v", err)
    }
}

func runApplication() error {
    // 應用程式邏輯
    return nil
}
```

### 調試資訊

```go
func printDebugInfo() {
    log.Infof("可執行文件: %s", runtime.ExecFile())
    log.Infof("目錄: %s", runtime.ExecDir())
    log.Infof("工作目錄: %s", runtime.Pwd())
    log.Infof("主目錄: %s", runtime.UserHomeDir())
    log.Infof("配置目錄: %s", runtime.UserConfigDir())
    log.Infof("緩存目錄: %s", runtime.UserCacheDir())
}
```

### 自定義 Panic 處理器

```go
func setupPanicHandler() {
    runtime.CachePanicWithHandle(func(err interface{}) {
        log.Errorf("發生 panic: %v", err)
        
        // 發送警報
        sendAlert(fmt.Sprintf("Panic: %v", err))
        
        // 保存堆棧跟踪
        saveStackTrace()
        
        // 優雅關閉
        gracefulShutdown()
    })
}
func sendAlert(message string) {
    // 發送警報到監控系統
}
func saveStackTrace() {
    // 保存堆棧跟踪到文件
    runtime.PrintStack()
}
func gracefulShutdown() {
    // 清理資源
    log.Info("執行優雅關閉...")
}
```

---

## 平台特定路徑

### Linux/Unix

```go
UserHomeDir()    // /home/username
UserConfigDir()  // /home/username/.config
UserCacheDir()   // /home/username/.cache
```

### macOS

```go
UserHomeDir()    // /Users/username
UserConfigDir()  // /Users/username/Library/Application Support
UserCacheDir()   // /Users/username/Library/Caches
```

### Windows

```go
UserHomeDir()    // C:\Users\username
UserConfigDir()  // C:\Users\username\AppData\Roaming
UserCacheDir()   // C:\Users\username\AppData\Local
```

---

## 最佳實踐

### Panic 處理

```go
// 好：使用 defer 進行 panic 恢復
func safeFunction() {
    defer runtime.CachePanic()
    
    // 可能 panic 的代碼
}

// 避免：不處理 panic
func unsafeFunction() {
    // 可能 panic 的代碼
}
```

### 路徑解析

```go
// 好：使用 runtime 函數獲取跨平台路徑
func getConfigPath() string {
    execDir := runtime.ExecDir()
    return filepath.Join(execDir, "config.json")
}

// 避免：硬編碼路徑
func getConfigPathBad() string {
    return "/usr/local/myapp/config.json"  // 不跨平台
}
```

### 調試資訊

```go
// 好：在啟動時打印調試資訊
func main() {
    printDebugInfo()
    
    if err := runApplication(); err != nil {
        log.Fatalf("應用程式錯誤: %v", err)
    }
}

func printDebugInfo() {
    log.Infof("可執行文件: %s", runtime.ExecFile())
    log.Infof("工作目錄: %s", runtime.Pwd())
}
```

---

## 相關文檔

- [osx](/zh-TW/modules/osx) - 操作系統操作
- [app](/zh-TW/modules/app) - 應用程式框架
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
