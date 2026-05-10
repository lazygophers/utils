package cryptox

import (
	"crypto/sha256"
	"testing"
)

// ============ ECDSA Sign Benchmark (10+ 方案) ============

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

// ============ ECDSA Verify Benchmark (10+ 方案) ============

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

// ============ ECDH ComputeShared Benchmark (10+ 方案) ============

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

// ============ ECDH ComputeSharedWithKDF Benchmark (10+ 方案) ============

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

// ============ 内存分配对比 ============

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

// ============ 并发安全性测试 ============

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

// ============ 功能正确性验证 ============

func TestECDSAOptimizationCorrectness(t *testing.T) {
	privateKey, _ := GenerateECDSAP256Key()
	data := []byte("test data for correctness check")

	// 原始实现
	r1, s1, _ := ECDSASign(privateKey.PrivateKey, data, sha256.New)
	valid1 := ECDSAVerify(privateKey.PublicKey, data, r1, s1, sha256.New)

	// 优化方案 1
	r2, s2, _ := ECDSASignOpt1(privateKey.PrivateKey, data)
	valid2 := ECDSAVerifyOpt6(privateKey.PublicKey, data, r2, s2)

	if !valid1 || !valid2 {
		t.Error("签名验证失败")
	}

	// ECDSA 签名是非确定性的，所以不比较签名值本身
	// 只要验证通过即可
}

func TestECDHOptimizationCorrectness(t *testing.T) {
	alicePrivate, _ := GenerateECDHP256Key()
	bobPublic, _ := GenerateECDHP256Key()

	// 原始实现
	secret1, _ := ECDHComputeShared(alicePrivate.PrivateKey, bobPublic.PublicKey)

	// 优化方案 1
	secret2, _ := ECDHComputeSharedOpt1(alicePrivate.PrivateKey, bobPublic.PublicKey)

	if len(secret1) != len(secret2) {
		t.Errorf("共享密钥长度不一致: %d vs %d", len(secret1), len(secret2))
	}

	for i := range secret1 {
		if secret1[i] != secret2[i] {
			t.Errorf("共享密钥内容不一致在位置 %d: %d vs %d", i, secret1[i], secret2[i])
		}
	}
}

func TestECDHKDFOptimizationCorrectness(t *testing.T) {
	alicePrivate, _ := GenerateECDHP256Key()
	bobPublic, _ := GenerateECDHP256Key()

	// 原始实现
	secret1, _ := ECDHComputeSharedWithKDF(alicePrivate.PrivateKey, bobPublic.PublicKey, 32, sha256.New)

	// 优化方案 7
	secret2, _ := ECDHComputeSharedWithKDFOpt7(alicePrivate.PrivateKey, bobPublic.PublicKey, 32)

	if len(secret1) != len(secret2) {
		t.Errorf("KDF 密钥长度不一致: %d vs %d", len(secret1), len(secret2))
	}

	for i := range secret1 {
		if secret1[i] != secret2[i] {
			t.Errorf("KDF 密钥内容不一致在位置 %d: %d vs %d", i, secret1[i], secret2[i])
		}
	}
}
