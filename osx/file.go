package osx

import (
	"io"
	"io/fs"
	"os"
)

// Deprecated: Use Exist instead. Will be removed in next major version.
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func Exist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func FsHasFile(fs fs.FS, path string) bool {
	f, err := fs.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()
	return true
}

func RenameForce(oldpath, newpath string) (err error) {
	if Exists(newpath) {
		err = os.RemoveAll(newpath)
		if err != nil {
			return err
		}
	}

	err = os.Rename(oldpath, newpath)
	if err != nil {
		return err
	}
	return nil
}

func Copy(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	stat, err := srcFile.Stat()
	if err != nil {
		return err
	}

	// 安全的权限复制：Perm() 清除了 setuid/setgid/sticky 等特殊位
	perm := stat.Mode().Perm() & 0777

	//nosec G302 -- perm sanitized via Perm(), safe for file copying
	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}
