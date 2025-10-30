package cryptox

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"strings"
	"testing"
)

// generateValidKey generates a valid 32-byte AES-256 key for testing
func generateValidKey() []byte {
	return []byte("12345678901234567890123456789012") // 32 bytes
}

// TestEncryptDecryptGCM tests AES-GCM encryption and decryption round-trip
func TestEncryptDecryptGCM(t *testing.T) {
	key := generateValidKey()

	testCases := []struct {
		name      string
		plaintext []byte
	}{
		{"empty data", []byte("")},
		{"single byte", []byte("A")},
		{"short message", []byte("Hello, World!")},
		{"medium message", []byte("The quick brown fox jumps over the lazy dog")},
		{"long message", bytes.Repeat([]byte("X"), 1000)},
		{"binary data", []byte{0x00, 0x01, 0x02, 0x03, 0xFF, 0xFE, 0xFD, 0xFC}},
		{"unicode data", []byte("你好世界🌍Hello World🚀")},
		{"all zeros", make([]byte, 100)},
		{"block aligned", make([]byte, aes.BlockSize*3)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Encrypt
			ciphertext, err := Encrypt(key, tc.plaintext)
			if err != nil {
				t.Fatalf("Encrypt failed: %v", err)
			}

			// Verify ciphertext is different from plaintext (unless empty)
			if len(tc.plaintext) > 0 && bytes.Equal(ciphertext, tc.plaintext) {
				t.Error("Ciphertext should be different from plaintext")
			}

			// Decrypt
			decrypted, err := Decrypt(key, ciphertext)
			if err != nil {
				t.Fatalf("Decrypt failed: %v", err)
			}

			// Verify round-trip
			if !bytes.Equal(decrypted, tc.plaintext) {
				t.Errorf("Decrypted data doesn't match original.\nGot:      %v\nExpected: %v", decrypted, tc.plaintext)
			}
		})
	}
}

// TestEncryptDecryptGCMErrors tests error cases for GCM mode
func TestEncryptDecryptGCMErrors(t *testing.T) {
	validKey := generateValidKey()
	plaintext := []byte("test message")

	t.Run("Encrypt with invalid key length", func(t *testing.T) {
		invalidKeys := [][]byte{
			[]byte("short"),                    // too short
			[]byte("1234567890123456"),         // 16 bytes (AES-128, but we require AES-256)
			[]byte("123456789012345678901234"), // 24 bytes (AES-192, but we require AES-256)
			make([]byte, 64),                   // too long
		}

		for _, key := range invalidKeys {
			_, err := Encrypt(key, plaintext)
			if err == nil {
				t.Errorf("Encrypt should fail with key length %d", len(key))
			}
			if err != nil && !strings.Contains(err.Error(), "invalid key length") {
				t.Errorf("Expected 'invalid key length' error, got: %v", err)
			}
		}
	})

	t.Run("Decrypt with invalid key length", func(t *testing.T) {
		ciphertext, _ := Encrypt(validKey, plaintext)

		invalidKey := []byte("short")
		_, err := Decrypt(invalidKey, ciphertext)
		if err == nil {
			t.Error("Decrypt should fail with invalid key length")
		}
		if err != nil && !strings.Contains(err.Error(), "invalid key length") {
			t.Errorf("Expected 'invalid key length' error, got: %v", err)
		}
	})

	t.Run("Decrypt with ciphertext too short", func(t *testing.T) {
		shortCiphertext := []byte("short")
		_, err := Decrypt(validKey, shortCiphertext)
		if err == nil {
			t.Error("Decrypt should fail with ciphertext too short")
		}
		if err != nil && !strings.Contains(err.Error(), "ciphertext too short") {
			t.Errorf("Expected 'ciphertext too short' error, got: %v", err)
		}
	})

	t.Run("Decrypt with corrupted ciphertext", func(t *testing.T) {
		ciphertext, _ := Encrypt(validKey, plaintext)
		// Corrupt the ciphertext
		ciphertext[len(ciphertext)-1] ^= 0xFF
		_, err := Decrypt(validKey, ciphertext)
		if err == nil {
			t.Error("Decrypt should fail with corrupted ciphertext")
		}
	})

	t.Run("Decrypt with wrong key", func(t *testing.T) {
		ciphertext, _ := Encrypt(validKey, plaintext)
		wrongKey := []byte("12345678901234567890123456789099") // different last byte
		_, err := Decrypt(wrongKey, ciphertext)
		if err == nil {
			t.Error("Decrypt should fail with wrong key")
		}
	})
}

