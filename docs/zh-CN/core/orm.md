---
    title: orm
    ---

    # orm

    根包中的数据库字段桥接辅助，重点是把 JSON / 文本字段和结构体互转。

    ## 适用场景

    - 数据库里存储 JSON blob，但业务中想直接读写结构体。
- 扫描空值时希望自动补默认值。

    ## 你会接触到什么

    - `utils.Scan(src, dst)`：把数据库字段内容解到目标结构体。
- `utils.Value(v)`：把结构体编码成数据库可写入值。

## 快速示例

```go
type Profile struct {
    Name string `json:"name"`
}

var profile Profile
utils.MustSuccess(utils.Scan(rowValue, &profile))
value := utils.Must(utils.Value(profile))
_ = value
```

    ## 使用建议

    - 适合数据库边界层，而不是通用 JSON 转换层。
- 如果字段为空，内部会结合 defaults 做默认值处理。

    ## 相关文档

    - [模块总览](/modules/overview)
    - [API 概览](/api/overview)
