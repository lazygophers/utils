package cryptox

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"errors"
	"strings"
	"testing"
)

// generateValidDESKey generates a valid 8-byte DES key for testing
func generateValidDESKey() []byte {
	return []byte("12345678") // 8 bytes
}

// generateValid3DESKey generates a valid 24-byte 3DES key for testing
func generateValid3DESKey() []byte {
	return []byte("123456789012345678901234") // 24 bytes
}

// TestDESEncryptDecryptECB tests DES ECB encryption and decryption round-trip
func TestDESEncryptDecryptECB(t *testing.T) {
	key := generateValidDESKey()

	testCases := []struct {
		name      string
		plaintext []byte
	}{
		{"empty data", []byte("")},
		{"single byte", []byte("A")},
		{"short message", []byte("Hello")},
		{"medium message", []byte("The quick brown fox")},
		{"long message", bytes.Repeat([]byte("X"), 100)},
		{"binary data", []byte{0x00, 0x01, 0x02, 0x03, 0xFF, 0xFE}},
		{"unicode data", []byte("你好世界")},
		{"block aligned", make([]byte, des.BlockSize*3)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Encrypt
			ciphertext, err := DESEncryptECB(key, tc.plaintext)
			if err != nil {
				t.Fatalf("DESEncryptECB failed: %v", err)
			}

			// Decrypt
			decrypted, err := DESDecryptECB(key, ciphertext)
			if err != nil {
				t.Fatalf("DESDecryptECB failed: %v", err)
			}

			// Verify round-trip
			if !bytes.Equal(decrypted, tc.plaintext) {
				t.Errorf("Decrypted data doesn't match original.\nGot:      %v\nExpected: %v", decrypted, tc.plaintext)
			}
		})
	}
}

// TestDESEncryptDecryptECBErrors tests error cases for DES ECB mode
func TestDESEncryptDecryptECBErrors(t *testing.T) {
	validKey := generateValidDESKey()
	plaintext := []byte("test message")

	t.Run("DESEncryptECB with invalid key length", func(t *testing.T) {
		invalidKeys := [][]byte{
			[]byte("short"),      // too short
			[]byte("123456789"),  // 9 bytes
			[]byte("1234567890"), // 10 bytes
		}

		for _, key := range invalidKeys {
			_, err := DESEncryptECB(key, plaintext)
			if err == nil {
				t.Errorf("DESEncryptECB should fail with key length %d", len(key))
			}
			if err != nil && !strings.Contains(err.Error(), "invalid key length") {
				t.Errorf("Expected 'invalid key length' error, got: %v", err)
			}
		}
	})

	t.Run("DESDecryptECB with invalid key length", func(t *testing.T) {
		ciphertext, _ := DESEncryptECB(validKey, plaintext)
		invalidKey := []byte("short")
		_, err := DESDecryptECB(invalidKey, ciphertext)
		if err == nil || !strings.Contains(err.Error(), "invalid key length") {
			t.Errorf("Expected 'invalid key length' error, got: %v", err)
		}
	})

	t.Run("DESDecryptECB with invalid block size", func(t *testing.T) {
		invalidCiphertext := make([]byte, des.BlockSize+1)
		_, err := DESDecryptECB(validKey, invalidCiphertext)
		if err == nil {
			t.Error("DESDecryptECB should fail with ciphertext not multiple of block size")
		}
		if err != nil && !strings.Contains(err.Error(), "not a multiple of the block size") {
			t.Errorf("Expected block size error, got: %v", err)
		}
	})
}

// TestDESEncryptDecryptCBC tests DES CBC encryption and decryption round-trip
func TestDESEncryptDecryptCBC(t *testing.T) {
	key := generateValidDESKey()

	testCases := []struct {
		name      string
		plaintext []byte
	}{
		{"empty data", []byte("")},
		{"single byte", []byte("A")},
		{"short message", []byte("Hello")},
		{"medium message", []byte("The quick brown fox")},
		{"long message", bytes.Repeat([]byte("X"), 100)},
		{"binary data", []byte{0x00, 0x01, 0x02, 0x03, 0xFF, 0xFE}},
		{"unicode data", []byte("你好世界")},
		{"block aligned", make([]byte, des.BlockSize*3)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Encrypt
			ciphertext, err := DESEncryptCBC(key, tc.plaintext)
			if err != nil {
				t.Fatalf("DESEncryptCBC failed: %v", err)
			}

			// Decrypt
			decrypted, err := DESDecryptCBC(key, ciphertext)
			if err != nil {
				t.Fatalf("DESDecryptCBC failed: %v", err)
			}

			// Verify round-trip
			if !bytes.Equal(decrypted, tc.plaintext) {
				t.Errorf("Decrypted data doesn't match original.\nGot:      %v\nExpected: %v", decrypted, tc.plaintext)
			}
		})
	}
}

