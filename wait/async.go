package wait

import (
	"sync"

	"github.com/lazygophers/log"
	"github.com/lazygophers/utils/routine"
)

// Wgp 是 sync.WaitGroup 的对象池，用于复用 WaitGroup 对象
var (
	Wgp = sync.Pool{
		New: func() interface{} {
			return &sync.WaitGroup{}
		},
	}
)

// Async 使用协程池处理任务
// 参数:
//
//	process: 并发处理的任务数量（协程数量）
//	push: 任务推送函数，接收一个通道参数用于发送任务
//	logic: 任务处理逻辑函数
//
// 协程池工作流程:
//  1. 创建缓冲通道用于任务传递
//  2. 从对象池获取 WaitGroup
//  3. 启动指定数量的工作协程
//  4. 调用 push 函数推送任务到通道
//  5. 关闭通道并等待所有任务完成
//  6. 将 WaitGroup 放回对象池
func Async[M any](process int, push func(chan M), logic func(M)) {
	if process <= 0 {
		return
	}
	
	c := make(chan M, process*2) // 增加缓冲区大小，避免阻塞

	w := Wgp.Get().(*sync.WaitGroup)
	defer func() {
		w.Wait()
		Wgp.Put(w)
	}()

	w.Add(process)
	for i := 0; i < process; i++ {
		routine.GoWithRecover(func() error {
			defer w.Done()

			for x := range c {
				logic(x)
			}

			return nil
		})
	}

	push(c)
	close(c)
}

// AsyncAlwaysWithChan 使用指定数量的协程持续处理通道中的任务
// 参数:
//
//	process: 并发处理的任务数量（协程数量）
//	c: 任务通道
//	logic: 任务处理逻辑函数
//
// 注意: 调用者需要负责关闭通道以停止协程
func AsyncAlwaysWithChan[M any](process int, c chan M, logic func(M)) {
	for i := 0; i < process; i++ {
		routine.GoWithRecover(func() error {
			var x M
			for x = range c {
				logic(x)
			}

			return nil
		})
	}
}

// UniqueTask 定义需要唯一性校验的任务接口
// UniqueKey() 返回任务的唯一标识键
type UniqueTask interface {
	UniqueKey() string
}

// AsyncUnique 使用带唯一性校验的协程池处理任务
// 参数:
//
//	process: 并发处理的任务数量（协程数量）
//	push: 任务推送函数
//	logic: 任务处理逻辑函数
//
// 唯一性校验逻辑:
//
//	使用 sync.Map 存储任务唯一键，确保相同任务不会并发执行
func AsyncUnique[M UniqueTask](process int, push func(chan M), logic func(M)) {
	c := make(chan M, process*2)

	// uniqueMap 用于存储任务唯一键，防止重复执行
	var uniqueMap sync.Map

	w := Wgp.Get().(*sync.WaitGroup)
	defer Wgp.Put(w)

	w.Add(process)
	for i := 0; i < process; i++ {
		routine.GoWithRecover(func() error {
			defer w.Done()

			var x M
			for x = range c {
				_, exist := uniqueMap.LoadOrStore(x.UniqueKey(), struct{}{})
				if exist {
					log.Warnf("task exist:%s", x.UniqueKey())
					continue
				}
				logic(x)
				uniqueMap.Delete(x.UniqueKey())
			}

			return nil
		})
	}

	push(c)
	close(c)

	w.Wait()
}

// AsyncAlwaysUnique 创建任务通道并启动带唯一性校验的协程
// 参数:
//
//	process: 并发处理的任务数量（协程数量）
//	logic: 任务处理逻辑函数
//
// 返回值: 任务通道
func AsyncAlwaysUnique[M UniqueTask](process int, logic func(M)) chan M {
	c := make(chan M, 20)
	AsyncAlwaysUniqueWithChan(c, process, logic)
	return c
}

// AsyncAlwaysUniqueWithChan 使用带唯一性校验的协程处理通道中的任务
// 参数:
//
//	c: 任务通道
//	process: 并发处理的任务数量（协程数量）
//	logic: 任务处理逻辑函数
//
// 唯一性校验逻辑:
//
//	使用 sync.Map 存储任务唯一键，确保相同任务不会并发执行
func AsyncAlwaysUniqueWithChan[M UniqueTask](c chan M, process int, logic func(M)) {
	// uniqueMap 用于存储任务唯一键，防止重复执行
	var uniqueMap sync.Map
	for i := 0; i < process; i++ {
		routine.GoWithRecover(func() error {
			var x M
			for x = range c {
				_, exist := uniqueMap.LoadOrStore(x.UniqueKey(), struct{}{})
				if exist {
					log.Warnf("task exist:%s", x.UniqueKey())
					continue
				}
				logic(x)
				uniqueMap.Delete(x.UniqueKey())
			}

			return nil
		})
	}
}
