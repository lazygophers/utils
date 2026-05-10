package cryptox

import (
	"crypto/sha1"
	"testing"
)

// TestSHA1_Coverage_String 测试 string 输入
func TestSHA1_Coverage_String(t *testing.T) {
	input := "hello world"
	result := SHA1(input)
	if len(result) != 40 {
		t.Errorf("SHA1 hash length should be 40, got %d", len(result))
	}
	// 验证结果一致性
	expected := sha1.Sum([]byte(input))
	expectedHex := encodeHexSHA1(expected[:])
	if result != expectedHex {
		t.Errorf("SHA1 hash mismatch: got %s, want %s", result, expectedHex)
	}
}

// TestSHA1_Coverage_Bytes 测试 []byte 输入
func TestSHA1_Coverage_Bytes(t *testing.T) {
	input := []byte("hello world")
	result := SHA1(input)
	if len(result) != 40 {
		t.Errorf("SHA1 hash length should be 40, got %d", len(result))
	}
}

// TestSHA1_Coverage_Empty 测试空输入
func TestSHA1_Coverage_Empty(t *testing.T) {
	result := SHA1("")
	if len(result) != 40 {
		t.Errorf("SHA1 hash length should be 40, got %d", len(result))
	}
	// 验证标准空字符串 SHA1
	expected := "da39a3ee5e6b4b0d3255bfef95601890afd80709"
	if result != expected {
		t.Errorf("Empty SHA1 hash: got %s, want %s", result, expected)
	}
}

// TestSHA1_Coverage_Large 测试大数据
func TestSHA1_Coverage_Large(t *testing.T) {
	input := make([]byte, 1024*1024)
	for i := range input {
		input[i] = byte(i % 256)
	}
	result := SHA1(input)
	if len(result) != 40 {
		t.Errorf("SHA1 hash length should be 40, got %d", len(result))
	}
}

// TestSHA1_Coverage_SpecialChars 测试特殊字符
func TestSHA1_Coverage_SpecialChars(t *testing.T) {
	inputs := []string{
		"!@#$%^&*()",
		"中文测试",
		"\n\r\t",
		"🔥🚀",
	}
	for _, input := range inputs {
		result := SHA1(input)
		if len(result) != 40 {
			t.Errorf("SHA1 hash length should be 40 for %q, got %d", input, len(result))
		}
	}
}

// TestSHA1_Coverage_Consistency 测试结果一致性
func TestSHA1_Coverage_Consistency(t *testing.T) {
	input := "consistency test"
	r1 := SHA1(input)
	r2 := SHA1(input)
	r3 := SHA1([]byte(input))
	if r1 != r2 || r2 != r3 {
		t.Errorf("SHA1 results inconsistent: %s, %s, %s", r1, r2, r3)
	}
}

// TestSHA1_Coverage_DifferentInputs 测试不同输入产生不同结果
func TestSHA1_Coverage_DifferentInputs(t *testing.T) {
	inputs := []string{"a", "b", "c"}
	hashes := make(map[string]bool)
	for _, input := range inputs {
		hash := SHA1(input)
		if hashes[hash] {
			t.Errorf("Duplicate hash for different input: %s", hash)
		}
		hashes[hash] = true
	}
}

// TestSHA1_Coverage_BinaryData 测试二进制数据
func TestSHA1_Coverage_BinaryData(t *testing.T) {
	input := []byte{0x00, 0x01, 0x02, 0xFF, 0xFE, 0xFD}
	result := SHA1(input)
	if len(result) != 40 {
		t.Errorf("SHA1 hash length should be 40, got %d", len(result))
	}
}

// TestSHA1_Coverage_Concurrent 测试并发安全
func TestSHA1_Coverage_Concurrent(t *testing.T) {
	input := "concurrent test"
	results := make(chan string, 100)
	for i := 0; i < 100; i++ {
		go func() {
			results <- SHA1(input)
		}()
	}
	expected := SHA1(input)
	for i := 0; i < 100; i++ {
		if result := <-results; result != expected {
			t.Errorf("Concurrent SHA1 mismatch: got %s, want %s", result, expected)
		}
	}
}

// TestSHA1_Coverage_KnownValues 测试已知值
func TestSHA1_Coverage_KnownValues(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", "da39a3ee5e6b4b0d3255bfef95601890afd80709"},
		{"a", "86f7e437faa5a7fce15d1ddcb9eaeaea377667b8"},
		{"abc", "a9993e364706816aba3e25717850c26c9cd0d89d"},
		{"message digest", "c12252ceda8be8994d5fa0290a47231c1d16aae3"},
		{"abcdefghijklmnopqrstuvwxyz", "32d10c7b8cf96570ca04ce37f2a19d84240d3a89"},
	}
	for _, tc := range testCases {
		result := SHA1(tc.input)
		if result != tc.expected {
			t.Errorf("SHA1(%q) = %s, want %s", tc.input, result, tc.expected)
		}
	}
}

// encodeHexSHA1 辅助函数：将字节转换为 hex 字符串
func encodeHexSHA1(data []byte) string {
	const hexTable = "0123456789abcdef"
	result := make([]byte, len(data)*2)
	for i, b := range data {
		result[i*2] = hexTable[b>>4]
		result[i*2+1] = hexTable[b&0x0f]
	}
	return string(result)
}
