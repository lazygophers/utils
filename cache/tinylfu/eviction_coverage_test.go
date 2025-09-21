package tinylfu

import (
	"testing"
)

// TestProtectedEvictionCoverage tests the demoteFromProtected and evictFromProtected functions
func TestProtectedEvictionCoverage(t *testing.T) {
	t.Run("Force protected segment operations", func(t *testing.T) {
		// Create a cache with a small capacity to force evictions
		cache, err := New[string, int](10) // Very small cache
		if err != nil {
			t.Fatalf("Failed to create cache: %v", err)
		}

		// Fill the window segment first
		for i := 0; i < 15; i++ {
			key := "window" + string(rune('A'+i))
			cache.Put(key, i)
		}

		// Access some items multiple times to promote them to protected
		// This should trigger moves from window to probation, then to protected
		for i := 0; i < 5; i++ {
			for j := 0; j < 3; j++ { // Multiple accesses to trigger promotion
				cache.Get("window" + string(rune('A'+i)))
			}
		}

		// Add more items to force the protected segment to become full
		// and trigger demoteFromProtected and evictFromProtected
		for i := 0; i < 20; i++ {
			key := "force" + string(rune('A'+i))
			cache.Put(key, i+100)

			// Access the new items to potentially promote them
			cache.Get(key)
			cache.Get(key)
		}

		// Verify the cache still works
		stats := cache.Stats()
		if stats.Size == 0 {
			t.Errorf("Expected some items in cache")
		}
	})

	t.Run("Force specific eviction scenarios", func(t *testing.T) {
		// Create an even smaller cache to force more aggressive evictions
		cache, err := New[int, string](5)
		if err != nil {
			t.Fatalf("Failed to create cache: %v", err)
		}

		// Fill and overflow to trigger various eviction paths
		for i := 0; i < 30; i++ {
			cache.Put(i, "value"+string(rune('0'+i%10)))

			// Access previous items to create access patterns
			if i > 0 {
				cache.Get(i - 1)
			}
			if i > 1 {
				cache.Get(i - 2)
			}
		}

		// Access items multiple times to force promotion and demotion
		for i := 25; i < 30; i++ {
			for j := 0; j < 5; j++ {
				cache.Get(i)
			}
		}

		// Continue adding items to force more evictions
		for i := 100; i < 120; i++ {
			cache.Put(i, "new"+string(rune('0'+i%10)))
		}

		// Verify cache functionality
		cache.Put(999, "test")
		value, ok := cache.Get(999)
		if !ok || value != "test" {
			t.Errorf("Cache should still function properly")
		}
	})

	t.Run("Complex access pattern for maximum coverage", func(t *testing.T) {
		cache, err := New[string, int](8) // Small cache to force evictions
		if err != nil {
			t.Fatalf("Failed to create cache: %v", err)
		}

		// Phase 1: Fill the cache
		keys := make([]string, 0, 20)
		for i := 0; i < 20; i++ {
			key := "key" + string(rune('0'+i%10)) + string(rune('A'+i%26))
			keys = append(keys, key)
			cache.Put(key, i)
		}

		// Phase 2: Create hot items (frequent access)
		hotKeys := keys[:5]
		for _, key := range hotKeys {
			for j := 0; j < 10; j++ { // Many accesses to promote to protected
				cache.Get(key)
			}
		}

		// Phase 3: Add more items to force protected segment to be full
		for i := 0; i < 15; i++ {
			key := "overflow" + string(rune('A'+i))
			cache.Put(key, i+1000)

			// Access the new item to try to promote it
			for j := 0; j < 3; j++ {
				cache.Get(key)
			}
		}

		// Phase 4: Add even more to force demotions and evictions
		for i := 0; i < 10; i++ {
			key := "final" + string(rune('A'+i))
			cache.Put(key, i+2000)
		}

		// Verify cache still works and has reasonable stats
		stats := cache.Stats()
		t.Logf("Final stats: Size=%d, Capacity=%d", stats.Size, stats.Capacity)

		// The cache should have processed many operations
		if stats.Capacity == 0 {
			t.Errorf("Expected valid cache configuration")
		}
	})
}

// TestMainCacheSegmentInteractions tests interactions between probation and protected segments
func TestMainCacheSegmentInteractions(t *testing.T) {
	t.Run("Probation to protected promotion", func(t *testing.T) {
		cache, err := New[int, string](6) // Small cache
		if err != nil {
			t.Fatalf("Failed to create cache: %v", err)
		}

		// Fill cache and create access patterns
		for i := 0; i < 12; i++ {
			cache.Put(i, "value"+string(rune('0'+i%10)))
		}

		// Create specific access patterns to trigger promotions
		// Access some items multiple times to promote from probation to protected
		for i := 8; i < 12; i++ {
			for j := 0; j < 5; j++ {
				cache.Get(i)
			}
		}

		// Add more items to force the segments to reorganize
		for i := 20; i < 30; i++ {
			cache.Put(i, "new"+string(rune('0'+i%10)))

			// Immediately access to influence admission
			cache.Get(i)
		}

		// Force more segment interactions
		for i := 35; i < 45; i++ {
			cache.Put(i, "latest"+string(rune('0'+i%10)))
		}

		// Verify cache functionality
		stats := cache.Stats()
		if stats.Capacity == 0 {
			t.Errorf("Expected valid cache")
		}
	})
}

// TestEvictionStressTest creates scenarios to maximize eviction coverage
func TestEvictionStressTest(t *testing.T) {
	t.Run("High pressure eviction test", func(t *testing.T) {
		cache, err := New[string, int](4) // Very small cache for maximum pressure
		if err != nil {
			t.Fatalf("Failed to create cache: %v", err)
		}

		// Continuous operations to force all eviction paths
		for round := 0; round < 5; round++ {
			// Add many items in this round
			for i := 0; i < 20; i++ {
				key := "r" + string(rune('0'+round)) + "i" + string(rune('0'+i%10))
				cache.Put(key, round*100+i)

				// Immediately access to affect admission decisions
				cache.Get(key)
			}

			// Access some previous round items if they still exist
			for i := 0; i < 10; i++ {
				if round > 0 {
					prevKey := "r" + string(rune('0'+round-1)) + "i" + string(rune('0'+i%10))
					cache.Get(prevKey) // May hit or miss
				}
			}
		}

		// Final verification
		cache.Put("final", 9999)
		value, ok := cache.Get("final")
		if !ok || value != 9999 {
			t.Errorf("Cache should handle final operation correctly")
		}

		stats := cache.Stats()
		t.Logf("Stress test stats: Size=%d, Capacity=%d", stats.Size, stats.Capacity)
	})
}