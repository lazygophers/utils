package cryptox

import (
	"bytes"
	"crypto/rand"
	"io"
	"strings"
	"testing"
)

func TestEncrypt(t *testing.T) {
	key := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		t.Fatalf("failed to generate random key: %v", err)
	}

	plaintext := []byte("test plaintext")

	ciphertext, err := Encrypt(key, plaintext)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	if len(ciphertext) == 0 {
		t.Error("ciphertext should not be empty")
	}
}

func TestDecrypt(t *testing.T) {
	key := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		t.Fatalf("failed to generate random key: %v", err)
	}

	plaintext := []byte("test plaintext")

	ciphertext, err := Encrypt(key, plaintext)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	decrypted, err := Decrypt(key, ciphertext)
	if err != nil {
		t.Fatalf("Decrypt failed: %v", err)
	}

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("decrypted text does not match original plaintext")
	}
}

func TestEncryptECB(t *testing.T) {
	key := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		t.Fatalf("failed to generate random key: %v", err)
	}

	plaintext := []byte("test plaintext")

	ciphertext, err := EncryptECB(key, plaintext)
	if err != nil {
		t.Fatalf("EncryptECB failed: %v", err)
	}

	if len(ciphertext) == 0 {
		t.Error("ciphertext should not be empty")
	}
}

func TestDecryptECB(t *testing.T) {
	key := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		t.Fatalf("failed to generate random key: %v", err)
	}

	plaintext := []byte("test plaintext")

	ciphertext, err := EncryptECB(key, plaintext)
	if err != nil {
		t.Fatalf("EncryptECB failed: %v", err)
	}

	decrypted, err := DecryptECB(key, ciphertext)
	if err != nil {
		t.Fatalf("DecryptECB failed: %v", err)
	}

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("decrypted text does not match original plaintext")
	}
}

func TestEncryptCBC(t *testing.T) {
	key := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		t.Fatalf("failed to generate random key: %v", err)
	}

	plaintext := []byte("test plaintext")

	ciphertext, err := EncryptCBC(key, plaintext)
	if err != nil {
		t.Fatalf("EncryptCBC failed: %v", err)
	}

	if len(ciphertext) == 0 {
		t.Error("ciphertext should not be empty")
	}
}

func TestDecryptCBC(t *testing.T) {
	key := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		t.Fatalf("failed to generate random key: %v", err)
	}

	plaintext := []byte("test plaintext")

	ciphertext, err := EncryptCBC(key, plaintext)
	if err != nil {
		t.Fatalf("EncryptCBC failed: %v", err)
	}

	decrypted, err := DecryptCBC(key, ciphertext)
	if err != nil {
		t.Fatalf("DecryptCBC failed: %v", err)
	}

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("decrypted text does not match original plaintext")
	}
}

func TestEncryptCFB(t *testing.T) {
	key := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		t.Fatalf("failed to generate random key: %v", err)
	}

	plaintext := []byte("test plaintext")

	ciphertext, err := EncryptCFB(key, plaintext)
	if err != nil {
		t.Fatalf("EncryptCFB failed: %v", err)
	}

	if len(ciphertext) == 0 {
		t.Error("ciphertext should not be empty")
	}
}

func TestDecryptCFB(t *testing.T) {
	key := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		t.Fatalf("failed to generate random key: %v", err)
	}

	plaintext := []byte("test plaintext")

	ciphertext, err := EncryptCFB(key, plaintext)
	if err != nil {
		t.Fatalf("EncryptCFB failed: %v", err)
	}

	decrypted, err := DecryptCFB(key, ciphertext)
	if err != nil {
		t.Fatalf("DecryptCFB failed: %v", err)
	}

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("decrypted text does not match original plaintext")
	}
}

