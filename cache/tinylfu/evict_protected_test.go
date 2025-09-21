package tinylfu

import (
	"testing"
)

// TestEvictFromProtectedSpecific creates the exact condition to trigger evictFromProtected
func TestEvictFromProtectedSpecific(t *testing.T) {
	t.Run("force_evictFromProtected_exact_condition", func(t *testing.T) {
		// Create cache and build a state where protected has items but window/probation are empty
		cache, err := New[string, int](20)
		if err != nil {
			t.Fatalf("Failed to create cache: %v", err)
		}

		// Step 1: Build high frequency for our protected keys
		protectedKeys := []string{"HIGH_FREQ_1", "HIGH_FREQ_2", "HIGH_FREQ_3", "HIGH_FREQ_4", "HIGH_FREQ_5"}

		// Build very high frequency in the sketch
		for round := 0; round < 50; round++ {
			for _, key := range protectedKeys {
				cache.Put(key, 1)
				for i := 0; i < 5; i++ {
					cache.Get(key)
				}
			}
		}

		// Step 2: Force these items into protected segment
		// Use the successful pattern from the demoteFromProtected test

		// Add victim items first
		for i := 0; i < 5; i++ {
			cache.Put("victim", i)
		}

		// Now force window evictions to get our high-freq items into probation
		for _, key := range protectedKeys {
			cache.Put(key, 100) // goes to window
		}

		// Force evictions from window to trigger admission policy
		for i := 0; i < 20; i++ {
			cache.Put("temp", i+200)
		}

		// Re-add our high-frequency keys - they should be admitted to probation
		for _, key := range protectedKeys {
			cache.Put(key, 100)
		}

		// Access them to promote from probation to protected
		for _, key := range protectedKeys {
			if cache.Contains(key) {
				cache.Get(key) // promote to protected
			}
		}

		stats := cache.Stats()
		t.Logf("Before clearing other segments: Protected=%d, Probation=%d, Window=%d",
			stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)

		// Step 3: Clear window and probation by adding many low-frequency items
		// These will fill window and then get evicted without being admitted
		for i := 0; i < 50; i++ {
			cache.Put("low_freq", i+1000) // Low frequency items that won't be admitted
		}

		midStats := cache.Stats()
		t.Logf("After clearing other segments: Protected=%d, Probation=%d, Window=%d",
			midStats.ProtectedSize, midStats.ProbationSize, midStats.WindowSize)

		// Step 4: Now resize to force eviction from protected
		// The resize loop should: try window (empty), try probation (empty), then evict from protected
		err = cache.Resize(2)
		if err != nil {
			t.Fatalf("Failed to resize: %v", err)
		}

		finalStats := cache.Stats()
		t.Logf("After resize triggering evictFromProtected: Protected=%d, Probation=%d, Window=%d",
			finalStats.ProtectedSize, finalStats.ProbationSize, finalStats.WindowSize)

		if cache.Len() > 2 {
			t.Errorf("Cache size %d exceeds capacity 2", cache.Len())
		}
	})

	t.Run("direct_protected_only_scenario", func(t *testing.T) {
		// Create a scenario where we have items only in protected
		cache, err := New[string, int](10)
		if err != nil {
			t.Fatalf("Failed to create cache: %v", err)
		}

		// Use the working pattern from our successful demoteFromProtected test
		keys := []string{"PROT_A", "PROT_B", "PROT_C", "PROT_D", "PROT_E"}

		// Build frequency
		for round := 0; round < 30; round++ {
			for _, key := range keys {
				cache.Put(key, 1)
				for i := 0; i < 3; i++ {
					cache.Get(key)
				}
			}
		}

		// Add victim
		for i := 0; i < 5; i++ {
			cache.Put("victim", i)
		}

		// Force window evictions to get our items into probation
		for i := 0; i < 50; i++ {
			cache.Put("temp", i)
		}

		// Re-add our keys - should be admitted to probation due to frequency
		for _, key := range keys {
			cache.Put(key, 100)
		}

		// Promote all to protected
		for _, key := range keys {
			if cache.Contains(key) {
				cache.Get(key)
			}
		}

		stats := cache.Stats()
		t.Logf("Built protected segment: Protected=%d, Probation=%d, Window=%d, Total=%d",
			stats.ProtectedSize, stats.ProbationSize, stats.WindowSize, cache.Len())

		// Now manually manipulate to ensure window and probation are empty
		// Add one final low-frequency item to window
		cache.Put("final_window_item", 999)

		// Clear window by adding another item (this one will stay in window but first gets evicted)
		cache.Put("stay_in_window", 888)

		// Now resize to trigger evictFromProtected path
		err = cache.Resize(1)
		if err != nil {
			t.Fatalf("Failed to resize: %v", err)
		}

		finalStats := cache.Stats()
		t.Logf("After resize to 1: Protected=%d, Probation=%d, Window=%d, Total=%d",
			finalStats.ProtectedSize, finalStats.ProbationSize, finalStats.WindowSize, cache.Len())

		if cache.Len() > 1 {
			t.Errorf("Cache size %d exceeds capacity 1", cache.Len())
		}
	})
}

