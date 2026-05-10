package cryptox

import (
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
)

// TestULID tests the ULID function
func TestULID(t *testing.T) {
	id := ULID()

	// ULID is 26 characters
	if len(id) != 26 {
		t.Errorf("Expected ULID length 26, got %d", len(id))
	}

	// Test that it's not empty
	if id == "" {
		t.Error("ULID should not be empty")
	}
}

// TestULID_Uniqueness tests that ULID generates unique values
func TestULID_Uniqueness(t *testing.T) {
	const iterations = 1000
	seen := make(map[string]bool, iterations)

	for i := 0; i < iterations; i++ {
		id := ULID()
		if seen[id] {
			t.Errorf("ULID collision detected: %s was generated twice", id)
		}
		seen[id] = true
	}

	if len(seen) != iterations {
		t.Errorf("Expected %d unique ULIDs, got %d", iterations, len(seen))
	}
}

// TestULID_Format tests ULID format correctness (Crockford Base32)
func TestULID_Format(t *testing.T) {
	// Crockford Base32 character set: 0123456789ABCDEFGHJKMNPQRSTVWXYZ
	crockfordPattern := regexp.MustCompile(`^[0123456789ABCDEFGHJKMNPQRSTVWXYZ]{26}$`)

	for i := 0; i < 10; i++ {
		id := ULID()

		if len(id) != 26 {
			t.Errorf("Iteration %d: Expected length 26, got %d", i, len(id))
		}

		if !crockfordPattern.MatchString(id) {
			t.Errorf("Iteration %d: ULID %s does not match Crockford Base32 format", i, id)
		}

		// Should not contain excluded characters (I, L, O, U)
		for _, c := range "ILOU" {
			if strings.ContainsRune(id, c) {
				t.Errorf("Iteration %d: ULID %s contains excluded Crockford Base32 character: %c", i, id, c)
			}
		}
	}
}

// TestULID_Concurrent tests ULID generation under concurrent access
func TestULID_Concurrent(t *testing.T) {
	const goroutines = 100
	const idsPerGoroutine = 10

	results := make(chan string, goroutines*idsPerGoroutine)

	for i := 0; i < goroutines; i++ {
		go func() {
			for j := 0; j < idsPerGoroutine; j++ {
				results <- ULID()
			}
		}()
	}

	seen := make(map[string]bool, goroutines*idsPerGoroutine)
	for i := 0; i < goroutines*idsPerGoroutine; i++ {
		id := <-results

		if len(id) != 26 {
			t.Errorf("ULID has wrong length: %d (expected 26)", len(id))
		}

		if seen[id] {
			t.Errorf("ULID collision in concurrent test: %s", id)
		}
		seen[id] = true
	}

	expectedCount := goroutines * idsPerGoroutine
	if len(seen) != expectedCount {
		t.Errorf("Expected %d unique ULIDs, got %d", expectedCount, len(seen))
	}
}

// TestULIDWithTimestamp tests ULIDWithTimestamp function
func TestULIDWithTimestamp(t *testing.T) {
	before := time.Now().UnixMilli()
	id, ts := ULIDWithTimestamp()
	after := time.Now().UnixMilli()

	if len(id) != 26 {
		t.Errorf("Expected ULID length 26, got %d", len(id))
	}

	if ts < before || ts > after {
		t.Errorf("Timestamp %d is not within expected range [%d, %d]", ts, before, after)
	}
}

// TestULIDWithTimestamp_Consistency tests that ULID and timestamp match
func TestULIDWithTimestamp_Consistency(t *testing.T) {
	id, ts := ULIDWithTimestamp()

	// Parse the ULID and extract its timestamp
	extractedTs, err := GetULIDTimestamp(id)
	if err != nil {
		t.Fatalf("Failed to extract timestamp from ULID: %v", err)
	}

	if extractedTs != ts {
		t.Errorf("Timestamp mismatch: ULIDWithTimestamp returned %d, GetULIDTimestamp extracted %d", ts, extractedTs)
	}
}

// TestGetULIDTimestamp tests extracting timestamp from a valid ULID
func TestGetULIDTimestamp(t *testing.T) {
	before := time.Now().UnixMilli()
	id := ULID()
	after := time.Now().UnixMilli()

	ts, err := GetULIDTimestamp(id)
	if err != nil {
		t.Fatalf("GetULIDTimestamp failed: %v", err)
	}

	if ts < before || ts > after {
		t.Errorf("Extracted timestamp %d is not within expected range [%d, %d]", ts, before, after)
	}
}