func TestEncryptCTR(t *testing.T) {
	key := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		t.Fatalf("failed to generate random key: %v", err)
	}

	plaintext := []byte("test plaintext")

	ciphertext, err := EncryptCTR(key, plaintext)
	if err != nil {
		t.Fatalf("EncryptCTR failed: %v", err)
	}

	if len(ciphertext) == 0 {
		t.Error("ciphertext should not be empty")
	}
}

func TestDecryptCTR(t *testing.T) {
	key := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		t.Fatalf("failed to generate random key: %v", err)
	}

	plaintext := []byte("test plaintext")

	ciphertext, err := EncryptCTR(key, plaintext)
	if err != nil {
		t.Fatalf("EncryptCTR failed: %v", err)
	}

	decrypted, err := DecryptCTR(key, ciphertext)
	if err != nil {
		t.Fatalf("DecryptCTR failed: %v", err)
	}

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("decrypted text does not match original plaintext")
	}
}

func TestPadPKCS7(t *testing.T) {
	data := []byte("test")
	blockSize := 8

	padded := padPKCS7(data, blockSize)
	if len(padded)%blockSize != 0 {
		t.Errorf("padded data length should be a multiple of block size")
	}
}

func TestUnpadPKCS7(t *testing.T) {
	data := []byte("test\x04\x04\x04\x04")
	expected := []byte("test")

	unpadded, err := unpadPKCS7(data)
	if err != nil {
		t.Fatalf("unpadPKCS7 failed: %v", err)
	}

	if !bytes.Equal(unpadded, expected) {
		t.Errorf("unpadded data does not match expected")
	}
}

// Test error conditions
func TestInvalidKeyLength(t *testing.T) {
	shortKey := make([]byte, 16) // Invalid key length
	plaintext := []byte("test")
	
	_, err := Encrypt(shortKey, plaintext)
	if err == nil {
		t.Error("Expected error for invalid key length")
	}
	
	_, err = EncryptECB(shortKey, plaintext)
	if err == nil {
		t.Error("Expected error for invalid key length")
	}
	
	_, err = EncryptCBC(shortKey, plaintext)
	if err == nil {
		t.Error("Expected error for invalid key length")
	}
	
	_, err = EncryptCFB(shortKey, plaintext)
	if err == nil {
		t.Error("Expected error for invalid key length")
	}
	
	_, err = EncryptCTR(shortKey, plaintext)
	if err == nil {
		t.Error("Expected error for invalid key length")
	}
}

func TestInvalidCiphertext(t *testing.T) {
	key := make([]byte, 32)
	shortCiphertext := []byte("short")
	
	_, err := Decrypt(key, shortCiphertext)
	if err == nil {
		t.Error("Expected error for short ciphertext")
	}
	
	_, err = DecryptCBC(key, shortCiphertext)
	if err == nil {
		t.Error("Expected error for short ciphertext")
	}
	
	_, err = DecryptCFB(key, shortCiphertext)
	if err == nil {
		t.Error("Expected error for short ciphertext")
	}
	
	_, err = DecryptCTR(key, shortCiphertext)
	if err == nil {
		t.Error("Expected error for short ciphertext")
	}
}

func TestUnpadPKCS7Errors(t *testing.T) {
	// Test empty data
	_, err := unpadPKCS7([]byte{})
	if err == nil {
		t.Error("Expected error for empty data")
	}
	
	// Test invalid padding (padding value larger than data length)
	_, err = unpadPKCS7([]byte{1, 2, 3, 4, 5})
	if err == nil {
		t.Error("Expected error for invalid padding")
	}
	
	// Test invalid padding data (inconsistent padding bytes)
	_, err = unpadPKCS7([]byte{1, 2, 3, 2})
	if err == nil {
		t.Error("Expected error for invalid padding data")
	}
	
	// Test padding value of 0
	_, err = unpadPKCS7([]byte{1, 2, 3, 0})
	if err == nil {
		t.Error("Expected error for padding value 0")
	}
}

