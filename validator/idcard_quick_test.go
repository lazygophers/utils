package validator

import (
	"testing"
)

// TestIDCard18Optimization 快速验证优化后的 validateIDCard18 函数
func TestIDCard18Optimization(t *testing.T) {
	validCases := []string{
		"110101199003072273", // 北京
		"310104199010017834", // 上海
		"44030819910403921X", // 广东（大X）
		"44030819910403921x", // 广东（小x）
		"11010119800101123X", // 测试用例
	}

	invalidCases := []string{
		"",                    // 空
		"12345678901234567",   // 17位
		"1234567890123456789", // 19位
		"abcdefghijklmnopqr",  // 全字母
		"110101A01003072273",  // 包含字母
		"110101 199003072273", // 包含空格
	}

	for _, tc := range validCases {
		if !validateIDCard18(tc) {
			t.Errorf("validateIDCard18(%s) = false, 期望 true", tc)
		}
	}

	for _, tc := range invalidCases {
		if validateIDCard18(tc) {
			t.Errorf("validateIDCard18(%s) = true, 期望 false", tc)
		}
	}
}

// BenchmarkIDCard18_Optimized 优化后的性能
func BenchmarkIDCard18_Optimized_Valid(b *testing.B) {
	testCard := "110101199003072273"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18(testCard)
	}
}

func BenchmarkIDCard18_Optimized_Invalid(b *testing.B) {
	testCard := "invalid"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		validateIDCard18(testCard)
	}
}

func BenchmarkIDCard18_Optimized_Alloc(b *testing.B) {
	testCard := "110101199003072273"
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		validateIDCard18(testCard)
	}
}
