---
title: routine - Goroutine 管理
---

# routine - Goroutine 管理

## 概述

routine 模組提供 goroutine 管理工具,包括工作池、任務調度和 panic 恢復。

## 函數

### Go()

在 goroutine 中執行函數,帶有 panic 恢復。

```go
func Go(f func() (err error))
```

**參數:**
- `f` - 要執行的函數

**行為:**
- 在 goroutine 中執行函數
- 如果函數返回錯誤則記錄日誌
- 自動管理 trace ID

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

在 goroutine 中執行函數,帶有完整的 panic 恢復。

```go
func GoWithRecover(f func() (err error))
```

**參數:**
- `f` - 要執行的函數

**行為:**
- 在 goroutine 中執行函數
- 捕獲 panic 並記錄堆棧跟蹤
- 如果函數返回錯誤則記錄日誌

**示例:**
```go
routine.GoWithRecover(func() error {
    // 這將被捕獲並記錄
    panic("Something went wrong")
    return nil
})
```

---

### GoWithMustSuccess()

在 goroutine 中執行函數,錯誤時 panic。

```go
func GoWithMustSuccess(f func() (err error))
```

**參數:**
- `f` - 要執行的函數

**行為:**
- 在 goroutine 中執行函數
- 如果函數返回錯誤則退出進程

**示例:**
```go
routine.GoWithMustSuccess(func() error {
    if err := criticalOperation(); err != nil {
        return err
    }
    return nil
})
// 如果 criticalOperation 失敗,進程將退出
```

---

### AddBeforeRoutine()

添加在 goroutine 啟動前執行的回調。

```go
func AddBeforeRoutine(f BeforeRoutine)
```

**參數:**
- `f` - 回調函數

**示例:**
```go
routine.AddBeforeRoutine(func(baseGid, currentGid int64) {
    log.Infof("Starting goroutine: %d -> %d", baseGid, currentGid)
})
```

---

### AddAfterRoutine()

添加在 goroutine 完成後執行的回調。

```go
func AddAfterRoutine(f AfterRoutine)
```

**參數:**
- `f` - 回調函數

**示例:**
```go
routine.AddAfterRoutine(func(currentGid int64) {
    log.Infof("Completed goroutine: %d", currentGid)
})
```

---

## 使用模式

### 後台任務

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

### 錯誤處理

```go
func safeAsyncOperation() {
    routine.GoWithRecover(func() error {
        // 這個 panic 將被捕獲
        if someCondition {
            panic("Unexpected error")
        }
        return nil
    })
}
```

### 任務調度

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

## 最佳實踐

### 錯誤恢復

```go
// 好的做法: 對關鍵 goroutine 使用 GoWithRecover
routine.GoWithRecover(func() error {
    criticalOperation()
    return nil
})

// 好的做法: 對簡單任務使用 Go
routine.Go(func() error {
    simpleOperation()
    return nil
})
```

---

## 相關文檔

- [wait](/zh-TW/modules/wait) - 流程控制
- [hystrix](/zh-TW/modules/hystrix) - 熔斷器
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
