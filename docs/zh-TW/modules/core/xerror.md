---
title: xerror
---

# xerror

帶錯誤碼、堆疊追蹤、i18n 本地化與多錯誤聚合的結構化錯誤處理套件，零第三方依賴。

## 適合什麼場景

- **需要對外回傳穩定錯誤碼**：API / RPC 邊界要用數字碼區分錯誤類型，又想讓同一個錯誤依使用者語言輸出對應文案。
- **要追根究柢的堆疊資訊**：希望錯誤自帶捕獲點堆疊，`%+v` 一次印出訊息、cause 鏈與呼叫堆疊，省去自己拼湊上下文。
- **批次處理要收集多個錯誤**：迴圈驗證、並發任務等場景，需要把多個錯誤聚合成一個再一次回傳，且相容標準庫 `errors.Is` / `errors.As` 穿透。
- **想把 panic 收斂成 error**：在邊界把可能 panic 的程式碼包起來，轉成帶堆疊的一般 error，不讓 panic 逸出。

不適合純內部、無需碼也無需堆疊的簡單流程——那種情況用標準庫 `errors` 或本庫 [must](/zh-TW/modules/core/must) 斷言更輕量。

## 常用入口

### 錯誤碼與 i18n 本地化

- `New(code, msg)` / `Newf(code, format, ...)`：建立帶錯誤碼的錯誤並當場捕獲堆疊。
- `Code(err) int64`：從任意 error 提取錯誤碼（非本套件錯誤回傳零值）。
- `RegisterMessage(*language.Tag, code, msg)`：為某語言註冊錯誤碼對應的本地化訊息。
- `(*Error).LocalizedError()`：依當前 goroutine 語言輸出本地化文案。
- `(*Error).WithMetadata(key, val)`：附加結構化元資料，鏈式回傳自身。

協程級語言由 [language](/zh-TW/modules/core/language) 套件提供（`Set` / `Get`），本套件不引入 `context`。

### 堆疊包裝

- `Wrap(err, msg)` / `Wrapf(err, format, ...)`：包裝既有錯誤並補上訊息與堆疊。
- `WithStack(err)`：僅補堆疊、不改訊息。
- `Cause(err)`：解開包裝鏈，取得最根部的原始錯誤。
- `(*Error).StackTrace() []Frame`：取得結構化堆疊框（檔案、行號、函式）。
- `%+v` 格式化：一次印出訊息、cause 鏈與完整堆疊。自有實作，不依賴 `pkg/errors`。

### 多錯誤聚合

- `Join(errs...)`：聚合多個錯誤；全為 nil 回 nil，僅一個時回傳原 error。
- `Append(dst, errs...)`：把若干錯誤追加到既有錯誤上。
- `Collector{Add, ErrorOrNil, Len}`：併發安全的收集器，適合並發任務累積錯誤。

聚合結果相容標準庫 `errors.Is` / `errors.As`，可正常穿透判斷。

### panic 輔助

- `Try(fn func()) error`：執行 `fn`，捕獲 panic 並轉成帶堆疊的 error。
- `TryE(fn func() error) error`：透傳 `fn` 回傳的 error，只有 panic 才轉換。
- `Recover(*error)`：在 defer 中呼叫，把 panic 回寫到指定的 error 指標。

底層為純標準庫 `recover`，不含任何魔法。

## 使用建議

- **碼的分配集中管理**：把錯誤碼定義成常數集中一處，避免散落各檔導致碰撞或語意漂移。
- **本地化訊息在初始化時註冊**：於程式啟動階段一次性 `RegisterMessage` 各語言文案，請求路徑上只負責 `LocalizedError()` 輸出。
- **包裝而非吞錯**：跨層傳遞時用 `Wrap` 補上上下文，保留原始 cause；要看全貌時以 `%+v` 列印。
- **只在邊界用 panic 輔助**：`Try` / `Recover` 用於系統邊界收斂 panic，內部流程仍走一般 error 回傳，符合不做防禦性編程的約定。
- **批次場景優先用 Collector**：併發累積錯誤時用 `Collector` 取代手動 slice 拼接，省去自行加鎖。

## 相關文檔

- [must](/zh-TW/modules/core/must) — 錯誤斷言
- [validator](/zh-TW/modules/core/validator) — 資料驗證
- [API 概覽](/zh-TW/api/overview)
