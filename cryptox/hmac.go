package cryptox

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
)

func hmacEncode(key, data string, h func() hash.Hash) string {
	sha := hmac.New(h, []byte(key))
	sha.Write([]byte(data))
	return string(sha.Sum(nil))
}

func HmacMd5(key, data string) string {
	return hmacEncode(key, data, md5.New)
}

func HmacSha256(key, data string) string {
	return hmacEncode(key, data, sha256.New)
}

func HmacSha224(key, data string) string {
	return hmacEncode(key, data, sha256.New224)
}

func HmacSha512(key, data string) string {
	return hmacEncode(key, data, sha512.New)
}

func HmacSha384(key, data string) string {
	return hmacEncode(key, data, sha512.New384)
}
