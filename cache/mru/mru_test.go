package mru

import (
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
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

func TestPutAndGet(t *testing.T) {
	cache := New[string, int](3)
	
	evicted := cache.Put("a", 1)
	if evicted {
		t.Error("Should not evict when cache is not full")
	}
	
	value, ok := cache.Get("a")
	if !ok || value != 1 {
		t.Errorf("Expected value 1, got %d, ok=%t", value, ok)
	}
}

func TestMRUEviction(t *testing.T) {
	cache := New[string, int](2)
	
	cache.Put("a", 1)
	cache.Put("b", 2)
	
	// Access "a" to make it most recently used
	cache.Get("a")
	
	// Add "c" - should evict "a" (most recently used)
	evicted := cache.Put("c", 3)
	if !evicted {
		t.Error("Expected eviction when cache is full")
	}
	
	// "a" should be evicted (MRU behavior)
	_, ok := cache.Get("a")
	if ok {
		t.Error("Expected 'a' to be evicted (most recently used)")
	}
	
	// "b" and "c" should still be there
	if value, ok := cache.Get("b"); !ok || value != 2 {
		t.Errorf("Expected 'b' to have value 2, got %d, ok=%t", value, ok)
	}
	if value, ok := cache.Get("c"); !ok || value != 3 {
		t.Errorf("Expected 'c' to have value 3, got %d, ok=%t", value, ok)
	}
}

func TestRemove(t *testing.T) {
	cache := New[string, int](3)
	
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
	
	value, ok := cache.Peek("a")
	if !ok || value != 1 {
		t.Errorf("Expected peek value 1, got %d, ok=%t", value, ok)
	}
	
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
}

func TestKeys(t *testing.T) {
	cache := New[string, int](3)
	
	cache.Put("c", 3)
	cache.Put("a", 1)
	cache.Put("b", 2)
	
	keys := cache.Keys()
	expectedKeys := []string{"b", "a", "c"} // Most to least recently used
	
	if len(keys) != len(expectedKeys) {
		t.Errorf("Expected %d keys, got %d", len(expectedKeys), len(keys))
	}
	
	for i, expectedKey := range expectedKeys {
		if i >= len(keys) || keys[i] != expectedKey {
			t.Errorf("Expected key[%d] = %s, got %s", i, expectedKey, keys[i])
		}
	}
}

func TestResize(t *testing.T) {
	cache := New[string, int](3)
	
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)
	
	cache.Resize(2)
	
	if cache.Cap() != 2 {
		t.Errorf("Expected capacity 2 after resize, got %d", cache.Cap())
	}
	
	if cache.Len() != 2 {
		t.Errorf("Expected length 2 after resize, got %d", cache.Len())
	}
	
	// Most recently used item "c" should be evicted
	_, ok := cache.Get("c")
	if ok {
		t.Error("Expected 'c' to be evicted after resize (most recently used)")
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

func TestStats(t *testing.T) {
	cache := New[string, int](5)
	
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
	cache := New[int, int](1000)
	
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
	
	cache.Put(9999, 9999)
	value, ok := cache.Get(9999)
	if !ok || value != 9999 {
		t.Errorf("Expected value 9999 after concurrent operations, got %d, ok=%t", value, ok)
	}
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
	cache.Put("c", 3) // Should evict most recently used
	
	if len(evictedKeys) != 1 {
		t.Errorf("Expected 1 evicted key, got %d", len(evictedKeys))
	}
	if len(evictedValues) != 1 {
		t.Errorf("Expected 1 evicted value, got %d", len(evictedValues))
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
	
	for i := 0; i < 1000; i++ {
		cache.Put(i, i)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Get(i % 1000)
	}
}

func TestValues(t *testing.T) {
	cache := New[string, int](3)
	
	// Test empty cache
	values := cache.Values()
	if len(values) != 0 {
		t.Errorf("Expected empty values slice, got length %d", len(values))
	}
	
	// Add items
	cache.Put("c", 3)
	cache.Put("a", 1) 
	cache.Put("b", 2)
	
	values = cache.Values()
	expectedValues := []int{2, 1, 3} // Most to least recently used
	
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
	cache := New[string, int](3)
	
	// Test empty cache
	items := cache.Items()
	if len(items) != 0 {
		t.Errorf("Expected empty items map, got length %d", len(items))
	}
	
	// Add items
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)
	
	items = cache.Items()
	expectedItems := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	
	if len(items) != len(expectedItems) {
		t.Errorf("Expected %d items, got %d", len(expectedItems), len(items))
	}
	
	for key, expectedValue := range expectedItems {
		if value, exists := items[key]; !exists || value != expectedValue {
			t.Errorf("Expected items[%s] = %d, got %d, exists=%t", key, expectedValue, value, exists)
		}
	}
}

func TestPutUpdate(t *testing.T) {
	cache := New[string, int](3)
	
	// Add initial item
	evicted := cache.Put("a", 1)
	if evicted {
		t.Error("Should not evict when cache is not full")
	}
	
	// Update existing item
	evicted = cache.Put("a", 10)
	if evicted {
		t.Error("Should not evict when updating existing item")
	}
	
	// Verify the value was updated
	value, ok := cache.Get("a")
	if !ok || value != 10 {
		t.Errorf("Expected updated value 10, got %d, ok=%t", value, ok)
	}
	
	// Verify cache length didn't change
	if cache.Len() != 1 {
		t.Errorf("Expected length 1 after update, got %d", cache.Len())
	}
}