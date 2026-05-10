package wait

import (
	"errors"
	"sync"
	"time"

	"github.com/lazygophers/log"
)

var ErrPoolNotReady = errors.New("wait: pool not ready (call Ready(key, max) first)")

// 全局读写锁，用于保护poolMap的并发访问
var (
	poolLock sync.RWMutex

	// poolMap 存储不同key对应的Pool实例
	poolMap = make(map[string]*Pool)
)

// Pool 是一个基于通道的信号量池，用于控制并发数量。
// 它使用一个有缓冲的通道来限制同时执行的操作数量。
type Pool struct {
	c chan struct{}
}

// Lock 获取一个信号量。如果池已满（即通道已满），则该方法会阻塞，直到有可用的信号量。
func (p *Pool) Lock() {
	p.c <- struct{}{}
}

// Unlock 释放一个信号量。如果池为空，则该方法会阻塞，直到有信号量被获取（通常不会发生，除非在未获取锁的情况下调用）。
func (p *Pool) Unlock() {
	<-p.c
}

// TryLock 尝试获取一个信号量，非阻塞。
// 返回 true 表示成功获取，false 表示池已满。
func (p *Pool) TryLock() bool {
	select {
	case p.c <- struct{}{}:
		return true
	default:
		return false
	}
}

// Available 返回当前可用的信号量数量。
func (p *Pool) Available() int {
	return len(p.c)
}

// Acquired 返回当前已获取的信号量数量。
func (p *Pool) Acquired() int {
	return cap(p.c) - len(p.c)
}

// Deprecated: Use Available instead.
func (p *Pool) Depth() int {
	return p.Available()
}

// getPool 根据key从poolMap中获取对应的Pool实例。
// 注意：调用此函数前必须持有poolLock的读锁。
func getPool(key string) *Pool {
	poolLock.RLock()
	defer poolLock.RUnlock()

	return poolMap[key]
}

// newPool 为指定的key创建一个新的Pool，并设置最大并发数max。
// 如果key对应的Pool已经存在，则不会重复创建。
func newPool(key string, max int) {
	if max <= 0 {
		max = 1
	}

	// 写锁下检查并创建（避免双重检查锁的间隙问题）
	poolLock.Lock()
	defer poolLock.Unlock()

	if poolMap[key] != nil {
		return
	}

	poolMap[key] = &Pool{
		c: make(chan struct{}, max),
	}
}

// Lock 获取指定key对应的Pool的锁。
// 如果key对应的Pool不存在，会panic。
func Lock(key string) {
	getPool(key).Lock()
}

// Unlock 释放指定key对应的Pool的锁。
// 如果key对应的Pool不存在，会panic。
func Unlock(key string) {
	getPool(key).Unlock()
}

// Depth 返回指定key对应的Pool的当前深度（已获取的信号量数量）。
// 如果key对应的Pool不存在，会panic。
func Depth(key string) int {
	return getPool(key).Depth()
}

// DepthOK 返回指定 key 对应 Pool 的深度及是否存在。
func DepthOK(key string) (depth int, ok bool) {
	pool := getPool(key)
	if pool == nil {
		return 0, false
	}
	return pool.Available(), true
}

// Resize 调整 Pool 的最大并发数。
// 如果新值小于当前已获取的信号量数量，则阻塞直到释放足够多的信号量。
func (p *Pool) Resize(newMax int) {
	if newMax <= 0 {
		newMax = 1
	}

	oldCap := cap(p.c)
	oldUsed := p.Acquired()

	if newMax == oldCap {
		return
	}

	// 创建新通道
	newC := make(chan struct{}, newMax)

	// 复制已使用的信号量到新通道
	for i := 0; i < oldUsed; i++ {
		newC <- struct{}{}
	}

	p.c = newC
}

// Sync 在指定key的Pool上同步执行逻辑函数logic。
// 它会自动获取锁，并在逻辑执行完成后释放锁，同时记录日志。
// 如果logic返回错误，该错误会被原样返回。
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

// Ready 初始化指定key的Pool，设置最大并发数max。
// 如果key对应的Pool已经存在，则不会重复创建。
func Ready(key string, max int) {
	newPool(key, max)
}

// TryLock 尝试获取指定key对应的Pool的锁，非阻塞。
// 返回 true 表示成功获取，false 表示池已满或不存在。
func TryLock(key string) bool {
	pool := getPool(key)
	if pool == nil {
		return false
	}
	return pool.TryLock()
}

// Resize 调整指定key的Pool的最大并发数。
func Resize(key string, newMax int) {
	pool := getPool(key)
	if pool == nil {
		return
	}
	pool.Resize(newMax)
}

// SyncTimeout 在指定key的Pool上同步执行逻辑函数，带超时控制。
func SyncTimeout(key string, timeout time.Duration, logic func() error) error {
	pool := getPool(key)
	if pool == nil {
		return ErrPoolNotReady
	}

	done := make(chan struct{}, 1)
	var err error

	log.Debugf("%s pool depth:%d", key, pool.Depth())
	pool.Lock()
	go func() {
		defer func() {
			pool.Unlock()
			log.Infof("%s pool depth:%d", key, pool.Depth())
			done <- struct{}{}
		}()
		err = logic()
	}()

	select {
	case <-done:
		return err
	case <-time.After(timeout):
		return errors.New("wait: timeout")
	}
}
