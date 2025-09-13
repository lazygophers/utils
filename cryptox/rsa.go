package cryptox

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

// Global variables for dependency injection during testing
var (
	rsaGenerateKey     = rsa.GenerateKey
	rsaEncryptOAEP     = rsa.EncryptOAEP
	rsaEncryptPKCS1v15 = rsa.EncryptPKCS1v15
	rsaDecryptOAEP     = rsa.DecryptOAEP
	rsaDecryptPKCS1v15 = rsa.DecryptPKCS1v15
	rsaSignPSS         = rsa.SignPSS
	rsaSignPKCS1v15    = rsa.SignPKCS1v15
	rsaVerifyPSS       = rsa.VerifyPSS
	rsaVerifyPKCS1v15  = rsa.VerifyPKCS1v15
	rsaRandReader      = rand.Reader
)

// RSAKeyPair 表示 RSA 公私钥对
type RSAKeyPair struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// GenerateRSAKeyPair 生成指定长度的 RSA 密钥对
// keySize: 密钥长度，建议使用 2048、3072 或 4096 位
func GenerateRSAKeyPair(keySize int) (*RSAKeyPair, error) {
	if keySize < 1024 {
		return nil, errors.New("RSA key size must be at least 1024 bits")
	}

	privateKey, err := rsaGenerateKey(rsaRandReader, keySize)
	if err != nil {
		return nil, fmt.Errorf("failed to generate RSA key pair: %w", err)
	}

	return &RSAKeyPair{
		PrivateKey: privateKey,
		PublicKey:  &privateKey.PublicKey,
	}, nil
}

// PrivateKeyToPEM 将私钥转换为 PEM 格式
func (kp *RSAKeyPair) PrivateKeyToPEM() ([]byte, error) {
	if kp.PrivateKey == nil {
		return nil, errors.New("private key is nil")
	}

	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(kp.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal private key: %w", err)
	}

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	return privateKeyPEM, nil
}

// PublicKeyToPEM 将公钥转换为 PEM 格式
func (kp *RSAKeyPair) PublicKeyToPEM() ([]byte, error) {
	if kp.PublicKey == nil {
		return nil, errors.New("public key is nil")
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(kp.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal public key: %w", err)
	}

	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	return publicKeyPEM, nil
}

// PrivateKeyFromPEM 从 PEM 格式加载私钥
func PrivateKeyFromPEM(pemData []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	if block.Type != "PRIVATE KEY" && block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("invalid PEM block type: %s", block.Type)
	}

	var privateKey *rsa.PrivateKey
	var err error

	if block.Type == "PRIVATE KEY" {
		// PKCS#8 format
		key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse PKCS8 private key: %w", err)
		}

		var ok bool
		privateKey, ok = key.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("key is not an RSA private key")
		}
	} else {
		// PKCS#1 format
		privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse PKCS1 private key: %w", err)
		}
	}

	return privateKey, nil
}

// PublicKeyFromPEM 从 PEM 格式加载公钥
func PublicKeyFromPEM(pemData []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	if block.Type != "PUBLIC KEY" && block.Type != "RSA PUBLIC KEY" {
		return nil, fmt.Errorf("invalid PEM block type: %s", block.Type)
	}

	var publicKey *rsa.PublicKey
	var err error

	if block.Type == "PUBLIC KEY" {
		// PKIX format
		key, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse PKIX public key: %w", err)
		}

		var ok bool
		publicKey, ok = key.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("key is not an RSA public key")
		}
	} else {
		// PKCS#1 format
		publicKey, err = x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse PKCS1 public key: %w", err)
		}
	}

	return publicKey, nil
}

// RSAEncryptOAEP 使用 OAEP 填充方式进行 RSA 加密
func RSAEncryptOAEP(publicKey *rsa.PublicKey, plaintext []byte) ([]byte, error) {
	if publicKey == nil {
		return nil, errors.New("public key is nil")
	}

	hash := sha256.New()
	ciphertext, err := rsaEncryptOAEP(hash, rsaRandReader, publicKey, plaintext, nil)
	if err != nil {
		return nil, fmt.Errorf("RSA OAEP encryption failed: %w", err)
	}

	return ciphertext, nil
}

