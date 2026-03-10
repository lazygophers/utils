---
    title: validator
    ---

    # validator

    结构体验证与本地化错误信息包，适合配置、输入模型和边界对象。

    ## 适用场景

    - 加载配置后校验字段完整性。
- 在请求入站时统一做结构体验证。
- 需要多语言错误消息或自定义规则。

    ## 你会接触到什么

    - `validator.Struct(v)`：对结构体执行校验。
- 包内包含 locale 与自定义规则能力，适合做统一验证入口。

## 快速示例

```go
type User struct {
    Name  string `validate:"required"`
    Email string `validate:"required,email"`
}

if err := validator.Struct(&User{Name: "LazyGophers", Email: "team@example.com"}); err != nil {
    panic(err)
}
```

    ## 使用建议

    - 如果文档或旧代码里出现 `utils.Validate`，请以当前源码为准，直接使用 `validator.Struct`。
- 建议把校验放在配置加载后、业务处理前的边界位置。

    ## 相关文档

    - [模块总览](/modules/overview)
    - [API 概览](/api/overview)
