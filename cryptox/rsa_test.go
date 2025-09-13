package cryptox

import (
	"bytes"
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
