package i18n

import "testing"

func BenchmarkInterpolate_NoBraceFastPath(b *testing.B) {
	tpl := "hello world, no placeholders here"
	args := []any{"name", "Alice"}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = interpolate(tpl, args)
	}
}

func BenchmarkInterpolate_NoArgs(b *testing.B) {
	tpl := "hello world"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = interpolate(tpl, nil)
	}
}

func BenchmarkInterpolate_NamedSingle(b *testing.B) {
	tpl := "hello {name}"
	args := []any{"name", "Alice"}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = interpolate(tpl, args)
	}
}

func BenchmarkInterpolate_NamedMulti(b *testing.B) {
	tpl := "{greet}, {name}! You have {n} messages."
	args := []any{"greet", "Hi", "name", "Bob", "n", 42}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = interpolate(tpl, args)
	}
}

func BenchmarkInterpolate_Positional(b *testing.B) {
	tpl := "{0} loves {1} (#{2})"
	args := []any{"Alice", "Go", 7}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = interpolate(tpl, args)
	}
}
