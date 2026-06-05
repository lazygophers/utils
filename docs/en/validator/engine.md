---
title: Validation Engine - Validator
---

# Engine — Validation Engine

For advanced scenarios requiring isolated validation rule sets.

## Create Engine

```go
e := validator.NewEngine()
```

## Register Field-Level Validator

```go
e.RegisterValidation("even", func(fl validator.FieldLevel) bool {
    return fl.Field().Int()%2 == 0
})
```

## Register Struct-Level Validator

Cross-field validation:

```go
validator.RegisterStructValidation(func(sl validator.StructLevel) bool {
    user := sl.Current().Interface().(User)
    return user.Password == user.ConfirmPassword
}, "User")
```

## FieldLevel Interface

Field-level validators access current field info via `FieldLevel`:

```go
type FieldLevel interface {
    // Current field
    Field() reflect.Value
    // Field name
    FieldName() string
    // Struct field name
    StructFieldName() string
    // Parameter (e.g., the 3 in min=3)
    Param() string
    // Field type
    GetTag() string
}
```

## StructLevel Interface

Struct-level validators for cross-field validation:

```go
type StructLevel interface {
    Current() reflect.Value
    ReportError(field interface{}, tag, message string)
}
```
