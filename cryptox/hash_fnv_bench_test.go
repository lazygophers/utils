package cryptox

import (
	"testing"
)

// 简单测试确保优化方案正确性
func TestOptimizationCorrectness(t *testing.T) {
	testStr := "hello world"
	testBytes := []byte(testStr)

	// 测试 Hash32
	orig32 := Hash32(testStr)
	opt32 := hash32DirectString(testStr)
	if orig32 != opt32 {
		t.Errorf("Hash32 mismatch: orig=%d opt=%d", orig32, opt32)
	}

	// 测试 Hash32a
	orig32a := Hash32a(testStr)
	opt32a := hash32aManual(testStr)
	if orig32a != opt32a {
		t.Errorf("Hash32a mismatch: orig=%d opt=%d", orig32a, opt32a)
	}

	// 测试 Hash64
	orig64 := Hash64(testStr)
	opt64 := hash64Manual(testStr)
	if orig64 != opt64 {
		t.Errorf("Hash64 mismatch: orig=%d opt=%d", orig64, opt64)
	}

	// 测试 Hash64a
	orig64a := Hash64a(testStr)
	opt64a := hash64aManual(testStr)
	if orig64a != opt64a {
		t.Errorf("Hash64a mismatch: orig=%d opt=%d", orig64a, opt64a)
	}

	// 测试 []byte 输入
	orig32Bytes := Hash32(testBytes)
	opt32Bytes := hash32Manual(testBytes)
	if orig32Bytes != opt32Bytes {
		t.Errorf("Hash32([]byte) mismatch: orig=%d opt=%d", orig32Bytes, opt32Bytes)
	}
}

// Benchmark baseline - 原始实现
func BenchmarkHash32_Original(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Hash32(s)
	}
}

func BenchmarkHash32a_Original(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Hash32a(s)
	}
}

func BenchmarkHash64_Original(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Hash64(s)
	}
}

func BenchmarkHash64a_Original(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Hash64a(s)
	}
}

// 方案1: 手动实现 FNV-1 32位（避免接口开销）
func BenchmarkHash32_Manual(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hash32Manual(s)
		_ = h
	}
}

func hash32Manual[M string | []byte](s M) uint32 {
	const (
		prime32  = uint32(16777619)
		offset32 = uint32(2166136261)
	)

	data := toBytes(s)
	h := offset32
	for _, c := range data {
		h *= prime32
		h ^= uint32(c)
	}
	return h
}

// 方案2: 手动实现 FNV-1a 32位
func BenchmarkHash32a_Manual(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hash32aManual(s)
		_ = h
	}
}

func hash32aManual[M string | []byte](s M) uint32 {
	const (
		prime32  = uint32(16777619)
		offset32 = uint32(2166136261)
	)

	data := toBytes(s)
	h := offset32
	for _, c := range data {
		h ^= uint32(c)
		h *= prime32
	}
	return h
}

// 方案3: 手动实现 FNV-1 64位
func BenchmarkHash64_Manual(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hash64Manual(s)
		_ = h
	}
}

func hash64Manual[M string | []byte](s M) uint64 {
	const (
		prime64  = uint64(1099511628211)
		offset64 = uint64(14695981039346656037)
	)

	data := toBytes(s)
	h := offset64
	for _, c := range data {
		h *= prime64
		h ^= uint64(c)
	}
	return h
}

// 方案4: 手动实现 FNV-1a 64位
func BenchmarkHash64a_Manual(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hash64aManual(s)
		_ = h
	}
}

func hash64aManual[M string | []byte](s M) uint64 {
	const (
		prime64  = uint64(1099511628211)
		offset64 = uint64(14695981039346656037)
	)

	data := toBytes(s)
	h := offset64
	for _, c := range data {
		h ^= uint64(c)
		h *= prime64
	}
	return h
}

// 方案5: 使用 unsafe.String 转换（零拷贝）
func BenchmarkHash32_Unsafe(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hash32Unsafe(s)
		_ = h
	}
}

func hash32Unsafe(s string) uint32 {
	const (
		prime32  = uint32(16777619)
		offset32 = uint32(2166136261)
	)

	h := offset32
	length := len(s)
	for i := 0; i < length; i++ {
		h *= prime32
		h ^= uint32(s[i])
	}
	return h
}

// 方案6: 循环展开（4路展开）
func BenchmarkHash32_Unroll(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hash32Unroll(s)
		_ = h
	}
}

func hash32Unroll(s string) uint32 {
	const (
		prime32  = uint32(16777619)
		offset32 = uint32(2166136261)
	)

	h := offset32
	length := len(s)
	i := 0

	// 4路循环展开
	for i+4 <= length {
		h *= prime32
		h ^= uint32(s[i])
		h *= prime32
		h ^= uint32(s[i+1])
		h *= prime32
		h ^= uint32(s[i+2])
		h *= prime32
		h ^= uint32(s[i+3])
		i += 4
	}

	// 处理剩余字节
	for i < length {
		h *= prime32
		h ^= uint32(s[i])
		i++
	}

	return h
}