// TestEncryptDecryptECB tests AES-ECB encryption and decryption round-trip
func TestEncryptDecryptECB(t *testing.T) {
	key := generateValidKey()

	testCases := []struct {
		name      string
		plaintext []byte
	}{
		{"empty data", []byte("")},
		{"single byte", []byte("A")},
		{"short message", []byte("Hello, World!")},
		{"medium message", []byte("The quick brown fox jumps over the lazy dog")},
		{"long message", bytes.Repeat([]byte("X"), 1000)},
		{"binary data", []byte{0x00, 0x01, 0x02, 0x03, 0xFF, 0xFE, 0xFD, 0xFC}},
		{"unicode data", []byte("你好世界🌍Hello World🚀")},
		{"block aligned", make([]byte, aes.BlockSize*3)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Encrypt
			ciphertext, err := EncryptECB(key, tc.plaintext)
			if err != nil {
				t.Fatalf("EncryptECB failed: %v", err)
			}

			// Decrypt
			decrypted, err := DecryptECB(key, ciphertext)
			if err != nil {
				t.Fatalf("DecryptECB failed: %v", err)
			}

			// Verify round-trip
			if !bytes.Equal(decrypted, tc.plaintext) {
				t.Errorf("Decrypted data doesn't match original.\nGot:      %v\nExpected: %v", decrypted, tc.plaintext)
			}
		})
	}
}

// TestEncryptDecryptECBErrors tests error cases for ECB mode
func TestEncryptDecryptECBErrors(t *testing.T) {
	validKey := generateValidKey()
	plaintext := []byte("test message")

	t.Run("EncryptECB with invalid key length", func(t *testing.T) {
		invalidKey := []byte("short")
		_, err := EncryptECB(invalidKey, plaintext)
		if err == nil || !strings.Contains(err.Error(), "invalid key length") {
			t.Errorf("Expected 'invalid key length' error, got: %v", err)
		}
	})

	t.Run("DecryptECB with invalid key length", func(t *testing.T) {
		ciphertext, _ := EncryptECB(validKey, plaintext)
		invalidKey := []byte("short")
		_, err := DecryptECB(invalidKey, ciphertext)
		if err == nil || !strings.Contains(err.Error(), "invalid key length") {
			t.Errorf("Expected 'invalid key length' error, got: %v", err)
		}
	})

	t.Run("DecryptECB with invalid block size", func(t *testing.T) {
		// Create ciphertext with invalid length (not multiple of block size)
		invalidCiphertext := make([]byte, aes.BlockSize+1)
		_, err := DecryptECB(validKey, invalidCiphertext)
		if err == nil {
			t.Error("DecryptECB should fail with ciphertext not multiple of block size")
		}
		if err != nil && !strings.Contains(err.Error(), "not a multiple of the block size") {
			t.Errorf("Expected block size error, got: %v", err)
		}
	})
}

// TestEncryptDecryptCBC tests AES-CBC encryption and decryption round-trip
func TestEncryptDecryptCBC(t *testing.T) {
	key := generateValidKey()

	testCases := []struct {
		name      string
		plaintext []byte
	}{
		{"empty data", []byte("")},
		{"single byte", []byte("A")},
		{"short message", []byte("Hello, World!")},
		{"medium message", []byte("The quick brown fox jumps over the lazy dog")},
		{"long message", bytes.Repeat([]byte("X"), 1000)},
		{"binary data", []byte{0x00, 0x01, 0x02, 0x03, 0xFF, 0xFE, 0xFD, 0xFC}},
		{"unicode data", []byte("你好世界🌍Hello World🚀")},
		{"block aligned", make([]byte, aes.BlockSize*3)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Encrypt
			ciphertext, err := EncryptCBC(key, tc.plaintext)
			if err != nil {
				t.Fatalf("EncryptCBC failed: %v", err)
			}

			// Decrypt
			decrypted, err := DecryptCBC(key, ciphertext)
			if err != nil {
				t.Fatalf("DecryptCBC failed: %v", err)
			}

			// Verify round-trip
			if !bytes.Equal(decrypted, tc.plaintext) {
				t.Errorf("Decrypted data doesn't match original.\nGot:      %v\nExpected: %v", decrypted, tc.plaintext)
			}
		})
	}
}

