//go:build windows

package runtime

func IsWindows() bool {
	return true
}

func IsDarwin() bool {
	return false
}

func IsLinux() bool {
	return false
}
