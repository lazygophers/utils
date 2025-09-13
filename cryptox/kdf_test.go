package cryptox

import (
	"bytes"
	"testing"
)

// Test data
const (
	testPassword = "test_password_123"
)

var (
	testSalt = []byte("test_salt_16byte")
)

// TestPBKDF2WithSHA256 tests PBKDF2 with SHA256
func TestPBKDF2WithSHA256(t *testing.T) {
	password := []byte(testPassword)
	salt := testSalt
	iterations := 1000
	keyLength := 32

	key := PBKDF2WithSHA256(password, salt, iterations, keyLength)

	if len(key) != keyLength {
		t.Errorf("Expected key length %d, got %d", keyLength, len(key))
	}

	// Test reproducibility
	key2 := PBKDF2WithSHA256(password, salt, iterations, keyLength)
	if !bytes.Equal(key, key2) {
		t.Error("PBKDF2WithSHA256 should produce consistent results")
	}

	// Test different parameters produce different results
	key3 := PBKDF2WithSHA256(password, salt, iterations+1, keyLength)
	if bytes.Equal(key, key3) {
		t.Error("Different iterations should produce different keys")
	}
}

// TestPBKDF2WithSHA1 tests PBKDF2 with SHA1
func TestPBKDF2WithSHA1(t *testing.T) {
	password := []byte(testPassword)
	salt := testSalt
	iterations := 1000
	keyLength := 20

	key := PBKDF2WithSHA1(password, salt, iterations, keyLength)

	if len(key) != keyLength {
		t.Errorf("Expected key length %d, got %d", keyLength, len(key))
	}

	// Test reproducibility
	key2 := PBKDF2WithSHA1(password, salt, iterations, keyLength)
	if !bytes.Equal(key, key2) {
		t.Error("PBKDF2WithSHA1 should produce consistent results")
	}
}

// TestPBKDF2WithSHA512 tests PBKDF2 with SHA512
func TestPBKDF2WithSHA512(t *testing.T) {
	password := []byte(testPassword)
	salt := testSalt
	iterations := 1000
	keyLength := 64

	key := PBKDF2WithSHA512(password, salt, iterations, keyLength)

	if len(key) != keyLength {
		t.Errorf("Expected key length %d, got %d", keyLength, len(key))
	}

	// Test reproducibility
	key2 := PBKDF2WithSHA512(password, salt, iterations, keyLength)
	if !bytes.Equal(key, key2) {
		t.Error("PBKDF2WithSHA512 should produce consistent results")
	}
}

// TestDefaultPBKDF2Config tests default PBKDF2 configuration
func TestDefaultPBKDF2Config(t *testing.T) {
	config := DefaultPBKDF2Config()

	if config.SaltLength <= 0 {
		t.Error("Default salt length should be positive")
	}
	if config.Iterations <= 0 {
		t.Error("Default iterations should be positive")
	}
	if config.KeyLength <= 0 {
		t.Error("Default key length should be positive")
	}
}

// TestPBKDF2Generate tests PBKDF2 key generation with random salt
func TestPBKDF2Generate(t *testing.T) {
	config := DefaultPBKDF2Config()

	key, salt, err := PBKDF2Generate(testPassword, config)
	if err != nil {
		t.Fatalf("PBKDF2Generate failed: %v", err)
	}

	if len(key) != config.KeyLength {
		t.Errorf("Expected key length %d, got %d", config.KeyLength, len(key))
	}
	if len(salt) != config.SaltLength {
		t.Errorf("Expected salt length %d, got %d", config.SaltLength, len(salt))
	}

	// Test that different calls produce different salts
	_, salt2, err := PBKDF2Generate(testPassword, config)
	if err != nil {
		t.Fatalf("Second PBKDF2Generate failed: %v", err)
	}
	if bytes.Equal(salt, salt2) {
		t.Error("Different calls should produce different salts")
	}
}

// TestPBKDF2Verify tests PBKDF2 password verification
func TestPBKDF2Verify(t *testing.T) {
	config := DefaultPBKDF2Config()

	key, salt, err := PBKDF2Generate(testPassword, config)
	if err != nil {
		t.Fatalf("PBKDF2Generate failed: %v", err)
	}

	// Test correct password verification
	if !PBKDF2Verify(testPassword, key, salt, config) {
		t.Error("Should verify correct password")
	}

	// Test incorrect password verification
	if PBKDF2Verify("wrong_password", key, salt, config) {
		t.Error("Should not verify incorrect password")
	}
}

