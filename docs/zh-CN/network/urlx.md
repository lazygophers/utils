---
    title: urlx
    ---

    # urlx

    一个很轻的 URL 辅助包，当前重点在 query 顺序规范化。

    ## 适用场景

    - 你需要把同一组查询参数稳定化，用于签名、缓存 key 或测试断言。

    ## 你会接触到什么

    - `urlx.SortQuery(query url.Values) url.Values`：按稳定顺序整理查询参数。

## 快速示例

```go
values := url.Values{"b": {"2"}, "a": {"1"}}
sorted := urlx.SortQuery(values)
_ = sorted
```

    ## 使用建议

    - 它的职责很聚焦，不建议把所有 URL 处理都堆进来。

    ## 相关文档

    - [模块总览](/modules/overview)
    - [API 概览](/api/overview)
