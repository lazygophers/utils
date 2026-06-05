---
title: 錯誤處理 - Validator
---

# 錯誤處理

驗證失敗返回 `ValidationErrors`，包含每個欄位的 `FieldError`。

## FieldError

```go
type FieldError struct {
    Field     string      // 欄位名
    Tag       string      // 驗證標籤
    Value     interface{} // 欄位值
    Param     string      // 標籤參數
    Message   string      // 錯誤訊息（已在地化）
}
```

## ValidationErrors

`ValidationErrors` 是 `[]FieldError`，可透過型別斷言取得：

```go
err := validator.Struct(&user)
if err != nil {
    if ve, ok := err.(validator.ValidationErrors); ok {
        for _, fe := range ve {
            // 逐欄位處理
        }
    }
}
```

## 完整範例

```go
err := validator.Struct(&user)
if err != nil {
    if ve, ok := err.(validator.ValidationErrors); ok {
        for _, fe := range ve {
            switch fe.Tag() {
            case "required":
                fmt.Printf("%s 是必填欄位\n", fe.Field())
            case "email":
                fmt.Printf("%s 郵箱格式不正確\n", fe.Field())
            case "min":
                fmt.Printf("%s 不能小於 %s\n", fe.Field(), fe.Param())
            default:
                fmt.Printf("%s 驗證失敗: %s\n", fe.Field(), fe.Tag())
            }
        }
    }
}
```
