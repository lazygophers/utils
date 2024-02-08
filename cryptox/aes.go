package cryptox

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

func aesEcbEncode(key []byte, data []byte, padding func(data []byte, size int) ([]byte, error)) ([]byte, error) {
	// 选择加密算法
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 填充原文
	data, err = padding(data, block.BlockSize())
	if err != nil {
		return nil, err
	}
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, key)
	// 加密
	crypted := make([]byte, len(data))
	blockMode.CryptBlocks(crypted, data)

	return []byte(hex.EncodeToString(crypted)), nil
}

func aesEcbDecode(key []byte, data []byte, unpadding func(data []byte) ([]byte, error)) ([]byte, error) {
	var err error
	data, err = hex.DecodeString(string(data))
	if err != nil {
		return nil, err
	}

	// 选择解密算法
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 解密模式
	blockMode := cipher.NewCBCDecrypter(block, key)
	// 解密
	origData := make([]byte, len(data))
	blockMode.CryptBlocks(origData, data)
	// 去除填充
	origData, err = unpadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, nil
}

// AesEcbPKCS7Decode AES解密
// 加密模式：ECB
// 填充：PKCS7
func AesEcbPKCS7Decode(key []byte, data []byte) ([]byte, error) {
	return aesEcbDecode(key, data, PKCS7UnPadding)
}

// AesEcbPKCS7Encode AES加密
// 加密模式：ECB
// 填充：PKCS7
func AesEcbPKCS7Encode(key []byte, data []byte) ([]byte, error) {
	return aesEcbEncode(key, data, PKCS7Padding)
}

// AesEcbZeroDecode AES解密
// 加密模式：ECB
// 填充：Zero
func AesEcbZeroDecode(key []byte, data []byte) ([]byte, error) {
	return aesEcbDecode(key, data, ZeroUnPadding)
}

// AesEcbZeroEncode AES加密
// 加密模式：ECB
// 填充：Zero
func AesEcbZeroEncode(key []byte, data []byte) ([]byte, error) {
	return aesEcbEncode(key, data, ZeroPadding)
}

// AesEcbNoDecode AES解密
// 加密模式：ECB
// 填充：No
func AesEcbNoDecode(key []byte, data []byte) ([]byte, error) {
	return aesEcbDecode(key, data, NoUnPadding)
}

// AesEcbNoEncode AES加密
// 加密模式：ECB
// 填充：No
func AesEcbNoEncode(key []byte, data []byte) ([]byte, error) {
	return aesEcbEncode(key, data, NoPadding)
}

// AesEcbPKCS5Decode AES解密
// 加密模式：ECB
// 填充：PKCS5
func AesEcbPKCS5Decode(key []byte, data []byte) ([]byte, error) {
	return aesEcbDecode(key, data, PKCS5UnPadding)
}

// AesEcbPKCS5Encode AES加密
// 加密模式：ECB
// 填充：PKCS5
func AesEcbPKCS5Encode(key []byte, data []byte) ([]byte, error) {
	return aesEcbEncode(key, data, PKCS5Padding)
}

func aesCbcEncode(key []byte, data []byte, padding func(data []byte, size int) ([]byte, error)) ([]byte, error) {
	// 选择加密算法
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 填充原文
	data, err = padding(data, block.BlockSize())
	if err != nil {
		return nil, err
	}
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, key)
	// 加密
	crypted := make([]byte, len(data))
	blockMode.CryptBlocks(crypted, data)
	return []byte(hex.EncodeToString(crypted)), nil
}

func aesCbcDecode(key []byte, data []byte, unpadding func(data []byte) ([]byte, error)) ([]byte, error) {
	var err error
	data, err = hex.DecodeString(string(data))
	if err != nil {
		return nil, err
	}

	// 选择解密算法
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 解密模式
	blockMode := cipher.NewCBCDecrypter(block, key)
	// 解密
	origData := make([]byte, len(data))
	blockMode.CryptBlocks(origData, data)
	// 去除填充
	origData, err = unpadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, nil
}

// AesCbcPKCS7Decode AES解密
// 加密模式：CBC
// 填充：PKCS7
func AesCbcPKCS7Decode(key []byte, data []byte) ([]byte, error) {
	return aesCbcDecode(key, data, PKCS7UnPadding)
}

// AesCbcPKCS7Encode AES加密
// 加密模式：CBC
// 填充：PKCS7
func AesCbcPKCS7Encode(key []byte, data []byte) ([]byte, error) {
	return aesCbcEncode(key, data, PKCS7Padding)
}

