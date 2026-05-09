package cryptox

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
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
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
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
	return fmt.Sprintf("%x", sha1.Sum([]byte(s)))
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
func Sha512[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(s)))
}

// Sha512_224 计算输入字符串或字节切片的 SHA-512/224 哈希值，并返回十六进制表示的字符串。
func Sha512_224[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha512.Sum512_224([]byte(s)))
}

// Sha512_256 计算输入字符串或字节切片的 SHA-512/256 哈希值，并返回十六进制表示的字符串。
func Sha512_256[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha512.Sum512_256([]byte(s)))
}
