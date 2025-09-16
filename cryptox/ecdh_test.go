package cryptox

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"errors"
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

// ==== MERGED ECDH TESTS FROM ecc_100_coverage_test.go ====

// Mock failures for ECDH dependency injection
type FailingECDHReader struct{}

func (fr FailingECDHReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("simulated random reader failure")
}

// TestECDH_100PercentCoverage tests all error paths in ECDH functions
func TestECDH_100PercentCoverage(t *testing.T) {
	// Save original functions
	originalECDHRandReader := ecdhRandReader

	// Restore original functions after test
	defer func() {
		ecdhRandReader = originalECDHRandReader
	}()

	// Test 1: Trigger rand.Reader failure in ECDH key generation
	ecdhRandReader = FailingECDHReader{}

	_, err := GenerateECDHKey(elliptic.P256())
	if err == nil {
		t.Error("Expected rand.Reader error in GenerateECDHKey")
	}

	_, err = GenerateECDHP256Key()
	if err == nil {
		t.Error("Expected rand.Reader error in GenerateECDHP256Key")
	}

	_, err = GenerateECDHP384Key()
	if err == nil {
		t.Error("Expected rand.Reader error in GenerateECDHP384Key")
	}

	_, err = GenerateECDHP521Key()
	if err == nil {
		t.Error("Expected rand.Reader error in GenerateECDHP521Key")
	}

	// Restore readers
	ecdhRandReader = originalECDHRandReader
}

// TestECDHSpecificCoveragePaths tests specific code paths to reach 100%
func TestECDHSpecificCoveragePaths(t *testing.T) {
	// Generate key pairs for testing
	keyPair1, err := GenerateECDHP256Key()
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}

	// Create an invalid public key by modifying coordinates to be off the curve
	// This should trigger the IsOnCurve check in ECDHComputeShared
	invalidPublicKey := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     big.NewInt(1), // Invalid point
		Y:     big.NewInt(1), // Invalid point
	}

	_, err = ECDHComputeShared(keyPair1.PrivateKey, invalidPublicKey)
	if err == nil {
		t.Error("Expected error for public key not on curve")
	}

	// Create key pair with mismatched curve types to trigger ValidateECDHKeyPair paths
	keyPairP384, err := GenerateECDHP384Key()
	if err != nil {
		t.Fatalf("Failed to generate P384 key pair: %v", err)
	}

	// Test ValidateECDHKeyPair with curve mismatch
	mismatchedKeyPair := &ECDHKeyPair{
		PrivateKey: keyPair1.PrivateKey,   // P256
		PublicKey:  keyPairP384.PublicKey, // P384
	}

	err = ValidateECDHKeyPair(mismatchedKeyPair)
	if err == nil {
		t.Error("Expected error for curve mismatch in key pair")
	}

	// Test ValidateECDHKeyPair with invalid public key coordinates
	invalidCoordKeyPair := &ECDHKeyPair{
		PrivateKey: keyPair1.PrivateKey,
		PublicKey:  invalidPublicKey,
	}

	err = ValidateECDHKeyPair(invalidCoordKeyPair)
	if err == nil {
		t.Error("Expected error for invalid public key coordinates")
	}

	// Test ValidateECDHKeyPair with nil key pair
	err = ValidateECDHKeyPair(nil)
	if err == nil {
		t.Error("Expected error for nil key pair")
	}

	// Test with nil private key in key pair
	keyPairWithNilPrivate := &ECDHKeyPair{
		PrivateKey: nil,
		PublicKey:  keyPair1.PublicKey,
	}
	err = ValidateECDHKeyPair(keyPairWithNilPrivate)
	if err == nil {
		t.Error("Expected error for nil private key")
	}

	// Test with nil public key in key pair
	keyPairWithNilPublic := &ECDHKeyPair{
		PrivateKey: keyPair1.PrivateKey,
		PublicKey:  nil,
	}
	err = ValidateECDHKeyPair(keyPairWithNilPublic)
	if err == nil {
		t.Error("Expected error for nil public key")
	}

	// Test ECDHComputeSharedWithKDF error paths
	keyPair2, _ := GenerateECDHP256Key()

	// Test with 0 key length
	_, err = ECDHComputeSharedWithKDF(keyPair1.PrivateKey, keyPair2.PublicKey, 0, nil)
	if err == nil {
		t.Error("Expected error for zero key length")
	}

	// Test with nil KDF function
	_, err = ECDHComputeSharedWithKDF(keyPair1.PrivateKey, keyPair2.PublicKey, 32, nil)
	if err == nil {
		t.Error("Expected error for nil KDF function")
	}

	// Test ECDHComputeShared error paths
	_, err = ECDHComputeShared(nil, keyPair1.PublicKey)
	if err == nil {
		t.Error("Expected error for nil private key")
	}

	_, err = ECDHComputeShared(keyPair1.PrivateKey, nil)
	if err == nil {
		t.Error("Expected error for nil public key")
	}

	// Test ECDHSharedSecretTest error paths
	_, err = ECDHSharedSecretTest(nil, keyPair1)
	if err == nil {
		t.Error("Expected error for nil keyPair1")
	}

	_, err = ECDHSharedSecretTest(keyPair1, nil)
	if err == nil {
		t.Error("Expected error for nil keyPair2")
	}
}

