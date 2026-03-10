---
    title: json
    ---

    # json

    围绕 JSON 编解码、字符串转换和文件读写的工程化封装。

    ## 适用场景

    - 你希望统一项目里的 JSON 入口。
- 你需要直接从文件读写 JSON。
- 你想在 panic 风格的初始化路径里快速使用 JSON。

    ## 你会接触到什么

    - `json.Marshal` / `json.Unmarshal`
- `json.MarshalString` / `json.UnmarshalString`
- `json.MarshalToFile` / `json.UnmarshalFromFile`
- `json.MustMarshal` / `json.MustMarshalString`

## 快速示例

```go
body := utils.Must(json.MarshalString(map[string]any{"name": "LazyGophers"}))
fmt.Println(body)
```

    ## 使用建议

    - 如果你要处理数据库字段，请优先看 `orm`；如果你只需要 JSON 编解码，直接使用本包即可。

    ## 相关文档

    - [模块总览](/modules/overview)
    - [API 概览](/api/overview)
