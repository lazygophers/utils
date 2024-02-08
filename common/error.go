package common

import (
	_ "embed"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/lazygophers/log"
	"gopkg.in/yaml.v3"
	"strings"
	"time"
)

//goland:noinspection GoCommentStart
const (
	SysError = -1
	Success  = 0

	// 100-999 http状态码
	StatusContinue           = 100
	StatusSwitchingProtocols = 101
	StatusProcessing         = 102
	StatusEarlyHints         = 103

	StatusCreated              = 201
	StatusAccepted             = 202
	StatusNonAuthoritativeInfo = 203
	StatusNoContent            = 204
	StatusResetContent         = 205
	StatusPartialContent       = 206
	StatusMultiStatus          = 207
	StatusAlreadyReported      = 208
	StatusIMUsed               = 226

	StatusMultipleChoices   = 300
	StatusMovedPermanently  = 301
	StatusFound             = 302
	StatusSeeOther          = 303
	StatusNotModified       = 304
	StatusUseProxy          = 305
	_                       = 306
	StatusTemporaryRedirect = 307
	StatusPermanentRedirect = 308

	StatusBadRequest                   = 400
	StatusUnauthorized                 = 401
	StatusPaymentRequired              = 402
	StatusForbidden                    = 403
	StatusNotFound                     = 404
	StatusMethodNotAllowed             = 405
	StatusNotAcceptable                = 406
	StatusProxyAuthRequired            = 407
	StatusRequestTimeout               = 408
	StatusConflict                     = 409
	StatusGone                         = 410
	StatusLengthRequired               = 411
	StatusPreconditionFailed           = 412
	StatusRequestEntityTooLarge        = 413
	StatusRequestURITooLong            = 414
	StatusUnsupportedMediaType         = 415
	StatusRequestedRangeNotSatisfiable = 416
	StatusExpectationFailed            = 417
	StatusTeapot                       = 418
	StatusMisdirectedRequest           = 421
	StatusUnprocessableEntity          = 422
	StatusLocked                       = 423
	StatusFailedDependency             = 424
	StatusTooEarly                     = 425
	StatusUpgradeRequired              = 426
	StatusPreconditionRequired         = 428
	StatusTooManyRequests              = 429
	StatusRequestHeaderFieldsTooLarge  = 431
	StatusUnavailableForLegalReasons   = 451

	StatusInternalServerError           = 500
	StatusNotImplemented                = 501
	StatusBadGateway                    = 502
	StatusServiceUnavailable            = 503
	StatusGatewayTimeout                = 504
	StatusHTTPVersionNotSupported       = 505
	StatusVariantAlsoNegotiates         = 506
	StatusInsufficientStorage           = 507
	StatusLoopDetected                  = 508
	StatusNotExtended                   = 510
	StatusNetworkAuthenticationRequired = 511

	// 系统错误
	InvalidError    = 1001
	QueueFull       = 1002
	TokenNotFound   = 1003
	OptionConflict  = 1004
	ThirdPartyError = 1005
	RetryError      = 1006

	// 业务错误
)

var (
	errMap = make(map[int32]*Error)

	//go:embed error.yml
	errBuf []byte
)

func init() {
	var errList []*Error
	err := yaml.Unmarshal(errBuf, &errList)
	if err != nil {
		log.Panicf("err:%v", err)
		return
	}

	for _, err := range errList {
		errMap[err.Code] = err
	}
}

func RegisterError(errs ...*Error) {
	for _, err := range errs {
		errMap[err.Code] = err
	}
}

func RegisterErrorWithYaml(buf []byte) {
	var errList []*Error
	err := yaml.Unmarshal(buf, &errList)
	if err != nil {
		log.Panicf("err:%v", err)
		return
	}

	for _, err := range errList {
		errMap[err.Code] = err
	}
}

type Error struct {
	Code int32  `json:"code,omitempty" yaml:"code,omitempty"`
	Msg  string `json:"msg,omitempty" yaml:"msg,omitempty"`

	SkipRetryCount bool          `json:"skip_retry_count,omitempty" yaml:"skip_retry_count,omitempty"`
	Retry          bool          `json:"retry,omitempty" yaml:"retry,omitempty"`
	RetryDelay     time.Duration `json:"retry_delay,omitempty" yaml:"retry_delay,omitempty"`
}

func (p *Error) Error() string {
	return fmt.Sprintf("code:%d,msg:%s", p.Code, p.Msg)
}

func (p *Error) SetSkipRetryCount(b ...bool) *Error {
	if len(b) == 0 {
		p.SkipRetryCount = true
	} else {
		p.SkipRetryCount = b[0]
	}

	return p
}

func (p *Error) SetRetry(b ...bool) *Error {
	if len(b) == 0 {
		p.Retry = true
	} else {
		p.Retry = b[0]
	}

	return p
}

func (p *Error) SetRetryDelay(delay time.Duration) *Error {
	p.Retry = true
	p.RetryDelay = delay
	return p
}

func (p *Error) NeedRetry() bool {
	return p.Retry
}

func (p *Error) SetErrCode(code int32) *Error {
	p.Code = code
	return p
}

func (p *Error) Is(err error) bool {
	if err == nil {
		return false
	}

	if x, ok := err.(*Error); ok {
		return x.Code == p.Code
	} else {
		return false
	}
}

func GetErrorMsgWithCode(code int32) string {
	if err, ok := errMap[code]; ok {
		return err.Msg
	}

	return ""
}

func NewError[M int | int32](code M) *Error {
	if err, ok := errMap[int32(code)]; ok {
		return err
	}

	return NewErrorWithMsg(code, "")
}

func NewErrorWithMsg[M int | int32](code M, format string, a ...interface{}) *Error {
	return &Error{
		Code: int32(code),
		Msg:  fmt.Sprintf(format, a...),
	}
}

func ErrInvalid(a ...interface{}) *Error {
	if len(a) == 0 {
		return NewError(InvalidError)
	}

	if format, ok := a[0].(string); ok {
		return NewErrorWithMsg(InvalidError, format, a[1:]...)
	}

	return NewErrorWithMsg(InvalidError, fmt.Sprint(a...))
}

func ErrOptionConflict(a ...interface{}) *Error {
	if len(a) == 0 {
		return NewError(OptionConflict)
	}

	if format, ok := a[0].(string); ok {
		return NewErrorWithMsg(OptionConflict, format, a[1:]...)
	}

	return NewErrorWithMsg(OptionConflict, fmt.Sprint(a...))
}

func ErrThirdParty(format string, a ...interface{}) *Error {
	if format == "" {
		return NewError(ThirdPartyError)
	}

	return NewErrorWithMsg(ThirdPartyError, format, a...)
}

func ErrRetry(format string, a ...interface{}) *Error {
	if format == "" {
		return NewError(RetryError).SetRetry(true)
	}

	return NewErrorWithMsg(RetryError, format, a...).SetRetry(true)
}

func NewErrorWithError(err error) *Error {
	if err == nil {
		return nil
	}

	switch x := err.(type) {
	case *Error:
		return x
	case validator.ValidationErrors:
		var items []string
		for _, v := range x {
			items = append(items, v.Field()+" "+v.Tag()+" "+v.Param())
		}
		return NewErrorWithMsg(InvalidError, strings.Join(items, ";"))
	default:
		return &Error{
			Code: SysError,
			Msg:  err.Error(),
		}
	}
}

func GetErrCode(err error) int32 {
	if err == nil {
		return 0
	}

	if x, ok := err.(*Error); ok {
		return x.Code
	} else {
		return -1
	}
}
