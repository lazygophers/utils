package cryptox

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"fmt"
	"hash"
	"math/big"
	"testing"
)

// Test key generation
func TestGenerateECDSAKey(t *testing.T) {
	curves := []struct {
		name  string
		curve elliptic.Curve
	}{
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

	tests := []struct {
		name     string
		pubKey   *ecdsa.PublicKey
		data     []byte
		r        *big.Int
		s        *big.Int
		hashFunc func() hash.Hash
	}{
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
	tests := []struct {
		name string
		data []byte
	}{
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
	tests := []struct {
		curve    elliptic.Curve
		expected string
	}{
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
	tests := []struct {
		name     string
		curve    elliptic.Curve
		expected bool
	}{
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
