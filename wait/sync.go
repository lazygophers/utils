package wait

import (
	"sync"

	"github.com/lazygophers/log"
)

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

// Depth 返回当前已获取的信号量数量，即通道中当前的元素数量。
func (p *Pool) Depth() int {
	return len(p.c)
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
	// 先尝试读锁下检查
	poolLock.RLock()
	p := poolMap[key]
	poolLock.RUnlock()

	if p != nil {
		return
	}

	// 写锁下再次检查并创建
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
