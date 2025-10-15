package cryptox

import (
	"bytes"
	"testing"
)

// TestCRC32 tests CRC32 function with various inputs
func TestCRC32(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected uint32
	}{
		{"empty string", "", 0},
		{"single byte", "a", 0xe8b7be43},
		{"short string", "hello", 0x3610a686},
		{"hello world", "hello world", 0x0d4a1185},
		{"numeric string", "123456789", 0xcbf43926},
		{"The quick brown fox", "The quick brown fox jumps over the lazy dog", 0x414fa339},
	}

	for _, tc := range testCases {
		t.Run(tc.name+" (string)", func(t *testing.T) {
			result := CRC32(tc.input)
			if result != tc.expected {
				t.Errorf("CRC32(string) = 0x%08x, expected 0x%08x", result, tc.expected)
			}
		})

		t.Run(tc.name+" ([]byte)", func(t *testing.T) {
			result := CRC32([]byte(tc.input))
			if result != tc.expected {
				t.Errorf("CRC32([]byte) = 0x%08x, expected 0x%08x", result, tc.expected)
			}
		})
	}
}

// TestCRC64 tests CRC64 function with various inputs
func TestCRC64(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected uint64
	}{
		{"empty string", "", 0},
		{"single byte", "a", 0x330284772e652b05},
		{"short string", "hello", 0x9b1edae5dbb937b1},
		{"hello world", "hello world", 0x53037ecdef2352da},
		{"numeric string", "123456789", 0x995dc9bbdf1939fa},
		{"The quick brown fox", "The quick brown fox jumps over the lazy dog", 0x5b5eb8c2e54aa1c4},
	}

	for _, tc := range testCases {
		t.Run(tc.name+" (string)", func(t *testing.T) {
			result := CRC64(tc.input)
			if result != tc.expected {
				t.Errorf("CRC64(string) = 0x%016x, expected 0x%016x", result, tc.expected)
			}
		})

		t.Run(tc.name+" ([]byte)", func(t *testing.T) {
			result := CRC64([]byte(tc.input))
			if result != tc.expected {
				t.Errorf("CRC64([]byte) = 0x%016x, expected 0x%016x", result, tc.expected)
			}
		})
	}
}

// TestCRCWithVariousDataTypes tests CRC functions with various data types
func TestCRCWithVariousDataTypes(t *testing.T) {
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
		{"newlines", []byte("line1\nline2\nline3\n")},
		{"tabs and spaces", []byte("\t  \t  \t  ")},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			// Just verify all functions return consistent results
			crc32_1 := CRC32(td.data)
			crc32_2 := CRC32(td.data)
			if crc32_1 != crc32_2 {
				t.Error("CRC32 should be deterministic")
			}

			crc64_1 := CRC64(td.data)
			crc64_2 := CRC64(td.data)
			if crc64_1 != crc64_2 {
				t.Error("CRC64 should be deterministic")
			}
		})
	}
}

// TestCRCConsistency tests that multiple calls with same input produce same output
func TestCRCConsistency(t *testing.T) {
	input := "consistency test data"

	// Test CRC32
	result1 := CRC32(input)
	result2 := CRC32(input)
	result3 := CRC32(input)
	if result1 != result2 || result2 != result3 {
		t.Error("CRC32 should produce consistent results")
	}

	// Test CRC64
	result1_64 := CRC64(input)
	result2_64 := CRC64(input)
	result3_64 := CRC64(input)
	if result1_64 != result2_64 || result2_64 != result3_64 {
		t.Error("CRC64 should produce consistent results")
	}
}

// TestCRCDifferentInputs tests that different inputs produce different outputs
func TestCRCDifferentInputs(t *testing.T) {
	inputs := []string{
		"input1",
		"input2",
		"input3",
		"different",
		"data",
		"test1",
		"test2",
		"test3",
	}

	// Test CRC32
	seen32 := make(map[uint32]bool)
	for _, input := range inputs {
		crc := CRC32(input)
		if seen32[crc] {
			t.Errorf("CRC32 collision detected for different inputs")
		}
		seen32[crc] = true
	}

	// Test CRC64
	seen64 := make(map[uint64]bool)
	for _, input := range inputs {
		crc := CRC64(input)
		if seen64[crc] {
			t.Errorf("CRC64 collision detected for different inputs")
		}
		seen64[crc] = true
	}
}

