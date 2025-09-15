# Event - 事件驱动编程工具

一个轻量级高效的 Go 包，提供同步和异步事件处理功能的事件驱动编程。`event` 包提供了一个简单而强大的事件系统，实现应用程序不同部分之间的解耦通信。

## 特性

- **同步事件**: 在当前 goroutine 中立即执行事件处理器
- **异步事件**: 在独立的 goroutine 中执行事件处理器，实现非阻塞操作
- **多个处理器**: 为同一事件注册多个处理器
- **线程安全**: 并发注册和发射事件
- **恐慌恢复**: 异步处理器内置恐慌恢复
- **零依赖**: 仅依赖其他 LazyGophers utils 包
- **高性能**: 最小开销的高效事件调度
- **全局管理器**: 便于使用的默认全局事件管理器

## 安装

```bash
go get github.com/lazygophers/utils/event
```

## 快速开始

```go
package main

import (
    "fmt"
    "time"

    "github.com/lazygophers/utils/event"
)

func main() {
    // 注册同步事件处理器
    event.Register("user.login", func(args any) {
        user := args.(string)
        fmt.Printf("用户 %s 同步登录\n", user)
    })

    // 注册异步事件处理器
    event.RegisterAsync("user.login", func(args any) {
        user := args.(string)
        fmt.Printf("异步: 正在为 %s 处理登录\n", user)
        time.Sleep(100 * time.Millisecond) // 模拟工作
        fmt.Printf("异步: %s 登录处理完成\n", user)
    })

    // 发射事件
    event.Emit("user.login", "john_doe")

    // 等待异步处理器完成
    time.Sleep(200 * time.Millisecond)
}
```

## API 参考

### 类型

#### `EventHandler func(args any)`

接收事件参数的事件处理器函数类型。

#### `Manager`

处理事件注册和发射的事件管理器。

### 全局函数

这些函数使用默认的全局事件管理器:

#### `Register(eventName string, handler EventHandler)`

为指定事件注册同步事件处理器。

**参数:**
- `eventName string`: 要处理的事件名称
- `handler EventHandler`: 事件发射时要执行的函数

**示例:**
```go
event.Register("order.created", func(args any) {
    order := args.(Order)
    fmt.Printf("订单 %d 已创建\n", order.ID)
})
```

#### `RegisterAsync(eventName string, handler EventHandler)`

为指定事件注册异步事件处理器。

**参数:**
- `eventName string`: 要处理的事件名称
- `handler EventHandler`: 事件发射时异步执行的函数

**示例:**
```go
event.RegisterAsync("email.send", func(args any) {
    email := args.(EmailData)
    sendEmail(email) // 这在独立的 goroutine 中运行
})
```

#### `Emit(eventName string, args any)`

向所有已注册的处理器发射带有指定参数的事件。

**参数:**
- `eventName string`: 要发射的事件名称
- `args any`: 传递给事件处理器的参数

**示例:**
```go
event.Emit("user.registered", UserRegisteredEvent{
    UserID: 123,
    Email:  "user@example.com",
})
```

### Manager 方法

用于创建自定义事件管理器:

#### `NewManager() *Manager`

创建新的事件管理器实例。

**返回值:**
- `*Manager`: 新的事件管理器实例

**示例:**
```go
manager := event.NewManager()
manager.Register("custom.event", handler)
```

#### `(*Manager) Register(eventName string, handler EventHandler)`

在特定管理器实例上注册同步处理器。

#### `(*Manager) RegisterAsync(eventName string, handler EventHandler)`

在特定管理器实例上注册异步处理器。

#### `(*Manager) Emit(eventName string, args any)`

在特定管理器实例上发射事件。

## 使用示例

### 基本事件系统

