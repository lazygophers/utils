---
    title: candy
    ---

    # candy

    仓库里最重的数据工具包之一，覆盖类型转换、集合操作、切片与映射辅助、深拷贝与深比较。

    ## 适用场景

    - 你需要把 string / []byte / number / bool 快速转为目标类型。
- 你在处理切片、映射、去重、过滤、排序或字段索引。
- 你想减少业务代码里零散的数据整理逻辑。

    ## 你会接触到什么

    - 常见转换：`ToInt`、`ToBool`、`ToFloat64`、`ToString`。
- 集合与切片：`Map`、`Filter`、`Contains`、`Unique`、`Chunk`、`Sort`。
- 结构辅助：`ToMap`、`DeepCopy`、`DeepEqual`、各类 `KeyBy...` / `SliceField2Map...`。

## 快速示例

```go
port := candy.ToInt("8080")
active := candy.ToBool("true")
ids := candy.Unique([]int{1, 1, 2, 3})
labels := candy.Map(ids, func(v int) string { return candy.ToString(v) })
```

    ## 使用建议

    - 如果只是一次性转换，直接调用即可；如果已经开始串联多个操作，建议把数据整理逻辑集中在一层里。
- candy 功能面很广，文档按“转换 / 集合 / 结构辅助”理解会更快。

    ## 相关文档

    - [模块总览](/modules/overview)
    - [API 概览](/api/overview)
