package tinylfu

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestDebugCoverage helps understand why functions aren't being called
func TestDebugCoverage(t *testing.T) {
	t.Run("debug_cache_structure", func(t *testing.T) {
		cache, err := New[string, string](10)
		assert.NoError(t, err)

		// Let's understand the cache structure first
		fmt.Printf("Initial state - Capacity: %d, Len: %d\n", cache.Cap(), cache.Len())

		// Add some items and see what happens
		for i := 0; i < 15; i++ {
			cache.Put("key"+strconv.Itoa(i), "value"+strconv.Itoa(i))
			fmt.Printf("After adding key%d - Len: %d\n", i, cache.Len())
		}

		// Try to promote items by accessing them
		fmt.Println("Starting to access items to promote them...")
		for i := 0; i < 10; i++ {
			for j := 0; j < 20; j++ {
				_, exists := cache.Get("key" + strconv.Itoa(i))
				if !exists {
					fmt.Printf("Key%d not found\n", i)
					break
				}
			}
		}

		fmt.Printf("After promotions - Len: %d\n", cache.Len())

		// Try to trigger resize-based eviction
		fmt.Println("Attempting resize to trigger evictFromProtected...")
		err = cache.Resize(3)
		assert.NoError(t, err)
		fmt.Printf("After resize to 3 - Len: %d, Cap: %d\n", cache.Len(), cache.Cap())
	})

	t.Run("manual_method_calls", func(t *testing.T) {
		// Since the functions might not be called due to complex logic,
		// let's test if we can call them manually through reflection
		// or by creating specific conditions

		cache, err := New[string, string](20)
		assert.NoError(t, err)

		// Fill the cache with items in protected segment
		for i := 0; i < 25; i++ {
			key := "protected" + strconv.Itoa(i)
			cache.Put(key, "value"+strconv.Itoa(i))

			// Access many times to get into protected
			for j := 0; j < 30; j++ {
				_, _ = cache.Get(key)
			}
		}

		fmt.Printf("Before resize - Len: %d\n", cache.Len())

		// Force a very aggressive resize
		err = cache.Resize(1)
		assert.NoError(t, err)

		fmt.Printf("After aggressive resize - Len: %d\n", cache.Len())
		assert.True(t, cache.Len() <= 1)
	})

	t.Run("test_promotion_logic", func(t *testing.T) {
		// Test the specific promotion logic that should trigger demoteFromProtected
		cache, err := New[string, string](100)
		assert.NoError(t, err)

		// Fill cache with many items
		for i := 0; i < 150; i++ {
			cache.Put("item"+strconv.Itoa(i), "value"+strconv.Itoa(i))
		}

		// Create hot items that should get promoted to protected
		hotItems := []string{"hot1", "hot2", "hot3", "hot4", "hot5"}
		for _, item := range hotItems {
			cache.Put(item, "hot_value")
			// Access very frequently to ensure promotion
			for j := 0; j < 100; j++ {
				_, _ = cache.Get(item)
			}
		}

		// Add more hot items to force protected segment overflow
		for i := 0; i < 50; i++ {
			newHot := "superhot" + strconv.Itoa(i)
			cache.Put(newHot, "superhot_value")
			// Access very frequently
			for j := 0; j < 150; j++ {
				_, _ = cache.Get(newHot)
			}
		}

		// Verify cache still works
		assert.True(t, cache.Len() <= 100)
	})
}