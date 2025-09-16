package cryptox

import (
	"crypto/elliptic"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"math/big"
	"strings"
	"testing"
)

// Test data
const (
	testECDSAMessage = "Hello, ECDSA test message!"
)

// TestGenerateECDSAKey tests ECDSA key generation with different curves
func TestGenerateECDSAKey(t *testing.T) {
	curves := []elliptic.Curve{
		elliptic.P224(),
		elliptic.P256(),
		elliptic.P384(),
		elliptic.P521(),
	}

	for _, curve := range curves {
		keyPair, err := GenerateECDSAKey(curve)
		if err != nil {
			t.Errorf("GenerateECDSAKey failed for curve %s: %v", GetCurveName(curve), err)
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

// TestGenerateECDSAP256Key tests P-256 key generation
func TestGenerateECDSAP256Key(t *testing.T) {
	keyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("GenerateECDSAP256Key failed: %v", err)
	}

	if keyPair.PrivateKey.Curve != elliptic.P256() {
		t.Error("Expected P-256 curve")
	}
}

// TestGenerateECDSAP384Key tests P-384 key generation
func TestGenerateECDSAP384Key(t *testing.T) {
	keyPair, err := GenerateECDSAP384Key()
	if err != nil {
		t.Fatalf("GenerateECDSAP384Key failed: %v", err)
	}

	if keyPair.PrivateKey.Curve != elliptic.P384() {
		t.Error("Expected P-384 curve")
	}
}

// TestGenerateECDSAP521Key tests P-521 key generation
func TestGenerateECDSAP521Key(t *testing.T) {
	keyPair, err := GenerateECDSAP521Key()
	if err != nil {
		t.Fatalf("GenerateECDSAP521Key failed: %v", err)
	}

	if keyPair.PrivateKey.Curve != elliptic.P521() {
		t.Error("Expected P-521 curve")
	}
}

// TestECDSASignAndVerify tests ECDSA signing and verification
func TestECDSASignAndVerify(t *testing.T) {
	keyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("GenerateECDSAP256Key failed: %v", err)
	}

	data := []byte(testECDSAMessage)

	// Test signing with SHA256
	r, s, err := ECDSASignSHA256(keyPair.PrivateKey, data)
	if err != nil {
		t.Fatalf("ECDSASignSHA256 failed: %v", err)
	}

	if r == nil || s == nil {
		t.Fatal("Signature components cannot be nil")
	}

	// Test verification with correct key
	if !ECDSAVerifySHA256(keyPair.PublicKey, data, r, s) {
		t.Error("ECDSAVerifySHA256 should verify correct signature")
	}

	// Test verification with wrong data
	wrongData := []byte("wrong message")
	if ECDSAVerifySHA256(keyPair.PublicKey, wrongData, r, s) {
		t.Error("ECDSAVerifySHA256 should not verify wrong data")
	}

	// Test verification with wrong signature
	wrongR := new(big.Int).Add(r, big.NewInt(1))
	if ECDSAVerifySHA256(keyPair.PublicKey, data, wrongR, s) {
		t.Error("ECDSAVerifySHA256 should not verify wrong signature")
	}
}

// TestECDSASignAndVerifySHA512 tests ECDSA signing and verification with SHA512
func TestECDSASignAndVerifySHA512(t *testing.T) {
	keyPair, err := GenerateECDSAP384Key()
	if err != nil {
		t.Fatalf("GenerateECDSAP384Key failed: %v", err)
	}

	data := []byte(testECDSAMessage)

	// Test signing with SHA512
	r, s, err := ECDSASignSHA512(keyPair.PrivateKey, data)
	if err != nil {
		t.Fatalf("ECDSASignSHA512 failed: %v", err)
	}

	// Test verification
	if !ECDSAVerifySHA512(keyPair.PublicKey, data, r, s) {
		t.Error("ECDSAVerifySHA512 should verify correct signature")
	}
}

// TestECDSAGenericSignAndVerify tests generic ECDSA sign/verify functions
func TestECDSAGenericSignAndVerify(t *testing.T) {
	keyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("GenerateECDSAP256Key failed: %v", err)
	}

	data := []byte(testECDSAMessage)

	// Test generic signing
	r, s, err := ECDSASign(keyPair.PrivateKey, data, sha256.New)
	if err != nil {
		t.Fatalf("ECDSASign failed: %v", err)
	}

	// Test generic verification
	if !ECDSAVerify(keyPair.PublicKey, data, r, s, sha256.New) {
		t.Error("ECDSAVerify should verify correct signature")
	}

	// Test with different hash function
	r2, s2, err := ECDSASign(keyPair.PrivateKey, data, sha512.New)
	if err != nil {
		t.Fatalf("ECDSASign with SHA512 failed: %v", err)
	}

	if !ECDSAVerify(keyPair.PublicKey, data, r2, s2, sha512.New) {
		t.Error("ECDSAVerify with SHA512 should verify correct signature")
	}

	// Test cross-hash verification (should fail)
	if ECDSAVerify(keyPair.PublicKey, data, r, s, sha512.New) {
		t.Error("Cross-hash verification should fail")
	}
}

