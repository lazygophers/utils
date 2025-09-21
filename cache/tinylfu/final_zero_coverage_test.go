package tinylfu

import (
	"fmt"
	"testing"
)

// TestFinalZeroCoverageFix creates realistic TinyLFU scenarios to trigger the 0% coverage functions
func TestFinalZeroCoverageFix(t *testing.T) {
	t.Run("demoteFromProtected_via_realistic_promotion", func(t *testing.T) {
		// Use cache size 25: window=1, main=24, protected capacity = int(24*0.8) = 19
		cache, err := New[string, int](25)
		if err != nil {
			t.Fatalf("Failed to create cache: %v", err)
		}

		// Step 1: Create multiple keys with varying frequencies
		// We need enough keys to fill protected + have some in probation
		keys := make([]string, 25)
		for i := 0; i < 25; i++ {
			keys[i] = fmt.Sprintf("key_%d", i)
		}

		// Step 2: Build frequency patterns - some keys more frequent than others
		for round := 0; round < 20; round++ {
			for i := 0; i < 25; i++ {
				key := keys[i]
				cache.Put(key, i)

				// Higher frequency for some keys (they should end up in protected)
				accessCount := 1
				if i < 20 {
					accessCount = 3 // High frequency keys
				}

				for j := 0; j < accessCount; j++ {
					cache.Get(key)
				}
			}
		}

		// Step 3: Add items that will become "victims" in probation
		for i := 0; i < 5; i++ {
			victimKey := fmt.Sprintf("victim_%d", i)
			cache.Put(victimKey, i+1000)
		}

		// Step 4: Force window evictions to get items into probation
		// Add many items to force evictions from window
		for i := 0; i < 50; i++ {
			tempKey := fmt.Sprintf("temp_%d", i)
			cache.Put(tempKey, i+2000)
		}

		// Step 5: Now access the high-frequency keys that should be in probation
		// This will promote them to protected, and should eventually trigger demoteFromProtected
		for i := 0; i < 20; i++ {
			key := keys[i]
			if cache.Contains(key) {
				cache.Get(key) // This should trigger promotion probation -> protected
			}
		}

		stats := cache.Stats()
		t.Logf("After promotions: Protected=%d, Probation=%d, Window=%d",
			stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)

		// Verify cache integrity
		if cache.Len() > cache.Cap() {
			t.Errorf("Cache overflow: %d > %d", cache.Len(), cache.Cap())
		}
	})

	t.Run("evictFromProtected_via_realistic_resize", func(t *testing.T) {
		// Create larger cache and fill it systematically
		cache, err := New[string, int](50)
		if err != nil {
			t.Fatalf("Failed to create cache: %v", err)
		}

		// Build a realistic cache state with items in all segments
		protectedKeys := make([]string, 0, 40)

		// Create high-frequency items that will end up in protected
		for i := 0; i < 40; i++ {
			key := fmt.Sprintf("protected_key_%d", i)
			protectedKeys = append(protectedKeys, key)

			// High frequency
			for round := 0; round < 15; round++ {
				cache.Put(key, i)
				cache.Get(key)
			}
		}

		// Add some medium frequency items for probation
		for i := 0; i < 10; i++ {
			key := fmt.Sprintf("probation_key_%d", i)
			for round := 0; round < 5; round++ {
				cache.Put(key, i+100)
				cache.Get(key)
			}
		}

		// Add low frequency items for window
		for i := 0; i < 10; i++ {
			key := fmt.Sprintf("window_key_%d", i)
			cache.Put(key, i+200)
		}

		initialStats := cache.Stats()
		t.Logf("Initial cache state: Protected=%d, Probation=%d, Window=%d, Total=%d",
			initialStats.ProtectedSize, initialStats.ProbationSize,
			initialStats.WindowSize, cache.Len())

		// Now resize to progressively smaller sizes to force evictions
		// This should eventually hit the evictFromProtected path

		// First resize
		err = cache.Resize(25)
		if err != nil {
			t.Fatalf("Failed to resize to 25: %v", err)
		}

		midStats := cache.Stats()
		t.Logf("After resize to 25: Protected=%d, Probation=%d, Window=%d, Total=%d",
			midStats.ProtectedSize, midStats.ProbationSize,
			midStats.WindowSize, cache.Len())

		// Second resize to force more evictions
		err = cache.Resize(10)
		if err != nil {
			t.Fatalf("Failed to resize to 10: %v", err)
		}

		// Final resize to very small size - this should force evictFromProtected
		err = cache.Resize(3)
		if err != nil {
			t.Fatalf("Failed to resize to 3: %v", err)
		}

		finalStats := cache.Stats()
		t.Logf("After resize to 3: Protected=%d, Probation=%d, Window=%d, Total=%d",
			finalStats.ProtectedSize, finalStats.ProbationSize,
			finalStats.WindowSize, cache.Len())

		if cache.Len() > 3 {
			t.Errorf("Cache size %d exceeds capacity 3", cache.Len())
		}
	})
}

