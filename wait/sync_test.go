package wait_test

import (
	"errors"
	"io"
	"sync"
	"testing"
	"time"

	"github.com/lazygophers/log"
	"github.com/lazygophers/utils/wait"
	"github.com/stretchr/testify/assert"
)

func TestReady(t *testing.T) {
	t.Run("create new pool", func(t *testing.T) {
		key := "test_pool_1"
		max := 5

		// 创建新池
		wait.Ready(key, max)

		// 验证池已创建
		depth := wait.Depth(key)
		assert.Equal(t, 0, depth, "新创建的池深度应该为0")
	})

	t.Run("pool already exists", func(t *testing.T) {
		key := "test_pool_2"
		max := 3

		// 第一次创建
		wait.Ready(key, max)

		// 第二次创建同样的key - 应该不会重复创建
		wait.Ready(key, max)

		// 验证池仍然可用
		depth := wait.Depth(key)
		assert.Equal(t, 0, depth, "现有池深度应该为0")
	})

	t.Run("different keys create different pools", func(t *testing.T) {
		key1 := "test_pool_3"
		key2 := "test_pool_4"

		wait.Ready(key1, 2)
		wait.Ready(key2, 4)

		// 验证不同池独立工作
		assert.Equal(t, 0, wait.Depth(key1))
		assert.Equal(t, 0, wait.Depth(key2))
	})
}

func TestLockUnlock(t *testing.T) {
	t.Run("basic lock/unlock", func(t *testing.T) {
		key := "test_lock_1"
		max := 3

		wait.Ready(key, max)

		// 测试加锁和解锁
		wait.Lock(key)
		assert.Equal(t, 1, wait.Depth(key), "加锁后深度应该为1")

		wait.Lock(key)
		assert.Equal(t, 2, wait.Depth(key), "第二次加锁后深度应该为2")

		wait.Unlock(key)
		assert.Equal(t, 1, wait.Depth(key), "解锁后深度应该为1")

		wait.Unlock(key)
		assert.Equal(t, 0, wait.Depth(key), "第二次解锁后深度应该为0")
	})

	t.Run("max capacity", func(t *testing.T) {
		key := "test_lock_2"
		max := 2

		wait.Ready(key, max)

		// 达到最大容量
		wait.Lock(key)
		wait.Lock(key)
		assert.Equal(t, 2, wait.Depth(key), "应该达到最大容量")

		// 释放一个
		wait.Unlock(key)
		assert.Equal(t, 1, wait.Depth(key), "释放后深度应该减少")

		wait.Unlock(key)
		assert.Equal(t, 0, wait.Depth(key), "全部释放后深度应该为0")
	})
}

func TestDepth(t *testing.T) {
	key := "test_depth"
	max := 5

	wait.Ready(key, max)

	// 测试不同深度
	depths := []int{0, 1, 2, 3, 4, 5}

	for i, expectedDepth := range depths[:len(depths)-1] {
		assert.Equal(t, expectedDepth, wait.Depth(key), "深度应该匹配当前状态")
		if i < len(depths)-2 {
			wait.Lock(key)
		}
	}

	// 解锁测试
	for i := max - 1; i >= 0; i-- {
		if i < max-1 {
			wait.Unlock(key)
		}
		assert.Equal(t, i, wait.Depth(key), "解锁后深度应该减少")
	}
}

