---
title: xerror
---

# xerror

`xerror` 包提供**错误码 + 本地化消息 + cause 链 + 多 error 聚合**的极简错误类型。Error 在构造时按当前 goroutine 语言完成翻译并冻结消息 —— 不同语言的 error 对象就是不同的实例。

## 适合什么场景

- 对外接口需要稳定的错误码 + 按用户语言渲染文案。
- 跨层传递保留 cause 链，配合 stdlib `errors.Is/As/Unwrap` 穿透。
- 把多个并发任务的失败合并为单一 error。
- 错误码集中、消息和码解耦（通过外部 Localizer 注入翻译，典型 `i18n.Default`）。

## 错误码

- `CodeSuccess int = 0` — 成功，非错误。
- `CodeSystem int = -1` — 系统级错误，`Wrap` / `Wraps` 默认 / 非 *Error 提取时使用。
- `Code(err) int` — `nil → 0`；`*Error → 其 code`；其他 error `→ -1`。

## 构造与方法

### 构造函数

```go
// 仅翻译，无 fallback（未命中时 Error() 返空字符串或 fmt.Sprint(args)）
New(code int, args ...any) *Error

// 带 fallback 消息；msg 是 Sprintf 模板，args 注入；未命中 Localizer 用 msg 兜底
NewWithMsg(code int, msg string, args ...any) *Error

// 显式指定语言（不读 goroutine-local）
NewWithLanguage(tag *language.Tag, code int, args ...any) *Error

// 包装 + 翻译；err 为 nil 透传 nil；默认 code = CodeSystem
Wrap(err error, msg string, args ...any) error

// 多 error 合并为 cause（通过 Join），errors.Is 可遍历子 error；全 nil 返 nil
Wraps(errs ...error) error
```

### 方法

```go
(*Error).Error() string       // 消息（空时回退 cause.Error()）
(*Error).Msg() string         // 原始 msg 字段（不回退 cause）
(*Error).Code() int           // 错误码
(*Error).Unwrap() error       // cause，供 stdlib errors.Is/As 遍历
(*Error).Wrap(cause error) *Error  // 流式设置 cause，链式 .Wrap(a).Wrap(b)（覆盖）
```

### 链穿透

- `Cause(err) error` — 沿 `Unwrap()` 链解到根；兼容任何实现 `Unwrap() error` 的类型（含 `fmt.Errorf("%w")`、第三方 errors）。
- 标准库 `errors.Is/As/Unwrap` 全链可用。

## 本地化（Localizer 接口）

```go
type Localizer interface {
    Localize(key string, args ...any) string
    LocalizeWithLang(tag *language.Tag, key string, args ...any) string
    Register(tag *language.Tag, key, value string)
    RegisterBatch(tag *language.Tag, data map[string]any)
}
```

- 未注入时 `New` / `NewWithMsg` / `Wrap` 直接用传入的 msg（Sprintf 风格 + args）。
- 本包不内置 Localizer 实现，保持轻量。
- `i18n.Default`（`*I18n`）天然满足接口，一行接入：

```go
import (
    "github.com/lazygophers/utils/i18n"
    "github.com/lazygophers/utils/xerror"
)

func init() {
    xerror.SetLocalizer(i18n.Default)
}
```

### 翻译规则

- 查询键 = `KeyPrefix() + strconv.Itoa(code)`，默认前缀 `"error."`，改：`SetKeyPrefix("biz.")`。
- 命中：用翻译结果作为 Sprintf 模板，`fmt.Sprintf(translated, args...)`。
- 未命中：用 msg 作为 Sprintf 模板。
- msg 为空且 args 非空：直接 `fmt.Sprint(args...)`，不走翻译。

### 包级 Register

```go
Register(tag *language.Tag, key, value string)
RegisterBatch(tag *language.Tag, data map[string]any)
RegisterMessage(tag *language.Tag, code int, msg string)  // 等价 Register(tag, errorKey(code), msg)
SetLocalizer(l Localizer)
GetLocalizer() Localizer
SetKeyPrefix(prefix string)
KeyPrefix() string
```

转发到当前 Localizer，未注入时静默忽略。

## 多 error 聚合

- `Join(errs...)` — 合并多 error（全 nil 返 nil，单个返回原 error）。
- `Append(dst, errs...)` — 追加。
- `Collector` — 并发安全收集器，`Add` / `ErrorOrNil` / `Len`。
- `Wraps(errs...) error` — 把多 error 作为单个 *Error 的 cause。

## 使用建议

- 错误码集中定义；启动期一次性 `RegisterMessage` 或加载 i18n 词条；运行时只构造 + 渲染。
- 语言切换走 `language.Set`，按请求级 goroutine-local 隔离；error 在构造时已绑定该语言。
- 跨层传递优先 `Wrap` 保留 cause 链；`errors.Is/As` 自动穿透。
- 多任务结果用 `Collector` / `Join` / `Wraps`。

## 相关文档

- [language](/modules/core/language)
- [i18n](/modules/core/i18n)
- [must](/modules/core/must)
- [API 概览](/api/overview)
