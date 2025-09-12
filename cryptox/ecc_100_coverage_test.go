package cryptox

import (
	"crypto/elliptic"
	"crypto/sha256"
	"errors"
	"math/big"
	"testing"
)

// Mock failures for ECC dependency injection
type FailingECCReader struct{}

func (fr FailingECCReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("simulated random reader failure")
}

// TestECC_100PercentCoverage tests all error paths in ECC functions
func TestECC_100PercentCoverage(t *testing.T) {
	// Save original functions
	originalECDSARandReader := ecdsaRandReader
	originalECDHRandReader := ecdhRandReader

	// Restore original functions after test
	defer func() {
		ecdsaRandReader = originalECDSARandReader
		ecdhRandReader = originalECDHRandReader
	}()

	// Test 1: Trigger rand.Reader failure in ECDSA key generation
	ecdsaRandReader = FailingECCReader{}
	ecdhRandReader = originalECDHRandReader

	_, err := GenerateECDSAKey(elliptic.P256())
	if err == nil {
		t.Error("Expected rand.Reader error in GenerateECDSAKey")
	}

	_, err = GenerateECDSAP256Key()
	if err == nil {
		t.Error("Expected rand.Reader error in GenerateECDSAP256Key")
	}

	_, err = GenerateECDSAP384Key()
	if err == nil {
		t.Error("Expected rand.Reader error in GenerateECDSAP384Key")
	}

	_, err = GenerateECDSAP521Key()
	if err == nil {
		t.Error("Expected rand.Reader error in GenerateECDSAP521Key")
	}

	// Test 2: Trigger rand.Reader failure in ECDH key generation
	ecdsaRandReader = originalECDSARandReader
	ecdhRandReader = FailingECCReader{}

	_, err = GenerateECDHKey(elliptic.P256())
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

	// Test 3: Trigger rand.Reader failure in ECDSA signing
	ecdsaRandReader = FailingECCReader{}
	ecdhRandReader = originalECDHRandReader

	// Generate a valid key first with original reader
	ecdsaRandReader = originalECDSARandReader
	keyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("Failed to generate test key: %v", err)
	}

	// Now set failing reader and test signing
	ecdsaRandReader = FailingECCReader{}

	_, _, err = ECDSASign(keyPair.PrivateKey, []byte("test"), sha256.New)
	if err == nil {
		t.Error("Expected rand.Reader error in ECDSASign")
	}

	_, _, err = ECDSASignSHA256(keyPair.PrivateKey, []byte("test"))
	if err == nil {
		t.Error("Expected rand.Reader error in ECDSASignSHA256")
	}

	_, _, err = ECDSASignSHA512(keyPair.PrivateKey, []byte("test"))
	if err == nil {
		t.Error("Expected rand.Reader error in ECDSASignSHA512")
	}

	// Restore readers
	ecdsaRandReader = originalECDSARandReader
	ecdhRandReader = originalECDHRandReader
}

