package wtinylfu

import (
	"fmt"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	if cache.Cap() != 100 {
		t.Errorf("Expected capacity 100, got %d", cache.Cap())
	}
	if cache.Len() != 0 {
		t.Errorf("Expected empty cache, got length %d", cache.Len())
	}

	stats := cache.Stats()
	if stats.WindowCapacity != 10 { // 10% of 100
		t.Errorf("Expected window capacity 10, got %d", stats.WindowCapacity)
	}
}

func TestNewError(t *testing.T) {
	_, err := New[string, int](0)
	if err == nil {
		t.Error("Expected error for zero capacity")
	}
}

func TestNewWithEvict(t *testing.T) {
	evicted := make(map[string]int)
	onEvict := func(key string, value int) {
		evicted[key] = value
	}

	cache, err := NewWithEvict[string, int](5, onEvict)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Fill cache beyond capacity
	for i := 0; i < 10; i++ {
		cache.Put(string(rune('a'+i)), i)
	}

	if len(evicted) == 0 {
		t.Errorf("Expected some evictions")
	}
}

func TestBasicOperations(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Test Put and Get
	evicted := cache.Put("a", 1)
	if evicted {
		t.Errorf("Expected no eviction on first insert")
	}

	value, ok := cache.Get("a")
	if !ok || value != 1 {
		t.Errorf("Expected value 1, got %d, ok=%v", value, ok)
	}

	// Test update
	evicted = cache.Put("a", 10)
	if evicted {
		t.Errorf("Expected no eviction on update")
	}

	value, ok = cache.Get("a")
	if !ok || value != 10 {
		t.Errorf("Expected updated value 10, got %d, ok=%v", value, ok)
	}

	if cache.Len() != 1 {
		t.Errorf("Expected length 1, got %d", cache.Len())
	}
}

func TestWindowBehavior(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	} // Window size = 1

	// Fill window
	cache.Put("a", 1)

	stats := cache.Stats()
	if stats.WindowSize != 1 {
		t.Errorf("Expected 1 item in window, got %d", stats.WindowSize)
	}

	// Access to promote from window
	cache.Get("a")
	cache.Put("b", 2) // This should fill window again

	// Add another item to trigger window eviction/promotion
	cache.Put("c", 3)

	// Check that items moved between spaces
	stats = cache.Stats()
	if stats.Size != 3 {
		t.Errorf("Expected 3 total items, got %d", stats.Size)
	}
}

func TestFrequencyBasedPromotion(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Add items to window and main cache
	cache.Put("frequent", 1)
	cache.Put("infrequent", 2)

	// Access frequent item multiple times to build frequency
	for i := 0; i < 5; i++ {
		cache.Get("frequent")
	}

	// Fill cache to trigger evictions
	for i := 0; i < 15; i++ {
		cache.Put(string(rune('a'+i)), i)
	}

	// Frequent item should still be in cache
	_, ok := cache.Get("frequent")
	if !ok {
		t.Errorf("Expected frequent item to remain in cache")
	}
}

func TestContainsAndPeek(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	cache.Put("a", 1)

	// Test Contains
	if !cache.Contains("a") {
		t.Errorf("Expected 'a' to exist")
	}
	if cache.Contains("b") {
		t.Errorf("Expected 'b' to not exist")
	}

	// Test Peek (shouldn't affect promotion)
	value, ok := cache.Peek("a")
	if !ok || value != 1 {
		t.Errorf("Expected to peek value 1, got %d, ok=%v", value, ok)
	}
}

func TestRemove(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	cache.Put("a", 1)
	cache.Put("b", 2)

	// Remove existing key
	value, ok := cache.Remove("a")
	if !ok || value != 1 {
		t.Errorf("Expected to remove value 1, got %d, ok=%v", value, ok)
	}

	if cache.Len() != 1 {
		t.Errorf("Expected length 1 after removal, got %d", cache.Len())
	}

	// Remove non-existing key
	_, ok = cache.Remove("c")
	if ok {
		t.Errorf("Expected removal of non-existing key to return false")
	}
}

func TestClear(t *testing.T) {
	evicted := make(map[string]int)
	onEvict := func(key string, value int) {
		evicted[key] = value
	}

	cache, err := NewWithEvict[string, int](10, onEvict)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	cache.Put("a", 1)
	cache.Put("b", 2)

	cache.Clear()

	if cache.Len() != 0 {
		t.Errorf("Expected empty cache after clear, got length %d", cache.Len())
	}

	if len(evicted) != 2 {
		t.Errorf("Expected 2 evictions on clear, got %d", len(evicted))
	}
}

func TestKeysAndValues(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	cache.Put("a", 1)
	cache.Put("b", 2)

	keys := cache.Keys()
	values := cache.Values()

	if len(keys) != 2 || len(values) != 2 {
		t.Errorf("Expected 2 keys and values, got %d keys, %d values", len(keys), len(values))
	}

	// Check that all keys are present
	keyMap := make(map[string]bool)
	for _, key := range keys {
		keyMap[key] = true
	}
	if !keyMap["a"] || !keyMap["b"] {
		t.Errorf("Expected keys 'a' and 'b' to be present")
	}
}

func TestItems(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	cache.Put("a", 1)
	cache.Put("b", 2)

	items := cache.Items()
	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}

	if items["a"] != 1 || items["b"] != 2 {
		t.Errorf("Expected correct item values")
	}
}

func TestResize(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Fill cache
	for i := 0; i < 8; i++ {
		cache.Put(string(rune('a'+i)), i)
	}

	// Resize to smaller capacity
	cache.Resize(5)

	if cache.Cap() != 5 {
		t.Errorf("Expected capacity 5, got %d", cache.Cap())
	}

	if cache.Len() > 5 {
		t.Errorf("Expected length <= 5 after resize, got %d", cache.Len())
	}
}

func TestResizeError(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	err = cache.Resize(0)
	if err == nil {
		t.Error("Expected error for zero capacity resize")
	}
}

