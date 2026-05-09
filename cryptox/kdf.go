package cryptox

import (
	"crypto/sha256"
	"crypto/sha512"
	"hash"

	"golang.org/x/crypto/pbkdf2"
)

// PBKDF2SHA256 使用 PBKDF2-HMAC-SHA256 从密码派生密钥
//
// 参数：
//   - password: 原始密码或密钥材料
//   - salt: 盐值，建议至少 16 字节，应使用 crypto/rand 生成
//   - iterations: 迭代次数，建议至少 100000，越高越安全但越慢
//   - keyLen: 派生密钥的长度（字节）
//
// 使用场景：
//   - 密码存储（配合随机盐）
//   - 密钥派生（从用户密码生成加密密钥）
//   - 密钥扩展（从短密钥生成长密钥）
//
// 安全建议：
//   - 盐值必须唯一且随机，每个密码使用不同的盐
//   - 迭代次数建议：桌面应用 ≥ 100,000，Web应用 ≥ 50,000
//   - 对于更高安全需求，考虑使用 Argon2id（需要外部依赖）
func PBKDF2SHA256(password, salt []byte, iterations, keyLen int) []byte {
	return pbkdf2.Key(password, salt, iterations, keyLen, sha256.New)
}

// PBKDF2SHA512 使用 PBKDF2-HMAC-SHA512 从密码派生密钥
//
// 参数与 PBKDF2SHA256 相同，但使用 SHA-512 作为底层哈希函数
// SHA-512 比 SHA-256 更慢，但在某些平台上可能更安全
func PBKDF2SHA512(password, salt []byte, iterations, keyLen int) []byte {
	return pbkdf2.Key(password, salt, iterations, keyLen, sha512.New)
}

// DeriveKey 通用的密钥派生函数，支持自定义哈希算法
//
// 参数：
//   - password: 原始密码或密钥材料
//   - salt: 盐值
//   - iterations: 迭代次数
//   - keyLen: 派生密钥的长度
//   - hashFunc: 哈希函数构造器（如 sha256.New）
//
// 示例：
//
//	key := DeriveKey([]byte("password"), salt, 100000, 32, sha256.New)
func DeriveKey(password, salt []byte, iterations, keyLen int, hashFunc func() hash.Hash) []byte {
	return pbkdf2.Key(password, salt, iterations, keyLen, hashFunc)
}

// PBKDF2Config PBKDF2 配置参数
type PBKDF2Config struct {
	Iterations int              // 迭代次数，建议 ≥ 100000
	SaltLen    int              // 盐值长度，建议 ≥ 16
	KeyLen     int              // 派生密钥长度
	HashFunc   func() hash.Hash // 哈希函数，默认 sha256.New
}

// DefaultPBKDF2Config 返回推荐的 PBKDF2 配置
//
// 配置说明：
//   - 迭代次数：100,000（适合桌面应用）
//   - 盐值长度：32 字节
//   - 密钥长度：32 字节（适用于 AES-256）
//   - 哈希函数：SHA-256
func DefaultPBKDF2Config() PBKDF2Config {
	return PBKDF2Config{
		Iterations: 100000,
		SaltLen:    32,
		KeyLen:     32,
		HashFunc:   sha256.New,
	}
}

// FastPBKDF2Config 返回快速但安全性较低的 PBKDF2 配置
//
// 配置说明：
//   - 迭代次数：50,000（适合 Web 应用，减少服务器负载）
//   - 盐值长度：16 字节
//   - 密钥长度：32 字节
//   - 哈希函数：SHA-256
//
// 警告：仅在性能要求高且威胁模型允许的情况下使用
func FastPBKDF2Config() PBKDF2Config {
	return PBKDF2Config{
		Iterations: 50000,
		SaltLen:    16,
		KeyLen:     32,
		HashFunc:   sha256.New,
	}
}

// StrongPBKDF2Config 返回高安全性的 PBKDF2 配置
//
// 配置说明：
//   - 迭代次数：200,000（更高安全性）
//   - 盐值长度：64 字节
//   - 密钥长度：64 字节
//   - 哈希函数：SHA-512
//
// 适用于高安全需求场景，但性能开销较大
func StrongPBKDF2Config() PBKDF2Config {
	return PBKDF2Config{
		Iterations: 200000,
		SaltLen:    64,
		KeyLen:     64,
		HashFunc:   sha512.New,
	}
}
