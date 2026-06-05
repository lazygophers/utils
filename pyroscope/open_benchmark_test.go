package pyroscope

import (
	"testing"
)

func BenchmarkLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Load("http://benchmark.server:4040")
	}
}

func BenchmarkLoadEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Load("")
	}
}
