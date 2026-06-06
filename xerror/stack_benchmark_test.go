package xerror

import (
	"errors"
	"testing"
)

func BenchmarkNew(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = New(1, "boom")
	}
}

func BenchmarkNewf(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Newf(1, "boom %d", i)
	}
}

func BenchmarkWrap(b *testing.B) {
	root := errors.New("root")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Wrap(root, "wrapped")
	}
}

func BenchmarkWithStack(b *testing.B) {
	b.Run("NewCapture", func(b *testing.B) {
		root := errors.New("root")
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = WithStack(root)
		}
	})
	b.Run("ReuseStack", func(b *testing.B) {
		e := New(1, "boom")
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = WithStack(e)
		}
	})
}

func BenchmarkStackTrace(b *testing.B) {
	e := New(1, "boom")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = e.StackTrace()
	}
}

func BenchmarkCause(b *testing.B) {
	root := errors.New("root")
	chain := Wrap(Wrap(Wrap(root, "l1"), "l2"), "l3")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Cause(chain)
	}
}
