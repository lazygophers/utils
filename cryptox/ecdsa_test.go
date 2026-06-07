package cryptox

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"hash"
	"math/big"
	"sync"
	"testing"
)

type ecdsaCurveCase struct {
	name  string
	curve elliptic.Curve
}

type ecdsaVerifyInputCase struct {
	name     string
	pubKey   *ecdsa.PublicKey
	data     []byte
	r        *big.Int
	s        *big.Int
	hashFunc func() hash.Hash
}

type ecdsaDataCase struct {
	name string
	data []byte
}

type ecdsaCurveNameCase struct {
	curve    elliptic.Curve
	expected string
}

type ecdsaCurveValidCase struct {
	name     string
	curve    elliptic.Curve
	expected bool
}

// Test key generation
func TestGenerateECDSAKey(t *testing.T) {
	curves := []ecdsaCurveCase{
		{"P224", elliptic.P224()},
		{"P256", elliptic.P256()},
		{"P384", elliptic.P384()},
		{"P521", elliptic.P521()},
	}

	for _, tc := range curves {
		t.Run(tc.name, func(t *testing.T) {
			keyPair, err := GenerateECDSAKey(tc.curve)
			if err != nil {
				t.Fatalf("GenerateECDSAKey(%s) failed: %v", tc.name, err)
			}
			if keyPair == nil {
				t.Fatal("keyPair is nil")
			}
			if keyPair.PrivateKey == nil {
				t.Fatal("PrivateKey is nil")
			}
			if keyPair.PublicKey == nil {
				t.Fatal("PublicKey is nil")
			}
			if keyPair.PrivateKey.PublicKey.Curve != tc.curve {
				t.Errorf("curve mismatch")
			}
		})
	}
}

func TestGenerateECDSAKey_NilCurve(t *testing.T) {
	_, err := GenerateECDSAKey(nil)
	if err == nil {
		t.Error("expected error for nil curve")
	}
	if err.Error() != "curve cannot be nil" {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestGenerateECDSAP256Key(t *testing.T) {
	keyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("GenerateECDSAP256Key failed: %v", err)
	}
	if keyPair == nil || keyPair.PrivateKey == nil {
		t.Fatal("keyPair is nil")
	}
	if keyPair.PrivateKey.Curve != elliptic.P256() {
		t.Error("expected P-256 curve")
	}
}

func TestGenerateECDSAP384Key(t *testing.T) {
	keyPair, err := GenerateECDSAP384Key()
	if err != nil {
		t.Fatalf("GenerateECDSAP384Key failed: %v", err)
	}
	if keyPair == nil || keyPair.PrivateKey == nil {
		t.Fatal("keyPair is nil")
	}
	if keyPair.PrivateKey.Curve != elliptic.P384() {
		t.Error("expected P-384 curve")
	}
}

func TestGenerateECDSAP521Key(t *testing.T) {
	keyPair, err := GenerateECDSAP521Key()
	if err != nil {
		t.Fatalf("GenerateECDSAP521Key failed: %v", err)
	}
	if keyPair == nil || keyPair.PrivateKey == nil {
		t.Fatal("keyPair is nil")
	}
	if keyPair.PrivateKey.Curve != elliptic.P521() {
		t.Error("expected P-521 curve")
	}
}

// Test signing and verification
func TestECDSASign(t *testing.T) {
	keyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}

	testData := []byte("test message")

	r, s, err := ECDSASign(keyPair.PrivateKey, testData, sha256.New)
	if err != nil {
		t.Fatalf("ECDSASign failed: %v", err)
	}

	if r == nil || s == nil {
		t.Fatal("signature components are nil")
	}

	if r.Sign() <= 0 || s.Sign() <= 0 {
		t.Error("signature components should be positive")
	}
}

func TestECDSASign_NilPrivateKey(t *testing.T) {
	_, _, err := ECDSASign(nil, []byte("test"), sha256.New)
	if err == nil {
		t.Error("expected error for nil private key")
	}
}

func TestECDSASign_NilHashFunc(t *testing.T) {
	keyPair, _ := GenerateECDSAP256Key()
	_, _, err := ECDSASign(keyPair.PrivateKey, []byte("test"), nil)
	if err == nil {
		t.Error("expected error for nil hash function")
	}
}

func TestECDSASignSHA256(t *testing.T) {
	keyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}

	testData := []byte("test message for SHA256")

	r, s, err := ECDSASignSHA256(keyPair.PrivateKey, testData)
	if err != nil {
		t.Fatalf("ECDSASignSHA256 failed: %v", err)
	}

	if r == nil || s == nil {
		t.Fatal("signature components are nil")
	}
}

func TestECDSASignSHA512(t *testing.T) {
	keyPair, err := GenerateECDSAP384Key()
	if err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}

	testData := []byte("test message for SHA512")

	r, s, err := ECDSASignSHA512(keyPair.PrivateKey, testData)
	if err != nil {
		t.Fatalf("ECDSASignSHA512 failed: %v", err)
	}

	if r == nil || s == nil {
		t.Fatal("signature components are nil")
	}
}

