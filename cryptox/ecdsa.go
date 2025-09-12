package cryptox

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"hash"
	"math/big"
)

// Global variables for dependency injection during testing
var (
	ecdsaRandReader = rand.Reader
)

// ECDSAKeyPair represents an ECDSA key pair
type ECDSAKeyPair struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
}

// GenerateECDSAKey 生成 ECDSA 密钥对
func GenerateECDSAKey(curve elliptic.Curve) (*ECDSAKeyPair, error) {
	if curve == nil {
		return nil, errors.New("curve cannot be nil")
	}

	privateKey, err := ecdsa.GenerateKey(curve, ecdsaRandReader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate ECDSA key: %w", err)
	}

	return &ECDSAKeyPair{
		PrivateKey: privateKey,
		PublicKey:  &privateKey.PublicKey,
	}, nil
}

// GenerateECDSAP256Key 生成 P-256 (secp256r1) ECDSA 密钥对
func GenerateECDSAP256Key() (*ECDSAKeyPair, error) {
	return GenerateECDSAKey(elliptic.P256())
}

// GenerateECDSAP384Key 生成 P-384 (secp384r1) ECDSA 密钥对
func GenerateECDSAP384Key() (*ECDSAKeyPair, error) {
	return GenerateECDSAKey(elliptic.P384())
}

// GenerateECDSAP521Key 生成 P-521 (secp521r1) ECDSA 密钥对
func GenerateECDSAP521Key() (*ECDSAKeyPair, error) {
	return GenerateECDSAKey(elliptic.P521())
}

// ECDSASign 使用私钥对数据进行 ECDSA 签名
func ECDSASign(privateKey *ecdsa.PrivateKey, data []byte, hashFunc func() hash.Hash) (r, s *big.Int, err error) {
	if privateKey == nil {
		return nil, nil, errors.New("private key cannot be nil")
	}
	if hashFunc == nil {
		return nil, nil, errors.New("hash function cannot be nil")
	}

	h := hashFunc()
	h.Write(data)
	hashed := h.Sum(nil)

	return ecdsa.Sign(ecdsaRandReader, privateKey, hashed)
}

// ECDSASignSHA256 使用 SHA256 哈希对数据进行 ECDSA 签名
func ECDSASignSHA256(privateKey *ecdsa.PrivateKey, data []byte) (r, s *big.Int, err error) {
	return ECDSASign(privateKey, data, sha256.New)
}

// ECDSASignSHA512 使用 SHA512 哈希对数据进行 ECDSA 签名
func ECDSASignSHA512(privateKey *ecdsa.PrivateKey, data []byte) (r, s *big.Int, err error) {
	return ECDSASign(privateKey, data, sha512.New)
}

// ECDSAVerify 使用公钥验证 ECDSA 签名
func ECDSAVerify(publicKey *ecdsa.PublicKey, data []byte, r, s *big.Int, hashFunc func() hash.Hash) bool {
	if publicKey == nil || r == nil || s == nil || hashFunc == nil {
		return false
	}

	h := hashFunc()
	h.Write(data)
	hashed := h.Sum(nil)

	return ecdsa.Verify(publicKey, hashed, r, s)
}

// ECDSAVerifySHA256 使用 SHA256 哈希验证 ECDSA 签名
func ECDSAVerifySHA256(publicKey *ecdsa.PublicKey, data []byte, r, s *big.Int) bool {
	return ECDSAVerify(publicKey, data, r, s, sha256.New)
}

// ECDSAVerifySHA512 使用 SHA512 哈希验证 ECDSA 签名
func ECDSAVerifySHA512(publicKey *ecdsa.PublicKey, data []byte, r, s *big.Int) bool {
	return ECDSAVerify(publicKey, data, r, s, sha512.New)
}

// ECDSAPrivateKeyToPEM 将 ECDSA 私钥转换为 PEM 格式
func ECDSAPrivateKeyToPEM(privateKey *ecdsa.PrivateKey) ([]byte, error) {
	if privateKey == nil {
		return nil, errors.New("private key cannot be nil")
	}

	x509Encoded, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal private key: %w", err)
	}

	pemEncoded := pem.EncodeToMemory(&pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: x509Encoded,
	})

	return pemEncoded, nil
}

// ECDSAPrivateKeyFromPEM 从 PEM 格式解析 ECDSA 私钥
func ECDSAPrivateKeyFromPEM(pemData []byte) (*ecdsa.PrivateKey, error) {
	if len(pemData) == 0 {
		return nil, errors.New("PEM data cannot be empty")
	}

	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	if block.Type != "EC PRIVATE KEY" {
		return nil, fmt.Errorf("invalid PEM block type: expected 'EC PRIVATE KEY', got '%s'", block.Type)
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse EC private key: %w", err)
	}

	return privateKey, nil
}

