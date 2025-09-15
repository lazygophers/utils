package cryptox

import (
	"fmt"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
)

// Global variables for dependency injection during testing
var (
	blake2bNew    = blake2b.New
	blake2sNew256 = blake2s.New256
)

// BLAKE2b 计算输入字符串或字节切片的 BLAKE2b 哈希值，并返回指定长度的十六进制表示的字符串。
func BLAKE2b[M string | []byte](s M, size int) (string, error) {
	if size <= 0 {
		return "", fmt.Errorf("size must be greater than 0")
	}
	var key []byte
	h, err := blake2b.New(size, key)
	if err != nil {
		return "", fmt.Errorf("failed to create BLAKE2b hash: %w", err)
	}
	_, _ = h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// BLAKE2s 计算输入字符串或字节切片的 BLAKE2s 哈希值，并返回256位十六进制表示的字符串。
// 注意：BLAKE2s 固定输出256位，size参数被忽略以保持兼容性。
func BLAKE2s[M string | []byte](s M, size int) (string, error) {
	if size <= 0 {
		return "", fmt.Errorf("size must be greater than 0")
	}
	var key []byte
	h, err := blake2sNew256(key)
	if err != nil {
		return "", fmt.Errorf("failed to create BLAKE2s hash: %w", err)
	}
	_, _ = h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// BLAKE2b512 计算输入字符串或字节切片的 BLAKE2b-512 哈希值，并返回十六进制表示的字符串。
func BLAKE2b512[M string | []byte](s M) string {
	h, _ := blake2b.New512(nil)
	_, _ = h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// BLAKE2b256 计算输入字符串或字节切片的 BLAKE2b-256 哈希值，并返回十六进制表示的字符串。
func BLAKE2b256[M string | []byte](s M) string {
	h, _ := blake2b.New256(nil)
	_, _ = h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// BLAKE2s256 计算输入字符串或字节切片的 BLAKE2s-256 哈希值，并返回十六进制表示的字符串。
func BLAKE2s256[M string | []byte](s M) string {
	h, _ := blake2s.New256(nil)
	_, _ = h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// BLAKE2bWithKey 使用密钥计算输入字符串或字节切片的 BLAKE2b 哈希值，并返回指定长度的十六进制表示的字符串。
func BLAKE2bWithKey[M string | []byte](s M, key []byte, size int) (string, error) {
	if size <= 0 || size > 64 {
		return "", fmt.Errorf("size must be between 1 and 64 bytes")
	}
	h, err := blake2bNew(size, key)
	if err != nil {
		return "", fmt.Errorf("failed to create BLAKE2b hash with key: %w", err)
	}
	_, _ = h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// BLAKE2sWithKey 使用密钥计算输入字符串或字节切片的 BLAKE2s 哈希值，并返回256位十六进制表示的字符串。
func BLAKE2sWithKey[M string | []byte](s M, key []byte) (string, error) {
	h, err := blake2sNew256(key)
	if err != nil {
		return "", fmt.Errorf("failed to create BLAKE2s hash with key: %w", err)
	}
	_, _ = h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
