package xerror

import (
	"testing"

	"github.com/lazygophers/utils/language"
)

// BenchmarkCode 测量从 *Error 提取错误码的开销。
func BenchmarkCode(b *testing.B) {
	err := New(1001, "internal server error")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Code(err)
	}
}

// BenchmarkRegisterMessage 测量注册一条本地化消息的开销（含写锁）。
func BenchmarkRegisterMessage(b *testing.B) {
	en := language.Make("en")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RegisterMessage(en, 1001, "internal server error")
	}
}

// BenchmarkLocalizedError 测量命中已注册消息的本地化查询开销（含 language.Get + map 查询）。
func BenchmarkLocalizedError(b *testing.B) {
	language.Set(language.Make("zh-CN"))
	defer language.Del()
	err := New(1001, "default")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = err.LocalizedError()
	}
}

// BenchmarkLocalizedErrorMiss 测量未注册消息回退默认 msg 的开销。
func BenchmarkLocalizedErrorMiss(b *testing.B) {
	language.Set(language.Make("zh-CN"))
	defer language.Del()
	err := New(999999, "default message")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = err.LocalizedError()
	}
}
