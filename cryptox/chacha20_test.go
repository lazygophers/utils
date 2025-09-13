package cryptox

import (
	"bytes"
	"testing"

	"golang.org/x/crypto/chacha20"
)

// Test data
const (
	testChaCha20Message = "Hello, ChaCha20 encryption test message!"
)

var (
	chacha20Key   = make([]byte, chacha20.KeySize)   // 32 bytes
	chacha20Nonce = make([]byte, chacha20.NonceSize) // 12 bytes
)

// TestChaCha20EncryptDecrypt tests ChaCha20 encryption and decryption
func TestChaCha20EncryptDecrypt(t *testing.T) {
	plaintext := []byte(testChaCha20Message)

	// Test encryption
	ciphertext, err := ChaCha20Encrypt(chacha20Key, plaintext)
	if err != nil {
		t.Fatalf("ChaCha20 encryption failed: %v", err)
	}

	// Test decryption
	decrypted, err := ChaCha20Decrypt(chacha20Key, ciphertext)
	if err != nil {
		t.Fatalf("ChaCha20 decryption failed: %v", err)
	}

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("ChaCha20: plaintext mismatch.\nExpected: %s\nGot: %s", plaintext, decrypted)
	}
}

// TestChaCha20Poly1305EncryptDecrypt tests ChaCha20-Poly1305 encryption and decryption
func TestChaCha20Poly1305EncryptDecrypt(t *testing.T) {
	plaintext := []byte(testChaCha20Message)

	// Test encryption
	ciphertext, err := ChaCha20Poly1305Encrypt(chacha20Key, plaintext)
	if err != nil {
		t.Fatalf("ChaCha20-Poly1305 encryption failed: %v", err)
	}

	// Test decryption
	decrypted, err := ChaCha20Poly1305Decrypt(chacha20Key, ciphertext)
	if err != nil {
		t.Fatalf("ChaCha20-Poly1305 decryption failed: %v", err)
	}

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("ChaCha20-Poly1305: plaintext mismatch.\nExpected: %s\nGot: %s", plaintext, decrypted)
	}
}

// TestChaCha20WithNonce tests ChaCha20 with specified nonce
func TestChaCha20WithNonce(t *testing.T) {
	plaintext := []byte(testChaCha20Message)

	// Test encryption with specified nonce
	ciphertext, err := ChaCha20WithNonce(chacha20Key, chacha20Nonce, plaintext)
	if err != nil {
		t.Fatalf("ChaCha20 with nonce encryption failed: %v", err)
	}

	// Test that encryption with same nonce produces same result
	ciphertext2, err := ChaCha20WithNonce(chacha20Key, chacha20Nonce, plaintext)
	if err != nil {
		t.Fatalf("ChaCha20 with nonce second encryption failed: %v", err)
	}

	if !bytes.Equal(ciphertext, ciphertext2) {
		t.Error("ChaCha20 with same nonce should produce identical ciphertext")
	}

	// Test decryption (ChaCha20 is symmetric, so we can decrypt with same function)
	decrypted, err := ChaCha20WithNonce(chacha20Key, chacha20Nonce, ciphertext)
	if err != nil {
		t.Fatalf("ChaCha20 with nonce decryption failed: %v", err)
	}

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("ChaCha20 with nonce: plaintext mismatch.\nExpected: %s\nGot: %s", plaintext, decrypted)
	}
}

// TestChaCha20Poly1305WithNonce tests ChaCha20-Poly1305 with specified nonce
func TestChaCha20Poly1305WithNonce(t *testing.T) {
	plaintext := []byte(testChaCha20Message)
	nonce := make([]byte, 12) // ChaCha20-Poly1305 uses 12-byte nonce

	// Test encryption with specified nonce
	ciphertext, err := ChaCha20Poly1305WithNonce(chacha20Key, nonce, plaintext)
	if err != nil {
		t.Fatalf("ChaCha20-Poly1305 with nonce encryption failed: %v", err)
	}

	// Test decryption
	decrypted, err := ChaCha20Poly1305WithNonceDecrypt(chacha20Key, nonce, ciphertext)
	if err != nil {
		t.Fatalf("ChaCha20-Poly1305 with nonce decryption failed: %v", err)
	}

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("ChaCha20-Poly1305 with nonce: plaintext mismatch.\nExpected: %s\nGot: %s", plaintext, decrypted)
	}
}