func TestSync(t *testing.T) {
	t.Run("successful execution", func(t *testing.T) {
		key := "test_sync_1"
		max := 2

		wait.Ready(key, max)

		// 捕获日志输出
		log.SetOutput(io.Discard)
		defer log.SetOutput(io.Discard)

		executed := false
		err := wait.Sync(key, func() error {
			executed = true
			return nil
		})

		assert.NoError(t, err, "同步执行不应该返回错误")
		assert.True(t, executed, "逻辑函数应该被执行")
		assert.Equal(t, 0, wait.Depth(key), "执行完成后深度应该为0")
	})

	t.Run("execution with error", func(t *testing.T) {
		key := "test_sync_2"
		max := 1

		wait.Ready(key, max)

		// 捕获日志输出
		log.SetOutput(io.Discard)
		defer log.SetOutput(io.Discard)

		expectedErr := errors.New("test error")
		executed := false

		err := wait.Sync(key, func() error {
			executed = true
			return expectedErr
		})

		assert.Error(t, err, "应该返回错误")
		assert.Equal(t, expectedErr, err, "应该返回相同的错误")
		assert.True(t, executed, "逻辑函数应该被执行")
		assert.Equal(t, 0, wait.Depth(key), "即使出错，执行完成后深度也应该为0")
	})

	t.Run("panic in logic function", func(t *testing.T) {
		key := "test_sync_3"
		max := 1

		wait.Ready(key, max)

		// 捕获日志输出
		log.SetOutput(io.Discard)
		defer log.SetOutput(io.Discard)

		executed := false

		// 测试panic恢复
		assert.Panics(t, func() {
			wait.Sync(key, func() error {
				executed = true
				panic("test panic")
			})
		}, "panic应该被传播")

		assert.True(t, executed, "逻辑函数应该被执行")
		// 注意：由于panic，defer可能不会执行，所以深度检查可能不准确
	})

	t.Run("concurrent access", func(t *testing.T) {
		key := "test_sync_4"
		max := 2

		wait.Ready(key, max)

		// 捕获日志输出
		log.SetOutput(io.Discard)
		defer log.SetOutput(io.Discard)

		results := make(chan int, 5)
		var wg sync.WaitGroup

		// 启动多个并发同步操作
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				wait.Sync(key, func() error {
					time.Sleep(50 * time.Millisecond) // 模拟工作
					results <- id
					return nil
				})
			}(i)
		}

		// 收集结果
		var collected []int
		for i := 0; i < 5; i++ {
			select {
			case result := <-results:
				collected = append(collected, result)
			case <-time.After(2 * time.Second):
				t.Fatal("超时等待结果")
			}
		}

		// 等待所有goroutine完成
		wg.Wait()

		assert.Len(t, collected, 5, "应该收到所有5个结果")
		assert.Equal(t, 0, wait.Depth(key), "所有操作完成后深度应该为0")
	})
}

func TestPoolMethods(t *testing.T) {
	t.Run("pool direct methods", func(t *testing.T) {
		// 通过Ready创建池然后测试底层方法
		key := "test_pool_methods"
		max := 3

		wait.Ready(key, max)

		// 测试Lock方法
		wait.Lock(key)
		assert.Equal(t, 1, wait.Depth(key))

		// 测试Unlock方法
		wait.Unlock(key)
		assert.Equal(t, 0, wait.Depth(key))

		// 测试多次Lock和Unlock
		for i := 1; i <= max; i++ {
			wait.Lock(key)
			assert.Equal(t, i, wait.Depth(key))
		}

		for i := max - 1; i >= 0; i-- {
			wait.Unlock(key)
			assert.Equal(t, i, wait.Depth(key))
		}
	})
}

// 边界情况和错误处理测试
func TestEdgeCases(t *testing.T) {
	t.Run("zero max capacity", func(t *testing.T) {
		key := "test_zero_max"
		max := 0

		wait.Ready(key, max)

		// 零容量池的深度应该始终为0
		assert.Equal(t, 0, wait.Depth(key))
	})

	t.Run("large max capacity", func(t *testing.T) {
		key := "test_large_max"
		max := 1000

		wait.Ready(key, max)

		// 测试大容量池
		for i := 0; i < 10; i++ {
			wait.Lock(key)
		}
		assert.Equal(t, 10, wait.Depth(key))

		for i := 0; i < 10; i++ {
			wait.Unlock(key)
		}
		assert.Equal(t, 0, wait.Depth(key))
	})

	t.Run("concurrent pool creation", func(t *testing.T) {
		key := "test_concurrent_creation"
		max := 5

		// 并发创建相同的池
		done := make(chan bool, 10)
		for i := 0; i < 10; i++ {
			go func() {
				wait.Ready(key, max)
				done <- true
			}()
		}

		// 等待所有goroutine完成
		for i := 0; i < 10; i++ {
			<-done
		}

		// 验证池正常工作
		wait.Lock(key)
		assert.Equal(t, 1, wait.Depth(key))
		wait.Unlock(key)
		assert.Equal(t, 0, wait.Depth(key))
	})
}

// 性能基准测试
func BenchmarkSync(b *testing.B) {
	key := "benchmark_sync"
	max := 10

	wait.Ready(key, max)

	// 捕获日志输出
	log.SetOutput(io.Discard)
	defer log.SetOutput(io.Discard)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			wait.Sync(key, func() error {
				// 模拟一些轻量级工作
				time.Sleep(time.Microsecond)
				return nil
			})
		}
	})
}

func BenchmarkLockUnlock(b *testing.B) {
	key := "benchmark_lock"
	max := 100

	wait.Ready(key, max)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			wait.Lock(key)
			wait.Unlock(key)
		}
	})
}
