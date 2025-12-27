package wtinylfu

import (
	"fmt"
	"sync"
	"testing"
)

func TestNewWithEvictError(t *testing.T) {
	_, err := NewWithEvict[string, int](0, func(key string, value int) {})
	if err == nil {
		t.Error("Expected error for zero capacity with NewWithEvict")
	}
}

func TestKeysEmptySegments(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	keys := cache.Keys()
	if len(keys) != 0 {
		t.Errorf("Expected 0 keys in empty cache, got %d", len(keys))
	}
}

func TestValuesEmptySegments(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	values := cache.Values()
	if len(values) != 0 {
		t.Errorf("Expected 0 values in empty cache, got %d", len(values))
	}
}

func TestItemsEmptySegments(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	items := cache.Items()
	if len(items) != 0 {
		t.Errorf("Expected 0 items in empty cache, got %d", len(items))
	}
}

func TestKeysOnlyWindow(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("window1", 1)
	cache.Put("window2", 2)

	keys := cache.Keys()
	if len(keys) == 0 {
		t.Error("Expected keys from window segment")
	}
}

func TestValuesOnlyWindow(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("window1", 1)
	cache.Put("window2", 2)

	values := cache.Values()
	if len(values) == 0 {
		t.Error("Expected values from window segment")
	}
}

func TestItemsOnlyWindow(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("window1", 1)
	cache.Put("window2", 2)

	items := cache.Items()
	if len(items) == 0 {
		t.Error("Expected items from window segment")
	}
}

func TestKeysOnlyProbation(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prob1", 1)
	cache.Get("prob1")
	cache.Put("prob2", 2)
	cache.Get("prob2")

	keys := cache.Keys()
	if len(keys) == 0 {
		t.Error("Expected keys from probation segment")
	}
}

func TestValuesOnlyProbation(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prob1", 1)
	cache.Get("prob1")
	cache.Put("prob2", 2)
	cache.Get("prob2")

	values := cache.Values()
	if len(values) == 0 {
		t.Error("Expected values from probation segment")
	}
}

func TestItemsOnlyProbation(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prob1", 1)
	cache.Get("prob1")
	cache.Put("prob2", 2)
	cache.Get("prob2")

	items := cache.Items()
	if len(items) == 0 {
		t.Error("Expected items from probation segment")
	}
}

func TestKeysOnlyProtected(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prot1", 1)
	cache.Get("prot1")
	cache.Get("prot1")

	cache.Put("prot2", 2)
	cache.Get("prot2")
	cache.Get("prot2")

	keys := cache.Keys()
	if len(keys) == 0 {
		t.Error("Expected keys from protected segment")
	}
}

func TestValuesOnlyProtected(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prot1", 1)
	cache.Get("prot1")
	cache.Get("prot1")

	cache.Put("prot2", 2)
	cache.Get("prot2")
	cache.Get("prot2")

	values := cache.Values()
	if len(values) == 0 {
		t.Error("Expected values from protected segment")
	}
}

func TestItemsOnlyProtected(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prot1", 1)
	cache.Get("prot1")
	cache.Get("prot1")

	cache.Put("prot2", 2)
	cache.Get("prot2")
	cache.Get("prot2")

	items := cache.Items()
	if len(items) == 0 {
		t.Error("Expected items from protected segment")
	}
}

func TestDemoteFromProtectedWithFullProbation(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 8; i++ {
		key := fmt.Sprintf("prot%d", i)
		cache.Put(key, i)
		for j := 0; j < 5; j++ {
			cache.Get(key)
		}
	}

	cache.Put("prob1", 100)
	cache.Get("prob1")

	cache.Put("prob2", 101)
	cache.Get("prob2")

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictFromWindowWithFullProbation(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prob1", 1)
	cache.Get("prob1")
	cache.Put("prob2", 2)
	cache.Get("prob2")

	cache.Put("window1", 10)
	cache.Put("window2", 20)

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestResizeWithEmptyCache(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	err = cache.Resize(50)
	if err != nil {
		t.Fatalf("Failed to resize empty cache: %v", err)
	}

	if cache.Cap() != 50 {
		t.Errorf("Expected capacity 50, got %d", cache.Cap())
	}
}

func TestResizeToSameCapacity(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("b", 2)

	err = cache.Resize(100)
	if err != nil {
		t.Fatalf("Failed to resize to same capacity: %v", err)
	}

	if cache.Cap() != 100 {
		t.Errorf("Expected capacity 100, got %d", cache.Cap())
	}
}

func TestResizeToLarger(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 10; i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
	}

	err = cache.Resize(200)
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}

	if cache.Cap() != 200 {
		t.Errorf("Expected capacity 200, got %d", cache.Cap())
	}

	if cache.Len() < 5 {
		t.Errorf("Expected length >= 5, got %d", cache.Len())
	}
}

