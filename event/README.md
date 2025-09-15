# Event - Event-Driven Programming Utilities

A lightweight and efficient Go package for event-driven programming with synchronous and asynchronous event handling capabilities. The `event` package provides a simple yet powerful event system that enables decoupled communication between different parts of your application.

## Features

- **Synchronous Events**: Execute event handlers immediately in the current goroutine
- **Asynchronous Events**: Execute event handlers in separate goroutines for non-blocking operation
- **Multiple Handlers**: Register multiple handlers for the same event
- **Thread-Safe**: Concurrent registration and emission of events
- **Panic Recovery**: Built-in panic recovery for async handlers
- **Zero Dependencies**: Only depends on other LazyGophers utils packages
- **High Performance**: Minimal overhead with efficient event dispatch
- **Global Manager**: Default global event manager for convenience

## Installation

```bash
go get github.com/lazygophers/utils/event
```

## Quick Start

```go
package main

import (
    "fmt"
    "time"

    "github.com/lazygophers/utils/event"
)

func main() {
    // Register synchronous event handler
    event.Register("user.login", func(args any) {
        user := args.(string)
        fmt.Printf("User %s logged in synchronously\n", user)
    })

    // Register asynchronous event handler
    event.RegisterAsync("user.login", func(args any) {
        user := args.(string)
        fmt.Printf("Async: Processing login for %s\n", user)
        time.Sleep(100 * time.Millisecond) // Simulate work
        fmt.Printf("Async: Login processing completed for %s\n", user)
    })

    // Emit event
    event.Emit("user.login", "john_doe")

    // Wait for async handlers to complete
    time.Sleep(200 * time.Millisecond)
}
```

## API Reference

### Types

#### `EventHandler func(args any)`

Function type for event handlers that receive event arguments.

#### `Manager`

Event manager that handles registration and emission of events.

### Global Functions

These functions use the default global event manager:

#### `Register(eventName string, handler EventHandler)`

Registers a synchronous event handler for the specified event.

**Parameters:**
- `eventName string`: Name of the event to handle
- `handler EventHandler`: Function to execute when event is emitted

**Example:**
```go
event.Register("order.created", func(args any) {
    order := args.(Order)
    fmt.Printf("Order %d created\n", order.ID)
})
```

#### `RegisterAsync(eventName string, handler EventHandler)`

Registers an asynchronous event handler for the specified event.

**Parameters:**
- `eventName string`: Name of the event to handle
- `handler EventHandler`: Function to execute asynchronously when event is emitted

**Example:**
```go
event.RegisterAsync("email.send", func(args any) {
    email := args.(EmailData)
    sendEmail(email) // This runs in a separate goroutine
})
```

#### `Emit(eventName string, args any)`

Emits an event with the specified arguments to all registered handlers.

**Parameters:**
- `eventName string`: Name of the event to emit
- `args any`: Arguments to pass to event handlers

**Example:**
```go
event.Emit("user.registered", UserRegisteredEvent{
    UserID: 123,
    Email:  "user@example.com",
})
```

### Manager Methods

For creating custom event managers:

#### `NewManager() *Manager`

Creates a new event manager instance.

**Returns:**
- `*Manager`: New event manager instance

**Example:**
```go
manager := event.NewManager()
manager.Register("custom.event", handler)
```

#### `(*Manager) Register(eventName string, handler EventHandler)`

Registers a synchronous handler on the specific manager instance.

#### `(*Manager) RegisterAsync(eventName string, handler EventHandler)`

Registers an asynchronous handler on the specific manager instance.

#### `(*Manager) Emit(eventName string, args any)`

Emits an event on the specific manager instance.

## Usage Examples

### Basic Event System

```go
package main

import (
    "fmt"
    "log"

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
    // Register handlers for user registration
    event.Register("user.registered", func(args any) {
        data := args.(UserRegisteredEvent)
        fmt.Printf("Welcome %s! Your account has been created.\n", data.User.Name)
    })

    event.RegisterAsync("user.registered", func(args any) {
        data := args.(UserRegisteredEvent)
        fmt.Printf("Sending welcome email to %s\n", data.User.Email)
        // Simulate email sending
        time.Sleep(50 * time.Millisecond)
        fmt.Printf("Welcome email sent to %s\n", data.User.Email)
    })

    // Register user
    user := User{ID: 1, Email: "john@example.com", Name: "John Doe"}
    event.Emit("user.registered", UserRegisteredEvent{
        User:      user,
        Timestamp: time.Now(),
    })

    time.Sleep(100 * time.Millisecond)
}
```

