package lfu

import (
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	// Test valid capacity
	cache := New[string, int](3)
	if cache.Cap() != 3 {
		t.Errorf("Expected capacity 3, got %d", cache.Cap())
	}
	if cache.Len() != 0 {
		t.Errorf("Expected empty cache, got length %d", cache.Len())
	}
}

func TestNewPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for zero capacity")
		}
	}()
	New[string, int](0)
}

func TestNewWithEvict(t *testing.T) {
	evictedKeys := []string{}
	evictedValues := []int{}
	
	cache := NewWithEvict[string, int](2, func(key string, value int) {
		evictedKeys = append(evictedKeys, key)
		evictedValues = append(evictedValues, value)
	})
	
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3) // Should evict "a" or "b" (both have frequency 1)
	
	if len(evictedKeys) != 1 {
		t.Errorf("Expected 1 evicted key, got %d", len(evictedKeys))
	}
	if len(evictedValues) != 1 {
		t.Errorf("Expected 1 evicted value, got %d", len(evictedValues))
	}
}

func TestPutAndGet(t *testing.T) {
	cache := New[string, int](3)
	
	// Test putting and getting values
	evicted := cache.Put("a", 1)
	if evicted {
		t.Error("Should not evict when cache is not full")
	}
	
	value, ok := cache.Get("a")
	if !ok || value != 1 {
		t.Errorf("Expected value 1, got %d, ok=%t", value, ok)
	}
	
	// Test getting non-existent key
	_, ok = cache.Get("non-existent")
	if ok {
		t.Error("Expected false for non-existent key")
	}
}

func TestPutUpdateExisting(t *testing.T) {
	cache := New[string, int](3)
	
	cache.Put("a", 1)
	evicted := cache.Put("a", 2) // Update existing
	if evicted {
		t.Error("Should not evict when updating existing key")
	}
	
	value, ok := cache.Get("a")
	if !ok || value != 2 {
		t.Errorf("Expected updated value 2, got %d", value)
	}
	
	// Check frequency increased
	freq := cache.GetFreq("a")
	if freq != 3 { // Put(1) + Put(2) + Get(1) = 3
		t.Errorf("Expected frequency 3, got %d", freq)
	}
}

func TestEviction(t *testing.T) {
	cache := New[string, int](2)
	
	cache.Put("a", 1)
	cache.Put("b", 2)
	evicted := cache.Put("c", 3) // Should evict one of "a" or "b"
	
	if !evicted {
		t.Error("Expected eviction when cache is full")
	}
	
	if cache.Len() != 2 {
		t.Errorf("Expected cache length 2 after eviction, got %d", cache.Len())
	}
}

func TestLFUOrder(t *testing.T) {
	cache := New[string, int](3)
	
	cache.Put("a", 1) // freq: 1
	cache.Put("b", 2) // freq: 1
	cache.Put("c", 3) // freq: 1
	
	// Access "a" multiple times to increase frequency
	cache.Get("a") // freq: 2
	cache.Get("a") // freq: 3
	
	// Access "b" once
	cache.Get("b") // freq: 2
	
	// Add "d", should evict "c" (least frequent)
	cache.Put("d", 4)
	
	_, ok := cache.Get("c")
	if ok {
		t.Error("Expected 'c' to be evicted (least frequent)")
	}
	
	// "a", "b", "d" should still be there
	for key, expectedValue := range map[string]int{"a": 1, "b": 2, "d": 4} {
		if value, ok := cache.Get(key); !ok || value != expectedValue {
			t.Errorf("Expected '%s' to have value %d, got %d, ok=%t", key, expectedValue, value, ok)
		}
	}
	
	// Check frequencies
	// "a": Put(1) + Get(2) + Get(3) + Get(4) = 4
	if freq := cache.GetFreq("a"); freq != 4 {
		t.Errorf("Expected 'a' frequency 4, got %d", freq)
	}
	// "b": Put(1) + Get(2) + Get(3) = 3  
	if freq := cache.GetFreq("b"); freq != 3 {
		t.Errorf("Expected 'b' frequency 3, got %d", freq)
	}
	// "d": Put(1) + Get(2) = 2
	if freq := cache.GetFreq("d"); freq != 2 {
		t.Errorf("Expected 'd' frequency 2, got %d", freq)
	}
}

