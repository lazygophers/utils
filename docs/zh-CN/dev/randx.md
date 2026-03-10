---
    title: randx
    ---

    # randx

    随机值与随机集合操作辅助，覆盖 bool、数值、集合选择和时间相关随机。

    ## 适用场景

    - 快速生成随机测试输入。
- 从集合中随机挑选元素或打乱顺序。

    ## 你会接触到什么

    - 布尔：`Bool`、`Booln`、`WeightedBool`。
- 数值：`Intn`、`Int64nRange`、`Float64Range`。
- 集合：`Choose`、`ChooseN`、`Shuffle`、`WeightedChoose`。

## 快速示例

```go
picked := randx.Choose([]string{"a", "b", "c"})
port := randx.IntnRange(2000, 9000)
_ = picked
_ = port
```

    ## 使用建议

    - 如果你需要“看起来像真实数据”的样本，优先配合 `fake`。

    ## 相关文档

    - [模块总览](/modules/overview)
    - [API 概览](/api/overview)
