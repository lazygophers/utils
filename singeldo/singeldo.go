package singeldo

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/lazygophers/log"
)

type call[T any] struct {
	wg  sync.WaitGroup
	val T
	err error
}

type Single[T any] struct {
	mux    sync.Mutex
	last   time.Time
	wait   time.Duration
	call   *call[T]
	result T
}

// Deprecated: Use DoCtx instead. Will be removed in next major version.
func (s *Single[T]) Do(fn func() (T, error)) (v T, err error) {
	return s.DoCtx(context.Background(), fn)
}

// DoCtx executes fn with context support.
// If ctx is cancelled before fn starts, returns ctx.Err() immediately.
// If ctx is cancelled while fn is executing, waits for fn to complete.
// Only one fn executes for a given key; concurrent calls wait for the result.
func (s *Single[T]) DoCtx(ctx context.Context, fn func() (T, error)) (v T, err error) {
	s.mux.Lock()
	now := time.Now()
	if now.Before(s.last.Add(s.wait)) {
		s.mux.Unlock()
		return s.result, nil
	}

	if callM := s.call; callM != nil {
		s.mux.Unlock()
		callM.wg.Wait()
		return callM.val, callM.err
	}

	// Check context after cache/running check
	if err := ctx.Err(); err != nil {
		s.mux.Unlock()
		return v, err
	}

	callM := &call[T]{}
	callM.wg.Add(1)
	s.call = callM
	s.mux.Unlock()

	defer callM.wg.Done()

	func() {
		defer func() {
			if r := recover(); r != nil {
				log.Errorf("singeldo: panic in fn: %v", r)
				callM.err = fmt.Errorf("panic: %v", r)
			}
		}()
		callM.val, callM.err = fn()
	}()

	s.mux.Lock()
	s.call = nil
	if callM.err == nil {
		s.last = now
		s.result = callM.val
	}
	s.mux.Unlock()

	return callM.val, callM.err
}

func (s *Single[T]) Reset() {
	s.last = time.Time{}
}

func NewSingle[T any](wait time.Duration) *Single[T] {
	return &Single[T]{wait: wait}
}
