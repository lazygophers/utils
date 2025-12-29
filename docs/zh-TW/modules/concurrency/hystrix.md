---
title: hystrix - 熔斷器
---

# hystrix - 熔斷器

## 概述

hystrix 模組提供熔斷器功能,用於容錯和優雅降級。它包括針對不同用例的優化實現。

## 類型

### State

熔斷器狀態。

```go
type State string

const (
    Closed   State = "closed"   // 服務可用
    Open     State = "open"      // 服務不可用
    HalfOpen State = "half-open" // 探測狀態
)
```

---

### CircuitBreakerConfig

熔斷器配置。

```go
type CircuitBreakerConfig struct {
    TimeWindow    time.Duration // 統計時間窗口
    OnStateChange StateChange   // 狀態變化回調
    ReadyToTrip   ReadyToTrip   // 跳閘條件函數
    Probe         Probe         // 半開探測函數
    BufferSize    int           // 請求結果緩存大小
}
```

---

## 熔斷器類型

### NewCircuitBreaker()

創建優化的熔斷器。

```go
func NewCircuitBreaker(c CircuitBreakerConfig) *CircuitBreaker
```

**參數:**
- `c` - 配置選項

**返回:**
- 熔斷器實例

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

創建超輕量級熔斷器。

```go
func NewFastCircuitBreaker(failureThreshold uint64, timeWindow time.Duration) *FastCircuitBreaker
```

**參數:**
- `failureThreshold` - 失敗閾值
- `timeWindow` - 時間窗口

**返回:**
- 快速熔斷器實例

**示例:**
```go
cb := hystrix.NewFastCircuitBreaker(5, time.Minute)
```

---

### NewBatchCircuitBreaker()

創建批處理熔斷器。

```go
func NewBatchCircuitBreaker(config CircuitBreakerConfig, batchSize int, batchTimeout time.Duration) *BatchCircuitBreaker
```

**參數:**
- `config` - 配置選項
- `batchSize` - 批次大小
- `batchTimeout` - 批次超時

**返回:**
- 批處理熔斷器實例

**示例:**
```go
cb := hystrix.NewBatchCircuitBreaker(hystrix.CircuitBreakerConfig{
    TimeWindow: time.Minute,
}, 100, time.Second*10)
```

---

## 熔斷器方法

### Before()

檢查是否允許請求。

```go
func (p *CircuitBreaker) Before() bool
```

**返回:**
- 如果允許請求返回 true
- 如果熔斷器打開返回 false

**示例:**
```go
if !cb.Before() {
    return errors.New("circuit breaker is open")
}

// 執行請求
return executeRequest()
```

---

### After()

記錄請求結果。

```go
func (p *CircuitBreaker) After(success bool)
```

**參數:**
- `success` - 請求是否成功

**示例:**
```go
err := executeRequest()
cb.After(err == nil)
```

---

### Call()

使用熔斷器保護執行函數。

```go
func (p *CircuitBreaker) Call(fn func() error) error
```

**參數:**
- `fn` - 要執行的函數

**返回:**
- 函數的錯誤
- 如果熔斷器打開則返回錯誤

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

獲取當前熔斷器狀態。

```go
func (p *CircuitBreaker) State() State
```

**返回:**
- 當前狀態

**示例:**
```go
state := cb.State()
log.Infof("Circuit breaker state: %s", state)
```

---

## 使用模式

### 服務調用保護

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

### 批處理

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

### 狀態監控

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

## 最佳實踐

### 熔斷器配置

```go
// 好的做法: 配置適當的閾值
cb := hystrix.NewCircuitBreaker(hystrix.CircuitBreakerConfig{
    TimeWindow: time.Minute,
    ReadyToTrip: func(successes, failures uint64) bool {
        // 如果失敗率 > 50% 則跳閘
        total := successes + failures
        return total >= 10 && failures > total/2
    },
})
```

### 錯誤處理

```go
// 好的做法: 處理熔斷器錯誤
func safeServiceCall() (string, error) {
    err := cb.Call(func() error {
        return callService()
    })
    
    if err != nil {
        if err.Error() == "circuit breaker is open" {
            return "", nil  // 返回默認值
        }
        return "", err
    }
    
    return "success", nil
}
```

---

## 相關文檔

- [routine](/zh-TW/modules/routine) - Goroutine 管理
- [wait](/zh-TW/modules/wait) - 流程控制
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
