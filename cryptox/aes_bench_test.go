package cryptox

import (
	"testing"
)

// Baseline benchmarks for AES ECB/CBC/CTR
func BenchmarkAESBaselineECB(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = EncryptECB(key, plaintext)
	}
}

func BenchmarkAESBaselineCBC(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = EncryptCBC(key, plaintext)
	}
}

func BenchmarkAESBaselineCTR(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = EncryptCTR(key, plaintext)
	}
}
