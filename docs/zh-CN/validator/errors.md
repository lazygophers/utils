---
title: 错误处理 - Validator
---

# 错误处理

验证失败返回 `ValidationErrors`，包含每个字段的 `FieldError`。

## FieldError

```go
type FieldError struct {
    Field     string      // 字段名
    Tag       string      // 验证标签
    Value     interface{} // 字段值
    Param     string      // 标签参数
    Message   string      // 错误消息（已本地化）
}
```

## ValidationErrors

`ValidationErrors` 是 `[]FieldError`，可通过类型断言获取：

```go
err := validator.Struct(&user)
if err != nil {
    if ve, ok := err.(validator.ValidationErrors); ok {
        for _, fe := range ve {
            // 逐字段处理
        }
    }
}
```

## 完整示例

```go
err := validator.Struct(&user)
if err != nil {
    if ve, ok := err.(validator.ValidationErrors); ok {
        for _, fe := range ve {
            switch fe.Tag() {
            case "required":
                fmt.Printf("%s 是必填字段\n", fe.Field())
            case "email":
                fmt.Printf("%s 邮箱格式不正确\n", fe.Field())
            case "min":
                fmt.Printf("%s 不能小于 %s\n", fe.Field(), fe.Param())
            default:
                fmt.Printf("%s 验证失败: %s\n", fe.Field(), fe.Tag())
            }
        }
    }
}
```
