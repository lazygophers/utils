package osx

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
)

func TestExists(t *testing.T) {
	// 注意：这个函数有Bug，第12行应该是!os.IsNotExist(err)
	tests := []struct {
		name     string
		path     string
		setup    func() (string, func())
		expected bool
	}{
		{
			name: "existing file",
			setup: func() (string, func()) {
				tmpFile, err := os.CreateTemp("", "exists_test_*.txt")
				if err != nil {
					t.Fatal(err)
				}
				path := tmpFile.Name()
				tmpFile.Close()
				return path, func() { os.Remove(path) }
			},
			expected: true,
		},
		{
			name: "existing directory",
			setup: func() (string, func()) {
				tmpDir, err := os.MkdirTemp("", "exists_test_*")
				if err != nil {
					t.Fatal(err)
				}
				return tmpDir, func() { os.RemoveAll(tmpDir) }
			},
			expected: true,
		},
		{
			name:     "non-existing path",
			path:     "/non/existent/path/12345",
			setup:    func() (string, func()) { return "", func() {} },
			expected: false, // 由于Bug，实际返回false（正确）
		},
		{
			name:     "empty path",
			path:     "",
			setup:    func() (string, func()) { return "", func() {} },
			expected: false,
		},
		{
			name:     "invalid path characters",
			path:     string([]byte{0}),
			setup:    func() (string, func()) { return "", func() {} },
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := tt.path
			if tt.setup != nil {
				setupPath, cleanup := tt.setup()
				if setupPath != "" {
					path = setupPath
				}
				defer cleanup()
			}

			result := Exists(path)
			if result != tt.expected {
				t.Errorf("Exists(%s) = %v, expected %v", path, result, tt.expected)
			}
		})
	}
}

func TestExist(t *testing.T) {
	// 测试正确实现的Exist函数
	tests := []struct {
		name     string
		path     string
		setup    func() (string, func())
		expected bool
	}{
		{
			name: "existing file",
			setup: func() (string, func()) {
				tmpFile, err := os.CreateTemp("", "exist_test_*.txt")
				if err != nil {
					t.Fatal(err)
				}
				path := tmpFile.Name()
				tmpFile.Close()
				return path, func() { os.Remove(path) }
			},
			expected: true,
		},
		{
			name: "existing directory",
			setup: func() (string, func()) {
				tmpDir, err := os.MkdirTemp("", "exist_test_*")
				if err != nil {
					t.Fatal(err)
				}
				return tmpDir, func() { os.RemoveAll(tmpDir) }
			},
			expected: true,
		},
		{
			name:     "non-existing path",
			path:     "/non/existent/path/67890",
			setup:    func() (string, func()) { return "", func() {} },
			expected: false,
		},
		{
			name:     "empty path",
			path:     "",
			setup:    func() (string, func()) { return "", func() {} },
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := tt.path
			if tt.setup != nil {
				setupPath, cleanup := tt.setup()
				if setupPath != "" {
					path = setupPath
				}
				defer cleanup()
			}

			result := Exist(path)
			if result != tt.expected {
				t.Errorf("Exist(%s) = %v, expected %v", path, result, tt.expected)
			}
		})
	}
}

func TestIsDir(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		setup    func() (string, func())
		expected bool
	}{
		{
			name: "existing directory",
			setup: func() (string, func()) {
				tmpDir, err := os.MkdirTemp("", "isdir_test_*")
				if err != nil {
					t.Fatal(err)
				}
				return tmpDir, func() { os.RemoveAll(tmpDir) }
			},
			expected: true,
		},
		{
			name: "existing file (not directory)",
			setup: func() (string, func()) {
				tmpFile, err := os.CreateTemp("", "isdir_test_*.txt")
				if err != nil {
					t.Fatal(err)
				}
				path := tmpFile.Name()
				tmpFile.Close()
				return path, func() { os.Remove(path) }
			},
			expected: false,
		},
		{
			name:     "non-existing path",
			path:     "/non/existent/directory",
			setup:    func() (string, func()) { return "", func() {} },
			expected: false,
		},
		{
			name:     "empty path",
			path:     "",
			setup:    func() (string, func()) { return "", func() {} },
			expected: false,
		},
		{
			name:     "current directory",
			path:     ".",
			setup:    func() (string, func()) { return "", func() {} },
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := tt.path
			if tt.setup != nil {
				setupPath, cleanup := tt.setup()
				if setupPath != "" {
					path = setupPath
				}
				defer cleanup()
			}

			result := IsDir(path)
			if result != tt.expected {
				t.Errorf("IsDir(%s) = %v, expected %v", path, result, tt.expected)
			}
		})
	}
}