// TestSystematicProtectedBuilding builds protected segment step by step
func TestSystematicProtectedBuilding(t *testing.T) {
	// Create cache with very controlled size
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Window=1, Main=9, Protected capacity = int(9*0.8) = 7

	// Phase 1: Create high-frequency keys that will be admitted
	keys := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}

	// Build frequency in sketch
	for round := 0; round < 10; round++ {
		for _, key := range keys {
			cache.Put(key, 1)
			cache.Get(key)
		}
	}

	// Phase 2: Systematically get items into probation and then promote to protected

	// Clear cache but keep frequency data
	cache.Clear()

	// Add victim with low frequency for admission policy
	cache.Put("victim", 999)

	// Get first 7 keys into protected (up to capacity)
	for i := 0; i < 7; i++ {
		key := keys[i]

		// First put - goes to window
		cache.Put(key, i+100)

		// Force eviction from window - should be admitted to probation due to frequency
		cache.Put("evict", 888)

		// Put again - should be admitted to probation
		cache.Put(key, i+100)

		// Access to promote from probation to protected
		cache.Get(key)
	}

	stats := cache.Stats()
	t.Logf("Protected segment at capacity: Protected=%d, Probation=%d, Window=%d",
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)

	// Phase 3: Add one more key to probation and promote - should trigger demoteFromProtected
	triggerKey := keys[7]

	// Get it into probation
	cache.Put(triggerKey, 800)
	cache.Put("evict2", 777)
	cache.Put(triggerKey, 800)

	// Promote to protected - this should trigger demoteFromProtected
	cache.Get(triggerKey)

	finalStats := cache.Stats()
	t.Logf("After triggering demotion: Protected=%d, Probation=%d, Window=%d",
		finalStats.ProtectedSize, finalStats.ProbationSize, finalStats.WindowSize)

	if cache.Len() > cache.Cap() {
		t.Errorf("Cache overflow: %d > %d", cache.Len(), cache.Cap())
	}
}

// TestResizePathEviction tests the specific resize eviction paths
func TestResizePathEviction(t *testing.T) {
	cache, err := New[string, int](12)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Build items with different access patterns
	protectedItems := []string{"P1", "P2", "P3", "P4", "P5", "P6", "P7"}
	probationItems := []string{"Q1", "Q2"}
	windowItems := []string{"W1"}

	// High frequency for protected items
	for _, key := range protectedItems {
		for i := 0; i < 20; i++ {
			cache.Put(key, 1)
			cache.Get(key)
		}
	}

	// Medium frequency for probation items
	for _, key := range probationItems {
		for i := 0; i < 5; i++ {
			cache.Put(key, 2)
			cache.Get(key)
		}
	}

	// Low frequency for window items
	for _, key := range windowItems {
		cache.Put(key, 3)
	}

	// Add low-frequency evictable items
	for i := 0; i < 10; i++ {
		cache.Put(fmt.Sprintf("evictable_%d", i), i)
	}

	initialStats := cache.Stats()
	t.Logf("Before resize: Protected=%d, Probation=%d, Window=%d, Total=%d",
		initialStats.ProtectedSize, initialStats.ProbationSize,
		initialStats.WindowSize, cache.Len())

	// Resize to force evictions through all segments
	err = cache.Resize(2)
	if err != nil {
		t.Fatalf("Failed to resize: %v", err)
	}

	finalStats := cache.Stats()
	t.Logf("After resize: Protected=%d, Probation=%d, Window=%d, Total=%d",
		finalStats.ProtectedSize, finalStats.ProbationSize,
		finalStats.WindowSize, cache.Len())

	if cache.Len() > 2 {
		t.Errorf("Cache size %d exceeds capacity 2", cache.Len())
	}
}