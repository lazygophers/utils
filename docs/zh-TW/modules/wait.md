---
title: wait - 流程控制
---

# wait - 流程控制

## 概述

wait 模組提供流程控制工具,包括信號量池、同步和超時管理。

## 函數

### Lock()

獲取指定 key 的鎖。

```go
func Lock(key string)
```

**參數:**
- `key` - 鎖標識符

**行為:**
- 阻塞直到鎖可用
- 如果 key 不存在則 panic

**示例:**
```go
wait.Ready("my-resource", 10)
wait.Lock("my-resource")
defer wait.Unlock("my-resource")

// 臨界區
processResource()
```

---

### Unlock()

釋放指定 key 的鎖。

```go
func Unlock(key string)
```

**參數:**
- `key` - 鎖標識符

**行為:**
- 釋放鎖
- 如果 key 不存在則 panic

**示例:**
```go
wait.Lock("my-resource")
defer wait.Unlock("my-resource")
```

---

### Depth()

獲取指定 key 的當前深度。

```go
func Depth(key string) int
```

**參數:**
- `key` - 鎖標識符

**返回:**
- 當前深度(已獲取鎖的數量)

**示例:**
```go
depth := wait.Depth("my-resource")
log.Infof("Current depth: %d", depth)
```

---

### Sync()

使用鎖執行邏輯函數。

```go
func Sync(key string, logic func() error) error
```

**參數:**
- `key` - 鎖標識符
- `logic` - 要執行的函數

**返回:**
- 邏輯函數的錯誤

**示例:**
```go
err := wait.Sync("database", func() error {
    return updateDatabase()
})
if err != nil {
    log.Errorf("Database update failed: %v", err)
}
```

---

### Ready()

為指定 key 初始化信號量。

```go
func Ready(key string, max int)
```

**參數:**
- `key` - 鎖標識符
- `max` - 最大並發數

**示例:**
```go
wait.Ready("api-requests", 10)
```

---

## 使用模式

### 速率限制

```go
func init() {
    wait.Ready("api-calls", 100)  // 最多 100 個並發調用
}

func makeAPICall() error {
    wait.Lock("api-calls")
    defer wait.Unlock("api-calls")
    
    return callAPI()
}
```

### 資源池

```go
func init() {
    wait.Ready("database-connections", 10)
}

func queryDatabase(query string) (*Result, error) {
    wait.Lock("database-connections")
    defer wait.Unlock("database-connections")
    
    conn := getDatabaseConnection()
    defer releaseDatabaseConnection(conn)
    
    return conn.Query(query)
}
```

### 臨界區

```go
func updateSharedResource() error {
    return wait.Sync("shared-resource", func() error {
        return performUpdate()
    })
}
```

### 並發控制

```go
func processItems(items []Item) error {
    wait.Ready("processing", 10)
    
    var wg sync.WaitGroup
    errors := make(chan error, len(items))
    
    for _, item := range items {
        wg.Add(1)
        go func(item Item) {
            defer wg.Done()
            
            wait.Lock("processing")
            defer wait.Unlock("processing")
            
            if err := processItem(item); err != nil {
                errors <- err
            }
        }(item)
    }
    
    wg.Wait()
    close(errors)
    
    for err := range errors {
        if err != nil {
            return err
        }
    }
    
    return nil
}
```

---

## 最佳實踐

### 鎖管理

```go
// 好的做法: 始終使用 defer 解鎖
func safeOperation() error {
    wait.Lock("resource")
    defer wait.Unlock("resource")
    
    return performOperation()
}

// 好的做法: 檢查鎖深度
func checkConcurrency() int {
    return wait.Depth("resource")
}
```

### 初始化

```go
// 好的做法: 在啟動期間初始化信號量
func init() {
    wait.Ready("api", 100)
    wait.Ready("database", 10)
    wait.Ready("cache", 50)
}
```

---

## 相關文檔

- [routine](/zh-TW/modules/routine) - Goroutine 管理
- [hystrix](/zh-TW/modules/hystrix) - 熔斷器
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
