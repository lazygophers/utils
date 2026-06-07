---
title: xerror
---

# xerror

The `xerror` package provides a minimal error type with **error codes + localized messages + cause chains + multi-error aggregation**. Errors translate and freeze their message **at construction time** under the current goroutine's language — different languages produce different error instances.

## When to reach for it

- An external API needs stable error codes plus per-user-language text.
- Cross-layer propagation must preserve the cause chain and stay compatible with stdlib `errors.Is/As/Unwrap`.
- Multiple concurrent task failures need to be merged into a single error.
- Error codes should be centralized; the message catalog is decoupled from the code via an injected Localizer (typically `i18n.Default`).

## Codes

- `CodeSuccess int = 0` — success, not an error.
- `CodeSystem int = -1` — system-level error; default for `Wrap` / `Wraps` and for `Code` on non-`*Error` values.
- `Code(err) int` — `nil → 0`; `*Error → its code`; other error `→ -1`.

## Constructors & methods

### Constructors

```go
// Translate only, no fallback (Error() may be empty or fmt.Sprint(args))
New(code int, args ...any) *Error

// With fallback message; msg is the Sprintf template, args are injected;
// the Localizer's translation (if hit) is used as the Sprintf template instead.
NewWithMsg(code int, msg string, args ...any) *Error

// Explicit language (ignores the goroutine-local language)
NewWithLanguage(tag xlanguage.Tag, code int, args ...any) *Error

// Wrap + translate; err=nil passes through; default code = CodeSystem
Wrap(err error, msg string, args ...any) error

// Merge multiple errors into one cause (via Join); errors.Is can traverse any child
Wraps(errs ...error) error
```

### Methods

```go
(*Error).Error() string                  // message (falls back to cause.Error() when empty)
(*Error).Msg() string                    // raw msg field (no cause fallback)
(*Error).Code() int                      // error code
(*Error).Unwrap() error                  // cause, for stdlib errors.Is/As traversal
(*Error).Wrap(cause error) *Error        // fluent: .Wrap(a).Wrap(b) overrides cause
```

### Chain traversal

- `Cause(err) error` — unwinds via `Unwrap()` to the root; compatible with anything implementing `Unwrap() error` (including `fmt.Errorf("%w")` and third-party errors).
- Stdlib `errors.Is/As/Unwrap` work end to end.

## Localization (Localizer interface)

```go
type Localizer interface {
    LocalizeWithLang(tag xlanguage.Tag, key string, args ...any) string
    Register(tag xlanguage.Tag, key, value string)
    RegisterBatch(tag xlanguage.Tag, data map[string]any)
}
```

- Without an injected Localizer, `New` / `NewWithMsg` / `Wrap` use the passed-in msg directly (Sprintf style + args).
- The package ships no built-in Localizer implementation — xerror stays lean.
- `i18n.Default` (`*I18n`) already satisfies the interface; swap in with one line:

```go
import (
    "github.com/lazygophers/utils/i18n"
    "github.com/lazygophers/utils/xerror"
)

func init() {
    xerror.SetLocalizer(i18n.Default)
}
```

### Translation rules

- Key = `KeyPrefix() + strconv.Itoa(code)`; default prefix `"error."`, change with `SetKeyPrefix("biz.")`.
- Hit: translated text is used as the Sprintf template — `fmt.Sprintf(translated, args...)`.
- Miss: msg is used as the Sprintf template.
- Empty msg + non-empty args: `fmt.Sprint(args...)` directly, no translation lookup.

### Package-level registry

```go
Register(tag xlanguage.Tag, key, value string)
RegisterBatch(tag xlanguage.Tag, data map[string]any)
RegisterMessage(tag xlanguage.Tag, code int, msg string)  // shortcut for Register(tag, errorKey(code), msg)
SetLocalizer(l Localizer)
GetLocalizer() Localizer
SetKeyPrefix(prefix string)
KeyPrefix() string
```

Forward to the current Localizer; silently no-op when none is set.

## Multi-error aggregation

- `Join(errs...)` — merge multiple errors (all-nil → nil, single → original error).
- `Append(dst, errs...)` — append to an aggregate.
- `Collector` — concurrency-safe collector, `Add` / `ErrorOrNil` / `Len`.
- `Wraps(errs...) error` — wrap multiple errors as the cause of one `*Error`.

## Guidelines

- Define codes centrally; register at startup (either `RegisterMessage` or load i18n catalogs); the request path only constructs and renders.
- Switch language via `language.Set`; isolation is per-goroutine. The error binds that language at construction.
- Cross-layer hops should `Wrap` to preserve the cause chain; `errors.Is/As` traverse transparently.
- Aggregate multi-task results with `Collector` / `Join` / `Wraps`.

## Related docs

- [language](/en/modules/core/language)
- [i18n](/en/modules/core/i18n)
- [must](/en/modules/core/must)
- [API overview](/en/api/overview)
