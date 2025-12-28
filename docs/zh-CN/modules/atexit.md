---
title: atexit - 优雅关闭
---

# atexit - 优雅关闭

## 概述

atexit 模块通过注册在应用程序终止时调用的退出处理器来提供优雅关闭功能。

## 函数

### Register()

注册在退出时调用的回调函数。

```go
func Register(callback func())
```

**参数：**
- `callback` - 在退出时调用的函数

**行为：**
- 注册回调以在退出时执行
- 在首次注册时初始化信号处理器
- 回调按注册顺序执行

**示例：**
```go
func main() {
    atexit.Register(cleanupResources)
    atexit.Register(closeConnections)
    atexit.Register(saveState)
    
    // 应用程序代码
    runApplication()
    
    // 退出处理器将自动调用
}

func cleanupResources() {
    log.Info("清理资源...")
}

func closeConnections() {
    log.Info("关闭连接...")
}

func saveState() {
    log.Info("保存状态...")
}
```

---

## 使用模式

### 资源清理

```go
func setupDatabase() *sql.DB {
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    
    atexit.Register(func() {
        log.Info("关闭数据库连接")
        db.Close()
    })
    
    return db
}

func setupHTTPServer() *http.Server {
    server := &http.Server{
        Addr:    ":8080",
        Handler: router,
    }
    
    atexit.Register(func() {
        log.Info("关闭 HTTP 服务器")
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()
        server.Shutdown(ctx)
    })
    
    go server.ListenAndServe()
    return server
}
```

### 信号处理

atexit 模块自动处理常见的终止信号：
- **类 Unix 系统**: SIGINT, SIGTERM
- **Windows**: 控制台事件

```go
func main() {
    atexit.Register(func() {
        log.Info("收到终止信号")
        gracefulShutdown()
    })
    
    // 应用程序将在 SIGINT/SIGTERM 时优雅退出
    select {}
}
```

### 多个处理器

```go
func main() {
    // 注册多个清理处理器
    atexit.Register(cleanupDatabase)
    atexit.Register(closeFiles)
    atexit.Register(flushLogs)
    atexit.Register(notifyMonitoring)
    
    // 应用程序代码
    runApplication()
}

func cleanupDatabase() {
    log.Info("清理数据库...")
}

func closeFiles() {
    log.Info("关闭打开的文件...")
}

func flushLogs() {
    log.Info("刷新日志...")
}

func notifyMonitoring() {
    log.Info("通知监控系统...")
}
```

---

## 最佳实践

### 处理器注册

```go
// 好：在初始化期间注册处理器
func init() {
    atexit.Register(cleanupResources)
}

// 好：注册带有错误恢复的处理器
func registerHandler() {
    atexit.Register(func() {
        defer func() {
            if r := recover(); r != nil {
                log.Errorf("退出处理器中发生 panic: %v", r)
            }
        }()
        
        cleanup()
    })
}
```

### 资源管理

```go
// 好：使用 defer 进行立即清理
func processFile(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()
    
    // 处理文件
    return nil
}

// 好：使用 atexit 进行应用程序级清理
func main() {
    db := setupDatabase()
    server := setupHTTPServer()
    
    atexit.Register(func() {
        db.Close()
        server.Shutdown(context.Background())
    })
    
    // 应用程序代码
}
```

---

## 相关文档

- [runtime](/zh-CN/modules/runtime) - 运行时信息
- [app](/zh-CN/modules/app) - 应用程序框架
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
