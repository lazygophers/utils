package singledo

import (
	"sync"
	"time"
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

func (s *Single[T]) Do(fn func() (T, error)) (v T, err error) {
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

	callM := &call[T]{}
	callM.wg.Add(1)
	s.call = callM
	s.mux.Unlock()

	callM.val, callM.err = fn()
	callM.wg.Done()

	s.mux.Lock()
	if callM.err == nil {
		s.last = now
		s.result = callM.val
		s.call = nil
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
