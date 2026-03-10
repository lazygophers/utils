---
    title: must
    ---

    # must

    根包中的 fail-fast 辅助，适合启动阶段、初始化阶段和必须成功的边界操作。

    ## 适用场景

    - 读取关键配置后必须继续启动。
- 依赖连接失败时不打算在当前层恢复。
- 处理 `(T, error)` 或 `(T, bool)` 返回值时希望快速收敛逻辑。

    ## 你会接触到什么

    - `utils.Must(value, err)`：返回值并在错误时直接失败。
- `utils.MustSuccess(err)`：只验证错误。
- `utils.MustOk(value, ok)`：处理 `(T, bool)` 风格返回值。
- `utils.Ignore(value, _)`：显式忽略不需要的第二返回值。

## 快速示例

```go
file := utils.Must(os.Open("app.yaml"))
defer file.Close()

port := utils.MustOk(ports["http"])
utils.MustSuccess(startServer())
```

    ## 使用建议

    - 把它们留在初始化和关键边界，不要替代所有业务错误处理。
- 如果失败后仍然需要补偿、重试或返回给调用方，就不要用 Must。

    ## 相关文档

    - [模块总览](/modules/overview)
    - [API 概览](/api/overview)
