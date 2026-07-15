# Logging Guidelines

> 使用 `github.com/lazygophers/log`，中文消息，结构化 trace。

---

## Library

`github.com/lazygophers/log` — 项目统一日志库。

## Levels

| Level | 用途 | 示例 |
|-------|------|------|
| `Debugf` | 开发诊断 | `log.Debugf("成功加密 %d 字节数据", len(data))` |
| `Infof` | 正常流程里程碑 | `log.Info("load config success")` |
| `Warnf` | 可恢复降级 | `log.Warnf("Config file not found, use default config")` |
| `Errorf` | 非预期但可继续 | `log.Errorf("task panic [key=%s]: %v", key, r)` |
| `Panicf` | 不可恢复，立即终止 | `log.Panicf("源类型 %s 与目标类型 %s 不匹配", v1.Type(), v2.Type())` |

## Trace System

支持 goroutine 级别 trace 链路追踪：

```go
log.SetTraceWithGID(gid, fmt.Sprintf("%s.%s", log.GetTraceWithGID(baseGid), log.GenTraceId()))
// ... do work ...
log.DelTraceWithGID(gid)
```

## Conventions

- 日志消息用中文（与代码注释风格一致）
- 包含上下文：key、长度、耗时等关键参数
- goroutine panic 必须 `recover` + `log.Errorf`
- 不在公共 API 返回路径打日志（由调用方决定如何处理）
- 使用 `log.PutBuffer(buf)` 回收 buffer（`sync.Pool` 模式）

## Buffer Pool

```go
b := log.GetBuffer()
defer log.PutBuffer(b)
// use b...
```