// TestECDSAPEMErrorPaths tests PEM parsing error paths
func TestECDSAPEMErrorPaths(t *testing.T) {
	// Test invalid PEM types
	invalidPrivateKeyPEM := `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA1234567890
-----END RSA PRIVATE KEY-----`

	_, err := ECDSAPrivateKeyFromPEM([]byte(invalidPrivateKeyPEM))
	if err == nil {
		t.Error("Expected error for wrong PEM type")
	}

	invalidPublicKeyPEM := `-----BEGIN RSA PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1234567890
-----END RSA PUBLIC KEY-----`

	_, err = ECDSAPublicKeyFromPEM([]byte(invalidPublicKeyPEM))
	if err == nil {
		t.Error("Expected error for wrong PEM type")
	}

	// Test valid PEM type but invalid key data
	invalidPrivateKeyData := `-----BEGIN EC PRIVATE KEY-----
aW52YWxpZCBkYXRh
-----END EC PRIVATE KEY-----`

	_, err = ECDSAPrivateKeyFromPEM([]byte(invalidPrivateKeyData))
	if err == nil {
		t.Error("Expected error for invalid private key data")
	}

	invalidPublicKeyData := `-----BEGIN PUBLIC KEY-----
aW52YWxpZCBkYXRh
-----END PUBLIC KEY-----`

	_, err = ECDSAPublicKeyFromPEM([]byte(invalidPublicKeyData))
	if err == nil {
		t.Error("Expected error for invalid public key data")
	}

	// Test public key with wrong key type - create a malformed PEM with correct header but invalid data
	malformedPublicKey := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7u25K9xP8lmE9v3U8OVt
w0N6HVc5tNOIgbzLFfvO8mWWnhxF11gK2JvKm9lGqLpbOy1rGtCNHsFGj5kD7MQX
-----END PUBLIC KEY-----`

	_, err = ECDSAPublicKeyFromPEM([]byte(malformedPublicKey))
	if err == nil {
		t.Error("Expected error for non-ECDSA public key")
	}
}

// TestECDSASignatureBytesErrorPaths tests signature byte encoding/decoding error paths
func TestECDSASignatureBytesErrorPaths(t *testing.T) {
	// Test decoding with various invalid DER data
	invalidSignatures := [][]byte{
		{}, // Empty
		{0x01}, // Too short
		{0x31, 0x04, 0x02, 0x01, 0x01}, // Wrong sequence tag
		{0x30, 0xFF, 0x02, 0x01, 0x01}, // Invalid sequence length
		{0x30, 0x04, 0x01, 0x01, 0x01}, // Missing INTEGER tag for r
		{0x30, 0x04, 0x02, 0xFF, 0x01}, // Invalid r length
		{0x30, 0x06, 0x02, 0x01, 0x01, 0x01}, // Missing INTEGER tag for s
		{0x30, 0x06, 0x02, 0x01, 0x01, 0x02, 0xFF}, // Invalid s length
		{0x30, 0x04, 0x02, 0x01}, // Incomplete data
	}

	for i, invalidSig := range invalidSignatures {
		_, _, err := ECDSASignatureFromBytes(invalidSig)
		if err == nil {
			t.Errorf("Expected error for invalid signature %d", i)
		}
	}

	// Test encoding with edge case values
	// Test with large numbers that might cause issues
	largeR := new(big.Int)
	largeR.SetString("123456789012345678901234567890123456789012345678901234567890", 10)
	largeS := new(big.Int)
	largeS.SetString("987654321098765432109876543210987654321098765432109876543210", 10)

	sigBytes, err := ECDSASignatureToBytes(largeR, largeS)
	if err != nil {
		t.Errorf("Should handle large signature components: %v", err)
	}

	decodedR, decodedS, err := ECDSASignatureFromBytes(sigBytes)
	if err != nil {
		t.Errorf("Should decode large signature components: %v", err)
	}

	if decodedR.Cmp(largeR) != 0 || decodedS.Cmp(largeS) != 0 {
		t.Error("Large signature components should round-trip correctly")
	}

	// Test with numbers that have high bit set (need 0x00 prefix)
	highBitR := new(big.Int).SetBytes([]byte{0xFF, 0xFF, 0xFF, 0xFF})
	highBitS := new(big.Int).SetBytes([]byte{0x80, 0x00, 0x00, 0x01})

	sigBytes2, err := ECDSASignatureToBytes(highBitR, highBitS)
	if err != nil {
		t.Errorf("Should handle high-bit signature components: %v", err)
	}

	decodedR2, decodedS2, err := ECDSASignatureFromBytes(sigBytes2)
	if err != nil {
		t.Errorf("Should decode high-bit signature components: %v", err)
	}

	if decodedR2.Cmp(highBitR) != 0 || decodedS2.Cmp(highBitS) != 0 {
		t.Error("High-bit signature components should round-trip correctly")
	}
}

// TestECDSAEdgeCases tests various edge cases
func TestECDSAEdgeCases(t *testing.T) {
	// Test with minimal data
	keyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("Failed to generate key: %v", err)
	}

	// Test with empty data
	r, s, err := ECDSASignSHA256(keyPair.PrivateKey, []byte{})
	if err != nil {
		t.Errorf("Should handle empty data: %v", err)
	}

	if !ECDSAVerifySHA256(keyPair.PublicKey, []byte{}, r, s) {
		t.Error("Should verify empty data signature")
	}

	// Test with single byte
	r2, s2, err := ECDSASignSHA256(keyPair.PrivateKey, []byte{0x42})
	if err != nil {
		t.Errorf("Should handle single byte: %v", err)
	}

	if !ECDSAVerifySHA256(keyPair.PublicKey, []byte{0x42}, r2, s2) {
		t.Error("Should verify single byte signature")
	}

	// Test signature round-trip with zero components
	zeroSig, err := ECDSASignatureToBytes(big.NewInt(0), big.NewInt(1))
	if err != nil {
		t.Errorf("Should handle zero R component: %v", err)
	}

	zeroR, zeroS, err := ECDSASignatureFromBytes(zeroSig)
	if err != nil {
		t.Errorf("Should decode zero R component: %v", err)
	}

	if zeroR.Cmp(big.NewInt(0)) != 0 || zeroS.Cmp(big.NewInt(1)) != 0 {
		t.Error("Zero component should round-trip correctly")
	}
}

// TestECDHEdgeCases tests ECDH edge cases
func TestECDHEdgeCases(t *testing.T) {
	keyPair, err := GenerateECDHP256Key()
	if err != nil {
		t.Fatalf("Failed to generate key: %v", err)
	}

	// Test self key exchange (Alice = Bob)
	sharedSecret, err := ECDHComputeShared(keyPair.PrivateKey, keyPair.PublicKey)
	if err != nil {
		t.Errorf("Self key exchange should work: %v", err)
	}

	if len(sharedSecret) == 0 {
		t.Error("Self shared secret should not be empty")
	}

	// Test KDF with very small key length
	derivedKey, err := ECDHComputeSharedWithKDF(keyPair.PrivateKey, keyPair.PublicKey, 1, sha256.New)
	if err != nil {
		t.Errorf("KDF should work with small key length: %v", err)
	}

	if len(derivedKey) != 1 {
		t.Errorf("Expected key length 1, got %d", len(derivedKey))
	}

	// Test with key length larger than hash output (will trigger counter mode)
	largeKey, err := ECDHComputeSharedWithKDF(keyPair.PrivateKey, keyPair.PublicKey, 100, sha256.New)
	if err != nil {
		t.Errorf("KDF should work with large key length: %v", err)
	}

	if len(largeKey) != 100 {
		t.Errorf("Expected key length 100, got %d", len(largeKey))
	}

	// Test coordinate conversion with curve edge points
	// Test with generator point
	gx, gy := elliptic.P256().Params().Gx, elliptic.P256().Params().Gy
	pubKey, err := ECDHPublicKeyFromCoordinates(elliptic.P256(), gx, gy)
	if err != nil {
		t.Errorf("Should handle generator point: %v", err)
	}

	x, y, err := ECDHPublicKeyToCoordinates(pubKey)
	if err != nil {
		t.Errorf("Should extract generator coordinates: %v", err)
	}

	if x.Cmp(gx) != 0 || y.Cmp(gy) != 0 {
		t.Error("Generator coordinates should round-trip correctly")
	}
}

// TestECCUtilityFunctions tests utility functions thoroughly
func TestECCUtilityFunctions(t *testing.T) {
	// Test GetCurveName with all supported curves
	testCases := []struct {
		curve elliptic.Curve
		name  string
	}{
		{elliptic.P224(), "P-224"},
		{elliptic.P256(), "P-256"},
		{elliptic.P384(), "P-384"},
		{elliptic.P521(), "P-521"},
	}

	for _, tc := range testCases {
		if GetCurveName(tc.curve) != tc.name {
			t.Errorf("Wrong curve name for %s", tc.name)
		}
	}

	// Test IsValidCurve with all cases
	for _, tc := range testCases {
		if !IsValidCurve(tc.curve) {
			t.Errorf("Curve %s should be valid", tc.name)
		}
	}

	// Test with nil curve
	if GetCurveName(nil) != "Unknown" {
		t.Error("nil curve should return 'Unknown'")
	}

	if IsValidCurve(nil) {
		t.Error("nil curve should be invalid")
	}
}

// TestECCRealWorldScenarios tests realistic usage scenarios
func TestECCRealWorldScenarios(t *testing.T) {
	// Test complete ECDSA workflow
	// 1. Generate key pair
	signer, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("Failed to generate signer key: %v", err)
	}

	// 2. Export keys to PEM
	privateKeyPEM, err := ECDSAPrivateKeyToPEM(signer.PrivateKey)
	if err != nil {
		t.Fatalf("Failed to export private key: %v", err)
	}

	publicKeyPEM, err := ECDSAPublicKeyToPEM(signer.PublicKey)
	if err != nil {
		t.Fatalf("Failed to export public key: %v", err)
	}

	// 3. Import keys from PEM (simulate key distribution)
	importedPrivateKey, err := ECDSAPrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		t.Fatalf("Failed to import private key: %v", err)
	}

	importedPublicKey, err := ECDSAPublicKeyFromPEM(publicKeyPEM)
	if err != nil {
		t.Fatalf("Failed to import public key: %v", err)
	}

	// 4. Sign document
	document := []byte("Important contract data")
	r, s, err := ECDSASignSHA256(importedPrivateKey, document)
	if err != nil {
		t.Fatalf("Failed to sign document: %v", err)
	}

	// 5. Export signature
	sigBytes, err := ECDSASignatureToBytes(r, s)
	if err != nil {
		t.Fatalf("Failed to export signature: %v", err)
	}

	// 6. Import signature (simulate transmission)
	importedR, importedS, err := ECDSASignatureFromBytes(sigBytes)
	if err != nil {
		t.Fatalf("Failed to import signature: %v", err)
	}

	// 7. Verify signature
	if !ECDSAVerifySHA256(importedPublicKey, document, importedR, importedS) {
		t.Error("Failed to verify document signature")
	}

	// 8. Test with tampered document
	tamperedDoc := []byte("Tampered contract data")
	if ECDSAVerifySHA256(importedPublicKey, tamperedDoc, importedR, importedS) {
		t.Error("Should not verify tampered document")
	}

	// Test complete ECDH workflow
	// 1. Generate key pairs for two parties
	alice, err := GenerateECDHP256Key()
	if err != nil {
		t.Fatalf("Failed to generate Alice's key: %v", err)
	}

	bob, err := GenerateECDHP256Key()
	if err != nil {
		t.Fatalf("Failed to generate Bob's key: %v", err)
	}

	// 2. Validate key pairs
	if err := ValidateECDHKeyPair(alice); err != nil {
		t.Errorf("Alice's key pair should be valid: %v", err)
	}

	if err := ValidateECDHKeyPair(bob); err != nil {
		t.Errorf("Bob's key pair should be valid: %v", err)
	}

	// 3. Exchange public keys and compute shared secrets
	aliceShared, err := ECDHKeyExchange(alice.PrivateKey, bob.PublicKey, 32)
	if err != nil {
		t.Fatalf("Alice's key exchange failed: %v", err)
	}

	bobShared, err := ECDHKeyExchange(bob.PrivateKey, alice.PublicKey, 32)
	if err != nil {
		t.Fatalf("Bob's key exchange failed: %v", err)
	}

	// 4. Verify shared secrets match
	match, err := ECDHSharedSecretTest(alice, bob)
	if err != nil {
		t.Errorf("ECDH shared secret test failed: %v", err)
	}
	if !match {
		t.Error("ECDH shared secret test should pass")
	}

	// 5. Use shared secret as encryption key (simulate)
	if len(aliceShared) != 32 || len(bobShared) != 32 {
		t.Error("Shared secrets should be 32 bytes")
	}

	// Test that different key pairs produce different secrets
	charlie, _ := GenerateECDHP256Key()
	charlieShared, _ := ECDHKeyExchange(alice.PrivateKey, charlie.PublicKey, 32)

	if string(aliceShared) == string(charlieShared) {
		t.Error("Different key pairs should produce different shared secrets")
	}
}