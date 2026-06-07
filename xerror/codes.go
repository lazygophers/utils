package xerror

// 框架内置错误码段（1001-10000 预留给框架；业务码建议 ≥ 10001）。
// 命名约定：Code<Concept>；构造器 New<Concept>(args ...any) *Error。
// 内置翻译按语言拆 codes_<lang>.go 文件，在各自 init() 中注册到 i18n.Default。
const (
	// CodeInvalidParam 请求参数无效或缺失。
	CodeInvalidParam = 1001

	// CodeNoAuth 未授权或认证失败。
	CodeNoAuth = 1002

	// CodeNoData 请求的数据不存在。
	CodeNoData = 1003

	// CodeConflict 数据冲突。
	CodeConflict = 1004

	// CodeNotLogin 登录状态异常。
	CodeNotLogin = 1005

	// CodeTimeout 操作超时。
	CodeTimeout = 1006

	// CodeRateLimited 触发限流。
	CodeRateLimited = 1007

	// CodeForbidden 操作被拒绝（已认证但无权限）。
	CodeForbidden = 1008

	// CodeUnavailable 服务暂不可用。
	CodeUnavailable = 1009

	// CodeDataCorrupted 数据损坏或不一致。
	CodeDataCorrupted = 1010
)

// NewSystemError 创建系统级错误（code = CodeSystem）。
func NewSystemError(args ...any) *Error {
	return New(CodeSystem, args...)
}

// NewInvalidParam 创建参数无效错误（code = CodeInvalidParam）。
func NewInvalidParam(args ...any) *Error {
	return New(CodeInvalidParam, args...)
}

// NewNoAuth 创建未授权错误（code = CodeNoAuth）。
func NewNoAuth(args ...any) *Error {
	return New(CodeNoAuth, args...)
}

// NewNoData 创建数据不存在错误（code = CodeNoData）。
func NewNoData(args ...any) *Error {
	return New(CodeNoData, args...)
}

// NewConflict 创建数据冲突错误（code = CodeConflict）。
func NewConflict(args ...any) *Error {
	return New(CodeConflict, args...)
}

// NewNotLogin 创建未登录错误（code = CodeNotLogin）。
func NewNotLogin(args ...any) *Error {
	return New(CodeNotLogin, args...)
}

// NewTimeout 创建超时错误（code = CodeTimeout）。
func NewTimeout(args ...any) *Error {
	return New(CodeTimeout, args...)
}

// NewRateLimited 创建限流错误（code = CodeRateLimited）。
func NewRateLimited(args ...any) *Error {
	return New(CodeRateLimited, args...)
}

// NewForbidden 创建禁止操作错误（code = CodeForbidden）。
func NewForbidden(args ...any) *Error {
	return New(CodeForbidden, args...)
}

// NewUnavailable 创建服务不可用错误（code = CodeUnavailable）。
func NewUnavailable(args ...any) *Error {
	return New(CodeUnavailable, args...)
}

// NewDataCorrupted 创建数据损坏错误（code = CodeDataCorrupted）。
func NewDataCorrupted(args ...any) *Error {
	return New(CodeDataCorrupted, args...)
}

// registerBuiltinLocale 把单语言内置错误码翻译表注册到 i18n.Default。
// 各 codes_<lang>.go 在 init() 中调用此函数，复用 errorKey 拼装与 i18n.Default 写入逻辑。
func registerBuiltinLocale(langStr string, codes map[int]string) {
	tag := makeLangTag(langStr)
	for code, msg := range codes {
		defaultRegister(tag, code, msg)
	}
}
