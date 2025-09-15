# Hystrix - 熔断器模式实现

一个高性能的 Go 熔断器模式实现，提供多种优化变体。`hystrix` 包通过熔断提供可靠的服务保护，帮助防止级联故障并提高系统弹性。

## 特性

- **多种实现**: 标准、快速和批量熔断器，适用于不同使用场景
- **高性能**: 无锁操作，原子计数器和优化的环形缓冲区
- **可配置状态**: 关闭、开启和半开状态，支持自定义转换
- **时间窗口支持**: 可配置的成功/失败跟踪窗口
- **恐慌恢复**: 内置恐慌恢复，确保稳定运行
- **内存优化**: 缓存对齐结构，防止伪共享
- **并发安全**: 适用于高并发环境的线程安全
- **零分配**: 在关键路径上优化最小内存分配

## 安装

```bash
go get github.com/lazygophers/utils/hystrix
```

## 快速开始

```go
package main

import (
    "fmt"
    "time"
    "errors"

    "github.com/lazygophers/utils/hystrix"
)

func main() {
    // 创建熔断器
    cb := hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
        TimeWindow: 10 * time.Second,
        ReadyToTrip: func(successes, failures uint64) bool {
            total := successes + failures
            return total >= 10 && failures > successes
        },
    })

    // 使用熔断器保护服务调用
    err := cb.Call(func() error {
        // 您的服务调用代码
        return callExternalService()
    })

    if err != nil {
        fmt.Printf("服务调用失败: %v\n", err)
    }

    fmt.Printf("熔断器状态: %s\n", cb.State())
}

func callExternalService() error {
    // 模拟服务调用
    return errors.New("服务不可用")
}
```

## 熔断器状态

### 关闭状态 (Closed)
- **描述**: 正常操作，请求通过
- **行为**: 监控成功/失败率
- **转换**: 失败阈值超过时转换到开启状态

### 开启状态 (Open)
- **描述**: 熔断器"开启"，请求立即被拒绝
- **行为**: 保护下游服务免受额外负载
- **转换**: 冷却期后转换到半开状态

### 半开状态 (Half-Open)
- **描述**: 允许有限请求测试服务恢复
- **行为**: 使用探测函数判断服务是否健康
- **转换**: 成功时转换到关闭状态，失败时返回开启状态

## API 参考

### 标准熔断器

#### `NewCircuitBreaker(config CircuitBreakerConfig) *CircuitBreaker`

创建具有高级功能和优化的新熔断器。

**配置:**
```go
type CircuitBreakerConfig struct {
    TimeWindow    time.Duration // 统计时间窗口
    OnStateChange StateChange   // 状态变化回调
    ReadyToTrip   ReadyToTrip   // 失败条件函数
    Probe         Probe         // 半开状态探测函数
    BufferSize    int          // 请求历史缓冲区大小
}
```

**示例:**
```go
cb := hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
    TimeWindow: 30 * time.Second,
    BufferSize: 1000,
    ReadyToTrip: func(successes, failures uint64) bool {
        total := successes + failures
        failureRate := float64(failures) / float64(total)
        return total >= 20 && failureRate > 0.6
    },
    OnStateChange: func(oldState, newState hystrix.State) {
        log.Printf("熔断器状态变化: %s -> %s", oldState, newState)
    },
})
```

#### 方法

**`Before() bool`**
检查请求是否应该通过熔断器。

**`After(success bool)`**
记录请求执行结果。

**`Call(fn func() error) error`**
使用熔断器保护执行函数。

**`State() State`**
返回当前熔断器状态。

**`Stat() (successes, failures uint64)`**
返回当前成功和失败计数。

### 快速熔断器

#### `NewFastCircuitBreaker(failureThreshold uint64, timeWindow time.Duration) *FastCircuitBreaker`

创建具有最小开销的轻量级熔断器。

**示例:**
```go
cb := hystrix.NewFastCircuitBreaker(5, 10*time.Second)

// 检查是否允许请求
if cb.AllowRequest() {
    err := makeServiceCall()
    cb.RecordResult(err == nil)
}
```

### 批量熔断器

#### `NewBatchCircuitBreaker(config CircuitBreakerConfig, batchSize int, batchTimeout time.Duration) *BatchCircuitBreaker`

创建为批处理场景优化的熔断器。

