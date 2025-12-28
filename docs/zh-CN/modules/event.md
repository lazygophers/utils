---
title: event - 事件系统
---

# event - 事件系统

## 概述

event 模块实现了发布/订阅模式,用于事件驱动架构,支持同步和异步事件处理。

## 类型

### EventHandler

事件处理函数。

```go
type EventHandler func(args any)
```

---

### Manager

事件管理器,用于处理事件。

```go
type Manager struct {
    eventMux sync.RWMutex
    events   map[string][]*eventItem
    c        chan *emitItem
}
```

---

## 函数

### Register()

注册事件处理器。

```go
func Register(eventName string, handler EventHandler)
```

**参数:**
- `eventName` - 事件名称
- `handler` - 事件处理函数

**示例:**
```go
event.Register("user.created", func(args any) {
    user := args.(User)
    log.Infof("User created: %s", user.Name)
})
```

---

### RegisterAsync()

注册异步事件处理器。

```go
func RegisterAsync(eventName string, handler EventHandler)
```

**参数:**
- `eventName` - 事件名称
- `handler` - 事件处理函数

**示例:**
```go
event.RegisterAsync("user.created", func(args any) {
    user := args.(User)
    sendWelcomeEmail(user)
})
```

---

### Emit()

发出事件并传递参数。

```go
func Emit(eventName string, args any)
```

**参数:**
- `eventName` - 事件名称
- `args` - 事件参数

**示例:**
```go
user := User{Name: "John", Email: "john@example.com"}
event.Emit("user.created", user)
```

---

## 使用模式

### 事件驱动架构

```go
func setupEventHandlers() {
    // 注册处理器
    event.Register("user.created", handleUserCreated)
    event.Register("user.updated", handleUserUpdated)
    event.Register("user.deleted", handleUserDeleted)
    
    event.RegisterAsync("user.created", sendWelcomeEmail)
    event.RegisterAsync("user.created", updateAnalytics)
}

func handleUserCreated(args any) {
    user := args.(User)
    log.Infof("User created: %s", user.Name)
    
    // 更新数据库
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

### 应用生命周期

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

### 组件通信

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
    
    // 更新缓存
    cache.Set(data.ID, data)
}
```

---

## 最佳实践

### 事件命名

```go
// 好的做法: 使用描述性的事件名称
event.Register("user.created", handler)
event.Register("user.updated", handler)
event.Register("user.deleted", handler)

// 好的做法: 使用分层命名
event.Register("order.created", handler)
event.Register("order.paid", handler)
event.Register("order.shipped", handler)
```

### 错误处理

```go
// 好的做法: 在事件处理器中处理错误
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

## 相关文档

- [routine](/zh-CN/modules/routine) - Goroutine 管理
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
