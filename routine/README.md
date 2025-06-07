# routine：工作流控制模块

🛠️ Go语言协程管理工具集，统一提供go routine启动、缓存、生命周期管理等能力。

## 🔧 核心用法

```go
// 基础协程启动
routine.Go(func() error {
    // your code
})

// 强行处理错误模式
routine.GoWithMustSuccess(func() error {
    // your code
})

// 带缓存的操作
cache := NewCache[string, string]()
cache.SetEx("key", "value", 5*time.Minute)
val, ok := cache.Get("key")
```

## 📜 函数说明

| 函数签名 | 描述 |
|---------|------|
`func Go(f func() (err error))` | 基础协程启动，异步执行函数<br>`before()`配置全局GID规则，`after()`用于资源回收 |

`func GoWithMustSuccess(f func() error)` | 保证函数不会因为未侦察错误而遗漏<br>会调用时自动记录到错误追踪系统 |

`func GoWithRecover(f func() error)` | 带panic捕获的进阶版本 |