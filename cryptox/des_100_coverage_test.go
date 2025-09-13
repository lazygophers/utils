package cryptox

import (
	"crypto/cipher"
	"errors"
	"testing"
)

// Mock failures for DES dependency injection
type FailingDESReader struct{}

func (fr FailingDESReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("simulated random reader failure")
}

func FailingDESNewCipher(key []byte) (cipher.Block, error) {
	return nil, errors.New("simulated DES NewCipher failure")
}

func FailingDESNewTripleDES(key []byte) (cipher.Block, error) {
	return nil, errors.New("simulated 3DES NewTripleDES failure")
}

// TestDES100PercentCoverage triggers all error paths using dependency injection
func TestDES100PercentCoverage(t *testing.T) {
	// Save original functions
	originalDesNewCipher := desNewCipher
	originalDesNewTripleDES := desNewTripleDES
	originalDesRandReader := desRandReader

	// Restore original functions after test
	defer func() {
		desNewCipher = originalDesNewCipher
		desNewTripleDES = originalDesNewTripleDES
		desRandReader = originalDesRandReader
	}()

	key8 := make([]byte, 8)
	key24 := make([]byte, 24)
	plaintext := []byte("test plaintext")

	// Test 1: Trigger DES NewCipher failure in all DES functions
	desNewCipher = FailingDESNewCipher
	desNewTripleDES = originalDesNewTripleDES
	desRandReader = originalDesRandReader

	_, err := DESEncryptECB(key8, plaintext)
	if err == nil {
		t.Error("Expected DES NewCipher error in DESEncryptECB")
	}

	_, err = DESDecryptECB(key8, []byte("fake ciphertext"))
	if err == nil {
		t.Error("Expected DES NewCipher error in DESDecryptECB")
	}

	_, err = DESEncryptCBC(key8, plaintext)
	if err == nil {
		t.Error("Expected DES NewCipher error in DESEncryptCBC")
	}

	_, err = DESDecryptCBC(key8, []byte("fake ciphertext that's long enough"))
	if err == nil {
		t.Error("Expected DES NewCipher error in DESDecryptCBC")
	}

	// Test 2: Trigger 3DES NewTripleDES failure in all 3DES functions
	desNewCipher = originalDesNewCipher
	desNewTripleDES = FailingDESNewTripleDES
	desRandReader = originalDesRandReader

	_, err = TripleDESEncryptECB(key24, plaintext)
	if err == nil {
		t.Error("Expected 3DES NewTripleDES error in TripleDESEncryptECB")
	}

	_, err = TripleDESDecryptECB(key24, []byte("fake ciphertext"))
	if err == nil {
		t.Error("Expected 3DES NewTripleDES error in TripleDESDecryptECB")
	}

	_, err = TripleDESEncryptCBC(key24, plaintext)
	if err == nil {
		t.Error("Expected 3DES NewTripleDES error in TripleDESEncryptCBC")
	}

	_, err = TripleDESDecryptCBC(key24, []byte("fake ciphertext that's long enough"))
	if err == nil {
		t.Error("Expected 3DES NewTripleDES error in TripleDESDecryptCBC")
	}

	// Test with 24-byte key as well
	_, err = TripleDESEncryptECB(key24, plaintext)
	if err == nil {
		t.Error("Expected 3DES NewTripleDES error in TripleDESEncryptECB (24-byte key)")
	}

	_, err = TripleDESDecryptECB(key24, []byte("fake ciphertext"))
	if err == nil {
		t.Error("Expected 3DES NewTripleDES error in TripleDESDecryptECB (24-byte key)")
	}

	_, err = TripleDESEncryptCBC(key24, plaintext)
	if err == nil {
		t.Error("Expected 3DES NewTripleDES error in TripleDESEncryptCBC (24-byte key)")
	}

	_, err = TripleDESDecryptCBC(key24, []byte("fake ciphertext that's long enough"))
	if err == nil {
		t.Error("Expected 3DES NewTripleDES error in TripleDESDecryptCBC (24-byte key)")
	}

	// Test 3: Trigger rand.Reader failure in CBC encryption functions
	desNewCipher = originalDesNewCipher
	desNewTripleDES = originalDesNewTripleDES
	desRandReader = FailingDESReader{}

	_, err = DESEncryptCBC(key8, plaintext)
	if err == nil {
		t.Error("Expected rand.Reader error in DESEncryptCBC")
	}

	_, err = TripleDESEncryptCBC(key24, plaintext)
	if err == nil {
		t.Error("Expected rand.Reader error in TripleDESEncryptCBC (24-byte key)")
	}
}

