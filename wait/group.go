package wait

import (
	"sync"

	"github.com/lazygophers/utils/routine"
	"github.com/lazygophers/utils/runtime"
)

// Worker 管理一组goroutine worker，通过任务队列和等待组协调任务执行
// 使用NewWorker创建，Add提交任务，Wait等待完成
type Worker struct {
	w *sync.WaitGroup // 用于等待所有任务完成的WaitGroup
	c chan func()     // 任务队列，接收待执行函数
}

// Add 向工作队列提交一个任务
// 任务是一个无参数的函数，将被Worker管理的goroutine执行
// 如果任务队列已满，该方法会阻塞，直到有可用空间
func (p *Worker) Add(fn func()) {
	p.c <- fn
}

// Wait 等待所有任务完成
// 注意：调用后不可再调用Add
// 内部会关闭通道并将WaitGroup放回对象池
func (p *Worker) Wait() {
	close(p.c)   // 关闭通道，停止接收新任务
	p.w.Wait()   // 等待所有正在执行的任务完成
	Wgp.Put(p.w) // 将WaitGroup放回对象池
}

// NewWorker 创建Worker实例
// max: 最大并发goroutine数量
func NewWorker(max int) *Worker {
	c := make(chan func(), max) // 创建带缓冲的任务通道

	// 从全局对象池获取WaitGroup（减少内存分配，优化性能）
	w := Wgp.Get().(*sync.WaitGroup)

	// 设置需要等待的goroutine数量
	w.Add(max)

	// 启动max个worker goroutine
	for i := 0; i < max; i++ {
		routine.GoWithRecover(func() error {
			defer w.Done() // goroutine结束时通知WaitGroup

			// 从任务通道不断获取任务执行
			for fn := range c {
				// 每个任务在独立闭包中执行，确保异常不会影响其他任务
				func() {
					// 捕获panic并记录日志（不中断程序）
					defer runtime.CachePanic()
					fn() // 执行实际任务函数
				}()
			}

			return nil
		})
	}

	return &Worker{
		c: c,
		w: w,
	}
}