```go
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/lazygophers/utils/event"
)

type User struct {
    ID    int
    Email string
    Name  string
}

type UserRegisteredEvent struct {
    User      User
    Timestamp time.Time
}

func main() {
    // 为用户注册注册处理器
    event.Register("user.registered", func(args any) {
        data := args.(UserRegisteredEvent)
        fmt.Printf("欢迎 %s！您的账户已创建。\n", data.User.Name)
    })

    event.RegisterAsync("user.registered", func(args any) {
        data := args.(UserRegisteredEvent)
        fmt.Printf("正在向 %s 发送欢迎邮件\n", data.User.Email)
        // 模拟邮件发送
        time.Sleep(50 * time.Millisecond)
        fmt.Printf("欢迎邮件已发送给 %s\n", data.User.Email)
    })

    // 注册用户
    user := User{ID: 1, Email: "john@example.com", Name: "John Doe"}
    event.Emit("user.registered", UserRegisteredEvent{
        User:      user,
        Timestamp: time.Now(),
    })

    time.Sleep(100 * time.Millisecond)
}
```

### 电商订单处理

```go
package main

import (
    "fmt"
    "time"

    "github.com/lazygophers/utils/event"
)

type Order struct {
    ID       int
    UserID   int
    Items    []string
    Total    float64
    Status   string
}

func main() {
    // 同步订单验证
    event.Register("order.created", func(args any) {
        order := args.(Order)
        fmt.Printf("订单 %d 验证通过（总价: $%.2f）\n", order.ID, order.Total)
    })

    // 异步库存更新
    event.RegisterAsync("order.created", func(args any) {
        order := args.(Order)
        fmt.Printf("正在为订单 %d 更新库存\n", order.ID)
        time.Sleep(30 * time.Millisecond) // 模拟数据库更新
        fmt.Printf("订单 %d 库存已更新\n", order.ID)
    })

    // 异步支付处理
    event.RegisterAsync("order.created", func(args any) {
        order := args.(Order)
        fmt.Printf("正在处理订单 %d 的支付\n", order.ID)
        time.Sleep(100 * time.Millisecond) // 模拟支付网关
        fmt.Printf("订单 %d 支付处理完成\n", order.ID)
    })

    // 异步通知
    event.RegisterAsync("order.created", func(args any) {
        order := args.(Order)
        fmt.Printf("正在发送订单 %d 的确认信息\n", order.ID)
        time.Sleep(20 * time.Millisecond) // 模拟邮件/短信
        fmt.Printf("订单 %d 确认信息已发送\n", order.ID)
    })

    // 创建并发射订单
    order := Order{
        ID:     12345,
        UserID: 1,
        Items:  []string{"笔记本电脑", "鼠标"},
        Total:  1299.99,
        Status: "待处理",
    }

    event.Emit("order.created", order)

    // 等待异步处理
    time.Sleep(200 * time.Millisecond)
}
```

### 应用程序生命周期事件

```go
package main

import (
    "fmt"
    "time"

    "github.com/lazygophers/utils/event"
)

func main() {
    // 应用启动事件
    event.Register("app.starting", func(args any) {
        config := args.(map[string]interface{})
        fmt.Printf("应用程序正在启动，配置: %v\n", config)
    })

    event.RegisterAsync("app.starting", func(args any) {
        fmt.Println("正在初始化后台服务...")
        time.Sleep(50 * time.Millisecond)
        fmt.Println("后台服务初始化完成")
    })

    // 应用就绪事件
    event.Register("app.ready", func(args any) {
        port := args.(int)
        fmt.Printf("应用程序在端口 %d 就绪\n", port)
    })

    // 应用关闭事件
    event.Register("app.shutdown", func(args any) {
        reason := args.(string)
        fmt.Printf("应用程序正在关闭: %s\n", reason)
    })

    // 模拟应用生命周期
    event.Emit("app.starting", map[string]interface{}{
        "port":     8080,
        "debug":    true,
        "database": "postgresql",
    })

    time.Sleep(100 * time.Millisecond)

    event.Emit("app.ready", 8080)

    time.Sleep(50 * time.Millisecond)

    event.Emit("app.shutdown", "用户请求")
}
```

### 自定义事件管理器

