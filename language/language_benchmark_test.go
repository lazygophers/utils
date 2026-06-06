package language

import (
	"testing"
)

func BenchmarkMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Make("zh-CN")
	}
}

func BenchmarkFallbackChain(b *testing.B) {
	t := Make("zh-Hant-TW")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = t.FallbackChain()
	}
}

func BenchmarkMatch_Hit(b *testing.B) {
	a := Make("zh-CN")
	c := Make("zh")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = a.Match(c)
	}
}

func BenchmarkMatch_Miss(b *testing.B) {
	a := Make("zh-CN")
	c := Make("en")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = a.Match(c)
	}
}

func BenchmarkParseAcceptLanguage_Single(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ParseAcceptLanguage("zh-CN")
	}
}

func BenchmarkParseAcceptLanguage_Multi(b *testing.B) {
	const h = "da, en-gb;q=0.8, en;q=0.7, zh-CN;q=0.9"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ParseAcceptLanguage(h)
	}
}

func BenchmarkDetect(b *testing.B) {
	supported := []*Tag{Make("en"), Make("zh"), Make("ja")}
	for i := 0; i < b.N; i++ {
		_, _ = Detect("zh-CN,en;q=0.8", supported)
	}
}

func BenchmarkGet_NoLocal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Get()
	}
}

func BenchmarkGet_WithLocal(b *testing.B) {
	Set(Make("zh-CN"))
	defer Del()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Get()
	}
}

func BenchmarkSet(b *testing.B) {
	tag := Make("zh-CN")
	for i := 0; i < b.N; i++ {
		Set(tag)
	}
	Del()
}
