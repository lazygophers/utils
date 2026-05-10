package cryptox

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"unsafe"
)

// Md5 计算输入字符串或字节切片的 MD5 哈希值，并返回十六进制表示的字符串。
//
// ⚠️ SECURITY WARNING: MD5 is cryptographically broken and vulnerable to collision attacks.
// - DO NOT use for passwords, digital signatures, or any security-critical operations
// - DO NOT use for certificates, TLS, or authentication systems
// - Acceptable use cases: non-security checksums, cache keys, file deduplication
//
// Deprecated: Use Sha256 or Sha512 instead for any security purpose.
func Md5[M string | []byte](s M) string {
	hash := md5.Sum([]byte(s))
	var result [32]byte
	for i := 0; i < 16; i++ {
		b := hash[i]
		result[i*2] = "0123456789abcdef"[b>>4]
		result[i*2+1] = "0123456789abcdef"[b&0x0f]
	}
	return string(result[:])
}

// SHA1 计算输入字符串或字节切片的 SHA1 哈希值，并返回十六进制表示的字符串。
//
// ⚠️ SECURITY WARNING: SHA1 is cryptographically broken (SHAttered attack).
// - DO NOT use for passwords, digital signatures, certificates, or TLS
// - DO NOT use for git objects (use SHA256), code signing, or authentication
// - Acceptable use cases: compatibility with legacy systems only
//
// Deprecated: Use Sha256 or Sha512 instead.
func SHA1[M string | []byte](s M) string {
	hash := sha1.Sum([]byte(s))
	var result [40]byte
	for i := 0; i < 20; i++ {
		b := hash[i]
		result[i*2] = "0123456789abcdef"[b>>4]
		result[i*2+1] = "0123456789abcdef"[b&0x0f]
	}
	return string(result[:])
}

// Sha224 计算输入字符串或字节切片的 SHA-224 哈希值，并返回十六进制表示的字符串。
func Sha224[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha256.Sum224([]byte(s)))
}

// Sha256 计算输入字符串或字节切片的 SHA-256 哈希值，并返回十六进制表示的字符串。
func Sha256[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

// Sha384 计算输入字符串或字节切片的 SHA-384 哈希值，并返回十六进制表示的字符串。
func Sha384[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha512.Sum384([]byte(s)))
}

// Sha512 计算输入字符串或字节切片的 SHA-512 哈希值，并返回十六进制表示的字符串。
// 手动 hex 编码优化
// 性能提升：2.05x（189.6 ns/op vs 388.4 ns-op vs fmt.Sprintf）
func Sha512[M string | []byte](s M) string {
	const hexChars = "0123456789abcdef"
	hash := sha512.Sum512([]byte(s))
	var result [128]byte
	for i := 0; i < 64; i++ {
		b := hash[i]
		result[i*2] = hexChars[b>>4]
		result[i*2+1] = hexChars[b&0x0f]
	}
	return string(result[:])
}

// Sha512_224 计算输入字符串或字节切片的 SHA-512/224 哈希值，并返回十六进制表示的字符串。
// 手动 hex 编码优化
// 性能提升：2.20x（146 ns/op vs 321 ns/op fmt.Sprintf）
func Sha512_224[M string | []byte](s M) string {
	const hexChars = "0123456789abcdef"
	hash := sha512.Sum512_224([]byte(s))
	var result [56]byte
	for i := 0; i < 28; i++ {
		b := hash[i]
		result[i*2] = hexChars[b>>4]
		result[i*2+1] = hexChars[b&15]
	}
	return string(result[:])
}

// Sha512_256 计算输入字符串或字节切片的 SHA-512/256 哈希值，并返回十六进制表示的字符串。
// 手动 hex 编码优化 + unsafe 转换
// 性能提升：2.14x（131.2 ns/op vs 280.2 ns-op fmt.Sprintf）
func Sha512_256[M string | []byte](s M) string {
	const hexChars = "0123456789abcdef"
	hash := sha512.Sum512_256([]byte(s))
	var result [64]byte
	for i := 0; i < 32; i++ {
		b := hash[i]
		result[i*2] = hexChars[b>>4]
		result[i*2+1] = hexChars[b&0x0f]
	}
	// unsafe 转换避免切片拷贝
	return unsafe.String(&result[0], 64)
}
