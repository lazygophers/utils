---
title: human
description: 人性化的數據格式化工具
---

# human

`human` 包提供將技術數據轉換為人類可讀格式的功能。

## 適用場景

- **監控展示**：將位元組數、速度、時長轉換為直觀的"KB/s"、"2分鐘前"等格式
- **日誌輸出**：提供易讀的時間、大小、速度信息
- **多語言應用**：需要根據地區顯示不同格式

## 主要功能

### 大小與速度格式化

```go
import "github.com/lazygophers/utils/human"

// 位元組大小格式化
human.ByteSize(1536)        // "1.5 KB"
human.ByteSize(1048576)     // "1.0 MB"

// 速度格式化（位元組/秒）
human.Speed(1048576)         // "1.0 MB/s"

// 比特速度格式化（比特/秒）
human.BitSpeed(1000000)      // "1.0 Mbps"
```

### 時間格式化

```go
import "time"

// 時長格式化
human.Duration(time.Hour)    // "1小時"
human.Duration(90*time.Minute) // "1小時30分鐘"

// 相對時間
human.RelativeTime(time.Now().Add(-time.Hour)) // "1小時前"
```

### 多語言支援

```go
// 設置預設語言
human.SetLocale("zh")  // 中文
human.SetLocale("en")  // 英文

// 獲取當前語言
lang := human.GetLocale()
```

### 精度控制

```go
// 設置預設精度（小數位數）
human.SetDefaultPrecision(2)

human.ByteSize(1536)  // "1.50 KB"（2位小數）
```

## 使用建議

1. **監控面板**：使用 `Speed` 和 `ByteSize` 替代原始數值
2. **使用者介面**：使用 `RelativeTime` 顯示"多久前"
3. **國際化**：配合 `SetLocale` 實現多語言切換

## 注意事項

- 預設語言為英文（`en`），中文應用需調用 `SetLocale("zh")`
- 精度設置是全局的，會影響所有後續調用
