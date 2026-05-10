package cryptox

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

// 打印性能对比结果
func TestPrintPerformanceResults(t *testing.T) {
	fmt.Println("\n========================================")
	fmt.Println("ECDSA/ECDH 性能优化测试结果")
	fmt.Println("========================================")
	fmt.Println()

	// ECDSA Sign 测试
	privateKey, _ := GenerateECDSAP256Key()
	data := []byte("test data")

	fmt.Println("ECDSA Sign (10,000 次迭代):")
	fmt.Println("-------------------------------------------")

	// 原始版本
	result := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ECDSASign(privateKey.PrivateKey, data, sha256.New)
		}
	})
	fmt.Printf("原始版本:       %10s ns/op, %10s allocs/op, %10s B/op\n",
		fmt.Sprintf("%.0f", float64(result.NsPerOp())),
		fmt.Sprintf("%.0f", float64(result.AllocsPerOp())),
		fmt.Sprintf("%.0f", float64(result.AllocedBytesPerOp())))

	// 优化版本 1
	result = testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ECDSASignOpt1(privateKey.PrivateKey, data)
		}
	})
	fmt.Printf("优化版本 1 (对象池):  %10s ns/op, %10s allocs/op, %10s B/op\n",
		fmt.Sprintf("%.0f", float64(result.NsPerOp())),
		fmt.Sprintf("%.0f", float64(result.AllocsPerOp())),
		fmt.Sprintf("%.0f", float64(result.AllocedBytesPerOp())))

	// ECDH ComputeShared 测试
	alicePrivate, _ := GenerateECDHP256Key()
	bobPublic, _ := GenerateECDHP256Key()

	fmt.Println("\nECDH ComputeShared (10,000 次迭代):")
	fmt.Println("-------------------------------------------")

	// 原始版本
	result = testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ECDHComputeShared(alicePrivate.PrivateKey, bobPublic.PublicKey)
		}
	})
	fmt.Printf("原始版本:       %10s ns/op, %10s allocs/op, %10s B/op\n",
		fmt.Sprintf("%.0f", float64(result.NsPerOp())),
		fmt.Sprintf("%.0f", float64(result.AllocsPerOp())),
		fmt.Sprintf("%.0f", float64(result.AllocedBytesPerOp())))

	// 优化版本 1
	result = testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ECDHComputeSharedOpt1(alicePrivate.PrivateKey, bobPublic.PublicKey)
		}
	})
	fmt.Printf("优化版本 1 (缓存曲线): %10s ns/op, %10s allocs/op, %10s B/op\n",
		fmt.Sprintf("%.0f", float64(result.NsPerOp())),
		fmt.Sprintf("%.0f", float64(result.AllocsPerOp())),
		fmt.Sprintf("%.0f", float64(result.AllocedBytesPerOp())))

	// ECDH ComputeSharedWithKDF 测试
	fmt.Println("\nECDH ComputeSharedWithKDF (5,000 次迭代):")
	fmt.Println("-------------------------------------------")

	// 原始版本
	result = testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ECDHComputeSharedWithKDF(alicePrivate.PrivateKey, bobPublic.PublicKey, 32, sha256.New)
		}
	})
	fmt.Printf("原始版本:       %10s ns/op, %10s allocs/op, %10s B/op\n",
		fmt.Sprintf("%.0f", float64(result.NsPerOp())),
		fmt.Sprintf("%.0f", float64(result.AllocsPerOp())),
		fmt.Sprintf("%.0f", float64(result.AllocedBytesPerOp())))

	// 优化版本 7
	result = testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ECDHComputeSharedWithKDFOpt7(alicePrivate.PrivateKey, bobPublic.PublicKey, 32)
		}
	})
	fmt.Printf("优化版本 7 (对象池): %10s ns/op, %10s allocs/op, %10s B/op\n",
		fmt.Sprintf("%.0f", float64(result.NsPerOp())),
		fmt.Sprintf("%.0f", float64(result.AllocsPerOp())),
		fmt.Sprintf("%.0f", float64(result.AllocedBytesPerOp())))

	fmt.Println("\n========================================")
}
