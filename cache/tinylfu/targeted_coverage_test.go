package tinylfu

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestTargetedCoverage specifically targets the uncovered functions
func TestTargetedCoverage(t *testing.T) {
	t.Run("trigger_demoteFromProtected", func(t *testing.T) {
		// Create a cache where we can precisely control the segments
		cache, err := New[string, string](20) // Small cache to make calculations easier
		assert.NoError(t, err)

		// According to the code:
		// - windowSize is typically ~10% of capacity = 2
		// - mainSize is the remaining = 18
		// - protectedCapacity is 80% of mainSize = ~14
		// - probationCapacity is 20% of mainSize = ~4

		// Step 1: Fill the cache to get items into different segments
		// Fill beyond window size to get items into main cache
		for i := 0; i < 30; i++ {
			cache.Put("key"+strconv.Itoa(i), "value"+strconv.Itoa(i))
		}

		// Step 2: Access specific items many times to get them promoted to protected
		// This should fill the protected segment
		hotKeys := make([]string, 15) // More than protected capacity
		for i := 0; i < 15; i++ {
			hotKeys[i] = "hot" + strconv.Itoa(i)
			cache.Put(hotKeys[i], "hot_value"+strconv.Itoa(i))
		}

		// Access these hot keys multiple times to promote them to protected
		for _, key := range hotKeys {
			for j := 0; j < 10; j++ { // Multiple accesses
				_, _ = cache.Get(key)
			}
		}

		// Step 3: Now add more hot items to force demotion from protected
		// This should trigger demoteFromProtected when protected becomes full
		for i := 15; i < 25; i++ {
			newHotKey := "superhot" + strconv.Itoa(i)
			cache.Put(newHotKey, "superhot_value"+strconv.Itoa(i))

			// Access this new key many times to promote it to protected
			for j := 0; j < 15; j++ {
				_, _ = cache.Get(newHotKey)
			}
		}

		// Verify cache is still functional
		assert.True(t, cache.Len() <= 20)
		assert.Greater(t, cache.Len(), 0)
	})

	t.Run("trigger_evictFromProtected_via_resize", func(t *testing.T) {
		// Create a larger cache and fill it
		cache, err := New[int, string](100)
		assert.NoError(t, err)

		// Fill the cache completely
		for i := 0; i < 120; i++ {
			cache.Put(i, "value"+strconv.Itoa(i))
		}

		// Promote many items to protected by accessing them frequently
		for i := 0; i < 50; i++ {
			for j := 0; j < 10; j++ {
				_, _ = cache.Get(i)
			}
		}

		// Now resize to a very small capacity to force eviction from protected
		// The resize logic will try to evict from window, then probation, then protected
		err = cache.Resize(10) // Much smaller than current size
		assert.NoError(t, err)

		// Verify the cache respected the new size
		assert.True(t, cache.Len() <= 10)
		assert.Equal(t, 10, cache.Cap())
	})

	t.Run("force_evictFromProtected_empty_other_segments", func(t *testing.T) {
		// Create a scenario where window and probation are empty but protected has items
		cache, err := New[string, string](20)
		assert.NoError(t, err)

		// Add items and promote them all to protected
		keys := make([]string, 15)
		for i := 0; i < 15; i++ {
			keys[i] = "protected" + strconv.Itoa(i)
			cache.Put(keys[i], "value"+strconv.Itoa(i))

			// Access multiple times to promote to protected
			for j := 0; j < 20; j++ {
				_, _ = cache.Get(keys[i])
			}
		}

		// Now resize to force eviction when other segments are empty
		err = cache.Resize(5)
		assert.NoError(t, err)

		assert.True(t, cache.Len() <= 5)
		assert.Equal(t, 5, cache.Cap())
	})

	t.Run("comprehensive_segment_transitions", func(t *testing.T) {
		// Test all possible transitions between segments
		cache, err := New[string, string](50)
		assert.NoError(t, err)

		// Phase 1: Fill window
		for i := 0; i < 10; i++ {
			cache.Put("window"+strconv.Itoa(i), "value"+strconv.Itoa(i))
		}

		// Phase 2: Overflow to probation
		for i := 10; i < 30; i++ {
			cache.Put("probation"+strconv.Itoa(i), "value"+strconv.Itoa(i))
		}

		// Phase 3: Promote some items to protected
		for i := 10; i < 25; i++ {
			for j := 0; j < 8; j++ {
				_, _ = cache.Get("probation" + strconv.Itoa(i))
			}
		}

		// Phase 4: Add more items to trigger all eviction paths
		for i := 30; i < 80; i++ {
			cache.Put("overflow"+strconv.Itoa(i), "value"+strconv.Itoa(i))

			// Occasionally access to create promotion pressure
			if i%3 == 0 {
				for j := 0; j < 5; j++ {
					_, _ = cache.Get("overflow" + strconv.Itoa(i))
				}
			}
		}

		// Phase 5: Resize to trigger protected eviction
		err = cache.Resize(15)
		assert.NoError(t, err)

		assert.True(t, cache.Len() <= 15)
		assert.Equal(t, 15, cache.Cap())

		// Verify cache still works
		cache.Put("final", "test")
		value, exists := cache.Get("final")
		assert.True(t, exists)
		assert.Equal(t, "test", value)
	})
}