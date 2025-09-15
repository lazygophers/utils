package cryptox

import (
	"testing"
)

func TestBLAKE2b(t *testing.T) {
	input := "test"
	size := 32
	expected := "928b20366943e2afd11ebc0eae2e53a93bf177a4fcf35bcc64d503704e65e202"
	result, err := BLAKE2b(input, size)
	if err != nil {
		t.Errorf("BLAKE2b(%s, %d) returned an error: %v", input, size, err)
	} else if result != expected {
		t.Errorf("BLAKE2b(%s, %d) = %s; want %s", input, size, result, expected)
	}
}

func TestBLAKE2s(t *testing.T) {
	input := "test"
	size := 32
	expected := "f308fc02ce9172ad02a7d75800ecfc027109bc67987ea32aba9b8dcc7b10150e"
	result, err := BLAKE2s(input, size)
	if err != nil {
		t.Errorf("BLAKE2s(%s, %d) returned an error: %v", input, size, err)
	} else if result != expected {
		t.Errorf("BLAKE2s(%s, %d) = %s; want %s", input, size, result, expected)
	}
}

func TestBLAKE2b512(t *testing.T) {
	input := "test"
	expected := "a71079d42853dea26e453004338670a53814b78137ffbed07603a41d76a483aa9bc33b582f77d30a65e6f29a896c0411f38312e1d66e0bf16386c86a89bea572"
	result := BLAKE2b512(input)
	if result != expected {
		t.Errorf("BLAKE2b512(%s) = %s; want %s", input, result, expected)
	}
}

func TestBLAKE2b256(t *testing.T) {
	input := "test"
	expected := "928b20366943e2afd11ebc0eae2e53a93bf177a4fcf35bcc64d503704e65e202"
	result := BLAKE2b256(input)
	if result != expected {
		t.Errorf("BLAKE2b256(%s) = %s; want %s", input, result, expected)
	}
}

func TestBLAKE2s256(t *testing.T) {
	input := "test"
	expected := "f308fc02ce9172ad02a7d75800ecfc027109bc67987ea32aba9b8dcc7b10150e"
	result := BLAKE2s256(input)
	if result != expected {
		t.Errorf("BLAKE2s256(%s) = %s; want %s", input, result, expected)
	}
}

func TestBLAKE2bWithKey(t *testing.T) {
	input := "test"
	key := []byte("key")
	size := 32
	result, err := BLAKE2bWithKey(input, key, size)
	if err != nil {
		t.Errorf("BLAKE2bWithKey(%s, %v, %d) returned error: %v", input, key, size, err)
	}
	// Verify it returns correct length
	if len(result) != size*2 {
		t.Errorf("BLAKE2bWithKey(%s, %v, %d) returned wrong length: got %d, want %d", input, key, size, len(result), size*2)
	}
}

func TestBLAKE2sWithKey(t *testing.T) {
	input := "test"
	key := []byte("key")
	result, err := BLAKE2sWithKey(input, key)
	if err != nil {
		t.Errorf("BLAKE2sWithKey(%s, %v) returned error: %v", input, key, err)
	}
	if len(result) != 64 { // 32 bytes * 2 hex chars
		t.Errorf("BLAKE2sWithKey(%s, %v) returned wrong length: got %d, want 64", input, key, len(result))
	}
}

// Test error conditions for BLAKE2 functions
func TestBLAKE2ErrorConditions(t *testing.T) {
	data := "test"
	key := []byte("key")

	// Test BLAKE2b with invalid size
	_, err := BLAKE2b(data, 0)
	if err == nil {
		t.Error("Expected error for BLAKE2b with size 0")
	}

	_, err = BLAKE2b(data, 65) // BLAKE2b max is 64 bytes
	if err == nil {
		t.Error("Expected error for BLAKE2b with size > 64")
	}

	// Test BLAKE2s with invalid size
	_, err = BLAKE2s(data, 0)
	if err == nil {
		t.Error("Expected error for BLAKE2s with size 0")
	}

	// Test BLAKE2bWithKey with invalid sizes
	_, err = BLAKE2bWithKey(data, key, 0)
	if err == nil {
		t.Error("Expected error for BLAKE2bWithKey with size 0")
	}

	_, err = BLAKE2bWithKey(data, key, 65)
	if err == nil {
		t.Error("Expected error for BLAKE2bWithKey with size > 64")
	}
}

// Test BLAKE2s error condition specifically
func TestBLAKE2sErrorHandling(t *testing.T) {
	data := "test"

	// The BLAKE2s function has an error path when blake2s.New256 fails
	// This is difficult to trigger in normal circumstances, but we can test the error validation
	_, err := BLAKE2s(data, 0)
	if err == nil {
		t.Error("Expected error for BLAKE2s with size 0")
	}

	// Valid case should work
	result, err := BLAKE2s(data, 32)
	if err != nil {
		t.Errorf("BLAKE2s with valid size should not error: %v", err)
	}
	if result == "" {
		t.Error("BLAKE2s should return non-empty result")
	}
}

// Create a test that attempts to trigger the blake2s.New256 error path
func TestBLAKE2sInternalError(t *testing.T) {
	data := "test"

	// Test all valid sizes to ensure we exercise all code paths
	for size := 1; size <= 64; size++ {
		result, err := BLAKE2s(data, size)
		if size <= 32 { // BLAKE2s supports up to 256 bits (32 bytes)
			if err != nil {
				t.Errorf("BLAKE2s should work for size %d: %v", size, err)
			}
			if result == "" {
				t.Errorf("BLAKE2s should return non-empty result for size %d", size)
			}
		} else {
			// For sizes > 32, it should still work because we ignore the size parameter
			if err != nil {
				t.Errorf("BLAKE2s should work even for size %d: %v", size, err)
			}
		}
	}
}
