package xerror

// Try 执行 fn，捕获其 panic 转为带栈 *Error；正常执行返回 nil。
func Try(fn func()) (err error) {
	defer Recover(&err)
	fn()
	return nil
}

// TryE 执行 fn，fn 返回的 error 直接透传，仅 panic 时转为带栈 *Error。
func TryE(fn func() error) (err error) {
	defer Recover(&err)
	return fn()
}

// Recover 供 defer 调用，将当前 panic 写入 *errp 并附加调用栈；无 panic 时不改动。
func Recover(errp *error) {
	r := recover()
	if r == nil {
		return
	}
	if e, ok := r.(error); ok {
		*errp = WithStack(e)
		return
	}
	*errp = Newf(0, "%v", r)
}
