# routine - Goroutine 管理和任务调度

`routine` 包提供增强的 goroutine 管理，具有自动错误处理、panic 恢复、跟踪上下文传播和生命周期钩子。它简化了并发编程，同时提供更好的可观察性和错误处理。

## 功能特性

- **增强的 Goroutine 启动**: 带自动错误处理的安全 goroutine 创建
- **Panic 恢复**: 自动 panic 恢复和堆栈跟踪日志
- **跟踪上下文传播**: 跨 goroutines 的自动跟踪 ID 传播
- **生命周期钩子**: goroutine 执行的前置/后置钩子
- **错误处理**: 结构化错误处理和日志记录
- **Goroutine 组**: 管理相关 goroutines 组
- **资源管理**: 自动清理和资源管理

## 安装

```bash
go get github.com/lazygophers/utils/routine
```

## 使用示例

### 基本 Goroutine 管理

```go
package main

import (
    "fmt"
    "time"
    "github.com/lazygophers/utils/routine"
)

func main() {
    // 启动带自动错误处理的简单 goroutine
    routine.Go(func() error {
        fmt.Println("来自 goroutine 的问候!")
        time.Sleep(1 * time.Second)
        return nil
    })

    // 启动可能有错误的 goroutine
    routine.Go(func() error {
        fmt.Println("处理数据...")
        return fmt.Errorf("出了点问题")
    })

    time.Sleep(2 * time.Second)
}
```

### Panic 恢复

```go
// 启动带自动 panic 恢复的 goroutine
routine.GoWithRecover(func() error {
    fmt.Println("这可能会 panic...")

    // 这将被捕获并记录
    panic("意外错误")

    return nil
})

// panic 被捕获、记录，不会使程序崩溃
time.Sleep(1 * time.Second)
```

### Goroutine 组

```go
// 创建 goroutine 组
group := routine.NewGroup()

// 向组中添加多个任务
for i := 0; i < 5; i++ {
    taskID := i
    group.Go(func() error {
        fmt.Printf("任务 %d 开始\n", taskID)
        time.Sleep(time.Duration(taskID) * time.Second)
        fmt.Printf("任务 %d 完成\n", taskID)
        return nil
    })
}

// 等待所有 goroutines 完成
err := group.Wait()
if err != nil {
    fmt.Printf("组执行失败: %v\n", err)
}
```

### 自定义生命周期钩子

```go
// 添加自定义前置钩子
routine.AddBeforeRoutine(func(baseGid, currentGid int64) {
    fmt.Printf("从父 %d 启动 goroutine %d\n", baseGid, currentGid)
})

// 添加自定义后置钩子
routine.AddAfterRoutine(func(currentGid int64) {
    fmt.Printf("完成 goroutine %d\n", currentGid)
})

// 启动 goroutine - 钩子将自动调用
routine.Go(func() error {
    fmt.Println("工作中...")
    return nil
})
```

### 后台任务

```go
// 启动后台工作者
routine.StartBackgroundWorker("data-processor", func() error {
    // 持续处理数据
    for {
        err := processData()
        if err != nil {
            return err
        }
        time.Sleep(10 * time.Second)
    }
})

routine.StartBackgroundWorker("health-checker", func() error {
    // 健康检查循环
    for {
        err := performHealthCheck()
        if err != nil {
            return err
        }
        time.Sleep(30 * time.Second)
    }
})

// 优雅地停止所有后台工作者
routine.StopAllBackgroundWorkers()
```

### 资源管理

```go
// 创建带自动资源清理的 routine
routine.GoWithCleanup(
    func() error {
        // 主要工作
        file, err := os.Open("data.txt")
        if err != nil {
            return err
        }

        // 处理文件
        return processFile(file)
    },
    func() {
        // 清理函数 - 总是被调用
        fmt.Println("清理资源...")
    },
)
```

## API 参考

### 核心函数

- `Go(f func() error)` - 启动带错误处理的 goroutine
- `GoWithRecover(f func() error)` - 启动带 panic 恢复的 goroutine
- `GoWithCleanup(work func() error, cleanup func())` - 启动带清理函数的 goroutine
- `GoWithTimeout(f func() error, timeout time.Duration) error` - 启动带超时的 goroutine

### 生命周期钩子

```go
type BeforeRoutine func(baseGid, currentGid int64)
type AfterRoutine func(currentGid int64)

// 函数
func AddBeforeRoutine(f BeforeRoutine)
func AddAfterRoutine(f AfterRoutine)
func RemoveBeforeRoutine(f BeforeRoutine)
func RemoveAfterRoutine(f AfterRoutine)
```

### Goroutine 组

```go
type Group struct {
    // 内部实现
}

// 函数
func NewGroup() *Group
func NewGroupWithLimit(limit int) *Group

// 方法
func (g *Group) Go(f func() error)
func (g *Group) GoWithRecover(f func() error)
func (g *Group) Wait() error
func (g *Group) WaitTimeout(timeout time.Duration) error
func (g *Group) Cancel()
func (g *Group) Size() int
```

### 后台工作者

- `StartBackgroundWorker(name string, f func() error)` - 启动命名后台工作者
- `StopBackgroundWorker(name string)` - 停止特定后台工作者
- `StopAllBackgroundWorkers()` - 停止所有后台工作者
- `GetBackgroundWorkerStatus(name string) WorkerStatus` - 获取工作者状态

