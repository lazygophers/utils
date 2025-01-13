package wait_test

import (
	"github.com/lazygophers/utils/wait"
	"testing"
	"time"
)

func TestWork(t *testing.T) {
	worker := wait.NewWorker(10)

	for i := 0; i < 1000; i++ {
		worker.Add(func() {
			time.Sleep(time.Millisecond * time.Duration(i))
			t.Log(i)
		})
	}

	worker.Wait()
}
