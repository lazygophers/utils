---
title: atexit - 優雅關閉
---

# atexit - 優雅關閉

## 概述

atexit 模組通過註冊在應用程式終止時調用的退出處理器來提供優雅關閉功能。

## 函數

### Register()

註冊在退出時調用的回調函數。

```go
func Register(callback func())
```

**參數：**
- `callback` - 在退出時調用的函數

**行為：**
- 註冊回調以在退出時執行
- 在首次註冊時初始化信號處理器
- 回調按註冊順序執行

**示例：**
```go
func main() {
    atexit.Register(cleanupResources)
    atexit.Register(closeConnections)
    atexit.Register(saveState)
    
    // 應用程式代碼
    runApplication()
    
    // 退出處理器將自動調用
}

func cleanupResources() {
    log.Info("清理資源...")
}

func closeConnections() {
    log.Info("關閉連接...")
}

func saveState() {
    log.Info("保存狀態...")
}
```

---

## 使用模式

### 資源清理

```go
func setupDatabase() *sql.DB {
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    
    atexit.Register(func() {
        log.Info("關閉數據庫連接")
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
        log.Info("關閉 HTTP 服務器")
        ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
        defer cancel()
        server.Shutdown(ctx)
    })
    
    go server.ListenAndServe()
    return server
}
```

### 信號處理

atexit 模組自動處理常見的終止信號：
- **類 Unix 系統**: SIGINT, SIGTERM
- **Windows**: 控制台事件

```go
func main() {
    atexit.Register(func() {
        log.Info("收到終止信號")
        gracefulShutdown()
    })
    
    // 應用程式將在 SIGINT/SIGTERM 時優雅退出
    select {}
}
```

### 多個處理器

```go
func main() {
    // 註冊多個清理處理器
    atexit.Register(cleanupDatabase)
    atexit.Register(closeFiles)
    atexit.Register(flushLogs)
    atexit.Register(notifyMonitoring)
    
    // 應用程式代碼
    runApplication()
}

func cleanupDatabase() {
    log.Info("清理數據庫...")
}

func closeFiles() {
    log.Info("關閉打開的文件...")
}

func flushLogs() {
    log.Info("刷新日誌...")
}

func notifyMonitoring() {
    log.Info("通知監控系統...")
}
```

---

## 最佳實踐

### 處理器註冊

```go
// 好：在初始化期間註冊處理器
func init() {
    atexit.Register(cleanupResources)
}

// 好：註冊帶有錯誤恢復的處理器
func registerHandler() {
    atexit.Register(func() {
        defer func() {
            if r := recover(); r != nil {
                log.Errorf("退出處理器中發生 panic: %v", r)
            }
        }()
        
        cleanup()
    })
}
```

### 資源管理

```go
// 好：使用 defer 進行立即清理
func processFile(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()
    
    // 處理文件
    return nil
}

// 好：使用 atexit 進行應用程式級清理
func main() {
    db := setupDatabase()
    server := setupHTTPServer()
    
    atexit.Register(func() {
        db.Close()
        server.Shutdown(context.Background())
    })
    
    // 應用程式代碼
}
```

---

## 相關文檔

- [runtime](/zh-TW/modules/runtime) - 運行時資訊
- [app](/zh-TW/modules/app) - 應用程式框架
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
