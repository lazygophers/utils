package wait

import (
	"github.com/lazygophers/log"
	"sync"
)

var (
	poolLock sync.RWMutex

	poolMap = make(map[string]*Pool)
)

type (
	Pool struct {
		c chan struct{}
	}
)

func (p *Pool) Lock() {
	p.c <- struct{}{}
}

func (p *Pool) Unlock() {
	<-p.c
}

func (p *Pool) Depth() int {
	return len(p.c)
}

func getPool(key string) *Pool {
	poolLock.RLock()
	defer poolLock.RUnlock()

	return poolMap[key]
}

func newPool(key string, max int) {
	poolLock.RLock()
	p := poolMap[key]
	poolLock.RUnlock()

	if p != nil {
		return
	}

	poolLock.Lock()
	defer poolLock.Unlock()
	p = poolMap[key]

	if p != nil {
		return
	}

	p = &Pool{
		c: make(chan struct{}, max),
	}
	poolMap[key] = p
}

func Lock(key string) {
	getPool(key).Lock()
}

func Unlock(key string) {
	getPool(key).Unlock()
}

func Depth(key string) int {
	return getPool(key).Depth()
}

func Sync(key string, logic func() error) error {
	pool := getPool(key)

	log.Debugf("%s pool depth:%d", key, pool.Depth())
	pool.Lock()
	defer func() {
		pool.Unlock()
		log.Infof("%s pool depth:%d", key, pool.Depth())
	}()

	return logic()
}

func Ready(key string, max int) {
	newPool(key, max)
}
