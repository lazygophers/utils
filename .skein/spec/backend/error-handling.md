# Error Handling

> Fail-fast with Must*系列，panic 用于断言，log 用于运行时异常。

---

## Core Pattern: Must*

库内部使用 fail-fast 模式，不堆 `if err != nil`：

```go
// ✅ Must* 系列断言
data := utils.Must(json.Marshal(payload))
utils.MustSuccess(db.Ping())

// MustOk 断言 bool 返回值
val := utils.MustOk(m[key])

// Must 文件操作
json.MustMarshalToFile(data, "config.json")
data := json.MustUnmarshalFromFile("config.json")
```

**根包 Must 函数签名**：
- `Must[T any](value T, err error) T` — 有 error 则 panic
- `MustOk[T any](value T, ok bool) T` — ok 为 false 则 panic
- `MustSuccess(err error)` — 有 error 则 panic
- `Ignore[T any](value T, _ any) T` — 忽略第二返回值

**子包 Must 变体**：
- `json.MustMarshal`, `json.MustMarshalString`, `json.MustMarshalToFile`, `json.MustUnmarshalFromFile`
- `xtime.MustParse`
- `app.MustParse`, `app.MustSetVersion`
- `cryptox.MustGetULIDTimestamp`

## Panic 使用规则

- **允许 panic**：Must* 断言、不可恢复的配置错误、类型断言失败
- **禁止 panic**：用户输入验证、可恢复的网络/IO 错误、公共 API 的返回路径
- **recover 模式**：goroutine 边界必须 recover + log（参考 `wait/async.go`, `singledo/singledo.go`）

```go
// ✅ goroutine 内 recover
defer func() {
    if r := recover(); r != nil {
        log.Errorf("task panic [key=%s]: %v", key, r)
    }
}()
```

## Log Integration

`github.com/lazygophers/log` 用于运行时异常记录：

- `log.Debugf` — 调试信息（加密字节数、连接池深度）
- `log.Infof` — 正常流程（配置加载成功）
- `log.Warnf` — 可恢复异常（配置文件不存在，降级默认值）
- `log.Errorf` — 非预期错误（task panic、接口获取失败）
- `log.Panicf` — 不可恢复（类型不匹配）
