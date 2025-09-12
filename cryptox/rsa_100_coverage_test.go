package cryptox

import (
	"crypto"
	"crypto/rsa"
	"errors"
	"hash"
	"io"
	"testing"
)

// Mock failures for dependency injection
type FailingRSAReader struct{}

func (fr FailingRSAReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("simulated random reader failure")
}

func FailingRSAGenerateKey(random io.Reader, bits int) (*rsa.PrivateKey, error) {
	return nil, errors.New("simulated RSA key generation failure")
}

func FailingRSAEncryptOAEP(hash hash.Hash, random io.Reader, pub *rsa.PublicKey, msg []byte, label []byte) ([]byte, error) {
	return nil, errors.New("simulated RSA OAEP encryption failure")
}

func FailingRSADecryptOAEP(hash hash.Hash, random io.Reader, priv *rsa.PrivateKey, ciphertext []byte, label []byte) ([]byte, error) {
	return nil, errors.New("simulated RSA OAEP decryption failure")
}

func FailingRSAEncryptPKCS1v15(random io.Reader, pub *rsa.PublicKey, msg []byte) ([]byte, error) {
	return nil, errors.New("simulated RSA PKCS1v15 encryption failure")
}

func FailingRSADecryptPKCS1v15(random io.Reader, priv *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	return nil, errors.New("simulated RSA PKCS1v15 decryption failure")
}

func FailingRSASignPSS(random io.Reader, priv *rsa.PrivateKey, hash crypto.Hash, hashed []byte, opts *rsa.PSSOptions) ([]byte, error) {
	return nil, errors.New("simulated RSA PSS signing failure")
}

func FailingRSASignPKCS1v15(random io.Reader, priv *rsa.PrivateKey, hash crypto.Hash, hashed []byte) ([]byte, error) {
	return nil, errors.New("simulated RSA PKCS1v15 signing failure")
}

func FailingRSAVerifyPSS(pub *rsa.PublicKey, hash crypto.Hash, hashed []byte, sig []byte, opts *rsa.PSSOptions) error {
	return errors.New("simulated RSA PSS verification failure")
}

func FailingRSAVerifyPKCS1v15(pub *rsa.PublicKey, hash crypto.Hash, hashed []byte, sig []byte) error {
	return errors.New("simulated RSA PKCS1v15 verification failure")
}

// TestRSA100PercentCoverage triggers all error paths using dependency injection
func TestRSA100PercentCoverage(t *testing.T) {
	// Save original functions
	originalRSAGenerateKey := rsaGenerateKey
	originalRSAEncryptOAEP := rsaEncryptOAEP
	originalRSAEncryptPKCS1v15 := rsaEncryptPKCS1v15
	originalRSADecryptOAEP := rsaDecryptOAEP
	originalRSADecryptPKCS1v15 := rsaDecryptPKCS1v15
	originalRSASignPSS := rsaSignPSS
	originalRSASignPKCS1v15 := rsaSignPKCS1v15
	originalRSAVerifyPSS := rsaVerifyPSS
	originalRSAVerifyPKCS1v15 := rsaVerifyPKCS1v15
	originalRSARandReader := rsaRandReader

	// Restore original functions after test
	defer func() {
		rsaGenerateKey = originalRSAGenerateKey
		rsaEncryptOAEP = originalRSAEncryptOAEP
		rsaEncryptPKCS1v15 = originalRSAEncryptPKCS1v15
		rsaDecryptOAEP = originalRSADecryptOAEP
		rsaDecryptPKCS1v15 = originalRSADecryptPKCS1v15
		rsaSignPSS = originalRSASignPSS
		rsaSignPKCS1v15 = originalRSASignPKCS1v15
		rsaVerifyPSS = originalRSAVerifyPSS
		rsaVerifyPKCS1v15 = originalRSAVerifyPKCS1v15
		rsaRandReader = originalRSARandReader
	}()

	// Test 1: Trigger RSA key generation failure
	rsaGenerateKey = FailingRSAGenerateKey
	_, err := GenerateRSAKeyPair(2048)
	if err == nil {
		t.Error("Expected RSA key generation error")
	}
	rsaGenerateKey = originalRSAGenerateKey

	// Create a valid key pair for other tests
	keyPair, err := GenerateRSAKeyPair(2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA key pair: %v", err)
	}

	message := []byte("test message")
	
	// Test 2: Test encryption failures
	rsaEncryptOAEP = FailingRSAEncryptOAEP
	_, err = RSAEncryptOAEP(keyPair.PublicKey, message)
	if err == nil {
		t.Error("Expected RSA OAEP encryption error")
	}
	rsaEncryptOAEP = originalRSAEncryptOAEP

	rsaEncryptPKCS1v15 = FailingRSAEncryptPKCS1v15
	_, err = RSAEncryptPKCS1v15(keyPair.PublicKey, message)
	if err == nil {
		t.Error("Expected RSA PKCS1v15 encryption error")
	}
	rsaEncryptPKCS1v15 = originalRSAEncryptPKCS1v15

	// Test 3: Test decryption failures
	// First encrypt a message normally
	ciphertextOAEP, _ := RSAEncryptOAEP(keyPair.PublicKey, message)
	ciphertextPKCS1v15, _ := RSAEncryptPKCS1v15(keyPair.PublicKey, message)

	rsaDecryptOAEP = FailingRSADecryptOAEP
	_, err = RSADecryptOAEP(keyPair.PrivateKey, ciphertextOAEP)
	if err == nil {
		t.Error("Expected RSA OAEP decryption error")
	}
	rsaDecryptOAEP = originalRSADecryptOAEP

	rsaDecryptPKCS1v15 = FailingRSADecryptPKCS1v15
	_, err = RSADecryptPKCS1v15(keyPair.PrivateKey, ciphertextPKCS1v15)
	if err == nil {
		t.Error("Expected RSA PKCS1v15 decryption error")
	}
	rsaDecryptPKCS1v15 = originalRSADecryptPKCS1v15

	// Test 4: Test signing failures
	rsaSignPSS = FailingRSASignPSS
	_, err = RSASignPSS(keyPair.PrivateKey, message)
	if err == nil {
		t.Error("Expected RSA PSS signing error")
	}
	rsaSignPSS = originalRSASignPSS

	rsaSignPKCS1v15 = FailingRSASignPKCS1v15
	_, err = RSASignPKCS1v15(keyPair.PrivateKey, message)
	if err == nil {
		t.Error("Expected RSA PKCS1v15 signing error")
	}
	rsaSignPKCS1v15 = originalRSASignPKCS1v15

	// Test 5: Test verification failures
	// First sign a message normally
	signaturePSS, _ := RSASignPSS(keyPair.PrivateKey, message)
	signaturePKCS1v15, _ := RSASignPKCS1v15(keyPair.PrivateKey, message)

	rsaVerifyPSS = FailingRSAVerifyPSS
	err = RSAVerifyPSS(keyPair.PublicKey, message, signaturePSS)
	if err == nil {
		t.Error("Expected RSA PSS verification error")
	}
	rsaVerifyPSS = originalRSAVerifyPSS

	rsaVerifyPKCS1v15 = FailingRSAVerifyPKCS1v15
	err = RSAVerifyPKCS1v15(keyPair.PublicKey, message, signaturePKCS1v15)
	if err == nil {
		t.Error("Expected RSA PKCS1v15 verification error")
	}
	rsaVerifyPKCS1v15 = originalRSAVerifyPKCS1v15
}

