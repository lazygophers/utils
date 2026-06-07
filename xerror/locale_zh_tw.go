//go:build lang_zh_tw || lang_all

package xerror

// 中文（繁体）内置错误码翻译；按 build tag 选择启用。
func init() {
	registerBuiltinLocale("zh-Hant", map[int]string{
		CodeSystem:        "系統錯誤",
		CodeInvalidParam:  "參數無效",
		CodeNoAuth:        "未授權",
		CodeNoData:        "資料不存在",
		CodeConflict:      "資料衝突",
		CodeNotLogin:      "未登入",
		CodeTimeout:       "請求逾時",
		CodeRateLimited:   "請求過於頻繁",
		CodeForbidden:     "禁止存取",
		CodeUnavailable:   "服務不可用",
		CodeDataCorrupted: "資料損壞",
	})
}
