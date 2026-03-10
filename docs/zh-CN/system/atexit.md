---
    title: atexit
    ---

    # atexit

    在进程退出时执行回调，适合放清理逻辑、资源释放与关闭前收尾动作。

    ## 适用场景

    - 需要在退出时关闭连接、刷新缓冲、打印最终状态。
- 希望多个回调按注册顺序执行。

    ## 你会接触到什么

    - `atexit.Register(func())`：注册退出回调。
- 部分平台存在显式 `Exit(code)`，但并不是所有平台都通用。

## 快速示例

```go
atexit.Register(func() {
    fmt.Println("cleanup before exit")
})
```

    ## 使用建议

    - 平台差异是真实存在的，不能把某个平台的退出模型当成通用 API。
- 回调 panic 不会阻止后续回调继续执行，但你仍然应保持回调幂等和简洁。

    ## 相关文档

    - [模块总览](/modules/overview)
    - [API 概览](/api/overview)
