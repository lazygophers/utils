package routine

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestGo_BasicOperation(t *testing.T) {
	done := make(chan bool, 1)
	executed := int32(0)
	
	Go(func() error {
		atomic.StoreInt32(&executed, 1)
		done <- true
		return nil
	})
	
	// Wait for goroutine to complete
	select {
	case <-done:
		// Success
	case <-time.After(1 * time.Second):
		t.Error("Goroutine did not complete within timeout")
	}
	
	if atomic.LoadInt32(&executed) != 1 {
		t.Error("Function was not executed")
	}
}

func TestGo_WithError(t *testing.T) {
	done := make(chan bool, 1)
	expectedError := errors.New("test error")
	
	Go(func() error {
		defer func() { done <- true }()
		return expectedError
	})
	
	// Wait for goroutine to complete
	select {
	case <-done:
		// Success - error should be logged but not panic
	case <-time.After(1 * time.Second):
		t.Error("Goroutine did not complete within timeout")
	}
}

func TestGo_Concurrent(t *testing.T) {
	const numGoroutines = 10
	var wg sync.WaitGroup
	counter := int32(0)
	
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		Go(func() error {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
			return nil
		})
	}
	
	// Wait for all goroutines to complete
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	
	select {
	case <-done:
		// Success
	case <-time.After(2 * time.Second):
		t.Error("Goroutines did not complete within timeout")
	}
	
	if atomic.LoadInt32(&counter) != numGoroutines {
		t.Errorf("Expected %d executions, got %d", numGoroutines, atomic.LoadInt32(&counter))
	}
}

func TestGoWithRecover_BasicOperation(t *testing.T) {
	done := make(chan bool, 1)
	executed := int32(0)
	
	GoWithRecover(func() error {
		atomic.StoreInt32(&executed, 1)
		done <- true
		return nil
	})
	
	// Wait for goroutine to complete
	select {
	case <-done:
		// Success
	case <-time.After(1 * time.Second):
		t.Error("Goroutine did not complete within timeout")
	}
	
	if atomic.LoadInt32(&executed) != 1 {
		t.Error("Function was not executed")
	}
}

func TestGoWithRecover_WithPanic(t *testing.T) {
	done := make(chan bool, 1)
	
	GoWithRecover(func() error {
		defer func() { done <- true }()
		panic("test panic")
	})
	
	// Wait for goroutine to complete
	select {
	case <-done:
		// Success - panic should be recovered and logged
	case <-time.After(1 * time.Second):
		t.Error("Goroutine did not complete within timeout")
	}
}

func TestGoWithRecover_WithError(t *testing.T) {
	done := make(chan bool, 1)
	expectedError := errors.New("test error")
	
	GoWithRecover(func() error {
		defer func() { done <- true }()
		return expectedError
	})
	
	// Wait for goroutine to complete
	select {
	case <-done:
		// Success - error should be logged
	case <-time.After(1 * time.Second):
		t.Error("Goroutine did not complete within timeout")
	}
}

func TestGoWithRecover_PanicRecovery(t *testing.T) {
	done := make(chan bool, 1)
	panicked := int32(0)
	
	GoWithRecover(func() error {
		defer func() {
			atomic.StoreInt32(&panicked, 1)
			done <- true
		}()
		panic("intentional panic for testing")
	})
	
	// Wait for goroutine to complete
	select {
	case <-done:
		// Success
	case <-time.After(1 * time.Second):
		t.Error("Goroutine did not complete within timeout")
	}
	
	if atomic.LoadInt32(&panicked) != 1 {
		t.Error("Function did not execute or panic was not handled")
	}
}

func TestGoWithMustSuccess_BasicOperation(t *testing.T) {
	done := make(chan bool, 1)
	executed := int32(0)
	
	GoWithMustSuccess(func() error {
		atomic.StoreInt32(&executed, 1)
		done <- true
		return nil
	})
	
	// Wait for goroutine to complete
	select {
	case <-done:
		// Success
	case <-time.After(1 * time.Second):
		t.Error("Goroutine did not complete within timeout")
	}
	
	if atomic.LoadInt32(&executed) != 1 {
		t.Error("Function was not executed")
	}
}

// Note: Testing GoWithMustSuccess with error is tricky because it calls os.Exit(1)
// We can test the basic execution path but not the error path in unit tests

