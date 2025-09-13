# Hystrix 模块文档

## 📋 概述

Hystrix 模块是 LazyGophers Utils 的高性能熔断器实现，基于 Netflix Hystrix 设计思想，专注于微服务架构中的故障隔离、快速失败和自动恢复。采用无锁算法和原子操作，提供极高的并发性能。

## 🎯 设计理念

### 故障隔离模式
- **快速失败**: 检测到故障时立即失败，避免级联故障
- **自动恢复**: 自动检测服务恢复并重新开启调用
- **资源保护**: 防止故障服务消耗过多系统资源

### 高性能架构
- **无锁设计**: 使用原子操作实现无锁并发
- **内存对齐**: 优化 CPU 缓存行，提升性能
- **零分配**: 核心路径实现零内存分配

## 🚀 核心功能

### 三种状态管理
- **Closed (关闭)**: 正常状态，请求正常通过
- **Open (开启)**: 熔断状态，请求直接失败
- **Half-Open (半开)**: 探测状态，允许部分请求测试服务恢复

### 灵活配置
- **时间窗口**: 可配置的统计时间窗口
- **熔断条件**: 自定义的熔断触发逻辑
- **状态回调**: 状态变化时的回调处理
- **探测策略**: 半开状态下的探测逻辑

### 三种实现变体
- **StandardCircuitBreaker**: 标准实现，平衡性能和功能
- **FastCircuitBreaker**: 高性能实现，最小化延迟
- **BulkCircuitBreaker**: 批量处理优化，适合高吞吐场景

## 📖 详细API文档

### 基础熔断器

#### NewCircuitBreaker()
```go
func NewCircuitBreaker(config CircuitBreakerConfig) *CircuitBreaker
```
**功能**: 创建标准熔断器实例

**配置参数**:
- `TimeWindow`: 统计时间窗口
- `OnStateChange`: 状态变化回调
- `ReadyToTrip`: 熔断条件判断
- `Probe`: 半开状态探测
- `BufferSize`: 结果缓存大小

**示例**:
```go
config := hystrix.CircuitBreakerConfig{
    TimeWindow: 30 * time.Second,
    BufferSize: 1000,
    ReadyToTrip: func(successes, failures uint64) bool {
        total := successes + failures
        return total >= 20 && float64(failures)/float64(total) >= 0.5
    },
    OnStateChange: func(oldState, newState hystrix.State) {
        log.Printf("熔断器状态变化: %s -> %s", oldState, newState)
    },
}

cb := hystrix.NewCircuitBreaker(config)
```

#### Call()
```go
func (cb *CircuitBreaker) Call(fn func() error) error
```
**功能**: 执行受熔断器保护的函数调用

**返回**: 函数执行结果或熔断器错误

**示例**:
```go
err := cb.Call(func() error {
    // 实际的业务逻辑调用
    return httpClient.Get("https://api.example.com/data")
})

if err != nil {
    if errors.Is(err, hystrix.ErrOpenState) {
        // 熔断器开启，服务暂时不可用
        log.Println("服务熔断中，请稍后重试")
    } else {
        // 其他错误
        log.Printf("调用失败: %v", err)
    }
}
```

#### State()
```go
func (cb *CircuitBreaker) State() State
```
**功能**: 获取当前熔断器状态

**示例**:
```go
switch cb.State() {
case hystrix.Closed:
    fmt.Println("服务正常")
case hystrix.Open:
    fmt.Println("服务熔断中")
case hystrix.HalfOpen:
    fmt.Println("服务恢复测试中")
}
```

### 高性能变体

#### NewFastCircuitBreaker()
```go
func NewFastCircuitBreaker(config FastConfig) *FastCircuitBreaker
```
**功能**: 创建高性能熔断器，最小化延迟

**特点**:
- 更少的原子操作
- 简化的状态机
- 针对低延迟优化

**示例**:
```go
fastCB := hystrix.NewFastCircuitBreaker(hystrix.FastConfig{
    FailureThreshold: 5,
    TimeoutDuration:  100 * time.Millisecond,
})

err := fastCB.Execute(func() error {
    return performCriticalOperation()
})
```

#### NewBulkCircuitBreaker()
```go
func NewBulkCircuitBreaker(config BulkConfig) *BulkCircuitBreaker
```
**功能**: 创建批量处理优化的熔断器

**特点**:
- 批量状态检查
- 减少锁竞争
- 适合高吞吐场景

