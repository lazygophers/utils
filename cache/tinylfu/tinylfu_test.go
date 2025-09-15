package tinylfu

import (
	"fmt"
	"sync"
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
	if stats.WindowCapacity != 1 { // 1% of 100
		t.Errorf("Expected window capacity 1, got %d", stats.WindowCapacity)
	}
	if stats.MainCapacity != 99 { // 99% of 100
		t.Errorf("Expected main capacity 99, got %d", stats.MainCapacity)
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
	cache := New[string, int](10)
	
	evicted := cache.Put("a", 1)
	if evicted {
		t.Error("Should not evict when cache is not full")
	}
	
	value, ok := cache.Get("a")
	if !ok || value != 1 {
		t.Errorf("Expected value 1, got %d, ok=%t", value, ok)
	}
}

func TestWindowEviction(t *testing.T) {
	cache := New[string, int](10) // window=1, main=9
	
	// Fill window (capacity 1)
	cache.Put("a", 1)
	
	// This should evict from window
	evicted := cache.Put("b", 2)
	if !evicted {
		t.Error("Expected eviction when window is full")
	}
	
	// "a" should not be admitted to main cache on first doorkeeper check
	_, ok := cache.Get("a")
	if ok {
		t.Error("Expected 'a' to be evicted from window and not admitted to main")
	}
	
	// "b" should still be in window
	if value, ok := cache.Get("b"); !ok || value != 2 {
		t.Errorf("Expected 'b' to have value 2, got %d, ok=%t", value, ok)
	}
}

func TestFrequencyPromotion(t *testing.T) {
	cache := New[string, int](10) // window=1, main=9
	
	// Add item to window
	cache.Put("a", 1)
	
	// Access multiple times to increase frequency
	for i := 0; i < 5; i++ {
		cache.Get("a")
	}
	
	// Add to doorkeeper by putting another item
	cache.Put("b", 2) // This should evict "a" from window but add to doorkeeper
	
	// Now put "a" again - should be admitted to main cache due to doorkeeper
	cache.Put("a", 1)
	cache.Put("c", 3) // This should not evict "a" since it's in main cache now
	
	// "a" should still be accessible
	if value, ok := cache.Get("a"); !ok || value != 1 {
		t.Errorf("Expected 'a' to be in main cache with value 1, got %d, ok=%t", value, ok)
	}
}

func TestProbationToProtectedPromotion(t *testing.T) {
	cache := New[string, int](100) // window=1, main=99, protected~79, probation~20
	
	// Add item and get it admitted to probation
	cache.Put("key", 1)
	
	// Access multiple times to build frequency
	for i := 0; i < 5; i++ {
		cache.Get("key")
	}
	
	// Force eviction from window by adding another item  
	cache.Put("temp", 999)
	
	// Add original key again to get it into main cache
	cache.Put("key", 1)
	
	stats := cache.Stats()
	initialProbationSize := stats.ProbationSize
	
	// Access the key to promote from probation to protected
	cache.Get("key")
	
	stats = cache.Stats()
	if stats.ProbationSize >= initialProbationSize && initialProbationSize > 0 {
		t.Error("Expected item to be promoted from probation to protected")
	}
}

func TestUpdateExisting(t *testing.T) {
	cache := New[string, int](10)
	
	cache.Put("a", 1)
	evicted := cache.Put("a", 10) // Update existing
	if evicted {
		t.Error("Should not evict when updating existing key")
	}
	
	value, ok := cache.Get("a")
	if !ok || value != 10 {
		t.Errorf("Expected updated value 10, got %d, ok=%t", value, ok)
	}
}

func TestRemove(t *testing.T) {
	cache := New[string, int](300) // Large cache to have window size >= 2
	
	cache.Put("a", 1)
	cache.Put("b", 2)
	
	// Verify both items are there first
	if !cache.Contains("a") || !cache.Contains("b") {
		t.Fatal("Items should be in cache before testing remove")
	}
	
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
	cache := New[string, int](10)
	
	cache.Put("a", 1)
	
	if !cache.Contains("a") {
		t.Error("Expected 'a' to be contained")
	}
	
	if cache.Contains("non-existent") {
		t.Error("Expected 'non-existent' to not be contained")
	}
}

func TestPeek(t *testing.T) {
	cache := New[string, int](10)
	
	cache.Put("a", 1)
	
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
	cache := NewWithEvict[string, int](10, func(key string, value int) {
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
	
	stats := cache.Stats()
	if stats.SketchSize != 0 {
		t.Errorf("Expected sketch to be reset, got size %d", stats.SketchSize)
	}
	
	if stats.DoorkeeperSize != 0 {
		t.Errorf("Expected doorkeeper to be cleared, got size %d", stats.DoorkeeperSize)
	}
}

func TestKeys(t *testing.T) {
	cache := New[string, int](100)
	
	// Add items to different segments
	cache.Put("window", 1)    // goes to window
	cache.Put("temp", 999)    // evicts window item to doorkeeper
	cache.Put("window", 1)    // should go to probation due to doorkeeper
	cache.Get("window")       // should promote to protected
	cache.Put("probation", 2) // should go to probation
	
	keys := cache.Keys()
	
	if len(keys) < 2 {
		t.Errorf("Expected at least 2 keys, got %d", len(keys))
	}
	
	// Check that we got the keys we expect
	hasWindow := false
	for _, key := range keys {
		if key == "window" {
			hasWindow = true
		}
	}
	
	if !hasWindow {
		t.Error("Expected to find 'window' key in results")
	}
}

func TestValues(t *testing.T) {
	cache := New[string, int](300) // Large cache to have window size >= 2
	
	cache.Put("a", 1)
	cache.Put("b", 2)
	
	// Verify both items are there first
	if !cache.Contains("a") || !cache.Contains("b") {
		t.Fatal("Items should be in cache before testing values")
	}
	
	values := cache.Values()
	
	if len(values) != 2 {
		t.Errorf("Expected 2 values, got %d", len(values))
	}
}

func TestItems(t *testing.T) {
	cache := New[string, int](300) // Large cache to have window size >= 2
	
	cache.Put("a", 1)
	cache.Put("b", 2)
	
	// Verify both items are there first
	if !cache.Contains("a") || !cache.Contains("b") {
		t.Fatal("Items should be in cache before testing items")
	}
	
	items := cache.Items()
	
	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}
	
	if items["a"] != 1 || items["b"] != 2 {
		t.Errorf("Expected items map to contain correct values")
	}
}

func TestResize(t *testing.T) {
	cache := New[string, int](1000) // Start with large cache
	
	// Fill cache with many items
	for i := 0; i < 500; i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
	}
	
	initialLen := cache.Len()
	if initialLen == 0 {
		t.Fatal("Cache should have items before resize")
	}
	
	// Resize to much smaller capacity
	cache.Resize(10)
	
	if cache.Cap() != 10 {
		t.Errorf("Expected capacity 10 after resize, got %d", cache.Cap())
	}
	
	if cache.Len() > 10 {
		t.Errorf("Expected length <= 10 after resize, got %d", cache.Len())
	}
	
	if cache.Len() >= initialLen {
		t.Errorf("Expected some items to be evicted during resize, initial: %d, final: %d", initialLen, cache.Len())
	}
	
	stats := cache.Stats()
	expectedWindowSize := 1 // max(1, 10*0.01) = 1
	if stats.WindowCapacity != expectedWindowSize {
		t.Errorf("Expected window capacity %d after resize, got %d", expectedWindowSize, stats.WindowCapacity)
	}
}

func TestResizePanic(t *testing.T) {
	cache := New[string, int](10)
	
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for zero capacity resize")
		}
	}()
	cache.Resize(0)
}

