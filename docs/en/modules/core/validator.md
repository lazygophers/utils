---
title: validator - Data Validation
---

# validator - Data Validation

## Overview

The validator module provides comprehensive data validation for Go structs with support for multiple locales and custom validation rules. It offers built-in validators for common data types and allows registration of custom validators.

## Core Functions

### Struct()

Validate struct data using struct tags.

```go
func Struct(s interface{}) error
```

**Parameters:**
- `s` - Struct pointer to validate

**Returns:**
- Validation error if validation fails
- nil if validation passes

**Example:**
```go
type User struct {
    Name  string `validate:"required,min=2,max=50"`
    Email string `validate:"required,email"`
    Age   int    `validate:"min=0,max=150"`
}

user := User{
    Name:  "John Doe",
    Email: "john@example.com",
    Age:   25,
}

if err := validator.Struct(&user); err != nil {
    fmt.Printf("Validation failed: %v\n", err)
}
```

---

### Var()

Validate a single variable with validation tags.

```go
func Var(field interface{}, tag string) error
```

**Parameters:**
- `field` - Value to validate
- `tag` - Validation rules

**Returns:**
- Validation error if validation fails
- nil if validation passes

**Example:**
```go
if err := validator.Var("john@example.com", "required,email"); err != nil {
    fmt.Printf("Invalid email: %v\n", err)
}

if err := validator.Var(25, "min=0,max=150"); err != nil {
    fmt.Printf("Invalid age: %v\n", err)
}
```

---

### SetLocale()

Set the locale for error messages.

```go
func SetLocale(locale string)
```

**Supported Locales:**
- `en` - English (default)
- `zh-CN` - Simplified Chinese
- `zh-TW` - Traditional Chinese
- `fr` - French
- `de` - German
- `es` - Spanish
- `it` - Italian
- `ja` - Japanese
- `ko` - Korean
- `pt` - Portuguese
- `ru` - Russian

**Example:**
```go
validator.SetLocale("zh-CN")
err := validator.Struct(&user)
// Error messages will be in Chinese
```

---

### RegisterValidation()

Register a custom validation rule.

```go
func RegisterValidation(tag string, fn ValidatorFunc) error
```

**Parameters:**
- `tag` - Validation tag name
- `fn` - Validation function

**Returns:**
- Error if registration fails

**Example:**
```go
func validatePassword(fl validator.FieldLevel) bool {
    password := fl.Field().String()
    return len(password) >= 8 &&
        strings.ContainsAny(password, "0123456789") &&
        strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") &&
        strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz")
}

validator.RegisterValidation("strong_password", validatePassword)

type User struct {
    Password string `validate:"required,strong_password"`
}
```

---

### RegisterTranslation()

Register a custom translation for validation errors.

```go
func RegisterTranslation(locale, tag, translation string)
```

**Parameters:**
- `locale` - Locale code
- `tag` - Validation tag
- `translation` - Translation template

**Example:**
```go
validator.RegisterTranslation("en", "custom_tag", "{field} must be custom")

type User struct {
    Field string `validate:"custom_tag"`
}
```

---

## Built-in Validation Rules

### Required Fields

```go
type User struct {
    Name  string `validate:"required"`
    Email string `validate:"required,email"`
}
```

### String Validation

```go
type User struct {
    Username string `validate:"required,min=3,max=20"`
    Bio      string `validate:"max=500"`
    Website  string `validate:"url"`
}
```

### Numeric Validation

```go
type Product struct {
    Price  float64 `validate:"required,gt=0"`
    Stock  int     `validate:"min=0,max=10000"`
    Rating float64 `validate:"min=0,max=5"`
}
```

### Date Validation

```go
type Event struct {
    StartDate time.Time `validate:"required"`
    EndDate   time.Time `validate:"required,gtfield=StartDate"`
}
```

### Custom Rules

```go
type User struct {
    Mobile  string `validate:"required,mobile"`
    IDCard  string `validate:"required,idcard"`
    BankCard string `validate:"required,bankcard"`
    Name    string `validate:"required,chinese_name"`
    Email   string `validate:"required,email"`
    URL     string `validate:"required,url"`
    IP      string `validate:"required,ipv4"`
    MAC     string `validate:"required,mac"`
    JSON    string `validate:"required,json"`
    UUID    string `validate:"required,uuid"`
}
```

