---
title: Custom Validators - Validator
---

# Custom Validators

When built-in tags don't cover your needs, register custom validators and translations.

## RegisterValidation

Register a custom field-level validator:

```go
func RegisterValidation(tag string, fn ValidatorFunc) error
```

```go
validator.RegisterValidation("mytag", func(fl validator.FieldLevel) bool {
    return fl.Field().String() != "bad"
})
```

Usage:

```go
type Form struct {
    Value string `validate:"required,mytag"`
}
```

## RegisterTranslation

Register error message translations for custom tags:

```go
func RegisterTranslation(locale, tag, translation string)
```

```go
validator.RegisterTranslation("en", "mytag", "{field} is invalid")
validator.RegisterTranslation("zh", "mytag", "{field} 不合法")
```

## Example: ISBN Validator

```go
// Register validator
validator.RegisterValidation("isbn", func(fl validator.FieldLevel) bool {
    s := fl.Field().String()
    // ISBN validation logic...
    return isValidISBN(s)
})

// Register translations
validator.RegisterTranslation("en", "isbn", "{field} is not a valid ISBN")
validator.RegisterTranslation("zh", "isbn", "{field} 不是有效的 ISBN")

// Usage
type Book struct {
    ISBN string `validate:"required,isbn"`
}
```