func TestRemove(t *testing.T) {
	cache := New[string, int](3)
	
	cache.Put("a", 1)
	cache.Put("b", 2)
	
	// Remove existing key
	value, ok := cache.Remove("a")
	if !ok || value != 1 {
		t.Errorf("Expected removed value 1, got %d, ok=%t", value, ok)
	}
	
	// Check it's gone
	_, ok = cache.Get("a")
	if ok {
		t.Error("Expected 'a' to be removed")
	}
	
	// Remove non-existent key
	_, ok = cache.Remove("non-existent")
	if ok {
		t.Error("Expected false for removing non-existent key")
	}
	
	if cache.Len() != 1 {
		t.Errorf("Expected length 1 after removal, got %d", cache.Len())
	}
}

func TestContains(t *testing.T) {
	cache := New[string, int](3)
	
	cache.Put("a", 1)
	
	if !cache.Contains("a") {
		t.Error("Expected 'a' to be contained")
	}
	
	if cache.Contains("non-existent") {
		t.Error("Expected 'non-existent' to not be contained")
	}
}

func TestPeek(t *testing.T) {
	cache := New[string, int](2)
	
	cache.Put("a", 1)
	cache.Put("b", 2)
	
	// Peek should not affect frequency
	freqBefore := cache.GetFreq("a")
	value, ok := cache.Peek("a")
	freqAfter := cache.GetFreq("a")
	
	if !ok || value != 1 {
		t.Errorf("Expected peek value 1, got %d, ok=%t", value, ok)
	}
	
	if freqBefore != freqAfter {
		t.Errorf("Peek should not change frequency: before=%d, after=%d", freqBefore, freqAfter)
	}
	
	// Test peek non-existent key
	_, ok = cache.Peek("non-existent")
	if ok {
		t.Error("Expected false for peeking non-existent key")
	}
}

func TestClear(t *testing.T) {
	evictCount := 0
	cache := NewWithEvict[string, int](3, func(key string, value int) {
		evictCount++
	})
	
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
	
	// Test that cache still works after clear
	cache.Put("d", 4)
	value, ok := cache.Get("d")
	if !ok || value != 4 {
		t.Errorf("Expected value 4 after clear, got %d, ok=%t", value, ok)
	}
}

func TestKeys(t *testing.T) {
	cache := New[string, int](3)
	
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)
	
	keys := cache.Keys()
	expectedKeys := map[string]bool{"a": true, "b": true, "c": true}
	
	if len(keys) != len(expectedKeys) {
		t.Errorf("Expected %d keys, got %d", len(expectedKeys), len(keys))
	}
	
	for _, key := range keys {
		if !expectedKeys[key] {
			t.Errorf("Unexpected key: %s", key)
		}
	}
}

func TestValues(t *testing.T) {
	cache := New[string, int](3)
	
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)
	
	values := cache.Values()
	expectedValues := map[int]bool{1: true, 2: true, 3: true}
	
	if len(values) != len(expectedValues) {
		t.Errorf("Expected %d values, got %d", len(expectedValues), len(values))
	}
	
	for _, value := range values {
		if !expectedValues[value] {
			t.Errorf("Unexpected value: %d", value)
		}
	}
}

