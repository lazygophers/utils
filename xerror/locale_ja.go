//go:build lang_ja || lang_all

package xerror

// 日文内置错误码翻译；按 build tag 选择启用。
func init() {
	registerBuiltinLocale("ja", map[int]string{
		CodeSystem:        "システムエラー",
		CodeInvalidParam:  "パラメータが無効です",
		CodeNoAuth:        "認証されていません",
		CodeNoData:        "データが見つかりません",
		CodeConflict:      "データの競合",
		CodeNotLogin:      "ログインしていません",
		CodeTimeout:       "タイムアウト",
		CodeRateLimited:   "リクエストが多すぎます",
		CodeForbidden:     "アクセス禁止",
		CodeUnavailable:   "サービス利用不可",
		CodeDataCorrupted: "データ破損",
	})
}
