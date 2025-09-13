package cryptox

import (
	"bytes"
	"testing"

	"golang.org/x/crypto/blowfish"
)

// Test data
const (
	testBlowfishMessage = "Hello, Blowfish encryption test message!"
)

var (
	blowfishKey8  = []byte("12345678")                                                 // 8 bytes
	blowfishKey16 = []byte("1234567890123456")                                         // 16 bytes
	blowfishKey32 = []byte("12345678901234567890123456789012")                         // 32 bytes
	blowfishKey56 = []byte("12345678901234567890123456789012345678901234567890123456") // 56 bytes
)

// TestBlowfishEncryptDecryptECB tests Blowfish ECB mode encryption and decryption
func TestBlowfishEncryptDecryptECB(t *testing.T) {
	plaintext := []byte(testBlowfishMessage)

	testCases := []struct {
		name string
		key  []byte
	}{
		{"8-byte key", blowfishKey8},
		{"16-byte key", blowfishKey16},
		{"32-byte key", blowfishKey32},
		{"56-byte key", blowfishKey56},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test encryption
			ciphertext, err := BlowfishEncryptECB(tc.key, plaintext)
			if err != nil {
				t.Fatalf("Blowfish ECB encryption failed: %v", err)
			}

			// Test decryption
			decrypted, err := BlowfishDecryptECB(tc.key, ciphertext)
			if err != nil {
				t.Fatalf("Blowfish ECB decryption failed: %v", err)
			}

			if !bytes.Equal(plaintext, decrypted) {
				t.Errorf("Blowfish ECB (%s): plaintext mismatch.\nExpected: %s\nGot: %s", tc.name, plaintext, decrypted)
			}
		})
	}
}

// TestBlowfishEncryptDecryptCBC tests Blowfish CBC mode encryption and decryption
func TestBlowfishEncryptDecryptCBC(t *testing.T) {
	plaintext := []byte(testBlowfishMessage)

	testCases := []struct {
		name string
		key  []byte
	}{
		{"8-byte key", blowfishKey8},
		{"16-byte key", blowfishKey16},
		{"32-byte key", blowfishKey32},
		{"56-byte key", blowfishKey56},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test encryption
			ciphertext, err := BlowfishEncryptCBC(tc.key, plaintext)
			if err != nil {
				t.Fatalf("Blowfish CBC encryption failed: %v", err)
			}

			// Test decryption
			decrypted, err := BlowfishDecryptCBC(tc.key, ciphertext)
			if err != nil {
				t.Fatalf("Blowfish CBC decryption failed: %v", err)
			}

			if !bytes.Equal(plaintext, decrypted) {
				t.Errorf("Blowfish CBC (%s): plaintext mismatch.\nExpected: %s\nGot: %s", tc.name, plaintext, decrypted)
			}
		})
	}
}

// TestBlowfishEncryptDecryptCFB tests Blowfish CFB mode encryption and decryption
func TestBlowfishEncryptDecryptCFB(t *testing.T) {
	plaintext := []byte(testBlowfishMessage)

	testCases := []struct {
		name string
		key  []byte
	}{
		{"8-byte key", blowfishKey8},
		{"16-byte key", blowfishKey16},
		{"32-byte key", blowfishKey32},
		{"56-byte key", blowfishKey56},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test encryption
			ciphertext, err := BlowfishEncryptCFB(tc.key, plaintext)
			if err != nil {
				t.Fatalf("Blowfish CFB encryption failed: %v", err)
			}

			// Test decryption
			decrypted, err := BlowfishDecryptCFB(tc.key, ciphertext)
			if err != nil {
				t.Fatalf("Blowfish CFB decryption failed: %v", err)
			}

			if !bytes.Equal(plaintext, decrypted) {
				t.Errorf("Blowfish CFB (%s): plaintext mismatch.\nExpected: %s\nGot: %s", tc.name, plaintext, decrypted)
			}
		})
	}
}

// TestBlowfishEncryptDecryptOFB tests Blowfish OFB mode encryption and decryption
func TestBlowfishEncryptDecryptOFB(t *testing.T) {
	plaintext := []byte(testBlowfishMessage)

	testCases := []struct {
		name string
		key  []byte
	}{
		{"8-byte key", blowfishKey8},
		{"16-byte key", blowfishKey16},
		{"32-byte key", blowfishKey32},
		{"56-byte key", blowfishKey56},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test encryption
			ciphertext, err := BlowfishEncryptOFB(tc.key, plaintext)
			if err != nil {
				t.Fatalf("Blowfish OFB encryption failed: %v", err)
			}

			// Test decryption
			decrypted, err := BlowfishDecryptOFB(tc.key, ciphertext)
			if err != nil {
				t.Fatalf("Blowfish OFB decryption failed: %v", err)
			}

			if !bytes.Equal(plaintext, decrypted) {
				t.Errorf("Blowfish OFB (%s): plaintext mismatch.\nExpected: %s\nGot: %s", tc.name, plaintext, decrypted)
			}
		})
	}
}

