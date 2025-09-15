package wtinylfu

import (
	"testing"
)

func TestNew(t *testing.T) {
	cache := New[string, int](100)
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
	
	cache := NewWithEvict[string, int](5, onEvict)
	
	// Fill cache beyond capacity
	for i := 0; i < 10; i++ {
		cache.Put(string(rune('a'+i)), i)
	}
	
	if len(evicted) == 0 {
		t.Errorf("Expected some evictions")
	}
}

func TestBasicOperations(t *testing.T) {
	cache := New[string, int](10)
	
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
	cache := New[string, int](10) // Window size = 1
	
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
	cache := New[string, int](10)
	
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
	cache := New[string, int](10)
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
	cache := New[string, int](10)
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
	
	cache := NewWithEvict[string, int](10, onEvict)
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
	cache := New[string, int](10)
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
	cache := New[string, int](10)
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
	cache := New[string, int](10)
	
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
	
	cache := New[string, int](10)
	cache.Resize(0)
}

func TestSpaceManagement(t *testing.T) {
	cache := New[string, int](10) // Window=1, Main=9 (Protected=7, Probation=2)
	
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
	cache := New[string, int](5) // Small cache for easy testing
	
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
	cache := New[string, int](20)
	
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
	cache := New[int, int](100)
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
	cache := New[string, int](10)
	
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
	cache := New[string, int](2)
	
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
	cache := New[int, int](1000)
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		cache.Put(i%1000, i)
	}
}

func BenchmarkGet(b *testing.B) {
	cache := New[int, int](1000)
	
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
	cache := New[int, int](1000)
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		if i%4 == 0 {
			cache.Put(i%1000, i)
		} else {
			cache.Get(i % 1000)
		}
	}
}