func TestECDSAVerify(t *testing.T) {
	keyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}

	testData := []byte("test message for verification")

	r, s, err := ECDSASign(keyPair.PrivateKey, testData, sha256.New)
	if err != nil {
		t.Fatalf("ECDSASign failed: %v", err)
	}

	// Valid signature
	valid := ECDSAVerify(keyPair.PublicKey, testData, r, s, sha256.New)
	if !valid {
		t.Error("valid signature should be verified as true")
	}

	// Invalid data
	invalidData := []byte("different message")
	valid = ECDSAVerify(keyPair.PublicKey, invalidData, r, s, sha256.New)
	if valid {
		t.Error("invalid data should be verified as false")
	}

	// Invalid signature
	invalidR := new(big.Int).Add(r, big.NewInt(1))
	valid = ECDSAVerify(keyPair.PublicKey, testData, invalidR, s, sha256.New)
	if valid {
		t.Error("invalid signature should be verified as false")
	}
}

func TestECDSAVerify_NilInputs(t *testing.T) {
	keyPair, _ := GenerateECDSAP256Key()
	r := big.NewInt(123)
	s := big.NewInt(456)
	data := []byte("test")

	tests := []ecdsaVerifyInputCase{
		{"nil public key", nil, data, r, s, sha256.New},
		{"nil r", keyPair.PublicKey, data, nil, s, sha256.New},
		{"nil s", keyPair.PublicKey, data, r, nil, sha256.New},
		{"nil hash func", keyPair.PublicKey, data, r, s, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid := ECDSAVerify(tt.pubKey, tt.data, tt.r, tt.s, tt.hashFunc)
			if valid {
				t.Error("expected false for nil input")
			}
		})
	}
}

func TestECDSAVerifySHA256(t *testing.T) {
	keyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}

	testData := []byte("test message for SHA256 verification")

	r, s, err := ECDSASignSHA256(keyPair.PrivateKey, testData)
	if err != nil {
		t.Fatalf("ECDSASignSHA256 failed: %v", err)
	}

	valid := ECDSAVerifySHA256(keyPair.PublicKey, testData, r, s)
	if !valid {
		t.Error("valid SHA256 signature should be verified as true")
	}
}

func TestECDSAVerifySHA512(t *testing.T) {
	keyPair, err := GenerateECDSAP384Key()
	if err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}

	testData := []byte("test message for SHA512 verification")

	r, s, err := ECDSASignSHA512(keyPair.PrivateKey, testData)
	if err != nil {
		t.Fatalf("ECDSASignSHA512 failed: %v", err)
	}

	valid := ECDSAVerifySHA512(keyPair.PublicKey, testData, r, s)
	if !valid {
		t.Error("valid SHA512 signature should be verified as true")
	}
}

// Test PEM encoding/decoding
func TestECDSAPrivateKeyToPEM(t *testing.T) {
	keyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}

	pemData, err := ECDSAPrivateKeyToPEM(keyPair.PrivateKey)
	if err != nil {
		t.Fatalf("ECDSAPrivateKeyToPEM failed: %v", err)
	}

	if len(pemData) == 0 {
		t.Fatal("PEM data is empty")
	}

	if !containsString(string(pemData), "-----BEGIN EC PRIVATE KEY-----") {
		t.Error("PEM data should contain BEGIN EC PRIVATE KEY header")
	}

	if !containsString(string(pemData), "-----END EC PRIVATE KEY-----") {
		t.Error("PEM data should contain END EC PRIVATE KEY footer")
	}
}

func TestECDSAPrivateKeyToPEM_NilKey(t *testing.T) {
	_, err := ECDSAPrivateKeyToPEM(nil)
	if err == nil {
		t.Error("expected error for nil private key")
	}
}

func TestECDSAPrivateKeyFromPEM(t *testing.T) {
	keyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}

	pemData, err := ECDSAPrivateKeyToPEM(keyPair.PrivateKey)
	if err != nil {
		t.Fatalf("ECDSAPrivateKeyToPEM failed: %v", err)
	}

	parsedKey, err := ECDSAPrivateKeyFromPEM(pemData)
	if err != nil {
		t.Fatalf("ECDSAPrivateKeyFromPEM failed: %v", err)
	}

	if parsedKey == nil {
		t.Fatal("parsed key is nil")
	}

	if parsedKey.D.Cmp(keyPair.PrivateKey.D) != 0 {
		t.Error("private key D component mismatch")
	}
}

func TestECDSAPrivateKeyFromPEM_EmptyData(t *testing.T) {
	_, err := ECDSAPrivateKeyFromPEM([]byte{})
	if err == nil {
		t.Error("expected error for empty PEM data")
	}
}

func TestECDSAPrivateKeyFromPEM_InvalidPEM(t *testing.T) {
	_, err := ECDSAPrivateKeyFromPEM([]byte("not valid pem"))
	if err == nil {
		t.Error("expected error for invalid PEM data")
	}
}

func TestECDSAPrivateKeyFromPEM_WrongType(t *testing.T) {
	invalidPEM := `-----BEGIN RSA PRIVATE KEY-----
MIIBOgIBAAJBAKj34GkxFhD90vcNLYLInFEX6Ppy1tPf9Cnzj4p4WGeKLs1Pt8Qu
-----END RSA PRIVATE KEY-----`

	_, err := ECDSAPrivateKeyFromPEM([]byte(invalidPEM))
	if err == nil {
		t.Error("expected error for wrong PEM block type")
	}
}

func TestECDSAPublicKeyToPEM(t *testing.T) {
	keyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}

	pemData, err := ECDSAPublicKeyToPEM(keyPair.PublicKey)
	if err != nil {
		t.Fatalf("ECDSAPublicKeyToPEM failed: %v", err)
	}

	if len(pemData) == 0 {
		t.Fatal("PEM data is empty")
	}

	if !containsString(string(pemData), "-----BEGIN PUBLIC KEY-----") {
		t.Error("PEM data should contain BEGIN PUBLIC KEY header")
	}

	if !containsString(string(pemData), "-----END PUBLIC KEY-----") {
		t.Error("PEM data should contain END PUBLIC KEY footer")
	}
}

