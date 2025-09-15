package fbr

import (
	"fmt"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	cache, err := New[string, int](5)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	if cache.Cap() != 5 {
		t.Errorf("Expected capacity 5, got %d", cache.Cap())
	}
	if cache.Len() != 0 {
		t.Errorf("Expected empty cache, got length %d", cache.Len())
	}

	stats := cache.Stats()
	if stats.MinFrequency != 1 {
		t.Errorf("Expected min frequency 1, got %d", stats.MinFrequency)
	}
	if stats.MaxFrequency != 1 {
		t.Errorf("Expected max frequency 1, got %d", stats.MaxFrequency)
	}
}

func TestNewError(t *testing.T) {
	_, err := New[string, int](0)
	if err == nil {
		t.Error("Expected error for zero capacity")
	}
}

func TestPutAndGet(t *testing.T) {
	cache, err := New[string, int](5)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	evicted := cache.Put("a", 1)
	if evicted {
		t.Error("Should not evict when cache is not full")
	}

	value, ok := cache.Get("a")
	if !ok || value != 1 {
		t.Errorf("Expected value 1, got %d, ok=%t", value, ok)
	}

	// After get, frequency should be incremented
	stats := cache.Stats()
	if stats.MaxFrequency != 2 {
		t.Errorf("Expected max frequency 2 after get, got %d", stats.MaxFrequency)
	}
}

func TestFrequencyBasedEviction(t *testing.T) {
	cache, err := New[string, int](3)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Add items with different access patterns
	cache.Put("low", 1)    // frequency 1
	cache.Put("medium", 2) // frequency 1
	cache.Put("high", 3)   // frequency 1

	// Access "high" multiple times
	for i := 0; i < 5; i++ {
		cache.Get("high")
	}

	// Access "medium" a few times
	for i := 0; i < 2; i++ {
		cache.Get("medium")
	}

	// "low" stays at frequency 1

	// Add new item - should evict "low" (lowest frequency)
	evicted := cache.Put("new", 4)
	if !evicted {
		t.Error("Expected eviction when cache is full")
	}

	// "low" should be evicted
	_, ok := cache.Get("low")
	if ok {
		t.Error("Expected 'low' to be evicted (lowest frequency)")
	}

	// "high", "medium", and "new" should still be there
	if value, ok := cache.Get("high"); !ok || value != 3 {
		t.Errorf("Expected 'high' to have value 3, got %d, ok=%t", value, ok)
	}
	if value, ok := cache.Get("medium"); !ok || value != 2 {
		t.Errorf("Expected 'medium' to have value 2, got %d, ok=%t", value, ok)
	}
	if value, ok := cache.Get("new"); !ok || value != 4 {
		t.Errorf("Expected 'new' to have value 4, got %d, ok=%t", value, ok)
	}
}

func TestUpdateExisting(t *testing.T) {
	cache, err := New[string, int](5)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Get("a") // increment frequency to 2

	evicted := cache.Put("a", 10) // Update existing
	if evicted {
		t.Error("Should not evict when updating existing key")
	}

	value, ok := cache.Get("a")
	if !ok || value != 10 {
		t.Errorf("Expected updated value 10, got %d, ok=%t", value, ok)
	}

	// Frequency should be incremented again
	stats := cache.Stats()
	if stats.MaxFrequency < 3 {
		t.Errorf("Expected max frequency >= 3 after update and get, got %d", stats.MaxFrequency)
	}
}

func TestRemove(t *testing.T) {
	cache, err := New[string, int](5)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("b", 2)

	value, ok := cache.Remove("a")
	if !ok || value != 1 {
		t.Errorf("Expected removed value 1, got %d, ok=%t", value, ok)
	}

	_, ok = cache.Get("a")
	if ok {
		t.Error("Expected 'a' to be removed")
	}

	_, ok = cache.Remove("non-existent")
	if ok {
		t.Error("Expected false for removing non-existent key")
	}
}

func TestContains(t *testing.T) {
	cache, err := New[string, int](5)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)

	if !cache.Contains("a") {
		t.Error("Expected 'a' to be contained")
	}

	if cache.Contains("non-existent") {
		t.Error("Expected 'non-existent' to not be contained")
	}
}

func TestPeek(t *testing.T) {
	cache, err := New[string, int](5)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)

	value, ok := cache.Peek("a")
	if !ok || value != 1 {
		t.Errorf("Expected peek value 1, got %d, ok=%t", value, ok)
	}

	// Peek should not increment frequency
	stats := cache.Stats()
	if stats.MaxFrequency != 1 {
		t.Errorf("Peek should not increment frequency, expected max freq 1, got %d", stats.MaxFrequency)
	}

	_, ok = cache.Peek("non-existent")
	if ok {
		t.Error("Expected false for peeking non-existent key")
	}
}