// AesCbcZeroDecode AES解密
// 加密模式：CBC
// 填充：Zero
func AesCbcZeroDecode(key []byte, data []byte) ([]byte, error) {
	return aesCbcDecode(key, data, ZeroUnPadding)
}

// AesCbcZeroEncode AES加密
// 加密模式：CBC
// 填充：Zero
func AesCbcZeroEncode(key []byte, data []byte) ([]byte, error) {
	return aesCbcEncode(key, data, ZeroPadding)
}

// AesCbcNoDecode AES解密
// 加密模式：CBC
// 填充：No
func AesCbcNoDecode(key []byte, data []byte) ([]byte, error) {
	return aesCbcDecode(key, data, NoUnPadding)
}

// AesCbcNoEncode AES加密
// 加密模式：CBC
// 填充：No
func AesCbcNoEncode(key []byte, data []byte) ([]byte, error) {
	return aesCbcEncode(key, data, NoPadding)
}

// AesCbcPKCS5Decode AES解密
// 加密模式：CBC
// 填充：PKCS5
func AesCbcPKCS5Decode(key []byte, data []byte) ([]byte, error) {
	return aesCbcDecode(key, data, PKCS5UnPadding)
}

// AesCbcPKCS5Encode AES加密
// 加密模式：CBC
// 填充：PKCS5
func AesCbcPKCS5Encode(key []byte, data []byte) ([]byte, error) {
	return aesCbcEncode(key, data, PKCS5Padding)
}

func aesCfbEncode(key []byte, data []byte, padding func(data []byte, size int) ([]byte, error)) ([]byte, error) {
	// 选择加密算法
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 填充原文
	data, err = padding(data, block.BlockSize())
	if err != nil {
		return nil, err
	}
	// 加密模式
	blockMode := cipher.NewCFBEncrypter(block, key)
	// 加密
	crypted := make([]byte, len(data))
	blockMode.XORKeyStream(crypted, data)
	return []byte(hex.EncodeToString(crypted)), nil
}

func aesCfbDecode(key []byte, data []byte, unpadding func(data []byte) ([]byte, error)) ([]byte, error) {
	var err error
	data, err = hex.DecodeString(string(data))
	if err != nil {
		return nil, err
	}

	// 选择解密算法
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 解密模式
	blockMode := cipher.NewCFBDecrypter(block, key)
	// 解密
	origData := make([]byte, len(data))
	blockMode.XORKeyStream(origData, data)
	// 去除填充
	origData, err = unpadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, nil
}

// AesCfbPKCS7Decode AES解密
// 加密模式：CFB
// 填充：PKCS7
func AesCfbPKCS7Decode(key []byte, data []byte) ([]byte, error) {
	return aesCfbDecode(key, data, PKCS7UnPadding)
}

// AesCfbPKCS7Encode AES加密
// 加密模式：CFB
// 填充：PKCS7
func AesCfbPKCS7Encode(key []byte, data []byte) ([]byte, error) {
	return aesCfbEncode(key, data, PKCS7Padding)
}

// AesCfbZeroDecode AES解密
// 加密模式：CFB
// 填充：Zero
func AesCfbZeroDecode(key []byte, data []byte) ([]byte, error) {
	return aesCfbDecode(key, data, ZeroUnPadding)
}

// AesCfbZeroEncode AES加密
// 加密模式：CFB
// 填充：Zero
func AesCfbZeroEncode(key []byte, data []byte) ([]byte, error) {
	return aesCfbEncode(key, data, ZeroPadding)
}

// AesCfbNoDecode AES解密
// 加密模式：CFB
// 填充：No
func AesCfbNoDecode(key []byte, data []byte) ([]byte, error) {
	return aesCfbDecode(key, data, NoUnPadding)
}

// AesCfbNoEncode AES加密
// 加密模式：CFB
// 填充：No
func AesCfbNoEncode(key []byte, data []byte) ([]byte, error) {
	return aesCfbEncode(key, data, NoPadding)
}

// AesCfbPKCS5Decode AES解密
// 加密模式：CFB
// 填充：PKCS5
func AesCfbPKCS5Decode(key []byte, data []byte) ([]byte, error) {
	return aesCfbDecode(key, data, PKCS5UnPadding)
}

// AesCfbPKCS5Encode AES加密
// 加密模式：CFB
// 填充：PKCS5
func AesCfbPKCS5Encode(key []byte, data []byte) ([]byte, error) {
	return aesCfbEncode(key, data, PKCS5Padding)
}