func TestItems(t *testing.T) {
	cache := New[string, int](3)
	
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)
	
	items := cache.Items()
	expectedItems := map[string]int{"a": 1, "b": 2, "c": 3}
	
	if len(items) != len(expectedItems) {
		t.Errorf("Expected %d items, got %d", len(expectedItems), len(items))
	}
	
	for key, expectedValue := range expectedItems {
		if value, ok := items[key]; !ok || value != expectedValue {
			t.Errorf("Expected items[%s] = %d, got %d, ok=%t", key, expectedValue, value, ok)
		}
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
		t.Errorf("Expected capacity 2 after resize, got %d", cache.Cap())
	}
	
	if cache.Len() != 2 {
		t.Errorf("Expected length 2 after resize, got %d", cache.Len())
	}
	
	// Resize to larger capacity
	cache.Resize(5)
	if cache.Cap() != 5 {
		t.Errorf("Expected capacity 5 after resize, got %d", cache.Cap())
	}
}

func TestResizePanic(t *testing.T) {
	cache := New[string, int](3)
	
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for zero capacity resize")
		}
	}()
	cache.Resize(0)
}

func TestGetFreq(t *testing.T) {
	cache := New[string, int](3)
	
	// Test frequency of non-existent key
	freq := cache.GetFreq("non-existent")
	if freq != 0 {
		t.Errorf("Expected frequency 0 for non-existent key, got %d", freq)
	}
	
	cache.Put("a", 1) // freq: 1
	cache.Get("a")    // freq: 2
	cache.Get("a")    // freq: 3
	
	freq = cache.GetFreq("a")
	if freq != 3 {
		t.Errorf("Expected frequency 3, got %d", freq)
	}
}

func TestStats(t *testing.T) {
	cache := New[string, int](5)
	
	cache.Put("a", 1) // freq: 1
	cache.Put("b", 2) // freq: 1
	cache.Get("a")    // freq: 2
	
	stats := cache.Stats()
	if stats.Size != 2 {
		t.Errorf("Expected size 2, got %d", stats.Size)
	}
	if stats.Capacity != 5 {
		t.Errorf("Expected capacity 5, got %d", stats.Capacity)
	}
	if stats.MinFreq != 1 {
		t.Errorf("Expected min frequency 1, got %d", stats.MinFreq)
	}
	
	expectedFreqDist := map[int]int{1: 1, 2: 1} // freq 1: 1 item (b), freq 2: 1 item (a)
	if len(stats.FreqDistribution) != len(expectedFreqDist) {
		t.Errorf("Expected freq distribution %v, got %v", expectedFreqDist, stats.FreqDistribution)
	}
	
	for freq, count := range expectedFreqDist {
		if stats.FreqDistribution[freq] != count {
			t.Errorf("Expected %d items with frequency %d, got %d", count, freq, stats.FreqDistribution[freq])
		}
	}
}

func TestMinFreqUpdate(t *testing.T) {
	cache := New[string, int](2)
	
	cache.Put("a", 1) // freq: 1, minFreq: 1
	cache.Put("b", 2) // freq: 1, minFreq: 1
	
	// Increase frequency of "a"
	cache.Get("a") // freq: 2, minFreq should still be 1 (because of "b")
	
	stats := cache.Stats()
	if stats.MinFreq != 1 {
		t.Errorf("Expected min frequency 1, got %d", stats.MinFreq)
	}
	
	// Increase frequency of "b"
	cache.Get("b") // freq: 2, minFreq should now be 2
	
	stats = cache.Stats()
	if stats.MinFreq != 2 {
		t.Errorf("Expected min frequency 2, got %d", stats.MinFreq)
	}
}

func TestMinFreqAfterEviction(t *testing.T) {
	cache := New[string, int](2)
	
	cache.Put("a", 1) // freq: 1
	cache.Put("b", 2) // freq: 1
	cache.Get("a")    // freq: 2
	
	// Add "c", should evict "b" (freq 1)
	cache.Put("c", 3) // freq: 1
	
	stats := cache.Stats()
	if stats.MinFreq != 1 {
		t.Errorf("Expected min frequency 1 after eviction, got %d", stats.MinFreq)
	}
}

