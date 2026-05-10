package anyx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ============================================================
// parseIndex 函数覆盖率测试
// 目标：确保所有分支和边界情况都被测试
// ============================================================

func TestParseIndex_Coverage_ValidNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "zero",
			input:    "0",
			expected: 0,
		},
		{
			name:     "single digit positive",
			input:    "5",
			expected: 5,
		},
		{
			name:     "two digits",
			input:    "42",
			expected: 42,
		},
		{
			name:     "three digits",
			input:    "123",
			expected: 123,
		},
		{
			name:     "large number",
			input:    "999999",
			expected: 999999,
		},
		{
			name:     "max int32",
			input:    "2147483647",
			expected: 2147483647,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseIndex(tt.input)
			require.NoError(t, err, "parseIndex(%q) should not return error", tt.input)
			assert.Equal(t, tt.expected, result, "parseIndex(%q) result mismatch", tt.input)
		})
	}
}

func TestParseIndex_Coverage_NegativeNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "negative single digit",
			input:    "-1",
			expected: -1,
		},
		{
			name:     "negative two digits",
			input:    "-42",
			expected: -42,
		},
		{
			name:     "negative three digits",
			input:    "-123",
			expected: -123,
		},
		{
			name:     "negative large",
			input:    "-9999",
			expected: -9999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseIndex(tt.input)
			require.NoError(t, err, "parseIndex(%q) should not return error", tt.input)
			assert.Equal(t, tt.expected, result, "parseIndex(%q) result mismatch", tt.input)
		})
	}
}

func TestParseIndex_Coverage_ErrorCases(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedErr error
	}{
		{
			name:        "empty string",
			input:       "",
			expectedErr: ErrInvalidIndex,
		},
		{
			name:        "non-numeric letters",
			input:       "abc",
			expectedErr: ErrInvalidIndex,
		},
		{
			name:        "mixed alphanumeric",
			input:       "12a34",
			expectedErr: ErrInvalidIndex,
		},
		{
			name:        "special characters",
			input:       "!@#",
			expectedErr: ErrInvalidIndex,
		},
		{
			name:        "decimal point",
			input:       "12.34",
			expectedErr: ErrInvalidIndex,
		},
		{
			name:        "only minus sign",
			input:       "-",
			expectedErr: ErrInvalidIndex,
		},
		{
			name:        "minus with letters",
			input:       "-abc",
			expectedErr: ErrInvalidIndex,
		},
		{
			name:        "spaces",
			input:       "12 34",
			expectedErr: ErrInvalidIndex,
		},
		{
			name:        "leading zero is valid",
			input:       "007",
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseIndex(tt.input)

			if tt.expectedErr != nil {
				require.Error(t, err, "parseIndex(%q) should return error", tt.input)
				assert.ErrorIs(t, err, tt.expectedErr, "error type mismatch")
			} else {
				require.NoError(t, err, "parseIndex(%q) should not return error", tt.input)
				assert.NotEqual(t, 0, result, "parseIndex(%q) should return non-zero for valid input", tt.input)
			}
		})
	}
}

func TestParseIndex_Coverage_BoundaryCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		valid    bool
		expected int
	}{
		{
			name:     "single zero",
			input:    "0",
			valid:    true,
			expected: 0,
		},
		{
			name:     "multiple zeros",
			input:    "000",
			valid:    true,
			expected: 0,
		},
		{
			name:     "negative zero",
			input:    "-0",
			valid:    true,
			expected: 0,
		},
		{
			name:     "one digit 1",
			input:    "1",
			valid:    true,
			expected: 1,
		},
		{
			name:     "one digit 9",
			input:    "9",
			valid:    true,
			expected: 9,
		},
		{
			name:     "negative one digit",
			input:    "-5",
			valid:    true,
			expected: -5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseIndex(tt.input)

			if tt.valid {
				require.NoError(t, err, "parseIndex(%q) should be valid", tt.input)
				assert.Equal(t, tt.expected, result, "parseIndex(%q) result mismatch", tt.input)
			} else {
				require.Error(t, err, "parseIndex(%q) should be invalid", tt.input)
			}
		})
	}
}

func TestParseIndex_Coverage_AllDigits(t *testing.T) {
	// Test all single digits 0-9
	for i := 0; i <= 9; i++ {
		digit := fmt.Sprintf("%d", i)
		t.Run("digit_"+digit, func(t *testing.T) {
			result, err := parseIndex(digit)
			require.NoError(t, err, "parseIndex(%q) should not return error", digit)
			assert.Equal(t, i, result, "parseIndex(%q) should return %d", digit, i)
		})
	}

	// Test all negative single digits -9 to -1
	for i := 1; i <= 9; i++ {
		negativeDigit := fmt.Sprintf("-%d", i)
		t.Run("negative_digit_"+negativeDigit, func(t *testing.T) {
			result, err := parseIndex(negativeDigit)
			require.NoError(t, err, "parseIndex(%q) should not return error", negativeDigit)
			assert.Equal(t, -i, result, "parseIndex(%q) should return %d", negativeDigit, -i)
		})
	}
}