// TestECDSAPrivateKeyPEM tests private key PEM encoding/decoding
func TestECDSAPrivateKeyPEM(t *testing.T) {
	keyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("GenerateECDSAP256Key failed: %v", err)
	}

	// Test encoding
	pemData, err := ECDSAPrivateKeyToPEM(keyPair.PrivateKey)
	if err != nil {
		t.Fatalf("ECDSAPrivateKeyToPEM failed: %v", err)
	}

	if len(pemData) == 0 {
		t.Fatal("PEM data cannot be empty")
	}

	if !strings.Contains(string(pemData), "BEGIN EC PRIVATE KEY") {
		t.Error("PEM data should contain EC PRIVATE KEY header")
	}

	// Test decoding
	decodedKey, err := ECDSAPrivateKeyFromPEM(pemData)
	if err != nil {
		t.Fatalf("ECDSAPrivateKeyFromPEM failed: %v", err)
	}

	// Verify the key is correct
	if decodedKey.D.Cmp(keyPair.PrivateKey.D) != 0 {
		t.Error("Decoded private key D component mismatch")
	}

	if decodedKey.PublicKey.X.Cmp(keyPair.PrivateKey.PublicKey.X) != 0 {
		t.Error("Decoded public key X component mismatch")
	}

	if decodedKey.PublicKey.Y.Cmp(keyPair.PrivateKey.PublicKey.Y) != 0 {
		t.Error("Decoded public key Y component mismatch")
	}
}

// TestECDSAPublicKeyPEM tests public key PEM encoding/decoding
func TestECDSAPublicKeyPEM(t *testing.T) {
	keyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("GenerateECDSAP256Key failed: %v", err)
	}

	// Test encoding
	pemData, err := ECDSAPublicKeyToPEM(keyPair.PublicKey)
	if err != nil {
		t.Fatalf("ECDSAPublicKeyToPEM failed: %v", err)
	}

	if len(pemData) == 0 {
		t.Fatal("PEM data cannot be empty")
	}

	if !strings.Contains(string(pemData), "BEGIN PUBLIC KEY") {
		t.Error("PEM data should contain PUBLIC KEY header")
	}

	// Test decoding
	decodedKey, err := ECDSAPublicKeyFromPEM(pemData)
	if err != nil {
		t.Fatalf("ECDSAPublicKeyFromPEM failed: %v", err)
	}

	// Verify the key is correct
	if decodedKey.X.Cmp(keyPair.PublicKey.X) != 0 {
		t.Error("Decoded public key X component mismatch")
	}

	if decodedKey.Y.Cmp(keyPair.PublicKey.Y) != 0 {
		t.Error("Decoded public key Y component mismatch")
	}

	if decodedKey.Curve != keyPair.PublicKey.Curve {
		t.Error("Decoded public key curve mismatch")
	}
}

// TestECDSASignatureBytes tests signature byte encoding/decoding
func TestECDSASignatureBytes(t *testing.T) {
	keyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("GenerateECDSAP256Key failed: %v", err)
	}

	data := []byte(testECDSAMessage)

	// Generate signature
	r, s, err := ECDSASignSHA256(keyPair.PrivateKey, data)
	if err != nil {
		t.Fatalf("ECDSASignSHA256 failed: %v", err)
	}

	// Test encoding
	sigBytes, err := ECDSASignatureToBytes(r, s)
	if err != nil {
		t.Fatalf("ECDSASignatureToBytes failed: %v", err)
	}

	if len(sigBytes) == 0 {
		t.Fatal("Signature bytes cannot be empty")
	}

	// Test decoding
	decodedR, decodedS, err := ECDSASignatureFromBytes(sigBytes)
	if err != nil {
		t.Fatalf("ECDSASignatureFromBytes failed: %v", err)
	}

	// Verify components match
	if decodedR.Cmp(r) != 0 {
		t.Error("Decoded R component mismatch")
	}

	if decodedS.Cmp(s) != 0 {
		t.Error("Decoded S component mismatch")
	}

	// Verify signature still works
	if !ECDSAVerifySHA256(keyPair.PublicKey, data, decodedR, decodedS) {
		t.Error("Decoded signature should still verify")
	}
}

// TestGetCurveName tests curve name function
func TestGetCurveName(t *testing.T) {
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
		name := GetCurveName(tc.curve)
		if name != tc.name {
			t.Errorf("Expected curve name %s, got %s", tc.name, name)
		}
	}

	// Test unknown curve
	if GetCurveName(nil) != "Unknown" {
		t.Error("Expected 'Unknown' for nil curve")
	}
}

