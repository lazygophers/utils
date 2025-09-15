package singledo

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

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
