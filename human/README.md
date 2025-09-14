# Human Package

一个提供人类友好格式化显示的 Go 包，支持多种数据类型的格式化和多语言国际化。

## 功能特性

- 📏 **大小格式化**: 字节、文件大小等的人类友好显示
- ⚡ **速度格式化**: 网络速度、处理速度等的格式化
- ⏰ **时间格式化**: 持续时间、相对时间的友好显示
- 🔢 **数字格式化**: 大数字的简化显示和千位分隔
- 🌍 **多语言支持**: 支持中文（简体/繁体）、英文、法语、俄语、阿拉伯语、西班牙语、日文、韩文等多种语言
- 🏷️ **Build Tags**: 通过构建标签选择需要的语言包

## 架构设计

### 核心组件
1. **Formatter Interface**: 统一的格式化接口
2. **Language System**: 多语言支持系统
3. **Unit Converters**: 各种单位转换器
4. **Locale Manager**: 地区设置管理器

### 支持的格式化类型
- `ByteSize`: 字节大小 (B, KB, MB, GB, TB, PB)
- `Speed`: 速度 (B/s, KB/s, MB/s, GB/s)
- `Duration`: 持续时间 (秒, 分钟, 小时, 天)
- `Number`: 数字 (千位分隔, 简化显示)
- `RelativeTime`: 相对时间 (几分钟前, 刚刚)

### Build Tags 语言支持
- `human_en`: 英文支持 (默认)
- `human_zh`: 简体中文支持
- `human_zh_tw`: 繁体中文支持  
- `human_fr`: 法语支持
- `human_ru`: 俄语支持
- `human_ar`: 阿拉伯语支持
- `human_es`: 西班牙语支持
- `human_ja`: 日文支持
- `human_ko`: 韩文支持
- `human_all`: 包含所有语言

## 使用示例

```go
package main

import (
    "fmt"
    "time"
    
    "github.com/lazygophers/utils/human"
)

func main() {
    // 设置语言
    human.SetLocale("zh-CN")
    
    // 字节大小格式化
    fmt.Println(human.ByteSize(1024))        // "1 KB"
    fmt.Println(human.ByteSize(1536))        // "1.5 KB"
    fmt.Println(human.ByteSize(1073741824))  // "1 GB"
    
    // 速度格式化
    fmt.Println(human.Speed(1024))           // "1 KB/s"
    fmt.Println(human.Speed(1048576))        // "1 MB/s"
    
    // 时间格式化
    fmt.Println(human.Duration(time.Hour))           // "1小时"
    fmt.Println(human.Duration(90*time.Minute))      // "1小时30分钟"
    
    // 相对时间
    fmt.Println(human.RelativeTime(time.Now().Add(-5*time.Minute)))  // "5分钟前"
}

// 多语言示例
func multiLanguageExamples() {
    // 繁体中文
    human.SetLocale("zh-TW")
    fmt.Println(human.Duration(90*time.Minute))      // "1小時30分鐘"
    
    // 法语
    human.SetLocale("fr")  
    fmt.Println(human.ByteSize(1024))                // "1 Ko"
    fmt.Println(human.RelativeTime(time.Now().Add(-5*time.Minute)))  // "il y a 5 minutes"
    
    // 俄语
    human.SetLocale("ru")
    fmt.Println(human.ByteSize(1024))                // "1 КБ"
    fmt.Println(human.BitSpeed(1000))                // "1 Кбит/с"
    
    // 阿拉伯语
    human.SetLocale("ar")
    fmt.Println(human.ByteSize(1024))                // "1 كب"
    
    // 西班牙语
    human.SetLocale("es")
    fmt.Println(human.Duration(90*time.Minute))      // "1 hora 30 minutos"
}
```

## 构建选项

```bash
# 仅包含英文
go build .

# 包含简体中文支持
go build -tags human_zh .

# 包含繁体中文支持  
go build -tags human_zh_tw .

# 包含法语支持
go build -tags human_fr .

# 包含所有语言
go build -tags human_all .

# 包含特定语言组合
go build -tags "human_zh human_fr human_es" .
```

## 支持的语言

| 语言代码 | 语言名称 | Build Tag | 特色功能 |
|---------|---------|-----------|----------|
| `en` | English | (默认包含) | 单复数变化, 标准单位 |
| `zh` | 简体中文 | `human_zh` | 中文数字单位 (万、亿) |
| `zh-CN` | 简体中文 | `human_zh` | 与 `zh` 相同 |
| `zh-TW` | 繁体中文 | `human_zh_tw` | 繁体字符, 台湾地区格式 |
| `fr` | Français | `human_fr` | 法语单位 (Ko, Mo), 逗号小数点 |
| `ru` | Русский | `human_ru` | 西里尔字母单位, 俄语复数规则 |
| `ar` | العربية | `human_ar` | 阿拉伯语单位, 从右到左显示 |
| `es` | Español | `human_es` | 西班牙语复数, 点分千位分隔 |
| `ja` | 日本語 | `human_ja` | 日文汉字时间单位 |
| `ko` | 한국어 | `human_ko` | 韩文时间单位 |

## 高级功能

### 选项配置
```go
// 使用功能选项模式
result := human.ByteSize(1536, 
    human.WithLocale("fr"),      // 设置语言
    human.WithPrecision(2),      // 精度
    human.WithCompact(),         // 紧凑模式
)

// 时钟格式
result := human.Duration(90*time.Second, 
    human.WithLocale("es"),
    human.WithClockFormat(),     // 时钟格式 "1:30"
)
```

### 可用选项
```go
// 选项函数
human.WithLocale(locale string)      // 设置语言地区
human.WithPrecision(precision int)   // 设置小数精度  
human.WithCompact()                  // 启用紧凑模式
human.WithClockFormat()              // 启用时钟格式
```