package anyx

import (
	"strconv"
	"testing"
)

// ============================================================
// Benchmark: parseIndex 不同场景性能测试（详细版本）
// ============================================================

// 1. 两位数
func BenchmarkParseIndex_TwoDigits(b *testing.B) {
	s := "42"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = parseIndex(s)
	}
}

// 2. 三位数（常见场景）
func BenchmarkParseIndex_ThreeDigits(b *testing.B) {
	s := "123"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = parseIndex(s)
	}
}

// 6. 负个位数
func BenchmarkParseIndex_NegativeSingle(b *testing.B) {
	s := "-1"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = parseIndex(s)
	}
}

// 7. 空字符串（错误路径）
func BenchmarkParseIndex_Empty(b *testing.B) {
	s := ""
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = parseIndex(s)
	}
}

// 8. 非数字字符串（错误路径）
func BenchmarkParseIndex_Invalid(b *testing.B) {
	s := "abc"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = parseIndex(s)
	}
}

// 9. 零
func BenchmarkParseIndex_Zero(b *testing.B) {
	s := "0"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = parseIndex(s)
	}
}

// 10. 带前导零
func BenchmarkParseIndex_LeadingZero(b *testing.B) {
	s := "007"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = parseIndex(s)
	}
}

// 11. 极大数字（边界测试）
func BenchmarkParseIndex_MaxInt(b *testing.B) {
	s := "2147483647" // math.MaxInt32
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = parseIndex(s)
	}
}

// 12. 混合测试（真实场景）
func BenchmarkParseIndex_Mixed(b *testing.B) {
	cases := []string{"0", "1", "10", "100", "-1", "-10"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range cases {
			_, _ = parseIndex(s)
		}
	}
}

// ============================================================
// 对比测试：当前实现 vs strconv.Atoi
// ============================================================

func BenchmarkParseIndex_Vs_Strconv_Valid_3Digits(b *testing.B) {
	s := "123"

	b.Run("parseIndex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = parseIndex(s)
		}
	})

	b.Run("strconv.Atoi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = strconv.Atoi(s)
		}
	})
}

func BenchmarkParseIndex_Vs_Strconv_Negative(b *testing.B) {
	s := "-456"

	b.Run("parseIndex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = parseIndex(s)
		}
	})

	b.Run("strconv.Atoi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = strconv.Atoi(s)
		}
	})
}

func BenchmarkParseIndex_Vs_Strconv_Single(b *testing.B) {
	s := "5"

	b.Run("parseIndex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = parseIndex(s)
		}
	})

	b.Run("strconv.Atoi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = strconv.Atoi(s)
		}
	})
}

// ============================================================
// 优化版本对比测试
// 注：优化实现在 standalone 文件中定义
// ============================================================

func BenchmarkParseIndex_Optimized_Valid_3Digits(b *testing.B) {
	s := "123"

	b.Run("Current", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = parseIndex(s)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = parseIndexOptimized(s)
		}
	})
}

func BenchmarkParseIndex_Optimized_Negative(b *testing.B) {
	s := "-456"

	b.Run("Current", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = parseIndex(s)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = parseIndexOptimized(s)
		}
	})
}

func BenchmarkParseIndex_Optimized_Single(b *testing.B) {
	s := "5"

	b.Run("Current", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = parseIndex(s)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = parseIndexOptimized(s)
		}
	})
}

func BenchmarkParseIndex_Optimized_Large(b *testing.B) {
	s := "999999"

	b.Run("Current", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = parseIndex(s)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = parseIndexOptimized(s)
		}
	})
}

// ============================================================
// 错误路径对比
// ============================================================

func BenchmarkParseIndex_Error_Empty(b *testing.B) {
	s := ""

	b.Run("Current", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = parseIndex(s)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = parseIndexOptimized(s)
		}
	})
}

func BenchmarkParseIndex_Error_Invalid(b *testing.B) {
	s := "abc"

	b.Run("Current", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = parseIndex(s)
		}
	})

	b.Run("Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = parseIndexOptimized(s)
		}
	})
}

// ============================================================
// 内存分配对比
// ============================================================

func BenchmarkParseIndex_Allocs_Single(b *testing.B) {
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
}

func BenchmarkParseIndex_Allocs_Negative(b *testing.B) {
	s := "-123"

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
}
