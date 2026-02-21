package wait

import (
	"context"
	"sync"

	"github.com/lazygophers/utils/routine"
	"github.com/lazygophers/utils/runtime"
)

// WorkerCtx 是 Worker 的 ctx 版本：可通过 ctx/Stop 提前结束，避免忘记关闭导致 goroutine 泄漏。
type WorkerCtx struct {
	ctx       context.Context
	cancel    context.CancelFunc
	c         chan func()
	closeOnce sync.Once
	done      chan struct{}
}

// NewWorkerCtx 创建带 ctx 的 Worker。
// 当 ctx 结束或调用 Stop() 时，worker 会尽快退出；未执行的任务会被丢弃。
func NewWorkerCtx(ctx context.Context, max int) *WorkerCtx {
	if ctx == nil {
		panic("wait: nil context")
	}
	if max <= 0 {
		max = 1
	}

	workerCtx, cancel := context.WithCancel(ctx)
	c := make(chan func(), max)
	done := make(chan struct{})

	w := Wgp.Get().(*sync.WaitGroup)
	w.Add(max)

	for i := 0; i < max; i++ {
		routine.GoWithRecover(func() error {
			defer w.Done()

			for {
				select {
				case <-workerCtx.Done():
					return nil
				case fn, ok := <-c:
					if !ok {
						return nil
					}

					func() {
						defer runtime.CachePanic()
						fn()
					}()
				}
			}
		})
	}

	routine.GoWithRecover(func() error {
		w.Wait()
		Wgp.Put(w)
		close(done)
		return nil
	})

	return &WorkerCtx{
		ctx:    workerCtx,
		cancel: cancel,
		c:      c,
		done:   done,
	}
}

// Add 提交一个任务；当 ctx 已结束/Stop 后返回 false。
func (p *WorkerCtx) Add(fn func()) bool {
	if fn == nil {
		return false
	}

	select {
	case <-p.ctx.Done():
		return false
	default:
	}

	select {
	case <-p.ctx.Done():
		return false
	case p.c <- fn:
		return true
	}
}

// Stop 取消 ctx 并关闭队列，触发 worker 退出。
func (p *WorkerCtx) Stop() {
	p.cancel()
	p.closeOnce.Do(func() {
		close(p.c)
	})
}

// Wait 关闭队列并等待所有 worker 退出。
func (p *WorkerCtx) Wait() {
	p.closeOnce.Do(func() {
		close(p.c)
	})
	<-p.done
}

// Done 返回 worker 全部退出后的信号。
func (p *WorkerCtx) Done() <-chan struct{} {
	return p.done
}

// Err 返回 ctx 的错误状态。
func (p *WorkerCtx) Err() error {
	return p.ctx.Err()
}