// TestIsValidCurve tests curve validation function
func TestIsValidCurve(t *testing.T) {
	validCurves := []elliptic.Curve{
		elliptic.P224(),
		elliptic.P256(),
		elliptic.P384(),
		elliptic.P521(),
	}

	for _, curve := range validCurves {
		if !IsValidCurve(curve) {
			t.Errorf("Curve %s should be valid", GetCurveName(curve))
		}
	}

	if IsValidCurve(nil) {
		t.Error("nil curve should be invalid")
	}
}

// TestECDSAErrorConditions tests error conditions
func TestECDSAErrorConditions(t *testing.T) {
	// Test GenerateECDSAKey with nil curve
	_, err := GenerateECDSAKey(nil)
	if err == nil {
		t.Error("Expected error for nil curve")
	}

	// Test ECDSASign with nil private key
	_, _, err = ECDSASign(nil, []byte("test"), sha256.New)
	if err == nil {
		t.Error("Expected error for nil private key")
	}

	// Test ECDSASign with nil hash function
	keyPair, _ := GenerateECDSAP256Key()
	_, _, err = ECDSASign(keyPair.PrivateKey, []byte("test"), nil)
	if err == nil {
		t.Error("Expected error for nil hash function")
	}

	// Test ECDSAVerify with nil components
	if ECDSAVerify(nil, []byte("test"), big.NewInt(1), big.NewInt(1), sha256.New) {
		t.Error("Should return false for nil public key")
	}

	if ECDSAVerify(keyPair.PublicKey, []byte("test"), nil, big.NewInt(1), sha256.New) {
		t.Error("Should return false for nil r")
	}

	if ECDSAVerify(keyPair.PublicKey, []byte("test"), big.NewInt(1), nil, sha256.New) {
		t.Error("Should return false for nil s")
	}

	if ECDSAVerify(keyPair.PublicKey, []byte("test"), big.NewInt(1), big.NewInt(1), nil) {
		t.Error("Should return false for nil hash function")
	}

	// Test PEM functions with nil keys
	_, err = ECDSAPrivateKeyToPEM(nil)
	if err == nil {
		t.Error("Expected error for nil private key")
	}

	_, err = ECDSAPublicKeyToPEM(nil)
	if err == nil {
		t.Error("Expected error for nil public key")
	}

	// Test PEM functions with invalid data
	_, err = ECDSAPrivateKeyFromPEM(nil)
	if err == nil {
		t.Error("Expected error for nil PEM data")
	}

	_, err = ECDSAPrivateKeyFromPEM([]byte(""))
	if err == nil {
		t.Error("Expected error for empty PEM data")
	}

	_, err = ECDSAPrivateKeyFromPEM([]byte("invalid pem"))
	if err == nil {
		t.Error("Expected error for invalid PEM data")
	}

	_, err = ECDSAPublicKeyFromPEM([]byte("invalid pem"))
	if err == nil {
		t.Error("Expected error for invalid PEM data")
	}

	// Test signature bytes functions with nil components
	_, err = ECDSASignatureToBytes(nil, big.NewInt(1))
	if err == nil {
		t.Error("Expected error for nil r")
	}

	_, err = ECDSASignatureToBytes(big.NewInt(1), nil)
	if err == nil {
		t.Error("Expected error for nil s")
	}

	// Test signature bytes decoding with invalid data
	_, _, err = ECDSASignatureFromBytes([]byte{})
	if err == nil {
		t.Error("Expected error for empty signature data")
	}

	_, _, err = ECDSASignatureFromBytes([]byte{0x01, 0x02})
	if err == nil {
		t.Error("Expected error for invalid signature data")
	}
}

// TestECDSADifferentDataSizes tests ECDSA with different data sizes
func TestECDSADifferentDataSizes(t *testing.T) {
	keyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("GenerateECDSAP256Key failed: %v", err)
	}

	testSizes := []int{0, 1, 16, 64, 256, 1024}

	for _, size := range testSizes {
		data := make([]byte, size)
		for i := range data {
			data[i] = byte(i % 256)
		}

		// Test signing
		r, s, err := ECDSASignSHA256(keyPair.PrivateKey, data)
		if err != nil {
			t.Errorf("ECDSASignSHA256 failed for data size %d: %v", size, err)
			continue
		}

		// Test verification
		if !ECDSAVerifySHA256(keyPair.PublicKey, data, r, s) {
			t.Errorf("ECDSAVerifySHA256 failed for data size %d", size)
		}
	}
}

// ==== MERGED ECDSA TESTS FROM ecc_100_coverage_test.go ====

// Mock failures for ECDSA dependency injection
type FailingECDSAReader struct{}

func (fr FailingECDSAReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("simulated random reader failure")
}