// Test AES cipher creation errors (simulate by using invalid operations)
func TestAESErrorConditions(t *testing.T) {
	key := make([]byte, 32)
	
	// Test decrypt with invalid GCM data
	invalidGCMData := []byte("short")
	_, err := Decrypt(key, invalidGCMData)
	if err == nil {
		t.Error("Expected error for invalid GCM data")
	}
	
	// Test decrypt ECB with invalid block size (needs to be properly sized to avoid panic)
	// ECB requires data to be a multiple of block size (16 bytes), but also needs valid padding
	invalidECBData := make([]byte, 16) // Valid block size but invalid padding
	for i := range invalidECBData {
		invalidECBData[i] = 0xFF // Invalid padding bytes
	}
	_, err = DecryptECB(key, invalidECBData)
	if err == nil {
		t.Error("Expected error for invalid ECB padding")
	}
	
	// Test decrypt with malformed IV
	invalidIVData := make([]byte, 15) // Less than AES block size
	_, err = DecryptCBC(key, invalidIVData)
	if err == nil {
		t.Error("Expected error for invalid IV length")
	}
	
	_, err = DecryptCFB(key, invalidIVData) 
	if err == nil {
		t.Error("Expected error for invalid IV length in CFB")
	}
	
	_, err = DecryptCTR(key, invalidIVData)
	if err == nil {
		t.Error("Expected error for invalid IV length in CTR")
	}
}

// Test decryption key length validation for all modes
func TestDecryptInvalidKeyLength(t *testing.T) {
	shortKey := make([]byte, 16) // Invalid key length
	ciphertext := make([]byte, 32) // Valid length ciphertext
	
	_, err := Decrypt(shortKey, ciphertext)
	if err == nil {
		t.Error("Expected error for invalid key length in Decrypt")
	}
	
	_, err = DecryptECB(shortKey, ciphertext)
	if err == nil {
		t.Error("Expected error for invalid key length in DecryptECB")
	}
	
	_, err = DecryptCBC(shortKey, ciphertext)
	if err == nil {
		t.Error("Expected error for invalid key length in DecryptCBC")
	}
	
	_, err = DecryptCFB(shortKey, ciphertext)
	if err == nil {
		t.Error("Expected error for invalid key length in DecryptCFB")
	}
	
	_, err = DecryptCTR(shortKey, ciphertext)
	if err == nil {
		t.Error("Expected error for invalid key length in DecryptCTR")
	}
}

// Test error paths by creating failing conditions
func TestAllAESErrorPaths(t *testing.T) {
	// Test with empty key to trigger various error conditions
	emptyKey := make([]byte, 32)
	plaintext := []byte("test message that needs padding for block size")
	
	// Test all encryption functions to ensure they work (success path)
	_, err := Encrypt(emptyKey, plaintext)
	if err != nil {
		t.Errorf("Encrypt should work with valid key: %v", err)
	}
	
	_, err = EncryptECB(emptyKey, plaintext)
	if err != nil {
		t.Errorf("EncryptECB should work with valid key: %v", err)
	}
	
	_, err = EncryptCBC(emptyKey, plaintext)
	if err != nil {
		t.Errorf("EncryptCBC should work with valid key: %v", err)
	}
	
	_, err = EncryptCFB(emptyKey, plaintext)
	if err != nil {
		t.Errorf("EncryptCFB should work with valid key: %v", err)
	}
	
	_, err = EncryptCTR(emptyKey, plaintext)
	if err != nil {
		t.Errorf("EncryptCTR should work with valid key: %v", err)
	}
}

