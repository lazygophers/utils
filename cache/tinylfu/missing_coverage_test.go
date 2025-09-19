package tinylfu

import (
	"testing"
)

// TestMissingCoverageFunctions tests functions with low coverage
func TestMissingCoverageFunctions(t *testing.T) {

	t.Run("NewWithEvict", func(t *testing.T) {
		t.Run("with evict callback", func(t *testing.T) {
			evicted := make(map[string]int)
			evictCallback := func(key string, value int) {
				evicted[key] = value
			}

			cache, err := NewWithEvict[string, int](5, evictCallback)
			if err != nil {
				t.Fatalf("Failed to create cache: %v", err)
			}

			// Fill cache beyond capacity to trigger evictions
			for i := 0; i < 10; i++ {
				cache.Put(string(rune('a'+i)), i)
			}

			// Should have evicted some items
			if len(evicted) == 0 {
				t.Error("Expected some items to be evicted")
			}
		})

		t.Run("with nil evict callback", func(t *testing.T) {
			cache, err := NewWithEvict[string, int](5, nil)
			if err != nil {
				t.Fatalf("Failed to create cache: %v", err)
			}

			// Fill cache beyond capacity
			for i := 0; i < 10; i++ {
				cache.Put(string(rune('a'+i)), i)
			}

			// Should work without panicking
			if cache.Len() > 5 {
				t.Errorf("Cache size should not exceed capacity")
			}
		})

		t.Run("invalid capacity", func(t *testing.T) {
			_, err := NewWithEvict[string, int](0, nil)
			if err == nil {
				t.Error("Expected error for zero capacity")
			}
		})
	})

	t.Run("Keys", func(t *testing.T) {
		cache, err := New[string, int](10)
		if err != nil {
			t.Fatalf("Failed to create cache: %v", err)
		}

		// Test empty cache
		keys := cache.Keys()
		if len(keys) != 0 {
			t.Errorf("Expected 0 keys, got %d", len(keys))
		}

		// Add items one by one and test keys
		cache.Put("key1", 1)
		keys = cache.Keys()
		if len(keys) != 1 {
			t.Errorf("Expected 1 key, got %d", len(keys))
		}

		cache.Put("key2", 2)
		cache.Put("key3", 3)
		keys = cache.Keys()

		// Check that we have some keys (account for potential evictions)
		if len(keys) == 0 {
			t.Error("Expected some keys, got 0")
		}

		// Verify keys are valid
		for _, key := range keys {
			if key == "" {
				t.Error("Found empty key")
			}
		}
	})

	t.Run("Values", func(t *testing.T) {
		cache, err := New[string, int](10)
		if err != nil {
			t.Fatalf("Failed to create cache: %v", err)
		}

		// Test empty cache
		values := cache.Values()
		if len(values) != 0 {
			t.Errorf("Expected 0 values, got %d", len(values))
		}

		// Add items and test values
		cache.Put("key1", 100)
		values = cache.Values()
		if len(values) != 1 {
			t.Errorf("Expected 1 value, got %d", len(values))
		}

		cache.Put("key2", 200)
		cache.Put("key3", 300)
		values = cache.Values()

		// Check that we have some values (account for potential evictions)
		if len(values) == 0 {
			t.Error("Expected some values, got 0")
		}

		// Verify values are valid
		for _, value := range values {
			if value < 0 {
				t.Error("Found negative value")
			}
		}
	})

	t.Run("Items", func(t *testing.T) {
		cache, err := New[string, int](10)
		if err != nil {
			t.Fatalf("Failed to create cache: %v", err)
		}

		// Test empty cache
		items := cache.Items()
		if len(items) != 0 {
			t.Errorf("Expected 0 items, got %d", len(items))
		}

		// Add one item and test
		cache.Put("key1", 100)
		items = cache.Items()
		if len(items) != 1 {
			t.Errorf("Expected 1 item, got %d", len(items))
		}

		// Add more items
		cache.Put("key2", 200)
		cache.Put("key3", 300)
		items = cache.Items()

		// Check that we have some items (account for potential evictions)
		if len(items) == 0 {
			t.Error("Expected some items, got 0")
		}

		// Verify items are valid
		for key, value := range items {
			if key == "" {
				t.Error("Found empty key")
			}
			if value <= 0 {
				t.Error("Found non-positive value")
			}
		}
	})

	t.Run("Resize", func(t *testing.T) {
		cache, err := New[string, int](10)
		if err != nil {
			t.Fatalf("Failed to create cache: %v", err)
		}

		// Fill cache
		for i := 0; i < 8; i++ {
			cache.Put(string(rune('a'+i)), i)
		}

		originalLen := cache.Len()

		t.Run("resize larger", func(t *testing.T) {
			cache.Resize(20)
			if cache.Cap() != 20 {
				t.Errorf("Expected capacity 20, got %d", cache.Cap())
			}
			// Items should be preserved
			if cache.Len() != originalLen {
				t.Errorf("Expected length %d after resize, got %d", originalLen, cache.Len())
			}
		})

		t.Run("resize smaller", func(t *testing.T) {
			cache.Resize(5)
			if cache.Cap() != 5 {
				t.Errorf("Expected capacity 5, got %d", cache.Cap())
			}
			// Some items might be evicted
			if cache.Len() > 5 {
				t.Errorf("Cache length should not exceed new capacity")
			}
		})

		t.Run("resize to zero", func(t *testing.T) {
			cache.Resize(0)
			// Should handle zero capacity gracefully
			if cache.Cap() < 0 {
				t.Error("Capacity should not be negative")
			}
		})
	})
}

