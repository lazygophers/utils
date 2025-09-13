package cryptox

import (
	"errors"
	"testing"
)

// Mock failures for KDF dependency injection
type FailingKDFReader struct{}

func (fr FailingKDFReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("simulated random reader failure")
}

// TestKDF_100PercentCoverage tests all error paths in KDF functions
func TestKDF_100PercentCoverage(t *testing.T) {
	// Save original functions
	originalKDFRandReader := kdfRandReader

	// Restore original functions after test
	defer func() {
		kdfRandReader = originalKDFRandReader
	}()

	validPassword := "test_password"
	validConfig := DefaultPBKDF2Config()

	// Test 1: Trigger rand.Reader failure in PBKDF2Generate
	kdfRandReader = FailingKDFReader{}

	_, _, err := PBKDF2Generate(validPassword, validConfig)
	if err == nil {
		t.Error("Expected rand.Reader error in PBKDF2Generate")
	}

	// Test 2: Trigger rand.Reader failure in ScryptGenerate
	scryptConfig := DefaultScryptConfig()
	_, _, err = ScryptGenerate(validPassword, scryptConfig)
	if err == nil {
		t.Error("Expected rand.Reader error in ScryptGenerate")
	}

	// Test 3: Trigger rand.Reader failure in Argon2Generate
	argon2Config := DefaultArgon2Config()
	_, _, err = Argon2Generate(validPassword, argon2Config)
	if err == nil {
		t.Error("Expected rand.Reader error in Argon2Generate")
	}

	// Test 4: Trigger rand.Reader failure in GenerateSalt
	_, err = GenerateSalt(16)
	if err == nil {
		t.Error("Expected rand.Reader error in GenerateSalt")
	}

	// Restore original reader for remaining tests
	kdfRandReader = originalKDFRandReader

	// Test 5: Test ScryptGenerate with Scrypt derivation failure
	// This tests the error path where ScryptDerive fails inside ScryptGenerate
	invalidScryptConfig := ScryptConfig{
		SaltLength: 16,
		N:          15, // Invalid: not a power of 2
		R:          8,
		P:          1,
		KeyLength:  32,
	}
	_, _, err = ScryptGenerate(validPassword, invalidScryptConfig)
	if err == nil {
		t.Error("Expected Scrypt derivation error in ScryptGenerate")
	}

	// Test 6: Test ScryptVerify with derivation failure
	// Create a valid key and salt first
	validScryptConfig := DefaultScryptConfig()
	key, salt, err := ScryptGenerate(validPassword, validScryptConfig)
	if err != nil {
		t.Fatalf("Failed to generate test key: %v", err)
	}

	// Now test verification with invalid config (should trigger error path in ScryptVerify)
	result := ScryptVerify(validPassword, key, salt, invalidScryptConfig)
	if result {
		t.Error("ScryptVerify should return false when derivation fails")
	}
}