// ECDSAPublicKeyToPEM 将 ECDSA 公钥转换为 PEM 格式
func ECDSAPublicKeyToPEM(publicKey *ecdsa.PublicKey) ([]byte, error) {
	if publicKey == nil {
		return nil, errors.New("public key cannot be nil")
	}

	x509Encoded, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal public key: %w", err)
	}

	pemEncoded := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509Encoded,
	})

	return pemEncoded, nil
}

// ECDSAPublicKeyFromPEM 从 PEM 格式解析 ECDSA 公钥
func ECDSAPublicKeyFromPEM(pemData []byte) (*ecdsa.PublicKey, error) {
	if len(pemData) == 0 {
		return nil, errors.New("PEM data cannot be empty")
	}

	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	if block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("invalid PEM block type: expected 'PUBLIC KEY', got '%s'", block.Type)
	}

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	publicKey, ok := publicKeyInterface.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("not an ECDSA public key")
	}

	return publicKey, nil
}

// ECDSASignature represents an ECDSA signature
type ECDSASignature struct {
	R *big.Int
	S *big.Int
}

// ECDSASignatureToBytes 将 ECDSA 签名转换为字节数组（DER 编码）
func ECDSASignatureToBytes(r, s *big.Int) ([]byte, error) {
	if r == nil || s == nil {
		return nil, errors.New("signature components cannot be nil")
	}

	// 简单的 DER 编码实现
	rBytes := r.Bytes()
	sBytes := s.Bytes()

	// 如果最高位是1，需要添加0x00前缀
	if len(rBytes) > 0 && rBytes[0]&0x80 != 0 {
		rBytes = append([]byte{0x00}, rBytes...)
	}
	if len(sBytes) > 0 && sBytes[0]&0x80 != 0 {
		sBytes = append([]byte{0x00}, sBytes...)
	}

	// 构建 DER 结构
	rDER := append([]byte{0x02, byte(len(rBytes))}, rBytes...)
	sDER := append([]byte{0x02, byte(len(sBytes))}, sBytes...)
	
	signature := append(rDER, sDER...)
	derEncoded := append([]byte{0x30, byte(len(signature))}, signature...)

	return derEncoded, nil
}

// ECDSASignatureFromBytes 从字节数组解析 ECDSA 签名（DER 解码）
func ECDSASignatureFromBytes(data []byte) (r, s *big.Int, err error) {
	if len(data) < 6 {
		return nil, nil, errors.New("signature data too short")
	}

	// 检查 DER 结构
	if data[0] != 0x30 {
		return nil, nil, errors.New("invalid DER signature: missing SEQUENCE tag")
	}

	seqLen := int(data[1])
	if len(data) < seqLen+2 {
		return nil, nil, errors.New("invalid DER signature: incorrect sequence length")
	}

	data = data[2:] // 跳过 SEQUENCE 头部

	// 解析 r
	if len(data) < 2 || data[0] != 0x02 {
		return nil, nil, errors.New("invalid DER signature: missing INTEGER tag for r")
	}

	rLen := int(data[1])
	if len(data) < rLen+2 {
		return nil, nil, errors.New("invalid DER signature: incorrect r length")
	}

	rBytes := data[2 : 2+rLen]
	r = new(big.Int).SetBytes(rBytes)
	data = data[2+rLen:]

	// 解析 s
	if len(data) < 2 || data[0] != 0x02 {
		return nil, nil, errors.New("invalid DER signature: missing INTEGER tag for s")
	}

	sLen := int(data[1])
	if len(data) < sLen+2 {
		return nil, nil, errors.New("invalid DER signature: incorrect s length")
	}

	sBytes := data[2 : 2+sLen]
	s = new(big.Int).SetBytes(sBytes)

	return r, s, nil
}

// GetCurveName 获取椭圆曲线的名称
func GetCurveName(curve elliptic.Curve) string {
	switch curve {
	case elliptic.P224():
		return "P-224"
	case elliptic.P256():
		return "P-256"
	case elliptic.P384():
		return "P-384"
	case elliptic.P521():
		return "P-521"
	default:
		return "Unknown"
	}
}

// IsValidCurve 检查椭圆曲线是否有效
func IsValidCurve(curve elliptic.Curve) bool {
	switch curve {
	case elliptic.P224(), elliptic.P256(), elliptic.P384(), elliptic.P521():
		return true
	default:
		return false
	}
}