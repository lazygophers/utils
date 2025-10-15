package cryptox

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"hash"
	"math/big"
	"testing"
)

// TestGenerateECDHKey tests ECDH key generation with various curves
func TestGenerateECDHKey(t *testing.T) {
	curves := []struct {
		name  string
		curve elliptic.Curve
	}{
		{"P-224", elliptic.P224()},
		{"P-256", elliptic.P256()},
		{"P-384", elliptic.P384()},
		{"P-521", elliptic.P521()},
	}

	for _, tc := range curves {
		t.Run(tc.name, func(t *testing.T) {
			keyPair, err := GenerateECDHKey(tc.curve)
			if err != nil {
				t.Fatalf("GenerateECDHKey failed: %v", err)
			}

			if keyPair == nil {
				t.Fatal("Key pair should not be nil")
			}
			if keyPair.PrivateKey == nil {
				t.Fatal("Private key should not be nil")
			}
			if keyPair.PublicKey == nil {
				t.Fatal("Public key should not be nil")
			}

			// Verify curve matches
			if keyPair.PrivateKey.Curve != tc.curve {
				t.Error("Private key curve doesn't match requested curve")
			}
			if keyPair.PublicKey.Curve != tc.curve {
				t.Error("Public key curve doesn't match requested curve")
			}

			// Verify public key is on curve
			if !tc.curve.IsOnCurve(keyPair.PublicKey.X, keyPair.PublicKey.Y) {
				t.Error("Public key is not on the curve")
			}

			// Verify key pair is valid
			if err := ValidateECDHKeyPair(keyPair); err != nil {
				t.Errorf("Generated key pair is not valid: %v", err)
			}
		})
	}
}

// TestGenerateECDHKeyErrors tests error cases for key generation
func TestGenerateECDHKeyErrors(t *testing.T) {
	t.Run("nil curve", func(t *testing.T) {
		_, err := GenerateECDHKey(nil)
		if err == nil {
			t.Error("GenerateECDHKey should fail with nil curve")
		}
	})
}

// TestGenerateECDHPxxxKey tests specific curve key generation functions
func TestGenerateECDHPxxxKey(t *testing.T) {
	testCases := []struct {
		name     string
		genFunc  func() (*ECDHKeyPair, error)
		expected elliptic.Curve
	}{
		{"P-256", GenerateECDHP256Key, elliptic.P256()},
		{"P-384", GenerateECDHP384Key, elliptic.P384()},
		{"P-521", GenerateECDHP521Key, elliptic.P521()},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			keyPair, err := tc.genFunc()
			if err != nil {
				t.Fatalf("Key generation failed: %v", err)
			}

			if keyPair.PrivateKey.Curve != tc.expected {
				t.Errorf("Expected curve %v, got %v", tc.expected, keyPair.PrivateKey.Curve)
			}

			if err := ValidateECDHKeyPair(keyPair); err != nil {
				t.Errorf("Generated key pair is not valid: %v", err)
			}
		})
	}
}

// TestECDHComputeShared tests shared secret computation
func TestECDHComputeShared(t *testing.T) {
	curves := []struct {
		name  string
		curve elliptic.Curve
	}{
		{"P-256", elliptic.P256()},
		{"P-384", elliptic.P384()},
		{"P-521", elliptic.P521()},
	}

	for _, tc := range curves {
		t.Run(tc.name, func(t *testing.T) {
			// Generate two key pairs
			alice, err := GenerateECDHKey(tc.curve)
			if err != nil {
				t.Fatalf("Failed to generate Alice's key: %v", err)
			}

			bob, err := GenerateECDHKey(tc.curve)
			if err != nil {
				t.Fatalf("Failed to generate Bob's key: %v", err)
			}

			// Compute shared secrets
			aliceShared, err := ECDHComputeShared(alice.PrivateKey, bob.PublicKey)
			if err != nil {
				t.Fatalf("Alice's shared secret computation failed: %v", err)
			}

			bobShared, err := ECDHComputeShared(bob.PrivateKey, alice.PublicKey)
			if err != nil {
				t.Fatalf("Bob's shared secret computation failed: %v", err)
			}

			// Verify shared secrets match
			if !bytes.Equal(aliceShared, bobShared) {
				t.Error("Shared secrets don't match")
			}

			// Verify shared secret is not empty
			if len(aliceShared) == 0 {
				t.Error("Shared secret should not be empty")
			}
		})
	}
}

