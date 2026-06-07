package xerror

import (
	"errors"

	xlanguage "golang.org/x/text/language"
)

// Unwrapper 等价 stdlib errors 包内部匿名的 Unwrap 单错误契约。
// stdlib 未导出命名版本（errors/wrap.go 全部匿名）；本包导出以便 var _ 编译期校验与外部 mock。
type Unwrapper interface {
	Unwrap() error
}

// MultiUnwrapper 等价 stdlib errors 包内部匿名的 Unwrap 多错误契约（Go 1.20+）。
type MultiUnwrapper interface {
	Unwrap() []error
}

// 编译期接口契约校验：*Error 必须满足 error + Unwrapper。
var (
	_ error     = (*Error)(nil)
	_ Unwrapper = (*Error)(nil)
)

// Error 是 xerror 的核心错误类型，组合错误码、消息与 cause 链。
type Error struct {
	code  int
	msg   string
	cause error
}

// Error 返回错误消息；无消息时回退 cause 的消息。
func (e *Error) Error() string {
	if e.msg != "" {
		return e.msg
	}
	if e.cause != nil {
		return e.cause.Error()
	}
	return ""
}

// Msg 返回构造时的原始消息字段（不回退 cause），用于调试或日志区分。
func (e *Error) Msg() string {
	return e.msg
}

// Unwrap 返回包装的下层 error，供标准库 errors.Is/As 遍历。
func (e *Error) Unwrap() error {
	return e.cause
}

// Code 返回业务错误码。CodeSuccess(0) 表示成功；CodeSystem(-1) 表示未指定。
func (e *Error) Code() int {
	return e.code
}

// WithCause 流式设置 cause 并返回自身；多次调用覆盖前次 cause。
// 区别于包级 Wrap：方法版仅挂载，不创建新实例；包级 Wrap 创建新 *Error 包装。
//
//	e := xerror.New(1001, "biz fail").WithCause(rootErr)
func (e *Error) WithCause(cause error) *Error {
	e.cause = cause
	return e
}

// New 创建带错误码的 *Error，按当前 goroutine 语言走 Localizer 翻译。
// args 注入翻译模板（命名 "k", v 或位置 {0}）。未注入 Localizer 或未命中翻译时 msg 为空字符串。
// 解析在构造时完成，*Error 绑定到构造时的 goroutine 语言。
func New(code int, args ...any) *Error {
	return &Error{code: code, msg: resolveMsg(code, "", args)}
}

// NewWithMsg 创建带 fallback 消息的 *Error。
// 若 Localizer 命中翻译用翻译结果（args 注入模板）；未命中则用 msg。
func NewWithMsg(code int, msg string, args ...any) *Error {
	return &Error{code: code, msg: resolveMsg(code, msg, args)}
}

// NewWithLanguage 创建带错误码的 *Error，按指定语言（golang.org/x/text/language.Tag 标准库类型）
// 走 Localizer 翻译；不读 goroutine-local 语言。args 注入翻译模板；未命中时 msg 为空。
func NewWithLanguage(tag xlanguage.Tag, code int, args ...any) *Error {
	return &Error{code: code, msg: resolveMsgWithLang(tag, code, "", args)}
}

// Wrap 用 msg 包装 err；err 为 nil 时透传 nil（避免 stdlib (*nil, false) 陷阱）。
// 默认 code = CodeSystem；msg + args 与 New 同语义（翻译命中则用翻译模板）。
// 返回类型显式为 error，避免 *Error nil 接口非 nil 的坑；
// 包装后 errors.Is/As/Unwrap 全链可穿透到原 err。
func Wrap(err error, msg string, args ...any) error {
	if err == nil {
		return nil
	}
	return &Error{code: CodeSystem, msg: resolveMsg(CodeSystem, msg, args), cause: err}
}

// Wraps 把多个 error 合并作为 cause 包装；全 nil 返 nil。
// cause 通过 Join 合并（语义对齐 stdlib errors.Join），
// errors.Is/As 沿 *Error.Unwrap → multiError.Unwrap() []error 可遍历任一子 error。
func Wraps(errs ...error) error {
	cause := Join(errs...)
	if cause == nil {
		return nil
	}
	return &Error{code: CodeSystem, cause: cause}
}

// Cause 沿 cause 链解到最底层根错误。
// 兼容任何实现 Unwrap() error 的类型（含 *Error / fmt.Errorf("%w") 等），
// 遇到 Unwrap() []error 的聚合错误时停在该层（多 cause 无单一根）。
func Cause(err error) error {
	for {
		next := errors.Unwrap(err)
		if next == nil {
			return err
		}
		err = next
	}
}
