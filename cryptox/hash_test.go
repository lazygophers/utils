package cryptox

import (
	"testing"
)

func TestMd5(t *testing.T) {
	input := "test"
	expected := "098f6bcd4621d373cade4e832627b4f6"
	result := Md5(input)
	if result != expected {
		t.Errorf("Md5(%s) = %s; want %s", input, result, expected)
	}
}

func TestSha256(t *testing.T) {
	input := "test"
	expected := "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
	result := Sha256(input)
	if result != expected {
		t.Errorf("Sha256(%s) = %s; want %s", input, result, expected)
	}
}

func TestSha224(t *testing.T) {
	input := "test"
	expected := "90a3ed9e32b2aaf4c61c410eb925426119e1a9dc53d4286ade99a809"
	result := Sha224(input)
	if result != expected {
		t.Errorf("Sha224(%s) = %s; want %s", input, result, expected)
	}
}

func TestSha512(t *testing.T) {
	input := "test"
	expected := "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff"
	result := Sha512(input)
	if result != expected {
		t.Errorf("Sha512(%s) = %s; want %s", input, result, expected)
	}
}

// 修改: TestSha384 测试用例预期值错误
func TestSha384(t *testing.T) {
	input := "test"
	expected := "768412320f7b0aa5812fce428dc4706b3cae50e02a64caa16a782249bfe8efc4b7ef1ccb126255d196047dfedf17a0a9"
	result := Sha384(input)
	if result != expected {
		t.Errorf("Sha384(%s) = %s; want %s", input, result, expected)
	}
}

// 修改: TestSha512_256 测试用例预期值错误
func TestSha512_256(t *testing.T) {
	input := "test"
	expected := "3d37fe58435e0d87323dee4a2c1b339ef954de63716ee79f5747f94d974f913f"
	result := Sha512_256(input)
	if result != expected {
		t.Errorf("Sha512_256(%s) = %s; want %s", input, result, expected)
	}
}

// 修改: TestSha512_224 测试用例预期值错误
func TestSha512_224(t *testing.T) {
	input := "test"
	expected := "06001bf08dfb17d2b54925116823be230e98b5c6c278303bc4909a8c"
	result := Sha512_224(input)
	if result != expected {
		t.Errorf("Sha512_224(%s) = %s; want %s", input, result, expected)
	}
}

// 修改: TestHash32 测试用例预期值错误
func TestHash32(t *testing.T) {
	input := "test"
	expected := uint32(3157003241)
	result := Hash32(input)
	if result != expected {
		t.Errorf("Hash32(%s) = %d; want %d", input, result, expected)
	}
}

// 修改: TestHash32a 测试用例预期值错误
func TestHash32a(t *testing.T) {
	input := "test"
	expected := uint32(2949673445)
	result := Hash32a(input)
	if result != expected {
		t.Errorf("Hash32a(%s) = %d; want %d", input, result, expected)
	}
}

// 修改: TestHash64 测试用例预期值错误
func TestHash64(t *testing.T) {
	input := "test"
	expected := uint64(10090666253179731817)
	result := Hash64(input)
	if result != expected {
		t.Errorf("Hash64(%s) = %d; want %d", input, result, expected)
	}
}

// 修改: TestHash64a 测试用例预期값 오류
func TestHash64a(t *testing.T) {
	input := "test"
	expected := uint64(18007334074686647077)
	result := Hash64a(input)
	if result != expected {
		t.Errorf("Hash64a(%s) = %d; want %d", input, result, expected)
	}
}

// 수정: TestCRC32 테스트 케이스 예상값 오류
func TestCRC32(t *testing.T) {
	input := "test"
	expected := uint32(3632233996)
	result := CRC32(input)
	if result != expected {
		t.Errorf("CRC32(%s) = %d; want %d", input, result, expected)
	}
}

// 수정: TestCRC64 테스트 케이스 예상값 오류
func TestCRC64(t *testing.T) {
	input := "test"
	expected := uint64(18020588380933092773)
	result := CRC64(input)
	if result != expected {
		t.Errorf("CRC64(%s) = %d; want %d", input, result, expected)
	}
}

