---
title: validator - 数据验证
---

# validator - 数据验证

## 概述

validator 模块为 Go 结构体提供全面的数据验证，支持多种语言环境和自定义验证规则。它为常见数据类型提供内置验证器，并允许注册自定义验证器。

## 核心函数

### Struct()

使用结构体标签验证结构体数据。

```go
func Struct(s interface{}) error
```

**参数：**
- `s` - 要验证的结构体指针

**返回值：**
- 如果验证失败，返回验证错误
- 如果验证通过，返回 nil

**示例：**
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
    fmt.Printf("验证失败: %v\n", err)
}
```

---

### Var()

使用验证标签验证单个变量。

```go
func Var(field interface{}, tag string) error
```

**参数：**
- `field` - 要验证的值
- `tag` - 验证规则

**返回值：**
- 如果验证失败，返回验证错误
- 如果验证通过，返回 nil

**示例：**
```go
if err := validator.Var("john@example.com", "required,email"); err != nil {
    fmt.Printf("无效的邮箱: %v\n", err)
}

if err := validator.Var(25, "min=0,max=150"); err != nil {
    fmt.Printf("无效的年龄: %v\n", err)
}
```

---

### SetLocale()

设置错误消息的语言环境。

```go
func SetLocale(locale string)
```

**支持的语言环境：**
- `en` - 英语（默认）
- `zh-CN` - 简体中文
- `zh-TW` - 繁体中文
- `fr` - 法语
- `de` - 德语
- `es` - 西班牙语
- `it` - 意大利语
- `ja` - 日语
- `ko` - 韩语
- `pt` - 葡萄牙语
- `ru` - 俄语

**示例：**
```go
validator.SetLocale("zh-CN")
err := validator.Struct(&user)
// 错误消息将是中文
```

---

### RegisterValidation()

注册自定义验证规则。

```go
func RegisterValidation(tag string, fn ValidatorFunc) error
```

**参数：**
- `tag` - 验证标签名称
- `fn` - 验证函数

**返回值：**
- 如果注册失败，返回错误

**示例：**
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

为验证错误注册自定义翻译。

```go
func RegisterTranslation(locale, tag, translation string)
```

**参数：**
- `locale` - 语言环境代码
- `tag` - 验证标签
- `translation` - 翻译模板

**示例：**
```go
validator.RegisterTranslation("en", "custom_tag", "{field} 必须是自定义的")

type User struct {
    Field string `validate:"custom_tag"`
}
```

---

## 内置验证规则

### 必填字段

```go
type User struct {
    Name  string `validate:"required"`
    Email string `validate:"required,email"`
}
```

### 字符串验证

```go
type User struct {
    Username string `validate:"required,min=3,max=20"`
    Bio      string `validate:"max=500"`
    Website  string `validate:"url"`
}
```

### 数值验证

```go
type Product struct {
    Price  float64 `validate:"required,gt=0"`
    Stock  int     `validate:"min=0,max=10000"`
    Rating float64 `validate:"min=0,max=5"`
}
```

### 日期验证

```go
type Event struct {
    StartDate time.Time `validate:"required"`
    EndDate   time.Time `validate:"required,gtfield=StartDate"`
}
```

### 自定义规则

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

## 使用模式

### 完整的用户验证

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
    // 设置语言环境以获取中文错误消息
    validator.SetLocale("zh-CN")
    
    // 验证用户
    if err := validator.Struct(user); err != nil {
        return fmt.Errorf("验证失败: %w", err)
    }
    
    // 在数据库中创建用户
    return db.Create(user)
}
```

### 条件验证

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
    // 处理订单
    return nil
}
```

### 带参数的自定义验证器

```go
func validatePasswordComplexity(fl validator.FieldLevel) bool {
    password := fl.Field().String()
    param := fl.Param() // 从标签获取参数
    
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

## 高级功能

### 自定义验证器

```go
type CustomValidator struct {
    validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
    v := validator.New()
    
    // 注册自定义验证器
    v.RegisterValidation("custom_rule", func(fl validator.FieldLevel) bool {
        value := fl.Field().String()
        // 自定义验证逻辑
        return len(value) > 0
    })
    
    return &CustomValidator{validator: v}
}

func (cv *CustomValidator) Validate(s interface{}) error {
    return cv.validator.Struct(s)
}
```

### 多语言支持

```go
func ValidateUser(user *User, lang string) error {
    // 根据用户偏好设置语言环境
    validator.SetLocale(lang)
    
    if err := validator.Struct(user); err != nil {
        // 错误消息将使用指定的语言
        return err
    }
    
    return nil
}

// 示例用法
ValidateUser(user, "en")  // 英文错误消息
ValidateUser(user, "zh-CN")  // 中文错误消息
ValidateUser(user, "fr")  // 法文错误消息
```

### 结构体级别验证

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

## 最佳实践

### 验证标签

```go
type User struct {
    // 好：具体且清晰的验证
    Name  string `validate:"required,min=2,max=50"`
    Email string `validate:"required,email"`
    
    // 坏：过于通用
    Name  string `validate:"required"`
    Email string `validate:"required"`
}
```

### 错误处理

```go
func HandleValidation(err error) error {
    if err == nil {
        return nil
    }
    
    // 检查是否为验证错误
    if validationErrors, ok := err.(validator.ValidationErrors); ok {
        var messages []string
        for _, e := range validationErrors {
            messages = append(messages, fmt.Sprintf(
                "%s: %s",
                e.Field(),
                e.Tag(),
            ))
        }
        return fmt.Errorf("验证错误: %s", strings.Join(messages, "; "))
    }
    
    return err
}
```

### 性能

```go
// 好：重用验证器实例
var v = validator.New()

func ValidateUser(user *User) error {
    return v.Struct(user)
}

// 避免：每次都创建新验证器
func ValidateUser(user *User) error {
    v := validator.New()  // 昂贵
    return v.Struct(user)
}
```

---

## 相关文档

- [must.go](/zh-CN/modules/must) - 错误断言
- [orm.go](/zh-CN/modules/orm) - 数据库操作
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