// TestEncryptDecryptCBCErrors tests error cases for CBC mode
func TestEncryptDecryptCBCErrors(t *testing.T) {
	validKey := generateValidKey()
	plaintext := []byte("test message")

	t.Run("EncryptCBC with invalid key length", func(t *testing.T) {
		invalidKey := []byte("short")
		_, err := EncryptCBC(invalidKey, plaintext)
		if err == nil || !strings.Contains(err.Error(), "invalid key length") {
			t.Errorf("Expected 'invalid key length' error, got: %v", err)
		}
	})

	t.Run("DecryptCBC with invalid key length", func(t *testing.T) {
		ciphertext, _ := EncryptCBC(validKey, plaintext)
		invalidKey := []byte("short")
		_, err := DecryptCBC(invalidKey, ciphertext)
		if err == nil || !strings.Contains(err.Error(), "invalid key length") {
			t.Errorf("Expected 'invalid key length' error, got: %v", err)
		}
	})

	t.Run("DecryptCBC with ciphertext too short", func(t *testing.T) {
		shortCiphertext := []byte("short")
		_, err := DecryptCBC(validKey, shortCiphertext)
		if err == nil {
			t.Error("DecryptCBC should fail with ciphertext too short")
		}
		if err != nil && !strings.Contains(err.Error(), "ciphertext too short") {
			t.Errorf("Expected 'ciphertext too short' error, got: %v", err)
		}
	})

	t.Run("DecryptCBC with corrupted padding", func(t *testing.T) {
		ciphertext, _ := EncryptCBC(validKey, plaintext)
		// Corrupt the last block (where padding is)
		ciphertext[len(ciphertext)-1] ^= 0xFF
		_, err := DecryptCBC(validKey, ciphertext)
		if err == nil {
			t.Error("DecryptCBC should fail with corrupted padding")
		}
	})
}

// TestEncryptDecryptCFB tests AES-CFB encryption and decryption round-trip
func TestEncryptDecryptCFB(t *testing.T) {
	key := generateValidKey()

	testCases := []struct {
		name      string
		plaintext []byte
	}{
		{"empty data", []byte("")},
		{"single byte", []byte("A")},
		{"short message", []byte("Hello, World!")},
		{"medium message", []byte("The quick brown fox jumps over the lazy dog")},
		{"long message", bytes.Repeat([]byte("X"), 1000)},
		{"binary data", []byte{0x00, 0x01, 0x02, 0x03, 0xFF, 0xFE, 0xFD, 0xFC}},
		{"unicode data", []byte("你好世界🌍Hello World🚀")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Encrypt
			ciphertext, err := EncryptCFB(key, tc.plaintext)
			if err != nil {
				t.Fatalf("EncryptCFB failed: %v", err)
			}

			// Decrypt
			decrypted, err := DecryptCFB(key, ciphertext)
			if err != nil {
				t.Fatalf("DecryptCFB failed: %v", err)
			}

			// Verify round-trip
			if !bytes.Equal(decrypted, tc.plaintext) {
				t.Errorf("Decrypted data doesn't match original.\nGot:      %v\nExpected: %v", decrypted, tc.plaintext)
			}
		})
	}
}

// TestEncryptDecryptCFBErrors tests error cases for CFB mode
func TestEncryptDecryptCFBErrors(t *testing.T) {
	validKey := generateValidKey()
	plaintext := []byte("test message")

	t.Run("EncryptCFB with invalid key length", func(t *testing.T) {
		invalidKey := []byte("short")
		_, err := EncryptCFB(invalidKey, plaintext)
		if err == nil || !strings.Contains(err.Error(), "invalid key length") {
			t.Errorf("Expected 'invalid key length' error, got: %v", err)
		}
	})

	t.Run("DecryptCFB with invalid key length", func(t *testing.T) {
		ciphertext, _ := EncryptCFB(validKey, plaintext)
		invalidKey := []byte("short")
		_, err := DecryptCFB(invalidKey, ciphertext)
		if err == nil || !strings.Contains(err.Error(), "invalid key length") {
			t.Errorf("Expected 'invalid key length' error, got: %v", err)
		}
	})

	t.Run("DecryptCFB with ciphertext too short", func(t *testing.T) {
		shortCiphertext := []byte("short")
		_, err := DecryptCFB(validKey, shortCiphertext)
		if err == nil {
			t.Error("DecryptCFB should fail with ciphertext too short")
		}
		if err != nil && !strings.Contains(err.Error(), "ciphertext too short") {
			t.Errorf("Expected 'ciphertext too short' error, got: %v", err)
		}
	})
}

// TestEncryptDecryptCTR tests AES-CTR encryption and decryption round-trip
func TestEncryptDecryptCTR(t *testing.T) {
	key := generateValidKey()

	testCases := []struct {
		name      string
		plaintext []byte
	}{
		{"empty data", []byte("")},
		{"single byte", []byte("A")},
		{"short message", []byte("Hello, World!")},
		{"medium message", []byte("The quick brown fox jumps over the lazy dog")},
		{"long message", bytes.Repeat([]byte("X"), 1000)},
		{"binary data", []byte{0x00, 0x01, 0x02, 0x03, 0xFF, 0xFE, 0xFD, 0xFC}},
		{"unicode data", []byte("你好世界🌍Hello World🚀")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Encrypt
			ciphertext, err := EncryptCTR(key, tc.plaintext)
			if err != nil {
				t.Fatalf("EncryptCTR failed: %v", err)
			}

			// Decrypt
			decrypted, err := DecryptCTR(key, ciphertext)
			if err != nil {
				t.Fatalf("DecryptCTR failed: %v", err)
			}

			// Verify round-trip
			if !bytes.Equal(decrypted, tc.plaintext) {
				t.Errorf("Decrypted data doesn't match original.\nGot:      %v\nExpected: %v", decrypted, tc.plaintext)
			}
		})
	}
}