// TestDESEncryptDecryptCBCErrors tests error cases for DES CBC mode
func TestDESEncryptDecryptCBCErrors(t *testing.T) {
	validKey := generateValidDESKey()
	plaintext := []byte("test message")

	t.Run("DESEncryptCBC with invalid key length", func(t *testing.T) {
		invalidKey := []byte("short")
		_, err := DESEncryptCBC(invalidKey, plaintext)
		if err == nil || !strings.Contains(err.Error(), "invalid key length") {
			t.Errorf("Expected 'invalid key length' error, got: %v", err)
		}
	})

	t.Run("DESDecryptCBC with invalid key length", func(t *testing.T) {
		ciphertext, _ := DESEncryptCBC(validKey, plaintext)
		invalidKey := []byte("short")
		_, err := DESDecryptCBC(invalidKey, ciphertext)
		if err == nil || !strings.Contains(err.Error(), "invalid key length") {
			t.Errorf("Expected 'invalid key length' error, got: %v", err)
		}
	})

	t.Run("DESDecryptCBC with ciphertext too short", func(t *testing.T) {
		shortCiphertext := []byte("short")
		_, err := DESDecryptCBC(validKey, shortCiphertext)
		if err == nil {
			t.Error("DESDecryptCBC should fail with ciphertext too short")
		}
		if err != nil && !strings.Contains(err.Error(), "ciphertext too short") {
			t.Errorf("Expected 'ciphertext too short' error, got: %v", err)
		}
	})

	t.Run("DESDecryptCBC with invalid block size", func(t *testing.T) {
		// Create ciphertext with IV + invalid data length
		invalidCiphertext := make([]byte, des.BlockSize+1)
		_, err := DESDecryptCBC(validKey, invalidCiphertext)
		if err == nil {
			t.Error("DESDecryptCBC should fail with ciphertext not multiple of block size")
		}
		if err != nil && !strings.Contains(err.Error(), "not a multiple of the block size") {
			t.Errorf("Expected block size error, got: %v", err)
		}
	})

	t.Run("DESDecryptCBC with corrupted padding", func(t *testing.T) {
		ciphertext, _ := DESEncryptCBC(validKey, plaintext)
		// Corrupt the last block (where padding is)
		ciphertext[len(ciphertext)-1] ^= 0xFF
		_, err := DESDecryptCBC(validKey, ciphertext)
		if err == nil {
			t.Error("DESDecryptCBC should fail with corrupted padding")
		}
	})
}

// TestTripleDESEncryptDecryptECB tests 3DES ECB encryption and decryption round-trip
func TestTripleDESEncryptDecryptECB(t *testing.T) {
	key := generateValid3DESKey()

	testCases := []struct {
		name      string
		plaintext []byte
	}{
		{"empty data", []byte("")},
		{"single byte", []byte("A")},
		{"short message", []byte("Hello")},
		{"medium message", []byte("The quick brown fox")},
		{"long message", bytes.Repeat([]byte("X"), 100)},
		{"binary data", []byte{0x00, 0x01, 0x02, 0x03, 0xFF, 0xFE}},
		{"unicode data", []byte("你好世界")},
		{"block aligned", make([]byte, des.BlockSize*3)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Encrypt
			ciphertext, err := TripleDESEncryptECB(key, tc.plaintext)
			if err != nil {
				t.Fatalf("TripleDESEncryptECB failed: %v", err)
			}

			// Decrypt
			decrypted, err := TripleDESDecryptECB(key, ciphertext)
			if err != nil {
				t.Fatalf("TripleDESDecryptECB failed: %v", err)
			}

			// Verify round-trip
			if !bytes.Equal(decrypted, tc.plaintext) {
				t.Errorf("Decrypted data doesn't match original.\nGot:      %v\nExpected: %v", decrypted, tc.plaintext)
			}
		})
	}
}

