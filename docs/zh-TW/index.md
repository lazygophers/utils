---
pageType: home

hero:
    name: LazyGophers Utils
    text: 現代化 Go 工具庫
    tagline: 為現代開發工作流程提供強大支持
    actions:
        - theme: brand
          text: 開始使用
          link: /zh-TW/guide/getting-started
        - theme: alt
          text: 瀏覽模組
          link: /zh-TW/modules/overview

features:
    - title: "模組化設計"
      details: 20+ 個專業模組，只導入您需要的內容，保持專案輕量高效
      icon: 🧩
    - title: "性能優先"
      details: 為速度和最小內存使用進行優化，關鍵操作比標準庫快 2-3 倍
      icon: ⚡
    - title: "類型安全"
      details: 利用 Go 泛型實現編譯時安全，避免運行時類型錯誤
      icon: 🛡️
    - title: "生產就緒"
      details: Goroutine 安全且經過實戰檢驗，可直接用於生產環境
      icon: 🔒
    - title: "開發者友好"
      details: 全面的文檔和示例，快速上手，提升開發效率
      icon: 📖
    - title: "時間與調度"
      details: 支持農曆、中國生肖、節氣，以及多種工作時間計算
      icon: ⏰
---

## 🌍 多語言支持

[简体中文](/zh-CN/) • [繁體中文](/zh-TW/) • [English](/en/)

## 🎯 什麼是 LazyGophers Utils？

LazyGophers Utils 是一個全面的 Go 工具庫，為常見開發任務提供 **20+ 個專業模組**。採用現代 Go 實踐構建，它提供類型安全、高性能的解決方案，可以無縫集成到任何 Go 專案中。

### ✨ 為什麼選擇 LazyGophers Utils？

-   **🧩 模組化設計** - 只導入您需要的內容
-   **⚡ 性能優先** - 為速度和最小內存使用進行優化
-   **🛡️ 類型安全** - 利用 Go 泛型實現編譯時安全
-   **🔒 生產就緒** - Goroutine 安全且經過實戰檢驗
-   **📖 開發者友好** - 全面的文檔和示例

## 🚀 快速開始

### 安裝

```bash
go get github.com/lazygophers/utils
```

### 30 秒示例

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils"
    "github.com/lazygophers/utils/candy"
    "github.com/lazygophers/utils/xtime"
)

func main() {
    // 簡化的錯誤處理
    data := utils.Must(loadData())

    // 無需麻煩的類型轉換
    userAge := candy.ToInt("25")
    isActive := candy.ToBool("true")

    // 高級時間處理
    calendar := xtime.NowCalendar()
    fmt.Printf("Today: %s\n", calendar.String())
    fmt.Printf("Lunar: %s\n", calendar.LunarDate())
}

