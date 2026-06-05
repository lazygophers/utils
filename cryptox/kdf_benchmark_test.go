package cryptox

import (
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"sync"
	"testing"

	"golang.org/x/crypto/pbkdf2"
)

// 基准测试配置
const (
	benchPassword   = "password123"
	benchSalt       = "randomsalt16byte"
	benchIterations = 10000 // 使用较小迭代次数加快测试
	benchKeyLen256  = 32
	benchKeyLen512  = 64
)

// ============================================================================
// 方案 0: 原始实现（基线）
// ============================================================================

func BenchmarkPBKDF2SHA256_Original(b *testing.B) {
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PBKDF2SHA256(password, salt, benchIterations, benchKeyLen256)
	}
}

func BenchmarkPBKDF2SHA512_Original(b *testing.B) {
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = PBKDF2SHA512(password, salt, benchIterations, benchKeyLen512)
	}
}

// ============================================================================
// 方案 1: 直接调用 pbkdf2.Key（减少一层函数调用）
// ============================================================================

func BenchmarkPBKDF2SHA256_Direct(b *testing.B) {
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pbkdf2.Key(password, salt, benchIterations, benchKeyLen256, sha256.New)
	}
}

func BenchmarkPBKDF2SHA512_Direct(b *testing.B) {
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pbkdf2.Key(password, salt, benchIterations, benchKeyLen512, sha512.New)
	}
}

// ============================================================================
// 方案 2: 预构造哈希实例（理论上无效，因为 pbkdf2.Key 内部会重新构造）
// ============================================================================

func BenchmarkPBKDF2SHA256_PreconstructedHash(b *testing.B) {
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	_ = sha256.New() // 预构造（但不会被使用，因为 pbkdf2.Key 需要构造器）
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 这个测试会失败或表现差，因为 pbkdf2.Key 需要构造器，不是实例
		// 但我们测试闭包捕获哈希实例的情况
		_ = pbkdf2.Key(password, salt, benchIterations, benchKeyLen256, func() hash.Hash {
			return sha256.New()
		})
	}
}

// ============================================================================
// 方案 3: 内联哈希函数构造
// ============================================================================

func BenchmarkPBKDF2SHA256_InlineHash(b *testing.B) {
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pbkdf2.Key(password, salt, benchIterations, benchKeyLen256, func() hash.Hash {
			return sha256.New()
		})
	}
}

// ============================================================================
// 方案 4: 使用 sync.Pool 复用 hash.Hash（理论上无效，因为每个 pbkdf2.Key 需要新实例）
// ============================================================================

var sha256Pool = sync.Pool{
	New: func() interface{} {
		return sha256.New()
	},
}

func BenchmarkPBKDF2SHA256_Pool(b *testing.B) {
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pbkdf2.Key(password, salt, benchIterations, benchKeyLen256, func() hash.Hash {
			return sha256Pool.Get().(hash.Hash)
		})
	}
}

// ============================================================================
// 方案 5: 减少闭包分配（使用全局函数）
// ============================================================================

func newSHA256Hash() hash.Hash {
	return sha256.New()
}

func newSHA512Hash() hash.Hash {
	return sha512.New()
}

func BenchmarkPBKDF2SHA256_GlobalFunc(b *testing.B) {
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pbkdf2.Key(password, salt, benchIterations, benchKeyLen256, newSHA256Hash)
	}
}

func BenchmarkPBKDF2SHA512_GlobalFunc(b *testing.B) {
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pbkdf2.Key(password, salt, benchIterations, benchKeyLen512, newSHA512Hash)
	}
}

// ============================================================================
// 方案 6: 测试不同迭代次数的性能比例
// ============================================================================

func BenchmarkPBKDF2SHA256_Iterations1K(b *testing.B) {
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pbkdf2.Key(password, salt, 1000, benchKeyLen256, sha256.New)
	}
}

func BenchmarkPBKDF2SHA256_Iterations10K(b *testing.B) {
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pbkdf2.Key(password, salt, 10000, benchKeyLen256, sha256.New)
	}
}

func BenchmarkPBKDF2SHA256_Iterations100K(b *testing.B) {
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pbkdf2.Key(password, salt, 100000, benchKeyLen256, sha256.New)
	}
}

