package runtime

import (
	"github.com/lazygophers/log"
	"github.com/lazygophers/utils/app"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
)

func CachePanic() {
	CachePanicWithHandle(nil)
}

func CachePanicWithHandle(handle func(err interface{})) {
	if err := recover(); err != nil {
		log.Errorf("PROCESS PANIC: err %s", err)
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
		if handle != nil {
			handle(err)
		}
	}
}

func PrintStack() {
	st := debug.Stack()
	if len(st) > 0 {
		log.Error("dump stack:")
		log.Error(string(st))
	} else {
		log.Error("stack is empty")
	}
}

func ExecDir() string {
	execPath, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(execPath)
}

func ExecFile() string {
	execPath, err := os.Executable()
	if err != nil {
		return ""
	}
	return execPath
}

func Pwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return pwd
}

func UserHomeDir() string {
	path, _ := os.UserHomeDir()
	return path
}

func UserConfigDir() string {
	path, _ := os.UserConfigDir()
	return path
}

func UserCacheDir() string {
	execPath, _ := os.UserCacheDir()
	return execPath
}

func LazyConfigDir() string {
	path, _ := os.UserConfigDir()
	return filepath.Join(path, app.Organization)
}