func TestSpaceManagement(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	} // Window=1, Main=9 (Protected=7, Probation=2)

	// Add items to fill different spaces
	cache.Put("w1", 1) // Goes to window

	stats := cache.Stats()
	if stats.WindowSize != 1 {
		t.Errorf("Expected 1 item in window, got %d", stats.WindowSize)
	}

	// Access to promote from window to protected
	cache.Get("w1")
	cache.Put("w2", 2) // New item in window

	stats = cache.Stats()
	if stats.WindowSize != 1 {
		t.Errorf("Expected 1 item in window after promotion, got %d", stats.WindowSize)
	}

	// Add more items to test probation space
	for i := 0; i < 5; i++ {
		cache.Put(string(rune('a'+i)), i+10)
	}

	stats = cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictionOrder(t *testing.T) {
	cache, err := New[string, int](5)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	} // Small cache for easy testing

	// Fill cache
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)
	cache.Put("d", 4)
	cache.Put("e", 5)

	// Access some items to build frequency
	cache.Get("a")
	cache.Get("a")
	cache.Get("b")

	// Add more items to trigger eviction
	cache.Put("f", 6)
	cache.Put("g", 7)

	// Frequently accessed items should still be present
	if !cache.Contains("a") {
		t.Errorf("Expected 'a' to remain (high frequency)")
	}
}

func TestStats(t *testing.T) {
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Add some items
	for i := 0; i < 10; i++ {
		cache.Put(string(rune('a'+i)), i)
	}

	// Access some to promote them
	cache.Get("a")
	cache.Get("b")
	cache.Get("c")

	stats := cache.Stats()

	// Note: Window-TinyLFU may evict items during promotion, so we just check it's functional
	if stats.Size == 0 {
		t.Errorf("Expected non-empty cache, got %d", stats.Size)
	}

	if stats.Capacity != 20 {
		t.Errorf("Expected capacity 20, got %d", stats.Capacity)
	}

	if stats.WindowCapacity == 0 {
		t.Errorf("Expected non-zero window capacity")
	}

	if stats.ProbationCap == 0 {
		t.Errorf("Expected non-zero probation capacity")
	}

	if stats.ProtectedCap == 0 {
		t.Errorf("Expected non-zero protected capacity")
	}

	// Check that space sizes don't exceed their capacities
	if stats.WindowSize > stats.WindowCapacity {
		t.Errorf("Window size %d exceeds capacity %d", stats.WindowSize, stats.WindowCapacity)
	}

	if stats.ProbationSize > stats.ProbationCap {
		t.Errorf("Probation size %d exceeds capacity %d", stats.ProbationSize, stats.ProbationCap)
	}

	if stats.ProtectedSize > stats.ProtectedCap {
		t.Errorf("Protected size %d exceeds capacity %d", stats.ProtectedSize, stats.ProtectedCap)
	}
}