func TestECDSAPublicKeyToPEM_NilKey(t *testing.T) {
	_, err := ECDSAPublicKeyToPEM(nil)
	if err == nil {
		t.Error("expected error for nil public key")
	}
}

func TestECDSAPublicKeyFromPEM(t *testing.T) {
	keyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}

	pemData, err := ECDSAPublicKeyToPEM(keyPair.PublicKey)
	if err != nil {
		t.Fatalf("ECDSAPublicKeyToPEM failed: %v", err)
	}

	parsedKey, err := ECDSAPublicKeyFromPEM(pemData)
	if err != nil {
		t.Fatalf("ECDSAPublicKeyFromPEM failed: %v", err)
	}

	if parsedKey == nil {
		t.Fatal("parsed key is nil")
	}

	if parsedKey.X.Cmp(keyPair.PublicKey.X) != 0 || parsedKey.Y.Cmp(keyPair.PublicKey.Y) != 0 {
		t.Error("public key coordinates mismatch")
	}
}

func TestECDSAPublicKeyFromPEM_EmptyData(t *testing.T) {
	_, err := ECDSAPublicKeyFromPEM([]byte{})
	if err == nil {
		t.Error("expected error for empty PEM data")
	}
}

func TestECDSAPublicKeyFromPEM_InvalidPEM(t *testing.T) {
	_, err := ECDSAPublicKeyFromPEM([]byte("not valid pem"))
	if err == nil {
		t.Error("expected error for invalid PEM data")
	}
}

func TestECDSAPublicKeyFromPEM_WrongType(t *testing.T) {
	invalidPEM := `-----BEGIN CERTIFICATE-----
MIIBOgIBAAJBAKj34GkxFhD90vcNLYLInFEX6Ppy1tPf9Cnzj4p4WGeKLs1Pt8Qu
-----END CERTIFICATE-----`

	_, err := ECDSAPublicKeyFromPEM([]byte(invalidPEM))
	if err == nil {
		t.Error("expected error for wrong PEM block type")
	}
}

// Test signature encoding/decoding
func TestECDSASignatureToBytes(t *testing.T) {
	r := big.NewInt(12345)
	s := big.NewInt(67890)

	bytes, err := ECDSASignatureToBytes(r, s)
	if err != nil {
		t.Fatalf("ECDSASignatureToBytes failed: %v", err)
	}

	if len(bytes) == 0 {
		t.Fatal("signature bytes are empty")
	}

	// DER signature should start with 0x30 (SEQUENCE)
	if bytes[0] != 0x30 {
		t.Error("DER signature should start with SEQUENCE tag (0x30)")
	}
}

func TestECDSASignatureToBytes_NilComponents(t *testing.T) {
	_, err := ECDSASignatureToBytes(nil, big.NewInt(1))
	if err == nil {
		t.Error("expected error for nil r")
	}

	_, err = ECDSASignatureToBytes(big.NewInt(1), nil)
	if err == nil {
		t.Error("expected error for nil s")
	}
}

func TestECDSASignatureFromBytes(t *testing.T) {
	r := big.NewInt(12345)
	s := big.NewInt(67890)

	bytes, err := ECDSASignatureToBytes(r, s)
	if err != nil {
		t.Fatalf("ECDSASignatureToBytes failed: %v", err)
	}

	parsedR, parsedS, err := ECDSASignatureFromBytes(bytes)
	if err != nil {
		t.Fatalf("ECDSASignatureFromBytes failed: %v", err)
	}

	if parsedR.Cmp(r) != 0 {
		t.Errorf("r mismatch: got %v, want %v", parsedR, r)
	}

	if parsedS.Cmp(s) != 0 {
		t.Errorf("s mismatch: got %v, want %v", parsedS, s)
	}
}

func TestECDSASignatureFromBytes_ShortData(t *testing.T) {
	_, _, err := ECDSASignatureFromBytes([]byte{0x30, 0x00})
	if err == nil {
		t.Error("expected error for too short data")
	}
}

func TestECDSASignatureFromBytes_InvalidDER(t *testing.T) {
	tests := []ecdsaDataCase{
		{"missing SEQUENCE tag", []byte{0x31, 0x06, 0x02, 0x01, 0x01, 0x02, 0x01, 0x02}},
		{"invalid sequence length", []byte{0x30, 0xFF, 0x02, 0x01, 0x01}},
		{"missing r INTEGER tag", []byte{0x30, 0x06, 0x03, 0x01, 0x01, 0x02, 0x01, 0x02}},
		{"invalid r length", []byte{0x30, 0x06, 0x02, 0xFF, 0x01, 0x02, 0x01, 0x02}},
		{"missing s INTEGER tag", []byte{0x30, 0x06, 0x02, 0x01, 0x01, 0x03, 0x01, 0x02}},
		{"invalid s length", []byte{0x30, 0x06, 0x02, 0x01, 0x01, 0x02, 0xFF, 0x02}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := ECDSASignatureFromBytes(tt.data)
			if err == nil {
				t.Errorf("expected error for %s", tt.name)
			}
		})
	}
}