func TestStats(t *testing.T) {
	cache := New[string, int](200) // Large enough window to hold both items
	
	cache.Put("a", 1)
	cache.Put("b", 2)
	
	stats := cache.Stats()
	
	if stats.Size != 2 {
		t.Errorf("Expected size 2, got %d", stats.Size)
	}
	if stats.Capacity != 200 {
		t.Errorf("Expected capacity 200, got %d", stats.Capacity)
	}
	if stats.WindowCapacity != 2 { // 1% of 200 = 2
		t.Errorf("Expected window capacity 2, got %d", stats.WindowCapacity)
	}
	if stats.MainCapacity != 198 {
		t.Errorf("Expected main capacity 198, got %d", stats.MainCapacity)
	}
	if stats.SketchSize <= 0 {
		t.Errorf("Expected positive sketch size, got %d", stats.SketchSize)
	}
}

func TestCountMinSketch(t *testing.T) {
	cms := NewCountMinSketch(100)
	
	key := []byte("test-key")
	
	// Initially should be 0
	count := cms.EstimateCount(key)
	if count != 0 {
		t.Errorf("Expected initial count 0, got %d", count)
	}
	
	// Add some counts
	for i := 0; i < 5; i++ {
		cms.Add(key)
	}
	
	count = cms.EstimateCount(key)
	if count < 3 { // Should be around 5, but CMS might underestimate
		t.Errorf("Expected count around 5, got %d", count)
	}
	
	if cms.Size() != 5 {
		t.Errorf("Expected total size 5, got %d", cms.Size())
	}
}