// TestTripleDESEncryptDecryptECBErrors tests error cases for 3DES ECB mode
func TestTripleDESEncryptDecryptECBErrors(t *testing.T) {
	validKey := generateValid3DESKey()
	plaintext := []byte("test message")

	t.Run("TripleDESEncryptECB with invalid key length", func(t *testing.T) {
		invalidKeys := [][]byte{
			[]byte("short"),               // too short
			generateValidDESKey(),         // 8 bytes (DES key)
			[]byte("1234567890123456"),    // 16 bytes
			bytes.Repeat([]byte("X"), 32), // 32 bytes
		}

		for _, key := range invalidKeys {
			_, err := TripleDESEncryptECB(key, plaintext)
			if err == nil {
				t.Errorf("TripleDESEncryptECB should fail with key length %d", len(key))
			}
			if err != nil && !strings.Contains(err.Error(), "invalid key length") {
				t.Errorf("Expected 'invalid key length' error, got: %v", err)
			}
		}
	})

	t.Run("TripleDESDecryptECB with invalid key length", func(t *testing.T) {
		ciphertext, _ := TripleDESEncryptECB(validKey, plaintext)
		invalidKey := []byte("short")
		_, err := TripleDESDecryptECB(invalidKey, ciphertext)
		if err == nil || !strings.Contains(err.Error(), "invalid key length") {
			t.Errorf("Expected 'invalid key length' error, got: %v", err)
		}
	})

	t.Run("TripleDESDecryptECB with invalid block size", func(t *testing.T) {
		invalidCiphertext := make([]byte, des.BlockSize+1)
		_, err := TripleDESDecryptECB(validKey, invalidCiphertext)
		if err == nil {
			t.Error("TripleDESDecryptECB should fail with ciphertext not multiple of block size")
		}
		if err != nil && !strings.Contains(err.Error(), "not a multiple of the block size") {
			t.Errorf("Expected block size error, got: %v", err)
		}
	})
}

// TestTripleDESEncryptDecryptCBC tests 3DES CBC encryption and decryption round-trip
func TestTripleDESEncryptDecryptCBC(t *testing.T) {
	key := generateValid3DESKey()

	testCases := []struct {
		name      string
		plaintext []byte
	}{
		{"empty data", []byte("")},
		{"single byte", []byte("A")},
		{"short message", []byte("Hello")},
		{"medium message", []byte("The quick brown fox")},
		{"long message", bytes.Repeat([]byte("X"), 100)},
		{"binary data", []byte{0x00, 0x01, 0x02, 0x03, 0xFF, 0xFE}},
		{"unicode data", []byte("你好世界")},
		{"block aligned", make([]byte, des.BlockSize*3)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Encrypt
			ciphertext, err := TripleDESEncryptCBC(key, tc.plaintext)
			if err != nil {
				t.Fatalf("TripleDESEncryptCBC failed: %v", err)
			}

			// Decrypt
			decrypted, err := TripleDESDecryptCBC(key, ciphertext)
			if err != nil {
				t.Fatalf("TripleDESDecryptCBC failed: %v", err)
			}

			// Verify round-trip
			if !bytes.Equal(decrypted, tc.plaintext) {
				t.Errorf("Decrypted data doesn't match original.\nGot:      %v\nExpected: %v", decrypted, tc.plaintext)
			}
		})
	}
}