// TestEncryptDecryptCTRErrors tests error cases for CTR mode
func TestEncryptDecryptCTRErrors(t *testing.T) {
	validKey := generateValidKey()
	plaintext := []byte("test message")

	t.Run("EncryptCTR with invalid key length", func(t *testing.T) {
		invalidKey := []byte("short")
		_, err := EncryptCTR(invalidKey, plaintext)
		if err == nil || !strings.Contains(err.Error(), "invalid key length") {
			t.Errorf("Expected 'invalid key length' error, got: %v", err)
		}
	})

	t.Run("DecryptCTR with invalid key length", func(t *testing.T) {
		ciphertext, _ := EncryptCTR(validKey, plaintext)
		invalidKey := []byte("short")
		_, err := DecryptCTR(invalidKey, ciphertext)
		if err == nil || !strings.Contains(err.Error(), "invalid key length") {
			t.Errorf("Expected 'invalid key length' error, got: %v", err)
		}
	})

	t.Run("DecryptCTR with ciphertext too short", func(t *testing.T) {
		shortCiphertext := []byte("short")
		_, err := DecryptCTR(validKey, shortCiphertext)
		if err == nil {
			t.Error("DecryptCTR should fail with ciphertext too short")
		}
		if err != nil && !strings.Contains(err.Error(), "ciphertext too short") {
			t.Errorf("Expected 'ciphertext too short' error, got: %v", err)
		}
	})
}

// TestEncryptDecryptOFB tests AES-OFB encryption and decryption round-trip
func TestEncryptDecryptOFB(t *testing.T) {
	key := generateValidKey()

	testCases := []struct {
		name      string
		plaintext []byte
	}{
		{"empty data", []byte("")},
		{"single byte", []byte("A")},
		{"short message", []byte("Hello, World!")},
		{"medium message", []byte("The quick brown fox jumps over the lazy dog")},
		{"long message", bytes.Repeat([]byte("X"), 1000)},
		{"binary data", []byte{0x00, 0x01, 0x02, 0x03, 0xFF, 0xFE, 0xFD, 0xFC}},
		{"unicode data", []byte("你好世界🌍Hello World🚀")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Encrypt
			ciphertext, err := EncryptOFB(key, tc.plaintext)
			if err != nil {
				t.Fatalf("EncryptOFB failed: %v", err)
			}

			// Decrypt
			decrypted, err := DecryptOFB(key, ciphertext)
			if err != nil {
				t.Fatalf("DecryptOFB failed: %v", err)
			}

			// Verify round-trip
			if !bytes.Equal(decrypted, tc.plaintext) {
				t.Errorf("Decrypted data doesn't match original.\nGot:      %v\nExpected: %v", decrypted, tc.plaintext)
			}
		})
	}
}

// TestEncryptDecryptOFBErrors tests error cases for OFB mode
func TestEncryptDecryptOFBErrors(t *testing.T) {
	validKey := generateValidKey()
	plaintext := []byte("test message")

	t.Run("EncryptOFB with invalid key length", func(t *testing.T) {
		invalidKey := []byte("short")
		_, err := EncryptOFB(invalidKey, plaintext)
		if err == nil || !strings.Contains(err.Error(), "invalid key length") {
			t.Errorf("Expected 'invalid key length' error, got: %v", err)
		}
	})

	t.Run("DecryptOFB with invalid key length", func(t *testing.T) {
		ciphertext, _ := EncryptOFB(validKey, plaintext)
		invalidKey := []byte("short")
		_, err := DecryptOFB(invalidKey, ciphertext)
		if err == nil || !strings.Contains(err.Error(), "invalid key length") {
			t.Errorf("Expected 'invalid key length' error, got: %v", err)
		}
	})

	t.Run("DecryptOFB with ciphertext too short", func(t *testing.T) {
		shortCiphertext := []byte("short")
		_, err := DecryptOFB(validKey, shortCiphertext)
		if err == nil {
			t.Error("DecryptOFB should fail with ciphertext too short")
		}
		if err != nil && !strings.Contains(err.Error(), "ciphertext too short") {
			t.Errorf("Expected 'ciphertext too short' error, got: %v", err)
		}
	})
}

