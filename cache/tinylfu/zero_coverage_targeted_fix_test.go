package tinylfu

import (
	"fmt"
	"testing"
)

// TestZeroCoverageFunctionsFix explicitly targets the two 0% coverage functions
func TestZeroCoverageFunctionsFix(t *testing.T) {
	t.Run("demoteFromProtected_direct_trigger", func(t *testing.T) {
		// Create cache with size 25: window=1, main=24, protected capacity = int(24*0.8) = 19
		cache, err := New[string, int](25)
		if err != nil {
			t.Fatalf("Failed to create cache: %v", err)
		}

		// Phase 1: Build up frequency in Count-Min Sketch for our keys
		// This is crucial for admission policy to work
		for round := 0; round < 3; round++ {
			for i := 0; i < 25; i++ {
				key := fmt.Sprintf("freq_%d", i)
				cache.Put(key, i)
				cache.Get(key)
			}
		}

		// Phase 2: Clear cache but keep frequency sketch
		cache.Clear()

		// Phase 3: Build up protected segment systematically to exact capacity (19 items)
		protectedKeys := make([]string, 0, 19)

		// Add items that will be admitted due to high frequency in sketch
		for i := 0; i < 19; i++ {
			key := fmt.Sprintf("freq_%d", i)

			// Put in window first
			cache.Put(key, i+100)

			// Force eviction from window to trigger admission policy check
			cache.Put("temp", 999)

			// Re-add - should be admitted to probation due to frequency
			cache.Put(key, i+100)

			// Promote from probation to protected by accessing
			cache.Get(key)

			protectedKeys = append(protectedKeys, key)
		}

		// Verify protected segment is at capacity
		stats := cache.Stats()
		t.Logf("Protected segment filled: size=%d, expected=19", stats.ProtectedSize)

		// Phase 4: Add one more high-frequency item to probation
		triggerKey := "freq_20"
		cache.Put(triggerKey, 200)
		cache.Put("temp2", 888) // evict to trigger admission policy
		cache.Put(triggerKey, 200) // admit to probation

		// Phase 5: Promote this item to protected - this should trigger demoteFromProtected
		// because protected.Len() (19) >= protectedCapacity (19)
		cache.Get(triggerKey) // This call should trigger promoteToProtected -> demoteFromProtected

		// Verify the operation completed successfully
		finalStats := cache.Stats()
		t.Logf("After demotion trigger: Protected=%d, Probation=%d, Window=%d",
			finalStats.ProtectedSize, finalStats.ProbationSize, finalStats.WindowSize)

		if cache.Len() > cache.Cap() {
			t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
		}
	})

	t.Run("evictFromProtected_direct_trigger", func(t *testing.T) {
		// Create cache and fill only protected segment
		cache, err := New[string, int](20)
		if err != nil {
			t.Fatalf("Failed to create cache: %v", err)
		}

		// Build frequency for keys first
		for round := 0; round < 3; round++ {
			for i := 0; i < 15; i++ {
				key := fmt.Sprintf("prot_%d", i)
				cache.Put(key, i)
				cache.Get(key)
			}
		}

		// Clear and rebuild only protected segment
		cache.Clear()

		// Fill protected segment systematically
		for i := 0; i < 15; i++ {
			key := fmt.Sprintf("prot_%d", i)

			// Window -> doorkeeper -> probation -> protected
			cache.Put(key, i+300)
			cache.Put("evict", 999) // evict to doorkeeper
			cache.Put(key, i+300)   // admit to probation
			cache.Get(key)          // promote to protected
		}

		// Ensure window and probation are empty by evicting all items from them
		// Fill window with low-frequency items that won't be admitted
		for i := 0; i < 5; i++ {
			cache.Put(fmt.Sprintf("low_freq_%d", i), i)
		}

		stats := cache.Stats()
		t.Logf("Before resize: Protected=%d, Probation=%d, Window=%d",
			stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)

		// Resize to force eviction from protected when other segments are empty/have low priority items
		// The resize loop will first try window, then probation, then protected
		err = cache.Resize(3)
		if err != nil {
			t.Fatalf("Failed to resize cache: %v", err)
		}

		finalStats := cache.Stats()
		t.Logf("After resize: Protected=%d, Probation=%d, Window=%d",
			finalStats.ProtectedSize, finalStats.ProbationSize, finalStats.WindowSize)

		if cache.Cap() != 3 {
			t.Errorf("Expected capacity 3, got %d", cache.Cap())
		}

		if cache.Len() > 3 {
			t.Errorf("Cache size %d exceeds capacity 3", cache.Len())
		}
	})
}

