package xerror

import (
	"testing"

	"github.com/lazygophers/utils/language"
	xlanguage "golang.org/x/text/language"
)

// BenchmarkCode 测量从 *Error 提取错误码的开销。
func BenchmarkCode(b *testing.B) {
	err := NewWithMsg(1001, "internal server error")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Code(err)
	}
}

// BenchmarkNewNoLocalizer 测量未注入 Localizer 时 New 的开销（短路）。
func BenchmarkNewNoLocalizer(b *testing.B) {
	SetLocalizer(nil)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewWithMsg(1001, "default")
	}
}

// BenchmarkNewLocalizerHit 注入 Localizer 且命中翻译时 New 的开销。
func BenchmarkNewLocalizerHit(b *testing.B) {
	stub := newStubLocalizer()
	stub.Register(xlanguage.Make("en"), "error.1001", "translated")
	SetLocalizer(stub)
	defer SetLocalizer(nil)
	language.Set(language.Make("en"))
	defer language.Del()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewWithMsg(1001, "default")
	}
}

// BenchmarkNewLocalizerMiss 注入 Localizer 但未命中（回退 fallback）。
func BenchmarkNewLocalizerMiss(b *testing.B) {
	stub := newStubLocalizer()
	SetLocalizer(stub)
	defer SetLocalizer(nil)
	language.Set(language.Make("en"))
	defer language.Del()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = NewWithMsg(999, "default")
	}
}