// TestECDHComputeSharedErrors tests error cases for shared secret computation
func TestECDHComputeSharedErrors(t *testing.T) {
	alice, _ := GenerateECDHP256Key()
	bob, _ := GenerateECDHP256Key()

	t.Run("nil private key", func(t *testing.T) {
		_, err := ECDHComputeShared(nil, bob.PublicKey)
		if err == nil {
			t.Error("ECDHComputeShared should fail with nil private key")
		}
	})

	t.Run("nil public key", func(t *testing.T) {
		_, err := ECDHComputeShared(alice.PrivateKey, nil)
		if err == nil {
			t.Error("ECDHComputeShared should fail with nil public key")
		}
	})

	t.Run("public key not on curve", func(t *testing.T) {
		invalidPublicKey := &ecdsa.PublicKey{
			Curve: elliptic.P256(),
			X:     big.NewInt(1),
			Y:     big.NewInt(1),
		}
		_, err := ECDHComputeShared(alice.PrivateKey, invalidPublicKey)
		if err == nil {
			t.Error("ECDHComputeShared should fail with public key not on curve")
		}
	})

	t.Run("curve mismatch", func(t *testing.T) {
		alice256, _ := GenerateECDHP256Key()
		bob384, _ := GenerateECDHP384Key()
		_, err := ECDHComputeShared(alice256.PrivateKey, bob384.PublicKey)
		if err == nil {
			t.Error("ECDHComputeShared should fail with curve mismatch")
		}
	})
}

// TestECDHComputeSharedWithKDF tests shared secret computation with KDF
func TestECDHComputeSharedWithKDF(t *testing.T) {
	alice, _ := GenerateECDHP256Key()
	bob, _ := GenerateECDHP256Key()

	testCases := []struct {
		name      string
		keyLength int
		kdf       func() hash.Hash
	}{
		{"SHA256 - 16 bytes", 16, sha256.New},
		{"SHA256 - 32 bytes", 32, sha256.New},
		{"SHA256 - 64 bytes", 64, sha256.New},
		{"SHA512 - 32 bytes", 32, sha512.New},
		{"SHA512 - 128 bytes", 128, sha512.New},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Compute shared secrets with KDF
			aliceShared, err := ECDHComputeSharedWithKDF(alice.PrivateKey, bob.PublicKey, tc.keyLength, tc.kdf)
			if err != nil {
				t.Fatalf("Alice's KDF computation failed: %v", err)
			}

			bobShared, err := ECDHComputeSharedWithKDF(bob.PrivateKey, alice.PublicKey, tc.keyLength, tc.kdf)
			if err != nil {
				t.Fatalf("Bob's KDF computation failed: %v", err)
			}

			// Verify shared secrets match
			if !bytes.Equal(aliceShared, bobShared) {
				t.Error("Shared secrets with KDF don't match")
			}

			// Verify key length
			if len(aliceShared) != tc.keyLength {
				t.Errorf("Expected key length %d, got %d", tc.keyLength, len(aliceShared))
			}
		})
	}
}

// TestECDHComputeSharedWithKDFErrors tests error cases for KDF computation
func TestECDHComputeSharedWithKDFErrors(t *testing.T) {
	alice, _ := GenerateECDHP256Key()
	bob, _ := GenerateECDHP256Key()

	t.Run("invalid key length", func(t *testing.T) {
		_, err := ECDHComputeSharedWithKDF(alice.PrivateKey, bob.PublicKey, 0, sha256.New)
		if err == nil {
			t.Error("Should fail with key length 0")
		}

		_, err = ECDHComputeSharedWithKDF(alice.PrivateKey, bob.PublicKey, -1, sha256.New)
		if err == nil {
			t.Error("Should fail with negative key length")
		}
	})

	t.Run("nil KDF function", func(t *testing.T) {
		_, err := ECDHComputeSharedWithKDF(alice.PrivateKey, bob.PublicKey, 32, nil)
		if err == nil {
			t.Error("Should fail with nil KDF function")
		}
	})

	t.Run("underlying ECDHComputeShared error propagates", func(t *testing.T) {
		// Use nil public key to trigger underlying error
		_, err := ECDHComputeSharedWithKDF(alice.PrivateKey, nil, 32, sha256.New)
		if err == nil {
			t.Error("Should fail when ECDHComputeShared fails")
		}
	})
}

