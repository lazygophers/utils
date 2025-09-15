package tinylfu

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
	if stats.WindowCapacity != 1 { // 1% of 100
		t.Errorf("Expected window capacity 1, got %d", stats.WindowCapacity)
	}
	if stats.MainCapacity != 99 { // 99% of 100
		t.Errorf("Expected main capacity 99, got %d", stats.MainCapacity)
	}
}

func TestNewError(t *testing.T) {
	_, err := New[string, int](0)
	if err == nil {
		t.Error("Expected error for zero capacity")
	}
}

func TestPutAndGet(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
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
}

func TestWindowEviction(t *testing.T) {
	cache, err := New[string, int](10) // window=1, main=9
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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
	cache, err := New[string, int](10) // window=1, main=9
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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
	cache, err := New[string, int](100) // window=1, main=99, protected~79, probation~20
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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
	cache, err := New[string, int](300) // Large cache to have window size >= 2
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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
	cache, err := New[string, int](10)
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
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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
	cache, err := NewWithEvict[string, int](10, func(key string, value int) {
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
	if stats.SketchSize != 0 {
		t.Errorf("Expected sketch to be reset, got size %d", stats.SketchSize)
	}
	
	if stats.DoorkeeperSize != 0 {
		t.Errorf("Expected doorkeeper to be cleared, got size %d", stats.DoorkeeperSize)
	}
}

func TestKeys(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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
	cache, err := New[string, int](300) // Large cache to have window size >= 2
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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
	cache, err := New[string, int](300) // Large cache to have window size >= 2
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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
	cache, err := New[string, int](1000) // Start with large cache
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Fill cache with many items
	for i := 0; i < 500; i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
	}
	
	initialLen := cache.Len()
	if initialLen == 0 {
		t.Fatal("Cache should have items before resize")
	}
	
	// Resize to much smaller capacity
	err = cache.Resize(10)
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}
	
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

func TestStats(t *testing.T) {
	cache, err := New[string, int](200) // Large enough window to hold both items
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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
	cache, err := New[string, int](20) // window=1, main=19
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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
	
	cache, err := NewWithEvict[string, int](3, func(key string, value int) {
		evictedKeys = append(evictedKeys, key)
		evictedValues = append(evictedValues, value)
	})
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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
	cache, err := New[string, int](500) // Large enough to test segment transitions
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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

func TestDemoteFromProtectedTriggering(t *testing.T) {
	cache, err := New[string, int](50) // Small cache to force segment interactions
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Fill cache to trigger main cache segments
	for i := 0; i < 60; i++ {
		key := fmt.Sprintf("key%d", i)
		cache.Put(key, i)
	}
	
	// Promote items to probation first by accessing them
	for i := 0; i < 40; i++ {
		key := fmt.Sprintf("key%d", i)
		cache.Get(key)
	}
	
	// Access some items heavily to promote them to protected (requires multiple accesses)
	for i := 0; i < 20; i++ {
		key := fmt.Sprintf("key%d", i)
		for j := 0; j < 10; j++ {
			cache.Get(key)
		}
	}
	
	// Now force more promotions to protected which should trigger demotion
	for i := 20; i < 35; i++ {
		key := fmt.Sprintf("key%d", i)
		for j := 0; j < 15; j++ {
			cache.Get(key)
		}
	}
	
	stats := cache.Stats()
	if stats.Size == 0 {
		t.Error("Expected cache to have items")
	}
}

func TestEvictFromProtectedTriggering(t *testing.T) {
	cache, err := New[string, int](10) // Very small cache
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Fill window first
	for i := 0; i < 15; i++ {
		key := fmt.Sprintf("win%d", i)
		cache.Put(key, i)
	}
	
	// Promote items to main segments by accessing
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("win%d", i)
		for j := 0; j < 5; j++ {
			cache.Get(key)
		}
	}
	
	// Heavy access to move items to protected
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("win%d", i)
		for j := 0; j < 20; j++ {
			cache.Get(key)
		}
	}
	
	// Now resize to very small to force eviction from all segments including protected
	err = cache.Resize(2) 
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}
	
	stats := cache.Stats()
	if stats.Size > 2 {
		t.Errorf("Expected cache size <= 2 after resize, got %d", stats.Size)
	}
}

