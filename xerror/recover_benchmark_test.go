package xerror

import (
	"errors"
	"testing"
)

var errBench = errors.New("bench error")

// BenchmarkTryNoPanic 衡量正常执行（无 panic）下 defer recover 的开销，期望接近 0 alloc。
func BenchmarkTryNoPanic(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Try(func() {})
	}
}

// BenchmarkTryPanic 衡量 panic 触发转 error 路径（含栈捕获分配）。
func BenchmarkTryPanic(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Try(func() { panic("boom") })
	}
}

// BenchmarkTryEPassthrough 衡量 fn 返回 error 直接透传路径。
func BenchmarkTryEPassthrough(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = TryE(func() error { return errBench })
	}
}

// BenchmarkTryENoError 衡量 fn 返回 nil 路径。
func BenchmarkTryENoError(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = TryE(func() error { return nil })
	}
}

// BenchmarkRecoverNoPanic 衡量 defer Recover 在无 panic 下的开销。
func BenchmarkRecoverNoPanic(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		func() (err error) {
			defer Recover(&err)
			return nil
		}()
	}
}
