package validator

import (
	"reflect"
	"testing"
)

// ========== Alpha 优化方案 ==========

func alphaManualRange(s string) bool {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')) {
			return false
		}
	}
	return true
}

func alphaSinglePass(s string) bool {
	if len(s) == 0 {
		return true
	}
	for i := 0; i < len(s); i++ {
		c := s[i] | 0x20
		if c < 'a' || c > 'z' {
			return false
		}
	}
	return true
}

func alphaBitOps(s string) bool {
	if len(s) == 0 {
		return true
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		isLower := (c | 0x20) == 'a' && c-'a' < 26
		if !isLower {
			return false
		}
	}
	return true
}

// ========== Alphanum 优化方案 ==========

func alphanumManualRange(s string) bool {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')) {
			return false
		}
	}
	return true
}

func alphanumSinglePass(s string) bool {
	if len(s) == 0 {
		return true
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		cl := c | 0x20
		if (cl < 'a' || cl > 'z') && (c < '0' || c > '9') {
			return false
		}
	}
	return true
}

func alphanumBitOps(s string) bool {
	if len(s) == 0 {
		return true
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		isLetter := (c | 0x20) == 'a' && c-'a' < 26
		isDigit := c-'0' < 10
		if !isLetter && !isDigit {
			return false
		}
	}
	return true
}

// ========== 基准测试 ==========

var testAlphaMedium = "HelloWorldTest"
var testAlphanumMedium = "User123456789"

func BenchmarkAlphaRegex(b *testing.B) {
	fl := &fieldLevel{field: reflect.ValueOf(testAlphaMedium)}
	validator := Alpha()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkAlphaManualRange(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		alphaManualRange(testAlphaMedium)
	}
}

func BenchmarkAlphaSinglePass(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		alphaSinglePass(testAlphaMedium)
	}
}

func BenchmarkAlphaBitOps(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		alphaBitOps(testAlphaMedium)
	}
}

func BenchmarkAlphanumRegex(b *testing.B) {
	fl := &fieldLevel{field: reflect.ValueOf(testAlphanumMedium)}
	validator := Alphanum()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator(fl)
	}
}

func BenchmarkAlphanumManualRange(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		alphanumManualRange(testAlphanumMedium)
	}
}

func BenchmarkAlphanumSinglePass(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		alphanumSinglePass(testAlphanumMedium)
	}
}

func BenchmarkAlphanumBitOps(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		alphanumBitOps(testAlphanumMedium)
	}
}