// TestTripleDESEncryptDecryptCBCErrors tests error cases for 3DES CBC mode
func TestTripleDESEncryptDecryptCBCErrors(t *testing.T) {
	validKey := generateValid3DESKey()
	plaintext := []byte("test message")

	t.Run("TripleDESEncryptCBC with invalid key length", func(t *testing.T) {
		invalidKey := []byte("short")
		_, err := TripleDESEncryptCBC(invalidKey, plaintext)
		if err == nil || !strings.Contains(err.Error(), "invalid key length") {
			t.Errorf("Expected 'invalid key length' error, got: %v", err)
		}
	})

	t.Run("TripleDESDecryptCBC with invalid key length", func(t *testing.T) {
		ciphertext, _ := TripleDESEncryptCBC(validKey, plaintext)
		invalidKey := []byte("short")
		_, err := TripleDESDecryptCBC(invalidKey, ciphertext)
		if err == nil || !strings.Contains(err.Error(), "invalid key length") {
			t.Errorf("Expected 'invalid key length' error, got: %v", err)
		}
	})

	t.Run("TripleDESDecryptCBC with ciphertext too short", func(t *testing.T) {
		shortCiphertext := []byte("short")
		_, err := TripleDESDecryptCBC(validKey, shortCiphertext)
		if err == nil {
			t.Error("TripleDESDecryptCBC should fail with ciphertext too short")
		}
		if err != nil && !strings.Contains(err.Error(), "ciphertext too short") {
			t.Errorf("Expected 'ciphertext too short' error, got: %v", err)
		}
	})

	t.Run("TripleDESDecryptCBC with invalid block size", func(t *testing.T) {
		// Create ciphertext with IV + invalid data length
		invalidCiphertext := make([]byte, des.BlockSize+1)
		_, err := TripleDESDecryptCBC(validKey, invalidCiphertext)
		if err == nil {
			t.Error("TripleDESDecryptCBC should fail with ciphertext not multiple of block size")
		}
		if err != nil && !strings.Contains(err.Error(), "not a multiple of the block size") {
			t.Errorf("Expected block size error, got: %v", err)
		}
	})

	t.Run("TripleDESDecryptCBC with corrupted padding", func(t *testing.T) {
		ciphertext, _ := TripleDESEncryptCBC(validKey, plaintext)
		// Corrupt the last block (where padding is)
		ciphertext[len(ciphertext)-1] ^= 0xFF
		_, err := TripleDESDecryptCBC(validKey, ciphertext)
		if err == nil {
			t.Error("TripleDESDecryptCBC should fail with corrupted padding")
		}
	})
}

