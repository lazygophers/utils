package cryptox

import (
	"strings"
	"testing"
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