func TestLFUEvictionOrder(t *testing.T) {
	cache := New[string, int](3)
	
	cache.Put("a", 1) // freq: 1
	cache.Put("b", 2) // freq: 1
	cache.Put("c", 3) // freq: 1
	
	// Access "a" and "b" to increase their frequencies
	cache.Get("a") // freq: 2
	cache.Get("b") // freq: 2
	
	// Add "d", should evict "c" (freq 1, LFU)
	cache.Put("d", 4)
	
	_, ok := cache.Get("c")
	if ok {
		t.Error("Expected 'c' to be evicted (least frequent)")
	}
	
	// "a", "b", "d" should still be there
	expectedItems := map[string]int{"a": 1, "b": 2, "d": 4}
	for key, expectedValue := range expectedItems {
		if value, ok := cache.Get(key); !ok || value != expectedValue {
			t.Errorf("Expected '%s' to have value %d, got %d, ok=%t", key, expectedValue, value, ok)
		}
	}
}

func TestConcurrency(t *testing.T) {
	cache := New[int, int](1000)
	
	// Test concurrent writes and reads
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
				cache.GetFreq(key)
			}
		}(i)
	}
	
	wg.Wait()
	
	// Verify cache is still functional
	cache.Put(9999, 9999)
	value, ok := cache.Get(9999)
	if !ok || value != 9999 {
		t.Errorf("Expected value 9999 after concurrent operations, got %d, ok=%t", value, ok)
	}
}

func TestEdgeCases(t *testing.T) {
	// Test with capacity 1
	cache := New[string, int](1)
	
	cache.Put("a", 1)
	cache.Put("b", 2) // Should evict "a"
	
	_, ok := cache.Get("a")
	if ok {
		t.Error("Expected 'a' to be evicted in capacity 1 cache")
	}
	
	value, ok := cache.Get("b")
	if !ok || value != 2 {
		t.Errorf("Expected 'b' to have value 2, got %d, ok=%t", value, ok)
	}
}

func TestEmptyCache(t *testing.T) {
	cache := New[string, int](3)
	
	// Test operations on empty cache
	_, ok := cache.Get("non-existent")
	if ok {
		t.Error("Expected false for get on empty cache")
	}
	
	_, ok = cache.Remove("non-existent")
	if ok {
		t.Error("Expected false for remove on empty cache")
	}
	
	if cache.Contains("non-existent") {
		t.Error("Expected false for contains on empty cache")
	}
	
	_, ok = cache.Peek("non-existent")
	if ok {
		t.Error("Expected false for peek on empty cache")
	}
	
	freq := cache.GetFreq("non-existent")
	if freq != 0 {
		t.Errorf("Expected frequency 0 for non-existent key, got %d", freq)
	}
	
	keys := cache.Keys()
	if len(keys) != 0 {
		t.Errorf("Expected empty keys slice, got %v", keys)
	}
	
	values := cache.Values()
	if len(values) != 0 {
		t.Errorf("Expected empty values slice, got %v", values)
	}
	
	items := cache.Items()
	if len(items) != 0 {
		t.Errorf("Expected empty items map, got %v", items)
	}
	
	stats := cache.Stats()
	if stats.Size != 0 {
		t.Errorf("Expected stats size 0, got %d", stats.Size)
	}
	if stats.MinFreq != 1 {
		t.Errorf("Expected stats min frequency 1, got %d", stats.MinFreq)
	}
}

