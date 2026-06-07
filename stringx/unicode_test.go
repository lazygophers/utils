package stringx

import (
	"testing"
)

func TestAllDigit(t *testing.T) {
	type allDigitCase struct {
		input    string
		expected bool
	}
	testCases := []allDigitCase{
		{"123", true},
		{"", true}, // Empty string case
		{"abc", false},
		{"123abc", false},
		{"1a2", false},
		{"0", true},
		{"９８７", true},   // Full-width digits
		{"123 ", false}, // With space
		{"12.3", false}, // With decimal point
		{"-123", false}, // With minus sign
		{"+123", false}, // With plus sign
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := AllDigit(tc.input)
			if result != tc.expected {
				t.Errorf("AllDigit(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestHasDigit(t *testing.T) {
	type hasDigitCase struct {
		input    string
		expected bool
	}
	testCases := []hasDigitCase{
		{"123", true},
		{"", false}, // Empty string case
		{"abc", false},
		{"123abc", true},
		{"a1b", true},
		{"hello", false},
		{"test9", true},
		{"９８７", true}, // Full-width digits
		{"hello world", false},
		{"version 2.0", true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := HasDigit(tc.input)
			if result != tc.expected {
				t.Errorf("HasDigit(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestAllLetter(t *testing.T) {
	type allLetterCase struct {
		input    string
		expected bool
	}
	testCases := []allLetterCase{
		{"abc", true},
		{"", true}, // Empty string case
		{"ABC", true},
		{"aBc", true},
		{"123", false},
		{"abc123", false},
		{"a1b", false},
		{"测试", true},           // Chinese characters are letters
		{"hello world", false}, // Space is not a letter
		{"Ñoël", true},         // Accented characters are letters
		{"hello!", false},      // Punctuation
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := AllLetter(tc.input)
			if result != tc.expected {
				t.Errorf("AllLetter(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestHasLetter(t *testing.T) {
	type hasLetterCase struct {
		input    string
		expected bool
	}
	testCases := []hasLetterCase{
		{"abc", true},
		{"", false}, // Empty string case
		{"123", false},
		{"abc123", true},
		{"123a", true},
		{"!!!!", false},
		{"测试123", true}, // Chinese characters are letters
		{"123@#$", false},
		{"version2", true},
		{"Ñoël123", true}, // Accented characters are letters
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := HasLetter(tc.input)
			if result != tc.expected {
				t.Errorf("HasLetter(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestAllSpace(t *testing.T) {
	type allSpaceCase struct {
		input    string
		expected bool
	}
	testCases := []allSpaceCase{
		{"   ", true},
		{"", true}, // Empty string case
		{"\t\n\r", true},
		{" a ", false},
		{"hello", false},
		{" ", true},
		{"\t", true},
		{"\n", true},
		{"\r", true},
		{" \t\n\r", true},
		{"a b", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := AllSpace(tc.input)
			if result != tc.expected {
				t.Errorf("AllSpace(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestHasSpace(t *testing.T) {
	type hasSpaceCase struct {
		input    string
		expected bool
	}
	testCases := []hasSpaceCase{
		{"   ", true},
		{"", false}, // Empty string case
		{"\t\n\r", true},
		{" a ", true},
		{"hello", false},
		{"hello world", true},
		{"a\tb", true},
		{"a\nb", true},
		{"a\rb", true},
		{"abc123", false},
		{"test string", true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := HasSpace(tc.input)
			if result != tc.expected {
				t.Errorf("HasSpace(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestAllSymbol(t *testing.T) {
	type allSymbolCase struct {
		input    string
		expected bool
	}
	testCases := []allSymbolCase{
		{"♦♠♣♥", true}, // Card symbols
		{"", true},     // Empty string case
		{"$", true},    // Currency symbol
		{"©®™", true},  // Copyright, registered, trademark symbols
		{"abc", false},
		{"123", false},
		{"$100", false}, // Mixed
		{"@#$", false},  // These are punctuation, not symbols
		{"😀", true},     // Emoji is classified as a symbol in Unicode
		{"∑∏∫", true},   // Mathematical symbols
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := AllSymbol(tc.input)
			if result != tc.expected {
				t.Errorf("AllSymbol(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestHasSymbol(t *testing.T) {
	type hasSymbolCase struct {
		input    string
		expected bool
	}
	testCases := []hasSymbolCase{
		{"♦♠♣♥", true}, // Card symbols
		{"", false},    // Empty string case
		{"$100", true}, // Contains currency symbol
		{"hello", false},
		{"test©", true}, // Contains copyright symbol
		{"abc123", false},
		{"Price: $50", true}, // Contains currency symbol
		{"∑total", true},     // Contains mathematical symbol
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := HasSymbol(tc.input)
			if result != tc.expected {
				t.Errorf("HasSymbol(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestAllMark(t *testing.T) {
	type allMarkCase struct {
		input    string
		expected bool
	}
	testCases := []allMarkCase{
		{"̀́̂̃", true}, // Combining marks
		{"", true},     // Empty string case
		{"é", false},   // This is a letter, not just a mark
		{"abc", false},
		{"123", false},
		{"◌́◌̀◌̂", false}, // These contain base characters with marks
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := AllMark(tc.input)
			if result != tc.expected {
				t.Errorf("AllMark(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestHasMark(t *testing.T) {
	type hasMarkCase struct {
		input    string
		expected bool
	}
	testCases := []hasMarkCase{
		{"é", false},      // Precomposed character, not a combining mark
		{"", false},       // Empty string case
		{"cafe", false},   // No combining marks
		{"café", false},   // Precomposed characters, no combining marks
		{"résumé", false}, // Precomposed characters, no combining marks
		{"hello", false},
		{"naïve", false}, // Precomposed characters, no combining marks
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := HasMark(tc.input)
			if result != tc.expected {
				t.Errorf("HasMark(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestAllPunct(t *testing.T) {
	type allPunctCase struct {
		input    string
		expected bool
	}
	testCases := []allPunctCase{
		{"!@#$%", false}, // $ is classified as symbol, not punctuation
		{"", true},       // Empty string case
		{".,;:", true},
		{"()[]{}", true},
		{"abc", false},
		{"123", false},
		{"!@#a", false}, // Mixed
		{"?!", true},
		{"\"'", true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := AllPunct(tc.input)
			if result != tc.expected {
				t.Errorf("AllPunct(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestHasPunct(t *testing.T) {
	type hasPunctCase struct {
		input    string
		expected bool
	}
	testCases := []hasPunctCase{
		{"!@#$%", true},
		{"", false}, // Empty string case
		{"hello!", true},
		{"test", false},
		{"version2.0", true},
		{"abc123", false},
		{"What?", true},
		{"email@domain.com", true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := HasPunct(tc.input)
			if result != tc.expected {
				t.Errorf("HasPunct(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestAllGraphic(t *testing.T) {
	type allGraphicCase struct {
		input    string
		expected bool
	}
	testCases := []allGraphicCase{
		{"abc123", true},
		{"", true}, // Empty string case
		{"hello!", true},
		{"测试", true},
		{"\t", false},           // Tab is not graphic
		{"\n", false},           // Newline is not graphic
		{"hello\tworld", false}, // Contains tab
		{"visible", true},
		{"😀", true},     // Emoji is graphic
		{"\x00", false}, // Null character is not graphic
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := AllGraphic(tc.input)
			if result != tc.expected {
				t.Errorf("AllGraphic(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestHasGraphic(t *testing.T) {
	type hasGraphicCase struct {
		input    string
		expected bool
	}
	testCases := []hasGraphicCase{
		{"abc123", true},
		{"", false},     // Empty string case
		{"\t\n", false}, // Only control characters
		{"\tabc", true}, // Has graphic characters
		{"hello world", true},
		{"\x00\x01", false}, // Only control characters
		{"测试", true},
		{"😀", true}, // Emoji is graphic
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := HasGraphic(tc.input)
			if result != tc.expected {
				t.Errorf("HasGraphic(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestAllPrint(t *testing.T) {
	type allPrintCase struct {
		input    string
		expected bool
	}
	testCases := []allPrintCase{
		{"abc123", true},
		{"", true}, // Empty string case
		{"hello!", true},
		{"hello world", true}, // Space is printable
		{"\t", false},         // Tab is not printable
		{"\n", false},         // Newline is not printable
		{"测试", true},
		{"😀", true},     // Emoji is printable
		{"\x00", false}, // Null character is not printable
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := AllPrint(tc.input)
			if result != tc.expected {
				t.Errorf("AllPrint(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestHasPrint(t *testing.T) {
	type hasPrintCase struct {
		input    string
		expected bool
	}
	testCases := []hasPrintCase{
		{"abc123", true},
		{"", false},     // Empty string case
		{"\t\n", false}, // Only control characters
		{"\tabc", true}, // Has printable characters
		{"hello world", true},
		{"\x00a", true}, // Has printable character
		{"测试", true},
		{"\x00\x01", false}, // Only control characters
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := HasPrint(tc.input)
			if result != tc.expected {
				t.Errorf("HasPrint(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestAllControl(t *testing.T) {
	type allControlCase struct {
		input    string
		expected bool
	}
	testCases := []allControlCase{
		{"\t\n\r", true},
		{"", true}, // Empty string case
		{"abc", false},
		{"\tabc", false},   // Mixed
		{"\x00\x01", true}, // Null and control characters
		{"hello", false},
		{" ", false}, // Space is not a control character
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := AllControl(tc.input)
			if result != tc.expected {
				t.Errorf("AllControl(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestHasControl(t *testing.T) {
	type hasControlCase struct {
		input    string
		expected bool
	}
	testCases := []hasControlCase{
		{"\t\n\r", true},
		{"", false}, // Empty string case
		{"abc", false},
		{"\tabc", true},   // Has control character
		{"hello\n", true}, // Has newline
		{"normal text", false},
		{"\x00test", true}, // Has null character
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := HasControl(tc.input)
			if result != tc.expected {
				t.Errorf("HasControl(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestAllUpper(t *testing.T) {
	type allUpperCase struct {
		input    string
		expected bool
	}
	testCases := []allUpperCase{
		{"ABC", true},
		{"", true}, // Empty string case
		{"abc", false},
		{"ABC123", false}, // Numbers don't have case
		{"HELLO", true},
		{"Hello", false},  // Mixed case
		{"ÑOËL", true},    // Accented uppercase
		{"测试", false},     // Chinese characters don't have case
		{"123", false},    // Numbers only
		{"HELLO!", false}, // Contains punctuation
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := AllUpper(tc.input)
			if result != tc.expected {
				t.Errorf("AllUpper(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestHasUpper(t *testing.T) {
	type hasUpperCase struct {
		input    string
		expected bool
	}
	testCases := []hasUpperCase{
		{"ABC", true},
		{"", false}, // Empty string case
		{"abc", false},
		{"Abc", true},         // Has uppercase
		{"hello World", true}, // Has uppercase
		{"hello", false},
		{"123ABC", true}, // Has uppercase
		{"测试", false},    // Chinese characters don't have case
		{"Ñoël", true},   // Has accented uppercase
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := HasUpper(tc.input)
			if result != tc.expected {
				t.Errorf("HasUpper(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestAllLower(t *testing.T) {
	type allLowerCase struct {
		input    string
		expected bool
	}
	testCases := []allLowerCase{
		{"abc", true},
		{"", true}, // Empty string case
		{"ABC", false},
		{"abc123", false}, // Numbers don't have case
		{"hello", true},
		{"Hello", false},  // Mixed case
		{"ñoël", true},    // Accented lowercase
		{"测试", false},     // Chinese characters don't have case
		{"123", false},    // Numbers only
		{"hello!", false}, // Contains punctuation
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := AllLower(tc.input)
			if result != tc.expected {
				t.Errorf("AllLower(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestHasLower(t *testing.T) {
	type hasLowerCase struct {
		input    string
		expected bool
	}
	testCases := []hasLowerCase{
		{"abc", true},
		{"", false}, // Empty string case
		{"ABC", false},
		{"Abc", true},         // Has lowercase
		{"HELLO world", true}, // Has lowercase
		{"HELLO", false},
		{"123abc", true}, // Has lowercase
		{"测试", false},    // Chinese characters don't have case
		{"Ñoël", true},   // Has accented lowercase
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := HasLower(tc.input)
			if result != tc.expected {
				t.Errorf("HasLower(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestAllTitle(t *testing.T) {
	type allTitleCase struct {
		input    string
		expected bool
	}
	testCases := []allTitleCase{
		{"ǅǈǋ", true},  // Title case characters (rare)
		{"", true},     // Empty string case
		{"ABC", false}, // Uppercase, not title case
		{"abc", false}, // Lowercase
		{"Abc", false}, // Not title case
		{"123", false}, // Numbers
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := AllTitle(tc.input)
			if result != tc.expected {
				t.Errorf("AllTitle(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestHasTitle(t *testing.T) {
	type hasTitleCase struct {
		input    string
		expected bool
	}
	testCases := []hasTitleCase{
		{"ǅǈǋ", true},  // Title case characters (rare)
		{"", false},    // Empty string case
		{"ABC", false}, // No title case
		{"abcǅ", true}, // Has title case
		{"hello", false},
		{"123", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := HasTitle(tc.input)
			if result != tc.expected {
				t.Errorf("HasTitle(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestAllLetterOrDigit(t *testing.T) {
	type allLetterOrDigitCase struct {
		input    string
		expected bool
	}
	testCases := []allLetterOrDigitCase{
		{"abc123", true},
		{"", true}, // Empty string case
		{"abc", true},
		{"123", true},
		{"abc!", false},        // Contains punctuation
		{"hello world", false}, // Contains space
		{"测试123", true},        // Chinese + numbers
		{"a1b2c3", true},
		{"hello@world", false}, // Contains punctuation
		{"Ñoël123", true},      // Accented letters + numbers
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := AllLetterOrDigit(tc.input)
			if result != tc.expected {
				t.Errorf("AllLetterOrDigit(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestHasLetterOrDigit(t *testing.T) {
	type hasLetterOrDigitCase struct {
		input    string
		expected bool
	}
	testCases := []hasLetterOrDigitCase{
		{"abc123", true},
		{"", false},      // Empty string case
		{"!!!", false},   // Only punctuation
		{"!@#a", true},   // Has letter
		{"!@#1", true},   // Has digit
		{" \t\n", false}, // Only whitespace
		{"测试", true},     // Chinese characters are letters
		{"😀", false},     // Emoji is not letter or digit
		{" a ", true},    // Has letter
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := HasLetterOrDigit(tc.input)
			if result != tc.expected {
				t.Errorf("HasLetterOrDigit(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkAllDigit(b *testing.B) {
	s := "1234567890123456789012345678901234567890"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AllDigit(s)
	}
}

func BenchmarkHasDigit(b *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyz1"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HasDigit(s)
	}
}

func BenchmarkAllLetter(b *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AllLetter(s)
	}
}

func BenchmarkAllLetterOrDigit(b *testing.B) {
	s := "abc123def456ghi789jklmnopqrstuvwxyz0123456789"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AllLetterOrDigit(s)
	}
}

// Edge case tests
func TestUnicodeEdgeCases(t *testing.T) {
	t.Run("empty_string_behavior", func(t *testing.T) {
		// All "All*" functions should return true for empty strings
		// All "Has*" functions should return false for empty strings

		allFunctions := []func(string) bool{
			AllDigit, AllLetter, AllSpace, AllSymbol, AllMark,
			AllPunct, AllGraphic, AllPrint, AllControl,
			AllUpper, AllLower, AllTitle, AllLetterOrDigit,
		}

		for i, fn := range allFunctions {
			if !fn("") {
				t.Errorf("All* function %d should return true for empty string", i)
			}
		}

		hasFunctions := []func(string) bool{
			HasDigit, HasLetter, HasSpace, HasSymbol, HasMark,
			HasPunct, HasGraphic, HasPrint, HasControl,
			HasUpper, HasLower, HasTitle, HasLetterOrDigit,
		}

		for i, fn := range hasFunctions {
			if fn("") {
				t.Errorf("Has* function %d should return false for empty string", i)
			}
		}
	})

	t.Run("unicode_normalization", func(t *testing.T) {
		// Test with different Unicode normalization forms
		s1 := "é"       // Single character with accent (precomposed)
		s2 := "e\u0301" // 'e' + combining acute accent (decomposed)

		// Both should be treated as having letters
		if !HasLetter(s1) || !HasLetter(s2) {
			t.Error("Both normalized forms should have letters")
		}

		// Only the decomposed form (s2) has actual combining marks
		// The precomposed form (s1) is treated as a single letter
		if !HasMark(s2) {
			t.Error("Decomposed form should have marks")
		}

		// The precomposed form may or may not have marks depending on Unicode implementation
		// This is acceptable behavior - test passes if either has marks
		if !HasMark(s1) && !HasMark(s2) {
			t.Error("At least one form should have detectable marks")
		}
	})

	t.Run("surrogate_pairs", func(t *testing.T) {
		// Test with emoji that use surrogate pairs
		emoji := "😀🎉🔥"

		// These should be treated as graphic and printable
		if !AllGraphic(emoji) {
			t.Error("Emoji should be graphic")
		}

		if !AllPrint(emoji) {
			t.Error("Emoji should be printable")
		}

		// But not letters or digits
		if AllLetter(emoji) {
			t.Error("Emoji should not be letters")
		}

		if AllDigit(emoji) {
			t.Error("Emoji should not be digits")
		}
	})
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