func TestIsFile(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		setup    func() (string, func())
		expected bool
	}{
		{
			name: "existing file",
			setup: func() (string, func()) {
				tmpFile, err := os.CreateTemp("", "isfile_test_*.txt")
				if err != nil {
					t.Fatal(err)
				}
				path := tmpFile.Name()
				tmpFile.Close()
				return path, func() { os.Remove(path) }
			},
			expected: true,
		},
		{
			name: "existing directory (not file)",
			setup: func() (string, func()) {
				tmpDir, err := os.MkdirTemp("", "isfile_test_*")
				if err != nil {
					t.Fatal(err)
				}
				return tmpDir, func() { os.RemoveAll(tmpDir) }
			},
			expected: false,
		},
		{
			name:     "non-existing path",
			path:     "/non/existent/file.txt",
			setup:    func() (string, func()) { return "", func() {} },
			expected: false,
		},
		{
			name:     "empty path",
			path:     "",
			setup:    func() (string, func()) { return "", func() {} },
			expected: false,
		},
		{
			name:     "current directory (not file)",
			path:     ".",
			setup:    func() (string, func()) { return "", func() {} },
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := tt.path
			if tt.setup != nil {
				setupPath, cleanup := tt.setup()
				if setupPath != "" {
					path = setupPath
				}
				defer cleanup()
			}

			result := IsFile(path)
			if result != tt.expected {
				t.Errorf("IsFile(%s) = %v, expected %v", path, result, tt.expected)
			}
		})
	}
}

func TestFsHasFile(t *testing.T) {
	// 创建测试文件系统
	testFS := fstest.MapFS{
		"file.txt":        &fstest.MapFile{Data: []byte("test content")},
		"dir/subfile.txt": &fstest.MapFile{Data: []byte("sub content")},
		"empty.txt":       &fstest.MapFile{Data: []byte("")},
	}

	tests := []struct {
		name     string
		fs       fs.FS
		path     string
		expected bool
	}{
		{
			name:     "existing file in root",
			fs:       testFS,
			path:     "file.txt",
			expected: true,
		},
		{
			name:     "existing file in subdirectory",
			fs:       testFS,
			path:     "dir/subfile.txt",
			expected: true,
		},
		{
			name:     "existing empty file",
			fs:       testFS,
			path:     "empty.txt",
			expected: true,
		},
		{
			name:     "non-existing file",
			fs:       testFS,
			path:     "nonexistent.txt",
			expected: false,
		},
		{
			name:     "non-existing subdirectory file",
			fs:       testFS,
			path:     "nondir/file.txt",
			expected: false,
		},
		{
			name:     "empty path",
			fs:       testFS,
			path:     "",
			expected: false, // 通常空路径会导致错误
		},
		{
			name:     "directory path (not file)",
			fs:       testFS,
			path:     "dir",
			expected: true, // MapFS中目录也可以"打开"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FsHasFile(tt.fs, tt.path)
			if result != tt.expected {
				t.Errorf("FsHasFile(fs, %s) = %v, expected %v", tt.path, result, tt.expected)
			}
		})
	}
}

func TestFsHasFile_RealFS(t *testing.T) {
	// 使用真实文件系统测试
	tmpDir, err := os.MkdirTemp("", "fshasfile_test_*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// 创建测试文件
	testFile := filepath.Join(tmpDir, "test.txt")
	err = os.WriteFile(testFile, []byte("test"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// 创建子目录和文件
	subDir := filepath.Join(tmpDir, "subdir")
	err = os.Mkdir(subDir, 0755)
	if err != nil {
		t.Fatal(err)
	}

	subFile := filepath.Join(subDir, "sub.txt")
	err = os.WriteFile(subFile, []byte("sub"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// 使用os.DirFS测试
	fsys := os.DirFS(tmpDir)

	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{"existing file", "test.txt", true},
		{"existing subdir file", "subdir/sub.txt", true},
		{"non-existing file", "nonexistent.txt", false},
		{"directory", "subdir", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FsHasFile(fsys, tt.path)
			if result != tt.expected {
				t.Errorf("FsHasFile(DirFS, %s) = %v, expected %v", tt.path, result, tt.expected)
			}
		})
	}
}