// TestEvictFromProtectedDirectPath uses the exact same pattern that worked for demoteFromProtected
func TestEvictFromProtectedDirectPath(t *testing.T) {
	// Use exact same approach as the successful demoteFromProtected test
	cache, err := New[string, int](25)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Step 1: Build frequency for 25 keys (same as successful test)
	keys := make([]string, 25)
	for i := 0; i < 25; i++ {
		keys[i] = "EVICT_TEST_KEY_" + string(rune('A'+i%26)) + string(rune('0'+(i/26)))
	}

	// Step 2: Build frequency patterns (exactly like successful test)
	for round := 0; round < 20; round++ {
		for i := 0; i < 25; i++ {
			key := keys[i]
			cache.Put(key, i)

			// Higher frequency for some keys
			accessCount := 1
			if i < 20 {
				accessCount = 3
			}

			for j := 0; j < accessCount; j++ {
				cache.Get(key)
			}
		}
	}

	// Step 3: Add victims
	for i := 0; i < 5; i++ {
		victimKey := "EVICT_VICTIM_" + string(rune('A'+i))
		cache.Put(victimKey, i+1000)
	}

	// Step 4: Force window evictions (same as successful test)
	for i := 0; i < 50; i++ {
		tempKey := "EVICT_TEMP_" + string(rune('A'+(i%26)))
		cache.Put(tempKey, i+2000)
	}

	// Step 5: Promote high-frequency keys (same as successful test)
	for i := 0; i < 20; i++ {
		key := keys[i]
		if cache.Contains(key) {
			cache.Get(key)
		}
	}

	stats := cache.Stats()
	t.Logf("After building protected like successful test: Protected=%d, Probation=%d, Window=%d",
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)

	// Step 6: Now resize aggressively to trigger evictFromProtected
	// Since we successfully built protected segment, now force resize eviction

	err = cache.Resize(10) // First resize
	if err != nil {
		t.Fatalf("Failed to resize to 10: %v", err)
	}

	midStats := cache.Stats()
	t.Logf("After resize to 10: Protected=%d, Probation=%d, Window=%d",
		midStats.ProtectedSize, midStats.ProbationSize, midStats.WindowSize)

	err = cache.Resize(5) // Second resize
	if err != nil {
		t.Fatalf("Failed to resize to 5: %v", err)
	}

	err = cache.Resize(2) // Final resize to force evictFromProtected
	if err != nil {
		t.Fatalf("Failed to resize to 2: %v", err)
	}

	finalStats := cache.Stats()
	t.Logf("After final resize to 2: Protected=%d, Probation=%d, Window=%d",
		finalStats.ProtectedSize, finalStats.ProbationSize, finalStats.WindowSize)

	if cache.Len() > 2 {
		t.Errorf("Cache size %d exceeds capacity 2", cache.Len())
	}
}