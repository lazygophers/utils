---
title: xerror
---

# xerror

`xerror` 包负责**错误的码化、堆栈、聚合与 panic 兜底**。当一个 error 需要跨层传递、对外暴露错误码、或随用户语言切换文案时，就在这里加工它，而不是各处手写 `fmt.Errorf`。

## 适合什么场景

- 对外接口要返回稳定的错误码，并且同一个码要按用户语言展示不同文案。
- 排查线上问题时需要保留出错位置的调用栈，而不仅是一句错误消息。
- 一次操作里有多个子任务可能失败，要把它们合并成一个 error 统一返回。
- 想把第三方库或自己代码里的 `panic` 收敛成普通 error 处理，不让它击穿整个流程。

## 常用入口

### 错误码与本地化

- `New(code, msg)` / `Newf(code, format, ...)`：创建带错误码的 error，并在创建处捕获堆栈。
- `Code(err) int64`：从 error 中提取错误码。
- `RegisterMessage(tag, code, msg)`：为某个错误码注册指定语言的本地化文案（en/zh 内置，其他语言走 build tag 扩展）。
- `(*Error).LocalizedError()`：按当前 goroutine 语言（`language.Get()`）输出本地化消息。
- `(*Error).WithMetadata(key, val)`：给 error 附加结构化元数据。

### 堆栈包装

- `Wrap(err, msg)` / `Wrapf`：包装底层 error 并附加堆栈（若已带栈则不重复抓取）。
- `WithStack(err)`：只附加堆栈，不改写消息。
- `Cause(err)`：剥离包装层，解到最根部的原始 error。
- `(*Error).StackTrace() []Frame`：取出帧列表；`%+v` 会打印消息 + cause 链 + 堆栈。

### 多错误聚合

- `Join(errs...)`：合并多个 error（全 nil 返 nil，单个返回原 error）。
- `Append(dst, errs...)`：在已有聚合上追加。
- `Collector`：并发安全收集器，提供 `Add` / `ErrorOrNil` / `Len`。

### panic 辅助

- `Try(fn func())`：执行 fn，把 panic 捕获并转成带栈 error。
- `TryE(fn func() error)`：透传 fn 返回的 error，仅当真正 panic 时才转换。
- `Recover(*error)`：在 defer 中调用，把 panic 回写到目标 error 指针。

## 使用建议

- 错误码集中定义，配合 `RegisterMessage` 把文案和码解耦；展示层只认 `LocalizedError()`，不要手拼字符串。
- 本地化随 goroutine 语言切换，语言用 `language` 包设置（禁用 `context`）；在请求入口处设好语言即可，无需逐层透传。
- 跨层传递时优先 `Wrap` 保留 cause 链，调试用 `%+v` 看完整堆栈；它兼容标准库 `errors.Is` / `As` 穿透判断。
- 聚合并发任务结果用 `Collector`，单协程批量场景用 `Join` / `Append` 更直接。
- 边界处用 `Try` / `Recover` 收敛 panic，内部逻辑保持 fail-fast，不要到处包 `recover`。

## 相关文档

- [must](/modules/core/must)
- [validator](/modules/core/validator)
- [API 概览](/api/overview)
