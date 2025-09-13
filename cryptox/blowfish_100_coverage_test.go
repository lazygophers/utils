package cryptox

import (
	"errors"
	"testing"

	"golang.org/x/crypto/blowfish"
)

// Mock failures for Blowfish dependency injection
type FailingBlowfishReader struct{}

func (fr FailingBlowfishReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("simulated random reader failure")
}

func FailingBlowfishNewCipher(key []byte) (*blowfish.Cipher, error) {
	return nil, errors.New("simulated Blowfish NewCipher failure")
}

// TestBlowfish100PercentCoverage triggers all error paths using dependency injection
func TestBlowfish100PercentCoverage(t *testing.T) {
	// Save original functions
	originalBlowfishNewCipher := blowfishNewCipher
	originalBlowfishRandReader := blowfishRandReader

	// Restore original functions after test
	defer func() {
		blowfishNewCipher = originalBlowfishNewCipher
		blowfishRandReader = originalBlowfishRandReader
	}()

	validKey := make([]byte, 8)
	plaintext := []byte("test plaintext")

	// Test 1: Trigger Blowfish NewCipher failure in all functions
	blowfishNewCipher = FailingBlowfishNewCipher
	blowfishRandReader = originalBlowfishRandReader

	_, err := BlowfishEncryptECB(validKey, plaintext)
	if err == nil {
		t.Error("Expected Blowfish NewCipher error in BlowfishEncryptECB")
	}

	_, err = BlowfishDecryptECB(validKey, []byte("fake ciphertext"))
	if err == nil {
		t.Error("Expected Blowfish NewCipher error in BlowfishDecryptECB")
	}

	_, err = BlowfishEncryptCBC(validKey, plaintext)
	if err == nil {
		t.Error("Expected Blowfish NewCipher error in BlowfishEncryptCBC")
	}

	_, err = BlowfishDecryptCBC(validKey, []byte("fake ciphertext that's long enough"))
	if err == nil {
		t.Error("Expected Blowfish NewCipher error in BlowfishDecryptCBC")
	}

	_, err = BlowfishEncryptCFB(validKey, plaintext)
	if err == nil {
		t.Error("Expected Blowfish NewCipher error in BlowfishEncryptCFB")
	}

	_, err = BlowfishDecryptCFB(validKey, []byte("fake ciphertext that's long enough"))
	if err == nil {
		t.Error("Expected Blowfish NewCipher error in BlowfishDecryptCFB")
	}

	_, err = BlowfishEncryptOFB(validKey, plaintext)
	if err == nil {
		t.Error("Expected Blowfish NewCipher error in BlowfishEncryptOFB")
	}

	_, err = BlowfishDecryptOFB(validKey, []byte("fake ciphertext that's long enough"))
	if err == nil {
		t.Error("Expected Blowfish NewCipher error in BlowfishDecryptOFB")
	}

	// Test 2: Trigger rand.Reader failure in functions that use random IV
	blowfishNewCipher = originalBlowfishNewCipher
	blowfishRandReader = FailingBlowfishReader{}

	_, err = BlowfishEncryptCBC(validKey, plaintext)
	if err == nil {
		t.Error("Expected rand.Reader error in BlowfishEncryptCBC")
	}

	_, err = BlowfishEncryptCFB(validKey, plaintext)
	if err == nil {
		t.Error("Expected rand.Reader error in BlowfishEncryptCFB")
	}

	_, err = BlowfishEncryptOFB(validKey, plaintext)
	if err == nil {
		t.Error("Expected rand.Reader error in BlowfishEncryptOFB")
	}
}

// TestBlowfishPaddingErrors tests PKCS#7 padding error scenarios
func TestBlowfishPaddingErrors(t *testing.T) {
	validKey := make([]byte, 8)

	// Test with invalid padding that will cause unpadPKCS7 to fail
	// Create a ciphertext with valid length but invalid padding
	validCiphertext, err := BlowfishEncryptECB(validKey, []byte("testdata"))
	if err != nil {
		t.Fatalf("Failed to create valid ciphertext: %v", err)
	}

	// Corrupt the last byte to create invalid padding
	corruptedCiphertext := make([]byte, len(validCiphertext))
	copy(corruptedCiphertext, validCiphertext)
	corruptedCiphertext[len(corruptedCiphertext)-1] = 0xFF // Invalid padding byte

	// This should fail during unpadding
	_, err = BlowfishDecryptECB(validKey, corruptedCiphertext)
	if err == nil {
		t.Error("Expected padding error in BlowfishDecryptECB")
	}
}

// TestBlowfishCBCPaddingErrors tests CBC mode padding errors
func TestBlowfishCBCPaddingErrors(t *testing.T) {
	validKey := make([]byte, 8)

	// Test with Blowfish CBC
	validCiphertext, err := BlowfishEncryptCBC(validKey, []byte("testdata"))
	if err != nil {
		t.Fatalf("Failed to create valid Blowfish CBC ciphertext: %v", err)
	}

	// Corrupt the last byte to create invalid padding
	corruptedCiphertext := make([]byte, len(validCiphertext))
	copy(corruptedCiphertext, validCiphertext)
	corruptedCiphertext[len(corruptedCiphertext)-1] = 0xFF

	_, err = BlowfishDecryptCBC(validKey, corruptedCiphertext)
	if err == nil {
		t.Error("Expected padding error in BlowfishDecryptCBC")
	}
}
