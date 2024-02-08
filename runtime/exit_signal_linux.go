//go:build linux && !mips64 && !mips64le && !mips && !mipsle && !arm64
// +build linux,!mips64,!mips64le,!mips,!mipsle,!arm64

package runtime

import (
	"os"
	"syscall"
)

var exitSignal = []os.Signal{
	syscall.SIGINT,  // 中断
	syscall.SIGQUIT, // 退出
	syscall.SIGABRT, // 中止
	syscall.SIGKILL, // kill信号
	syscall.SIGTERM, // 终止信号
	syscall.SIGSTOP, // 进程停止
	syscall.SIGTRAP, // 陷阱
	syscall.SIGTSTP, // 终端停止
}