// TestChaCha20InvalidKeyLength tests ChaCha20 functions with invalid key lengths
func TestChaCha20InvalidKeyLength(t *testing.T) {
	plaintext := []byte("test")
	invalidKey := []byte("invalidkey") // Wrong length

	// Test ChaCha20
	_, err := ChaCha20Encrypt(invalidKey, plaintext)
	if err == nil || err.Error() != "invalid key length: must be 32 bytes for ChaCha20" {
		t.Error("Expected invalid key length error for ChaCha20 encryption")
	}

	_, err = ChaCha20Decrypt(invalidKey, []byte("fake ciphertext"))
	if err == nil || err.Error() != "invalid key length: must be 32 bytes for ChaCha20" {
		t.Error("Expected invalid key length error for ChaCha20 decryption")
	}

	_, err = ChaCha20WithNonce(invalidKey, chacha20Nonce, plaintext)
	if err == nil || err.Error() != "invalid key length: must be 32 bytes for ChaCha20" {
		t.Error("Expected invalid key length error for ChaCha20WithNonce")
	}

	// Test ChaCha20-Poly1305
	_, err = ChaCha20Poly1305Encrypt(invalidKey, plaintext)
	if err == nil || err.Error() != "invalid key length: must be 32 bytes for ChaCha20-Poly1305" {
		t.Error("Expected invalid key length error for ChaCha20-Poly1305 encryption")
	}

	_, err = ChaCha20Poly1305Decrypt(invalidKey, []byte("fake ciphertext"))
	if err == nil || err.Error() != "invalid key length: must be 32 bytes for ChaCha20-Poly1305" {
		t.Error("Expected invalid key length error for ChaCha20-Poly1305 decryption")
	}

	_, err = ChaCha20Poly1305WithNonce(invalidKey, chacha20Nonce, plaintext)
	if err == nil || err.Error() != "invalid key length: must be 32 bytes for ChaCha20-Poly1305" {
		t.Error("Expected invalid key length error for ChaCha20Poly1305WithNonce")
	}

	_, err = ChaCha20Poly1305WithNonceDecrypt(invalidKey, chacha20Nonce, []byte("fake"))
	if err == nil || err.Error() != "invalid key length: must be 32 bytes for ChaCha20-Poly1305" {
		t.Error("Expected invalid key length error for ChaCha20Poly1305WithNonceDecrypt")
	}
}

// TestChaCha20InvalidNonceLength tests functions with invalid nonce lengths
func TestChaCha20InvalidNonceLength(t *testing.T) {
	plaintext := []byte("test")
	invalidNonce := []byte("invalid") // Wrong length

	// Test ChaCha20WithNonce
	_, err := ChaCha20WithNonce(chacha20Key, invalidNonce, plaintext)
	if err == nil || err.Error() != "invalid nonce length: must be 12 bytes for ChaCha20" {
		t.Error("Expected invalid nonce length error for ChaCha20WithNonce")
	}

	// Test ChaCha20-Poly1305 with nonce functions
	_, err = ChaCha20Poly1305WithNonce(chacha20Key, invalidNonce, plaintext)
	if err == nil || err.Error() != "invalid nonce length: must be 12 bytes for ChaCha20-Poly1305" {
		t.Error("Expected invalid nonce length error for ChaCha20Poly1305WithNonce")
	}

	_, err = ChaCha20Poly1305WithNonceDecrypt(chacha20Key, invalidNonce, []byte("fake"))
	if err == nil || err.Error() != "invalid nonce length: must be 12 bytes for ChaCha20-Poly1305" {
		t.Error("Expected invalid nonce length error for ChaCha20Poly1305WithNonceDecrypt")
	}
}

// TestChaCha20ShortCiphertext tests functions with short ciphertext
func TestChaCha20ShortCiphertext(t *testing.T) {
	shortCiphertext := make([]byte, chacha20.NonceSize-1) // Less than nonce size

	// Test ChaCha20Decrypt
	_, err := ChaCha20Decrypt(chacha20Key, shortCiphertext)
	if err == nil || err.Error() != "ciphertext too short" {
		t.Error("Expected 'ciphertext too short' error for ChaCha20Decrypt")
	}

	// Test ChaCha20Poly1305Decrypt
	shortPoly1305Ciphertext := make([]byte, 11) // Less than 12 bytes (nonce size)
	_, err = ChaCha20Poly1305Decrypt(chacha20Key, shortPoly1305Ciphertext)
	if err == nil || err.Error() != "ciphertext too short" {
		t.Error("Expected 'ciphertext too short' error for ChaCha20Poly1305Decrypt")
	}
}

