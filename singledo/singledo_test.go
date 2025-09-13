package singledo

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestNewSingle(t *testing.T) {
	wait := 100 * time.Millisecond
	single := NewSingle[int](wait)
	
	if single == nil {
		t.Error("NewSingle() returned nil")
	}
	
	if single.wait != wait {
		t.Errorf("NewSingle() wait = %v, expected %v", single.wait, wait)
	}
}

func TestSingleDo_BasicOperation(t *testing.T) {
	single := NewSingle[string](100 * time.Millisecond)
	
	callCount := int32(0)
	fn := func() (string, error) {
		atomic.AddInt32(&callCount, 1)
		return "result", nil
	}
	
	result, err := single.Do(fn)
	
	if err != nil {
		t.Errorf("Single.Do() error = %v", err)
	}
	
	if result != "result" {
		t.Errorf("Single.Do() result = %q, expected %q", result, "result")
	}
	
	if atomic.LoadInt32(&callCount) != 1 {
		t.Errorf("Function called %d times, expected 1", atomic.LoadInt32(&callCount))
	}
}

func TestSingleDo_WithError(t *testing.T) {
	single := NewSingle[int](100 * time.Millisecond)
	
	expectedError := errors.New("test error")
	fn := func() (int, error) {
		return 0, expectedError
	}
	
	result, err := single.Do(fn)
	
	if err != expectedError {
		t.Errorf("Single.Do() error = %v, expected %v", err, expectedError)
	}
	
	if result != 0 {
		t.Errorf("Single.Do() result = %d, expected 0", result)
	}
}

func TestSingleDo_CacheWithinWaitPeriod(t *testing.T) {
	wait := 200 * time.Millisecond
	single := NewSingle[int](wait)
	
	callCount := int32(0)
	fn := func() (int, error) {
		return int(atomic.AddInt32(&callCount, 1)), nil
	}
	
	// First call
	result1, err1 := single.Do(fn)
	if err1 != nil {
		t.Errorf("First Single.Do() error = %v", err1)
	}
	
	// Second call immediately after (should use cached result)
	result2, err2 := single.Do(fn)
	if err2 != nil {
		t.Errorf("Second Single.Do() error = %v", err2)
	}
	
	// Both results should be the same (cached)
	if result1 != result2 {
		t.Errorf("Cached result mismatch: first = %d, second = %d", result1, result2)
	}
	
	if result1 != 1 {
		t.Errorf("Result = %d, expected 1", result1)
	}
	
	// Function should have been called only once
	if atomic.LoadInt32(&callCount) != 1 {
		t.Errorf("Function called %d times, expected 1", atomic.LoadInt32(&callCount))
	}
}

func TestSingleDo_ExpiredCache(t *testing.T) {
	wait := 50 * time.Millisecond
	single := NewSingle[int](wait)
	
	callCount := int32(0)
	fn := func() (int, error) {
		return int(atomic.AddInt32(&callCount, 1)), nil
	}
	
	// First call
	result1, err1 := single.Do(fn)
	if err1 != nil {
		t.Errorf("First Single.Do() error = %v", err1)
	}
	
	// Wait for cache to expire
	time.Sleep(wait + 10*time.Millisecond)
	
	// Second call after cache expiry
	result2, err2 := single.Do(fn)
	if err2 != nil {
		t.Errorf("Second Single.Do() error = %v", err2)
	}
	
	// Results should be different now
	if result1 == result2 {
		t.Errorf("Results should be different after cache expiry: first = %d, second = %d", result1, result2)
	}
	
	if result1 != 1 || result2 != 2 {
		t.Errorf("Results = %d, %d, expected 1, 2", result1, result2)
	}
	
	// Function should have been called twice
	if atomic.LoadInt32(&callCount) != 2 {
		t.Errorf("Function called %d times, expected 2", atomic.LoadInt32(&callCount))
	}
}

