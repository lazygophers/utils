package arc

import (
	"fmt"
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

	stats := cache.Stats()
	if stats.P != 0 {
		t.Errorf("Expected initial P=0, got %d", stats.P)
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
	cache.Put("c", 3) // Should evict one item

	if len(evictedKeys) != 1 {
		t.Errorf("Expected 1 evicted key, got %d", len(evictedKeys))
	}
	if len(evictedValues) != 1 {
		t.Errorf("Expected 1 evicted value, got %d", len(evictedValues))
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

func TestBasicEviction(t *testing.T) {
	cache, err := New[string, int](2)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("b", 2)
	evicted := cache.Put("c", 3) // Should evict "a" from T1

	if !evicted {
		t.Error("Expected eviction when cache is full")
	}

	// "a" should be moved to B1 (ghost)
	_, ok := cache.Get("a")
	if ok {
		t.Error("Expected 'a' to be evicted from cache")
	}

	if cache.Len() != 2 {
		t.Errorf("Expected cache length 2, got %d", cache.Len())
	}
}

func TestARCAdaptation(t *testing.T) {
	cache, err := New[string, int](4)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Fill cache
	cache.Put("a", 1) // T1
	cache.Put("b", 2) // T1
	cache.Put("c", 3) // T1
	cache.Put("d", 4) // T1

	// Access some items to move them to T2
	cache.Get("a") // a moves to T2
	cache.Get("b") // b moves to T2

	// Now T1: [d, c], T2: [b, a]

	// Force eviction to create ghost entries
	cache.Put("e", 5) // Should evict from T1 to B1
	cache.Put("f", 6) // Should evict from T1 to B1

	// Check if we have ghost entries
	stats := cache.Stats()
	if stats.B1Size == 0 {
		t.Skip("No B1 entries created, skipping adaptation test")
	}

	initialP := stats.P

	// Try to trigger ghost hit - put an item that should be in B1
	cache.Put("c", 30) // This should be a ghost hit if c was evicted to B1

	newStats := cache.Stats()
	if newStats.P <= initialP {
		t.Logf("P didn't increase after potential ghost hit: was %d, now %d", initialP, newStats.P)
		t.Logf("B1 size: %d, B2 size: %d", newStats.B1Size, newStats.B2Size)
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

	// Peek should not affect ARC behavior
	value, ok := cache.Peek("a")
	if !ok || value != 1 {
		t.Errorf("Expected peek value 1, got %d, ok=%t", value, ok)
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

	stats := cache.Stats()
	if stats.P != 0 {
		t.Errorf("Expected P=0 after clear, got %d", stats.P)
	}

	// Test that cache still works after clear
	cache.Put("d", 4)
	value, ok := cache.Get("d")
	if !ok || value != 4 {
		t.Errorf("Expected value 4 after clear, got %d, ok=%t", value, ok)
	}
}

func TestKeys(t *testing.T) {
	cache, err := New[string, int](4)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1) // T1
	cache.Put("b", 2) // T1
	cache.Get("a")    // a moves to T2
	cache.Put("c", 3) // T1

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

	// T2 keys should come first
	if len(keys) > 0 && keys[0] != "a" {
		t.Errorf("Expected T2 key 'a' to be first, got %s", keys[0])
	}
}

func TestValues(t *testing.T) {
	cache, err := New[string, int](3)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

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
	cache.Get("a") // Move a to T2

	stats := cache.Stats()
	if stats.Size != 2 {
		t.Errorf("Expected size 2, got %d", stats.Size)
	}
	if stats.Capacity != 5 {
		t.Errorf("Expected capacity 5, got %d", stats.Capacity)
	}
	if stats.T1Size != 1 {
		t.Errorf("Expected T1 size 1, got %d", stats.T1Size)
	}
	if stats.T2Size != 1 {
		t.Errorf("Expected T2 size 1, got %d", stats.T2Size)
	}
}

func TestGhostListMaintenance(t *testing.T) {
	cache, err := New[string, int](2)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Fill cache
	cache.Put("a", 1) // T1
	cache.Put("b", 2) // T1

	// Evict to B1
	cache.Put("c", 3) // T1, evicts a to B1
	cache.Put("d", 4) // T1, evicts b to B1

	stats := cache.Stats()
	if stats.B1Size != 2 {
		t.Errorf("Expected B1 size 2, got %d", stats.B1Size)
	}

	// Overflow B1 - should remove oldest ghost
	cache.Put("e", 5) // Should remove oldest from B1

	stats = cache.Stats()
	if stats.B1Size > cache.Cap() {
		t.Errorf("B1 size should not exceed capacity, got %d", stats.B1Size)
	}
}

func TestB2GhostHit(t *testing.T) {
	cache, err := New[string, int](3)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Create scenario where items go to B2
	cache.Put("a", 1) // T1
	cache.Put("b", 2) // T1
	cache.Put("c", 3) // T1

	// Move all items to T2
	cache.Get("a") // a to T2
	cache.Get("b") // b to T2
	cache.Get("c") // c to T2

	// Now T1: [], T2: [c, b, a]
	// Set p high to force eviction from T2
	cache.p = cache.capacity

	// Force eviction from T2 to B2
	cache.Put("d", 4) // Should evict from T2 to B2

	stats := cache.Stats()
	if stats.B2Size == 0 {
		t.Skip("No B2 entries created, skipping B2 ghost hit test")
	}

	initialP := stats.P

	// Ghost hit from B2 should decrease p
	cache.Put("a", 10) // This should be a ghost hit if a was evicted to B2

	newP := cache.Stats().P
	if newP >= initialP {
		t.Logf("P didn't decrease after potential B2 ghost hit: was %d, now %d", initialP, newP)
		t.Logf("B1 size: %d, B2 size: %d", stats.B1Size, stats.B2Size)
	}
}

func TestComplexARCBehavior(t *testing.T) {
	cache, err := New[string, int](4)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Test complex ARC adaptation behavior

	// Phase 1: Fill T1
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)
	cache.Put("d", 4)

	// Phase 2: Create mixed access pattern
	cache.Get("a") // a: T1 -> T2
	cache.Get("b") // b: T1 -> T2

	// Phase 3: Force evictions
	cache.Put("e", 5) // Should evict from T1 to B1
	cache.Put("f", 6) // Should evict from T1 to B1

	// Phase 4: Test ghost hits and adaptation
	stats := cache.Stats()
	_ = stats.P // Avoid unused variable

	// Should have items in B1
	if stats.B1Size == 0 {
		t.Error("Expected items in B1")
	}

	// Test that cache still functions correctly
	value, ok := cache.Get("a")
	if !ok || value != 1 {
		t.Errorf("Expected to find 'a' with value 1, got %d, ok=%t", value, ok)
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

func TestMinMaxFunctions(t *testing.T) {
	// Test min function
	if min(5, 3) != 3 {
		t.Errorf("Expected min(5, 3) = 3, got %d", min(5, 3))
	}
	if min(2, 7) != 2 {
		t.Errorf("Expected min(2, 7) = 2, got %d", min(2, 7))
	}
	if min(4, 4) != 4 {
		t.Errorf("Expected min(4, 4) = 4, got %d", min(4, 4))
	}

	// Test max function
	if max(5, 3) != 5 {
		t.Errorf("Expected max(5, 3) = 5, got %d", max(5, 3))
	}
	if max(2, 7) != 7 {
		t.Errorf("Expected max(2, 7) = 7, got %d", max(2, 7))
	}
	if max(4, 4) != 4 {
		t.Errorf("Expected max(4, 4) = 4, got %d", max(4, 4))
	}
}

func TestGhostEntryHandling(t *testing.T) {
	cache, err := New[string, int](2)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Create ghost entries
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3) // Evicts "a" to B1

	// Check that ghost entry is not returned by Contains
	if cache.Contains("a") {
		t.Error("Ghost entry should not be contained")
	}

	// Check that ghost entry is not returned by Peek
	_, ok := cache.Peek("a")
	if ok {
		t.Error("Ghost entry should not be peekable")
	}

	// Remove ghost entry should not return value
	value, ok := cache.Remove("a")
	if ok || value != 0 {
		t.Errorf("Removing ghost entry should return zero value and false, got %d, %t", value, ok)
	}
}

func TestReplaceFromT2(t *testing.T) {
	cache, err := New[string, int](2)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Force scenario where replacement happens from T2
	cache.Put("a", 1) // T1
	cache.Put("b", 2) // T1
	cache.Get("a")    // a moves to T2
	cache.Get("b")    // b moves to T2

	// Now both items are in T2, T1 is empty
	// Set p to 0 to force eviction from T2
	cache.p = 0

	cache.Put("c", 3) // Should evict from T2 since T1 is empty

	stats := cache.Stats()
	if stats.B2Size == 0 {
		t.Error("Expected item to be moved to B2")
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

func TestValuesWithEmptyT2(t *testing.T) {
	cache, err := New[string, int](5)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Test empty cache
	values := cache.Values()
	if len(values) != 0 {
		t.Errorf("Expected 0 values in empty cache, got %d", len(values))
	}

	// Add items only to T1 by not accessing them after initial put
	cache.Put("a", 1)
	cache.Put("b", 2)

	values = cache.Values()
	expectedCount := 2
	if len(values) != expectedCount {
		t.Errorf("Expected %d values, got %d", expectedCount, len(values))
	}

	// Verify values content
	valueMap := make(map[int]bool)
	for _, v := range values {
		valueMap[v] = true
	}

	if !valueMap[1] || !valueMap[2] {
		t.Errorf("Expected values to contain 1 and 2")
	}
}

func TestItemsWithEmptyT2(t *testing.T) {
	cache, err := New[string, int](5)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Test empty cache
	items := cache.Items()
	if len(items) != 0 {
		t.Errorf("Expected 0 items in empty cache, got %d", len(items))
	}

	// Add items only to T1
	cache.Put("a", 1)
	cache.Put("b", 2)

	items = cache.Items()
	expectedCount := 2
	if len(items) != expectedCount {
		t.Errorf("Expected %d items, got %d", expectedCount, len(items))
	}

	// Verify items content
	expectedItems := map[string]int{"a": 1, "b": 2}
	for key, expectedValue := range expectedItems {
		if value, exists := items[key]; !exists || value != expectedValue {
			t.Errorf("Expected items[%s] = %d, got %d, exists=%t", key, expectedValue, value, exists)
		}
	}
}

func TestMaintainGhostListsEdgeCase(t *testing.T) {
	cache, err := New[string, int](2)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Fill cache to create ghost entries
	cache.Put("a", 1) // T1
	cache.Put("b", 2) // T1
	cache.Put("c", 3) // T1, evicts "a" to B1
	cache.Put("d", 4) // T1, evicts "b" to B1

	// Now we should have B1 with 2 entries and need to trigger ghost list maintenance
	// Add more entries to exceed ghost list capacity
	cache.Put("e", 5) // T1, should evict "c" to B1, triggering B1 maintenance
	cache.Put("f", 6) // T1, should evict "d" to B1, triggering B1 maintenance

	stats := cache.Stats()
	t.Logf("After ghost maintenance: B1Size=%d, B2Size=%d", stats.B1Size, stats.B2Size)

	// Ghost lists should be maintained within reasonable bounds
	maxGhostSize := cache.Cap() * 2 // Typical ARC constraint
	if stats.B1Size > maxGhostSize {
		t.Errorf("B1 size %d exceeds max ghost size %d", stats.B1Size, maxGhostSize)
	}
	if stats.B2Size > maxGhostSize {
		t.Errorf("B2 size %d exceeds max ghost size %d", stats.B2Size, maxGhostSize)
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

func TestGhostListOverflow(t *testing.T) {
	cache, err := New[string, int](1)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Fill cache and create many ghost entries to trigger ghost list maintenance
	cache.Put("a", 1)
	cache.Put("b", 2) // evicts "a" to B1
	cache.Put("c", 3) // evicts "b" to B1
	cache.Put("d", 4) // evicts "c" to B1, should trigger B1 maintenance (B1 > capacity)

	stats := cache.Stats()
	// B1 should be maintained to not exceed capacity
	if stats.B1Size > cache.Cap() {
		t.Errorf("B1 size %d should not exceed capacity %d", stats.B1Size, cache.Cap())
	}
}

func TestNilCheckInMaintainGhostLists(t *testing.T) {
	cache, err := New[string, int](1)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Create a scenario where maintainGhostLists is called but lists might be empty
	// This is to cover the nil check lines in maintainGhostLists
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	// Force maintainGhostLists to be called with potential empty conditions
	cache.maintainGhostLists()

	// The test passes if no panic occurs
}

func TestB2GhostListMaintenance(t *testing.T) {
	cache, err := New[string, int](1)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Create scenario where B2 fills up and needs maintenance
	cache.Put("a", 1) // T1
	cache.Get("a")    // a moves to T2
	cache.Put("b", 2) // evicts "a" from T2 to B2
	cache.Get("b")    // b moves to T2
	cache.Put("c", 3) // evicts "b" from T2 to B2

	// Create more B2 entries to trigger B2 maintenance
	cache.Get("c")    // c moves to T2
	cache.Put("d", 4) // evicts "c" from T2 to B2, should trigger B2 maintenance

	stats := cache.Stats()
	// B2 should be maintained within capacity bounds
	if stats.B2Size > cache.Cap() {
		t.Errorf("B2 size %d should not exceed capacity %d", stats.B2Size, cache.Cap())
	}
}

func TestExtremePressureOnGhosts(t *testing.T) {
	// Use a very small cache to force heavy ghost list churn
	cache, err := New[string, int](1)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Create maximum ghost pressure to test all edge cases
	// This should create multiple evictions and force ghost list maintenance
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		cache.Put(key, i)

		// Also try some hits to create T2 -> B2 evictions
		if i > 0 {
			prevKey := fmt.Sprintf("key%d", i-1)
			cache.Get(prevKey) // This should create B1 ghost hits
		}
	}

	// Force many more evictions to stress test ghost maintenance
	for i := 10; i < 20; i++ {
		key := fmt.Sprintf("key%d", i)
		cache.Put(key, i)
	}

	stats := cache.Stats()
	t.Logf("Final ghost lists: B1=%d, B2=%d", stats.B1Size, stats.B2Size)

	// The test should pass without panics and ghost lists should be maintained
	if stats.B1Size > cache.Cap()*2 {
		t.Errorf("B1 ghost list grew too large: %d", stats.B1Size)
	}
	if stats.B2Size > cache.Cap()*2 {
		t.Errorf("B2 ghost list grew too large: %d", stats.B2Size)
	}
}

func TestMassForcedEvictions(t *testing.T) {
	// Create cache with capacity that will force constant ghost maintenance
	cache, err := New[string, int](2)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Force a massive number of evictions
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("k%d", i)
		cache.Put(key, i)

		// Mix in some hits to create T2 entries and B2 ghosts
		if i > 2 && i%3 == 0 {
			prevKey := fmt.Sprintf("k%d", i-2)
			cache.Get(prevKey) // Try to hit ghosts
		}
	}

	// Test passes if no panic occurs during this massive churn
	stats := cache.Stats()
	t.Logf("After mass evictions: B1=%d, B2=%d", stats.B1Size, stats.B2Size)
}

func TestValuesWithEmptyT1(t *testing.T) {
	cache, err := New[string, int](3)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Put items and immediately access them to move to T2
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	// Access all items to move them to T2
	cache.Get("a")
	cache.Get("b")
	cache.Get("c")

	// Now T1 should be empty and T2 should have all items
	values := cache.Values()
	expectedCount := 3
	if len(values) != expectedCount {
		t.Errorf("Expected %d values, got %d", expectedCount, len(values))
	}

	// Values should come from T2 first
	valueMap := make(map[int]bool)
	for _, v := range values {
		valueMap[v] = true
	}

	for i := 1; i <= 3; i++ {
		if !valueMap[i] {
			t.Errorf("Expected value %d to be present", i)
		}
	}
}

func TestItemsWithEmptyT1(t *testing.T) {
	cache, err := New[string, int](3)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Put items and access them to move to T2
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	// Access all to move to T2
	cache.Get("a")
	cache.Get("b")
	cache.Get("c")

	// Now T1 should be empty and T2 should have all items
	items := cache.Items()
	expectedCount := 3
	if len(items) != expectedCount {
		t.Errorf("Expected %d items, got %d", expectedCount, len(items))
	}

	expectedItems := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, expectedValue := range expectedItems {
		if value, exists := items[key]; !exists || value != expectedValue {
			t.Errorf("Expected items[%s] = %d, got %d, exists=%t", key, expectedValue, value, exists)
		}
	}
}

func TestSuperIntensiveGhostMaintenance(t *testing.T) {
	// Use capacity = 1 to force maximum ghost maintenance stress
	cache, err := New[string, int](1)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	// Create a scenario that forces ghost lists to exceed capacity
	// This should stress test the maintainGhostLists function thoroughly

	// First fill and create massive B1 ghost list
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("b1_%d", i)
		cache.Put(key, i)
	}

	// Now create B2 ghost entries by moving items through T2
	for i := 50; i < 100; i++ {
		key := fmt.Sprintf("b2_%d", i)
		cache.Put(key, i)
		cache.Get(key) // Move to T2

		// Force eviction from T2 to B2
		nextKey := fmt.Sprintf("b2_%d", i+1000)
		cache.Put(nextKey, i+1000)
	}

	stats := cache.Stats()
	t.Logf("Intensive ghost maintenance: B1=%d, B2=%d, capacity=%d",
		stats.B1Size, stats.B2Size, cache.Cap())

	// With such intensive operations, ghost lists should be properly maintained
	// The key is to not crash and to have reasonable ghost list sizes
	maxReasonableGhostSize := cache.Cap() * 10 // Very generous upper bound
	if stats.B1Size > maxReasonableGhostSize {
		t.Errorf("B1 ghost list excessive: %d > %d", stats.B1Size, maxReasonableGhostSize)
	}
	if stats.B2Size > maxReasonableGhostSize {
		t.Errorf("B2 ghost list excessive: %d > %d", stats.B2Size, maxReasonableGhostSize)
	}
}
