package cryptox

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"errors"
	"io"
)

// Global variables for dependency injection during testing
var (
	desNewCipher    = des.NewCipher
	desNewTripleDES = des.NewTripleDESCipher
	desRandReader   = rand.Reader
)

// DESEncryptECB 使用 DES 在 ECB 模式下加密明文
// 警告：DES 已被认为是不安全的，仅用于兼容性目的。推荐使用 AES。
// 警告：ECB 模式在密码学上是不安全的，相同的明文块会产生相同的密文块。
func DESEncryptECB(key, plaintext []byte) ([]byte, error) {
	if len(key) != 8 {
		return nil, errors.New("invalid key length: must be 8 bytes for DES")
	}

	block, err := desNewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext = padPKCS7(plaintext, des.BlockSize)
	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i += des.BlockSize {
		block.Encrypt(ciphertext[i:i+des.BlockSize], plaintext[i:i+des.BlockSize])
	}
	return ciphertext, nil
}

// DESDecryptECB 使用 DES 在 ECB 模式下解密密文
// 警告：DES 已被认为是不安全的，仅用于兼容性目的。推荐使用 AES。
// 警告：ECB 模式在密码学上是不安全的，相同的明文块会产生相同的密文块。
func DESDecryptECB(key, ciphertext []byte) ([]byte, error) {
	if len(key) != 8 {
		return nil, errors.New("invalid key length: must be 8 bytes for DES")
	}

	block, err := desNewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext)%des.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	plaintext := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += des.BlockSize {
		block.Decrypt(plaintext[i:i+des.BlockSize], ciphertext[i:i+des.BlockSize])
	}
	return unpadPKCS7(plaintext)
}

// DESEncryptCBC 使用 DES 在 CBC 模式下加密明文
// 警告：DES 已被认为是不安全的，仅用于兼容性目的。推荐使用 AES。
func DESEncryptCBC(key, plaintext []byte) ([]byte, error) {
	if len(key) != 8 {
		return nil, errors.New("invalid key length: must be 8 bytes for DES")
	}

	block, err := desNewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext = padPKCS7(plaintext, des.BlockSize)
	ciphertext := make([]byte, des.BlockSize+len(plaintext))
	iv := ciphertext[:des.BlockSize]
	_, err = io.ReadFull(desRandReader, iv)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[des.BlockSize:], plaintext)
	return ciphertext, nil
}

// DESDecryptCBC 使用 DES 在 CBC 模式下解密密文
// 警告：DES 已被认为是不安全的，仅用于兼容性目的。推荐使用 AES。
func DESDecryptCBC(key, ciphertext []byte) ([]byte, error) {
	if len(key) != 8 {
		return nil, errors.New("invalid key length: must be 8 bytes for DES")
	}

	block, err := desNewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < des.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:des.BlockSize]
	ciphertext = ciphertext[des.BlockSize:]

	if len(ciphertext)%des.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return unpadPKCS7(ciphertext)
}

// TripleDESEncryptECB 使用 3DES 在 ECB 模式下加密明文
// 警告：ECB 模式在密码学上是不安全的，相同的明文块会产生相同的密文块。
func TripleDESEncryptECB(key, plaintext []byte) ([]byte, error) {
	if len(key) != 24 {
		return nil, errors.New("invalid key length: must be 24 bytes for 3DES")
	}

	block, err := desNewTripleDES(key)
	if err != nil {
		return nil, err
	}

	plaintext = padPKCS7(plaintext, des.BlockSize)
	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i += des.BlockSize {
		block.Encrypt(ciphertext[i:i+des.BlockSize], plaintext[i:i+des.BlockSize])
	}
	return ciphertext, nil
}

// TripleDESDecryptECB 使用 3DES 在 ECB 模式下解密密文
// 警告：ECB 模式在密码学上是不安全的，相同的明文块会产生相同的密文块。
func TripleDESDecryptECB(key, ciphertext []byte) ([]byte, error) {
	if len(key) != 24 {
		return nil, errors.New("invalid key length: must be 24 bytes for 3DES")
	}

	block, err := desNewTripleDES(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext)%des.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	plaintext := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += des.BlockSize {
		block.Decrypt(plaintext[i:i+des.BlockSize], ciphertext[i:i+des.BlockSize])
	}
	return unpadPKCS7(plaintext)
}

// TripleDESEncryptCBC 使用 3DES 在 CBC 模式下加密明文
func TripleDESEncryptCBC(key, plaintext []byte) ([]byte, error) {
	if len(key) != 24 {
		return nil, errors.New("invalid key length: must be 24 bytes for 3DES")
	}

	block, err := desNewTripleDES(key)
	if err != nil {
		return nil, err
	}

	plaintext = padPKCS7(plaintext, des.BlockSize)
	ciphertext := make([]byte, des.BlockSize+len(plaintext))
	iv := ciphertext[:des.BlockSize]
	_, err = io.ReadFull(desRandReader, iv)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[des.BlockSize:], plaintext)
	return ciphertext, nil
}

// TripleDESDecryptCBC 使用 3DES 在 CBC 模式下解密密文
func TripleDESDecryptCBC(key, ciphertext []byte) ([]byte, error) {
	if len(key) != 24 {
		return nil, errors.New("invalid key length: must be 24 bytes for 3DES")
	}

	block, err := desNewTripleDES(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < des.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:des.BlockSize]
	ciphertext = ciphertext[des.BlockSize:]

	if len(ciphertext)%des.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return unpadPKCS7(ciphertext)
}