**示例:**
```go
cb := hystrix.NewBatchCircuitBreaker(
    hystrix.CircuitBreakerConfig{TimeWindow: 30 * time.Second},
    100,                    // 批量大小
    1 * time.Second,       // 批量超时
)

// 批量记录结果以获得更好的性能
cb.AfterBatch(true)  // 成功
cb.AfterBatch(false) // 失败
```

## 使用示例

### 基本服务保护

```go
package main

import (
    "fmt"
    "time"
    "math/rand"

    "github.com/lazygophers/utils/hystrix"
)

type ExternalService struct {
    cb *hystrix.CircuitBreaker
}

func NewExternalService() *ExternalService {
    return &ExternalService{
        cb: hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
            TimeWindow: 30 * time.Second,
            ReadyToTrip: func(successes, failures uint64) bool {
                total := successes + failures
                return total >= 10 && failures*2 > successes
            },
        }),
    }
}

func (s *ExternalService) CallAPI() (string, error) {
    var result string
    err := s.cb.Call(func() error {
        // 模拟 API 调用
        if rand.Float32() < 0.3 { // 30% 失败率
            return fmt.Errorf("API 调用失败")
        }
        result = "成功响应"
        return nil
    })

    if err != nil {
        return "", err
    }
    return result, nil
}

func main() {
    service := NewExternalService()

    for i := 0; i < 50; i++ {
        result, err := service.CallAPI()
        if err != nil {
            fmt.Printf("调用 %d 失败: %v (状态: %s)\n",
                i+1, err, service.cb.State())
        } else {
            fmt.Printf("调用 %d 成功: %s (状态: %s)\n",
                i+1, result, service.cb.State())
        }

        time.Sleep(100 * time.Millisecond)
    }
}
```

### 数据库连接保护

```go
package main

import (
    "database/sql"
    "fmt"
    "time"

    "github.com/lazygophers/utils/hystrix"
)

type Database struct {
    db *sql.DB
    cb *hystrix.CircuitBreaker
}

func NewDatabase(db *sql.DB) *Database {
    return &Database{
        db: db,
        cb: hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
            TimeWindow: 60 * time.Second,
            BufferSize: 500,
            ReadyToTrip: func(successes, failures uint64) bool {
                total := successes + failures
                if total < 5 {
                    return false
                }
                failureRate := float64(failures) / float64(total)
                return failureRate > 0.5
            },
            OnStateChange: func(oldState, newState hystrix.State) {
                fmt.Printf("数据库熔断器: %s -> %s\n",
                    oldState, newState)
            },
        }),
    }
}

func (d *Database) QueryUser(userID int) (*User, error) {
    var user User
    err := d.cb.Call(func() error {
        return d.db.QueryRow(
            "SELECT id, name, email FROM users WHERE id = ?",
            userID,
        ).Scan(&user.ID, &user.Name, &user.Email)
    })

    if err != nil {
        return nil, err
    }
    return &user, nil
}

type User struct {
    ID    int
    Name  string
    Email string
}
```

### HTTP 客户端保护

```go
package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/lazygophers/utils/hystrix"
)

type HTTPClient struct {
    client *http.Client
    cb     *hystrix.CircuitBreaker
}

func NewHTTPClient() *HTTPClient {
    return &HTTPClient{
        client: &http.Client{Timeout: 5 * time.Second},
        cb: hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
            TimeWindow: 30 * time.Second,
            ReadyToTrip: func(successes, failures uint64) bool {
                total := successes + failures
                return total >= 5 && failures*3 > successes
            },
            Probe: hystrix.ProbeWithChance(25), // 25% 概率探测
        }),
    }
}

func (h *HTTPClient) Get(url string) (*http.Response, error) {
    var resp *http.Response
    err := h.cb.Call(func() error {
        var err error
        resp, err = h.client.Get(url)
        if err != nil {
            return err
        }
        if resp.StatusCode >= 500 {
            return fmt.Errorf("服务器错误: %d", resp.StatusCode)
        }
        return nil
    })

    return resp, err
}

func main() {
    client := NewHTTPClient()

    urls := []string{
        "https://httpbin.org/status/200",
        "https://httpbin.org/status/500",
        "https://httpbin.org/delay/6", // 会超时
    }

    for i := 0; i < 20; i++ {
        url := urls[i%len(urls)]
        resp, err := client.Get(url)

        if err != nil {
            fmt.Printf("请求 %d 失败: %v (状态: %s)\n",
                i+1, err, client.cb.State())
        } else {
            fmt.Printf("请求 %d 成功: %d (状态: %s)\n",
                i+1, resp.StatusCode, client.cb.State())
            resp.Body.Close()
        }

        time.Sleep(500 * time.Millisecond)
    }
}
```

