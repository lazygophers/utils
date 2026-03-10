---
    title: config
    ---

    # config

    多格式配置加载入口，支持文件探测、环境变量覆盖和可选验证。

    ## 适用场景

    - 应用启动时希望自动寻找配置文件。
- 需要同时兼容 JSON、YAML、TOML、INI、HCL、XML、Properties、ENV 等格式。
- 希望加载后再统一交给 validator 校验。

    ## 你会接触到什么

    - `config.LoadConfig(c, paths...)`：加载并校验。
- `config.LoadConfigSkipValidate(c, paths...)`：只加载，不校验。
- `config.SetConfig(c)`：回写配置。
- `config.RegisterParser(ext, m, u)`：扩展新格式解析器。

## 快速示例

```go
type AppConfig struct {
    Name string `json:"name" validate:"required"`
    Port int    `json:"port" validate:"min=1,max=65535"`
}

var cfg AppConfig
utils.MustSuccess(config.LoadConfig(&cfg, "config.yaml"))
```

    ## 使用建议

    - 如果你需要更精细地控制“加载”和“校验”的时机，先用 `LoadConfigSkipValidate`，再单独调用 `validator.Struct`。
- 旧文档中出现的 `config.Load` 已不是当前源码入口。

    ## 相关文档

    - [模块总览](/modules/overview)
    - [API 概览](/api/overview)