// TestDESErrorPaths tests error paths using dependency injection
func TestDESErrorPaths(t *testing.T) {
	desKey := generateValidDESKey()
	tripleKey := generateValid3DESKey()
	plaintext := []byte("test message")

	t.Run("DESEncryptECB fails when desNewCipher returns error", func(t *testing.T) {
		originalDesNewCipher := desNewCipher
		desNewCipher = func(key []byte) (cipher.Block, error) {
			return nil, errors.New("mock cipher error")
		}
		defer func() { desNewCipher = originalDesNewCipher }()

		_, err := DESEncryptECB(desKey, plaintext)
		if err == nil || !strings.Contains(err.Error(), "mock cipher error") {
			t.Errorf("Expected mock cipher error, got: %v", err)
		}
	})

	t.Run("DESDecryptECB fails when desNewCipher returns error", func(t *testing.T) {
		ciphertext, _ := DESEncryptECB(desKey, plaintext)

		originalDesNewCipher := desNewCipher
		desNewCipher = func(key []byte) (cipher.Block, error) {
			return nil, errors.New("mock cipher error")
		}
		defer func() { desNewCipher = originalDesNewCipher }()

		_, err := DESDecryptECB(desKey, ciphertext)
		if err == nil || !strings.Contains(err.Error(), "mock cipher error") {
			t.Errorf("Expected mock cipher error, got: %v", err)
		}
	})

	t.Run("DESEncryptCBC fails when desNewCipher returns error", func(t *testing.T) {
		originalDesNewCipher := desNewCipher
		desNewCipher = func(key []byte) (cipher.Block, error) {
			return nil, errors.New("mock cipher error")
		}
		defer func() { desNewCipher = originalDesNewCipher }()

		_, err := DESEncryptCBC(desKey, plaintext)
		if err == nil || !strings.Contains(err.Error(), "mock cipher error") {
			t.Errorf("Expected mock cipher error, got: %v", err)
		}
	})

	t.Run("DESEncryptCBC fails when desRandReader returns error", func(t *testing.T) {
		originalDesRandReader := desRandReader
		desRandReader = &desFailingReader{}
		defer func() { desRandReader = originalDesRandReader }()

		_, err := DESEncryptCBC(desKey, plaintext)
		if err == nil || !strings.Contains(err.Error(), "mock random error") {
			t.Errorf("Expected mock random error, got: %v", err)
		}
	})

	t.Run("DESDecryptCBC fails when desNewCipher returns error", func(t *testing.T) {
		ciphertext, _ := DESEncryptCBC(desKey, plaintext)

		originalDesNewCipher := desNewCipher
		desNewCipher = func(key []byte) (cipher.Block, error) {
			return nil, errors.New("mock cipher error")
		}
		defer func() { desNewCipher = originalDesNewCipher }()

		_, err := DESDecryptCBC(desKey, ciphertext)
		if err == nil || !strings.Contains(err.Error(), "mock cipher error") {
			t.Errorf("Expected mock cipher error, got: %v", err)
		}
	})

	t.Run("TripleDESEncryptECB fails when desNewTripleDES returns error", func(t *testing.T) {
		originalDesNewTripleDES := desNewTripleDES
		desNewTripleDES = func(key []byte) (cipher.Block, error) {
			return nil, errors.New("mock 3DES error")
		}
		defer func() { desNewTripleDES = originalDesNewTripleDES }()

		_, err := TripleDESEncryptECB(tripleKey, plaintext)
		if err == nil || !strings.Contains(err.Error(), "mock 3DES error") {
			t.Errorf("Expected mock 3DES error, got: %v", err)
		}
	})

	t.Run("TripleDESDecryptECB fails when desNewTripleDES returns error", func(t *testing.T) {
		ciphertext, _ := TripleDESEncryptECB(tripleKey, plaintext)

		originalDesNewTripleDES := desNewTripleDES
		desNewTripleDES = func(key []byte) (cipher.Block, error) {
			return nil, errors.New("mock 3DES error")
		}
		defer func() { desNewTripleDES = originalDesNewTripleDES }()

		_, err := TripleDESDecryptECB(tripleKey, ciphertext)
		if err == nil || !strings.Contains(err.Error(), "mock 3DES error") {
			t.Errorf("Expected mock 3DES error, got: %v", err)
		}
	})

	t.Run("TripleDESEncryptCBC fails when desNewTripleDES returns error", func(t *testing.T) {
		originalDesNewTripleDES := desNewTripleDES
		desNewTripleDES = func(key []byte) (cipher.Block, error) {
			return nil, errors.New("mock 3DES error")
		}
		defer func() { desNewTripleDES = originalDesNewTripleDES }()

		_, err := TripleDESEncryptCBC(tripleKey, plaintext)
		if err == nil || !strings.Contains(err.Error(), "mock 3DES error") {
			t.Errorf("Expected mock 3DES error, got: %v", err)
		}
	})

	t.Run("TripleDESEncryptCBC fails when desRandReader returns error", func(t *testing.T) {
		originalDesRandReader := desRandReader
		desRandReader = &desFailingReader{}
		defer func() { desRandReader = originalDesRandReader }()

		_, err := TripleDESEncryptCBC(tripleKey, plaintext)
		if err == nil || !strings.Contains(err.Error(), "mock random error") {
			t.Errorf("Expected mock random error, got: %v", err)
		}
	})

	t.Run("TripleDESDecryptCBC fails when desNewTripleDES returns error", func(t *testing.T) {
		ciphertext, _ := TripleDESEncryptCBC(tripleKey, plaintext)

		originalDesNewTripleDES := desNewTripleDES
		desNewTripleDES = func(key []byte) (cipher.Block, error) {
			return nil, errors.New("mock 3DES error")
		}
		defer func() { desNewTripleDES = originalDesNewTripleDES }()

		_, err := TripleDESDecryptCBC(tripleKey, ciphertext)
		if err == nil || !strings.Contains(err.Error(), "mock 3DES error") {
			t.Errorf("Expected mock 3DES error, got: %v", err)
		}
	})
}

// desFailingReader is a mock io.Reader that always returns an error
type desFailingReader struct{}

func (r *desFailingReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("mock random error")
}

// TestDESCBCConsistency tests that CBC mode produces different ciphertexts with random IV
func TestDESCBCConsistency(t *testing.T) {
	desKey := generateValidDESKey()
	tripleKey := generateValid3DESKey()
	plaintext := []byte("test message for consistency check")

	t.Run("DES CBC produces different ciphertexts", func(t *testing.T) {
		ct1, _ := DESEncryptCBC(desKey, plaintext)
		ct2, _ := DESEncryptCBC(desKey, plaintext)
		if bytes.Equal(ct1, ct2) {
			t.Error("DES CBC should produce different ciphertexts with random IV")
		}
	})

	t.Run("3DES CBC produces different ciphertexts", func(t *testing.T) {
		ct1, _ := TripleDESEncryptCBC(tripleKey, plaintext)
		ct2, _ := TripleDESEncryptCBC(tripleKey, plaintext)
		if bytes.Equal(ct1, ct2) {
			t.Error("3DES CBC should produce different ciphertexts with random IV")
		}
	})
}