// TestChaCha20EmptyPlaintext tests functions with empty plaintext
func TestChaCha20EmptyPlaintext(t *testing.T) {
	emptyPlaintext := []byte("")

	// Test ChaCha20
	ciphertext, err := ChaCha20Encrypt(chacha20Key, emptyPlaintext)
	if err != nil {
		t.Fatalf("ChaCha20 encryption of empty plaintext failed: %v", err)
	}

	decrypted, err := ChaCha20Decrypt(chacha20Key, ciphertext)
	if err != nil {
		t.Fatalf("ChaCha20 decryption failed: %v", err)
	}

	if !bytes.Equal(emptyPlaintext, decrypted) {
		t.Error("ChaCha20: empty plaintext mismatch")
	}

	// Test ChaCha20-Poly1305
	ciphertext, err = ChaCha20Poly1305Encrypt(chacha20Key, emptyPlaintext)
	if err != nil {
		t.Fatalf("ChaCha20-Poly1305 encryption of empty plaintext failed: %v", err)
	}

	decrypted, err = ChaCha20Poly1305Decrypt(chacha20Key, ciphertext)
	if err != nil {
		t.Fatalf("ChaCha20-Poly1305 decryption failed: %v", err)
	}

	if !bytes.Equal(emptyPlaintext, decrypted) {
		t.Error("ChaCha20-Poly1305: empty plaintext mismatch")
	}

	// Test ChaCha20WithNonce
	ciphertext, err = ChaCha20WithNonce(chacha20Key, chacha20Nonce, emptyPlaintext)
	if err != nil {
		t.Fatalf("ChaCha20WithNonce encryption of empty plaintext failed: %v", err)
	}

	decrypted, err = ChaCha20WithNonce(chacha20Key, chacha20Nonce, ciphertext)
	if err != nil {
		t.Fatalf("ChaCha20WithNonce decryption failed: %v", err)
	}

	if !bytes.Equal(emptyPlaintext, decrypted) {
		t.Error("ChaCha20WithNonce: empty plaintext mismatch")
	}
}

// TestChaCha20LargePlaintext tests functions with large plaintext
func TestChaCha20LargePlaintext(t *testing.T) {
	// Create a large plaintext
	largePlaintext := bytes.Repeat([]byte("This is a large plaintext for testing ChaCha20 encryption. "), 100)

	// Test ChaCha20
	ciphertext, err := ChaCha20Encrypt(chacha20Key, largePlaintext)
	if err != nil {
		t.Fatalf("ChaCha20 encryption of large plaintext failed: %v", err)
	}

	decrypted, err := ChaCha20Decrypt(chacha20Key, ciphertext)
	if err != nil {
		t.Fatalf("ChaCha20 decryption failed: %v", err)
	}

	if !bytes.Equal(largePlaintext, decrypted) {
		t.Error("ChaCha20: large plaintext mismatch")
	}

	// Test ChaCha20-Poly1305
	ciphertext, err = ChaCha20Poly1305Encrypt(chacha20Key, largePlaintext)
	if err != nil {
		t.Fatalf("ChaCha20-Poly1305 encryption of large plaintext failed: %v", err)
	}

	decrypted, err = ChaCha20Poly1305Decrypt(chacha20Key, ciphertext)
	if err != nil {
		t.Fatalf("ChaCha20-Poly1305 decryption failed: %v", err)
	}

	if !bytes.Equal(largePlaintext, decrypted) {
		t.Error("ChaCha20-Poly1305: large plaintext mismatch")
	}
}

// TestChaCha20Poly1305Authentication tests ChaCha20-Poly1305 authentication
func TestChaCha20Poly1305Authentication(t *testing.T) {
	plaintext := []byte(testChaCha20Message)

	// Test encryption
	ciphertext, err := ChaCha20Poly1305Encrypt(chacha20Key, plaintext)
	if err != nil {
		t.Fatalf("ChaCha20-Poly1305 encryption failed: %v", err)
	}

	// Corrupt the ciphertext (excluding nonce)
	if len(ciphertext) > 13 { // nonce(12) + at least 1 byte of ciphertext
		corruptedCiphertext := make([]byte, len(ciphertext))
		copy(corruptedCiphertext, ciphertext)
		corruptedCiphertext[13] ^= 0xFF // Flip bits in the first ciphertext byte

		// Decryption should fail due to authentication failure
		_, err = ChaCha20Poly1305Decrypt(chacha20Key, corruptedCiphertext)
		if err == nil {
			t.Error("Expected authentication failure for corrupted ChaCha20-Poly1305 ciphertext")
		}
	}
}

// TestChaCha20StreamCipher tests that ChaCha20 behaves as a stream cipher
func TestChaCha20StreamCipher(t *testing.T) {
	plaintext1 := []byte("Hello")
	plaintext2 := []byte("World")
	combinedPlaintext := append(plaintext1, plaintext2...)

	// Encrypt first part separately
	ciphertext1, err := ChaCha20WithNonce(chacha20Key, chacha20Nonce, plaintext1)
	if err != nil {
		t.Fatalf("Failed to encrypt first part: %v", err)
	}

	// Encrypt combined plaintext
	combinedCiphertext, err := ChaCha20WithNonce(chacha20Key, chacha20Nonce, combinedPlaintext)
	if err != nil {
		t.Fatalf("Failed to encrypt combined plaintext: %v", err)
	}

	// Verify that the first part matches (stream cipher property)
	if !bytes.Equal(ciphertext1, combinedCiphertext[:len(ciphertext1)]) {
		t.Error("ChaCha20 stream cipher property test failed")
	}
}
