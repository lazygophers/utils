package stringx

import (
	"testing"
	"unicode/utf16"
)

func TestUtf16Len(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{} // Can be string, []rune, or []byte
		expected int
	}{
		// String inputs
		{"empty_string", "", 0},
		{"ascii_string", "hello", 5},
		{"ascii_with_numbers", "hello123", 8},
		{"unicode_basic", "测试", 2},
		{"mixed_ascii_unicode", "hello测试", 7},
		
		// Emoji and surrogate pairs
		{"simple_emoji", "😀", 2}, // This emoji requires surrogate pair
		{"multiple_emoji", "😀😁😂", 6}, // 3 emoji, each needs 2 UTF-16 code units
		{"emoji_with_text", "hello😀world", 12}, // 5 + 2 + 5
		
		// Special Unicode characters
		{"combining_characters", "é", 1}, // Single precomposed character
		{"combining_separate", "e\u0301", 2}, // 'e' + combining acute accent
		{"high_unicode", "𝕳𝖊𝖑𝖑𝖔", 10}, // Mathematical alphanumeric symbols (surrogate pairs)
		
		// []rune inputs
		{"rune_slice_empty", []rune{}, 0},
		{"rune_slice_ascii", []rune("hello"), 5},
		{"rune_slice_unicode", []rune("测试"), 2},
		{"rune_slice_emoji", []rune("😀"), 2},
		
		// []byte inputs
		{"byte_slice_empty", []byte{}, 0},
		{"byte_slice_ascii", []byte("hello"), 5},
		{"byte_slice_unicode", []byte("测试"), 2},
		{"byte_slice_emoji", []byte("😀"), 2},
		
		// Edge cases
		{"null_character", "\x00", 1},
		{"control_characters", "\t\n\r", 3},
		{"long_ascii", "abcdefghijklmnopqrstuvwxyz", 26},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result int
			switch v := tc.input.(type) {
			case string:
				result = Utf16Len(v)
			case []rune:
				result = Utf16Len(v)
			case []byte:
				result = Utf16Len(v)
			default:
				t.Fatalf("Unsupported input type: %T", v)
			}
			
			if result != tc.expected {
				t.Errorf("Utf16Len(%v) = %d, expected %d", tc.input, result, tc.expected)
			}
		})
	}
}

func TestUtf16LenConsistency(t *testing.T) {
	testStrings := []string{
		"",
		"hello",
		"测试",
		"😀😁😂",
		"hello😀world",
		"mixed测试content😀",
		"𝕳𝖊𝖑𝖑𝖔", // Mathematical alphanumeric
		"Ñoël café",
		"русский язык",
		"العربية",
	}

	for _, s := range testStrings {
		t.Run(s, func(t *testing.T) {
			// Test that all three input types give the same result
			strResult := Utf16Len(s)
			runeResult := Utf16Len([]rune(s))
			byteResult := Utf16Len([]byte(s))
			
			if strResult != runeResult {
				t.Errorf("String and rune results differ: string=%d, rune=%d", strResult, runeResult)
			}
			
			if strResult != byteResult {
				t.Errorf("String and byte results differ: string=%d, byte=%d", strResult, byteResult)
			}
			
			// Also verify against standard library
			expected := len(utf16.Encode([]rune(s)))
			if strResult != expected {
				t.Errorf("Result differs from stdlib: got=%d, expected=%d", strResult, expected)
			}
		})
	}
}

func TestUtf16LenAgainstStdlib(t *testing.T) {
	testCases := []string{
		"",
		"a",
		"hello world",
		"测试字符串",
		"🙂😀🎉",
		"hello🙂world",
		"𝒽𝑒𝓁𝓁𝑜", // Mathematical script characters
		"🇺🇸🇨🇳", // Flag emojis (each is 2 surrogate pairs = 4 UTF-16 units)
		"👨‍👩‍👧‍👦", // Complex emoji sequence
		"Åpfel", // With combining characters
	}

	for _, s := range testCases {
		t.Run("", func(t *testing.T) {
			ourResult := Utf16Len(s)
			stdResult := len(utf16.Encode([]rune(s)))
			
			if ourResult != stdResult {
				t.Errorf("Utf16Len(%q) = %d, stdlib result = %d", s, ourResult, stdResult)
			}
		})
	}
}

func TestUtf16LenEdgeCases(t *testing.T) {
	t.Run("nil_byte_slice", func(t *testing.T) {
		result := Utf16Len([]byte(nil))
		if result != 0 {
			t.Errorf("Utf16Len(nil []byte) = %d, expected 0", result)
		}
	})

	t.Run("nil_rune_slice", func(t *testing.T) {
		result := Utf16Len([]rune(nil))
		if result != 0 {
			t.Errorf("Utf16Len(nil []rune) = %d, expected 0", result)
		}
	})

	t.Run("single_rune", func(t *testing.T) {
		// Test various single runes
		testRunes := []struct {
			r        rune
			expected int
		}{
			{'a', 1},          // ASCII
			{'测', 1},          // CJK
			{'😀', 2},          // Emoji (surrogate pair)
			{'𝒽', 2},          // Mathematical script (surrogate pair)
			{'\x00', 1},       // Null
			{'\u0301', 1},     // Combining acute accent
		}

		for _, tr := range testRunes {
			s := string(tr.r)
			result := Utf16Len(s)
			if result != tr.expected {
				t.Errorf("Utf16Len(%q) = %d, expected %d", s, result, tr.expected)
			}
		}
	})

	t.Run("very_long_string", func(t *testing.T) {
		// Test with a very long string
		longStr := ""
		for i := 0; i < 10000; i++ {
			longStr += "a"
		}
		
		result := Utf16Len(longStr)
		expected := 10000
		if result != expected {
			t.Errorf("Utf16Len(long string) = %d, expected %d", result, expected)
		}
	})

	t.Run("mixed_surrogate_pairs", func(t *testing.T) {
		// String with mix of 1-unit and 2-unit characters
		s := "a😀b测c🎉d"
		// a(1) + 😀(2) + b(1) + 测(1) + c(1) + 🎉(2) + d(1) = 9
		expected := 9
		result := Utf16Len(s)
		if result != expected {
			t.Errorf("Utf16Len(%q) = %d, expected %d", s, result, expected)
		}
	})
}

