package cryptox

import (
	"errors"
	"hash"
	"testing"
)

// Mock failure functions
func FailingBLAKE2bNew(size int, key []byte) (hash.Hash, error) {
	return nil, errors.New("simulated BLAKE2b New failure")
}

func FailingBLAKE2sNew256(key []byte) (hash.Hash, error) {
	return nil, errors.New("simulated BLAKE2s New256 failure")
}

// TestHash_100PercentCoverage tests all error paths in hash functions
func TestHash_100PercentCoverage(t *testing.T) {
	// Save original functions
	originalBLAKE2bNew := blake2bNew
	originalBLAKE2sNew256 := blake2sNew256

	// Restore original functions after test
	defer func() {
		blake2bNew = originalBLAKE2bNew
		blake2sNew256 = originalBLAKE2sNew256
	}()

	validData := "test"
	validKey := []byte("key")

	// Test 1: Trigger BLAKE2b New failure in BLAKE2bWithKey
	blake2bNew = FailingBLAKE2bNew
	blake2sNew256 = originalBLAKE2sNew256

	_, err := BLAKE2bWithKey(validData, validKey, 32)
	if err == nil {
		t.Error("Expected BLAKE2b New error in BLAKE2bWithKey")
	}

	// Test 2: Trigger BLAKE2s New256 failure in BLAKE2sWithKey
	blake2bNew = originalBLAKE2bNew
	blake2sNew256 = FailingBLAKE2sNew256

	_, err = BLAKE2sWithKey(validData, validKey)
	if err == nil {
		t.Error("Expected BLAKE2s New256 error in BLAKE2sWithKey")
	}

	// Test 2a: Trigger BLAKE2s New256 failure in original BLAKE2s function
	// This will test the error path in the original BLAKE2s function
	_, err = BLAKE2s(validData, 32)
	if err == nil {
		t.Error("Expected BLAKE2s New256 error in BLAKE2s function")
	}

	// Test 3: Trigger existing BLAKE2s error path in original BLAKE2s function
	// Reset functions to original
	blake2bNew = originalBLAKE2bNew
	blake2sNew256 = originalBLAKE2sNew256

	// Test with invalid size (already tested in main test file, but ensure coverage)
	_, err = BLAKE2s(validData, 0)
	if err == nil {
		t.Error("Expected error for BLAKE2s with size 0")
	}

	_, err = BLAKE2s(validData, -1)
	if err == nil {
		t.Error("Expected error for BLAKE2s with negative size")
	}

	// Test 4: Test BLAKE2b function error paths
	_, err = BLAKE2b(validData, 0)
	if err == nil {
		t.Error("Expected error for BLAKE2b with size 0")
	}

	_, err = BLAKE2b(validData, 65)
	if err == nil {
		t.Error("Expected error for BLAKE2b with size > 64")
	}

	// Test 5: Test SHAKE functions error paths
	_, err = SHAKE128(validData, 0)
	if err == nil {
		t.Error("Expected error for SHAKE128 with size 0")
	}

	_, err = SHAKE128(validData, -1)
	if err == nil {
		t.Error("Expected error for SHAKE128 with negative size")
	}

	_, err = SHAKE256(validData, 0)
	if err == nil {
		t.Error("Expected error for SHAKE256 with size 0")
	}

	_, err = SHAKE256(validData, -1)
	if err == nil {
		t.Error("Expected error for SHAKE256 with negative size")
	}

	// Test 6: Test BLAKE2bWithKey error paths
	_, err = BLAKE2bWithKey(validData, validKey, 0)
	if err == nil {
		t.Error("Expected error for BLAKE2bWithKey with size 0")
	}

	_, err = BLAKE2bWithKey(validData, validKey, -1)
	if err == nil {
		t.Error("Expected error for BLAKE2bWithKey with negative size")
	}

	_, err = BLAKE2bWithKey(validData, validKey, 65)
	if err == nil {
		t.Error("Expected error for BLAKE2bWithKey with size > 64")
	}
}

