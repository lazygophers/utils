---
title: 自訂驗證器 - Validator
---

# 自訂驗證器

當內建標籤無法滿足需求時，可註冊自訂驗證器和翻譯。

## RegisterValidation

註冊自訂欄位級驗證器：

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

為自訂標籤註冊錯誤訊息翻譯：

```go
func RegisterTranslation(locale, tag, translation string)
```

```go
validator.RegisterTranslation("zh-TW", "mytag", "{field} 不合法")
validator.RegisterTranslation("en", "mytag", "{field} is invalid")
```

## 範例：ISBN 驗證器

```go
// 註冊驗證器
validator.RegisterValidation("isbn", func(fl validator.FieldLevel) bool {
    s := fl.Field().String()
    // ISBN 校驗邏輯...
    return isValidISBN(s)
})

// 註冊翻譯
validator.RegisterTranslation("zh-TW", "isbn", "{field} 不是有效的 ISBN")
validator.RegisterTranslation("en", "isbn", "{field} is not a valid ISBN")

// 使用
type Book struct {
    ISBN string `validate:"required,isbn"`
}
```
