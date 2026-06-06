---
title: 模組概覽
---

# 模組概覽

LazyGophers Utils 提供 20+ 個專業模組，涵蓋 Go 開發的各個方面。所有模組已按功能分類，方便您快速找到需要的工具。

## 八個主題分組

| 分組 | 什麼時候先看它 | 模組 |
| --- | --- | --- |
| [核心工具](/zh-TW/modules/core/) | 想減少初始化樣板代碼，或處理數據庫字段映射 | [must](/zh-TW/modules/core/must)、[orm](/zh-TW/modules/core/orm) |
| [數據處理](/zh-TW/modules/data/) | 想做類型轉換、集合操作、字符串規範化、JSON 編解碼 | [candy](/zh-TW/modules/data/candy)、[json](/zh-TW/modules/data/json)、[stringx](/zh-TW/modules/data/stringx)、[anyx](/zh-TW/modules/data/anyx) |
| [緩存策略](/zh-TW/modules/cache/) | 想在多種淘汰策略裡選一個更適合實際負載的緩存 | [緩存概覽](/zh-TW/modules/cache/)、[LRU](/zh-TW/modules/cache/lru)、[LFU](/zh-TW/modules/cache/lfu)、[TinyLFU](/zh-TW/modules/cache/tinylfu)、[SLRU](/zh-TW/modules/cache/slru)、[MRU](/zh-TW/modules/cache/mru)、[ALFU](/zh-TW/modules/cache/alfu)、[ARC](/zh-TW/modules/cache/arc)、[LRU-K](/zh-TW/modules/cache/lruk)、[W-TinyLFU](/zh-TW/modules/cache/wtinylfu)、[FBR](/zh-TW/modules/cache/fbr)、[Optimal](/zh-TW/modules/cache/optimal) |
| [時間與調度](/zh-TW/modules/time/) | 想處理農曆、節氣、日曆或固定排班規則 | [xtime](/zh-TW/modules/time/xtime)、[xtime996](/zh-TW/modules/time/xtime996)、[xtime955](/zh-TW/modules/time/xtime955)、[xtime007](/zh-TW/modules/time/xtime007) |
| [系統與配置](/zh-TW/modules/system/) | 想做配置加載、路徑定位、應用初始化與退出清理 | [config](/zh-TW/modules/system/config)、[runtime](/zh-TW/modules/system/runtime)、[osx](/zh-TW/modules/system/osx)、[app](/zh-TW/modules/system/app)、[atexit](/zh-TW/modules/system/atexit) |
| [網絡與安全](/zh-TW/modules/network/) | 想處理網絡輔助、加密、簽名或 URL 規範化 | [network](/zh-TW/modules/network/network)、[cryptox](/zh-TW/modules/network/cryptox)、[pgp](/zh-TW/modules/network/pgp)、[urlx](/zh-TW/modules/network/urlx) |
| [並發與控制流](/zh-TW/modules/concurrency/) | 想組織任務執行、等待條件、熔斷或去重 | [routine](/zh-TW/modules/concurrency/routine)、[wait](/zh-TW/modules/concurrency/wait)、[hystrix](/zh-TW/modules/concurrency/hystrix)、[singledo](/zh-TW/modules/concurrency/singledo)、[event](/zh-TW/modules/concurrency/event) |
| [開發與測試](/zh-TW/modules/dev/) | 想補默認值、製造隨機/假數據或接入採樣工具 | [fake](/zh-TW/modules/dev/fake)、[randx](/zh-TW/modules/dev/randx)、[defaults](/zh-TW/modules/dev/defaults)、[pyroscope](/zh-TW/modules/dev/pyroscope) |

## 獨立模組

| 模組 | 說明 |
| --- | --- |
| [Validator](/zh-TW/validator/) | 數據校驗——169 個內建驗證器，100% 覆蓋 go-playground/validator v10 全部規則 |

## 選擇順序建議

### 先判斷是「基礎設施問題」還是「業務輔助問題」

- 如果你在做啟動、配置、校驗、數據庫字段映射：先看 **核心工具** 與 **系統與配置**。
- 如果你在做集合處理、字符串規整、JSON、隨機數據：先看 **數據處理** 與 **開發與測試**。
- 如果你在做緩存、並發、重試、定時規則：先看 **緩存策略**、**並發與控制流**、**時間與調度**。

### 再判斷是否存在局部規則

以下主題在選型前最好先讀頁面說明，而不是直接猜 API：

- **緩存**：每種策略的淘汰邏輯不同，默認線程安全語義也要單獨看。
- **xtime**：不僅有時間幫助函數，還包含農曆、節氣與排班規則。
- **atexit**：不同平台退出行為並不完全一樣。

## 推薦閱讀路徑

1. 新項目接入：`must` → `config` → `validator` → 對應業務主題模組。
2. 已有項目補能力：從對應分類頁進入，再看單模組頁的適用場景。
3. 需要精確簽名：模組頁讀完後轉到 [pkg.go.dev](https://pkg.go.dev/github.com/lazygophers/utils)。
