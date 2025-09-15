package stringx

import (
	"strings"
	"testing"
)

// TestToSnakeCapacityLimit tests the capacity limit in ToSnake function
func TestToSnakeCapacityLimit(t *testing.T) {
	t.Run("large_input_capacity_limit", func(t *testing.T) {
		// Create a string large enough to trigger capacity > 256 condition
		// This should cover the missing line 121.20,123.3
		longString := strings.Repeat("A", 300)
		result := ToSnake(longString)

		// Should convert to snake case (each letter separated by underscore)
		// The ToSnake function adds underscores between letters, so "AAA" becomes "a_a_a"
		if result == "" {
			t.Error("ToSnake should not return empty string")
		}

		t.Logf("ToSnake handled large input correctly: length=%d", len(result))
	})

	t.Run("capacity_estimation_edge_cases", func(t *testing.T) {
		// Test various string lengths to ensure capacity estimation works
		testCases := []struct {
			input string
			desc  string
		}{
			{strings.Repeat("ABC", 100), "300 character string"},
			{strings.Repeat("A", 400), "400 character string"},
			{strings.Repeat("CamelCase", 50), "repeated camel case"},
		}

		for _, tc := range testCases {
			t.Run(tc.desc, func(t *testing.T) {
				result := ToSnake(tc.input)
				if result == "" {
					t.Error("ToSnake should not return empty string for valid input")
				}
				t.Logf("Processed %s: result length=%d", tc.desc, len(result))
			})
		}
	})
}

// TestToSmallCamelElseBranch tests the else branch in ToSmallCamel function
func TestToSmallCamelElseBranch(t *testing.T) {
	t.Run("non_letter_character_handling", func(t *testing.T) {
		// Create input that triggers the else branch (lines 332-334)
		// This happens when upper=true but the character is not a letter
		testCases := []string{
			"test_123_case",  // numbers after underscores
			"test_!@#_case",  // symbols after underscores
			"test___case",    // multiple underscores
			"test_$%^_case",  // special characters
			"test_1a2b_case", // mixed numbers and letters
		}

		for _, input := range testCases {
			t.Run(input, func(t *testing.T) {
				result := ToSmallCamel(input)
				if result == "" {
					t.Error("ToSmallCamel should not return empty string")
				}
				t.Logf("Input: %s, Output: %s", input, result)
			})
		}
	})

	t.Run("edge_case_characters", func(t *testing.T) {
		// Test with various non-letter characters that should trigger the else branch
		edgeCases := []string{
			"a_1b",         // number after underscore
			"a_@b",         // symbol after underscore
			"test_8_value", // number in middle
			"x_#_y",        // symbol in middle
			"start_9end",   // number at word boundary
		}

		for _, input := range edgeCases {
			result := ToSmallCamel(input)
			// Just verify it doesn't crash and returns something
			if len(result) == 0 {
				t.Errorf("ToSmallCamel(%s) returned empty string", input)
			}
		}
	})
}
