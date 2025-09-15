# wait - 超时、重试和速率限制工具

`wait` 包提供了控制并发操作、实现超时、重试机制和速率限制的工具。它包括信号量池、等待组和异步操作管理。

## 功能特性

- **信号量池**: 使用命名池控制并发操作
- **等待组**: 具有超时支持的增强等待组操作
- **异步操作**: 带结果处理的异步任务执行
- **速率限制**: 内置速率限制功能
- **超时管理**: 所有操作的超时支持
- **线程安全**: 所有操作都是 goroutine 安全的

## 安装

```bash
go get github.com/lazygophers/utils/wait
```

## 使用示例

### 信号量池操作

```go
package main

import (
    "fmt"
    "sync"
    "time"
    "github.com/lazygophers/utils/wait"
)

func main() {
    // 创建最多 3 个并发操作的信号量池
    poolName := "api_requests"
    maxConcurrent := 3

    // 启动并发工作者
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()

            // 获取信号量
            wait.Lock(poolName, maxConcurrent)
            defer wait.Unlock(poolName)

            fmt.Printf("工作者 %d 开始\n", id)
            time.Sleep(time.Second) // 模拟工作
            fmt.Printf("工作者 %d 完成\n", id)
        }(i)
    }

    wg.Wait()
}
```

### 带超时的等待组

```go
// 创建带超时的等待组
group := wait.NewGroup()

// 添加任务
for i := 0; i < 5; i++ {
    group.Add(1)
    go func(id int) {
        defer group.Done()

        // 模拟工作
        time.Sleep(time.Duration(id) * time.Second)
        fmt.Printf("任务 %d 完成\n", id)
    }(i)
}

// 带超时等待
timeout := 3 * time.Second
if group.WaitTimeout(timeout) {
    fmt.Println("所有任务在超时内完成")
} else {
    fmt.Println("达到超时时间，某些任务可能仍在运行")
}
```

### 异步操作

```go
// 定义异步操作
operation := func() (interface{}, error) {
    time.Sleep(2 * time.Second)
    return "操作完成", nil
}

// 异步执行
asyncResult := wait.Async(operation)

// 在操作运行时做其他工作
fmt.Println("执行其他工作...")
time.Sleep(1 * time.Second)

// 带超时获取结果
result, err := asyncResult.GetWithTimeout(5 * time.Second)
if err != nil {
    fmt.Printf("错误: %v\n", err)
} else {
    fmt.Printf("结果: %v\n", result)
}
```

### 池管理

```go
// 检查池状态
poolName := "database_connections"
maxConnections := 10

// 创建池
wait.NewPool(poolName, maxConnections)

// 检查当前使用情况
depth := wait.GetPoolDepth(poolName)
fmt.Printf("当前连接: %d/%d\n", depth, maxConnections)

// 锁定多个资源
wait.LockMultiple([]string{"db_pool", "cache_pool"}, []int{5, 3})
defer wait.UnlockMultiple([]string{"db_pool", "cache_pool"})

// 执行需要两种资源的操作
fmt.Println("使用数据库和缓存...")
```

## API 参考

### 信号量池函数

- `Lock(key string, max int)` - 获取信号量（如果需要则创建池）
- `Unlock(key string)` - 释放信号量
- `TryLock(key string, max int) bool` - 尝试获取信号量而不阻塞
- `LockWithTimeout(key string, max int, timeout time.Duration) bool` - 带超时获取

### 池管理

- `NewPool(key string, max int)` - 创建新的信号量池
- `GetPool(key string) *Pool` - 获取现有池
- `GetPoolDepth(key string) int` - 获取当前池使用情况
- `DestroyPool(key string)` - 移除池并释放资源

### 多池操作

- `LockMultiple(keys []string, maxes []int)` - 原子性锁定多个池
- `UnlockMultiple(keys []string)` - 解锁多个池
- `TryLockMultiple(keys []string, maxes []int) bool` - 尝试锁定多个池

### 等待组操作

```go
type Group struct {
    // 内部实现
}

// 方法
func NewGroup() *Group
func (g *Group) Add(delta int)
func (g *Group) Done()
func (g *Group) Wait()
func (g *Group) WaitTimeout(timeout time.Duration) bool
func (g *Group) WaitContext(ctx context.Context) error
```

### 异步操作

```go
type AsyncResult struct {
    // 内部实现
}

// 函数
func Async(fn func() (interface{}, error)) *AsyncResult
func AsyncWithContext(ctx context.Context, fn func() (interface{}, error)) *AsyncResult

// 方法
func (ar *AsyncResult) Get() (interface{}, error)
func (ar *AsyncResult) GetWithTimeout(timeout time.Duration) (interface{}, error)
func (ar *AsyncResult) IsReady() bool
func (ar *AsyncResult) Cancel()
```

### Pool 类型

