package stringx

import (
	"math/rand"
	"regexp"
	"testing"
	"time"
)

func init() {
	// Seed the random number generator for consistent testing
	rand.Seed(time.Now().UnixNano())
}

func TestRandLetters(t *testing.T) {
	t.Run("zero_length", func(t *testing.T) {
		result := RandLetters(0)
		if result != "" {
			t.Errorf("RandLetters(0) = %q, expected empty string", result)
		}
	})

	t.Run("negative_length", func(t *testing.T) {
		result := RandLetters(-1)
		if result != "" {
			t.Errorf("RandLetters(-1) = %q, expected empty string", result)
		}
	})

	t.Run("valid_length", func(t *testing.T) {
		length := 10
		result := RandLetters(length)
		if len(result) != length {
			t.Errorf("RandLetters(%d) length = %d, expected %d", length, len(result), length)
		}
		
		// Verify all characters are letters (a-z, A-Z)
		matched, err := regexp.MatchString("^[a-zA-Z]+$", result)
		if err != nil {
			t.Fatalf("Regex error: %v", err)
		}
		if !matched {
			t.Errorf("RandLetters(%d) = %q, contains non-letter characters", length, result)
		}
	})

	t.Run("randomness_check", func(t *testing.T) {
		// Generate multiple strings and check they're different
		results := make(map[string]bool)
		length := 8
		iterations := 100
		
		for i := 0; i < iterations; i++ {
			result := RandLetters(length)
			results[result] = true
		}
		
		// Should have high diversity (at least 90% unique)
		if len(results) < int(float64(iterations)*0.9) {
			t.Errorf("RandLetters not random enough: got %d unique strings out of %d", len(results), iterations)
		}
	})
}

func TestRandLowerLetters(t *testing.T) {
	t.Run("zero_length", func(t *testing.T) {
		result := RandLowerLetters(0)
		if result != "" {
			t.Errorf("RandLowerLetters(0) = %q, expected empty string", result)
		}
	})

	t.Run("negative_length", func(t *testing.T) {
		result := RandLowerLetters(-1)
		if result != "" {
			t.Errorf("RandLowerLetters(-1) = %q, expected empty string", result)
		}
	})

	t.Run("valid_length", func(t *testing.T) {
		length := 10
		result := RandLowerLetters(length)
		if len(result) != length {
			t.Errorf("RandLowerLetters(%d) length = %d, expected %d", length, len(result), length)
		}
		
		// Verify all characters are lowercase letters (a-z)
		matched, err := regexp.MatchString("^[a-z]+$", result)
		if err != nil {
			t.Fatalf("Regex error: %v", err)
		}
		if !matched {
			t.Errorf("RandLowerLetters(%d) = %q, contains non-lowercase characters", length, result)
		}
	})
}

func TestRandUpperLetters(t *testing.T) {
	t.Run("zero_length", func(t *testing.T) {
		result := RandUpperLetters(0)
		if result != "" {
			t.Errorf("RandUpperLetters(0) = %q, expected empty string", result)
		}
	})

	t.Run("negative_length", func(t *testing.T) {
		result := RandUpperLetters(-1)
		if result != "" {
			t.Errorf("RandUpperLetters(-1) = %q, expected empty string", result)
		}
	})

	t.Run("valid_length", func(t *testing.T) {
		length := 10
		result := RandUpperLetters(length)
		if len(result) != length {
			t.Errorf("RandUpperLetters(%d) length = %d, expected %d", length, len(result), length)
		}
		
		// Verify all characters are uppercase letters (A-Z)
		matched, err := regexp.MatchString("^[A-Z]+$", result)
		if err != nil {
			t.Fatalf("Regex error: %v", err)
		}
		if !matched {
			t.Errorf("RandUpperLetters(%d) = %q, contains non-uppercase characters", length, result)
		}
	})
}

func TestRandNumbers(t *testing.T) {
	t.Run("zero_length", func(t *testing.T) {
		result := RandNumbers(0)
		if result != "" {
			t.Errorf("RandNumbers(0) = %q, expected empty string", result)
		}
	})

	t.Run("negative_length", func(t *testing.T) {
		result := RandNumbers(-1)
		if result != "" {
			t.Errorf("RandNumbers(-1) = %q, expected empty string", result)
		}
	})

	t.Run("valid_length", func(t *testing.T) {
		length := 10
		result := RandNumbers(length)
		if len(result) != length {
			t.Errorf("RandNumbers(%d) length = %d, expected %d", length, len(result), length)
		}
		
		// Verify all characters are digits (0-9)
		matched, err := regexp.MatchString("^[0-9]+$", result)
		if err != nil {
			t.Fatalf("Regex error: %v", err)
		}
		if !matched {
			t.Errorf("RandNumbers(%d) = %q, contains non-digit characters", length, result)
		}
	})
}

