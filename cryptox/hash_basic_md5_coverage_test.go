package cryptox

import (
	"crypto/md5"
	"testing"
)

// TestMd5_Coverage_String 测试 string 输入
func TestMd5_Coverage_String(t *testing.T) {
	input := "hello world"
	result := Md5(input)
	if len(result) != 32 {
		t.Errorf("MD5 hash length should be 32, got %d", len(result))
	}
	// 验证结果一致性
	expected := md5.Sum([]byte(input))
	expectedHex := encodeHex(expected[:])
	if result != expectedHex {
		t.Errorf("MD5 hash mismatch: got %s, want %s", result, expectedHex)
	}
}

// TestMd5_Coverage_Bytes 测试 []byte 输入
func TestMd5_Coverage_Bytes(t *testing.T) {
	input := []byte("hello world")
	result := Md5(input)
	if len(result) != 32 {
		t.Errorf("MD5 hash length should be 32, got %d", len(result))
	}
	// 验证结果一致性
	expected := md5.Sum(input)
	expectedHex := encodeHex(expected[:])
	if result != expectedHex {
		t.Errorf("MD5 hash mismatch: got %s, want %s", result, expectedHex)
	}
}

// TestMd5_Coverage_Empty 测试空输入
func TestMd5_Coverage_Empty(t *testing.T) {
	result := Md5("")
	if len(result) != 32 {
		t.Errorf("MD5 hash length should be 32, got %d", len(result))
	}
	// 验证标准空字符串 MD5
	expected := "d41d8cd98f00b204e9800998ecf8427e"
	if result != expected {
		t.Errorf("Empty MD5 hash: got %s, want %s", result, expected)
	}
}

// TestMd5_Coverage_Large 测试大数据
func TestMd5_Coverage_Large(t *testing.T) {
	input := make([]byte, 1024*1024)
	for i := range input {
		input[i] = byte(i % 256)
	}
	result := Md5(input)
	if len(result) != 32 {
		t.Errorf("MD5 hash length should be 32, got %d", len(result))
	}
}

// TestMd5_Coverage_SpecialChars 测试特殊字符
func TestMd5_Coverage_SpecialChars(t *testing.T) {
	inputs := []string{
		"!@#$%^&*()",
		"中文测试",
		"\n\r\t",
		"🔥🚀",
	}
	for _, input := range inputs {
		result := Md5(input)
		if len(result) != 32 {
			t.Errorf("MD5 hash length should be 32 for %q, got %d", input, len(result))
		}
	}
}

// TestMd5_Coverage_Consistency 测试结果一致性
func TestMd5_Coverage_Consistency(t *testing.T) {
	input := "consistency test"
	r1 := Md5(input)
	r2 := Md5(input)
	r3 := Md5([]byte(input))
	if r1 != r2 || r2 != r3 {
		t.Errorf("MD5 results inconsistent: %s, %s, %s", r1, r2, r3)
	}
}

// TestMd5_Coverage_DifferentInputs 测试不同输入产生不同结果
func TestMd5_Coverage_DifferentInputs(t *testing.T) {
	inputs := []string{"a", "b", "c"}
	hashes := make(map[string]bool)
	for _, input := range inputs {
		hash := Md5(input)
		if hashes[hash] {
			t.Errorf("Duplicate hash for different input: %s", hash)
		}
		hashes[hash] = true
	}
}

// TestMd5_Coverage_BinaryData 测试二进制数据
func TestMd5_Coverage_BinaryData(t *testing.T) {
	input := []byte{0x00, 0x01, 0x02, 0xFF, 0xFE, 0xFD}
	result := Md5(input)
	if len(result) != 32 {
		t.Errorf("MD5 hash length should be 32, got %d", len(result))
	}
}

// TestMd5_Coverage_Concurrent 测试并发安全
func TestMd5_Coverage_Concurrent(t *testing.T) {
	input := "concurrent test"
	results := make(chan string, 100)
	for i := 0; i < 100; i++ {
		go func() {
			results <- Md5(input)
		}()
	}
	expected := Md5(input)
	for i := 0; i < 100; i++ {
		if result := <-results; result != expected {
			t.Errorf("Concurrent MD5 mismatch: got %s, want %s", result, expected)
		}
	}
}

// TestMd5_Coverage_KnownValues 测试已知值
func TestMd5_Coverage_KnownValues(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", "d41d8cd98f00b204e9800998ecf8427e"},
		{"a", "0cc175b9c0f1b6a831c399e269772661"},
		{"abc", "900150983cd24fb0d6963f7d28e17f72"},
		{"message digest", "f96b697d7cb7938d525a2f31aaf161d0"},
		{"abcdefghijklmnopqrstuvwxyz", "c3fcd3d76192e4007dfb496cca67e13b"},
	}
	for _, tc := range testCases {
		result := Md5(tc.input)
		if result != tc.expected {
			t.Errorf("MD5(%q) = %s, want %s", tc.input, result, tc.expected)
		}
	}
}

// encodeHex 辅助函数：将字节转换为 hex 字符串
func encodeHex(data []byte) string {
	const hexTable = "0123456789abcdef"
	result := make([]byte, len(data)*2)
	for i, b := range data {
		result[i*2] = hexTable[b>>4]
		result[i*2+1] = hexTable[b&0x0f]
	}
	return string(result)
}
