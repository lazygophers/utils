package cryptox

import (
	"bytes"
	"crypto/elliptic"
	"crypto/sha256"
	"math/big"
	"testing"
)

// TestGenerateECDHKey tests ECDH key generation with different curves
func TestGenerateECDHKey(t *testing.T) {
	curves := []elliptic.Curve{
		elliptic.P224(),
		elliptic.P256(),
		elliptic.P384(),
		elliptic.P521(),
	}

	for _, curve := range curves {
		keyPair, err := GenerateECDHKey(curve)
		if err != nil {
			t.Errorf("GenerateECDHKey failed for curve %s: %v", GetCurveName(curve), err)
			continue
		}

		if keyPair.PrivateKey == nil {
			t.Errorf("Private key is nil for curve %s", GetCurveName(curve))
		}

		if keyPair.PublicKey == nil {
			t.Errorf("Public key is nil for curve %s", GetCurveName(curve))
		}

		if keyPair.PrivateKey.Curve != curve {
			t.Errorf("Private key curve mismatch for curve %s", GetCurveName(curve))
		}

		if keyPair.PublicKey.Curve != curve {
			t.Errorf("Public key curve mismatch for curve %s", GetCurveName(curve))
		}
	}
}

// TestGenerateECDHP256Key tests P-256 ECDH key generation
func TestGenerateECDHP256Key(t *testing.T) {
	keyPair, err := GenerateECDHP256Key()
	if err != nil {
		t.Fatalf("GenerateECDHP256Key failed: %v", err)
	}

	if keyPair.PrivateKey.Curve != elliptic.P256() {
		t.Error("Expected P-256 curve")
	}
}

// TestGenerateECDHP384Key tests P-384 ECDH key generation
func TestGenerateECDHP384Key(t *testing.T) {
	keyPair, err := GenerateECDHP384Key()
	if err != nil {
		t.Fatalf("GenerateECDHP384Key failed: %v", err)
	}

	if keyPair.PrivateKey.Curve != elliptic.P384() {
		t.Error("Expected P-384 curve")
	}
}

// TestGenerateECDHP521Key tests P-521 ECDH key generation
func TestGenerateECDHP521Key(t *testing.T) {
	keyPair, err := GenerateECDHP521Key()
	if err != nil {
		t.Fatalf("GenerateECDHP521Key failed: %v", err)
	}

	if keyPair.PrivateKey.Curve != elliptic.P521() {
		t.Error("Expected P-521 curve")
	}
}

// TestECDHComputeShared tests basic ECDH shared secret computation
func TestECDHComputeShared(t *testing.T) {
	// Generate two key pairs
	alice, err := GenerateECDHP256Key()
	if err != nil {
		t.Fatalf("Failed to generate Alice's key: %v", err)
	}

	bob, err := GenerateECDHP256Key()
	if err != nil {
		t.Fatalf("Failed to generate Bob's key: %v", err)
	}

	// Compute shared secrets
	aliceShared, err := ECDHComputeShared(alice.PrivateKey, bob.PublicKey)
	if err != nil {
		t.Fatalf("Alice's ECDH computation failed: %v", err)
	}

	bobShared, err := ECDHComputeShared(bob.PrivateKey, alice.PublicKey)
	if err != nil {
		t.Fatalf("Bob's ECDH computation failed: %v", err)
	}

	// Verify shared secrets are identical
	if !bytes.Equal(aliceShared, bobShared) {
		t.Error("ECDH shared secrets should be identical")
	}

	if len(aliceShared) == 0 {
		t.Error("Shared secret cannot be empty")
	}
}