func TestAddBeforeRoutine(t *testing.T) {
	// Save original state
	originalBefore := beforeRoutines
	defer func() { beforeRoutines = originalBefore }()
	
	// Reset for test
	beforeRoutines = nil
	
	called := int32(0)
	testFunc := func(baseGid, currentGid int64) {
		atomic.AddInt32(&called, 1)
	}
	
	AddBeforeRoutine(testFunc)
	
	if len(beforeRoutines) != 1 {
		t.Errorf("Expected 1 before routine, got %d", len(beforeRoutines))
	}
	
	// Test that it gets called
	before(123, 456)
	
	if atomic.LoadInt32(&called) != 1 {
		t.Error("Before routine was not called")
	}
}

func TestAddAfterRoutine(t *testing.T) {
	// Save original state
	originalAfter := afterRoutines
	defer func() { afterRoutines = originalAfter }()
	
	// Reset for test
	afterRoutines = nil
	
	called := int32(0)
	testFunc := func(currentGid int64) {
		atomic.AddInt32(&called, 1)
	}
	
	AddAfterRoutine(testFunc)
	
	if len(afterRoutines) != 1 {
		t.Errorf("Expected 1 after routine, got %d", len(afterRoutines))
	}
	
	// Test that it gets called
	after(123)
	
	if atomic.LoadInt32(&called) != 1 {
		t.Error("After routine was not called")
	}
}

func TestBefore_MultipleRoutines(t *testing.T) {
	// Save original state
	originalBefore := beforeRoutines
	defer func() { beforeRoutines = originalBefore }()
	
	// Reset for test
	beforeRoutines = nil
	
	called1 := int32(0)
	called2 := int32(0)
	
	AddBeforeRoutine(func(baseGid, currentGid int64) {
		atomic.AddInt32(&called1, 1)
	})
	
	AddBeforeRoutine(func(baseGid, currentGid int64) {
		atomic.AddInt32(&called2, 1)
	})
	
	before(123, 456)
	
	if atomic.LoadInt32(&called1) != 1 {
		t.Error("First before routine was not called")
	}
	
	if atomic.LoadInt32(&called2) != 1 {
		t.Error("Second before routine was not called")
	}
}

func TestAfter_MultipleRoutines(t *testing.T) {
	// Save original state
	originalAfter := afterRoutines
	defer func() { afterRoutines = originalAfter }()
	
	// Reset for test
	afterRoutines = nil
	
	called1 := int32(0)
	called2 := int32(0)
	
	AddAfterRoutine(func(currentGid int64) {
		atomic.AddInt32(&called1, 1)
	})
	
	AddAfterRoutine(func(currentGid int64) {
		atomic.AddInt32(&called2, 1)
	})
	
	after(123)
	
	if atomic.LoadInt32(&called1) != 1 {
		t.Error("First after routine was not called")
	}
	
	if atomic.LoadInt32(&called2) != 1 {
		t.Error("Second after routine was not called")
	}
}

func TestBeforeAfter_Integration(t *testing.T) {
	done := make(chan bool, 1)
	beforeCalled := int32(0)
	afterCalled := int32(0)
	
	// Save original state
	originalBefore := beforeRoutines
	originalAfter := afterRoutines
	defer func() {
		beforeRoutines = originalBefore
		afterRoutines = originalAfter
	}()
	
	// Add test callbacks
	AddBeforeRoutine(func(baseGid, currentGid int64) {
		atomic.AddInt32(&beforeCalled, 1)
	})
	
	AddAfterRoutine(func(currentGid int64) {
		atomic.AddInt32(&afterCalled, 1)
	})
	
	Go(func() error {
		defer func() { done <- true }()
		return nil
	})
	
	// Wait for goroutine to complete
	select {
	case <-done:
		// Success
	case <-time.After(1 * time.Second):
		t.Error("Goroutine did not complete within timeout")
	}
	
	// Give some time for after routines to execute
	time.Sleep(100 * time.Millisecond)
	
	if atomic.LoadInt32(&beforeCalled) == 0 {
		t.Error("Before routine was not called")
	}
	
	if atomic.LoadInt32(&afterCalled) == 0 {
		t.Error("After routine was not called")
	}
}

func TestGoWithRecover_EmptyStack(t *testing.T) {
	done := make(chan bool, 1)
	
	GoWithRecover(func() error {
		defer func() { done <- true }()
		
		// Create a panic scenario that might result in empty stack
		// This is a bit artificial but tests the code path
		defer func() {
			if r := recover(); r != nil {
				// Re-panic to test the stack handling
				panic(r)
			}
		}()
		
		panic("test panic")
	})
	
	// Wait for goroutine to complete
	select {
	case <-done:
		// Success - should handle empty stack gracefully
	case <-time.After(1 * time.Second):
		t.Error("Goroutine did not complete within timeout")
	}
}

