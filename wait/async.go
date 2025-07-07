package wait

import (
	"sync"

	"github.com/lazygophers/log"
	"github.com/lazygophers/utils/routine"
)

var (
	Wgp = sync.Pool{
		New: func() interface{} {
			return &sync.WaitGroup{}
		},
	}
)

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

type UniqueTask interface {
	UniqueKey() string
}

func AsyncUnique[M UniqueTask](process int, push func(chan M), logic func(M)) {
	c := make(chan M, process*2)

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