// TestECDSA_100PercentCoverage tests all error paths in ECDSA functions
func TestECDSA_100PercentCoverage(t *testing.T) {
	// Save original functions
	originalECDSARandReader := ecdsaRandReader

	// Restore original functions after test
	defer func() {
		ecdsaRandReader = originalECDSARandReader
	}()

	// Test 1: Trigger rand.Reader failure in ECDSA key generation
	ecdsaRandReader = FailingECDSAReader{}

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

	// Test 2: Trigger rand.Reader failure in ECDSA signing
	// Generate a valid key first with original reader
	ecdsaRandReader = originalECDSARandReader
	keyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("Failed to generate test key: %v", err)
	}

	// Now set failing reader and test signing
	ecdsaRandReader = FailingECDSAReader{}

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
}

// TestECDSASpecificCoveragePaths tests ECDSA specific paths for 100% coverage
func TestECDSASpecificCoveragePaths(t *testing.T) {
	// Test successful PEM operations to trigger success paths
	keyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("Failed to generate ECDSA key pair: %v", err)
	}

	// Test ECDSAPrivateKeyToPEM success path
	privPEM, err := ECDSAPrivateKeyToPEM(keyPair.PrivateKey)
	if err != nil {
		t.Errorf("ECDSAPrivateKeyToPEM should succeed: %v", err)
	}
	if len(privPEM) == 0 {
		t.Error("Private PEM should not be empty")
	}

	// Test ECDSAPublicKeyToPEM success path
	pubPEM, err := ECDSAPublicKeyToPEM(keyPair.PublicKey)
	if err != nil {
		t.Errorf("ECDSAPublicKeyToPEM should succeed: %v", err)
	}
	if len(pubPEM) == 0 {
		t.Error("Public PEM should not be empty")
	}

	// Test ECDSAPrivateKeyFromPEM success path
	decodedPriv, err := ECDSAPrivateKeyFromPEM(privPEM)
	if err != nil {
		t.Errorf("ECDSAPrivateKeyFromPEM should succeed: %v", err)
	}
	if decodedPriv == nil {
		t.Error("Decoded private key should not be nil")
	}

	// Test ECDSAPublicKeyFromPEM success path
	decodedPub, err := ECDSAPublicKeyFromPEM(pubPEM)
	if err != nil {
		t.Errorf("ECDSAPublicKeyFromPEM should succeed: %v", err)
	}
	if decodedPub == nil {
		t.Error("Decoded public key should not be nil")
	}

	// Test various error paths for signature parsing
	_, _, err = ECDSASignatureFromBytes([]byte{})
	if err == nil {
		t.Error("Expected error for empty signature bytes")
	}

	// Test nil parameter errors
	_, err = ECDSAPrivateKeyToPEM(nil)
	if err == nil {
		t.Error("Expected error for nil private key")
	}

	_, err = ECDSAPublicKeyToPEM(nil)
	if err == nil {
		t.Error("Expected error for nil public key")
	}

	_, _, err = ECDSASign(nil, []byte("test"), sha256.New)
	if err == nil {
		t.Error("Expected error for nil private key in sign")
	}

	_, _, err = ECDSASign(keyPair.PrivateKey, []byte("test"), nil)
	if err == nil {
		t.Error("Expected error for nil hash function")
	}

	if ECDSAVerify(nil, []byte("test"), big.NewInt(1), big.NewInt(1), sha256.New) {
		t.Error("Should return false for nil public key")
	}

	if ECDSAVerify(keyPair.PublicKey, []byte("test"), nil, big.NewInt(1), sha256.New) {
		t.Error("Should return false for nil r")
	}

	if ECDSAVerify(keyPair.PublicKey, []byte("test"), big.NewInt(1), nil, sha256.New) {
		t.Error("Should return false for nil s")
	}

	if ECDSAVerify(keyPair.PublicKey, []byte("test"), big.NewInt(1), big.NewInt(1), nil) {
		t.Error("Should return false for nil hash function")
	}

	_, err = ECDSASignatureToBytes(nil, big.NewInt(1))
	if err == nil {
		t.Error("Expected error for nil r")
	}

	_, err = ECDSASignatureToBytes(big.NewInt(1), nil)
	if err == nil {
		t.Error("Expected error for nil s")
	}

	_, err = ECDSAPrivateKeyFromPEM(nil)
	if err == nil {
		t.Error("Expected error for nil PEM data")
	}

	_, err = ECDSAPrivateKeyFromPEM([]byte(""))
	if err == nil {
		t.Error("Expected error for empty PEM data")
	}

	_, err = ECDSAPrivateKeyFromPEM([]byte("invalid pem"))
	if err == nil {
		t.Error("Expected error for invalid PEM data")
	}

	_, err = ECDSAPublicKeyFromPEM([]byte("invalid pem"))
	if err == nil {
		t.Error("Expected error for invalid PEM data")
	}
}