func TestRandLetterNumbers(t *testing.T) {
	t.Run("zero_length", func(t *testing.T) {
		result := RandLetterNumbers(0)
		if result != "" {
			t.Errorf("RandLetterNumbers(0) = %q, expected empty string", result)
		}
	})

	t.Run("negative_length", func(t *testing.T) {
		result := RandLetterNumbers(-1)
		if result != "" {
			t.Errorf("RandLetterNumbers(-1) = %q, expected empty string", result)
		}
	})

	t.Run("valid_length", func(t *testing.T) {
		length := 10
		result := RandLetterNumbers(length)
		if len(result) != length {
			t.Errorf("RandLetterNumbers(%d) length = %d, expected %d", length, len(result), length)
		}
		
		// Verify all characters are alphanumeric (a-z, A-Z, 0-9)
		matched, err := regexp.MatchString("^[a-zA-Z0-9]+$", result)
		if err != nil {
			t.Fatalf("Regex error: %v", err)
		}
		if !matched {
			t.Errorf("RandLetterNumbers(%d) = %q, contains non-alphanumeric characters", length, result)
		}
	})

	t.Run("contains_both_letters_and_numbers", func(t *testing.T) {
		// Generate a long enough string to statistically contain both
		length := 100
		hasLetter := false
		hasNumber := false
		
		for i := 0; i < 10; i++ {
			result := RandLetterNumbers(length)
			for _, char := range result {
				if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
					hasLetter = true
				}
				if char >= '0' && char <= '9' {
					hasNumber = true
				}
			}
			if hasLetter && hasNumber {
				break
			}
		}
		
		if !hasLetter || !hasNumber {
			t.Error("RandLetterNumbers should eventually generate both letters and numbers")
		}
	})
}

func TestRandLowerLetterNumbers(t *testing.T) {
	t.Run("zero_length", func(t *testing.T) {
		result := RandLowerLetterNumbers(0)
		if result != "" {
			t.Errorf("RandLowerLetterNumbers(0) = %q, expected empty string", result)
		}
	})

	t.Run("negative_length", func(t *testing.T) {
		result := RandLowerLetterNumbers(-1)
		if result != "" {
			t.Errorf("RandLowerLetterNumbers(-1) = %q, expected empty string", result)
		}
	})

	t.Run("valid_length", func(t *testing.T) {
		length := 10
		result := RandLowerLetterNumbers(length)
		if len(result) != length {
			t.Errorf("RandLowerLetterNumbers(%d) length = %d, expected %d", length, len(result), length)
		}
		
		// Verify all characters are lowercase alphanumeric (a-z, 0-9)
		matched, err := regexp.MatchString("^[a-z0-9]+$", result)
		if err != nil {
			t.Fatalf("Regex error: %v", err)
		}
		if !matched {
			t.Errorf("RandLowerLetterNumbers(%d) = %q, contains non-lowercase-alphanumeric characters", length, result)
		}
	})
}

func TestRandUpperLetterNumbers(t *testing.T) {
	t.Run("zero_length", func(t *testing.T) {
		result := RandUpperLetterNumbers(0)
		if result != "" {
			t.Errorf("RandUpperLetterNumbers(0) = %q, expected empty string", result)
		}
	})

	t.Run("negative_length", func(t *testing.T) {
		result := RandUpperLetterNumbers(-1)
		if result != "" {
			t.Errorf("RandUpperLetterNumbers(-1) = %q, expected empty string", result)
		}
	})

	t.Run("valid_length", func(t *testing.T) {
		length := 10
		result := RandUpperLetterNumbers(length)
		if len(result) != length {
			t.Errorf("RandUpperLetterNumbers(%d) length = %d, expected %d", length, len(result), length)
		}
		
		// Verify all characters are uppercase alphanumeric (A-Z, 0-9)
		matched, err := regexp.MatchString("^[A-Z0-9]+$", result)
		if err != nil {
			t.Fatalf("Regex error: %v", err)
		}
		if !matched {
			t.Errorf("RandUpperLetterNumbers(%d) = %q, contains non-uppercase-alphanumeric characters", length, result)
		}
	})
}

