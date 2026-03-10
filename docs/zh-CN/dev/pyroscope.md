---
    title: pyroscope
    ---

    # pyroscope

    接入 Pyroscope 采样分析的轻量入口。

    ## 适用场景

    - 想在服务启动时快速挂上性能采样。

    ## 你会接触到什么

    - `pyroscope.Load(address string)`：按地址接入 Pyroscope。

## 快速示例

```go
pyroscope.Load("http://127.0.0.1:4040")
```

    ## 使用建议

    - 它的目标是降低接入门槛，不是替代完整的性能治理流程。

    ## 相关文档

    - [模块总览](/modules/overview)
    - [API 概览](/api/overview)