### E-Commerce Order Processing

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
    // Synchronous order validation
    event.Register("order.created", func(args any) {
        order := args.(Order)
        fmt.Printf("Order %d validated (Total: $%.2f)\n", order.ID, order.Total)
    })

    // Asynchronous inventory update
    event.RegisterAsync("order.created", func(args any) {
        order := args.(Order)
        fmt.Printf("Updating inventory for order %d\n", order.ID)
        time.Sleep(30 * time.Millisecond) // Simulate database update
        fmt.Printf("Inventory updated for order %d\n", order.ID)
    })

    // Asynchronous payment processing
    event.RegisterAsync("order.created", func(args any) {
        order := args.(Order)
        fmt.Printf("Processing payment for order %d\n", order.ID)
        time.Sleep(100 * time.Millisecond) // Simulate payment gateway
        fmt.Printf("Payment processed for order %d\n", order.ID)
    })

    // Asynchronous notification
    event.RegisterAsync("order.created", func(args any) {
        order := args.(Order)
        fmt.Printf("Sending order confirmation for order %d\n", order.ID)
        time.Sleep(20 * time.Millisecond) // Simulate email/SMS
        fmt.Printf("Order confirmation sent for order %d\n", order.ID)
    })

    // Create and emit order
    order := Order{
        ID:     12345,
        UserID: 1,
        Items:  []string{"laptop", "mouse"},
        Total:  1299.99,
        Status: "pending",
    }

    event.Emit("order.created", order)

    // Wait for async processing
    time.Sleep(200 * time.Millisecond)
}
```

### Application Lifecycle Events

```go
package main

import (
    "fmt"
    "time"

    "github.com/lazygophers/utils/event"
)

func main() {
    // Application startup events
    event.Register("app.starting", func(args any) {
        config := args.(map[string]interface{})
        fmt.Printf("Application starting with config: %v\n", config)
    })

    event.RegisterAsync("app.starting", func(args any) {
        fmt.Println("Initializing background services...")
        time.Sleep(50 * time.Millisecond)
        fmt.Println("Background services initialized")
    })

    // Application ready events
    event.Register("app.ready", func(args any) {
        port := args.(int)
        fmt.Printf("Application ready on port %d\n", port)
    })

    // Application shutdown events
    event.Register("app.shutdown", func(args any) {
        reason := args.(string)
        fmt.Printf("Application shutting down: %s\n", reason)
    })

    // Simulate application lifecycle
    event.Emit("app.starting", map[string]interface{}{
        "port":     8080,
        "debug":    true,
        "database": "postgresql",
    })

    time.Sleep(100 * time.Millisecond)

    event.Emit("app.ready", 8080)

    time.Sleep(50 * time.Millisecond)

    event.Emit("app.shutdown", "user requested")
}
```

### Custom Event Manager

```go
package main

import (
    "fmt"
    "time"

    "github.com/lazygophers/utils/event"
)

func main() {
    // Create separate event managers for different modules
    userManager := event.NewManager()
    orderManager := event.NewManager()

    // User events
    userManager.Register("login", func(args any) {
        username := args.(string)
        fmt.Printf("[USER] %s logged in\n", username)
    })

    userManager.RegisterAsync("login", func(args any) {
        username := args.(string)
        fmt.Printf("[USER] Recording login for %s\n", username)
        time.Sleep(10 * time.Millisecond)
        fmt.Printf("[USER] Login recorded for %s\n", username)
    })

    // Order events
    orderManager.Register("created", func(args any) {
        orderID := args.(int)
        fmt.Printf("[ORDER] Order %d created\n", orderID)
    })

    orderManager.RegisterAsync("created", func(args any) {
        orderID := args.(int)
        fmt.Printf("[ORDER] Processing order %d\n", orderID)
        time.Sleep(30 * time.Millisecond)
        fmt.Printf("[ORDER] Order %d processed\n", orderID)
    })

    // Emit events to different managers
    userManager.Emit("login", "john_doe")
    orderManager.Emit("created", 12345)

    time.Sleep(100 * time.Millisecond)
}
```

### Error Handling and Recovery

```go
package main

import (
    "fmt"
    "time"

    "github.com/lazygophers/utils/event"
)