func TestECDSASignatureRoundTrip(t *testing.T) {
	keyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}

	testData := []byte("test message for signature round trip")

	r, s, err := ECDSASignSHA256(keyPair.PrivateKey, testData)
	if err != nil {
		t.Fatalf("ECDSASignSHA256 failed: %v", err)
	}

	// Encode to bytes
	sigBytes, err := ECDSASignatureToBytes(r, s)
	if err != nil {
		t.Fatalf("ECDSASignatureToBytes failed: %v", err)
	}

	// Decode from bytes
	parsedR, parsedS, err := ECDSASignatureFromBytes(sigBytes)
	if err != nil {
		t.Fatalf("ECDSASignatureFromBytes failed: %v", err)
	}

	// Verify with decoded signature
	valid := ECDSAVerifySHA256(keyPair.PublicKey, testData, parsedR, parsedS)
	if !valid {
		t.Error("signature verification failed after round trip")
	}
}

// Test curve utilities
func TestGetCurveName(t *testing.T) {
	tests := []ecdsaCurveNameCase{
		{elliptic.P224(), "P-224"},
		{elliptic.P256(), "P-256"},
		{elliptic.P384(), "P-384"},
		{elliptic.P521(), "P-521"},
		{nil, "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			name := GetCurveName(tt.curve)
			if name != tt.expected {
				t.Errorf("GetCurveName() = %s, want %s", name, tt.expected)
			}
		})
	}
}

func TestIsValidCurve(t *testing.T) {
	tests := []ecdsaCurveValidCase{
		{"P224", elliptic.P224(), true},
		{"P256", elliptic.P256(), true},
		{"P384", elliptic.P384(), true},
		{"P521", elliptic.P521(), true},
		{"nil", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid := IsValidCurve(tt.curve)
			if valid != tt.expected {
				t.Errorf("IsValidCurve(%s) = %v, want %v", tt.name, valid, tt.expected)
			}
		})
	}
}

// Test complete workflow with different curves
func TestECDSAWorkflowP256(t *testing.T) {
	testECDSAWorkflow(t, elliptic.P256(), "P-256")
}

func TestECDSAWorkflowP384(t *testing.T) {
	testECDSAWorkflow(t, elliptic.P384(), "P-384")
}

func TestECDSAWorkflowP521(t *testing.T) {
	testECDSAWorkflow(t, elliptic.P521(), "P-521")
}

func testECDSAWorkflow(t *testing.T, curve elliptic.Curve, curveName string) {
	// Generate key pair
	keyPair, err := GenerateECDSAKey(curve)
	if err != nil {
		t.Fatalf("GenerateECDSAKey(%s) failed: %v", curveName, err)
	}

	// Test data
	testData := []byte("test message for " + curveName)

	// Sign with SHA256
	r, s, err := ECDSASignSHA256(keyPair.PrivateKey, testData)
	if err != nil {
		t.Fatalf("ECDSASignSHA256 failed: %v", err)
	}

	// Verify signature
	valid := ECDSAVerifySHA256(keyPair.PublicKey, testData, r, s)
	if !valid {
		t.Error("signature verification failed")
	}

	// Export and import private key
	privPEM, err := ECDSAPrivateKeyToPEM(keyPair.PrivateKey)
	if err != nil {
		t.Fatalf("ECDSAPrivateKeyToPEM failed: %v", err)
	}

	importedPrivKey, err := ECDSAPrivateKeyFromPEM(privPEM)
	if err != nil {
		t.Fatalf("ECDSAPrivateKeyFromPEM failed: %v", err)
	}

	// Sign with imported private key
	r2, s2, err := ECDSASignSHA256(importedPrivKey, testData)
	if err != nil {
		t.Fatalf("ECDSASignSHA256 with imported key failed: %v", err)
	}

	// Verify with original public key
	valid = ECDSAVerifySHA256(keyPair.PublicKey, testData, r2, s2)
	if !valid {
		t.Error("signature verification with imported key failed")
	}

	// Export and import public key
	pubPEM, err := ECDSAPublicKeyToPEM(keyPair.PublicKey)
	if err != nil {
		t.Fatalf("ECDSAPublicKeyToPEM failed: %v", err)
	}

	importedPubKey, err := ECDSAPublicKeyFromPEM(pubPEM)
	if err != nil {
		t.Fatalf("ECDSAPublicKeyFromPEM failed: %v", err)
	}

	// Verify with imported public key
	valid = ECDSAVerifySHA256(importedPubKey, testData, r, s)
	if !valid {
		t.Error("signature verification with imported public key failed")
	}
}

// Test with various data sizes
func TestECDSAWithVariousDataSizes(t *testing.T) {
	keyPair, err := GenerateECDSAP256Key()
	if err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}

	dataSizes := []int{0, 1, 16, 256, 1024, 4096, 65536}

	for _, size := range dataSizes {
		t.Run(fmt.Sprintf("size_%d", size), func(t *testing.T) {
			data := make([]byte, size)
			for i := range data {
				data[i] = byte(i % 256)
			}

			r, s, err := ECDSASignSHA256(keyPair.PrivateKey, data)
			if err != nil {
				t.Fatalf("ECDSASignSHA256 failed for size %d: %v", size, err)
			}

			valid := ECDSAVerifySHA256(keyPair.PublicKey, data, r, s)
			if !valid {
				t.Errorf("signature verification failed for size %d", size)
			}
		})
	}
}

