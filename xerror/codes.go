package xerror

// 框架内置错误码段（1001-10000 预留给框架；业务码建议 ≥ 10001）。
// 命名约定：Code<Concept>；构造器 New<Concept>(args ...any) *Error。
const (
	// CodeInvalidParam 请求参数无效或缺失。
	// 场景：参数校验失败 / 必填缺失 / 格式错误。
	CodeInvalidParam = 1001

	// CodeNoAuth 未授权或认证失败。
	// 场景：token 无效 / 未登录 / 权限不足。
	CodeNoAuth = 1002

	// CodeNoData 请求的数据不存在。
	// 场景:查询记录不存在 / 资源未找到。
	CodeNoData = 1003

	// CodeConflict 数据冲突。
	// 场景：并发更新冲突 / 唯一键冲突 / 版本不匹配。
	CodeConflict = 1004

	// CodeNotLogin 登录状态异常。
	// 场景：未获取到登录态 / 登录态过期。
	CodeNotLogin = 1005

	// CodeTimeout 操作超时。
	// 场景：下游 RPC 超时 / 数据库慢查询 / 第三方接口未响应。
	CodeTimeout = 1006

	// CodeRateLimited 触发限流。
	// 场景：请求过于频繁 / 配额耗尽。
	CodeRateLimited = 1007

	// CodeForbidden 操作被拒绝（已认证但无权限）。
	// 场景：用户已登录但缺少特定权限位 / 资源访问被策略禁止。
	CodeForbidden = 1008

	// CodeUnavailable 服务暂不可用。
	// 场景：依赖下游故障 / 维护中 / 熔断打开。
	CodeUnavailable = 1009

	// CodeDataCorrupted 数据损坏或不一致。
	// 场景：反序列化失败 / 校验和不匹配 / 数据完整性错误。
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
