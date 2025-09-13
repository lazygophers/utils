package stringx

import (
	"strings"
	"testing"
	"unicode"
)

// Enhanced string conversion benchmarks
func BenchmarkStringConversion(b *testing.B) {
	data := []byte("Hello, 世界! This is a test string with Unicode characters.")
	str := "Hello, 世界! This is a test string with Unicode characters."

	b.Run("ToString", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ToString(data)
		}
	})

	b.Run("ToBytes", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ToBytes(str)
		}
	})

	b.Run("StandardToString", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = string(data)
		}
	})

	b.Run("StandardToBytes", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = []byte(str)
		}
	})
}

// Case conversion benchmarks
func BenchmarkCaseConversions(b *testing.B) {
	camelCases := []string{"CamelCase", "HTTPSConnection", "XMLHttpParser", "myVariableName", "iPhone15Pro"}
	snakeCases := []string{"snake_case", "http_connection", "xml_parser", "my_variable_name", "_leading_underscore"}
	upperCases := []string{"UPPER_CASE", "HTTP_CONNECTION", "XML_PARSER"}

	b.Run("Camel2Snake", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range camelCases {
				_ = Camel2Snake(tc)
			}
		}
	})

	b.Run("Snake2Camel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range snakeCases {
				_ = Snake2Camel(tc)
			}
		}
	})

	b.Run("Snake2SmallCamel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range upperCases {
				_ = Snake2SmallCamel(tc)
			}
		}
	})
}

func BenchmarkAdvancedConversions(b *testing.B) {
	camelInputs := []string{"SimpleTest", "HTTPSConnection", "XMLHttpParser", "iPhone15Pro", "myVariable123Name"}
	mixedInputs := []string{"simple_test", "my-variable-name", "http.connection", "123numbers", "@symbol#test"}

	b.Run("ToSnake", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range camelInputs {
				_ = ToSnake(tc)
			}
		}
	})

	b.Run("ToKebab", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range camelInputs {
				_ = ToKebab(tc)
			}
		}
	})

	b.Run("ToCamel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range mixedInputs {
				_ = ToCamel(tc)
			}
		}
	})

	b.Run("ToSmallCamel", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range mixedInputs {
				_ = ToSmallCamel(tc)
			}
		}
	})
}

// String utility benchmarks
func BenchmarkStringUtilities(b *testing.B) {
	unicodeStr := "这是一个很长的Unicode测试字符串，用于测试字符串分割功能的性能表现。"
	longStr := "This is a very long test string for benchmarking the shorten function performance."
	reverseStrs := []string{"hello", "你好世界", "Hello_World", "ab😀cd"}

	b.Run("SplitLen", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SplitLen(unicodeStr, 10)
		}
	})

	b.Run("Shorten", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Shorten(longStr, 20)
		}
	})

	b.Run("ShortenShow", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ShortenShow(longStr, 20)
		}
	})

	b.Run("Reverse", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range reverseStrs {
				_ = Reverse(tc)
			}
		}
	})
}

// Random string generation benchmarks
func BenchmarkRandomGeneration(b *testing.B) {
	lengths := []int{10, 50, 100}

	for _, length := range lengths {
		b.Run("Letters_"+string(rune(length/10+'0')), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = RandLetters(length)
			}
		})

		b.Run("Numbers_"+string(rune(length/10+'0')), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = RandNumbers(length)
			}
		})

		b.Run("LetterNumbers_"+string(rune(length/10+'0')), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = RandLetterNumbers(length)
			}
		})
	}
}

// Unicode functions benchmarks
func BenchmarkUnicodeClassification(b *testing.B) {
	digitCases := []string{"123456789", "1234567890123456789012345678901234567890", "12345abc", ""}
	letterCases := []string{"abcdefghijklmnopqrstuvwxyz", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "HelloWorld", "测试字符串"}
	mixedCases := []string{"abcdef", "abc123def", "hello world", "测试字符串123"}
	upperCases := []string{"hello world", "Hello World", "HELLO WORLD", "测试字符串"}

	b.Run("AllDigit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range digitCases {
				_ = AllDigit(tc)
			}
		}
	})

	b.Run("HasDigit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range mixedCases {
				_ = HasDigit(tc)
			}
		}
	})

	b.Run("AllLetter", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range letterCases {
				_ = AllLetter(tc)
			}
		}
	})

	b.Run("HasUpper", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range upperCases {
				_ = HasUpper(tc)
			}
		}
	})
}

// UTF-16 length benchmark
func BenchmarkUtf16Length(b *testing.B) {
	testCases := []string{
		"ASCII text",
		"Unicode 测试",
		"Emoji 😀😁😂",
		"Mixed ASCII + Unicode 测试 + Emoji 😀",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			_ = Utf16Len(tc)
		}
	}
}

// Comparison benchmarks with standard library
func BenchmarkToSnakeVsRegex(b *testing.B) {
	testStr := "HTTPSConnectionXMLParser"

	b.Run("ToSnake", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ToSnake(testStr)
		}
	})

	// Simple implementation using standard library
	b.Run("StandardLib", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := ""
			for i, r := range testStr {
				if unicode.IsUpper(r) && i > 0 {
					result += "_"
				}
				result += strings.ToLower(string(r))
			}
			_ = result
		}
	})
}

func BenchmarkReverseComparison(b *testing.B) {
	testStr := "Hello, 世界! 😀"

	b.Run("Reverse", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Reverse(testStr)
		}
	})

	// Standard library approach
	b.Run("StandardLib", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			runes := []rune(testStr)
			for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
				runes[i], runes[j] = runes[j], runes[i]
			}
			_ = string(runes)
		}
	})
}

// Memory allocation benchmarks
func BenchmarkMemoryAllocation(b *testing.B) {
	testStr := "TestStringForMemoryBenchmark"

	b.Run("ToString", func(b *testing.B) {
		data := []byte(testStr)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ToString(data)
		}
	})

	b.Run("StandardConversion", func(b *testing.B) {
		data := []byte(testStr)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = string(data)
		}
	})
}

// Large data benchmarks
func BenchmarkLargeData(b *testing.B) {
	// Generate large test strings
	largeStr := strings.Repeat("Hello World 你好世界 ", 1000) // ~20KB
	veryLargeStr := strings.Repeat("Test String ", 10000) // ~120KB

	b.Run("ToSnake_Large", func(b *testing.B) {
		testStr := strings.Repeat("CamelCaseString", 100)
		for i := 0; i < b.N; i++ {
			_ = ToSnake(testStr)
		}
	})

	b.Run("Reverse_Large", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Reverse(largeStr)
		}
	})

	b.Run("SplitLen_VeryLarge", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = SplitLen(veryLargeStr, 50)
		}
	})
}