**示例**:
```go
bulkCB := hystrix.NewBulkCircuitBreaker(hystrix.BulkConfig{
    BatchSize:       100,
    FlushInterval:   10 * time.Millisecond,
    FailureRatio:    0.5,
})

// 批量处理
results := bulkCB.ExecuteBatch(requests)
```

## 🔧 高级特性

### 自定义熔断条件

#### 基于失败率的熔断
```go
config.ReadyToTrip = func(successes, failures uint64) bool {
    total := successes + failures
    if total < 10 {
        return false // 样本不足，不熔断
    }
    
    failureRate := float64(failures) / float64(total)
    return failureRate >= 0.5 // 失败率超过50%时熔断
}
```

#### 基于响应时间的熔断
```go
var slowRequests uint64

config.ReadyToTrip = func(successes, failures uint64) bool {
    slow := atomic.LoadUint64(&slowRequests)
    total := successes + failures + slow
    
    if total < 20 {
        return false
    }
    
    slowRate := float64(slow) / float64(total)
    return slowRate >= 0.3 // 慢请求超过30%时熔断
}
```

#### 复合条件熔断
```go
config.ReadyToTrip = func(successes, failures uint64) bool {
    total := successes + failures
    
    // 条件1: 最小请求数
    if total < 50 {
        return false
    }
    
    // 条件2: 失败率
    failureRate := float64(failures) / float64(total)
    if failureRate >= 0.5 {
        return true
    }
    
    // 条件3: 连续失败次数
    if failures >= 10 {
        return true
    }
    
    return false
}
```

### 智能探测策略

#### 渐进式探测
```go
var probeAttempts uint64

config.Probe = func() bool {
    attempts := atomic.AddUint64(&probeAttempts, 1)
    
    // 渐进式增加探测频率
    switch {
    case attempts <= 3:
        return attempts%1 == 0  // 每次都探测
    case attempts <= 10:
        return attempts%2 == 0  // 每2次探测一次
    default:
        return attempts%5 == 0  // 每5次探测一次
    }
}
```

#### 时间窗口探测
```go
var lastProbeTime int64

config.Probe = func() bool {
    now := time.Now().UnixNano()
    last := atomic.LoadInt64(&lastProbeTime)
    
    // 每5秒最多探测一次
    if now-last >= 5*1e9 {
        atomic.StoreInt64(&lastProbeTime, now)
        return true
    }
    return false
}
```

### 状态变化处理

#### 监控和告警
```go
config.OnStateChange = func(oldState, newState hystrix.State) {
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    
    switch newState {
    case hystrix.Open:
        // 熔断开启告警
        alertManager.SendAlert(fmt.Sprintf(
            "[%s] 服务熔断开启: %s -> %s", 
            timestamp, oldState, newState))
        
        // 记录监控指标
        metrics.Counter("circuit_breaker_open").Inc()
        
    case hystrix.Closed:
        // 服务恢复通知
        log.Printf("[%s] 服务已恢复正常", timestamp)
        metrics.Counter("circuit_breaker_recovered").Inc()
        
    case hystrix.HalfOpen:
        // 探测状态
        log.Printf("[%s] 开始服务恢复探测", timestamp)
        metrics.Counter("circuit_breaker_probe").Inc()
    }
}
```

#### 自适应配置调整
```go
var failureHistory []float64

config.OnStateChange = func(oldState, newState hystrix.State) {
    if newState == hystrix.Open {
        // 记录失败率历史
        rate := calculateCurrentFailureRate()
        failureHistory = append(failureHistory, rate)
        
        // 保持最近10次记录
        if len(failureHistory) > 10 {
            failureHistory = failureHistory[1:]
        }
        
        // 根据历史调整熔断阈值
        avgRate := calculateAverage(failureHistory)
        if avgRate > 0.7 {
            // 提高敏感度
            adjustThreshold(0.4)
        } else if avgRate < 0.3 {
            // 降低敏感度
            adjustThreshold(0.6)
        }
    }
}
```

## 🚀 实际应用场景

### 微服务调用保护

#### HTTP 客户端保护
```go
// 创建针对特定服务的熔断器
userServiceCB := hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
    TimeWindow: 30 * time.Second,
    BufferSize: 1000,
    ReadyToTrip: func(successes, failures uint64) bool {
        total := successes + failures
        return total >= 20 && float64(failures)/float64(total) >= 0.5
    },
})

// 包装HTTP调用
func GetUser(userID string) (*User, error) {
    var user *User
    var err error
    
    cbErr := userServiceCB.Call(func() error {
        user, err = httpClient.GetUser(userID)
        return err
    })
    
    if cbErr != nil {
        if errors.Is(cbErr, hystrix.ErrOpenState) {
            // 返回缓存数据或默认数据
            return getCachedUser(userID), nil
        }
        return nil, cbErr
    }
    
    return user, err
}
```

