package cryptox

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"errors"
	"testing"
)

// Test data
const (
	testDESMessage       = "Hello, DES encryption!"
	testTripleDESMessage = "Hello, 3DES encryption test message!"
)

var (
	desKey8  = []byte("12345678")                 // 8 bytes for DES
	desKey24 = []byte("123456789012345678901234") // 24 bytes for 3DES
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

// ==== MERGED FROM des_100_coverage_test.go ====

// FailingDESReader implements io.Reader but always fails
type FailingDESReader struct{}

func (fr FailingDESReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("simulated random reader failure")
}

// FailingDESFunc simulates des.NewCipher failure
func FailingDESFunc(key []byte) (cipher.Block, error) {
	return nil, errors.New("simulated des.NewCipher failure")
}

// FailingTripleDESFunc simulates des.NewTripleDESCipher failure
func FailingTripleDESFunc(key []byte) (cipher.Block, error) {
	return nil, errors.New("simulated des.NewTripleDESCipher failure")
}

// TestDES_100PercentCoverage triggers all error paths using dependency injection
func TestDES_100PercentCoverage(t *testing.T) {
	// Save original functions
	originalNewDESCipher := desNewCipher
	originalNewTripleDESCipher := desNewTripleDES
	originalDESRandReader := desRandReader

	// Restore original functions after test
	defer func() {
		desNewCipher = originalNewDESCipher
		desNewTripleDES = originalNewTripleDESCipher
		desRandReader = originalDESRandReader
	}()

	key8 := make([]byte, 8)   // DES key
	key24 := make([]byte, 24) // 3DES key
	plaintext := []byte("test plaintext for DES")

	// Test 1: Trigger des.NewCipher failure in DES functions
	desNewCipher = FailingDESFunc
	desNewTripleDES = originalNewTripleDESCipher
	desRandReader = originalDESRandReader

	_, err := DESEncryptECB(key8, plaintext)
	if err == nil {
		t.Error("Expected des.NewCipher error in DESEncryptECB")
	}

	_, err = DESDecryptECB(key8, []byte("fake ciphertext"))
	if err == nil {
		t.Error("Expected des.NewCipher error in DESDecryptECB")
	}

	_, err = DESEncryptCBC(key8, plaintext)
	if err == nil {
		t.Error("Expected des.NewCipher error in DESEncryptCBC")
	}

	_, err = DESDecryptCBC(key8, []byte("fake ciphertext that's long enough"))
	if err == nil {
		t.Error("Expected des.NewCipher error in DESDecryptCBC")
	}

	// Test 2: Trigger des.NewTripleDESCipher failure in 3DES functions
	desNewCipher = originalNewDESCipher
	desNewTripleDES = FailingTripleDESFunc
	desRandReader = originalDESRandReader

	_, err = TripleDESEncryptECB(key24, plaintext)
	if err == nil {
		t.Error("Expected des.NewTripleDESCipher error in TripleDESEncryptECB")
	}

	_, err = TripleDESDecryptECB(key24, []byte("fake ciphertext"))
	if err == nil {
		t.Error("Expected des.NewTripleDESCipher error in TripleDESDecryptECB")
	}

	_, err = TripleDESEncryptCBC(key24, plaintext)
	if err == nil {
		t.Error("Expected des.NewTripleDESCipher error in TripleDESEncryptCBC")
	}

	_, err = TripleDESDecryptCBC(key24, []byte("fake ciphertext that's long enough"))
	if err == nil {
		t.Error("Expected des.NewTripleDESCipher error in TripleDESDecryptCBC")
	}

	// Test 3: Trigger rand.Reader failure in functions that use random IV
	desNewCipher = originalNewDESCipher
	desNewTripleDES = originalNewTripleDESCipher
	desRandReader = FailingDESReader{}

	_, err = DESEncryptCBC(key8, plaintext)
	if err == nil {
		t.Error("Expected rand.Reader error in DESEncryptCBC")
	}

	_, err = TripleDESEncryptCBC(key24, plaintext)
	if err == nil {
		t.Error("Expected rand.Reader error in TripleDESEncryptCBC")
	}
}

// TestInvalidCiphertextForAllModes tests ciphertext validation for all DES modes
func TestInvalidCiphertextForAllModes(t *testing.T) {
	// Test with valid keys
	key8 := desKey8
	key24 := desKey24

	// Test ECB with non-block-size ciphertext
	invalidECBCiphertext := make([]byte, des.BlockSize+1) // 9 bytes (not multiple of 8)

	_, err := DESDecryptECB(key8, invalidECBCiphertext)
	if err == nil || err.Error() != "ciphertext is not a multiple of the block size" {
		t.Error("Expected block size error for DES ECB")
	}

	_, err = TripleDESDecryptECB(key24, invalidECBCiphertext)
	if err == nil || err.Error() != "ciphertext is not a multiple of the block size" {
		t.Error("Expected block size error for 3DES ECB")
	}

	// Test CBC with non-block-size data portion (after IV)
	invalidCBCCiphertext := make([]byte, des.BlockSize+des.BlockSize+1) // IV + 9 bytes data

	_, err = DESDecryptCBC(key8, invalidCBCCiphertext)
	if err == nil || err.Error() != "ciphertext is not a multiple of the block size" {
		t.Error("Expected block size error for DES CBC")
	}

	_, err = TripleDESDecryptCBC(key24, invalidCBCCiphertext)
	if err == nil || err.Error() != "ciphertext is not a multiple of the block size" {
		t.Error("Expected block size error for 3DES CBC")
	}
}

// TestRoundTripConsistency ensures all DES modes produce consistent results
func TestRoundTripConsistency(t *testing.T) {
	testCases := []struct {
		name      string
		plaintext []byte
	}{
		{"empty", []byte("")},
		{"single_block", []byte("12345678")},
		{"multiple_blocks", []byte("This is a test message that spans multiple DES blocks for testing.")},
		{"binary_data", []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test DES ECB
			ciphertext, err := DESEncryptECB(desKey8, tc.plaintext)
			if err != nil {
				t.Fatalf("DES ECB encryption failed: %v", err)
			}
			decrypted, err := DESDecryptECB(desKey8, ciphertext)
			if err != nil {
				t.Fatalf("DES ECB decryption failed: %v", err)
			}
			if !bytes.Equal(tc.plaintext, decrypted) {
				t.Error("DES ECB round-trip failed")
			}

			// Test DES CBC
			ciphertext, err = DESEncryptCBC(desKey8, tc.plaintext)
			if err != nil {
				t.Fatalf("DES CBC encryption failed: %v", err)
			}
			decrypted, err = DESDecryptCBC(desKey8, ciphertext)
			if err != nil {
				t.Fatalf("DES CBC decryption failed: %v", err)
			}
			if !bytes.Equal(tc.plaintext, decrypted) {
				t.Error("DES CBC round-trip failed")
			}

			// Test 3DES ECB
			ciphertext, err = TripleDESEncryptECB(desKey24, tc.plaintext)
			if err != nil {
				t.Fatalf("3DES ECB encryption failed: %v", err)
			}
			decrypted, err = TripleDESDecryptECB(desKey24, ciphertext)
			if err != nil {
				t.Fatalf("3DES ECB decryption failed: %v", err)
			}
			if !bytes.Equal(tc.plaintext, decrypted) {
				t.Error("3DES ECB round-trip failed")
			}

			// Test 3DES CBC
			ciphertext, err = TripleDESEncryptCBC(desKey24, tc.plaintext)
			if err != nil {
				t.Fatalf("3DES CBC encryption failed: %v", err)
			}
			decrypted, err = TripleDESDecryptCBC(desKey24, ciphertext)
			if err != nil {
				t.Fatalf("3DES CBC decryption failed: %v", err)
			}
			if !bytes.Equal(tc.plaintext, decrypted) {
				t.Error("3DES CBC round-trip failed")
			}
		})
	}
}
