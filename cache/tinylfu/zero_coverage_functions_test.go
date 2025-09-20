package tinylfu

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestZeroCoverageFunctions tests the two functions with 0% coverage
func TestZeroCoverageFunctions(t *testing.T) {
	t.Run("demoteFromProtected", func(t *testing.T) {
		// Create a cache that will trigger protected -> probation demotion
		cache, err := New[string, string](100)
		assert.NoError(t, err)

		// Fill the protected segment
		// First we need to get items into protected by promoting from probation

		// Step 1: Fill window to capacity to trigger eviction to probation
		for i := 0; i < 25; i++ { // More than window size to trigger eviction
			cache.Put(string(rune('A'+i)), string(rune('A'+i)))
		}

		// Step 2: Access some items to promote them to protected
		// Access items that made it to probation multiple times to promote to protected
		for i := 0; i < 10; i++ {
			for j := 0; j < 5; j++ { // Multiple accesses to increase frequency
				_, _ = cache.Get(string(rune('A' + i)))
			}
		}

		// Step 3: Fill up more to force protected to be full and trigger demotion
		initialSize := cache.Len()

		// Add many more items to force the cache to demote from protected
		for i := 0; i < 200; i++ {
			cache.Put(string(rune('Z'-i%26))+strconv.Itoa(i), "value"+string(rune('0'+i%10)))
		}

		// The cache should have triggered demoteFromProtected during this process
		// Verify the cache is still functional
		assert.True(t, cache.Len() <= 100, "Cache should not exceed capacity")
		assert.Greater(t, cache.Len(), 0, "Cache should not be empty")

		// Verify we can still get and set
		cache.Put("test-demote", "test-value")
		value, exists := cache.Get("test-demote")
		assert.True(t, exists)
		assert.Equal(t, "test-value", value)

		_ = initialSize // Use the variable to avoid unused warning
	})

	t.Run("evictFromProtected", func(t *testing.T) {
		// Create a small cache to easily trigger eviction from protected
		cache, err := New[int, string](10)
		assert.NoError(t, err)

		// Fill the cache beyond capacity to trigger all eviction paths
		for i := 0; i < 20; i++ {
			cache.Put(i, "value"+strconv.Itoa(i%10))

			// Access recent items multiple times to get them into protected
			if i > 0 {
				for j := 0; j < 3; j++ {
					_, _ = cache.Get(i - 1)
				}
			}
		}

		// At this point, the cache should have triggered evictFromProtected
		// Verify the cache is functioning correctly
		assert.True(t, cache.Len() <= 10, "Cache should not exceed total capacity")
		assert.Greater(t, cache.Len(), 0, "Cache should not be empty")

		// Test that we can still perform operations
		cache.Put(999, "test-evict")
		value, exists := cache.Get(999)
		assert.True(t, exists)
		assert.Equal(t, "test-evict", value)

		// Clear and verify
		cache.Clear()
		assert.Equal(t, 0, cache.Len())
	})

	t.Run("stress_test_all_eviction_paths", func(t *testing.T) {
		// Comprehensive test to ensure both functions are called under various scenarios
		cache, err := New[string, int](50)
		assert.NoError(t, err)

		// Scenario 1: Normal fill and access pattern
		keys := make([]string, 100)
		for i := 0; i < 100; i++ {
			keys[i] = "key" + string(rune('0'+i%10)) + string(rune('A'+i%26))
			cache.Put(keys[i], i)
		}

		// Scenario 2: Heavy access pattern to create hot items
		for i := 0; i < 30; i++ {
			for j := 0; j < 10; j++ {
				if i < len(keys) {
					_, _ = cache.Get(keys[i])
				}
			}
		}

		// Scenario 3: Mixed operations to trigger different eviction scenarios
		for i := 100; i < 200; i++ {
			key := "stress" + strconv.Itoa(i%10)
			cache.Put(key, i)

			// Occasionally access old keys to maintain some in protected
			if i%5 == 0 && i-50 >= 0 && i-50 < len(keys) {
				_, _ = cache.Get(keys[i-50])
			}
		}

		// Verify cache integrity
		assert.True(t, cache.Len() <= 50, "Cache should respect capacity")
		assert.Greater(t, cache.Len(), 0, "Cache should not be empty")

		// Test final operations
		cache.Put("final-test", 999)
		value, exists := cache.Get("final-test")
		assert.True(t, exists)
		assert.Equal(t, 999, value)
	})

	t.Run("edge_case_empty_segments", func(t *testing.T) {
		// Test behavior when trying to demote/evict from empty segments
		cache, err := New[string, string](5)
		assert.NoError(t, err)

		// Test with empty cache
		cache.Put("test1", "value1")
		assert.Equal(t, 1, cache.Len())

		// Fill and overflow to trigger evictions
		for i := 0; i < 10; i++ {
			cache.Put("key"+strconv.Itoa(i), "value"+strconv.Itoa(i))
		}

		// Verify cache is still functional
		assert.True(t, cache.Len() <= 5)
		assert.Greater(t, cache.Len(), 0)
	})
}