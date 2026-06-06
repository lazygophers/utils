---
title: xerror
---

# xerror

The `xerror` package handles **error coding, stack capture, aggregation, and panic fallback**. When an error needs to cross layers, expose a stable code to callers, or switch its text by the user's language, shape it here instead of hand-writing `fmt.Errorf` everywhere.

## When to use

- Your public API must return stable error codes, and the same code needs to display different text per the user's language.
- Diagnosing production issues requires the call stack where the error originated, not just a message string.
- A single operation has several sub-tasks that may fail, and you want to merge them into one error to return together.
- You want to converge a `panic` (from a third-party library or your own code) into an ordinary error rather than let it tear through the whole flow.

## Common entry points

### Error codes and localization

- `New(code, msg)` / `Newf(code, format, ...)`: build a coded error and capture the stack at the creation site.
- `Code(err) int64`: extract the error code from an error.
- `RegisterMessage(tag, code, msg)`: register localized text for a code in a given language (en/zh built in, other languages via build-tag extension).
- `(*Error).LocalizedError()`: output the localized message per the current goroutine language (`language.Get()`).
- `(*Error).WithMetadata(key, val)`: attach structured metadata to the error.

### Stack wrapping

- `Wrap(err, msg)` / `Wrapf`: wrap a lower-level error and attach a stack (no duplicate capture if one is already present).
- `WithStack(err)`: attach a stack only, without rewriting the message.
- `Cause(err)`: strip wrapping layers and unwind to the root original error.
- `(*Error).StackTrace() []Frame`: get the frame list; `%+v` prints the message + cause chain + stack.

### Multi-error aggregation

- `Join(errs...)`: merge multiple errors (all-nil returns nil, a single error returns the original).
- `Append(dst, errs...)`: append onto an existing aggregate.
- `Collector`: a concurrency-safe collector providing `Add` / `ErrorOrNil` / `Len`.

### panic helpers

- `Try(fn func())`: run fn, recover any panic, and convert it into a stacked error.
- `TryE(fn func() error)`: pass through the error fn returns, converting only on an actual panic.
- `Recover(*error)`: call in a defer to write a panic back into the target error pointer.

## Recommendations

- Define error codes centrally and pair them with `RegisterMessage` to decouple text from code; the presentation layer only reads `LocalizedError()` and never concatenates strings by hand.
- Localization follows the goroutine language, set via the `language` package (no `context`); set the language once at the request entry point, no need to thread it through every layer.
- Prefer `Wrap` when crossing layers to preserve the cause chain, and use `%+v` to see the full stack when debugging; it stays compatible with the standard library's `errors.Is` / `As` traversal.
- Use `Collector` to aggregate results from concurrent tasks; `Join` / `Append` are more direct for single-goroutine batch scenarios.
- Converge panics at boundaries with `Try` / `Recover`; keep internal logic fail-fast and avoid wrapping `recover` everywhere.

## Related docs

- [must](/en/modules/core/must)
- [validator](/en/modules/core/validator)
- [API Overview](/en/api/overview)
