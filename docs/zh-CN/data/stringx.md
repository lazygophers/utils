---
    title: stringx
    ---

    # stringx

    偏命名规范转换与字符串辅助，适合接口协议、代码生成和展示层字符串整理。

    ## 适用场景

    - 在 snake_case、camelCase、kebab-case 之间转换。
- 需要安全地做 `[]byte` / `string` 互转。
- 需要截断、分段、反转或判断字符形态。

    ## 你会接触到什么

    - 命名转换：`Camel2Snake`、`Snake2Camel`、`ToKebab`、`ToDot`、`ToSlash`。
- 基础辅助：`ToString`、`ToBytes`、`SplitLen`、`Shorten`、`Reverse`。

## 快速示例

```go
fmt.Println(stringx.Camel2Snake("HTTPServer"))
fmt.Println(stringx.ToKebab("CreateUserProfile"))
```

    ## 使用建议

    - 适合做命名规整，不建议在业务核心里滥用各种大小写转换掩盖建模问题。

    ## 相关文档

    - [模块总览](/modules/overview)
    - [API 概览](/api/overview)