func aesCtrEncode(key []byte, data []byte, padding func(data []byte, size int) ([]byte, error)) ([]byte, error) {
	// 选择加密算法
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 填充原文
	data, err = padding(data, block.BlockSize())
	if err != nil {
		return nil, err
	}
	// 加密模式
	blockMode := cipher.NewCTR(block, key)
	// 加密
	crypted := make([]byte, len(data))
	blockMode.XORKeyStream(crypted, data)
	return []byte(hex.EncodeToString(crypted)), nil
}

func aesCtrDecode(key []byte, data []byte, unpadding func(data []byte) ([]byte, error)) ([]byte, error) {
	var err error
	data, err = hex.DecodeString(string(data))
	if err != nil {
		return nil, err
	}

	// 选择解密算法
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 解密模式
	blockMode := cipher.NewCTR(block, key)
	// 解密
	origData := make([]byte, len(data))
	blockMode.XORKeyStream(origData, data)
	// 去除填充
	origData, err = unpadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, nil
}

// AesCtrPKCS7Decode AES解密
// 加密模式：CTR
// 填充：PKCS7
func AesCtrPKCS7Decode(key []byte, data []byte) ([]byte, error) {
	return aesCtrDecode(key, data, PKCS7UnPadding)
}

// AesCtrPKCS7Encode AES加密
// 加密模式：CTR
// 填充：PKCS7
func AesCtrPKCS7Encode(key []byte, data []byte) ([]byte, error) {
	return aesCtrEncode(key, data, PKCS7Padding)
}

// AesCtrZeroDecode AES解密
// 加密模式：CTR
// 填充：Zero
func AesCtrZeroDecode(key []byte, data []byte) ([]byte, error) {
	return aesCtrDecode(key, data, ZeroUnPadding)
}

// AesCtrZeroEncode AES加密
// 加密模式：CTR
// 填充：Zero
func AesCtrZeroEncode(key []byte, data []byte) ([]byte, error) {
	return aesCtrEncode(key, data, ZeroPadding)
}

// AesCtrNoDecode AES解密
// 加密模式：CTR
// 填充：No
func AesCtrNoDecode(key []byte, data []byte) ([]byte, error) {
	return aesCtrDecode(key, data, NoUnPadding)
}

// AesCtrNoEncode AES加密
// 加密模式：CTR
// 填充：No
func AesCtrNoEncode(key []byte, data []byte) ([]byte, error) {
	return aesCtrEncode(key, data, NoPadding)
}

// AesCtrPKCS5Decode AES解密
// 加密模式：CTR
// 填充：PKCS5
func AesCtrPKCS5Decode(key []byte, data []byte) ([]byte, error) {
	return aesCtrDecode(key, data, PKCS5UnPadding)
}

// AesCtrPKCS5Encode AES加密
// 加密模式：CTR
// 填充：PKCS5
func AesCtrPKCS5Encode(key []byte, data []byte) ([]byte, error) {
	return aesCtrEncode(key, data, PKCS5Padding)
}

func aesOfbEncode(key []byte, data []byte, padding func(data []byte, size int) ([]byte, error)) ([]byte, error) {
	// 选择加密算法
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 填充原文
	data, err = padding(data, block.BlockSize())
	if err != nil {
		return nil, err
	}
	// 加密模式
	blockMode := cipher.NewOFB(block, key)
	// 加密
	crypted := make([]byte, len(data))
	blockMode.XORKeyStream(crypted, data)
	return []byte(hex.EncodeToString(crypted)), nil
}

func aesOfbDecode(key []byte, data []byte, unpadding func(data []byte) ([]byte, error)) ([]byte, error) {
	var err error
	data, err = hex.DecodeString(string(data))
	if err != nil {
		return nil, err
	}

	// 选择解密算法
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 解密模式
	blockMode := cipher.NewOFB(block, key)
	// 解密
	origData := make([]byte, len(data))
	blockMode.XORKeyStream(origData, data)
	// 去除填充
	origData, err = unpadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, nil
}

// AesOfbPKCS7Decode AES解密
// 加密模式：OFB
// 填充：PKCS7
func AesOfbPKCS7Decode(key []byte, data []byte) ([]byte, error) {
	return aesOfbDecode(key, data, PKCS7UnPadding)
}

// AesOfbPKCS7Encode AES加密
// 加密模式：OFB
// 填充：PKCS7
func AesOfbPKCS7Encode(key []byte, data []byte) ([]byte, error) {
	return aesOfbEncode(key, data, PKCS7Padding)
}

