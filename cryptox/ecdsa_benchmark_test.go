package cryptox

import (
	"crypto/sha256"
	"testing"
)

func BenchmarkECDSASignOriginal(b *testing.B) {
	privateKey, _ := GenerateECDSAP256Key()
	data := []byte("test data for benchmarking ECDSA signature operations")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = ECDSASign(privateKey.PrivateKey, data, sha256.New)
	}
}

func BenchmarkECDSASignOpt1(b *testing.B) {
	privateKey, _ := GenerateECDSAP256Key()
	data := []byte("test data for benchmarking ECDSA signature operations")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = ECDSASignOpt1(privateKey.PrivateKey, data)
	}
}

func BenchmarkECDSASignOpt2(b *testing.B) {
	privateKey, _ := GenerateECDSAP256Key()
	data := []byte("test data for benchmarking ECDSA signature operations")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = ECDSASignOpt2(privateKey.PrivateKey, data)
	}
}

func BenchmarkECDSASignOpt3(b *testing.B) {
	privateKey, _ := GenerateECDSAP256Key()
	data := []byte("test data for benchmarking ECDSA signature operations")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = ECDSASignOpt3(privateKey.PrivateKey, data, sha256.New)
	}
}

func BenchmarkECDSASignOpt4(b *testing.B) {
	privateKey, _ := GenerateECDSAP256Key()
	data := []byte("test data for benchmarking ECDSA signature operations")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = ECDSASignOpt4(privateKey.PrivateKey, data)
	}
}

func BenchmarkECDSASignOpt5(b *testing.B) {
	privateKey, _ := GenerateECDSAP256Key()
	data := []byte("test data for benchmarking ECDSA signature operations")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = ECDSASignOpt5(privateKey.PrivateKey, data)
	}
}

func BenchmarkECDSAVerifyOriginal(b *testing.B) {
	privateKey, _ := GenerateECDSAP256Key()
	data := []byte("test data for benchmarking ECDSA verify operations")
	r, s, _ := ECDSASign(privateKey.PrivateKey, data, sha256.New)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ECDSAVerify(privateKey.PublicKey, data, r, s, sha256.New)
	}
}

func BenchmarkECDSAVerifyOpt6(b *testing.B) {
	privateKey, _ := GenerateECDSAP256Key()
	data := []byte("test data for benchmarking ECDSA verify operations")
	r, s, _ := ECDSASign(privateKey.PrivateKey, data, sha256.New)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ECDSAVerifyOpt6(privateKey.PublicKey, data, r, s)
	}
}

func BenchmarkECDSAVerifyOpt7(b *testing.B) {
	privateKey, _ := GenerateECDSAP256Key()
	data := []byte("test data for benchmarking ECDSA verify operations")
	r, s, _ := ECDSASign(privateKey.PrivateKey, data, sha256.New)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ECDSAVerifyOpt7(privateKey.PublicKey, data, r, s)
	}
}

func BenchmarkECDSAVerifyOpt8(b *testing.B) {
	privateKey, _ := GenerateECDSAP256Key()
	data := []byte("test data for benchmarking ECDSA verify operations")
	r, s, _ := ECDSASign(privateKey.PrivateKey, data, sha256.New)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ECDSAVerifyOpt8(privateKey.PublicKey, data, r, s, sha256.New)
	}
}

func BenchmarkECDSAVerifyOpt9(b *testing.B) {
	privateKey, _ := GenerateECDSAP256Key()
	data := []byte("test data for benchmarking ECDSA verify operations")
	r, s, _ := ECDSASign(privateKey.PrivateKey, data, sha256.New)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ECDSAVerifyOpt9(privateKey.PublicKey, data, r, s)
	}
}

func BenchmarkECDSAVerifyOpt10(b *testing.B) {
	privateKey, _ := GenerateECDSAP256Key()
	data := []byte("test data for benchmarking ECDSA verify operations")
	r, s, _ := ECDSASign(privateKey.PrivateKey, data, sha256.New)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ECDSAVerifyOpt10(privateKey.PublicKey, data, r, s)
	}
}

func BenchmarkECDHComputeSharedOriginal(b *testing.B) {
	alicePrivate, _ := GenerateECDHP256Key()
	bobPublic, _ := GenerateECDHP256Key()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ECDHComputeShared(alicePrivate.PrivateKey, bobPublic.PublicKey)
	}
}

func BenchmarkECDHComputeSharedOpt1(b *testing.B) {
	alicePrivate, _ := GenerateECDHP256Key()
	bobPublic, _ := GenerateECDHP256Key()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ECDHComputeSharedOpt1(alicePrivate.PrivateKey, bobPublic.PublicKey)
	}
}

func BenchmarkECDHComputeSharedOpt2(b *testing.B) {
	alicePrivate, _ := GenerateECDHP256Key()
	bobPublic, _ := GenerateECDHP256Key()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ECDHComputeSharedOpt2(alicePrivate.PrivateKey, bobPublic.PublicKey)
	}
}

func BenchmarkECDHComputeSharedOpt3(b *testing.B) {
	alicePrivate, _ := GenerateECDHP256Key()
	bobPublic, _ := GenerateECDHP256Key()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ECDHComputeSharedOpt3(alicePrivate.PrivateKey, bobPublic.PublicKey)
	}
}