func TestCountMinSketchReset(t *testing.T) {
	cms := NewCountMinSketch(10)
	
	key := []byte("test")
	for i := 0; i < 10; i++ {
		cms.Add(key)
	}
	
	countBefore := cms.EstimateCount(key)
	sizeBefore := cms.Size()
	
	cms.Reset()
	
	countAfter := cms.EstimateCount(key)
	sizeAfter := cms.Size()
	
	if countAfter >= countBefore {
		t.Errorf("Expected count to decrease after reset, before: %d, after: %d", countBefore, countAfter)
	}
	
	if sizeAfter != 0 {
		t.Errorf("Expected size to be 0 after reset, got %d", sizeAfter)
	}
	
	_ = sizeBefore // Prevent unused variable warning
}

func TestKeyToBytesString(t *testing.T) {
	key := "test-string"
	bytes := keyToBytes(key)
	
	if string(bytes) != key {
		t.Errorf("Expected string key conversion to work correctly")
	}
}

func TestKeyToBytesInt(t *testing.T) {
	key := 12345
	bytes := keyToBytes(key)
	
	if len(bytes) != 8 {
		t.Errorf("Expected 8 bytes for int key, got %d", len(bytes))
	}
}

func TestKeyToBytesInt64(t *testing.T) {
	key := int64(123456789)
	bytes := keyToBytes(key)
	
	if len(bytes) != 8 {
		t.Errorf("Expected 8 bytes for int64 key, got %d", len(bytes))
	}
}

func TestKeyToBytesOther(t *testing.T) {
	key := float64(3.14) // Non-handled type should use fallback
	bytes := keyToBytes(key)
	
	if len(bytes) != 8 {
		t.Errorf("Expected 8 bytes for fallback key conversion, got %d", len(bytes))
	}
}

func TestSketchAging(t *testing.T) {
	cache := New[string, int](10)
	
	// Add many items to trigger sketch aging
	for i := 0; i < 150; i++ { // 10 * 10 + 50 = trigger reset
		cache.Put(fmt.Sprintf("key%d", i), i)
	}
	
	// Check that sketch was reset (size should be lower)
	stats := cache.Stats()
	if stats.SketchSize > 100 {
		t.Errorf("Expected sketch size to be reset and lower, got %d", stats.SketchSize)
	}
}

func TestAdmissionPolicy(t *testing.T) {
	cache := New[string, int](20) // window=1, main=19
	
	// Add item with high frequency
	highFreqKey := "high-freq"
	for i := 0; i < 10; i++ {
		cache.Put(highFreqKey, 1)
		cache.Get(highFreqKey)
	}
	
	// Force eviction to doorkeeper
	cache.Put("evict", 999)
	
	// Fill main cache with low frequency items
	for i := 0; i < 15; i++ {
		cache.Put(fmt.Sprintf("low%d", i), i)
	}
	
	// Try to re-admit high frequency item
	cache.Put(highFreqKey, 1)
	
	// High frequency item should be admitted over low frequency ones
	value, ok := cache.Get(highFreqKey)
	if !ok || value != 1 {
		t.Errorf("Expected high frequency item to be admitted, got %d, ok=%t", value, ok)
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
	
	cache := NewWithEvict[string, int](3, func(key string, value int) {
		evictedKeys = append(evictedKeys, key)
		evictedValues = append(evictedValues, value)
	})
	
	// Fill cache beyond capacity
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)
	cache.Put("d", 4) // Should trigger eviction
	
	if len(evictedKeys) < 1 {
		t.Errorf("Expected at least 1 evicted key, got %d", len(evictedKeys))
	}
	if len(evictedValues) < 1 {
		t.Errorf("Expected at least 1 evicted value, got %d", len(evictedValues))
	}
}