// TestDESPaddingErrors tests PKCS#7 padding error scenarios
func TestDESPaddingErrors(t *testing.T) {
	key8 := make([]byte, 8)

	// Test with invalid padding that will cause unpadPKCS7 to fail
	// Create a ciphertext with valid length but invalid padding
	validCiphertext, err := DESEncryptECB(key8, []byte("testdata"))
	if err != nil {
		t.Fatalf("Failed to create valid ciphertext: %v", err)
	}

	// Corrupt the last byte to create invalid padding
	corruptedCiphertext := make([]byte, len(validCiphertext))
	copy(corruptedCiphertext, validCiphertext)
	corruptedCiphertext[len(corruptedCiphertext)-1] = 0xFF // Invalid padding byte

	// This should fail during unpadding
	_, err = DESDecryptECB(key8, corruptedCiphertext)
	if err == nil {
		t.Error("Expected padding error in DESDecryptECB")
	}

	// Test with 3DES as well
	key24DES := make([]byte, 24)
	valid3DESCiphertext, err := TripleDESEncryptECB(key24DES, []byte("testdata"))
	if err != nil {
		t.Fatalf("Failed to create valid 3DES ciphertext: %v", err)
	}

	corrupted3DESCiphertext := make([]byte, len(valid3DESCiphertext))
	copy(corrupted3DESCiphertext, valid3DESCiphertext)
	corrupted3DESCiphertext[len(corrupted3DESCiphertext)-1] = 0xFF

	_, err = TripleDESDecryptECB(key24DES, corrupted3DESCiphertext)
	if err == nil {
		t.Error("Expected padding error in TripleDESDecryptECB")
	}
}

// TestDESCBCPaddingErrors tests CBC mode padding errors
func TestDESCBCPaddingErrors(t *testing.T) {
	key8 := make([]byte, 8)
	key24CBC := make([]byte, 24)

	// Test with DES CBC
	validCiphertext, err := DESEncryptCBC(key8, []byte("testdata"))
	if err != nil {
		t.Fatalf("Failed to create valid DES CBC ciphertext: %v", err)
	}

	// Corrupt the last byte to create invalid padding
	corruptedCiphertext := make([]byte, len(validCiphertext))
	copy(corruptedCiphertext, validCiphertext)
	corruptedCiphertext[len(corruptedCiphertext)-1] = 0xFF

	_, err = DESDecryptCBC(key8, corruptedCiphertext)
	if err == nil {
		t.Error("Expected padding error in DESDecryptCBC")
	}

	// Test with 3DES CBC
	valid3DESCiphertext, err := TripleDESEncryptCBC(key24CBC, []byte("testdata"))
	if err != nil {
		t.Fatalf("Failed to create valid 3DES CBC ciphertext: %v", err)
	}

	corrupted3DESCiphertext := make([]byte, len(valid3DESCiphertext))
	copy(corrupted3DESCiphertext, valid3DESCiphertext)
	corrupted3DESCiphertext[len(corrupted3DESCiphertext)-1] = 0xFF

	_, err = TripleDESDecryptCBC(key24CBC, corrupted3DESCiphertext)
	if err == nil {
		t.Error("Expected padding error in TripleDESDecryptCBC")
	}
}
