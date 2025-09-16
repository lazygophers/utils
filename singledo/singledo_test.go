package singledo

import (
	"errors"
	"runtime"
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

// Race condition tests - originally from forced_race_test.go

// TestGetOrAddSingleForcedRaceCondition forces the double-checked locking scenario
func TestGetOrAddSingleForcedRaceCondition(t *testing.T) {
	t.Run("force_double_check_locking_scenario", func(t *testing.T) {
		// Use many attempts with different timing to maximize chance of race
		for attempt := 0; attempt < 1000; attempt++ {
			group := NewSingleGroup[int](100 * time.Millisecond)

			var wg sync.WaitGroup
			const numGoroutines = 50
			results := make([]*Single[int], numGoroutines)

			// Use a channel that will be closed to start all goroutines
			start := make(chan struct{})
			ready := make(chan struct{}, numGoroutines)

			// Create goroutines that will all try to access the same key
			for i := 0; i < numGoroutines; i++ {
				wg.Add(1)
				go func(index int) {
					defer wg.Done()

					// Signal ready and wait for start
					ready <- struct{}{}
					<-start

					// Force a small delay to create contention
					runtime.Gosched()

					// Access the same key - this should trigger the race
					results[index] = group.getOrAddSingle("contested-key")
				}(i)
			}

			// Wait for all goroutines to be ready
			for i := 0; i < numGoroutines; i++ {
				<-ready
			}

			// Start all goroutines simultaneously
			close(start)
			wg.Wait()

			// Verify all results are the same instance
			for i := 1; i < numGoroutines; i++ {
				if results[i] != results[0] {
					t.Errorf("Attempt %d: Different instances returned", attempt)
					return
				}
			}

			if results[0] == nil {
				t.Errorf("Attempt %d: nil result", attempt)
				return
			}
		}

		t.Log("Completed forced race condition test")
	})

	t.Run("force_timing_based_race", func(t *testing.T) {
		// Try different approaches to force the race condition
		for timing := 0; timing < 10; timing++ {
			for attempt := 0; attempt < 100; attempt++ {
				group := NewSingleGroup[string](10 * time.Millisecond)

				var mu sync.Mutex
				var firstInside, secondInside bool
				results := make([]*Single[string], 2)

				var wg sync.WaitGroup
				wg.Add(2)

				// First goroutine
				go func() {
					defer wg.Done()

					// Add varying delays to create different timing
					time.Sleep(time.Duration(timing) * time.Microsecond)

					mu.Lock()
					firstInside = true
					mu.Unlock()

					results[0] = group.getOrAddSingle("timing-key")
				}()

				// Second goroutine
				go func() {
					defer wg.Done()

					// Slightly different delay
					time.Sleep(time.Duration(timing+1) * time.Microsecond)

					mu.Lock()
					secondInside = true
					mu.Unlock()

					results[1] = group.getOrAddSingle("timing-key")
				}()

				wg.Wait()

				mu.Lock()
				bothInside := firstInside && secondInside
				mu.Unlock()

				if !bothInside {
					continue // Try again
				}

				// Verify both got the same instance
				if results[0] != results[1] {
					t.Errorf("Timing %d, Attempt %d: Different instances", timing, attempt)
				}

				if results[0] == nil || results[1] == nil {
					t.Errorf("Timing %d, Attempt %d: nil results", timing, attempt)
				}
			}
		}

		t.Log("Completed timing-based race condition test")
	})
}

// Double-checked locking tests - originally from double_check_locking_test.go

// TestGroupGetOrAddSingleDoubleCheckLocking tests the double-checked locking pattern
// to achieve 100% coverage of the getOrAddSingle function
func TestGroupGetOrAddSingleDoubleCheckLocking(t *testing.T) {
	t.Run("double_check_locking_race_condition", func(t *testing.T) {
		// Create many groups to increase the chance of hitting the race condition
		for attempt := 0; attempt < 100; attempt++ {
			group := NewSingleGroup[int](100 * time.Millisecond)

			const numGoroutines = 100
			var wg sync.WaitGroup
			results := make([]*Single[int], numGoroutines)

			// Create a channel to synchronize the start
			start := make(chan struct{})

			// Start multiple goroutines trying to get the same key simultaneously
			for i := 0; i < numGoroutines; i++ {
				wg.Add(1)
				go func(index int) {
					defer wg.Done()

					// Wait for start signal
					<-start

					// This should trigger the double-checked locking pattern
					results[index] = group.getOrAddSingle("race-key")
				}(i)
			}

			// Release all goroutines at once
			close(start)

			// Wait for all goroutines to complete
			wg.Wait()

			// All results should point to the same Single instance
			for i := 1; i < numGoroutines; i++ {
				if results[i] != results[0] {
					t.Errorf("Attempt %d, Goroutine %d got different Single instance", attempt, i)
				}
			}

			// Verify the Single is not nil
			if results[0] == nil {
				t.Errorf("Attempt %d: getOrAddSingle returned nil", attempt)
			}
		}

		t.Log("Successfully tested double-checked locking with intensive race conditions")
	})

	t.Run("double_check_locking_sequential_access", func(t *testing.T) {
		group := NewSingleGroup[string](50 * time.Millisecond)

		// First call - should create new Single (covers creation path)
		single1 := group.getOrAddSingle("sequential-key")
		if single1 == nil {
			t.Error("First getOrAddSingle returned nil")
		}

		// Second call - should return existing Single (covers return existing path)
		single2 := group.getOrAddSingle("sequential-key")
		if single2 != single1 {
			t.Error("Second getOrAddSingle returned different Single")
		}

		// Third call - should still return same Single
		single3 := group.getOrAddSingle("sequential-key")
		if single3 != single1 {
			t.Error("Third getOrAddSingle returned different Single")
		}
	})

	t.Run("double_check_locking_mixed_access", func(t *testing.T) {
		group := NewSingleGroup[int](75 * time.Millisecond)

		const numKeys = 3
		const goroutinesPerKey = 5
		var wg sync.WaitGroup
		var startWg sync.WaitGroup

		results := make(map[string][]*Single[int])
		resultMutex := sync.Mutex{}

		// Initialize results map
		for i := 0; i < numKeys; i++ {
			key := string(rune('A' + i))
			results[key] = make([]*Single[int], goroutinesPerKey)
		}

		startWg.Add(1)

		// Launch goroutines for each key
		for keyIndex := 0; keyIndex < numKeys; keyIndex++ {
			key := string(rune('A' + keyIndex))

			for goroutineIndex := 0; goroutineIndex < goroutinesPerKey; goroutineIndex++ {
				wg.Add(1)

				go func(k string, gIndex int) {
					defer wg.Done()

					// Wait for all goroutines to be ready
					startWg.Wait()

					// Get or create Single for this key
					single := group.getOrAddSingle(k)

					// Store result safely
					resultMutex.Lock()
					results[k][gIndex] = single
					resultMutex.Unlock()
				}(key, goroutineIndex)
			}
		}

		// Start all goroutines simultaneously
		startWg.Done()
		wg.Wait()

		// Verify results for each key
		for keyIndex := 0; keyIndex < numKeys; keyIndex++ {
			key := string(rune('A' + keyIndex))
			keyResults := results[key]

			// All Singles for the same key should be identical
			for i := 1; i < goroutinesPerKey; i++ {
				if keyResults[i] != keyResults[0] {
					t.Errorf("Key %s: goroutine %d got different Single instance", key, i)
				}
			}

			if keyResults[0] == nil {
				t.Errorf("Key %s: getOrAddSingle returned nil", key)
			}
		}

		// Verify that different keys have different Singles
		keyA := results["A"][0]
		keyB := results["B"][0]
		keyC := results["C"][0]

		if keyA == keyB || keyA == keyC || keyB == keyC {
			t.Error("Different keys should have different Single instances")
		}

		t.Logf("Successfully tested mixed access with %d keys and %d goroutines per key", numKeys, goroutinesPerKey)
	})
}

// TestGroupGetOrAddSingleCoverageEdgeCases covers additional edge cases
func TestGroupGetOrAddSingleCoverageEdgeCases(t *testing.T) {
	t.Run("high_contention_scenario", func(t *testing.T) {
		group := NewSingleGroup[int](25 * time.Millisecond)

		const numGoroutines = 50
		const numIterations = 10

		var wg sync.WaitGroup
		successes := make([]int, numGoroutines)

		for goroutine := 0; goroutine < numGoroutines; goroutine++ {
			wg.Add(1)

			go func(gID int) {
				defer wg.Done()

				for iter := 0; iter < numIterations; iter++ {
					key := "high-contention"
					single := group.getOrAddSingle(key)

					if single != nil {
						successes[gID]++
					}

					// Small delay to create more contention
					time.Sleep(time.Microsecond)
				}
			}(goroutine)
		}

		wg.Wait()

		// Verify all operations succeeded
		totalSuccess := 0
		for i := 0; i < numGoroutines; i++ {
			totalSuccess += successes[i]
		}

		expectedTotal := numGoroutines * numIterations
		if totalSuccess != expectedTotal {
			t.Errorf("Expected %d successful operations, got %d", expectedTotal, totalSuccess)
		}

		// Final verification that the key exists and is consistent
		finalSingle1 := group.getOrAddSingle("high-contention")
		finalSingle2 := group.getOrAddSingle("high-contention")

		if finalSingle1 != finalSingle2 {
			t.Error("Final verification failed: Singles are not identical")
		}
	})

	t.Run("alternating_keys_scenario", func(t *testing.T) {
		group := NewSingleGroup[string](30 * time.Millisecond)

		const numRounds = 20
		keys := []string{"alt1", "alt2"}

		var wg sync.WaitGroup

		for round := 0; round < numRounds; round++ {
			for _, key := range keys {
				wg.Add(1)

				go func(k string, r int) {
					defer wg.Done()

					single := group.getOrAddSingle(k)
					if single == nil {
						t.Errorf("Round %d, Key %s: getOrAddSingle returned nil", r, k)
					}
				}(key, round)
			}
		}

		wg.Wait()

		// Verify final consistency
		single1a := group.getOrAddSingle("alt1")
		single1b := group.getOrAddSingle("alt1")
		single2a := group.getOrAddSingle("alt2")
		single2b := group.getOrAddSingle("alt2")

		if single1a != single1b {
			t.Error("Key alt1: inconsistent Singles")
		}
		if single2a != single2b {
			t.Error("Key alt2: inconsistent Singles")
		}
		if single1a == single2a {
			t.Error("Different keys should have different Singles")
		}
	})
}

// Additional race condition tests - originally from race_condition_test.go

// TestGetOrAddSingleRaceCondition tests the race condition in getOrAddSingle
// This test tries to trigger the second nil check after acquiring the write lock
func TestGetOrAddSingleRaceCondition(t *testing.T) {
	group := NewSingleGroup[int](100 * time.Millisecond)

	// Use a large number of goroutines to increase the chance of hitting the race condition
	const numGoroutines = 100
	var wg sync.WaitGroup

	// Create a barrier to make all goroutines start at roughly the same time
	startBarrier := make(chan struct{})

	results := make([]*Single[int], numGoroutines)

	// Launch multiple goroutines that all try to get the same key
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			// Wait for all goroutines to be ready
			<-startBarrier

			// All goroutines try to get the same key at the same time
			single := group.getOrAddSingle("race-test-key")
			results[index] = single
		}(i)
	}

	// Release all goroutines at once to maximize race condition chances
	close(startBarrier)

	// Wait for all goroutines to complete
	wg.Wait()

	// Verify that all goroutines got the same Single instance
	firstSingle := results[0]
	if firstSingle == nil {
		t.Error("First Single is nil")
		return
	}

	for i := 1; i < numGoroutines; i++ {
		if results[i] != firstSingle {
			t.Errorf("Goroutine %d got different Single instance", i)
		}
	}

	// Verify that the Single is properly stored in the map
	storedSingle := group.getOrAddSingle("race-test-key")
	if storedSingle != firstSingle {
		t.Error("Stored Single is different from returned Single")
	}
}