// TestHashEdgeCases tests additional edge cases for complete coverage
func TestHashEdgeCases(t *testing.T) {
	// Test with empty input
	emptyInput := ""
	emptyBytes := []byte("")

	// Test all hash functions with empty input
	result := Md5(emptyInput)
	if len(result) != 32 {
		t.Error("MD5 should return 32 character hex string")
	}

	result = SHA1(emptyInput)
	if len(result) != 40 {
		t.Error("SHA1 should return 40 character hex string")
	}

	result = Sha256(emptyInput)
	if len(result) != 64 {
		t.Error("SHA256 should return 64 character hex string")
	}

	result = RIPEMD160(emptyInput)
	if len(result) != 40 {
		t.Error("RIPEMD160 should return 40 character hex string")
	}

	result = Keccak256(emptyInput)
	if len(result) != 64 {
		t.Error("Keccak256 should return 64 character hex string")
	}

	result = BLAKE2b512(emptyInput)
	if len(result) != 128 {
		t.Error("BLAKE2b512 should return 128 character hex string")
	}

	result = BLAKE2b256(emptyInput)
	if len(result) != 64 {
		t.Error("BLAKE2b256 should return 64 character hex string")
	}

	result = BLAKE2s256(emptyInput)
	if len(result) != 64 {
		t.Error("BLAKE2s256 should return 64 character hex string")
	}

	// Test HMAC functions with empty input
	emptyKey := ""
	result = HMACMd5(emptyKey, emptyInput)
	if len(result) != 32 {
		t.Error("HMACMd5 should return 32 character hex string")
	}

	result = HMACSHA1(emptyKey, emptyInput)
	if len(result) != 40 {
		t.Error("HMACSHA1 should return 40 character hex string")
	}

	result = HMACSHA256(emptyKey, emptyInput)
	if len(result) != 64 {
		t.Error("HMACSHA256 should return 64 character hex string")
	}

	result = HMACSHA384(emptyKey, emptyInput)
	if len(result) != 96 {
		t.Error("HMACSHA384 should return 96 character hex string")
	}

	result = HMACSHA512(emptyKey, emptyInput)
	if len(result) != 128 {
		t.Error("HMACSHA512 should return 128 character hex string")
	}

	// Test with byte slice input
	result = Md5(emptyBytes)
	if len(result) != 32 {
		t.Error("MD5 with bytes should return 32 character hex string")
	}

	// Test hash functions with various sizes
	for size := 1; size <= 64; size++ {
		result, err := BLAKE2b("test", size)
		if err != nil {
			t.Errorf("BLAKE2b should work for size %d: %v", size, err)
		}
		if len(result) != size*2 {
			t.Errorf("BLAKE2b size %d should return %d hex characters, got %d", size, size*2, len(result))
		}
	}

	// Test SHAKE functions with various sizes
	for size := 1; size <= 100; size++ {
		result, err := SHAKE128("test", size)
		if err != nil {
			t.Errorf("SHAKE128 should work for size %d: %v", size, err)
		}
		if len(result) != size*2 {
			t.Errorf("SHAKE128 size %d should return %d hex characters, got %d", size, size*2, len(result))
		}

		result, err = SHAKE256("test", size)
		if err != nil {
			t.Errorf("SHAKE256 should work for size %d: %v", size, err)
		}
		if len(result) != size*2 {
			t.Errorf("SHAKE256 size %d should return %d hex characters, got %d", size, size*2, len(result))
		}
	}

	// Test BLAKE2bWithKey with various key sizes
	for keySize := 0; keySize <= 64; keySize++ {
		key := make([]byte, keySize)
		result, err := BLAKE2bWithKey("test", key, 32)
		if err != nil {
			t.Errorf("BLAKE2bWithKey should work with key size %d: %v", keySize, err)
		}
		if len(result) != 64 {
			t.Errorf("BLAKE2bWithKey should return 64 hex characters, got %d", len(result))
		}
	}

	// Test BLAKE2sWithKey with various key sizes
	for keySize := 0; keySize <= 32; keySize++ {
		key := make([]byte, keySize)
		result, err := BLAKE2sWithKey("test", key)
		if err != nil {
			t.Errorf("BLAKE2sWithKey should work with key size %d: %v", keySize, err)
		}
		if len(result) != 64 {
			t.Errorf("BLAKE2sWithKey should return 64 hex characters, got %d", len(result))
		}
	}

	// Test numeric hash functions
	testData := "test data"
	hash32 := Hash32(testData)
	if hash32 == 0 {
		t.Error("Hash32 should not return 0 for non-empty input")
	}

	hash32a := Hash32a(testData)
	if hash32a == 0 {
		t.Error("Hash32a should not return 0 for non-empty input")
	}

	hash64 := Hash64(testData)
	if hash64 == 0 {
		t.Error("Hash64 should not return 0 for non-empty input")
	}

	hash64a := Hash64a(testData)
	if hash64a == 0 {
		t.Error("Hash64a should not return 0 for non-empty input")
	}

	crc32 := CRC32(testData)
	if crc32 == 0 {
		t.Error("CRC32 should not return 0 for non-empty input")
	}

	crc64 := CRC64(testData)
	if crc64 == 0 {
		t.Error("CRC64 should not return 0 for non-empty input")
	}
}

// TestHashTypeVariations tests string vs []byte type variations
func TestHashTypeVariations(t *testing.T) {
	testString := "test"
	testBytes := []byte("test")

	// Verify string and bytes produce same results
	if Md5(testString) != Md5(testBytes) {
		t.Error("MD5 should produce same result for string and []byte")
	}

	if SHA1(testString) != SHA1(testBytes) {
		t.Error("SHA1 should produce same result for string and []byte")
	}

	if Sha256(testString) != Sha256(testBytes) {
		t.Error("SHA256 should produce same result for string and []byte")
	}

	if RIPEMD160(testString) != RIPEMD160(testBytes) {
		t.Error("RIPEMD160 should produce same result for string and []byte")
	}

	if Keccak256(testString) != Keccak256(testBytes) {
		t.Error("Keccak256 should produce same result for string and []byte")
	}

	if BLAKE2b512(testString) != BLAKE2b512(testBytes) {
		t.Error("BLAKE2b512 should produce same result for string and []byte")
	}

	if BLAKE2b256(testString) != BLAKE2b256(testBytes) {
		t.Error("BLAKE2b256 should produce same result for string and []byte")
	}

	if BLAKE2s256(testString) != BLAKE2s256(testBytes) {
		t.Error("BLAKE2s256 should produce same result for string and []byte")
	}

	// Test HMAC functions with mixed types
	keyString := "key"
	keyBytes := []byte("key")

	if HMACMd5(keyString, testString) != HMACMd5(keyBytes, testBytes) {
		t.Error("HMACMd5 should produce same result for string and []byte")
	}

	if HMACSHA1(keyString, testString) != HMACSHA1(keyBytes, testBytes) {
		t.Error("HMACSHA1 should produce same result for string and []byte")
	}

	if HMACSHA256(keyString, testString) != HMACSHA256(keyBytes, testBytes) {
		t.Error("HMACSHA256 should produce same result for string and []byte")
	}

	if HMACSHA384(keyString, testString) != HMACSHA384(keyBytes, testBytes) {
		t.Error("HMACSHA384 should produce same result for string and []byte")
	}

	if HMACSHA512(keyString, testString) != HMACSHA512(keyBytes, testBytes) {
		t.Error("HMACSHA512 should produce same result for string and []byte")
	}
}