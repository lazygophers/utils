package xerror

// 中文（简体）内置错误码翻译；en/zh 始终注册，不加 build tag。
func init() {
	registerBuiltinLocale("zh", map[int]string{
		CodeSystem:        "系统错误",
		CodeInvalidParam:  "参数无效",
		CodeNoAuth:        "未授权",
		CodeNoData:        "数据不存在",
		CodeConflict:      "数据冲突",
		CodeNotLogin:      "未登录",
		CodeTimeout:       "请求超时",
		CodeRateLimited:   "请求过于频繁",
		CodeForbidden:     "禁止访问",
		CodeUnavailable:   "服务不可用",
		CodeDataCorrupted: "数据损坏",
	})
}