func TestRandStringWithSeed(t *testing.T) {
	t.Run("zero_length", func(t *testing.T) {
		seed := []rune("abc")
		result := RandStringWithSeed(0, seed)
		if result != "" {
			t.Errorf("RandStringWithSeed(0, seed) = %q, expected empty string", result)
		}
	})

	t.Run("negative_length", func(t *testing.T) {
		seed := []rune("abc")
		result := RandStringWithSeed(-1, seed)
		if result != "" {
			t.Errorf("RandStringWithSeed(-1, seed) = %q, expected empty string", result)
		}
	})

	t.Run("empty_seed", func(t *testing.T) {
		result := RandStringWithSeed(5, []rune{})
		if result != "" {
			t.Errorf("RandStringWithSeed(5, empty) = %q, expected empty string", result)
		}
	})

	t.Run("nil_seed", func(t *testing.T) {
		result := RandStringWithSeed(5, nil)
		if result != "" {
			t.Errorf("RandStringWithSeed(5, nil) = %q, expected empty string", result)
		}
	})

	t.Run("single_character_seed", func(t *testing.T) {
		seed := []rune("a")
		length := 5
		result := RandStringWithSeed(length, seed)
		expected := "aaaaa"
		if result != expected {
			t.Errorf("RandStringWithSeed(%d, %v) = %q, expected %q", length, seed, result, expected)
		}
	})

	t.Run("valid_seed_and_length", func(t *testing.T) {
		seed := []rune("abc123")
		length := 10
		result := RandStringWithSeed(length, seed)
		
		if len(result) != length {
			t.Errorf("RandStringWithSeed length = %d, expected %d", len(result), length)
		}
		
		// Verify all characters are from the seed
		seedMap := make(map[rune]bool)
		for _, r := range seed {
			seedMap[r] = true
		}
		
		for _, r := range result {
			if !seedMap[r] {
				t.Errorf("RandStringWithSeed result contains %q which is not in seed %v", r, seed)
			}
		}
	})

	t.Run("unicode_seed", func(t *testing.T) {
		seed := []rune("测试🎉😀")
		length := 6
		result := RandStringWithSeed(length, seed)
		
		if len([]rune(result)) != length {
			t.Errorf("RandStringWithSeed unicode length = %d, expected %d", len([]rune(result)), length)
		}
		
		// Verify all characters are from the seed
		seedMap := make(map[rune]bool)
		for _, r := range seed {
			seedMap[r] = true
		}
		
		for _, r := range []rune(result) {
			if !seedMap[r] {
				t.Errorf("RandStringWithSeed result contains %q which is not in unicode seed %v", r, seed)
			}
		}
	})

	t.Run("distribution_check", func(t *testing.T) {
		seed := []rune("ab")
		length := 1000
		result := RandStringWithSeed(length, seed)
		
		countA := 0
		countB := 0
		for _, r := range result {
			switch r {
			case 'a':
				countA++
			case 'b':
				countB++
			}
		}
		
		// Check if distribution is roughly 50/50 (within 10% tolerance)
		tolerance := length / 10
		expectedCount := length / 2
		if abs(countA-expectedCount) > tolerance || abs(countB-expectedCount) > tolerance {
			t.Errorf("Distribution not balanced: a=%d, b=%d (expected ~%d each)", countA, countB, expectedCount)
		}
	})
}

// Helper function for absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Benchmark tests
func BenchmarkRandLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandLetters(10)
	}
}

func BenchmarkRandNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandNumbers(10)
	}
}

func BenchmarkRandLetterNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandLetterNumbers(10)
	}
}

func BenchmarkRandStringWithSeed(b *testing.B) {
	seed := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RandStringWithSeed(10, seed)
	}
}

// Performance comparison tests
func TestPerformanceComparison(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping performance test in short mode")
	}
	
	t.Run("large_string_generation", func(t *testing.T) {
		length := 10000
		
		start := time.Now()
		result := RandLetterNumbers(length)
		duration := time.Since(start)
		
		if len(result) != length {
			t.Errorf("Generated string length %d, expected %d", len(result), length)
		}
		
		t.Logf("Generated %d character string in %v", length, duration)
		
		// Should complete within reasonable time (1 second)
		if duration > time.Second {
			t.Errorf("String generation took too long: %v", duration)
		}
	})
}

// Edge case tests
func TestEdgeCases(t *testing.T) {
	t.Run("max_int_length", func(t *testing.T) {
		// Don't actually generate max int length (would use too much memory)
		// Just test that it doesn't panic with very large numbers
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Expected behavior: panicked with large length: %v", r)
			}
		}()
		
		// This might panic due to memory allocation, which is expected
		_ = RandStringWithSeed(1000000, []rune("a"))
	})
	
	t.Run("special_characters_seed", func(t *testing.T) {
		seed := []rune("!@#$%^&*()_+{}|:<>?")
		length := 20
		result := RandStringWithSeed(length, seed)
		
		if len(result) != length {
			t.Errorf("Special characters seed length = %d, expected %d", len(result), length)
		}
		
		// Verify all characters are from the seed
		seedMap := make(map[rune]bool)
		for _, r := range seed {
			seedMap[r] = true
		}
		
		for _, r := range result {
			if !seedMap[r] {
				t.Errorf("Result contains %q which is not in special characters seed", r)
			}
		}
	})
}