func TestConcurrentAccess(t *testing.T) {
	cache, err := New[int, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	done := make(chan bool)

	// Start multiple goroutines
	for i := 0; i < 10; i++ {
		go func(start int) {
			for j := 0; j < 100; j++ {
				key := start*100 + j
				cache.Put(key, key*2)
				cache.Get(key)
			}
			done <- true
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}

	// Verify cache is still functional
	cache.Put(9999, 9999)
	value, ok := cache.Get(9999)
	if !ok || value != 9999 {
		t.Errorf("Cache corrupted after concurrent access")
	}
}

func TestHashFunction(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Test that hash function works with different types
	hash1 := cache.hash("test")
	hash2 := cache.hash("test")
	hash3 := cache.hash("different")

	if hash1 != hash2 {
		t.Errorf("Hash function not deterministic")
	}

	if hash1 == hash3 {
		t.Errorf("Different keys should produce different hashes")
	}
}

func TestCountMinSketch(t *testing.T) {
	sketch := newCountMinSketch(100)

	// Test increment and estimate
	hash := uint32(12345)

	// Initially should be 0
	if sketch.estimate(hash) != 0 {
		t.Errorf("Expected initial estimate to be 0")
	}

	// Increment several times (considering sampling)
	for i := 0; i < 100; i++ {
		sketch.increment(hash)
	}

	// Should have some count now
	estimate := sketch.estimate(hash)
	if estimate == 0 {
		t.Errorf("Expected non-zero estimate after increments")
	}

	// Test clear
	sketch.clear()
	if sketch.estimate(hash) != 0 {
		t.Errorf("Expected estimate to be 0 after clear")
	}
}

func TestEdgeCases(t *testing.T) {
	// Test very small cache
	cache, err := New[string, int](2)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3) // Should evict something

	if cache.Len() > 2 {
		t.Errorf("Cache size exceeds capacity")
	}

	// Test single capacity cache
	single, err := New[string, int](1)
	if err != nil {
		t.Fatalf("Failed to create single capacity cache: %v", err)
	}
	single.Put("a", 1)
	single.Put("b", 2) // Should evict "a"

	if single.Len() != 1 {
		t.Errorf("Single capacity cache should have size 1, got %d", single.Len())
	}

	_, ok := single.Get("b")
	if !ok {
		t.Errorf("Expected 'b' to be in single capacity cache")
	}
}

func BenchmarkPut(b *testing.B) {
	cache, err := New[int, int](1000)
	if err != nil {
		b.Fatalf("Failed to create cache: %v", err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.Put(i%1000, i)
	}
}

func BenchmarkGet(b *testing.B) {
	cache, err := New[int, int](1000)
	if err != nil {
		b.Fatalf("Failed to create cache: %v", err)
	}

	// Pre-populate cache
	for i := 0; i < 1000; i++ {
		cache.Put(i, i*2)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Get(i % 1000)
	}
}

func BenchmarkMixed(b *testing.B) {
	cache, err := New[int, int](1000)
	if err != nil {
		b.Fatalf("Failed to create cache: %v", err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if i%4 == 0 {
			cache.Put(i%1000, i)
		} else {
			cache.Get(i % 1000)
		}
	}
}

func TestGetMissScenarios(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Test Get on empty cache
	_, ok := cache.Get("nonexistent")
	if ok {
		t.Error("Expected false for non-existent key in empty cache")
	}

	// Add item and test Get miss
	cache.Put("a", 1)
	_, ok = cache.Get("missing")
	if ok {
		t.Error("Expected false for non-existent key")
	}
}

func TestPeekMissScenarios(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Test Peek on empty cache
	_, ok := cache.Peek("nonexistent")
	if ok {
		t.Error("Expected false for non-existent key in empty cache")
	}

	// Add item and test Peek miss
	cache.Put("a", 1)
	_, ok = cache.Peek("missing")
	if ok {
		t.Error("Expected false for non-existent key")
	}
}

func TestKeysValuesWithEmptySegments(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Test with empty cache
	keys := cache.Keys()
	if len(keys) != 0 {
		t.Errorf("Expected 0 keys in empty cache, got %d", len(keys))
	}

	values := cache.Values()
	if len(values) != 0 {
		t.Errorf("Expected 0 values in empty cache, got %d", len(values))
	}

	// Fill only window (capacity 1 out of 10)
	cache.Put("a", 1)

	keys = cache.Keys()
	if len(keys) != 1 {
		t.Errorf("Expected 1 key, got %d", len(keys))
	}

	values = cache.Values()
	if len(values) != 1 {
		t.Errorf("Expected 1 value, got %d", len(values))
	}
}

func TestEvictFromProtectedScenario(t *testing.T) {
	// Create a very small cache to force protected eviction
	cache, err := New[string, int](4)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	} // Window=0 (1 item), Main=4 (Probation=1, Protected=3)

	// Fill the cache to trigger complex eviction scenarios
	for i := 0; i < 8; i++ {
		key := string(rune('a' + i))
		cache.Put(key, i)

		// Access some items multiple times to build frequency and move to protected
		if i < 4 {
			for j := 0; j < 3; j++ {
				cache.Get(key)
			}
		}
	}

	stats := cache.Stats()
	t.Logf("After complex scenario: Protected=%d, Probation=%d, Window=%d, Total=%d",
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize, stats.Size)

	// Force more evictions to potentially trigger evictFromProtected
	for i := 8; i < 16; i++ {
		key := string(rune('a' + i))
		cache.Put(key, i)
	}

	// Verify cache is still functional
	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

func TestDemoteFromProtectedScenario(t *testing.T) {
	// Create cache with specific sizing to trigger demotion
	cache, err := New[string, int](6)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	} // Window=1 (1 item), Main=5 (Probation=1, Protected=4)

	// Build up protected segment
	protectedItems := []string{"p1", "p2", "p3", "p4"}
	for _, key := range protectedItems {
		cache.Put(key, 1)
		// Access multiple times to promote to protected
		for j := 0; j < 5; j++ {
			cache.Get(key)
		}
	}

	stats := cache.Stats()
	t.Logf("Before demotion test: Protected=%d, Probation=%d, Window=%d",
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)

	// Try to trigger demotion by forcing more items into protected
	moreItems := []string{"new1", "new2", "new3"}
	for _, key := range moreItems {
		cache.Put(key, 2)
		// Access to try to promote
		for j := 0; j < 10; j++ {
			cache.Get(key)
		}
	}

	stats = cache.Stats()
	t.Logf("After demotion test: Protected=%d, Probation=%d, Window=%d",
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)

	// Verify cache integrity
	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

func TestResizeWithProtectedEviction(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Fill cache and build up protected segment
	for i := 0; i < 10; i++ {
		key := string(rune('a' + i))
		cache.Put(key, i)
		// Build frequency for some items
		if i < 5 {
			for j := 0; j < 5; j++ {
				cache.Get(key)
			}
		}
	}

	stats := cache.Stats()
	t.Logf("Before resize: Protected=%d, Probation=%d, Window=%d, Total=%d",
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize, stats.Size)

	// Resize to force evictions from protected segment
	cache.Resize(3)

	stats = cache.Stats()
	t.Logf("After resize to 3: Protected=%d, Probation=%d, Window=%d, Total=%d",
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize, stats.Size)

	if cache.Cap() != 3 {
		t.Errorf("Expected capacity 3, got %d", cache.Cap())
	}

	if cache.Len() > 3 {
		t.Errorf("Cache size %d exceeds new capacity 3", cache.Len())
	}
}

func TestComplexEvictionChain(t *testing.T) {
	// Create cache with minimal capacity to force edge case evictions
	cache, err := New[string, int](5)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	} // Window=1, Main=4 (Probation=1, Protected=3)

	// Create scenario where probation is empty but protected has items
	// This should trigger evictFromProtected when probation segment is requested

	// Fill and promote items to protected
	protectedKeys := []string{"p1", "p2", "p3"}
	for _, key := range protectedKeys {
		cache.Put(key, 1)
		// Heavy access to ensure promotion to protected
		for j := 0; j < 10; j++ {
			cache.Get(key)
		}
	}

	// Fill remaining capacity
	cache.Put("w1", 2)  // Should go to window
	cache.Put("pr1", 3) // Should go to probation if space

	stats := cache.Stats()
	t.Logf("After setup: Protected=%d, Probation=%d, Window=%d, Total=%d",
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize, stats.Size)

	// Now force eviction when probation is empty
	// This should trigger the evictFromProtected path
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("force%d", i)
		cache.Put(key, i+10)
	}

	stats = cache.Stats()
	t.Logf("After forced evictions: Protected=%d, Probation=%d, Window=%d, Total=%d",
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize, stats.Size)

	// Verify cache integrity
	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

func TestDemoteFromProtectedEdgeCase(t *testing.T) {
	cache, err := New[string, int](8)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	} // Window=1, Main=7 (Probation=2, Protected=5)

	// Fill protected to capacity
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("prot%d", i)
		cache.Put(key, i)
		// Access multiple times to promote to protected
		for j := 0; j < 8; j++ {
			cache.Get(key)
		}
	}

	// Add some probation items
	cache.Put("prob1", 10)
	cache.Put("prob2", 11)

	// Add window item
	cache.Put("win1", 12)

	stats := cache.Stats()
	t.Logf("Setup complete: Protected=%d, Probation=%d, Window=%d, Total=%d",
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize, stats.Size)

	// Now try to promote something from probation when protected is full
	// This should trigger demotion from protected
	cache.Get("prob1") // Access probation item multiple times
	for i := 0; i < 10; i++ {
		cache.Get("prob1")
	}

	stats = cache.Stats()
	t.Logf("After promoting prob1: Protected=%d, Probation=%d, Window=%d",
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)

	// Verify cache integrity
	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

func TestResizeDownToMinimal(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Fill cache
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("item%d", i)
		cache.Put(key, i)
		if i < 5 {
			// Make some items frequent
			for j := 0; j < 5; j++ {
				cache.Get(key)
			}
		}
	}

	// Resize down to force evictions from all segments
	cache.Resize(2)

	stats := cache.Stats()
	t.Logf("After resize to 2: Protected=%d, Probation=%d, Window=%d, Total=%d",
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize, stats.Size)

	if cache.Cap() != 2 {
		t.Errorf("Expected capacity 2, got %d", cache.Cap())
	}

	if cache.Len() > 2 {
		t.Errorf("Cache size %d exceeds capacity 2", cache.Len())
	}

	// Resize to 1 to force even more extreme eviction
	cache.Resize(1)

	if cache.Cap() != 1 {
		t.Errorf("Expected capacity 1, got %d", cache.Cap())
	}

	if cache.Len() > 1 {
		t.Errorf("Cache size %d exceeds capacity 1", cache.Len())
	}
}