func BenchmarkECDHComputeSharedOpt4(b *testing.B) {
	alicePrivate, _ := GenerateECDHP256Key()
	bobPublic, _ := GenerateECDHP256Key()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ECDHComputeSharedOpt4(alicePrivate.PrivateKey, bobPublic.PublicKey)
	}
}

func BenchmarkECDHComputeSharedOpt5(b *testing.B) {
	alicePrivate, _ := GenerateECDHP256Key()
	bobPublic, _ := GenerateECDHP256Key()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ECDHComputeSharedOpt5(alicePrivate.PrivateKey, bobPublic.PublicKey)
	}
}

func BenchmarkECDHComputeSharedWithKDFOriginal(b *testing.B) {
	alicePrivate, _ := GenerateECDHP256Key()
	bobPublic, _ := GenerateECDHP256Key()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ECDHComputeSharedWithKDF(alicePrivate.PrivateKey, bobPublic.PublicKey, 32, sha256.New)
	}
}

func BenchmarkECDHComputeSharedWithKDFOpt6(b *testing.B) {
	alicePrivate, _ := GenerateECDHP256Key()
	bobPublic, _ := GenerateECDHP256Key()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ECDHComputeSharedWithKDFOpt6(alicePrivate.PrivateKey, bobPublic.PublicKey, 32, sha256.New)
	}
}

func BenchmarkECDHComputeSharedWithKDFOpt7(b *testing.B) {
	alicePrivate, _ := GenerateECDHP256Key()
	bobPublic, _ := GenerateECDHP256Key()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ECDHComputeSharedWithKDFOpt7(alicePrivate.PrivateKey, bobPublic.PublicKey, 32)
	}
}

func BenchmarkECDHComputeSharedWithKDFOpt8(b *testing.B) {
	alicePrivate, _ := GenerateECDHP256Key()
	bobPublic, _ := GenerateECDHP256Key()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ECDHComputeSharedWithKDFOpt8(alicePrivate.PrivateKey, bobPublic.PublicKey, 32, sha256.New)
	}
}

func BenchmarkECDHComputeSharedWithKDFOpt9(b *testing.B) {
	alicePrivate, _ := GenerateECDHP256Key()
	bobPublic, _ := GenerateECDHP256Key()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ECDHComputeSharedWithKDFOpt9(alicePrivate.PrivateKey, bobPublic.PublicKey, 32)
	}
}

func BenchmarkECDHComputeSharedWithKDFOpt10(b *testing.B) {
	alicePrivate, _ := GenerateECDHP256Key()
	bobPublic, _ := GenerateECDHP256Key()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ECDHComputeSharedWithKDFOpt10(alicePrivate.PrivateKey, bobPublic.PublicKey, 32)
	}
}

func BenchmarkECDSASignOriginalAllocs(b *testing.B) {
	privateKey, _ := GenerateECDSAP256Key()
	data := []byte("test data")

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = ECDSASign(privateKey.PrivateKey, data, sha256.New)
	}
}

func BenchmarkECDSASignOpt1Allocs(b *testing.B) {
	privateKey, _ := GenerateECDSAP256Key()
	data := []byte("test data")

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = ECDSASignOpt1(privateKey.PrivateKey, data)
	}
}

func BenchmarkECDHComputeSharedOriginalAllocs(b *testing.B) {
	alicePrivate, _ := GenerateECDHP256Key()
	bobPublic, _ := GenerateECDHP256Key()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ECDHComputeShared(alicePrivate.PrivateKey, bobPublic.PublicKey)
	}
}

func BenchmarkECDHComputeSharedOpt1Allocs(b *testing.B) {
	alicePrivate, _ := GenerateECDHP256Key()
	bobPublic, _ := GenerateECDHP256Key()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ECDHComputeSharedOpt1(alicePrivate.PrivateKey, bobPublic.PublicKey)
	}
}

func BenchmarkECDHComputeSharedWithKDFOriginalAllocs(b *testing.B) {
	alicePrivate, _ := GenerateECDHP256Key()
	bobPublic, _ := GenerateECDHP256Key()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ECDHComputeSharedWithKDF(alicePrivate.PrivateKey, bobPublic.PublicKey, 32, sha256.New)
	}
}

func BenchmarkECDHComputeSharedWithKDFOpt7Allocs(b *testing.B) {
	alicePrivate, _ := GenerateECDHP256Key()
	bobPublic, _ := GenerateECDHP256Key()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ECDHComputeSharedWithKDFOpt7(alicePrivate.PrivateKey, bobPublic.PublicKey, 32)
	}
}

func BenchmarkECDSASignOpt1Parallel(b *testing.B) {
	privateKey, _ := GenerateECDSAP256Key()
	data := []byte("test data")

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _, _ = ECDSASignOpt1(privateKey.PrivateKey, data)
		}
	})
}

func BenchmarkECDHComputeSharedWithKDFOpt7Parallel(b *testing.B) {
	alicePrivate, _ := GenerateECDHP256Key()
	bobPublic, _ := GenerateECDHP256Key()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = ECDHComputeSharedWithKDFOpt7(alicePrivate.PrivateKey, bobPublic.PublicKey, 32)
		}
	})
}