// TestECDHComputeSharedWithKDF tests ECDH with KDF
func TestECDHComputeSharedWithKDF(t *testing.T) {
	alice, err := GenerateECDHP256Key()
	if err != nil {
		t.Fatalf("Failed to generate Alice's key: %v", err)
	}

	bob, err := GenerateECDHP256Key()
	if err != nil {
		t.Fatalf("Failed to generate Bob's key: %v", err)
	}

	keyLength := 32

	// Compute derived keys
	aliceKey, err := ECDHComputeSharedWithKDF(alice.PrivateKey, bob.PublicKey, keyLength, sha256.New)
	if err != nil {
		t.Fatalf("Alice's ECDH KDF computation failed: %v", err)
	}

	bobKey, err := ECDHComputeSharedWithKDF(bob.PrivateKey, alice.PublicKey, keyLength, sha256.New)
	if err != nil {
		t.Fatalf("Bob's ECDH KDF computation failed: %v", err)
	}

	// Verify derived keys are identical
	if !bytes.Equal(aliceKey, bobKey) {
		t.Error("ECDH derived keys should be identical")
	}

	if len(aliceKey) != keyLength {
		t.Errorf("Expected key length %d, got %d", keyLength, len(aliceKey))
	}

	// Test with different key lengths
	testLengths := []int{16, 32, 48, 64, 128}
	for _, length := range testLengths {
		key1, err := ECDHComputeSharedWithKDF(alice.PrivateKey, bob.PublicKey, length, sha256.New)
		if err != nil {
			t.Errorf("ECDH KDF failed for length %d: %v", length, err)
			continue
		}

		if len(key1) != length {
			t.Errorf("Expected key length %d, got %d", length, len(key1))
		}
	}
}

// TestECDHComputeSharedSHA256 tests ECDH with SHA256 KDF
func TestECDHComputeSharedSHA256(t *testing.T) {
	alice, err := GenerateECDHP256Key()
	if err != nil {
		t.Fatalf("Failed to generate Alice's key: %v", err)
	}

	bob, err := GenerateECDHP256Key()
	if err != nil {
		t.Fatalf("Failed to generate Bob's key: %v", err)
	}

	keyLength := 32

	aliceKey, err := ECDHComputeSharedSHA256(alice.PrivateKey, bob.PublicKey, keyLength)
	if err != nil {
		t.Fatalf("Alice's ECDH SHA256 computation failed: %v", err)
	}

	bobKey, err := ECDHComputeSharedSHA256(bob.PrivateKey, alice.PublicKey, keyLength)
	if err != nil {
		t.Fatalf("Bob's ECDH SHA256 computation failed: %v", err)
	}

	if !bytes.Equal(aliceKey, bobKey) {
		t.Error("ECDH SHA256 derived keys should be identical")
	}

	if len(aliceKey) != keyLength {
		t.Errorf("Expected key length %d, got %d", keyLength, len(aliceKey))
	}
}

// TestECDHKeyExchange tests complete ECDH key exchange
func TestECDHKeyExchange(t *testing.T) {
	alice, err := GenerateECDHP256Key()
	if err != nil {
		t.Fatalf("Failed to generate Alice's key: %v", err)
	}

	bob, err := GenerateECDHP256Key()
	if err != nil {
		t.Fatalf("Failed to generate Bob's key: %v", err)
	}

	keyLength := 32

	// Alice computes shared key
	aliceShared, err := ECDHKeyExchange(alice.PrivateKey, bob.PublicKey, keyLength)
	if err != nil {
		t.Fatalf("Alice's key exchange failed: %v", err)
	}

	// Bob computes shared key
	bobShared, err := ECDHKeyExchange(bob.PrivateKey, alice.PublicKey, keyLength)
	if err != nil {
		t.Fatalf("Bob's key exchange failed: %v", err)
	}

	if !bytes.Equal(aliceShared, bobShared) {
		t.Error("Key exchange should produce identical keys")
	}

	if len(aliceShared) != keyLength {
		t.Errorf("Expected key length %d, got %d", keyLength, len(aliceShared))
	}
}

// TestValidateECDHKeyPair tests ECDH key pair validation
func TestValidateECDHKeyPair(t *testing.T) {
	keyPair, err := GenerateECDHP256Key()
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}

	// Test valid key pair
	err = ValidateECDHKeyPair(keyPair)
	if err != nil {
		t.Errorf("Valid key pair should pass validation: %v", err)
	}

	// Test with nil key pair
	err = ValidateECDHKeyPair(nil)
	if err == nil {
		t.Error("Expected error for nil key pair")
	}

	// Test with nil private key
	invalidKeyPair := &ECDHKeyPair{
		PrivateKey: nil,
		PublicKey:  keyPair.PublicKey,
	}
	err = ValidateECDHKeyPair(invalidKeyPair)
	if err == nil {
		t.Error("Expected error for nil private key")
	}

	// Test with nil public key
	invalidKeyPair = &ECDHKeyPair{
		PrivateKey: keyPair.PrivateKey,
		PublicKey:  nil,
	}
	err = ValidateECDHKeyPair(invalidKeyPair)
	if err == nil {
		t.Error("Expected error for nil public key")
	}

	// Test with mismatched public key
	otherKeyPair, _ := GenerateECDHP256Key()
	invalidKeyPair = &ECDHKeyPair{
		PrivateKey: keyPair.PrivateKey,
		PublicKey:  otherKeyPair.PublicKey,
	}
	err = ValidateECDHKeyPair(invalidKeyPair)
	if err == nil {
		t.Error("Expected error for mismatched keys")
	}
}

