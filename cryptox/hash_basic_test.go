package cryptox

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
	"unsafe"
)

// Test vectors for hash functions
const (
	testEmpty   = ""
	testABC     = "abc"
	testQuick   = "The quick brown fox jumps over the lazy dog"
	testUnicode = "Hello世界🌍"
)

var (
	testBinary = []byte{0x00, 0x01, 0x02, 0x03, 0xFF, 0xFE, 0xFD}
	testLong   = strings.Repeat("a", 10000)
)

// MD5 test vectors (from RFC 1321)
func TestMd5(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", testEmpty, "d41d8cd98f00b204e9800998ecf8427e"},
		{"abc", testABC, "900150983cd24fb0d6963f7d28e17f72"},
		{"quick fox", testQuick, "9e107d9d372bb6826bd81d3542a419d6"},
		{"unicode", testUnicode, "9af5f6e9dc774fe9cf02e5981ff3fa1a"},
		{"single char", "a", "0cc175b9c0f1b6a831c399e269772661"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Md5(tt.input)
			if result != tt.expected {
				t.Errorf("Md5(%q) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMd5Bytes(t *testing.T) {
	result := Md5([]byte("abc"))
	expected := "900150983cd24fb0d6963f7d28e17f72"
	if result != expected {
		t.Errorf("Md5([]byte) = %s, want %s", result, expected)
	}

	result = Md5(testBinary)
	if len(result) != 32 {
		t.Errorf("Md5(binary) returned invalid hex length: %d", len(result))
	}
}

// SHA1 test vectors (from FIPS 180-1)
func TestSHA1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", testEmpty, "da39a3ee5e6b4b0d3255bfef95601890afd80709"},
		{"abc", testABC, "a9993e364706816aba3e25717850c26c9cd0d89d"},
		{"quick fox", testQuick, "2fd4e1c67a2d28fced849ee1bb76e7391b93eb12"},
		{"unicode", testUnicode, "23ed166736a301d418b5d7a402d41b862db8cd27"},
		{"single char", "a", "86f7e437faa5a7fce15d1ddcb9eaeaea377667b8"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SHA1(tt.input)
			if result != tt.expected {
				t.Errorf("SHA1(%q) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSHA1Bytes(t *testing.T) {
	result := SHA1([]byte("abc"))
	expected := "a9993e364706816aba3e25717850c26c9cd0d89d"
	if result != expected {
		t.Errorf("SHA1([]byte) = %s, want %s", result, expected)
	}
}

// SHA-224 test vectors (from FIPS 180-4)
func TestSha224(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", testEmpty, "d14a028c2a3a2bc9476102bb288234c415a2b01f828ea62ac5b3e42f"},
		{"abc", testABC, "23097d223405d8228642a477bda255b32aadbce4bda0b3f7e36c9da7"},
		{"quick fox", testQuick, "730e109bd7a8a32b1cb9d9a09aa2325d2430587ddbc0c38bad911525"},
		{"single char", "a", "abd37534c7d9a2efb9465de931cd7055ffdb8879563ae98078d6d6d5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sha224(tt.input)
			if result != tt.expected {
				t.Errorf("Sha224(%q) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSha224Bytes(t *testing.T) {
	result := Sha224([]byte("abc"))
	expected := "23097d223405d8228642a477bda255b32aadbce4bda0b3f7e36c9da7"
	if result != expected {
		t.Errorf("Sha224([]byte) = %s, want %s", result, expected)
	}
}

// SHA-256 test vectors (from FIPS 180-4)
func TestSha256(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", testEmpty, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"abc", testABC, "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"},
		{"quick fox", testQuick, "d7a8fbb307d7809469ca9abcb0082e4f8d5651e46d3cdb762d02d0bf37c9e592"},
		{"single char", "a", "ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sha256(tt.input)
			if result != tt.expected {
				t.Errorf("Sha256(%q) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSha256Bytes(t *testing.T) {
	result := Sha256([]byte("abc"))
	expected := "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"
	if result != expected {
		t.Errorf("Sha256([]byte) = %s, want %s", result, expected)
	}
}

// SHA-384 test vectors (from FIPS 180-4)
func TestSha384(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", testEmpty, "38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b"},
		{"abc", testABC, "cb00753f45a35e8bb5a03d699ac65007272c32ab0eded1631a8b605a43ff5bed8086072ba1e7cc2358baeca134c825a7"},
		{"quick fox", testQuick, "ca737f1014a48f4c0b6dd43cb177b0afd9e5169367544c494011e3317dbf9a509cb1e5dc1e85a941bbee3d7f2afbc9b1"},
		{"single char", "a", "54a59b9f22b0b80880d8427e548b7c23abd873486e1f035dce9cd697e85175033caa88e6d57bc35efae0b5afd3145f31"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sha384(tt.input)
			if result != tt.expected {
				t.Errorf("Sha384(%q) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSha384Bytes(t *testing.T) {
	result := Sha384([]byte("abc"))
	expected := "cb00753f45a35e8bb5a03d699ac65007272c32ab0eded1631a8b605a43ff5bed8086072ba1e7cc2358baeca134c825a7"
	if result != expected {
		t.Errorf("Sha384([]byte) = %s, want %s", result, expected)
	}
}

// SHA-512 test vectors (from FIPS 180-4)
func TestSha512(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", testEmpty, "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"},
		{"abc", testABC, "ddaf35a193617abacc417349ae20413112e6fa4e89a97ea20a9eeee64b55d39a2192992a274fc1a836ba3c23a3feebbd454d4423643ce80e2a9ac94fa54ca49f"},
		{"quick fox", testQuick, "07e547d9586f6a73f73fbac0435ed76951218fb7d0c8d788a309d785436bbb642e93a252a954f23912547d1e8a3b5ed6e1bfd7097821233fa0538f3db854fee6"},
		{"single char", "a", "1f40fc92da241694750979ee6cf582f2d5d7d28e18335de05abc54d0560e0f5302860c652bf08d560252aa5e74210546f369fbbbce8c12cfc7957b2652fe9a75"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sha512(tt.input)
			if result != tt.expected {
				t.Errorf("Sha512(%q) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSha512Bytes(t *testing.T) {
	result := Sha512([]byte("abc"))
	expected := "ddaf35a193617abacc417349ae20413112e6fa4e89a97ea20a9eeee64b55d39a2192992a274fc1a836ba3c23a3feebbd454d4423643ce80e2a9ac94fa54ca49f"
	if result != expected {
		t.Errorf("Sha512([]byte) = %s, want %s", result, expected)
	}
}

// SHA-512/224 test vectors
func TestSha512_224(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", testEmpty, "6ed0dd02806fa89e25de060c19d3ac86cabb87d6a0ddd05c333b84f4"},
		{"abc", testABC, "4634270f707b6a54daae7530460842e20e37ed265ceee9a43e8924aa"},
		{"quick fox", testQuick, "944cd2847fb54558d4775db0485a50003111c8e5daa63fe722c6aa37"},
		{"single char", "a", "d5cdb9ccc769a5121d4175f2bfdd13d6310e0d3d361ea75d82108327"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sha512_224(tt.input)
			if result != tt.expected {
				t.Errorf("Sha512_224(%q) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSha512_224Bytes(t *testing.T) {
	result := Sha512_224([]byte("abc"))
	expected := "4634270f707b6a54daae7530460842e20e37ed265ceee9a43e8924aa"
	if result != expected {
		t.Errorf("Sha512_224([]byte) = %s, want %s", result, expected)
	}
}

// SHA-512/256 test vectors
func TestSha512_256(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", testEmpty, "c672b8d1ef56ed28ab87c3622c5114069bdd3ad7b8f9737498d0c01ecef0967a"},
		{"abc", testABC, "53048e2681941ef99b2e29b76b4c7dabe4c2d0c634fc6d46e0e2f13107e7af23"},
		{"quick fox", testQuick, "dd9d67b371519c339ed8dbd25af90e976a1eeefd4ad3d889005e532fc5bef04d"},
		{"single char", "a", "455e518824bc0601f9fb858ff5c37d417d67c2f8e0df2babe4808858aea830f8"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sha512_256(tt.input)
			if result != tt.expected {
				t.Errorf("Sha512_256(%q) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSha512_256Bytes(t *testing.T) {
	result := Sha512_256([]byte("abc"))
	expected := "53048e2681941ef99b2e29b76b4c7dabe4c2d0c634fc6d46e0e2f13107e7af23"
	if result != expected {
		t.Errorf("Sha512_256([]byte) = %s, want %s", result, expected)
	}
}

// Additional edge case tests
func TestHashFunctionsWithBinaryData(t *testing.T) {
	testData := testBinary

	md5Result := Md5(testData)
	if len(md5Result) != 32 {
		t.Errorf("Md5(binary) returned invalid hex length: %d", len(md5Result))
	}

	sha1Result := SHA1(testData)
	if len(sha1Result) != 40 {
		t.Errorf("SHA1(binary) returned invalid hex length: %d", len(sha1Result))
	}

	sha224Result := Sha224(testData)
	if len(sha224Result) != 56 {
		t.Errorf("Sha224(binary) returned invalid hex length: %d", len(sha224Result))
	}

	sha256Result := Sha256(testData)
	if len(sha256Result) != 64 {
		t.Errorf("Sha256(binary) returned invalid hex length: %d", len(sha256Result))
	}

	sha384Result := Sha384(testData)
	if len(sha384Result) != 96 {
		t.Errorf("Sha384(binary) returned invalid hex length: %d", len(sha384Result))
	}

	sha512Result := Sha512(testData)
	if len(sha512Result) != 128 {
		t.Errorf("Sha512(binary) returned invalid hex length: %d", len(sha512Result))
	}

	sha512_224Result := Sha512_224(testData)
	if len(sha512_224Result) != 56 {
		t.Errorf("Sha512_224(binary) returned invalid hex length: %d", len(sha512_224Result))
	}

	sha512_256Result := Sha512_256(testData)
	if len(sha512_256Result) != 64 {
		t.Errorf("Sha512_256(binary) returned invalid hex length: %d", len(sha512_256Result))
	}
}

func TestHashFunctionsWithLongInput(t *testing.T) {
	longData := testLong

	md5Result := Md5(longData)
	if len(md5Result) != 32 {
		t.Errorf("Md5(long) returned invalid hex length: %d", len(md5Result))
	}

	sha256Result := Sha256(longData)
	if len(sha256Result) != 64 {
		t.Errorf("Sha256(long) returned invalid hex length: %d", len(sha256Result))
	}

	sha512Result := Sha512(longData)
	if len(sha512Result) != 128 {
		t.Errorf("Sha512(long) returned invalid hex length: %d", len(sha512Result))
	}
}

// Benchmarks
func BenchmarkMd5(b *testing.B) {
	data := []byte(testQuick)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Md5(data)
	}
}

func BenchmarkSHA1(b *testing.B) {
	data := []byte(testQuick)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SHA1(data)
	}
}

func BenchmarkSha224(b *testing.B) {
	data := []byte(testQuick)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Sha224(data)
	}
}

func BenchmarkSha256(b *testing.B) {
	data := []byte(testQuick)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Sha256(data)
	}
}

func BenchmarkSha384(b *testing.B) {
	data := []byte(testQuick)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Sha384(data)
	}
}

func BenchmarkSha512(b *testing.B) {
	data := []byte(testQuick)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Sha512(data)
	}
}

func BenchmarkSha512_224(b *testing.B) {
	data := []byte(testQuick)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Sha512_224(data)
	}
}

func BenchmarkSha512_256(b *testing.B) {
	data := []byte(testQuick)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Sha512_256(data)
	}
}

// Benchmark with different input sizes
func BenchmarkMd5_Short(b *testing.B) {
	data := []byte("a")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Md5(data)
	}
}

func BenchmarkMd5_Long(b *testing.B) {
	data := []byte(testLong)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Md5(data)
	}
}

func BenchmarkSha256_Short(b *testing.B) {
	data := []byte("a")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Sha256(data)
	}
}

func BenchmarkSha256_Long(b *testing.B) {
	data := []byte(testLong)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Sha256(data)
	}
}

// ============================================================
// HMACSHA1 Benchmark Tests
// ============================================================

func BenchmarkHMACSHA1_Optimized(b *testing.B) {
	key := "test_key"
	msg := "hello"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = HMACSHA1(key, msg)
	}
}

func BenchmarkHMACSHA1_Bytes(b *testing.B) {
	key := []byte("test_key")
	msg := []byte("hello")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = HMACSHA1(key, msg)
	}
}

// TestMd5_Coverage_String 测试 string 输入
func TestMd5_Coverage_String(t *testing.T) {
	input := "hello world"
	result := Md5(input)
	if len(result) != 32 {
		t.Errorf("MD5 hash length should be 32, got %d", len(result))
	}
	// 验证结果一致性
	expected := md5.Sum([]byte(input))
	expectedHex := encodeHex(expected[:])
	if result != expectedHex {
		t.Errorf("MD5 hash mismatch: got %s, want %s", result, expectedHex)
	}
}

// TestMd5_Coverage_Bytes 测试 []byte 输入
func TestMd5_Coverage_Bytes(t *testing.T) {
	input := []byte("hello world")
	result := Md5(input)
	if len(result) != 32 {
		t.Errorf("MD5 hash length should be 32, got %d", len(result))
	}
	// 验证结果一致性
	expected := md5.Sum(input)
	expectedHex := encodeHex(expected[:])
	if result != expectedHex {
		t.Errorf("MD5 hash mismatch: got %s, want %s", result, expectedHex)
	}
}

// TestMd5_Coverage_Empty 测试空输入
func TestMd5_Coverage_Empty(t *testing.T) {
	result := Md5("")
	if len(result) != 32 {
		t.Errorf("MD5 hash length should be 32, got %d", len(result))
	}
	// 验证标准空字符串 MD5
	expected := "d41d8cd98f00b204e9800998ecf8427e"
	if result != expected {
		t.Errorf("Empty MD5 hash: got %s, want %s", result, expected)
	}
}

// TestMd5_Coverage_Large 测试大数据
func TestMd5_Coverage_Large(t *testing.T) {
	input := make([]byte, 1024*1024)
	for i := range input {
		input[i] = byte(i % 256)
	}
	result := Md5(input)
	if len(result) != 32 {
		t.Errorf("MD5 hash length should be 32, got %d", len(result))
	}
}

// TestMd5_Coverage_SpecialChars 测试特殊字符
func TestMd5_Coverage_SpecialChars(t *testing.T) {
	inputs := []string{
		"!@#$%^&*()",
		"中文测试",
		"\n\r\t",
		"🔥🚀",
	}
	for _, input := range inputs {
		result := Md5(input)
		if len(result) != 32 {
			t.Errorf("MD5 hash length should be 32 for %q, got %d", input, len(result))
		}
	}
}

// TestMd5_Coverage_Consistency 测试结果一致性
func TestMd5_Coverage_Consistency(t *testing.T) {
	input := "consistency test"
	r1 := Md5(input)
	r2 := Md5(input)
	r3 := Md5([]byte(input))
	if r1 != r2 || r2 != r3 {
		t.Errorf("MD5 results inconsistent: %s, %s, %s", r1, r2, r3)
	}
}

// TestMd5_Coverage_DifferentInputs 测试不同输入产生不同结果
func TestMd5_Coverage_DifferentInputs(t *testing.T) {
	inputs := []string{"a", "b", "c"}
	hashes := make(map[string]bool)
	for _, input := range inputs {
		hash := Md5(input)
		if hashes[hash] {
			t.Errorf("Duplicate hash for different input: %s", hash)
		}
		hashes[hash] = true
	}
}

// TestMd5_Coverage_BinaryData 测试二进制数据
func TestMd5_Coverage_BinaryData(t *testing.T) {
	input := []byte{0x00, 0x01, 0x02, 0xFF, 0xFE, 0xFD}
	result := Md5(input)
	if len(result) != 32 {
		t.Errorf("MD5 hash length should be 32, got %d", len(result))
	}
}

// TestMd5_Coverage_Concurrent 测试并发安全
func TestMd5_Coverage_Concurrent(t *testing.T) {
	input := "concurrent test"
	results := make(chan string, 100)
	for i := 0; i < 100; i++ {
		go func() {
			results <- Md5(input)
		}()
	}
	expected := Md5(input)
	for i := 0; i < 100; i++ {
		if result := <-results; result != expected {
			t.Errorf("Concurrent MD5 mismatch: got %s, want %s", result, expected)
		}
	}
}

// TestMd5_Coverage_KnownValues 测试已知值
func TestMd5_Coverage_KnownValues(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", "d41d8cd98f00b204e9800998ecf8427e"},
		{"a", "0cc175b9c0f1b6a831c399e269772661"},
		{"abc", "900150983cd24fb0d6963f7d28e17f72"},
		{"message digest", "f96b697d7cb7938d525a2f31aaf161d0"},
		{"abcdefghijklmnopqrstuvwxyz", "c3fcd3d76192e4007dfb496cca67e13b"},
	}
	for _, tc := range testCases {
		result := Md5(tc.input)
		if result != tc.expected {
			t.Errorf("MD5(%q) = %s, want %s", tc.input, result, tc.expected)
		}
	}
}

// encodeHex 辅助函数：将字节转换为 hex 字符串
func encodeHex(data []byte) string {
	const hexTable = "0123456789abcdef"
	result := make([]byte, len(data)*2)
	for i, b := range data {
		result[i*2] = hexTable[b>>4]
		result[i*2+1] = hexTable[b&0x0f]
	}
	return string(result)
}

// Sha256Original 原始实现（使用 fmt.Sprintf）
func Sha256Original[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

// Sha256V1 手动 hex 编码（参考 MD5/SHA1 实现）
func Sha256V1[M string | []byte](s M) string {
	hash := sha256.Sum256([]byte(s))
	var result [64]byte
	for i := 0; i < 32; i++ {
		b := hash[i]
		result[i*2] = "0123456789abcdef"[b>>4]
		result[i*2+1] = "0123456789abcdef"[b&0x0f]
	}
	return string(result[:])
}

// Sha256V2 使用 encoding/hex.EncodeToString
func Sha256V2[M string | []byte](s M) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}

// Sha256V3 预分配 hex 字符串常量
func Sha256V3[M string | []byte](s M) string {
	const hexChars = "0123456789abcdef"
	hash := sha256.Sum256([]byte(s))
	var result [64]byte
	for i := 0; i < 32; i++ {
		b := hash[i]
		result[i*2] = hexChars[b>>4]
		result[i*2+1] = hexChars[b&0x0f]
	}
	return string(result[:])
}

// Sha256V4 循环展开 4 次
func Sha256V4[M string | []byte](s M) string {
	const hexChars = "0123456789abcdef"
	hash := sha256.Sum256([]byte(s))
	var result [64]byte

	for i := 0; i < 32; i += 4 {
		b1 := hash[i]
		result[i*2] = hexChars[b1>>4]
		result[i*2+1] = hexChars[b1&0x0f]

		if i+1 < 32 {
			b2 := hash[i+1]
			result[(i+1)*2] = hexChars[b2>>4]
			result[(i+1)*2+1] = hexChars[b2&0x0f]
		}

		if i+2 < 32 {
			b3 := hash[i+2]
			result[(i+2)*2] = hexChars[b3>>4]
			result[(i+2)*2+1] = hexChars[b3&0x0f]
		}

		if i+3 < 32 {
			b4 := hash[i+3]
			result[(i+3)*2] = hexChars[b4>>4]
			result[(i+3)*2+1] = hexChars[b4&0x0f]
		}
	}
	return string(result[:])
}

// Sha256V5 循环展开 8 次
func Sha256V5[M string | []byte](s M) string {
	const hexChars = "0123456789abcdef"
	hash := sha256.Sum256([]byte(s))
	var result [64]byte

	for i := 0; i < 32; i += 8 {
		b1 := hash[i]
		result[i*2] = hexChars[b1>>4]
		result[i*2+1] = hexChars[b1&0x0f]

		b2 := hash[i+1]
		result[(i+1)*2] = hexChars[b2>>4]
		result[(i+1)*2+1] = hexChars[b2&0x0f]

		b3 := hash[i+2]
		result[(i+2)*2] = hexChars[b3>>4]
		result[(i+2)*2+1] = hexChars[b3&0x0f]

		b4 := hash[i+3]
		result[(i+3)*2] = hexChars[b4>>4]
		result[(i+3)*2+1] = hexChars[b4&0x0f]

		b5 := hash[i+4]
		result[(i+4)*2] = hexChars[b5>>4]
		result[(i+4)*2+1] = hexChars[b5&0x0f]

		b6 := hash[i+5]
		result[(i+5)*2] = hexChars[b6>>4]
		result[(i+5)*2+1] = hexChars[b6&0x0f]

		b7 := hash[i+6]
		result[(i+6)*2] = hexChars[b7>>4]
		result[(i+6)*2+1] = hexChars[b7&0x0f]

		b8 := hash[i+7]
		result[(i+7)*2] = hexChars[b8>>4]
		result[(i+7)*2+1] = hexChars[b8&0x0f]
	}
	return string(result[:])
}

// Sha256V6 使用 unsafe 转换（避免边界检查）
func Sha256V6[M string | []byte](s M) string {
	const hexChars = "0123456789abcdef"
	hash := sha256.Sum256([]byte(s))
	var result [64]byte

	// 使用 unsafe 避免边界检查
	src := unsafe.Slice(&hash[0], 32)
	dst := unsafe.Slice(&result[0], 64)

	for i := 0; i < 32; i++ {
		b := src[i]
		dst[i*2] = hexChars[b>>4]
		dst[i*2+1] = hexChars[b&0x0f]
	}
	return string(result[:])
}

// Sha256V7 查表优化（使用 16 字节查找表）
func Sha256V7[M string | []byte](s M) string {
	var hexTable = [16]string{
		"0", "1", "2", "3", "4", "5", "6", "7",
		"8", "9", "a", "b", "c", "d", "e", "f",
	}

	hash := sha256.Sum256([]byte(s))
	result := make([]byte, 0, 64)

	for i := 0; i < 32; i++ {
		b := hash[i]
		result = append(result, hexTable[b>>4]...)
		result = append(result, hexTable[b&0x0f]...)
	}
	return string(result)
}

// Sha256V8 查表优化（使用 512 字节全局查找表）
func Sha256V8[M string | []byte](s M) string {
	hash := sha256.Sum256([]byte(s))
	var result [64]byte

	for i := 0; i < 32; i++ {
		b := hash[i]
		// 直接计算 hex 字符，避免查表开销
		result[i*2] = "0123456789abcdef"[b>>4]
		result[i*2+1] = "0123456789abcdef"[b&0x0f]
	}
	return string(result[:])
}

// Sha256V9 使用 16 位查找表（一次处理一个字节）
func Sha256V9[M string | []byte](s M) string {
	// 预生成 16 位查找表，每个元素是 2 字节
	var hexTable [256][2]byte
	for i := 0; i < 256; i++ {
		hexTable[i][0] = "0123456789abcdef"[i>>4]
		hexTable[i][1] = "0123456789abcdef"[i&0x0f]
	}

	hash := sha256.Sum256([]byte(s))
	var result [64]byte

	for i := 0; i < 32; i++ {
		b := hash[i]
		pair := hexTable[b]
		result[i*2] = pair[0]
		result[i*2+1] = pair[1]
	}
	return string(result[:])
}

// Sha256V10 完全展开循环（32 次，无循环）
func Sha256V10[M string | []byte](s M) string {
	const hexChars = "0123456789abcdef"
	hash := sha256.Sum256([]byte(s))
	var result [64]byte

	// 完全展开
	result[0] = hexChars[hash[0]>>4]
	result[1] = hexChars[hash[0]&0x0f]
	result[2] = hexChars[hash[1]>>4]
	result[3] = hexChars[hash[1]&0x0f]
	result[4] = hexChars[hash[2]>>4]
	result[5] = hexChars[hash[2]&0x0f]
	result[6] = hexChars[hash[3]>>4]
	result[7] = hexChars[hash[3]&0x0f]
	result[8] = hexChars[hash[4]>>4]
	result[9] = hexChars[hash[4]&0x0f]
	result[10] = hexChars[hash[5]>>4]
	result[11] = hexChars[hash[5]&0x0f]
	result[12] = hexChars[hash[6]>>4]
	result[13] = hexChars[hash[6]&0x0f]
	result[14] = hexChars[hash[7]>>4]
	result[15] = hexChars[hash[7]&0x0f]
	result[16] = hexChars[hash[8]>>4]
	result[17] = hexChars[hash[8]&0x0f]
	result[18] = hexChars[hash[9]>>4]
	result[19] = hexChars[hash[9]&0x0f]
	result[20] = hexChars[hash[10]>>4]
	result[21] = hexChars[hash[10]&0x0f]
	result[22] = hexChars[hash[11]>>4]
	result[23] = hexChars[hash[11]&0x0f]
	result[24] = hexChars[hash[12]>>4]
	result[25] = hexChars[hash[12]&0x0f]
	result[26] = hexChars[hash[13]>>4]
	result[27] = hexChars[hash[13]&0x0f]
	result[28] = hexChars[hash[14]>>4]
	result[29] = hexChars[hash[14]&0x0f]
	result[30] = hexChars[hash[15]>>4]
	result[31] = hexChars[hash[15]&0x0f]
	result[32] = hexChars[hash[16]>>4]
	result[33] = hexChars[hash[16]&0x0f]
	result[34] = hexChars[hash[17]>>4]
	result[35] = hexChars[hash[17]&0x0f]
	result[36] = hexChars[hash[18]>>4]
	result[37] = hexChars[hash[18]&0x0f]
	result[38] = hexChars[hash[19]>>4]
	result[39] = hexChars[hash[19]&0x0f]
	result[40] = hexChars[hash[20]>>4]
	result[41] = hexChars[hash[20]&0x0f]
	result[42] = hexChars[hash[21]>>4]
	result[43] = hexChars[hash[21]&0x0f]
	result[44] = hexChars[hash[22]>>4]
	result[45] = hexChars[hash[22]&0x0f]
	result[46] = hexChars[hash[23]>>4]
	result[47] = hexChars[hash[23]&0x0f]
	result[48] = hexChars[hash[24]>>4]
	result[49] = hexChars[hash[24]&0x0f]
	result[50] = hexChars[hash[25]>>4]
	result[51] = hexChars[hash[25]&0x0f]
	result[52] = hexChars[hash[26]>>4]
	result[53] = hexChars[hash[26]&0x0f]
	result[54] = hexChars[hash[27]>>4]
	result[55] = hexChars[hash[27]&0x0f]
	result[56] = hexChars[hash[28]>>4]
	result[57] = hexChars[hash[28]&0x0f]
	result[58] = hexChars[hash[29]>>4]
	result[59] = hexChars[hash[29]&0x0f]
	result[60] = hexChars[hash[30]>>4]
	result[61] = hexChars[hash[30]&0x0f]
	result[62] = hexChars[hash[31]>>4]
	result[63] = hexChars[hash[31]&0x0f]

	return string(result[:])
}

// Sha256V11 混合优化：4 次循环展开 + 预分配 hexChars + 内联
func Sha256V11[M string | []byte](s M) string {
	const hexChars = "0123456789abcdef"
	hash := sha256.Sum256([]byte(s))
	var result [64]byte

	// 4 次循环展开，无边界检查（已知 32 字节）
	i := 0
	b1 := hash[i]
	result[i*2] = hexChars[b1>>4]
	result[i*2+1] = hexChars[b1&0x0f]

	b2 := hash[i+1]
	result[(i+1)*2] = hexChars[b2>>4]
	result[(i+1)*2+1] = hexChars[b2&0x0f]

	b3 := hash[i+2]
	result[(i+2)*2] = hexChars[b3>>4]
	result[(i+2)*2+1] = hexChars[b3&0x0f]

	b4 := hash[i+3]
	result[(i+3)*2] = hexChars[b4>>4]
	result[(i+3)*2+1] = hexChars[b4&0x0f]

	i = 4
	b1 = hash[i]
	result[i*2] = hexChars[b1>>4]
	result[i*2+1] = hexChars[b1&0x0f]

	b2 = hash[i+1]
	result[(i+1)*2] = hexChars[b2>>4]
	result[(i+1)*2+1] = hexChars[b2&0x0f]

	b3 = hash[i+2]
	result[(i+2)*2] = hexChars[b3>>4]
	result[(i+2)*2+1] = hexChars[b3&0x0f]

	b4 = hash[i+3]
	result[(i+3)*2] = hexChars[b4>>4]
	result[(i+3)*2+1] = hexChars[b4&0x0f]

	i = 8
	b1 = hash[i]
	result[i*2] = hexChars[b1>>4]
	result[i*2+1] = hexChars[b1&0x0f]

	b2 = hash[i+1]
	result[(i+1)*2] = hexChars[b2>>4]
	result[(i+1)*2+1] = hexChars[b2&0x0f]

	b3 = hash[i+2]
	result[(i+2)*2] = hexChars[b3>>4]
	result[(i+2)*2+1] = hexChars[b3&0x0f]

	b4 = hash[i+3]
	result[(i+3)*2] = hexChars[b4>>4]
	result[(i+3)*2+1] = hexChars[b4&0x0f]

	i = 12
	b1 = hash[i]
	result[i*2] = hexChars[b1>>4]
	result[i*2+1] = hexChars[b1&0x0f]

	b2 = hash[i+1]
	result[(i+1)*2] = hexChars[b2>>4]
	result[(i+1)*2+1] = hexChars[b2&0x0f]

	b3 = hash[i+2]
	result[(i+2)*2] = hexChars[b3>>4]
	result[(i+2)*2+1] = hexChars[b3&0x0f]

	b4 = hash[i+3]
	result[(i+3)*2] = hexChars[b4>>4]
	result[(i+3)*2+1] = hexChars[b4&0x0f]

	i = 16
	b1 = hash[i]
	result[i*2] = hexChars[b1>>4]
	result[i*2+1] = hexChars[b1&0x0f]

	b2 = hash[i+1]
	result[(i+1)*2] = hexChars[b2>>4]
	result[(i+1)*2+1] = hexChars[b2&0x0f]

	b3 = hash[i+2]
	result[(i+2)*2] = hexChars[b3>>4]
	result[(i+2)*2+1] = hexChars[b3&0x0f]

	b4 = hash[i+3]
	result[(i+3)*2] = hexChars[b4>>4]
	result[(i+3)*2+1] = hexChars[b4&0x0f]

	i = 20
	b1 = hash[i]
	result[i*2] = hexChars[b1>>4]
	result[i*2+1] = hexChars[b1&0x0f]

	b2 = hash[i+1]
	result[(i+1)*2] = hexChars[b2>>4]
	result[(i+1)*2+1] = hexChars[b2&0x0f]

	b3 = hash[i+2]
	result[(i+2)*2] = hexChars[b3>>4]
	result[(i+2)*2+1] = hexChars[b3&0x0f]

	b4 = hash[i+3]
	result[(i+3)*2] = hexChars[b4>>4]
	result[(i+3)*2+1] = hexChars[b4&0x0f]

	i = 24
	b1 = hash[i]
	result[i*2] = hexChars[b1>>4]
	result[i*2+1] = hexChars[b1&0x0f]

	b2 = hash[i+1]
	result[(i+1)*2] = hexChars[b2>>4]
	result[(i+1)*2+1] = hexChars[b2&0x0f]

	b3 = hash[i+2]
	result[(i+2)*2] = hexChars[b3>>4]
	result[(i+2)*2+1] = hexChars[b3&0x0f]

	b4 = hash[i+3]
	result[(i+3)*2] = hexChars[b4>>4]
	result[(i+3)*2+1] = hexChars[b4&0x0f]

	i = 28
	b1 = hash[i]
	result[i*2] = hexChars[b1>>4]
	result[i*2+1] = hexChars[b1&0x0f]

	b2 = hash[i+1]
	result[(i+1)*2] = hexChars[b2>>4]
	result[(i+1)*2+1] = hexChars[b2&0x0f]

	b3 = hash[i+2]
	result[(i+2)*2] = hexChars[b3>>4]
	result[(i+2)*2+1] = hexChars[b3&0x0f]

	b4 = hash[i+3]
	result[(i+3)*2] = hexChars[b4>>4]
	result[(i+3)*2+1] = hexChars[b4&0x0f]

	return string(result[:])
}

// Sha256V12 使用 bytes.Builder
func Sha256V12[M string | []byte](s M) string {
	const hexChars = "0123456789abcdef"
	hash := sha256.Sum256([]byte(s))

	var builder [64]byte
	for i := 0; i < 32; i++ {
		b := hash[i]
		builder[i*2] = hexChars[b>>4]
		builder[i*2+1] = hexChars[b&0x0f]
	}
	return string(builder[:])
}

// TestSHA1_Coverage_String 测试 string 输入
func TestSHA1_Coverage_String(t *testing.T) {
	input := "hello world"
	result := SHA1(input)
	if len(result) != 40 {
		t.Errorf("SHA1 hash length should be 40, got %d", len(result))
	}
	// 验证结果一致性
	expected := sha1.Sum([]byte(input))
	expectedHex := encodeHexSHA1(expected[:])
	if result != expectedHex {
		t.Errorf("SHA1 hash mismatch: got %s, want %s", result, expectedHex)
	}
}

// TestSHA1_Coverage_Bytes 测试 []byte 输入
func TestSHA1_Coverage_Bytes(t *testing.T) {
	input := []byte("hello world")
	result := SHA1(input)
	if len(result) != 40 {
		t.Errorf("SHA1 hash length should be 40, got %d", len(result))
	}
}

// TestSHA1_Coverage_Empty 测试空输入
func TestSHA1_Coverage_Empty(t *testing.T) {
	result := SHA1("")
	if len(result) != 40 {
		t.Errorf("SHA1 hash length should be 40, got %d", len(result))
	}
	// 验证标准空字符串 SHA1
	expected := "da39a3ee5e6b4b0d3255bfef95601890afd80709"
	if result != expected {
		t.Errorf("Empty SHA1 hash: got %s, want %s", result, expected)
	}
}

// TestSHA1_Coverage_Large 测试大数据
func TestSHA1_Coverage_Large(t *testing.T) {
	input := make([]byte, 1024*1024)
	for i := range input {
		input[i] = byte(i % 256)
	}
	result := SHA1(input)
	if len(result) != 40 {
		t.Errorf("SHA1 hash length should be 40, got %d", len(result))
	}
}

// TestSHA1_Coverage_SpecialChars 测试特殊字符
func TestSHA1_Coverage_SpecialChars(t *testing.T) {
	inputs := []string{
		"!@#$%^&*()",
		"中文测试",
		"\n\r\t",
		"🔥🚀",
	}
	for _, input := range inputs {
		result := SHA1(input)
		if len(result) != 40 {
			t.Errorf("SHA1 hash length should be 40 for %q, got %d", input, len(result))
		}
	}
}

// TestSHA1_Coverage_Consistency 测试结果一致性
func TestSHA1_Coverage_Consistency(t *testing.T) {
	input := "consistency test"
	r1 := SHA1(input)
	r2 := SHA1(input)
	r3 := SHA1([]byte(input))
	if r1 != r2 || r2 != r3 {
		t.Errorf("SHA1 results inconsistent: %s, %s, %s", r1, r2, r3)
	}
}

// TestSHA1_Coverage_DifferentInputs 测试不同输入产生不同结果
func TestSHA1_Coverage_DifferentInputs(t *testing.T) {
	inputs := []string{"a", "b", "c"}
	hashes := make(map[string]bool)
	for _, input := range inputs {
		hash := SHA1(input)
		if hashes[hash] {
			t.Errorf("Duplicate hash for different input: %s", hash)
		}
		hashes[hash] = true
	}
}

// TestSHA1_Coverage_BinaryData 测试二进制数据
func TestSHA1_Coverage_BinaryData(t *testing.T) {
	input := []byte{0x00, 0x01, 0x02, 0xFF, 0xFE, 0xFD}
	result := SHA1(input)
	if len(result) != 40 {
		t.Errorf("SHA1 hash length should be 40, got %d", len(result))
	}
}

// TestSHA1_Coverage_Concurrent 测试并发安全
func TestSHA1_Coverage_Concurrent(t *testing.T) {
	input := "concurrent test"
	results := make(chan string, 100)
	for i := 0; i < 100; i++ {
		go func() {
			results <- SHA1(input)
		}()
	}
	expected := SHA1(input)
	for i := 0; i < 100; i++ {
		if result := <-results; result != expected {
			t.Errorf("Concurrent SHA1 mismatch: got %s, want %s", result, expected)
		}
	}
}

// TestSHA1_Coverage_KnownValues 测试已知值
func TestSHA1_Coverage_KnownValues(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", "da39a3ee5e6b4b0d3255bfef95601890afd80709"},
		{"a", "86f7e437faa5a7fce15d1ddcb9eaeaea377667b8"},
		{"abc", "a9993e364706816aba3e25717850c26c9cd0d89d"},
		{"message digest", "c12252ceda8be8994d5fa0290a47231c1d16aae3"},
		{"abcdefghijklmnopqrstuvwxyz", "32d10c7b8cf96570ca04ce37f2a19d84240d3a89"},
	}
	for _, tc := range testCases {
		result := SHA1(tc.input)
		if result != tc.expected {
			t.Errorf("SHA1(%q) = %s, want %s", tc.input, result, tc.expected)
		}
	}
}

// encodeHexSHA1 辅助函数：将字节转换为 hex 字符串
func encodeHexSHA1(data []byte) string {
	const hexTable = "0123456789abcdef"
	result := make([]byte, len(data)*2)
	for i, b := range data {
		result[i*2] = hexTable[b>>4]
		result[i*2+1] = hexTable[b&0x0f]
	}
	return string(result)
}
