package cryptox

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/sha3"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
)

// Md5 计算输入字符串或字节切片的 MD5 哈希值，并返回十六进制表示的字符串。
func Md5[M string | []byte](s M) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// Sha256 计算输入字符串或字节切片的 SHA-256 哈希值，并返回十六进制表示的字符串。
func Sha256[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

// Sha224 计算输入字符串或字节切片的 SHA-224 哈希值，并返回十六进制表示的字符串。
func Sha224[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha256.Sum224([]byte(s)))
}

// Sha512 计算输入字符串或字节切片的 SHA-512 哈希值，并返回十六进制表示的字符串。
func Sha512[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(s)))
}

// Sha384 计算输入字符串或字节切片的 SHA-384 哈希值，并返回十六进制表示的字符串。
func Sha384[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha512.Sum384([]byte(s)))
}

// Sha512_256 计算输入字符串或字节切片的 SHA-512/256 哈希值，并返回十六进制表示的字符串。
func Sha512_256[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha512.Sum512_256([]byte(s)))
}

// Sha512_224 计算输入字符串或字节切片的 SHA-512/224 哈希值，并返回十六进制表示的字符串。
func Sha512_224[M string | []byte](s M) string {
	return fmt.Sprintf("%x", sha512.Sum512_224([]byte(s)))
}

// Hash32 使用 FNV-1 算法计算输入字符串或字节切片的 32 位哈希值。
func Hash32[M string | []byte](s M) uint32 {
	h := fnv.New32()
	_, _ = h.Write([]byte(s))
	return h.Sum32()
}

// Hash32a 使用 FNV-1a 算法计算输入字符串或字节切片的 32 位哈希值。
func Hash32a[M string | []byte](s M) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return h.Sum32()
}

// Hash64 使用 FNV-1 算法计算输入字符串或字节切片的 64 位哈希值。
func Hash64[M string | []byte](s M) uint64 {
	h := fnv.New64()
	_, _ = h.Write([]byte(s))
	return h.Sum64()
}

// Hash64a 使用 FNV-1a 算法计算输入字符串或字节切片的 64 位哈希值。
func Hash64a[M string | []byte](s M) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s))
	return h.Sum64()
}

// CRC32 计算输入字符串或字节切片的 CRC32 校验值。
func CRC32[M string | []byte](s M) uint32 {
	return crc32.ChecksumIEEE([]byte(s))
}

// CRC64 计算输入字符串或字节切片的 CRC64 校验值。
func CRC64[M string | []byte](s M) uint64 {
	table := crc64.MakeTable(crc64.ECMA)
	return crc64.Checksum([]byte(s), table)
}

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
	h, err := blake2s.New256(key)
	if err != nil {
		return "", fmt.Errorf("failed to create BLAKE2s hash: %w", err)
	}
	_, _ = h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}