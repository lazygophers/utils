//go:build plan9
// +build plan9

package runtime

import (
	"os"
	"syscall"
)

var exitSignal = []os.Signal{
	syscall.SIGINT,
	syscall.SIGABRT,
	syscall.SIGKILL,
	syscall.SIGTERM,
}
