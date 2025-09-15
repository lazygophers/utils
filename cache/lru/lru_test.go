package lru

import (
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	// Test valid capacity
	cache, err := New[string, int](3)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	if cache.Cap() != 3 {
		t.Errorf("Expected capacity 3, got %d", cache.Cap())
	}
	if cache.Len() != 0 {
		t.Errorf("Expected empty cache, got length %d", cache.Len())
	}
}

func TestNewError(t *testing.T) {
	_, err := New[string, int](0)
	if err == nil {
		t.Error("Expected error for zero capacity")
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
	cache.Put("c", 3) // Should evict "a"
	
	if len(evictedKeys) != 1 || evictedKeys[0] != "a" {
		t.Errorf("Expected evicted key 'a', got %v", evictedKeys)
	}
	if len(evictedValues) != 1 || evictedValues[0] != 1 {
		t.Errorf("Expected evicted value 1, got %v", evictedValues)
	}
}

func TestPutAndGet(t *testing.T) {
	cache, err := New[string, int](3)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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
	cache, err := New[string, int](3)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	cache.Put("a", 1)
	evicted := cache.Put("a", 2) // Update existing
	if evicted {
		t.Error("Should not evict when updating existing key")
	}
	
	value, ok := cache.Get("a")
	if !ok || value != 2 {
		t.Errorf("Expected updated value 2, got %d", value)
	}
}

func TestEviction(t *testing.T) {
	cache, err := New[string, int](2)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	cache.Put("a", 1)
	cache.Put("b", 2)
	evicted := cache.Put("c", 3) // Should evict "a"
	
	if !evicted {
		t.Error("Expected eviction when cache is full")
	}
	
	// "a" should be evicted
	_, ok := cache.Get("a")
	if ok {
		t.Error("Expected 'a' to be evicted")
	}
	
	// "b" and "c" should still be there
	if value, ok := cache.Get("b"); !ok || value != 2 {
		t.Errorf("Expected 'b' to have value 2, got %d, ok=%t", value, ok)
	}
	if value, ok := cache.Get("c"); !ok || value != 3 {
		t.Errorf("Expected 'c' to have value 3, got %d, ok=%t", value, ok)
	}
}

func TestLRUOrder(t *testing.T) {
	cache, err := New[string, int](3)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)
	
	// Access "a" to make it most recently used
	cache.Get("a")
	
	// Add "d", should evict "b" (least recently used)
	cache.Put("d", 4)
	
	_, ok := cache.Get("b")
	if ok {
		t.Error("Expected 'b' to be evicted")
	}
	
	// "a", "c", "d" should still be there
	for key, expectedValue := range map[string]int{"a": 1, "c": 3, "d": 4} {
		if value, ok := cache.Get(key); !ok || value != expectedValue {
			t.Errorf("Expected '%s' to have value %d, got %d, ok=%t", key, expectedValue, value, ok)
		}
	}
}

func TestRemove(t *testing.T) {
	cache, err := New[string, int](3)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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
	cache, err := New[string, int](3)
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
	cache, err := New[string, int](2)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	cache.Put("a", 1)
	cache.Put("b", 2)
	
	// Peek should not affect LRU order
	value, ok := cache.Peek("a")
	if !ok || value != 1 {
		t.Errorf("Expected peek value 1, got %d, ok=%t", value, ok)
	}
	
	// Add "c", "a" should still be evicted despite peek
	cache.Put("c", 3)
	
	_, ok = cache.Get("a")
	if ok {
		t.Error("Expected 'a' to be evicted even after peek")
	}
	
	// Test peek non-existent key
	_, ok = cache.Peek("non-existent")
	if ok {
		t.Error("Expected false for peeking non-existent key")
	}
}

func TestClear(t *testing.T) {
	evictCount := 0
	cache, err := NewWithEvict[string, int](3, func(key string, value int) {
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
	
	// Test that cache still works after clear
	cache.Put("d", 4)
	value, ok := cache.Get("d")
	if !ok || value != 4 {
		t.Errorf("Expected value 4 after clear, got %d, ok=%t", value, ok)
	}
}

func TestKeys(t *testing.T) {
	cache, err := New[string, int](3)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	cache.Put("c", 3)
	cache.Put("a", 1)
	cache.Put("b", 2)
	
	// Access "a" to make it most recently used
	cache.Get("a")
	
	keys := cache.Keys()
	expectedKeys := []string{"a", "b", "c"} // Most to least recently used
	
	if len(keys) != len(expectedKeys) {
		t.Errorf("Expected %d keys, got %d", len(expectedKeys), len(keys))
	}
	
	for i, expectedKey := range expectedKeys {
		if i >= len(keys) || keys[i] != expectedKey {
			t.Errorf("Expected key[%d] = %s, got %s", i, expectedKey, keys[i])
		}
	}
}

func TestValues(t *testing.T) {
	cache, err := New[string, int](3)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	cache.Put("c", 3)
	cache.Put("a", 1)
	cache.Put("b", 2)
	
	// Access "a" to make it most recently used
	cache.Get("a")
	
	values := cache.Values()
	expectedValues := []int{1, 2, 3} // Most to least recently used
	
	if len(values) != len(expectedValues) {
		t.Errorf("Expected %d values, got %d", len(expectedValues), len(values))
	}
	
	for i, expectedValue := range expectedValues {
		if i >= len(values) || values[i] != expectedValue {
			t.Errorf("Expected value[%d] = %d, got %d", i, expectedValue, values[i])
		}
	}
}

func TestItems(t *testing.T) {
	cache, err := New[string, int](3)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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
	cache, err := New[string, int](3)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)
	
	// Resize to smaller capacity
	err = cache.Resize(2)
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}
	
	if cache.Cap() != 2 {
		t.Errorf("Expected capacity 2 after resize, got %d", cache.Cap())
	}
	
	if cache.Len() != 2 {
		t.Errorf("Expected length 2 after resize, got %d", cache.Len())
	}
	
	// "a" should be evicted (least recently used)
	_, ok := cache.Get("a")
	if ok {
		t.Error("Expected 'a' to be evicted after resize")
	}
	
	// Resize to larger capacity
	err = cache.Resize(5)
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}
	if cache.Cap() != 5 {
		t.Errorf("Expected capacity 5 after resize, got %d", cache.Cap())
	}
}

func TestResizeError(t *testing.T) {
	cache, err := New[string, int](3)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	err = cache.Resize(0)
	if err == nil {
		t.Error("Expected error for zero capacity resize")
	}
}

func TestStats(t *testing.T) {
	cache, err := New[string, int](5)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	cache.Put("a", 1)
	cache.Put("b", 2)
	
	stats := cache.Stats()
	if stats.Size != 2 {
		t.Errorf("Expected size 2, got %d", stats.Size)
	}
	if stats.Capacity != 5 {
		t.Errorf("Expected capacity 5, got %d", stats.Capacity)
	}
}

func TestConcurrency(t *testing.T) {
	cache, err := New[int, int](1000)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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
	cache, err := New[string, int](1)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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
	cache, err := New[string, int](3)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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
		cache.Put(i, i)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Get(i % 1000)
	}
}

func BenchmarkPutGet(b *testing.B) {
	cache, err := New[int, int](1000)
	if err != nil {
		b.Fatalf("Failed to create cache: %v", err)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := i % 1000
		cache.Put(key, i)
		cache.Get(key)
	}
}