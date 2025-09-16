package cryptox

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

// Test data
const (
	testMessage = "Hello, RSA encryption and digital signature!"
	testKeySize = 2048
)

// Sample RSA keys for testing (2048-bit)
const (
	testPrivateKeyPEM = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC/9n7w8rHQm/wM
L8XnHt8lqz5fRzp6YkrYrZPNDfgKDGwjQGkjQ3UtGS6j7j7yO3vNTgv5BRJ8DfmO
bSqL3G7hKvYnF2cF8zB5gF8F8oG5cV7xF2YyR3R7nD8Q5dHtFzq3R5nNtF5Q5fQ5
fG3K8z3F5dF3Q8z8G5F2Q5yG3dF8Q5zD3F5dQ8F3zG8F5Q3zF8Q5zF3G8Q5zF3Q8
F5zG3Q8F5zF3G8Q5zF3Q8F5zG3Q8F5zF3G8Q5zF3Q8F5zG3Q8F5zF3G8Q5zF3Q8F
5zG3Q8F5zF3G8Q5zF3Q8F5zG3Q8F5zF3G8Q5zF3Q8F5zG3Q8F5zF3G8Q5zF3Q8Fy
AgMBAAECggEAArQHjk8WJyVu4V4U5V8JvHx5D8F5H8F5D8F5H8F5D8F5H8F5D8F5
H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5
H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5
H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5
H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5
H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5
QKBgQDqGqF8dF8Q5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5
H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5
wKBgQDSgH8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5
H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5
AoIBAQDqF8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5H8F5D8F5
-----END PRIVATE KEY-----`

	testPublicKeyPEM = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAv/Z+8PKx0Jv8DC/F5x7f
Jas+X0c6emJK2K2TzQ34CgxsI0BpI0N1LRkuo+4+8jt7zU4L+QUSfA35jm0qi9xu
4Sr2JxdnBfMweYBfBfKBuXFe8RdmMkd0e5w/EOXR7Rc6t0eZzbReUOX0OXxtyvM9
xeXRd0PM/BuRdkOcht3RfEOcw9xeXUPBd8xvBeUN8xfEOcxdxvEOcxd0PBeXBt0P
BeXBdxvEOcxd0PBeXBt0PBeXBdxvEOcxd0PBeXBt0PBeXBdxvEOcxd0PBeXBt0PB
eXBdxvEOcxd0PBeXBt0PBeXBdxvEOcxd0PBeXBt0PBeXBdxvEOcxd0PBeXBdxvEO
cxd0PAcgFwIDAQAB
-----END PUBLIC KEY-----`
)

func TestGenerateRSAKeyPair(t *testing.T) {
	keyPair, err := GenerateRSAKeyPair(testKeySize)
	if err != nil {
		t.Fatalf("Failed to generate RSA key pair: %v", err)
	}

	if keyPair.PrivateKey == nil {
		t.Error("Private key is nil")
	}

	if keyPair.PublicKey == nil {
		t.Error("Public key is nil")
	}

	// Verify key size
	keySize := GetRSAKeySize(keyPair.PublicKey)
	if keySize != testKeySize {
		t.Errorf("Expected key size %d, got %d", testKeySize, keySize)
	}
}

