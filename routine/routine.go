package routine

import (
	"fmt"
	"github.com/lazygophers/log"
	"os"
	"runtime/debug"
	"strings"
)

func Go(f func() (err error)) {
	traceId := log.GetTrace()
	go func() {
		log.SetTrace(fmt.Sprintf("%s.%s", traceId, log.GenTraceId()))
		defer log.DelTrace()

		err := f()
		if err != nil {
			log.Errorf("err:%v", err)
		}
	}()
}

func GoWithRecover(f func() (err error)) {
	traceId := log.GetTrace()
	go func() {
		log.SetTrace(fmt.Sprintf("%s.%s", traceId, log.GenTraceId()))
		defer log.DelTrace()

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
	traceId := log.GetTrace()
	go func() {
		log.SetTrace(fmt.Sprintf("%s.%s", traceId, log.GenTraceId()))
		defer log.DelTrace()

		err := f()
		if err != nil {
			log.Errorf("err:%v", err)
			os.Exit(1)
		}
	}()
}