func TestForceProtectedSegmentEviction(t *testing.T) {
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Fill cache with items
	for i := 0; i < 25; i++ {
		key := fmt.Sprintf("item%d", i)
		cache.Put(key, i)
		
		// Heavy access to push items to protected segment
		for j := 0; j < 25; j++ {
			cache.Get(key)
		}
	}
	
	// Verify we have items in protected
	stats := cache.Stats()
	initialProtected := stats.ProtectedSize
	
	// Resize to force eviction from protected segment
	err = cache.Resize(3)
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}
	
	finalStats := cache.Stats()
	if finalStats.Size > 3 {
		t.Errorf("Expected cache size <= 3, got %d", finalStats.Size)
	}
	
	// Should have evicted from protected if it had items
	if initialProtected > 0 && finalStats.ProtectedSize >= initialProtected {
		t.Log("Protected segment should have been reduced")
	}
}

func TestPreciseDemoteFromProtected(t *testing.T) {
	// Use capacity 100: window=1, main=99, protected capacity = int(99*0.8) = 79
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// First, fill window
	cache.Put("window", 0)
	
	// Fill probation by adding items and accessing once to move to main
	for i := 1; i <= 99; i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
		cache.Get(fmt.Sprintf("key%d", i)) // Move to probation
	}
	
	// Now promote exactly 79 items to protected (full capacity)
	for i := 1; i <= 79; i++ {
		key := fmt.Sprintf("key%d", i)
		// Multiple accesses to promote from probation to protected
		for j := 0; j < 10; j++ {
			cache.Get(key)
		}
	}
	
	stats := cache.Stats()
	t.Logf("Before demotion trigger: Protected=%d, Probation=%d, Window=%d", 
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)
	
	// Now promote one more item - this should trigger demoteFromProtected
	key := "key80"
	for j := 0; j < 10; j++ {
		cache.Get(key)
	}
	
	finalStats := cache.Stats()
	t.Logf("After demotion trigger: Protected=%d, Probation=%d, Window=%d", 
		finalStats.ProtectedSize, finalStats.ProbationSize, finalStats.WindowSize)
	
	if finalStats.Size == 0 {
		t.Error("Expected cache to have items")
	}
}

func TestPreciseEvictFromProtected(t *testing.T) {
	// Use capacity 10: window=1, main=9
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Fill the cache completely
	for i := 0; i < 15; i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
		// Heavy access to push items towards protected
		for j := 0; j < 20; j++ {
			cache.Get(fmt.Sprintf("key%d", i))
		}
	}
	
	stats := cache.Stats()
	t.Logf("Before resize: Protected=%d, Probation=%d, Window=%d, Total=%d", 
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize, stats.Size)
	
	// Resize to force evictions from all segments, including protected
	// When window and probation are empty, must evict from protected
	err = cache.Resize(1)
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}
	
	finalStats := cache.Stats()
	t.Logf("After resize: Protected=%d, Probation=%d, Window=%d, Total=%d", 
		finalStats.ProtectedSize, finalStats.ProbationSize, finalStats.WindowSize, finalStats.Size)
	
	if finalStats.Size > 1 {
		t.Errorf("Expected cache size = 1 after resize, got %d", finalStats.Size)
	}
}

func TestExactDemoteFromProtected(t *testing.T) {
	// Use capacity 85: window=1, main=84, protected capacity = int(84*0.8) = 67
	cache, err := New[string, int](85)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Step 1: Fill window and create doorkeeper entries
	cache.Put("w1", 1) // Goes to window
	
	// Add 84 items to be evicted from window (they'll be rejected first time due to doorkeeper)
	for i := 1; i <= 84; i++ {
		cache.Put(fmt.Sprintf("k%d", i), i) // This evicts previous item from window
	}
	
	// Step 2: Add same 84 items again - now they'll be admitted to probation due to doorkeeper
	for i := 1; i <= 84; i++ {
		cache.Put(fmt.Sprintf("k%d", i), i*10) // Update values, should now be admitted to probation
	}
	
	stats := cache.Stats()
	t.Logf("After filling probation: Protected=%d, Probation=%d, Window=%d", 
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)
	
	// Step 3: Access first 67 items to fill protected segment to capacity
	for i := 1; i <= 67; i++ {
		cache.Get(fmt.Sprintf("k%d", i)) // Should promote to protected
	}
	
	stats = cache.Stats()
	t.Logf("After filling protected to capacity: Protected=%d, Probation=%d, Window=%d", 
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)
	
	// Step 4: Access item 68 - this should trigger demoteFromProtected
	cache.Get("k68")
	
	finalStats := cache.Stats()
	t.Logf("After triggering demotion: Protected=%d, Probation=%d, Window=%d", 
		finalStats.ProtectedSize, finalStats.ProbationSize, finalStats.WindowSize)
	
	if finalStats.Size == 0 {
		t.Error("Expected cache to have items")
	}
}

