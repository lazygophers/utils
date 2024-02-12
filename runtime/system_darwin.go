//go:build darwin

package runtime

func IsWindows() bool {
	return false
}

func IsDarwin() bool {
	return true
}

func IsLinux() bool {
	return false
}
