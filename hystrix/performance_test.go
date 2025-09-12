package hystrix

import (
	"errors"
	"sync"
	"testing"
	"time"
)

// BenchmarkOriginalVsOptimized 对比原版和优化版性能
func BenchmarkOriginalVsOptimized(b *testing.B) {
	b.Run("Original", func(b *testing.B) {
		cb := NewCircuitBreaker(CircuitBreakerConfig{
			TimeWindow: time.Second,
			ReadyToTrip: func(s, f uint64) bool {
				return false // 禁用熔断以专注于性能测试
			},
		})
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cb.Call(func() error { return nil })
		}
	})
	
	b.Run("Optimized", func(b *testing.B) {
		cb := NewOptimizedCircuitBreaker(CircuitBreakerConfig{
			TimeWindow: time.Second,
			ReadyToTrip: func(s, f uint64) bool {
				return false
			},
		})
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cb.Call(func() error { return nil })
		}
	})
	
	b.Run("Fast", func(b *testing.B) {
		cb := NewFastCircuitBreaker(100, time.Second)
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cb.CallFast(func() error { return nil })
		}
	})
}

// BenchmarkBeforeComparison Before方法性能对比
func BenchmarkBeforeComparison(b *testing.B) {
	b.Run("Original_Before", func(b *testing.B) {
		cb := NewCircuitBreaker(CircuitBreakerConfig{
			TimeWindow: time.Second,
			ReadyToTrip: func(s, f uint64) bool { return false },
		})
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cb.Before()
		}
	})
	
	b.Run("Optimized_Before", func(b *testing.B) {
		cb := NewOptimizedCircuitBreaker(CircuitBreakerConfig{
			TimeWindow: time.Second,
			ReadyToTrip: func(s, f uint64) bool { return false },
		})
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cb.Before()
		}
	})
	
	b.Run("Fast_AllowRequest", func(b *testing.B) {
		cb := NewFastCircuitBreaker(100, time.Second)
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cb.AllowRequest()
		}
	})
}

// BenchmarkAfterComparison After方法性能对比
func BenchmarkAfterComparison(b *testing.B) {
	b.Run("Original_After", func(b *testing.B) {
		cb := NewCircuitBreaker(CircuitBreakerConfig{
			TimeWindow: time.Second,
		})
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cb.After(i%2 == 0)
		}
	})
	
	b.Run("Optimized_After", func(b *testing.B) {
		cb := NewOptimizedCircuitBreaker(CircuitBreakerConfig{
			TimeWindow: time.Second,
		})
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cb.After(i%2 == 0)
		}
	})
	
	b.Run("Fast_RecordResult", func(b *testing.B) {
		cb := NewFastCircuitBreaker(100, time.Second)
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cb.RecordResult(i%2 == 0)
		}
	})
	
	b.Run("Batch_After", func(b *testing.B) {
		cb := NewBatchCircuitBreaker(CircuitBreakerConfig{
			TimeWindow: time.Second,
		}, 100, time.Millisecond*10)
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cb.AfterBatch(i%2 == 0)
		}
	})
}

// BenchmarkStateQueryComparison 状态查询性能对比
func BenchmarkStateQueryComparison(b *testing.B) {
	b.Run("Original_State", func(b *testing.B) {
		cb := NewCircuitBreaker(CircuitBreakerConfig{
			TimeWindow: time.Second,
		})
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cb.State()
		}
	})
	
	b.Run("Optimized_State", func(b *testing.B) {
		cb := NewOptimizedCircuitBreaker(CircuitBreakerConfig{
			TimeWindow: time.Second,
		})
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cb.State()
		}
	})
	
	b.Run("Optimized_GetState", func(b *testing.B) {
		cb := NewOptimizedCircuitBreaker(CircuitBreakerConfig{
			TimeWindow: time.Second,
		})
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cb.GetState()
		}
	})
	
	b.Run("Optimized_IsOpen", func(b *testing.B) {
		cb := NewOptimizedCircuitBreaker(CircuitBreakerConfig{
			TimeWindow: time.Second,
		})
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cb.IsOpen()
		}
	})
}

// BenchmarkConcurrentPerformance 并发性能对比
func BenchmarkConcurrentPerformance(b *testing.B) {
	b.Run("Original_Concurrent", func(b *testing.B) {
		cb := NewCircuitBreaker(CircuitBreakerConfig{
			TimeWindow: time.Second,
			ReadyToTrip: func(s, f uint64) bool { return false },
		})
		
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				cb.Call(func() error { return nil })
			}
		})
	})
	
	b.Run("Optimized_Concurrent", func(b *testing.B) {
		cb := NewOptimizedCircuitBreaker(CircuitBreakerConfig{
			TimeWindow: time.Second,
			ReadyToTrip: func(s, f uint64) bool { return false },
		})
		
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				cb.Call(func() error { return nil })
			}
		})
	})
	
	b.Run("Fast_Concurrent", func(b *testing.B) {
		cb := NewFastCircuitBreaker(1000, time.Second)
		
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				cb.CallFast(func() error { return nil })
			}
		})
	})
}

// BenchmarkMemoryUsage 内存使用对比
func BenchmarkMemoryUsage(b *testing.B) {
	b.Run("Original_Memory", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cb := NewCircuitBreaker(CircuitBreakerConfig{
				TimeWindow: time.Second,
			})
			// 执行一些操作来触发内存分配
			cb.After(true)
			cb.Before()
		}
	})
	
	b.Run("Optimized_Memory", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cb := NewOptimizedCircuitBreaker(CircuitBreakerConfig{
				TimeWindow: time.Second,
			})
			cb.After(true)
			cb.Before()
		}
	})
	
	b.Run("Fast_Memory", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cb := NewFastCircuitBreaker(100, time.Second)
			cb.RecordResult(true)
			cb.AllowRequest()
		}
	})
}