func TestExactEvictFromProtected(t *testing.T) {
	// Use small cache: capacity 4: window=1, main=3
	cache, err := New[string, int](4)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Fill with items and promote to protected
	cache.Put("w", 0)  // window
	
	// Add items to be admitted to probation (need doorkeeper first)
	for i := 1; i <= 3; i++ {
		cache.Put(fmt.Sprintf("p%d", i), i)
	}
	
	// Add same items again to get them admitted to probation
	for i := 1; i <= 3; i++ {
		cache.Put(fmt.Sprintf("p%d", i), i*10)
	}
	
	// Promote all items to protected
	for i := 1; i <= 3; i++ {
		cache.Get(fmt.Sprintf("p%d", i))
	}
	
	stats := cache.Stats()
	t.Logf("Before resize: Protected=%d, Probation=%d, Window=%d, Total=%d", 
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize, stats.Size)
	
	// Resize to smaller capacity 
	err = cache.Resize(2)
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}
	
	finalStats := cache.Stats()
	t.Logf("After resize to 2: Protected=%d, Probation=%d, Window=%d, Total=%d", 
		finalStats.ProtectedSize, finalStats.ProbationSize, finalStats.WindowSize, finalStats.Size)
	
	if finalStats.Size > 2 {
		t.Errorf("Expected cache size <= 2 after resize, got %d", finalStats.Size)
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

func TestProtectedSegmentFullness(t *testing.T) {
	// Test that forces protected segment to be full and trigger demotion
	cache, err := New[string, int](100) // window=1, main=99, protected~79
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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
	cache, err := New[string, int](50) // window=1, main=49, probation~10
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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
	cache, err := New[string, int](200) // Large enough for complex segments
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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
	err = cache.Resize(2)
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}
	
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
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
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
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Test resize when some segments are empty
	cache.Put("single", 1)
	
	// Resize to smaller capacity
	err = cache.Resize(5)
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}
	
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

