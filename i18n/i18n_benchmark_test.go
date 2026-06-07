package i18n

import (
	"testing"

	xlanguage "golang.org/x/text/language"
)

func BenchmarkLocalizeHit(b *testing.B) {
	p := New()
	en := xlanguage.Make("en")
	p.Register(en, "hello", "Hello")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.LocalizeWithLang(en, "hello")
	}
}

func BenchmarkLocalizeMiss(b *testing.B) {
	p := New()
	en := xlanguage.Make("en")
	p.Register(en, "hello", "Hello")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.LocalizeWithLang(en, "missing")
	}
}

func BenchmarkLocalizeTemplate(b *testing.B) {
	p := New()
	en := xlanguage.Make("en")
	p.Register(en, "greet", "Hello {{.Name}}, you have {{.Count}} msg")
	data := map[string]any{"Name": "Alice", "Count": 3}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.LocalizeWithLang(en, "greet", data)
	}
}

func BenchmarkLocalizeFallback(b *testing.B) {
	p := New(WithDefaultLang(xlanguage.Make("en")))
	p.Register(xlanguage.Make("en"), "hello", "Hello")
	zhCN := xlanguage.Make("zh-CN")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// zh-CN → zh → en
		_ = p.LocalizeWithLang(zhCN, "hello")
	}
}
