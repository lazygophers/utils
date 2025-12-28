---
title: 入门指南
---

# 入门指南

本指南将帮助您快速开始使用 LazyGophers Utils。

## 安装

使用 Go 模块安装 LazyGophers Utils：

```bash
go get github.com/lazygophers/utils
```

## 基本用法

### 错误处理

LazyGophers Utils 提供简化的错误处理：

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils"
)

func main() {
    // 使用 Must 简化错误处理
    data := utils.Must(loadData())
    fmt.Println(data)
}

func loadData() (string, error) {
    return "Hello, World!", nil
}
```

### 类型转换

使用 `candy` 模块进行类型转换：

```go
import "github.com/lazygophers/utils/candy"

// 字符串转整数
age := candy.ToInt("25")

// 字符串转布尔值
active := candy.ToBool("true")

// 字符串转浮点数
price := candy.ToFloat("99.99")
```

### 时间处理

使用 `xtime` 模块进行时间处理：

```go
import "github.com/lazygophers/utils/xtime"

// 获取当前日历
cal := xtime.NowCalendar()

// 格式化日期
fmt.Printf("Today: %s\n", cal.String())

// 获取农历日期
fmt.Printf("Lunar: %s\n", cal.LunarDate())

// 获取生肖
fmt.Printf("Animal: %s\n", cal.Animal())

// 获取节气
fmt.Printf("Solar Term: %s\n", cal.CurrentSolarTerm())
```

### 配置管理

使用 `config` 模块加载配置：

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

### 数据验证

使用 `validator` 模块验证数据：

```go
import "github.com/lazygophers/utils/validator"

type User struct {
    Name  string `validate:"required"`
    Email string `validate:"required,email"`
    Age   int    `validate:"min=0,max=150"`
}

func main() {
    user := User{
        Name:  "John Doe",
        Email: "john@example.com",
        Age:   25,
    }

    if err := utils.Validate(&user); err != nil {
        fmt.Printf("Validation failed: %v\n", err)
    } else {
        fmt.Println("Validation successful")
    }
}
```

## 下一步

-   查看 [模块概览](/zh-CN/modules/overview) 了解所有可用模块
-   阅读 [API 文档](/zh-CN/api/overview) 获取详细的 API 信息
-   访问 [GitHub 仓库](https://github.com/lazygophers/utils) 查看更多示例

## 获取帮助

-   📖 [完整 API 参考](https://pkg.go.dev/github.com/lazygophers/utils)
-   🐛 [报告问题](https://github.com/lazygophers/utils/issues)
-   💬 [GitHub 讨论](https://github.com/lazygophers/utils/discussions)