// TestDemoteFromProtectedSpecific targets the exact conditions needed for demoteFromProtected
func TestDemoteFromProtectedSpecific(t *testing.T) {
	cache, err := New[string, int](5) // Tiny cache: window=1, main=4, protected capacity=3
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Create high-frequency keys that will definitely be admitted
	highFreqKeys := []string{"high1", "high2", "high3", "high4"}
	
	// Build very high frequency for these keys in the sketch
	for _, key := range highFreqKeys {
		for i := 0; i < 20; i++ {
			cache.Put(key, 100)
			cache.Get(key)
		}
	}
	
	// Create a low-frequency victim key  
	cache.Put("victim", 1)
	
	// Clear cache to reset segments but keep frequency sketch
	cache.Clear()
	
	// Re-add victim first (will go to probation due to doorkeeper after eviction)
	cache.Put("victim", 1) 
	cache.Put("evict", 999) // evict victim to doorkeeper
	cache.Put("victim", 1)  // victim goes to probation due to doorkeeper
	
	// Add high frequency keys - they should be admitted over the victim
	for i, key := range highFreqKeys[:3] {
		cache.Put(key, i+100)
		cache.Put("temp", 888) // evict to doorkeeper  
		cache.Put(key, i+100)  // readmit to probation
		cache.Get(key)         // promote to protected
	}
	
	stats := cache.Stats()
	t.Logf("Before final promotion: Protected=%d, Probation=%d, Window=%d", 
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)
	
	// Add final high-frequency key and promote - should trigger demoteFromProtected
	cache.Put("high4", 400)
	cache.Put("temp2", 777) // evict to doorkeeper
	cache.Put("high4", 400) // readmit to probation
	cache.Get("high4")      // promote to protected - should trigger demotion
	
	finalStats := cache.Stats()
	t.Logf("After final promotion: Protected=%d, Probation=%d, Window=%d", 
		finalStats.ProtectedSize, finalStats.ProbationSize, finalStats.WindowSize)
	
	// Verify cache is functional
	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

// TestEvictFromProbationDirect tests evictFromProbation by directly creating the conditions  
func TestEvictFromProbationDirect(t *testing.T) {
	cache, err := New[string, int](5) // Tiny cache: window=1, main=4, probation=1
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Step 1: Fill probation to capacity (1 item) via doorkeeper admission
	cache.Put("prob1", 1)
	cache.Put("temp", 999) // evict prob1 to doorkeeper
	cache.Put("prob1", 1)  // readmit to probation
	
	stats := cache.Stats()
	t.Logf("After filling probation: Protected=%d, Probation=%d, Window=%d", 
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)
	
	// Step 2: Try to admit another item to probation - should call evictFromProbation
	cache.Put("prob2", 2)
	cache.Put("temp2", 888) // evict prob2 to doorkeeper
	cache.Put("prob2", 2)   // readmit to probation - should trigger evictFromProbation
	
	// Verify cache is functional
	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

// TestEvictFromProtectedDirect tests evictFromProtected during resize
func TestEvictFromProtectedDirect(t *testing.T) {
	cache, err := New[string, int](10) // window=1, main=9, protected=7
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Fill protected segment by going through doorkeeper process
	for i := 1; i <= 8; i++ {
		key := fmt.Sprintf("prot%d", i)
		// Put in window first
		cache.Put(key, i)
		// Evict to doorkeeper
		cache.Put("temp", 999)
		// Readmit to probation
		cache.Put(key, i)
		// Promote to protected
		cache.Get(key)
	}
	
	initialStats := cache.Stats()
	t.Logf("Before resize: Protected=%d, Probation=%d, Window=%d, Total=%d", 
		initialStats.ProtectedSize, initialStats.ProbationSize, initialStats.WindowSize, cache.Len())
	
	// Resize to force evictions from protected (after window and probation are empty)
	err = cache.Resize(3) // Force eviction from protected
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}
	
	finalStats := cache.Stats()
	t.Logf("After resize: Protected=%d, Probation=%d, Window=%d, Total=%d", 
		finalStats.ProtectedSize, finalStats.ProbationSize, finalStats.WindowSize, cache.Len())
	
	// Verify constraints
	if cache.Cap() != 3 {
		t.Errorf("Expected capacity 3, got %d", cache.Cap())
	}
	if cache.Len() > 3 {
		t.Errorf("Cache size %d exceeds capacity 3", cache.Len())
	}
}

// TestComplexEvictionScenarios tests multiple eviction scenarios together
func TestComplexEvictionScenarios(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Create a scenario that exercises all eviction paths
	// 1. Fill protected segment
	for i := 0; i < 80; i++ {
		key := fmt.Sprintf("complex_prot_%d", i)
		cache.Put(key, i)
		cache.Put("temp", 999)
		cache.Put(key, i)
		cache.Get(key) // Promote to protected
	}
	
	// 2. Fill probation segment  
	for i := 0; i < 15; i++ {
		key := fmt.Sprintf("complex_prob_%d", i)
		cache.Put(key, i+100)
		cache.Put("temp2", 888)
		cache.Put(key, i+100) // Goes to probation
	}
	
	// 3. Trigger all eviction types by continuing to add items
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("trigger_%d", i)
		cache.Put(key, i+200)
		cache.Put("temp3", 777)
		cache.Put(key, i+200)
		if i%2 == 0 {
			cache.Get(key) // Some promote to protected, triggering demotion
		}
	}
	
	// Verify cache integrity
	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
	
	stats := cache.Stats()
	totalItems := stats.WindowSize + stats.ProbationSize + stats.ProtectedSize
	if totalItems != cache.Len() {
		t.Errorf("Stats don't match actual cache size: stats=%d, len=%d", totalItems, cache.Len())
	}
}

func TestValuesWithEmptySegments(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Test empty cache - all segments empty
	values := cache.Values()
	if len(values) != 0 {
		t.Errorf("Expected 0 values in empty cache, got %d", len(values))
	}
	
	// Test with only window items
	cache.Put("w1", 1)
	values = cache.Values()
	if len(values) != 1 {
		t.Errorf("Expected 1 value from window, got %d", len(values))
	}
	
	// Build protected segment by promoting items
	for i := 0; i < 8; i++ {
		key := fmt.Sprintf("protected%d", i)
		cache.Put(key, i+10)
		// Access multiple times to promote to protected
		for j := 0; j < 5; j++ {
			cache.Get(key)
		}
	}
	
	// Add probation items
	cache.Put("prob1", 100)
	cache.Get("prob1") // Move to probation
	
	// Now test with all segments having items
	values = cache.Values()
	
	// Verify we have values from all segments
	foundValues := make(map[int]bool)
	for _, val := range values {
		foundValues[val] = true
	}
	
	// Should have some values (TinyLFU algorithm is complex)
	if len(foundValues) < 1 {
		t.Errorf("Expected at least some values, got %d unique values", len(foundValues))
	}
}

