package runtime

import (
	"fmt"
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
		// 使用标准错误输出，避免日志系统的递归调用
		fmt.Fprintf(os.Stderr, "PROCESS PANIC: err %v\n", err)
		st := debug.Stack()
		if len(st) > 0 {
			fmt.Fprintf(os.Stderr, "dump stack (%v):\n", err)
			lines := strings.Split(string(st), "\n")
			for _, line := range lines {
				fmt.Fprintf(os.Stderr, "  %s\n", line)
			}
		} else {
			fmt.Fprintf(os.Stderr, "stack is empty (%v)\n", err)
		}
		if handle != nil {
			handle(err)
		}
		// 不再重新panic，真正"缓存"（消化）panic
	}
}

func PrintStack() {
	st := debug.Stack()
	if len(st) > 0 {
		// 使用标准错误输出，避免日志系统的递归调用
		fmt.Fprintf(os.Stderr, "dump stack:\n%s\n", string(st))
	} else {
		fmt.Fprintf(os.Stderr, "stack is empty\n")
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
	path, _ := os.UserCacheDir()
	return path
}

func LazyConfigDir() string {
	path, _ := os.UserConfigDir()
	return filepath.Join(path, app.Organization)
}

func LazyCacheDir() string {
	path, _ := os.UserCacheDir()
	return filepath.Join(path, app.Organization)
}
