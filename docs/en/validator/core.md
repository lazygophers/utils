---
title: Core Functions - Validator
---

# Core Functions

Basic APIs: struct validation, single-field validation, locale setting, and instance management.

## Struct — Struct Validation

Validates a struct using tag rules. Returns `ValidationErrors` on failure.

```go
func Struct(s interface{}) error
```

```go
type Config struct {
    Host  string `validate:"required,url"`
    Port  int    `validate:"required,min=1,max=65535"`
    Debug bool   `validate:"required"`
}

err := validator.Struct(&cfg)
```

## Var — Single Field Validation

Validates a single value, useful for non-struct scenarios.

```go
func Var(field interface{}, tag string) error
```

```go
err := validator.Var("user@example.com", "required,email")
err := validator.Var(42, "min=0,max=100")
```

## SetLocale — Set Language

Sets the global error message language for all subsequent validations.

```go
func SetLocale(locale string)
```

Supported: `en` (default), `zh`, `zh-TW`, `ja`, `ko`, `ar`, `de`, `es`, `fr`, `it`, `pt`, `ru`.

```go
validator.SetLocale("en")
err := validator.Struct(&user)
```

## Validator Instance

### Default — Global Singleton

```go
v := validator.Default()
err := v.Struct(&data)
```

### New — Custom Configuration

```go
v, err := validator.New(
    validator.WithLocale("en"),
    validator.WithUseJSON(true),
    validator.WithTranslations(map[string]string{
        "custom": "Custom message",
    }),
    validator.WithCustomValidator("mytag", func(v interface{}) bool {
        return v.(string) != ""
    }),
)
```

### Config Struct

```go
type Config struct {
    Locale        string
    UseJSON       bool
    Translations  map[string]string
}
```

Configure all at once with `WithConfig`:

```go
v, _ := validator.New(validator.WithConfig(validator.Config{
    Locale:  "en",
    UseJSON: true,
}))
```