// TestECDHComputeSharedSHA256 tests SHA256 KDF variant
func TestECDHComputeSharedSHA256(t *testing.T) {
	alice, _ := GenerateECDHP256Key()
	bob, _ := GenerateECDHP256Key()

	keyLengths := []int{16, 32, 48, 64, 128}

	for _, keyLength := range keyLengths {
		t.Run("keylen_"+string(rune(keyLength+'0')), func(t *testing.T) {
			aliceShared, err := ECDHComputeSharedSHA256(alice.PrivateKey, bob.PublicKey, keyLength)
			if err != nil {
				t.Fatalf("Failed: %v", err)
			}

			bobShared, err := ECDHComputeSharedSHA256(bob.PrivateKey, alice.PublicKey, keyLength)
			if err != nil {
				t.Fatalf("Failed: %v", err)
			}

			if !bytes.Equal(aliceShared, bobShared) {
				t.Error("Shared secrets don't match")
			}

			if len(aliceShared) != keyLength {
				t.Errorf("Expected length %d, got %d", keyLength, len(aliceShared))
			}
		})
	}
}

// TestECDHKeyExchange tests full key exchange
func TestECDHKeyExchange(t *testing.T) {
	alice, _ := GenerateECDHP256Key()
	bob, _ := GenerateECDHP256Key()

	keyLength := 32

	aliceShared, err := ECDHKeyExchange(alice.PrivateKey, bob.PublicKey, keyLength)
	if err != nil {
		t.Fatalf("Alice's key exchange failed: %v", err)
	}

	bobShared, err := ECDHKeyExchange(bob.PrivateKey, alice.PublicKey, keyLength)
	if err != nil {
		t.Fatalf("Bob's key exchange failed: %v", err)
	}

	if !bytes.Equal(aliceShared, bobShared) {
		t.Error("Key exchange results don't match")
	}

	if len(aliceShared) != keyLength {
		t.Errorf("Expected key length %d, got %d", keyLength, len(aliceShared))
	}
}

// TestValidateECDHKeyPair tests key pair validation
func TestValidateECDHKeyPair(t *testing.T) {
	t.Run("valid key pair", func(t *testing.T) {
		keyPair, _ := GenerateECDHP256Key()
		if err := ValidateECDHKeyPair(keyPair); err != nil {
			t.Errorf("Valid key pair should pass validation: %v", err)
		}
	})

	t.Run("nil key pair", func(t *testing.T) {
		if err := ValidateECDHKeyPair(nil); err == nil {
			t.Error("Should fail with nil key pair")
		}
	})

	t.Run("nil private key", func(t *testing.T) {
		keyPair, _ := GenerateECDHP256Key()
		keyPair.PrivateKey = nil
		if err := ValidateECDHKeyPair(keyPair); err == nil {
			t.Error("Should fail with nil private key")
		}
	})

	t.Run("nil public key", func(t *testing.T) {
		keyPair, _ := GenerateECDHP256Key()
		keyPair.PublicKey = nil
		if err := ValidateECDHKeyPair(keyPair); err == nil {
			t.Error("Should fail with nil public key")
		}
	})

	t.Run("public key not on curve", func(t *testing.T) {
		keyPair, _ := GenerateECDHP256Key()
		keyPair.PublicKey.X = big.NewInt(1)
		keyPair.PublicKey.Y = big.NewInt(1)
		if err := ValidateECDHKeyPair(keyPair); err == nil {
			t.Error("Should fail with public key not on curve")
		}
	})

	t.Run("curve mismatch", func(t *testing.T) {
		keyPair, _ := GenerateECDHP256Key()
		keyPair.PublicKey.Curve = elliptic.P384()
		if err := ValidateECDHKeyPair(keyPair); err == nil {
			t.Error("Should fail with curve mismatch")
		}
	})

	t.Run("public key doesn't match private key", func(t *testing.T) {
		keyPair1, _ := GenerateECDHP256Key()
		keyPair2, _ := GenerateECDHP256Key()
		keyPair1.PublicKey = keyPair2.PublicKey
		if err := ValidateECDHKeyPair(keyPair1); err == nil {
			t.Error("Should fail when public key doesn't match private key")
		}
	})
}

