# hystrix

Hystrix 熔断器实现模块，提供分布式系统中的容错保护机制，防止级联故障。

## 功能特性

### 核心功能
- **熔断器模式**：基于请求成功/失败率自动熔断
- **三状态管理**：Open（正常）、Closed（熔断）、Half-Open（探测）
- **滑动窗口统计**：基于时间窗口的成功率统计
- **自动恢复**：支持自动探测服务恢复
- **并发安全**：支持高并发场景下的安全使用

### 高级特性
- **自定义熔断条件**：支持自定义熔断判断逻辑
- **状态回调通知**：状态变化时触发回调函数
- **可配置探测策略**：支持自定义探测概率
- **环形缓冲区优化**：使用内存对优化的高性能数据结构
- **自动过期清理**：自动清理过期的请求记录

## 快速开始

### 基本用法

```go
package main

import (
    "fmt"
    "time"
    
    "github.com/lazygophers/utils/hystrix"
)

func main() {
    // 创建熔断器实例
    cb := hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
        TimeWindow: time.Minute, // 1分钟统计窗口
        BufferSize: 1000,        // 缓存1000个请求结果
        OnStateChange: func(oldState, newState hystrix.State) {
            fmt.Printf("状态变化: %s -> %s\n", oldState, newState)
        },
    })
    
    // 使用熔断器保护的服务调用
    err := cb.Call(func() error {
        // 这里是受保护的服务调用
        return callExternalService()
    })
    
    if err != nil {
        fmt.Println("调用失败:", err)
    }
}

func callExternalService() error {
    // 模拟外部服务调用
    return nil
}
```

### 自定义配置

```go
// 创建带有自定义熔断条件的熔断器
cb := hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
    TimeWindow: 30 * time.Second,
    ReadyToTrip: func(successes, failures uint64) bool {
        // 当失败率超过50%时熔断
        total := successes + failures
        if total == 0 {
            return false
        }
        return failures*100/total > 50
    },
    Probe: func() bool {
        // 半开状态下，10%的概率进行探测
        return hystrix.ProbeWithChance(10)()
    },
    OnStateChange: func(oldState, newState hystrix.State) {
        log.Printf("熔断器状态: %s -> %s", oldState, newState)
    },
})
```

### 手动控制模式

```go
// 手动判断是否允许请求
if cb.Before() {
    // 执行请求
    err := doSomething()
    // 记录结果
    cb.After(err == nil)
} else {
    // 熔断器开启，拒绝请求
    fmt.Println("服务不可用，熔断器开启")
}

// 获取当前状态
state := cb.State()
fmt.Printf("当前状态: %s\n", state)

// 获取统计信息
successes, failures := cb.Stat()
fmt.Printf("成功: %d, 失败: %d\n", successes, failures)
```

## API 文档

### 类型定义

#### State
熔断器状态类型：
- `Open`：正常状态，允许请求通过
- `Closed`：熔断状态，拒绝所有请求
- `HalfOpen`：半开状态，允许部分探测请求

#### CircuitBreakerConfig
熔断器配置结构：
```go
type CircuitBreakerConfig struct {
    TimeWindow    time.Duration // 统计时间窗口
    OnStateChange StateChange   // 状态变化回调
    ReadyToTrip   ReadyToTrip   // 熔断条件判断函数
    Probe         Probe         // 探测函数
    BufferSize    int           // 请求结果缓存大小
}
```

### 主要方法

#### NewCircuitBreaker
创建熔断器实例
```go
func NewCircuitBreaker(c CircuitBreakerConfig) *CircuitBreaker
```

#### Call
执行受保护的函数调用
```go
func (p *CircuitBreaker) Call(fn func() error) error
```

#### Before
判断是否允许执行请求
```go
func (p *CircuitBreaker) Before() bool
```

#### After
记录请求执行结果
```go
func (p *CircuitBreaker) After(success bool)
```

#### State
获取当前状态
```go
func (p *CircuitBreaker) State() State
```

#### Stat
获取成功和失败计数
```go
func (p *CircuitBreaker) Stat() (successes, failures uint64)
```

### 工具函数

#### ProbeWithChance
创建指定概率的探测函数
```go
func ProbeWithChance(percent float64) Probe
```

## 设计原理

### 熔断器状态转换
```
Open (正常)
    ↓ 失败率超过阈值
Half-Open (半开)
    ↓ 探测成功/失败
    ↑ 服务恢复    ↓ 继续失败
Open (正常)    Closed (熔断)
```

### 性能优化
- 使用环形缓冲区存储请求结果，避免频繁内存分配
- 采用原子计数器，减少锁竞争
- 内存对齐优化，避免伪共享问题
- 延迟清理策略，提高并发性能

## 注意事项

1. **合理配置时间窗口**：根据业务特点设置合适的统计窗口
2. **监控状态变化**：建议通过回调监控熔断器状态
3. **配合降级策略**：熔断时应配合适当的降级方案
4. **避免频繁探测**：半开状态下的探测频率不宜过高

## 许可证

本项目采用 MIT 许可证，详见 [LICENSE](../../LICENSE) 文件。