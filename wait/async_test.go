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
	time.Sleep(500 * time.Millisecond)
	close(c)
	time.Sleep(500 * time.Millisecond) // 给协程完成时间

	// 验证结果
	expected := []int{2, 4, 6, 8, 10}
	assert.Len(t, results, len(expected))
	assert.ElementsMatch(t, expected, results)
}
