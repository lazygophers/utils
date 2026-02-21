package wait

import (
	"context"
	"sync"

	"github.com/lazygophers/utils/routine"
)

// Sender 用于向异步处理器提交任务。
// 返回 false 表示 ctx 已结束，任务未被接收（调用方应尽快停止提交）。
type Sender[M any] func(M) bool

// AsyncCtx 是 Async 的 ctx 版本：允许通过 ctx 提前结束，避免因未关闭通道导致 goroutine 泄漏。
//
// 注意：
//   - push 必须使用 send 提交任务，并在 send 返回 false 时尽快返回。
//   - ctx 结束后，worker 会停止继续处理；未处理完的任务会被丢弃。
func AsyncCtx[M any](ctx context.Context, process int, push func(context.Context, Sender[M]), logic func(context.Context, M)) error {
	if ctx == nil {
		panic("wait: nil context")
	}
	if process <= 0 {
		return nil
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	c := make(chan M, process*2)

	w := Wgp.Get().(*sync.WaitGroup)
	defer func() {
		w.Wait()
		Wgp.Put(w)
	}()

	w.Add(process)
	for i := 0; i < process; i++ {
		routine.GoWithRecover(func() error {
			defer w.Done()

			for {
				select {
				case <-ctx.Done():
					return nil
				case x, ok := <-c:
					if !ok {
						return nil
					}
					logic(ctx, x)
				}
			}
		})
	}

	send := func(m M) bool {
		select {
		case <-ctx.Done():
			return false
		case c <- m:
			return true
		}
	}

	func() {
		defer func() {
			if r := recover(); r != nil {
				cancel()
				panic(r)
			}
		}()
		push(ctx, send)
	}()

	close(c)
	return ctx.Err()
}

// AsyncUniqueCtx 是 AsyncUnique 的 ctx 版本。
// 语义与 AsyncCtx 相同，但会基于 UniqueKey() 做并发去重，避免同一 key 并发执行。
func AsyncUniqueCtx[M UniqueTask](ctx context.Context, process int, push func(context.Context, Sender[M]), logic func(context.Context, M)) error {
	if ctx == nil {
		panic("wait: nil context")
	}
	if process <= 0 {
		return nil
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	c := make(chan M, process*2)
	var uniqueMap sync.Map

	w := Wgp.Get().(*sync.WaitGroup)
	defer func() {
		w.Wait()
		Wgp.Put(w)
	}()

	w.Add(process)
	for i := 0; i < process; i++ {
		routine.GoWithRecover(func() error {
			defer w.Done()

			for {
				select {
				case <-ctx.Done():
					return nil
				case x, ok := <-c:
					if !ok {
						return nil
					}

					key := x.UniqueKey()
					_, exist := uniqueMap.LoadOrStore(key, struct{}{})
					if exist {
						continue
					}
					func() {
						defer uniqueMap.Delete(key)
						logic(ctx, x)
					}()
				}
			}
		})
	}

	send := func(m M) bool {
		select {
		case <-ctx.Done():
			return false
		case c <- m:
			return true
		}
	}

	func() {
		defer func() {
			if r := recover(); r != nil {
				cancel()
				panic(r)
			}
		}()
		push(ctx, send)
	}()

	close(c)
	return ctx.Err()
}

// AsyncAlwaysWithChanCtx 是 AsyncAlwaysWithChan 的 ctx 版本。
// 即使调用者不关闭通道，也可以通过取消 ctx 让 worker 退出，避免泄漏。
//
// 返回的 done 会在所有 worker 退出后被关闭，便于调用方等待回收。
func AsyncAlwaysWithChanCtx[M any](ctx context.Context, process int, c <-chan M, logic func(context.Context, M)) <-chan struct{} {
	done := make(chan struct{})

	if ctx == nil {
		panic("wait: nil context")
	}
	if process <= 0 {
		close(done)
		return done
	}

	w := Wgp.Get().(*sync.WaitGroup)
	w.Add(process)

	for i := 0; i < process; i++ {
		routine.GoWithRecover(func() error {
			defer w.Done()

			for {
				select {
				case <-ctx.Done():
					return nil
				case x, ok := <-c:
					if !ok {
						return nil
					}
					logic(ctx, x)
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

	return done
}

// AsyncAlwaysUniqueWithChanCtx 是 AsyncAlwaysUniqueWithChan 的 ctx 版本。
// 返回 done，用于等待所有 worker 退出。
func AsyncAlwaysUniqueWithChanCtx[M UniqueTask](ctx context.Context, c <-chan M, process int, logic func(context.Context, M)) <-chan struct{} {
	done := make(chan struct{})

	if ctx == nil {
		panic("wait: nil context")
	}
	if process <= 0 {
		close(done)
		return done
	}

	var uniqueMap sync.Map
	w := Wgp.Get().(*sync.WaitGroup)
	w.Add(process)

	for i := 0; i < process; i++ {
		routine.GoWithRecover(func() error {
			defer w.Done()

			for {
				select {
				case <-ctx.Done():
					return nil
				case x, ok := <-c:
					if !ok {
						return nil
					}

					key := x.UniqueKey()
					_, exist := uniqueMap.LoadOrStore(key, struct{}{})
					if exist {
						continue
					}
					func() {
						defer uniqueMap.Delete(key)
						logic(ctx, x)
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

	return done
}