// Test signature with high bit set (requires 0x00 padding in DER)
func TestECDSASignatureWithHighBit(t *testing.T) {
	// Create a big.Int with high bit set
	r := new(big.Int)
	r.SetString("FF00000000000000000000000000000000000000000000000000000000000001", 16)
	s := new(big.Int)
	s.SetString("8000000000000000000000000000000000000000000000000000000000000002", 16)

	bytes, err := ECDSASignatureToBytes(r, s)
	if err != nil {
		t.Fatalf("ECDSASignatureToBytes failed: %v", err)
	}

	parsedR, parsedS, err := ECDSASignatureFromBytes(bytes)
	if err != nil {
		t.Fatalf("ECDSASignatureFromBytes failed: %v", err)
	}

	if parsedR.Cmp(r) != 0 {
		t.Error("r mismatch with high bit")
	}

	if parsedS.Cmp(s) != 0 {
		t.Error("s mismatch with high bit")
	}
}

// Benchmarks
func BenchmarkGenerateECDSAP256Key(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateECDSAP256Key()
	}
}

func BenchmarkGenerateECDSAP384Key(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GenerateECDSAP384Key()
	}
}

func BenchmarkECDSASignSHA256(b *testing.B) {
	keyPair, _ := GenerateECDSAP256Key()
	data := []byte("benchmark test data")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, _ = ECDSASignSHA256(keyPair.PrivateKey, data)
	}
}

func BenchmarkECDSAVerifySHA256(b *testing.B) {
	keyPair, _ := GenerateECDSAP256Key()
	data := []byte("benchmark test data")
	r, s, _ := ECDSASignSHA256(keyPair.PrivateKey, data)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = ECDSAVerifySHA256(keyPair.PublicKey, data, r, s)
	}
}

func BenchmarkECDSAPrivateKeyToPEM(b *testing.B) {
	keyPair, _ := GenerateECDSAP256Key()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = ECDSAPrivateKeyToPEM(keyPair.PrivateKey)
	}
}

func BenchmarkECDSAPrivateKeyFromPEM(b *testing.B) {
	keyPair, _ := GenerateECDSAP256Key()
	pemData, _ := ECDSAPrivateKeyToPEM(keyPair.PrivateKey)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = ECDSAPrivateKeyFromPEM(pemData)
	}
}

func BenchmarkECDSASignatureToBytes(b *testing.B) {
	keyPair, _ := GenerateECDSAP256Key()
	data := []byte("benchmark test data")
	r, s, _ := ECDSASignSHA256(keyPair.PrivateKey, data)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = ECDSASignatureToBytes(r, s)
	}
}

func BenchmarkECDSASignatureFromBytes(b *testing.B) {
	keyPair, _ := GenerateECDSAP256Key()
	data := []byte("benchmark test data")
	r, s, _ := ECDSASignSHA256(keyPair.PrivateKey, data)
	sigBytes, _ := ECDSASignatureToBytes(r, s)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, _ = ECDSASignatureFromBytes(sigBytes)
	}
}

// Helper function
func containsString(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && containsStringHelper(s, substr))
}

func containsStringHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// ============ ECDSA Sign Benchmark (10+ 方案) ============

// ============ ECDSA Verify Benchmark (10+ 方案) ============

// ============ ECDH ComputeShared Benchmark (10+ 方案) ============

// ============ ECDH ComputeSharedWithKDF Benchmark (10+ 方案) ============

// ============ 内存分配对比 ============

// ============ 并发安全性测试 ============

// ============ 功能正确性验证 ============

func TestECDSAOptimizationCorrectness(t *testing.T) {
	privateKey, _ := GenerateECDSAP256Key()
	data := []byte("test data for correctness check")

	// 原始实现
	r1, s1, _ := ECDSASign(privateKey.PrivateKey, data, sha256.New)
	valid1 := ECDSAVerify(privateKey.PublicKey, data, r1, s1, sha256.New)

	// 优化方案 1
	r2, s2, _ := ECDSASignOpt1(privateKey.PrivateKey, data)
	valid2 := ECDSAVerifyOpt6(privateKey.PublicKey, data, r2, s2)

	if !valid1 || !valid2 {
		t.Error("签名验证失败")
	}

	// ECDSA 签名是非确定性的，所以不比较签名值本身
	// 只要验证通过即可
}

func TestECDHOptimizationCorrectness(t *testing.T) {
	alicePrivate, _ := GenerateECDHP256Key()
	bobPublic, _ := GenerateECDHP256Key()

	// 原始实现
	secret1, _ := ECDHComputeShared(alicePrivate.PrivateKey, bobPublic.PublicKey)

	// 优化方案 1
	secret2, _ := ECDHComputeSharedOpt1(alicePrivate.PrivateKey, bobPublic.PublicKey)

	if len(secret1) != len(secret2) {
		t.Errorf("共享密钥长度不一致: %d vs %d", len(secret1), len(secret2))
	}

	for i := range secret1 {
		if secret1[i] != secret2[i] {
			t.Errorf("共享密钥内容不一致在位置 %d: %d vs %d", i, secret1[i], secret2[i])
		}
	}
}

func TestECDHKDFOptimizationCorrectness(t *testing.T) {
	alicePrivate, _ := GenerateECDHP256Key()
	bobPublic, _ := GenerateECDHP256Key()

	// 原始实现
	secret1, _ := ECDHComputeSharedWithKDF(alicePrivate.PrivateKey, bobPublic.PublicKey, 32, sha256.New)

	// 优化方案 7
	secret2, _ := ECDHComputeSharedWithKDFOpt7(alicePrivate.PrivateKey, bobPublic.PublicKey, 32)

	if len(secret1) != len(secret2) {
		t.Errorf("KDF 密钥长度不一致: %d vs %d", len(secret1), len(secret2))
	}

	for i := range secret1 {
		if secret1[i] != secret2[i] {
			t.Errorf("KDF 密钥内容不一致在位置 %d: %d vs %d", i, secret1[i], secret2[i])
		}
	}
}

