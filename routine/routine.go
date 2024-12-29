package routine

import (
	"fmt"
	"github.com/lazygophers/log"
	"github.com/petermattis/goid"
	"os"
	"runtime/debug"
	"strings"
)

type BeforeRoutine func(baseGid, currentGid int64)
type AfterRoutine func(currentGid int64)

var (
	beforeRoutines []BeforeRoutine
	afterRoutines  []AfterRoutine
)

func before(baseGid, currentGid int64) {
	for _, f := range beforeRoutines {
		f(baseGid, currentGid)
	}
}

func after(currentGid int64) {
	for _, f := range afterRoutines {
		f(currentGid)
	}
}

func AddBeforeRoutine(f BeforeRoutine) {
	beforeRoutines = append(beforeRoutines, f)
}

func AddAfterRoutine(f AfterRoutine) {
	afterRoutines = append(afterRoutines, f)
}

func init() {
	AddBeforeRoutine(func(baseGid, currentGid int64) {
		log.SetTraceWithGID(currentGid, fmt.Sprintf("%s.%s", log.GetTraceWithGID(baseGid), log.GenTraceId()))
	})

	AddAfterRoutine(func(currentGid int64) {
		log.DelTraceWithGID(currentGid)
	})
}

func Go(f func() (err error)) {
	baseGid := goid.Get()
	go func() {
		currentGid := goid.Get()
		before(baseGid, currentGid)
		defer func() {
			after(currentGid)
		}()

		err := f()
		if err != nil {
			log.Errorf("err:%v", err)
		}
	}()
}

func GoWithRecover(f func() (err error)) {
	baseGid := goid.Get()
	go func() {
		currentGid := goid.Get()
		before(baseGid, currentGid)
		defer func() {
			after(currentGid)
		}()

		defer func() {
			if err := recover(); err != nil {
				log.Errorf("err:%v", err)
				st := debug.Stack()
				if len(st) > 0 {
					log.Errorf("dump stack (%s):", err)
					lines := strings.Split(string(st), "\n")
					for _, line := range lines {
						log.Error("  ", line)
					}
				} else {
					log.Errorf("stack is empty (%s)", err)
				}
			}
		}()

		err := f()
		if err != nil {
			log.Errorf("err:%v", err)
		}
	}()
}

func GoWithMustSuccess(f func() (err error)) {
	baseGid := goid.Get()
	go func() {
		currentGid := goid.Get()
		before(baseGid, currentGid)
		defer func() {
			after(currentGid)
		}()

		err := f()
		if err != nil {
			log.Errorf("err:%v", err)
			os.Exit(1)
		}
	}()
}
