---
title: 核心函数 - Validator
---

# 核心函数

Validator 的基础 API：结构体验证、单字段验证、语言设置、实例管理。

## Struct — 结构体验证

对结构体执行标签规则校验，返回 `ValidationErrors`。

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

## Var — 单字段验证

验证单个值，适合非结构体场景。

```go
func Var(field interface{}, tag string) error
```

```go
err := validator.Var("user@example.com", "required,email")
err := validator.Var(42, "min=0,max=100")
```

## SetLocale — 设置语言

设置全局错误消息语言，影响后续所有验证结果。

```go
func SetLocale(locale string)
```

支持：`en`（默认）、`zh`、`zh-TW`、`ja`、`ko`、`ar`、`de`、`es`、`fr`、`it`、`pt`、`ru`。

```go
validator.SetLocale("zh")
err := validator.Struct(&user)
// 错误消息为中文
```

## Validator 实例

### Default — 全局单例

```go
v := validator.Default()
err := v.Struct(&data)
```

### New — 自定义配置

```go
v, err := validator.New(
    validator.WithLocale("zh"),
    validator.WithUseJSON(true),
    validator.WithTranslations(map[string]string{
        "custom": "自定义消息",
    }),
    validator.WithCustomValidator("mytag", func(v interface{}) bool {
        return v.(string) != ""
    }),
)
```

### Config 结构

```go
type Config struct {
    Locale        string
    UseJSON       bool
    Translations  map[string]string
}
```

通过 `WithConfig` 一次性配置：

```go
v, _ := validator.New(validator.WithConfig(validator.Config{
    Locale:  "en",
    UseJSON: true,
}))
```
