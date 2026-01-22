package runtime

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"

	"github.com/lazygophers/log"
	"github.com/lazygophers/utils/app"
)

func CachePanic() {
	CachePanicWithHandle(nil)
}

type panicHandler func(err interface{})

var panicHandlerList []panicHandler

func OnPanic(logic panicHandler) {
	panicHandlerList = append(panicHandlerList, logic)
}

func CachePanicWithHandle(handle func(err interface{})) {
	if err := recover(); err != nil {
		// 使用最基础的系统调用避免栈溢出
		os.Stderr.WriteString("PROCESS PANIC: err ")
		os.Stderr.WriteString(fmt.Sprintf("%v", err))
		os.Stderr.WriteString("\n")

		st := debug.Stack()
		if len(st) > 0 {
			os.Stderr.WriteString("dump stack (")
			os.Stderr.WriteString(fmt.Sprintf("%v", err))
			os.Stderr.WriteString("):\n")

			// 分块处理栈跟踪信息，添加缩进
			lines := strings.Split(string(st), "\n")
			for _, line := range lines {
				os.Stderr.WriteString("  ")
				os.Stderr.WriteString(line)
				os.Stderr.WriteString("\n")
			}
		} else {
			os.Stderr.WriteString("stack is empty (")
			os.Stderr.WriteString(fmt.Sprintf("%v", err))
			os.Stderr.WriteString(")\n")
		}
		if handle != nil {
			handle(err)
		}

		for _, handler := range panicHandlerList {
			handler(err)
		}
	}
}

func PrintStack() {
	st := debug.Stack()
	if len(st) > 0 {
		// 使用最基础的系统调用避免栈溢出
		os.Stderr.WriteString("dump stack:\n")
		// 分块写入大的栈跟踪信息
		const chunkSize = 1024
		for i := 0; i < len(st); i += chunkSize {
			end := i + chunkSize
			if end > len(st) {
				end = len(st)
			}
			os.Stderr.Write(st[i:end])
		}
		os.Stderr.WriteString("\n")
	} else {
		os.Stderr.WriteString("stack is empty\n")
	}
}

func GetStack() string {
	b := log.GetBuffer()
	defer log.PutBuffer(b)

	b.WriteString("dump stack:\n")
	b.Write(debug.Stack())

	return b.String()
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