func TestGenerateRSAKeyPairInvalidSize(t *testing.T) {
	_, err := GenerateRSAKeyPair(512) // Too small
	if err == nil {
		t.Error("Expected error for small key size")
	}

	expectedMsg := "RSA key size must be at least 1024 bits"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestRSAKeyPairToPEM(t *testing.T) {
	keyPair, err := GenerateRSAKeyPair(testKeySize)
	if err != nil {
		t.Fatalf("Failed to generate RSA key pair: %v", err)
	}

	// Test private key to PEM
	privatePEM, err := keyPair.PrivateKeyToPEM()
	if err != nil {
		t.Fatalf("Failed to convert private key to PEM: %v", err)
	}

	if !bytes.Contains(privatePEM, []byte("-----BEGIN PRIVATE KEY-----")) {
		t.Error("Private key PEM format is invalid")
	}

	// Test public key to PEM
	publicPEM, err := keyPair.PublicKeyToPEM()
	if err != nil {
		t.Fatalf("Failed to convert public key to PEM: %v", err)
	}

	if !bytes.Contains(publicPEM, []byte("-----BEGIN PUBLIC KEY-----")) {
		t.Error("Public key PEM format is invalid")
	}
}

func TestRSAKeyPairToPEMNilKeys(t *testing.T) {
	keyPair := &RSAKeyPair{}

	// Test nil private key
	_, err := keyPair.PrivateKeyToPEM()
	if err == nil {
		t.Error("Expected error for nil private key")
	}

	// Test nil public key
	_, err = keyPair.PublicKeyToPEM()
	if err == nil {
		t.Error("Expected error for nil public key")
	}
}

func TestPrivateKeyFromPEM(t *testing.T) {
	keyPair, err := GenerateRSAKeyPair(testKeySize)
	if err != nil {
		t.Fatalf("Failed to generate RSA key pair: %v", err)
	}

	// Convert to PEM and back
	privatePEM, err := keyPair.PrivateKeyToPEM()
	if err != nil {
		t.Fatalf("Failed to convert private key to PEM: %v", err)
	}

	loadedPrivateKey, err := PrivateKeyFromPEM(privatePEM)
	if err != nil {
		t.Fatalf("Failed to load private key from PEM: %v", err)
	}

	// Compare key sizes
	if keyPair.PrivateKey.Size() != loadedPrivateKey.Size() {
		t.Error("Loaded private key size doesn't match original")
	}
}

func TestPublicKeyFromPEM(t *testing.T) {
	keyPair, err := GenerateRSAKeyPair(testKeySize)
	if err != nil {
		t.Fatalf("Failed to generate RSA key pair: %v", err)
	}

	// Convert to PEM and back
	publicPEM, err := keyPair.PublicKeyToPEM()
	if err != nil {
		t.Fatalf("Failed to convert public key to PEM: %v", err)
	}

	loadedPublicKey, err := PublicKeyFromPEM(publicPEM)
	if err != nil {
		t.Fatalf("Failed to load public key from PEM: %v", err)
	}

	// Compare key sizes
	if keyPair.PublicKey.Size() != loadedPublicKey.Size() {
		t.Error("Loaded public key size doesn't match original")
	}
}

func TestPEMInvalidFormat(t *testing.T) {
	// Test invalid PEM data
	invalidPEM := []byte("invalid PEM data")

	_, err := PrivateKeyFromPEM(invalidPEM)
	if err == nil {
		t.Error("Expected error for invalid private key PEM")
	}

	_, err = PublicKeyFromPEM(invalidPEM)
	if err == nil {
		t.Error("Expected error for invalid public key PEM")
	}

	// Test wrong PEM block type
	wrongTypePEM := []byte(`-----BEGIN CERTIFICATE-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA
-----END CERTIFICATE-----`)

	_, err = PrivateKeyFromPEM(wrongTypePEM)
	if err == nil {
		t.Error("Expected error for wrong private key PEM type")
	}

	_, err = PublicKeyFromPEM(wrongTypePEM)
	if err == nil {
		t.Error("Expected error for wrong public key PEM type")
	}
}

func TestRSAEncryptDecryptOAEP(t *testing.T) {
	keyPair, err := GenerateRSAKeyPair(testKeySize)
	if err != nil {
		t.Fatalf("Failed to generate RSA key pair: %v", err)
	}

	message := []byte(testMessage)

	// Encrypt
	ciphertext, err := RSAEncryptOAEP(keyPair.PublicKey, message)
	if err != nil {
		t.Fatalf("Failed to encrypt with OAEP: %v", err)
	}

	// Decrypt
	decrypted, err := RSADecryptOAEP(keyPair.PrivateKey, ciphertext)
	if err != nil {
		t.Fatalf("Failed to decrypt with OAEP: %v", err)
	}

	// Verify
	if !bytes.Equal(message, decrypted) {
		t.Error("Decrypted message doesn't match original")
	}
}

func TestRSAEncryptDecryptPKCS1v15(t *testing.T) {
	keyPair, err := GenerateRSAKeyPair(testKeySize)
	if err != nil {
		t.Fatalf("Failed to generate RSA key pair: %v", err)
	}

	message := []byte(testMessage)

	// Encrypt
	ciphertext, err := RSAEncryptPKCS1v15(keyPair.PublicKey, message)
	if err != nil {
		t.Fatalf("Failed to encrypt with PKCS1v15: %v", err)
	}

	// Decrypt
	decrypted, err := RSADecryptPKCS1v15(keyPair.PrivateKey, ciphertext)
	if err != nil {
		t.Fatalf("Failed to decrypt with PKCS1v15: %v", err)
	}

	// Verify
	if !bytes.Equal(message, decrypted) {
		t.Error("Decrypted message doesn't match original")
	}
}

func TestRSAEncryptDecryptNilKeys(t *testing.T) {
	message := []byte(testMessage)

	// Test nil public key
	_, err := RSAEncryptOAEP(nil, message)
	if err == nil {
		t.Error("Expected error for nil public key in OAEP encrypt")
	}

	_, err = RSAEncryptPKCS1v15(nil, message)
	if err == nil {
		t.Error("Expected error for nil public key in PKCS1v15 encrypt")
	}

	// Test nil private key
	_, err = RSADecryptOAEP(nil, message)
	if err == nil {
		t.Error("Expected error for nil private key in OAEP decrypt")
	}

	_, err = RSADecryptPKCS1v15(nil, message)
	if err == nil {
		t.Error("Expected error for nil private key in PKCS1v15 decrypt")
	}
}

func TestRSASignVerifyPSS(t *testing.T) {
	keyPair, err := GenerateRSAKeyPair(testKeySize)
	if err != nil {
		t.Fatalf("Failed to generate RSA key pair: %v", err)
	}

	message := []byte(testMessage)

	// Sign
	signature, err := RSASignPSS(keyPair.PrivateKey, message)
	if err != nil {
		t.Fatalf("Failed to sign with PSS: %v", err)
	}

	// Verify
	err = RSAVerifyPSS(keyPair.PublicKey, message, signature)
	if err != nil {
		t.Fatalf("Failed to verify PSS signature: %v", err)
	}

	// Test verification with wrong message
	wrongMessage := []byte("Wrong message")
	err = RSAVerifyPSS(keyPair.PublicKey, wrongMessage, signature)
	if err == nil {
		t.Error("PSS signature verification should fail with wrong message")
	}
}

func TestRSASignVerifyPKCS1v15(t *testing.T) {
	keyPair, err := GenerateRSAKeyPair(testKeySize)
	if err != nil {
		t.Fatalf("Failed to generate RSA key pair: %v", err)
	}

	message := []byte(testMessage)

	// Sign
	signature, err := RSASignPKCS1v15(keyPair.PrivateKey, message)
	if err != nil {
		t.Fatalf("Failed to sign with PKCS1v15: %v", err)
	}

	// Verify
	err = RSAVerifyPKCS1v15(keyPair.PublicKey, message, signature)
	if err != nil {
		t.Fatalf("Failed to verify PKCS1v15 signature: %v", err)
	}

	// Test verification with wrong message
	wrongMessage := []byte("Wrong message")
	err = RSAVerifyPKCS1v15(keyPair.PublicKey, wrongMessage, signature)
	if err == nil {
		t.Error("PKCS1v15 signature verification should fail with wrong message")
	}
}

func TestRSASignVerifyNilKeys(t *testing.T) {
	message := []byte(testMessage)
	signature := []byte("dummy signature")

	// Test nil private key for signing
	_, err := RSASignPSS(nil, message)
	if err == nil {
		t.Error("Expected error for nil private key in PSS sign")
	}

	_, err = RSASignPKCS1v15(nil, message)
	if err == nil {
		t.Error("Expected error for nil private key in PKCS1v15 sign")
	}

	// Test nil public key for verification
	err = RSAVerifyPSS(nil, message, signature)
	if err == nil {
		t.Error("Expected error for nil public key in PSS verify")
	}

	err = RSAVerifyPKCS1v15(nil, message, signature)
	if err == nil {
		t.Error("Expected error for nil public key in PKCS1v15 verify")
	}
}

func TestGetRSAKeySize(t *testing.T) {
	keyPair, err := GenerateRSAKeyPair(testKeySize)
	if err != nil {
		t.Fatalf("Failed to generate RSA key pair: %v", err)
	}

	keySize := GetRSAKeySize(keyPair.PublicKey)
	if keySize != testKeySize {
		t.Errorf("Expected key size %d, got %d", testKeySize, keySize)
	}

	// Test nil key
	nilKeySize := GetRSAKeySize(nil)
	if nilKeySize != 0 {
		t.Errorf("Expected 0 for nil key, got %d", nilKeySize)
	}
}

func TestRSAMaxMessageLength(t *testing.T) {
	keyPair, err := GenerateRSAKeyPair(testKeySize)
	if err != nil {
		t.Fatalf("Failed to generate RSA key pair: %v", err)
	}

	// Test OAEP
	maxLenOAEP, err := RSAMaxMessageLength(keyPair.PublicKey, "OAEP")
	if err != nil {
		t.Errorf("Failed to get max message length for OAEP: %v", err)
	}

	expectedOAEP := keyPair.PublicKey.Size() - 2*32 - 2 // SHA256 hash length is 32
	if maxLenOAEP != expectedOAEP {
		t.Errorf("Expected OAEP max length %d, got %d", expectedOAEP, maxLenOAEP)
	}

	// Test PKCS1v15
	maxLenPKCS1v15, err := RSAMaxMessageLength(keyPair.PublicKey, "PKCS1v15")
	if err != nil {
		t.Errorf("Failed to get max message length for PKCS1v15: %v", err)
	}

	expectedPKCS1v15 := keyPair.PublicKey.Size() - 11
	if maxLenPKCS1v15 != expectedPKCS1v15 {
		t.Errorf("Expected PKCS1v15 max length %d, got %d", expectedPKCS1v15, maxLenPKCS1v15)
	}

	// Test unsupported padding
	_, err = RSAMaxMessageLength(keyPair.PublicKey, "INVALID")
	if err == nil {
		t.Error("Expected error for unsupported padding")
	}

	// Test nil key
	_, err = RSAMaxMessageLength(nil, "OAEP")
	if err == nil {
		t.Error("Expected error for nil public key")
	}
}

func TestRSALargeMessage(t *testing.T) {
	keyPair, err := GenerateRSAKeyPair(testKeySize)
	if err != nil {
		t.Fatalf("Failed to generate RSA key pair: %v", err)
	}

	// Create a message that's too large
	maxLen, _ := RSAMaxMessageLength(keyPair.PublicKey, "OAEP")
	largeMessage := make([]byte, maxLen+1)

	// Should fail to encrypt
	_, err = RSAEncryptOAEP(keyPair.PublicKey, largeMessage)
	if err == nil {
		t.Error("Expected error for message too large")
	}
}

func TestRSAEdgeCases(t *testing.T) {
	keyPair, err := GenerateRSAKeyPair(testKeySize)
	if err != nil {
		t.Fatalf("Failed to generate RSA key pair: %v", err)
	}

	// Test empty message
	emptyMessage := []byte("")

	ciphertext, err := RSAEncryptOAEP(keyPair.PublicKey, emptyMessage)
	if err != nil {
		t.Errorf("Failed to encrypt empty message: %v", err)
	}

	decrypted, err := RSADecryptOAEP(keyPair.PrivateKey, ciphertext)
	if err != nil {
		t.Errorf("Failed to decrypt empty message: %v", err)
	}

	if !bytes.Equal(emptyMessage, decrypted) {
		t.Error("Decrypted empty message doesn't match")
	}

	// Test single byte message
	singleByte := []byte("A")

	ciphertext, err = RSAEncryptPKCS1v15(keyPair.PublicKey, singleByte)
	if err != nil {
		t.Errorf("Failed to encrypt single byte: %v", err)
	}

	decrypted, err = RSADecryptPKCS1v15(keyPair.PrivateKey, ciphertext)
	if err != nil {
		t.Errorf("Failed to decrypt single byte: %v", err)
	}

	if !bytes.Equal(singleByte, decrypted) {
		t.Error("Decrypted single byte doesn't match")
	}
}

func TestRSADifferentKeySizes(t *testing.T) {
	keySizes := []int{1024, 2048, 3072, 4096}

	for _, keySize := range keySizes {
		t.Run(fmt.Sprintf("KeySize_%d", keySize), func(t *testing.T) {
			keyPair, err := GenerateRSAKeyPair(keySize)
			if err != nil {
				t.Fatalf("Failed to generate %d-bit RSA key pair: %v", keySize, err)
			}

			actualSize := GetRSAKeySize(keyPair.PublicKey)
			if actualSize != keySize {
				t.Errorf("Expected key size %d, got %d", keySize, actualSize)
			}

			// Test encryption/decryption with this key size
			message := []byte(fmt.Sprintf("Test message for %d-bit key", keySize))

			ciphertext, err := RSAEncryptOAEP(keyPair.PublicKey, message)
			if err != nil {
				t.Errorf("Failed to encrypt with %d-bit key: %v", keySize, err)
			}

			decrypted, err := RSADecryptOAEP(keyPair.PrivateKey, ciphertext)
			if err != nil {
				t.Errorf("Failed to decrypt with %d-bit key: %v", keySize, err)
			}

			if !bytes.Equal(message, decrypted) {
				t.Errorf("Decryption failed for %d-bit key", keySize)
			}
		})
	}
}

// ==== MERGED FROM rsa_100_coverage_test.go ====

// FailingRSAReader implements io.Reader but always fails
type FailingRSAReader struct{}

func (fr FailingRSAReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("simulated random reader failure")
}

// TestRSA_100PercentCoverage triggers all error paths using dependency injection
func TestRSA_100PercentCoverage(t *testing.T) {
	// Save original functions
	originalRSARandReader := rsaRandReader

	// Restore original functions after test
	defer func() {
		rsaRandReader = originalRSARandReader
	}()

	message := []byte("test message for RSA")

	// Test 1: Trigger rand.Reader failure in key generation
	rsaRandReader = FailingRSAReader{}

	_, err := GenerateRSAKeyPair(2048)
	if err == nil {
		t.Error("Expected rand.Reader error in GenerateRSAKeyPair")
	}

	// Restore rand.Reader for key generation in next tests
	rsaRandReader = originalRSARandReader

	// Generate a valid key pair for encryption/decryption tests
	keyPair, err := GenerateRSAKeyPair(2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA key pair: %v", err)
	}

	// Test 2: Trigger rand.Reader failure in encryption
	rsaRandReader = FailingRSAReader{}

	_, err = RSAEncryptOAEP(keyPair.PublicKey, message)
	if err == nil {
		t.Error("Expected rand.Reader error in RSAEncryptOAEP")
	}

	_, err = RSAEncryptPKCS1v15(keyPair.PublicKey, message)
	if err == nil {
		t.Error("Expected rand.Reader error in RSAEncryptPKCS1v15")
	}

	// Test 3: Trigger rand.Reader failure in signing
	_, err = RSASignPSS(keyPair.PrivateKey, message)
	if err == nil {
		t.Error("Expected rand.Reader error in RSASignPSS")
	}

	_, err = RSASignPKCS1v15(keyPair.PrivateKey, message)
	if err == nil {
		t.Error("Expected rand.Reader error in RSASignPKCS1v15")
	}

	// Restore rand.Reader for the rest of tests
	rsaRandReader = originalRSARandReader
}

// TestRSAEncryptionErrorPaths tests various encryption error conditions
func TestRSAEncryptionErrorPaths(t *testing.T) {
	keyPair, err := GenerateRSAKeyPair(2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA key pair: %v", err)
	}

	// Test message too long for OAEP
	maxOAEPLen, _ := RSAMaxMessageLength(keyPair.PublicKey, "OAEP")
	tooLongMessage := make([]byte, maxOAEPLen+1)

	_, err = RSAEncryptOAEP(keyPair.PublicKey, tooLongMessage)
	if err == nil {
		t.Error("Expected error for message too long in OAEP")
	}

	// Test message too long for PKCS1v15
	maxPKCS1v15Len, _ := RSAMaxMessageLength(keyPair.PublicKey, "PKCS1v15")
	tooLongMessage = make([]byte, maxPKCS1v15Len+1)

	_, err = RSAEncryptPKCS1v15(keyPair.PublicKey, tooLongMessage)
	if err == nil {
		t.Error("Expected error for message too long in PKCS1v15")
	}
}

// TestRSADecryptionErrorPaths tests various decryption error conditions
func TestRSADecryptionErrorPaths(t *testing.T) {
	keyPair, err := GenerateRSAKeyPair(2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA key pair: %v", err)
	}

	// Test decryption with invalid ciphertext
	invalidCiphertext := []byte("invalid ciphertext")

	_, err = RSADecryptOAEP(keyPair.PrivateKey, invalidCiphertext)
	if err == nil {
		t.Error("Expected error for invalid ciphertext in OAEP decrypt")
	}

	_, err = RSADecryptPKCS1v15(keyPair.PrivateKey, invalidCiphertext)
	if err == nil {
		t.Error("Expected error for invalid ciphertext in PKCS1v15 decrypt")
	}

	// Test decryption with wrong key (different key pair)
	otherKeyPair, _ := GenerateRSAKeyPair(2048)
	message := []byte("test message")
	ciphertext, _ := RSAEncryptOAEP(keyPair.PublicKey, message)

	_, err = RSADecryptOAEP(otherKeyPair.PrivateKey, ciphertext)
	if err == nil {
		t.Error("Expected error for wrong private key in OAEP decrypt")
	}

	ciphertext, _ = RSAEncryptPKCS1v15(keyPair.PublicKey, message)
	_, err = RSADecryptPKCS1v15(otherKeyPair.PrivateKey, ciphertext)
	if err == nil {
		t.Error("Expected error for wrong private key in PKCS1v15 decrypt")
	}
}

// TestRSASignatureErrorPaths tests various signature verification error conditions
func TestRSASignatureErrorPaths(t *testing.T) {
	keyPair, err := GenerateRSAKeyPair(2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA key pair: %v", err)
	}

	message := []byte("test message for signing")

	// Test verification with invalid signature
	invalidSignature := []byte("invalid signature")

	err = RSAVerifyPSS(keyPair.PublicKey, message, invalidSignature)
	if err == nil {
		t.Error("Expected error for invalid signature in PSS verify")
	}

	err = RSAVerifyPKCS1v15(keyPair.PublicKey, message, invalidSignature)
	if err == nil {
		t.Error("Expected error for invalid signature in PKCS1v15 verify")
	}

	// Test verification with wrong key (different key pair)
	otherKeyPair, _ := GenerateRSAKeyPair(2048)
	signature, _ := RSASignPSS(keyPair.PrivateKey, message)

	err = RSAVerifyPSS(otherKeyPair.PublicKey, message, signature)
	if err == nil {
		t.Error("Expected error for wrong public key in PSS verify")
	}

	signature, _ = RSASignPKCS1v15(keyPair.PrivateKey, message)
	err = RSAVerifyPKCS1v15(otherKeyPair.PublicKey, message, signature)
	if err == nil {
		t.Error("Expected error for wrong public key in PKCS1v15 verify")
	}

	// Test verification with modified message
	modifiedMessage := []byte("modified test message for signing")
	signature, _ = RSASignPSS(keyPair.PrivateKey, message)

	err = RSAVerifyPSS(keyPair.PublicKey, modifiedMessage, signature)
	if err == nil {
		t.Error("Expected error for modified message in PSS verify")
	}

	signature, _ = RSASignPKCS1v15(keyPair.PrivateKey, message)
	err = RSAVerifyPKCS1v15(keyPair.PublicKey, modifiedMessage, signature)
	if err == nil {
		t.Error("Expected error for modified message in PKCS1v15 verify")
	}
}

// TestRSAPEMErrorPaths tests PEM encoding/decoding error conditions
func TestRSAPEMErrorPaths(t *testing.T) {
	// Test with nil keys
	nilKeyPair := &RSAKeyPair{}

	_, err := nilKeyPair.PrivateKeyToPEM()
	if err == nil {
		t.Error("Expected error for nil private key in PEM conversion")
	}

	_, err = nilKeyPair.PublicKeyToPEM()
	if err == nil {
		t.Error("Expected error for nil public key in PEM conversion")
	}

	// Test with invalid PEM data
	invalidPEMData := []byte("not a PEM block")

	_, err = PrivateKeyFromPEM(invalidPEMData)
	if err == nil {
		t.Error("Expected error for invalid PEM data in private key parsing")
	}

	_, err = PublicKeyFromPEM(invalidPEMData)
	if err == nil {
		t.Error("Expected error for invalid PEM data in public key parsing")
	}

	// Test with empty PEM data
	emptyPEMData := []byte("")

	_, err = PrivateKeyFromPEM(emptyPEMData)
	if err == nil {
		t.Error("Expected error for empty PEM data in private key parsing")
	}

	_, err = PublicKeyFromPEM(emptyPEMData)
	if err == nil {
		t.Error("Expected error for empty PEM data in public key parsing")
	}

	// Test with wrong PEM block type
	wrongTypePEM := []byte(`-----BEGIN CERTIFICATE-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA
-----END CERTIFICATE-----`)

	_, err = PrivateKeyFromPEM(wrongTypePEM)
	if err == nil {
		t.Error("Expected error for wrong PEM block type in private key parsing")
	}

	_, err = PublicKeyFromPEM(wrongTypePEM)
	if err == nil {
		t.Error("Expected error for wrong PEM block type in public key parsing")
	}

	// Test with malformed PEM content
	malformedPEM := []byte(`-----BEGIN PRIVATE KEY-----
invalid base64 content
-----END PRIVATE KEY-----`)

	_, err = PrivateKeyFromPEM(malformedPEM)
	if err == nil {
		t.Error("Expected error for malformed PEM content in private key parsing")
	}

	malformedPEM = []byte(`-----BEGIN PUBLIC KEY-----
invalid base64 content
-----END PUBLIC KEY-----`)

	_, err = PublicKeyFromPEM(malformedPEM)
	if err == nil {
		t.Error("Expected error for malformed PEM content in public key parsing")
	}

	// Test with valid PEM but wrong key type (ECDSA key as RSA)
	// Create an ECDSA key and try to parse it as RSA
	ecdsaPEM := []byte(`-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE...
-----END PUBLIC KEY-----`) // Truncated ECDSA key

	_, err = PublicKeyFromPEM(ecdsaPEM)
	if err == nil {
		t.Error("Expected error for ECDSA key in RSA parser")
	}
}

// TestRSAKeyGenerationEdgeCases tests edge cases in key generation
func TestRSAKeyGenerationEdgeCases(t *testing.T) {
	// Test with minimum allowed key size
	_, err := GenerateRSAKeyPair(1024)
	if err != nil {
		t.Errorf("Should succeed with 1024-bit key: %v", err)
	}

	// Test with various invalid key sizes
	invalidSizes := []int{0, 512, 1023}
	for _, size := range invalidSizes {
		_, err := GenerateRSAKeyPair(size)
		if err == nil {
			t.Errorf("Expected error for invalid key size %d", size)
		}
	}

	// Test with odd key sizes (should still work for valid sizes)
	oddSizes := []int{1025, 2049}
	for _, size := range oddSizes {
		_, err := GenerateRSAKeyPair(size)
		if err != nil {
			t.Errorf("Should handle odd key size %d: %v", size, err)
		}
	}
}

// TestRSAMaxMessageLengthEdgeCases tests edge cases in max message length calculation
func TestRSAMaxMessageLengthEdgeCases(t *testing.T) {
	keyPair, err := GenerateRSAKeyPair(2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA key pair: %v", err)
	}

	// Test with various padding types
	paddingTypes := []string{"OAEP", "PKCS1v15", "oaep", "pkcs1v15"}
	for _, padding := range paddingTypes {
		length, err := RSAMaxMessageLength(keyPair.PublicKey, padding)
		if err != nil {
			t.Errorf("Should handle padding type %s: %v", padding, err)
		} else if length <= 0 {
			t.Errorf("Max message length should be positive for %s, got %d", padding, length)
		}
	}

	// Test with invalid padding types
	invalidPaddings := []string{"", "INVALID", "RSA", "AES"}
	for _, padding := range invalidPaddings {
		_, err := RSAMaxMessageLength(keyPair.PublicKey, padding)
		if err == nil {
			t.Errorf("Expected error for invalid padding type %s", padding)
		}
	}

	// Test with nil public key
	_, err = RSAMaxMessageLength(nil, "OAEP")
	if err == nil {
		t.Error("Expected error for nil public key")
	}
}

// TestRSAWithDifferentMessageSizes tests encryption/decryption with various message sizes
func TestRSAWithDifferentMessageSizes(t *testing.T) {
	keyPair, err := GenerateRSAKeyPair(2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA key pair: %v", err)
	}

	// Get max message lengths
	maxOAEP, _ := RSAMaxMessageLength(keyPair.PublicKey, "OAEP")
	maxPKCS1v15, _ := RSAMaxMessageLength(keyPair.PublicKey, "PKCS1v15")

	// Test with various message sizes for OAEP
	oaepSizes := []int{0, 1, 10, 50, maxOAEP}
	for _, size := range oaepSizes {
		message := make([]byte, size)
		for i := range message {
			message[i] = byte(i % 256)
		}

		ciphertext, err := RSAEncryptOAEP(keyPair.PublicKey, message)
		if err != nil {
			t.Errorf("OAEP encryption failed for size %d: %v", size, err)
			continue
		}

		decrypted, err := RSADecryptOAEP(keyPair.PrivateKey, ciphertext)
		if err != nil {
			t.Errorf("OAEP decryption failed for size %d: %v", size, err)
			continue
		}

		if !bytes.Equal(message, decrypted) {
			t.Errorf("OAEP round-trip failed for size %d", size)
		}
	}

	// Test with various message sizes for PKCS1v15
	pkcs1v15Sizes := []int{0, 1, 10, 50, maxPKCS1v15}
	for _, size := range pkcs1v15Sizes {
		message := make([]byte, size)
		for i := range message {
			message[i] = byte(i % 256)
		}

		ciphertext, err := RSAEncryptPKCS1v15(keyPair.PublicKey, message)
		if err != nil {
			t.Errorf("PKCS1v15 encryption failed for size %d: %v", size, err)
			continue
		}

		decrypted, err := RSADecryptPKCS1v15(keyPair.PrivateKey, ciphertext)
		if err != nil {
			t.Errorf("PKCS1v15 decryption failed for size %d: %v", size, err)
			continue
		}

		if !bytes.Equal(message, decrypted) {
			t.Errorf("PKCS1v15 round-trip failed for size %d", size)
		}
	}
}

// TestRSAPEMRoundTrip tests complete PEM encoding/decoding round-trip
func TestRSAPEMRoundTrip(t *testing.T) {
	originalKeyPair, err := GenerateRSAKeyPair(2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA key pair: %v", err)
	}

	// Test private key round-trip
	privatePEM, err := originalKeyPair.PrivateKeyToPEM()
	if err != nil {
		t.Fatalf("Failed to encode private key to PEM: %v", err)
	}

	decodedPrivateKey, err := PrivateKeyFromPEM(privatePEM)
	if err != nil {
		t.Fatalf("Failed to decode private key from PEM: %v", err)
	}

	// Verify private keys are equivalent (same modulus and exponent)
	if originalKeyPair.PrivateKey.N.Cmp(decodedPrivateKey.N) != 0 {
		t.Error("Private key modulus doesn't match after PEM round-trip")
	}

	if originalKeyPair.PrivateKey.E != decodedPrivateKey.E {
		t.Error("Private key exponent doesn't match after PEM round-trip")
	}

	// Test public key round-trip
	publicPEM, err := originalKeyPair.PublicKeyToPEM()
	if err != nil {
		t.Fatalf("Failed to encode public key to PEM: %v", err)
	}

	decodedPublicKey, err := PublicKeyFromPEM(publicPEM)
	if err != nil {
		t.Fatalf("Failed to decode public key from PEM: %v", err)
	}

	// Verify public keys are equivalent
	if originalKeyPair.PublicKey.N.Cmp(decodedPublicKey.N) != 0 {
		t.Error("Public key modulus doesn't match after PEM round-trip")
	}

	if originalKeyPair.PublicKey.E != decodedPublicKey.E {
		t.Error("Public key exponent doesn't match after PEM round-trip")
	}

	// Test that the decoded keys can be used for encryption/decryption
	message := []byte("test message for PEM round-trip")

	ciphertext, err := RSAEncryptOAEP(decodedPublicKey, message)
	if err != nil {
		t.Fatalf("Failed to encrypt with decoded public key: %v", err)
	}

	decrypted, err := RSADecryptOAEP(decodedPrivateKey, ciphertext)
	if err != nil {
		t.Fatalf("Failed to decrypt with decoded private key: %v", err)
	}

	if !bytes.Equal(message, decrypted) {
		t.Error("Message doesn't match after encryption/decryption with decoded keys")
	}
}

// TestRSASignatureIntegrity tests signature integrity across different scenarios
func TestRSASignatureIntegrity(t *testing.T) {
	keyPair, err := GenerateRSAKeyPair(2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA key pair: %v", err)
	}

	testMessages := [][]byte{
		[]byte(""),                                    // Empty message
		[]byte("A"),                                   // Single character
		[]byte("Hello, World!"),                       // Short message
		[]byte("This is a longer message for testing RSA signature integrity with various message lengths."), // Long message
		bytes.Repeat([]byte("X"), 1000),               // Very long message
	}

	for i, message := range testMessages {
		t.Run(fmt.Sprintf("Message_%d", i), func(t *testing.T) {
			// Test PSS signatures
			pssSignature, err := RSASignPSS(keyPair.PrivateKey, message)
			if err != nil {
				t.Fatalf("PSS signing failed: %v", err)
			}

			err = RSAVerifyPSS(keyPair.PublicKey, message, pssSignature)
			if err != nil {
				t.Errorf("PSS verification failed: %v", err)
			}

			// Test PKCS1v15 signatures
			pkcs1v15Signature, err := RSASignPKCS1v15(keyPair.PrivateKey, message)
			if err != nil {
				t.Fatalf("PKCS1v15 signing failed: %v", err)
			}

			err = RSAVerifyPKCS1v15(keyPair.PublicKey, message, pkcs1v15Signature)
			if err != nil {
				t.Errorf("PKCS1v15 verification failed: %v", err)
			}

			// Verify cross-verification fails (PSS signature with PKCS1v15 verify)
			err = RSAVerifyPKCS1v15(keyPair.PublicKey, message, pssSignature)
			if err == nil {
				t.Error("Cross-verification should fail (PSS sig with PKCS1v15 verify)")
			}

			err = RSAVerifyPSS(keyPair.PublicKey, message, pkcs1v15Signature)
			if err == nil {
				t.Error("Cross-verification should fail (PKCS1v15 sig with PSS verify)")
			}
		})
	}
}
