---
title: routine - Goroutine 管理
---

# routine - Goroutine 管理

## 概述

routine 模块提供 goroutine 管理工具,包括工作池、任务调度和 panic 恢复。

## 函数

### Go()

在 goroutine 中执行函数,带有 panic 恢复。

```go
func Go(f func() (err error))
```

**参数:**
- `f` - 要执行的函数

**行为:**
- 在 goroutine 中执行函数
- 如果函数返回错误则记录日志
- 自动管理 trace ID

**示例:**
```go
routine.Go(func() error {
    if err := processData(); err != nil {
        return err
    }
    return nil
})
```

---

### GoWithRecover()

在 goroutine 中执行函数,带有完整的 panic 恢复。

```go
func GoWithRecover(f func() (err error))
```

**参数:**
- `f` - 要执行的函数

**行为:**
- 在 goroutine 中执行函数
- 捕获 panic 并记录堆栈跟踪
- 如果函数返回错误则记录日志

**示例:**
```go
routine.GoWithRecover(func() error {
    // 这将被捕获并记录
    panic("Something went wrong")
    return nil
})
```

---

### GoWithMustSuccess()

在 goroutine 中执行函数,错误时 panic。

```go
func GoWithMustSuccess(f func() (err error))
```

**参数:**
- `f` - 要执行的函数

**行为:**
- 在 goroutine 中执行函数
- 如果函数返回错误则退出进程

**示例:**
```go
routine.GoWithMustSuccess(func() error {
    if err := criticalOperation(); err != nil {
        return err
    }
    return nil
})
// 如果 criticalOperation 失败,进程将退出
```

---

### AddBeforeRoutine()

添加在 goroutine 启动前执行的回调。

```go
func AddBeforeRoutine(f BeforeRoutine)
```

**参数:**
- `f` - 回调函数

**示例:**
```go
routine.AddBeforeRoutine(func(baseGid, currentGid int64) {
    log.Infof("Starting goroutine: %d -> %d", baseGid, currentGid)
})
```

---

### AddAfterRoutine()

添加在 goroutine 完成后执行的回调。

```go
func AddAfterRoutine(f AfterRoutine)
```

**参数:**
- `f` - 回调函数

**示例:**
```go
routine.AddAfterRoutine(func(currentGid int64) {
    log.Infof("Completed goroutine: %d", currentGid)
})
```

---

## 使用模式

### 后台任务

```go
func startBackgroundTasks() {
    routine.Go(func() error {
        ticker := time.NewTicker(time.Minute)
        defer ticker.Stop()
        
        for range ticker.C {
            if err := performMaintenance(); err != nil {
                log.Errorf("Maintenance failed: %v", err)
            }
        }
        return nil
    })
}
```

### 错误处理

```go
func safeAsyncOperation() {
    routine.GoWithRecover(func() error {
        // 这个 panic 将被捕获
        if someCondition {
            panic("Unexpected error")
        }
        return nil
    })
}
```

### 任务调度

```go
func scheduleTask(delay time.Duration, task func()) {
    routine.Go(func() error {
        time.Sleep(delay)
        task()
        return nil
    })
}
```

---

## 最佳实践

### 错误恢复

```go
// 好的做法: 对关键 goroutine 使用 GoWithRecover
routine.GoWithRecover(func() error {
    criticalOperation()
    return nil
})

// 好的做法: 对简单任务使用 Go
routine.Go(func() error {
    simpleOperation()
    return nil
})
```

---

## 相关文档

- [wait](/zh-CN/modules/wait) - 流程控制
- [hystrix](/zh-CN/modules/hystrix) - 熔断器
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