### 缓存和上下文

- `GetCache(key string) interface{}` - 获取 goroutine 本地缓存值
- `SetCache(key string, value interface{})` - 设置 goroutine 本地缓存值
- `ClearCache()` - 清除 goroutine 本地缓存
- `GetGoroutineID() int64` - 获取当前 goroutine ID

## 高级使用示例

### 任务处理管道

```go
// 创建处理管道
pipeline := routine.NewPipeline()

// 添加处理阶段
pipeline.AddStage("validate", func(data interface{}) (interface{}, error) {
    // 验证输入数据
    return validateData(data), nil
})

pipeline.AddStage("transform", func(data interface{}) (interface{}, error) {
    // 转换数据
    return transformData(data), nil
})

pipeline.AddStage("save", func(data interface{}) (interface{}, error) {
    // 保存到数据库
    return saveData(data), nil
})

// 通过管道处理数据
data := []interface{}{item1, item2, item3}
results, err := pipeline.Process(data)
if err != nil {
    fmt.Printf("管道错误: %v\n", err)
}
```

### 工作者池

```go
// 创建工作者池
pool := routine.NewWorkerPool(10) // 10 个工作者

// 提交作业
for i := 0; i < 100; i++ {
    jobID := i
    pool.Submit(func() error {
        fmt.Printf("处理作业 %d\n", jobID)
        time.Sleep(100 * time.Millisecond)
        return nil
    })
}

// 等待所有作业完成
pool.Wait()

// 关闭池
pool.Shutdown()
```

### 计划任务

```go
// 创建任务调度器
scheduler := routine.NewScheduler()

// 调度周期性任务
scheduler.SchedulePeriodic("backup", 1*time.Hour, func() error {
    return performBackup()
})

// 调度一次性任务
scheduler.ScheduleOnce("cleanup", 5*time.Minute, func() error {
    return performCleanup()
})

// 调度 cron 风格任务
scheduler.ScheduleCron("report", "0 0 * * *", func() error {
    return generateDailyReport()
})

// 启动调度器
scheduler.Start()

// 停止调度器
defer scheduler.Stop()
```

### 错误处理和监控

```go
// 设置错误处理器
routine.SetErrorHandler(func(err error, gid int64) {
    fmt.Printf("Goroutine %d 错误: %v\n", gid, err)
    // 发送到监控系统
    sendToMonitoring(err, gid)
})

routine.SetPanicHandler(func(panicValue interface{}, stack []byte, gid int64) {
    fmt.Printf("Goroutine %d panic: %v\n", gid, panicValue)
    fmt.Printf("堆栈跟踪:\n%s\n", stack)
    // 发送警报
    sendAlert(panicValue, stack, gid)
})

// 监控 goroutine 统计
stats := routine.GetStats()
fmt.Printf("活跃 goroutines: %d\n", stats.ActiveGoroutines)
fmt.Printf("总启动数: %d\n", stats.TotalLaunched)
fmt.Printf("错误数: %d\n", stats.Errors)
fmt.Printf("Panics: %d\n", stats.Panics)
```

## 最佳实践

1. **使用错误处理**: 总是从 goroutine 函数返回错误而不是 panic
2. **资源清理**: 使用清理函数或 defer 语句进行资源管理
3. **避免阻塞**: 不要在 goroutines 中无限期阻塞而没有超时
4. **监控资源**: 监控 goroutine 数量和资源使用
5. **使用组**: 对应该一起完成的相关任务使用 goroutine 组

## 性能考虑

- **Goroutine 开销**: 每个 goroutine 有 ~8KB 初始堆栈大小
- **上下文切换**: 太多 goroutines 可能导致过度的上下文切换
- **内存使用**: 监控大量 goroutines 的内存使用
- **错误处理**: 与 panics 相比，错误处理增加的开销很小

## 错误处理模式

```go
// 优雅的错误处理
routine.Go(func() error {
    if err := doSomething(); err != nil {
        return fmt.Errorf("doSomething 失败: %w", err)
    }
    return nil
})

// 带重试逻辑
routine.GoWithRetry(func() error {
    return doSomethingThatMightFail()
}, 3, time.Second) // 3 次重试，间隔 1 秒

// 带熔断器
routine.GoWithCircuitBreaker(func() error {
    return callExternalService()
}, "external-service")
```

## 集成示例

### HTTP 服务器

```go
// 在 HTTP 处理器中
func handleRequest(w http.ResponseWriter, r *http.Request) {
    // 异步处理请求
    routine.Go(func() error {
        return processRequestAsync(r)
    })

    w.WriteHeader(http.StatusAccepted)
}
```

### 消息队列消费者

```go
// 消息队列消费者
routine.StartBackgroundWorker("message-consumer", func() error {
    for message := range messageChannel {
        routine.Go(func() error {
            return processMessage(message)
        })
    }
    return nil
})
```

### 数据库操作

```go
// 批量数据库操作
group := routine.NewGroup()

for _, record := range records {
    group.Go(func() error {
        return db.Insert(record)
    })
}

if err := group.Wait(); err != nil {
    // 处理批量操作失败
    return fmt.Errorf("批量插入失败: %w", err)
}
```

## 相关包

- `wait` - 超时、重试和速率限制工具
- `event` - 事件驱动编程工具
- `hystrix` - 熔断器模式实现