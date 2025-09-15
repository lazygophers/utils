package singledo

import (
	"sync"
	"testing"
	"time"
)

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