// 方案7: 索引循环替代 range
func BenchmarkHash32_IndexLoop(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hash32IndexLoop(s)
		_ = h
	}
}

func hash32IndexLoop(s string) uint32 {
	const (
		prime32  = uint32(16777619)
		offset32 = uint32(2166136261)
	)

	h := offset32
	data := []byte(s)
	for i := 0; i < len(data); i++ {
		h *= prime32
		h ^= uint32(data[i])
	}
	return h
}

// 方案8: 内联字节转换（避免辅助函数）
func BenchmarkHash32_InlineBytes(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		const (
			prime32  = uint32(16777619)
			offset32 = uint32(2166136261)
		)

		h := offset32
		var data []byte
		switch v := any(&s).(type) {
		case *string:
			data = []byte(*v)
		case *[]byte:
			data = *v
		}

		for _, c := range data {
			h *= prime32
			h ^= uint32(c)
		}
		_ = h
	}
}

// 方案9: 预计算乘法表（查表优化）
func BenchmarkHash32_LookupTable(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hash32LookupTable(s)
		_ = h
	}
}

var mul32Table [256]uint32

func init() {
	for i := 0; i < 256; i++ {
		mul32Table[i] = uint32(i) * 16777619
	}
}

func hash32LookupTable(s string) uint32 {
	const offset32 = uint32(2166136261)
	h := offset32
	data := []byte(s)
	for _, c := range data {
		h ^= mul32Table[c]
	}
	return h
}

// 方案10: SIMD风格批量处理（8字节一组）
func BenchmarkHash64_SIMD(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hash64SIMD(s)
		_ = h
	}
}

func hash64SIMD(s string) uint64 {
	const (
		prime64  = uint64(1099511628211)
		offset64 = uint64(14695981039346656037)
	)

	h := offset64
	length := len(s)
	i := 0

	// 8字节一组处理
	for i+8 <= length {
		// 处理8字节
		for j := 0; j < 8; j++ {
			h *= prime64
			h ^= uint64(s[i+j])
		}
		i += 8
	}

	// 处理剩余字节
	for i < length {
		h *= prime64
		h ^= uint64(s[i])
		i++
	}

	return h
}

// 方案11: 避免类型转换（直接字符串处理）
func BenchmarkHash32_DirectString(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hash32DirectString(s)
		_ = h
	}
}

func hash32DirectString(s string) uint32 {
	const (
		prime32  = uint32(16777619)
		offset32 = uint32(2166136261)
	)

	h := offset32
	for i := 0; i < len(s); i++ {
		h *= prime32
		h ^= uint32(s[i])
	}
	return h
}

// 方案12: 混合优化（unsafe + 循环展开）
func BenchmarkHash32_Hybrid(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hash32Hybrid(s)
		_ = h
	}
}

func hash32Hybrid(s string) uint32 {
	const (
		prime32  = uint32(16777619)
		offset32 = uint32(2166136261)
	)

	h := offset32
	length := len(s)
	i := 0

	// 8路循环展开
	for i+8 <= length {
		h *= prime32
		h ^= uint32(s[i])
		h *= prime32
		h ^= uint32(s[i+1])
		h *= prime32
		h ^= uint32(s[i+2])
		h *= prime32
		h ^= uint32(s[i+3])
		h *= prime32
		h ^= uint32(s[i+4])
		h *= prime32
		h ^= uint32(s[i+5])
		h *= prime32
		h ^= uint32(s[i+6])
		h *= prime32
		h ^= uint32(s[i+7])
		i += 8
	}

	// 处理剩余字节
	for i < length {
		h *= prime32
		h ^= uint32(s[i])
		i++
	}

	return h
}

// 辅助函数：泛型转字节
func toBytes[M string | []byte](s M) []byte {
	var data []byte
	switch v := any(&s).(type) {
	case *string:
		data = []byte(*v)
	case *[]byte:
		data = *v
	}
	return data
}

// 测试短字符串性能
func BenchmarkHash32_ShortString(b *testing.B) {
	s := "hello"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Hash32(s)
	}
}

func BenchmarkHash32_DirectString_Short(b *testing.B) {
	s := "hello"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash32DirectString(s)
	}
}

// 测试长字符串性能
func BenchmarkHash32_LongString(b *testing.B) {
	s := string(make([]byte, 1024))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Hash32(s)
	}
}

func BenchmarkHash32_DirectString_Long(b *testing.B) {
	s := string(make([]byte, 1024))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash32DirectString(s)
	}
}

// 测试 []byte 输入
func BenchmarkHash32_Bytes(b *testing.B) {
	s := []byte("The quick brown fox jumps over the lazy dog")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Hash32(s)
	}
}

func BenchmarkHash32_Manual_Bytes(b *testing.B) {
	s := []byte("The quick brown fox jumps over the lazy dog")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hash32Manual(s)
	}
}