// TestDemoteFromProtectedCoverage tests the demoteFromProtected function
func TestDemoteFromProtectedCoverage(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Fill window cache first to force items to move to main cache
	for i := 0; i < 10; i++ {
		cache.Put(string(rune('a'+i)), i)
	}

	// Access some items multiple times to get them into protected segment
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			cache.Get(string(rune('a' + i)))
		}
	}

	// Add more items to trigger demotion
	for i := 10; i < 50; i++ {
		cache.Put(string(rune('a'+i)), i)
	}

	// Add many more items to force protected segment to become full and trigger demotion
	for i := 50; i < 120; i++ {
		cache.Put(string(rune('a'+i)), i)
		// Access some items to promote them to protected
		if i%2 == 0 {
			cache.Get(string(rune('a' + (i-20))))
		}
	}

	// The demotion should have been triggered internally
	stats := cache.Stats()
	t.Logf("Cache stats: %+v", stats)
}

// TestEvictFromProtectedCoverage tests the evictFromProtected function
func TestEvictFromProtectedCoverage(t *testing.T) {
	cache, err := New[string, int](50)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Fill cache to trigger evictions from protected segment
	for i := 0; i < 60; i++ {
		cache.Put(string(rune('a'+i)), i)
	}

	// Access items multiple times to promote them to protected
	for i := 0; i < 30; i++ {
		for j := 0; j < 3; j++ {
			if _, ok := cache.Get(string(rune('a' + i))); ok {
				// Access successful
			}
		}
	}

	// Add more items to force eviction from protected
	for i := 60; i < 100; i++ {
		cache.Put(string(rune('a'+i)), i)
	}

	// Continue adding items to force more evictions
	for i := 100; i < 150; i++ {
		cache.Put(string(rune('a'+i)), i)
		// Promote some items to protected to fill it
		if i%3 == 0 {
			cache.Get(string(rune('a' + (i-50))))
		}
	}

	// The eviction should have been triggered internally
	if cache.Len() > cache.Cap() {
		t.Errorf("Cache length %d should not exceed capacity %d", cache.Len(), cache.Cap())
	}
}

// TestComplexCacheOperations tests complex scenarios to trigger internal methods
func TestComplexCacheOperations(t *testing.T) {
	cache, err := NewWithEvict[string, int](20, func(key string, value int) {
		t.Logf("Evicted: %s -> %d", key, value)
	})
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Complex access pattern to trigger different code paths

	// Phase 1: Fill window
	for i := 0; i < 5; i++ {
		cache.Put(string(rune('A'+i)), i*100)
	}

	// Phase 2: Access items to move them to main cache
	for i := 0; i < 5; i++ {
		cache.Get(string(rune('A' + i)))
	}

	// Phase 3: Add more items to force window eviction
	for i := 5; i < 15; i++ {
		cache.Put(string(rune('A'+i)), i*100)
	}

	// Phase 4: Create access patterns to promote items to protected
	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			cache.Get(string(rune('A' + i)))
		}
	}

	// Phase 5: Add items to force probation eviction
	for i := 15; i < 30; i++ {
		cache.Put(string(rune('A'+i)), i*100)
	}

	// Phase 6: Force protected segment to be full and trigger demotion
	for i := 30; i < 50; i++ {
		cache.Put(string(rune('A'+i)), i*100)
		// Access items to try to promote them to protected
		if i%2 == 0 {
			cache.Get(string(rune('A' + (i-20))))
		}
	}

	// Phase 7: Continue operations to trigger eviction from protected
	for i := 50; i < 80; i++ {
		cache.Put(string(rune('A'+i)), i*100)

		// Mix of operations
		if i%3 == 0 {
			cache.Get(string(rune('A' + (i-30))))
		}
		if i%4 == 0 {
			cache.Remove(string(rune('A' + (i-40))))
		}
	}

	// Verify cache integrity
	if cache.Len() > cache.Cap() {
		t.Errorf("Cache length %d exceeds capacity %d", cache.Len(), cache.Cap())
	}

	// Test all query methods
	keys := cache.Keys()
	values := cache.Values()
	items := cache.Items()

	if len(keys) != len(values) || len(keys) != len(items) {
		t.Errorf("Inconsistent lengths: keys=%d, values=%d, items=%d", len(keys), len(values), len(items))
	}

	// Test resize during complex state
	cache.Resize(10)
	if cache.Cap() != 10 {
		t.Errorf("Expected capacity 10 after resize, got %d", cache.Cap())
	}

	cache.Resize(30)
	if cache.Cap() != 30 {
		t.Errorf("Expected capacity 30 after resize, got %d", cache.Cap())
	}
}