func TestSingleDo_ConcurrentCalls(t *testing.T) {
	single := NewSingle[int](100 * time.Millisecond)
	
	callCount := int32(0)
	callStarted := make(chan struct{})
	callCanFinish := make(chan struct{})
	
	fn := func() (int, error) {
		count := atomic.AddInt32(&callCount, 1)
		close(callStarted)
		<-callCanFinish
		return int(count), nil
	}
	
	var wg sync.WaitGroup
	results := make([]int, 5)
	errors := make([]error, 5)
	
	// Start 5 concurrent calls
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			results[index], errors[index] = single.Do(fn)
		}(i)
	}
	
	// Wait for the first call to start
	<-callStarted
	
	// Allow the function to finish
	close(callCanFinish)
	
	// Wait for all goroutines to finish
	wg.Wait()
	
	// All results should be the same (from single execution)
	expectedResult := 1
	for i := 0; i < 5; i++ {
		if errors[i] != nil {
			t.Errorf("Goroutine %d error = %v", i, errors[i])
		}
		if results[i] != expectedResult {
			t.Errorf("Goroutine %d result = %d, expected %d", i, results[i], expectedResult)
		}
	}
	
	// Function should have been called only once
	if atomic.LoadInt32(&callCount) != 1 {
		t.Errorf("Function called %d times, expected 1", atomic.LoadInt32(&callCount))
	}
}

func TestSingleDo_ErrorKeepsCall(t *testing.T) {
	single := NewSingle[int](100 * time.Millisecond)
	
	callCount := int32(0)
	fn := func() (int, error) {
		atomic.AddInt32(&callCount, 1)
		return 0, errors.New("persistent error")
	}
	
	// First call should return error
	result1, err1 := single.Do(fn)
	if err1 == nil {
		t.Error("First call should have returned error")
	}
	if result1 != 0 {
		t.Errorf("First result = %d, expected 0", result1)
	}
	
	// Second call immediately after should return same error (call object persists)
	result2, err2 := single.Do(fn)
	if err2 == nil {
		t.Error("Second call should have returned error")
	}
	if result2 != 0 {
		t.Errorf("Second result = %d, expected 0", result2)
	}
	
	// Function should have been called only once (second call waits on first)
	if atomic.LoadInt32(&callCount) != 1 {
		t.Errorf("Function called %d times, expected 1", atomic.LoadInt32(&callCount))
	}
}

func TestSingleReset(t *testing.T) {
	wait := 200 * time.Millisecond
	single := NewSingle[int](wait)
	
	callCount := int32(0)
	fn := func() (int, error) {
		return int(atomic.AddInt32(&callCount, 1)), nil
	}
	
	// First call
	result1, err1 := single.Do(fn)
	if err1 != nil {
		t.Errorf("First Single.Do() error = %v", err1)
	}
	
	// Reset the cache
	single.Reset()
	
	// Second call immediately after reset (should not use cache)
	result2, err2 := single.Do(fn)
	if err2 != nil {
		t.Errorf("Second Single.Do() error = %v", err2)
	}
	
	// Results should be different after reset
	if result1 == result2 {
		t.Errorf("Results should be different after reset: first = %d, second = %d", result1, result2)
	}
	
	if result1 != 1 || result2 != 2 {
		t.Errorf("Results = %d, %d, expected 1, 2", result1, result2)
	}
	
	// Function should have been called twice
	if atomic.LoadInt32(&callCount) != 2 {
		t.Errorf("Function called %d times, expected 2", atomic.LoadInt32(&callCount))
	}
}

func TestNewSingleGroup(t *testing.T) {
	wait := 100 * time.Millisecond
	group := NewSingleGroup[string](wait)
	
	if group == nil {
		t.Error("NewSingleGroup() returned nil")
	}
	
	if group.wait != wait {
		t.Errorf("NewSingleGroup() wait = %v, expected %v", group.wait, wait)
	}
	
	if group.singleMap == nil {
		t.Error("NewSingleGroup() singleMap is nil")
	}
}