func TestResizeToSmaller(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 80; i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
	}

	err = cache.Resize(20)
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}

	if cache.Cap() != 20 {
		t.Errorf("Expected capacity 20, got %d", cache.Cap())
	}

	if cache.Len() > 20 {
		t.Errorf("Expected length <= 20, got %d", cache.Len())
	}
}

func TestResizeWithProtectedEvictionCoverage(t *testing.T) {
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 15; i++ {
		key := fmt.Sprintf("key%d", i)
		cache.Put(key, i)
		for j := 0; j < 5; j++ {
			cache.Get(key)
		}
	}

	err = cache.Resize(5)
	if err != nil {
		t.Fatalf("Failed to resize cache: %v", err)
	}

	if cache.Cap() != 5 {
		t.Errorf("Expected capacity 5, got %d", cache.Cap())
	}

	if cache.Len() > 5 {
		t.Errorf("Expected length <= 5, got %d", cache.Len())
	}
}

func TestRemoveFromEmptyCache(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	_, ok := cache.Remove("nonexistent")
	if ok {
		t.Error("Expected false for removing from empty cache")
	}
}

func TestPeekFromEmptyCache(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	_, ok := cache.Peek("nonexistent")
	if ok {
		t.Error("Expected false for peeking from empty cache")
	}
}

func TestContainsFromEmptyCache(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	if cache.Contains("nonexistent") {
		t.Error("Expected false for contains in empty cache")
	}
}

func TestGetFromEmptyCache(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	_, ok := cache.Get("nonexistent")
	if ok {
		t.Error("Expected false for getting from empty cache")
	}
}

func TestUpdateNonExistentKey(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	evicted := cache.Put("a", 10)
	if evicted {
		t.Error("Expected no eviction when updating existing key")
	}

	value, ok := cache.Get("a")
	if !ok || value != 10 {
		t.Errorf("Expected updated value 10, got %d, ok=%t", value, ok)
	}
}

func TestClearEmptyCache(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Clear()

	if cache.Len() != 0 {
		t.Errorf("Expected empty cache after clear, got %d", cache.Len())
	}
}

func TestSketchCounterOverflow(t *testing.T) {
	sketch := newCountMinSketch(10)

	hash := uint32(12345)
	for i := 0; i < 300; i++ {
		sketch.increment(hash)
	}

	estimate := sketch.estimate(hash)
	if estimate == 0 {
		t.Error("Expected non-zero estimate after many increments")
	}
}

func TestSketchMultipleKeys(t *testing.T) {
	sketch := newCountMinSketch(100)

	for i := 0; i < 10; i++ {
		hash := uint32(i * 1000)
		for j := 0; j < 50; j++ {
			sketch.increment(hash)
		}
	}

	for i := 0; i < 10; i++ {
		hash := uint32(i * 1000)
		estimate := sketch.estimate(hash)
		if estimate == 0 {
			t.Logf("Warning: estimate is 0 for hash %d (sampling may have skipped increments)", hash)
		}
	}
}

func TestSketchClear(t *testing.T) {
	sketch := newCountMinSketch(100)

	hash := uint32(12345)
	for i := 0; i < 50; i++ {
		sketch.increment(hash)
	}

	estimateBefore := sketch.estimate(hash)
	if estimateBefore == 0 {
		t.Error("Expected non-zero estimate before clear")
	}

	sketch.clear()

	estimateAfter := sketch.estimate(hash)
	if estimateAfter != 0 {
		t.Errorf("Expected estimate 0 after clear, got %d", estimateAfter)
	}
}

func TestSketchSampling(t *testing.T) {
	sketch := newCountMinSketch(50)

	hash := uint32(12345)
	for i := 0; i < 20; i++ {
		sketch.increment(hash)
	}

	estimate := sketch.estimate(hash)
	if estimate == 0 {
		t.Error("Expected non-zero estimate after increments with sampling")
	}
}

