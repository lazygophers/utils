---
title: pyroscope - 性能分析
---

# pyroscope - 性能分析

## 概述

pyroscope 模块提供与 Pyroscope 的集成,用于生产监控和性能分析。

## 函数

### load()

加载 Pyroscope 服务器地址并开始分析。

```go
func load(address string)
```

**参数:**
- `address` - Pyroscope 服务器地址

**行为:**
- 启动 Pyroscope 客户端
- 为应用程序配置分析

**示例:**
```go
pyroscope.load("http://localhost:4040")
```

---

## 使用模式

### 应用集成

```go
func main() {
    // 开始分析
    pyroscope.load("http://localhost:4040")
    
    // 应用程序代码
    runApplication()
}
```

### 生产监控

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

## 最佳实践

### 配置

```go
// 好的做法: 配置 Pyroscope 地址
address := os.Getenv("PYROSCOPE_ADDRESS")
if address == "" {
    address = "http://localhost:4040"
}

// 好的做法: 处理连接错误
pyroscope.load(address)
// 连接错误会被记录
```

---

## 相关文档

- [runtime](/zh-CN/modules/runtime) - 运行时信息
- [app](/zh-CN/modules/app) - 应用框架
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
