---
title: Error Handling - Validator
---

# Error Handling

Validation failure returns `ValidationErrors`, containing a `FieldError` per field.

## FieldError

```go
type FieldError struct {
    Field     string      // Field name
    Tag       string      // Validation tag
    Value     interface{} // Field value
    Param     string      // Tag parameter
    Message   string      // Error message (localized)
}
```

## ValidationErrors

`ValidationErrors` is `[]FieldError`, accessible via type assertion:

```go
err := validator.Struct(&user)
if err != nil {
    if ve, ok := err.(validator.ValidationErrors); ok {
        for _, fe := range ve {
            // Process per field
        }
    }
}
```

## Complete Example

```go
err := validator.Struct(&user)
if err != nil {
    if ve, ok := err.(validator.ValidationErrors); ok {
        for _, fe := range ve {
            switch fe.Tag() {
            case "required":
                fmt.Printf("%s is required\n", fe.Field())
            case "email":
                fmt.Printf("%s has invalid email format\n", fe.Field())
            case "min":
                fmt.Printf("%s must be at least %s\n", fe.Field(), fe.Param())
            default:
                fmt.Printf("%s failed: %s\n", fe.Field(), fe.Tag())
            }
        }
    }
}
```
