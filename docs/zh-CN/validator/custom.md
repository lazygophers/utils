---
title: 自定义验证器 - Validator
---

# 自定义验证器

当内置标签无法满足需求时，可注册自定义验证器和翻译。

## RegisterValidation

注册自定义字段级验证器：

```go
func RegisterValidation(tag string, fn ValidatorFunc) error
```

```go
validator.RegisterValidation("mytag", func(fl validator.FieldLevel) bool {
    return fl.Field().String() != "bad"
})
```

使用：

```go
type Form struct {
    Value string `validate:"required,mytag"`
}
```

## RegisterTranslation

为自定义标签注册错误消息翻译：

```go
func RegisterTranslation(locale, tag, translation string)
```

```go
validator.RegisterTranslation("zh", "mytag", "{field} 不合法")
validator.RegisterTranslation("en", "mytag", "{field} is invalid")
```

## 示例：ISBN 验证器

```go
// 注册验证器
validator.RegisterValidation("isbn", func(fl validator.FieldLevel) bool {
    s := fl.Field().String()
    // ISBN 校验逻辑...
    return isValidISBN(s)
})

// 注册翻译
validator.RegisterTranslation("zh", "isbn", "{field} 不是有效的 ISBN")
validator.RegisterTranslation("en", "isbn", "{field} is not a valid ISBN")

// 使用
type Book struct {
    ISBN string `validate:"required,isbn"`
}
```