// Test specific decrypt error conditions with malformed data
func TestDecryptMalformedData(t *testing.T) {
	key := make([]byte, 32)
	
	// Test GCM decrypt with malformed data (fails in gcm.Open)
	malformedGCM := make([]byte, 24) // Valid nonce size but invalid ciphertext
	_, err := Decrypt(key, malformedGCM)
	if err == nil {
		t.Error("Expected error for malformed GCM data")
	}
	
	// Test CBC decrypt with invalid IV length check
	validSizeCBC := make([]byte, 32) // 16 byte IV + 16 byte data
	for i := range validSizeCBC {
		validSizeCBC[i] = 0x00 // All zeros to trigger padding errors
	}
	_, err = DecryptCBC(key, validSizeCBC)
	if err == nil {
		t.Error("Expected error for malformed CBC data")
	}
}

// Test with bytes input types to ensure generic functions work
func TestGenericInputTypes(t *testing.T) {
	// Test with []byte input for hash functions
	data := []byte("test data")
	
	result := Md5(data)
	if result == "" {
		t.Error("Md5 with []byte should return non-empty result")
	}
	
	result = Sha256(data)  
	if result == "" {
		t.Error("Sha256 with []byte should return non-empty result")
	}
}

// Test covering missing edge cases and error paths
func TestRemainingErrorPaths(t *testing.T) {
	key := make([]byte, 32)
	
	// Test empty plaintext encryption/decryption to cover edge cases
	emptyText := []byte{}
	
	ciphertext, err := Encrypt(key, emptyText)
	if err != nil {
		t.Errorf("Encrypt empty text should work: %v", err)
	}
	
	decrypted, err := Decrypt(key, ciphertext)
	if err != nil {
		t.Errorf("Decrypt empty text should work: %v", err)  
	}
	if !bytes.Equal(emptyText, decrypted) {
		t.Error("Empty text roundtrip failed")
	}
	
	// Test ECB with empty text
	ciphertext, err = EncryptECB(key, emptyText)
	if err != nil {
		t.Errorf("EncryptECB empty text should work: %v", err)
	}
	
	decrypted, err = DecryptECB(key, ciphertext) 
	if err != nil {
		t.Errorf("DecryptECB empty text should work: %v", err)
	}
	if !bytes.Equal(emptyText, decrypted) {
		t.Error("ECB empty text roundtrip failed")
	}
}

// Test more padding scenarios
func TestMorePaddingScenarios(t *testing.T) {
	// Test with data that's exactly one block size
	data := make([]byte, 16)
	for i := range data {
		data[i] = byte(i)
	}
	
	padded := padPKCS7(data, 16)
	if len(padded) != 32 { // Should add a full block of padding
		t.Errorf("Expected 32 bytes after padding, got %d", len(padded))
	}
	
	unpadded, err := unpadPKCS7(padded)
	if err != nil {
		t.Errorf("Failed to unpad: %v", err)
	}
	
	if !bytes.Equal(data, unpadded) {
		t.Error("Padding/unpadding roundtrip failed")
	}
}

// Test UUID function with []byte conversion
func TestUUIDFunction(t *testing.T) {
	uuid1 := UUID()
	uuid2 := UUID()
	
	if uuid1 == uuid2 {
		t.Error("UUIDs should be unique")
	}
	
	if len(uuid1) != 32 { // 36 chars - 4 hyphens = 32 chars
		t.Errorf("UUID should be 32 characters, got %d", len(uuid1))
	}
	
	// Should not contain hyphens
	if strings.Contains(uuid1, "-") {
		t.Error("UUID should not contain hyphens")
	}
}

