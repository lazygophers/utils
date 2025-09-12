package cryptox

import (
	"crypto/cipher"
	"errors"
	"testing"
)

// FailingReader implements io.Reader but always fails
type FailingReader struct{}

func (fr FailingReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("simulated random reader failure")
}

// FailingCipherFunc simulates aes.NewCipher failure
func FailingCipherFunc(key []byte) (cipher.Block, error) {
	return nil, errors.New("simulated aes.NewCipher failure")
}

// FailingGCMFunc simulates cipher.NewGCM failure  
func FailingGCMFunc(block cipher.Block) (cipher.AEAD, error) {
	return nil, errors.New("simulated cipher.NewGCM failure")
}

// Test100PercentCoverage triggers all error paths using dependency injection
func Test100PercentCoverage(t *testing.T) {
	// Save original functions
	originalNewCipher := newCipherFunc
	originalNewGCM := newGCMFunc
	originalRandReader := randReader
	
	// Restore original functions after test
	defer func() {
		newCipherFunc = originalNewCipher
		newGCMFunc = originalNewGCM
		randReader = originalRandReader
	}()
	
	key := make([]byte, 32)
	plaintext := []byte("test plaintext")
	
	// Test 1: Trigger aes.NewCipher failure in all functions
	newCipherFunc = FailingCipherFunc
	newGCMFunc = originalNewGCM
	randReader = originalRandReader
	
	_, err := Encrypt(key, plaintext)
	if err == nil {
		t.Error("Expected aes.NewCipher error in Encrypt")
	}
	
	_, err = Decrypt(key, []byte("fake ciphertext that's long enough"))
	if err == nil {
		t.Error("Expected aes.NewCipher error in Decrypt")
	}
	
	_, err = EncryptECB(key, plaintext)
	if err == nil {
		t.Error("Expected aes.NewCipher error in EncryptECB")
	}
	
	_, err = DecryptECB(key, []byte("fake ciphertext"))
	if err == nil {
		t.Error("Expected aes.NewCipher error in DecryptECB")
	}
	
	_, err = EncryptCBC(key, plaintext)
	if err == nil {
		t.Error("Expected aes.NewCipher error in EncryptCBC")
	}
	
	_, err = DecryptCBC(key, []byte("fake ciphertext that's long enough"))
	if err == nil {
		t.Error("Expected aes.NewCipher error in DecryptCBC")
	}
	
	_, err = EncryptCFB(key, plaintext)
	if err == nil {
		t.Error("Expected aes.NewCipher error in EncryptCFB")
	}
	
	_, err = DecryptCFB(key, []byte("fake ciphertext that's long enough"))
	if err == nil {
		t.Error("Expected aes.NewCipher error in DecryptCFB")
	}
	
	_, err = EncryptCTR(key, plaintext)
	if err == nil {
		t.Error("Expected aes.NewCipher error in EncryptCTR")
	}
	
	_, err = DecryptCTR(key, []byte("fake ciphertext that's long enough"))
	if err == nil {
		t.Error("Expected aes.NewCipher error in DecryptCTR")
	}
	
	_, err = EncryptOFB(key, plaintext)
	if err == nil {
		t.Error("Expected aes.NewCipher error in EncryptOFB")
	}
	
	_, err = DecryptOFB(key, []byte("fake ciphertext that's long enough"))
	if err == nil {
		t.Error("Expected aes.NewCipher error in DecryptOFB")
	}
	
	// Test 2: Trigger cipher.NewGCM failure in GCM functions
	newCipherFunc = originalNewCipher
	newGCMFunc = FailingGCMFunc
	randReader = originalRandReader
	
	_, err = Encrypt(key, plaintext)
	if err == nil {
		t.Error("Expected cipher.NewGCM error in Encrypt")
	}
	
	_, err = Decrypt(key, []byte("fake ciphertext that's long enough"))
	if err == nil {
		t.Error("Expected cipher.NewGCM error in Decrypt")
	}
	
	// Test 3: Trigger rand.Reader failure in functions that use random IV
	newCipherFunc = originalNewCipher
	newGCMFunc = originalNewGCM
	randReader = FailingReader{}
	
	_, err = Encrypt(key, plaintext)
	if err == nil {
		t.Error("Expected rand.Reader error in Encrypt")
	}
	
	_, err = EncryptCBC(key, plaintext)
	if err == nil {
		t.Error("Expected rand.Reader error in EncryptCBC")
	}
	
	_, err = EncryptCFB(key, plaintext)
	if err == nil {
		t.Error("Expected rand.Reader error in EncryptCFB")
	}
	
	_, err = EncryptCTR(key, plaintext)
	if err == nil {
		t.Error("Expected rand.Reader error in EncryptCTR")
	}
	
	_, err = EncryptOFB(key, plaintext)
	if err == nil {
		t.Error("Expected rand.Reader error in EncryptOFB")
	}
}

// TestShortCiphertext tests ciphertext length validation
func TestShortCiphertext(t *testing.T) {
	key := make([]byte, 32)
	
	// Test all decrypt functions with short ciphertext (less than block size)
	shortCiphertext := make([]byte, 15) // Less than aes.BlockSize (16 bytes)
	
	_, err := DecryptCBC(key, shortCiphertext)
	if err == nil || err.Error() != "ciphertext too short" {
		t.Error("Expected 'ciphertext too short' error in DecryptCBC")
	}
	
	_, err = DecryptCFB(key, shortCiphertext)
	if err == nil || err.Error() != "ciphertext too short" {
		t.Error("Expected 'ciphertext too short' error in DecryptCFB")
	}
	
	_, err = DecryptCTR(key, shortCiphertext)
	if err == nil || err.Error() != "ciphertext too short" {
		t.Error("Expected 'ciphertext too short' error in DecryptCTR")
	}
	
	_, err = DecryptOFB(key, shortCiphertext)
	if err == nil || err.Error() != "ciphertext too short" {
		t.Error("Expected 'ciphertext too short' error in DecryptOFB")
	}
}