// TestECDHPublicKeyCoordinates tests public key coordinate conversion
func TestECDHPublicKeyCoordinates(t *testing.T) {
	keyPair, err := GenerateECDHP256Key()
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}

	// Test coordinate extraction
	x, y, err := ECDHPublicKeyToCoordinates(keyPair.PublicKey)
	if err != nil {
		t.Fatalf("ECDHPublicKeyToCoordinates failed: %v", err)
	}

	if x == nil || y == nil {
		t.Fatal("Coordinates cannot be nil")
	}

	// Test coordinate reconstruction
	reconstructedKey, err := ECDHPublicKeyFromCoordinates(elliptic.P256(), x, y)
	if err != nil {
		t.Fatalf("ECDHPublicKeyFromCoordinates failed: %v", err)
	}

	// Verify reconstructed key matches original
	if reconstructedKey.X.Cmp(keyPair.PublicKey.X) != 0 {
		t.Error("Reconstructed X coordinate mismatch")
	}

	if reconstructedKey.Y.Cmp(keyPair.PublicKey.Y) != 0 {
		t.Error("Reconstructed Y coordinate mismatch")
	}

	if reconstructedKey.Curve != keyPair.PublicKey.Curve {
		t.Error("Reconstructed curve mismatch")
	}
}

// TestECDHSharedSecretTest tests the shared secret test function
func TestECDHSharedSecretTest(t *testing.T) {
	keyPair1, err := GenerateECDHP256Key()
	if err != nil {
		t.Fatalf("Failed to generate key pair 1: %v", err)
	}

	keyPair2, err := GenerateECDHP256Key()
	if err != nil {
		t.Fatalf("Failed to generate key pair 2: %v", err)
	}

	// Test with valid key pairs
	match, err := ECDHSharedSecretTest(keyPair1, keyPair2)
	if err != nil {
		t.Fatalf("ECDHSharedSecretTest failed: %v", err)
	}

	if !match {
		t.Error("ECDH shared secrets should match")
	}

	// Test with self (should always work)
	match, err = ECDHSharedSecretTest(keyPair1, keyPair1)
	if err != nil {
		t.Fatalf("Self ECDH test failed: %v", err)
	}

	if !match {
		t.Error("Self ECDH should always work")
	}
}

// TestECDHErrorConditions tests error conditions for ECDH functions
func TestECDHErrorConditions(t *testing.T) {
	// Test GenerateECDHKey with nil curve
	_, err := GenerateECDHKey(nil)
	if err == nil {
		t.Error("Expected error for nil curve")
	}

	// Test ECDHComputeShared with nil keys
	keyPair, _ := GenerateECDHP256Key()

	_, err = ECDHComputeShared(nil, keyPair.PublicKey)
	if err == nil {
		t.Error("Expected error for nil private key")
	}

	_, err = ECDHComputeShared(keyPair.PrivateKey, nil)
	if err == nil {
		t.Error("Expected error for nil public key")
	}

	// Test with point not on curve (create invalid public key)
	invalidPublicKey := *keyPair.PublicKey
	invalidPublicKey.X = big.NewInt(1)
	invalidPublicKey.Y = big.NewInt(1)

	_, err = ECDHComputeShared(keyPair.PrivateKey, &invalidPublicKey)
	if err == nil {
		t.Error("Expected error for point not on curve")
	}

	// Test with curve mismatch
	keyPairP384, _ := GenerateECDHP384Key()
	_, err = ECDHComputeShared(keyPair.PrivateKey, keyPairP384.PublicKey)
	if err == nil {
		t.Error("Expected error for curve mismatch")
	}

	// Test ECDHComputeSharedWithKDF with invalid parameters
	_, err = ECDHComputeSharedWithKDF(keyPair.PrivateKey, keyPair.PublicKey, 0, sha256.New)
	if err == nil {
		t.Error("Expected error for zero key length")
	}

	_, err = ECDHComputeSharedWithKDF(keyPair.PrivateKey, keyPair.PublicKey, 32, nil)
	if err == nil {
		t.Error("Expected error for nil KDF function")
	}

	// Test ECDHPublicKeyFromCoordinates with nil parameters
	_, err = ECDHPublicKeyFromCoordinates(nil, big.NewInt(1), big.NewInt(1))
	if err == nil {
		t.Error("Expected error for nil curve")
	}

	_, err = ECDHPublicKeyFromCoordinates(elliptic.P256(), nil, big.NewInt(1))
	if err == nil {
		t.Error("Expected error for nil x coordinate")
	}

	_, err = ECDHPublicKeyFromCoordinates(elliptic.P256(), big.NewInt(1), nil)
	if err == nil {
		t.Error("Expected error for nil y coordinate")
	}

	_, err = ECDHPublicKeyFromCoordinates(elliptic.P256(), big.NewInt(1), big.NewInt(1))
	if err == nil {
		t.Error("Expected error for invalid point")
	}

	// Test ECDHPublicKeyToCoordinates with nil key
	_, _, err = ECDHPublicKeyToCoordinates(nil)
	if err == nil {
		t.Error("Expected error for nil public key")
	}

	// Test ECDHSharedSecretTest with nil key pairs
	_, err = ECDHSharedSecretTest(nil, keyPair)
	if err == nil {
		t.Error("Expected error for nil key pair 1")
	}

	_, err = ECDHSharedSecretTest(keyPair, nil)
	if err == nil {
		t.Error("Expected error for nil key pair 2")
	}
}

