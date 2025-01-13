package singledo

import (
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

func (p *Group[T]) Do(key string, fn func() (T, error)) (v T, err error) {
	return p.getOrAddSingle(key).Do(fn)
}

func NewSingleGroup[T any](wait time.Duration) *Group[T] {
	return &Group[T]{
		wait:      wait,
		singleMap: make(map[string]*Single[T]),
	}
}
