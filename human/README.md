# Human Package

一个提供人类友好格式化显示的 Go 包，支持多种数据类型的格式化和多语言国际化。

## 功能特性

- 📏 **大小格式化**: 字节、文件大小等的人类友好显示
- ⚡ **速度格式化**: 网络速度、处理速度等的格式化
- ⏰ **时间格式化**: 持续时间、相对时间的友好显示
- 🔢 **数字格式化**: 大数字的简化显示和千位分隔
- 🌍 **多语言支持**: 支持中文、英文等多种语言
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
- `human_zh`: 中文支持
- `human_en`: 英文支持 (默认)
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
    
    // 数字格式化
    fmt.Println(human.Number(1234567))       // "123.5万"
    fmt.Println(human.Comma(1234567))        // "1,234,567"
}
```

## 构建选项

```bash
# 仅包含英文
go build .

# 包含中文支持
go build -tags human_zh .

# 包含所有语言
go build -tags human_all .

# 包含特定语言组合
go build -tags "human_zh human_ja" .
```