// TestECDHDifferentCurves tests ECDH with different elliptic curves
func TestECDHDifferentCurves(t *testing.T) {
	curves := []elliptic.Curve{
		elliptic.P224(),
		elliptic.P256(),
		elliptic.P384(),
		elliptic.P521(),
	}

	for _, curve := range curves {
		// Generate two key pairs on the same curve
		keyPair1, err := GenerateECDHKey(curve)
		if err != nil {
			t.Errorf("Failed to generate key pair 1 for %s: %v", GetCurveName(curve), err)
			continue
		}

		keyPair2, err := GenerateECDHKey(curve)
		if err != nil {
			t.Errorf("Failed to generate key pair 2 for %s: %v", GetCurveName(curve), err)
			continue
		}

		// Test shared secret computation
		secret1, err := ECDHComputeShared(keyPair1.PrivateKey, keyPair2.PublicKey)
		if err != nil {
			t.Errorf("ECDH computation failed for %s: %v", GetCurveName(curve), err)
			continue
		}

		secret2, err := ECDHComputeShared(keyPair2.PrivateKey, keyPair1.PublicKey)
		if err != nil {
			t.Errorf("ECDH computation failed for %s: %v", GetCurveName(curve), err)
			continue
		}

		if !bytes.Equal(secret1, secret2) {
			t.Errorf("ECDH shared secrets should be identical for %s", GetCurveName(curve))
		}

		if len(secret1) == 0 {
			t.Errorf("Shared secret should not be empty for %s", GetCurveName(curve))
		}
	}
}

// TestECDHKeyDerivationWithDifferentLengths tests KDF with various key lengths
func TestECDHKeyDerivationWithDifferentLengths(t *testing.T) {
	alice, err := GenerateECDHP256Key()
	if err != nil {
		t.Fatalf("Failed to generate Alice's key: %v", err)
	}

	bob, err := GenerateECDHP256Key()
	if err != nil {
		t.Fatalf("Failed to generate Bob's key: %v", err)
	}

	// Test various key lengths
	testLengths := []int{1, 8, 16, 32, 48, 64, 100, 128, 256}

	for _, length := range testLengths {
		key1, err := ECDHComputeSharedWithKDF(alice.PrivateKey, bob.PublicKey, length, sha256.New)
		if err != nil {
			t.Errorf("KDF failed for length %d: %v", length, err)
			continue
		}

		key2, err := ECDHComputeSharedWithKDF(bob.PrivateKey, alice.PublicKey, length, sha256.New)
		if err != nil {
			t.Errorf("KDF failed for length %d: %v", length, err)
			continue
		}

		if len(key1) != length {
			t.Errorf("Expected key length %d, got %d", length, len(key1))
		}

		if !bytes.Equal(key1, key2) {
			t.Errorf("Derived keys should be identical for length %d", length)
		}
	}
}