package xerror

import "testing"

func BenchmarkError(b *testing.B) {
	e := NewWithMsg(1, "boom")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = e.Error()
	}
}

func BenchmarkErrorWrap(b *testing.B) {
	e := NewWithMsg(1, "boom")
	root := NewWithMsg(2, "root")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = e.WithCause(root)
	}
}
