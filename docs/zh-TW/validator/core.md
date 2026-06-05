---
title: 核心函數 - Validator
---

# 核心函數

Validator 的基礎 API：結構體驗證、單欄位驗證、語言設定、實例管理。

## Struct — 結構體驗證

對結構體執行標籤規則校驗，返回 `ValidationErrors`。

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

## Var — 單欄位驗證

驗證單一值，適合非結構體場景。

```go
func Var(field interface{}, tag string) error
```

```go
err := validator.Var("user@example.com", "required,email")
err := validator.Var(42, "min=0,max=100")
```

## SetLocale — 設定語言

設定全域錯誤訊息語言，影響後續所有驗證結果。

```go
func SetLocale(locale string)
```

支援：`en`（預設）、`zh`、`zh-TW`、`ja`、`ko`、`ar`、`de`、`es`、`fr`、`it`、`pt`、`ru`。

```go
validator.SetLocale("zh-TW")
err := validator.Struct(&user)
// 錯誤訊息為繁體中文
```

## Validator 實例

### Default — 全域單例

```go
v := validator.Default()
err := v.Struct(&data)
```

### New — 自訂配置

```go
v, err := validator.New(
    validator.WithLocale("zh-TW"),
    validator.WithUseJSON(true),
    validator.WithTranslations(map[string]string{
        "custom": "自訂訊息",
    }),
    validator.WithCustomValidator("mytag", func(v interface{}) bool {
        return v.(string) != ""
    }),
)
```

### Config 結構

```go
type Config struct {
    Locale        string
    UseJSON       bool
    Translations  map[string]string
}
```

透過 `WithConfig` 一次性配置：

```go
v, _ := validator.New(validator.WithConfig(validator.Config{
    Locale:  "zh-TW",
    UseJSON: true,
}))
```
