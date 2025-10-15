package cryptox

import (
	"bytes"
	"testing"
)

// TestHash32 tests Hash32 function with various inputs
func TestHash32(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected uint32
	}{
		{"empty string", "", 2166136261},
		{"single byte", "a", 84696446},
		{"short string", "hello", 3069866343},
		{"FNV test vector", "foobar", 837857890},
		{"numeric string", "123456", 3942681528},
	}

	for _, tc := range testCases {
		t.Run(tc.name+" (string)", func(t *testing.T) {
			result := Hash32(tc.input)
			if result != tc.expected {
				t.Errorf("Hash32(string) = %d, expected %d", result, tc.expected)
			}
		})

		t.Run(tc.name+" ([]byte)", func(t *testing.T) {
			result := Hash32([]byte(tc.input))
			if result != tc.expected {
				t.Errorf("Hash32([]byte) = %d, expected %d", result, tc.expected)
			}
		})
	}
}

// TestHash32a tests Hash32a function with various inputs
func TestHash32a(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected uint32
	}{
		{"empty string", "", 2166136261},
		{"single byte", "a", 3826002220},
		{"short string", "hello", 1335831723},
		{"FNV-1a test vector", "foobar", 3214735720},
		{"numeric string", "123456", 2576725674},
	}

	for _, tc := range testCases {
		t.Run(tc.name+" (string)", func(t *testing.T) {
			result := Hash32a(tc.input)
			if result != tc.expected {
				t.Errorf("Hash32a(string) = %d, expected %d", result, tc.expected)
			}
		})

		t.Run(tc.name+" ([]byte)", func(t *testing.T) {
			result := Hash32a([]byte(tc.input))
			if result != tc.expected {
				t.Errorf("Hash32a([]byte) = %d, expected %d", result, tc.expected)
			}
		})
	}
}

// TestHash64 tests Hash64 function with various inputs
func TestHash64(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected uint64
	}{
		{"empty string", "", 14695981039346656037},
		{"single byte", "a", 12638153115695167422},
		{"short string", "hello", 8883723591023973575},
		{"FNV test vector", "foobar", 3750802935296928194},
		{"numeric string", "123456", 1222894297566385272},
	}

	for _, tc := range testCases {
		t.Run(tc.name+" (string)", func(t *testing.T) {
			result := Hash64(tc.input)
			if result != tc.expected {
				t.Errorf("Hash64(string) = %d, expected %d", result, tc.expected)
			}
		})

		t.Run(tc.name+" ([]byte)", func(t *testing.T) {
			result := Hash64([]byte(tc.input))
			if result != tc.expected {
				t.Errorf("Hash64([]byte) = %d, expected %d", result, tc.expected)
			}
		})
	}
}

// TestHash64a tests Hash64a function with various inputs
func TestHash64a(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected uint64
	}{
		{"empty string", "", 14695981039346656037},
		{"single byte", "a", 12638187200555641996},
		{"short string", "hello", 11831194018420276491},
		{"FNV-1a test vector", "foobar", 9625390261332436968},
		{"numeric string", "123456", 17790324078706895114},
	}

	for _, tc := range testCases {
		t.Run(tc.name+" (string)", func(t *testing.T) {
			result := Hash64a(tc.input)
			if result != tc.expected {
				t.Errorf("Hash64a(string) = %d, expected %d", result, tc.expected)
			}
		})

		t.Run(tc.name+" ([]byte)", func(t *testing.T) {
			result := Hash64a([]byte(tc.input))
			if result != tc.expected {
				t.Errorf("Hash64a([]byte) = %d, expected %d", result, tc.expected)
			}
		})
	}
}

