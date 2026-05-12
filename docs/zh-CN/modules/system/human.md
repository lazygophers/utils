---
title: human
description: 人性化的数据格式化工具
---

# human

`human` 包提供将技术数据转换为人类可读格式的功能。

## 适用场景

- **监控展示**：将字节数、速度、时长转换为直观的"KB/s"、"2分钟前"等格式
- **日志输出**：提供易读的时间、大小、速度信息
- **多语言应用**：需要根据地区显示不同格式

## 主要功能

### 大小与速度格式化

```go
import "github.com/lazygophers/utils/human"

// 字节大小格式化
human.ByteSize(1536)        // "1.5 KB"
human.ByteSize(1048576)     // "1.0 MB"

// 速度格式化（字节/秒）
human.Speed(1048576)         // "1.0 MB/s"

// 比特速度格式化（比特/秒）
human.BitSpeed(1000000)      // "1.0 Mbps"
```

### 时间格式化

```go
import "time"

// 时长格式化
human.Duration(time.Hour)    // "1小时"
human.Duration(90*time.Minute) // "1小时30分钟"

// 相对时间
human.RelativeTime(time.Now().Add(-time.Hour)) // "1小时前"
```

### 多语言支持

```go
// 设置默认语言
human.SetLocale("zh")  // 中文
human.SetLocale("en")  // 英文

// 获取当前语言
lang := human.GetLocale()
```

### 精度控制

```go
// 设置默认精度（小数位数）
human.SetDefaultPrecision(2)

human.ByteSize(1536)  // "1.50 KB"（2位小数）
```

## 使用建议

1. **监控面板**：使用 `Speed` 和 `ByteSize` 替代原始数值
2. **用户界面**：使用 `RelativeTime` 显示"多久前"
3. **国际化**：配合 `SetLocale` 实现多语言切换

## 注意事项

- 默认语言为英文（`en`），中文应用需调用 `SetLocale("zh")`
- 精度设置是全局的，会影响所有后续调用
