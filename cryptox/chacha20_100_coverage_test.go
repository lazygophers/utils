package cryptox

import (
	"crypto/cipher"
	"errors"
	"testing"

	"golang.org/x/crypto/chacha20"
	"golang.org/x/crypto/chacha20poly1305"
)

// Mock failures for ChaCha20 dependency injection
type FailingChaCha20Reader struct{}

func (fr FailingChaCha20Reader) Read(p []byte) (n int, err error) {
	return 0, errors.New("simulated random reader failure")
}

func FailingChaCha20NewUnauthenticatedCipher(key, nonce []byte) (*chacha20.Cipher, error) {
	return nil, errors.New("simulated ChaCha20 NewUnauthenticatedCipher failure")
}

func FailingChaCha20Poly1305New(key []byte) (cipher.AEAD, error) {
	return nil, errors.New("simulated ChaCha20Poly1305 New failure")
}

// TestChaCha20_100PercentCoverage triggers all error paths using dependency injection
func TestChaCha20_100PercentCoverage(t *testing.T) {
	// Save original functions
	originalChaCha20NewUnauthenticatedCipher := chacha20NewUnauthenticatedCipher
	originalChaCha20Poly1305New := chacha20poly1305New
	originalChaCha20RandReader := chacha20RandReader

	// Restore original functions after test
	defer func() {
		chacha20NewUnauthenticatedCipher = originalChaCha20NewUnauthenticatedCipher
		chacha20poly1305New = originalChaCha20Poly1305New
		chacha20RandReader = originalChaCha20RandReader
	}()

	validKey := make([]byte, chacha20.KeySize)
	validNonce := make([]byte, chacha20.NonceSize)
	plaintext := []byte("test plaintext")

	// Test 1: Trigger rand.Reader failure in functions that use random nonce
	chacha20NewUnauthenticatedCipher = originalChaCha20NewUnauthenticatedCipher
	chacha20poly1305New = originalChaCha20Poly1305New
	chacha20RandReader = FailingChaCha20Reader{}

	_, err := ChaCha20Encrypt(validKey, plaintext)
	if err == nil {
		t.Error("Expected rand.Reader error in ChaCha20Encrypt")
	}

	_, err = ChaCha20Poly1305Encrypt(validKey, plaintext)
	if err == nil {
		t.Error("Expected rand.Reader error in ChaCha20Poly1305Encrypt")
	}

	// Test 2: Trigger ChaCha20 NewUnauthenticatedCipher failure
	chacha20NewUnauthenticatedCipher = FailingChaCha20NewUnauthenticatedCipher
	chacha20poly1305New = originalChaCha20Poly1305New
	chacha20RandReader = originalChaCha20RandReader

	_, err = ChaCha20Encrypt(validKey, plaintext)
	if err == nil {
		t.Error("Expected ChaCha20 NewUnauthenticatedCipher error in ChaCha20Encrypt")
	}

	_, err = ChaCha20Decrypt(validKey, append(validNonce, []byte("fake ciphertext")...))
	if err == nil {
		t.Error("Expected ChaCha20 NewUnauthenticatedCipher error in ChaCha20Decrypt")
	}

	_, err = ChaCha20WithNonce(validKey, validNonce, plaintext)
	if err == nil {
		t.Error("Expected ChaCha20 NewUnauthenticatedCipher error in ChaCha20WithNonce")
	}

	// Test 3: Trigger ChaCha20Poly1305 New failure
	chacha20NewUnauthenticatedCipher = originalChaCha20NewUnauthenticatedCipher
	chacha20poly1305New = FailingChaCha20Poly1305New
	chacha20RandReader = originalChaCha20RandReader

	_, err = ChaCha20Poly1305Encrypt(validKey, plaintext)
	if err == nil {
		t.Error("Expected ChaCha20Poly1305 New error in ChaCha20Poly1305Encrypt")
	}

	_, err = ChaCha20Poly1305Decrypt(validKey, append(validNonce, []byte("fake ciphertext")...))
	if err == nil {
		t.Error("Expected ChaCha20Poly1305 New error in ChaCha20Poly1305Decrypt")
	}

	_, err = ChaCha20Poly1305WithNonce(validKey, validNonce, plaintext)
	if err == nil {
		t.Error("Expected ChaCha20Poly1305 New error in ChaCha20Poly1305WithNonce")
	}

	_, err = ChaCha20Poly1305WithNonceDecrypt(validKey, validNonce, []byte("fake ciphertext"))
	if err == nil {
		t.Error("Expected ChaCha20Poly1305 New error in ChaCha20Poly1305WithNonceDecrypt")
	}
}

// TestChaCha20Poly1305AuthenticationFailure tests authentication failure scenarios
func TestChaCha20Poly1305AuthenticationFailure(t *testing.T) {
	validKey := make([]byte, chacha20poly1305.KeySize)
	validNonce := make([]byte, 12)
	plaintext := []byte("test message")

	// Create valid ciphertext first
	aead, err := chacha20poly1305.New(validKey)
	if err != nil {
		t.Fatalf("Failed to create AEAD: %v", err)
	}

	validCiphertext := aead.Seal(nil, validNonce, plaintext, nil)

	// Test 1: Corrupt the ciphertext to trigger authentication failure
	corruptedCiphertext := make([]byte, len(validCiphertext))
	copy(corruptedCiphertext, validCiphertext)
	if len(corruptedCiphertext) > 0 {
		corruptedCiphertext[0] ^= 0xFF // Flip bits
	}

	_, err = ChaCha20Poly1305WithNonceDecrypt(validKey, validNonce, corruptedCiphertext)
	if err == nil {
		t.Error("Expected authentication failure for corrupted ciphertext")
	}

	// Test 2: Test with empty ciphertext that would cause authentication failure
	_, err = ChaCha20Poly1305WithNonceDecrypt(validKey, validNonce, []byte(""))
	if err == nil {
		t.Error("Expected authentication failure for empty ciphertext")
	}

	// Test 3: Test with short but non-empty ciphertext
	shortCiphertext := []byte("short")
	_, err = ChaCha20Poly1305WithNonceDecrypt(validKey, validNonce, shortCiphertext)
	if err == nil {
		t.Error("Expected authentication failure for short ciphertext")
	}
}

// TestChaCha20Poly1305DecryptAuthFailure tests Open failure scenarios
func TestChaCha20Poly1305DecryptAuthFailure(t *testing.T) {
	validKey := make([]byte, chacha20poly1305.KeySize)
	plaintext := []byte("test message")

	// Create valid ciphertext first
	validCiphertext, err := ChaCha20Poly1305Encrypt(validKey, plaintext)
	if err != nil {
		t.Fatalf("Failed to create valid ciphertext: %v", err)
	}

	// Corrupt the ciphertext portion (not the nonce) to cause authentication failure
	corruptedCiphertext := make([]byte, len(validCiphertext))
	copy(corruptedCiphertext, validCiphertext)

	// Corrupt a byte in the ciphertext portion (after the nonce)
	if len(corruptedCiphertext) > 13 { // 12 bytes nonce + at least 1 byte ciphertext + tag
		corruptedCiphertext[13] ^= 0xFF // Flip bits in the ciphertext
	}

	_, err = ChaCha20Poly1305Decrypt(validKey, corruptedCiphertext)
	if err == nil {
		t.Error("Expected authentication failure in ChaCha20Poly1305Decrypt")
	}
}
