package candy

import "testing"

func BenchmarkSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ToPtr(42)
	}
}