// TestKDFEdgeCasesAndBoundaries tests edge cases and boundary conditions
func TestKDFEdgeCasesAndBoundaries(t *testing.T) {
	// Test minimum valid values
	minConfig := PBKDF2Config{
		SaltLength: 1,
		Iterations: 1,
		KeyLength:  1,
	}
	key, salt, err := PBKDF2Generate("test", minConfig)
	if err != nil {
		t.Errorf("PBKDF2Generate should work with minimum config: %v", err)
	}
	if len(key) != 1 || len(salt) != 1 {
		t.Error("Should generate correct lengths for minimum config")
	}

	// Test Scrypt with minimum valid values
	minScryptConfig := ScryptConfig{
		SaltLength: 1,
		N:          2, // Minimum power of 2 > 1
		R:          1,
		P:          1,
		KeyLength:  1,
	}
	key2, salt2, err := ScryptGenerate("test", minScryptConfig)
	if err != nil {
		t.Errorf("ScryptGenerate should work with minimum config: %v", err)
	}
	if len(key2) != 1 || len(salt2) != 1 {
		t.Error("Should generate correct lengths for minimum Scrypt config")
	}

	// Test Argon2 with minimum valid values
	minArgon2Config := Argon2Config{
		SaltLength: 1,
		Time:       1,
		Memory:     1,
		Threads:    1,
		KeyLength:  1,
	}
	key3, salt3, err := Argon2Generate("test", minArgon2Config)
	if err != nil {
		t.Errorf("Argon2Generate should work with minimum config: %v", err)
	}
	if len(key3) != 1 || len(salt3) != 1 {
		t.Error("Should generate correct lengths for minimum Argon2 config")
	}

	// Test large values
	largeConfig := PBKDF2Config{
		SaltLength: 128,
		Iterations: 1000,
		KeyLength:  128,
	}
	key4, salt4, err := PBKDF2Generate("test", largeConfig)
	if err != nil {
		t.Errorf("PBKDF2Generate should work with large config: %v", err)
	}
	if len(key4) != 128 || len(salt4) != 128 {
		t.Error("Should generate correct lengths for large config")
	}

	// Test GenerateSalt with various sizes
	for _, size := range []int{1, 8, 16, 32, 64, 128} {
		salt, err := GenerateSalt(size)
		if err != nil {
			t.Errorf("GenerateSalt should work with size %d: %v", size, err)
		}
		if len(salt) != size {
			t.Errorf("Expected salt length %d, got %d", size, len(salt))
		}
	}
}

// TestConstantTimeCompareEdgeCases tests constant time compare with edge cases
func TestConstantTimeCompareEdgeCases(t *testing.T) {
	// Test with nil slices
	if constantTimeCompare(nil, nil) != true {
		t.Error("Two nil slices should be equal")
	}

	if constantTimeCompare(nil, []byte{}) != true {
		t.Error("nil and empty slice should be equal")
	}

	if constantTimeCompare([]byte{}, nil) != true {
		t.Error("empty slice and nil should be equal")
	}

	if constantTimeCompare([]byte{}, []byte{}) != true {
		t.Error("Two empty slices should be equal")
	}

	// Test with single byte
	if constantTimeCompare([]byte{1}, []byte{1}) != true {
		t.Error("Single identical bytes should be equal")
	}

	if constantTimeCompare([]byte{1}, []byte{2}) != false {
		t.Error("Single different bytes should not be equal")
	}

	// Test with all zero bytes
	zeros1 := make([]byte, 16)
	zeros2 := make([]byte, 16)
	if constantTimeCompare(zeros1, zeros2) != true {
		t.Error("All zero bytes should be equal")
	}

	// Test with all 0xFF bytes
	ones1 := make([]byte, 16)
	ones2 := make([]byte, 16)
	for i := range ones1 {
		ones1[i] = 0xFF
		ones2[i] = 0xFF
	}
	if constantTimeCompare(ones1, ones2) != true {
		t.Error("All 0xFF bytes should be equal")
	}

	// Test with different lengths
	short := []byte{1, 2, 3}
	long := []byte{1, 2, 3, 4}
	if constantTimeCompare(short, long) != false {
		t.Error("Different length slices should not be equal")
	}
}