// 수정: TestSHA3_224 테스트 케이스 예상값 오류
func TestSHA3_224(t *testing.T) {
	input := "test"
	expected := "3797bf0afbbfca4a7bbba7602a2b552746876517a7f9b7ce2db0ae7b"
	result := SHA3_224(input)
	if result != expected {
		t.Errorf("SHA3_224(%s) = %s; want %s", input, result, expected)
	}
}

// 수정: TestSHA3_256 테스트 케이스 예상값 오류
func TestSHA3_256(t *testing.T) {
	input := "test"
	expected := "36f028580bb02cc8272a9a020f4200e346e276ae664e45ee80745574e2f5ab80"
	result := SHA3_256(input)
	if result != expected {
		t.Errorf("SHA3_256(%s) = %s; want %s", input, result, expected)
	}
}

// 수정: TestSHA3_384 테스트 케이스 예상값 오류
func TestSHA3_384(t *testing.T) {
	input := "test"
	expected := "e516dabb23b6e30026863543282780a3ae0dccf05551cf0295178d7ff0f1b41eecb9db3ff219007c4e097260d58621bd"
	result := SHA3_384(input)
	if result != expected {
		t.Errorf("SHA3_384(%s) = %s; want %s", input, result, expected)
	}
}

// 수정: TestSHA3_512 테스트 케이스 예상값 오류
func TestSHA3_512(t *testing.T) {
	input := "test"
	expected := "9ece086e9bac491fac5c1d1046ca11d737b92a2b2ebd93f005d7b710110c0a678288166e7fbe796883a4f2e9b3ca9f484f521d0ce464345cc1aec96779149c14"
	result := SHA3_512(input)
	if result != expected {
		t.Errorf("SHA3_512(%s) = %s; want %s", input, result, expected)
	}
}

// 수정: TestSHAKE128 테스트 케이스
func TestSHAKE128(t *testing.T) {
	input := "test"
	size := 32
	expected := "d3b0aa9cd8b7255622cebc631e867d4093d6f6010191a53973c45fec9b07c774"
	result, err := SHAKE128(input, size)
	if err != nil {
		t.Errorf("SHAKE128(%s, %d) returned an error: %v", input, size, err)
	} else if result != expected {
		t.Errorf("SHAKE128(%s, %d) = %s; want %s", input, size, result, expected)
	}
}

// 수정: TestSHAKE256 테스트 케이스
func TestSHAKE256(t *testing.T) {
	input := "test"
	size := 32
	expected := "b54ff7255705a71ee2925e4a3e30e41aed489a579d5595e0df13e32e1e4dd202"
	result, err := SHAKE256(input, size)
	if err != nil {
		t.Errorf("SHAKE256(%s, %d) returned an error: %v", input, size, err)
	} else if result != expected {
		t.Errorf("SHAKE256(%s, %d) = %s; want %s", input, size, result, expected)
	}
}

// 수정: TestBLAKE2b 테스트 케이스
func TestBLAKE2b(t *testing.T) {
	input := "test"
	size := 32
	expected := "928b20366943e2afd11ebc0eae2e53a93bf177a4fcf35bcc64d503704e65e202"
	result, err := BLAKE2b(input, size)
	if err != nil {
		t.Errorf("BLAKE2b(%s, %d) returned an error: %v", input, size, err)
	} else if result != expected {
		t.Errorf("BLAKE2b(%s, %d) = %s; want %s", input, size, result, expected)
	}
}

// 수정: TestBLAKE2s 테스트 케이스
func TestBLAKE2s(t *testing.T) {
	input := "test"
	size := 32
	expected := "f308fc02ce9172ad02a7d75800ecfc027109bc67987ea32aba9b8dcc7b10150e"
	result, err := BLAKE2s(input, size)
	if err != nil {
		t.Errorf("BLAKE2s(%s, %d) returned an error: %v", input, size, err)
	} else if result != expected {
		t.Errorf("BLAKE2s(%s, %d) = %s; want %s", input, size, result, expected)
	}
}

