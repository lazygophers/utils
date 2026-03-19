package stringx

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"math/bits"
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
// 优化版本：批量读取随机字节，减少 crypto/rand 调用次数
func secureRandStringWithCharset(n int, charset string) (string, error) {
	if n <= 0 {
		return "", nil
	}
	if len(charset) == 0 {
		return "", nil
	}

	charsetLen := len(charset)
	result := make([]byte, n)

	// 计算需要多少随机字节来覆盖 n 个字符
	// 每个随机字节可以提供 8 位熵，但我们需要保证均匀分布
	// 使用拒绝采样来避免偏差

	// 对于小字符集（<=256），可以直接使用字节索引
	if charsetLen <= 256 {
		// 计算需要多少随机字节
		// 使用乘法避免浮点运算
		bytesNeeded := n
		if bytesNeeded < 32 {
			bytesNeeded = 32 // 最小读取量以减少系统调用
		}

		// 批量读取随机字节
		randomBytes := make([]byte, bytesNeeded)
		if _, err := rand.Read(randomBytes); err != nil {
			return "", err
		}

		// 使用拒绝采样确保均匀分布
		// 计算大于charsetLen的最小2的幂
		threshold := (256 / charsetLen) * charsetLen

		pos := 0
		bytePos := 0
		for pos < n {
			if bytePos >= len(randomBytes) {
				// 需要更多随机字节
				if _, err := rand.Read(randomBytes); err != nil {
					return "", err
				}
				bytePos = 0
			}

			b := randomBytes[bytePos]
			bytePos++

			// 拒绝采样：只接受小于threshold的值
			if int(b) < threshold {
				result[pos] = charset[int(b)%charsetLen]
				pos++
			}
			// 如果b >= threshold，丢弃这个字节继续循环
		}

		return string(result), nil
	}

	// 对于大字符集（>256），回退到原始方法
	// 但使用批量读取来优化
	buf := make([]byte, n*4) // 每个字符需要多个字节
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}

	charsetRunes := []rune(charset)
	resultRunes := make([]rune, n)

	bitPos := 0
	bufPos := 0
	for i := 0; i < n; i++ {
		// 从buf中提取足够的位数
		var index uint64
		bitsNeeded := bits.Len64(uint64(charsetLen - 1))

		for bitsNeeded > 0 {
			if bufPos >= len(buf) {
				if _, err := rand.Read(buf); err != nil {
					return "", err
				}
				bufPos = 0
			}

			bitsAvailable := 8 - bitPos
			bitsToTake := bitsNeeded
			if bitsToTake > bitsAvailable {
				bitsToTake = bitsAvailable
			}

			index = (index << uint(bitsToTake)) | uint64((buf[bufPos]>>bitPos)&((1<<bitsToTake)-1))
			bitPos += bitsToTake
			bitsNeeded -= bitsToTake

			if bitPos >= 8 {
				bitPos = 0
				bufPos++
			}
		}

		resultRunes[i] = charsetRunes[index%uint64(charsetLen)]
	}

	return string(resultRunes), nil
}
