package wait

import (
	"github.com/lazygophers/utils/routine"
	"github.com/lazygophers/utils/runtime"
	"sync"
)

type Worker struct {
	w *sync.WaitGroup
	c chan func()
}

func (p *Worker) Add(fn func()) {
	p.c <- fn
}

func (p *Worker) Wait() {
	close(p.c)
	p.w.Wait()
}

func NewWorker(max int) *Worker {
	c := make(chan func(), max)

	w := Wgp.Get().(*sync.WaitGroup)
	defer Wgp.Put(w)

	w.Add(max)
	for i := 0; i < max; i++ {
		routine.GoWithRecover(func() error {
			defer w.Done()

			for fn := range c {
				func() {
					defer runtime.CachePanic()
					fn()
				}()
			}

			return nil
		})
	}

	return &Worker{
		c: c,
		w: w,
	}
}