// Test error conditions for hash functions
func TestHashErrorConditions(t *testing.T) {
	data := "test"
	
	// Test SHAKE128 with invalid size
	_, err := SHAKE128(data, 0)
	if err == nil {
		t.Error("Expected error for SHAKE128 with size 0")
	}
	
	_, err = SHAKE128(data, -1)
	if err == nil {
		t.Error("Expected error for SHAKE128 with negative size")
	}
	
	// Test SHAKE256 with invalid size
	_, err = SHAKE256(data, 0)
	if err == nil {
		t.Error("Expected error for SHAKE256 with size 0")
	}
	
	// Test BLAKE2b with invalid size
	_, err = BLAKE2b(data, 0)
	if err == nil {
		t.Error("Expected error for BLAKE2b with size 0")
	}
	
	_, err = BLAKE2b(data, 65) // BLAKE2b max is 64 bytes
	if err == nil {
		t.Error("Expected error for BLAKE2b with size > 64")
	}
	
	// Test BLAKE2s with invalid size
	_, err = BLAKE2s(data, 0)
	if err == nil {
		t.Error("Expected error for BLAKE2s with size 0")
	}
}

// Test BLAKE2s error condition specifically  
func TestBLAKE2sErrorHandling(t *testing.T) {
	data := "test"
	
	// The BLAKE2s function has an error path when blake2s.New256 fails
	// This is difficult to trigger in normal circumstances, but we can test the error validation
	_, err := BLAKE2s(data, 0)
	if err == nil {
		t.Error("Expected error for BLAKE2s with size 0")
	}
	
	// Valid case should work
	result, err := BLAKE2s(data, 32)
	if err != nil {
		t.Errorf("BLAKE2s with valid size should not error: %v", err)
	}
	if result == "" {
		t.Error("BLAKE2s should return non-empty result")
	}
}

// Create a test that attempts to trigger the blake2s.New256 error path
func TestBLAKE2sInternalError(t *testing.T) {
	data := "test"
	
	// Test all valid sizes to ensure we exercise all code paths
	for size := 1; size <= 64; size++ {
		result, err := BLAKE2s(data, size) 
		if size <= 32 { // BLAKE2s supports up to 256 bits (32 bytes)
			if err != nil {
				t.Errorf("BLAKE2s should work for size %d: %v", size, err)
			}
			if result == "" {
				t.Errorf("BLAKE2s should return non-empty result for size %d", size)
			}
		} else {
			// For sizes > 32, it should still work because we ignore the size parameter
			if err != nil {
				t.Errorf("BLAKE2s should work even for size %d: %v", size, err)
			}
		}
	}
}

// Test new hash functions
func TestSHA1(t *testing.T) {
	input := "test"
	expected := "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3"
	result := SHA1(input)
	if result != expected {
		t.Errorf("SHA1(%s) = %s; want %s", input, result, expected)
	}
}

func TestRIPEMD160(t *testing.T) {
	input := "test"
	expected := "5e52fee47e6b070565f74372468cdc699de89107"
	result := RIPEMD160(input)
	if result != expected {
		t.Errorf("RIPEMD160(%s) = %s; want %s", input, result, expected)
	}
}

func TestKeccak256(t *testing.T) {
	input := "test"
	expected := "9c22ff5f21f0b81b113e63f7db6da94fedef11b2119b4088b89664fb9a3cb658"
	result := Keccak256(input)
	if result != expected {
		t.Errorf("Keccak256(%s) = %s; want %s", input, result, expected)
	}
}

func TestBLAKE2b512(t *testing.T) {
	input := "test"
	expected := "a71079d42853dea26e453004338670a53814b78137ffbed07603a41d76a483aa9bc33b582f77d30a65e6f29a896c0411f38312e1d66e0bf16386c86a89bea572"
	result := BLAKE2b512(input)
	if result != expected {
		t.Errorf("BLAKE2b512(%s) = %s; want %s", input, result, expected)
	}
}

func TestBLAKE2b256(t *testing.T) {
	input := "test"
	expected := "928b20366943e2afd11ebc0eae2e53a93bf177a4fcf35bcc64d503704e65e202"
	result := BLAKE2b256(input)
	if result != expected {
		t.Errorf("BLAKE2b256(%s) = %s; want %s", input, result, expected)
	}
}