```go
package main

import (
    "fmt"
    "time"

    "github.com/lazygophers/utils/event"
)

func main() {
    // 为不同模块创建独立的事件管理器
    userManager := event.NewManager()
    orderManager := event.NewManager()

    // 用户事件
    userManager.Register("login", func(args any) {
        username := args.(string)
        fmt.Printf("[用户] %s 已登录\n", username)
    })

    userManager.RegisterAsync("login", func(args any) {
        username := args.(string)
        fmt.Printf("[用户] 正在记录 %s 的登录\n", username)
        time.Sleep(10 * time.Millisecond)
        fmt.Printf("[用户] %s 的登录已记录\n", username)
    })

    // 订单事件
    orderManager.Register("created", func(args any) {
        orderID := args.(int)
        fmt.Printf("[订单] 订单 %d 已创建\n", orderID)
    })

    orderManager.RegisterAsync("created", func(args any) {
        orderID := args.(int)
        fmt.Printf("[订单] 正在处理订单 %d\n", orderID)
        time.Sleep(30 * time.Millisecond)
        fmt.Printf("[订单] 订单 %d 处理完成\n", orderID)
    })

    // 向不同管理器发射事件
    userManager.Emit("login", "john_doe")
    orderManager.Emit("created", 12345)

    time.Sleep(100 * time.Millisecond)
}
```

### 错误处理和恢复

```go
package main

import (
    "fmt"
    "time"

    "github.com/lazygophers/utils/event"
)

func main() {
    // 可能恐慌的处理器（异步处理器有恐慌恢复）
    event.RegisterAsync("risky.operation", func(args any) {
        data := args.(string)
        if data == "bad_data" {
            panic("出了问题!")
        }
        fmt.Printf("处理完成: %s\n", data)
    })

    // 安全处理器
    event.Register("risky.operation", func(args any) {
        data := args.(string)
        fmt.Printf("同步处理器处理: %s\n", data)
    })

    // 发射正常数据
    fmt.Println("发射正常数据:")
    event.Emit("risky.operation", "good_data")

    time.Sleep(50 * time.Millisecond)

    // 发射异常数据（异步处理器会恐慌但被恢复）
    fmt.Println("发射异常数据:")
    event.Emit("risky.operation", "bad_data")

    time.Sleep(50 * time.Millisecond)

    fmt.Println("尽管异步处理器发生恐慌，应用程序继续运行")
}
```

### 多种事件类型

```go
package main

import (
    "fmt"
    "time"

    "github.com/lazygophers/utils/event"
)

type BlogPost struct {
    ID     int
    Title  string
    Author string
}

type Comment struct {
    ID     int
    PostID int
    Author string
    Text   string
}

func main() {
    // 博客文章事件
    event.Register("blog.post.created", func(args any) {
        post := args.(BlogPost)
        fmt.Printf("新博客文章: '%s' 作者 %s\n", post.Title, post.Author)
    })

    event.RegisterAsync("blog.post.created", func(args any) {
        post := args.(BlogPost)
        fmt.Printf("正在为搜索索引文章 %d\n", post.ID)
        time.Sleep(20 * time.Millisecond)
        fmt.Printf("文章 %d 已索引\n", post.ID)
    })

    // 评论事件
    event.Register("blog.comment.added", func(args any) {
        comment := args.(Comment)
        fmt.Printf("%s 在文章 %d 上发表了新评论\n", comment.Author, comment.PostID)
    })

    event.RegisterAsync("blog.comment.added", func(args any) {
        comment := args.(Comment)
        fmt.Printf("正在发送评论 %d 的通知\n", comment.ID)
        time.Sleep(15 * time.Millisecond)
        fmt.Printf("评论 %d 的通知已发送\n", comment.ID)
    })

    // 创建博客内容
    post := BlogPost{
        ID:     1,
        Title:  "Go 中的事件驱动架构",
        Author: "Jane Developer",
    }

    comment := Comment{
        ID:     1,
        PostID: 1,
        Author: "John Reader",
        Text:   "很棒的文章！",
    }

    event.Emit("blog.post.created", post)
    time.Sleep(50 * time.Millisecond)

    event.Emit("blog.comment.added", comment)
    time.Sleep(50 * time.Millisecond)
}
```

## 事件命名约定

### 推荐模式

