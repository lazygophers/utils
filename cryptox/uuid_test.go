package cryptox

import (
	"strings"
	"testing"

	"github.com/google/uuid"
)

// TestUUID tests the UUID function
func TestUUID(t *testing.T) {
	uuid := UUID()

	// Test length (UUID v4 without dashes is 32 characters)
	if len(uuid) != 32 {
		t.Errorf("Expected UUID length 32, got %d", len(uuid))
	}

	// Test that it contains no dashes
	if strings.Contains(uuid, "-") {
		t.Error("UUID should not contain dashes")
	}

	// Test that it's not empty
	if uuid == "" {
		t.Error("UUID should not be empty")
	}

	// Test that it only contains hexadecimal characters
	for _, c := range uuid {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			t.Errorf("UUID contains non-hexadecimal character: %c", c)
		}
	}
}

// TestUUIDUniqueness tests that UUID generates unique values
func TestUUIDUniqueness(t *testing.T) {
	const iterations = 1000
	seen := make(map[string]bool, iterations)

	for i := 0; i < iterations; i++ {
		uuid := UUID()
		if seen[uuid] {
			t.Errorf("UUID collision detected: %s was generated twice", uuid)
		}
		seen[uuid] = true
	}

	// Verify we actually generated all UUIDs
	if len(seen) != iterations {
		t.Errorf("Expected %d unique UUIDs, got %d", iterations, len(seen))
	}
}

// TestUUIDFormat tests UUID format correctness
func TestUUIDFormat(t *testing.T) {
	// UUID v4 format (with dashes): xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx
	// Without dashes: 32 hex characters
	// We should verify it's a valid hex string

	// Test multiple times to ensure consistency
	for i := 0; i < 10; i++ {
		uuid := UUID()

		// Length check
		if len(uuid) != 32 {
			t.Errorf("Iteration %d: Expected length 32, got %d", i, len(uuid))
		}

		// All characters should be valid hex
		validHex := true
		for _, c := range uuid {
			if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
				validHex = false
				t.Errorf("Iteration %d: Invalid hex character %c in UUID %s", i, c, uuid)
				break
			}
		}

		if !validHex {
			t.Errorf("Iteration %d: UUID contains invalid hex characters: %s", i, uuid)
		}
	}
}

// TestUUIDConcurrency tests UUID generation under concurrent access
func TestUUIDConcurrency(t *testing.T) {
	const goroutines = 100
	const uuidsPerGoroutine = 10

	results := make(chan string, goroutines*uuidsPerGoroutine)

	// Launch multiple goroutines generating UUIDs
	for i := 0; i < goroutines; i++ {
		go func() {
			for j := 0; j < uuidsPerGoroutine; j++ {
				results <- UUID()
			}
		}()
	}

	// Collect all UUIDs
	seen := make(map[string]bool, goroutines*uuidsPerGoroutine)
	for i := 0; i < goroutines*uuidsPerGoroutine; i++ {
		uuid := <-results

		// Check basic properties
		if len(uuid) != 32 {
			t.Errorf("UUID has wrong length: %d (expected 32)", len(uuid))
		}

		if strings.Contains(uuid, "-") {
			t.Error("UUID contains dashes")
		}

		// Check for collisions
		if seen[uuid] {
			t.Errorf("UUID collision in concurrent test: %s", uuid)
		}
		seen[uuid] = true
	}

	// Verify we got all expected UUIDs
	expectedCount := goroutines * uuidsPerGoroutine
	if len(seen) != expectedCount {
		t.Errorf("Expected %d unique UUIDs, got %d", expectedCount, len(seen))
	}
}

// TestUUIDNotEmpty tests that UUID never returns empty string
func TestUUIDNotEmpty(t *testing.T) {
	for i := 0; i < 100; i++ {
		uuid := UUID()
		if uuid == "" {
			t.Errorf("Iteration %d: UUID returned empty string", i)
		}
	}
}

// BenchmarkUUID benchmarks UUID generation performance
func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = UUID()
	}
}

// BenchmarkUUIDParallel benchmarks UUID generation under parallel load
func BenchmarkUUIDParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = UUID()
		}
	})
}

// ============== 性能优化基准测试 ==============

