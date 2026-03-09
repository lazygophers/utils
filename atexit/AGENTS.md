# Atexit Package - Shutdown Callback Notes

**Package:** `github.com/lazygophers/utils/atexit`

## OVERVIEW
- `atexit` registers callbacks that run during process shutdown.
- The package is platform-aware: most OS families use signal handlers; `plan9`, `js/wasm`, and `wasip1/wasm` expose explicit `Exit(code)`.
- Read local platform files before changing semantics; this package is not a single generic implementation.

## WHERE TO LOOK
- `atexit.go`: fallback implementation for unsupported/future platforms.
- `atexit_linux.go`, `atexit_linux_arm.go`: Linux/Android-specific behavior.
- `atexit_darwin.go`, `atexit_bsd.go`, `atexit_solaris.go`, `atexit_windows.go`, `atexit_aix.go`: OS-family implementations.
- `atexit_plan9.go`, `atexit_js.go`, `atexit_wasip1.go`: explicit-exit platforms.
- `PLATFORM_SUPPORT.md`: platform matrix and intended support claims.

## LOCAL RULES
- Public entrypoint is `Register(func())`; nil callbacks are ignored.
- Callback execution preserves registration order.
- Callback panics are isolated so one bad callback does not block later callbacks.
- Signal handler setup is guarded by `sync.Once`; avoid duplicate init paths.
- When editing behavior, keep the platform split aligned with build tags instead of pushing everything into one file.

## GOTCHAS
- Not every platform has identical signal support.
- `Exit(code)` is only present on explicit-exit platforms; do not document it as universal API without checking target file.
- Shutdown paths call `os.Exit`, so code after signal-triggered exit is unreachable.

## TESTING
```bash
go test ./atexit/...
make test
```
