//go:build !plan9 && !windows && !darwin && !dragonfly && !freebsd && !linux && !solaris && !openbsd && !netbsd
// +build !plan9,!windows,!darwin,!dragonfly,!freebsd,!linux,!solaris,!openbsd,!netbsd

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
