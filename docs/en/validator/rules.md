---
title: Built-in Rules - Validator
---

# Built-in Validation Tags

Rich set of built-in validation tags covering common scenarios.

## Basic Rules

| Tag | Description | Example |
|-----|-------------|---------|
| `required` | Non-zero value | `validate:"required"` |
| `email` | Email format | `validate:"email"` |
| `url` | URL format | `validate:"url"` |
| `alpha` | Letters only | `validate:"alpha"` |
| `alphanum` | Letters + digits | `validate:"alphanum"` |
| `json` | JSON format | `validate:"json"` |
| `uuid` | UUID format | `validate:"uuid"` |

## Numeric / Length

| Tag | Description | Example |
|-----|-------------|---------|
| `min=N` | Min value/length | `validate:"min=3"` |
| `max=N` | Max value/length | `validate:"max=100"` |
| `len=N` | Exact length | `validate:"len=11"` |
| `minlen=N` | Min length | `validate:"minlen=6"` |
| `maxlen=N` | Max length | `validate:"maxlen=20"` |
| `range=min,max` | Numeric range | `validate:"range=0.0,5.0"` |

## China-Specific

| Tag | Description | Example |
|-----|-------------|---------|
| `mobile` | Phone number (11 digits) | `validate:"mobile"` |
| `bankcard` | Bank card (Luhn) | `validate:"bankcard"` |
| `idcard` | ID card (18 digits) | `validate:"idcard"` |
| `chinese_name` | Chinese name | `validate:"chinese_name"` |

## Collection / Pattern

| Tag | Description | Example |
|-----|-------------|---------|
| `in=v1,v2,...` | Enum values | `validate:"in=male,female"` |
| `notin=v1,v2,...` | Exclude enum | `validate:"notin=admin,root"` |
| `pattern=regex` | Regex match | `validate:"pattern=^[A-Z]"` |
| `containspecial` | Has special chars | `validate:"containspecial"` |
| `strong_password` | Strong password | `validate:"strong_password"` |

## Logical Composition

```go
// And — all must pass
combined := validator.And(
    validator.Required(),
    validator.MinLength(5),
)

// Or — any must pass
either := validator.Or(
    validator.Email(),
    validator.Pattern(`^\d+$`),
)

// Not — negate
negated := validator.Not(validator.In("a", "b"))
```
