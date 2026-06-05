---
title: 验证引擎 - Validator
---

# Engine — 验证引擎

适合需要多个隔离验证规则集的高级场景。

## 创建引擎

```go
e := validator.NewEngine()
```

## 注册字段级验证器

```go
e.RegisterValidation("even", func(fl validator.FieldLevel) bool {
    return fl.Field().Int()%2 == 0
})
```

## 注册结构体级验证器

跨字段校验：

```go
validator.RegisterStructValidation(func(sl validator.StructLevel) bool {
    user := sl.Current().Interface().(User)
    return user.Password == user.ConfirmPassword
}, "User")
```

## FieldLevel 接口

字段级验证器通过 `FieldLevel` 访问当前字段信息：

```go
type FieldLevel interface {
    // 当前字段
    Field() reflect.Value
    // 字段名
    FieldName() string
    // 结构体名
    StructFieldName() string
    // 参数（如 min=3 中的 3）
    Param() string
    // 字段类型
    GetTag() string
}
```

## StructLevel 接口

结构体级验证器用于跨字段校验：

```go
type StructLevel interface {
    Current() reflect.Value
    ReportError(field interface{}, tag, message string)
}
```
