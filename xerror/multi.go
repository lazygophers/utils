package xerror

import (
	"strings"
	"sync"
)

// 编译期接口契约校验：*multiError 满足 error + MultiUnwrapper（errors.Is/As 多 cause 遍历）。
var (
	_ error          = (*multiError)(nil)
	_ MultiUnwrapper = (*multiError)(nil)
)

// multiError 聚合多个 error，实现标准库多错误语义。
type multiError struct {
	errs []error
}

// Error 用换行连接所有子错误消息。
func (m *multiError) Error() string {
	if len(m.errs) == 1 {
		return m.errs[0].Error()
	}
	msgs := make([]string, len(m.errs))
	total := len(m.errs) - 1
	for i, err := range m.errs {
		msgs[i] = err.Error()
		total += len(msgs[i])
	}
	var b strings.Builder
	b.Grow(total)
	for i, msg := range msgs {
		if i > 0 {
			b.WriteByte('\n')
		}
		b.WriteString(msg)
	}
	return b.String()
}

// Unwrap 返回子错误切片，使 errors.Is/As 自动遍历。
func (m *multiError) Unwrap() []error {
	return m.errs
}

// Join 合并多个 error：过滤 nil；全为 nil 返 nil；仅一个非 nil 直接返回该 error。
func Join(errs ...error) error {
	merged := make([]error, 0, len(errs))
	for _, err := range errs {
		if err != nil {
			merged = append(merged, err)
		}
	}
	switch len(merged) {
	case 0:
		return nil
	case 1:
		return merged[0]
	}
	return &multiError{errs: merged}
}

// Append 把 errs 追加到 dst：dst 为聚合错误时原地扩展，否则新建合并。
func Append(dst error, errs ...error) error {
	if m, ok := dst.(*multiError); ok {
		for _, err := range errs {
			if err != nil {
				m.errs = append(m.errs, err)
			}
		}
		return m
	}
	all := make([]error, 0, len(errs)+1)
	all = append(all, dst)
	all = append(all, errs...)
	return Join(all...)
}

// Collector 是并发安全的错误收集器，多 goroutine 可并发 Add。
type Collector struct {
	mu   sync.Mutex
	errs []error
}

// Add 收集一个 error，nil 忽略。
func (c *Collector) Add(err error) {
	if err == nil {
		return
	}
	c.mu.Lock()
	c.errs = append(c.errs, err)
	c.mu.Unlock()
}

// ErrorOrNil 返回聚合错误，无错时返 nil。
// 锁内只做切片快照，Join 在锁外执行减少持有时间。
func (c *Collector) ErrorOrNil() error {
	c.mu.Lock()
	if len(c.errs) == 0 {
		c.mu.Unlock()
		return nil
	}
	snap := make([]error, len(c.errs))
	copy(snap, c.errs)
	c.mu.Unlock()
	return Join(snap...)
}

// Len 返回已收集的错误数量。
func (c *Collector) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.errs)
}