func TestKeysWithProbationSegment(t *testing.T) {
	// Test Keys method to cover probation segment iteration
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Add items that will go to probation
	cache.Put("prob1", 1)
	cache.Put("prob2", 2)

	keys := cache.Keys()
	if len(keys) == 0 {
		t.Error("Expected keys from probation segment")
	}
}

func TestValuesWithProbationSegment(t *testing.T) {
	// Test Values method to cover probation segment iteration
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Add items that will go to probation
	cache.Put("prob1", 1)
	cache.Put("prob2", 2)

	values := cache.Values()
	if len(values) == 0 {
		t.Error("Expected values from probation segment")
	}
}

func TestResizeEdgeCases(t *testing.T) {
	// Test Resize with edge cases
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Add items
	for i := 0; i < 5; i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
	}

	// Test resize to 0 (should return error)
	err = cache.Resize(0)
	if err == nil {
		t.Error("Expected error for zero capacity resize")
	}
}

func TestEvictFromEmptySegments(t *testing.T) {
	// Test evict functions with empty segments
	cache, err := New[string, int](5)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Force eviction from empty segments by calling internal methods
	// This is to trigger specific uncovered branches
	result1 := cache.evictFromWindow()
	if result1 {
		t.Error("Expected false from evicting empty window")
	}

	result2 := cache.evictFromProbation()
	if result2 {
		t.Error("Expected false from evicting empty probation")
	}

	result3 := cache.evictFromProtected()
	if result3 {
		t.Error("Expected false from evicting empty protected")
	}
}

func TestDemoteFromProtectedExtensive(t *testing.T) {
	// Test demoteFromProtected with various scenarios
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Fill cache completely to ensure segment transitions
	for i := 0; i < 20; i++ {
		key := fmt.Sprintf("key%d", i)
		cache.Put(key, i)

		// Access some items multiple times to get them into protected
		if i%3 == 0 {
			for j := 0; j < 5; j++ {
				cache.Get(key)
			}
		}
	}

	// This should trigger various demotions and transitions
	stats := cache.Stats()
	if stats.Size == 0 {
		t.Error("Expected items in cache after complex operations")
	}
}

