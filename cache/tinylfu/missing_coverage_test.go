package tinylfu

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMissingCoveragePaths targets the uncovered functions to improve overall coverage
func TestMissingCoveragePaths(t *testing.T) {
	t.Run("demoteFromProtected coverage", func(t *testing.T) {
		// Create a cache with small capacity to trigger segment transitions
		cache, err := New[string, int](10) // Total capacity of 10
		require.NoError(t, err)

		// Fill the window first (windowSize = 10% of total = 1)
		// Then fill probation + protected to force demotion scenarios

		// Step 1: Fill window
		cache.Put("w1", 1)

		// Step 2: Add more items to push items through the segments
		// These will go to window and potentially move to main cache
		for i := 2; i <= 15; i++ {
			cache.Put("key"+string(rune(i)), i*100)
		}

		// Step 3: Access items multiple times to promote them to protected
		// which will eventually trigger demoteFromProtected when protected is full
		for i := 0; i < 10; i++ {
			for j := 10; j <= 15; j++ {
				cache.Get("key" + string(rune(j)))
			}
		}

		// Step 4: Add more items to force evictions and potentially trigger demotion
		for i := 20; i <= 30; i++ {
			cache.Put("newkey"+string(rune(i)), i*100)
			// Access them to promote to protected and force demotion of existing items
			cache.Get("newkey" + string(rune(i)))
		}

		// The above operations should have triggered demoteFromProtected
		// at some point due to segment capacity constraints
		assert.True(t, cache.Len() <= 10, "Cache should respect capacity limit")
	})

	t.Run("evictFromProtected coverage", func(t *testing.T) {
		// Create a cache and force it into a state where evictFromProtected is called
		cache, err := New[string, int](5) // Small capacity for easier testing
		require.NoError(t, err)

		// Fill up the cache completely
		for i := 1; i <= 10; i++ {
			cache.Put("item"+string(rune(i+48)), i) // Convert to char
		}

		// Access some items multiple times to promote them to protected
		for i := 0; i < 5; i++ {
			for j := 1; j <= 5; j++ {
				cache.Get("item" + string(rune(j+48)))
			}
		}

		// Now resize to a smaller capacity to force eviction from protected
		err = cache.Resize(2)
		require.NoError(t, err)

		// This resize should have triggered evictFromProtected
		assert.Equal(t, 2, cache.Cap(), "Cache capacity should be resized to 2")
		assert.True(t, cache.Len() <= 2, "Cache length should not exceed new capacity")
	})

	t.Run("Keys method coverage", func(t *testing.T) {
		cache, err := New[string, int](10)
		require.NoError(t, err)

		// Test Keys() with empty cache
		keys := cache.Keys()
		assert.Empty(t, keys, "Keys should be empty for empty cache")

		// Add items to different segments and test Keys()
		testData := map[string]int{
			"key1": 100,
			"key2": 200,
			"key3": 300,
			"key4": 400,
			"key5": 500,
		}

		for k, v := range testData {
			cache.Put(k, v)
		}

		// Access some items to promote them through segments
		cache.Get("key1")
		cache.Get("key2")
		cache.Get("key1") // Multiple accesses to promote to protected

		keys = cache.Keys()
		assert.Greater(t, len(keys), 0, "Keys should not be empty")

		// Verify all keys are present
		keySet := make(map[string]bool)
		for _, key := range keys {
			keySet[key] = true
		}

		for key := range testData {
			if cache.Contains(key) {
				assert.True(t, keySet[key], "Key %s should be in keys list", key)
			}
		}
	})

	t.Run("Values method coverage", func(t *testing.T) {
		cache, err := New[string, int](10)
		require.NoError(t, err)

		// Test Values() with empty cache
		values := cache.Values()
		assert.Empty(t, values, "Values should be empty for empty cache")

		// Add items to different segments
		testData := map[string]int{
			"a": 100,
			"b": 200,
			"c": 300,
		}

		for k, v := range testData {
			cache.Put(k, v)
		}

		values = cache.Values()
		assert.Greater(t, len(values), 0, "Values should not be empty")

		// Verify values are reasonable (exact values may vary due to eviction)
		valueSet := make(map[int]bool)
		for _, value := range values {
			valueSet[value] = true
		}

		// At least some of our test values should be present
		foundValues := 0
		for _, expectedValue := range testData {
			if valueSet[expectedValue] {
				foundValues++
			}
		}
		assert.Greater(t, foundValues, 0, "Should find at least some expected values")
	})

	t.Run("Items method coverage", func(t *testing.T) {
		cache, err := New[string, int](10)
		require.NoError(t, err)

		// Test Items() with empty cache
		items := cache.Items()
		assert.Empty(t, items, "Items should be empty for empty cache")

		// Add items across different segments
		testData := map[string]int{
			"x": 10,
			"y": 20,
			"z": 30,
		}

		for k, v := range testData {
			cache.Put(k, v)
		}

		// Access to distribute across segments
		cache.Get("x")
		cache.Get("y")

		items = cache.Items()
		assert.Greater(t, len(items), 0, "Items should not be empty")

		// Verify the structure of returned items
		for key, value := range items {
			assert.NotEmpty(t, key, "Item key should not be empty")
			// Value can be any int, so just verify it's accessible
			_ = value
		}
	})

	t.Run("Resize method coverage", func(t *testing.T) {
		cache, err := New[string, int](10)
		require.NoError(t, err)

		// Fill cache with items
		for i := 1; i <= 8; i++ {
			cache.Put("resize"+string(rune(i+48)), i*10)
		}

		originalLen := cache.Len()
		assert.Greater(t, originalLen, 0, "Cache should have items before resize")

		// Test resize to larger capacity
		err = cache.Resize(15)
		require.NoError(t, err)
		assert.Equal(t, 15, cache.Cap(), "Cache capacity should be increased")

		// Test resize to smaller capacity (should trigger evictions)
		err = cache.Resize(3)
		require.NoError(t, err)
		assert.Equal(t, 3, cache.Cap(), "Cache capacity should be decreased")
		assert.True(t, cache.Len() <= 3, "Cache length should respect new capacity")

		// Test resize to zero (error case)
		err = cache.Resize(0)
		assert.Error(t, err, "Resize to zero should return an error")

		// Test resize to negative (error case)
		err = cache.Resize(-1)
		assert.Error(t, err, "Resize to negative should return an error")
	})

	t.Run("Complex segment interactions", func(t *testing.T) {
		// This test aims to exercise multiple segment transitions to improve coverage
		cache, err := New[string, int](8)
		require.NoError(t, err)

		// Phase 1: Fill window and push to main cache
		phase1Keys := []string{"p1a", "p1b", "p1c", "p1d"}
		for i, key := range phase1Keys {
			cache.Put(key, (i+1)*100)
		}

		// Phase 2: Access pattern to promote items through segments
		// Multiple accesses should move items from probation -> protected
		for i := 0; i < 3; i++ {
			for _, key := range phase1Keys {
				cache.Get(key)
			}
		}

		// Phase 3: Add more items to force evictions
		phase2Keys := []string{"p2a", "p2b", "p2c", "p2d", "p2e"}
		for i, key := range phase2Keys {
			cache.Put(key, (i+10)*100)
			// Immediate access to try to promote
			cache.Get(key)
		}

		// Phase 4: Force protected segment to be full and trigger demotion
		for i := 0; i < 5; i++ {
			for _, key := range phase2Keys {
				cache.Get(key)
			}
		}

		// Phase 5: Add even more items to force various eviction paths
		for i := 0; i < 10; i++ {
			key := "final" + string(rune(i+65)) // A, B, C, etc.
			cache.Put(key, i*1000)
			cache.Get(key) // Try to promote immediately
		}

		// Verify cache is still within capacity and functioning
		assert.True(t, cache.Len() <= 8, "Cache should respect capacity")
		assert.Greater(t, cache.Len(), 0, "Cache should not be empty")

		// Test that we can still perform basic operations
		cache.Put("final_test", 9999)
		value, exists := cache.Get("final_test")
		if exists {
			assert.Equal(t, 9999, value, "Should be able to retrieve recently added item")
		}
	})
}

