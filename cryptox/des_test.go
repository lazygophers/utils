package cryptox

import (
	"bytes"
	"crypto/des"
	"testing"
)

// Test data
const (
	testDESMessage    = "Hello, DES encryption!"
	testTripleDESMessage = "Hello, 3DES encryption test message!"
)

var (
	desKey8   = []byte("12345678")        // 8 bytes for DES
	desKey24  = []byte("123456789012345678901234") // 24 bytes for 3DES
)

// TestDESEncryptDecryptECB tests DES ECB mode encryption and decryption
func TestDESEncryptDecryptECB(t *testing.T) {
	plaintext := []byte(testDESMessage)

	// Test encryption
	ciphertext, err := DESEncryptECB(desKey8, plaintext)
	if err != nil {
		t.Fatalf("DES ECB encryption failed: %v", err)
	}

	// Test decryption
	decrypted, err := DESDecryptECB(desKey8, ciphertext)
	if err != nil {
		t.Fatalf("DES ECB decryption failed: %v", err)
	}

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("DES ECB: plaintext mismatch.\nExpected: %s\nGot: %s", plaintext, decrypted)
	}
}

// TestDESEncryptDecryptCBC tests DES CBC mode encryption and decryption
func TestDESEncryptDecryptCBC(t *testing.T) {
	plaintext := []byte(testDESMessage)

	// Test encryption
	ciphertext, err := DESEncryptCBC(desKey8, plaintext)
	if err != nil {
		t.Fatalf("DES CBC encryption failed: %v", err)
	}

	// Test decryption
	decrypted, err := DESDecryptCBC(desKey8, ciphertext)
	if err != nil {
		t.Fatalf("DES CBC decryption failed: %v", err)
	}

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("DES CBC: plaintext mismatch.\nExpected: %s\nGot: %s", plaintext, decrypted)
	}
}

// TestTripleDESEncryptDecryptECB tests 3DES ECB mode encryption and decryption
func TestTripleDESEncryptDecryptECB(t *testing.T) {
	plaintext := []byte(testTripleDESMessage)

	// Test with 24-byte key
	t.Run("24-byte key", func(t *testing.T) {
		ciphertext, err := TripleDESEncryptECB(desKey24, plaintext)
		if err != nil {
			t.Fatalf("3DES ECB encryption failed: %v", err)
		}

		decrypted, err := TripleDESDecryptECB(desKey24, ciphertext)
		if err != nil {
			t.Fatalf("3DES ECB decryption failed: %v", err)
		}

		if !bytes.Equal(plaintext, decrypted) {
			t.Errorf("3DES ECB (24-byte): plaintext mismatch.\nExpected: %s\nGot: %s", plaintext, decrypted)
		}
	})
}

// TestTripleDESEncryptDecryptCBC tests 3DES CBC mode encryption and decryption
func TestTripleDESEncryptDecryptCBC(t *testing.T) {
	plaintext := []byte(testTripleDESMessage)

	// Test with 24-byte key
	t.Run("24-byte key", func(t *testing.T) {
		ciphertext, err := TripleDESEncryptCBC(desKey24, plaintext)
		if err != nil {
			t.Fatalf("3DES CBC encryption failed: %v", err)
		}

		decrypted, err := TripleDESDecryptCBC(desKey24, ciphertext)
		if err != nil {
			t.Fatalf("3DES CBC decryption failed: %v", err)
		}

		if !bytes.Equal(plaintext, decrypted) {
			t.Errorf("3DES CBC (24-byte): plaintext mismatch.\nExpected: %s\nGot: %s", plaintext, decrypted)
		}
	})
}

