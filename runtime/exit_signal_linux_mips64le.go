//go:build linux && mips64le
// +build linux,mips64le

package runtime

import (
	"os"
	"syscall"
)

var exitSignal = []os.Signal{
	syscall.SIGINT,
	syscall.SIGQUIT,
	syscall.SIGABRT,
	syscall.SIGKILL,
	syscall.SIGTERM,
	syscall.SIGSTOP,
	syscall.SIGTRAP,
	syscall.SIGTSTP,
}