// TestECDHPublicKeyFromCoordinates tests creating public key from coordinates
func TestECDHPublicKeyFromCoordinates(t *testing.T) {
	t.Run("valid coordinates", func(t *testing.T) {
		original, _ := GenerateECDHP256Key()
		x, y, _ := ECDHPublicKeyToCoordinates(original.PublicKey)

		reconstructed, err := ECDHPublicKeyFromCoordinates(elliptic.P256(), x, y)
		if err != nil {
			t.Fatalf("Failed to create public key from coordinates: %v", err)
		}

		if reconstructed.X.Cmp(original.PublicKey.X) != 0 {
			t.Error("Reconstructed X coordinate doesn't match")
		}
		if reconstructed.Y.Cmp(original.PublicKey.Y) != 0 {
			t.Error("Reconstructed Y coordinate doesn't match")
		}
	})

	t.Run("nil curve", func(t *testing.T) {
		_, err := ECDHPublicKeyFromCoordinates(nil, big.NewInt(1), big.NewInt(1))
		if err == nil {
			t.Error("Should fail with nil curve")
		}
	})

	t.Run("nil x coordinate", func(t *testing.T) {
		_, err := ECDHPublicKeyFromCoordinates(elliptic.P256(), nil, big.NewInt(1))
		if err == nil {
			t.Error("Should fail with nil x coordinate")
		}
	})

	t.Run("nil y coordinate", func(t *testing.T) {
		_, err := ECDHPublicKeyFromCoordinates(elliptic.P256(), big.NewInt(1), nil)
		if err == nil {
			t.Error("Should fail with nil y coordinate")
		}
	})

	t.Run("point not on curve", func(t *testing.T) {
		_, err := ECDHPublicKeyFromCoordinates(elliptic.P256(), big.NewInt(1), big.NewInt(1))
		if err == nil {
			t.Error("Should fail with point not on curve")
		}
	})
}

// TestECDHPublicKeyToCoordinates tests extracting coordinates from public key
func TestECDHPublicKeyToCoordinates(t *testing.T) {
	t.Run("valid public key", func(t *testing.T) {
		keyPair, _ := GenerateECDHP256Key()
		x, y, err := ECDHPublicKeyToCoordinates(keyPair.PublicKey)
		if err != nil {
			t.Fatalf("Failed to extract coordinates: %v", err)
		}

		if x.Cmp(keyPair.PublicKey.X) != 0 {
			t.Error("X coordinate doesn't match")
		}
		if y.Cmp(keyPair.PublicKey.Y) != 0 {
			t.Error("Y coordinate doesn't match")
		}

		// Verify coordinates are copies, not references
		x.Add(x, big.NewInt(1))
		if x.Cmp(keyPair.PublicKey.X) == 0 {
			t.Error("X coordinate should be a copy, not a reference")
		}
	})

	t.Run("nil public key", func(t *testing.T) {
		_, _, err := ECDHPublicKeyToCoordinates(nil)
		if err == nil {
			t.Error("Should fail with nil public key")
		}
	})
}

// TestECDHSharedSecretTest tests the shared secret testing function
func TestECDHSharedSecretTest(t *testing.T) {
	t.Run("valid key exchange", func(t *testing.T) {
		alice, _ := GenerateECDHP256Key()
		bob, _ := GenerateECDHP256Key()

		match, err := ECDHSharedSecretTest(alice, bob)
		if err != nil {
			t.Fatalf("Shared secret test failed: %v", err)
		}
		if !match {
			t.Error("Shared secrets should match")
		}
	})

	t.Run("nil key pair 1", func(t *testing.T) {
		bob, _ := GenerateECDHP256Key()
		_, err := ECDHSharedSecretTest(nil, bob)
		if err == nil {
			t.Error("Should fail with nil key pair")
		}
	})

	t.Run("nil key pair 2", func(t *testing.T) {
		alice, _ := GenerateECDHP256Key()
		_, err := ECDHSharedSecretTest(alice, nil)
		if err == nil {
			t.Error("Should fail with nil key pair")
		}
	})

	t.Run("different curves", func(t *testing.T) {
		alice, _ := GenerateECDHP256Key()
		bob, _ := GenerateECDHP384Key()

		match, err := ECDHSharedSecretTest(alice, bob)
		if err == nil {
			t.Error("Should fail with different curves")
		}
		if match {
			t.Error("Match should be false")
		}
	})
}