func TestUtf16LenPerformance(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping performance test in short mode")
	}

	// Test performance with various string types
	testCases := []struct {
		name string
		str  string
	}{
		{"ascii_1000", generateString("abcdefghijklmnopqrstuvwxyz", 1000)},
		{"unicode_1000", generateString("测试字符串", 1000)},
		{"emoji_100", generateString("😀🎉🔥", 100)},
		{"mixed_500", generateString("hello测试😀world", 500)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Warm up
			for i := 0; i < 10; i++ {
				Utf16Len(tc.str)
			}
			
			// Measure
			iterations := 1000
			for i := 0; i < iterations; i++ {
				Utf16Len(tc.str)
			}
			
			t.Logf("Completed %d iterations for %s", iterations, tc.name)
		})
	}
}

// Helper function to generate test strings
func generateString(pattern string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += pattern
	}
	return result
}

// Benchmark tests
func BenchmarkUtf16Len(b *testing.B) {
	testCases := []struct {
		name string
		str  string
	}{
		{"ascii_short", "hello world"},
		{"ascii_long", generateString("hello world", 100)},
		{"unicode_short", "测试字符串"},
		{"unicode_long", generateString("测试字符串", 100)},
		{"emoji_short", "😀🎉🔥"},
		{"emoji_long", generateString("😀🎉🔥", 50)},
		{"mixed", "hello测试😀world"},
		{"mixed_long", generateString("hello测试😀world", 50)},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Utf16Len(tc.str)
			}
		})
	}
}

func BenchmarkUtf16LenVsStdlib(b *testing.B) {
	testStr := "hello测试😀world🎉"
	
	b.Run("our_implementation", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Utf16Len(testStr)
		}
	})
	
	b.Run("stdlib_direct", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = len(utf16.Encode([]rune(testStr)))
		}
	})
}

func BenchmarkUtf16LenTypes(b *testing.B) {
	testStr := "hello测试😀world🎉"
	testRunes := []rune(testStr)
	testBytes := []byte(testStr)
	
	b.Run("string_input", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Utf16Len(testStr)
		}
	})
	
	b.Run("rune_slice_input", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Utf16Len(testRunes)
		}
	})
	
	b.Run("byte_slice_input", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Utf16Len(testBytes)
		}
	})
}

// Comprehensive Unicode test
func TestUtf16LenUnicodeCompliance(t *testing.T) {
	// Test various Unicode categories and edge cases
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		// Basic Multilingual Plane (BMP)
		{"latin_basic", "Hello", 5},
		{"latin_extended", "Héllö", 5},
		{"cyrillic", "Привет", 6},
		{"arabic", "مرحبا", 5},
		{"hebrew", "שלום", 4},
		{"cjk", "你好", 2},
		{"hiragana", "こんにちは", 5},
		{"katakana", "コンニチハ", 5},
		
		// Supplementary planes (require surrogate pairs)
		{"emoji_faces", "😀😃😄", 6},    // 3 emoji × 2 units each
		{"emoji_objects", "🏠🚗✈️", 6},  // house(2) + car(2) + plane(2)
		{"mathematical", "𝒽𝑒𝓁𝓁𝑜", 10}, // 5 characters × 2 units each
		{"musical", "𝄞𝄢𝄪", 6},          // 3 musical symbols × 2 units each
		
		// Complex emoji sequences
		{"flag_emoji", "🇺🇸", 4},        // US flag = 2 regional indicator symbols × 2 units each
		{"skin_tone", "👋🏽", 4},        // Wave + skin tone modifier = 2 + 2 units
		{"zwj_sequence", "👨‍💻", 5},     // Man + ZWJ + Computer = 2 + 1 + 2 units
		
		// Combining characters
		{"combining_acute", "é", 1},     // Single precomposed
		{"combining_separate", "e\u0301", 2}, // e + combining acute
		{"multiple_combining", "ê̂", 2}, // e + circumflex (combined)
		
		// Control and special characters
		{"control_chars", "\t\n\r", 3},
		{"bom", "\uFEFF", 1},           // Byte Order Mark
		{"replacement", "\uFFFD", 1},    // Replacement character
		
		// Mixed content
		{"mixed_simple", "Hello世界", 7},
		{"mixed_complex", "Hi👋测试😀", 8}, // Hi(2) + wave(2) + test(2) + smile(2)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Utf16Len(tc.input)
			if result != tc.expected {
				// Also show the standard library result for comparison
				stdResult := len(utf16.Encode([]rune(tc.input)))
				t.Errorf("Utf16Len(%q) = %d, expected %d (stdlib: %d)", 
					tc.input, result, tc.expected, stdResult)
			}
		})
	}
}