// ========================================
// ECDSA/ECDH 性能优化方案
// ========================================

// 优化对象池
var (
	ecdsaSha256Pool = sync.Pool{
		New: func() any { return sha256.New() },
	}
	ecdsaSha512Pool = sync.Pool{
		New: func() any { return sha512.New() },
	}
)

// ============ ECDSA 优化方案 1-5：签名 ============

// ECDSASignOpt1: 对象池优化（SHA256）
func ECDSASignOpt1(privateKey *ecdsa.PrivateKey, data []byte) (r, s *big.Int, err error) {
	if privateKey == nil {
		return nil, nil, errors.New("private key cannot be nil")
	}

	h := ecdsaSha256Pool.Get().(hash.Hash)
	h.Reset()
	h.Write(data)
	hashed := h.Sum(nil)
	ecdsaSha256Pool.Put(h)

	return ecdsa.Sign(ecdsaRandReader, privateKey, hashed)
}

// ECDSASignOpt2: 对象池优化（SHA512）
func ECDSASignOpt2(privateKey *ecdsa.PrivateKey, data []byte) (r, s *big.Int, err error) {
	if privateKey == nil {
		return nil, nil, errors.New("private key cannot be nil")
	}

	h := ecdsaSha512Pool.Get().(hash.Hash)
	h.Reset()
	h.Write(data)
	hashed := h.Sum(nil)
	ecdsaSha512Pool.Put(h)

	return ecdsa.Sign(ecdsaRandReader, privateKey, hashed)
}

// ECDSASignOpt3: 预分配哈希缓冲区
func ECDSASignOpt3(privateKey *ecdsa.PrivateKey, data []byte, hashFunc func() hash.Hash) (r, s *big.Int, err error) {
	if privateKey == nil {
		return nil, nil, errors.New("private key cannot be nil")
	}
	if hashFunc == nil {
		return nil, nil, errors.New("hash function cannot be nil")
	}

	h := hashFunc()
	size := h.Size()
	buf := make([]byte, 0, size)

	h.Write(data)
	hashed := h.Sum(buf)

	return ecdsa.Sign(ecdsaRandReader, privateKey, hashed)
}

// ECDSASignOpt4: 减少函数调用（内联哈希）
func ECDSASignOpt4(privateKey *ecdsa.PrivateKey, data []byte) (r, s *big.Int, err error) {
	if privateKey == nil {
		return nil, nil, errors.New("private key cannot be nil")
	}

	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)

	return ecdsa.Sign(ecdsaRandReader, privateKey, hashed)
}

// ECDSASignOpt5: 组合优化（对象池 + 预分配）
func ECDSASignOpt5(privateKey *ecdsa.PrivateKey, data []byte) (r, s *big.Int, err error) {
	if privateKey == nil {
		return nil, nil, errors.New("private key cannot be nil")
	}

	h := ecdsaSha256Pool.Get().(hash.Hash)
	h.Reset()

	buf := make([]byte, 0, 32)
	h.Write(data)
	hashed := h.Sum(buf)

	ecdsaSha256Pool.Put(h)

	return ecdsa.Sign(ecdsaRandReader, privateKey, hashed)
}

// ============ ECDSA 优化方案 6-10：验证 ============

// ECDSAVerifyOpt6: 对象池优化（SHA256）
func ECDSAVerifyOpt6(publicKey *ecdsa.PublicKey, data []byte, r, s *big.Int) bool {
	if publicKey == nil || r == nil || s == nil {
		return false
	}

	h := ecdsaSha256Pool.Get().(hash.Hash)
	h.Reset()
	h.Write(data)
	hashed := h.Sum(nil)
	ecdsaSha256Pool.Put(h)

	return ecdsa.Verify(publicKey, hashed, r, s)
}

// ECDSAVerifyOpt7: 对象池优化（SHA512）
func ECDSAVerifyOpt7(publicKey *ecdsa.PublicKey, data []byte, r, s *big.Int) bool {
	if publicKey == nil || r == nil || s == nil {
		return false
	}

	h := ecdsaSha512Pool.Get().(hash.Hash)
	h.Reset()
	h.Write(data)
	hashed := h.Sum(nil)
	ecdsaSha512Pool.Put(h)

	return ecdsa.Verify(publicKey, hashed, r, s)
}

// ECDSAVerifyOpt8: 预分配哈希缓冲区
func ECDSAVerifyOpt8(publicKey *ecdsa.PublicKey, data []byte, r, s *big.Int, hashFunc func() hash.Hash) bool {
	if publicKey == nil || r == nil || s == nil || hashFunc == nil {
		return false
	}

	h := hashFunc()
	size := h.Size()
	buf := make([]byte, 0, size)

	h.Write(data)
	hashed := h.Sum(buf)

	return ecdsa.Verify(publicKey, hashed, r, s)
}

// ECDSAVerifyOpt9: 减少函数调用（内联哈希）
func ECDSAVerifyOpt9(publicKey *ecdsa.PublicKey, data []byte, r, s *big.Int) bool {
	if publicKey == nil || r == nil || s == nil {
		return false
	}

	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)

	return ecdsa.Verify(publicKey, hashed, r, s)
}