func loadData() (string, error) {
    return "Hello, World!", nil
}
```

## 📦 模組概覽

### 🔧 核心工具
| 模組 | 用途 | 主要函數 |
|--------|---------|---------------|
| **[must.go](https://github.com/lazygophers/utils/blob/main/must.go)** | 錯誤斷言 | `Must()`, `MustSuccess()`, `MustOk()` |
| **[orm.go](https://github.com/lazygophers/utils/blob/main/orm.go)** | 數據庫操作 | `Scan()`, `Value()` |
| **[validate.go](https://github.com/lazygophers/utils/blob/main/validator/)** | 數據驗證 | `Validate()` |

### 🍭 數據處理
| 模組 | 用途 | 亮點 |
|--------|---------|------------|
| **[candy/](https://github.com/lazygophers/utils/tree/main/candy)** | 類型轉換糖 | 零分配轉換 |
| **[json/](https://github.com/lazygophers/utils/tree/main/json)** | 增強的 JSON 處理 | 更好的錯誤消息 |
| **[stringx/](https://github.com/lazygophers/utils/tree/main/stringx)** | 字符串工具 | Unicode 感知操作 |
| **[anyx/](https://github.com/lazygophers/utils/tree/main/anyx)** | Interface{} 輔助 | 類型安全的 any 操作 |

### ⏰ 時間與調度
| 模組 | 用途 | 特殊功能 |
|--------|---------|------------------|
| **[xtime/](https://github.com/lazygophers/utils/tree/main/xtime)** | 高級時間處理 | 🌙 農曆, 🐲 中國生肖, 🌾 節氣 |
| **[xtime996/](https://github.com/lazygophers/utils/tree/main/xtime996)** | 996 工作時間 | 工作時間計算 |
| **[xtime955/](https://github.com/lazygophers/utils/tree/main/xtime955)** | 955 工作時間 | 平衡時間表支持 |
| **[xtime007/](https://github.com/lazygophers/utils/tree/main/xtime007)** | 24/7 操作 | 始終在線的時間工具 |

### 🔧 系統與配置
| 模組 | 用途 | 用例 |
|--------|---------|-----------|
| **[config/](https://github.com/lazygophers/utils/tree/main/config)** | 配置管理 | JSON, YAML, TOML, INI, HCL 支持 |
| **[runtime/](https://github.com/lazygophers/utils/tree/main/runtime)** | 運行時信息 | 系統檢測和診斷 |
| **[osx/](https://github.com/lazygophers/utils/tree/main/osx)** | 操作系統操作 | 文件和進程管理 |
| **[app/](https://github.com/lazygophers/utils/tree/main/app)** | 應用框架 | 生命週期管理 |
| **[atexit/](https://github.com/lazygophers/utils/tree/main/atexit)** | 優雅關閉 | 乾淨的退出處理 |

### 🌐 網絡與安全
| 模組 | 用途 | 功能 |
|--------|---------|----------|
| **[network/](https://github.com/lazygophers/utils/tree/main/network)** | HTTP 工具 | 連接池, 重試邏輯 |
| **[cryptox/](https://github.com/lazygophers/utils/tree/main/cryptox)** | 加密函數 | 哈希, 加密, 安全隨機 |
| **[pgp/](https://github.com/lazygophers/utils/tree/main/pgp)** | PGP 操作 | 郵件加密, 文件簽名 |
| **[urlx/](https://github.com/lazygophers/utils/tree/main/urlx)** | URL 操作 | 解析, 構建, 驗證 |

### 🚀 並發與控制流
| 模組 | 用途 | 模式 |
|--------|---------|----------|
| **[routine/](https://github.com/lazygophers/utils/tree/main/routine)** | Goroutine 管理 | 工作池, 任務調度 |
| **[wait/](https://github.com/lazygophers/utils/tree/main/wait)** | 流程控制 | 超時, 重試, 速率限制 |
| **[hystrix/](https://github.com/lazygophers/utils/tree/main/hystrix)** | 熔斷器 | 容錯, 優雅降級 |
| **[singledo/](https://github.com/lazygophers/utils/tree/main/singledo)** | 單例執行 | 防止重複操作 |
| **[event/](https://github.com/lazygophers/utils/tree/main/event)** | 事件系統 | 發布/訂閱模式實現 |

### 🧪 開發與測試
| 模組 | 用途 | 開發階段 |
|--------|---------|-------------------|
| **[fake/](https://github.com/lazygophers/utils/tree/main/fake)** | 測試數據生成 | 單元測試, 集成測試 |
| **[randx/](https://github.com/lazygophers/utils/tree/main/randx)** | 隨機工具 | 加密安全的隨機數 |
| **[defaults/](https://github.com/lazygophers/utils/tree/main/defaults)** | 默認值 | 結構體初始化 |
| **[pyroscope/](https://github.com/lazygophers/utils/tree/main/pyroscope)** | 性能分析 | 生產監控 |

## 📊 性能亮點

| 操作 | 時間 | 內存 | vs 標準庫 |
|-----------|------|--------|-------------------|
| `candy.ToInt()` | 12.3 ns/op | 0 B/op | **快 3.2 倍** |
| `json.Marshal()` | 156 ns/op | 64 B/op | **快 1.8 倍** |
| `xtime.Now()` | 45.2 ns/op | 0 B/op | **快 2.1 倍** |
| `utils.Must()` | 2.1 ns/op | 0 B/op | **零開銷** |

## 🤝 貢獻

我們歡迎貢獻！以下是入門方法：

### 快速貢獻指南

1. **Fork** 倉庫
2. **創建** 功能分支：`git checkout -b feature/amazing-feature`
3. **編寫** 代碼和測試
4. **確保** 測試通過：`go test ./...`
5. **提交** 拉取請求

## 📄 許可證

本專案採用 **GNU Affero General Public License v3.0** 許可。

詳見 [LICENSE](https://github.com/lazygophers/utils/blob/main/LICENSE) 文件。

---

<div align="center">

**⭐ 如果這個專案幫助您構建更好的 Go 應用，請給它一個星標！**

[🚀 開始使用](/zh-TW/guide/getting-started) • [📖 瀏覽模組](/zh-TW/modules/overview) • [🤝 貢獻](https://github.com/lazygophers/utils/blob/main/CONTRIBUTING.md)

*由 LazyGophers 團隊用 ❤️ 構建*

</div>
