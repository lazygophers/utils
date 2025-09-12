package cryptox

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"hash"
	"math/big"
)

// Global variables for dependency injection during testing
var (
	ecdhRandReader = rand.Reader
)

// ECDHKeyPair represents an ECDH key pair (same as ECDSA but used for key exchange)
type ECDHKeyPair struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
}

// GenerateECDHKey 生成 ECDH 密钥对
func GenerateECDHKey(curve elliptic.Curve) (*ECDHKeyPair, error) {
	if curve == nil {
		return nil, errors.New("curve cannot be nil")
	}

	privateKey, err := ecdsa.GenerateKey(curve, ecdhRandReader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate ECDH key: %w", err)
	}

	return &ECDHKeyPair{
		PrivateKey: privateKey,
		PublicKey:  &privateKey.PublicKey,
	}, nil
}

// GenerateECDHP256Key 生成 P-256 ECDH 密钥对
func GenerateECDHP256Key() (*ECDHKeyPair, error) {
	return GenerateECDHKey(elliptic.P256())
}

// GenerateECDHP384Key 生成 P-384 ECDH 密钥对
func GenerateECDHP384Key() (*ECDHKeyPair, error) {
	return GenerateECDHKey(elliptic.P384())
}

// GenerateECDHP521Key 生成 P-521 ECDH 密钥对
func GenerateECDHP521Key() (*ECDHKeyPair, error) {
	return GenerateECDHKey(elliptic.P521())
}

// ECDHComputeShared 计算 ECDH 共享密钥
func ECDHComputeShared(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) ([]byte, error) {
	if privateKey == nil {
		return nil, errors.New("private key cannot be nil")
	}
	if publicKey == nil {
		return nil, errors.New("public key cannot be nil")
	}

	// 验证公钥是否在曲线上
	if !privateKey.Curve.IsOnCurve(publicKey.X, publicKey.Y) {
		return nil, errors.New("public key is not on the curve")
	}

	// 验证曲线是否匹配
	if privateKey.Curve != publicKey.Curve {
		return nil, errors.New("curve mismatch between private and public keys")
	}

	// 计算共享点
	x, _ := privateKey.Curve.ScalarMult(publicKey.X, publicKey.Y, privateKey.D.Bytes())

	// 返回 x 坐标作为共享密钥
	return x.Bytes(), nil
}

// ECDHComputeSharedWithKDF 计算 ECDH 共享密钥并使用 KDF 派生最终密钥
func ECDHComputeSharedWithKDF(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, keyLength int, kdf func() hash.Hash) ([]byte, error) {
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

	// 使用简单的 KDF (Hash-based)
	h := kdf()
	h.Write(sharedSecret)
	derivedKey := h.Sum(nil)

	// 如果需要的密钥长度超过哈希输出长度，截断或重复
	if len(derivedKey) >= keyLength {
		return derivedKey[:keyLength], nil
	}

	// 如果需要更长的密钥，使用计数器模式扩展
	result := make([]byte, 0, keyLength)
	counter := uint32(0)
	
	for len(result) < keyLength {
		h := kdf()
		h.Write(sharedSecret)
		h.Write([]byte{byte(counter >> 24), byte(counter >> 16), byte(counter >> 8), byte(counter)})
		block := h.Sum(nil)
		result = append(result, block...)
		counter++
	}

	return result[:keyLength], nil
}

// ECDHComputeSharedSHA256 使用 SHA256 KDF 计算 ECDH 共享密钥
func ECDHComputeSharedSHA256(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, keyLength int) ([]byte, error) {
	return ECDHComputeSharedWithKDF(privateKey, publicKey, keyLength, sha256.New)
}

// ECDHKeyExchange 执行完整的 ECDH 密钥交换
func ECDHKeyExchange(alicePrivateKey *ecdsa.PrivateKey, bobPublicKey *ecdsa.PublicKey, keyLength int) ([]byte, error) {
	return ECDHComputeSharedSHA256(alicePrivateKey, bobPublicKey, keyLength)
}

// ValidateECDHKeyPair 验证 ECDH 密钥对的有效性
func ValidateECDHKeyPair(keyPair *ECDHKeyPair) error {
	if keyPair == nil {
		return errors.New("key pair cannot be nil")
	}
	if keyPair.PrivateKey == nil {
		return errors.New("private key cannot be nil")
	}
	if keyPair.PublicKey == nil {
		return errors.New("public key cannot be nil")
	}

	// 验证公钥是否在曲线上
	if !keyPair.PrivateKey.Curve.IsOnCurve(keyPair.PublicKey.X, keyPair.PublicKey.Y) {
		return errors.New("public key is not on the curve")
	}

	// 验证曲线是否匹配
	if keyPair.PrivateKey.Curve != keyPair.PublicKey.Curve {
		return errors.New("curve mismatch between private and public keys")
	}

	// 验证公钥是否与私钥匹配
	expectedX, expectedY := keyPair.PrivateKey.Curve.ScalarBaseMult(keyPair.PrivateKey.D.Bytes())
	if keyPair.PublicKey.X.Cmp(expectedX) != 0 || keyPair.PublicKey.Y.Cmp(expectedY) != 0 {
		return errors.New("public key does not match private key")
	}

	return nil
}

// ECDHPublicKeyFromCoordinates 从 x, y 坐标创建 ECDH 公钥
func ECDHPublicKeyFromCoordinates(curve elliptic.Curve, x, y *big.Int) (*ecdsa.PublicKey, error) {
	if curve == nil {
		return nil, errors.New("curve cannot be nil")
	}
	if x == nil || y == nil {
		return nil, errors.New("coordinates cannot be nil")
	}

	// 验证点是否在曲线上
	if !curve.IsOnCurve(x, y) {
		return nil, errors.New("point is not on the curve")
	}

	return &ecdsa.PublicKey{
		Curve: curve,
		X:     new(big.Int).Set(x),
		Y:     new(big.Int).Set(y),
	}, nil
}

// ECDHPublicKeyToCoordinates 将 ECDH 公钥转换为 x, y 坐标
func ECDHPublicKeyToCoordinates(publicKey *ecdsa.PublicKey) (x, y *big.Int, err error) {
	if publicKey == nil {
		return nil, nil, errors.New("public key cannot be nil")
	}

	return new(big.Int).Set(publicKey.X), new(big.Int).Set(publicKey.Y), nil
}

// ECDHSharedSecretTest 测试两个密钥对是否能生成相同的共享密钥（用于测试）
func ECDHSharedSecretTest(keyPair1, keyPair2 *ECDHKeyPair) (bool, error) {
	if keyPair1 == nil || keyPair2 == nil {
		return false, errors.New("key pairs cannot be nil")
	}

	// Alice 使用她的私钥和 Bob 的公钥
	secret1, err := ECDHComputeShared(keyPair1.PrivateKey, keyPair2.PublicKey)
	if err != nil {
		return false, err
	}

	// Bob 使用他的私钥和 Alice 的公钥
	secret2, err := ECDHComputeShared(keyPair2.PrivateKey, keyPair1.PublicKey)
	if err != nil {
		return false, err
	}

	// 比较共享密钥是否相同
	if len(secret1) != len(secret2) {
		return false, nil
	}

	for i := 0; i < len(secret1); i++ {
		if secret1[i] != secret2[i] {
			return false, nil
		}
	}

	return true, nil
}