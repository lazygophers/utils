package anyx

import (
	"fmt"
	"strconv"
	"testing"
)

// parseIndexOptimized 优化版本：使用 byte 索引代替 rune range
func parseIndexOptimized(s string) (int, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("%w: empty index", ErrInvalidIndex)
	}

	start := 0
	negative := false
	if s[0] == '-' {
		negative = true
		start = 1
		// Bug fix: "-" should return error, not 0
		if len(s) == 1 {
			return 0, fmt.Errorf("%w: '%s'", ErrInvalidIndex, s)
		}
	}

	var result int
	for i := start; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, fmt.Errorf("%w: '%s'", ErrInvalidIndex, s)
		}
		result = result*10 + int(c-'0')
	}

	if negative {
		result = -result
	}

	return result, nil
}

// 正确性验证
func TestParseIndexOptimized_Correctness(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		wantErr  bool
	}{
		{"0", 0, false},
		{"5", 5, false},
		{"123", 123, false},
		{"-1", -1, false},
		{"-456", -456, false},
		{"", 0, true},
		{"abc", 0, true},
		{"-", 0, true}, // Bug fix
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := parseIndexOptimized(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("parseIndexOptimized(%q) expected error, got nil", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("parseIndexOptimized(%q) unexpected error: %v", tt.input, err)
				}
				if got != tt.expected {
					t.Errorf("parseIndexOptimized(%q) = %d, want %d", tt.input, got, tt.expected)
				}
			}
		})
	}
}

// 核心性能对比
func BenchmarkParseIndex_Compare(b *testing.B) {
	cases := []struct {
		name string
		s    string
	}{
		{"SingleDigit", "5"},
		{"TwoDigits", "42"},
		{"ThreeDigits", "123"},
		{"Large", "999999"},
		{"Negative", "-456"},
		{"NegativeSingle", "-1"},
		{"Empty", ""},
		{"Invalid", "abc"},
	}

	for _, tt := range cases {
		b.Run(tt.name, func(b *testing.B) {
			b.Run("Current", func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_, _ = parseIndex(tt.s)
				}
			})

			b.Run("Optimized", func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_, _ = parseIndexOptimized(tt.s)
				}
			})

			b.Run("Strconv", func(b *testing.B) {
				b.ReportAllocs()
				for i := 0; i < b.N; i++ {
					_, _ = strconv.Atoi(tt.s)
				}
			})
		})
	}
}

// 内存分配分析
func BenchmarkParseIndex_Allocs(b *testing.B) {
	b.Run("Positive_Single", func(b *testing.B) {
		s := "5"
		b.Run("Current", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, _ = parseIndex(s)
			}
		})
		b.Run("Optimized", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, _ = parseIndexOptimized(s)
			}
		})
	})

	b.Run("Negative_Single", func(b *testing.B) {
		s := "-1"
		b.Run("Current", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, _ = parseIndex(s)
			}
		})
		b.Run("Optimized", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, _ = parseIndexOptimized(s)
			}
		})
	})

	b.Run("ThreeDigits", func(b *testing.B) {
		s := "123"
		b.Run("Current", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, _ = parseIndex(s)
			}
		})
		b.Run("Optimized", func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, _ = parseIndexOptimized(s)
			}
		})
	})
}