1. **点符号**: 使用点来创建分层事件名称
   ```go
   event.Emit("user.login.success", userData)
   event.Emit("user.login.failed", errorData)
   event.Emit("order.payment.completed", paymentData)
   ```

2. **过去时**: 对表示已完成操作的事件使用过去时
   ```go
   event.Emit("user.created", user)      // ✓ 好
   event.Emit("user.create", user)       // ✗ 避免
   ```

3. **资源-操作模式**: 结构化为 `资源.操作`
   ```go
   event.Emit("email.sent", emailData)
   event.Emit("file.uploaded", fileData)
   event.Emit("cache.cleared", cacheInfo)
   ```

## 性能考虑

### 同步与异步处理器

**同步处理器:**
- 在发射 goroutine 中立即执行
- 阻塞直到完成
- 用于关键验证或立即响应
- 简单操作的低延迟

**异步处理器:**
- 在独立的 goroutine 中执行
- 对发射者非阻塞
- 用于重型处理、I/O 操作或非关键任务
- 并发操作的高吞吐量

### 通道缓冲区大小

默认管理器为异步处理器使用大小为 10 的缓冲通道。对于高吞吐量应用程序，考虑创建具有更大缓冲区的自定义管理器:

```go
// 注意: 通道大小目前在 API 中不可配置
// 这是一个潜在的增强领域
manager := event.NewManager()
```

### 内存使用

- 事件处理器存储在内存中，直到应用程序关闭
- 每个管理器实例维护自己的处理器注册表
- 对于长期运行的应用程序，考虑处理器生命周期管理

## 最佳实践

### 1. 事件设计

**清晰的事件名称:**
```go
// 好
event.Emit("user.password.changed", userID)
event.Emit("order.status.updated", orderUpdate)

// 避免
event.Emit("change", data)
event.Emit("update", info)
```

**结构化事件数据:**
```go
type UserLoginEvent struct {
    UserID    int       `json:"user_id"`
    IP        string    `json:"ip"`
    UserAgent string    `json:"user_agent"`
    Timestamp time.Time `json:"timestamp"`
}

event.Emit("user.login", UserLoginEvent{
    UserID:    123,
    IP:        "192.168.1.1",
    UserAgent: "Mozilla/5.0...",
    Timestamp: time.Now(),
})
```

### 2. 处理器注册

**早期注册:**
```go
func init() {
    // 在应用程序初始化期间注册处理器
    event.Register("app.started", handleAppStarted)
    event.RegisterAsync("user.created", sendWelcomeEmail)
}
```

**错误处理:**
```go
event.Register("order.created", func(args any) {
    order, ok := args.(Order)
    if !ok {
        log.Printf("无效的订单数据: %v", args)
        return
    }

    if err := validateOrder(order); err != nil {
        log.Printf("订单验证失败: %v", err)
        return
    }

    processOrder(order)
})
```

### 3. 测试

**测试用的模拟处理器:**
```go
func TestUserRegistration(t *testing.T) {
    var emailSent bool

    // 用测试处理器替换异步邮件处理器
    event.RegisterAsync("user.registered", func(args any) {
        emailSent = true
    })

    registerUser("test@example.com")

    // 等待异步处理器
    time.Sleep(10 * time.Millisecond)

    assert.True(t, emailSent)
}
```

## 线程安全

event 包完全线程安全:

- **注册**: 多个 goroutine 可以并发注册处理器
- **发射**: 事件可以从多个 goroutine 安全发射
- **处理器执行**: 异步处理器在独立的 goroutine 中运行，包含恐慌恢复

## 限制

1. **无事件历史**: 默认不存储或记录事件
2. **无处理器优先级**: 处理器按注册顺序执行
3. **无条件处理器**: 所有已注册的处理器都会为每个事件执行
4. **无内置过滤**: 没有基于条件过滤事件的内置方式

## 贡献

欢迎贡献！增强领域:

1. 事件过滤和条件处理器
2. 处理器优先级和排序
3. 事件持久化和重放
4. 指标和监控集成
5. 处理器生命周期管理

## 许可证

此包是 LazyGophers Utils 库的一部分，遵循相同的许可条款。