// TestPadPKCS7 tests PKCS#7 padding function
func TestPadPKCS7(t *testing.T) {
	testCases := []struct {
		name      string
		data      []byte
		blockSize int
		expected  int // expected length after padding
	}{
		{"empty data", []byte{}, 16, 16},
		{"single byte", []byte{0x01}, 16, 16},
		{"15 bytes", bytes.Repeat([]byte{0x01}, 15), 16, 16},
		{"16 bytes (full block)", bytes.Repeat([]byte{0x01}, 16), 16, 32},
		{"17 bytes", bytes.Repeat([]byte{0x01}, 17), 16, 32},
		{"31 bytes", bytes.Repeat([]byte{0x01}, 31), 16, 32},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			padded := padPKCS7(tc.data, tc.blockSize)
			if len(padded) != tc.expected {
				t.Errorf("Expected length %d, got %d", tc.expected, len(padded))
			}

			// Verify padding is correct
			paddingLen := int(padded[len(padded)-1])
			if paddingLen == 0 || paddingLen > tc.blockSize {
				t.Errorf("Invalid padding length: %d", paddingLen)
			}

			// Verify all padding bytes are the same
			for i := len(padded) - paddingLen; i < len(padded); i++ {
				if padded[i] != byte(paddingLen) {
					t.Errorf("Padding byte at position %d is %d, expected %d", i, padded[i], paddingLen)
				}
			}
		})
	}
}

// TestUnpadPKCS7 tests PKCS#7 unpadding function
func TestUnpadPKCS7(t *testing.T) {
	blockSize := 16

	t.Run("valid padding", func(t *testing.T) {
		testCases := []struct {
			name     string
			original []byte
		}{
			{"empty data", []byte{}},
			{"single byte", []byte{0x01}},
			{"15 bytes", bytes.Repeat([]byte{0x01}, 15)},
			{"16 bytes", bytes.Repeat([]byte{0x01}, 16)},
			{"32 bytes", bytes.Repeat([]byte{0x01}, 32)},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				padded := padPKCS7(tc.original, blockSize)
				unpadded, err := unpadPKCS7(padded)
				if err != nil {
					t.Fatalf("unpadPKCS7 failed: %v", err)
				}
				if !bytes.Equal(unpadded, tc.original) {
					t.Errorf("Unpadded data doesn't match original.\nGot:      %v\nExpected: %v", unpadded, tc.original)
				}
			})
		}
	})

	t.Run("invalid padding", func(t *testing.T) {
		invalidCases := []struct {
			name string
			data []byte
		}{
			{"empty data", []byte{}},
			{"zero padding", []byte{0x01, 0x02, 0x03, 0x00}},
			{"padding too large", []byte{0x01, 0x02, 0x03, 0x20}},
			{"inconsistent padding", []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x03}},
		}

		for _, tc := range invalidCases {
			t.Run(tc.name, func(t *testing.T) {
				_, err := unpadPKCS7(tc.data)
				if err == nil {
					t.Error("unpadPKCS7 should fail with invalid padding")
				}
			})
		}
	})
}

// TestDecryptECBBlockSizeValidation tests block size validation in DecryptECB
func TestDecryptECBBlockSizeValidation(t *testing.T) {
	key := generateValidKey()

	invalidSizes := []int{1, 7, 15, 17, 23, 31, 33}
	for _, size := range invalidSizes {
		t.Run("size_"+string(rune(size+'0')), func(t *testing.T) {
			invalidCiphertext := make([]byte, size)
			_, err := DecryptECB(key, invalidCiphertext)
			if err == nil {
				t.Errorf("DecryptECB should fail with ciphertext size %d", size)
			}
			if err != nil && !strings.Contains(err.Error(), "not a multiple of the block size") {
				t.Errorf("Expected block size error, got: %v", err)
			}
		})
	}
}

// TestAESModeConsistency tests that same plaintext with same key produces different ciphertexts
// due to random IV/nonce (except for deterministic ECB mode)
func TestAESModeConsistency(t *testing.T) {
	key := generateValidKey()
	plaintext := []byte("test message for consistency check")

	t.Run("GCM produces different ciphertexts", func(t *testing.T) {
		ct1, _ := Encrypt(key, plaintext)
		ct2, _ := Encrypt(key, plaintext)
		if bytes.Equal(ct1, ct2) {
			t.Error("GCM should produce different ciphertexts with random nonce")
		}
	})

	t.Run("CBC produces different ciphertexts", func(t *testing.T) {
		ct1, _ := EncryptCBC(key, plaintext)
		ct2, _ := EncryptCBC(key, plaintext)
		if bytes.Equal(ct1, ct2) {
			t.Error("CBC should produce different ciphertexts with random IV")
		}
	})

	t.Run("CFB produces different ciphertexts", func(t *testing.T) {
		ct1, _ := EncryptCFB(key, plaintext)
		ct2, _ := EncryptCFB(key, plaintext)
		if bytes.Equal(ct1, ct2) {
			t.Error("CFB should produce different ciphertexts with random IV")
		}
	})

	t.Run("CTR produces different ciphertexts", func(t *testing.T) {
		ct1, _ := EncryptCTR(key, plaintext)
		ct2, _ := EncryptCTR(key, plaintext)
		if bytes.Equal(ct1, ct2) {
			t.Error("CTR should produce different ciphertexts with random IV")
		}
	})

	t.Run("OFB produces different ciphertexts", func(t *testing.T) {
		ct1, _ := EncryptOFB(key, plaintext)
		ct2, _ := EncryptOFB(key, plaintext)
		if bytes.Equal(ct1, ct2) {
			t.Error("OFB should produce different ciphertexts with random IV")
		}
	})
}

