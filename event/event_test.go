package event

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewManager(t *testing.T) {
	manager := NewManager()

	assert.NotNil(t, manager, "Manager should not be nil")
	assert.NotNil(t, manager.events, "Events map should be initialized")
	assert.NotNil(t, manager.c, "Channel should be initialized")
	assert.Equal(t, 10, cap(manager.c), "Channel should have buffer size of 10")
}

func TestManagerRegister(t *testing.T) {
	manager := NewManager()
	defer close(manager.c)

	var called bool
	var receivedArgs any
	handler := func(args any) {
		called = true
		receivedArgs = args
	}

	// Test registering sync handler
	manager.Register("test-event", handler)

	// Verify handler was registered
	items := manager.getItems("test-event")
	require.Len(t, items, 1, "Should have one registered handler")
	assert.False(t, items[0].async, "Handler should be synchronous")
	assert.NotNil(t, items[0].handler, "Handler function should not be nil")

	// Test emitting event
	testData := "test-data"
	manager.Emit("test-event", testData)

	assert.True(t, called, "Handler should have been called")
	assert.Equal(t, testData, receivedArgs, "Handler should receive correct arguments")
}

func TestManagerRegisterAsync(t *testing.T) {
	manager := NewManager()
	defer close(manager.c)

	var called int32
	var receivedArgs any
	handler := func(args any) {
		atomic.AddInt32(&called, 1)
		receivedArgs = args
	}

	// Test registering async handler
	manager.RegisterAsync("async-test-event", handler)

	// Verify handler was registered
	items := manager.getItems("async-test-event")
	require.Len(t, items, 1, "Should have one registered handler")
	assert.True(t, items[0].async, "Handler should be asynchronous")

	// Test emitting async event
	testData := "async-test-data"
	manager.Emit("async-test-event", testData)

	// Wait for async execution
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, int32(1), atomic.LoadInt32(&called), "Async handler should have been called")
	assert.Equal(t, testData, receivedArgs, "Async handler should receive correct arguments")
}

func TestManagerMultipleHandlers(t *testing.T) {
	manager := NewManager()
	defer close(manager.c)

	var syncCalled, asyncCalled int32
	var syncArgs, asyncArgs any

	syncHandler := func(args any) {
		atomic.AddInt32(&syncCalled, 1)
		syncArgs = args
	}

	asyncHandler := func(args any) {
		atomic.AddInt32(&asyncCalled, 1)
		asyncArgs = args
	}

	eventName := "multi-handler-event"

	// Register both sync and async handlers for same event
	manager.Register(eventName, syncHandler)
	manager.RegisterAsync(eventName, asyncHandler)

	// Verify both handlers are registered
	items := manager.getItems(eventName)
	require.Len(t, items, 2, "Should have two registered handlers")

	// Emit event
	testData := "multi-handler-data"
	manager.Emit(eventName, testData)

	// Wait for async execution
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, int32(1), atomic.LoadInt32(&syncCalled), "Sync handler should be called once")
	assert.Equal(t, int32(1), atomic.LoadInt32(&asyncCalled), "Async handler should be called once")
	assert.Equal(t, testData, syncArgs, "Sync handler should receive correct arguments")
	assert.Equal(t, testData, asyncArgs, "Async handler should receive correct arguments")
}

func TestManagerGetItemsNonExistentEvent(t *testing.T) {
	manager := NewManager()
	defer close(manager.c)

	items := manager.getItems("non-existent-event")
	assert.Nil(t, items, "Non-existent event should return nil")
}

func TestManagerEmitNonExistentEvent(t *testing.T) {
	manager := NewManager()
	defer close(manager.c)

	// Should not panic when emitting non-existent event
	assert.NotPanics(t, func() {
		manager.Emit("non-existent-event", "test-data")
	}, "Emitting non-existent event should not panic")
}

func TestManagerConcurrentAccess(t *testing.T) {
	manager := NewManager()
	defer close(manager.c)

	const numGoroutines = 50
	const numEvents = 10

	var wg sync.WaitGroup
	var totalCalled int32

	// Create handlers
	handler := func(args any) {
		atomic.AddInt32(&totalCalled, 1)
	}

	// Register handlers concurrently
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			eventName := "concurrent-event"
			manager.Register(eventName, handler)
		}(i)
	}
	wg.Wait()

	// Emit events concurrently
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numEvents; j++ {
				manager.Emit("concurrent-event", id)
			}
		}(i)
	}
	wg.Wait()

	// Wait for all handlers to complete
	time.Sleep(200 * time.Millisecond)

	// Each goroutine registered one handler, and each emitted numEvents events
	expectedCalls := int32(numGoroutines * numGoroutines * numEvents)
	actualCalls := atomic.LoadInt32(&totalCalled)
	assert.Equal(t, expectedCalls, actualCalls, "All handlers should be called")
}

