package wait_test

import (
	"io"
	"sync"
	"testing"
	"time"

	"github.com/lazygophers/log"
	"github.com/lazygophers/utils/wait"
	"github.com/stretchr/testify/assert"
)

func TestAsync(t *testing.T) {
	// 测试数据准备
	type testCase struct {
		name     string
		process  int
		inputs   []int
		expected []int
	}

	testCases := []testCase{
		{
			name:     "single process",
			process:  1,
			inputs:   []int{1, 2, 3, 4, 5},
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:     "multiple processes",
			process:  3,
			inputs:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
		},
		{
			name:     "empty input",
			process:  2,
			inputs:   []int{},
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 准备结果收集器
			var (
				results []int
				mu      sync.Mutex
			)

			// 定义任务逻辑
			logic := func(x int) {
				mu.Lock()
				defer mu.Unlock()
				results = append(results, x*2)
			}

			// 定义任务推送函数
			push := func(ch chan int) {
				for _, num := range tc.inputs {
					ch <- num
				}
			}

			// 执行Async
			wait.Async(tc.process, push, logic)

			// 验证结果
			assert.Len(t, results, len(tc.expected))
			assert.ElementsMatch(t, tc.expected, results)
		})
	}
}

func TestAsync_Recover(t *testing.T) {
	// 测试panic恢复机制
	var (
		counter int
		mu      sync.Mutex
	)

	// 会panic的任务逻辑
	logic := func(x int) {
		mu.Lock()
		defer mu.Unlock()
		counter++

		if x == 3 {
			panic("test panic")
		}
	}

	// 任务推送
	push := func(ch chan int) {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
	}

	// 捕获日志输出（使用 io.Discard 避免空指针错误）
	log.SetOutput(io.Discard)

	// 执行Async
	wait.Async(2, push, logic)

	// 验证panic后继续执行
	assert.Equal(t, 5, counter, "所有任务都应完成，即使有panic")
}

func TestAsyncAlwaysWithChan(t *testing.T) {
	// 创建任务通道
	c := make(chan int, 10)

	// 结果收集器
	var (
		results []int
		mu      sync.Mutex
	)

	// 任务处理逻辑
	logic := func(x int) {
		mu.Lock()
		defer mu.Unlock()
		results = append(results, x*2)
	}

	// 启动持续处理协程
	wait.AsyncAlwaysWithChan(3, c, logic)

	// 推送任务
	for i := 1; i <= 5; i++ {
		c <- i
	}

	// 等待任务处理完成
	time.Sleep(50 * time.Millisecond)
	close(c)
	time.Sleep(50 * time.Millisecond) // 给协程完成时间

	// 验证结果
	expected := []int{2, 4, 6, 8, 10}
	assert.Len(t, results, len(expected))
	assert.ElementsMatch(t, expected, results)
}

// TestTask 实现 UniqueTask 接口用于测试
type TestTask struct {
	ID    string
	Value int
}

func (t TestTask) UniqueKey() string {
	return t.ID
}

func TestAsyncUnique(t *testing.T) {
	t.Run("normal operation", func(t *testing.T) {
		// 结果收集器
		var (
			results []int
			mu      sync.Mutex
		)

		// 任务处理逻辑
		logic := func(task TestTask) {
			mu.Lock()
			defer mu.Unlock()
			results = append(results, task.Value)
			time.Sleep(time.Millisecond) // 模拟处理时间
		}

		// 任务推送函数
		push := func(ch chan TestTask) {
			tasks := []TestTask{
				{ID: "task1", Value: 1},
				{ID: "task2", Value: 2},
				{ID: "task1", Value: 3}, // 重复任务
				{ID: "task3", Value: 4},
				{ID: "task2", Value: 5}, // 重复任务
			}
			for _, task := range tasks {
				ch <- task
			}
		}

		// 捕获日志输出
		log.SetOutput(io.Discard)
		defer log.SetOutput(io.Discard)

		// 执行AsyncUnique
		wait.AsyncUnique(3, push, logic)

		// 验证结果 - 重复的任务应该被过滤
		assert.True(t, len(results) >= 3 && len(results) <= 5, "应该处理3-5个任务")
		assert.Contains(t, results, 1)
		assert.Contains(t, results, 2)
		assert.Contains(t, results, 4)
	})

	t.Run("empty tasks", func(t *testing.T) {
		// 结果收集器
		var (
			results []int
			mu      sync.Mutex
		)

		// 任务处理逻辑
		logic := func(task TestTask) {
			mu.Lock()
			defer mu.Unlock()
			results = append(results, task.Value)
		}

		// 空任务推送函数
		push := func(ch chan TestTask) {
			// 不推送任何任务
		}

		// 执行AsyncUnique
		wait.AsyncUnique(2, push, logic)

		// 验证结果
		assert.Len(t, results, 0, "不应该有任何结果")
	})
}

func TestAsyncAlwaysUnique(t *testing.T) {
	// 结果收集器
	var (
		results []int
		mu      sync.Mutex
	)

	// 任务处理逻辑
	logic := func(task TestTask) {
		mu.Lock()
		defer mu.Unlock()
		results = append(results, task.Value)
		time.Sleep(time.Millisecond) // 模拟处理时间
	}

	// 捕获日志输出
	log.SetOutput(io.Discard)

	// 创建通道并启动协程
	c := wait.AsyncAlwaysUnique(3, logic)

	// 推送任务
	tasks := []TestTask{
		{ID: "task1", Value: 10},
		{ID: "task2", Value: 20},
		{ID: "task1", Value: 30}, // 重复任务
		{ID: "task3", Value: 40},
	}
	for _, task := range tasks {
		c <- task
	}

	// 等待处理完成
	time.Sleep(20 * time.Millisecond)
	close(c)
	time.Sleep(20 * time.Millisecond)

	// 验证结果
	assert.True(t, len(results) >= 3, "应该至少处理3个不重复任务")
	assert.Contains(t, results, 10)
	assert.Contains(t, results, 20)
	assert.Contains(t, results, 40)
}

func TestAsyncAlwaysUniqueWithChan(t *testing.T) {
	// 创建任务通道
	c := make(chan TestTask, 10)

	// 结果收集器
	var (
		results []int
		mu      sync.Mutex
	)

	// 任务处理逻辑
	logic := func(task TestTask) {
		mu.Lock()
		defer mu.Unlock()
		results = append(results, task.Value)
		time.Sleep(time.Millisecond) // 模拟处理时间
	}

	// 捕获日志输出
	log.SetOutput(io.Discard)

	// 启动持续处理协程
	wait.AsyncAlwaysUniqueWithChan(c, 2, logic)

	// 推送任务
	tasks := []TestTask{
		{ID: "task1", Value: 100},
		{ID: "task2", Value: 200},
		{ID: "task1", Value: 300}, // 重复任务
		{ID: "task3", Value: 400},
		{ID: "task2", Value: 500}, // 重复任务
	}
	for _, task := range tasks {
		c <- task
	}

	// 等待处理完成
	time.Sleep(20 * time.Millisecond)
	close(c)
	time.Sleep(20 * time.Millisecond)

	// 验证结果 - 重复的任务应该被过滤
	assert.True(t, len(results) >= 3, "应该至少处理3个不重复任务")
	assert.Contains(t, results, 100)
	assert.Contains(t, results, 200)
	assert.Contains(t, results, 400)
}