// TestScryptDerive tests Scrypt key derivation
func TestScryptDerive(t *testing.T) {
	password := []byte(testPassword)
	salt := testSalt
	config := DefaultScryptConfig()

	key, err := ScryptDerive(password, salt, config)
	if err != nil {
		t.Fatalf("ScryptDerive failed: %v", err)
	}

	if len(key) != config.KeyLength {
		t.Errorf("Expected key length %d, got %d", config.KeyLength, len(key))
	}

	// Test reproducibility
	key2, err := ScryptDerive(password, salt, config)
	if err != nil {
		t.Fatalf("Second ScryptDerive failed: %v", err)
	}
	if !bytes.Equal(key, key2) {
		t.Error("ScryptDerive should produce consistent results")
	}
}

// TestDefaultScryptConfig tests default Scrypt configuration
func TestDefaultScryptConfig(t *testing.T) {
	config := DefaultScryptConfig()

	if config.SaltLength <= 0 {
		t.Error("Default salt length should be positive")
	}
	if config.N <= 0 || (config.N&(config.N-1)) != 0 {
		t.Error("Default N should be a positive power of 2")
	}
	if config.R <= 0 {
		t.Error("Default R should be positive")
	}
	if config.P <= 0 {
		t.Error("Default P should be positive")
	}
	if config.KeyLength <= 0 {
		t.Error("Default key length should be positive")
	}
}

// TestScryptGenerate tests Scrypt key generation with random salt
func TestScryptGenerate(t *testing.T) {
	config := DefaultScryptConfig()

	key, salt, err := ScryptGenerate(testPassword, config)
	if err != nil {
		t.Fatalf("ScryptGenerate failed: %v", err)
	}

	if len(key) != config.KeyLength {
		t.Errorf("Expected key length %d, got %d", config.KeyLength, len(key))
	}
	if len(salt) != config.SaltLength {
		t.Errorf("Expected salt length %d, got %d", config.SaltLength, len(salt))
	}

	// Test that different calls produce different salts
	_, salt2, err := ScryptGenerate(testPassword, config)
	if err != nil {
		t.Fatalf("Second ScryptGenerate failed: %v", err)
	}
	if bytes.Equal(salt, salt2) {
		t.Error("Different calls should produce different salts")
	}
}

// TestScryptVerify tests Scrypt password verification
func TestScryptVerify(t *testing.T) {
	config := DefaultScryptConfig()

	key, salt, err := ScryptGenerate(testPassword, config)
	if err != nil {
		t.Fatalf("ScryptGenerate failed: %v", err)
	}

	// Test correct password verification
	if !ScryptVerify(testPassword, key, salt, config) {
		t.Error("Should verify correct password")
	}

	// Test incorrect password verification
	if ScryptVerify("wrong_password", key, salt, config) {
		t.Error("Should not verify incorrect password")
	}
}

// TestArgon2IDDerive tests Argon2id key derivation
func TestArgon2IDDerive(t *testing.T) {
	password := []byte(testPassword)
	salt := testSalt
	config := DefaultArgon2Config()

	key := Argon2IDDerive(password, salt, config)

	if len(key) != config.KeyLength {
		t.Errorf("Expected key length %d, got %d", config.KeyLength, len(key))
	}

	// Test reproducibility
	key2 := Argon2IDDerive(password, salt, config)
	if !bytes.Equal(key, key2) {
		t.Error("Argon2IDDerive should produce consistent results")
	}
}

// TestArgon2IDerive tests Argon2i key derivation
func TestArgon2IDerive(t *testing.T) {
	password := []byte(testPassword)
	salt := testSalt
	config := DefaultArgon2Config()

	key := Argon2IDerive(password, salt, config)

	if len(key) != config.KeyLength {
		t.Errorf("Expected key length %d, got %d", config.KeyLength, len(key))
	}

	// Test reproducibility
	key2 := Argon2IDerive(password, salt, config)
	if !bytes.Equal(key, key2) {
		t.Error("Argon2IDerive should produce consistent results")
	}

	// Test that Argon2i and Argon2id produce different results
	keyID := Argon2IDDerive(password, salt, config)
	if bytes.Equal(key, keyID) {
		t.Error("Argon2i and Argon2id should produce different results")
	}
}

