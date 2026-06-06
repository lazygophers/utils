package xerror

import (
	"errors"
	"testing"
)

func BenchmarkJoinAllNil(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Join(nil, nil, nil)
	}
}

func BenchmarkJoinSingle(b *testing.B) {
	e := errors.New("x")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Join(nil, e)
	}
}

func BenchmarkJoinMultiple(b *testing.B) {
	e1, e2, e3 := errors.New("a"), errors.New("b"), errors.New("c")
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Join(e1, e2, e3)
	}
}

func BenchmarkCollectorAdd(b *testing.B) {
	e := errors.New("x")
	var c Collector
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c.Add(e)
	}
}

func BenchmarkCollectorConcurrentAdd(b *testing.B) {
	e := errors.New("x")
	var c Collector
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Add(e)
		}
	})
}