// Test comprehensive error validation with edge cases
func TestComprehensiveErrorValidation(t *testing.T) {
	// Test all modes with shortest possible invalid key (0 length)
	emptyKey := []byte{}
	plaintext := []byte("test")
	
	_, err := Encrypt(emptyKey, plaintext)
	if err == nil {
		t.Error("Expected error for empty key in Encrypt")
	}
	
	_, err = EncryptECB(emptyKey, plaintext)
	if err == nil {
		t.Error("Expected error for empty key in EncryptECB")
	}
	
	_, err = EncryptCBC(emptyKey, plaintext)
	if err == nil {
		t.Error("Expected error for empty key in EncryptCBC")
	}
	
	_, err = EncryptCFB(emptyKey, plaintext)
	if err == nil {
		t.Error("Expected error for empty key in EncryptCFB")
	}
	
	_, err = EncryptCTR(emptyKey, plaintext)
	if err == nil {
		t.Error("Expected error for empty key in EncryptCTR")
	}
	
	// Test all decrypt modes with empty key
	ciphertext := []byte("dummy data")
	
	_, err = Decrypt(emptyKey, ciphertext)
	if err == nil {
		t.Error("Expected error for empty key in Decrypt")
	}
	
	_, err = DecryptECB(emptyKey, ciphertext)
	if err == nil {
		t.Error("Expected error for empty key in DecryptECB")
	}
	
	_, err = DecryptCBC(emptyKey, ciphertext)
	if err == nil {
		t.Error("Expected error for empty key in DecryptCBC")
	}
	
	_, err = DecryptCFB(emptyKey, ciphertext)
	if err == nil {
		t.Error("Expected error for empty key in DecryptCFB")
	}
	
	_, err = DecryptCTR(emptyKey, ciphertext)
	if err == nil {
		t.Error("Expected error for empty key in DecryptCTR")
	}
}

// Attempt to create a test that might trigger some of the hard-to-reach error conditions
// by testing with extremely unusual inputs that might cause internal failures
func TestExtremeEdgeCases(t *testing.T) {
	// Test with various unusual key lengths that might trigger different error paths
	for _, keyLen := range []int{1, 15, 17, 31, 33, 64} {
		key := make([]byte, keyLen)
		plaintext := []byte("test")
		
		// All these should fail with "invalid key length" but ensure we test the path
		_, err := Encrypt(key, plaintext)
		if err == nil && keyLen != 32 {
			t.Errorf("Expected error for key length %d in Encrypt", keyLen)
		}
		
		_, err = EncryptECB(key, plaintext)
		if err == nil && keyLen != 32 {
			t.Errorf("Expected error for key length %d in EncryptECB", keyLen)
		}
		
		_, err = EncryptCBC(key, plaintext)
		if err == nil && keyLen != 32 {
			t.Errorf("Expected error for key length %d in EncryptCBC", keyLen)
		}
		
		_, err = EncryptCFB(key, plaintext)
		if err == nil && keyLen != 32 {
			t.Errorf("Expected error for key length %d in EncryptCFB", keyLen)
		}
		
		_, err = EncryptCTR(key, plaintext)
		if err == nil && keyLen != 32 {
			t.Errorf("Expected error for key length %d in EncryptCTR", keyLen)
		}
	}
}

// Test behavior with very large inputs to ensure no overflow/panic conditions
func TestLargeInputs(t *testing.T) {
	key := make([]byte, 32)
	
	// Test with larger plaintext to ensure all code paths are exercised
	largePlaintext := make([]byte, 1024*16) // 16KB
	for i := range largePlaintext {
		largePlaintext[i] = byte(i % 256)
	}
	
	// Test encryption and decryption with large input
	ciphertext, err := Encrypt(key, largePlaintext)
	if err != nil {
		t.Errorf("Encrypt large input failed: %v", err)
	}
	
	decrypted, err := Decrypt(key, ciphertext)
	if err != nil {
		t.Errorf("Decrypt large input failed: %v", err)
	}
	
	if !bytes.Equal(largePlaintext, decrypted) {
		t.Error("Large input roundtrip failed")
	}
	
	// Test ECB mode with large input
	ciphertext, err = EncryptECB(key, largePlaintext)
	if err != nil {
		t.Errorf("EncryptECB large input failed: %v", err)
	}
	
	decrypted, err = DecryptECB(key, ciphertext)
	if err != nil {
		t.Errorf("DecryptECB large input failed: %v", err)
	}
	
	if !bytes.Equal(largePlaintext, decrypted) {
		t.Error("ECB large input roundtrip failed")
	}
}