// TestDESInvalidKeyLength tests DES functions with invalid key lengths
func TestDESInvalidKeyLength(t *testing.T) {
	plaintext := []byte("test")
	invalidKey := []byte("invalid")

	// Test DES ECB with invalid key
	_, err := DESEncryptECB(invalidKey, plaintext)
	if err == nil || err.Error() != "invalid key length: must be 8 bytes for DES" {
		t.Error("Expected invalid key length error for DES ECB encryption")
	}

	_, err = DESDecryptECB(invalidKey, plaintext)
	if err == nil || err.Error() != "invalid key length: must be 8 bytes for DES" {
		t.Error("Expected invalid key length error for DES ECB decryption")
	}

	// Test DES CBC with invalid key
	_, err = DESEncryptCBC(invalidKey, plaintext)
	if err == nil || err.Error() != "invalid key length: must be 8 bytes for DES" {
		t.Error("Expected invalid key length error for DES CBC encryption")
	}

	_, err = DESDecryptCBC(invalidKey, plaintext)
	if err == nil || err.Error() != "invalid key length: must be 8 bytes for DES" {
		t.Error("Expected invalid key length error for DES CBC decryption")
	}
}

// TestTripleDESInvalidKeyLength tests 3DES functions with invalid key lengths
func TestTripleDESInvalidKeyLength(t *testing.T) {
	plaintext := []byte("test")
	invalidKey := []byte("invalid")

	// Test 3DES ECB with invalid key
	_, err := TripleDESEncryptECB(invalidKey, plaintext)
	if err == nil || err.Error() != "invalid key length: must be 24 bytes for 3DES" {
		t.Error("Expected invalid key length error for 3DES ECB encryption")
	}

	_, err = TripleDESDecryptECB(invalidKey, plaintext)
	if err == nil || err.Error() != "invalid key length: must be 24 bytes for 3DES" {
		t.Error("Expected invalid key length error for 3DES ECB decryption")
	}

	// Test 3DES CBC with invalid key
	_, err = TripleDESEncryptCBC(invalidKey, plaintext)
	if err == nil || err.Error() != "invalid key length: must be 24 bytes for 3DES" {
		t.Error("Expected invalid key length error for 3DES CBC encryption")
	}

	_, err = TripleDESDecryptCBC(invalidKey, plaintext)
	if err == nil || err.Error() != "invalid key length: must be 24 bytes for 3DES" {
		t.Error("Expected invalid key length error for 3DES CBC decryption")
	}
}

// TestDESCBCShortCiphertext tests DES CBC with short ciphertext
func TestDESCBCShortCiphertext(t *testing.T) {
	shortCiphertext := make([]byte, des.BlockSize-1) // Less than block size
	
	_, err := DESDecryptCBC(desKey8, shortCiphertext)
	if err == nil || err.Error() != "ciphertext too short" {
		t.Error("Expected 'ciphertext too short' error for DES CBC")
	}
}

// TestTripleDESCBCShortCiphertext tests 3DES CBC with short ciphertext
func TestTripleDESCBCShortCiphertext(t *testing.T) {
	shortCiphertext := make([]byte, des.BlockSize-1) // Less than block size
	
	_, err := TripleDESDecryptCBC(desKey24, shortCiphertext)
	if err == nil || err.Error() != "ciphertext too short" {
		t.Error("Expected 'ciphertext too short' error for 3DES CBC")
	}
}

// TestDESInvalidCiphertext tests DES functions with invalid ciphertext lengths
func TestDESInvalidCiphertext(t *testing.T) {
	// Create ciphertext that's not a multiple of block size
	invalidCiphertext := make([]byte, des.BlockSize+1) // 9 bytes (not multiple of 8)
	
	// Test DES ECB
	_, err := DESDecryptECB(desKey8, invalidCiphertext)
	if err == nil || err.Error() != "ciphertext is not a multiple of the block size" {
		t.Error("Expected block size error for DES ECB decryption")
	}

	// Test 3DES ECB
	_, err = TripleDESDecryptECB(desKey24, invalidCiphertext)
	if err == nil || err.Error() != "ciphertext is not a multiple of the block size" {
		t.Error("Expected block size error for 3DES ECB decryption")
	}
}

// TestDESCBCInvalidCiphertext tests DES CBC functions with invalid ciphertext lengths
func TestDESCBCInvalidCiphertext(t *testing.T) {
	// Create ciphertext with valid IV but invalid data length
	invalidCiphertext := make([]byte, des.BlockSize+des.BlockSize+1) // IV + 9 bytes data
	
	// Test DES CBC
	_, err := DESDecryptCBC(desKey8, invalidCiphertext)
	if err == nil || err.Error() != "ciphertext is not a multiple of the block size" {
		t.Error("Expected block size error for DES CBC decryption")
	}

	// Test 3DES CBC
	_, err = TripleDESDecryptCBC(desKey24, invalidCiphertext)
	if err == nil || err.Error() != "ciphertext is not a multiple of the block size" {
		t.Error("Expected block size error for 3DES CBC decryption")
	}
}

