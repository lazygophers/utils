package xtime

import (
	"testing"
)

func TestSimple(t *testing.T) {
	t.Log("This is a simple test")
}

func BenchmarkSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = i
	}
}
