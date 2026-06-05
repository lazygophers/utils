---
title: Validator - Data Validation
---

# Validator

Go struct data validation library. Supports tag-based rules, custom validators, and error messages in 12 languages.

## Installation

```bash
go get github.com/lazygophers/utils/validator
```

## Quick Start

```go
import "github.com/lazygophers/utils/validator"

type User struct {
    Name     string `validate:"required,min=2,max=50"`
    Email    string `validate:"required,email"`
    Age      int    `validate:"min=0,max=150"`
    Phone    string `validate:"required,mobile"`
    BankCard string `validate:"required,bankcard"`
    IDCard   string `validate:"required,idcard"`
}

err := validator.Struct(&user)
if err != nil {
    if ve, ok := err.(validator.ValidationErrors); ok {
        for _, fe := range ve {
            fmt.Printf("field %s failed rule %s\n", fe.Field(), fe.Tag())
        }
    }
}
```

## Documentation

| Topic | Description |
|-------|-------------|
| [Core Functions](/en/validator/core) | Struct, Var, SetLocale and other basic APIs |
| [Built-in Rules](/en/validator/rules) | All built-in validation tags: required, email, min, etc. |
| [Custom Validators](/en/validator/custom) | Register custom validators and translations |
| [Validation Engine](/en/validator/engine) | Engine, FieldLevel, StructLevel for advanced use |
| [Error Handling](/en/validator/errors) | ValidationErrors and FieldError types |
| [i18n](/en/validator/i18n) | 12-language support and switching |
| [Performance & Best Practices](/en/validator/performance) | Performance optimization and usage tips |
