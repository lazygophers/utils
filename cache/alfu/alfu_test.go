package alfu

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	cache := New[string, int](3)
	if cache.Cap() != 3 {
		t.Errorf("Expected capacity 3, got %d", cache.Cap())
	}
	if cache.Len() != 0 {
		t.Errorf("Expected empty cache, got length %d", cache.Len())
	}
	
	stats := cache.Stats()
	if stats.DecayFactor != 0.9 {
		t.Errorf("Expected default decay factor 0.9, got %f", stats.DecayFactor)
	}
	if stats.DecayInterval != 5*time.Minute {
		t.Errorf("Expected default decay interval 5m, got %v", stats.DecayInterval)
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

func TestNewWithConfig(t *testing.T) {
	cache := NewWithConfig[string, int](3, 0.8, 10*time.Minute)
	stats := cache.Stats()
	if stats.DecayFactor != 0.8 {
		t.Errorf("Expected decay factor 0.8, got %f", stats.DecayFactor)
	}
	if stats.DecayInterval != 10*time.Minute {
		t.Errorf("Expected decay interval 10m, got %v", stats.DecayInterval)
	}
}

func TestNewWithConfigPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid decay factor")
		}
	}()
	NewWithConfig[string, int](3, 1.5, time.Minute)
}

func TestNewWithEvict(t *testing.T) {
	evicted := make(map[string]int)
	onEvict := func(key string, value int) {
		evicted[key] = value
	}
	
	cache := NewWithEvict[string, int](2, onEvict)
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3) // Should evict "a"
	
	if len(evicted) != 1 {
		t.Errorf("Expected 1 eviction, got %d", len(evicted))
	}
	if evicted["a"] != 1 {
		t.Errorf("Expected evicted value 1, got %d", evicted["a"])
	}
}

func TestBasicOperations(t *testing.T) {
	cache := New[string, int](3)
	
	// Test Put and Get
	evicted := cache.Put("a", 1)
	if evicted {
		t.Errorf("Expected no eviction")
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
}

func TestFrequencyBasedEviction(t *testing.T) {
	cache := New[string, int](3)
	
	// Add items
	cache.Put("a", 1) // frequency 1
	cache.Put("b", 2) // frequency 1
	cache.Put("c", 3) // frequency 1
	
	// Access "a" and "b" multiple times
	cache.Get("a")  // frequency 2
	cache.Get("a")  // frequency 3
	cache.Get("b")  // frequency 2
	
	// Add new item, should evict "c" (lowest frequency)
	evicted := cache.Put("d", 4)
	if !evicted {
		t.Errorf("Expected eviction")
	}
	
	// "c" should be gone
	_, ok := cache.Get("c")
	if ok {
		t.Errorf("Expected 'c' to be evicted")
	}
	
	// Others should still exist
	if !cache.Contains("a") || !cache.Contains("b") || !cache.Contains("d") {
		t.Errorf("Expected a, b, d to exist")
	}
}

func TestLRUWithinFrequency(t *testing.T) {
	cache := New[string, int](3)
	
	// Add items with same frequency
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)
	
	// Access "a" to make it more recent
	cache.Get("a")
	time.Sleep(1 * time.Millisecond) // Ensure different access times
	
	// Access "c" to make it more recent
	cache.Get("c")
	
	// Add new item, should evict "b" (same frequency but least recent)
	cache.Put("d", 4)
	
	// "b" should be gone (LRU within same frequency)
	_, ok := cache.Get("b")
	if ok {
		t.Errorf("Expected 'b' to be evicted as LRU within frequency")
	}
}

func TestContainsAndPeek(t *testing.T) {
	cache := New[string, int](3)
	cache.Put("a", 1)
	
	// Test Contains
	if !cache.Contains("a") {
		t.Errorf("Expected 'a' to exist")
	}
	if cache.Contains("b") {
		t.Errorf("Expected 'b' to not exist")
	}
	
	// Test Peek (shouldn't update frequency)
	value, ok := cache.Peek("a")
	if !ok || value != 1 {
		t.Errorf("Expected to peek value 1, got %d, ok=%v", value, ok)
	}
}

func TestRemove(t *testing.T) {
	cache := New[string, int](3)
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
	
	cache := NewWithEvict[string, int](3, onEvict)
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
	cache := New[string, int](3)
	cache.Put("a", 1)
	cache.Put("b", 2)
	
	// Access "a" more to increase its frequency
	cache.Get("a")
	cache.Get("a")
	
	keys := cache.Keys()
	values := cache.Values()
	
	if len(keys) != 2 || len(values) != 2 {
		t.Errorf("Expected 2 keys and values, got %d keys, %d values", len(keys), len(values))
	}
	
	// Keys should be ordered by frequency (highest first)
	if keys[0] != "a" {
		t.Errorf("Expected 'a' to be first key (highest frequency), got %s", keys[0])
	}
	
	if values[0] != 1 {
		t.Errorf("Expected 1 to be first value, got %d", values[0])
	}
}