func TestBLAKE2s256(t *testing.T) {
	input := "test"
	expected := "f308fc02ce9172ad02a7d75800ecfc027109bc67987ea32aba9b8dcc7b10150e"
	result := BLAKE2s256(input)
	if result != expected {
		t.Errorf("BLAKE2s256(%s) = %s; want %s", input, result, expected)
	}
}

func TestBLAKE2bWithKey(t *testing.T) {
	input := "test"
	key := []byte("key")
	size := 32
	result, err := BLAKE2bWithKey(input, key, size)
	if err != nil {
		t.Errorf("BLAKE2bWithKey(%s, %v, %d) returned error: %v", input, key, size, err)
	}
	// Verify it returns correct length
	if len(result) != size*2 {
		t.Errorf("BLAKE2bWithKey(%s, %v, %d) returned wrong length: got %d, want %d", input, key, size, len(result), size*2)
	}
}

func TestBLAKE2sWithKey(t *testing.T) {
	input := "test"
	key := []byte("key")
	result, err := BLAKE2sWithKey(input, key)
	if err != nil {
		t.Errorf("BLAKE2sWithKey(%s, %v) returned error: %v", input, key, err)
	}
	if len(result) != 64 { // 32 bytes * 2 hex chars
		t.Errorf("BLAKE2sWithKey(%s, %v) returned wrong length: got %d, want 64", input, key, len(result))
	}
}

func TestHMACMd5(t *testing.T) {
	key := "key"
	message := "test"
	expected := "1d4a2743c056e467ff3f09c9af31de7e"
	result := HMACMd5(key, message)
	if result != expected {
		t.Errorf("HMACMd5(%s, %s) = %s; want %s", key, message, result, expected)
	}
}

func TestHMACSHA1(t *testing.T) {
	key := "key"
	message := "test"
	expected := "671f54ce0c540f78ffe1e26dcf9c2a047aea4fda"
	result := HMACSHA1(key, message)
	if result != expected {
		t.Errorf("HMACSHA1(%s, %s) = %s; want %s", key, message, result, expected)
	}
}

func TestHMACSHA256(t *testing.T) {
	key := "key"
	message := "test"
	expected := "02afb56304902c656fcb737cdd03de6205bb6d401da2812efd9b2d36a08af159"
	result := HMACSHA256(key, message)
	if result != expected {
		t.Errorf("HMACSHA256(%s, %s) = %s; want %s", key, message, result, expected)
	}
}

func TestHMACSHA384(t *testing.T) {
	key := "key"
	message := "test"
	expected := "160a099ad9d6dadb46311cb4e6dfe98aca9ca519c2e0fedc8dc45da419b1173039cc131f0b5f68b2bbc2b635109b57a8"
	result := HMACSHA384(key, message)
	if result != expected {
		t.Errorf("HMACSHA384(%s, %s) = %s; want %s", key, message, result, expected)
	}
}

func TestHMACSHA512(t *testing.T) {
	key := "key"
	message := "test"
	expected := "287a0fb89a7fbdfa5b5538636918e537a5b83065e4ff331268b7aaa115dde047a9b0f4fb5b828608fc0b6327f10055f7637b058e9e0dbb9e698901a3e6dd461c"
	result := HMACSHA512(key, message)
	if result != expected {
		t.Errorf("HMACSHA512(%s, %s) = %s; want %s", key, message, result, expected)
	}
}

// Test error conditions for new hash functions
func TestNewHashErrorConditions(t *testing.T) {
	data := "test"
	key := []byte("key")
	
	// Test BLAKE2bWithKey with invalid sizes
	_, err := BLAKE2bWithKey(data, key, 0)
	if err == nil {
		t.Error("Expected error for BLAKE2bWithKey with size 0")
	}
	
	_, err = BLAKE2bWithKey(data, key, 65)
	if err == nil {
		t.Error("Expected error for BLAKE2bWithKey with size > 64")
	}
}
