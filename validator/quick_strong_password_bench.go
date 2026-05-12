package validator

import (
	"fmt"
	"testing"
)

// validateStrongPasswordOld 旧版本实现
func validateStrongPasswordOld(password string) bool {
	if len(password) < 8 {
		return false
	}

	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, char := range password {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= '0' && char <= '9':
			hasNumber = true
		case (char >= '!' && char <= '/') || (char >= ':' && char <= '@') || (char >= '[' && char <= '`') || (char >= '{' && char <= '~'):
			hasSpecial = true
		}
	}

	count := 0
	if hasUpper {
		count++
	}
	if hasLower {
		count++
	}
	if hasNumber {
		count++
	}
	if hasSpecial {
		count++
	}

	return count >= 3
}

// TestStrongPassword_Performance 对比新旧版本性能
func TestStrongPassword_Performance(t *testing.T) {
	passwords := []string{
		"Abc123!@",
		"Password123!",
		"SecurePass#2024",
		"MyP@ssw0rd",
		"Test@1234",
	}

	// 测试旧版本
	oldResult := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, pwd := range passwords {
				_ = validateStrongPasswordOld(pwd)
			}
		}
	})

	// 测试新版本（使用相同的逻辑）
	newResult := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, pwd := range passwords {
				_ = validateStrongPasswordFast(pwd)
			}
		}
	})

	improvement := ((float64(oldResult.NsPerOp()) - float64(newResult.NsPerOp())) / float64(oldResult.NsPerOp())) * 100

	fmt.Printf("\n=== Strong Password Performance Comparison ===\n")
	fmt.Printf("Old version (rune loop): %d ns/op, %d allocs/op\n", oldResult.NsPerOp(), oldResult.AllocsPerOp())
	fmt.Printf("New version (byte loop + fast fail): %d ns/op, %d allocs/op\n", newResult.NsPerOp(), newResult.AllocsPerOp())
	fmt.Printf("Performance improvement: %.1f%%\n", improvement)
	fmt.Printf("Memory allocation: %s\n", map[bool]string{true: "✅ Zero allocation", false: "❌ Has allocation"}[newResult.AllocsPerOp() == 0])

	// 验证正确性
	for _, pwd := range passwords {
		oldResult := validateStrongPasswordOld(pwd)
		newResult := validateStrongPasswordFast(pwd)
		if oldResult != newResult {
			t.Errorf("结果不一致: password=%s, old=%v, new=%v", pwd, oldResult, newResult)
		}
	}
}

// validateStrongPasswordFast 新版本实现
func validateStrongPasswordFast(password string) bool {
	if len(password) < 8 {
		return false
	}

	var (
		hasUpper   uint8
		hasLower   uint8
		hasNumber  uint8
		hasSpecial uint8
	)

	for i := 0; i < len(password); i++ {
		c := password[i]
		switch {
		case c >= 'A' && c <= 'Z':
			hasUpper = 1
		case c >= 'a' && c <= 'z':
			hasLower = 1
		case c >= '0' && c <= '9':
			hasNumber = 1
		default:
			// 可打印 ASCII 字符视为特殊字符
			if c >= 32 && c <= 126 {
				hasSpecial = 1
			}
		}

		// 快速失败：已经找到3种类型
		if hasUpper+hasLower+hasNumber+hasSpecial >= 3 {
			return true
		}
	}

	// 至少包含大写字母、小写字母、数字、特殊字符中的三种
	return hasUpper+hasLower+hasNumber+hasSpecial >= 3
}

// BenchmarkStrongPassword_Old 旧版本基准测试
func BenchmarkStrongPassword_Old(b *testing.B) {
	passwords := []string{
		"Abc123!@",
		"Password123!",
		"SecurePass#2024",
		"MyP@ssw0rd",
		"Test@1234",
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, pwd := range passwords {
			_ = validateStrongPasswordOld(pwd)
		}
	}
}

// BenchmarkStrongPassword_New 新版本基准测试
func BenchmarkStrongPassword_New(b *testing.B) {
	passwords := []string{
		"Abc123!@",
		"Password123!",
		"SecurePass#2024",
		"MyP@ssw0rd",
		"Test@1234",
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, pwd := range passwords {
			_ = validateStrongPasswordFast(pwd)
		}
	}
}