// TestDESEmptyPlaintext tests DES functions with empty plaintext
func TestDESEmptyPlaintext(t *testing.T) {
	emptyPlaintext := []byte("")

	// Test DES ECB
	ciphertext, err := DESEncryptECB(desKey8, emptyPlaintext)
	if err != nil {
		t.Fatalf("DES ECB encryption of empty plaintext failed: %v", err)
	}

	decrypted, err := DESDecryptECB(desKey8, ciphertext)
	if err != nil {
		t.Fatalf("DES ECB decryption failed: %v", err)
	}

	if !bytes.Equal(emptyPlaintext, decrypted) {
		t.Error("DES ECB: empty plaintext mismatch")
	}

	// Test DES CBC
	ciphertext, err = DESEncryptCBC(desKey8, emptyPlaintext)
	if err != nil {
		t.Fatalf("DES CBC encryption of empty plaintext failed: %v", err)
	}

	decrypted, err = DESDecryptCBC(desKey8, ciphertext)
	if err != nil {
		t.Fatalf("DES CBC decryption failed: %v", err)
	}

	if !bytes.Equal(emptyPlaintext, decrypted) {
		t.Error("DES CBC: empty plaintext mismatch")
	}
}

// TestTripleDESEmptyPlaintext tests 3DES functions with empty plaintext
func TestTripleDESEmptyPlaintext(t *testing.T) {
	emptyPlaintext := []byte("")

	// Test 3DES ECB
	ciphertext, err := TripleDESEncryptECB(desKey24, emptyPlaintext)
	if err != nil {
		t.Fatalf("3DES ECB encryption of empty plaintext failed: %v", err)
	}

	decrypted, err := TripleDESDecryptECB(desKey24, ciphertext)
	if err != nil {
		t.Fatalf("3DES ECB decryption failed: %v", err)
	}

	if !bytes.Equal(emptyPlaintext, decrypted) {
		t.Error("3DES ECB: empty plaintext mismatch")
	}

	// Test 3DES CBC
	ciphertext, err = TripleDESEncryptCBC(desKey24, emptyPlaintext)
	if err != nil {
		t.Fatalf("3DES CBC encryption of empty plaintext failed: %v", err)
	}

	decrypted, err = TripleDESDecryptCBC(desKey24, ciphertext)
	if err != nil {
		t.Fatalf("3DES CBC decryption failed: %v", err)
	}

	if !bytes.Equal(emptyPlaintext, decrypted) {
		t.Error("3DES CBC: empty plaintext mismatch")
	}
}

// TestDESLargePlaintext tests DES functions with large plaintext
func TestDESLargePlaintext(t *testing.T) {
	// Create a large plaintext (multiple blocks)
	largePlaintext := bytes.Repeat([]byte("This is a large plaintext for testing DES encryption with multiple blocks. "), 10)

	// Test DES ECB
	ciphertext, err := DESEncryptECB(desKey8, largePlaintext)
	if err != nil {
		t.Fatalf("DES ECB encryption of large plaintext failed: %v", err)
	}

	decrypted, err := DESDecryptECB(desKey8, ciphertext)
	if err != nil {
		t.Fatalf("DES ECB decryption failed: %v", err)
	}

	if !bytes.Equal(largePlaintext, decrypted) {
		t.Error("DES ECB: large plaintext mismatch")
	}

	// Test DES CBC
	ciphertext, err = DESEncryptCBC(desKey8, largePlaintext)
	if err != nil {
		t.Fatalf("DES CBC encryption of large plaintext failed: %v", err)
	}

	decrypted, err = DESDecryptCBC(desKey8, ciphertext)
	if err != nil {
		t.Fatalf("DES CBC decryption failed: %v", err)
	}

	if !bytes.Equal(largePlaintext, decrypted) {
		t.Error("DES CBC: large plaintext mismatch")
	}
}