// TestDefaultArgon2Config tests default Argon2 configuration
func TestDefaultArgon2Config(t *testing.T) {
	config := DefaultArgon2Config()

	if config.SaltLength <= 0 {
		t.Error("Default salt length should be positive")
	}
	if config.Time == 0 {
		t.Error("Default time should be positive")
	}
	if config.Memory == 0 {
		t.Error("Default memory should be positive")
	}
	if config.Threads == 0 {
		t.Error("Default threads should be positive")
	}
	if config.KeyLength <= 0 {
		t.Error("Default key length should be positive")
	}
}

// TestArgon2Generate tests Argon2 key generation with random salt
func TestArgon2Generate(t *testing.T) {
	config := DefaultArgon2Config()

	key, salt, err := Argon2Generate(testPassword, config)
	if err != nil {
		t.Fatalf("Argon2Generate failed: %v", err)
	}

	if len(key) != config.KeyLength {
		t.Errorf("Expected key length %d, got %d", config.KeyLength, len(key))
	}
	if len(salt) != config.SaltLength {
		t.Errorf("Expected salt length %d, got %d", config.SaltLength, len(salt))
	}

	// Test that different calls produce different salts
	_, salt2, err := Argon2Generate(testPassword, config)
	if err != nil {
		t.Fatalf("Second Argon2Generate failed: %v", err)
	}
	if bytes.Equal(salt, salt2) {
		t.Error("Different calls should produce different salts")
	}
}

// TestArgon2Verify tests Argon2 password verification
func TestArgon2Verify(t *testing.T) {
	config := DefaultArgon2Config()

	key, salt, err := Argon2Generate(testPassword, config)
	if err != nil {
		t.Fatalf("Argon2Generate failed: %v", err)
	}

	// Test correct password verification
	if !Argon2Verify(testPassword, key, salt, config) {
		t.Error("Should verify correct password")
	}

	// Test incorrect password verification
	if Argon2Verify("wrong_password", key, salt, config) {
		t.Error("Should not verify incorrect password")
	}
}

// TestConstantTimeCompare tests constant time comparison
func TestConstantTimeCompare(t *testing.T) {
	a := []byte("test")
	b := []byte("test")
	c := []byte("diff")
	d := []byte("test2") // different length

	if !constantTimeCompare(a, b) {
		t.Error("Should return true for equal slices")
	}

	if constantTimeCompare(a, c) {
		t.Error("Should return false for different slices")
	}

	if constantTimeCompare(a, d) {
		t.Error("Should return false for different length slices")
	}
}

// TestGenerateSalt tests salt generation
func TestGenerateSalt(t *testing.T) {
	length := 16
	salt, err := GenerateSalt(length)
	if err != nil {
		t.Fatalf("GenerateSalt failed: %v", err)
	}

	if len(salt) != length {
		t.Errorf("Expected salt length %d, got %d", length, len(salt))
	}

	// Test that different calls produce different salts
	salt2, err := GenerateSalt(length)
	if err != nil {
		t.Fatalf("Second GenerateSalt failed: %v", err)
	}
	if bytes.Equal(salt, salt2) {
		t.Error("Different calls should produce different salts")
	}
}

