# 貢獻指南

歡迎參與 LazyGophers Utils 專案的貢獻！我們非常感謝社群的每一份貢獻。

[![Contributors](https://img.shields.io/badge/Contributors-Welcome-brightgreen.svg)](#如何貢獻)
[![Code Style](https://img.shields.io/badge/Code%20Style-Go%20Standard-blue.svg)](#程式碼規範)

## 🤝 如何貢獻

### 貢獻類型

我們歡迎以下類型的貢獻：

- 🐛 **錯誤修復** - 修復已知問題
- ✨ **新功能** - 新增實用函數或模組
- 📚 **文件改進** - 完善文件、新增範例
- 🎨 **程式碼優化** - 效能優化、重構
- 🧪 **測試改進** - 增加測試覆蓋率、修復測試問題
- 🌐 **國際化** - 新增多語言支援

### 貢獻流程

#### 1. 準備工作

**Fork 專案**
```bash
# 1. Fork 此專案到您的 GitHub 帳號
# 2. 在本地複製您的 fork
git clone https://github.com/YOUR_USERNAME/utils.git
cd utils

# 3. 新增原專案為上游倉庫
git remote add upstream https://github.com/lazygophers/utils.git

# 4. 建立新的功能分支
git checkout -b feature/your-awesome-feature
```

**設定開發環境**
```bash
# 安裝依賴項
go mod tidy

# 驗證環境
go version  # 需要 Go 1.24.0+
go test ./... # 確保所有測試通過
```

#### 2. 開發階段

**編寫程式碼**
- 遵循 [程式碼規範](#程式碼規範)
- 為新功能編寫測試案例
- 確保測試覆蓋率不低於當前水準
- 新增必要的文件註解

**提交規範**
```bash
# 使用標準化的提交訊息格式
git commit -m "feat(module): 新增新的實用函數

- 新增 FormatDuration 函數
- 支援多種時間格式輸出
- 新增全面的測試案例
- 更新相關文件

Closes #123"
```

**提交訊息格式**：
```
<type>(<scope>): <subject>

<body>

<footer>
```

**類型分類**：
- `feat`: 新功能
- `fix`: 錯誤修復  
- `docs`: 文件更新
- `style`: 程式碼格式調整
- `refactor`: 程式碼重構
- `perf`: 效能優化
- `test`: 測試相關
- `chore`: 建置工具或依賴項更新

**範圍範圍** (可選)：
- `candy`: candy 模組
- `xtime`: xtime 模組
- `config`: config 模組
- `cryptox`: cryptox 模組
- 等等...

#### 3. 測試和驗證

**執行測試**
```bash
# 執行所有測試
go test -v ./...

# 檢查測試覆蓋率
go test -cover -v ./...

# 執行基準測試
go test -bench=. ./...

# 檢查程式碼格式
go fmt ./...

# 靜態分析
go vet ./...
```

**效能測試**
```bash
# 執行效能測試
go test -bench=BenchmarkYourFunction -benchmem ./...

# 確保沒有顯著的效能回歸
```

#### 4. 建立 Pull Request

**推送到您的 Fork**
```bash
git push origin feature/your-awesome-feature
```

**建立 PR**
1. 訪問 GitHub 上的專案頁面
2. 點選 "New Pull Request"
3. 選擇您的分支
4. 填寫 PR 描述 (參考 [PR 模板](#pr-模板))
5. 確保所有檢查通過

#### 5. 程式碼審查

- 維護者將審查您的程式碼
- 根據反饋進行修改
- 保持溝通和合作態度
- 測試通過後將被合併

## 📝 程式碼規範

### Go 程式碼風格

**基本規範**
```go
// ✅ 良好範例
package candy

import (
    "context"
    "fmt"
    "time"
    
    "github.com/lazygophers/log"
)

// FormatDuration 將時間間隔格式化為人類可讀的字符串
// 支援多種精度級別，自動選擇合適的單位
//
// 參數：
//   - duration: 要格式化的時間間隔
//   - precision: 精度級別 (1-3)
//
// 回傳：
//   - string: 格式化後的字符串，如 "2 小時 30 分鐘"
//
// 範例：
//   FormatDuration(90*time.Minute, 2) // 回傳 "1 小時 30 分鐘"
//   FormatDuration(45*time.Second, 1) // 回傳 "45 秒"
func FormatDuration(duration time.Duration, precision int) string {
    if duration == 0 {
        return "0 秒"
    }
    
    // 實現邏輯...
    return result
}
```

**命名慣例**
- 使用 CamelCase
- 函數名稱以動詞開頭：`Get`, `Set`, `Format`, `Parse`
- 常數使用 ALL_CAPS：`const MaxRetries = 3`
- 私有成員使用小寫：`internalHelper`
- 套件名稱使用小寫單字：`candy`, `xtime`

**註解規範**
- 所有公開函數必須有註解
- 註解以函數名稱開頭
- 包含參數和回傳值說明  
- 提供使用範例
- 英文註解，簡潔明瞭

**錯誤處理**
```go
// ✅ 推薦的錯誤處理方式
func ProcessData(data []byte) (*Result, error) {
    if len(data) == 0 {
        log.Warn("提供了空資料")
        return nil, fmt.Errorf("資料不能為空")
    }
    
    result, err := parseData(data)
    if err != nil {
        log.Error("解析資料失敗", log.Error(err))
        return nil, fmt.Errorf("解析資料失敗: %w", err)
    }
    
    return result, nil
}
```

### 專案結構規範

**模組組織**
```
utils/
├── README.md           # 專案概述
├── CONTRIBUTING.md     # 貢獻指南  
├── SECURITY.md        # 安全政策
├── go.mod             # Go 模組定義
├── must.go            # 核心實用函數
├── candy/             # 資料處理工具
│   ├── README.md      # 模組文件
│   ├── to_string.go   # 型別轉換
│   └── to_string_test.go
├── xtime/             # 時間處理工具  
│   ├── README.md      # 詳細使用文件
│   ├── TESTING.md     # 測試報告
│   ├── PERFORMANCE.md # 效能報告
│   ├── calendar.go    # 日曆功能
│   └── calendar_test.go
└── ...
```

**檔案命名**
- 使用小寫字母和底線：`to_string.go`
- 測試檔案後綴：`_test.go`
- 基準測試：`_benchmark_test.go`
- 文件檔案：`README.md`, `TESTING.md`

### 測試規範

**測試覆蓋率要求**
- 新功能測試覆蓋率必須 ≥ 90%
- 不能降低整體測試覆蓋率
- 包含正常情況和邊界情況
- 錯誤處理路徑必須測試

**測試範例**
```go
func TestFormatDuration(t *testing.T) {
    testCases := []struct {
        name      string
        duration  time.Duration
        precision int
        want      string
    }{
        {
            name:      "零時間",
            duration:  0,
            precision: 1,
            want:      "0 秒",
        },
        {
            name:      "90 分鐘高精度",
            duration:  90 * time.Minute,
            precision: 2,
            want:      "1 小時 30 分鐘",
        },
        // 更多測試案例...
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            got := FormatDuration(tc.duration, tc.precision)
            assert.Equal(t, tc.want, got)
        })
    }
}

// 基準測試
func BenchmarkFormatDuration(b *testing.B) {
    duration := 90 * time.Minute
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = FormatDuration(duration, 2)
    }
}
```

## 🎯 重點開發領域

### 高優先順序

1. **xtime 模組增強**
   - 農曆和節氣功能改進
   - 效能優化
   - 更多文化特定功能

2. **candy 模組擴展**  
   - 型別轉換函數
   - 資料處理工具
   - 效能優化

3. **測試覆蓋率改進**
   - 目標：所有模組 > 90%
   - 邊界情況補充
   - 效能測試改進

### 中等優先順序

4. **新實用模組**
   - AI/ML 實用函數
   - 雲服務整合工具
   - 微服務工具

5. **文件增強**
   - API 參考文件
   - 最佳實務指南
   - 效能優化指南

### 歡迎貢獻

- 🌏 **多語言支援** - 英文文件、錯誤訊息國際化
- 📊 **更多資料格式支援** - XML、YAML、TOML 處理
- 🔧 **開發工具** - 程式碼生成、配置管理
- 🎨 **UI/UX 工具** - 顏色處理、格式化輸出
- 🔐 **安全工具** - 加密/解密、簽名驗證

## 📋 PR 模板

建立 PR 時請使用以下模板：

```markdown
## 變更描述

簡要描述此變更的內容和目的。

## 變更類型

- [ ] 錯誤修復
- [ ] 新功能
- [ ] 文件更新
- [ ] 效能優化  
- [ ] 程式碼重構
- [ ] 測試改進

## 詳細變更

### 新功能
- 新增 `FormatDuration` 函數
- 支援多種精度級別
- 新增中文時間單位顯示

### 修復問題  
- 修復時區轉換錯誤 (#123)
- 解決記憶體洩漏問題

### 效能優化
- 優化字符串連接效能
- 減少 30% 記憶體分配

## 測試描述

- [ ] 所有測試通過
- [ ] 新增新測試案例
- [ ] 測試覆蓋率 ≥ 90%
- [ ] 基準測試通過

**測試覆蓋率**: 92.5%

## 文件更新

- [ ] 更新 README.md
- [ ] 新增函數註解
- [ ] 更新範例程式碼

## 相容性

- [ ] 向後相容
- [ ] 需要版本升級 (說明原因)
- [ ] 破壞性變更 (詳細說明)

## 檢查清單

- [ ] 程式碼遵循專案規範
- [ ] 通過 `go fmt` 格式檢查
- [ ] 通過 `go vet` 靜態檢查
- [ ] 所有測試通過
- [ ] 文件已更新
- [ ] 提交訊息遵循規範

## 相關問題

Closes #123
Refs #456

## 螢幕截圖/演示

如有必要，請提供螢幕截圖或演示。
```

## 🐛 錯誤報告

發現錯誤？請使用以下模板建立 Issue：

```markdown
## 錯誤描述

簡要描述遇到的問題。

## 重現步驟

1. 執行步驟 1
2. 執行步驟 2  
3. 觀察結果

## 預期行為

描述您期望看到的正確行為。

## 實際行為

描述實際觀察到的錯誤行為。

## 環境資訊

- **作業系統**: macOS 12.0
- **Go 版本**: 1.24.0
- **Utils 版本**: v1.2.0
- **其他相關資訊**:

## 錯誤日誌

```
在此粘貼錯誤日誌
```

## 最小重現範例

```go
package main

import (
    "github.com/lazygophers/utils/xtime"
)

func main() {
    // 最小錯誤重現程式碼
}
```
```

## ✨ 功能請求

想要新功能？請使用以下模板：

```markdown
## 功能描述

描述您想要新增的功能。

## 使用案例

描述何時會使用此功能。

## 建議的 API 設計

```go
// 建議的函數簽名和使用方式
func NewAwesomeFunction(param string) (Result, error) {
    // ...
}
```

## 替代解決方案

您是否考慮過其他解決方案？

## 其他資訊

其他相關資訊或參考。
```

## 🏆 貢獻者認可

### 貢獻類型認可

我們將根據貢獻類型給予不同認可：

- 🥇 **核心貢獻者** - 長期活躍，重要功能貢獻
- 🥈 **活躍貢獻者** - 多次有價值的貢獻  
- 🥉 **社群貢獻者** - 錯誤修復、文件改進
- 🌟 **首次貢獻者** - 歡迎首次貢獻

### 貢獻統計

我們將在以下地方展示貢獻者：

- README.md 貢獻者列表
- 發布說明中的致謝
- 專案網站 (如有)
- 年度貢獻者報告

## 💬 溝通

### 獲得幫助

- 📖 **文件問題**: 檢查各模組的 README.md
- 🐛 **錯誤報告**: [GitHub Issues](https://github.com/lazygophers/utils/issues)
- 💡 **功能討論**: [GitHub Discussions](https://github.com/lazygophers/utils/discussions)
- ❓ **使用問題**: [GitHub Discussions Q&A](https://github.com/lazygophers/utils/discussions/categories/q-a)

### 討論規範

請遵循以下溝通規範：

- 使用友好和專業的語言
- 提供詳細的問題描述和建議
- 提供充分的上下文資訊
- 尊重不同的觀點和意見
- 積極參與建設性討論

## 📜 授權

此專案採用 [GNU Affero General Public License v3.0](LICENSE) 授權。

**貢獻即表示同意**：
- 您擁有提交程式碼的版權
- 同意在 AGPL v3.0 授權下發布程式碼
- 遵循專案的貢獻者行為準則

## 🙏 致謝

感謝所有為 LazyGophers Utils 專案做出貢獻的開發者！

**特別感謝**：
- 所有提交 Issues 和 PRs 的貢獻者
- 提供建議和反饋的社群成員
- 幫助改進文件的志願者

---

**其他語言版本：** [English](CONTRIBUTING.md) | [简体中文](CONTRIBUTING_zh.md) | [Français](CONTRIBUTING_fr.md) | [Русский](CONTRIBUTING_ru.md) | [Español](CONTRIBUTING_es.md) | [العربية](CONTRIBUTING_ar.md)

**Happy Coding! 🎉**

如有問題，隨時聯絡維護者團隊。我們樂意幫助您開始貢獻之旅！