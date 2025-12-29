---
title: singledo - 單例執行
---

# singledo - 單例執行

## 概述

singledo 模組確保操作只執行一次,防止重複計算或執行。

## 類型

### Single[T]

泛型單例執行器。

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

## 構造函數

### NewSingle[T]()

創建新的單例執行器。

```go
func NewSingle[T any](wait time.Duration) *Single[T]
```

**參數:**
- `wait` - 允許重新執行前的等待時間

**返回:**
- 單例執行器實例

**示例:**
```go
single := singledo.NewSingle[string](time.Minute)
```

---

## 方法

### Do()

使用單例保證執行函數。

```go
func (s *Single[T]) Do(fn func() (T, error)) (v T, err error)
```

**參數:**
- `fn` - 要執行的函數

**返回:**
- 函數的結果
- 函數的錯誤

**行為:**
- 如果在等待時間內,返回緩存結果
- 否則,執行函數並緩存結果

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

重置單例執行器。

```go
func (s *Single[T]) Reset()
```

**示例:**
```go
single.Reset()
```

---

## 使用模式

### 昂貴計算

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

### 數據加載

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

### 配置加載

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

## 最佳實踐

### 等待時間

```go
// 好的做法: 使用適當的等待時間
var shortCache = singledo.NewSingle[string](time.Minute)      // 1 分鐘
var mediumCache = singledo.NewSingle[string](time.Hour)      // 1 小時
var longCache = singledo.NewSingle[string](24 * time.Hour) // 24 小時
```

### 錯誤處理

```go
// 好的做法: 正確處理錯誤
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

## 相關文檔

- [routine](/zh-TW/modules/routine) - Goroutine 管理
- [wait](/zh-TW/modules/wait) - 流程控制
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
