---
title: pyroscope - 性能分析
---

# pyroscope - 性能分析

## 概述

pyroscope 模組提供與 Pyroscope 的集成,用於生產監控和性能分析。

## 函數

### load()

加載 Pyroscope 服務器地址並開始分析。

```go
func load(address string)
```

**參數:**
- `address` - Pyroscope 服務器地址

**行為:**
- 啟動 Pyroscope 客戶端
- 為應用程序配置分析

**示例:**
```go
pyroscope.load("http://localhost:4040")
```

---

## 使用模式

### 應用集成

```go
func main() {
    // 開始分析
    pyroscope.load("http://localhost:4040")
    
    // 應用程序代碼
    runApplication()
}
```

### 生產監控

```go
func setupMonitoring() {
    address := os.Getenv("PYROSCOPE_ADDRESS")
    if address == "" {
        address = "http://localhost:4040"
    }
    
    pyroscope.load(address)
    
    log.Info("Pyroscope profiling enabled")
}
```

---

## 最佳實踐

### 配置

```go
// 好的做法: 配置 Pyroscope 地址
address := os.Getenv("PYROSCOPE_ADDRESS")
if address == "" {
    address = "http://localhost:4040"
}

// 好的做法: 處理連接錯誤
pyroscope.load(address)
// 連接錯誤會被記錄
```

---

## 相關文檔

- [runtime](/zh-TW/modules/runtime) - 運行時信息
- [app](/zh-TW/modules/app) - 應用框架
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
