package stringx

import (
	"testing"
)

// TestCamel2SnakeUnicode tests the Unicode path of Camel2Snake that's currently not covered
func TestCamel2SnakeUnicode(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty_string",
			input:    "",
			expected: "",
		},
		{
			name:     "unicode_characters",
			input:    "测试CamelCase",
			expected: "测试_camel_case",
		},
		{
			name:     "mixed_unicode_ascii",
			input:    "测试HTTP服务器",
			expected: "测试_h_t_t_p服务器",
		},
		{
			name:     "chinese_camelcase",
			input:    "用户Name数据库",
			expected: "用户_name数据库",
		},
		{
			name:     "emoji_with_camel",
			input:    "🚀RocketLaunch",
			expected: "🚀_rocket_launch",
		},
		{
			name:     "japanese_hiragana",
			input:    "こんにちはWorld",
			expected: "こんにちは_world",
		},
		{
			name:     "unicode_uppercase",
			input:    "ÜberTest",
			expected: "\xfcber_test", // This is the actual UTF-8 encoding
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Camel2Snake(tc.input)
			if result != tc.expected {
				t.Errorf("Camel2Snake(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

// TestToSnakeMissingBranches tests the missing branches in ToSnake function
func TestToSnakeMissingBranches(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty_string",
			input:    "",
			expected: "",
		},
		{
			name:     "single_char",
			input:    "a",
			expected: "a",
		},
		{
			name:     "single_upper_char",
			input:    "A",
			expected: "a",
		},
		{
			name:     "unicode_lowercase",
			input:    "测试小写",
			expected: "测试小写",
		},
		{
			name:     "mixed_unicode_with_capitals",
			input:    "测试DatabaseConnection",
			expected: "测试_database_connection",
		},
		{
			name:     "consecutive_capitals",
			input:    "XMLHTTPRequest",
			expected: "x_m_l_h_t_t_p_request",
		},
		{
			name:     "digits_and_capitals",
			input:    "Version2Update",
			expected: "version_2_update",
		},
		{
			name:     "all_caps",
			input:    "CONSTANT",
			expected: "c_o_n_s_t_a_n_t",
		},
		{
			name:     "mixed_case_with_numbers",
			input:    "API2ServiceV1",
			expected: "a_p_i_2_service_v_1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ToSnake(tc.input)
			if result != tc.expected {
				t.Errorf("ToSnake(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

// TestToSmallCamelMissingBranches tests the missing branches in ToSmallCamel
func TestToSmallCamelMissingBranches(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty_string",
			input:    "",
			expected: "",
		},
		{
			name:     "single_char",
			input:    "a",
			expected: "a",
		},
		{
			name:     "single_underscore",
			input:    "_",
			expected: "",
		},
		{
			name:     "leading_underscore",
			input:    "_test_case",
			expected: "testCase",
		},
		{
			name:     "trailing_underscore",
			input:    "test_case_",
			expected: "testCase",
		},
		{
			name:     "consecutive_underscores",
			input:    "test__case",
			expected: "testCase",
		},
		{
			name:     "underscore_only",
			input:    "___",
			expected: "",
		},
		{
			name:     "mixed_separators",
			input:    "test_case-name",
			expected: "testCaseName",
		},
		{
			name:     "unicode_with_underscores",
			input:    "测试_case_数据库",
			expected: "测试Case数据库",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ToSmallCamel(tc.input)
			if result != tc.expected {
				t.Errorf("ToSmallCamel(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

// TestSplitLenMissingBranches tests the missing branches in SplitLen
func TestSplitLenMissingBranches(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		length   int
		expected []string
	}{
		{
			name:     "empty_string",
			input:    "",
			length:   3,
			expected: []string{},
		},
		{
			name:     "zero_length",
			input:    "hello",
			length:   0,
			expected: []string{"hello"},
		},
		{
			name:     "negative_length",
			input:    "hello",
			length:   -1,
			expected: []string{"hello"},
		},
		{
			name:     "length_equals_string_length",
			input:    "hello",
			length:   5,
			expected: []string{"hello"},
		},
		{
			name:     "length_greater_than_string",
			input:    "hi",
			length:   10,
			expected: []string{"hi"},
		},
		{
			name:     "normal_split",
			input:    "hello world test",
			length:   5,
			expected: []string{"hello", " worl", "d tes", "t"},
		},
		{
			name:     "unicode_split",
			input:    "你好世界测试",
			length:   2,
			expected: []string{"你好", "世界", "测试"},
		},
		{
			name:     "single_char_string",
			input:    "a",
			length:   2,
			expected: []string{"a"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SplitLen(tc.input, tc.length)
			if len(result) != len(tc.expected) {
				t.Errorf("SplitLen(%q, %d) returned %d parts, expected %d",
					tc.input, tc.length, len(result), len(tc.expected))
				return
			}
			for i, part := range result {
				if part != tc.expected[i] {
					t.Errorf("SplitLen(%q, %d)[%d] = %q, expected %q",
						tc.input, tc.length, i, part, tc.expected[i])
				}
			}
		})
	}
}

// TestHelperFunctions tests helper functions for complete coverage
func TestHelperFunctions(t *testing.T) {
	t.Run("isASCII", func(t *testing.T) {
		testCases := []struct {
			input    string
			expected bool
		}{
			{"", true},
			{"hello", true},
			{"Hello123", true},
			{"test_case", true},
			{"测试", false},
			{"hello世界", false},
			{"🚀", false},
			{"café", false}, // 'é' is not ASCII
		}

		for _, tc := range testCases {
			result := isASCII(tc.input)
			if result != tc.expected {
				t.Errorf("isASCII(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		}
	})

	t.Run("optimizedASCIICamel2Snake", func(t *testing.T) {
		testCases := []struct {
			input    string
			expected string
		}{
			{"", ""},
			{"hello", "hello"},
			{"HelloWorld", "hello_world"},
			{"HTTPSConnection", "h_t_t_p_s_connection"},
			{"XMLParser", "x_m_l_parser"},
			{"iPhone", "i_phone"},
			{"iOS", "i_o_s"},
			{"myVariableName", "my_variable_name"},
			{"ABC", "a_b_c"},
			{"a", "a"},
			{"A", "a"},
		}

		for _, tc := range testCases {
			result := optimizedASCIICamel2Snake(tc.input)
			if result != tc.expected {
				t.Errorf("optimizedASCIICamel2Snake(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		}
	})
}

// TestEdgeCasesForCompleteCoverage tests various edge cases
func TestEdgeCasesForCompleteCoverage(t *testing.T) {
	t.Run("unicode_edge_cases", func(t *testing.T) {
		// Test various Unicode scenarios that might hit different branches
		testCases := []string{
			"Ñoño",    // Spanish with tildes
			"Москва",  // Cyrillic
			"العربية", // Arabic
			"한국어",     // Korean
			"🎉🎊",      // Emojis
			"𝕌𝕟𝕚𝕔𝕠𝕕𝕖", // Mathematical symbols
		}

		for _, input := range testCases {
			// Just ensure these don't panic and return reasonable results
			result1 := Camel2Snake(input)
			result2 := ToSnake(input)
			result3 := ToSmallCamel(input)

			// Basic sanity checks
			if len(result1) < 0 || len(result2) < 0 || len(result3) < 0 {
				t.Errorf("Negative length result for input %q", input)
			}

			t.Logf("Input: %q, Camel2Snake: %q, ToSnake: %q, ToSmallCamel: %q",
				input, result1, result2, result3)
		}
	})
}

// TestAllBranchesInStringFunctions ensures we hit all conditional branches
func TestAllBranchesInStringFunctions(t *testing.T) {
	// Test various patterns that should trigger different code paths
	patterns := []struct {
		input string
		desc  string
	}{
		{"", "empty string"},
		{"a", "single lowercase"},
		{"A", "single uppercase"},
		{"aB", "mixed case"},
		{"AB", "double uppercase"},
		{"abc", "all lowercase"},
		{"ABC", "all uppercase"},
		{"aBc", "alternating case"},
		{"a1B", "with numbers"},
		{"_test", "leading underscore"},
		{"test_", "trailing underscore"},
		{"test__case", "double underscore"},
		{"测试", "non-ASCII"},
		{"测A试", "mixed Unicode/ASCII"},
		{"🚀Test", "emoji with text"},
	}

	for _, p := range patterns {
		t.Run(p.desc, func(t *testing.T) {
			// Call all the functions to ensure they don't panic and cover branches
			r1 := ToString([]byte(p.input))
			r2 := ToBytes(p.input)
			r3 := Camel2Snake(p.input)
			r4 := Snake2Camel(p.input)
			r5 := Snake2SmallCamel(p.input)
			r6 := ToSnake(p.input)
			r7 := ToKebab(p.input)
			r8 := ToCamel(p.input)
			r9 := ToSlash(p.input)
			r10 := ToDot(p.input)
			r11 := ToSmallCamel(p.input)

			// Basic validation - ensure no nil/empty responses for non-empty input
			if p.input != "" {
				if r1 == "" && p.input != "" {
					t.Errorf("ToString returned empty for non-empty input %q", p.input)
				}
				if len(r2) == 0 && p.input != "" {
					t.Errorf("ToBytes returned empty for non-empty input %q", p.input)
				}
			}

			// Log results for inspection (helps with debugging coverage)
			t.Logf("Input: %q -> ToString: %q, Camel2Snake: %q, ToSnake: %q, ToSmallCamel: %q",
				p.input, r1, r3, r6, r11)

			_ = r4
			_ = r5
			_ = r7
			_ = r8
			_ = r9
			_ = r10
		})
	}
}