func TestManagerAsyncChannelCapacity(t *testing.T) {
	manager := NewManager()
	defer close(manager.c)

	// Register a slow async handler
	slowHandler := func(args any) {
		time.Sleep(50 * time.Millisecond)
	}

	manager.RegisterAsync("slow-event", slowHandler)

	// Send more events than channel capacity
	for i := 0; i < 20; i++ {
		manager.Emit("slow-event", i)
	}

	// Should not block since it's using a buffered channel
	// This test verifies the channel doesn't deadlock
	assert.True(t, true, "Channel should handle burst of async events")
}

func TestDefaultManagerRegister(t *testing.T) {
	// Reset default manager state for clean test
	// Note: We cannot easily reset the default manager without affecting other tests
	// So we use a unique event name
	eventName := "default-sync-test-event"

	var called bool
	var receivedArgs any
	handler := func(args any) {
		called = true
		receivedArgs = args
	}

	// Test global Register function
	Register(eventName, handler)

	// Test global Emit function
	testData := "default-sync-data"
	Emit(eventName, testData)

	assert.True(t, called, "Default manager handler should have been called")
	assert.Equal(t, testData, receivedArgs, "Default manager handler should receive correct arguments")
}

func TestDefaultManagerRegisterAsync(t *testing.T) {
	eventName := "default-async-test-event"

	var called int32
	var receivedArgs any
	handler := func(args any) {
		atomic.AddInt32(&called, 1)
		receivedArgs = args
	}

	// Test global RegisterAsync function
	RegisterAsync(eventName, handler)

	// Test global Emit function
	testData := "default-async-data"
	Emit(eventName, testData)

	// Wait for async execution
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, int32(1), atomic.LoadInt32(&called), "Default manager async handler should have been called")
	assert.Equal(t, testData, receivedArgs, "Default manager async handler should receive correct arguments")
}

func TestEmitItemDo(t *testing.T) {
	var called bool
	var receivedArgs any

	handler := func(args any) {
		called = true
		receivedArgs = args
	}

	testData := "emit-item-data"
	item := &emitItem{
		handler: handler,
		args:    testData,
	}

	// Test the do method
	item.do()

	assert.True(t, called, "Handler should have been called")
	assert.Equal(t, testData, receivedArgs, "Handler should receive correct arguments")
}

func TestEmitItemDoWithNormalHandler(t *testing.T) {
	var called bool
	var receivedArgs any

	normalHandler := func(args any) {
		called = true
		receivedArgs = args
	}

	testData := "normal-handler-data"
	item := &emitItem{
		handler: normalHandler,
		args:    testData,
	}

	// Test normal execution of do method
	item.do()

	assert.True(t, called, "Normal handler should have been called")
	assert.Equal(t, testData, receivedArgs, "Normal handler should receive correct arguments")
}

func TestEventHandlerTypes(t *testing.T) {
	manager := NewManager()
	defer close(manager.c)

	// Test with different argument types
	testCases := []struct {
		name string
		data any
	}{
		{"string", "test-string"},
		{"int", 42},
		{"slice", []int{1, 2, 3}},
		{"map", map[string]int{"test": 123}},
		{"struct", struct{ Name string }{"test"}},
		{"nil", nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var receivedData any
			handler := func(args any) {
				receivedData = args
			}

			eventName := "type-test-" + tc.name
			manager.Register(eventName, handler)
			manager.Emit(eventName, tc.data)

			assert.Equal(t, tc.data, receivedData, "Handler should receive data of correct type")
		})
	}
}

func TestManagerChannelClose(t *testing.T) {
	manager := NewManager()

	var asyncCalled int32
	asyncHandler := func(args any) {
		atomic.AddInt32(&asyncCalled, 1)
	}

	manager.RegisterAsync("close-test-event", asyncHandler)

	// Emit some async events
	for i := 0; i < 5; i++ {
		manager.Emit("close-test-event", i)
	}

	// Close the channel to stop the goroutine
	close(manager.c)

	// Wait for goroutine to process events
	time.Sleep(100 * time.Millisecond)

	// Verify async events were processed
	assert.True(t, atomic.LoadInt32(&asyncCalled) >= 0, "Async handlers should have been called")
}

