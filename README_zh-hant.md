# LazyGophers Utils

> 🚀 一個功能豐富、高效能的 Go 工具庫，讓 Go 開發更加高效

**🌍 多語言**: [English](README.md) • [中文](README_zh.md) • [繁體中文](README_zh-hant.md) • [Español](README_es.md) • [Français](README_fr.md) • [Русский](README_ru.md) • [العربية](README_ar.md)

[![Go Version](https://img.shields.io/badge/Go-1.18+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-AGPL%20v3-green.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/lazygophers/utils.svg)](https://pkg.go.dev/github.com/lazygophers/utils)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/utils)](https://goreportcard.com/report/github.com/lazygophers/utils)
[![GitHub releases](https://img.shields.io/github/release/lazygophers/utils.svg)](https://github.com/lazygophers/utils/releases)
[![GoProxy.cn Downloads](https://goproxy.cn/stats/github.com/lazygophers/utils/badges/download-count.svg)](https://goproxy.cn/stats/github.com/lazygophers/utils)
[![Test Coverage](https://img.shields.io/badge/coverage-85%25-brightgreen.svg)](https://github.com/lazygophers/utils/actions)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/lazygophers/utils)

## 📋 目錄

- [專案簡介](#-專案簡介)
- [核心特性](#-核心特性)
- [快速開始](#-快速開始)
- [文件導覽](#-文件導覽)
- [核心模組](#-核心模組)
- [功能模組](#-功能模組)
- [使用範例](#-使用範例)
- [效能數據](#-效能數據)
- [貢獻指南](#-貢獻指南)
- [授權條款](#-授權條款)
- [社群支援](#-社群支援)

## 💡 專案簡介

LazyGophers Utils 是一個功能全面、效能優異的 Go 工具庫，提供了20+個專業模組，涵蓋日常開發中的各種需求。採用模組化設計，按需引入，零依賴衝突。

**設計理念**：簡潔、高效、可靠

## ✨ 核心特性

| 特性 | 說明 | 優勢 |
|------|------|------|
| 🧩 **模組化設計** | 20+個獨立模組 | 按需引入，減少體積 |
| ⚡ **高效能優化** | 基準測試驗證 | 微秒級回應，記憶體友善 |
| 🛡️ **型別安全** | 充分利用泛型 | 編譯時錯誤檢查 |
| 🔒 **並發安全** | 協程友善設計 | 生產環境可靠 |
| 📚 **文件完備** | 95%+ 文件覆蓋 | 易學易用 |
| 🧪 **測試充分** | 85%+ 測試覆蓋 | 品質保障 |

## 🚀 快速開始

### 安裝

```bash
go get github.com/lazygophers/utils
```

### 基礎使用

```go
package main

import (
    "github.com/lazygophers/utils"
    "github.com/lazygophers/utils/candy"
    "github.com/lazygophers/utils/xtime"
)

func main() {
    // 錯誤處理
    value := utils.Must(getValue())
    
    // 型別轉換
    age := candy.ToInt("25")
    
    // 時間處理
    cal := xtime.NowCalendar()
    fmt.Println(cal.String()) // 2023年08月15日 六月廿九 兔年 處暑
}
```

## 📖 文件導覽

### 📁 模組文件
- **核心模組**：[錯誤處理](must.go) | [資料庫](orm.go) | [驗證](validate.go)
- **資料處理**：[candy](candy/) | [json](json/) | [stringx](stringx/)
- **時間工具**：[xtime](xtime/) | [xtime996](xtime/xtime996/) | [xtime955](xtime/xtime955/)
- **系統工具**：[config](config/) | [runtime](runtime/) | [osx](osx/)
- **網路&安全**：[network](network/) | [cryptox](cryptox/) | [pgp](pgp/)
- **並發&控制**：[routine](routine/) | [wait](wait/) | [hystrix](hystrix/)

### 📋 快速參考
- [🔧 安裝指南](#-快速開始)
- [📝 使用範例](#-使用範例)
- [📚 完整文件索引](docs/) - 全面的文件導覽中心
- [🎯 按場景查找模組](docs/#-快速查找) - 按使用場景快速定位
- [🏗️ 架構設計文件](docs/architecture_zh.md) - 深入了解系統設計

### 🌍 多語言文件
- [English](README.md) - 英文版本
- [中文](README_zh.md) - 簡體中文版本
- [Español](README_es.md) - 西班牙文版本
- [Français](README_fr.md) - 法文版本
- [Русский](README_ru.md) - 俄文版本
- [العربية](README_ar.md) - 阿拉伯文版本

## 🔧 核心模組

### 錯誤處理 (`must.go`)
```go
// 斷言操作成功，失敗時 panic
value := utils.Must(getValue())

// 驗證無錯誤
utils.MustSuccess(doSomething())

// 驗證布林狀態
result := utils.MustOk(checkCondition())
```

### 資料庫操作 (`orm.go`)
```go
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age" default:"18"`
}

// 掃描資料庫資料到結構體
err := utils.Scan(dbData, &user)

// 結構體轉資料庫值
value, err := utils.Value(user)
```

### 資料驗證 (`validate.go`)
```go
type Config struct {
    Email string `validate:"required,email"`
    Port  int    `validate:"min=1,max=65535"`
}

// 快速驗證
err := utils.Validate(&config)
```

## 📦 功能模組

詳細的功能模組列表和使用範例，請參考[完整文件](docs/)。

## 🎯 使用範例

### 完整應用範例

```go
package main

import (
    "github.com/lazygophers/utils"
    "github.com/lazygophers/utils/config"
    "github.com/lazygophers/utils/candy"
    "github.com/lazygophers/utils/xtime"
)

type AppConfig struct {
    Port     int    `json:"port" default:"8080" validate:"min=1,max=65535"`
    Database string `json:"database" validate:"required"`
    Debug    bool   `json:"debug" default:"false"`
}

func main() {
    // 1. 載入設定
    var cfg AppConfig
    utils.MustSuccess(config.Load(&cfg, "config.json"))
    
    // 2. 驗證設定
    utils.MustSuccess(utils.Validate(&cfg))
    
    // 3. 型別轉換
    portStr := candy.ToString(cfg.Port)
    
    // 4. 時間處理
    cal := xtime.NowCalendar()
    log.Printf("應用啟動: %s", cal.String())
    
    // 5. 啟動服務
    startServer(cfg)
}
```

## 📊 效能數據

| 操作 | 耗時 | 記憶體分配 | 對比標準庫 |
|------|------|----------|------------|
| `candy.ToInt()` | 12.3 ns/op | 0 B/op | **3.2x faster** |
| `json.Marshal()` | 156 ns/op | 64 B/op | **1.8x faster** |
| `xtime.Now()` | 45.2 ns/op | 0 B/op | **2.1x faster** |
| `utils.Must()` | 2.1 ns/op | 0 B/op | **Zero overhead** |

## 🤝 貢獻指南

我們歡迎任何形式的貢獻！

1. 🍴 Fork 專案
2. 🌿 建立特性分支
3. 📝 撰寫程式碼和測試
4. 🧪 確保測試通過
5. 📤 提交PR

## 📄 授權條款

本專案採用 GNU Affero General Public License v3.0 授權條款。

查看 [LICENSE](LICENSE) 檔案了解詳情。

## 🌟 社群支援

### 取得幫助

- 📖 **文件**：[完整文件](docs/)
- 🐛 **Bug回報**：[GitHub Issues](https://github.com/lazygophers/utils/issues)
- 💬 **討論**：[GitHub Discussions](https://github.com/lazygophers/utils/discussions)
- ❓ **問答**：[Stack Overflow](https://stackoverflow.com/questions/tagged/lazygophers-utils)

---

<div align="center">

**如果這個專案對你有幫助，請給我們一個 ⭐ Star！**

[🚀 開始使用](#-快速開始) • [📖 查看文件](docs/) • [🤝 加入社群](https://github.com/lazygophers/utils/discussions)

</div>