// TestCRC32vs64 tests that 32-bit and 64-bit versions produce different ranges
func TestCRC32vs64(t *testing.T) {
	input := "test data for 32 vs 64 bit comparison"

	crc32 := CRC32(input)
	crc64 := CRC64(input)

	// Verify crc32 fits in 32 bits (always true by type)
	// and crc64 can use the full 64-bit range
	if crc32 == 0 && crc64 == 0 {
		t.Error("Both CRCs should not be zero for non-empty input")
	}

	// Verify they produce different values (as they use different polynomials)
	if uint64(crc32) == crc64 {
		t.Log("CRC32 and CRC64 happened to produce same value (rare but possible)")
	}
}

// TestCRCEmptyInput tests behavior with empty input
func TestCRCEmptyInput(t *testing.T) {
	t.Run("CRC32 empty", func(t *testing.T) {
		result := CRC32("")
		expected := uint32(0)
		if result != expected {
			t.Errorf("CRC32('') = 0x%08x, expected 0x%08x", result, expected)
		}
	})

	t.Run("CRC32 empty bytes", func(t *testing.T) {
		result := CRC32([]byte{})
		expected := uint32(0)
		if result != expected {
			t.Errorf("CRC32([]byte{}) = 0x%08x, expected 0x%08x", result, expected)
		}
	})

	t.Run("CRC64 empty", func(t *testing.T) {
		result := CRC64("")
		expected := uint64(0)
		if result != expected {
			t.Errorf("CRC64('') = 0x%016x, expected 0x%016x", result, expected)
		}
	})

	t.Run("CRC64 empty bytes", func(t *testing.T) {
		result := CRC64([]byte{})
		expected := uint64(0)
		if result != expected {
			t.Errorf("CRC64([]byte{}) = 0x%016x, expected 0x%016x", result, expected)
		}
	})
}

// TestCRCSingleByteSequence tests behavior with sequences of single bytes
func TestCRCSingleByteSequence(t *testing.T) {
	// Test that hashing byte by byte produces different results
	crc32s := make([]uint32, 256)
	crc64s := make([]uint64, 256)

	for i := 0; i < 256; i++ {
		b := byte(i)
		crc32s[i] = CRC32([]byte{b})
		crc64s[i] = CRC64([]byte{b})
	}

	// Verify all are unique (CRC should have no collisions for single bytes)
	seen32 := make(map[uint32]bool)
	for _, c := range crc32s {
		if seen32[c] {
			t.Error("CRC32 collision detected for single byte values")
			break
		}
		seen32[c] = true
	}

	seen64 := make(map[uint64]bool)
	for _, c := range crc64s {
		if seen64[c] {
			t.Error("CRC64 collision detected for single byte values")
			break
		}
		seen64[c] = true
	}
}

// TestCRCStringVsByteSlice tests that string and []byte produce same results
func TestCRCStringVsByteSlice(t *testing.T) {
	testInputs := []string{
		"test",
		"hello world",
		"123456",
		"你好",
		"",
		"a",
		"The quick brown fox",
	}

	for _, input := range testInputs {
		t.Run(input, func(t *testing.T) {
			if CRC32(input) != CRC32([]byte(input)) {
				t.Error("CRC32 should produce same result for string and []byte")
			}
			if CRC64(input) != CRC64([]byte(input)) {
				t.Error("CRC64 should produce same result for string and []byte")
			}
		})
	}
}

// TestCRCLongMessages tests behavior with very long messages
func TestCRCLongMessages(t *testing.T) {
	sizes := []int{1000, 10000, 100000}

	for _, size := range sizes {
		t.Run("size_"+string(rune(size/1000+'0'))+"k", func(t *testing.T) {
			data := bytes.Repeat([]byte("A"), size)

			// Just verify they complete without error and produce consistent results
			crc32_1 := CRC32(data)
			crc32_2 := CRC32(data)
			if crc32_1 != crc32_2 {
				t.Error("CRC32 should be consistent for long messages")
			}

			crc64_1 := CRC64(data)
			crc64_2 := CRC64(data)
			if crc64_1 != crc64_2 {
				t.Error("CRC64 should be consistent for long messages")
			}
		})
	}
}

// TestCRCBitFlip tests that single bit flip changes the checksum
func TestCRCBitFlip(t *testing.T) {
	original := []byte("test data for bit flip detection")
	originalCRC32 := CRC32(original)
	originalCRC64 := CRC64(original)

	// Flip a bit in the middle
	modified := make([]byte, len(original))
	copy(modified, original)
	modified[len(modified)/2] ^= 0x01 // Flip the lowest bit

	modifiedCRC32 := CRC32(modified)
	modifiedCRC64 := CRC64(modified)

	// CRC should detect the bit flip
	if originalCRC32 == modifiedCRC32 {
		t.Error("CRC32 should detect single bit flip")
	}
	if originalCRC64 == modifiedCRC64 {
		t.Error("CRC64 should detect single bit flip")
	}
}