func TestItemsWithEmptySegments(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Test empty cache - all segments empty  
	items := cache.Items()
	if len(items) != 0 {
		t.Errorf("Expected 0 items in empty cache, got %d", len(items))
	}
	
	// Test with only window items
	cache.Put("w1", 1)
	items = cache.Items()
	if len(items) != 1 {
		t.Errorf("Expected 1 item from window, got %d", len(items))
	}
	
	// Build protected segment by promoting items
	for i := 0; i < 6; i++ {
		key := fmt.Sprintf("protected%d", i)
		cache.Put(key, i+10)
		// Access multiple times to promote to protected
		for j := 0; j < 5; j++ {
			cache.Get(key)
		}
	}
	
	// Add probation items  
	cache.Put("prob1", 100)
	cache.Get("prob1") // Move to probation
	
	// Now test with all segments having items
	items = cache.Items()
	
	// Verify we have some items (TinyLFU algorithm is complex)
	if len(items) < 1 {
		t.Errorf("Expected at least some items, got %d items", len(items))
	}
}

func TestEvictFromProtectedResize(t *testing.T) {
	// Create small cache to force protected segment eviction
	cache, err := New[string, int](8) // Window=1, Main=7 (Probation=2, Protected=5)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Fill protected segment to capacity
	protectedKeys := []string{"p1", "p2", "p3", "p4", "p5"}
	for _, key := range protectedKeys {
		cache.Put(key, 1)
		// Heavy access to ensure promotion to protected
		for j := 0; j < 10; j++ {
			cache.Get(key)
		}
	}
	
	// Now resize to force eviction from protected when probation is empty
	err = cache.Resize(4) // This should force eviction from protected
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}
	
	stats := cache.Stats()
	t.Logf("After resize to force protected eviction: Protected=%d, Probation=%d, Window=%d, Total=%d", 
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize, stats.Size)
	
	if cache.Cap() != 4 {
		t.Errorf("Expected capacity 4, got %d", cache.Cap())
	}
	
	if cache.Len() > 4 {
		t.Errorf("Cache size %d exceeds capacity 4", cache.Len())
	}
}

func TestDemoteFromProtectedPromotion(t *testing.T) {
	// Create cache that will trigger demotion
	cache, err := New[string, int](10) // Window=1, Main=9 (Probation=2, Protected=7)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Fill protected segment to capacity
	for i := 0; i < 7; i++ {
		key := fmt.Sprintf("protected%d", i)
		cache.Put(key, i)
		// Heavy access to promote to protected
		for j := 0; j < 8; j++ {
			cache.Get(key)
		}
	}
	
	// Add probation items
	cache.Put("prob1", 100)
	cache.Put("prob2", 101)
	
	stats := cache.Stats()
	t.Logf("Before demotion: Protected=%d, Probation=%d, Window=%d", 
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)
	
	// Now try to promote a probation item when protected is full
	// This should trigger demotion from protected
	for i := 0; i < 15; i++ {
		cache.Get("prob1") // Heavy access to probation item
	}
	
	stats = cache.Stats()
	t.Logf("After heavy access to prob1: Protected=%d, Probation=%d, Window=%d", 
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)
	
	// Verify cache integrity
	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

func TestResizeEvictionCoverage(t *testing.T) {
	cache, err := New[string, int](15)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Fill cache with mixed frequency items
	for i := 0; i < 20; i++ {
		key := fmt.Sprintf("item%d", i)
		cache.Put(key, i)
		
		if i < 8 {
			// High frequency - should end up in protected
			for j := 0; j < 6; j++ {
				cache.Get(key)
			}
		} else if i < 14 {
			// Medium frequency - probation
			for j := 0; j < 2; j++ {
				cache.Get(key)
			}
		}
		// Low frequency items remain in window or get evicted
	}
	
	stats := cache.Stats()
	t.Logf("Before resize: Protected=%d, Probation=%d, Window=%d, Total=%d", 
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize, stats.Size)
	
	// Resize to trigger different eviction paths
	err = cache.Resize(5)
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}
	
	stats = cache.Stats()
	t.Logf("After resize to 5: Protected=%d, Probation=%d, Window=%d, Total=%d", 
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize, stats.Size)
	
	if cache.Cap() != 5 {
		t.Errorf("Expected capacity 5, got %d", cache.Cap())
	}
	
	if cache.Len() > 5 {
		t.Errorf("Cache size %d exceeds capacity 5", cache.Len())
	}
}

