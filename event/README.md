# event

事件驱动编程模块，提供强大而灵活的事件发布订阅机制。

## 特性

- 高性能的事件分发系统
- 支持异步事件处理
- 事件优先级管理
- 事件过滤器
- 支持多种事件传递模式
- 优雅的错误处理机制

## 安装

```bash
go get github.com/lazygophers/utils/event
```

## 快速开始

### 基本事件发布订阅

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/event"
)

type UserCreated struct {
    ID       int
    Username string
    Email    string
}

func main() {
    // 创建事件管理器
    em := event.NewEventManager()
    
    // 订阅事件
    em.Subscribe("user.created", func(e *event.Event) {
        if user, ok := e.Data.(*UserCreated); ok {
            fmt.Printf("新用户创建: %s (%s)\n", user.Username, user.Email)
        }
    })
    
    // 发布事件
    user := &UserCreated{
        ID:       1,
        Username: "johndoe",
        Email:    "john@example.com",
    }
    
    em.Publish("user.created", user)
}
```

### 异步事件处理

```go
func main() {
    em := event.NewEventManager()
    
    // 异步订阅
    em.SubscribeAsync("order.placed", func(e *event.Event) {
        // 处理订单逻辑
        fmt.Printf("处理订单: %+v\n", e.Data)
    })
    
    // 发布事件（异步处理）
    em.Publish("order.placed", &Order{
        ID:     "ORD-001",
        Amount: 99.99,
    })
}
```

### 带优先级的事件

```go
func main() {
    em := event.NewEventManager()
    
    // 高优先级处理器
    em.SubscribeWithPriority("payment.processed", func(e *event.Event) {
        fmt.Println("高优先级：记录日志")
    }, event.HighPriority)
    
    // 低优先级处理器
    em.SubscribeWithPriority("payment.processed", func(e *event.Event) {
        fmt.Println("低优先级：发送通知")
    }, event.LowPriority)
    
    // 发布事件
    em.Publish("payment.processed", &Payment{Amount: 100.0})
}
```

## 文档

详细的 API 文档和更多示例，请参考 [GoDoc](https://pkg.go.dev/github.com/lazygophers/utils/event)。

## 许可证

本项目采用 AGPL-3.0 许可证。详情请参阅 [LICENSE](../LICENSE) 文件。