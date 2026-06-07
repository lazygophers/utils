---
title: xerror
---

# xerror

`xerror` 套件提供**錯誤碼 + 本地化訊息 + cause 鏈 + 多 error 聚合**的極簡錯誤類型。Error 在建構時按當前 goroutine 語言完成翻譯並凍結訊息 —— 不同語言的 error 物件就是不同的實例。

## 適合什麼場景

- 對外介面需要穩定錯誤碼 + 按使用者語言渲染文案。
- 跨層傳遞保留 cause 鏈，配合 stdlib `errors.Is/As/Unwrap` 穿透。
- 把多個並行任務的失敗合併為單一 error。
- 錯誤碼集中、訊息和碼解耦（透過外部 Localizer 注入翻譯，典型 `i18n.Default`）。

## 錯誤碼

- `CodeSuccess int = 0` — 成功，非錯誤。
- `CodeSystem int = -1` — 系統級錯誤，`Wrap` / `Wraps` 預設 / 非 *Error 提取時使用。
- `Code(err) int` — `nil → 0`；`*Error → 其 code`；其他 error `→ -1`。

## 建構與方法

### 建構函式

```go
New(code int, args ...any) *Error
NewWithMsg(code int, msg string, args ...any) *Error
NewWithLanguage(tag xlanguage.Tag, code int, args ...any) *Error
Wrap(err error, msg string, args ...any) error
Wraps(errs ...error) error
```

### 方法

```go
(*Error).Error() string
(*Error).Msg() string
(*Error).Code() int
(*Error).Unwrap() error
(*Error).Wrap(cause error) *Error  // 流式：.Wrap(a).Wrap(b) 覆蓋前次 cause
```

### 鏈穿透

- `Cause(err) error` — 沿 `Unwrap()` 鏈解到根；相容任何實作 `Unwrap() error` 的類型（含 `fmt.Errorf("%w")`、第三方 errors）。
- 標準庫 `errors.Is/As/Unwrap` 全鏈可用。

## 本地化（Localizer 介面）

```go
type Localizer interface {
    LocalizeWithLang(tag xlanguage.Tag, key string, args ...any) string
    Register(tag xlanguage.Tag, key, value string)
    RegisterBatch(tag xlanguage.Tag, data map[string]any)
}
```

- 未注入時 `New` / `NewWithMsg` / `Wrap` 直接用傳入的 msg（Sprintf 風格 + args）。
- 本套件不內建 Localizer 實作，保持輕量。
- `i18n.Default`（`*I18n`）天然滿足介面，一行接入：

```go
import (
    "github.com/lazygophers/utils/i18n"
    "github.com/lazygophers/utils/xerror"
)

func init() {
    xerror.SetLocalizer(i18n.Default)
}
```

### 翻譯規則

- 查詢鍵 = `KeyPrefix() + strconv.Itoa(code)`，預設前綴 `"error."`，改：`SetKeyPrefix("biz.")`。
- 命中：用翻譯結果作為 Sprintf 模板，`fmt.Sprintf(translated, args...)`。
- 未命中：用 msg 作為 Sprintf 模板。
- msg 為空且 args 非空：直接 `fmt.Sprint(args...)`，不走翻譯。

### 套件級 Register

```go
Register(tag xlanguage.Tag, key, value string)
RegisterBatch(tag xlanguage.Tag, data map[string]any)
RegisterMessage(tag xlanguage.Tag, code int, msg string)
SetLocalizer(l Localizer)
GetLocalizer() Localizer
SetKeyPrefix(prefix string)
KeyPrefix() string
```

轉發到當前 Localizer，未注入時靜默忽略。

## 多 error 聚合

- `Join(errs...)` — 合併多 error（全 nil 回 nil，單個回傳原 error）。
- `Append(dst, errs...)` — 追加。
- `Collector` — 並發安全收集器，`Add` / `ErrorOrNil` / `Len`。
- `Wraps(errs...) error` — 把多 error 作為單個 *Error 的 cause。

## 使用建議

- 錯誤碼集中定義；啟動期一次性 `RegisterMessage` 或載入 i18n 詞條；執行階段只建構 + 渲染。
- 語言切換走 `language.Set`，按請求級 goroutine-local 隔離；error 在建構時已綁定該語言。
- 跨層傳遞優先 `Wrap` 保留 cause 鏈；`errors.Is/As` 自動穿透。
- 多任務結果用 `Collector` / `Join` / `Wraps`。

## 相關文件

- [language](/zh-TW/modules/core/language)
- [i18n](/zh-TW/modules/core/i18n)
- [must](/zh-TW/modules/core/must)
- [API 概覽](/zh-TW/api/overview)