// TestECDHDeterminism tests that same key pairs produce same shared secrets
func TestECDHDeterminism(t *testing.T) {
	alice, _ := GenerateECDHP256Key()
	bob, _ := GenerateECDHP256Key()

	// Compute shared secret multiple times
	shared1, _ := ECDHComputeShared(alice.PrivateKey, bob.PublicKey)
	shared2, _ := ECDHComputeShared(alice.PrivateKey, bob.PublicKey)
	shared3, _ := ECDHComputeShared(alice.PrivateKey, bob.PublicKey)

	if !bytes.Equal(shared1, shared2) || !bytes.Equal(shared2, shared3) {
		t.Error("ECDH should be deterministic - same inputs should produce same outputs")
	}
}

// TestECDHWithDifferentKeyPairs tests that different key pairs produce different shared secrets
func TestECDHWithDifferentKeyPairs(t *testing.T) {
	alice, _ := GenerateECDHP256Key()
	bob1, _ := GenerateECDHP256Key()
	bob2, _ := GenerateECDHP256Key()

	shared1, _ := ECDHComputeShared(alice.PrivateKey, bob1.PublicKey)
	shared2, _ := ECDHComputeShared(alice.PrivateKey, bob2.PublicKey)

	if bytes.Equal(shared1, shared2) {
		t.Error("Different key pairs should produce different shared secrets")
	}
}

// TestECDHErrorPathWithMockReader tests error path when rand.Reader fails
func TestECDHErrorPathWithMockReader(t *testing.T) {
	t.Run("GenerateECDHKey fails when rand.Reader fails", func(t *testing.T) {
		originalReader := ecdhRandReader
		ecdhRandReader = &ecdhFailingReader{}
		defer func() { ecdhRandReader = originalReader }()

		_, err := GenerateECDHKey(elliptic.P256())
		if err == nil {
			t.Error("GenerateECDHKey should fail when rand.Reader fails")
		}
	})
}

// ecdhFailingReader is a mock io.Reader that always returns an error
type ecdhFailingReader struct{}

func (r *ecdhFailingReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("mock random error")
}

// BenchmarkECDH benchmarks ECDH operations
func BenchmarkGenerateECDHP256Key(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateECDHP256Key()
	}
}

func BenchmarkGenerateECDHP384Key(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateECDHP384Key()
	}
}

func BenchmarkGenerateECDHP521Key(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateECDHP521Key()
	}
}

func BenchmarkECDHComputeSharedP256(b *testing.B) {
	alice, _ := GenerateECDHP256Key()
	bob, _ := GenerateECDHP256Key()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ECDHComputeShared(alice.PrivateKey, bob.PublicKey)
	}
}

func BenchmarkECDHComputeSharedP384(b *testing.B) {
	alice, _ := GenerateECDHP384Key()
	bob, _ := GenerateECDHP384Key()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ECDHComputeShared(alice.PrivateKey, bob.PublicKey)
	}
}

func BenchmarkECDHComputeSharedP521(b *testing.B) {
	alice, _ := GenerateECDHP521Key()
	bob, _ := GenerateECDHP521Key()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ECDHComputeShared(alice.PrivateKey, bob.PublicKey)
	}
}

func BenchmarkECDHComputeSharedSHA256(b *testing.B) {
	alice, _ := GenerateECDHP256Key()
	bob, _ := GenerateECDHP256Key()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ECDHComputeSharedSHA256(alice.PrivateKey, bob.PublicKey, 32)
	}
}

func BenchmarkECDHKeyExchange(b *testing.B) {
	alice, _ := GenerateECDHP256Key()
	bob, _ := GenerateECDHP256Key()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ECDHKeyExchange(alice.PrivateKey, bob.PublicKey, 32)
	}
}