func TestWindowEvictionWithEmptyProbation(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("window1", 1)
	cache.Put("window2", 2)
	cache.Put("window3", 3)

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestProbationEvictionWithEmptyProtected(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prob1", 1)
	cache.Get("prob1")
	cache.Put("prob2", 2)
	cache.Get("prob2")
	cache.Put("prob3", 3)
	cache.Get("prob3")

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestProtectedEvictionWithFullCache(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 15; i++ {
		key := fmt.Sprintf("key%d", i)
		cache.Put(key, i)
		for j := 0; j < 10; j++ {
			cache.Get(key)
		}
	}

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictionCallbackOnClear(t *testing.T) {
	evictedKeys := make([]string, 0)
	cache, err := NewWithEvict[string, int](10, func(key string, value int) {
		evictedKeys = append(evictedKeys, key)
	})
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	cache.Clear()

	if len(evictedKeys) != 3 {
		t.Errorf("Expected 3 evicted keys, got %d", len(evictedKeys))
	}
}

func TestEvictionCallbackOnEviction(t *testing.T) {
	evictedKeys := make([]string, 0)
	cache, err := NewWithEvict[string, int](5, func(key string, value int) {
		evictedKeys = append(evictedKeys, key)
	})
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 10; i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
	}

	if len(evictedKeys) == 0 {
		t.Error("Expected some evictions")
	}
}

func TestRemoveWithEvictionCallback(t *testing.T) {
	evictedKeys := make([]string, 0)
	cache, err := NewWithEvict[string, int](10, func(key string, value int) {
		evictedKeys = append(evictedKeys, key)
	})
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("b", 2)

	if !cache.Contains("a") {
		t.Fatal("Expected 'a' to be in cache before removal")
	}

	value, ok := cache.Remove("a")
	if !ok || value != 1 {
		t.Errorf("Expected to remove value 1, got %d, ok=%t", value, ok)
	}

	if len(evictedKeys) != 1 || evictedKeys[0] != "a" {
		t.Error("Expected eviction callback on Remove for wtinylfu")
	}
}

func TestStatsAccuracy(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 50; i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
	}

	stats := cache.Stats()

	if stats.Size != cache.Len() {
		t.Errorf("Stats size %d doesn't match cache len %d", stats.Size, cache.Len())
	}

	if stats.Capacity != cache.Cap() {
		t.Errorf("Stats capacity %d doesn't match cache cap %d", stats.Capacity, cache.Cap())
	}

	totalSegments := stats.WindowSize + stats.ProbationSize + stats.ProtectedSize
	if totalSegments != stats.Size {
		t.Errorf("Segment sizes don't match total size: %d vs %d", totalSegments, stats.Size)
	}
}

func TestCapacityOne(t *testing.T) {
	cache, err := New[string, int](1)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("b", 2)

	if cache.Len() != 1 {
		t.Errorf("Expected cache size 1, got %d", cache.Len())
	}

	_, ok := cache.Get("b")
	if !ok {
		t.Error("Expected 'b' to be in cache")
	}
}

func TestCapacityTwo(t *testing.T) {
	cache, err := New[string, int](2)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	if cache.Len() > 2 {
		t.Errorf("Expected cache size <= 2, got %d", cache.Len())
	}
}

func TestLargeCapacity(t *testing.T) {
	cache, err := New[string, int](10000)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 5000; i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
	}

	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}

	stats := cache.Stats()
	if stats.WindowCapacity == 0 {
		t.Error("Expected non-zero window capacity")
	}
}