// TestAESWithDifferentKeys tests that different keys produce different ciphertexts
func TestAESWithDifferentKeys(t *testing.T) {
	key1 := generateValidKey()
	key2 := []byte("98765432109876543210987654321098") // different key
	plaintext := []byte("test message")

	modes := []struct {
		name    string
		encrypt func([]byte, []byte) ([]byte, error)
	}{
		{"GCM", Encrypt},
		{"ECB", EncryptECB},
		{"CBC", EncryptCBC},
		{"CFB", EncryptCFB},
		{"CTR", EncryptCTR},
		{"OFB", EncryptOFB},
	}

	for _, mode := range modes {
		t.Run(mode.name, func(t *testing.T) {
			ct1, _ := mode.encrypt(key1, plaintext)
			ct2, _ := mode.encrypt(key2, plaintext)

			// For ECB, ciphertexts should be completely different
			// For others, at least the encrypted part should differ (IV/nonce will be random)
			if mode.name == "ECB" {
				if bytes.Equal(ct1, ct2) {
					t.Error("Different keys should produce different ciphertexts")
				}
			} else {
				// For modes with IV, check that they're not identical
				if bytes.Equal(ct1, ct2) {
					t.Error("Different keys should produce different ciphertexts")
				}
			}
		})
	}
}