// TestProtectedSegmentSpecific specifically targets protected segment operations
func TestProtectedSegmentSpecific(t *testing.T) {
	t.Run("Force demoteFromProtected execution", func(t *testing.T) {
		// Create a cache with specific size ratios to force demotion
		cache, err := New[int, string](10)
		require.NoError(t, err)

		// Fill the cache with items and promote many to protected
		for i := 1; i <= 15; i++ {
			cache.Put(i, "value"+string(rune(i+48)))
		}

		// Access items multiple times to promote to protected
		// The protected segment has limited capacity, so this should force demotion
		for round := 0; round < 5; round++ {
			for i := 10; i <= 15; i++ {
				cache.Get(i)
			}
		}

		// Add more items to force segment pressure
		for i := 20; i <= 35; i++ {
			cache.Put(i, "new"+string(rune(i+48)))
			// Multiple accesses to force promotion and create pressure
			for j := 0; j < 3; j++ {
				cache.Get(i)
			}
		}

		// At this point, demoteFromProtected should have been called
		// due to protected segment capacity constraints
		assert.True(t, cache.Len() <= 10, "Cache should maintain capacity constraints")
	})

	t.Run("Force evictFromProtected via capacity reduction", func(t *testing.T) {
		cache, err := New[string, int](20)
		require.NoError(t, err)

		// Fill cache and promote items to protected
		keys := make([]string, 0, 15)
		for i := 0; i < 15; i++ {
			key := "protected" + string(rune(i+65))
			keys = append(keys, key)
			cache.Put(key, i*100)
		}

		// Heavily access these keys to promote them to protected segment
		for round := 0; round < 10; round++ {
			for _, key := range keys {
				cache.Get(key)
			}
		}

		// Now drastically reduce capacity to force eviction from protected
		err = cache.Resize(3)
		require.NoError(t, err)

		// This should have forced evictFromProtected to be called
		assert.Equal(t, 3, cache.Cap(), "Capacity should be reduced")
		assert.True(t, cache.Len() <= 3, "Length should not exceed new capacity")

		// Verify cache still works after protected eviction
		cache.Put("test_after_eviction", 999)
		value, exists := cache.Get("test_after_eviction")
		if exists {
			assert.Equal(t, 999, value, "Cache should still work after evictions")
		}
	})
}