// TestGetULIDTimestamp_InvalidInput tests GetULIDTimestamp with invalid input
func TestGetULIDTimestamp_InvalidInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"short string", "abc"},
		{"overflow value", "8ZZZZZZZZZZZZZZZZZZZZZZZZZ"},
		{"too long", "0000000000000000000000000000"},
		{"random text", "not-a-valid-ulid"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetULIDTimestamp(tt.input)
			if err == nil {
				t.Errorf("Expected error for input %q, got nil", tt.input)
			}
		})
	}
}

// TestGetULIDTimestamp_EmptyString tests GetULIDTimestamp with empty string
func TestGetULIDTimestamp_EmptyString(t *testing.T) {
	_, err := GetULIDTimestamp("")
	if err == nil {
		t.Error("Expected error for empty string, got nil")
	}
}

// TestMustGetULIDTimestamp tests MustGetULIDTimestamp with valid input
func TestMustGetULIDTimestamp(t *testing.T) {
	before := time.Now().UnixMilli()
	id := ULID()
	after := time.Now().UnixMilli()

	ts := MustGetULIDTimestamp(id)

	if ts < before || ts > after {
		t.Errorf("Extracted timestamp %d is not within expected range [%d, %d]", ts, before, after)
	}
}

// TestMustGetULIDTimestamp_Panic tests MustGetULIDTimestamp panics on invalid input
func TestMustGetULIDTimestamp_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for invalid ULID, but did not panic")
		}
	}()
	_ = MustGetULIDTimestamp("invalid")
}

// BenchmarkULID benchmarks ULID generation performance
func BenchmarkULID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ULID()
	}
}

// BenchmarkULIDWithTimestamp benchmarks ULIDWithTimestamp performance
func BenchmarkULIDWithTimestamp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ULIDWithTimestamp()
	}
}

// BenchmarkGetULIDTimestamp benchmarks GetULIDTimestamp performance
func BenchmarkGetULIDTimestamp(b *testing.B) {
	id := ULID()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = GetULIDTimestamp(id)
	}
}

// ============== 性能优化基准测试 ==============

// BenchmarkULID_Opt1_Monotonic 使用单调熵源优化
func BenchmarkULID_Opt1_Monotonic(b *testing.B) {
	entropy := ulid.Monotonic(randReader, 0)
	for i := 0; i < b.N; i++ {
		_ = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
	}
}

// BenchmarkULID_Opt2_PreAllocBuffer 预分配缓冲区优化
func BenchmarkULID_Opt2_PreAllocBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		id := ulid.Make()
		buf := make([]byte, 26)
		_ = id.MarshalTextTo(buf)
		_ = string(buf)
	}
}

// BenchmarkULID_Opt3_ArrayBuffer 数组缓冲优化
func BenchmarkULID_Opt3_ArrayBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		id := ulid.Make()
		var buf [26]byte
		_ = id.MarshalTextTo(buf[:])
		_ = string(buf[:])
	}
}

// BenchmarkULIDWithTimestamp_Opt4_OptimizedOrder 优化顺序
func BenchmarkULIDWithTimestamp_Opt4_OptimizedOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		id := ulid.Make()
		timestamp := int64(id.Time())
		str := id.String()
		_, _ = str, timestamp
	}
}

// BenchmarkULIDWithTimestamp_Opt5_SingleEncode 单次编码优化
func BenchmarkULIDWithTimestamp_Opt5_SingleEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		id := ulid.Make()
		var buf [26]byte
		_ = id.MarshalTextTo(buf[:])
		timestamp := int64(id.Time())
		_ = string(buf[:])
		_ = timestamp
	}
}

// BenchmarkGetULIDTimestamp_Opt6_MustParse 直接解析优化
func BenchmarkGetULIDTimestamp_Opt6_MustParse(b *testing.B) {
	testID := ulid.Make().String()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		parsedId := ulid.MustParse(testID)
		_ = int64(parsedId.Time())
	}
}

// BenchmarkGetULIDTimestamp_Opt7_CachedParsed 缓存解析结果优化
func BenchmarkGetULIDTimestamp_Opt7_CachedParsed(b *testing.B) {
	testID := ulid.MustParse(ulid.Make().String())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = int64(testID.Time())
	}
}

// BenchmarkGetULIDTimestamp_Opt8_DirectBytes 从字节提取优化
func BenchmarkGetULIDTimestamp_Opt8_DirectBytes(b *testing.B) {
	testID := ulid.Make()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes := testID.Bytes()
		_ = int64(bytesToTimestamp(bytes))
	}
}

// 辅助函数：从字节提取时间戳
func bytesToTimestamp(b []byte) uint64 {
	return uint64(b[0])<<40 |
		uint64(b[1])<<32 |
		uint64(b[2])<<24 |
		uint64(b[3])<<16 |
		uint64(b[4])<<8 |
		uint64(b[5])
}