// TestEncryptErrorPaths tests error paths in encryption functions using dependency injection
func TestEncryptErrorPaths(t *testing.T) {
	key := generateValidKey()
	plaintext := []byte("test message")

	t.Run("Encrypt fails when newCipherFunc returns error", func(t *testing.T) {
		originalNewCipherFunc := newCipherFunc
		newCipherFunc = func(key []byte) (cipher.Block, error) {
			return nil, errors.New("mock cipher error")
		}
		defer func() { newCipherFunc = originalNewCipherFunc }()

		_, err := Encrypt(key, plaintext)
		if err == nil {
			t.Error("Encrypt should fail when newCipherFunc returns error")
		}
		if err != nil && !strings.Contains(err.Error(), "mock cipher error") {
			t.Errorf("Expected mock cipher error, got: %v", err)
		}
	})

	t.Run("Encrypt fails when newGCMFunc returns error", func(t *testing.T) {
		originalNewGCMFunc := newGCMFunc
		newGCMFunc = func(cipher cipher.Block) (cipher.AEAD, error) {
			return nil, errors.New("mock GCM error")
		}
		defer func() { newGCMFunc = originalNewGCMFunc }()

		_, err := Encrypt(key, plaintext)
		if err == nil {
			t.Error("Encrypt should fail when newGCMFunc returns error")
		}
		if err != nil && !strings.Contains(err.Error(), "mock GCM error") {
			t.Errorf("Expected mock GCM error, got: %v", err)
		}
	})

	t.Run("Encrypt fails when randReader returns error", func(t *testing.T) {
		originalRandReader := randReader
		randReader = &failingReader{}
		defer func() { randReader = originalRandReader }()

		_, err := Encrypt(key, plaintext)
		if err == nil {
			t.Error("Encrypt should fail when randReader returns error")
		}
		if err != nil && !strings.Contains(err.Error(), "mock random error") {
			t.Errorf("Expected mock random error, got: %v", err)
		}
	})

	t.Run("EncryptECB fails when newCipherFunc returns error", func(t *testing.T) {
		originalNewCipherFunc := newCipherFunc
		newCipherFunc = func(key []byte) (cipher.Block, error) {
			return nil, errors.New("mock cipher error")
		}
		defer func() { newCipherFunc = originalNewCipherFunc }()

		_, err := EncryptECB(key, plaintext)
		if err == nil {
			t.Error("EncryptECB should fail when newCipherFunc returns error")
		}
	})

	t.Run("EncryptCBC fails when newCipherFunc returns error", func(t *testing.T) {
		originalNewCipherFunc := newCipherFunc
		newCipherFunc = func(key []byte) (cipher.Block, error) {
			return nil, errors.New("mock cipher error")
		}
		defer func() { newCipherFunc = originalNewCipherFunc }()

		_, err := EncryptCBC(key, plaintext)
		if err == nil {
			t.Error("EncryptCBC should fail when newCipherFunc returns error")
		}
	})

	t.Run("EncryptCBC fails when randReader returns error", func(t *testing.T) {
		originalRandReader := randReader
		randReader = &failingReader{}
		defer func() { randReader = originalRandReader }()

		_, err := EncryptCBC(key, plaintext)
		if err == nil {
			t.Error("EncryptCBC should fail when randReader returns error")
		}
	})

	t.Run("EncryptCFB fails when newCipherFunc returns error", func(t *testing.T) {
		originalNewCipherFunc := newCipherFunc
		newCipherFunc = func(key []byte) (cipher.Block, error) {
			return nil, errors.New("mock cipher error")
		}
		defer func() { newCipherFunc = originalNewCipherFunc }()

		_, err := EncryptCFB(key, plaintext)
		if err == nil {
			t.Error("EncryptCFB should fail when newCipherFunc returns error")
		}
	})

	t.Run("EncryptCFB fails when randReader returns error", func(t *testing.T) {
		originalRandReader := randReader
		randReader = &failingReader{}
		defer func() { randReader = originalRandReader }()

		_, err := EncryptCFB(key, plaintext)
		if err == nil {
			t.Error("EncryptCFB should fail when randReader returns error")
		}
	})

	t.Run("EncryptCTR fails when newCipherFunc returns error", func(t *testing.T) {
		originalNewCipherFunc := newCipherFunc
		newCipherFunc = func(key []byte) (cipher.Block, error) {
			return nil, errors.New("mock cipher error")
		}
		defer func() { newCipherFunc = originalNewCipherFunc }()

		_, err := EncryptCTR(key, plaintext)
		if err == nil {
			t.Error("EncryptCTR should fail when newCipherFunc returns error")
		}
	})

	t.Run("EncryptCTR fails when randReader returns error", func(t *testing.T) {
		originalRandReader := randReader
		randReader = &failingReader{}
		defer func() { randReader = originalRandReader }()

		_, err := EncryptCTR(key, plaintext)
		if err == nil {
			t.Error("EncryptCTR should fail when randReader returns error")
		}
	})

	t.Run("EncryptOFB fails when newCipherFunc returns error", func(t *testing.T) {
		originalNewCipherFunc := newCipherFunc
		newCipherFunc = func(key []byte) (cipher.Block, error) {
			return nil, errors.New("mock cipher error")
		}
		defer func() { newCipherFunc = originalNewCipherFunc }()

		_, err := EncryptOFB(key, plaintext)
		if err == nil {
			t.Error("EncryptOFB should fail when newCipherFunc returns error")
		}
	})

	t.Run("EncryptOFB fails when randReader returns error", func(t *testing.T) {
		originalRandReader := randReader
		randReader = &failingReader{}
		defer func() { randReader = originalRandReader }()

		_, err := EncryptOFB(key, plaintext)
		if err == nil {
			t.Error("EncryptOFB should fail when randReader returns error")
		}
	})
}

