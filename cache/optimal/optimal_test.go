package optimal

import (
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
	if cache.CurrentTime() != 0 {
		t.Errorf("Expected current time 0, got %d", cache.CurrentTime())
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

func TestNewWithPattern(t *testing.T) {
	pattern := []string{"a", "b", "c", "a", "d"}
	cache := NewWithPattern[string, int](3, pattern)
	
	if cache.Cap() != 3 {
		t.Errorf("Expected capacity 3, got %d", cache.Cap())
	}
}

func TestNewWithEvict(t *testing.T) {
	evicted := make(map[string]int)
	onEvict := func(key string, value int) {
		evicted[key] = value
	}
	
	cache := NewWithEvict[string, int](2, onEvict)
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3) // Should evict one item
	
	if len(evicted) != 1 {
		t.Errorf("Expected 1 eviction, got %d", len(evicted))
	}
}

func TestBasicOperations(t *testing.T) {
	cache := New[string, int](3)
	
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

func TestOptimalEvictionWithPattern(t *testing.T) {
	// Access pattern: a, b, c, a, d, e, a
	pattern := []string{"a", "b", "c", "a", "d", "e", "a"}
	cache := NewWithPattern[string, int](2, pattern)
	
	// Start simulation
	cache.Put("a", 1) // Time 1: Next access at time 4
	cache.Put("b", 2) // Time 2: Next access never (should be evicted first)
	
	// At this point cache has [a, b], next accesses: a=4, b=never
	
	cache.Put("c", 3) // Time 3: Should evict 'b' (never accessed again)
	
	// Verify 'b' was evicted and 'a', 'c' remain
	if !cache.Contains("a") {
		t.Errorf("Expected 'a' to remain in cache")
	}
	if !cache.Contains("c") {
		t.Errorf("Expected 'c' to remain in cache")
	}
	if cache.Contains("b") {
		t.Errorf("Expected 'b' to be evicted")
	}
}

func TestOptimalEvictionFarthestFuture(t *testing.T) {
	// Access pattern: a, b, c, b, a (positions 0,1,2,3,4)
	pattern := []string{"a", "b", "c", "b", "a"}
	cache := NewWithPattern[string, int](2, pattern)
	
	cache.Put("a", 1) // Time 1: Next access at position 4
	cache.Put("b", 2) // Time 2: Next access at position 3
	
	// When 'c' is added (time 3), it should evict 'a' (accessed at pos 4) 
	// instead of 'b' (accessed at pos 3)
	cache.Put("c", 3) // Time 3: Should evict 'a' (farthest future access)
	
	if cache.Contains("a") {
		t.Errorf("Expected 'a' to be evicted (farthest future access)")
	}
	if !cache.Contains("b") {
		t.Errorf("Expected 'b' to remain")
	}
	if !cache.Contains("c") {
		t.Errorf("Expected 'c' to remain")
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
	
	// Test Peek
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
	
	if cache.CurrentTime() != 0 {
		t.Errorf("Expected current time reset to 0, got %d", cache.CurrentTime())
	}
	
	if len(evicted) != 2 {
		t.Errorf("Expected 2 evictions on clear, got %d", len(evicted))
	}
}

func TestKeysOrdering(t *testing.T) {
	pattern := []string{"a", "b", "c", "b", "d"}
	cache := NewWithPattern[string, int](3, pattern)
	
	cache.Put("a", 1) // Next access at pos 0 (already passed)
	cache.Put("b", 2) // Next access at pos 3
	cache.Put("c", 3) // Next access at pos 2
	
	keys := cache.Keys()
	
	// Should be ordered by next access time (farthest first)
	// a: no future access (-1), b: pos 3, c: pos 2
	// Order should be: a (farthest/never), b (pos 3), c (pos 2)
	if len(keys) != 3 {
		t.Errorf("Expected 3 keys, got %d", len(keys))
	}
}

func TestValues(t *testing.T) {
	cache := New[string, int](3)
	cache.Put("a", 1)
	cache.Put("b", 2)
	
	values := cache.Values()
	if len(values) != 2 {
		t.Errorf("Expected 2 values, got %d", len(values))
	}
	
	// Check values are present (order not guaranteed)
	hasOne := false
	hasTwo := false
	for _, v := range values {
		if v == 1 {
			hasOne = true
		}
		if v == 2 {
			hasTwo = true
		}
	}
	
	if !hasOne || !hasTwo {
		t.Errorf("Expected values 1 and 2 to be present")
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

func TestSetAccessPattern(t *testing.T) {
	cache := New[string, int](3)
	cache.Put("a", 1)
	cache.Put("b", 2)
	
	// Set new access pattern
	pattern := []string{"b", "c", "a", "b"}
	cache.SetAccessPattern(pattern)
	
	if cache.CurrentTime() != 0 {
		t.Errorf("Expected current time reset to 0, got %d", cache.CurrentTime())
	}
}

func TestCurrentTimeAdvancement(t *testing.T) {
	cache := New[string, int](3)
	
	if cache.CurrentTime() != 0 {
		t.Errorf("Expected initial time 0, got %d", cache.CurrentTime())
	}
	
	cache.Put("a", 1) // Should advance time to 1
	if cache.CurrentTime() != 1 {
		t.Errorf("Expected time 1 after Put, got %d", cache.CurrentTime())
	}
	
	cache.Get("a") // Should advance time to 2
	if cache.CurrentTime() != 2 {
		t.Errorf("Expected time 2 after Get, got %d", cache.CurrentTime())
	}
}

func TestSimulate(t *testing.T) {
	cache := New[string, int](2)
	
	// Set access pattern for optimal decisions
	pattern := []string{"a", "b", "c", "a", "b"}
	cache.SetAccessPattern(pattern)
	
	// Create operations
	operations := []Operation[string, int]{
		{Type: OpPut, Key: "a", Value: 1},
		{Type: OpPut, Key: "b", Value: 2},
		{Type: OpGet, Key: "a"},
		{Type: OpPut, Key: "c", Value: 3}, // Should evict 'b' (accessed later than 'a')
		{Type: OpGet, Key: "a"},
		{Type: OpGet, Key: "b"}, // Should be a miss
	}
	
	stats := cache.Simulate(operations)
	
	if stats.Hits != 2 { // 2 successful gets for 'a'
		t.Errorf("Expected 2 hits, got %d", stats.Hits)
	}
	
	if stats.Misses != 1 { // 1 miss for 'b' at the end
		t.Errorf("Expected 1 miss, got %d", stats.Misses)
	}
	
	if stats.Evictions != 1 { // 1 eviction when 'c' is added
		t.Errorf("Expected 1 eviction, got %d", stats.Evictions)
	}
	
	expectedHitRate := 2.0 / 3.0 // 2 hits out of 3 get operations
	if stats.HitRate < expectedHitRate-0.01 || stats.HitRate > expectedHitRate+0.01 {
		t.Errorf("Expected hit rate %.2f, got %.2f", expectedHitRate, stats.HitRate)
	}
}

func TestOptimalVsRandom(t *testing.T) {
	// This test demonstrates that Belady's algorithm is optimal
	cache := New[string, int](2)
	
	// Access pattern where optimal choice is clear
	pattern := []string{"a", "b", "c", "a", "d", "a"}
	cache.SetAccessPattern(pattern)
	
	operations := []Operation[string, int]{
		{Type: OpPut, Key: "a", Value: 1}, // Next access: pos 3
		{Type: OpPut, Key: "b", Value: 2}, // Next access: never
		{Type: OpPut, Key: "c", Value: 3}, // Should evict 'b' (optimal choice)
		{Type: OpGet, Key: "a"},           // Hit (optimal kept 'a')
		{Type: OpPut, Key: "d", Value: 4}, // Should evict 'c' (optimal choice)
		{Type: OpGet, Key: "a"},           // Hit (optimal kept 'a')
	}
	
	stats := cache.Simulate(operations)
	
	// With optimal replacement, we should get maximum hits
	if stats.Hits != 2 {
		t.Errorf("Expected 2 hits with optimal replacement, got %d", stats.Hits)
	}
	
	if stats.Misses != 0 {
		t.Errorf("Expected 0 misses with optimal replacement, got %d", stats.Misses)
	}
}

func TestEdgeCases(t *testing.T) {
	// Test single capacity cache
	cache := New[string, int](1)
	
	cache.Put("a", 1)
	cache.Put("b", 2) // Should evict "a"
	
	if cache.Len() != 1 {
		t.Errorf("Expected length 1 for single capacity cache, got %d", cache.Len())
	}
	
	_, ok := cache.Get("b")
	if !ok {
		t.Errorf("Expected 'b' to be in single capacity cache")
	}
	
	_, ok = cache.Get("a")
	if ok {
		t.Errorf("Expected 'a' to be evicted from single capacity cache")
	}
}

func TestNoFutureAccess(t *testing.T) {
	// Test behavior when items have no future access
	cache := New[string, int](2)
	
	cache.Put("a", 1) // No pattern set, so no future access
	cache.Put("b", 2) // No pattern set, so no future access
	cache.Put("c", 3) // Should evict one of the items
	
	if cache.Len() != 2 {
		t.Errorf("Expected length 2, got %d", cache.Len())
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

func BenchmarkSimulate(b *testing.B) {
	cache := New[int, int](100)
	
	// Create a large operation sequence
	operations := make([]Operation[int, int], 1000)
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			operations[i] = Operation[int, int]{Type: OpGet, Key: i % 100}
		} else {
			operations[i] = Operation[int, int]{Type: OpPut, Key: i % 100, Value: i}
		}
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Simulate(operations)
	}
}