// TestBlowfishInvalidKeyLength tests Blowfish functions with invalid key lengths
func TestBlowfishInvalidKeyLength(t *testing.T) {
	plaintext := []byte("test")

	// Test with empty key
	emptyKey := []byte("")
	longKey := make([]byte, 57) // 57 bytes, too long

	testFunctions := []struct {
		name    string
		encFunc func([]byte, []byte) ([]byte, error)
		decFunc func([]byte, []byte) ([]byte, error)
	}{
		{"ECB", BlowfishEncryptECB, BlowfishDecryptECB},
		{"CBC", BlowfishEncryptCBC, BlowfishDecryptCBC},
		{"CFB", BlowfishEncryptCFB, BlowfishDecryptCFB},
		{"OFB", BlowfishEncryptOFB, BlowfishDecryptOFB},
	}

	for _, tf := range testFunctions {
		t.Run(tf.name, func(t *testing.T) {
			// Test empty key
			_, err := tf.encFunc(emptyKey, plaintext)
			if err == nil || err.Error() != "invalid key length: must be between 1 and 56 bytes for Blowfish" {
				t.Error("Expected invalid key length error for empty key in encryption")
			}

			_, err = tf.decFunc(emptyKey, plaintext)
			if err == nil || err.Error() != "invalid key length: must be between 1 and 56 bytes for Blowfish" {
				t.Error("Expected invalid key length error for empty key in decryption")
			}

			// Test long key
			_, err = tf.encFunc(longKey, plaintext)
			if err == nil || err.Error() != "invalid key length: must be between 1 and 56 bytes for Blowfish" {
				t.Error("Expected invalid key length error for long key in encryption")
			}

			_, err = tf.decFunc(longKey, plaintext)
			if err == nil || err.Error() != "invalid key length: must be between 1 and 56 bytes for Blowfish" {
				t.Error("Expected invalid key length error for long key in decryption")
			}
		})
	}
}

// TestBlowfishShortCiphertext tests Blowfish with short ciphertext
func TestBlowfishShortCiphertext(t *testing.T) {
	shortCiphertext := make([]byte, blowfish.BlockSize-1) // Less than block size
	key := blowfishKey8

	testFunctions := []struct {
		name    string
		decFunc func([]byte, []byte) ([]byte, error)
	}{
		{"CBC", BlowfishDecryptCBC},
		{"CFB", BlowfishDecryptCFB},
		{"OFB", BlowfishDecryptOFB},
	}

	for _, tf := range testFunctions {
		t.Run(tf.name, func(t *testing.T) {
			_, err := tf.decFunc(key, shortCiphertext)
			if err == nil || err.Error() != "ciphertext too short" {
				t.Error("Expected 'ciphertext too short' error")
			}
		})
	}
}

// TestBlowfishInvalidCiphertext tests Blowfish functions with invalid ciphertext lengths
func TestBlowfishInvalidCiphertext(t *testing.T) {
	// Create ciphertext that's not a multiple of block size
	invalidCiphertext := make([]byte, blowfish.BlockSize+1) // 9 bytes (not multiple of 8)
	key := blowfishKey8

	// Test ECB decryption
	_, err := BlowfishDecryptECB(key, invalidCiphertext)
	if err == nil || err.Error() != "ciphertext is not a multiple of the block size" {
		t.Error("Expected block size error for Blowfish ECB decryption")
	}
}

// TestBlowfishCBCInvalidCiphertext tests Blowfish CBC functions with invalid ciphertext lengths
func TestBlowfishCBCInvalidCiphertext(t *testing.T) {
	// Create ciphertext with valid IV but invalid data length
	invalidCiphertext := make([]byte, blowfish.BlockSize+blowfish.BlockSize+1) // IV + 9 bytes data
	key := blowfishKey8

	// Test CBC decryption
	_, err := BlowfishDecryptCBC(key, invalidCiphertext)
	if err == nil || err.Error() != "ciphertext is not a multiple of the block size" {
		t.Error("Expected block size error for Blowfish CBC decryption")
	}
}