func TestGroupDo_BasicOperation(t *testing.T) {
	group := NewSingleGroup[int](100 * time.Millisecond)
	
	callCount := int32(0)
	fn := func() (int, error) {
		return int(atomic.AddInt32(&callCount, 1)), nil
	}
	
	result, err := group.Do("key1", fn)
	
	if err != nil {
		t.Errorf("Group.Do() error = %v", err)
	}
	
	if result != 1 {
		t.Errorf("Group.Do() result = %d, expected 1", result)
	}
	
	if atomic.LoadInt32(&callCount) != 1 {
		t.Errorf("Function called %d times, expected 1", atomic.LoadInt32(&callCount))
	}
}

func TestGroupDo_DifferentKeys(t *testing.T) {
	group := NewSingleGroup[int](100 * time.Millisecond)
	
	callCount := int32(0)
	fn := func() (int, error) {
		return int(atomic.AddInt32(&callCount, 1)), nil
	}
	
	// Call with different keys should execute function for each key
	result1, err1 := group.Do("key1", fn)
	result2, err2 := group.Do("key2", fn)
	
	if err1 != nil {
		t.Errorf("First Group.Do() error = %v", err1)
	}
	if err2 != nil {
		t.Errorf("Second Group.Do() error = %v", err2)
	}
	
	if result1 != 1 || result2 != 2 {
		t.Errorf("Results = %d, %d, expected 1, 2", result1, result2)
	}
	
	// Function should have been called twice (once per key)
	if atomic.LoadInt32(&callCount) != 2 {
		t.Errorf("Function called %d times, expected 2", atomic.LoadInt32(&callCount))
	}
}

func TestGroupDo_SameKey(t *testing.T) {
	group := NewSingleGroup[int](100 * time.Millisecond)
	
	callCount := int32(0)
	fn := func() (int, error) {
		return int(atomic.AddInt32(&callCount, 1)), nil
	}
	
	// Call with same key should cache result
	result1, err1 := group.Do("key1", fn)
	result2, err2 := group.Do("key1", fn)
	
	if err1 != nil {
		t.Errorf("First Group.Do() error = %v", err1)
	}
	if err2 != nil {
		t.Errorf("Second Group.Do() error = %v", err2)
	}
	
	if result1 != result2 {
		t.Errorf("Results should be same for same key: first = %d, second = %d", result1, result2)
	}
	
	if result1 != 1 {
		t.Errorf("Result = %d, expected 1", result1)
	}
	
	// Function should have been called only once
	if atomic.LoadInt32(&callCount) != 1 {
		t.Errorf("Function called %d times, expected 1", atomic.LoadInt32(&callCount))
	}
}

func TestGroupDo_ConcurrentSameKey(t *testing.T) {
	group := NewSingleGroup[int](100 * time.Millisecond)
	
	callCount := int32(0)
	callStarted := make(chan struct{})
	callCanFinish := make(chan struct{})
	
	fn := func() (int, error) {
		count := atomic.AddInt32(&callCount, 1)
		if count == 1 {
			close(callStarted)
		}
		<-callCanFinish
		return int(count), nil
	}
	
	var wg sync.WaitGroup
	results := make([]int, 5)
	errors := make([]error, 5)
	
	// Start 5 concurrent calls with same key
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			results[index], errors[index] = group.Do("same-key", fn)
		}(i)
	}
	
	// Wait for the first call to start
	<-callStarted
	
	// Allow the function to finish
	close(callCanFinish)
	
	// Wait for all goroutines to finish
	wg.Wait()
	
	// All results should be the same (from single execution)
	expectedResult := 1
	for i := 0; i < 5; i++ {
		if errors[i] != nil {
			t.Errorf("Goroutine %d error = %v", i, errors[i])
		}
		if results[i] != expectedResult {
			t.Errorf("Goroutine %d result = %d, expected %d", i, results[i], expectedResult)
		}
	}
	
	// Function should have been called only once
	if atomic.LoadInt32(&callCount) != 1 {
		t.Errorf("Function called %d times, expected 1", atomic.LoadInt32(&callCount))
	}
}

