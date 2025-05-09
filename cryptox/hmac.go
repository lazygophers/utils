package cryptox

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"golang.org/x/crypto/sha3"
	"hash"
)

// hmacEncode 是一个通用的 HMAC 编码函数，支持传入不同的哈希函数
func hmacEncode(key, data string, h func() hash.Hash) string {
	sha := hmac.New(h, []byte(key))
	sha.Write([]byte(data))
	return string(sha.Sum(nil))
}

// HmacMd5 使用 MD5 哈希算法进行 HMAC 编码
func HmacMd5(key, data string) string {
	return hmacEncode(key, data, md5.New)
}

// HmacSha1 使用 SHA-1 哈希算法进行 HMAC 编码
func HmacSha1(key, data string) string {
	return hmacEncode(key, data, sha1.New)
}

// HmacSha256 使用 SHA-256 哈希算法进行 HMAC 编码
func HmacSha256(key, data string) string {
	return hmacEncode(key, data, sha256.New)
}

// HmacSha224 使用 SHA-224 哈希算法进行 HMAC 编码
func HmacSha224(key, data string) string {
	return hmacEncode(key, data, sha256.New224)
}

// HmacSha512 使用 SHA-512 哈希算法进行 HMAC 编码
func HmacSha512(key, data string) string {
	return hmacEncode(key, data, sha512.New)
}

// HmacSha384 使用 SHA-384 哈希算法进行 HMAC 编码
func HmacSha384(key, data string) string {
	return hmacEncode(key, data, sha512.New384)
}

// HmacSha3_256 使用 SHA3-256 哈希算法进行 HMAC 编码
func HmacSha3_256(key, data string) string {
	return hmacEncode(key, data, sha3.New256)
}

// HmacSha3_384 使用 SHA3-384 哈希算法进行 HMAC 编码
func HmacSha3_384(key, data string) string {
	return hmacEncode(key, data, sha3.New384)
}

// HmacSha3_512 使用 SHA3-512 哈希算法进行 HMAC 编码
func HmacSha3_512(key, data string) string {
	return hmacEncode(key, data, sha3.New512)
}