// BenchmarkUUID_Opt1_FixedIndexes 固定索引位置优化
func BenchmarkUUID_Opt1_FixedIndexes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := uuid.New().String()
		result := make([]byte, 32)
		copy(result[0:8], s[0:8])
		copy(result[8:12], s[9:13])
		copy(result[12:16], s[14:18])
		copy(result[16:20], s[19:23])
		copy(result[20:32], s[24:36])
		_ = string(result)
	}
}

// BenchmarkUUID_Opt2_StringConcat 字符串拼接优化
func BenchmarkUUID_Opt2_StringConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := uuid.New().String()
		_ = s[0:8] + s[9:13] + s[14:18] + s[19:23] + s[24:36]
	}
}

// BenchmarkUUID_Opt3_ArrayCopy 数组复制优化
func BenchmarkUUID_Opt3_ArrayCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := uuid.New().String()
		var result [32]byte
		copy(result[0:8], s[0:8])
		copy(result[8:12], s[9:13])
		copy(result[12:16], s[14:18])
		copy(result[16:20], s[19:23])
		copy(result[20:32], s[24:36])
		_ = string(result[:])
	}
}

// BenchmarkUUID_Opt4_HexEncode 直接 hex 编码优化
func BenchmarkUUID_Opt4_HexEncode(b *testing.B) {
	hex := "0123456789abcdef"
	for i := 0; i < b.N; i++ {
		id := uuid.New()
		bytes := id[:]
		result := make([]byte, 32)
		for j := 0; j < 16; j++ {
			result[j*2] = hex[bytes[j]>>4]
			result[j*2+1] = hex[bytes[j]&0x0F]
		}
		_ = string(result)
	}
}

// BenchmarkUUID_Opt5_PreAllocAppend 预分配 append 优化
func BenchmarkUUID_Opt5_PreAllocAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := uuid.New().String()
		result := make([]byte, 0, 32)
		result = append(result, s[0:8]...)
		result = append(result, s[9:13]...)
		result = append(result, s[14:18]...)
		result = append(result, s[19:23]...)
		result = append(result, s[24:36]...)
		_ = string(result)
	}
}

// BenchmarkUUID_Opt6_ByteLoop 字节循环优化
func BenchmarkUUID_Opt6_ByteLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := uuid.New().String()
		result := make([]byte, 0, 32)
		for j := 0; j < len(s); j++ {
			if s[j] != '-' {
				result = append(result, s[j])
			}
		}
		_ = string(result)
	}
}

// BenchmarkUUID_Opt7_StringBuilder strings.Builder 优化
func BenchmarkUUID_Opt7_StringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := uuid.New().String()
		var builder strings.Builder
		builder.Grow(32)
		for j := 0; j < len(s); j++ {
			if s[j] != '-' {
				builder.WriteByte(s[j])
			}
		}
		_ = builder.String()
	}
}

// BenchmarkUUID_Opt8_HybridOpt 混合优化
func BenchmarkUUID_Opt8_HybridOpt(b *testing.B) {
	hexTable := [16]byte{
		'0', '1', '2', '3', '4', '5', '6', '7',
		'8', '9', 'a', 'b', 'c', 'd', 'e', 'f',
	}
	for i := 0; i < b.N; i++ {
		id := uuid.New()
		bytes := id[:]
		var result [32]byte
		for j := 0; j < 16; j++ {
			v := bytes[j]
			result[j*2] = hexTable[v>>4]
			result[j*2+1] = hexTable[v&0x0F]
		}
		_ = string(result[:])
	}
}

// BenchmarkUUID_Opt9_NewVsNewString uuid.New() vs uuid.NewString()
func BenchmarkUUID_Opt9_NewVsNewString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		id := uuid.New()
		s := id.String()
		result := make([]byte, 32)
		copy(result[0:8], s[0:8])
		copy(result[8:12], s[9:13])
		copy(result[12:16], s[14:18])
		copy(result[16:20], s[19:23])
		copy(result[20:32], s[24:36])
		_ = string(result)
	}
}

// BenchmarkUUID_Opt10_ReplaceAll 当前实现（对比基线）
func BenchmarkUUID_Opt10_ReplaceAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.ReplaceAll(uuid.NewString(), "-", "")
	}
}