// AesOfbZeroDecode AES解密
// 加密模式：OFB
// 填充：Zero
func AesOfbZeroDecode(key []byte, data []byte) ([]byte, error) {
	return aesOfbDecode(key, data, ZeroUnPadding)
}

// AesOfbZeroEncode AES加密
// 加密模式：OFB
// 填充：Zero
func AesOfbZeroEncode(key []byte, data []byte) ([]byte, error) {
	return aesOfbEncode(key, data, ZeroPadding)
}

// AesOfbNoDecode AES解密
// 加密模式：OFB
// 填充：No
func AesOfbNoDecode(key []byte, data []byte) ([]byte, error) {
	return aesOfbDecode(key, data, NoUnPadding)
}

// AesOfbNoEncode AES加密
// 加密模式：OFB
// 填充：No
func AesOfbNoEncode(key []byte, data []byte) ([]byte, error) {
	return aesOfbEncode(key, data, NoPadding)
}

// AesOfbPKCS5Decode AES解密
// 加密模式：OFB
// 填充：PKCS5
func AesOfbPKCS5Decode(key []byte, data []byte) ([]byte, error) {
	return aesOfbDecode(key, data, PKCS5UnPadding)
}

// AesOfbPKCS5Encode AES加密
// 加密模式：OFB
// 填充：PKCS5
func AesOfbPKCS5Encode(key []byte, data []byte) ([]byte, error) {
	return aesOfbEncode(key, data, PKCS5Padding)
}

func aesGcmEncode(key []byte, data []byte, padding func(data []byte, size int) ([]byte, error)) ([]byte, error) {
	// 选择加密算法
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 填充原文
	data, err = padding(data, block.BlockSize())
	if err != nil {
		return nil, err
	}
	// 加密模式
	blockMode, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	// 加密
	nonce := make([]byte, blockMode.NonceSize())
	crypted := blockMode.Seal(nil, nonce, data, nil)
	return []byte(hex.EncodeToString(crypted)), nil
}

func aesGcmDecode(key []byte, data []byte, unpadding func(data []byte) ([]byte, error)) ([]byte, error) {
	var err error
	data, err = hex.DecodeString(string(data))
	if err != nil {
		return nil, err
	}

	// 选择解密算法
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 解密模式
	blockMode, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	// 解密
	nonceSize := blockMode.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	origData, err := blockMode.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}
	// 去除填充
	origData, err = unpadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, nil
}

// AesGcmPKCS7Decode AES解密
// 加密模式：GCM
// 填充：PKCS7
func AesGcmPKCS7Decode(key []byte, data []byte) ([]byte, error) {
	return aesGcmDecode(key, data, PKCS7UnPadding)
}

// AesGcmPKCS7Encode AES加密
// 加密模式：GCM
// 填充：PKCS7
func AesGcmPKCS7Encode(key []byte, data []byte) ([]byte, error) {
	return aesGcmEncode(key, data, PKCS7Padding)
}

// AesGcmZeroDecode AES解密
// 加密模式：GCM
// 填充：Zero
func AesGcmZeroDecode(key []byte, data []byte) ([]byte, error) {
	return aesGcmDecode(key, data, ZeroUnPadding)
}

// AesGcmZeroEncode AES加密
// 加密模式：GCM
// 填充：Zero
func AesGcmZeroEncode(key []byte, data []byte) ([]byte, error) {
	return aesGcmEncode(key, data, ZeroPadding)
}

// AesGcmNoDecode AES解密
// 加密模式：GCM
// 填充：No
func AesGcmNoDecode(key []byte, data []byte) ([]byte, error) {
	return aesGcmDecode(key, data, NoUnPadding)
}

// AesGcmNoEncode AES加密
// 加密模式：GCM
// 填充：No
func AesGcmNoEncode(key []byte, data []byte) ([]byte, error) {
	return aesGcmEncode(key, data, NoPadding)
}

// AesGcmPKCS5Decode AES解密
// 加密模式：GCM
// 填充：PKCS5
func AesGcmPKCS5Decode(key []byte, data []byte) ([]byte, error) {
	return aesGcmDecode(key, data, PKCS5UnPadding)
}

// AesGcmPKCS5Encode AES加密
// 加密模式：GCM
// 填充：PKCS5
func AesGcmPKCS5Encode(key []byte, data []byte) ([]byte, error) {
	return aesGcmEncode(key, data, PKCS5Padding)
}