---

## Usage Patterns

### Complete User Validation

```go
type User struct {
    Name     string `validate:"required,min=2,max=50,chinese_name"`
    Email    string `validate:"required,email"`
    Mobile   string `validate:"required,mobile"`
    IDCard   string `validate:"required,idcard"`
    BankCard string `validate:"required,bankcard"`
    Age      int    `validate:"required,min=18,max=120"`
    Password string `validate:"required,min=8,strong_password"`
}

func CreateUser(user *User) error {
    // Set locale for Chinese error messages
    validator.SetLocale("zh-CN")
    
    // Validate user
    if err := validator.Struct(user); err != nil {
        return fmt.Errorf("validation failed: %w", err)
    }
    
    // Create user in database
    return db.Create(user)
}
```

### Conditional Validation

```go
type Order struct {
    PaymentMethod string `validate:"required,oneof=credit_card paypal bank_transfer"`
    CardNumber   string `validate:"required_if=PaymentMethod credit_card,credit_card"`
    PayPalEmail  string `validate:"required_if=PaymentMethod paypal,email"`
}

func CreateOrder(order *Order) error {
    if err := validator.Struct(order); err != nil {
        return err
    }
    // Process order
    return nil
}
```

### Custom Validator with Parameters

```go
func validatePasswordComplexity(fl validator.FieldLevel) bool {
    password := fl.Field().String()
    param := fl.Param() // Get parameter from tag
    
    minLen := 8
    if param != "" {
        minLen, _ = strconv.Atoi(param)
    }
    
    if len(password) < minLen {
        return false
    }
    
    hasUpper := strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
    hasLower := strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz")
    hasDigit := strings.ContainsAny(password, "0123456789")
    
    return hasUpper && hasLower && hasDigit
}

validator.RegisterValidation("password_complexity", validatePasswordComplexity)

type User struct {
    Password string `validate:"required,password_complexity=12"`
}
```

---

## Advanced Features

### Custom Validator

```go
type CustomValidator struct {
    validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
    v := validator.New()
    
    // Register custom validators
    v.RegisterValidation("custom_rule", func(fl validator.FieldLevel) bool {
        value := fl.Field().String()
        // Custom validation logic
        return len(value) > 0
    })
    
    return &CustomValidator{validator: v}
}

func (cv *CustomValidator) Validate(s interface{}) error {
    return cv.validator.Struct(s)
}
```

### Multi-language Support

```go
func ValidateUser(user *User, lang string) error {
    // Set locale based on user preference
    validator.SetLocale(lang)
    
    if err := validator.Struct(user); err != nil {
        // Error messages will be in the specified language
        return err
    }
    
    return nil
}

// Example usage
ValidateUser(user, "en")  // English error messages
ValidateUser(user, "zh-CN")  // Chinese error messages
ValidateUser(user, "fr")  // French error messages
```

### Struct Level Validation

```go
type User struct {
    Password string `validate:"required,min=8"`
    Confirm  string `validate:"required,eqfield=Password"`
}

func (u User) Validate() error {
    return validator.Struct(&u)
}
```

---

## Best Practices

### Validation Tags

```go
type User struct {
    // Good: Specific and clear validation
    Name  string `validate:"required,min=2,max=50"`
    Email string `validate:"required,email"`
    
    // Bad: Too generic
    Name  string `validate:"required"`
    Email string `validate:"required"`
}
```

### Error Handling

```go
func HandleValidation(err error) error {
    if err == nil {
        return nil
    }
    
    // Check if it's a validation error
    if validationErrors, ok := err.(validator.ValidationErrors); ok {
        var messages []string
        for _, e := range validationErrors {
            messages = append(messages, fmt.Sprintf(
                "%s: %s",
                e.Field(),
                e.Tag(),
            ))
        }
        return fmt.Errorf("validation errors: %s", strings.Join(messages, "; "))
    }
    
    return err
}
```

### Performance

```go
// Good: Reuse validator instance
var v = validator.New()

func ValidateUser(user *User) error {
    return v.Struct(user)
}

// Avoid: Creating new validator each time
func ValidateUser(user *User) error {
    v := validator.New()  // Expensive
    return v.Struct(user)
}
```

---

## Related Documentation

- [must.go](/en/modules/must) - Error assertion
- [orm.go](/en/modules/orm) - Database operations
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
