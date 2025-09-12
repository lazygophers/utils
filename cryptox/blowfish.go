package cryptox

import (
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"

	"golang.org/x/crypto/blowfish"
)

// Global variables for dependency injection during testing
var (
	blowfishNewCipher  func([]byte) (*blowfish.Cipher, error) = blowfish.NewCipher
	blowfishRandReader = rand.Reader
)

// BlowfishEncryptECB 使用 Blowfish 在 ECB 模式下加密明文
// 警告：ECB 模式在密码学上是不安全的，相同的明文块会产生相同的密文块。
func BlowfishEncryptECB(key, plaintext []byte) ([]byte, error) {
	if len(key) < 1 || len(key) > 56 {
		return nil, errors.New("invalid key length: must be between 1 and 56 bytes for Blowfish")
	}

	block, err := blowfishNewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext = padPKCS7(plaintext, blowfish.BlockSize)
	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i += blowfish.BlockSize {
		block.Encrypt(ciphertext[i:i+blowfish.BlockSize], plaintext[i:i+blowfish.BlockSize])
	}
	return ciphertext, nil
}

// BlowfishDecryptECB 使用 Blowfish 在 ECB 模式下解密密文
// 警告：ECB 模式在密码学上是不安全的，相同的明文块会产生相同的密文块。
func BlowfishDecryptECB(key, ciphertext []byte) ([]byte, error) {
	if len(key) < 1 || len(key) > 56 {
		return nil, errors.New("invalid key length: must be between 1 and 56 bytes for Blowfish")
	}

	block, err := blowfishNewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext)%blowfish.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	plaintext := make([]byte, len(ciphertext))
	for i := 0; i < len(ciphertext); i += blowfish.BlockSize {
		block.Decrypt(plaintext[i:i+blowfish.BlockSize], ciphertext[i:i+blowfish.BlockSize])
	}
	return unpadPKCS7(plaintext)
}

// BlowfishEncryptCBC 使用 Blowfish 在 CBC 模式下加密明文
func BlowfishEncryptCBC(key, plaintext []byte) ([]byte, error) {
	if len(key) < 1 || len(key) > 56 {
		return nil, errors.New("invalid key length: must be between 1 and 56 bytes for Blowfish")
	}

	block, err := blowfishNewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext = padPKCS7(plaintext, blowfish.BlockSize)
	ciphertext := make([]byte, blowfish.BlockSize+len(plaintext))
	iv := ciphertext[:blowfish.BlockSize]
	_, err = io.ReadFull(blowfishRandReader, iv)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[blowfish.BlockSize:], plaintext)
	return ciphertext, nil
}

// BlowfishDecryptCBC 使用 Blowfish 在 CBC 模式下解密密文
func BlowfishDecryptCBC(key, ciphertext []byte) ([]byte, error) {
	if len(key) < 1 || len(key) > 56 {
		return nil, errors.New("invalid key length: must be between 1 and 56 bytes for Blowfish")
	}

	block, err := blowfishNewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < blowfish.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:blowfish.BlockSize]
	ciphertext = ciphertext[blowfish.BlockSize:]

	if len(ciphertext)%blowfish.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return unpadPKCS7(ciphertext)
}

// BlowfishEncryptCFB 使用 Blowfish 在 CFB 模式下加密明文
func BlowfishEncryptCFB(key, plaintext []byte) ([]byte, error) {
	if len(key) < 1 || len(key) > 56 {
		return nil, errors.New("invalid key length: must be between 1 and 56 bytes for Blowfish")
	}

	block, err := blowfishNewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, blowfish.BlockSize+len(plaintext))
	iv := ciphertext[:blowfish.BlockSize]
	_, err = io.ReadFull(blowfishRandReader, iv)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[blowfish.BlockSize:], plaintext)
	return ciphertext, nil
}

// BlowfishDecryptCFB 使用 Blowfish 在 CFB 模式下解密密文
func BlowfishDecryptCFB(key, ciphertext []byte) ([]byte, error) {
	if len(key) < 1 || len(key) > 56 {
		return nil, errors.New("invalid key length: must be between 1 and 56 bytes for Blowfish")
	}

	block, err := blowfishNewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < blowfish.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:blowfish.BlockSize]
	ciphertext = ciphertext[blowfish.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return ciphertext, nil
}

// BlowfishEncryptOFB 使用 Blowfish 在 OFB 模式下加密明文
func BlowfishEncryptOFB(key, plaintext []byte) ([]byte, error) {
	if len(key) < 1 || len(key) > 56 {
		return nil, errors.New("invalid key length: must be between 1 and 56 bytes for Blowfish")
	}

	block, err := blowfishNewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, blowfish.BlockSize+len(plaintext))
	iv := ciphertext[:blowfish.BlockSize]
	_, err = io.ReadFull(blowfishRandReader, iv)
	if err != nil {
		return nil, err
	}

	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(ciphertext[blowfish.BlockSize:], plaintext)
	return ciphertext, nil
}

// BlowfishDecryptOFB 使用 Blowfish 在 OFB 模式下解密密文
func BlowfishDecryptOFB(key, ciphertext []byte) ([]byte, error) {
	if len(key) < 1 || len(key) > 56 {
		return nil, errors.New("invalid key length: must be between 1 and 56 bytes for Blowfish")
	}

	block, err := blowfishNewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < blowfish.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:blowfish.BlockSize]
	ciphertext = ciphertext[blowfish.BlockSize:]

	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return ciphertext, nil
}