// TestDESWithDifferentKeys tests that different keys produce different ciphertexts
func TestDESWithDifferentKeys(t *testing.T) {
	desKey1 := []byte("12345678")
	desKey2 := []byte("87654321")
	tripleKey1 := []byte("123456789012345678901234")
	tripleKey2 := []byte("432109876543210987654321")
	plaintext := []byte("test message")

	t.Run("DES ECB with different keys", func(t *testing.T) {
		ct1, _ := DESEncryptECB(desKey1, plaintext)
		ct2, _ := DESEncryptECB(desKey2, plaintext)
		if bytes.Equal(ct1, ct2) {
			t.Error("Different DES keys should produce different ciphertexts in ECB mode")
		}
	})

	t.Run("3DES ECB with different keys", func(t *testing.T) {
		ct1, _ := TripleDESEncryptECB(tripleKey1, plaintext)
		ct2, _ := TripleDESEncryptECB(tripleKey2, plaintext)
		if bytes.Equal(ct1, ct2) {
			t.Error("Different 3DES keys should produce different ciphertexts in ECB mode")
		}
	})
}

// TestDESvsTripleDES tests that DES and 3DES produce different results
func TestDESvsTripleDES(t *testing.T) {
	// Use different keys to ensure different results
	desKey := []byte("12345678")
	tripleKey := []byte("123456789012345678901234") // Different key parts
	plaintext := []byte("test message")

	t.Run("DES and 3DES ECB produce different ciphertexts", func(t *testing.T) {
		desCT, _ := DESEncryptECB(desKey, plaintext)
		tripleCT, _ := TripleDESEncryptECB(tripleKey, plaintext)
		// DES and 3DES should produce different ciphertexts
		// (they have different key lengths and algorithm structures)
		if bytes.Equal(desCT, tripleCT) {
			t.Error("DES and 3DES should produce different ciphertexts")
		}
	})
}

// BenchmarkDES benchmarks all DES encryption/decryption modes
func BenchmarkDESEncryptECB(b *testing.B) {
	key := generateValidDESKey()
	plaintext := []byte("benchmark message for DES testing")
	for i := 0; i < b.N; i++ {
		_, _ = DESEncryptECB(key, plaintext)
	}
}

func BenchmarkDESDecryptECB(b *testing.B) {
	key := generateValidDESKey()
	plaintext := []byte("benchmark message for DES testing")
	ciphertext, _ := DESEncryptECB(key, plaintext)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = DESDecryptECB(key, ciphertext)
	}
}

func BenchmarkDESEncryptCBC(b *testing.B) {
	key := generateValidDESKey()
	plaintext := []byte("benchmark message for DES testing")
	for i := 0; i < b.N; i++ {
		_, _ = DESEncryptCBC(key, plaintext)
	}
}

func BenchmarkDESDecryptCBC(b *testing.B) {
	key := generateValidDESKey()
	plaintext := []byte("benchmark message for DES testing")
	ciphertext, _ := DESEncryptCBC(key, plaintext)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = DESDecryptCBC(key, ciphertext)
	}
}

func BenchmarkTripleDESEncryptECB(b *testing.B) {
	key := generateValid3DESKey()
	plaintext := []byte("benchmark message for 3DES testing")
	for i := 0; i < b.N; i++ {
		_, _ = TripleDESEncryptECB(key, plaintext)
	}
}

func BenchmarkTripleDESDecryptECB(b *testing.B) {
	key := generateValid3DESKey()
	plaintext := []byte("benchmark message for 3DES testing")
	ciphertext, _ := TripleDESEncryptECB(key, plaintext)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = TripleDESDecryptECB(key, ciphertext)
	}
}

func BenchmarkTripleDESEncryptCBC(b *testing.B) {
	key := generateValid3DESKey()
	plaintext := []byte("benchmark message for 3DES testing")
	for i := 0; i < b.N; i++ {
		_, _ = TripleDESEncryptCBC(key, plaintext)
	}
}

func BenchmarkTripleDESDecryptCBC(b *testing.B) {
	key := generateValid3DESKey()
	plaintext := []byte("benchmark message for 3DES testing")
	ciphertext, _ := TripleDESEncryptCBC(key, plaintext)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = TripleDESDecryptCBC(key, ciphertext)
	}
}
