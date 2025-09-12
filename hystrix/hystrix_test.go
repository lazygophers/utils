package hystrix

import (
	"errors"
	"testing"
	"time"
)

func TestNewCircuitBreaker(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		TimeWindow: time.Second * 2,
	})
	if cb.State() != Closed {
		t.Errorf("Expected initial state to be Closed but got %s", cb.State())
	}
}

func TestBefore(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		TimeWindow: time.Second * 2,
	})
	if !cb.Before() {
		t.Error("Before() should return true when state is Closed")
	}

	cb.state.Store(stateOpenOpt)
	if cb.Before() {
		t.Error("Before() should return false when state is Open")
	}

	cb.state.Store(stateHalfOpenOpt)
	cb.probe = func() bool { return true }
	if !cb.Before() {
		t.Error("Before() should return true when Probe returns true")
	}

	cb.probe = func() bool { return false }
	if cb.Before() {
		t.Error("Before() should return false when Probe returns false")
	}
}

func TestState(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		TimeWindow: time.Second * 2,
	})
	if cb.State() != Closed {
		t.Errorf("Expected state to be Closed but got %s", cb.State())
	}

	cb.state.Store(stateOpenOpt)
	if cb.State() != Open {
		t.Errorf("Expected state to be Open but got %s", cb.State())
	}
}

func TestStat(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		TimeWindow: time.Second * 2,
	})
	if successes, failures := cb.Stat(); successes != 0 || failures != 0 {
		t.Errorf("Expected initial stats to be 0, got %d/%d", successes, failures)
	}

	cb.stats.successes.Store(10)
	cb.stats.failures.Store(5)
	if successes, failures := cb.Stat(); successes != 10 || failures != 5 {
		t.Errorf("Expected stats 10/5, got %d/%d", successes, failures)
	}
}

func TestTotal(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		TimeWindow: time.Second * 2,
	})
	if total := cb.Total(); total != 0 {
		t.Errorf("Expected total to be 0, got %d", total)
	}

	cb.stats.successes.Store(10)
	cb.stats.failures.Store(5)
	if total := cb.Total(); total != 15 {
		t.Errorf("Expected total to be 15, got %d", total)
	}
}

func TestAfter(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		TimeWindow: time.Second * 2,
	})

	cb.After(true)
	if successes, _ := cb.Stat(); successes != 1 {
		t.Error("Success count should be 1 after After(true)")
	}

	cb.After(false)
	if _, failures := cb.Stat(); failures != 1 {
		t.Error("Failure count should be 1 after After(false)")
	}
}

func TestCall(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		TimeWindow: time.Second * 2,
	})

	// 测试成功调用
	err := cb.Call(func() error {
		return nil
	})
	if err != nil {
		t.Errorf("Expected no error for successful call, got %v", err)
	}

	// 测试失败调用
	err = cb.Call(func() error {
		return errors.New("test error")
	})
	if err == nil {
		t.Error("Expected error for failed call")
	}

	// 测试熔断状态下的调用
	cb.state.Store(stateOpenOpt)
	// 防止状态被updateStateOptimized重置，通过增加失败计数
	cb.stats.failures.Store(10)
	cb.stats.changed.Store(0) // 防止状态更新
	err = cb.Call(func() error {
		return nil
	})
	if err == nil || err.Error() != "circuit breaker is open" {
		t.Error("Expected circuit breaker open error")
	}
}

func TestCircuitBreakerStateTransition(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		TimeWindow: time.Millisecond * 100,
		ReadyToTrip: func(successes, failures uint64) bool {
			return failures >= 3 // 3个失败触发熔断
		},
	})

	// 初始状态应该是 Closed
	if cb.State() != Closed {
		t.Errorf("Expected initial state Closed, got %s", cb.State())
	}

	// 连续失败触发熔断
	for i := 0; i < 5; i++ {
		cb.After(false)
	}
	
	// 触发状态更新
	cb.Before()
	
	if cb.State() != Open {
		t.Errorf("Expected state Open after failures, got %s", cb.State())
	}
}

func TestFastCircuitBreaker(t *testing.T) {
	cb := NewFastCircuitBreaker(3, time.Millisecond*100)

	// 初始状态应该允许请求
	if !cb.AllowRequest() {
		t.Error("Should allow requests initially")
	}

	// 记录失败
	cb.RecordResult(false)
	cb.RecordResult(false)
	cb.RecordResult(false)

	// 应该进入熔断状态
	if cb.AllowRequest() {
		t.Error("Should not allow requests after threshold failures")
	}

	// 等待时间窗口重置
	time.Sleep(time.Millisecond * 150)

	// 应该重新允许请求
	if !cb.AllowRequest() {
		t.Error("Should allow requests after time window reset")
	}
}

func TestBatchCircuitBreaker(t *testing.T) {
	config := CircuitBreakerConfig{
		TimeWindow: time.Second,
	}
	
	cb := NewBatchCircuitBreaker(config, 10, time.Millisecond*50)
	
	// 测试批量记录
	for i := 0; i < 5; i++ {
		cb.AfterBatch(true)
	}
	
	// 强制刷新
	cb.flush()
	
	// 检查统计
	successes, failures := cb.Stat()
	if successes != 5 || failures != 0 {
		t.Errorf("Expected 5 successes, 0 failures, got %d/%d", successes, failures)
	}
}

func TestCircuitBreakerOptimizations(t *testing.T) {
	cb := NewCircuitBreaker(CircuitBreakerConfig{
		TimeWindow: time.Millisecond * 100,
	})

	// 测试零分配状态查询
	if cb.GetState() != stateClosedOpt {
		t.Error("GetState should return stateClosedOpt initially")
	}

	if !cb.IsClosed() {
		t.Error("IsClosed should return true initially")
	}

	if cb.IsOpen() {
		t.Error("IsOpen should return false initially")
	}

	// 测试状态变更
	cb.state.Store(stateOpenOpt)
	if !cb.IsOpen() {
		t.Error("IsOpen should return true after setting to open")
	}

	if cb.IsClosed() {
		t.Error("IsClosed should return false after setting to open")
	}
}

func TestRingBufferCompatibility(t *testing.T) {
	rb := newRingBuffer(10)

	// 测试添加和长度
	if rb.len() != 0 {
		t.Error("Ring buffer should be empty initially")
	}

	// 添加请求结果
	rb.add(&requestResult{success: true, time: time.Now()})
	if rb.len() != 1 {
		t.Error("Ring buffer should have 1 element after add")
	}

	// 测试最后一个结果
	last := rb.last()
	if last == nil || !last.success {
		t.Error("Last result should exist and be successful")
	}

	// 测试重置
	rb.reset()
	if rb.len() != 0 {
		t.Error("Ring buffer should be empty after reset")
	}
}

func TestMemoryPoolOptimization(t *testing.T) {
	// 测试内存池
	result1 := getRequestResult()
	if result1 == nil {
		t.Error("getRequestResult should return non-nil")
	}

	result1.success = true
	result1.time = time.Now()

	putRequestResult(result1)

	result2 := getRequestResult()
	if result2 == nil {
		t.Error("getRequestResult should return non-nil after put")
	}

	// 应该被重置
	if result2.success != false || !result2.time.IsZero() {
		t.Error("Result should be reset when retrieved from pool")
	}
}