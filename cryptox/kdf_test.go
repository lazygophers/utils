package cryptox

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"testing"
)

func TestPBKDF2SHA256(t *testing.T) {
	tests := []struct {
		name       string
		password   []byte
		salt       []byte
		iterations int
		keyLen     int
		wantNil    bool
	}{
		{
			name:       "标准参数",
			password:   []byte("password123"),
			salt:       []byte("randomsalt16byte"),
			iterations: 10000,
			keyLen:     32,
			wantNil:    false,
		},
		{
			name:       "高迭代次数",
			password:   []byte("secure_password"),
			salt:       []byte("longsalt32bytes1234567890abcd"),
			iterations: 100000,
			keyLen:     32,
			wantNil:    false,
		},
		{
			name:       "短密钥",
			password:   []byte("pass"),
			salt:       []byte("salt"),
			iterations: 1000,
			keyLen:     16,
			wantNil:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PBKDF2SHA256(tt.password, tt.salt, tt.iterations, tt.keyLen)
			if (got == nil) != tt.wantNil {
				t.Errorf("PBKDF2SHA256() = %v, wantNil %v", got, tt.wantNil)
			}
			if len(got) != tt.keyLen {
				t.Errorf("PBKDF2SHA256() length = %v, want %v", len(got), tt.keyLen)
			}
		})
	}

	// 测试确定性：相同输入应产生相同输出
	key1 := PBKDF2SHA256([]byte("password"), []byte("salt"), 10000, 32)
	key2 := PBKDF2SHA256([]byte("password"), []byte("salt"), 10000, 32)
	if !bytes.Equal(key1, key2) {
		t.Error("PBKDF2SHA256() should be deterministic")
	}

	// 测试不同盐产生不同密钥
	key3 := PBKDF2SHA256([]byte("password"), []byte("salt1"), 10000, 32)
	key4 := PBKDF2SHA256([]byte("password"), []byte("salt2"), 10000, 32)
	if bytes.Equal(key3, key4) {
		t.Error("PBKDF2SHA256() should produce different keys for different salts")
	}
}

func TestPBKDF2SHA512(t *testing.T) {
	tests := []struct {
		name       string
		password   []byte
		salt       []byte
		iterations int
		keyLen     int
		wantNil    bool
	}{
		{
			name:       "标准参数",
			password:   []byte("password123"),
			salt:       []byte("randomsalt16byte"),
			iterations: 10000,
			keyLen:     64,
			wantNil:    false,
		},
		{
			name:       "32字节密钥",
			password:   []byte("secure_password"),
			salt:       []byte("longsalt32bytes1234567890abcd"),
			iterations: 50000,
			keyLen:     32,
			wantNil:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PBKDF2SHA512(tt.password, tt.salt, tt.iterations, tt.keyLen)
			if (got == nil) != tt.wantNil {
				t.Errorf("PBKDF2SHA512() = %v, wantNil %v", got, tt.wantNil)
			}
			if len(got) != tt.keyLen {
				t.Errorf("PBKDF2SHA512() length = %v, want %v", len(got), tt.keyLen)
			}
		})
	}

	// 测试确定性
	key1 := PBKDF2SHA512([]byte("password"), []byte("salt"), 10000, 64)
	key2 := PBKDF2SHA512([]byte("password"), []byte("salt"), 10000, 64)
	if !bytes.Equal(key1, key2) {
		t.Error("PBKDF2SHA512() should be deterministic")
	}
}

func TestDeriveKey(t *testing.T) {
	password := []byte("testpassword")
	salt := []byte("testsalt12345678")
	iterations := 10000
	keyLen := 32

	// 测试使用 SHA-256
	key256 := DeriveKey(password, salt, iterations, keyLen, sha256.New)
	if len(key256) != keyLen {
		t.Errorf("DeriveKey(SHA256) length = %v, want %v", len(key256), keyLen)
	}

	// 应该与 PBKDF2SHA256 相同
	expected := PBKDF2SHA256(password, salt, iterations, keyLen)
	if !bytes.Equal(key256, expected) {
		t.Error("DeriveKey(SHA256) should match PBKDF2SHA256")
	}

	// 测试使用 SHA-512
	key512 := DeriveKey(password, salt, iterations, keyLen, sha512.New)
	if len(key512) != keyLen {
		t.Errorf("DeriveKey(SHA512) length = %v, want %v", len(key512), keyLen)
	}

	// SHA-256 和 SHA-512 应该产生不同的密钥
	if bytes.Equal(key256, key512) {
		t.Error("DeriveKey() should produce different keys for different hash functions")
	}
}

