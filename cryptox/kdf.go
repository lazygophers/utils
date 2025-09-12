package cryptox

import (
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"io"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
)

// Global variables for dependency injection during testing
var (
	kdfRandReader = rand.Reader
)

// PBKDF2WithSHA256 使用 PBKDF2 和 SHA256 从密码派生密钥
func PBKDF2WithSHA256(password, salt []byte, iterations, keyLength int) []byte {
	return pbkdf2.Key(password, salt, iterations, keyLength, sha256.New)
}

// PBKDF2WithSHA1 使用 PBKDF2 和 SHA1 从密码派生密钥
// 注意：SHA1 已被认为不安全，仅用于兼容性目的
func PBKDF2WithSHA1(password, salt []byte, iterations, keyLength int) []byte {
	return pbkdf2.Key(password, salt, iterations, keyLength, sha1.New)
}

// PBKDF2WithSHA512 使用 PBKDF2 和 SHA512 从密码派生密钥
func PBKDF2WithSHA512(password, salt []byte, iterations, keyLength int) []byte {
	return pbkdf2.Key(password, salt, iterations, keyLength, sha512.New)
}

// PBKDF2Config PBKDF2 配置参数
type PBKDF2Config struct {
	SaltLength int // 盐长度（字节）
	Iterations int // 迭代次数
	KeyLength  int // 密钥长度（字节）
}

// DefaultPBKDF2Config 返回推荐的 PBKDF2 配置
func DefaultPBKDF2Config() PBKDF2Config {
	return PBKDF2Config{
		SaltLength: 16,     // 128位盐
		Iterations: 100000, // 10万次迭代
		KeyLength:  32,     // 256位密钥
	}
}

// PBKDF2Generate 生成随机盐并使用 PBKDF2-SHA256 派生密钥
func PBKDF2Generate(password string, config PBKDF2Config) (key, salt []byte, err error) {
	if config.SaltLength <= 0 {
		return nil, nil, errors.New("salt length must be greater than 0")
	}
	if config.Iterations <= 0 {
		return nil, nil, errors.New("iterations must be greater than 0")
	}
	if config.KeyLength <= 0 {
		return nil, nil, errors.New("key length must be greater than 0")
	}

	salt = make([]byte, config.SaltLength)
	_, err = io.ReadFull(kdfRandReader, salt)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate salt: %w", err)
	}

	key = PBKDF2WithSHA256([]byte(password), salt, config.Iterations, config.KeyLength)
	return key, salt, nil
}

// PBKDF2Verify 验证密码是否匹配给定的密钥和盐
func PBKDF2Verify(password string, key, salt []byte, config PBKDF2Config) bool {
	derivedKey := PBKDF2WithSHA256([]byte(password), salt, config.Iterations, config.KeyLength)
	return len(key) == len(derivedKey) && constantTimeCompare(key, derivedKey)
}

// ScryptConfig Scrypt 配置参数
type ScryptConfig struct {
	SaltLength int // 盐长度（字节）
	N          int // CPU/内存成本参数（必须是2的幂）
	R          int // 块大小参数
	P          int // 并行化参数
	KeyLength  int // 密钥长度（字节）
}

// DefaultScryptConfig 返回推荐的 Scrypt 配置
func DefaultScryptConfig() ScryptConfig {
	return ScryptConfig{
		SaltLength: 16,    // 128位盐
		N:          32768, // 2^15, 推荐值
		R:          8,     // 推荐值
		P:          1,     // 推荐值
		KeyLength:  32,    // 256位密钥
	}
}

// ScryptDerive 使用 Scrypt 从密码派生密钥
func ScryptDerive(password, salt []byte, config ScryptConfig) ([]byte, error) {
	if config.N <= 0 || (config.N&(config.N-1)) != 0 {
		return nil, errors.New("N must be a positive power of 2")
	}
	if config.R <= 0 {
		return nil, errors.New("R must be greater than 0")
	}
	if config.P <= 0 {
		return nil, errors.New("P must be greater than 0")
	}
	if config.KeyLength <= 0 {
		return nil, errors.New("key length must be greater than 0")
	}

	return scrypt.Key(password, salt, config.N, config.R, config.P, config.KeyLength)
}

