package cryptox

import (
	"strings"
	"testing"
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