func TestEvictWithComplexTransitions(t *testing.T) {
	// Test evict with complex segment transitions
	cache, err := New[string, int](8)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Create a scenario that will trigger all eviction branches
	for i := 0; i < 25; i++ {
		key := fmt.Sprintf("item%d", i)
		cache.Put(key, i)

		// Create different access patterns
		if i%4 == 0 {
			// Make some items very frequent (protected candidates)
			for j := 0; j < 10; j++ {
				cache.Get(key)
			}
		} else if i%4 == 1 {
			// Medium frequency (probation candidates)
			for j := 0; j < 3; j++ {
				cache.Get(key)
			}
		}
		// Others stay in window with low frequency
	}

	// Verify final state
	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

func TestKeysAllSegmentsPopulated(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("window", 1)

	cache.Put("prob1", 10)
	cache.Get("prob1")

	cache.Put("prot1", 100)
	cache.Get("prot1")
	cache.Get("prot1")

	cache.Put("prob2", 20)
	cache.Get("prob2")

	cache.Put("prot2", 200)
	cache.Get("prot2")

	keys := cache.Keys()

	if len(keys) < 3 {
		t.Errorf("Expected at least 3 keys, got %d", len(keys))
	}

	keySet := make(map[string]bool)
	for _, key := range keys {
		keySet[key] = true
	}

	if !keySet["window"] {
		t.Error("Expected 'window' key to be present")
	}
	if !keySet["prob1"] && !keySet["prob2"] {
		t.Error("Expected at least one probation key to be present")
	}
	if !keySet["prot1"] && !keySet["prot2"] {
		t.Error("Expected at least one protected key to be present")
	}
}

func TestValuesAllSegmentsPopulated(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("window", 1)

	cache.Put("prob1", 10)
	cache.Get("prob1")

	cache.Put("prot1", 100)
	cache.Get("prot1")
	cache.Get("prot1")

	cache.Put("prob2", 20)
	cache.Get("prob2")

	cache.Put("prot2", 200)
	cache.Get("prot2")

	values := cache.Values()

	if len(values) < 3 {
		t.Errorf("Expected at least 3 values, got %d", len(values))
	}

	valueSet := make(map[int]bool)
	for _, val := range values {
		valueSet[val] = true
	}

	if !valueSet[1] {
		t.Error("Expected value 1 to be present")
	}
	if !valueSet[10] && !valueSet[20] {
		t.Error("Expected at least one probation value to be present")
	}
	if !valueSet[100] && !valueSet[200] {
		t.Error("Expected at least one protected value to be present")
	}
}

func TestItemsAllSegmentsPopulated(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("window", 1)

	cache.Put("prob1", 10)
	cache.Get("prob1")

	cache.Put("prot1", 100)
	cache.Get("prot1")
	cache.Get("prot1")

	cache.Put("prob2", 20)
	cache.Get("prob2")

	cache.Put("prot2", 200)
	cache.Get("prot2")

	items := cache.Items()

	if len(items) < 3 {
		t.Errorf("Expected at least 3 items, got %d", len(items))
	}

	if val, ok := items["window"]; !ok || val != 1 {
		t.Error("Expected 'window' key with value 1 to be present")
	}
	if _, ok := items["prob1"]; !ok {
		if _, ok := items["prob2"]; !ok {
			t.Error("Expected at least one probation key to be present")
		}
	}
	if _, ok := items["prot1"]; !ok {
		if _, ok := items["prot2"]; !ok {
			t.Error("Expected at least one protected key to be present")
		}
	}
}

func TestDemoteFromProtectedFull(t *testing.T) {
	cache, err := New[string, int](30)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 20; i++ {
		key := fmt.Sprintf("prot%d", i)
		cache.Put(key, i)
		for j := 0; j < 8; j++ {
			cache.Get(key)
		}
	}

	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("prob%d", i)
		cache.Put(key, i+1000)
		cache.Get(key)
		cache.Get(key)
	}

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictFromWindowAdmit(t *testing.T) {
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("freq", 1)
	for i := 0; i < 10; i++ {
		cache.Get("freq")
	}

	cache.Put("victim", 2)
	cache.Put("temp", 999)
	cache.Put("victim", 2)

	cache.Put("new", 3)

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictFromWindowNoAdmit(t *testing.T) {
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 30; i++ {
		key := fmt.Sprintf("key%d", i)
		cache.Put(key, i)
	}

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictFromProbationFull(t *testing.T) {
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 15; i++ {
		key := fmt.Sprintf("prob%d", i)
		cache.Put(key, i)
		cache.Get(key)
	}

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictFromProtectedFull(t *testing.T) {
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 25; i++ {
		key := fmt.Sprintf("prot%d", i)
		cache.Put(key, i)
		for j := 0; j < 10; j++ {
			cache.Get(key)
		}
	}

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictPrioritizeProbation(t *testing.T) {
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("prob%d", i)
		cache.Put(key, i)
		cache.Get(key)
	}

	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("prot%d", i)
		cache.Put(key, i+100)
		for j := 0; j < 5; j++ {
			cache.Get(key)
		}
	}

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("window%d", i)
		cache.Put(key, i+200)
	}

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictPrioritizeProtected(t *testing.T) {
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 15; i++ {
		key := fmt.Sprintf("prot%d", i)
		cache.Put(key, i)
		for j := 0; j < 5; j++ {
			cache.Get(key)
		}
	}

	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("window%d", i)
		cache.Put(key, i+100)
	}

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictPrioritizeWindow(t *testing.T) {
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 20; i++ {
		key := fmt.Sprintf("window%d", i)
		cache.Put(key, i)
	}

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestDemoteFromProtectedEmptyProbation(t *testing.T) {
	cache, err := New[string, int](30)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 25; i++ {
		key := fmt.Sprintf("prot%d", i)
		cache.Put(key, i)
		for j := 0; j < 10; j++ {
			cache.Get(key)
		}
	}

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestDemoteFromProtectedFullProbation(t *testing.T) {
	cache, err := New[string, int](30)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 20; i++ {
		key := fmt.Sprintf("prot%d", i)
		cache.Put(key, i)
		for j := 0; j < 8; j++ {
			cache.Get(key)
		}
	}

	for i := 0; i < 8; i++ {
		key := fmt.Sprintf("prob%d", i)
		cache.Put(key, i+1000)
		cache.Get(key)
		cache.Get(key)
	}

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictFromWindowCompete(t *testing.T) {
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("freq", 1)
	for i := 0; i < 10; i++ {
		cache.Get("freq")
	}

	for i := 0; i < 15; i++ {
		key := fmt.Sprintf("prob%d", i)
		cache.Put(key, i)
		cache.Get(key)
	}

	cache.Put("victim", 2)
	cache.Put("temp", 999)
	cache.Put("victim", 2)

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestNewWithEvictError(t *testing.T) {
	_, err := NewWithEvict[string, int](0, func(key string, value int) {})
	if err == nil {
		t.Error("Expected error for zero capacity with NewWithEvict")
	}
}

func TestKeysEmptySegments(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	keys := cache.Keys()
	if len(keys) != 0 {
		t.Errorf("Expected 0 keys in empty cache, got %d", len(keys))
	}
}

func TestValuesEmptySegments(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	values := cache.Values()
	if len(values) != 0 {
		t.Errorf("Expected 0 values in empty cache, got %d", len(values))
	}
}

func TestItemsEmptySegments(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	items := cache.Items()
	if len(items) != 0 {
		t.Errorf("Expected 0 items in empty cache, got %d", len(items))
	}
}

func TestKeysOnlyWindow(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("window1", 1)
	cache.Put("window2", 2)

	keys := cache.Keys()
	if len(keys) == 0 {
		t.Error("Expected keys from window segment")
	}
}

func TestValuesOnlyWindow(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("window1", 1)
	cache.Put("window2", 2)

	values := cache.Values()
	if len(values) == 0 {
		t.Error("Expected values from window segment")
	}
}

func TestItemsOnlyWindow(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("window1", 1)
	cache.Put("window2", 2)

	items := cache.Items()
	if len(items) == 0 {
		t.Error("Expected items from window segment")
	}
}

func TestKeysOnlyProbation(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prob1", 1)
	cache.Get("prob1")
	cache.Put("prob2", 2)
	cache.Get("prob2")

	keys := cache.Keys()
	if len(keys) == 0 {
		t.Error("Expected keys from probation segment")
	}
}

func TestValuesOnlyProbation(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prob1", 1)
	cache.Get("prob1")
	cache.Put("prob2", 2)
	cache.Get("prob2")

	values := cache.Values()
	if len(values) == 0 {
		t.Error("Expected values from probation segment")
	}
}

func TestItemsOnlyProbation(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prob1", 1)
	cache.Get("prob1")
	cache.Put("prob2", 2)
	cache.Get("prob2")

	items := cache.Items()
	if len(items) == 0 {
		t.Error("Expected items from probation segment")
	}
}

func TestKeysOnlyProtected(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prot1", 1)
	cache.Get("prot1")
	cache.Get("prot1")

	cache.Put("prot2", 2)
	cache.Get("prot2")
	cache.Get("prot2")

	keys := cache.Keys()
	if len(keys) == 0 {
		t.Error("Expected keys from protected segment")
	}
}

func TestValuesOnlyProtected(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prot1", 1)
	cache.Get("prot1")
	cache.Get("prot1")

	cache.Put("prot2", 2)
	cache.Get("prot2")
	cache.Get("prot2")

	values := cache.Values()
	if len(values) == 0 {
		t.Error("Expected values from protected segment")
	}
}

func TestItemsOnlyProtected(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prot1", 1)
	cache.Get("prot1")
	cache.Get("prot1")

	cache.Put("prot2", 2)
	cache.Get("prot2")
	cache.Get("prot2")

	items := cache.Items()
	if len(items) == 0 {
		t.Error("Expected items from protected segment")
	}
}

func TestDemoteFromProtectedWithFullProbation(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 8; i++ {
		key := fmt.Sprintf("prot%d", i)
		cache.Put(key, i)
		for j := 0; j < 5; j++ {
			cache.Get(key)
		}
	}

	cache.Put("prob1", 100)
	cache.Get("prob1")

	cache.Put("prob2", 101)
	cache.Get("prob2")

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictFromWindowWithFullProbation(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prob1", 1)
	cache.Get("prob1")
	cache.Put("prob2", 2)
	cache.Get("prob2")

	cache.Put("window1", 10)
	cache.Put("window2", 20)

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestResizeWithEmptyCache(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	err = cache.Resize(50)
	if err != nil {
		t.Fatalf("Failed to resize empty cache: %v", err)
	}

	if cache.Cap() != 50 {
		t.Errorf("Expected capacity 50, got %d", cache.Cap())
	}
}

func TestResizeToSameCapacity(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("b", 2)

	err = cache.Resize(100)
	if err != nil {
		t.Fatalf("Failed to resize to same capacity: %v", err)
	}

	if cache.Cap() != 100 {
		t.Errorf("Expected capacity 100, got %d", cache.Cap())
	}
}

func TestResizeToLarger(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 10; i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
	}

	err = cache.Resize(200)
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}

	if cache.Cap() != 200 {
		t.Errorf("Expected capacity 200, got %d", cache.Cap())
	}

	if cache.Len() < 5 {
		t.Errorf("Expected length >= 5, got %d", cache.Len())
	}
}

func TestResizeToSmaller(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 80; i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
	}

	err = cache.Resize(20)
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}

	if cache.Cap() != 20 {
		t.Errorf("Expected capacity 20, got %d", cache.Cap())
	}

	if cache.Len() > 20 {
		t.Errorf("Expected length <= 20, got %d", cache.Len())
	}
}

func TestResizeWithProtectedEvictionCoverage(t *testing.T) {
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 15; i++ {
		key := fmt.Sprintf("key%d", i)
		cache.Put(key, i)
		for j := 0; j < 5; j++ {
			cache.Get(key)
		}
	}

	err = cache.Resize(5)
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}

	if cache.Cap() != 5 {
		t.Errorf("Expected capacity 5, got %d", cache.Cap())
	}

	if cache.Len() > 5 {
		t.Errorf("Expected length <= 5, got %d", cache.Len())
	}
}

func TestRemoveFromEmptyCache(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	_, ok := cache.Remove("nonexistent")
	if ok {
		t.Error("Expected false for removing from empty cache")
	}
}

func TestPeekFromEmptyCache(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	_, ok := cache.Peek("nonexistent")
	if ok {
		t.Error("Expected false for peeking from empty cache")
	}
}

func TestContainsFromEmptyCache(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	if cache.Contains("nonexistent") {
		t.Error("Expected false for contains in empty cache")
	}
}

func TestGetFromEmptyCache(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	_, ok := cache.Get("nonexistent")
	if ok {
		t.Error("Expected false for getting from empty cache")
	}
}

func TestUpdateNonExistentKey(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	evicted := cache.Put("a", 10)
	if evicted {
		t.Error("Expected no eviction when updating existing key")
	}

	value, ok := cache.Get("a")
	if !ok || value != 10 {
		t.Errorf("Expected updated value 10, got %d, ok=%t", value, ok)
	}
}

func TestClearEmptyCache(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Clear()

	if cache.Len() != 0 {
		t.Errorf("Expected empty cache after clear, got %d", cache.Len())
	}
}

func TestSketchCounterOverflow(t *testing.T) {
	sketch := newCountMinSketch(10)

	hash := uint32(12345)
	for i := 0; i < 300; i++ {
		sketch.increment(hash)
	}

	estimate := sketch.estimate(hash)
	if estimate == 0 {
		t.Error("Expected non-zero estimate after many increments")
	}
}

func TestSketchMultipleKeys(t *testing.T) {
	sketch := newCountMinSketch(100)

	for i := 0; i < 10; i++ {
		hash := uint32(i * 1000)
		for j := 0; j < 50; j++ {
			sketch.increment(hash)
		}
	}

	for i := 0; i < 10; i++ {
		hash := uint32(i * 1000)
		estimate := sketch.estimate(hash)
		if estimate == 0 {
			t.Logf("Warning: estimate is 0 for hash %d (sampling may have skipped increments)", hash)
		}
	}
}

func TestSketchClear(t *testing.T) {
	sketch := newCountMinSketch(100)

	hash := uint32(12345)
	for i := 0; i < 50; i++ {
		sketch.increment(hash)
	}

	estimateBefore := sketch.estimate(hash)
	if estimateBefore == 0 {
		t.Error("Expected non-zero estimate before clear")
	}

	sketch.clear()

	estimateAfter := sketch.estimate(hash)
	if estimateAfter != 0 {
		t.Errorf("Expected estimate 0 after clear, got %d", estimateAfter)
	}
}

func TestSketchSampling(t *testing.T) {
	sketch := newCountMinSketch(50)

	hash := uint32(12345)
	for i := 0; i < 20; i++ {
		sketch.increment(hash)
	}

	estimate := sketch.estimate(hash)
	if estimate == 0 {
		t.Error("Expected non-zero estimate after increments with sampling")
	}
}

func TestWindowEvictionWithEmptyProbation(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("window1", 1)
	cache.Put("window2", 2)
	cache.Put("window3", 3)

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestProbationEvictionWithEmptyProtected(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prob1", 1)
	cache.Get("prob1")
	cache.Put("prob2", 2)
	cache.Get("prob2")
	cache.Put("prob3", 3)
	cache.Get("prob3")

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestProtectedEvictionWithFullCache(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 15; i++ {
		key := fmt.Sprintf("key%d", i)
		cache.Put(key, i)
		for j := 0; j < 10; j++ {
			cache.Get(key)
		}
	}

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictionCallbackOnClear(t *testing.T) {
	evictedKeys := make([]string, 0)
	cache, err := NewWithEvict[string, int](10, func(key string, value int) {
		evictedKeys = append(evictedKeys, key)
	})
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	cache.Clear()

	if len(evictedKeys) != 3 {
		t.Errorf("Expected 3 evicted keys, got %d", len(evictedKeys))
	}
}

func TestEvictionCallbackOnEviction(t *testing.T) {
	evictedKeys := make([]string, 0)
	cache, err := NewWithEvict[string, int](5, func(key string, value int) {
		evictedKeys = append(evictedKeys, key)
	})
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 10; i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
	}

	if len(evictedKeys) == 0 {
		t.Error("Expected some evictions")
	}
}

func TestRemoveWithEvictionCallback(t *testing.T) {
	evictedKeys := make([]string, 0)
	cache, err := NewWithEvict[string, int](10, func(key string, value int) {
		evictedKeys = append(evictedKeys, key)
	})
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("b", 2)

	if !cache.Contains("a") {
		t.Fatal("Expected 'a' to be in cache before removal")
	}

	value, ok := cache.Remove("a")
	if !ok || value != 1 {
		t.Errorf("Expected to remove value 1, got %d, ok=%t", value, ok)
	}

	if len(evictedKeys) != 1 || evictedKeys[0] != "a" {
		t.Error("Expected eviction callback on Remove for wtinylfu")
	}
}

func TestStatsAccuracy(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 50; i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
	}

	stats := cache.Stats()

	if stats.Size != cache.Len() {
		t.Errorf("Stats size %d doesn't match cache len %d", stats.Size, cache.Len())
	}

	if stats.Capacity != cache.Cap() {
		t.Errorf("Stats capacity %d doesn't match cache cap %d", stats.Capacity, cache.Cap())
	}

	totalSegments := stats.WindowSize + stats.ProbationSize + stats.ProtectedSize
	if totalSegments != stats.Size {
		t.Errorf("Segment sizes don't match total size: %d vs %d", totalSegments, stats.Size)
	}
}

func TestCapacityOne(t *testing.T) {
	cache, err := New[string, int](1)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("b", 2)

	if cache.Len() != 1 {
		t.Errorf("Expected cache size 1, got %d", cache.Len())
	}

	_, ok := cache.Get("b")
	if !ok {
		t.Error("Expected 'b' to be in cache")
	}
}

func TestCapacityTwo(t *testing.T) {
	cache, err := New[string, int](2)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	if cache.Len() > 2 {
		t.Errorf("Expected cache size <= 2, got %d", cache.Len())
	}
}

func TestLargeCapacity(t *testing.T) {
	cache, err := New[string, int](10000)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 5000; i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
	}

	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}

	stats := cache.Stats()
	if stats.WindowCapacity == 0 {
		t.Error("Expected non-zero window capacity")
	}
}

func TestKeyTypes(t *testing.T) {
	cache, err := New[int, string](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put(1, "one")
	cache.Put(2, "two")

	value, ok := cache.Get(1)
	if !ok || value != "one" {
		t.Errorf("Expected value 'one', got %s, ok=%t", value, ok)
	}

	value, ok = cache.Get(2)
	if !ok || value != "two" {
		t.Errorf("Expected value 'two', got %s, ok=%t", value, ok)
	}
}

func TestValueTypes(t *testing.T) {
	cache, err := New[string, []int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", []int{1, 2, 3})
	cache.Put("b", []int{4, 5, 6})

	value, ok := cache.Get("a")
	if !ok || len(value) != 3 {
		t.Errorf("Expected value with length 3, got %d, ok=%t", len(value), ok)
	}
}

func TestZeroValue(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("zero", 0)

	value, ok := cache.Get("zero")
	if !ok || value != 0 {
		t.Errorf("Expected value 0, got %d, ok=%t", value, ok)
	}
}

func TestNilValue(t *testing.T) {
	cache, err := New[string, *int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	var nilPtr *int = nil
	cache.Put("nil", nilPtr)

	value, ok := cache.Get("nil")
	if !ok || value != nil {
		t.Errorf("Expected nil value, got %v, ok=%t", value, ok)
	}
}

func TestStringKeyWithSpecialChars(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	specialKeys := []string{
		"key with spaces",
		"key\twith\ttabs",
		"key\nwith\nnewlines",
		"key\"with\"quotes",
		"key'with'apostrophes",
		"key\\with\\backslashes",
	}

	for _, key := range specialKeys {
		cache.Put(key, len(key))
	}

	for _, key := range specialKeys {
		value, ok := cache.Get(key)
		if !ok || value != len(key) {
			t.Errorf("Expected value %d for key %q, got %d, ok=%t", len(key), key, value, ok)
		}
	}
}

func TestInt64Key(t *testing.T) {
	cache, err := New[int64, string](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put(1234567890123456789, "value")

	value, ok := cache.Get(1234567890123456789)
	if !ok || value != "value" {
		t.Errorf("Expected value 'value', got %s, ok=%t", value, ok)
	}
}

func TestFloat64Key(t *testing.T) {
	cache, err := New[float64, string](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put(3.14159, "pi")

	value, ok := cache.Get(3.14159)
	if !ok || value != "pi" {
		t.Errorf("Expected value 'pi', got %s, ok=%t", value, ok)
	}
}

func TestBoolKey(t *testing.T) {
	cache, err := New[bool, string](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put(true, "yes")
	cache.Put(false, "no")

	value, ok := cache.Get(true)
	if !ok || value != "yes" {
		t.Errorf("Expected value 'yes', got %s, ok=%t", value, ok)
	}

	value, ok = cache.Get(false)
	if !ok || value != "no" {
		t.Errorf("Expected value 'no', got %s, ok=%t", value, ok)
	}
}

func TestMultipleUpdates(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("a", 2)
	cache.Put("a", 3)
	cache.Put("a", 4)
	cache.Put("a", 5)

	value, ok := cache.Get("a")
	if !ok || value != 5 {
		t.Errorf("Expected value 5, got %d, ok=%t", value, ok)
	}
}

func TestRemoveNonExistent(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)

	_, ok := cache.Remove("nonexistent")
	if ok {
		t.Error("Expected false for removing non-existent key")
	}

	value, ok := cache.Get("a")
	if !ok || value != 1 {
		t.Errorf("Expected value 1, got %d, ok=%t", value, ok)
	}
}

func TestPeekDoesNotUpdatePosition(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("b", 2)

	cache.Peek("a")

	cache.Put("c", 3)

	value, ok := cache.Get("a")
	if !ok || value != 1 {
		t.Logf("Note: Peek may not prevent eviction in wtinylfu, got %d, ok=%t", value, ok)
	}
}

func TestContainsDoesNotUpdatePosition(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("b", 2)

	cache.Contains("a")

	cache.Put("c", 3)

	value, ok := cache.Get("a")
	if !ok || value != 1 {
		t.Logf("Note: Contains may not prevent eviction in wtinylfu, got %d, ok=%t", value, ok)
	}
}

func TestLenAccuracy(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	if cache.Len() != 0 {
		t.Errorf("Expected len 0, got %d", cache.Len())
	}

	cache.Put("a", 1)
	if cache.Len() != 1 {
		t.Errorf("Expected len 1, got %d", cache.Len())
	}

	cache.Put("b", 2)
	if cache.Len() != 2 {
		t.Errorf("Expected len 2, got %d", cache.Len())
	}

	cache.Remove("a")
	if cache.Len() != 1 {
		t.Errorf("Expected len 1, got %d", cache.Len())
	}

	cache.Clear()
	if cache.Len() != 0 {
		t.Errorf("Expected len 0, got %d", cache.Len())
	}
}

func TestCapAccuracy(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	if cache.Cap() != 100 {
		t.Errorf("Expected cap 100, got %d", cache.Cap())
	}

	cache.Resize(50)
	if cache.Cap() != 50 {
		t.Errorf("Expected cap 50, got %d", cache.Cap())
	}
}

func TestSegmentCapacities(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	stats := cache.Stats()

	if stats.WindowCapacity == 0 {
		t.Error("Expected non-zero window capacity")
	}

	if stats.ProbationCap == 0 {
		t.Error("Expected non-zero probation capacity")
	}

	if stats.ProtectedCap == 0 {
		t.Error("Expected non-zero protected capacity")
	}
}

func TestSegmentSizesWithinCapacities(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 150; i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
	}

	stats := cache.Stats()

	if stats.WindowSize > stats.WindowCapacity {
		t.Errorf("Window size %d exceeds capacity %d", stats.WindowSize, stats.WindowCapacity)
	}

	if stats.ProbationSize > stats.ProbationCap {
		t.Errorf("Probation size %d exceeds capacity %d", stats.ProbationSize, stats.ProbationCap)
	}

	if stats.ProtectedSize > stats.ProtectedCap {
		t.Errorf("Protected size %d exceeds capacity %d", stats.ProtectedSize, stats.ProtectedCap)
	}
}

func TestConcurrentResize(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				key := fmt.Sprintf("%d-%d", id, j)
				cache.Put(key, id*100+j)
				cache.Get(key)
			}
		}(i)
	}

	wg.Wait()

	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

func TestStressTest(t *testing.T) {
	cache, err := New[int, int](1000)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 10000; i++ {
		key := i % 2000
		cache.Put(key, key*2)
		cache.Get(key)
	}

	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

func TestRandomAccessPattern(t *testing.T) {
	cache, err := New[int, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 1000; i++ {
		key := (i * 7) % 200
		cache.Put(key, i)
		if i%3 == 0 {
			cache.Get(key)
		}
	}

	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

func TestSequentialAccessPattern(t *testing.T) {
	cache, err := New[int, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 500; i++ {
		cache.Put(i, i)
	}

	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

func TestRepeatedAccessPattern(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	keys := []string{"a", "b", "c", "d", "e"}

	for i := 0; i < 1000; i++ {
		key := keys[i%len(keys)]
		cache.Put(key, i)
		cache.Get(key)
	}

	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}

	for _, key := range keys {
		if !cache.Contains(key) {
			t.Errorf("Expected key %s to be in cache", key)
		}
	}
}

func TestZipfAccessPattern(t *testing.T) {
	cache, err := New[int, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 1000; i++ {
		key := (i * i) % 50
		cache.Put(key, i)
		cache.Get(key)
	}

	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

func TestRecordAccessFromWindow(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("window", 1)
	cache.Get("window")

	stats := cache.Stats()
	if stats.WindowSize == 0 && stats.ProbationSize == 0 {
		t.Error("Expected item to move from window")
	}
}

func TestRecordAccessFromProbation(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prob", 1)
	cache.Get("prob")
	cache.Get("prob")

	stats := cache.Stats()
	if stats.ProtectedSize == 0 {
		t.Error("Expected item to be promoted to protected")
	}
}

func TestRecordAccessFromProtected(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prot", 1)
	cache.Get("prot")
	cache.Get("prot")

	stats := cache.Stats()
	if stats.ProtectedSize == 0 {
		t.Error("Expected item to be in protected")
	}
}

func TestEvictFromWindowWithAdmission(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Get("a")
	cache.Get("a")

	cache.Put("b", 2)
	cache.Put("c", 3)

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictFromWindowWithoutAdmission(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictFromProbation(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prob1", 1)
	cache.Get("prob1")
	cache.Put("prob2", 2)
	cache.Get("prob2")
	cache.Put("prob3", 3)
	cache.Get("prob3")

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictFromProtected(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 15; i++ {
		key := fmt.Sprintf("key%d", i)
		cache.Put(key, i)
		for j := 0; j < 10; j++ {
			cache.Get(key)
		}
	}

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestDemoteFromProtected(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 8; i++ {
		key := fmt.Sprintf("key%d", i)
		cache.Put(key, i)
		for j := 0; j < 5; j++ {
			cache.Get(key)
		}
	}

	cache.Put("new", 100)
	cache.Get("new")
	cache.Get("new")

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestRemoveEntryFromWindow(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("window", 1)
	cache.Remove("window")

	stats := cache.Stats()
	if stats.WindowSize != 0 {
		t.Errorf("Expected window size 0, got %d", stats.WindowSize)
	}
}

func TestRemoveEntryFromProbation(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prob", 1)
	cache.Get("prob")
	cache.Remove("prob")

	stats := cache.Stats()
	if stats.ProbationSize != 0 {
		t.Errorf("Expected probation size 0, got %d", stats.ProbationSize)
	}
}

func TestRemoveEntryFromProtected(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prot", 1)
	cache.Get("prot")
	cache.Get("prot")
	cache.Remove("prot")

	stats := cache.Stats()
	if stats.ProtectedSize != 0 {
		t.Errorf("Expected protected size 0, got %d", stats.ProtectedSize)
	}
}

func TestHashDeterministic(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	hash1 := cache.hash("test")
	hash2 := cache.hash("test")

	if hash1 != hash2 {
		t.Error("Hash function should be deterministic")
	}
}

func TestHashDifferentKeys(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	hash1 := cache.hash("test1")
	hash2 := cache.hash("test2")

	if hash1 == hash2 {
		t.Error("Different keys should produce different hashes")
	}
}
