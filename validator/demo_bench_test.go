package validator

import (
	"testing"
)

func BenchmarkDemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = i * 2
	}
}