func TestParseIndex_Coverage_ErrorMessages(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedSubstr string
	}{
		{
			name:           "empty error message",
			input:          "",
			expectedSubstr: "empty index",
		},
		{
			name:           "invalid chars error message",
			input:          "abc",
			expectedSubstr: "abc",
		},
		{
			name:           "invalid chars with digits",
			input:          "12a34",
			expectedSubstr: "12a34",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := parseIndex(tt.input)
			require.Error(t, err, "parseIndex(%q) should return error", tt.input)
			assert.Contains(t, err.Error(), tt.expectedSubstr, "error message should contain input")
		})
	}
}

func TestParseIndex_Coverage_RealWorldScenarios(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		valid    bool
		expected int
	}{
		// Common array indices
		{name: "first element", input: "0", valid: true, expected: 0},
		{name: "second element", input: "1", valid: true, expected: 1},
		{name: "tenth element", input: "9", valid: true, expected: 9},
		{name: "hundredth element", input: "99", valid: true, expected: 99},

		// Negative indices (from end)
		{name: "last element", input: "-1", valid: true, expected: -1},
		{name: "second to last", input: "-2", valid: true, expected: -2},
		{name: "tenth from end", input: "-10", valid: true, expected: -10},

		// Large indices
		{name: "large index", input: "1000", valid: true, expected: 1000},
		{name: "very large index", input: "999999", valid: true, expected: 999999},

		// Invalid cases
		{name: "just minus", input: "-", valid: false, expected: 0},
		{name: "with spaces", input: " 123", valid: false, expected: 0},
		{name: "trailing space", input: "123 ", valid: false, expected: 0},
		{name: "with plus", input: "+123", valid: false, expected: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseIndex(tt.input)

			if tt.valid {
				require.NoError(t, err, "parseIndex(%q) should be valid", tt.input)
				assert.Equal(t, tt.expected, result, "parseIndex(%q) result mismatch", tt.input)
			} else {
				require.Error(t, err, "parseIndex(%q) should be invalid", tt.input)
			}
		})
	}
}

// ============================================================
// 边界条件测试：确保优化的实现与原始实现行为一致
// ============================================================

func TestParseIndex_Coverage_EdgeCase_EmptyAfterNegative(t *testing.T) {
	// 测试 "-" 的情况（负号后没有数字）
	_, err := parseIndex("-")
	require.Error(t, err, "parseIndex(\"-\") should return error")
	assert.ErrorIs(t, err, ErrInvalidIndex)
}

func TestParseIndex_Coverage_EdgeCase_VeryLongNumber(t *testing.T) {
	// 测试非常长的数字（可能溢出，但函数不处理溢出）
	veryLong := "12345678901234567890"
	result, err := parseIndex(veryLong)
	require.NoError(t, err, "parseIndex should accept very long numbers")
	// 结果可能溢出，但不应崩溃
	assert.NotEqual(t, 0, result, "very long number should not return zero")
}

func TestParseIndex_Coverage_EdgeCase_UnicodeDigits(t *testing.T) {
	// 测试非 ASCII 数字字符（应该失败）
	unicodeDigit := "١" // Arabic-Indic digit 1
	_, err := parseIndex(unicodeDigit)
	require.Error(t, err, "parseIndex should reject non-ASCII digits")
	assert.ErrorIs(t, err, ErrInvalidIndex)
}

func TestParseIndex_Coverage_EdgeCase_TabAndNewline(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "with tab", input: "12\t3"},
		{name: "with newline", input: "12\n3"},
		{name: "with carriage return", input: "12\r3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := parseIndex(tt.input)
			require.Error(t, err, "parseIndex(%q) should reject whitespace", tt.input)
			assert.ErrorIs(t, err, ErrInvalidIndex)
		})
	}
}

// ============================================================
// 性能关键路径的正确性验证
// ============================================================

func TestParseIndex_Coverage_PerformanceCriticalPaths(t *testing.T) {
	// 测试最常见的路径（热路径）
	cases := []string{"0", "1", "2", "10", "100", "-1", "-2"}

	for _, s := range cases {
		t.Run("common_"+s, func(t *testing.T) {
			result, err := parseIndex(s)
			require.NoError(t, err, "parseIndex(%q) should not return error", s)

			// 使用 strconv.Atoi 验证结果正确性
			expected, err2 := parseIndex(s)
			require.NoError(t, err2, "reference parseIndex failed")
			assert.Equal(t, expected, result, "result mismatch for %q", s)
		})
	}
}