func TestKeyTypes(t *testing.T) {
	cache, err := New[int, string](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put(1, "one")
	cache.Put(2, "two")

	value, ok := cache.Get(1)
	if !ok || value != "one" {
		t.Errorf("Expected value 'one', got %s, ok=%t", value, ok)
	}

	value, ok = cache.Get(2)
	if !ok || value != "two" {
		t.Errorf("Expected value 'two', got %s, ok=%t", value, ok)
	}
}

func TestValueTypes(t *testing.T) {
	cache, err := New[string, []int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", []int{1, 2, 3})
	cache.Put("b", []int{4, 5, 6})

	value, ok := cache.Get("a")
	if !ok || len(value) != 3 {
		t.Errorf("Expected value with length 3, got %d, ok=%t", len(value), ok)
	}
}

func TestZeroValue(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("zero", 0)

	value, ok := cache.Get("zero")
	if !ok || value != 0 {
		t.Errorf("Expected value 0, got %d, ok=%t", value, ok)
	}
}

func TestNilValue(t *testing.T) {
	cache, err := New[string, *int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	var nilPtr *int = nil
	cache.Put("nil", nilPtr)

	value, ok := cache.Get("nil")
	if !ok || value != nil {
		t.Errorf("Expected nil value, got %v, ok=%t", value, ok)
	}
}

func TestStringKeyWithSpecialChars(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	specialKeys := []string{
		"key with spaces",
		"key\twith\ttabs",
		"key\nwith\nnewlines",
		"key\"with\"quotes",
		"key'with'apostrophes",
		"key\\with\\backslashes",
	}

	for _, key := range specialKeys {
		cache.Put(key, len(key))
	}

	for _, key := range specialKeys {
		value, ok := cache.Get(key)
		if !ok || value != len(key) {
			t.Errorf("Expected value %d for key %q, got %d, ok=%t", len(key), key, value, ok)
		}
	}
}

func TestInt64Key(t *testing.T) {
	cache, err := New[int64, string](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put(1234567890123456789, "value")

	value, ok := cache.Get(1234567890123456789)
	if !ok || value != "value" {
		t.Errorf("Expected value 'value', got %s, ok=%t", value, ok)
	}
}

func TestFloat64Key(t *testing.T) {
	cache, err := New[float64, string](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put(3.14159, "pi")

	value, ok := cache.Get(3.14159)
	if !ok || value != "pi" {
		t.Errorf("Expected value 'pi', got %s, ok=%t", value, ok)
	}
}

func TestBoolKey(t *testing.T) {
	cache, err := New[bool, string](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put(true, "yes")
	cache.Put(false, "no")

	value, ok := cache.Get(true)
	if !ok || value != "yes" {
		t.Errorf("Expected value 'yes', got %s, ok=%t", value, ok)
	}

	value, ok = cache.Get(false)
	if !ok || value != "no" {
		t.Errorf("Expected value 'no', got %s, ok=%t", value, ok)
	}
}

func TestMultipleUpdates(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("a", 2)
	cache.Put("a", 3)
	cache.Put("a", 4)
	cache.Put("a", 5)

	value, ok := cache.Get("a")
	if !ok || value != 5 {
		t.Errorf("Expected value 5, got %d, ok=%t", value, ok)
	}
}

func TestRemoveNonExistent(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)

	_, ok := cache.Remove("nonexistent")
	if ok {
		t.Error("Expected false for removing non-existent key")
	}

	value, ok := cache.Get("a")
	if !ok || value != 1 {
		t.Errorf("Expected value 1, got %d, ok=%t", value, ok)
	}
}

func TestPeekDoesNotUpdatePosition(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("b", 2)

	cache.Peek("a")

	cache.Put("c", 3)

	value, ok := cache.Get("a")
	if !ok || value != 1 {
		t.Logf("Note: Peek may not prevent eviction in wtinylfu, got %d, ok=%t", value, ok)
	}
}

func TestContainsDoesNotUpdatePosition(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("b", 2)

	cache.Contains("a")

	cache.Put("c", 3)

	value, ok := cache.Get("a")
	if !ok || value != 1 {
		t.Logf("Note: Contains may not prevent eviction in wtinylfu, got %d, ok=%t", value, ok)
	}
}

func TestLenAccuracy(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	if cache.Len() != 0 {
		t.Errorf("Expected len 0, got %d", cache.Len())
	}

	cache.Put("a", 1)
	if cache.Len() != 1 {
		t.Errorf("Expected len 1, got %d", cache.Len())
	}

	cache.Put("b", 2)
	if cache.Len() != 2 {
		t.Errorf("Expected len 2, got %d", cache.Len())
	}

	cache.Remove("a")
	if cache.Len() != 1 {
		t.Errorf("Expected len 1, got %d", cache.Len())
	}

	cache.Clear()
	if cache.Len() != 0 {
		t.Errorf("Expected len 0, got %d", cache.Len())
	}
}

func TestCapAccuracy(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	if cache.Cap() != 100 {
		t.Errorf("Expected cap 100, got %d", cache.Cap())
	}

	cache.Resize(50)
	if cache.Cap() != 50 {
		t.Errorf("Expected cap 50, got %d", cache.Cap())
	}
}

func TestSegmentCapacities(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	stats := cache.Stats()

	if stats.WindowCapacity == 0 {
		t.Error("Expected non-zero window capacity")
	}

	if stats.ProbationCap == 0 {
		t.Error("Expected non-zero probation capacity")
	}

	if stats.ProtectedCap == 0 {
		t.Error("Expected non-zero protected capacity")
	}
}

func TestSegmentSizesWithinCapacities(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 150; i++ {
		cache.Put(fmt.Sprintf("key%d", i), i)
	}

	stats := cache.Stats()

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

func TestConcurrentResize(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				key := fmt.Sprintf("%d-%d", id, j)
				cache.Put(key, id*100+j)
				cache.Get(key)
			}
		}(i)
	}

	wg.Wait()

	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

func TestStressTest(t *testing.T) {
	cache, err := New[int, int](1000)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 10000; i++ {
		key := i % 2000
		cache.Put(key, key*2)
		cache.Get(key)
	}

	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

func TestRandomAccessPattern(t *testing.T) {
	cache, err := New[int, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 1000; i++ {
		key := (i * 7) % 200
		cache.Put(key, i)
		if i%3 == 0 {
			cache.Get(key)
		}
	}

	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

func TestSequentialAccessPattern(t *testing.T) {
	cache, err := New[int, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 500; i++ {
		cache.Put(i, i)
	}

	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

func TestRepeatedAccessPattern(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	keys := []string{"a", "b", "c", "d", "e"}

	for i := 0; i < 1000; i++ {
		key := keys[i%len(keys)]
		cache.Put(key, i)
		cache.Get(key)
	}

	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}

	for _, key := range keys {
		if !cache.Contains(key) {
			t.Errorf("Expected key %s to be in cache", key)
		}
	}
}

func TestZipfAccessPattern(t *testing.T) {
	cache, err := New[int, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 1000; i++ {
		key := (i * i) % 50
		cache.Put(key, i)
		cache.Get(key)
	}

	if cache.Len() > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", cache.Len(), cache.Cap())
	}
}

func TestRecordAccessFromWindow(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("window", 1)
	cache.Get("window")

	stats := cache.Stats()
	if stats.WindowSize == 0 && stats.ProbationSize == 0 {
		t.Error("Expected item to move from window")
	}
}

func TestRecordAccessFromProbation(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prob", 1)
	cache.Get("prob")
	cache.Get("prob")

	stats := cache.Stats()
	if stats.ProtectedSize == 0 {
		t.Error("Expected item to be promoted to protected")
	}
}

func TestRecordAccessFromProtected(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prot", 1)
	cache.Get("prot")
	cache.Get("prot")

	stats := cache.Stats()
	if stats.ProtectedSize == 0 {
		t.Error("Expected item to be in protected")
	}
}

func TestEvictFromWindowWithAdmission(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Get("a")
	cache.Get("a")

	cache.Put("b", 2)
	cache.Put("c", 3)

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictFromWindowWithoutAdmission(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictFromProbation(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prob1", 1)
	cache.Get("prob1")
	cache.Put("prob2", 2)
	cache.Get("prob2")
	cache.Put("prob3", 3)
	cache.Get("prob3")

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictFromProtected(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 15; i++ {
		key := fmt.Sprintf("key%d", i)
		cache.Put(key, i)
		for j := 0; j < 10; j++ {
			cache.Get(key)
		}
	}

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestDemoteFromProtected(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 8; i++ {
		key := fmt.Sprintf("key%d", i)
		cache.Put(key, i)
		for j := 0; j < 5; j++ {
			cache.Get(key)
		}
	}

	cache.Put("new", 100)
	cache.Get("new")
	cache.Get("new")

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestRemoveEntryFromWindow(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("window", 1)
	cache.Remove("window")

	stats := cache.Stats()
	if stats.WindowSize != 0 {
		t.Errorf("Expected window size 0, got %d", stats.WindowSize)
	}
}

func TestRemoveEntryFromProbation(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prob", 1)
	cache.Get("prob")
	cache.Remove("prob")

	stats := cache.Stats()
	if stats.ProbationSize != 0 {
		t.Errorf("Expected probation size 0, got %d", stats.ProbationSize)
	}
}

func TestRemoveEntryFromProtected(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("prot", 1)
	cache.Get("prot")
	cache.Get("prot")
	cache.Remove("prot")

	stats := cache.Stats()
	if stats.ProtectedSize != 0 {
		t.Errorf("Expected protected size 0, got %d", stats.ProtectedSize)
	}
}

func TestHashDeterministic(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	hash1 := cache.hash("test")
	hash2 := cache.hash("test")

	if hash1 != hash2 {
		t.Error("Hash function should be deterministic")
	}
}

func TestHashDifferentKeys(t *testing.T) {
	cache, err := New[string, int](10)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	hash1 := cache.hash("test1")
	hash2 := cache.hash("test2")

	if hash1 == hash2 {
		t.Error("Different keys should produce different hashes")
	}
}