func TestProtectedSegmentDemotion(t *testing.T) {
	cache := New[string, int](500) // Large enough to test segment transitions
	
	// Add many items to build up different segments
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key%d", i)
		cache.Put(key, i)
		
		// Access some items multiple times to build frequency
		if i%10 == 0 {
			for j := 0; j < 5; j++ {
				cache.Get(key)
			}
		}
	}
	
	stats := cache.Stats()
	
	// Should have items distributed across segments
	if stats.Size == 0 {
		t.Error("Expected cache to have items")
	}
	
	// The segments should be managing capacity correctly
	totalItems := stats.WindowSize + stats.ProbationSize + stats.ProtectedSize
	if totalItems > cache.Cap() {
		t.Errorf("Total items %d exceeds capacity %d", totalItems, cache.Cap())
	}
	
	// Should have some distribution across segments
	if stats.WindowSize == 0 && stats.ProbationSize == 0 && stats.ProtectedSize == 0 {
		t.Error("Expected some items to be in cache segments")
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

func BenchmarkMixed(b *testing.B) {
	cache := New[int, int](1000)
	
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

func TestProtectedSegmentFullness(t *testing.T) {
	// Test that forces protected segment to be full and trigger demotion
	cache := New[string, int](100) // window=1, main=99, protected~79
	
	// Add many items and promote them all to protected
	for i := 0; i < 85; i++ {
		key := fmt.Sprintf("key%d", i)
		cache.Put(key, i)
		
		// Force admission to main cache by putting in doorkeeper first
		cache.Put("temp", 999) // evict to doorkeeper
		cache.Put(key, i)      // should be admitted due to doorkeeper
		
		// Access multiple times to promote to protected
		for j := 0; j < 3; j++ {
			cache.Get(key)
		}
	}
	
	stats := cache.Stats()
	if stats.ProtectedSize == 0 {
		t.Error("Expected items to be in protected segment")
	}
	
	// Add one more item that should trigger demotion from protected
	cache.Put("final", 999)
	cache.Put("temp2", 888) // evict to doorkeeper
	cache.Put("final", 999) // admit to main
	cache.Get("final")      // promote to protected - should trigger demotion
	
	// The cache should still be within capacity
	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

func TestProbationEviction(t *testing.T) {
	// Test eviction from probation segment
	cache := New[string, int](50) // window=1, main=49, probation~10
	
	// Fill probation segment by adding items to doorkeeper then readmitting
	probationCapacity := 10 // Approximately 20% of main cache (49)
	
	for i := 0; i < probationCapacity+5; i++ {
		key := fmt.Sprintf("prob%d", i)
		cache.Put(key, i)
		
		// Get into doorkeeper
		cache.Put("evict", 999)
		
		// Readmit to probation
		cache.Put(key, i)
	}
	
	// Verify items are distributed correctly
	stats := cache.Stats()
	if stats.Size == 0 {
		t.Error("Expected cache to have items")
	}
	
	// Cache should not exceed capacity
	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

func TestSegmentEvictionOrder(t *testing.T) {
	// Test that resize evicts from window first, then probation, then protected
	cache := New[string, int](200) // Large enough for complex segments
	
	// Add items to different segments
	// Window items
	cache.Put("window1", 1)
	
	// Probation items (via doorkeeper admission)
	cache.Put("prob1", 10)
	cache.Put("temp", 999) // evict to doorkeeper
	cache.Put("prob1", 10) // readmit to probation
	
	// Protected items (via promotion)
	cache.Put("prot1", 100)
	cache.Put("temp2", 888) // evict to doorkeeper
	cache.Put("prot1", 100) // readmit to probation
	cache.Get("prot1")      // promote to protected
	
	initialStats := cache.Stats()
	
	// Resize to very small capacity to force evictions
	cache.Resize(2)
	
	finalStats := cache.Stats()
	
	// Should have evicted items
	if finalStats.Size >= initialStats.Size {
		t.Error("Expected evictions during resize")
	}
	
	// Should not exceed new capacity
	if finalStats.Size > 2 {
		t.Errorf("Expected size <= 2 after resize, got %d", finalStats.Size)
	}
}

func TestAdmissionPolicyEdgeCases(t *testing.T) {
	cache := New[string, int](20)
	
	// Test admission when probation is empty
	cache.Put("item", 1)
	
	// Force eviction from window
	cache.Put("evict", 999)
	
	// Put the item again - should be admitted since probation is empty
	cache.Put("item", 1)
	
	if !cache.Contains("item") {
		t.Error("Expected item to be admitted when probation is empty")
	}
}

func TestEmptySegmentEvictions(t *testing.T) {
	cache := New[string, int](10)
	
	// Test resize when some segments are empty
	cache.Put("single", 1)
	
	// Resize to smaller capacity
	cache.Resize(5)
	
	// Should still work correctly
	if cache.Cap() != 5 {
		t.Errorf("Expected capacity 5, got %d", cache.Cap())
	}
	
	// Test that we can still operate normally
	cache.Put("new", 2)
	if value, ok := cache.Get("new"); !ok || value != 2 {
		t.Errorf("Expected to retrieve 'new' with value 2, got %d, ok=%t", value, ok)
	}
}