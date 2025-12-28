---
title: validator - 資料驗證
---

# validator - 資料驗證

## 概述

validator 模組為 Go 結構體提供全面的資料驗證，支援多種語言環境和自定義驗證規則。它為常見資料類型提供內置驗證器，並允許註冊自定義驗證器。

## 核心函數

### Struct()

使用結構體標籤驗證結構體資料。

```go
func Struct(s interface{}) error
```

**參數：**
- `s` - 要驗證的結構體指針

**返回值：**
- 如果驗證失敗，返回驗證錯誤
- 如果驗證通過，返回 nil

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
    fmt.Printf("驗證失敗: %v\n", err)
}
```

---

### Var()

使用驗證標籤驗證單個變數。

```go
func Var(field interface{}, tag string) error
```

**參數：**
- `field` - 要驗證的值
- `tag` - 驗證規則

**返回值：**
- 如果驗證失敗，返回驗證錯誤
- 如果驗證通過，返回 nil

**示例：**
```go
if err := validator.Var("john@example.com", "required,email"); err != nil {
    fmt.Printf("無效的電子郵件: %v\n", err)
}

if err := validator.Var(25, "min=0,max=150"); err != nil {
    fmt.Printf("無效的年齡: %v\n", err)
}
```

---

### SetLocale()

設置錯誤訊息的語言環境。

```go
func SetLocale(locale string)
```

**支援的語言環境：**
- `en` - 英語（預設）
- `zh-CN` - 簡體中文
- `zh-TW` - 繁體中文
- `fr` - 法語
- `de` - 德語
- `es` - 西班牙語
- `it` - 義大利語
- `ja` - 日語
- `ko` - 韓語
- `pt` - 葡萄牙語
- `ru` - 俄語

**示例：**
```go
validator.SetLocale("zh-TW")
err := validator.Struct(&user)
// 錯誤訊息將是繁體中文
```

---

### RegisterValidation()

註冊自定義驗證規則。

```go
func RegisterValidation(tag string, fn ValidatorFunc) error
```

**參數：**
- `tag` - 驗證標籤名稱
- `fn` - 驗證函數

**返回值：**
- 如果註冊失敗，返回錯誤

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

為驗證錯誤註冊自定義翻譯。

```go
func RegisterTranslation(locale, tag, translation string)
```

**參數：**
- `locale` - 語言環境代碼
- `tag` - 驗證標籤
- `translation` - 翻譯模板

**示例：**
```go
validator.RegisterTranslation("en", "custom_tag", "{field} 必須是自定義的")

type User struct {
    Field string `validate:"custom_tag"`
}
```

---

## 內置驗證規則

### 必填欄位

```go
type User struct {
    Name  string `validate:"required"`
    Email string `validate:"required,email"`
}
```

### 字串驗證

```go
type User struct {
    Username string `validate:"required,min=3,max=20"`
    Bio      string `validate:"max=500"`
    Website  string `validate:"url"`
}
```

### 數值驗證

```go
type Product struct {
    Price  float64 `validate:"required,gt=0"`
    Stock  int     `validate:"min=0,max=10000"`
    Rating float64 `validate:"min=0,max=5"`
}
```

### 日期驗證

```go
type Event struct {
    StartDate time.Time `validate:"required"`
    EndDate   time.Time `validate:"required,gtfield=StartDate"`
}
```

### 自定義規則

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

### 完整的使用者驗證

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
    // 設置語言環境以獲取繁體中文錯誤訊息
    validator.SetLocale("zh-TW")
    
    // 驗證使用者
    if err := validator.Struct(user); err != nil {
        return fmt.Errorf("驗證失敗: %w", err)
    }
    
    // 在資料庫中建立使用者
    return db.Create(user)
}
```

### 條件驗證

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
    // 處理訂單
    return nil
}
```

### 帶參數的自定義驗證器

```go
func validatePasswordComplexity(fl validator.FieldLevel) bool {
    password := fl.Field().String()
    param := fl.Param() // 從標籤獲取參數
    
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

## 高級功能

### 自定義驗證器

```go
type CustomValidator struct {
    validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
    v := validator.New()
    
    // 註冊自定義驗證器
    v.RegisterValidation("custom_rule", func(fl validator.FieldLevel) bool {
        value := fl.Field().String()
        // 自定義驗證邏輯
        return len(value) > 0
    })
    
    return &CustomValidator{validator: v}
}

func (cv *CustomValidator) Validate(s interface{}) error {
    return cv.validator.Struct(s)
}
```

### 多語言支援

```go
func ValidateUser(user *User, lang string) error {
    // 根據使用者偏好設置語言環境
    validator.SetLocale(lang)
    
    if err := validator.Struct(user); err != nil {
        // 錯誤訊息將使用指定的語言
        return err
    }
    
    return nil
}

// 示例用法
ValidateUser(user, "en")  // 英文錯誤訊息
ValidateUser(user, "zh-TW")  // 繁體中文錯誤訊息
ValidateUser(user, "fr")  // 法文錯誤訊息
```

### 結構體級別驗證

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

## 最佳實踐

### 驗證標籤

```go
type User struct {
    // 好：具體且清晰的驗證
    Name  string `validate:"required,min=2,max=50"`
    Email string `validate:"required,email"`
    
    // 壞：過於通用
    Name  string `validate:"required"`
    Email string `validate:"required"`
}
```

### 錯誤處理

```go
func HandleValidation(err error) error {
    if err == nil {
        return nil
    }
    
    // 檢查是否為驗證錯誤
    if validationErrors, ok := err.(validator.ValidationErrors); ok {
        var messages []string
        for _, e := range validationErrors {
            messages = append(messages, fmt.Sprintf(
                "%s: %s",
                e.Field(),
                e.Tag(),
            ))
        }
        return fmt.Errorf("驗證錯誤: %s", strings.Join(messages, "; "))
    }
    
    return err
}
```

### 效能

```go
// 好：重用驗證器實例
var v = validator.New()

func ValidateUser(user *User) error {
    return v.Struct(user)
}

// 避免：每次都建立新驗證器
func ValidateUser(user *User) error {
    v := validator.New()  // 昂貴
    return v.Struct(user)
}
```

---

## 相關文檔

- [must.go](/zh-TW/modules/must) - 錯誤斷言
- [orm.go](/zh-TW/modules/orm) - 資料庫操作
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
