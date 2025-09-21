package tinylfu

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestZeroCoverageTargeted focuses specifically on hitting the 0% coverage functions
func TestZeroCoverageTargeted(t *testing.T) {
	t.Run("demoteFromProtected direct trigger", func(t *testing.T) {
		// Create a cache with specific capacity to trigger demote logic
		cache, err := New[string, int](10) // capacity=10, window=1, main=9, protected=~7
		require.NoError(t, err)

		// Step 1: Fill window to capacity
		cache.Put("window1", 1)

		// Step 2: Add item to probation (should move window1 to main->probation)
		cache.Put("prob1", 2) // This should move window1 to probation

		// Step 3: Access prob1 to get it into protected, but first need to get it into probation
		// Let's add more items to push through the system
		cache.Put("item1", 10)
		cache.Put("item2", 20)
		cache.Put("item3", 30)

		// Access items to promote them to protected
		// We need to carefully promote enough items to protected to trigger demotion
		cache.Get("item1")
		cache.Get("item1") // Access multiple times to ensure promotion
		cache.Get("item2")
		cache.Get("item2")

		// Now add more items and promote them to protected until we hit the capacity limit
		mainSize := 9 // capacity - windowSize
		protectedCapacity := int(float64(mainSize) * 0.8) // 9 * 0.8 = 7.2 -> 7

		// Fill protected to capacity by promoting items
		for i := 0; i < protectedCapacity + 2; i++ {
			key := "protected" + string(rune('A'+i))
			cache.Put(key, i*100)

			// Access multiple times to ensure promotion to protected
			for j := 0; j < 5; j++ {
				cache.Get(key)
			}
		}

		// At this point, protected should be at capacity
		// Adding one more item and promoting it should trigger demoteFromProtected
		cache.Put("trigger_demote", 999)
		for i := 0; i < 5; i++ {
			cache.Get("trigger_demote") // This should trigger promoteToProtected -> demoteFromProtected
		}

		// Verify cache still works
		assert.True(t, cache.Len() <= 10, "Cache should respect capacity")
	})

	t.Run("evictFromProtected direct trigger", func(t *testing.T) {
		// Create a larger cache and fill protected segment
		cache, err := New[string, int](20) // capacity=20, window=1, main=19, protected=~15
		require.NoError(t, err)

		// Fill the cache with items that will end up in protected
		keys := make([]string, 0, 15)
		for i := 0; i < 15; i++ {
			key := "prot" + string(rune('A'+i))
			keys = append(keys, key)
			cache.Put(key, i*10)

			// Access heavily to ensure they get into protected
			for j := 0; j < 10; j++ {
				cache.Get(key)
			}
		}

		// At this point, protected should have items
		// Now resize to a very small capacity to force evictFromProtected
		// The resize logic will evict from window first, then probation, then protected
		err = cache.Resize(2) // Very small capacity to force protected eviction
		require.NoError(t, err)

		// This should have triggered evictFromProtected
		assert.Equal(t, 2, cache.Cap(), "Capacity should be 2")
		assert.True(t, cache.Len() <= 2, "Length should not exceed capacity")

		// Verify cache still works after eviction
		cache.Put("final", 123)
		value, exists := cache.Get("final")
		if exists {
			assert.Equal(t, 123, value)
		}
	})

	t.Run("demoteFromProtected via precise probation promotion", func(t *testing.T) {
		// Very targeted approach to trigger demotion
		cache, err := New[int, string](5) // Small cache: capacity=5, window=1, main=4, protected=~3
		require.NoError(t, err)

		// Fill window
		cache.Put(1, "window")

		// Add item to get window item into probation
		cache.Put(2, "item2")

		// Promote item2 to protected
		cache.Get(2)
		cache.Get(2)

		// Add and promote more items to fill protected
		cache.Put(3, "item3")
		cache.Get(3)
		cache.Get(3)

		cache.Put(4, "item4")
		cache.Get(4)
		cache.Get(4)

		// At this point protected should be near capacity (3 items for capacity 5)
		// Add one more and promote it to trigger demotion
		cache.Put(5, "item5")

		// This promotion should trigger demoteFromProtected
		for i := 0; i < 3; i++ {
			cache.Get(5)
		}

		assert.True(t, cache.Len() <= 5, "Cache should respect capacity")
	})

	t.Run("evictFromProtected with empty window and probation", func(t *testing.T) {
		// Create specific conditions where only protected has items during resize
		cache, err := New[string, string](10)
		require.NoError(t, err)

		// Add items and promote them all to protected
		for i := 0; i < 8; i++ {
			key := "item" + string(rune('0'+i))
			cache.Put(key, key+"_value")

			// Heavy access to ensure promotion to protected
			for j := 0; j < 10; j++ {
				cache.Get(key)
			}
		}

		// At this point, most/all items should be in protected
		// Now resize to force eviction specifically from protected
		err = cache.Resize(3)
		require.NoError(t, err)

		assert.Equal(t, 3, cache.Cap())
		assert.True(t, cache.Len() <= 3)
	})
}