func TestClear(t *testing.T) {
	evictCount := 0
	cache, err := NewWithEvict[string, int](5, func(key string, value int) {
		evictCount++
	})
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	cache.Clear()

	if cache.Len() != 0 {
		t.Errorf("Expected empty cache after clear, got length %d", cache.Len())
	}

	if evictCount != 3 {
		t.Errorf("Expected 3 evictions on clear, got %d", evictCount)
	}

	stats := cache.Stats()
	if len(stats.FrequencyDistribution) != 0 {
		t.Error("Expected empty frequency distribution after clear")
	}
}

func TestKeys(t *testing.T) {
	cache, err := New[string, int](5)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("low", 1)    // frequency 1
	cache.Put("medium", 2) // frequency 1
	cache.Put("high", 3)   // frequency 1

	// Create different frequencies
	cache.Get("high")   // frequency 2
	cache.Get("high")   // frequency 3
	cache.Get("medium") // frequency 2

	keys := cache.Keys()

	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}

	// Keys should be ordered by frequency (highest first)
	// "high" should be first (freq 3), then "medium" (freq 2), then "low" (freq 1)
	if keys[0] != "high" {
		t.Errorf("Expected first key to be 'high', got %s", keys[0])
	}
}

func TestValues(t *testing.T) {
	cache, err := New[string, int](5)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	values := cache.Values()

	if len(values) != 3 {
		t.Errorf("Expected 3 values, got %d", len(values))
	}

	// Check that all values are present
	valueSet := make(map[int]bool)
	for _, v := range values {
		valueSet[v] = true
	}

	for _, expected := range []int{1, 2, 3} {
		if !valueSet[expected] {
			t.Errorf("Expected value %d to be present", expected)
		}
	}
}

func TestItems(t *testing.T) {
	cache, err := New[string, int](5)
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
		t.Error("Expected items map to contain correct values")
	}
}

func TestResize(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Fill cache
	for i := 0; i < 8; i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
	}

	initialLen := cache.Len()

	// Resize to smaller capacity
	cache.Resize(5)

	if cache.Cap() != 5 {
		t.Errorf("Expected capacity 5 after resize, got %d", cache.Cap())
	}

	if cache.Len() > 5 {
		t.Errorf("Expected length <= 5 after resize, got %d", cache.Len())
	}

	if cache.Len() >= initialLen {
		t.Error("Expected some items to be evicted during resize")
	}
}

func TestResizeError(t *testing.T) {
	cache, err := New[string, int](5)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	err = cache.Resize(0)
	if err == nil {
		t.Error("Expected error for zero capacity resize")
	}
}

func TestStats(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Get("a") // increment frequency

	stats := cache.Stats()

	if stats.Size != 2 {
		t.Errorf("Expected size 2, got %d", stats.Size)
	}
	if stats.Capacity != 10 {
		t.Errorf("Expected capacity 10, got %d", stats.Capacity)
	}
	if stats.MinFrequency != 1 {
		t.Errorf("Expected min frequency 1, got %d", stats.MinFrequency)
	}
	if stats.MaxFrequency != 2 {
		t.Errorf("Expected max frequency 2, got %d", stats.MaxFrequency)
	}

	// Check frequency distribution
	if len(stats.FrequencyDistribution) == 0 {
		t.Error("Expected non-empty frequency distribution")
	}
}

func TestFrequencyDistribution(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Add items with different access patterns
	cache.Put("once", 1)
	cache.Put("twice", 2)
	cache.Put("thrice", 3)

	cache.Get("twice")  // freq 2
	cache.Get("thrice") // freq 2
	cache.Get("thrice") // freq 3

	stats := cache.Stats()

	// Should have: 1 item with freq 1, 1 item with freq 2, 1 item with freq 3
	if stats.FrequencyDistribution[1] != 1 {
		t.Errorf("Expected 1 item with frequency 1, got %d", stats.FrequencyDistribution[1])
	}
	if stats.FrequencyDistribution[2] != 1 {
		t.Errorf("Expected 1 item with frequency 2, got %d", stats.FrequencyDistribution[2])
	}
	if stats.FrequencyDistribution[3] != 1 {
		t.Errorf("Expected 1 item with frequency 3, got %d", stats.FrequencyDistribution[3])
	}
}

