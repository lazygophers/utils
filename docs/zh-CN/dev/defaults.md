---
    title: defaults
    ---

    # defaults

    按标签和规则为结构体填充默认值，适合配置对象和测试夹具。

    ## 适用场景

    - 配置文件缺少部分字段时想补默认值。
- 测试中想快速拿到一组完整结构体。

    ## 你会接触到什么

    - `defaults.SetDefaults(value)`：以默认策略填充。
- `defaults.SetDefaultsWithOptions(value, opts)`：用选项控制行为。
- `defaults.RegisterCustomDefault(typeName, fn)`：注册自定义默认值。

## 快速示例

```go
type AppConfig struct {
    Host string `default:"127.0.0.1"`
    Port int    `default:"8080"`
}

cfg := &AppConfig{}
defaults.SetDefaults(cfg)
```

    ## 使用建议

    - 和 `config` 一起使用时，要先明确“文件加载优先”还是“默认值优先”。

    ## 相关文档

    - [模块总览](/modules/overview)
    - [API 概览](/api/overview)