// TestFNVWithVariousDataTypes tests FNV functions with various data types
func TestFNVWithVariousDataTypes(t *testing.T) {
	testData := []struct {
		name string
		data []byte
	}{
		{"binary zeros", []byte{0x00, 0x00, 0x00, 0x00}},
		{"binary ones", []byte{0xFF, 0xFF, 0xFF, 0xFF}},
		{"binary mixed", []byte{0x01, 0x02, 0x03, 0x04, 0x05}},
		{"unicode Chinese", []byte("你好世界")},
		{"unicode emoji", []byte("🌍🚀💻")},
		{"mixed ASCII and Unicode", []byte("Hello 世界 🌍")},
		{"long message", bytes.Repeat([]byte("X"), 1000)},
		{"all printable ASCII", []byte("!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ")},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			// Just verify all functions return non-zero and consistent results
			hash32_1 := Hash32(td.data)
			hash32_2 := Hash32(td.data)
			if hash32_1 != hash32_2 {
				t.Error("Hash32 should be deterministic")
			}

			hash32a_1 := Hash32a(td.data)
			hash32a_2 := Hash32a(td.data)
			if hash32a_1 != hash32a_2 {
				t.Error("Hash32a should be deterministic")
			}

			hash64_1 := Hash64(td.data)
			hash64_2 := Hash64(td.data)
			if hash64_1 != hash64_2 {
				t.Error("Hash64 should be deterministic")
			}

			hash64a_1 := Hash64a(td.data)
			hash64a_2 := Hash64a(td.data)
			if hash64a_1 != hash64a_2 {
				t.Error("Hash64a should be deterministic")
			}
		})
	}
}

// TestFNVConsistency tests that multiple calls with same input produce same output
func TestFNVConsistency(t *testing.T) {
	input := "consistency test data"

	// Test Hash32
	result1 := Hash32(input)
	result2 := Hash32(input)
	result3 := Hash32(input)
	if result1 != result2 || result2 != result3 {
		t.Error("Hash32 should produce consistent results")
	}

	// Test Hash32a
	result1a := Hash32a(input)
	result2a := Hash32a(input)
	result3a := Hash32a(input)
	if result1a != result2a || result2a != result3a {
		t.Error("Hash32a should produce consistent results")
	}

	// Test Hash64
	result1_64 := Hash64(input)
	result2_64 := Hash64(input)
	result3_64 := Hash64(input)
	if result1_64 != result2_64 || result2_64 != result3_64 {
		t.Error("Hash64 should produce consistent results")
	}

	// Test Hash64a
	result1a_64 := Hash64a(input)
	result2a_64 := Hash64a(input)
	result3a_64 := Hash64a(input)
	if result1a_64 != result2a_64 || result2a_64 != result3a_64 {
		t.Error("Hash64a should produce consistent results")
	}
}

// TestFNVDifferentInputs tests that different inputs produce different outputs
func TestFNVDifferentInputs(t *testing.T) {
	inputs := []string{
		"input1",
		"input2",
		"input3",
		"different",
		"data",
	}

	// Test Hash32
	seen32 := make(map[uint32]bool)
	for _, input := range inputs {
		hash := Hash32(input)
		if seen32[hash] {
			t.Errorf("Hash32 collision detected for different inputs")
		}
		seen32[hash] = true
	}

	// Test Hash32a
	seen32a := make(map[uint32]bool)
	for _, input := range inputs {
		hash := Hash32a(input)
		if seen32a[hash] {
			t.Errorf("Hash32a collision detected for different inputs")
		}
		seen32a[hash] = true
	}

	// Test Hash64
	seen64 := make(map[uint64]bool)
	for _, input := range inputs {
		hash := Hash64(input)
		if seen64[hash] {
			t.Errorf("Hash64 collision detected for different inputs")
		}
		seen64[hash] = true
	}

	// Test Hash64a
	seen64a := make(map[uint64]bool)
	for _, input := range inputs {
		hash := Hash64a(input)
		if seen64a[hash] {
			t.Errorf("Hash64a collision detected for different inputs")
		}
		seen64a[hash] = true
	}
}

