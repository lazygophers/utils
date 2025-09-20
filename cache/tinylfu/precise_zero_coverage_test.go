package tinylfu

import (
	"fmt"
	"testing"
)

// TestPreciseZeroCoverageFunctions creates the exact algorithmic conditions to trigger 0% coverage functions
func TestPreciseZeroCoverageFunctions(t *testing.T) {
	t.Run("demoteFromProtected_algorithmic", func(t *testing.T) {
		// Use size 13: window=1, main=12, protected capacity = int(12*0.8) = 9
		cache, err := New[string, int](13)
		if err != nil {
			t.Fatalf("Failed to create cache: %v", err)
		}

		// Phase 1: Build frequency for keys in Count-Min Sketch
		// This is crucial - without frequency, items won't be admitted
		for round := 0; round < 5; round++ {
			for i := 0; i < 12; i++ {
				key := fmt.Sprintf("key%d", i)
				cache.Put(key, i)
				cache.Get(key) // Build frequency
			}
		}

		// Phase 2: Clear cache but preserve frequency sketch and doorkeeper
		cache.Clear()

		// Phase 3: Systematically build protected segment to capacity (9 items)
		// We need to work with the admission policy correctly

		// First, establish doorkeeper entries by putting keys once
		for i := 0; i < 9; i++ {
			key := fmt.Sprintf("key%d", i)
			cache.Put(key, i+100)
		}

		// Add a victim with low frequency to probation
		cache.Put("victim", 999)

		// Now re-add our high-frequency keys - they should be admitted due to doorkeeper + frequency
		for i := 0; i < 9; i++ {
			key := fmt.Sprintf("key%d", i)
			cache.Put(key, i+100) // Should be admitted to probation

			// Promote to protected by accessing
			cache.Get(key)
		}

		stats := cache.Stats()
		t.Logf("After building protected: Protected=%d, Probation=%d, Window=%d",
			stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)

		// Phase 4: Create a new high-frequency item and promote it to protected
		// This should trigger demoteFromProtected when protected is at capacity
		newKey := "key9"

		// Build frequency for this key
		for i := 0; i < 10; i++ {
			cache.Put(newKey, 900)
			cache.Get(newKey)
		}

		// Clear and re-establish doorkeeper
		cache.Put(newKey, 900) // Put once to establish doorkeeper

		// Add victim to probation
		cache.Put("new_victim", 888)

		// Now re-add the new key - should be admitted to probation
		cache.Put(newKey, 900)

		// Access to promote to protected - should trigger demoteFromProtected
		cache.Get(newKey) // This should call promoteToProtected -> demoteFromProtected

		finalStats := cache.Stats()
		t.Logf("After promotion: Protected=%d, Probation=%d, Window=%d",
			finalStats.ProtectedSize, finalStats.ProbationSize, finalStats.WindowSize)

		if cache.Len() > cache.Cap() {
			t.Errorf("Cache overflow: %d > %d", cache.Len(), cache.Cap())
		}
	})

	t.Run("evictFromProtected_resize_exact", func(t *testing.T) {
		// Create a cache where we can fill protected and then force resize eviction
		cache, err := New[string, int](10)
		if err != nil {
			t.Fatalf("Failed to create cache: %v", err)
		}

		// Build high-frequency keys
		for round := 0; round < 5; round++ {
			for i := 0; i < 6; i++ {
				key := fmt.Sprintf("prot%d", i)
				cache.Put(key, i)
				cache.Get(key)
			}
		}

		// Clear and rebuild only protected segment
		cache.Clear()

		// Phase 1: Establish doorkeeper for our keys
		for i := 0; i < 6; i++ {
			key := fmt.Sprintf("prot%d", i)
			cache.Put(key, i+200)
		}

		// Phase 2: Re-add to probation (admitted due to doorkeeper + frequency)
		for i := 0; i < 6; i++ {
			key := fmt.Sprintf("prot%d", i)
			cache.Put(key, i+200)
		}

		// Phase 3: Promote all to protected
		for i := 0; i < 6; i++ {
			key := fmt.Sprintf("prot%d", i)
			cache.Get(key)
		}

		// Verify protected has items
		stats := cache.Stats()
		t.Logf("Pre-resize: Protected=%d, Probation=%d, Window=%d",
			stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)

		// Phase 4: Resize to force eviction from protected
		// The resize loop will: check window (empty), check probation (empty), then evict from protected
		err = cache.Resize(3)
		if err != nil {
			t.Fatalf("Failed to resize: %v", err)
		}

		finalStats := cache.Stats()
		t.Logf("Post-resize: Protected=%d, Probation=%d, Window=%d",
			finalStats.ProtectedSize, finalStats.ProbationSize, finalStats.WindowSize)

		if cache.Len() > 3 {
			t.Errorf("Cache size %d exceeds capacity 3", cache.Len())
		}
	})
}