func TestDefaultPBKDF2Config(t *testing.T) {
	config := DefaultPBKDF2Config()

	if config.Iterations < 100000 {
		t.Errorf("DefaultPBKDF2Config.Iterations = %v, want >= 100000", config.Iterations)
	}
	if config.SaltLen < 16 {
		t.Errorf("DefaultPBKDF2Config.SaltLen = %v, want >= 16", config.SaltLen)
	}
	if config.KeyLen < 32 {
		t.Errorf("DefaultPBKDF2Config.KeyLen = %v, want >= 32", config.KeyLen)
	}
	if config.HashFunc == nil {
		t.Error("DefaultPBKDF2Config.HashFunc should not be nil")
	}
}

func TestFastPBKDF2Config(t *testing.T) {
	config := FastPBKDF2Config()

	if config.Iterations < 50000 {
		t.Errorf("FastPBKDF2Config.Iterations = %v, want >= 50000", config.Iterations)
	}
	if config.SaltLen < 16 {
		t.Errorf("FastPBKDF2Config.SaltLen = %v, want >= 16", config.SaltLen)
	}
	if config.KeyLen < 32 {
		t.Errorf("FastPBKDF2Config.KeyLen = %v, want >= 32", config.KeyLen)
	}
	if config.HashFunc == nil {
		t.Error("FastPBKDF2Config.HashFunc should not be nil")
	}
}

func TestStrongPBKDF2Config(t *testing.T) {
	config := StrongPBKDF2Config()

	if config.Iterations < 100000 {
		t.Errorf("StrongPBKDF2Config.Iterations = %v, want >= 100000", config.Iterations)
	}
	if config.SaltLen < 32 {
		t.Errorf("StrongPBKDF2Config.SaltLen = %v, want >= 32", config.SaltLen)
	}
	if config.KeyLen < 32 {
		t.Errorf("StrongPBKDF2Config.KeyLen = %v, want >= 32", config.KeyLen)
	}
	if config.HashFunc == nil {
		t.Error("StrongPBKDF2Config.HashFunc should not be nil")
	}

	// StrongPBKDF2Config 应该比 DefaultPBKDF2Config 更强
	defaultConfig := DefaultPBKDF2Config()
	if config.Iterations <= defaultConfig.Iterations {
		t.Error("StrongPBKDF2Config should have more iterations than DefaultPBKDF2Config")
	}
}

func TestPBKDF2EdgeCases(t *testing.T) {
	// 空密码
	key := PBKDF2SHA256([]byte(""), []byte("salt"), 10000, 32)
	if len(key) != 32 {
		t.Error("PBKDF2SHA256 should handle empty password")
	}

	// 空盐
	key = PBKDF2SHA256([]byte("password"), []byte(""), 10000, 32)
	if len(key) != 32 {
		t.Error("PBKDF2SHA256 should handle empty salt")
	}

	// 低迭代次数（虽然不推荐，但应该能工作）
	key = PBKDF2SHA256([]byte("password"), []byte("salt"), 1, 32)
	if len(key) != 32 {
		t.Error("PBKDF2SHA256 should handle low iteration count")
	}
}

func BenchmarkPBKDF2SHA256_10000(b *testing.B) {
	password := []byte("benchmarkpassword")
	salt := []byte("benchmarksalt123")
	for i := 0; i < b.N; i++ {
		_ = PBKDF2SHA256(password, salt, 10000, 32)
	}
}

func BenchmarkPBKDF2SHA256_100000(b *testing.B) {
	password := []byte("benchmarkpassword")
	salt := []byte("benchmarksalt123")
	for i := 0; i < b.N; i++ {
		_ = PBKDF2SHA256(password, salt, 100000, 32)
	}
}

func BenchmarkPBKDF2SHA512_10000(b *testing.B) {
	password := []byte("benchmarkpassword")
	salt := []byte("benchmarksalt123")
	for i := 0; i < b.N; i++ {
		_ = PBKDF2SHA512(password, salt, 10000, 64)
	}
}

func BenchmarkPBKDF2SHA512_100000(b *testing.B) {
	password := []byte("benchmarkpassword")
	salt := []byte("benchmarksalt123")
	for i := 0; i < b.N; i++ {
		_ = PBKDF2SHA512(password, salt, 100000, 64)
	}
}