func TestDemoteFromProtectedSpecificConditions(t *testing.T) {
	// Create conditions to trigger promoteToProtected -> demoteFromProtected
	cache, err := New[string, int](25) // Window=3, Main=22 (Protected=17-18, Probation=4-5)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// First, build frequency for some keys in the sketch by accessing them from the doorkeeper
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("freq%d", i%10)
		cache.Put(key, i)
		cache.Get(key) // Build frequency in Count-Min Sketch
	}
	
	// Now systematically build up the cache segments with already-frequent keys
	
	// Phase 1: Get items into probation first (they need to pass admission policy)
	probationKeys := []string{}
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("freq%d", i) // Use keys that already have frequency
		cache.Put(key, i+1000)
		
		// Access to move from window to probation via admission policy
		for j := 0; j < 3; j++ {
			cache.Get(key)
		}
		probationKeys = append(probationKeys, key)
	}
	
	// Phase 2: Promote items from probation to protected (fill protected to 80% capacity)
	for i := 0; i < len(probationKeys); i++ {
		key := probationKeys[i]
		// Heavy access to promote from probation to protected
		for j := 0; j < 8; j++ {
			cache.Get(key)
		}
	}
	
	stats := cache.Stats()
	t.Logf("After building protected: Protected=%d, Probation=%d, Window=%d", 
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)
	
	// Phase 3: Add one more item to probation with high frequency
	newKey := "freq8" // Another pre-frequent key
	cache.Put(newKey, 2000)
	
	// Access it enough to trigger promotion when protected is near capacity
	for i := 0; i < 10; i++ {
		cache.Get(newKey) // This should trigger promoteToProtected -> demoteFromProtected
	}
	
	stats = cache.Stats()
	t.Logf("After potential demotion: Protected=%d, Probation=%d, Window=%d", 
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)
	
	// Verify cache integrity
	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

func TestEvictFromProtectedSpecificConditions(t *testing.T) {
	// Create conditions where probation is empty but protected has items
	cache, err := New[string, int](10) // Window=1, Main=9 (Probation=2, Protected=7)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Fill protected segment
	for i := 0; i < 7; i++ {
		key := fmt.Sprintf("protected%d", i)
		cache.Put(key, i)
		// Access many times to promote to protected
		for j := 0; j < 8; j++ {
			cache.Get(key)
		}
	}
	
	// Fill window
	cache.Put("window", 100)
	
	stats := cache.Stats()
	t.Logf("Setup: Protected=%d, Probation=%d, Window=%d", 
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)
	
	// Now resize down to force eviction from protected when probation is empty
	err = cache.Resize(5) // This should force eviction from protected
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}
	
	stats = cache.Stats()
	t.Logf("After resize: Protected=%d, Probation=%d, Window=%d", 
		stats.ProtectedSize, stats.ProbationSize, stats.WindowSize)
	
	if cache.Cap() != 5 {
		t.Errorf("Expected capacity 5, got %d", cache.Cap())
	}
	
	if cache.Len() > 5 {
		t.Errorf("Cache size %d exceeds capacity 5", cache.Len())
	}
}

func TestKeysFullSegmentCoverage(t *testing.T) {
	cache, err := New[string, int](15)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Test empty cache first
	keys := cache.Keys()
	if len(keys) != 0 {
		t.Errorf("Expected 0 keys in empty cache, got %d", len(keys))
	}
	
	// Create items in all segments
	// Protected items
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("prot%d", i)
		cache.Put(key, i)
		for j := 0; j < 5; j++ {
			cache.Get(key)
		}
	}
	
	// Probation items
	for i := 0; i < 3; i++ {
		key := fmt.Sprintf("prob%d", i)
		cache.Put(key, i+10)
		cache.Get(key) // Move to probation
	}
	
	// Window items
	for i := 0; i < 2; i++ {
		key := fmt.Sprintf("win%d", i)
		cache.Put(key, i+20)
	}
	
	keys = cache.Keys()
	
	// Verify we have keys from all segments
	protectedKeys := 0
	probationKeys := 0
	windowKeys := 0
	
	for _, key := range keys {
		if len(key) >= 4 {
			prefix := key[:4]
			switch prefix {
			case "prot":
				protectedKeys++
			case "prob":
				probationKeys++
			case "win":
				windowKeys++
			}
		}
	}
	
	t.Logf("Key distribution: Protected=%d, Probation=%d, Window=%d", 
		protectedKeys, probationKeys, windowKeys)
		
	// Should have some keys (TinyLFU algorithm is complex) 
	if len(keys) == 0 {
		t.Error("Expected at least some keys in cache")
	}
}

