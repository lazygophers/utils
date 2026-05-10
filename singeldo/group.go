package singeldo

import (
	"context"
	"sync"
	"time"
)

type Group[T any] struct {
	wait time.Duration

	mux       sync.RWMutex
	singleMap map[string]*Single[T]
}

func (p *Group[T]) getOrAddSingle(key string) *Single[T] {
	p.mux.RLock()
	single := p.singleMap[key]
	p.mux.RUnlock()

	if single != nil {
		return single
	}

	p.mux.Lock()
	defer p.mux.Unlock()

	single = p.singleMap[key]

	if single != nil {
		return single
	}

	single = NewSingle[T](p.wait)
	p.singleMap[key] = single

	return single
}

// Deprecated: Use DoCtx instead. Will be removed in next major version.
func (p *Group[T]) Do(key string, fn func() (T, error)) (v T, err error) {
	return p.DoCtx(context.Background(), key, fn)
}

// DoCtx executes fn with context support for the given key.
// If ctx is cancelled before fn starts, returns ctx.Err() immediately.
// If ctx is cancelled while fn is executing, waits for fn to complete.
// Only one fn executes per key; concurrent calls with the same key wait for the result.
func (p *Group[T]) DoCtx(ctx context.Context, key string, fn func() (T, error)) (v T, err error) {
	return p.getOrAddSingle(key).DoCtx(ctx, fn)
}

func NewSingleGroup[T any](wait time.Duration) *Group[T] {
	return &Group[T]{
		wait:      wait,
		singleMap: make(map[string]*Single[T]),
	}
}
