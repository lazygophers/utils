package stringx

import (
	"encoding/base64"
	"encoding/hex"
	"testing"
)

func TestSecureRandBytes(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		wantLen int
		wantErr bool
	}{
		{"正常情况", 16, 16, false},
		{"零长度", 0, 0, false},
		{"负数长度", -1, 0, false},
		{"大长度", 1024, 1024, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SecureRandBytes(tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("SecureRandBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("SecureRandBytes() length = %v, want %v", len(got), tt.wantLen)
			}
		})
	}

	// 测试随机性：生成两次应该不同
	b1, _ := SecureRandBytes(16)
	b2, _ := SecureRandBytes(16)
	if string(b1) == string(b2) {
		t.Error("SecureRandBytes() generated identical results, should be random")
	}
}

func TestSecureRandString(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		wantErr bool
	}{
		{"正常情况", 16, false},
		{"零长度", 0, false},
		{"负数长度", -1, false},
		{"大长度", 256, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SecureRandString(tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("SecureRandString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.n > 0 && len(got) > tt.n {
				t.Errorf("SecureRandString() length = %v, want <= %v", len(got), tt.n)
			}
			// 验证是否是有效的base64
			if tt.n > 0 && got != "" {
				_, err := base64.RawURLEncoding.DecodeString(got)
				if err != nil {
					t.Errorf("SecureRandString() result is not valid base64: %v", err)
				}
			}
		})
	}

	// 测试随机性
	s1, _ := SecureRandString(16)
	s2, _ := SecureRandString(16)
	if s1 == s2 {
		t.Error("SecureRandString() generated identical results, should be random")
	}
}

func TestSecureRandHex(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		wantErr bool
	}{
		{"正常情况", 16, false},
		{"零长度", 0, false},
		{"负数长度", -1, false},
		{"奇数长度", 15, false},
		{"大长度", 256, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SecureRandHex(tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("SecureRandHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.n > 0 && len(got) > tt.n {
				t.Errorf("SecureRandHex() length = %v, want <= %v", len(got), tt.n)
			}
			// 验证是否是有效的十六进制（只验证偶数长度）
			if tt.n > 0 && got != "" && len(got)%2 == 0 {
				_, err := hex.DecodeString(got)
				if err != nil {
					t.Errorf("SecureRandHex() result is not valid hex: %v", err)
				}
			}
		})
	}

	// 测试随机性
	h1, _ := SecureRandHex(16)
	h2, _ := SecureRandHex(16)
	if h1 == h2 {
		t.Error("SecureRandHex() generated identical results, should be random")
	}
}

func TestSecureRandLetters(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		wantErr bool
	}{
		{"正常情况", 16, false},
		{"零长度", 0, false},
		{"负数长度", -1, false},
		{"大长度", 256, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SecureRandLetters(tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("SecureRandLetters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			expectedLen := tt.n
			if tt.n < 0 {
				expectedLen = 0
			}
			if len(got) != expectedLen {
				t.Errorf("SecureRandLetters() length = %v, want %v", len(got), expectedLen)
			}
			// 验证只包含字母
			for _, c := range got {
				if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')) {
					t.Errorf("SecureRandLetters() contains non-letter character: %c", c)
				}
			}
		})
	}

	// 测试随机性
	s1, _ := SecureRandLetters(16)
	s2, _ := SecureRandLetters(16)
	if s1 == s2 {
		t.Error("SecureRandLetters() generated identical results, should be random")
	}
}

func TestSecureRandLowerLetters(t *testing.T) {
	got, err := SecureRandLowerLetters(16)
	if err != nil {
		t.Fatalf("SecureRandLowerLetters() error = %v", err)
	}
	if len(got) != 16 {
		t.Errorf("SecureRandLowerLetters() length = %v, want 16", len(got))
	}
	for _, c := range got {
		if c < 'a' || c > 'z' {
			t.Errorf("SecureRandLowerLetters() contains non-lowercase character: %c", c)
		}
	}
}

func TestSecureRandUpperLetters(t *testing.T) {
	got, err := SecureRandUpperLetters(16)
	if err != nil {
		t.Fatalf("SecureRandUpperLetters() error = %v", err)
	}
	if len(got) != 16 {
		t.Errorf("SecureRandUpperLetters() length = %v, want 16", len(got))
	}
	for _, c := range got {
		if c < 'A' || c > 'Z' {
			t.Errorf("SecureRandUpperLetters() contains non-uppercase character: %c", c)
		}
	}
}

func TestSecureRandNumbers(t *testing.T) {
	got, err := SecureRandNumbers(16)
	if err != nil {
		t.Fatalf("SecureRandNumbers() error = %v", err)
	}
	if len(got) != 16 {
		t.Errorf("SecureRandNumbers() length = %v, want 16", len(got))
	}
	for _, c := range got {
		if c < '0' || c > '9' {
			t.Errorf("SecureRandNumbers() contains non-digit character: %c", c)
		}
	}
}

func TestSecureRandLetterNumbers(t *testing.T) {
	got, err := SecureRandLetterNumbers(16)
	if err != nil {
		t.Fatalf("SecureRandLetterNumbers() error = %v", err)
	}
	if len(got) != 16 {
		t.Errorf("SecureRandLetterNumbers() length = %v, want 16", len(got))
	}
	for _, c := range got {
		valid := (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
		if !valid {
			t.Errorf("SecureRandLetterNumbers() contains invalid character: %c", c)
		}
	}
}

func TestSecureRandLowerLetterNumbers(t *testing.T) {
	got, err := SecureRandLowerLetterNumbers(16)
	if err != nil {
		t.Fatalf("SecureRandLowerLetterNumbers() error = %v", err)
	}
	if len(got) != 16 {
		t.Errorf("SecureRandLowerLetterNumbers() length = %v, want 16", len(got))
	}
	for _, c := range got {
		valid := (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z')
		if !valid {
			t.Errorf("SecureRandLowerLetterNumbers() contains invalid character: %c", c)
		}
	}
}

func TestSecureRandUpperLetterNumbers(t *testing.T) {
	got, err := SecureRandUpperLetterNumbers(16)
	if err != nil {
		t.Fatalf("SecureRandUpperLetterNumbers() error = %v", err)
	}
	if len(got) != 16 {
		t.Errorf("SecureRandUpperLetterNumbers() length = %v, want 16", len(got))
	}
	for _, c := range got {
		valid := (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z')
		if !valid {
			t.Errorf("SecureRandUpperLetterNumbers() contains invalid character: %c", c)
		}
	}
}

func BenchmarkSecureRandBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = SecureRandBytes(32)
	}
}

func BenchmarkSecureRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = SecureRandString(32)
	}
}

func BenchmarkSecureRandHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = SecureRandHex(32)
	}
}

func BenchmarkSecureRandLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = SecureRandLetters(32)
	}
}