```go
type Pool struct {
    // 内部基于通道的信号量
}

// 方法
func (p *Pool) Lock()
func (p *Pool) Unlock()
func (p *Pool) TryLock() bool
func (p *Pool) LockWithTimeout(timeout time.Duration) bool
func (p *Pool) Depth() int
func (p *Pool) Cap() int
```

## 高级使用示例

### HTTP 请求速率限制

```go
func makeAPIRequests(urls []string) {
    poolName := "api_rate_limit"
    maxConcurrent := 5 // 最多 5 个并发请求

    var wg sync.WaitGroup
    for _, url := range urls {
        wg.Add(1)
        go func(u string) {
            defer wg.Done()

            // 速率限制请求
            wait.Lock(poolName, maxConcurrent)
            defer wait.Unlock(poolName)

            resp, err := http.Get(u)
            if err != nil {
                fmt.Printf("获取 %s 时出错: %v\n", u, err)
                return
            }
            defer resp.Body.Close()

            fmt.Printf("获取 %s: %d\n", u, resp.StatusCode)
        }(url)
    }

    wg.Wait()
}
```

### 数据库连接池

```go
type DatabaseManager struct {
    poolName string
    maxConns int
}

func NewDatabaseManager(maxConnections int) *DatabaseManager {
    poolName := "database_connections"
    wait.NewPool(poolName, maxConnections)

    return &DatabaseManager{
        poolName: poolName,
        maxConns: maxConnections,
    }
}

func (dm *DatabaseManager) ExecuteQuery(query string) error {
    // 从池中获取连接
    if !wait.LockWithTimeout(dm.poolName, dm.maxConns, 5*time.Second) {
        return fmt.Errorf("获取数据库连接超时")
    }
    defer wait.Unlock(dm.poolName)

    // 执行查询
    fmt.Printf("执行查询: %s\n", query)
    time.Sleep(100 * time.Millisecond) // 模拟查询执行

    return nil
}

func (dm *DatabaseManager) GetStats() (current, max int) {
    return wait.GetPoolDepth(dm.poolName), dm.maxConns
}
```

### 带超时的批处理

```go
func processBatchWithTimeout(items []string, timeout time.Duration) []string {
    results := make([]string, 0, len(items))
    resultsChan := make(chan string, len(items))

    group := wait.NewGroup()

    // 启动工作者
    for _, item := range items {
        group.Add(1)
        go func(data string) {
            defer group.Done()

            // 模拟处理
            processed := fmt.Sprintf("processed_%s", data)
            resultsChan <- processed
        }(item)
    }

    // 等待完成或超时
    done := make(chan bool)
    go func() {
        group.Wait()
        close(done)
    }()

    select {
    case <-done:
        // 全部完成
        close(resultsChan)
        for result := range resultsChan {
            results = append(results, result)
        }
    case <-time.After(timeout):
        // 发生超时
        fmt.Println("批处理超时")
        close(resultsChan)
        for result := range resultsChan {
            results = append(results, result)
        }
    }

    return results
}
```

### 资源管理

```go
type ResourceManager struct {
    pools map[string]int
}

func NewResourceManager() *ResourceManager {
    return &ResourceManager{
        pools: map[string]int{
            "cpu_intensive":    2,
            "memory_intensive": 3,
            "io_operations":    10,
        },
    }
}

func (rm *ResourceManager) ExecuteTask(taskType string, task func()) error {
    maxConcurrent, exists := rm.pools[taskType]
    if !exists {
        return fmt.Errorf("未知任务类型: %s", taskType)
    }

    // 尝试带超时获取资源
    if !wait.LockWithTimeout(taskType, maxConcurrent, 30*time.Second) {
        return fmt.Errorf("获取 %s 资源超时", taskType)
    }
    defer wait.Unlock(taskType)

    // 执行任务
    task()
    return nil
}
```

## 最佳实践

1. **池命名**: 使用指示资源类型的描述性池名称
2. **超时管理**: 对长时间运行的操作始终使用超时
3. **资源清理**: 确保在 defer 语句中调用 `Unlock()`
4. **合理调整大小**: 根据实际资源约束调整池大小
5. **监控**: 监控池使用情况以识别瓶颈

## 性能考虑

- **基于通道的信号量**: 使用缓冲通道实现高效的信号量
- **无锁操作**: 池深度检查在可能的情况下是无锁的
- **内存效率**: 每个池的内存开销最小
- **可扩展性**: 支持数千个并发操作

## 错误处理

该包使用无 panic 设计模式：

```go
// 不会 panic 的安全操作
success := wait.TryLock("mypool", 10)
if !success {
    fmt.Println("无法获取锁")
}

// 基于超时的操作
if wait.LockWithTimeout("mypool", 10, 5*time.Second) {
    defer wait.Unlock("mypool")
    // 执行工作
} else {
    fmt.Println("获取锁超时")
}
```

## 相关包

- `routine` - Goroutine 管理和任务调度
- `event` - 事件驱动编程工具
- `hystrix` - 熔断器模式实现