// TestKDFVerificationEdgeCases tests verification functions with edge cases
func TestKDFVerificationEdgeCases(t *testing.T) {
	config := DefaultPBKDF2Config()

	// Generate a valid key and salt
	key, salt, err := PBKDF2Generate("test", config)
	if err != nil {
		t.Fatalf("Failed to generate test key: %v", err)
	}

	// Test verification with empty password
	if PBKDF2Verify("", key, salt, config) {
		t.Error("Empty password should not verify against non-empty password key")
	}

	// Test with modified key (single bit flip)
	modifiedKey := make([]byte, len(key))
	copy(modifiedKey, key)
	modifiedKey[0] ^= 0x01 // Flip one bit

	if PBKDF2Verify("test", modifiedKey, salt, config) {
		t.Error("Modified key should not verify")
	}

	// Test with modified salt
	modifiedSalt := make([]byte, len(salt))
	copy(modifiedSalt, salt)
	modifiedSalt[0] ^= 0x01 // Flip one bit

	if PBKDF2Verify("test", key, modifiedSalt, config) {
		t.Error("Modified salt should not verify")
	}

	// Test Scrypt verification edge cases
	scryptConfig := DefaultScryptConfig()
	scryptKey, scryptSalt, err := ScryptGenerate("test", scryptConfig)
	if err != nil {
		t.Fatalf("Failed to generate Scrypt test key: %v", err)
	}

	// Test with truncated key
	truncatedKey := scryptKey[:len(scryptKey)-1]
	if ScryptVerify("test", truncatedKey, scryptSalt, scryptConfig) {
		t.Error("Truncated key should not verify")
	}

	// Test Argon2 verification edge cases
	argon2Config := DefaultArgon2Config()
	argon2Key, argon2Salt, err := Argon2Generate("test", argon2Config)
	if err != nil {
		t.Fatalf("Failed to generate Argon2 test key: %v", err)
	}

	// Test with extended key
	extendedKey := make([]byte, len(argon2Key)+1)
	copy(extendedKey, argon2Key)
	if Argon2Verify("test", extendedKey, argon2Salt, argon2Config) {
		t.Error("Extended key should not verify")
	}
}

// TestKDFConfigurationVariations tests different parameter combinations
func TestKDFConfigurationVariations(t *testing.T) {
	// Test PBKDF2 with different iteration counts
	baseConfig := DefaultPBKDF2Config()

	for _, iterations := range []int{1, 10, 100, 1000, 10000} {
		config := baseConfig
		config.Iterations = iterations

		key, _, err := PBKDF2Generate("test", config)
		if err != nil {
			t.Errorf("PBKDF2 should work with %d iterations: %v", iterations, err)
		}
		if len(key) != config.KeyLength {
			t.Errorf("Should generate correct key length for %d iterations", iterations)
		}
	}

	// Test Scrypt with different N values (powers of 2)
	scryptConfig := DefaultScryptConfig()

	for _, n := range []int{2, 4, 8, 16, 32, 64} {
		config := scryptConfig
		config.N = n

		key, _, err := ScryptGenerate("test", config)
		if err != nil {
			t.Errorf("Scrypt should work with N=%d: %v", n, err)
		}
		if len(key) != config.KeyLength {
			t.Errorf("Should generate correct key length for N=%d", n)
		}
	}

	// Test Argon2 with different memory values
	argon2Config := DefaultArgon2Config()

	for _, memory := range []uint32{1, 8, 64, 512, 1024} {
		config := argon2Config
		config.Memory = memory

		key, _, err := Argon2Generate("test", config)
		if err != nil {
			t.Errorf("Argon2 should work with memory=%d: %v", memory, err)
		}
		if len(key) != config.KeyLength {
			t.Errorf("Should generate correct key length for memory=%d", memory)
		}
	}
}

// TestKDFWithSpecialPasswords tests KDF functions with special password cases
func TestKDFWithSpecialPasswords(t *testing.T) {
	config := DefaultPBKDF2Config()

	specialPasswords := []string{
		"",     // Empty
		"a",    // Single character
		"123",  // Numeric
		"🔐🗝️🔒", // Emoji only
		"password with spaces",
		"パスワード",                         // Japanese
		"كلمة المرور",                   // Arabic
		string([]byte{0, 1, 2, 3, 255}), // Binary data
	}

	keys := make([][]byte, len(specialPasswords))

	for i, password := range specialPasswords {
		key, salt, err := PBKDF2Generate(password, config)
		if err != nil {
			t.Errorf("PBKDF2 should work with special password %d: %v", i, err)
			continue
		}

		if len(key) != config.KeyLength {
			t.Errorf("Should generate correct key length for special password %d", i)
		}

		// Verify the password
		if !PBKDF2Verify(password, key, salt, config) {
			t.Errorf("Should verify special password %d", i)
		}

		keys[i] = key
	}

	// Ensure different passwords produce different keys
	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			if keys[i] != nil && keys[j] != nil && constantTimeCompare(keys[i], keys[j]) {
				t.Errorf("Different passwords (indices %d and %d) should produce different keys", i, j)
			}
		}
	}
}