// ============================================================================
// 方案 7: 测试不同密钥长度的性能影响
// ============================================================================

func BenchmarkPBKDF2SHA256_KeyLen16(b *testing.B) {
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pbkdf2.Key(password, salt, benchIterations, 16, sha256.New)
	}
}

func BenchmarkPBKDF2SHA256_KeyLen32(b *testing.B) {
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pbkdf2.Key(password, salt, benchIterations, 32, sha256.New)
	}
}

func BenchmarkPBKDF2SHA256_KeyLen64(b *testing.B) {
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pbkdf2.Key(password, salt, benchIterations, 64, sha256.New)
	}
}

// ============================================================================
// 方案 8: SHA256 vs SHA512 性能对比
// ============================================================================

func BenchmarkPBKDF2SHA512_vs_SHA256(b *testing.B) {
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pbkdf2.Key(password, salt, benchIterations, benchKeyLen256, sha512.New)
	}
}

// ============================================================================
// 方案 9: 内存分配优化（预分配切片，但 pbkdf2.Key 内部已处理）
// ============================================================================

func BenchmarkPBKDF2SHA256_Preallocate(b *testing.B) {
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// pbkdf2.Key 已经返回新切片，我们无法控制其内部分配
		// 这个测试确认我们的假设
		result := make([]byte, benchKeyLen256)
		_ = pbkdf2.Key(password, salt, benchIterations, benchKeyLen256, sha256.New)
		_ = result
	}
}

// ============================================================================
// 方案 10: 批量处理（不适用于 PBKDF2，但测试确认）
// ============================================================================

func BenchmarkPBKDF2SHA256_Sequential(b *testing.B) {
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 连续调用两次，测试是否有批处理优化空间
		_ = pbkdf2.Key(password, salt, benchIterations, benchKeyLen256, sha256.New)
		_ = pbkdf2.Key(password, salt, benchIterations, benchKeyLen256, sha256.New)
	}
}

// ============================================================================
// 方案 11: 检查编译器内联优化
// ============================================================================

func pbkdf2SHA256Inline(password, salt []byte, iterations, keyLen int) []byte {
	return pbkdf2.Key(password, salt, iterations, keyLen, sha256.New)
}

func BenchmarkPBKDF2SHA256_InlinedFunc(b *testing.B) {
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pbkdf2SHA256Inline(password, salt, benchIterations, benchKeyLen256)
	}
}

// ============================================================================
// 方案 12: 使用 DeriveKey 通用函数的性能开销
// ============================================================================

func BenchmarkDeriveKey_SHA256(b *testing.B) {
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = DeriveKey(password, salt, benchIterations, benchKeyLen256, sha256.New)
	}
}

func BenchmarkDeriveKey_SHA512(b *testing.B) {
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = DeriveKey(password, salt, benchIterations, benchKeyLen512, sha512.New)
	}
}

// ============================================================================
// 方案 13: 测试零分配优化（使用 escape analysis）
// ============================================================================

func BenchmarkPBKDF2SHA256_NoEscape(b *testing.B) {
	// 这个测试验证是否有参数逃逸到堆
	b.ReportAllocs()
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pbkdf2.Key(password, salt, benchIterations, benchKeyLen256, sha256.New)
	}
}

// ============================================================================
// 方案 14: 并行安全测试（确认 PBKDF2 是否可以并行化）
// ============================================================================

func BenchmarkPBKDF2SHA256_Parallel(b *testing.B) {
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = pbkdf2.Key(password, salt, benchIterations, benchKeyLen256, sha256.New)
		}
	})
}

// ============================================================================
// 方案 15: 测试缓存哈希构造器的效果
// ============================================================================

type cachedHashFunc struct {
	hashFunc func() hash.Hash
}

func BenchmarkPBKDF2SHA256_CachedStruct(b *testing.B) {
	password := []byte(benchPassword)
	salt := []byte(benchSalt)
	cached := cachedHashFunc{hashFunc: sha256.New}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = pbkdf2.Key(password, salt, benchIterations, benchKeyLen256, cached.hashFunc)
	}
}