// RSADecryptOAEP 使用 OAEP 填充方式进行 RSA 解密
func RSADecryptOAEP(privateKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	if privateKey == nil {
		return nil, errors.New("private key is nil")
	}

	hash := sha256.New()
	plaintext, err := rsaDecryptOAEP(hash, rsaRandReader, privateKey, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("RSA OAEP decryption failed: %w", err)
	}

	return plaintext, nil
}

// RSAEncryptPKCS1v15 使用 PKCS1v15 填充方式进行 RSA 加密
func RSAEncryptPKCS1v15(publicKey *rsa.PublicKey, plaintext []byte) ([]byte, error) {
	if publicKey == nil {
		return nil, errors.New("public key is nil")
	}

	ciphertext, err := rsaEncryptPKCS1v15(rsaRandReader, publicKey, plaintext)
	if err != nil {
		return nil, fmt.Errorf("RSA PKCS1v15 encryption failed: %w", err)
	}

	return ciphertext, nil
}

// RSADecryptPKCS1v15 使用 PKCS1v15 填充方式进行 RSA 解密
func RSADecryptPKCS1v15(privateKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	if privateKey == nil {
		return nil, errors.New("private key is nil")
	}

	plaintext, err := rsaDecryptPKCS1v15(rsaRandReader, privateKey, ciphertext)
	if err != nil {
		return nil, fmt.Errorf("RSA PKCS1v15 decryption failed: %w", err)
	}

	return plaintext, nil
}

// RSASignPSS 使用 PSS 填充方式进行数字签名
func RSASignPSS(privateKey *rsa.PrivateKey, message []byte) ([]byte, error) {
	if privateKey == nil {
		return nil, errors.New("private key is nil")
	}

	hash := sha256.Sum256(message)
	signature, err := rsaSignPSS(rsaRandReader, privateKey, crypto.SHA256, hash[:], nil)
	if err != nil {
		return nil, fmt.Errorf("RSA PSS signing failed: %w", err)
	}

	return signature, nil
}

// RSAVerifyPSS 使用 PSS 填充方式验证数字签名
func RSAVerifyPSS(publicKey *rsa.PublicKey, message []byte, signature []byte) error {
	if publicKey == nil {
		return errors.New("public key is nil")
	}

	hash := sha256.Sum256(message)
	err := rsaVerifyPSS(publicKey, crypto.SHA256, hash[:], signature, nil)
	if err != nil {
		return fmt.Errorf("RSA PSS signature verification failed: %w", err)
	}

	return nil
}

// RSASignPKCS1v15 使用 PKCS1v15 填充方式进行数字签名
func RSASignPKCS1v15(privateKey *rsa.PrivateKey, message []byte) ([]byte, error) {
	if privateKey == nil {
		return nil, errors.New("private key is nil")
	}

	hash := sha256.Sum256(message)
	signature, err := rsaSignPKCS1v15(rsaRandReader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return nil, fmt.Errorf("RSA PKCS1v15 signing failed: %w", err)
	}

	return signature, nil
}

// RSAVerifyPKCS1v15 使用 PKCS1v15 填充方式验证数字签名
func RSAVerifyPKCS1v15(publicKey *rsa.PublicKey, message []byte, signature []byte) error {
	if publicKey == nil {
		return errors.New("public key is nil")
	}

	hash := sha256.Sum256(message)
	err := rsaVerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		return fmt.Errorf("RSA PKCS1v15 signature verification failed: %w", err)
	}

	return nil
}

// GetRSAKeySize 获取 RSA 密钥长度（位）
func GetRSAKeySize(key *rsa.PublicKey) int {
	if key == nil {
		return 0
	}
	return key.Size() * 8
}

// RSAMaxMessageLength 计算 RSA 加密时的最大消息长度
func RSAMaxMessageLength(publicKey *rsa.PublicKey, padding string) (int, error) {
	if publicKey == nil {
		return 0, errors.New("public key is nil")
	}

	keySize := publicKey.Size()

	switch padding {
	case "OAEP":
		// OAEP: keySize - 2*hashLen - 2 (SHA256 hash length is 32 bytes)
		return keySize - 2*32 - 2, nil
	case "PKCS1v15":
		// PKCS1v15: keySize - 11
		return keySize - 11, nil
	default:
		return 0, fmt.Errorf("unsupported padding: %s", padding)
	}
}
