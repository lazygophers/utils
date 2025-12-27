---
title: 快速開始
---

# 快速開始

本指南將幫助您快速開始使用 LazyGophers Utils。

## 安裝

使用 Go 模組安裝 LazyGophers Utils：

```bash
go get github.com/lazygophers/utils
```

## 基本用法

### 錯誤處理

LazyGophers Utils 提供了簡化的錯誤處理方式：

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils"
)

func main() {
    // 使用 Must 簡化錯誤處理
    data := utils.Must(loadData())
    fmt.Println(data)
}

func loadData() (string, error) {
    return "Hello, World!", nil
}
```

### 類型轉換

使用 `candy` 模組進行類型轉換：

```go
import "github.com/lazygophers/utils/candy"

// 字串轉整數
age := candy.ToInt("25")

// 字串轉布林值
active := candy.ToBool("true")

// 字串轉浮點數
price := candy.ToFloat("99.99")
```

### 時間處理

使用 `xtime` 模組處理時間：

```go
import "github.com/lazygophers/utils/xtime"

// 獲取當前日曆
cal := xtime.NowCalendar()

// 格式化日期
fmt.Printf("今天: %s\n", cal.String())

// 獲取農曆日期
fmt.Printf("農曆: %s\n", cal.LunarDate())

// 獲取生肖
fmt.Printf("生肖: %s\n", cal.Animal())

// 獲取節氣
fmt.Printf("節氣: %s\n", cal.CurrentSolarTerm())
```

### 配置管理

使用 `config` 模組加載配置：

```go
import "github.com/lazygophers/utils/config"

type Config struct {
    Database string `json:"database"`
    Port     int    `json:"port"`
    Debug    bool   `json:"debug"`
}

func main() {
    var cfg Config
    utils.MustSuccess(config.Load(&cfg, "config.json"))
    fmt.Printf("Config: %+v\n", cfg)
}
```

### 資料驗證

使用 `validator` 模組驗證資料：

```go
import "github.com/lazygophers/utils/validator"

type User struct {
    Name  string `validate:"required"`
    Email string `validate:"required,email"`
    Age   int    `validate:"min=0,max=150"`
}

func main() {
    user := User{
        Name:  "張三",
        Email: "zhangsan@example.com",
        Age:   25,
    }

    if err := utils.Validate(&user); err != nil {
        fmt.Printf("驗證失敗: %v\n", err)
    } else {
        fmt.Println("驗證成功")
    }
}
```

## 下一步

- 查看 [模組概覽](/zh-TW/modules/overview) 了解所有可用模組
- 閱讀 [API 文檔](/zh-TW/api/overview) 了解詳細 API
- 查看 [GitHub 倉庫](https://github.com/lazygophers/utils) 獲取更多示例

## 獲取幫助

- 📖 [完整 API 參考](https://pkg.go.dev/github.com/lazygophers/utils)
- 🐛 [提交問題](https://github.com/lazygophers/utils/issues)
- 💬 [GitHub Discussions](https://github.com/lazygophers/utils/discussions)