// BenchmarkHighLoad 高负载场景测试
func BenchmarkHighLoad(b *testing.B) {
	const workers = 100
	const opsPerWorker = 10000
	
	b.Run("Original_HighLoad", func(b *testing.B) {
		cb := NewCircuitBreaker(CircuitBreakerConfig{
			TimeWindow: time.Millisecond * 100,
			ReadyToTrip: func(s, f uint64) bool {
				return f > s && f+s > 1000
			},
		})
		
		b.ResetTimer()
		
		var wg sync.WaitGroup
		for i := 0; i < workers; i++ {
			wg.Add(1)
			go func(workerID int) {
				defer wg.Done()
				for j := 0; j < opsPerWorker; j++ {
					if workerID%10 == 0 {
						cb.Call(func() error { return errors.New("error") })
					} else {
						cb.Call(func() error { return nil })
					}
				}
			}(i)
		}
		wg.Wait()
	})
	
	b.Run("Optimized_HighLoad", func(b *testing.B) {
		cb := NewOptimizedCircuitBreaker(CircuitBreakerConfig{
			TimeWindow: time.Millisecond * 100,
			ReadyToTrip: func(s, f uint64) bool {
				return f > s && f+s > 1000
			},
		})
		
		b.ResetTimer()
		
		var wg sync.WaitGroup
		for i := 0; i < workers; i++ {
			wg.Add(1)
			go func(workerID int) {
				defer wg.Done()
				for j := 0; j < opsPerWorker; j++ {
					if workerID%10 == 0 {
						cb.Call(func() error { return errors.New("error") })
					} else {
						cb.Call(func() error { return nil })
					}
				}
			}(i)
		}
		wg.Wait()
	})
	
	b.Run("Fast_HighLoad", func(b *testing.B) {
		cb := NewFastCircuitBreaker(10000, time.Millisecond*100)
		
		b.ResetTimer()
		
		var wg sync.WaitGroup
		for i := 0; i < workers; i++ {
			wg.Add(1)
			go func(workerID int) {
				defer wg.Done()
				for j := 0; j < opsPerWorker; j++ {
					if workerID%10 == 0 {
						cb.CallFast(func() error { return errors.New("error") })
					} else {
						cb.CallFast(func() error { return nil })
					}
				}
			}(i)
		}
		wg.Wait()
	})
}

// BenchmarkStateTransitions 状态转换性能测试
func BenchmarkStateTransitions(b *testing.B) {
	b.Run("Original_StateTransitions", func(b *testing.B) {
		cb := NewCircuitBreaker(CircuitBreakerConfig{
			TimeWindow: time.Millisecond * 10,
			ReadyToTrip: func(s, f uint64) bool {
				return f > s
			},
		})
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if i%2 == 0 {
				cb.Call(func() error { return nil })
			} else {
				cb.Call(func() error { return errors.New("error") })
			}
		}
	})
	
	b.Run("Optimized_StateTransitions", func(b *testing.B) {
		cb := NewOptimizedCircuitBreaker(CircuitBreakerConfig{
			TimeWindow: time.Millisecond * 10,
			ReadyToTrip: func(s, f uint64) bool {
				return f > s
			},
		})
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if i%2 == 0 {
				cb.Call(func() error { return nil })
			} else {
				cb.Call(func() error { return errors.New("error") })
			}
		}
	})
}

// TestOptimizedFunctionality 测试优化版本的功能正确性
func TestOptimizedFunctionality(t *testing.T) {
	t.Run("OptimizedCircuitBreaker", func(t *testing.T) {
		cb := NewOptimizedCircuitBreaker(CircuitBreakerConfig{
			TimeWindow: time.Second,
			ReadyToTrip: func(s, f uint64) bool {
				return f >= 3 && s+f >= 3 // 至少3个样本且失败次数>=3
			},
		})
		
		// 测试初始状态
		if cb.State() != Closed {
			t.Errorf("Expected initial state to be Closed, got %s", cb.State())
		}
		
		// 测试成功调用
		err := cb.Call(func() error { return nil })
		if err != nil {
			t.Errorf("Expected successful call, got error: %v", err)
		}
		
		// 测试失败调用触发熔断
		for i := 0; i < 3; i++ {
			cb.Call(func() error { return errors.New("test error") })
		}
		
		// 触发状态更新检查 (调用Before会触发updateStateOptimized)
		cb.Before()
		
		if cb.State() != Open {
			t.Errorf("Expected state to be Open after failures, got %s", cb.State())
		}
		
		// 测试熔断状态下的调用被拒绝
		err = cb.Call(func() error { return nil })
		if err == nil {
			t.Error("Expected call to be rejected when circuit is open")
		}
	})
	
	t.Run("FastCircuitBreaker", func(t *testing.T) {
		cb := NewFastCircuitBreaker(2, time.Second)
		
		// 测试正常请求
		if !cb.AllowRequest() {
			t.Error("Should allow initial requests")
		}
		
		// 记录成功
		cb.RecordResult(true)
		
		// 记录失败直到触发熔断
		cb.RecordResult(false)
		cb.RecordResult(false)
		
		if cb.AllowRequest() {
			t.Error("Should not allow requests after reaching failure threshold")
		}
	})
}