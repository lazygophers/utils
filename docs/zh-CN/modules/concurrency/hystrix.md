---
title: hystrix - 熔断器
---

# hystrix - 熔断器

## 概述

hystrix 模块提供熔断器功能,用于容错和优雅降级。它包括针对不同用例的优化实现。

## 类型

### State

熔断器状态。

```go
type State string

const (
    Closed   State = "closed"   // 服务可用
    Open     State = "open"      // 服务不可用
    HalfOpen State = "half-open" // 探测状态
)
```

---

### CircuitBreakerConfig

熔断器配置。

```go
type CircuitBreakerConfig struct {
    TimeWindow    time.Duration // 统计时间窗口
    OnStateChange StateChange   // 状态变化回调
    ReadyToTrip   ReadyToTrip   // 跳闸条件函数
    Probe         Probe         // 半开探测函数
    BufferSize    int           // 请求结果缓存大小
}
```

---

## 熔断器类型

### NewCircuitBreaker()

创建优化的熔断器。

```go
func NewCircuitBreaker(c CircuitBreakerConfig) *CircuitBreaker
```

**参数:**
- `c` - 配置选项

**返回:**
- 熔断器实例

**示例:**
```go
cb := hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
    TimeWindow: time.Minute,
    OnStateChange: func(oldState, newState hystrix.State) {
        log.Infof("State changed: %s -> %s", oldState, newState)
    },
    ReadyToTrip: func(successes, failures uint64) bool {
        total := successes + failures
        return total >= 10 && failures > successes
    },
})
```

---

### NewFastCircuitBreaker()

创建超轻量级熔断器。

```go
func NewFastCircuitBreaker(failureThreshold uint64, timeWindow time.Duration) *FastCircuitBreaker
```

**参数:**
- `failureThreshold` - 失败阈值
- `timeWindow` - 时间窗口

**返回:**
- 快速熔断器实例

**示例:**
```go
cb := hystrix.NewFastCircuitBreaker(5, time.Minute)
```

---

### NewBatchCircuitBreaker()

创建批处理熔断器。

```go
func NewBatchCircuitBreaker(config CircuitBreakerConfig, batchSize int, batchTimeout time.Duration) *BatchCircuitBreaker
```

**参数:**
- `config` - 配置选项
- `batchSize` - 批次大小
- `batchTimeout` - 批次超时

**返回:**
- 批处理熔断器实例

**示例:**
```go
cb := hystrix.NewBatchCircuitBreaker(hystrix.CircuitBreakerConfig{
    TimeWindow: time.Minute,
}, 100, time.Second*10)
```

---

## 熔断器方法

### Before()

检查是否允许请求。

```go
func (p *CircuitBreaker) Before() bool
```

**返回:**
- 如果允许请求返回 true
- 如果熔断器打开返回 false

**示例:**
```go
if !cb.Before() {
    return errors.New("circuit breaker is open")
}

// 执行请求
return executeRequest()
```

---

### After()

记录请求结果。

```go
func (p *CircuitBreaker) After(success bool)
```

**参数:**
- `success` - 请求是否成功

**示例:**
```go
err := executeRequest()
cb.After(err == nil)
```

---

### Call()

使用熔断器保护执行函数。

```go
func (p *CircuitBreaker) Call(fn func() error) error
```

**参数:**
- `fn` - 要执行的函数

**返回:**
- 函数的错误
- 如果熔断器打开则返回错误

**示例:**
```go
err := cb.Call(func() error {
    return callExternalService()
})
if err != nil {
    log.Errorf("Service call failed: %v", err)
}
```

---

### State()

获取当前熔断器状态。

```go
func (p *CircuitBreaker) State() State
```

**返回:**
- 当前状态

**示例:**
```go
state := cb.State()
log.Infof("Circuit breaker state: %s", state)
```

---

## 使用模式

### 服务调用保护

```go
func callExternalService() (string, error) {
    err := cb.Call(func() error {
        return makeHTTPRequest()
    })
    
    if err != nil {
        return "", err
    }
    
    return "success", nil
}
```

### 批处理

```go
func processBatch(items []Item) error {
    for _, item := range items {
        err := cb.Call(func() error {
            return processItem(item)
        })
        
        cb.After(err == nil)
    }
    
    return nil
}
```

### 状态监控

```go
func monitorCircuitBreaker() {
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()
    
    for range ticker.C {
        state := cb.State()
        successes, failures := cb.Stat()
        total := cb.Total()
        
        log.Infof("State: %s, Success: %d, Failures: %d, Total: %d",
            state, successes, failures, total)
    }
}
```

---

## 最佳实践

### 熔断器配置

```go
// 好的做法: 配置适当的阈值
cb := hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
    TimeWindow: time.Minute,
    ReadyToTrip: func(successes, failures uint64) bool {
        // 如果失败率 > 50% 则跳闸
        total := successes + failures
        return total >= 10 && failures > total/2
    },
})
```

### 错误处理

```go
// 好的做法: 处理熔断器错误
func safeServiceCall() (string, error) {
    err := cb.Call(func() error {
        return callService()
    })
    
    if err != nil {
        if err.Error() == "circuit breaker is open" {
            return "", nil  // 返回默认值
        }
        return "", err
    }
    
    return "success", nil
}
```

---

## 相关文档

- [routine](/zh-CN/modules/routine) - Goroutine 管理
- [wait](/zh-CN/modules/wait) - 流程控制
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