// ECDSAVerifyOpt10: 合并 nil 检查
func ECDSAVerifyOpt10(publicKey *ecdsa.PublicKey, data []byte, r, s *big.Int) bool {
	if publicKey == nil || r == nil || s == nil {
		return false
	}

	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)

	return ecdsa.Verify(publicKey, hashed, r, s)
}

// ============ ECDH 优化方案 1-5：ComputeShared ============

// ECDHComputeSharedOpt1: 缓存曲线参数
func ECDHComputeSharedOpt1(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) ([]byte, error) {
	if privateKey == nil || publicKey == nil {
		return nil, errors.New("keys cannot be nil")
	}

	curve := privateKey.Curve

	if !curve.IsOnCurve(publicKey.X, publicKey.Y) {
		return nil, errors.New("public key is not on the curve")
	}

	if curve != publicKey.Curve {
		return nil, errors.New("curve mismatch between private and public keys")
	}

	x, _ := curve.ScalarMult(publicKey.X, publicKey.Y, privateKey.D.Bytes())
	return x.Bytes(), nil
}

// ECDHComputeSharedOpt2: 减少错误检查分支
func ECDHComputeSharedOpt2(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) ([]byte, error) {
	if privateKey == nil || publicKey == nil {
		return nil, errors.New("keys cannot be nil")
	}

	if !privateKey.Curve.IsOnCurve(publicKey.X, publicKey.Y) {
		return nil, errors.New("public key is not on the curve")
	}

	if privateKey.Curve != publicKey.Curve {
		return nil, errors.New("curve mismatch between private and public keys")
	}

	x, _ := privateKey.Curve.ScalarMult(publicKey.X, publicKey.Y, privateKey.D.Bytes())
	return x.Bytes(), nil
}

// ECDHComputeSharedOpt3: 预分配结果切片
func ECDHComputeSharedOpt3(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) ([]byte, error) {
	if privateKey == nil || publicKey == nil {
		return nil, errors.New("keys cannot be nil")
	}

	curve := privateKey.Curve

	if !curve.IsOnCurve(publicKey.X, publicKey.Y) {
		return nil, errors.New("public key is not on the curve")
	}

	if curve != publicKey.Curve {
		return nil, errors.New("curve mismatch between private and public keys")
	}

	x, _ := curve.ScalarMult(publicKey.X, publicKey.Y, privateKey.D.Bytes())
	byteLen := (curve.Params().BitSize + 7) / 8

	result := make([]byte, byteLen)
	copy(result, x.Bytes())

	return result, nil
}

// ECDHComputeSharedOpt4: 组合优化
func ECDHComputeSharedOpt4(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) ([]byte, error) {
	if privateKey == nil || publicKey == nil {
		return nil, errors.New("keys cannot be nil")
	}

	curve := privateKey.Curve

	if !curve.IsOnCurve(publicKey.X, publicKey.Y) {
		return nil, errors.New("public key is not on the curve")
	}

	if curve != publicKey.Curve {
		return nil, errors.New("curve mismatch between private and public keys")
	}

	x, _ := curve.ScalarMult(publicKey.X, publicKey.Y, privateKey.D.Bytes())
	return x.Bytes(), nil
}

// ECDHComputeSharedOpt5: 内联优化
func ECDHComputeSharedOpt5(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) ([]byte, error) {
	if privateKey == nil || publicKey == nil {
		return nil, errors.New("keys cannot be nil")
	}

	curve := privateKey.Curve

	if !curve.IsOnCurve(publicKey.X, publicKey.Y) || curve != publicKey.Curve {
		return nil, errors.New("invalid public key or curve mismatch")
	}

	x, _ := curve.ScalarMult(publicKey.X, publicKey.Y, privateKey.D.Bytes())
	return x.Bytes(), nil
}

// ============ ECDH 优化方案 6-10：ComputeSharedWithKDF ============

// ECDHComputeSharedWithKDFOpt6: 预分配迭代缓冲区
func ECDHComputeSharedWithKDFOpt6(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, keyLength int, kdf func() hash.Hash) ([]byte, error) {
	if keyLength <= 0 {
		return nil, errors.New("key length must be greater than 0")
	}
	if kdf == nil {
		return nil, errors.New("KDF function cannot be nil")
	}

	sharedSecret, err := ECDHComputeShared(privateKey, publicKey)
	if err != nil {
		return nil, err
	}

	h := kdf()
	h.Write(sharedSecret)
	derivedKey := h.Sum(nil)

	if len(derivedKey) >= keyLength {
		return derivedKey[:keyLength], nil
	}

	hashSize := len(derivedKey)
	iterations := (keyLength + hashSize - 1) / hashSize

	result := make([]byte, 0, keyLength)
	counter := uint32(0)
	counterBuf := make([]byte, 4)

	for i := 0; i < iterations; i++ {
		h.Reset()
		h.Write(sharedSecret)

		counterBuf[0] = byte(counter >> 24)
		counterBuf[1] = byte(counter >> 16)
		counterBuf[2] = byte(counter >> 8)
		counterBuf[3] = byte(counter)

		h.Write(counterBuf)
		block := h.Sum(nil)
		result = append(result, block...)
		counter++
	}

	return result[:keyLength], nil
}

