package validator

import "testing"

func TestSimple(t *testing.T) {
	t.Log("test runs")
}

func BenchmarkSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = i
	}
}