func main() {
    // Handler that might panic (async handlers have panic recovery)
    event.RegisterAsync("risky.operation", func(args any) {
        data := args.(string)
        if data == "bad_data" {
            panic("Something went wrong!")
        }
        fmt.Printf("Processed: %s\n", data)
    })

    // Safe handler
    event.Register("risky.operation", func(args any) {
        data := args.(string)
        fmt.Printf("Sync handler processed: %s\n", data)
    })

    // Emit with good data
    fmt.Println("Emitting good data:")
    event.Emit("risky.operation", "good_data")

    time.Sleep(50 * time.Millisecond)

    // Emit with bad data (async handler will panic but be recovered)
    fmt.Println("Emitting bad data:")
    event.Emit("risky.operation", "bad_data")

    time.Sleep(50 * time.Millisecond)

    fmt.Println("Application continues running despite panic in async handler")
}
```

### Multiple Event Types

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
    // Blog post events
    event.Register("blog.post.created", func(args any) {
        post := args.(BlogPost)
        fmt.Printf("New blog post: '%s' by %s\n", post.Title, post.Author)
    })

    event.RegisterAsync("blog.post.created", func(args any) {
        post := args.(BlogPost)
        fmt.Printf("Indexing post %d for search\n", post.ID)
        time.Sleep(20 * time.Millisecond)
        fmt.Printf("Post %d indexed\n", post.ID)
    })

    // Comment events
    event.Register("blog.comment.added", func(args any) {
        comment := args.(Comment)
        fmt.Printf("New comment by %s on post %d\n", comment.Author, comment.PostID)
    })

    event.RegisterAsync("blog.comment.added", func(args any) {
        comment := args.(Comment)
        fmt.Printf("Sending notification for comment %d\n", comment.ID)
        time.Sleep(15 * time.Millisecond)
        fmt.Printf("Notification sent for comment %d\n", comment.ID)
    })

    // Create blog content
    post := BlogPost{
        ID:     1,
        Title:  "Event-Driven Architecture in Go",
        Author: "Jane Developer",
    }

    comment := Comment{
        ID:     1,
        PostID: 1,
        Author: "John Reader",
        Text:   "Great article!",
    }

    event.Emit("blog.post.created", post)
    time.Sleep(50 * time.Millisecond)

    event.Emit("blog.comment.added", comment)
    time.Sleep(50 * time.Millisecond)
}
```

## Event Naming Conventions

### Recommended Patterns

1. **Dot Notation**: Use dots to create hierarchical event names
   ```go
   event.Emit("user.login.success", userData)
   event.Emit("user.login.failed", errorData)
   event.Emit("order.payment.completed", paymentData)
   ```

2. **Past Tense**: Use past tense for events that represent completed actions
   ```go
   event.Emit("user.created", user)      // ✓ Good
   event.Emit("user.create", user)       // ✗ Avoid
   ```

3. **Resource-Action Pattern**: Structure as `resource.action`
   ```go
   event.Emit("email.sent", emailData)
   event.Emit("file.uploaded", fileData)
   event.Emit("cache.cleared", cacheInfo)
   ```

## Performance Considerations

### Synchronous vs Asynchronous Handlers

**Synchronous Handlers:**
- Execute immediately in the emitting goroutine
- Block until completion
- Use for critical validation or immediate responses
- Lower latency for simple operations

**Asynchronous Handlers:**
- Execute in separate goroutines
- Non-blocking for the emitter
- Use for heavy processing, I/O operations, or non-critical tasks
- Higher throughput for concurrent operations

### Channel Buffer Size

The default manager uses a buffered channel with size 10 for async handlers. For high-throughput applications, consider creating custom managers with larger buffers:

```go
// Note: Channel size is currently not configurable in the API
// This is an area for potential enhancement
manager := event.NewManager()
```

### Memory Usage

- Event handlers are stored in memory until the application shuts down
- Each manager instance maintains its own handler registry
- Consider handler lifecycle management for long-running applications

## Best Practices

### 1. Event Design

**Clear Event Names:**
```go
// Good
event.Emit("user.password.changed", userID)
event.Emit("order.status.updated", orderUpdate)

// Avoid
event.Emit("change", data)
event.Emit("update", info)
```

**Structured Event Data:**
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

### 2. Handler Registration

**Register Early:**
```go
func init() {
    // Register handlers during application initialization
    event.Register("app.started", handleAppStarted)
    event.RegisterAsync("user.created", sendWelcomeEmail)
}
```

**Error Handling:**
```go
event.Register("order.created", func(args any) {
    order, ok := args.(Order)
    if !ok {
        log.Printf("Invalid order data: %v", args)
        return
    }

    if err := validateOrder(order); err != nil {
        log.Printf("Order validation failed: %v", err)
        return
    }

    processOrder(order)
})
```

### 3. Testing

**Mock Handlers for Testing:**
```go
func TestUserRegistration(t *testing.T) {
    var emailSent bool

    // Replace async email handler with test handler
    event.RegisterAsync("user.registered", func(args any) {
        emailSent = true
    })

    registerUser("test@example.com")

    // Wait for async handler
    time.Sleep(10 * time.Millisecond)

    assert.True(t, emailSent)
}
```

## Thread Safety

The event package is fully thread-safe:

- **Registration**: Multiple goroutines can register handlers concurrently
- **Emission**: Events can be emitted from multiple goroutines safely
- **Handler Execution**: Async handlers run in separate goroutines with panic recovery

## Limitations

1. **No Event History**: Events are not stored or logged by default
2. **No Handler Prioritization**: Handlers execute in registration order
3. **No Conditional Handlers**: All registered handlers execute for each event
4. **No Built-in Filtering**: No built-in way to filter events based on criteria

## Contributing

Contributions are welcome! Areas for enhancement:

1. Event filtering and conditional handlers
2. Handler prioritization and ordering
3. Event persistence and replay
4. Metrics and monitoring integration
5. Handler lifecycle management

## License

This package is part of the LazyGophers Utils library and follows the same licensing terms.