// ScryptGenerate 生成随机盐并使用 Scrypt 派生密钥
func ScryptGenerate(password string, config ScryptConfig) (key, salt []byte, err error) {
	if config.SaltLength <= 0 {
		return nil, nil, errors.New("salt length must be greater than 0")
	}

	salt = make([]byte, config.SaltLength)
	_, err = io.ReadFull(kdfRandReader, salt)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate salt: %w", err)
	}

	key, err = ScryptDerive([]byte(password), salt, config)
	if err != nil {
		return nil, nil, fmt.Errorf("scrypt derivation failed: %w", err)
	}

	return key, salt, nil
}

// ScryptVerify 验证密码是否匹配给定的密钥和盐
func ScryptVerify(password string, key, salt []byte, config ScryptConfig) bool {
	derivedKey, err := ScryptDerive([]byte(password), salt, config)
	if err != nil {
		return false
	}
	return len(key) == len(derivedKey) && constantTimeCompare(key, derivedKey)
}

// Argon2Config Argon2 配置参数
type Argon2Config struct {
	SaltLength int    // 盐长度（字节）
	Time       uint32 // 时间参数（迭代次数）
	Memory     uint32 // 内存参数（KB）
	Threads    uint8  // 并行度参数
	KeyLength  int    // 密钥长度（字节）
}

// DefaultArgon2Config 返回推荐的 Argon2 配置
func DefaultArgon2Config() Argon2Config {
	return Argon2Config{
		SaltLength: 16,      // 128位盐
		Time:       1,       // 1次迭代（推荐值）
		Memory:     64 * 1024, // 64MB 内存
		Threads:    4,       // 4线程
		KeyLength:  32,      // 256位密钥
	}
}

// Argon2IDDerive 使用 Argon2id 从密码派生密钥
func Argon2IDDerive(password, salt []byte, config Argon2Config) []byte {
	return argon2.IDKey(password, salt, config.Time, config.Memory, config.Threads, uint32(config.KeyLength))
}

// Argon2IDerive 使用 Argon2i 从密码派生密钥
func Argon2IDerive(password, salt []byte, config Argon2Config) []byte {
	return argon2.Key(password, salt, config.Time, config.Memory, config.Threads, uint32(config.KeyLength))
}

// Argon2Generate 生成随机盐并使用 Argon2id 派生密钥
func Argon2Generate(password string, config Argon2Config) (key, salt []byte, err error) {
	if config.SaltLength <= 0 {
		return nil, nil, errors.New("salt length must be greater than 0")
	}
	if config.Time == 0 {
		return nil, nil, errors.New("time parameter must be greater than 0")
	}
	if config.Memory == 0 {
		return nil, nil, errors.New("memory parameter must be greater than 0")
	}
	if config.Threads == 0 {
		return nil, nil, errors.New("threads parameter must be greater than 0")
	}
	if config.KeyLength <= 0 {
		return nil, nil, errors.New("key length must be greater than 0")
	}

	salt = make([]byte, config.SaltLength)
	_, err = io.ReadFull(kdfRandReader, salt)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate salt: %w", err)
	}

	key = Argon2IDDerive([]byte(password), salt, config)
	return key, salt, nil
}

// Argon2Verify 验证密码是否匹配给定的密钥和盐
func Argon2Verify(password string, key, salt []byte, config Argon2Config) bool {
	derivedKey := Argon2IDDerive([]byte(password), salt, config)
	return len(key) == len(derivedKey) && constantTimeCompare(key, derivedKey)
}

// constantTimeCompare 恒定时间比较，防止时间攻击
func constantTimeCompare(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	
	var result byte
	for i := 0; i < len(a); i++ {
		result |= a[i] ^ b[i]
	}
	return result == 0
}

// GenerateSalt 生成指定长度的随机盐
func GenerateSalt(length int) ([]byte, error) {
	if length <= 0 {
		return nil, errors.New("salt length must be greater than 0")
	}
	
	salt := make([]byte, length)
	_, err := io.ReadFull(kdfRandReader, salt)
	if err != nil {
		return nil, fmt.Errorf("failed to generate salt: %w", err)
	}
	
	return salt, nil
}