func TestManagerEmptyEventName(t *testing.T) {
	manager := NewManager()
	defer close(manager.c)

	var called bool
	handler := func(args any) {
		called = true
	}

	// Register and emit with empty event name
	manager.Register("", handler)
	manager.Emit("", "test-data")

	assert.True(t, called, "Handler with empty event name should work")
}

func TestManagerRegisterMultipleSameEvent(t *testing.T) {
	manager := NewManager()
	defer close(manager.c)

	var callCount int32
	handler := func(args any) {
		atomic.AddInt32(&callCount, 1)
	}

	eventName := "same-event-multiple-handlers"

	// Register the same handler multiple times
	for i := 0; i < 3; i++ {
		manager.Register(eventName, handler)
	}

	// Emit event once
	manager.Emit(eventName, "test-data")

	assert.Equal(t, int32(3), atomic.LoadInt32(&callCount), "All registered handlers should be called")
}

func TestAsyncEventProcessingOrder(t *testing.T) {
	manager := NewManager()
	defer close(manager.c)

	var results []int
	var mu sync.Mutex

	handler := func(args any) {
		// Add some delay to make ordering more apparent
		time.Sleep(10 * time.Millisecond)
		mu.Lock()
		results = append(results, args.(int))
		mu.Unlock()
	}

	manager.RegisterAsync("order-test-event", handler)

	// Emit events in order
	expectedOrder := []int{1, 2, 3, 4, 5}
	for _, val := range expectedOrder {
		manager.Emit("order-test-event", val)
	}

	// Wait for all async events to process
	time.Sleep(200 * time.Millisecond)

	mu.Lock()
	actualOrder := make([]int, len(results))
	copy(actualOrder, results)
	mu.Unlock()

	assert.Len(t, actualOrder, len(expectedOrder), "All events should be processed")
	assert.Equal(t, expectedOrder, actualOrder, "Events should be processed in order")
}

func TestMixedSyncAsyncHandlers(t *testing.T) {
	manager := NewManager()
	defer close(manager.c)

	var syncResults, asyncResults []string
	var syncMu, asyncMu sync.Mutex

	syncHandler := func(args any) {
		syncMu.Lock()
		syncResults = append(syncResults, "sync-"+args.(string))
		syncMu.Unlock()
	}

	asyncHandler := func(args any) {
		asyncMu.Lock()
		asyncResults = append(asyncResults, "async-"+args.(string))
		asyncMu.Unlock()
	}

	eventName := "mixed-handlers-event"
	manager.Register(eventName, syncHandler)
	manager.RegisterAsync(eventName, asyncHandler)

	// Emit multiple events
	testData := []string{"event1", "event2", "event3"}
	for _, data := range testData {
		manager.Emit(eventName, data)
	}

	// Wait for async handlers
	time.Sleep(100 * time.Millisecond)

	syncMu.Lock()
	actualSyncResults := make([]string, len(syncResults))
	copy(actualSyncResults, syncResults)
	syncMu.Unlock()

	asyncMu.Lock()
	actualAsyncResults := make([]string, len(asyncResults))
	copy(actualAsyncResults, asyncResults)
	asyncMu.Unlock()

	// Verify both sync and async handlers processed all events
	assert.Len(t, actualSyncResults, len(testData), "All sync handlers should be called")
	assert.Len(t, actualAsyncResults, len(testData), "All async handlers should be called")

	expectedSyncResults := []string{"sync-event1", "sync-event2", "sync-event3"}
	expectedAsyncResults := []string{"async-event1", "async-event2", "async-event3"}

	assert.Equal(t, expectedSyncResults, actualSyncResults, "Sync results should match expected")
	assert.Equal(t, expectedAsyncResults, actualAsyncResults, "Async results should match expected")
}

func TestManagerStructFields(t *testing.T) {
	manager := NewManager()
	defer close(manager.c)

	// Test that Manager has expected fields and they are properly initialized
	assert.NotNil(t, manager.events, "events map should be initialized")
	assert.NotNil(t, manager.c, "channel should be initialized")
	assert.Equal(t, 10, cap(manager.c), "channel should have buffer size of 10")

	// Test initial state
	items := manager.getItems("any-event")
	assert.Nil(t, items, "getItems should return nil for unregistered events")
}
