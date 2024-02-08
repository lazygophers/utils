//go:build !linux

package atexit

import "modernc.org/libc"

func Register(callback func()) {
	libc.AtExit(callback)
}
