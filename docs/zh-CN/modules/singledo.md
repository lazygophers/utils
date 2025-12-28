---
title: singledo - 单例执行
---

# singledo - 单例执行

## 概述

singledo 模块确保操作只执行一次,防止重复计算或执行。

## 类型

### Single[T]

泛型单例执行器。

```go
type Single[T any] struct {
    mux    sync.Mutex
    last   time.Time
    wait   time.Duration
    call   *call[T]
    result T
}
```

---

## 构造函数

### NewSingle[T]()

创建新的单例执行器。

```go
func NewSingle[T any](wait time.Duration) *Single[T]
```

**参数:**
- `wait` - 允许重新执行前的等待时间

**返回:**
- 单例执行器实例

**示例:**
```go
single := singledo.NewSingle[string](time.Minute)
```

---

## 方法

### Do()

使用单例保证执行函数。

```go
func (s *Single[T]) Do(fn func() (T, error)) (v T, err error)
```

**参数:**
- `fn` - 要执行的函数

**返回:**
- 函数的结果
- 函数的错误

**行为:**
- 如果在等待时间内,返回缓存结果
- 否则,执行函数并缓存结果

**示例:**
```go
single := singledo.NewSingle[string](time.Minute)

result, err := single.Do(func() (string, error) {
    log.Info("Executing expensive operation...")
    
    data, err := fetchExpensiveData()
    if err != nil {
        return "", err
    }
    
    return data, nil
})

if err != nil {
    log.Errorf("Operation failed: %v", err)
} else {
    log.Infof("Result: %s", result)
}
```

---

### Reset()

重置单例执行器。

```go
func (s *Single[T]) Reset()
```

**示例:**
```go
single.Reset()
```

---

## 使用模式

### 昂贵计算

```go
var calculationSingle = singledo.NewSingle[int64](time.Hour)

func getExpensiveResult() (int64, error) {
    return calculationSingle.Do(func() (int64, error) {
        log.Info("Performing expensive calculation...")
        
        result, err := performExpensiveCalculation()
        if err != nil {
            return 0, err
        }
        
        return result, nil
    })
}
```

### 数据加载

```go
var dataLoaderSingle = singledo.NewSingle[[]byte](5 * time.Minute)

func loadData() ([]byte, error) {
    return dataLoaderSingle.Do(func() ([]byte, error) {
        log.Info("Loading data from remote...")
        
        data, err := http.Get("https://api.example.com/data")
        if err != nil {
            return nil, err
        }
        
        return data, nil
    })
}
```

### 配置加载

```go
var configSingle = singledo.NewSingle[Config](time.Hour)

func getConfig() (*Config, error) {
    return configSingle.Do(func() (*Config, error) {
        log.Info("Loading configuration...")
        
        cfg, err := loadConfiguration()
        if err != nil {
            return nil, err
        }
        
        return cfg, nil
    })
}
```

---

## 最佳实践

### 等待时间

```go
// 好的做法: 使用适当的等待时间
var shortCache = singledo.NewSingle[string](time.Minute)      // 1 分钟
var mediumCache = singledo.NewSingle[string](time.Hour)      // 1 小时
var longCache = singledo.NewSingle[string](24 * time.Hour) // 24 小时
```

### 错误处理

```go
// 好的做法: 正确处理错误
result, err := single.Do(func() (string, error) {
    data, err := fetchData()
    if err != nil {
        return "", err
    }
    
    return processData(data), nil
})

if err != nil {
    log.Errorf("Operation failed: %v", err)
} else {
    log.Infof("Result: %s", result)
}
```

---

## 相关文档

- [routine](/zh-CN/modules/routine) - Goroutine 管理
- [wait](/zh-CN/modules/wait) - 流程控制
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
