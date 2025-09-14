package singledo

import (
	"sync"
	"testing"
	"time"
)

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