package stringx

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"math/big"
)

// SecureRandBytes 生成指定长度的安全随机字节序列
// 使用 crypto/rand 生成密码学安全的随机数
func SecureRandBytes(n int) ([]byte, error) {
	if n <= 0 {
		return nil, nil
	}
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// SecureRandString 生成指定长度的安全随机字符串（base64编码）
// 适用于生成会话ID、临时密钥等安全敏感场景
func SecureRandString(n int) (string, error) {
	if n <= 0 {
		return "", nil
	}
	// base64编码后长度会增加，所以计算所需的原始字节数
	byteLen := (n*3 + 3) / 4
	b, err := SecureRandBytes(byteLen)
	if err != nil {
		return "", err
	}
	encoded := base64.RawURLEncoding.EncodeToString(b)
	if len(encoded) > n {
		encoded = encoded[:n]
	}
	return encoded, nil
}

// SecureRandHex 生成指定长度的安全随机十六进制字符串
// 适用于生成token、密钥等
func SecureRandHex(n int) (string, error) {
	if n <= 0 {
		return "", nil
	}
	// 每个字节编码为2个十六进制字符
	byteLen := (n + 1) / 2
	b, err := SecureRandBytes(byteLen)
	if err != nil {
		return "", err
	}
	encoded := hex.EncodeToString(b)
	if len(encoded) > n {
		encoded = encoded[:n]
	}
	return encoded, nil
}

// SecureRandLetters 生成指定长度的安全随机字母字符串（大小写混合）
// 使用 crypto/rand 生成，适用于密码、验证码等安全场景
func SecureRandLetters(n int) (string, error) {
	return secureRandStringWithCharset(n, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

// SecureRandLowerLetters 生成指定长度的安全随机小写字母字符串
func SecureRandLowerLetters(n int) (string, error) {
	return secureRandStringWithCharset(n, "abcdefghijklmnopqrstuvwxyz")
}

// SecureRandUpperLetters 生成指定长度的安全随机大写字母字符串
func SecureRandUpperLetters(n int) (string, error) {
	return secureRandStringWithCharset(n, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

// SecureRandNumbers 生成指定长度的安全随机数字字符串
func SecureRandNumbers(n int) (string, error) {
	return secureRandStringWithCharset(n, "0123456789")
}

// SecureRandLetterNumbers 生成指定长度的安全随机字母数字字符串（大小写混合）
func SecureRandLetterNumbers(n int) (string, error) {
	return secureRandStringWithCharset(n, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

// SecureRandLowerLetterNumbers 生成指定长度的安全随机小写字母数字字符串
func SecureRandLowerLetterNumbers(n int) (string, error) {
	return secureRandStringWithCharset(n, "0123456789abcdefghijklmnopqrstuvwxyz")
}

// SecureRandUpperLetterNumbers 生成指定长度的安全随机大写字母数字字符串
func SecureRandUpperLetterNumbers(n int) (string, error) {
	return secureRandStringWithCharset(n, "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

// secureRandStringWithCharset 从指定字符集生成安全随机字符串
func secureRandStringWithCharset(n int, charset string) (string, error) {
	if n <= 0 {
		return "", nil
	}
	if len(charset) == 0 {
		return "", nil
	}

	charsetRunes := []rune(charset)
	charsetLen := big.NewInt(int64(len(charsetRunes)))
	result := make([]rune, n)

	for i := 0; i < n; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		result[i] = charsetRunes[randomIndex.Int64()]
	}

	return string(result), nil
}