// TestCRCByteOrderSensitivity tests that byte order matters
func TestCRCByteOrderSensitivity(t *testing.T) {
	data1 := []byte{0x01, 0x02, 0x03, 0x04}
	data2 := []byte{0x04, 0x03, 0x02, 0x01}

	crc32_1 := CRC32(data1)
	crc32_2 := CRC32(data2)
	if crc32_1 == crc32_2 {
		t.Error("CRC32 should be sensitive to byte order")
	}

	crc64_1 := CRC64(data1)
	crc64_2 := CRC64(data2)
	if crc64_1 == crc64_2 {
		t.Error("CRC64 should be sensitive to byte order")
	}
}

// TestCRCKnownVectors tests against known CRC values
func TestCRCKnownVectors(t *testing.T) {
	// Known test vectors for "123456789"
	t.Run("CRC32 known vector", func(t *testing.T) {
		input := "123456789"
		expected := uint32(0xcbf43926) // Known CRC32 for "123456789"
		result := CRC32(input)
		if result != expected {
			t.Errorf("CRC32('%s') = 0x%08x, expected 0x%08x", input, result, expected)
		}
	})

	t.Run("CRC64 known vector", func(t *testing.T) {
		input := "123456789"
		expected := uint64(0x995dc9bbdf1939fa) // Known CRC64-ECMA for "123456789"
		result := CRC64(input)
		if result != expected {
			t.Errorf("CRC64('%s') = 0x%016x, expected 0x%016x", input, result, expected)
		}
	})
}

// TestCRCIncrementalVsOneShot tests that incremental and one-shot produce different results
// (This test documents that our functions are one-shot only)
func TestCRCIncrementalVsOneShot(t *testing.T) {
	part1 := "Hello "
	part2 := "World"
	combined := "Hello World"

	// One-shot CRC
	crc32Combined := CRC32(combined)
	crc64Combined := CRC64(combined)

	// Separate CRCs (NOT incremental - just for comparison)
	crc32Part1 := CRC32(part1)
	crc32Part2 := CRC32(part2)
	crc64Part1 := CRC64(part1)
	crc64Part2 := CRC64(part2)

	// Document that simple XOR doesn't work for combining CRCs
	crc32Xor := crc32Part1 ^ crc32Part2
	if crc32Xor == crc32Combined {
		t.Log("Note: CRC32 XOR of parts happened to equal combined (very unlikely)")
	}

	crc64Xor := crc64Part1 ^ crc64Part2
	if crc64Xor == crc64Combined {
		t.Log("Note: CRC64 XOR of parts happened to equal combined (very unlikely)")
	}

	// Verify that the combined result is different from the parts
	if crc32Combined == crc32Part1 || crc32Combined == crc32Part2 {
		t.Log("Note: Combined CRC32 happened to equal one of the parts (unlikely)")
	}
	if crc64Combined == crc64Part1 || crc64Combined == crc64Part2 {
		t.Log("Note: Combined CRC64 happened to equal one of the parts (unlikely)")
	}
}

// BenchmarkCRC32 benchmarks CRC32 function
func BenchmarkCRC32String(b *testing.B) {
	input := "benchmark test data for CRC32 checksum"
	for i := 0; i < b.N; i++ {
		_ = CRC32(input)
	}
}

func BenchmarkCRC32Bytes(b *testing.B) {
	input := []byte("benchmark test data for CRC32 checksum")
	for i := 0; i < b.N; i++ {
		_ = CRC32(input)
	}
}

func BenchmarkCRC64String(b *testing.B) {
	input := "benchmark test data for CRC64 checksum"
	for i := 0; i < b.N; i++ {
		_ = CRC64(input)
	}
}

func BenchmarkCRC64Bytes(b *testing.B) {
	input := []byte("benchmark test data for CRC64 checksum")
	for i := 0; i < b.N; i++ {
		_ = CRC64(input)
	}
}

func BenchmarkCRC32LongMessage(b *testing.B) {
	input := bytes.Repeat([]byte("X"), 10000)
	for i := 0; i < b.N; i++ {
		_ = CRC32(input)
	}
}

func BenchmarkCRC64LongMessage(b *testing.B) {
	input := bytes.Repeat([]byte("X"), 10000)
	for i := 0; i < b.N; i++ {
		_ = CRC64(input)
	}
}

func BenchmarkCRC32SmallMessage(b *testing.B) {
	input := []byte("test")
	for i := 0; i < b.N; i++ {
		_ = CRC32(input)
	}
}

func BenchmarkCRC64SmallMessage(b *testing.B) {
	input := []byte("test")
	for i := 0; i < b.N; i++ {
		_ = CRC64(input)
	}
}
