package cryptox

import (
	"fmt"
	"golang.org/x/crypto/sha3"
)

// SHA3_224 计算输入字符串或字节切片的 SHA3-224 哈希值，并返回十六进制表示的字符串。
func SHA3_224[M string | []byte](s M) string {
	h := sha3.New224()
	_, _ = h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA3_256 计算输入字符串或字节切片的 SHA3-256 哈希值，并返回十六进制表示的字符串。
func SHA3_256[M string | []byte](s M) string {
	h := sha3.New256()
	_, _ = h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA3_384 计算输入字符串或字节切片的 SHA3-384 哈希值，并返回十六进制表示的字符串。
func SHA3_384[M string | []byte](s M) string {
	h := sha3.New384()
	_, _ = h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA3_512 计算输入字符串或字节切片的 SHA3-512 哈希值，并返回十六进制表示的字符串。
func SHA3_512[M string | []byte](s M) string {
	h := sha3.New512()
	_, _ = h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHAKE128 计算输入字符串或字节切片的 SHAKE128 哈希值，并返回指定长度的十六进制表示的字符串。
func SHAKE128[M string | []byte](s M, size int) (string, error) {
	if size <= 0 {
		return "", fmt.Errorf("size must be greater than 0")
	}
	h := sha3.NewShake128()
	_, _ = h.Write([]byte(s))
	buf := make([]byte, size)
	_, _ = h.Read(buf)
	return fmt.Sprintf("%x", buf), nil
}

// SHAKE256 计算输入字符串或字节切片的 SHAKE256 哈希值，并返回指定长度的十六进制表示的字符串。
func SHAKE256[M string | []byte](s M, size int) (string, error) {
	if size <= 0 {
		return "", fmt.Errorf("size must be greater than 0")
	}
	h := sha3.NewShake256()
	_, _ = h.Write([]byte(s))
	buf := make([]byte, size)
	_, _ = h.Read(buf)
	return fmt.Sprintf("%x", buf), nil
}

// Keccak256 计算输入字符串或字节切片的 Keccak-256 哈希值，并返回十六进制表示的字符串。
// 注意：这是原始的 Keccak，不是 NIST 标准化的 SHA3。
func Keccak256[M string | []byte](s M) string {
	h := sha3.NewLegacyKeccak256()
	_, _ = h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}