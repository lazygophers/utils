package xerror

import (
	"fmt"
	"path"
	"runtime"
	"strconv"
)

// stack 保存调用栈的程序计数器，懒解析为 Frame。
type stack struct {
	pcs []uintptr
}

// Frame 表示调用栈中的单帧，包含函数名、文件与行号。
type Frame struct {
	Function string
	File     string
	Line     int
}

// callers 捕获当前调用栈，skip 为跳过的栈帧数（含 callers 自身）。
func callers(skip int) *stack {
	var pcs [32]uintptr
	n := runtime.Callers(skip+1, pcs[:])
	s := make([]uintptr, n)
	copy(s, pcs[:n])
	return &stack{pcs: s}
}

// frames 将程序计数器解析为可读的 Frame 列表。
func (s *stack) frames() []Frame {
	frames := runtime.CallersFrames(s.pcs)
	out := make([]Frame, 0, len(s.pcs))
	for {
		f, more := frames.Next()
		out = append(out, Frame{Function: f.Function, File: f.File, Line: f.Line})
		if !more {
			break
		}
	}
	return out
}

// format 将堆栈帧按 `\nfunc\n\tfile:line` 形式写入。
func (s *stack) format(f fmt.State) {
	for _, fr := range s.frames() {
		fmt.Fprintf(f, "\n%s\n\t%s:%s", fr.Function, trimPath(fr.File), strconv.Itoa(fr.Line))
	}
}

// StackTrace 返回错误捕获时的调用栈帧；未捕获栈时返回 nil。
func (e *Error) StackTrace() []Frame {
	if e.stack == nil {
		return nil
	}
	return e.stack.frames()
}

// New 创建带错误码与消息的 *Error，并捕获当前调用栈。
func New(code int64, msg string) *Error {
	return &Error{code: code, msg: msg, stack: callers(2)}
}

// Newf 创建带错误码的 *Error，消息按 format 格式化，并捕获当前调用栈。
func Newf(code int64, format string, a ...any) *Error {
	return &Error{code: code, msg: fmt.Sprintf(format, a...), stack: callers(2)}
}

// Wrap 用 msg 包装 err 并附加调用栈；err 为 nil 时透传 nil。
func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	return &Error{msg: msg, cause: err, stack: stackFor(err, 2)}
}

// Wrapf 用格式化消息包装 err 并附加调用栈；err 为 nil 时透传 nil。
func Wrapf(err error, format string, a ...any) error {
	if err == nil {
		return nil
	}
	return &Error{msg: fmt.Sprintf(format, a...), cause: err, stack: stackFor(err, 2)}
}

// WithStack 为 err 附加调用栈而不添加消息；err 为 nil 时透传 nil。
func WithStack(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(*Error); ok && e.stack != nil {
		return e
	}
	return &Error{cause: err, stack: callers(2)}
}

// Cause 沿 cause 链解到最底层的根错误。
func Cause(err error) error {
	for err != nil {
		e, ok := err.(*Error)
		if !ok || e.cause == nil {
			break
		}
		err = e.cause
	}
	return err
}

// stackFor 复用已含栈的 *Error 栈，否则在 skip 处捕获新栈。
func stackFor(err error, skip int) *stack {
	if e, ok := err.(*Error); ok && e.stack != nil {
		return e.stack
	}
	return callers(skip + 1)
}

// trimPath 返回文件的短路径（包名/文件名），用于堆栈打印。
func trimPath(file string) string {
	return path.Join(path.Base(path.Dir(file)), path.Base(file))
}