#### 数据库连接保护
```go
dbCircuitBreaker := hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
    TimeWindow: 10 * time.Second,
    ReadyToTrip: func(successes, failures uint64) bool {
        return failures >= 5 // 连续5次失败则熔断
    },
})

func QueryDatabase(query string) ([]Row, error) {
    var rows []Row
    var err error
    
    cbErr := dbCircuitBreaker.Call(func() error {
        rows, err = db.Query(query)
        return err
    })
    
    if cbErr != nil {
        if errors.Is(cbErr, hystrix.ErrOpenState) {
            // 数据库不可用，使用只读副本
            return queryReadReplica(query)
        }
        return nil, cbErr
    }
    
    return rows, err
}
```

### 资源限制和保护

#### 第三方API调用限制
```go
// 限制第三方API调用频率
apiRateLimiter := hystrix.NewFastCircuitBreaker(hystrix.FastConfig{
    FailureThreshold: 3,
    TimeoutDuration:  5 * time.Second,
})

func CallThirdPartyAPI(request *APIRequest) (*APIResponse, error) {
    var response *APIResponse
    var err error
    
    cbErr := apiRateLimiter.Execute(func() error {
        // 添加超时控制
        ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
        defer cancel()
        
        response, err = thirdPartyClient.CallWithContext(ctx, request)
        return err
    })
    
    if cbErr != nil {
        // 记录失败原因
        log.Printf("第三方API调用失败: %v", cbErr)
        return nil, cbErr
    }
    
    return response, err
}
```

#### 内存密集型操作保护
```go
memoryIntensiveCB := hystrix.NewBulkCircuitBreaker(hystrix.BulkConfig{
    BatchSize:     10,
    FlushInterval: 100 * time.Millisecond,
    FailureRatio:  0.3,
})

func ProcessLargeData(data []byte) (*ProcessResult, error) {
    // 检查系统内存使用率
    if getMemoryUsage() > 0.8 {
        return nil, hystrix.ErrOpenState
    }
    
    var result *ProcessResult
    var err error
    
    cbErr := memoryIntensiveCB.Execute(func() error {
        result, err = heavyProcessing(data)
        return err
    })
    
    return result, cbErr
}
```

## 📊 性能特点

### 基准测试结果

| 操作 | 标准实现 | Fast实现 | Bulk实现 | 性能提升 |
|------|----------|----------|----------|----------|
| **状态检查** | 2 ns/op | 1 ns/op | 1.5 ns/op | 2x |
| **成功调用** | 25 ns/op | 15 ns/op | 20 ns/op | 1.7x |
| **失败调用** | 30 ns/op | 20 ns/op | 25 ns/op | 1.5x |
| **状态变化** | 100 ns/op | 80 ns/op | 90 ns/op | 1.3x |

### 内存使用优化

```go
// 内存对齐优化
type CircuitBreaker struct {
    // 热点字段放在同一缓存行
    state    uint32  // 原子操作
    failures uint64  // 原子操作
    successes uint64 // 原子操作
    
    // 配置字段（冷数据）
    config Config
    // ...
}
```

### 并发性能

- **无锁设计**: 使用 `sync/atomic` 包避免锁竞争
- **读写分离**: 热点读操作使用原子加载，写操作使用 CAS
- **缓存友好**: 相关字段内存对齐，提升缓存命中率

## 🚨 使用注意事项

### 配置合理性

1. **时间窗口设置**
   ```go
   // ❌ 时间窗口过短，可能导致误判
   TimeWindow: 1 * time.Second
   
   // ✅ 合理的时间窗口
   TimeWindow: 30 * time.Second
   ```

2. **缓冲区大小**
   ```go
   // ❌ 缓冲区过小，统计不准确
   BufferSize: 10
   
   // ✅ 合理的缓冲区大小
   BufferSize: 1000
   ```

### 熔断条件设计

1. **避免过度敏感**
   ```go
   // ❌ 过度敏感，可能误熔断
   ReadyToTrip: func(successes, failures uint64) bool {
       return failures >= 1
   }
   
   // ✅ 合理的熔断条件
   ReadyToTrip: func(successes, failures uint64) bool {
       total := successes + failures
       return total >= 10 && float64(failures)/float64(total) >= 0.5
   }
   ```

