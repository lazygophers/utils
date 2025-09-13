package wait_test

import (
	"sync"
	"testing"
	"time"

	"github.com/lazygophers/utils/wait"
	"github.com/stretchr/testify/assert"
)

func TestNewWorker(t *testing.T) {
	t.Run("basic functionality", func(t *testing.T) {
		worker := wait.NewWorker(3)

		var (
			results []int
			mu      sync.Mutex
		)

		// 添加多个任务
		for i := 0; i < 10; i++ {
			i := i // 捕获循环变量
			worker.Add(func() {
				mu.Lock()
				defer mu.Unlock()
				results = append(results, i)
			})
		}

		worker.Wait()

		// 验证所有任务都执行了
		assert.Len(t, results, 10, "应该执行10个任务")

		// 验证包含所有预期值
		for i := 0; i < 10; i++ {
			assert.Contains(t, results, i, "结果应该包含%d", i)
		}
	})

	t.Run("single worker", func(t *testing.T) {
		worker := wait.NewWorker(1)

		var (
			counter int
			mu      sync.Mutex
		)

		// 串行执行任务
		for i := 0; i < 5; i++ {
			worker.Add(func() {
				mu.Lock()
				counter++
				mu.Unlock()
				time.Sleep(10 * time.Millisecond)
			})
		}

		worker.Wait()

		mu.Lock()
		finalCounter := counter
		mu.Unlock()

		assert.Equal(t, 5, finalCounter, "应该执行5个任务")
	})

	t.Run("many workers", func(t *testing.T) {
		worker := wait.NewWorker(100)

		var (
			counter int
			mu      sync.Mutex
		)

		// 大量并发任务
		for i := 0; i < 1000; i++ {
			worker.Add(func() {
				mu.Lock()
				counter++
				mu.Unlock()
			})
		}

		worker.Wait()

		mu.Lock()
		finalCounter := counter
		mu.Unlock()

		assert.Equal(t, 1000, finalCounter, "应该执行1000个任务")
	})

	t.Run("zero workers", func(t *testing.T) {
		worker := wait.NewWorker(0)

		executed := false
		worker.Add(func() {
			executed = true
		})

		// 立即等待，因为没有worker
		worker.Wait()

		// 零worker情况下任务不会被执行
		assert.False(t, executed, "零worker时任务不应该被执行")
	})
}

func TestWorkerPanicRecovery(t *testing.T) {
	worker := wait.NewWorker(2)

	var (
		results []string
		mu      sync.Mutex
	)

	// 添加正常任务
	worker.Add(func() {
		mu.Lock()
		defer mu.Unlock()
		results = append(results, "task1")
	})

	// 添加会panic的任务
	worker.Add(func() {
		mu.Lock()
		results = append(results, "before_panic")
		mu.Unlock()
		panic("test panic")
	})

	// 添加另一个正常任务
	worker.Add(func() {
		time.Sleep(50 * time.Millisecond) // 确保panic任务先执行
		mu.Lock()
		defer mu.Unlock()
		results = append(results, "task3")
	})

	worker.Wait()

	mu.Lock()
	finalResults := results
	mu.Unlock()

	// 验证panic不会阻止其他任务执行
	assert.Contains(t, finalResults, "task1", "正常任务1应该执行")
	assert.Contains(t, finalResults, "before_panic", "panic任务的前半部分应该执行")
	assert.Contains(t, finalResults, "task3", "正常任务3应该执行")
}

func TestWorkerConcurrentAdd(t *testing.T) {
	worker := wait.NewWorker(5)

	var (
		counter int
		mu      sync.Mutex
	)

	// 并发添加任务
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				worker.Add(func() {
					mu.Lock()
					counter++
					mu.Unlock()
				})
			}
		}()
	}

	// 等待所有添加操作完成
	wg.Wait()

	// 等待所有任务执行完成
	worker.Wait()

	mu.Lock()
	finalCounter := counter
	mu.Unlock()

	assert.Equal(t, 100, finalCounter, "应该执行100个任务")
}

func TestWorkerTaskExecution(t *testing.T) {
	t.Run("task execution order with buffered channel", func(t *testing.T) {
		// 测试缓冲通道的任务执行
		worker := wait.NewWorker(2)

		var (
			execution []int
			mu        sync.Mutex
		)

		// 快速添加多个任务
		for i := 0; i < 6; i++ {
			i := i
			worker.Add(func() {
				mu.Lock()
				execution = append(execution, i)
				mu.Unlock()
				time.Sleep(10 * time.Millisecond)
			})
		}

		worker.Wait()

		mu.Lock()
		finalExecution := execution
		mu.Unlock()

		assert.Len(t, finalExecution, 6, "应该执行6个任务")

		// 验证所有任务都执行了（不关心顺序）
		for i := 0; i < 6; i++ {
			assert.Contains(t, finalExecution, i, "应该包含任务%d", i)
		}
	})

	t.Run("long running tasks", func(t *testing.T) {
		worker := wait.NewWorker(3)

		var (
			results []time.Time
			mu      sync.Mutex
		)

		start := time.Now()

		// 添加一些耗时任务
		for i := 0; i < 5; i++ {
			worker.Add(func() {
				time.Sleep(100 * time.Millisecond)
				mu.Lock()
				results = append(results, time.Now())
				mu.Unlock()
			})
		}

		worker.Wait()

		duration := time.Since(start)

		mu.Lock()
		finalResults := results
		mu.Unlock()

		assert.Len(t, finalResults, 5, "应该执行5个任务")

		// 由于有3个worker，5个任务应该在大约200ms内完成（两批）
		assert.Less(t, duration, 300*time.Millisecond, "任务应该并发执行")
		assert.Greater(t, duration, 150*time.Millisecond, "任务应该确实执行了")
	})
}

// 基准测试
func BenchmarkWorker(b *testing.B) {
	b.Run("worker_creation", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			worker := wait.NewWorker(10)
			worker.Wait()
		}
	})

	b.Run("task_execution", func(b *testing.B) {
		worker := wait.NewWorker(10)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			worker.Add(func() {
				// 模拟轻量级任务
				time.Sleep(time.Microsecond)
			})
		}

		worker.Wait()
	})

	b.Run("concurrent_task_addition", func(b *testing.B) {
		worker := wait.NewWorker(10)

		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				worker.Add(func() {
					// 轻量级任务
				})
			}
		})

		worker.Wait()
	})
}