// TestDecryptErrorPaths tests error paths in decryption functions using dependency injection
func TestDecryptErrorPaths(t *testing.T) {
	key := generateValidKey()
	plaintext := []byte("test message")

	t.Run("Decrypt fails when newCipherFunc returns error", func(t *testing.T) {
		// First create valid ciphertext
		ciphertext, _ := Encrypt(key, plaintext)

		// Then inject error
		originalNewCipherFunc := newCipherFunc
		newCipherFunc = func(key []byte) (cipher.Block, error) {
			return nil, errors.New("mock cipher error")
		}
		defer func() { newCipherFunc = originalNewCipherFunc }()

		_, err := Decrypt(key, ciphertext)
		if err == nil {
			t.Error("Decrypt should fail when newCipherFunc returns error")
		}
	})

	t.Run("Decrypt fails when newGCMFunc returns error", func(t *testing.T) {
		ciphertext, _ := Encrypt(key, plaintext)

		originalNewGCMFunc := newGCMFunc
		newGCMFunc = func(cipher cipher.Block) (cipher.AEAD, error) {
			return nil, errors.New("mock GCM error")
		}
		defer func() { newGCMFunc = originalNewGCMFunc }()

		_, err := Decrypt(key, ciphertext)
		if err == nil {
			t.Error("Decrypt should fail when newGCMFunc returns error")
		}
	})

	t.Run("DecryptECB fails when newCipherFunc returns error", func(t *testing.T) {
		ciphertext, _ := EncryptECB(key, plaintext)

		originalNewCipherFunc := newCipherFunc
		newCipherFunc = func(key []byte) (cipher.Block, error) {
			return nil, errors.New("mock cipher error")
		}
		defer func() { newCipherFunc = originalNewCipherFunc }()

		_, err := DecryptECB(key, ciphertext)
		if err == nil {
			t.Error("DecryptECB should fail when newCipherFunc returns error")
		}
	})

	t.Run("DecryptCBC fails when newCipherFunc returns error", func(t *testing.T) {
		ciphertext, _ := EncryptCBC(key, plaintext)

		originalNewCipherFunc := newCipherFunc
		newCipherFunc = func(key []byte) (cipher.Block, error) {
			return nil, errors.New("mock cipher error")
		}
		defer func() { newCipherFunc = originalNewCipherFunc }()

		_, err := DecryptCBC(key, ciphertext)
		if err == nil {
			t.Error("DecryptCBC should fail when newCipherFunc returns error")
		}
	})

	t.Run("DecryptCFB fails when newCipherFunc returns error", func(t *testing.T) {
		ciphertext, _ := EncryptCFB(key, plaintext)

		originalNewCipherFunc := newCipherFunc
		newCipherFunc = func(key []byte) (cipher.Block, error) {
			return nil, errors.New("mock cipher error")
		}
		defer func() { newCipherFunc = originalNewCipherFunc }()

		_, err := DecryptCFB(key, ciphertext)
		if err == nil {
			t.Error("DecryptCFB should fail when newCipherFunc returns error")
		}
	})

	t.Run("DecryptCTR fails when newCipherFunc returns error", func(t *testing.T) {
		ciphertext, _ := EncryptCTR(key, plaintext)

		originalNewCipherFunc := newCipherFunc
		newCipherFunc = func(key []byte) (cipher.Block, error) {
			return nil, errors.New("mock cipher error")
		}
		defer func() { newCipherFunc = originalNewCipherFunc }()

		_, err := DecryptCTR(key, ciphertext)
		if err == nil {
			t.Error("DecryptCTR should fail when newCipherFunc returns error")
		}
	})

	t.Run("DecryptOFB fails when newCipherFunc returns error", func(t *testing.T) {
		ciphertext, _ := EncryptOFB(key, plaintext)

		originalNewCipherFunc := newCipherFunc
		newCipherFunc = func(key []byte) (cipher.Block, error) {
			return nil, errors.New("mock cipher error")
		}
		defer func() { newCipherFunc = originalNewCipherFunc }()

		_, err := DecryptOFB(key, ciphertext)
		if err == nil {
			t.Error("DecryptOFB should fail when newCipherFunc returns error")
		}
	})
}

// failingReader is a mock io.Reader that always returns an error
type failingReader struct{}

func (r *failingReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("mock random error")
}

// BenchmarkAESEncrypt benchmarks all encryption modes
func BenchmarkAESEncryptGCM(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	for i := 0; i < b.N; i++ {
		_, _ = Encrypt(key, plaintext)
	}
}

func BenchmarkAESEncryptECB(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	for i := 0; i < b.N; i++ {
		_, _ = EncryptECB(key, plaintext)
	}
}

func BenchmarkAESEncryptCBC(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	for i := 0; i < b.N; i++ {
		_, _ = EncryptCBC(key, plaintext)
	}
}

func BenchmarkAESEncryptCFB(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	for i := 0; i < b.N; i++ {
		_, _ = EncryptCFB(key, plaintext)
	}
}

func BenchmarkAESEncryptCTR(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	for i := 0; i < b.N; i++ {
		_, _ = EncryptCTR(key, plaintext)
	}
}

func BenchmarkAESEncryptOFB(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	for i := 0; i < b.N; i++ {
		_, _ = EncryptOFB(key, plaintext)
	}
}

// BenchmarkAESDecrypt benchmarks all decryption modes
func BenchmarkAESDecryptGCM(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	ciphertext, _ := Encrypt(key, plaintext)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Decrypt(key, ciphertext)
	}
}

func BenchmarkAESDecryptECB(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	ciphertext, _ := EncryptECB(key, plaintext)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = DecryptECB(key, ciphertext)
	}
}

func BenchmarkAESDecryptCBC(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	ciphertext, _ := EncryptCBC(key, plaintext)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = DecryptCBC(key, ciphertext)
	}
}

func BenchmarkAESDecryptCFB(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	ciphertext, _ := EncryptCFB(key, plaintext)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = DecryptCFB(key, ciphertext)
	}
}

func BenchmarkAESDecryptCTR(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	ciphertext, _ := EncryptCTR(key, plaintext)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = DecryptCTR(key, ciphertext)
	}
}

func BenchmarkAESDecryptOFB(b *testing.B) {
	key := generateValidKey()
	plaintext := []byte("benchmark message for AES testing")
	ciphertext, _ := EncryptOFB(key, plaintext)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = DecryptOFB(key, ciphertext)
	}
}
