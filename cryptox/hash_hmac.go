package cryptox

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

// HMACMd5 使用 MD5 作为底层哈希函数计算 HMAC 值，并返回十六进制表示的字符串。
//
// ⚠️ SECURITY WARNING: MD5 is cryptographically broken.
// - DO NOT use for any security purpose (authentication, tokens, API signatures)
// - Use HMACSHA256 instead
//
// Deprecated: Use HMACSHA256 for security-critical applications.
func HMACMd5[M string | []byte](key, message M) string {
	h := hmac.New(md5.New, []byte(key))
	_, _ = h.Write([]byte(message))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// HMACSHA1 使用 SHA1 作为底层哈希函数计算 HMAC 值，并返回十六进制表示的字符串。
//
// ⚠️ SECURITY WARNING: SHA1 is cryptographically broken (SHAttered attack).
// - DO NOT use for API signatures, JWT, or authentication tokens
// - Use HMACSHA256 instead
//
// Deprecated: Use HMACSHA256 for security-critical applications.
func HMACSHA1[M string | []byte](key, message M) string {
	h := hmac.New(sha1.New, []byte(key))
	_, _ = h.Write([]byte(message))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// HMACSHA256 使用 SHA256 作为底层哈希函数计算 HMAC 值，并返回十六进制表示的字符串。
func HMACSHA256[M string | []byte](key, message M) string {
	h := hmac.New(sha256.New, []byte(key))
	_, _ = h.Write([]byte(message))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// HMACSHA384 使用 SHA384 作为底层哈希函数计算 HMAC 值，并返回十六进制表示的字符串。
func HMACSHA384[M string | []byte](key, message M) string {
	h := hmac.New(sha512.New384, []byte(key))
	_, _ = h.Write([]byte(message))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// HMACSHA512 使用 SHA512 作为底层哈希函数计算 HMAC 值，并返回十六进制表示的字符串。
func HMACSHA512[M string | []byte](key, message M) string {
	h := hmac.New(sha512.New, []byte(key))
	_, _ = h.Write([]byte(message))
	return fmt.Sprintf("%x", h.Sum(nil))
}