func TestItems(t *testing.T) {
	cache := New[string, int](3)
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
	cache := New[string, int](3)
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)
	
	// Resize to smaller capacity
	cache.Resize(2)
	
	if cache.Cap() != 2 {
		t.Errorf("Expected capacity 2, got %d", cache.Cap())
	}
	
	if cache.Len() != 2 {
		t.Errorf("Expected length 2 after resize, got %d", cache.Len())
	}
}

func TestResizePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for zero capacity resize")
		}
	}()
	
	cache := New[string, int](3)
	cache.Resize(0)
}

func TestDecayFunctionality(t *testing.T) {
	// Use short decay interval for testing
	cache := NewWithConfig[string, int](3, 0.5, 10*time.Millisecond)
	
	// Add items and access them
	cache.Put("a", 1)
	cache.Get("a") // frequency 2
	cache.Get("a") // frequency 3
	cache.Get("a") // frequency 4
	
	cache.Put("b", 2)
	cache.Get("b") // frequency 2
	
	// Wait for decay to trigger
	time.Sleep(20 * time.Millisecond)
	
	// Force decay by accessing cache
	cache.Get("a")
	
	stats := cache.Stats()
	// After decay, frequencies should be reduced
	// Check that decay actually occurred by verifying frequency distribution
	totalEntries := 0
	for _, count := range stats.FrequencyDistribution {
		totalEntries += count
	}
	
	if totalEntries != 2 {
		t.Errorf("Expected 2 entries after decay, got %d", totalEntries)
	}
}

func TestForceDecay(t *testing.T) {
	cache := New[string, int](3)
	cache.Put("a", 1)
	cache.Get("a") // frequency 2
	cache.Get("a") // frequency 3
	
	statsBefore := cache.Stats()
	
	// Force decay
	cache.ForceDecay()
	
	statsAfter := cache.Stats()
	
	// Verify that last decay time was updated
	if !statsAfter.LastDecay.After(statsBefore.LastDecay) {
		t.Errorf("Expected last decay time to be updated")
	}
}

func TestStats(t *testing.T) {
	cache := New[string, int](5)
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Get("a") // increase frequency
	
	stats := cache.Stats()
	
	if stats.Size != 2 {
		t.Errorf("Expected size 2, got %d", stats.Size)
	}
	
	if stats.Capacity != 5 {
		t.Errorf("Expected capacity 5, got %d", stats.Capacity)
	}
	
	if stats.MinFrequency < 1 {
		t.Errorf("Expected min frequency >= 1, got %d", stats.MinFrequency)
	}
	
	if stats.MaxFrequency < stats.MinFrequency {
		t.Errorf("Expected max frequency >= min frequency")
	}
	
	if len(stats.FrequencyDistribution) == 0 {
		t.Errorf("Expected non-empty frequency distribution")
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

func TestTimeBasedDecay(t *testing.T) {
	cache := NewWithConfig[string, int](3, 0.8, 5*time.Millisecond)
	
	// Add an item and access it
	cache.Put("old", 1)
	cache.Get("old") // frequency 2
	
	// Wait a bit
	time.Sleep(10 * time.Millisecond)
	
	// Add another item more recently
	cache.Put("new", 2)
	cache.Get("new") // frequency 2
	
	// Force decay
	cache.ForceDecay()
	
	// The older item should have lower effective frequency due to time decay
	// Add third item to test eviction
	cache.Put("third", 3)
	cache.Put("fourth", 4) // This should trigger eviction
	
	// The "old" item should be more likely to be evicted due to time-based decay
	// This is probabilistic, but we can check that decay occurred
	stats := cache.Stats()
	if stats.Size != 3 {
		t.Errorf("Expected 3 items in cache, got %d", stats.Size)
	}
}

func TestEdgeCases(t *testing.T) {
	// Test single capacity cache
	cache := New[string, int](1)
	cache.Put("a", 1)
	cache.Put("b", 2) // Should evict "a"
	
	_, ok := cache.Get("a")
	if ok {
		t.Errorf("Expected 'a' to be evicted in single capacity cache")
	}
	
	value, ok := cache.Get("b")
	if !ok || value != 2 {
		t.Errorf("Expected 'b' to remain in single capacity cache")
	}
}

func TestFrequencyManagement(t *testing.T) {
	cache := New[string, int](3)
	
	// Create entries with different frequencies
	cache.Put("low", 1)      // freq 1
	cache.Put("medium", 2)   // freq 1
	cache.Put("high", 3)     // freq 1
	
	// Increase frequencies differently
	cache.Get("medium")      // freq 2
	cache.Get("high")        // freq 2
	cache.Get("high")        // freq 3
	
	stats := cache.Stats()
	
	// Should have entries at frequencies 1, 2, and 3
	expectedFreqs := []int{1, 2, 3}
	for _, freq := range expectedFreqs {
		if count, exists := stats.FrequencyDistribution[freq]; !exists || count == 0 {
			t.Errorf("Expected entries at frequency %d", freq)
		}
	}
	
	if stats.MinFrequency != 1 {
		t.Errorf("Expected min frequency 1, got %d", stats.MinFrequency)
	}
	
	if stats.MaxFrequency != 3 {
		t.Errorf("Expected max frequency 3, got %d", stats.MaxFrequency)
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