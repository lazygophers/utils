package cryptox

import (
	"crypto/rand"
	"errors"
	"io"

	"golang.org/x/crypto/chacha20"
	"golang.org/x/crypto/chacha20poly1305"
)

// Global variables for dependency injection during testing
var (
	chacha20NewUnauthenticatedCipher = chacha20.NewUnauthenticatedCipher
	chacha20poly1305New              = chacha20poly1305.New
	chacha20RandReader               = rand.Reader
)

// ChaCha20Encrypt 使用 ChaCha20 流密码加密明文
func ChaCha20Encrypt(key, plaintext []byte) ([]byte, error) {
	if len(key) != chacha20.KeySize {
		return nil, errors.New("invalid key length: must be 32 bytes for ChaCha20")
	}

	// Generate random nonce
	nonce := make([]byte, chacha20.NonceSize)
	_, err := io.ReadFull(chacha20RandReader, nonce)
	if err != nil {
		return nil, err
	}

	cipher, err := chacha20NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, chacha20.NonceSize+len(plaintext))
	copy(ciphertext[:chacha20.NonceSize], nonce)
	
	cipher.XORKeyStream(ciphertext[chacha20.NonceSize:], plaintext)
	return ciphertext, nil
}

// ChaCha20Decrypt 使用 ChaCha20 流密码解密密文
func ChaCha20Decrypt(key, ciphertext []byte) ([]byte, error) {
	if len(key) != chacha20.KeySize {
		return nil, errors.New("invalid key length: must be 32 bytes for ChaCha20")
	}

	if len(ciphertext) < chacha20.NonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce := ciphertext[:chacha20.NonceSize]
	ciphertext = ciphertext[chacha20.NonceSize:]

	cipher, err := chacha20NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, len(ciphertext))
	cipher.XORKeyStream(plaintext, ciphertext)
	return plaintext, nil
}

// ChaCha20Poly1305Encrypt 使用 ChaCha20-Poly1305 AEAD 加密明文
func ChaCha20Poly1305Encrypt(key, plaintext []byte) ([]byte, error) {
	if len(key) != chacha20poly1305.KeySize {
		return nil, errors.New("invalid key length: must be 32 bytes for ChaCha20-Poly1305")
	}

	aead, err := chacha20poly1305New(key)
	if err != nil {
		return nil, err
	}

	// Generate random nonce
	nonce := make([]byte, aead.NonceSize())
	_, err = io.ReadFull(chacha20RandReader, nonce)
	if err != nil {
		return nil, err
	}

	// Seal encrypts and authenticates plaintext
	ciphertext := aead.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// ChaCha20Poly1305Decrypt 使用 ChaCha20-Poly1305 AEAD 解密密文
func ChaCha20Poly1305Decrypt(key, ciphertext []byte) ([]byte, error) {
	if len(key) != chacha20poly1305.KeySize {
		return nil, errors.New("invalid key length: must be 32 bytes for ChaCha20-Poly1305")
	}

	aead, err := chacha20poly1305New(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aead.NonceSize() {
		return nil, errors.New("ciphertext too short")
	}

	nonce := ciphertext[:aead.NonceSize()]
	ciphertext = ciphertext[aead.NonceSize():]

	// Open decrypts and verifies ciphertext
	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// ChaCha20WithNonce 使用指定的 nonce 进行 ChaCha20 加密（用于测试或特殊需求）
func ChaCha20WithNonce(key, nonce, plaintext []byte) ([]byte, error) {
	if len(key) != chacha20.KeySize {
		return nil, errors.New("invalid key length: must be 32 bytes for ChaCha20")
	}

	if len(nonce) != chacha20.NonceSize {
		return nil, errors.New("invalid nonce length: must be 12 bytes for ChaCha20")
	}

	cipher, err := chacha20NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, len(plaintext))
	cipher.XORKeyStream(ciphertext, plaintext)
	return ciphertext, nil
}

// ChaCha20Poly1305WithNonce 使用指定的 nonce 进行 ChaCha20-Poly1305 加密（用于测试或特殊需求）
func ChaCha20Poly1305WithNonce(key, nonce, plaintext []byte) ([]byte, error) {
	if len(key) != chacha20poly1305.KeySize {
		return nil, errors.New("invalid key length: must be 32 bytes for ChaCha20-Poly1305")
	}

	aead, err := chacha20poly1305New(key)
	if err != nil {
		return nil, err
	}

	if len(nonce) != aead.NonceSize() {
		return nil, errors.New("invalid nonce length: must be 12 bytes for ChaCha20-Poly1305")
	}

	// Seal encrypts and authenticates plaintext
	ciphertext := aead.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nil
}

// ChaCha20Poly1305WithNonceDecrypt 使用指定的 nonce 进行 ChaCha20-Poly1305 解密
func ChaCha20Poly1305WithNonceDecrypt(key, nonce, ciphertext []byte) ([]byte, error) {
	if len(key) != chacha20poly1305.KeySize {
		return nil, errors.New("invalid key length: must be 32 bytes for ChaCha20-Poly1305")
	}

	aead, err := chacha20poly1305New(key)
	if err != nil {
		return nil, err
	}

	if len(nonce) != aead.NonceSize() {
		return nil, errors.New("invalid nonce length: must be 12 bytes for ChaCha20-Poly1305")
	}

	// Open decrypts and verifies ciphertext
	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}