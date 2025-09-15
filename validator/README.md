# Validator 包

一个功能强大的验证器包，兼容 `github.com/go-playground/validator/v10` 的协议，支持更多验证规则和多语言错误消息。

## 特性

- ✅ 完全兼容 `github.com/go-playground/validator/v10` API
- ✅ 支持多语言错误消息（中文、英文等）
- ✅ 优先使用 JSON tag 作为字段名
- ✅ 内置常用的中国本土化验证规则（手机号、身份证、银行卡等）
- ✅ 强类型错误处理
- ✅ 丰富的错误处理方法
- ✅ 支持自定义验证规则和翻译
- ✅ 高性能，线程安全

## 快速开始

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/validator"
)

type User struct {
    Name     string `json:"name" validate:"required,min=2,max=50"`
    Email    string `json:"email" validate:"required,email"`
    Age      int    `json:"age" validate:"required,min=18,max=120"`
    Phone    string `json:"phone" validate:"required,mobile"`
    Password string `json:"password" validate:"required,strong_password"`
}

func main() {
    // 创建验证器
    v := validator.New()
    
    user := User{
        Name:     "张三",
        Email:    "zhangsan@example.com", 
        Age:      25,
        Phone:    "13812345678",
        Password: "MyPass123!",
    }
    
    // 验证
    err := v.Struct(user)
    if err != nil {
        fmt.Printf("验证失败: %s\n", err)
        return
    }
    
    fmt.Println("验证通过")
}
```

### 多语言支持

```go
// 中文验证器
zhValidator := validator.New(validator.WithLocale("zh"))

// 英文验证器  
enValidator := validator.New(validator.WithLocale("en"))

user := User{Name: ""}

// 中文错误信息
err := zhValidator.Struct(user)
fmt.Println(err) // name不能为空

// 英文错误信息
err = enValidator.Struct(user) 
fmt.Println(err) // name is required
```

### 全局函数

```go
// 设置全局配置
validator.SetLocale("zh")
validator.SetUseJSON(true)

// 使用全局函数
err := validator.Struct(user)
err = validator.Var("test@example.com", "email")
```

## 内置验证规则

### 标准验证规则

支持所有 `validator/v10` 的标准规则：
- `required` - 必填
- `email` - 邮箱格式
- `url` - URL格式
- `min`, `max` - 最小值/最大值
- `len` - 长度
- `oneof` - 枚举值
- `alpha`, `alphanum`, `numeric` - 字符类型
- 等等...

### 中国本土化验证规则

#### `mobile` - 手机号验证
```go
Phone string `validate:"mobile"`
// 验证中国大陆手机号：1[3-9]\d{9}
```

#### `idcard` - 身份证验证
```go
IDCard string `validate:"idcard"`
// 支持15位和18位身份证，包含校验码验证
```

#### `bankcard` - 银行卡验证
```go
BankCard string `validate:"bankcard"`
// 使用Luhn算法验证银行卡号
```

#### `chinese_name` - 中文姓名验证
```go
Name string `validate:"chinese_name"`
// 验证2-4个中文字符的姓名，支持少数民族姓名（·）
```

#### `strong_password` - 强密码验证
```go
Password string `validate:"strong_password"`
// 至少8位，包含大小写字母、数字、特殊字符中的至少3种
```

## 错误处理

### ValidationErrors 类型

```go
type ValidationErrors []*FieldError

// 方法
First() *FieldError                    // 获取第一个错误
FirstError() string                    // 获取第一个错误消息
ByField(field string) *FieldError      // 根据字段获取错误
HasField(field string) bool            // 检查字段是否有错误
Fields() []string                      // 获取所有错误字段
Messages() []string                    // 获取所有错误消息
ToMap() map[string]string              // 转换为字段->消息映射
ToDetailMap() map[string]map[string]interface{} // 详细错误信息
JSON() map[string]interface{}          // JSON格式错误信息
```

### FieldError 类型

```go
type FieldError struct {
    Field       string      // 字段名（优先JSON tag）
    Tag         string      // 验证标签
    Value       interface{} // 字段值
    Param       string      // 验证参数
    Message     string      // 错误消息
    // ...
}
```

### 错误处理示例

```go
err := v.Struct(user)
if err != nil {
    if validationErrors, ok := err.(validator.ValidationErrors); ok {
        // 获取第一个错误
        fmt.Println("第一个错误:", validationErrors.FirstError())
        
        // 检查特定字段
        if validationErrors.HasField("email") {
            emailErr := validationErrors.ByField("email")
            fmt.Printf("邮箱错误: %s\n", emailErr.Message)
        }
        
        // 转换为映射用于API响应
        errMap := validationErrors.ToMap()
        // {"email": "email必须是有效的邮箱地址", "name": "name不能为空"}
        
        // 获取详细信息
        detailMap := validationErrors.ToDetailMap()
        
        // 自定义格式化
        formatted := validationErrors.Format("{field}: {message}")
    }
}
```

## 配置选项

### WithLocale(locale string)
设置语言地区
```go
v := validator.New(validator.WithLocale("zh"))
```

### WithUseJSON(useJSON bool)  
设置是否优先使用JSON字段名
```go
v := validator.New(validator.WithUseJSON(true))
```

### WithConfig(config Config)
批量配置
```go
config := validator.Config{
    Locale:  "zh",
    UseJSON: true,
    Translations: map[string]string{
        "zh.custom": "自定义错误消息",
    },
}
v := validator.New(validator.WithConfig(config))
```

## 自定义验证

### 注册自定义验证规则

```go
v := validator.New()

// 注册验证函数
err := v.RegisterValidation("custom_tag", func(fl validator.FieldLevel) bool {
    value := fl.Field().String()
    return value == "expected_value"
})

// 注册翻译
v.RegisterTranslation("zh", "custom_tag", "{field}必须是期望的值")
v.RegisterTranslation("en", "custom_tag", "{field} must be expected value")
```

### 使用选项注册

```go
v := validator.New(
    validator.WithCustomValidator("even_number", func(value interface{}) bool {
        if num, ok := value.(int); ok {
            return num%2 == 0
        }
        return false
    }),
)
```

## JSON 字段名优先级

当 `UseJSON` 为 `true` 时（默认），验证器会优先使用 `json` tag 作为字段名：

```go
type User struct {
    UserName string `json:"user_name" validate:"required"` // 错误消息中显示 "user_name"
    Email    string `validate:"required"`                  // 错误消息中显示 "Email"
}
```

## 支持的语言

- `en` - English
- `zh` - 中文简体
- `zh-CN` - 中文（中国）

可以通过实现 `LocaleConfig` 添加更多语言支持。

## 性能

该包基于 `validator/v10` 构建，保持了原有的高性能特性：
- 使用反射缓存
- 并发安全
- 零内存分配的路径优化

## 迁移指南

### 从 validator/v10 迁移

直接替换导入即可，API完全兼容：

```go
// 原来
import "github.com/go-playground/validator/v10"

// 现在  
import "github.com/lazygophers/utils/validator"
```

### 从项目原有的 validate.go 迁移

```go
// 原来
utils.Validate(user)

// 现在
validator.Struct(user)
```

## 构建标签

支持条件编译，可以选择性包含语言支持：

```bash
# 包含中文支持
go build -tags validator_zh

# 包含所有语言支持  
go build -tags validator_all

# 默认只包含英文
go build
```