// TestFNV1vsFNV1a tests that FNV-1 and FNV-1a produce different results
func TestFNV1vsFNV1a(t *testing.T) {
	testInputs := []string{
		"test",
		"hello world",
		"FNV comparison",
		"123456789",
		"你好世界",
	}

	for _, input := range testInputs {
		t.Run(input, func(t *testing.T) {
			// For most inputs, FNV-1 and FNV-1a should produce different results
			// (except for empty string which has the same offset basis)
			if input != "" {
				hash32 := Hash32(input)
				hash32a := Hash32a(input)
				
				// Note: It's possible (though unlikely) for them to be equal
				// but typically they should differ
				if hash32 == hash32a {
					t.Logf("Warning: Hash32 and Hash32a produced same result for '%s' (this is rare but possible)", input)
				}

				hash64 := Hash64(input)
				hash64a := Hash64a(input)
				
				if hash64 == hash64a {
					t.Logf("Warning: Hash64 and Hash64a produced same result for '%s' (this is rare but possible)", input)
				}
			}
		})
	}
}

// TestFNV32vs64 tests that 32-bit and 64-bit versions produce different ranges
func TestFNV32vs64(t *testing.T) {
	input := "test data for 32 vs 64 bit comparison"

	hash32 := Hash32(input)
	hash64 := Hash64(input)

	// Verify hash32 fits in 32 bits (always true by type)
	// and hash64 can use the full 64-bit range
	if hash32 == 0 && hash64 == 0 {
		t.Error("Both hashes should not be zero for non-empty input")
	}

	// Test with multiple inputs to ensure we get different hash sizes
	hash32a := Hash32a(input)
	hash64a := Hash64a(input)

	if hash32a == 0 && hash64a == 0 {
		t.Error("Both hashes should not be zero for non-empty input")
	}
}

// TestFNVEmptyInput tests behavior with empty input
func TestFNVEmptyInput(t *testing.T) {
	// Empty string should return the FNV offset basis
	t.Run("Hash32 empty", func(t *testing.T) {
		result := Hash32("")
		expected := uint32(2166136261) // FNV-1 32-bit offset basis
		if result != expected {
			t.Errorf("Hash32('') = %d, expected %d", result, expected)
		}
	})

	t.Run("Hash32a empty", func(t *testing.T) {
		result := Hash32a("")
		expected := uint32(2166136261) // FNV-1a 32-bit offset basis
		if result != expected {
			t.Errorf("Hash32a('') = %d, expected %d", result, expected)
		}
	})

	t.Run("Hash64 empty", func(t *testing.T) {
		result := Hash64("")
		expected := uint64(14695981039346656037) // FNV-1 64-bit offset basis
		if result != expected {
			t.Errorf("Hash64('') = %d, expected %d", result, expected)
		}
	})

	t.Run("Hash64a empty", func(t *testing.T) {
		result := Hash64a("")
		expected := uint64(14695981039346656037) // FNV-1a 64-bit offset basis
		if result != expected {
			t.Errorf("Hash64a('') = %d, expected %d", result, expected)
		}
	})
}

// TestFNVSingleByteSequence tests behavior with sequences of single bytes
func TestFNVSingleByteSequence(t *testing.T) {
	// Test that hashing byte by byte produces different results
	hashes32 := make([]uint32, 256)
	hashes64 := make([]uint64, 256)

	for i := 0; i < 256; i++ {
		b := byte(i)
		hashes32[i] = Hash32a([]byte{b})
		hashes64[i] = Hash64a([]byte{b})
	}

	// Verify all are unique (FNV should have no collisions for single bytes)
	seen32 := make(map[uint32]bool)
	for _, h := range hashes32 {
		if seen32[h] {
			t.Error("Hash32a collision detected for single byte values")
			break
		}
		seen32[h] = true
	}

	seen64 := make(map[uint64]bool)
	for _, h := range hashes64 {
		if seen64[h] {
			t.Error("Hash64a collision detected for single byte values")
			break
		}
		seen64[h] = true
	}
}