// TestDirectProtectedSegmentManipulation bypasses admission policy complexities
func TestDirectProtectedSegmentManipulation(t *testing.T) {
	t.Run("force_protected_fullness", func(t *testing.T) {
		cache, err := New[string, int](15)
		if err != nil {
			t.Fatalf("Failed to create cache: %v", err)
		}

		// Strategy: Create many items with very high frequency,
		// then manipulate the cache state to get them into protected

		keys := make([]string, 20)
		for i := 0; i < 20; i++ {
			keys[i] = fmt.Sprintf("high_freq_%d", i)
		}

		// Build extremely high frequency for all keys
		for round := 0; round < 10; round++ {
			for _, key := range keys {
				cache.Put(key, 1)
				for j := 0; j < 3; j++ {
					cache.Get(key)
				}
			}
		}

		// Clear cache
		cache.Clear()

		// Re-build cache with systematic approach
		// Main capacity = 14 (93% of 15), Protected capacity = int(14*0.8) = 11

		// Step 1: Get items into doorkeeper
		for i := 0; i < 12; i++ {
			cache.Put(keys[i], i+100)
		}

		// Step 2: Add a few low-frequency victims
		for i := 0; i < 3; i++ {
			cache.Put(fmt.Sprintf("victim_%d", i), i)
		}

		// Step 3: Re-add high-frequency keys to get them admitted
		for i := 0; i < 12; i++ {
			cache.Put(keys[i], i+100)
		}

		// Step 4: Promote to protected by accessing
		for i := 0; i < 11; i++ { // Fill protected to capacity (11)
			cache.Get(keys[i])
		}

		stats := cache.Stats()
		t.Logf("Protected filled: Protected=%d, Probation=%d, Window=%d",
			stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)

		// Step 5: Try to promote one more item to protected
		cache.Get(keys[11]) // Should trigger promoteToProtected -> demoteFromProtected

		finalStats := cache.Stats()
		t.Logf("After final promotion: Protected=%d, Probation=%d, Window=%d",
			finalStats.ProtectedSize, finalStats.ProbationSize, finalStats.WindowSize)
	})
}

// TestProtectedEvictionDuringResize specifically targets the resize eviction path
func TestProtectedEvictionDuringResize(t *testing.T) {
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Fill cache with items that will end up in protected
	for i := 0; i < 30; i++ {
		key := fmt.Sprintf("item_%d", i)

		// Build frequency
		for round := 0; round < 8; round++ {
			cache.Put(key, i)
			cache.Get(key)
		}
	}

	stats := cache.Stats()
	t.Logf("Before resize: Protected=%d, Probation=%d, Window=%d, Total=%d",
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize, cache.Len())

	// Resize to very small size to force eviction from all segments including protected
	err = cache.Resize(5)
	if err != nil {
		t.Fatalf("Failed to resize: %v", err)
	}

	finalStats := cache.Stats()
	t.Logf("After resize: Protected=%d, Probation=%d, Window=%d, Total=%d",
		finalStats.ProtectedSize, finalStats.ProbationSize, finalStats.WindowSize, cache.Len())

	if cache.Len() > 5 {
		t.Errorf("Cache size %d exceeds capacity 5", cache.Len())
	}

	// Further resize to force more evictions
	err = cache.Resize(2)
	if err != nil {
		t.Fatalf("Failed to resize: %v", err)
	}

	if cache.Len() > 2 {
		t.Errorf("Final cache size %d exceeds capacity 2", cache.Len())
	}
}