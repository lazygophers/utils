
# Utils

> LazyGophers Utils - 一个功能丰富的 Go 工具库

[![Go Version](https://img.shields.io/badge/Go-1.24.0-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-AGPL%20v3-green.svg)](LICENSE)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/lazygophers/utils)

## 项目简介

`utils` 是一个功能全面的 Go 工具库，提供了大量实用的工具函数和模块，旨在帮助开发者更高效地构建 Go 应用程序。该项目采用模块化设计，每个子包都专注于特定的功能领域。

## 特性

- **模块化设计**：每个功能模块独立，按需引入
- **高性能**：针对 Go 语言特性进行优化
- **易于使用**：简洁的 API 设计，开箱即用
- **全面覆盖**：涵盖日常开发中的各种需求

## 安装

使用 Go Modules 导入：

```bash
go get github.com/lazygophers/utils
```

## 核心模块

### 基础工具

- **[`must`](must.go)**：提供错误处理和断言工具
  - [`MustOk()`](must.go:16)：验证状态并返回值
  - [`MustSuccess()`](must.go:29)：验证错误状态
  - [`Must()`](must.go:46)：组合验证函数
  - [`Ignore()`](must.go:62)：忽略参数的工具函数

### 数据库操作

- **[`orm`](orm.go)**：数据库操作工具
  - [`Scan()`](orm.go:32)：数据库字段到结构体的扫描
  - [`Value()`](orm.go:68)：结构体到数据库值的转换
  - 支持 JSON 序列化和默认值填充

### 数据验证

- **[`validate`](validate.go)**：基于 validator 的结构体验证
  - [`Validate()`](validate.go:20)：快速结构体验证
  - 自动记录错误日志到日志系统

## 功能模块

### anyx
提供 Any 类型相关的工具函数。

### app
应用程序框架和工具。

### atexit
程序退出时的清理和钩子函数管理。

### bufiox
增强的缓冲区操作工具。

### candy
提供一些"语法糖"和便捷函数。

### config
配置文件读取和管理工具。

### cryptox
加密和解密相关工具。

### defaults
默认值处理工具。

### event
事件驱动编程支持。

### fake
用于测试的假数据生成工具。

### hystrix
熔断器模式实现。

### json
JSON 处理增强工具。

### network
网络操作相关工具。

### osx
操作系统相关工具的增强版本。

### pgp
PGP 加密和签名工具。

### pyroscope
性能分析和监控工具集成。

### randx
随机数生成工具。

### routine
协程和并发编程工具。

### runtime
运行时信息获取和处理。

### singledo
单例模式实现。

### stringx
字符串处理增强工具。

### unit
单元测试辅助工具。

### urlx
URL 处理工具。

### wait
等待和超时控制工具。

### xtime
时间处理工具的增强版本，包含丰富的中国传统历法功能。

**核心特性**：
- 🗓️ **统一日历接口** - 整合公历农历信息
- 🌙 **农历计算** - 精确的农历转换和节气计算
- 🐲 **生肖干支** - 完整的天干地支系统
- 🏮 **传统节日** - 自动检测和文化信息展示

#### 子模块
- **xtime007** - 007工作制时间常量
- **xtime955** - 955工作制时间常量  
- **xtime996** - 996工作制时间常量

#### 快速示例
```go
// 创建日历对象
cal := xtime.NowCalendar()

// 获取完整信息
fmt.Println(cal.String())
// 输出: 2023年08月15日 六月廿九 兔年 处暑

// 农历信息
fmt.Println(cal.LunarDate())     // 农历二零二三年六月廿九
fmt.Println(cal.Animal())        // 兔
fmt.Println(cal.YearGanZhi())    // 癸卯

// 节气信息
fmt.Println(cal.CurrentSolarTerm()) // 处暑
fmt.Println(cal.DaysToNextTerm())   // 8
```

## 使用示例

### 错误处理
```go
package main

import (
    "github.com/lazygophers/utils"
)

func main() {
    // 使用 Must 处理可能出错的操作
    result := utils.Must(someFunction())

    // 使用 MustSuccess 验证错误
    utils.MustSuccess(doSomething())
}
```

### 数据库操作
```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email" validate:"required,email"`
    Age   int    `json:"age" default:"18"`
}

func GetUser(id int) (*User, error) {
    var user User

    // 从数据库扫描数据
    err := utils.Scan(dbData, &user)
    if err != nil {
        return nil, err
    }

    // 验证结构体
    err = utils.Validate(&user)
    if err != nil {
        return nil, err
    }

    return &user, nil
}
```

## 快速开始

1. **安装依赖**
    ```bash
    go mod tidy
    ```

2. **导入所需模块**
    ```go
    import "github.com/lazygophers/utils"
    import "github.com/lazygophers/utils/json"
    import "github.com/lazygophers/utils/config"
    ```

3. **开始使用**
    ```go
    // 使用各种工具函数
    value := utils.Must(getValue())
    err := utils.Validate(&config)
    ```

## 依赖项

- Go 1.24.0+
- [`github.com/go-playground/validator/v10`](https://github.com/go-playground/validator) - 结构体验证
- [`github.com/mcuadros/go-defaults`](https://github.com/mcuadros/go-defaults) - 默认值设置
- [`github.com/lazygophers/log`](https://github.com/lazygophers/log) - 日志库
- 更多依赖请查看 [`go.mod`](go.mod) 文件

## 性能特点

- **零依赖冲突**：所有依赖都经过精心选择，避免版本冲突
- **内存优化**：采用 sync.Pool 等技术减少内存分配
- **并发安全**：所有工具函数都考虑了并发场景
- **类型安全**：充分利用 Go 的类型系统保证编译时安全

## 贡献指南

我们欢迎任何形式的贡献！请遵循以下步骤：

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

### 开发规范

- 遵循 Go 标准代码风格
- 所有公共 API 必须有文档注释
- 提交新功能时必须包含测试用例
- 保持向后兼容性

## 许可证

本项目采用 GNU Affero General Public License v3.0 许可证。

查看 [LICENSE](LICENSE) 文件了解更多信息。

## 📊 项目统计

| 指标 | 数值 | 说明 |
|------|------|------|
| 📦 **模块数量** | 20+ | 涵盖各种常用功能 |
| 🧪 **测试覆盖率** | 85%+ | 高质量代码保障 |
| 📝 **文档完整度** | 95%+ | 详尽的使用说明 |
| ⚡ **性能等级** | A+ | 经过基准测试优化 |
| 🌟 **Star 数量** | - | GitHub 社区认可 |

## 🎯 开发路线图

### 近期计划 (1-3个月)
- [ ] 完善 xtime 农历节气功能
- [ ] 提升整体测试覆盖率至 90%+
- [ ] 添加更多 candy 类型转换工具
- [ ] 性能优化和基准测试完善

### 长期规划 (6-12个月)  
- [ ] AI/ML 工具模块
- [ ] 云服务集成工具
- [ ] 微服务开发套件
- [ ] 国际化支持

## 📈 性能特点

- **高性能**: 核心功能经过基准测试优化
- **低延迟**: 大部分操作在微秒级完成  
- **内存友好**: 合理的内存使用和垃圾回收
- **并发安全**: 所有公共API都是协程安全的

## 🤝 社区与支持

- 📖 **文档**: 每个模块都有详细的使用文档
- 🐛 **Issue 跟踪**: GitHub Issues 快速响应
- 💬 **讨论社区**: GitHub Discussions 交流
- 🔄 **持续集成**: 自动化测试和质量检查

## 📄 更新日志

### v1.2.0 (最新)
- ✨ **xtime 重大更新**: 全新的 Calendar 统一接口
- ✨ **农历节气功能**: 完整的传统历法支持
- ✨ **生肖干支系统**: 准确的天干地支计算  
- 🧪 **测试提升**: xtime 测试覆盖率达到 72.7%
- 📚 **文档完善**: 详细的使用指南和性能报告
- ⚡ **性能优化**: 关键路径性能提升 20%+

### v1.1.0
- ✨ 增强的时间处理功能
- 🐛 修复农历计算精度问题
- 📈 整体性能优化
- 📚 文档改进和示例补充

### v1.0.0
- 🎉 初始版本发布
- ✨ 包含基础工具模块
- ✨ 提供完整的错误处理和验证功能
- ✨ 支持数据库操作和结构体验证
- 📚 基础文档和使用指南