func TestLRUWithinFrequency(t *testing.T) {
	cache, err := New[string, int](3)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Add items all with same frequency (1)
	cache.Put("first", 1)
	cache.Put("second", 2)
	cache.Put("third", 3)

	// All should have frequency 1
	// When adding fourth item, "first" should be evicted (LRU within frequency 1)
	evicted := cache.Put("fourth", 4)
	if !evicted {
		t.Error("Expected eviction when cache is full")
	}

	// "first" should be evicted (least recently used among frequency 1)
	_, ok := cache.Get("first")
	if ok {
		t.Error("Expected 'first' to be evicted (LRU within frequency)")
	}

	// Others should still be there
	for _, key := range []string{"second", "third", "fourth"} {
		if !cache.Contains(key) {
			t.Errorf("Expected '%s' to still be in cache", key)
		}
	}
}

func TestMinFrequencyUpdate(t *testing.T) {
	cache, err := New[string, int](3)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1) // freq 1
	cache.Put("b", 2) // freq 1
	cache.Put("c", 3) // freq 1

	stats := cache.Stats()
	if stats.MinFrequency != 1 {
		t.Errorf("Expected min frequency 1, got %d", stats.MinFrequency)
	}

	// Increment frequency of all items
	cache.Get("a") // freq 2
	cache.Get("b") // freq 2
	cache.Get("c") // freq 2

	stats = cache.Stats()
	if stats.MinFrequency != 2 {
		t.Errorf("Expected min frequency 2 after incrementing all, got %d", stats.MinFrequency)
	}

	// Add new item
	cache.Put("new", 4) // freq 1, should evict one of the freq 2 items

	stats = cache.Stats()
	if stats.MinFrequency != 1 {
		t.Errorf("Expected min frequency 1 after adding new item, got %d", stats.MinFrequency)
	}
}

func TestConcurrency(t *testing.T) {
	cache, err := New[int, int](1000)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	var wg sync.WaitGroup
	numGoroutines := 10
	numOperations := 100

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				key := id*numOperations + j
				cache.Put(key, key*2)
				cache.Get(key)
				cache.Contains(key)
				cache.Peek(key)
			}
		}(i)
	}

	wg.Wait()

	// Test that cache still works correctly after concurrent operations
	cache.Put(9999, 9999)
	value, ok := cache.Get(9999)
	if !ok || value != 9999 {
		t.Errorf("Expected value 9999 after concurrent operations, got %d, ok=%t", value, ok)
	}
}

func TestNewWithEvict(t *testing.T) {
	evictedKeys := []string{}
	evictedValues := []int{}

	cache, err := NewWithEvict[string, int](2, func(key string, value int) {
		evictedKeys = append(evictedKeys, key)
		evictedValues = append(evictedValues, value)
	})
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3) // Should evict one item

	if len(evictedKeys) != 1 {
		t.Errorf("Expected 1 evicted key, got %d", len(evictedKeys))
	}
	if len(evictedValues) != 1 {
		t.Errorf("Expected 1 evicted value, got %d", len(evictedValues))
	}
}

func TestEmptyFrequencyListHandling(t *testing.T) {
	cache, err := New[string, int](2)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1) // freq 1
	cache.Put("b", 2) // freq 1

	// Increment frequency of "a"
	cache.Get("a") // freq 2

	// Now we have: "a" at freq 2, "b" at freq 1
	stats := cache.Stats()
	if stats.MinFrequency != 1 {
		t.Errorf("Expected min frequency 1, got %d", stats.MinFrequency)
	}

	// Add new item, should evict "b" (freq 1)
	cache.Put("c", 3)

	// Check that "b" was evicted
	if cache.Contains("b") {
		t.Error("Expected 'b' to be evicted")
	}

	// Min frequency should now be 1 (new item "c")
	stats = cache.Stats()
	if stats.MinFrequency != 1 {
		t.Errorf("Expected min frequency 1 after eviction, got %d", stats.MinFrequency)
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

	for i := 0; i < 1000; i++ {
		cache.Put(i, i)
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
		key := i % 1000
		if i%2 == 0 {
			cache.Put(key, i)
		} else {
			cache.Get(key)
		}
	}
}

func TestEvictLeastFrequentEdgeCase(t *testing.T) {
	cache, err := New[string, int](1)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Test evictLeastFrequent on empty cache
	// This should cover the return false case when no items can be evicted
	result := cache.evictLeastFrequent()
	if result {
		t.Errorf("Expected evictLeastFrequent to return false for empty cache")
	}

	// Also test with no evictable items due to empty frequency lists
	// This edge case happens when all frequency lists are empty but the loop still runs
	cache.Put("test", 1)
	cache.Remove("test") // This leaves empty frequency lists but might not reset minFreq correctly

	// Now try eviction again - should return false as no items exist to evict
	result = cache.evictLeastFrequent()
	if result {
		t.Errorf("Expected evictLeastFrequent to return false when no items exist to evict")
	}
}
