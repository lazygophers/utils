package validator

import (
	"testing"
)

type QuickStruct struct {
	Name string `validate:"required"`
}

func BenchmarkQuick(b *testing.B) {
	v, _ := New()
	s := QuickStruct{Name: "test"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Struct(s)
	}
}
