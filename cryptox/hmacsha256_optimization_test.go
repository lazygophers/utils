package cryptox

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"testing"
)

// TestHMACSHA256Optimization 验证优化后的 HMACSHA256 函数正确性
func TestHMACSHA256Optimization(t *testing.T) {
	// 使用标准库计算期望值
	tests := []struct {
		name    string
		key     string
		message string
	}{
		{
			name:    "基本测试",
			key:     "key",
			message: "message",
		},
		{
			name:    "空字符串",
			key:     "",
			message: "",
		},
		{
			name:    "长消息",
			key:     "secret-key",
			message: "This is a longer message to test the HMAC-SHA256 implementation",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 使用标准库计算期望值
			h := hmac.New(sha256.New, []byte(tt.key))
			h.Write([]byte(tt.message))
			expected := fmt.Sprintf("%x", h.Sum(nil))

			result := HMACSHA256(tt.key, tt.message)
			if result != expected {
				t.Errorf("HMACSHA256() = %v, want %v", result, expected)
			}
		})
	}
}

// TestHMACSHA256Generics 测试泛型约束
func TestHMACSHA256Generics(t *testing.T) {
	key := "test-key"
	message := "test-message"

	// 测试 string 参数
	result1 := HMACSHA256(key, message)
	if len(result1) != 64 {
		t.Errorf("HMACSHA256(string) returned length %d, expected 64", len(result1))
	}

	// 测试 []byte 参数
	result2 := HMACSHA256([]byte(key), []byte(message))
	if len(result2) != 64 {
		t.Errorf("HMACSHA256([]byte) returned length %d, expected 64", len(result2))
	}

	// 验证两种方式结果一致
	if result1 != result2 {
		t.Errorf("HMACSHA256(string) != HMACSHA256([]byte)")
	}
}

// BenchmarkHMACSHA256_Old 旧实现（fmt.Sprintf）
func BenchmarkHMACSHA256_Old(b *testing.B) {
	key, message := "test-key", "test-message"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := hmac.New(sha256.New, []byte(key))
		_, _ = h.Write([]byte(message))
		_ = fmt.Sprintf("%x", h.Sum(nil))
	}
}

// BenchmarkHMACSHA256_New 新实现（手动 hex 编码）
func BenchmarkHMACSHA256_New(b *testing.B) {
	key, message := "test-key", "test-message"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = HMACSHA256(key, message)
	}
}

// BenchmarkHMACSHA256_New_Bytes 新实现（使用 []byte 参数）
func BenchmarkHMACSHA256_New_Bytes(b *testing.B) {
	key, message := []byte("test-key"), []byte("test-message")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = HMACSHA256(key, message)
	}
}
