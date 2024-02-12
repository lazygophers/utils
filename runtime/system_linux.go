//go:build linux

package runtime

func IsWindows() bool {
	return false
}

func IsDarwin() bool {
	return false
}

func IsLinux() bool {
	return true
}