// TestConcurrentGetOrAddSingleDifferentKeys tests concurrent access with different keys
func TestConcurrentGetOrAddSingleDifferentKeys(t *testing.T) {
	group := NewSingleGroup[string](50 * time.Millisecond)

	const numKeys = 50
	var wg sync.WaitGroup

	results := make(map[string]*Single[string])
	var mutex sync.Mutex

	// Create goroutines for different keys
	for i := 0; i < numKeys; i++ {
		wg.Add(1)
		go func(keyIndex int) {
			defer wg.Done()

			key := string(rune('A' + keyIndex))
			single := group.getOrAddSingle(key)

			mutex.Lock()
			results[key] = single
			mutex.Unlock()
		}(i)
	}

	wg.Wait()

	// Verify each key has its own Single instance
	if len(results) != numKeys {
		t.Errorf("Expected %d results, got %d", numKeys, len(results))
	}

	// Verify all Singles are different
	singles := make([]*Single[string], 0, numKeys)
	for _, single := range results {
		singles = append(singles, single)
	}

	for i := 0; i < len(singles); i++ {
		for j := i + 1; j < len(singles); j++ {
			if singles[i] == singles[j] {
				t.Error("Found duplicate Single instances for different keys")
			}
		}
	}
}

// TestHighContention tests the locking mechanism under high contention
func TestHighContention(t *testing.T) {
	group := NewSingleGroup[int](10 * time.Millisecond)

	const numGoroutines = 200
	const numKeys = 10

	var wg sync.WaitGroup

	var mutex sync.Mutex
	results := make(map[string][]*Single[int])

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()

			// Each goroutine accesses multiple keys
			keyIndex := goroutineID % numKeys
			key := string(rune('0' + keyIndex))

			single := group.getOrAddSingle(key)

			mutex.Lock()
			results[key] = append(results[key], single)
			mutex.Unlock()
		}(i)
	}

	wg.Wait()

	// Verify that all accesses to the same key returned the same Single
	for key, singles := range results {
		if len(singles) == 0 {
			continue
		}

		firstSingle := singles[0]
		for i, single := range singles {
			if single != firstSingle {
				t.Errorf("Key %s: access %d returned different Single", key, i)
			}
		}
	}
}

