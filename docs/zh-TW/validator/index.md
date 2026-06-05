---
title: Validator - 資料驗證
---

# Validator

Go 結構體資料驗證庫。支援標籤規則、自訂驗證器、12 種語言錯誤訊息。

## 安裝

```bash
go get github.com/lazygophers/utils/validator
```

## 快速開始

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
            fmt.Printf("欄位 %s 不滿足規則 %s\n", fe.Field(), fe.Tag())
        }
    }
}
```

## 文件導航

| 主題 | 說明 |
|------|------|
| [核心函數](/zh-TW/validator/core) | Struct、Var、SetLocale 等基礎 API |
| [內建規則](/zh-TW/validator/rules) | required、email、min 等全部內建驗證標籤 |
| [自訂驗證器](/zh-TW/validator/custom) | 註冊自訂驗證器和翻譯 |
| [驗證引擎](/zh-TW/validator/engine) | Engine、FieldLevel、StructLevel 進階場景 |
| [錯誤處理](/zh-TW/validator/errors) | ValidationErrors、FieldError 型別 |
| [多語言](/zh-TW/validator/i18n) | 12 種語言支援與切換 |
| [效能與最佳實踐](/zh-TW/validator/performance) | 效能最佳化與使用建議 |
