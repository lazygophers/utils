---
title: 效能與最佳實踐 - Validator
---

# 效能與最佳實踐

## 效能

Validator 內建多種演算法變體，透過基準測試選擇最佳實現：

| 驗證項 | 最佳化方向 |
|--------|-----------|
| 訊息格式化 | 5 種字串模板實現（原始/Builder/編譯/位元組切片/無 fmt） |
| 郵箱 | 多種正規表示式/非正規表示式方案對比 |

執行基準測試：

```bash
go test -bench=. ./validator/
```

## 最佳實踐

**標籤設計**：每個欄位用 2-3 個標籤組合，而非一個萬能標籤。

```go
// ✅ 推薦
Email string `validate:"required,email"`
Name  string `validate:"required,min=2,max=50"`

// ❌ 太粗
Email string `validate:"required"`
Name  string `validate:"required"`
```

**驗證時機**：在資料進入系統的第一站驗證（配置載入、請求解析），不要等到業務深處。

**錯誤處理**：用 `ValidationErrors` 型別斷言，逐欄位生成使用者友好的訊息。

**複用實例**：用 `Default()` 或快取 `New()` 實例，避免每次建立。
