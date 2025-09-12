package stringx

import (
	"testing"
)

// 优化版本 vs 原版本性能对比
func BenchmarkOptimizedVsOriginal(b *testing.B) {
	// 测试用例
	asciiStr := "HTTPSConnectionXMLParser"
	unicodeStr := "HTTP连接XML解析器"
	longStr := "ThisIsAVeryLongStringForPerformanceTesting"
	
	b.Run("ToSnake_Original_ASCII", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ToSnake(asciiStr)
		}
	})
	
	b.Run("ToSnake_Optimized_ASCII", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = OptimizedToSnake(asciiStr)
		}
	})
	
	b.Run("ToSnake_Original_Unicode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ToSnake(unicodeStr)
		}
	})
	
	b.Run("ToSnake_Optimized_Unicode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = OptimizedToSnake(unicodeStr)
		}
	})
	
	b.Run("Camel2Snake_Original_ASCII", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Camel2Snake(asciiStr)
		}
	})
	
	b.Run("Camel2Snake_Optimized_ASCII", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = OptimizedCamel2Snake(asciiStr)
		}
	})
	
	b.Run("SplitLen_Original", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SplitLen(longStr, 10)
		}
	})
	
	b.Run("SplitLen_Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = OptimizedSplitLen(longStr, 10)
		}
	})
	
	b.Run("Reverse_Original_ASCII", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Reverse(asciiStr)
		}
	})
	
	b.Run("Reverse_Optimized_ASCII", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = OptimizedReverse(asciiStr)
		}
	})
	
	b.Run("Reverse_Original_Unicode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Reverse(unicodeStr)
		}
	})
	
	b.Run("Reverse_Optimized_Unicode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = OptimizedReverse(unicodeStr)
		}
	})
}

// 内存分配对比测试
func BenchmarkMemoryOptimization(b *testing.B) {
	testStr := "HTTPSConnectionXMLParserWithVeryLongName"
	
	b.Run("ToSnake_Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ToSnake(testStr)
		}
	})
	
	b.Run("ToSnake_Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = OptimizedToSnake(testStr)
		}
	})
	
	b.Run("Camel2Snake_Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = Camel2Snake(testStr)
		}
	})
	
	b.Run("Camel2Snake_Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = OptimizedCamel2Snake(testStr)
		}
	})
	
	b.Run("SplitLen_Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = SplitLen(testStr, 8)
		}
	})
	
	b.Run("SplitLen_Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = OptimizedSplitLen(testStr, 8)
		}
	})
}

// 大数据量性能测试
func BenchmarkLargeDataOptimization(b *testing.B) {
	// 生成大型测试数据
	largeASCII := ""
	for i := 0; i < 1000; i++ {
		largeASCII += "CamelCaseStringWithNumbers123"
	}
	
	largeUnicode := ""
	for i := 0; i < 500; i++ {
		largeUnicode += "骆驼命名法字符串测试123"
	}
	
	b.Run("ToSnake_Large_Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ToSnake(largeASCII)
		}
	})
	
	b.Run("ToSnake_Large_Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = OptimizedToSnake(largeASCII)
		}
	})
	
	b.Run("Reverse_Large_Original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = Reverse(largeASCII[:1000]) // 限制测试大小
		}
	})
	
	b.Run("Reverse_Large_Optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = OptimizedReverse(largeASCII[:1000])
		}
	})
}

// CPU密集型操作测试
func BenchmarkCPUIntensive(b *testing.B) {
	strs := []string{
		"SimpleTest",
		"HTTPSConnection", 
		"XMLHttpParser",
		"myVariable123Name",
		"VeryLongCamelCaseStringWithMultipleWordsAndNumbers123456",
	}
	
	b.Run("BatchProcess_Original", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, s := range strs {
				_ = ToSnake(s)
				_ = Camel2Snake(s)
				_ = Reverse(s)
			}
		}
	})
	
	b.Run("BatchProcess_Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, s := range strs {
				_ = OptimizedToSnake(s)
				_ = OptimizedCamel2Snake(s)
				_ = OptimizedReverse(s)
			}
		}
	})
}

// 特定场景优化测试
func BenchmarkSpecialCases(b *testing.B) {
	// 短字符串
	shortStr := "iOS"
	// 纯ASCII
	asciiStr := "HTTPConnection"
	// 混合字符
	mixedStr := "测试String123"
	// 长字符串
	longStr := "ThisIsAVeryVeryLongCamelCaseStringForPerformanceTestingWithLotsOfWords"
	
	b.Run("Short_Original", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ToSnake(shortStr)
		}
	})
	
	b.Run("Short_Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = OptimizedToSnake(shortStr)
		}
	})
	
	b.Run("ASCII_Original", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Camel2Snake(asciiStr)
		}
	})
	
	b.Run("ASCII_Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = OptimizedCamel2Snake(asciiStr)
		}
	})
	
	b.Run("Mixed_Original", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ToSnake(mixedStr)
		}
	})
	
	b.Run("Mixed_Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = OptimizedToSnake(mixedStr)
		}
	})
	
	b.Run("Long_Original", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ToSnake(longStr)
		}
	})
	
	b.Run("Long_Optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = OptimizedToSnake(longStr)
		}
	})
}