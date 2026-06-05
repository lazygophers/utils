---
title: Validator - 数据验证
---

# Validator

Go 结构体数据验证库。支持标签规则、自定义验证器、12 种语言错误消息。

## 安装

```bash
go get github.com/lazygophers/utils/validator
```

## 快速开始

```go
import "github.com/lazygophers/utils/validator"

type User struct {
    Name     string `validate:"required,min=2,max=50"`
    Email    string `validate:"required,email"`
    Age      int    `validate:"min=0,max=150"`
    Password string `validate:"required,strong_password"`
}

err := validator.Struct(&user)
if err != nil {
    if ve, ok := err.(validator.ValidationErrors); ok {
        for _, fe := range ve {
            fmt.Printf("字段 %s 不满足规则 %s\n", fe.Field(), fe.Tag())
        }
    }
}
```

## 文档导航

| 主题 | 说明 |
|------|------|
| [核心函数](/validator/core) | Struct、Var、SetLocale 等基础 API |
| [内置规则](/validator/rules) | required、email、min 等全部内置验证标签 |
| [自定义验证器](/validator/custom) | 注册自定义验证器和翻译 |
| [验证引擎](/validator/engine) | Engine、FieldLevel、StructLevel 高级场景 |
| [错误处理](/validator/errors) | ValidationErrors、FieldError 类型 |
| [多语言](/validator/i18n) | 12 种语言支持与切换 |
| [性能与最佳实践](/validator/performance) | 性能优化与使用建议 |