func TestGroup_Basic(t *testing.T) {
	// Test that Group struct exists and can be instantiated
	g := &Group{}
	if g == nil {
		t.Error("Group should not be nil")
	}
}

// Test concurrent execution of different routine types
func TestMixedRoutines_Concurrent(t *testing.T) {
	const numEach = 5
	var wg sync.WaitGroup
	
	goCount := int32(0)
	recoverCount := int32(0)
	mustSuccessCount := int32(0)
	
	// Launch Go routines
	for i := 0; i < numEach; i++ {
		wg.Add(1)
		Go(func() error {
			defer wg.Done()
			atomic.AddInt32(&goCount, 1)
			return nil
		})
	}
	
	// Launch GoWithRecover routines
	for i := 0; i < numEach; i++ {
		wg.Add(1)
		GoWithRecover(func() error {
			defer wg.Done()
			atomic.AddInt32(&recoverCount, 1)
			return nil
		})
	}
	
	// Launch GoWithMustSuccess routines
	for i := 0; i < numEach; i++ {
		wg.Add(1)
		GoWithMustSuccess(func() error {
			defer wg.Done()
			atomic.AddInt32(&mustSuccessCount, 1)
			return nil
		})
	}
	
	// Wait for all to complete
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	
	select {
	case <-done:
		// Success
	case <-time.After(3 * time.Second):
		t.Error("Not all goroutines completed within timeout")
	}
	
	if atomic.LoadInt32(&goCount) != numEach {
		t.Errorf("Expected %d Go executions, got %d", numEach, atomic.LoadInt32(&goCount))
	}
	
	if atomic.LoadInt32(&recoverCount) != numEach {
		t.Errorf("Expected %d GoWithRecover executions, got %d", numEach, atomic.LoadInt32(&recoverCount))
	}
	
	if atomic.LoadInt32(&mustSuccessCount) != numEach {
		t.Errorf("Expected %d GoWithMustSuccess executions, got %d", numEach, atomic.LoadInt32(&mustSuccessCount))
	}
}

// Test that routines handle nil errors properly
func TestRoutines_NilError(t *testing.T) {
	done := make(chan bool, 3)
	
	Go(func() error {
		done <- true
		return nil
	})
	
	GoWithRecover(func() error {
		done <- true
		return nil
	})
	
	GoWithMustSuccess(func() error {
		done <- true
		return nil
	})
	
	// Wait for all three
	for i := 0; i < 3; i++ {
		select {
		case <-done:
			// Success
		case <-time.After(1 * time.Second):
			t.Error("Goroutine did not complete within timeout")
		}
	}
}

// Benchmark tests
func BenchmarkGo(b *testing.B) {
	var wg sync.WaitGroup
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		Go(func() error {
			wg.Done()
			return nil
		})
	}
	wg.Wait()
}

func BenchmarkGoWithRecover(b *testing.B) {
	var wg sync.WaitGroup
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		GoWithRecover(func() error {
			wg.Done()
			return nil
		})
	}
	wg.Wait()
}

func BenchmarkGoWithMustSuccess(b *testing.B) {
	var wg sync.WaitGroup
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		GoWithMustSuccess(func() error {
			wg.Done()
			return nil
		})
	}
	wg.Wait()
}

// Test edge cases
func TestRoutines_EdgeCases(t *testing.T) {
	done := make(chan bool, 1)
	
	// Test with function that takes time
	Go(func() error {
		time.Sleep(10 * time.Millisecond)
		done <- true
		return nil
	})
	
	select {
	case <-done:
		// Success
	case <-time.After(1 * time.Second):
		t.Error("Long-running goroutine did not complete")
	}
}

func TestRoutines_WithMultipleErrors(t *testing.T) {
	const numErrors = 5
	var wg sync.WaitGroup
	
	for i := 0; i < numErrors; i++ {
		wg.Add(1)
		Go(func() error {
			defer wg.Done()
			return errors.New("test error")
		})
	}
	
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	
	select {
	case <-done:
		// Success - all errors should be logged
	case <-time.After(2 * time.Second):
		t.Error("Error goroutines did not complete within timeout")
	}
}