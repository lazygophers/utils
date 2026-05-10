package cryptox

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"hash"
	"math/big"
	"sync"
)

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