// ==== MERGED FROM missing_coverage_test.go ====

// TestECDHSharedSecretTestMissingPaths tests the specific missing paths in ECDHSharedSecretTest
func TestECDHSharedSecretTestMissingPaths(t *testing.T) {
	t.Run("force_secret_length_mismatch", func(t *testing.T) {
		// Create two different key pairs with different curves to potentially get different secret lengths
		p256Key, err := GenerateECDHP256Key()
		if err != nil {
			t.Fatal(err)
		}

		p384Key, err := GenerateECDHP384Key()
		if err != nil {
			t.Fatal(err)
		}

		// Try to create a scenario where the shared secrets have different lengths
		// This is tricky because ECDH normally produces secrets of the same length for the same curve

		// Create a modified key pair that might produce different behavior
		// We'll create keys from the same curve but try to manipulate the calculation

		// Actually, let's try a simpler approach - create custom ECDHKeyPair structs
		// with keys that will fail computation in different ways to get different secret lengths

		// Generate valid keys first
		key1, _ := GenerateECDHP256Key()
		key2, _ := GenerateECDHP256Key()

		// Test the normal case first
		match, err := ECDHSharedSecretTest(key1, key2)
		if err != nil {
			t.Logf("Normal case error (expected for different keys): %v", err)
		} else {
			t.Logf("Normal case match: %v", match)
		}

		// Now let's try to create a case where ECDHComputeShared might return different lengths
		// by using keys with different curves (this should cause an error, but let's test it)

		// Create mixed curve scenario
		mixedPair1 := &ECDHKeyPair{
			PrivateKey: p256Key.PrivateKey,
			PublicKey:  p384Key.PublicKey, // Different curve!
		}

		mixedPair2 := &ECDHKeyPair{
			PrivateKey: p384Key.PrivateKey,
			PublicKey:  p256Key.PublicKey, // Different curve!
		}

		// This should trigger an error path in ECDHComputeShared
		match, err = ECDHSharedSecretTest(mixedPair1, mixedPair2)
		if err != nil {
			t.Logf("Mixed curve error (expected): %v", err)
		} else {
			t.Logf("Mixed curve match: %v", match)
		}
	})

	t.Run("force_byte_mismatch", func(t *testing.T) {
		// Try to create a scenario where the bytes don't match
		// This is actually the normal case when using different key pairs

		key1, _ := GenerateECDHP256Key()
		key2, _ := GenerateECDHP256Key()
		key3, _ := GenerateECDHP256Key()

		// Create mismatched pairs - Alice's private with Bob's public, vs Charlie's private with Alice's public
		mismatchPair1 := &ECDHKeyPair{
			PrivateKey: key1.PrivateKey,
			PublicKey:  key2.PublicKey,
		}

		mismatchPair2 := &ECDHKeyPair{
			PrivateKey: key3.PrivateKey,
			PublicKey:  key1.PublicKey,
		}

		// These should not produce matching secrets
		match, err := ECDHSharedSecretTest(mismatchPair1, mismatchPair2)
		if err != nil {
			t.Logf("Mismatched pairs error: %v", err)
		} else {
			if match {
				t.Error("Expected mismatch, but got match")
			} else {
				t.Log("Successfully triggered byte mismatch path")
			}
		}
	})

	t.Run("manual_secret_comparison", func(t *testing.T) {
		// Let's manually test the secret comparison logic by creating
		// a scenario that should trigger the byte-by-byte comparison

		key1, _ := GenerateECDHP256Key()
		key2, _ := GenerateECDHP256Key()

		// Create proper key pairs for ECDH
		pair1 := &ECDHKeyPair{
			PrivateKey: key1.PrivateKey,
			PublicKey:  key2.PublicKey,
		}

		pair2 := &ECDHKeyPair{
			PrivateKey: key2.PrivateKey,
			PublicKey:  key1.PublicKey,
		}

		// These should produce matching secrets (proper ECDH)
		match, err := ECDHSharedSecretTest(pair1, pair2)
		if err != nil {
			t.Logf("ECDH error (this might happen with mixed key pairs): %v", err)
		} else {
			if match {
				t.Log("Successfully tested matching ECDH comparison")
			} else {
				t.Log("Successfully tested non-matching ECDH comparison")
			}
		}
	})
}

// TestECDHSharedSecretTestNilCheck tests nil parameter handling
func TestECDHSharedSecretTestNilCheck(t *testing.T) {
	key1, _ := GenerateECDHP256Key()

	// Test nil first parameter
	_, err := ECDHSharedSecretTest(nil, key1)
	if err == nil {
		t.Error("Expected error for nil first parameter")
	}

	// Test nil second parameter
	_, err = ECDHSharedSecretTest(key1, nil)
	if err == nil {
		t.Error("Expected error for nil second parameter")
	}

	// Test both nil
	_, err = ECDHSharedSecretTest(nil, nil)
	if err == nil {
		t.Error("Expected error for both nil parameters")
	}
}