// ECDHComputeSharedWithKDFOpt7: 使用对象池
func ECDHComputeSharedWithKDFOpt7(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, keyLength int) ([]byte, error) {
	if keyLength <= 0 {
		return nil, errors.New("key length must be greater than 0")
	}

	sharedSecret, err := ECDHComputeShared(privateKey, publicKey)
	if err != nil {
		return nil, err
	}

	h := ecdsaSha256Pool.Get().(hash.Hash)
	h.Reset()
	h.Write(sharedSecret)
	derivedKey := h.Sum(nil)
	ecdsaSha256Pool.Put(h)

	if len(derivedKey) >= keyLength {
		return derivedKey[:keyLength], nil
	}

	hashSize := len(derivedKey)
	iterations := (keyLength + hashSize - 1) / hashSize

	result := make([]byte, 0, keyLength)
	counter := uint32(0)
	counterBuf := make([]byte, 4)

	for i := 0; i < iterations; i++ {
		h := ecdsaSha256Pool.Get().(hash.Hash)
		h.Reset()
		h.Write(sharedSecret)

		counterBuf[0] = byte(counter >> 24)
		counterBuf[1] = byte(counter >> 16)
		counterBuf[2] = byte(counter >> 8)
		counterBuf[3] = byte(counter)

		h.Write(counterBuf)
		block := h.Sum(nil)
		ecdsaSha256Pool.Put(h)
		result = append(result, block...)
		counter++
	}

	return result[:keyLength], nil
}

// ECDHComputeSharedWithKDFOpt8: 优化循环内分配
func ECDHComputeSharedWithKDFOpt8(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, keyLength int, kdf func() hash.Hash) ([]byte, error) {
	if keyLength <= 0 {
		return nil, errors.New("key length must be greater than 0")
	}
	if kdf == nil {
		return nil, errors.New("KDF function cannot be nil")
	}

	sharedSecret, err := ECDHComputeShared(privateKey, publicKey)
	if err != nil {
		return nil, err
	}

	h := kdf()
	hashSize := h.Size()

	h.Reset()
	h.Write(sharedSecret)
	derivedKey := h.Sum(nil)

	if len(derivedKey) >= keyLength {
		return derivedKey[:keyLength], nil
	}

	iterations := (keyLength + hashSize - 1) / hashSize
	result := make([]byte, 0, keyLength)
	counter := uint32(0)
	counterBuf := make([]byte, 4)

	for i := 0; i < iterations; i++ {
		h.Reset()
		h.Write(sharedSecret)

		counterBuf[0] = byte(counter >> 24)
		counterBuf[1] = byte(counter >> 16)
		counterBuf[2] = byte(counter >> 8)
		counterBuf[3] = byte(counter)

		h.Write(counterBuf)
		block := h.Sum(nil)
		result = append(result, block...)
		counter++
	}

	return result[:keyLength], nil
}

// ECDHComputeSharedWithKDFOpt9: 组合优化（预分配 + 对象池）
func ECDHComputeSharedWithKDFOpt9(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, keyLength int) ([]byte, error) {
	if keyLength <= 0 {
		return nil, errors.New("key length must be greater than 0")
	}

	sharedSecret, err := ECDHComputeShared(privateKey, publicKey)
	if err != nil {
		return nil, err
	}

	h := ecdsaSha256Pool.Get().(hash.Hash)
	hashSize := h.Size()

	h.Reset()
	h.Write(sharedSecret)
	derivedKey := h.Sum(nil)
	ecdsaSha256Pool.Put(h)

	if len(derivedKey) >= keyLength {
		return derivedKey[:keyLength], nil
	}

	iterations := (keyLength + hashSize - 1) / hashSize
	result := make([]byte, 0, keyLength)
	counter := uint32(0)
	counterBuf := make([]byte, 4)

	for i := 0; i < iterations; i++ {
		h := ecdsaSha256Pool.Get().(hash.Hash)
		h.Reset()
		h.Write(sharedSecret)

		counterBuf[0] = byte(counter >> 24)
		counterBuf[1] = byte(counter >> 16)
		counterBuf[2] = byte(counter >> 8)
		counterBuf[3] = byte(counter)

		h.Write(counterBuf)
		block := h.Sum(nil)
		ecdsaSha256Pool.Put(h)
		result = append(result, block...)
		counter++
	}

	return result[:keyLength], nil
}

// ECDHComputeSharedWithKDFOpt10: 内联所有优化
func ECDHComputeSharedWithKDFOpt10(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, keyLength int) ([]byte, error) {
	if keyLength <= 0 {
		return nil, errors.New("key length must be greater than 0")
	}

	sharedSecret, err := ECDHComputeSharedOpt1(privateKey, publicKey)
	if err != nil {
		return nil, err
	}

	h := ecdsaSha256Pool.Get().(hash.Hash)
	hashSize := h.Size()

	h.Reset()
	h.Write(sharedSecret)
	derivedKey := h.Sum(nil)
	ecdsaSha256Pool.Put(h)

	if len(derivedKey) >= keyLength {
		return derivedKey[:keyLength], nil
	}

	iterations := (keyLength + hashSize - 1) / hashSize
	result := make([]byte, 0, keyLength)
	counter := uint32(0)
	counterBuf := make([]byte, 4)

	for i := 0; i < iterations; i++ {
		h := ecdsaSha256Pool.Get().(hash.Hash)
		h.Reset()
		h.Write(sharedSecret)

		counterBuf[0] = byte(counter >> 24)
		counterBuf[1] = byte(counter >> 16)
		counterBuf[2] = byte(counter >> 8)
		counterBuf[3] = byte(counter)

		h.Write(counterBuf)
		block := h.Sum(nil)
		ecdsaSha256Pool.Put(h)
		result = append(result, block...)
		counter++
	}

	return result[:keyLength], nil
}