2. **考虑业务特性**
   ```go
   // 为不同服务设置不同的熔断策略
   func createServiceCircuitBreaker(serviceType string) *hystrix.CircuitBreaker {
       var config hystrix.CircuitBreakerConfig
       
       switch serviceType {
       case "critical":
           // 关键服务：更宽松的熔断条件
           config.ReadyToTrip = func(s, f uint64) bool {
               return f >= 10 && float64(f)/(float64(s+f)) >= 0.7
           }
       case "optional":
           // 可选服务：更严格的熔断条件
           config.ReadyToTrip = func(s, f uint64) bool {
               return f >= 3 && float64(f)/(float64(s+f)) >= 0.3
           }
       }
       
       return hystrix.NewCircuitBreaker(config)
   }
   ```

### 状态恢复策略

```go
// 渐进式恢复
var consecutiveSuccesses uint64

config.OnStateChange = func(oldState, newState hystrix.State) {
    if newState == hystrix.Closed {
        atomic.StoreUint64(&consecutiveSuccesses, 0)
    }
}

config.Probe = func() bool {
    successes := atomic.LoadUint64(&consecutiveSuccesses)
    
    // 需要连续成功才完全恢复
    if successes < 5 {
        return true // 继续探测
    }
    
    return false // 稳定后减少探测频率
}
```

## 💡 最佳实践

### 1. 分层熔断策略
```go
// 为不同层级设置不同的熔断器
type ServiceLayer struct {
    dbCircuitBreaker   *hystrix.CircuitBreaker
    cacheCircuitBreaker *hystrix.CircuitBreaker
    apiCircuitBreaker   *hystrix.CircuitBreaker
}

func (s *ServiceLayer) GetData(id string) (*Data, error) {
    // 首先尝试缓存
    var data *Data
    var err error
    
    cbErr := s.cacheCircuitBreaker.Call(func() error {
        data, err = s.getFromCache(id)
        return err
    })
    
    if cbErr == nil && data != nil {
        return data, nil
    }
    
    // 缓存失败，尝试数据库
    cbErr = s.dbCircuitBreaker.Call(func() error {
        data, err = s.getFromDB(id)
        return err
    })
    
    return data, cbErr
}
```

### 2. 监控和指标收集
```go
type CircuitBreakerMetrics struct {
    totalRequests   uint64
    successRequests uint64
    failedRequests  uint64
    rejectedRequests uint64
    stateChanges    uint64
}

func (m *CircuitBreakerMetrics) OnCall(success bool) {
    atomic.AddUint64(&m.totalRequests, 1)
    if success {
        atomic.AddUint64(&m.successRequests, 1)
    } else {
        atomic.AddUint64(&m.failedRequests, 1)
    }
}

func (m *CircuitBreakerMetrics) OnReject() {
    atomic.AddUint64(&m.rejectedRequests, 1)
}

func (m *CircuitBreakerMetrics) OnStateChange() {
    atomic.AddUint64(&m.stateChanges, 1)
}
```

### 3. 优雅降级
```go
func GetUserWithFallback(userID string) (*User, error) {
    var user *User
    var err error
    
    // 尝试主服务
    cbErr := primaryServiceCB.Call(func() error {
        user, err = primaryService.GetUser(userID)
        return err
    })
    
    if cbErr != nil {
        // 主服务不可用，尝试备用服务
        cbErr = fallbackServiceCB.Call(func() error {
            user, err = fallbackService.GetUser(userID)
            return err
        })
        
        if cbErr != nil {
            // 备用服务也不可用，返回缓存数据
            if cachedUser := getUserFromCache(userID); cachedUser != nil {
                return cachedUser, nil
            }
            
            // 最后降级：返回默认用户数据
            return &User{
                ID:   userID,
                Name: "未知用户",
            }, nil
        }
    }
    
    return user, err
}
```

## 🔗 相关模块

- **[wait](../wait/)**: 并发控制和超时管理
- **[retry](../retry/)**: 重试机制（与熔断器互补）
- **[ratelimit](../ratelimit/)**: 限流控制

## 📚 更多资源

- **[熔断器模式详解](./patterns.md)**: 熔断器设计模式
- **[性能调优指南](./performance.md)**: 性能优化技巧
- **[监控最佳实践](./monitoring.md)**: 监控和告警设置
- **[示例代码](./examples/)**: 丰富的使用示例