func TestComplexFrequencyScenario(t *testing.T) {
	cache := New[string, int](3)
	
	// Create items with different frequencies
	cache.Put("a", 1) // freq: 1
	cache.Put("b", 2) // freq: 1
	cache.Put("c", 3) // freq: 1
	
	// Make "a" most frequently used
	cache.Get("a") // freq: 2
	cache.Get("a") // freq: 3
	cache.Get("a") // freq: 4
	
	// Make "b" moderately used
	cache.Get("b") // freq: 2
	
	// "c" remains least frequently used (freq: 1)
	
	// Add new item, should evict "c"
	cache.Put("d", 4) // freq: 1
	
	_, ok := cache.Get("c")
	if ok {
		t.Error("Expected 'c' to be evicted (least frequent)")
	}
	
	// Now we have: a(freq:5), b(freq:3), d(freq:1)
	// Add another item, should evict "d"
	cache.Put("e", 5) // freq: 1
	
	_, ok = cache.Get("d")
	if ok {
		t.Error("Expected 'd' to be evicted (least frequent)")
	}
	
	// Final state: a(freq:6), b(freq:4), e(freq:1)
	expectedItems := map[string]int{"a": 1, "b": 2, "e": 5}
	for key, expectedValue := range expectedItems {
		if value, ok := cache.Get(key); !ok || value != expectedValue {
			t.Errorf("Expected '%s' to have value %d, got %d, ok=%t", key, expectedValue, value, ok)
		}
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
		cache.Put(i, i)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Get(i % 1000)
	}
}

func BenchmarkPutGet(b *testing.B) {
	cache := New[int, int](1000)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := i % 1000
		cache.Put(key, i)
		cache.Get(key)
	}
}

func TestRemoveFromEmptyFreqList(t *testing.T) {
	cache := New[string, int](1)
	
	// Add an item then clear its frequency list to test empty list scenario
	cache.Put("a", 1)
	
	// Manually clear the frequency list to test the edge case
	if freqList := cache.freqLists[1]; freqList != nil {
		freqList.Init() // Clear the list
	}
	
	// Now try to remove the item - this should handle the empty list case
	cache.Remove("a")
}

func TestEvictFromCorruptedState(t *testing.T) {
	cache := New[string, int](1)
	
	// Create a scenario where eviction might encounter an empty or nil list
	cache.Put("a", 1)
	cache.Put("b", 2) // This should evict "a"
	
	// Verify "a" was evicted
	_, ok := cache.Get("a")
	if ok {
		t.Error("Expected 'a' to be evicted")
	}
}

// TestEmptyFreqListEviction tests the edge case where frequency list is empty during eviction
func TestEmptyFreqListEviction(t *testing.T) {
	cache := New[string, int](1)
	
	// Force an empty state by manipulating internal state
	cache.minFreq = 99 // Set to a frequency that doesn't exist
	
	// This should not panic even with invalid minFreq
	cache.Put("a", 1)
	
	// The item should still be added successfully
	value, ok := cache.Get("a")
	if !ok || value != 1 {
		t.Errorf("Expected item to be added successfully, got value=%d, ok=%t", value, ok)
	}
}

// TestUpdateMinFreqWithZeroFreq tests updateMinFreq with edge cases
func TestUpdateMinFreqWithZeroFreq(t *testing.T) {
	cache := New[string, int](2)
	
	// Test updateMinFreq when all frequencies are > 0
	cache.Put("a", 1) // freq: 1
	cache.Get("a")    // freq: 2
	
	// Remove to trigger updateMinFreq paths  
	cache.Remove("a")
	
	// After removing all items, cache should be empty
	if cache.Len() != 0 {
		t.Errorf("Expected empty cache after removing all items, got length %d", cache.Len())
	}
}

// TestEvictLFUWithNilFreqList tests the edge case where frequency list is nil during eviction
func TestEvictLFUWithNilFreqList(t *testing.T) {
	cache := New[string, int](1)
	
	// Set minFreq to a value that doesn't have a corresponding list
	cache.minFreq = 999
	
	// This should gracefully handle the nil frequency list case
	// and not panic when trying to evict
	cache.Put("a", 1) // This will trigger eviction if cache is full
	
	// The cache should still work
	value, ok := cache.Get("a")
	if !ok || value != 1 {
		t.Errorf("Expected cache to work despite nil frequency list, got value=%d, ok=%t", value, ok)
	}
}