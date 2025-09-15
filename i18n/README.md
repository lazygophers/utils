# I18n - 国际化支持

> 🌍 统一的多语言支持框架，为整个项目提供国际化功能

[English](README.md) • [简体中文](README_zh_CN.md)

## 概览

`i18n` 包是 LazyGophers Utils 的全局国际化支持框架，提供统一的多语言管理和本地化功能。

## 支持的语言

| 语言 | 代码 | 状态 | Build Tag |
|------|------|------|-----------|
| 英语 | `en` | ✅ 默认 | 无需 |
| 简体中文 | `zh-CN` | ✅ | `i18n_zh_cn` |
| 繁体中文 | `zh-TW` | ✅ | `i18n_zh_tw` |
| 日语 | `ja` | ✅ | `i18n_ja` |
| 韩语 | `ko` | ✅ | `i18n_ko` |
| 法语 | `fr` | ✅ | `i18n_fr` |
| 西班牙语 | `es` | ✅ | `i18n_es` |
| 阿拉伯语 | `ar` | ✅ | `i18n_ar` |
| 俄语 | `ru` | ✅ | `i18n_ru` |
| 意大利语 | `it` | ✅ | `i18n_it` |
| 葡萄牙语 | `pt` | ✅ | `i18n_pt` |
| 德语 | `de` | ✅ | `i18n_de` |

## 快速开始

### 基本用法

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/i18n"
)

func main() {
    // 设置默认语言
    i18n.SetDefaultLocale(i18n.ChineseSimplified)
    
    // 翻译消息
    msg := i18n.Translate(i18n.ChineseSimplified, "error")
    fmt.Println(msg) // 输出: 错误
    
    // 使用默认语言翻译
    msg2 := i18n.TranslateDefault("loading")
    fmt.Println(msg2) // 输出: 加载中...
    
    // 获取语言配置
    locale, ok := i18n.GetLocale(i18n.ChineseSimplified)
    if ok {
        fmt.Println("语言:", locale.Name)
        fmt.Println("货币格式:", locale.Formats.CurrencyFormat)
    }
}
```

### 构建特定语言支持

```bash
# 构建所有语言支持
go build -tags="i18n_all" ./...

# 构建简体中文支持
go build -tags="i18n_zh_cn" ./...

# 构建多种语言支持
go build -tags="i18n_zh_cn,i18n_ja,i18n_fr" ./...

# 默认构建（仅英语）
go build ./...
```

## API 参考

### 核心函数

#### 语言管理

```go
// 设置默认语言
func SetDefaultLocale(locale string)

// 获取默认语言
func GetDefaultLocale() string

// 注册语言配置
func RegisterLocale(language string, locale *Locale)

// 获取语言配置
func GetLocale(language string) (*Locale, bool)

// 获取所有可用语言
func GetAvailableLocales() []string
```

#### 翻译功能

```go
// 翻译消息（指定语言）
func Translate(language, key string, args ...interface{}) string

// 翻译消息（使用默认语言）
func TranslateDefault(key string, args ...interface{}) string
```

#### 实用函数

```go
// 检查语言是否被支持
func IsSupported(language string) bool

