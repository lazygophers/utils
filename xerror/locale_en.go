package xerror

// 英文内置错误码翻译；en/zh 始终注册，不加 build tag。
func init() {
	registerBuiltinLocale("en", map[int]string{
		CodeSystem:        "system error",
		CodeInvalidParam:  "invalid parameter",
		CodeNoAuth:        "unauthorized",
		CodeNoData:        "not found",
		CodeConflict:      "conflict",
		CodeNotLogin:      "not logged in",
		CodeTimeout:       "timeout",
		CodeRateLimited:   "too many requests",
		CodeForbidden:     "forbidden",
		CodeUnavailable:   "service unavailable",
		CodeDataCorrupted: "data corrupted",
	})
}
