---
title: event - Event System
---

# event - Event System

## Overview

The event module implements a publish/subscribe pattern for event-driven architecture with support for both synchronous and asynchronous event handling.

## Types

### EventHandler

Event handler function.

```go
type EventHandler func(args any)
```

---

### Manager

Event manager for handling events.

```go
type Manager struct {
    eventMux sync.RWMutex
    events   map[string][]*eventItem
    c        chan *emitItem
}
```

---

## Functions

### Register()

Register event handler.

```go
func Register(eventName string, handler EventHandler)
```

**Parameters:**
- `eventName` - Event name
- `handler` - Event handler function

**Example:**
```go
event.Register("user.created", func(args any) {
    user := args.(User)
    log.Infof("User created: %s", user.Name)
})
```

---

### RegisterAsync()

Register asynchronous event handler.

```go
func RegisterAsync(eventName string, handler EventHandler)
```

**Parameters:**
- `eventName` - Event name
- `handler` - Event handler function

**Example:**
```go
event.RegisterAsync("user.created", func(args any) {
    user := args.(User)
    sendWelcomeEmail(user)
})
```

---

### Emit()

Emit event with arguments.

```go
func Emit(eventName string, args any)
```

**Parameters:**
- `eventName` - Event name
- `args` - Event arguments

**Example:**
```go
user := User{Name: "John", Email: "john@example.com"}
event.Emit("user.created", user)
```

---

## Usage Patterns

### Event-Driven Architecture

```go
func setupEventHandlers() {
    // Register handlers
    event.Register("user.created", handleUserCreated)
    event.Register("user.updated", handleUserUpdated)
    event.Register("user.deleted", handleUserDeleted)
    
    event.RegisterAsync("user.created", sendWelcomeEmail)
    event.RegisterAsync("user.created", updateAnalytics)
}

func handleUserCreated(args any) {
    user := args.(User)
    log.Infof("User created: %s", user.Name)
    
    // Update database
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

### Application Lifecycle

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

### Component Communication

```go
func setupComponentEvents() {
    event.Register("data.updated", handleDataUpdate)
    event.Register("data.loaded", handleDataLoad)
}

func handleDataUpdate(args any) {
    data := args.(Data)
    log.Infof("Data updated: %s", data.ID)
    
    // Refresh UI
    ui.Refresh()
}

func handleDataLoad(args any) {
    data := args.(Data)
    log.Infof("Data loaded: %s", data.ID)
    
    // Update cache
    cache.Set(data.ID, data)
}
```

---

## Best Practices

### Event Naming

```go
// Good: Use descriptive event names
event.Register("user.created", handler)
event.Register("user.updated", handler)
event.Register("user.deleted", handler)

// Good: Use hierarchical naming
event.Register("order.created", handler)
event.Register("order.paid", handler)
event.Register("order.shipped", handler)
```

### Error Handling

```go
// Good: Handle errors in event handlers
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

## Related Documentation

- [routine](/en/modules/routine) - Goroutine management
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