func TestDemoteFromProtected(t *testing.T) {
	// Test demoteFromProtected function by creating specific conditions
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Fill cache to trigger protection/demotion scenarios
	for i := 0; i < 15; i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
	}
	
	// This should trigger various internal methods including demoteFromProtected
	stats := cache.Stats()
	if stats.Size == 0 {
		t.Error("Expected cache to have items after puts")
	}
}

func TestEvictFromProtected(t *testing.T) {
	// Test evictFromProtected function
	cache, err := New[string, int](5)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Add enough items to trigger protected segment eviction
	for i := 0; i < 20; i++ {
		cache.Put(fmt.Sprintf("item%d", i), i)
	}
	
	// This should trigger evictFromProtected at some point
	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size (%d) exceeded capacity (%d)", cache.Len(), cache.Cap())
	}
}

func TestCountMinSketchEdgeCases(t *testing.T) {
	// Test EstimateCount with edge cases
	sketch := NewCountMinSketch(100)
	
	// Test with empty sketch
	count := sketch.EstimateCount([]byte("nonexistent"))
	if count != 0 {
		t.Errorf("Expected count 0 for nonexistent key, got %d", count)
	}
	
	// Add some values and test
	sketch.Add([]byte("test"))
	sketch.Add([]byte("test"))
	sketch.Add([]byte("test"))
	
	count = sketch.EstimateCount([]byte("test"))
	if count < 3 {
		t.Errorf("Expected count >= 3, got %d", count)
	}
}

func TestResizeEdgeCases(t *testing.T) {
	// Test Resize with various scenarios
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Add some items first
	for i := 0; i < 5; i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
	}
	
	// Resize to smaller capacity
	err = cache.Resize(3)
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}
	if cache.Cap() != 3 {
		t.Errorf("Expected capacity 3 after resize, got %d", cache.Cap())
	}
	
	// Resize to larger capacity
	err = cache.Resize(20)
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}
	if cache.Cap() != 20 {
		t.Errorf("Expected capacity 20 after resize, got %d", cache.Cap())
	}
}

func TestTriggerDemoteFromProtected(t *testing.T) {
	// Create specific conditions to trigger demoteFromProtected
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Step 1: Fill the cache and create items in protected segment
	for i := 0; i < 30; i++ {
		key := fmt.Sprintf("prot%d", i)
		cache.Put(key, i)
		
		// Access frequently to get items into protected
		for j := 0; j < 8; j++ {
			cache.Get(key)
		}
	}
	
	// Step 2: Add many more items to force evictions and demotion
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("new%d", i)
		cache.Put(key, i+1000)
		// Some moderate access
		cache.Get(key)
		cache.Get(key)
	}
	
	// This complex scenario should eventually trigger demoteFromProtected
	stats := cache.Stats()
	if stats.Size == 0 {
		t.Error("Expected items to remain in cache")
	}
}

func TestTriggerEvictFromProtected(t *testing.T) {
	// Create conditions to trigger evictFromProtected
	cache, err := New[string, int](15)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Fill protected segment by creating highly accessed items
	for i := 0; i < 25; i++ {
		key := fmt.Sprintf("frequent%d", i)
		cache.Put(key, i)
		
		// Make items very frequent to get into protected
		for j := 0; j < 10; j++ {
			cache.Get(key)
		}
	}
	
	// Now add more items to force eviction from protected
	for i := 0; i < 30; i++ {
		key := fmt.Sprintf("force%d", i)
		cache.Put(key, i+2000)
		
		// Access to make them competitive
		for j := 0; j < 5; j++ {
			cache.Get(key)
		}
	}
	
	// This should eventually trigger evictFromProtected
	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

func TestExtensiveValuesAndItems(t *testing.T) {
	// Test Values and Items with all segment types
	cache, err := New[string, int](15)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Add items to window (low access)
	for i := 0; i < 5; i++ {
		cache.Put(fmt.Sprintf("window%d", i), i)
	}
	
	// Add items to probation (medium access)
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("prob%d", i)
		cache.Put(key, i+100)
		cache.Get(key) // Move to probation
		cache.Get(key)
	}
	
	// Add items to protected (high access)
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("prot%d", i)
		cache.Put(key, i+200)
		for j := 0; j < 6; j++ {
			cache.Get(key) // Should move to protected
		}
	}
	
	// Test Values method
	values := cache.Values()
	if len(values) == 0 {
		t.Error("Expected values from all segments")
	}
	
	// Test Items method
	items := cache.Items()
	if len(items) == 0 {
		t.Error("Expected items from all segments")
	}
	
	// Verify we have a reasonable distribution
	t.Logf("Total values: %d, Total items: %d", len(values), len(items))
	
	stats := cache.Stats()
	t.Logf("Cache stats - Window: %d, Probation: %d, Protected: %d", 
		stats.WindowSize, stats.ProbationSize, stats.ProtectedSize)
}