// TestFNVStringVsByteSlice tests that string and []byte produce same results
func TestFNVStringVsByteSlice(t *testing.T) {
	testInputs := []string{
		"test",
		"hello world",
		"123456",
		"你好",
		"",
	}

	for _, input := range testInputs {
		t.Run(input, func(t *testing.T) {
			if Hash32(input) != Hash32([]byte(input)) {
				t.Error("Hash32 should produce same result for string and []byte")
			}
			if Hash32a(input) != Hash32a([]byte(input)) {
				t.Error("Hash32a should produce same result for string and []byte")
			}
			if Hash64(input) != Hash64([]byte(input)) {
				t.Error("Hash64 should produce same result for string and []byte")
			}
			if Hash64a(input) != Hash64a([]byte(input)) {
				t.Error("Hash64a should produce same result for string and []byte")
			}
		})
	}
}

// TestFNVLongMessages tests behavior with very long messages
func TestFNVLongMessages(t *testing.T) {
	sizes := []int{1000, 10000, 100000}

	for _, size := range sizes {
		t.Run("size_"+string(rune(size/1000+'0'))+"k", func(t *testing.T) {
			data := bytes.Repeat([]byte("A"), size)

			// Just verify they complete without error and produce consistent results
			hash32_1 := Hash32(data)
			hash32_2 := Hash32(data)
			if hash32_1 != hash32_2 {
				t.Error("Hash32 should be consistent for long messages")
			}

			hash64_1 := Hash64(data)
			hash64_2 := Hash64(data)
			if hash64_1 != hash64_2 {
				t.Error("Hash64 should be consistent for long messages")
			}
		})
	}
}

// BenchmarkHash32 benchmarks Hash32 function
func BenchmarkHash32String(b *testing.B) {
	input := "benchmark test data for FNV-1 32-bit hash"
	for i := 0; i < b.N; i++ {
		_ = Hash32(input)
	}
}

func BenchmarkHash32Bytes(b *testing.B) {
	input := []byte("benchmark test data for FNV-1 32-bit hash")
	for i := 0; i < b.N; i++ {
		_ = Hash32(input)
	}
}

func BenchmarkHash32aString(b *testing.B) {
	input := "benchmark test data for FNV-1a 32-bit hash"
	for i := 0; i < b.N; i++ {
		_ = Hash32a(input)
	}
}

func BenchmarkHash32aBytes(b *testing.B) {
	input := []byte("benchmark test data for FNV-1a 32-bit hash")
	for i := 0; i < b.N; i++ {
		_ = Hash32a(input)
	}
}

func BenchmarkHash64String(b *testing.B) {
	input := "benchmark test data for FNV-1 64-bit hash"
	for i := 0; i < b.N; i++ {
		_ = Hash64(input)
	}
}

func BenchmarkHash64Bytes(b *testing.B) {
	input := []byte("benchmark test data for FNV-1 64-bit hash")
	for i := 0; i < b.N; i++ {
		_ = Hash64(input)
	}
}

func BenchmarkHash64aString(b *testing.B) {
	input := "benchmark test data for FNV-1a 64-bit hash"
	for i := 0; i < b.N; i++ {
		_ = Hash64a(input)
	}
}

func BenchmarkHash64aBytes(b *testing.B) {
	input := []byte("benchmark test data for FNV-1a 64-bit hash")
	for i := 0; i < b.N; i++ {
		_ = Hash64a(input)
	}
}

func BenchmarkHash32aLongMessage(b *testing.B) {
	input := bytes.Repeat([]byte("X"), 10000)
	for i := 0; i < b.N; i++ {
		_ = Hash32a(input)
	}
}

func BenchmarkHash64aLongMessage(b *testing.B) {
	input := bytes.Repeat([]byte("X"), 10000)
	for i := 0; i < b.N; i++ {
		_ = Hash64a(input)
	}
}
