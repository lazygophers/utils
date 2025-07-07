// Package wait 提供并发控制和异步任务管理工具
//
// 使用WaitGroup池和channel实现任务分发与回收
// 主要功能：
//   - 并发任务执行(Async)
//   - 持续任务处理(AsyncAlwaysWithChan)
//   - 唯一性任务保证(AsyncUnique/AsyncAlwaysUniqueWithChan)
package wait

import (
	"sync"

	"github.com/lazygophers/log"
	"github.com/lazygophers/utils/routine"
)

// Wgp WaitGroup对象池，用于管理goroutine同步
//
// 实现细节：
//   - 使用sync.Pool减少内存分配
//   - WaitGroup生命周期由Get/Put管理
//   - 每个WaitGroup可支持最大并发量由调用方指定
var (
	Wgp = sync.Pool{
		New: func() interface{} {
			return &sync.WaitGroup{}
		},
	}
)

// Async 并发处理数据流
//
// 参数:
//   - process: 并发处理数量
//   - push: 数据生成函数，向channel推送数据
//   - logic: 数据处理逻辑函数
//
// 执行流程:
// 1. 创建带缓冲的channel
// 2. 从对象池获取WaitGroup
// 3. 启动指定数量的goroutine消费channel
// 4. 等待所有goroutine完成
//
// 该函数通过创建指定数量的goroutine来并发处理数据，使用WaitGroup确保所有任务完成后再返回
// 注意：routine.GoWithRecover会自动捕获panic并转换为error处理
func Async[M any](process int, push func(chan M), logic func(M)) {
	c := make(chan M, process)

	w := Wgp.Get().(*sync.WaitGroup)
	defer Wgp.Put(w)

	w.Add(process)
	for i := 0; i < process; i++ {
		routine.GoWithRecover(func() error {
			defer w.Done()

			var x M
			for x = range c {
				logic(x)
			}

			return nil
		})
	}

	push(c)
	close(c)

	w.Wait()
}

// AsyncAlwaysWithChan 持续处理指定channel中的所有数据项
//
// @param process 并发处理数量
// @param c 数据源channel
// @param logic 每个数据项的处理逻辑
//
// 函数会创建process个goroutine，持续从channel读取数据直到关闭
// 适用于需要完全处理channel中所有数据的场景
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

// UniqueTask 唯一性任务接口
//
// 必须实现UniqueKey方法返回任务唯一标识
// 用于防止重复处理相同任务
type UniqueTask interface {
	// UniqueKey 返回任务的唯一键
	// 用于去重判断
	UniqueKey() string
}

// AsyncUnique 并发处理具有唯一性的任务
//
// @param process 并发数
// @param push 数据推送函数
// @param logic 数据处理逻辑
//
// 通过UniqueKey()方法确保每个任务的唯一性，重复任务会被跳过
// 使用sync.Map存储已处理的任务键，防止并发写入冲突
func AsyncUnique[M UniqueTask](process int, push func(chan M), logic func(M)) {
	c := make(chan M, process*2)
	// uniqueMap 用于存储已处理的任务唯一键
	// Key: 任务的UniqueKey()返回值
	// Value: 空结构体占位符，仅用于判断是否存在
	// uniqueMap 用于存储已处理的任务唯一键
	// Key: 任务的UniqueKey()返回值
	// Value: 空结构体占位符，仅用于判断是否存在
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

func AsyncAlwaysUnique[M UniqueTask](process int, logic func(M)) chan M {
	c := make(chan M, 20)
	AsyncAlwaysUniqueWithChan(c, process, logic)
	return c
}

// AsyncAlwaysUniqueWithChan 持续处理具有唯一性的数据流
//
// @param c 数据源channel
// @param process 并发处理数量
// @param logic 处理逻辑函数
//
// 与AsyncAlwaysWithChan类似，但增加了任务唯一性校验
// 通过UniqueKey()方法识别重复项，重复数据会触发警告日志并跳过
func AsyncAlwaysUniqueWithChan[M UniqueTask](c chan M, process int, logic func(M)) {
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