func TestGroupDo_ConcurrentDifferentKeys(t *testing.T) {
	group := NewSingleGroup[string](100 * time.Millisecond)
	
	callCount := int32(0)
	fn := func() (string, error) {
		count := atomic.AddInt32(&callCount, 1)
		time.Sleep(10 * time.Millisecond) // Small delay to ensure concurrency
		return "result-" + string(rune('0'+count)), nil
	}
	
	var wg sync.WaitGroup
	results := make(map[string]string)
	errors := make(map[string]error)
	var mutex sync.Mutex
	
	keys := []string{"key1", "key2", "key3", "key4", "key5"}
	
	// Start concurrent calls with different keys
	for _, key := range keys {
		wg.Add(1)
		go func(k string) {
			defer wg.Done()
			result, err := group.Do(k, fn)
			mutex.Lock()
			results[k] = result
			errors[k] = err
			mutex.Unlock()
		}(key)
	}
	
	wg.Wait()
	
	// Check all calls succeeded
	for _, key := range keys {
		if errors[key] != nil {
			t.Errorf("Key %s error = %v", key, errors[key])
		}
		if results[key] == "" {
			t.Errorf("Key %s result is empty", key)
		}
	}
	
	// Function should have been called once per key
	if atomic.LoadInt32(&callCount) != int32(len(keys)) {
		t.Errorf("Function called %d times, expected %d", atomic.LoadInt32(&callCount), len(keys))
	}
}

func TestGroupGetOrAddSingle(t *testing.T) {
	group := NewSingleGroup[int](100 * time.Millisecond)
	
	// First call should create new Single
	single1 := group.getOrAddSingle("key1")
	if single1 == nil {
		t.Error("getOrAddSingle() returned nil")
	}
	
	// Second call with same key should return same Single
	single2 := group.getOrAddSingle("key1")
	if single2 != single1 {
		t.Error("getOrAddSingle() returned different Single for same key")
	}
	
	// Call with different key should return different Single
	single3 := group.getOrAddSingle("key2")
	if single3 == single1 {
		t.Error("getOrAddSingle() returned same Single for different key")
	}
}

// Test with different types
func TestSingleDo_StringType(t *testing.T) {
	single := NewSingle[string](100 * time.Millisecond)
	
	fn := func() (string, error) {
		return "test-string", nil
	}
	
	result, err := single.Do(fn)
	if err != nil {
		t.Errorf("Single.Do() error = %v", err)
	}
	if result != "test-string" {
		t.Errorf("Single.Do() result = %q, expected %q", result, "test-string")
	}
}

func TestSingleDo_StructType(t *testing.T) {
	type testStruct struct {
		ID   int
		Name string
	}
	
	single := NewSingle[testStruct](100 * time.Millisecond)
	
	expected := testStruct{ID: 42, Name: "test"}
	fn := func() (testStruct, error) {
		return expected, nil
	}
	
	result, err := single.Do(fn)
	if err != nil {
		t.Errorf("Single.Do() error = %v", err)
	}
	if result != expected {
		t.Errorf("Single.Do() result = %+v, expected %+v", result, expected)
	}
}

// Benchmark tests
func BenchmarkSingleDo(b *testing.B) {
	single := NewSingle[int](100 * time.Millisecond)
	fn := func() (int, error) {
		return 42, nil
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = single.Do(fn)
	}
}

func BenchmarkGroupDo(b *testing.B) {
	group := NewSingleGroup[int](100 * time.Millisecond)
	fn := func() (int, error) {
		return 42, nil
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = group.Do("benchmark-key", fn)
	}
}

func BenchmarkGroupDoDifferentKeys(b *testing.B) {
	group := NewSingleGroup[int](100 * time.Millisecond)
	fn := func() (int, error) {
		return 42, nil
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := "key-" + string(rune('0'+(i%10)))
		_, _ = group.Do(key, fn)
	}
}