// TestInvalidPEMData tests PEM parsing with invalid data
func TestInvalidPEMData(t *testing.T) {
	// Test with invalid PEM data (not actually PEM format)
	invalidPEMData := []byte("not a valid PEM block")
	
	_, err := PrivateKeyFromPEM(invalidPEMData)
	if err == nil || err.Error() != "failed to decode PEM block" {
		t.Error("Expected 'failed to decode PEM block' error")
	}
	
	_, err = PublicKeyFromPEM(invalidPEMData)
	if err == nil || err.Error() != "failed to decode PEM block" {
		t.Error("Expected 'failed to decode PEM block' error")
	}
	
	// Test with valid PEM block but wrong type
	wrongTypePEM := []byte(`-----BEGIN CERTIFICATE-----
MIICljCCAX4CCQCKxWPjvkCk0jANBgkqhkiG9w0BAQsFADA
-----END CERTIFICATE-----`)
	
	_, err = PrivateKeyFromPEM(wrongTypePEM)
	if err == nil {
		t.Error("Expected error for wrong PEM block type")
	}
	
	_, err = PublicKeyFromPEM(wrongTypePEM)
	if err == nil {
		t.Error("Expected error for wrong PEM block type")
	}
	
	// Test with valid PEM block but corrupted data
	corruptedPrivateKeyPEM := []byte(`-----BEGIN PRIVATE KEY-----
invalidbase64data
-----END PRIVATE KEY-----`)
	
	_, err = PrivateKeyFromPEM(corruptedPrivateKeyPEM)
	if err == nil {
		t.Error("Expected error for corrupted private key PEM")
	}
	
	corruptedPublicKeyPEM := []byte(`-----BEGIN PUBLIC KEY-----
invalidbase64data
-----END PUBLIC KEY-----`)
	
	_, err = PublicKeyFromPEM(corruptedPublicKeyPEM)
	if err == nil {
		t.Error("Expected error for corrupted public key PEM")
	}
	
	// Test PKCS#1 format
	rsaPrivateKeyPEM := []byte(`-----BEGIN RSA PRIVATE KEY-----
invalidbase64data
-----END RSA PRIVATE KEY-----`)
	
	_, err = PrivateKeyFromPEM(rsaPrivateKeyPEM)
	if err == nil {
		t.Error("Expected error for corrupted RSA private key PEM")
	}
	
	rsaPublicKeyPEM := []byte(`-----BEGIN RSA PUBLIC KEY-----
invalidbase64data
-----END RSA PUBLIC KEY-----`)
	
	_, err = PublicKeyFromPEM(rsaPublicKeyPEM)
	if err == nil {
		t.Error("Expected error for corrupted RSA public key PEM")
	}
}

// TestPEMErrorPaths tests error handling in PEM functions
func TestPEMErrorPaths(t *testing.T) {
	
	// Test nil key scenarios for PEM conversion
	nilKeyPair := &RSAKeyPair{PrivateKey: nil, PublicKey: nil}
	
	_, err := nilKeyPair.PrivateKeyToPEM()
	if err == nil || err.Error() != "private key is nil" {
		t.Error("Expected 'private key is nil' error")
	}
	
	_, err = nilKeyPair.PublicKeyToPEM()
	if err == nil || err.Error() != "public key is nil" {
		t.Error("Expected 'public key is nil' error")
	}
	
	// Test with a non-RSA key embedded in PKCS#8 format (simulated by creating invalid PEM)
	nonRSAKeyPEM := []byte(`-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQg
-----END PRIVATE KEY-----`)
	
	_, err = PrivateKeyFromPEM(nonRSAKeyPEM)
	if err == nil {
		t.Error("Expected error for non-RSA key in PKCS#8 format")
	}
	
	nonRSAPublicKeyPEM := []byte(`-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE
-----END PUBLIC KEY-----`)
	
	_, err = PublicKeyFromPEM(nonRSAPublicKeyPEM)
	if err == nil {
		t.Error("Expected error for non-RSA key in PKIX format")
	}
}