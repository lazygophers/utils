package cryptox

import (
	"crypto/ecdsa"
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

// TestECDHMissingCoverage tests specific ECDH error paths to reach 100% coverage
func TestECDHMissingCoverage(t *testing.T) {
	// Test ECDHComputeShared with public key not on curve (artificial case)
	// This will test the curve validation path
	keyPair1, err := GenerateECDHP256Key()
	if err != nil {
		t.Fatalf("Failed to generate key pair 1: %v", err)
	}

	keyPair2, err := GenerateECDHP256Key() 
	if err != nil {
		t.Fatalf("Failed to generate key pair 2: %v", err)
	}

	// Test different curves mismatch
	keyPairP384, err := GenerateECDHP384Key()
	if err != nil {
		t.Fatalf("Failed to generate P384 key: %v", err)
	}

	_, err = ECDHComputeShared(keyPair1.PrivateKey, keyPairP384.PublicKey)
	if err == nil {
		t.Error("Expected error for curve mismatch")
	}

	// Test ValidateECDHKeyPair with mismatched keys
	// Create key pair with wrong public key
	wrongKeyPair := &ECDHKeyPair{
		PrivateKey: keyPair1.PrivateKey,
		PublicKey:  keyPair2.PublicKey, // Different public key
	}

	err = ValidateECDHKeyPair(wrongKeyPair)
	if err == nil {
		t.Error("Expected error for mismatched key pair")
	}

	// Test ECDHSharedSecretTest to ensure it executes without error
	// We don't assert on the match result since different keys could theoretically match
	_, err = ECDHSharedSecretTest(keyPair1, keyPair2)
	if err != nil {
		t.Errorf("ECDHSharedSecretTest should not error with valid keys: %v", err)
	}

	// Test with same key pair (should match)
	match, err := ECDHSharedSecretTest(keyPair1, keyPair1)
	if err != nil {
		t.Errorf("ECDHSharedSecretTest should not error with same key: %v", err)
	}
	if !match {
		t.Error("Same key pair should have matching secrets")
	}

	// Test ECDHComputeSharedWithKDF with key length exactly equal to hash length
	derivedKey, err := ECDHComputeSharedWithKDF(keyPair1.PrivateKey, keyPair2.PublicKey, 32, sha256.New)
	if err != nil {
		t.Errorf("KDF should work with key length equal to hash length: %v", err)
	}
	if len(derivedKey) != 32 {
		t.Errorf("Expected key length 32, got %d", len(derivedKey))
	}

	// Test with key length less than hash length (truncation path)
	derivedKey2, err := ECDHComputeSharedWithKDF(keyPair1.PrivateKey, keyPair2.PublicKey, 16, sha256.New)
	if err != nil {
		t.Errorf("KDF should work with key length less than hash length: %v", err)
	}
	if len(derivedKey2) != 16 {
		t.Errorf("Expected key length 16, got %d", len(derivedKey2))
	}
}

// TestPEMEncodingErrorPaths tests the remaining PEM encoding/decoding paths
func TestPEMEncodingErrorPaths(t *testing.T) {
	// Generate test keys
	ecdsaKeyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("Failed to generate ECDSA key pair: %v", err)
	}

	// Test ECDSAPrivateKeyToPEM success path and internal error handling
	pemData, err := ECDSAPrivateKeyToPEM(ecdsaKeyPair.PrivateKey)
	if err != nil {
		t.Errorf("ECDSAPrivateKeyToPEM should succeed: %v", err)
	}
	if len(pemData) == 0 {
		t.Error("PEM data should not be empty")
	}

	// Test ECDSAPublicKeyToPEM success path
	pubPemData, err := ECDSAPublicKeyToPEM(ecdsaKeyPair.PublicKey)
	if err != nil {
		t.Errorf("ECDSAPublicKeyToPEM should succeed: %v", err)
	}
	if len(pubPemData) == 0 {
		t.Error("Public PEM data should not be empty")
	}

	// Test PEM parsing with edge cases that trigger different code paths

	// Test malformed ASN.1 in private key PEM
	malformedPrivatePEM := `-----BEGIN EC PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQg
-----END EC PRIVATE KEY-----`

	_, err = ECDSAPrivateKeyFromPEM([]byte(malformedPrivatePEM))
	if err == nil {
		t.Error("Expected error for malformed private key ASN.1")
	}

	// Test malformed ASN.1 in public key PEM  
	malformedPublicPEM := `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE
-----END PUBLIC KEY-----`

	_, err = ECDSAPublicKeyFromPEM([]byte(malformedPublicPEM))
	if err == nil {
		t.Error("Expected error for malformed public key ASN.1")
	}

	// Test with non-EC algorithms (should fail the type assertion)
	rsaPublicPEM := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1hKkprG8j0jLH9s7IqiK
6VgKtCWNOhd2HWqe1bF/Xy4yU1R1ZRswRWuAHVEJl2Gx4H2Xc7k/lP7I0nN1Q8Xs
F9Pl5b3r6ZkQXVmV1tTc8GmZ7u4mQ3b1x8cNlmPX7r7a6Q5H8VvT1K2tG3k6w8Kn
8r7S7Gx1V8R5L5S8s2r5qT7QGk7o5R8L8F2I7d2uV1X7cM4W8K8r7S7K7I1M7V7Q
7K7J3Q8O7p5f2c4D3b8r7X6K8s7S7Q7K7k7G1V8L8F2I7t4n7w7S4U7o3h6G3s3S3K
5Q4I3j4K3m7p7N5T7w7R9s7T7K7Q8K7G3s7S7K8n8s7w6h7I4K2Q4K2A6J4K2E5E7X
2Q3W3E2G3X7N5C7V3K5L8K5Q7K2A5D7P1K5Q3K3A5X3K5N5c2L5T5w7I2I2K8K8K5A
5D7K7J7w7T4w7r7b3j4k7s5j7k5s7s5v7e7t3e6h2c4n5j7h1j2z4m8h6f7j8i5l7l
5k7s5j7k5s7s5v7e7t3e6h2c4n5j7h1j2z4m8h6f7j8i5l7l5k7s5j7k5s7s5v7e7t
3e6h2c4n5j7h1j2z4m8h6f7j8i5l7l5k7s5j7k5s7s5v7e7t3e6h2c4n5j7h1j2z4m
8h6f7j8i5l7l5k7s5j7k5s7s5v7e7t3e6h2c4n5j7h1j2z4m8h6f7j8i5l7lQIDAQAB
-----END PUBLIC KEY-----`

	_, err = ECDSAPublicKeyFromPEM([]byte(rsaPublicPEM))
	if err == nil {
		t.Error("Expected error for RSA public key in ECDSA parser")
	}
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
		PrivateKey: keyPair1.PrivateKey,       // P256
		PublicKey:  keyPairP384.PublicKey,     // P384
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
}

// TestECDSASpecificCoveragePaths tests ECDSA specific paths for 100% coverage
func TestECDSASpecificCoveragePaths(t *testing.T) {
	// Test successful PEM operations to trigger success paths
	keyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("Failed to generate ECDSA key pair: %v", err)
	}

	// Test ECDSAPrivateKeyToPEM success path (85.7% -> 100%)
	privPEM, err := ECDSAPrivateKeyToPEM(keyPair.PrivateKey)
	if err != nil {
		t.Errorf("ECDSAPrivateKeyToPEM should succeed: %v", err)
	}
	if len(privPEM) == 0 {
		t.Error("Private PEM should not be empty")
	}

	// Test ECDSAPublicKeyToPEM success path (85.7% -> 100%)
	pubPEM, err := ECDSAPublicKeyToPEM(keyPair.PublicKey)
	if err != nil {
		t.Errorf("ECDSAPublicKeyToPEM should succeed: %v", err)
	}
	if len(pubPEM) == 0 {
		t.Error("Public PEM should not be empty")
	}

	// Test ECDSAPrivateKeyFromPEM success path (90.9% -> 100%)
	decodedPriv, err := ECDSAPrivateKeyFromPEM(privPEM)
	if err != nil {
		t.Errorf("ECDSAPrivateKeyFromPEM should succeed: %v", err)
	}
	if decodedPriv == nil {
		t.Error("Decoded private key should not be nil")
	}

	// Test ECDSAPublicKeyFromPEM success path (78.6% -> 100%)
	decodedPub, err := ECDSAPublicKeyFromPEM(pubPEM)
	if err != nil {
		t.Errorf("ECDSAPublicKeyFromPEM should succeed: %v", err)
	}
	if decodedPub == nil {
		t.Error("Decoded public key should not be nil")
	}

	// Test with malformed DER signature to trigger ECDSASignatureFromBytes error paths
	
	// DER with sequence tag but wrong type for first integer
	malformedDER1 := []byte{0x30, 0x08, 0x04, 0x02, 0x01, 0x01, 0x02, 0x02, 0x01, 0x01}
	_, _, err = ECDSASignatureFromBytes(malformedDER1)
	if err == nil {
		t.Error("Expected error for malformed DER with wrong type")
	}

	// DER with sequence but truncated after r length
	malformedDER2 := []byte{0x30, 0x08, 0x02, 0x02, 0x01}
	_, _, err = ECDSASignatureFromBytes(malformedDER2)
	if err == nil {
		t.Error("Expected error for truncated DER after r length")
	}

	// DER with sequence but wrong type for second integer (s)
	malformedDER3 := []byte{0x30, 0x08, 0x02, 0x02, 0x01, 0x01, 0x04, 0x02, 0x01, 0x01}
	_, _, err = ECDSASignatureFromBytes(malformedDER3)
	if err == nil {
		t.Error("Expected error for malformed DER with wrong s type")
	}

	// DER with sequence but truncated after s length
	malformedDER4 := []byte{0x30, 0x08, 0x02, 0x02, 0x01, 0x01, 0x02, 0x02, 0x01}
	_, _, err = ECDSASignatureFromBytes(malformedDER4)
	if err == nil {
		t.Error("Expected error for truncated DER after s length")
	}

	// DER with truncated sequence length
	malformedDER5 := []byte{0x30, 0xFF, 0x02, 0x02, 0x01, 0x01}
	_, _, err = ECDSASignatureFromBytes(malformedDER5)
	if err == nil {
		t.Error("Expected error for truncated sequence")
	}

	// Test ECDSAPrivateKeyFromPEM with invalid ASN.1 structure
	invalidASN1PEM := []byte(`-----BEGIN EC PRIVATE KEY-----
MIGUAgEAMBAGByqGSM49AgEGBSuBBAAKBHQwcgIBAQQgTest
-----END EC PRIVATE KEY-----`)
	
	_, err = ECDSAPrivateKeyFromPEM(invalidASN1PEM)
	if err == nil {
		t.Error("Expected error for invalid ASN.1 structure")
	}

	// Test ECDSAPublicKeyFromPEM with invalid algorithm identifier
	invalidAlgoIdPEM := []byte(`-----BEGIN PUBLIC KEY-----
MFMwDQYJKoZIhvcNAQEBBQADQgAwPwIBAAKCAQEAxxx
-----END PUBLIC KEY-----`)
	
	_, err = ECDSAPublicKeyFromPEM(invalidAlgoIdPEM)
	if err == nil {
		t.Error("Expected error for invalid algorithm identifier")
	}

	// Test with completely invalid PEM content
	invalidPEMContent := []byte(`-----BEGIN EC PRIVATE KEY-----
invalid base64 content here!!!
-----END EC PRIVATE KEY-----`)
	
	_, err = ECDSAPrivateKeyFromPEM(invalidPEMContent)
	if err == nil {
		t.Error("Expected error for invalid PEM content")
	}

	// Test public key PEM with invalid content
	invalidPublicPEMContent := []byte(`-----BEGIN PUBLIC KEY-----
invalid base64 content here!!!
-----END PUBLIC KEY-----`)
	
	_, err = ECDSAPublicKeyFromPEM(invalidPublicPEMContent)
	if err == nil {
		t.Error("Expected error for invalid public PEM content")
	}
}

// TestECDHAndECDSAMissingCoverage specifically targets the remaining uncovered lines
func TestECDHAndECDSAMissingCoverage(t *testing.T) {
	// Generate key pairs for testing
	keyPair1, err := GenerateECDHP256Key()
	if err != nil {
		t.Fatalf("Failed to generate ECDH key pair 1: %v", err)
	}

	keyPair2, err := GenerateECDHP256Key()
	if err != nil {
		t.Fatalf("Failed to generate ECDH key pair 2: %v", err)
	}

	// Test ECDHSharedSecretTest with error path
	// Create an invalid key pair that will cause ECDHComputeShared to fail
	invalidPublicKey := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     big.NewInt(0), // Point at infinity (invalid)
		Y:     big.NewInt(1),
	}

	invalidKeyPair := &ECDHKeyPair{
		PrivateKey: keyPair1.PrivateKey,
		PublicKey:  invalidPublicKey,
	}

	// This should trigger the error path in ECDHSharedSecretTest (71.4% -> 100%)
	_, err = ECDHSharedSecretTest(keyPair1, invalidKeyPair)
	if err == nil {
		t.Error("Expected error in ECDHSharedSecretTest with invalid public key")
	}

	// Test second error path in ECDHSharedSecretTest
	_, err = ECDHSharedSecretTest(invalidKeyPair, keyPair1)
	if err == nil {
		t.Error("Expected error in ECDHSharedSecretTest with invalid key pair as first param")
	}

	// Test ValidateECDHKeyPair with all error conditions (92.9% -> 100%)
	
	// Test with nil key pair
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

	// Test ValidateECDHKeyPair with the ECDHKeyPair struct method
	validPair := &ECDHKeyPair{
		PrivateKey: keyPair1.PrivateKey,
		PublicKey:  keyPair1.PublicKey,
	}
	err = ValidateECDHKeyPair(validPair)
	if err != nil {
		t.Errorf("ValidateECDHKeyPair should succeed with valid pair: %v", err)
	}

	// Test ECDHComputeSharedWithKDF with all error paths (95.5% -> 100%)
	
	// Test with 0 key length
	_, err = ECDHComputeSharedWithKDF(keyPair1.PrivateKey, keyPair2.PublicKey, 0, nil)
	if err == nil {
		t.Error("Expected error for zero key length in ECDHComputeSharedWithKDF")
	}

	// Test with negative key length
	_, err = ECDHComputeSharedWithKDF(keyPair1.PrivateKey, keyPair2.PublicKey, -1, nil)
	if err == nil {
		t.Error("Expected error for negative key length in ECDHComputeSharedWithKDF")
	}

	// Test with nil KDF function (line 89)
	_, err = ECDHComputeSharedWithKDF(keyPair1.PrivateKey, keyPair2.PublicKey, 32, nil)
	if err == nil {
		t.Error("Expected error for nil KDF function in ECDHComputeSharedWithKDF")
	}

	// Test with invalid public key (should trigger ECDHComputeShared error)
	_, err = ECDHComputeSharedWithKDF(keyPair1.PrivateKey, invalidPublicKey, 32, sha256.New)
	if err == nil {
		t.Error("Expected error for invalid public key in ECDHComputeSharedWithKDF")
	}

	// Test counter mode expansion (key length > hash output length)
	largeKey, err := ECDHComputeSharedWithKDF(keyPair1.PrivateKey, keyPair2.PublicKey, 100, sha256.New)
	if err != nil {
		t.Errorf("ECDHComputeSharedWithKDF should work with large key length: %v", err)
	}
	if len(largeKey) != 100 {
		t.Errorf("Expected key length 100, got %d", len(largeKey))
	}

	// Test ECDHComputeShared with all error paths (90.0% -> 100%)
	
	// Test with nil private key
	_, err = ECDHComputeShared(nil, keyPair1.PublicKey)
	if err == nil {
		t.Error("Expected error for nil private key in ECDHComputeShared")
	}

	// Test with nil public key
	_, err = ECDHComputeShared(keyPair1.PrivateKey, nil)
	if err == nil {
		t.Error("Expected error for nil public key in ECDHComputeShared")
	}

	// Test with invalid public key coordinates (not on curve)
	_, err = ECDHComputeShared(keyPair1.PrivateKey, invalidPublicKey)
	if err == nil {
		t.Error("Expected error for public key not on curve in ECDHComputeShared")
	}

	// Test curve mismatch
	keyPairP384, err := GenerateECDHP384Key()
	if err != nil {
		t.Fatalf("Failed to generate P384 key pair: %v", err)
	}
	
	_, err = ECDHComputeShared(keyPair1.PrivateKey, keyPairP384.PublicKey)
	if err == nil {
		t.Error("Expected error for curve mismatch in ECDHComputeShared")
	}

	// Test successful ECDHComputeShared to hit the success path (line 77-80)
	sharedSecret, err := ECDHComputeShared(keyPair1.PrivateKey, keyPair2.PublicKey)
	if err != nil {
		t.Errorf("ECDHComputeShared should succeed with valid keys: %v", err)
	}
	if len(sharedSecret) == 0 {
		t.Error("Shared secret should not be empty")
	}

	// Test ECDHSharedSecretTest with nil key pairs (to trigger nil check)
	_, err = ECDHSharedSecretTest(nil, keyPair1)
	if err == nil {
		t.Error("Expected error for nil keyPair1 in ECDHSharedSecretTest")
	}

	_, err = ECDHSharedSecretTest(keyPair1, nil)
	if err == nil {
		t.Error("Expected error for nil keyPair2 in ECDHSharedSecretTest")
	}

	// Test ECDHSharedSecretTest to trigger the secret length comparison and byte comparison
	// Create two valid but different key pairs - this should produce different secrets
	match, err := ECDHSharedSecretTest(keyPair1, keyPair2)
	if err != nil {
		t.Errorf("ECDHSharedSecretTest with valid but different keys should not error: %v", err)
	}
	// These should not match since they're different key pairs
	if match {
		t.Log("Note: Different key pairs produced matching secrets (unlikely but possible)")
	}

	// Test ECDHSharedSecretTest to trigger specific byte comparison path  
	// This tests the return false path at line 219
	keyPair3, _ := GenerateECDHP256Key()
	keyPair4, _ := GenerateECDHP256Key()
	
	// These should not match 
	match, err = ECDHSharedSecretTest(keyPair3, keyPair4)
	if err != nil {
		t.Errorf("ECDHSharedSecretTest should not error: %v", err)
	}
	// The secrets from different key pairs should not match
	if match {
		t.Log("Note: Different key pairs produced matching secrets (very unlikely)")
	}

	// Test ValidateECDHKeyPair with key pair where public key doesn't match private key
	// Create a mismatched key pair (same curve, but wrong public key)
	mismatchedPubKeyPair := &ECDHKeyPair{
		PrivateKey: keyPair1.PrivateKey,
		PublicKey:  keyPair2.PublicKey,
	}
	err = ValidateECDHKeyPair(mismatchedPubKeyPair)
	if err == nil {
		t.Error("Expected error for public key not matching private key")
	}
}

// TestFinalCoverageEdgeCases targets the last remaining uncovered lines
func TestFinalCoverageEdgeCases(t *testing.T) {
	// Generate ECDSA key pairs for PEM testing
	ecdsaKeyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("Failed to generate ECDSA key pair: %v", err)
	}

	// Test the success path of ECDSAPrivateKeyToPEM to reach 100%
	privPEM, err := ECDSAPrivateKeyToPEM(ecdsaKeyPair.PrivateKey)
	if err != nil {
		t.Errorf("ECDSAPrivateKeyToPEM should succeed: %v", err)
	}

	// Test the success path of ECDSAPublicKeyToPEM to reach 100%
	pubPEM, err := ECDSAPublicKeyToPEM(ecdsaKeyPair.PublicKey)
	if err != nil {
		t.Errorf("ECDSAPublicKeyToPEM should succeed: %v", err)
	}

	// Test ECDSA PEM round-trip to trigger all success paths
	decodedPriv, err := ECDSAPrivateKeyFromPEM(privPEM)
	if err != nil {
		t.Errorf("ECDSAPrivateKeyFromPEM should succeed: %v", err)
	}
	if decodedPriv == nil {
		t.Error("Decoded private key should not be nil")
	}

	decodedPub, err := ECDSAPublicKeyFromPEM(pubPEM)
	if err != nil {
		t.Errorf("ECDSAPublicKeyFromPEM should succeed: %v", err)
	}
	if decodedPub == nil {
		t.Error("Decoded public key should not be nil")
	}

	// Test ECDSASignatureFromBytes with various malformed inputs to trigger all error paths
	
	// Empty input
	_, _, err = ECDSASignatureFromBytes([]byte{})
	if err == nil {
		t.Error("Expected error for empty signature bytes")
	}

	// Non-sequence tag
	malformedDER := []byte{0x04, 0x08, 0x02, 0x02, 0x01, 0x01, 0x02, 0x02, 0x01, 0x01}
	_, _, err = ECDSASignatureFromBytes(malformedDER)
	if err == nil {
		t.Error("Expected error for non-sequence tag")
	}

	// Too short to read sequence length
	malformedDER = []byte{0x30}
	_, _, err = ECDSASignatureFromBytes(malformedDER)
	if err == nil {
		t.Error("Expected error for truncated sequence length")
	}

	// Sequence length longer than remaining bytes
	malformedDER = []byte{0x30, 0xFF, 0x02, 0x02}
	_, _, err = ECDSASignatureFromBytes(malformedDER)
	if err == nil {
		t.Error("Expected error for sequence length exceeding data")
	}

	// Test successful signature round-trip to ensure all paths are hit
	// Create a signature and then decode it to hit success paths
	r, s, err := ECDSASignSHA256(ecdsaKeyPair.PrivateKey, []byte("test message"))
	if err != nil {
		t.Errorf("ECDSASignSHA256 should succeed: %v", err)
	}

	sigBytes, err := ECDSASignatureToBytes(r, s)
	if err != nil {
		t.Errorf("ECDSASignatureToBytes should succeed: %v", err)
	}

	decodedR, decodedS, err := ECDSASignatureFromBytes(sigBytes)
	if err != nil {
		t.Errorf("ECDSASignatureFromBytes should succeed: %v", err)
	}
	if decodedR.Cmp(r) != 0 || decodedS.Cmp(s) != 0 {
		t.Error("Decoded signature should match original")
	}
}