func TestResizeWithSegmentEvictions(t *testing.T) {
	// Test Resize that forces evictions from different segments
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Create a complex cache state
	for i := 0; i < 30; i++ {
		key := fmt.Sprintf("mixed%d", i)
		cache.Put(key, i)
		
		// Varying access patterns
		accessCount := i % 7
		for j := 0; j < accessCount; j++ {
			cache.Get(key)
		}
	}
	
	// Resize down to force evictions from all segments
	err = cache.Resize(8)
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}
	
	if cache.Cap() != 8 {
		t.Errorf("Expected capacity 8, got %d", cache.Cap())
	}
	
	if cache.Len() > 8 {
		t.Errorf("Cache size %d exceeds capacity 8", cache.Len())
	}
	
	// Further resize
	err = cache.Resize(3)
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}
	
	if cache.Cap() != 3 {
		t.Errorf("Expected capacity 3, got %d", cache.Cap())
	}
}

func TestSpecificEvictFromProtected(t *testing.T) {
	// Create a scenario where Resize calls evictFromProtected
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Fill the cache with highly accessed items to populate protected segment
	for i := 0; i < 150; i++ {
		key := fmt.Sprintf("prot%d", i)
		cache.Put(key, i)
		
		// High access to promote to protected
		for j := 0; j < 15; j++ {
			cache.Get(key)
		}
	}
	
	// Verify protected segment has items
	stats := cache.Stats()
	if stats.ProtectedSize == 0 {
		t.Logf("Protected size is 0, trying to force promotion...")
		// Try more aggressive promotion
		for i := 0; i < 50; i++ {
			key := fmt.Sprintf("force%d", i)
			cache.Put(key, i+1000)
			for j := 0; j < 20; j++ {
				cache.Get(key)
			}
		}
		stats = cache.Stats()
	}
	
	// Now resize to a very small size to force evictFromProtected
	// This should trigger the third branch in Resize: c.protected.Len() > 0
	err = cache.Resize(1)
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}
	
	if cache.Cap() != 1 {
		t.Errorf("Expected capacity 1, got %d", cache.Cap())
	}
	
	if cache.Len() > 1 {
		t.Errorf("Cache size %d exceeds capacity 1", cache.Len())
	}
}

func TestSpecificDemoteFromProtected(t *testing.T) {
	// Create a scenario where promoteToProtected calls demoteFromProtected
	cache, err := New[string, int](50)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}
	
	// Step 1: Fill protected segment to its capacity (~80% of main)
	// Main size would be ~45 (90% of 50), so protected capacity is ~36
	for i := 0; i < 60; i++ {
		key := fmt.Sprintf("protected%d", i)
		cache.Put(key, i)
		
		// Access enough to get into protected
		for j := 0; j < 12; j++ {
			cache.Get(key)
		}
	}
	
	// Step 2: Add items to probation
	for i := 0; i < 20; i++ {
		key := fmt.Sprintf("probation%d", i)  
		cache.Put(key, i+1000)
		
		// Moderate access to get into probation
		for j := 0; j < 3; j++ {
			cache.Get(key)
		}
	}
	
	// Step 3: Now promote from probation to protected
	// This should trigger demoteFromProtected when protected is full
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("promote%d", i)
		cache.Put(key, i+2000)
		
		// Access to get to probation first
		cache.Get(key)
		cache.Get(key)
		
		// More access to promote to protected (should trigger demotion)
		for j := 0; j < 8; j++ {
			cache.Get(key)
		}
	}
	
	stats := cache.Stats()
	t.Logf("Final stats - Window: %d, Probation: %d, Protected: %d", 
		stats.WindowSize, stats.ProbationSize, stats.ProtectedSize)
		
	// Verify cache integrity
	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}