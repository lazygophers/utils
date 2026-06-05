---
title: 驗證引擎 - Validator
---

# Engine — 驗證引擎

適合需要多個隔離驗證規則集的進階場景。

## 建立引擎

```go
e := validator.NewEngine()
```

## 註冊欄位級驗證器

```go
e.RegisterValidation("even", func(fl validator.FieldLevel) bool {
    return fl.Field().Int()%2 == 0
})
```

## 註冊結構體級驗證器

跨欄位校驗：

```go
validator.RegisterStructValidation(func(sl validator.StructLevel) bool {
    user := sl.Current().Interface().(User)
    return user.Password == user.ConfirmPassword
}, "User")
```

## FieldLevel 介面

欄位級驗證器透過 `FieldLevel` 存取當前欄位資訊：

```go
type FieldLevel interface {
    // 當前欄位
    Field() reflect.Value
    // 欄位名
    FieldName() string
    // 結構體名
    StructFieldName() string
    // 參數（如 min=3 中的 3）
    Param() string
    // 欄位型別
    GetTag() string
}
```

## StructLevel 介面

結構體級驗證器用於跨欄位校驗：

```go
type StructLevel interface {
    Current() reflect.Value
    ReportError(field interface{}, tag, message string)
}
```