// TestDoubleCheckedLocking specifically targets the double-checked locking pattern
func TestDoubleCheckedLocking(t *testing.T) {
	group := NewSingleGroup[int](100 * time.Millisecond)

	// This test attempts to trigger the scenario where:
	// 1. Multiple goroutines pass the first nil check
	// 2. They all wait for the write lock
	// 3. The first one creates the Single
	// 4. The subsequent ones should hit the second nil check and return early

	const numGoroutines = 50
	var wg sync.WaitGroup

	// Synchronized start
	startChan := make(chan struct{})
	results := make([]*Single[int], numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			<-startChan
			single := group.getOrAddSingle("double-check-key")
			results[index] = single
		}(i)
	}

	// Start all goroutines simultaneously
	close(startChan)
	wg.Wait()

	// All results should be the same Single instance
	if results[0] == nil {
		t.Error("First result is nil")
		return
	}

	for i := 1; i < numGoroutines; i++ {
		if results[i] != results[0] {
			t.Errorf("Result %d is different from first result", i)
		}
		if results[i] == nil {
			t.Errorf("Result %d is nil", i)
		}
	}

	// Verify the Single is correctly stored
	finalSingle := group.getOrAddSingle("double-check-key")
	if finalSingle != results[0] {
		t.Error("Final retrieval returned different Single")
	}
}