### 微服务集成

```go
package main

import (
    "fmt"
    "time"

    "github.com/lazygophers/utils/hystrix"
)

type UserService struct {
    cb *hystrix.CircuitBreaker
}

type OrderService struct {
    userService *UserService
    cb          *hystrix.CircuitBreaker
}

func NewUserService() *UserService {
    return &UserService{
        cb: hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
            TimeWindow: 30 * time.Second,
            ReadyToTrip: func(successes, failures uint64) bool {
                return failures >= 3
            },
        }),
    }
}

func NewOrderService(userService *UserService) *OrderService {
    return &OrderService{
        userService: userService,
        cb: hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
            TimeWindow: 30 * time.Second,
            ReadyToTrip: func(successes, failures uint64) bool {
                return failures >= 5
            },
        }),
    }
}

func (u *UserService) GetUser(userID int) (*User, error) {
    var user User
    err := u.cb.Call(func() error {
        // 模拟用户服务调用
        if userID <= 0 {
            return fmt.Errorf("无效的用户 ID")
        }
        user = User{ID: userID, Name: fmt.Sprintf("用户%d", userID)}
        return nil
    })
    return &user, err
}

func (o *OrderService) CreateOrder(userID int, items []string) (*Order, error) {
    var order Order
    err := o.cb.Call(func() error {
        // 首先获取用户（可能由于用户服务熔断器而失败）
        user, err := o.userService.GetUser(userID)
        if err != nil {
            return fmt.Errorf("获取用户失败: %w", err)
        }

        // 创建订单
        order = Order{
            ID:     time.Now().Unix(),
            UserID: user.ID,
            Items:  items,
            Status: "已创建",
        }
        return nil
    })
    return &order, err
}

type Order struct {
    ID     int64    `json:"id"`
    UserID int      `json:"user_id"`
    Items  []string `json:"items"`
    Status string   `json:"status"`
}

func main() {
    userService := NewUserService()
    orderService := NewOrderService(userService)

    // 模拟订单创建
    for i := 0; i < 10; i++ {
        userID := i%5 - 2 // 一些无效 ID 触发失败

        order, err := orderService.CreateOrder(userID, []string{"商品1", "商品2"})
        if err != nil {
            fmt.Printf("订单创建 %d 失败: %v\n", i+1, err)
            fmt.Printf("  用户服务状态: %s\n", userService.cb.State())
            fmt.Printf("  订单服务状态: %s\n", orderService.cb.State())
        } else {
            fmt.Printf("订单已创建: %+v\n", order)
        }

        time.Sleep(200 * time.Millisecond)
    }
}
```

### 快速熔断器高性能场景

```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"

    "github.com/lazygophers/utils/hystrix"
)

func main() {
    // 高吞吐量场景使用快速熔断器
    cb := hystrix.NewFastCircuitBreaker(100, 10*time.Second)

    var (
        allowed    uint64
        rejected   uint64
        successful uint64
        failed     uint64
    )

    // 模拟高负载
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()

            for j := 0; j < 1000; j++ {
                if cb.AllowRequest() {
                    atomic.AddUint64(&allowed, 1)

                    // 模拟偶尔失败的工作
                    success := j%10 != 0 // 90% 成功率
                    cb.RecordResult(success)

                    if success {
                        atomic.AddUint64(&successful, 1)
                    } else {
                        atomic.AddUint64(&failed, 1)
                    }
                } else {
                    atomic.AddUint64(&rejected, 1)
                }
            }
        }(i)
    }

    wg.Wait()

    fmt.Printf("结果:\n")
    fmt.Printf("  允许:    %d\n", allowed)
    fmt.Printf("  拒绝:   %d\n", rejected)
    fmt.Printf("  成功: %d\n", successful)
    fmt.Printf("  失败:     %d\n", failed)
    fmt.Printf("  总计:      %d\n", allowed+rejected)
}
```

## 高级配置

### 自定义熔断条件函数

```go
cb := hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
    TimeWindow: 60 * time.Second,
    ReadyToTrip: func(successes, failures uint64) bool {
        total := successes + failures

        // 要求最小样本量
        if total < 50 {
            return false
        }

        // 多种条件
        failureRate := float64(failures) / float64(total)
        consecutiveFailures := failures >= 10

        return failureRate > 0.7 || consecutiveFailures
    },
})
```