// TestKDFErrorConditions tests error conditions for KDF functions
func TestKDFErrorConditions(t *testing.T) {
	// Test PBKDF2Generate error conditions
	_, _, err := PBKDF2Generate("test", PBKDF2Config{SaltLength: 0, Iterations: 1000, KeyLength: 32})
	if err == nil {
		t.Error("Expected error for zero salt length")
	}

	_, _, err = PBKDF2Generate("test", PBKDF2Config{SaltLength: 16, Iterations: 0, KeyLength: 32})
	if err == nil {
		t.Error("Expected error for zero iterations")
	}

	_, _, err = PBKDF2Generate("test", PBKDF2Config{SaltLength: 16, Iterations: 1000, KeyLength: 0})
	if err == nil {
		t.Error("Expected error for zero key length")
	}

	// Test ScryptDerive error conditions
	config := ScryptConfig{N: 0, R: 8, P: 1, KeyLength: 32}
	_, err = ScryptDerive([]byte("test"), []byte("salt"), config)
	if err == nil {
		t.Error("Expected error for invalid N")
	}

	config = ScryptConfig{N: 15, R: 8, P: 1, KeyLength: 32} // Not a power of 2
	_, err = ScryptDerive([]byte("test"), []byte("salt"), config)
	if err == nil {
		t.Error("Expected error for N not being power of 2")
	}

	config = ScryptConfig{N: 16, R: 0, P: 1, KeyLength: 32}
	_, err = ScryptDerive([]byte("test"), []byte("salt"), config)
	if err == nil {
		t.Error("Expected error for zero R")
	}

	config = ScryptConfig{N: 16, R: 8, P: 0, KeyLength: 32}
	_, err = ScryptDerive([]byte("test"), []byte("salt"), config)
	if err == nil {
		t.Error("Expected error for zero P")
	}

	config = ScryptConfig{N: 16, R: 8, P: 1, KeyLength: 0}
	_, err = ScryptDerive([]byte("test"), []byte("salt"), config)
	if err == nil {
		t.Error("Expected error for zero key length")
	}

	// Test ScryptGenerate error conditions
	_, _, err = ScryptGenerate("test", ScryptConfig{SaltLength: 0, N: 16, R: 8, P: 1, KeyLength: 32})
	if err == nil {
		t.Error("Expected error for zero salt length")
	}

	// Test Argon2Generate error conditions
	_, _, err = Argon2Generate("test", Argon2Config{SaltLength: 0, Time: 1, Memory: 1024, Threads: 1, KeyLength: 32})
	if err == nil {
		t.Error("Expected error for zero salt length")
	}

	_, _, err = Argon2Generate("test", Argon2Config{SaltLength: 16, Time: 0, Memory: 1024, Threads: 1, KeyLength: 32})
	if err == nil {
		t.Error("Expected error for zero time")
	}

	_, _, err = Argon2Generate("test", Argon2Config{SaltLength: 16, Time: 1, Memory: 0, Threads: 1, KeyLength: 32})
	if err == nil {
		t.Error("Expected error for zero memory")
	}

	_, _, err = Argon2Generate("test", Argon2Config{SaltLength: 16, Time: 1, Memory: 1024, Threads: 0, KeyLength: 32})
	if err == nil {
		t.Error("Expected error for zero threads")
	}

	_, _, err = Argon2Generate("test", Argon2Config{SaltLength: 16, Time: 1, Memory: 1024, Threads: 1, KeyLength: 0})
	if err == nil {
		t.Error("Expected error for zero key length")
	}

	// Test GenerateSalt error conditions
	_, err = GenerateSalt(0)
	if err == nil {
		t.Error("Expected error for zero salt length")
	}

	_, err = GenerateSalt(-1)
	if err == nil {
		t.Error("Expected error for negative salt length")
	}
}

// TestKDFWithDifferentInputs tests KDF functions with various inputs
func TestKDFWithDifferentInputs(t *testing.T) {
	// Test with empty password
	emptyPassword := ""
	config := DefaultPBKDF2Config()

	key, _, err := PBKDF2Generate(emptyPassword, config)
	if err != nil {
		t.Errorf("PBKDF2Generate should work with empty password: %v", err)
	}
	if len(key) != config.KeyLength {
		t.Error("Should generate correct key length for empty password")
	}

	// Test with long password
	longPassword := string(make([]byte, 1000))
	key2, _, err := PBKDF2Generate(longPassword, config)
	if err != nil {
		t.Errorf("PBKDF2Generate should work with long password: %v", err)
	}
	if bytes.Equal(key, key2) {
		t.Error("Different passwords should produce different keys")
	}

	// Test with Unicode password
	unicodePassword := "测试密码🔒"
	key3, _, err := PBKDF2Generate(unicodePassword, config)
	if err != nil {
		t.Errorf("PBKDF2Generate should work with Unicode password: %v", err)
	}
	if bytes.Equal(key, key3) {
		t.Error("Different passwords should produce different keys")
	}
}
