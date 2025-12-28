---
title: event - 事件系統
---

# event - 事件系統

## 概述

event 模組實現了發布/訂閱模式,用於事件驅動架構,支持同步和異步事件處理。

## 類型

### EventHandler

事件處理函數。

```go
type EventHandler func(args any)
```

---

### Manager

事件管理器,用於處理事件。

```go
type Manager struct {
    eventMux sync.RWMutex
    events   map[string][]*eventItem
    c        chan *emitItem
}
```

---

## 函數

### Register()

註冊事件處理器。

```go
func Register(eventName string, handler EventHandler)
```

**參數:**
- `eventName` - 事件名稱
- `handler` - 事件處理函數

**示例:**
```go
event.Register("user.created", func(args any) {
    user := args.(User)
    log.Infof("User created: %s", user.Name)
})
```

---

### RegisterAsync()

註冊異步事件處理器。

```go
func RegisterAsync(eventName string, handler EventHandler)
```

**參數:**
- `eventName` - 事件名稱
- `handler` - 事件處理函數

**示例:**
```go
event.RegisterAsync("user.created", func(args any) {
    user := args.(User)
    sendWelcomeEmail(user)
})
```

---

### Emit()

發出事件並傳遞參數。

```go
func Emit(eventName string, args any)
```

**參數:**
- `eventName` - 事件名稱
- `args` - 事件參數

**示例:**
```go
user := User{Name: "John", Email: "john@example.com"}
event.Emit("user.created", user)
```

---

## 使用模式

### 事件驅動架構

```go
func setupEventHandlers() {
    // 註冊處理器
    event.Register("user.created", handleUserCreated)
    event.Register("user.updated", handleUserUpdated)
    event.Register("user.deleted", handleUserDeleted)
    
    event.RegisterAsync("user.created", sendWelcomeEmail)
    event.RegisterAsync("user.created", updateAnalytics)
}

func handleUserCreated(args any) {
    user := args.(User)
    log.Infof("User created: %s", user.Name)
    
    // 更新數據庫
    db.Create(&user)
}

func sendWelcomeEmail(args any) {
    user := args.(User)
    emailService.SendWelcome(user.Email)
}

func updateAnalytics(args any) {
    user := args.(User)
    analytics.TrackUserCreated(user.ID)
}
```

### 應用生命週期

```go
func setupApplicationEvents() {
    event.Register("app.start", handleAppStart)
    event.Register("app.stop", handleAppStop)
    event.Register("app.error", handleAppError)
}

func handleAppStart(args any) {
    log.Info("Application started")
    initializeServices()
}

func handleAppStop(args any) {
    log.Info("Application stopping")
    cleanupResources()
}

func handleAppError(args any) {
    err := args.(error)
    log.Errorf("Application error: %v", err)
    alertService.SendAlert(err)
}
```

### 組件通信

```go
func setupComponentEvents() {
    event.Register("data.updated", handleDataUpdate)
    event.Register("data.loaded", handleDataLoad)
}

func handleDataUpdate(args any) {
    data := args.(Data)
    log.Infof("Data updated: %s", data.ID)
    
    // 刷新 UI
    ui.Refresh()
}

func handleDataLoad(args any) {
    data := args.(Data)
    log.Infof("Data loaded: %s", data.ID)
    
    // 更新緩存
    cache.Set(data.ID, data)
}
```

---

## 最佳實踐

### 事件命名

```go
// 好的做法: 使用描述性的事件名稱
event.Register("user.created", handler)
event.Register("user.updated", handler)
event.Register("user.deleted", handler)

// 好的做法: 使用分層命名
event.Register("order.created", handler)
event.Register("order.paid", handler)
event.Register("order.shipped", handler)
```

### 錯誤處理

```go
// 好的做法: 在事件處理器中處理錯誤
event.Register("user.created", func(args any) {
    defer func() {
        if r := recover(); r != nil {
            log.Errorf("Panic in event handler: %v", r)
        }
    }()
    
    user := args.(User)
    if err := db.Create(&user); err != nil {
        log.Errorf("Failed to create user: %v", err)
    }
})
```

---

## 相關文檔

- [routine](/zh-TW/modules/routine) - Goroutine 管理
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