// 标准化语言代码
func NormalizeLanguage(language string) string
```

### 类型定义

#### Locale 结构体

```go
type Locale struct {
    Language     string            // 语言代码 (ISO 639-1)
    Region       string            // 地区代码 (ISO 3166-1 alpha-2)
    Name         string            // 语言本地化名称
    EnglishName  string            // 英语名称
    Messages     map[string]string // 消息映射
    Formats      *Formats          // 格式化配置
}
```

#### Formats 结构体

```go
type Formats struct {
    DateFormat        string // 日期格式
    TimeFormat        string // 时间格式
    DateTimeFormat    string // 日期时间格式
    NumberFormat      string // 数字格式
    CurrencyFormat    string // 货币格式
    DecimalSeparator  string // 小数分隔符
    ThousandSeparator string // 千位分隔符
    Units            *Units  // 单位配置
}
```

## 高级用法

### 自定义语言配置

```go
// 创建自定义语言配置
customLocale := &i18n.Locale{
    Language:     "custom",
    Region:       "XX",
    Name:         "Custom Language",
    EnglishName:  "Custom Language",
    Messages: map[string]string{
        "hello":   "Custom Hello",
        "goodbye": "Custom Goodbye",
    },
    Formats: &i18n.Formats{
        DateFormat:        "2006-01-02",
        CurrencyFormat:    "%.2f $",
        DecimalSeparator:  ".",
        ThousandSeparator: ",",
        Units: &i18n.Units{
            ByteUnits: []string{"B", "KB", "MB", "GB"},
        },
    },
}

// 注册自定义语言
i18n.RegisterLocale("custom", customLocale)
```

### 参数化翻译

```go
// 带参数的翻译
msg := i18n.Translate("zh-CN", "welcome_%s", "张三")
// 如果消息模板为 "welcome_%s": "欢迎%s"，则输出: "欢迎张三"

// 多参数翻译
msg2 := i18n.Translate("en", "user_info_%s_%d", "John", 25)
// 如果消息模板为 "user_info_%s_%d": "User: %s, Age: %d"，则输出: "User: John, Age: 25"
```

### 格式化功能

```go
locale, _ := i18n.GetLocale(i18n.ChineseSimplified)

// 使用本地化的单位
fmt.Printf("文件大小: 1024 %s\\n", locale.Formats.Units.ByteUnits[1]) // 输出: 文件大小: 1024 KB

// 使用本地化的时间单位
fmt.Printf("持续时间: 5 %s\\n", locale.Formats.Units.TimeUnits["minutes"]) // 输出: 持续时间: 5 分钟
```

## 集成到其他包

### 为现有包添加国际化支持

1. 在包中导入 i18n：
```go
import "github.com/lazygophers/utils/i18n"
```

2. 使用 i18n 进行消息翻译：
```go
func ErrorMessage(lang string) string {
    return i18n.Translate(lang, "validation_failed")
}
```

3. 为包创建特定的语言文件（可选）：
```go
//go:build i18n_zh_cn || i18n_all

package yourpackage

import "github.com/lazygophers/utils/i18n"

func init() {
    // 添加包特定的翻译
    locale, _ := i18n.GetLocale(i18n.ChineseSimplified)
    locale.Messages["your_package_error"] = "您的包错误"
}
```

## 最佳实践

### 1. 消息键命名规范

```go
// ✅ 推荐：使用模块.功能.类型的命名方式
"validator.email.invalid"
"network.connection.timeout"
"file.read.error"

// ❌ 避免：过于简单的键名
"error"
"msg"
```

### 2. 回退策略

```go
func GetMessage(lang, key string) string {
    // 尝试获取指定语言的消息
    if msg := i18n.Translate(lang, key); msg != key {
        return msg
    }
    
    // 回退到英语
    if msg := i18n.Translate(i18n.English, key); msg != key {
        return msg
    }
    
    // 最终回退到键名
    return key
}
```

### 3. 性能优化

```go
// 缓存语言配置以避免重复查找
var cachedLocale *i18n.Locale

func init() {
    cachedLocale, _ = i18n.GetLocale(i18n.ChineseSimplified)
}

func GetLocalizedMessage(key string) string {
    if msg, exists := cachedLocale.Messages[key]; exists {
        return msg
    }
    return key
}
```

## 测试

运行测试：

```bash
# 测试所有语言
go test -tags="i18n_all" ./...

# 测试特定语言
go test -tags="i18n_zh_cn" ./...

# 基准测试
go test -tags="i18n_all" -bench=. ./...
```

## 许可证

MIT License - 详见 [LICENSE](../LICENSE) 文件。