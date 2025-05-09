package cryptox

import (
	"bytes"
	"crypto/rand"
	"io"
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