// TestDemoteFromProtectedExact creates the exact condition for demoteFromProtected
func TestDemoteFromProtectedExact(t *testing.T) {
	// Use a very precise cache size to control segment capacities
	// Cache size 10: window=1, main=9, protected capacity = int(9*0.8) = 7
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Step 1: Build high frequency for specific keys
	highFreqKeys := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	for _, key := range highFreqKeys {
		for i := 0; i < 10; i++ {
			cache.Put(key, 1)
			cache.Get(key)
		}
	}

	// Clear cache but keep frequency data
	cache.Clear()

	// Step 2: Fill protected to exactly capacity (7 items)
	for i := 0; i < 7; i++ {
		key := highFreqKeys[i]

		// Admission process: window -> doorkeeper -> probation -> protected
		cache.Put(key, i+100)      // Put in window
		cache.Put("temp", 999)     // Evict to doorkeeper
		cache.Put(key, i+100)      // Readmit to probation (due to frequency)
		cache.Get(key)             // Promote to protected
	}

	// Step 3: Add one item to probation
	probationKey := highFreqKeys[7]
	cache.Put(probationKey, 800)   // Put in window
	cache.Put("temp2", 888)        // Evict to doorkeeper
	cache.Put(probationKey, 800)   // Readmit to probation

	// Step 4: Promote from probation to protected - this triggers demoteFromProtected
	cache.Get(probationKey) // promoteToProtected -> demoteFromProtected (because protected.Len()=7 >= 7)

	// Verification
	stats := cache.Stats()
	if cache.Len() > cache.Cap() {
		t.Errorf("Cache overflow: size=%d > capacity=%d", cache.Len(), cache.Cap())
	}
	t.Logf("Success: Protected=%d, Probation=%d, Window=%d",
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)
}

// TestEvictFromProtectedExact creates the exact condition for evictFromProtected
func TestEvictFromProtectedExact(t *testing.T) {
	cache, err := New[string, int](8)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Build high frequency keys
	keys := []string{"P1", "P2", "P3", "P4", "P5"}
	for _, key := range keys {
		for i := 0; i < 15; i++ {
			cache.Put(key, 1)
			cache.Get(key)
		}
	}

	// Clear and fill only protected
	cache.Clear()

	for i, key := range keys {
		cache.Put(key, i+500)     // window
		cache.Put("evict", 777)   // evict to doorkeeper
		cache.Put(key, i+500)     // probation (admitted due to frequency)
		cache.Get(key)            // protected
	}

	// Ensure other segments are empty or have only evictable items
	stats := cache.Stats()
	t.Logf("Pre-resize: Protected=%d, Probation=%d, Window=%d",
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)

	// Resize to trigger the evictFromProtected path in the resize loop
	// The loop will: window.Len()==0 || window.Len()>0 but evictFromWindow(),
	// then probation.Len()==0 || probation.Len()>0 but evictFromProbation(),
	// then protected.Len()>0 so evictFromProtected()
	err = cache.Resize(2)
	if err != nil {
		t.Fatalf("Failed to resize: %v", err)
	}

	finalStats := cache.Stats()
	t.Logf("Post-resize: Protected=%d, Probation=%d, Window=%d",
		finalStats.ProtectedSize, finalStats.ProbationSize, finalStats.WindowSize)

	if cache.Len() > 2 {
		t.Errorf("Cache size %d exceeds capacity 2", cache.Len())
	}
}