### 自定义探测函数

```go
// 具有指数退避的自定义探测
var probeAttempts uint64

cb := hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
    TimeWindow: 30 * time.Second,
    Probe: func() bool {
        attempts := atomic.AddUint64(&probeAttempts, 1)
        // 指数退避: 1%, 2%, 4%, 8%, 最大 25%
        chance := min(25, int(attempts))
        return rand.Intn(100) < chance
    },
})
```

### 状态变化监控

```go
cb := hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
    TimeWindow: 30 * time.Second,
    OnStateChange: func(oldState, newState hystrix.State) {
        // 记录状态变化
        log.Printf("熔断器状态: %s -> %s", oldState, newState)

        // 发送指标
        metrics.Counter("circuit_breaker.state_change").
            With("from", string(oldState), "to", string(newState)).
            Increment()

        // 为开启状态触发警报
        if newState == hystrix.Open {
            alerting.Send("熔断器开启", "检测到高失败率")
        }
    },
})
```

## 性能特征

### 内存使用
- **标准熔断器**: ~200 字节 + (BufferSize * 8) 字节用于环形缓冲区
- **快速熔断器**: ~64 字节
- **批量熔断器**: 标准熔断器 + (BatchSize * 1) 字节

### 延迟
- **Before() 检查**: ~50ns（无锁原子操作）
- **After() 记录**: ~100ns（包括环形缓冲区更新）
- **Call() 开销**: 总计 ~200ns

### 吞吐量
- **标准熔断器**: >1000万操作/秒
- **快速熔断器**: >5000万操作/秒
- **并发访问**: 线性扩展到 CPU 核心数

## 最佳实践

### 1. 配置指南

**时间窗口选择:**
```go
// 短期请求（API）
TimeWindow: 10 * time.Second

// 批处理
TimeWindow: 5 * time.Minute

// 长时间运行的操作
TimeWindow: 30 * time.Minute
```

**失败阈值:**
```go
// 保守（避免误判）
ReadyToTrip: func(successes, failures uint64) bool {
    total := successes + failures
    return total >= 20 && failures*4 > successes // 80% 失败率
}

// 激进（快速失败检测）
ReadyToTrip: func(successes, failures uint64) bool {
    return failures >= 3 // 3 次连续失败后熔断
}
```

### 2. 集成模式

**优雅降级:**
```go
func GetUserProfile(userID int) *UserProfile {
    profile, err := userService.GetProfile(userID)
    if err != nil {
        // 返回缓存或默认配置文件
        return getCachedProfile(userID)
    }
    return profile
}
```

**带熔断器的重试:**
```go
func CallWithRetry(cb *hystrix.CircuitBreaker, fn func() error) error {
    for i := 0; i < 3; i++ {
        err := cb.Call(fn)
        if err == nil {
            return nil
        }

        if err.Error() == "circuit breaker is open" {
            return err // 熔断器开启时不重试
        }

        time.Sleep(time.Duration(i+1) * 100 * time.Millisecond)
    }
    return fmt.Errorf("3 次重试后失败")
}
```

### 3. 监控和警报

**指标收集:**
```go
// 定期收集和报告指标
go func() {
    ticker := time.NewTicker(10 * time.Second)
    for range ticker.C {
        successes, failures := cb.Stat()
        total := successes + failures

        metrics.Gauge("circuit_breaker.requests.total").Set(float64(total))
        metrics.Gauge("circuit_breaker.requests.successes").Set(float64(successes))
        metrics.Gauge("circuit_breaker.requests.failures").Set(float64(failures))
        metrics.Gauge("circuit_breaker.state").Set(stateToFloat(cb.State()))
    }
}()
```

## 线程安全

所有熔断器实现都完全线程安全:

- **原子操作**: 所有计数器使用原子操作
- **无锁**: 热点路径无互斥锁
- **缓存对齐**: 内存布局防止伪共享
- **并发访问**: 安全地跨多个 goroutine 使用

## 贡献

欢迎贡献！改进领域:

1. 额外的探测策略
2. 指标集成
3. 配置验证
4. 性能优化
5. 额外的状态转换策略

## 许可证

此包是 LazyGophers Utils 库的一部分，遵循相同的许可条款。