// TestBlowfishEmptyPlaintext tests Blowfish functions with empty plaintext
func TestBlowfishEmptyPlaintext(t *testing.T) {
	emptyPlaintext := []byte("")
	key := blowfishKey8

	// Test ECB
	ciphertext, err := BlowfishEncryptECB(key, emptyPlaintext)
	if err != nil {
		t.Fatalf("Blowfish ECB encryption of empty plaintext failed: %v", err)
	}

	decrypted, err := BlowfishDecryptECB(key, ciphertext)
	if err != nil {
		t.Fatalf("Blowfish ECB decryption failed: %v", err)
	}

	if !bytes.Equal(emptyPlaintext, decrypted) {
		t.Error("Blowfish ECB: empty plaintext mismatch")
	}

	// Test CBC
	ciphertext, err = BlowfishEncryptCBC(key, emptyPlaintext)
	if err != nil {
		t.Fatalf("Blowfish CBC encryption of empty plaintext failed: %v", err)
	}

	decrypted, err = BlowfishDecryptCBC(key, ciphertext)
	if err != nil {
		t.Fatalf("Blowfish CBC decryption failed: %v", err)
	}

	if !bytes.Equal(emptyPlaintext, decrypted) {
		t.Error("Blowfish CBC: empty plaintext mismatch")
	}

	// Test CFB
	ciphertext, err = BlowfishEncryptCFB(key, emptyPlaintext)
	if err != nil {
		t.Fatalf("Blowfish CFB encryption of empty plaintext failed: %v", err)
	}

	decrypted, err = BlowfishDecryptCFB(key, ciphertext)
	if err != nil {
		t.Fatalf("Blowfish CFB decryption failed: %v", err)
	}

	if !bytes.Equal(emptyPlaintext, decrypted) {
		t.Error("Blowfish CFB: empty plaintext mismatch")
	}

	// Test OFB
	ciphertext, err = BlowfishEncryptOFB(key, emptyPlaintext)
	if err != nil {
		t.Fatalf("Blowfish OFB encryption of empty plaintext failed: %v", err)
	}

	decrypted, err = BlowfishDecryptOFB(key, ciphertext)
	if err != nil {
		t.Fatalf("Blowfish OFB decryption failed: %v", err)
	}

	if !bytes.Equal(emptyPlaintext, decrypted) {
		t.Error("Blowfish OFB: empty plaintext mismatch")
	}
}

// TestBlowfishLargePlaintext tests Blowfish functions with large plaintext
func TestBlowfishLargePlaintext(t *testing.T) {
	// Create a large plaintext (multiple blocks)
	largePlaintext := bytes.Repeat([]byte("This is a large plaintext for testing Blowfish encryption with multiple blocks. "), 10)
	key := blowfishKey16

	// Test ECB
	ciphertext, err := BlowfishEncryptECB(key, largePlaintext)
	if err != nil {
		t.Fatalf("Blowfish ECB encryption of large plaintext failed: %v", err)
	}

	decrypted, err := BlowfishDecryptECB(key, ciphertext)
	if err != nil {
		t.Fatalf("Blowfish ECB decryption failed: %v", err)
	}

	if !bytes.Equal(largePlaintext, decrypted) {
		t.Error("Blowfish ECB: large plaintext mismatch")
	}

	// Test CBC
	ciphertext, err = BlowfishEncryptCBC(key, largePlaintext)
	if err != nil {
		t.Fatalf("Blowfish CBC encryption of large plaintext failed: %v", err)
	}

	decrypted, err = BlowfishDecryptCBC(key, ciphertext)
	if err != nil {
		t.Fatalf("Blowfish CBC decryption failed: %v", err)
	}

	if !bytes.Equal(largePlaintext, decrypted) {
		t.Error("Blowfish CBC: large plaintext mismatch")
	}

	// Test CFB
	ciphertext, err = BlowfishEncryptCFB(key, largePlaintext)
	if err != nil {
		t.Fatalf("Blowfish CFB encryption of large plaintext failed: %v", err)
	}

	decrypted, err = BlowfishDecryptCFB(key, ciphertext)
	if err != nil {
		t.Fatalf("Blowfish CFB decryption failed: %v", err)
	}

	if !bytes.Equal(largePlaintext, decrypted) {
		t.Error("Blowfish CFB: large plaintext mismatch")
	}

	// Test OFB
	ciphertext, err = BlowfishEncryptOFB(key, largePlaintext)
	if err != nil {
		t.Fatalf("Blowfish OFB encryption of large plaintext failed: %v", err)
	}

	decrypted, err = BlowfishDecryptOFB(key, ciphertext)
	if err != nil {
		t.Fatalf("Blowfish OFB decryption failed: %v", err)
	}

	if !bytes.Equal(largePlaintext, decrypted) {
		t.Error("Blowfish OFB: large plaintext mismatch")
	}
}

// TestBlowfishMinimalKey tests Blowfish with 1-byte key
func TestBlowfishMinimalKey(t *testing.T) {
	minimalKey := []byte("a") // 1 byte key
	plaintext := []byte("test message")

	// Test ECB
	ciphertext, err := BlowfishEncryptECB(minimalKey, plaintext)
	if err != nil {
		t.Fatalf("Blowfish ECB encryption with minimal key failed: %v", err)
	}

	decrypted, err := BlowfishDecryptECB(minimalKey, ciphertext)
	if err != nil {
		t.Fatalf("Blowfish ECB decryption with minimal key failed: %v", err)
	}

	if !bytes.Equal(plaintext, decrypted) {
		t.Error("Blowfish ECB: minimal key plaintext mismatch")
	}

	// Test CBC
	ciphertext, err = BlowfishEncryptCBC(minimalKey, plaintext)
	if err != nil {
		t.Fatalf("Blowfish CBC encryption with minimal key failed: %v", err)
	}

	decrypted, err = BlowfishDecryptCBC(minimalKey, ciphertext)
	if err != nil {
		t.Fatalf("Blowfish CBC decryption with minimal key failed: %v", err)
	}

	if !bytes.Equal(plaintext, decrypted) {
		t.Error("Blowfish CBC: minimal key plaintext mismatch")
	}
}
