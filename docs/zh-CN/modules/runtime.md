---
title: runtime - 运行时信息
---

# runtime - 运行时信息

## 概述

runtime 模块为 Go 应用程序提供系统信息、运行时诊断和路径工具。

## 函数

### CachePanic()

捕获 panic 并防止堆栈溢出。

```go
func CachePanic()
```

**行为：**
- 捕获 panic 并防止堆栈溢出
- 将 panic 信息写入 stderr
- 转储堆栈跟踪

---

### CachePanicWithHandle()

使用自定义处理器捕获 panic。

```go
func CachePanicWithHandle(handle func(err interface{}))
```

**参数：**
- `handle` - 自定义 panic 处理器函数

**示例：**
```go
runtime.CachePanicWithHandle(func(err interface{}) {
    log.Errorf("发生 panic: %v", err)
    // 自定义错误处理
})
```

---

### PrintStack()

打印当前堆栈跟踪。

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

获取可执行文件目录。

```go
func ExecDir() string
```

**返回值：**
- 包含可执行文件的目录
- 如果发生错误，返回空字符串

**示例：**
```go
execDir := runtime.ExecDir()
configPath := filepath.Join(execDir, "config.json")
```

---

### ExecFile()

获取可执行文件路径。

```go
func ExecFile() string
```

**返回值：**
- 可执行文件的完整路径
- 如果发生错误，返回空字符串

**示例：**
```go
execFile := runtime.ExecFile()
log.Infof("运行自: %s", execFile)
```

---

### Pwd()

获取当前工作目录。

```go
func Pwd() string
```

**返回值：**
- 当前工作目录
- 如果发生错误，返回空字符串

**示例：**
```go
cwd := runtime.Pwd()
log.Infof("当前目录: %s", cwd)
```

---

### UserHomeDir()

获取用户主目录。

```go
func UserHomeDir() string
```

**返回值：**
- 用户主目录
- 如果发生错误，返回空字符串

**示例：**
```go
homeDir := runtime.UserHomeDir()
configPath := filepath.Join(homeDir, ".myapp", "config.json")
```

---

### UserConfigDir()

获取用户配置目录。

```go
func UserConfigDir() string
```

**返回值：**
- 平台特定的用户配置目录
- 如果发生错误，返回空字符串

**示例：**
```go
configDir := runtime.UserConfigDir()
appConfigDir := filepath.Join(configDir, "myapp")
```

---

### UserCacheDir()

获取用户缓存目录。

```go
func UserCacheDir() string
```

**返回值：**
- 平台特定的用户缓存目录
- 如果发生错误，返回空字符串

**示例：**
```go
cacheDir := runtime.UserCacheDir()
appCacheDir := filepath.Join(cacheDir, "myapp")
```

---

### LazyConfigDir()

获取 lazygophers 配置目录。

```go
func LazyConfigDir() string
```

**返回值：**
- 带有 lazygophers 组织的用户配置目录

**示例：**
```go
lazyConfigDir := runtime.LazyConfigDir()
configPath := filepath.Join(lazyConfigDir, "config.json")
```

---

### LazyCacheDir()

获取 lazygophers 缓存目录。

```go
func LazyCacheDir() string
```

**返回值：**
- 带有 lazygophers 组织的用户缓存目录

**示例：**
```go
lazyCacheDir := runtime.LazyCacheDir()
cachePath := filepath.Join(lazyCacheDir, "cache.db")
```

---

## 使用模式

### 应用程序初始化

```go
func initApp() {
    // 获取可执行文件目录
    execDir := runtime.ExecDir()
    
    // 获取配置路径
    configPath := filepath.Join(execDir, "config.json")
    
    // 加载配置
    var cfg Config
    if err := config.LoadConfig(&cfg, configPath); err != nil {
        log.Fatalf("加载配置失败: %v", err)
    }
    
    // 获取缓存目录
    cacheDir := runtime.LazyCacheDir()
    os.MkdirAll(cacheDir, 0755)
    
    // 初始化应用程序
    app.Init(&cfg, cacheDir)
}
```

### Panic 恢复

```go
func main() {
    defer runtime.CachePanic()
    
    // 应用程序代码
    if err := runApplication(); err != nil {
        log.Fatalf("应用程序错误: %v", err)
    }
}

func runApplication() error {
    // 应用程序逻辑
    return nil
}
```

### 调试信息

```go
func printDebugInfo() {
    log.Infof("可执行文件: %s", runtime.ExecFile())
    log.Infof("目录: %s", runtime.ExecDir())
    log.Infof("工作目录: %s", runtime.Pwd())
    log.Infof("主目录: %s", runtime.UserHomeDir())
    log.Infof("配置目录: %s", runtime.UserConfigDir())
    log.Infof("缓存目录: %s", runtime.UserCacheDir())
}
```

### 自定义 Panic 处理器

```go
func setupPanicHandler() {
    runtime.CachePanicWithHandle(func(err interface{}) {
        log.Errorf("发生 panic: %v", err)
        
        // 发送警报
        sendAlert(fmt.Sprintf("Panic: %v", err))
        
        // 保存堆栈跟踪
        saveStackTrace()
        
        // 优雅关闭
        gracefulShutdown()
    })
}
func sendAlert(message string) {
    // 发送警报到监控系统
}
func saveStackTrace() {
    // 保存堆栈跟踪到文件
    runtime.PrintStack()
}
func gracefulShutdown() {
    // 清理资源
    log.Info("执行优雅关闭...")
}
```

---

## 平台特定路径

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

## 最佳实践

### Panic 处理

```go
// 好：使用 defer 进行 panic 恢复
func safeFunction() {
    defer runtime.CachePanic()
    
    // 可能 panic 的代码
}

// 避免：不处理 panic
func unsafeFunction() {
    // 可能 panic 的代码
}
```

### 路径解析

```go
// 好：使用 runtime 函数获取跨平台路径
func getConfigPath() string {
    execDir := runtime.ExecDir()
    return filepath.Join(execDir, "config.json")
}

// 避免：硬编码路径
func getConfigPathBad() string {
    return "/usr/local/myapp/config.json"  // 不跨平台
}
```

### 调试信息

```go
// 好：在启动时打印调试信息
func main() {
    printDebugInfo()
    
    if err := runApplication(); err != nil {
        log.Fatalf("应用程序错误: %v", err)
    }
}

func printDebugInfo() {
    log.Infof("可执行文件: %s", runtime.ExecFile())
    log.Infof("工作目录: %s", runtime.Pwd())
}
```

---

## 相关文档

- [osx](/zh-CN/modules/osx) - 操作系统操作
- [app](/zh-CN/modules/app) - 应用程序框架
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
