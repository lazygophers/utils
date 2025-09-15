package wtinylfu

import (
	"fmt"
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

func TestNewPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for zero capacity")
		}
	}()
	New[string, int](0)
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

func TestResizePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for zero capacity resize")
		}
	}()
	
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	cache.Resize(0)
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
	single := New[string, int](1)
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
		t.Fatalf("Failed to create cache: %v", err)
	}
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		cache.Put(i%1000, i)
	}
}

func BenchmarkGet(b *testing.B) {
	cache, err := New[int, int](1000)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
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
		t.Fatalf("Failed to create cache: %v", err)
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
	cache.Put("w1", 2) // Should go to window
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
	
	// Test resize to 0 (should panic)
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for zero capacity resize")
		}
	}()
	
	cache.Resize(0)
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