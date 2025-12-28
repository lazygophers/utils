---
title: wait - 流程控制
---

# wait - 流程控制

## 概述

wait 模块提供流程控制工具,包括信号量池、同步和超时管理。

## 函数

### Lock()

获取指定 key 的锁。

```go
func Lock(key string)
```

**参数:**
- `key` - 锁标识符

**行为:**
- 阻塞直到锁可用
- 如果 key 不存在则 panic

**示例:**
```go
wait.Ready("my-resource", 10)
wait.Lock("my-resource")
defer wait.Unlock("my-resource")

// 临界区
processResource()
```

---

### Unlock()

释放指定 key 的锁。

```go
func Unlock(key string)
```

**参数:**
- `key` - 锁标识符

**行为:**
- 释放锁
- 如果 key 不存在则 panic

**示例:**
```go
wait.Lock("my-resource")
defer wait.Unlock("my-resource")
```

---

### Depth()

获取指定 key 的当前深度。

```go
func Depth(key string) int
```

**参数:**
- `key` - 锁标识符

**返回:**
- 当前深度(已获取锁的数量)

**示例:**
```go
depth := wait.Depth("my-resource")
log.Infof("Current depth: %d", depth)
```

---

### Sync()

使用锁执行逻辑函数。

```go
func Sync(key string, logic func() error) error
```

**参数:**
- `key` - 锁标识符
- `logic` - 要执行的函数

**返回:**
- 逻辑函数的错误

**示例:**
```go
err := wait.Sync("database", func() error {
    return updateDatabase()
})
if err != nil {
    log.Errorf("Database update failed: %v", err)
}
```

---

### Ready()

为指定 key 初始化信号量。

```go
func Ready(key string, max int)
```

**参数:**
- `key` - 锁标识符
- `max` - 最大并发数

**示例:**
```go
wait.Ready("api-requests", 10)
```

---

## 使用模式

### 速率限制

```go
func init() {
    wait.Ready("api-calls", 100)  // 最多 100 个并发调用
}

func makeAPICall() error {
    wait.Lock("api-calls")
    defer wait.Unlock("api-calls")
    
    return callAPI()
}
```

### 资源池

```go
func init() {
    wait.Ready("database-connections", 10)
}

func queryDatabase(query string) (*Result, error) {
    wait.Lock("database-connections")
    defer wait.Unlock("database-connections")
    
    conn := getDatabaseConnection()
    defer releaseDatabaseConnection(conn)
    
    return conn.Query(query)
}
```

### 临界区

```go
func updateSharedResource() error {
    return wait.Sync("shared-resource", func() error {
        return performUpdate()
    })
}
```

### 并发控制

```go
func processItems(items []Item) error {
    wait.Ready("processing", 10)
    
    var wg sync.WaitGroup
    errors := make(chan error, len(items))
    
    for _, item := range items {
        wg.Add(1)
        go func(item Item) {
            defer wg.Done()
            
            wait.Lock("processing")
            defer wait.Unlock("processing")
            
            if err := processItem(item); err != nil {
                errors <- err
            }
        }(item)
    }
    
    wg.Wait()
    close(errors)
    
    for err := range errors {
        if err != nil {
            return err
        }
    }
    
    return nil
}
```

---

## 最佳实践

### 锁管理

```go
// 好的做法: 始终使用 defer 解锁
func safeOperation() error {
    wait.Lock("resource")
    defer wait.Unlock("resource")
    
    return performOperation()
}

// 好的做法: 检查锁深度
func checkConcurrency() int {
    return wait.Depth("resource")
}
```

### 初始化

```go
// 好的做法: 在启动期间初始化信号量
func init() {
    wait.Ready("api", 100)
    wait.Ready("database", 10)
    wait.Ready("cache", 50)
}
```

---

## 相关文档

- [routine](/zh-CN/modules/routine) - Goroutine 管理
- [hystrix](/zh-CN/modules/hystrix) - 熔断器
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
