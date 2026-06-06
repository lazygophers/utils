package xerror

import (
	"fmt"
	"io"
	"testing"
)

func BenchmarkError(b *testing.B) {
	e := New(1, "boom")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = e.Error()
	}
}

func BenchmarkWithMetadata(b *testing.B) {
	b.Run("FirstAlloc", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			e := &Error{code: 1, msg: "boom"}
			e.WithMetadata("k", "v")
		}
	})
	b.Run("Allocated", func(b *testing.B) {
		e := New(1, "boom").WithMetadata("k", "v")
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			e.WithMetadata("k", "v")
		}
	})
}

func BenchmarkFormatV(b *testing.B) {
	e := New(1, "boom")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(io.Discard, "%v", e)
	}
}

func BenchmarkFormatPlusV(b *testing.B) {
	e := Wrap(New(1, "root"), "wrapped")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fmt.Fprintf(io.Discard, "%+v", e)
	}
}
