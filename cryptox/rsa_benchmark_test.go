package cryptox

import (
	"testing"
)

func BenchmarkRSAPrivateKeyToPEM(b *testing.B) {
	keyPair, _ := GenerateRSAKeyPair(2048)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = keyPair.PrivateKeyToPEM()
	}
}

func BenchmarkRSAPublicKeyToPEM(b *testing.B) {
	keyPair, _ := GenerateRSAKeyPair(2048)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = keyPair.PublicKeyToPEM()
	}
}

func BenchmarkRSAEncryptOAEP(b *testing.B) {
	keyPair, _ := GenerateRSAKeyPair(2048)
	plaintext := []byte("Hello, World!")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = RSAEncryptOAEP(keyPair.PublicKey, plaintext)
	}
}

func BenchmarkRSADecryptOAEP(b *testing.B) {
	keyPair, _ := GenerateRSAKeyPair(2048)
	plaintext := []byte("Hello, World!")
	ciphertext, _ := RSAEncryptOAEP(keyPair.PublicKey, plaintext)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = RSADecryptOAEP(keyPair.PrivateKey, ciphertext)
	}
}
