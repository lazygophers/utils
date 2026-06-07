package xerror

import (
	"fmt"
	"io"
)

// Error 是 xerror 的核心错误类型，组合错误码、消息、元数据、cause 链与堆栈。
type Error struct {
	code  int64
	msg   string
	meta  map[string]string
	cause error
	stack *stack
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

// Unwrap 返回包装的下层 error，供标准库 errors.Is/As 遍历。
func (e *Error) Unwrap() error {
	return e.cause
}

// Is 在 code 非 0 时按 code 比较是否同类错误。
func (e *Error) Is(target error) bool {
	if e.code == 0 {
		return false
	}
	t, ok := target.(*Error)
	return ok && t.code == e.code
}

// Code 返回业务错误码，0 表示无码。
func (e *Error) Code() int64 {
	return e.code
}

// WithMetadata 附加一对元数据并返回自身，meta 懒分配。
func (e *Error) WithMetadata(key, val string) *Error {
	if e.meta == nil {
		e.meta = make(map[string]string, 1)
	}
	e.meta[key] = val
	return e
}

// Format 实现 fmt.Formatter：%v/%s 仅消息，%+v 附加 cause 链与堆栈。
func (e *Error) Format(f fmt.State, verb rune) {
	switch verb {
	case 'v':
		if f.Flag('+') {
			_, _ = io.WriteString(f, e.Error())
			for cause := e.cause; cause != nil; {
				_, _ = io.WriteString(f, "\n")
				_, _ = io.WriteString(f, cause.Error())
				c, ok := cause.(*Error)
				if !ok {
					break
				}
				cause = c.cause
			}
			if e.stack != nil {
				e.stack.format(f)
			}
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(f, e.Error())
	case 'q':
		fmt.Fprintf(f, "%q", e.Error())
	}
}
