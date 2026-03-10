---
title: 快速开始
---

# 快速开始

这一页只做一件事：帮助你在最短时间内确认导入方式、理解常见入口，并知道接下来该读哪里。

## 安装

```bash
go get github.com/lazygophers/utils
```

## 导入策略

这个仓库更适合按需导入，而不是一次性把所有能力都放进根包：

- 根包 `github.com/lazygophers/utils`：少量通用辅助，例如 `Must`、`MustSuccess`、`Scan`、`Value`。
- 子包：面向明确主题，如 `candy`、`config`、`xtime`、`validator`、`cache/...`。

## 一分钟上手示例

```go
package main

import (
    "fmt"

    utils "github.com/lazygophers/utils"
    "github.com/lazygophers/utils/candy"
    "github.com/lazygophers/utils/config"
    "github.com/lazygophers/utils/validator"
    "github.com/lazygophers/utils/xtime"
)

type AppConfig struct {
    Name string `json:"name" validate:"required"`
    Port int    `json:"port" validate:"min=1,max=65535"`
}

func main() {
    // 1. 初始化阶段可以用 Must 系列快速失败
    fmt.Println(utils.Must(loadMessage()))

    // 2. candy 负责常见转换与集合辅助
    fmt.Println(candy.ToInt("8080"))

    // 3. config 负责多格式配置加载
    var cfg AppConfig
    utils.MustSuccess(config.LoadConfig(&cfg, "config.yaml"))

    // 4. validator 负责结构体验证
    utils.MustSuccess(validator.Struct(&cfg))

    // 5. xtime 提供日历、农历和节气信息
    cal := xtime.NowCalendar()
    fmt.Println(cal.String())
    fmt.Println(cal.LunarDate())
}

func loadMessage() (string, error) {
    return "LazyGophers Utils", nil
}
```

## 从哪里继续

### 如果你已经知道问题类型

- 错误快速失败：看 [must](/modules/core/must)
- 数据库 JSON 字段：看 [orm](/modules/core/orm)
- 配置：看 [config](/modules/system/config)
- 时间与农历：看 [xtime](/modules/time/xtime)
- 缓存选型：看 [缓存策略](/modules/cache/)

### 如果你还不确定该用哪个包

先看 [模块总览](/modules/overview)。它按场景把整个仓库拆成了八个主题分组，比直接浏览目录更容易定位。

## 使用这套库时的两个建议

1. **初始化路径和业务路径分开看**：`Must` 很适合启动阶段，不适合替代所有业务错误处理。
2. **优先按包职责理解，而不是只看函数名**：比如 `xtime` 并不是简单的时间格式化，它包含日历、农历与排班规则；`cache` 也不是一个实现，而是一组策略。
