package osx

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRenameForce(t *testing.T) {
	tests := []struct {
		name      string
		setup     func() (oldpath, newpath string, cleanup func())
		expectErr bool
	}{
		{
			name: "rename file to non-existing path",
			setup: func() (string, string, func()) {
				tmpDir, err := os.MkdirTemp("", "rename_test_*")
				if err != nil {
					t.Fatal(err)
				}

				oldFile := filepath.Join(tmpDir, "old.txt")
				newFile := filepath.Join(tmpDir, "new.txt")

				err = os.WriteFile(oldFile, []byte("test content"), 0644)
				if err != nil {
					t.Fatal(err)
				}

				return oldFile, newFile, func() { os.RemoveAll(tmpDir) }
			},
			expectErr: false,
		},
		{
			name: "rename file overwriting existing file",
			setup: func() (string, string, func()) {
				tmpDir, err := os.MkdirTemp("", "rename_test_*")
				if err != nil {
					t.Fatal(err)
				}

				oldFile := filepath.Join(tmpDir, "old.txt")
				newFile := filepath.Join(tmpDir, "existing.txt")

				err = os.WriteFile(oldFile, []byte("old content"), 0644)
				if err != nil {
					t.Fatal(err)
				}

				err = os.WriteFile(newFile, []byte("existing content"), 0644)
				if err != nil {
					t.Fatal(err)
				}

				return oldFile, newFile, func() { os.RemoveAll(tmpDir) }
			},
			expectErr: false,
		},
		{
			name: "rename file overwriting existing directory",
			setup: func() (string, string, func()) {
				tmpDir, err := os.MkdirTemp("", "rename_test_*")
				if err != nil {
					t.Fatal(err)
				}

				oldFile := filepath.Join(tmpDir, "old.txt")
				newDir := filepath.Join(tmpDir, "existing_dir")

				err = os.WriteFile(oldFile, []byte("old content"), 0644)
				if err != nil {
					t.Fatal(err)
				}

				err = os.Mkdir(newDir, 0755)
				if err != nil {
					t.Fatal(err)
				}

				// 在目录中创建一个文件，使其非空
				subFile := filepath.Join(newDir, "sub.txt")
				err = os.WriteFile(subFile, []byte("sub content"), 0644)
				if err != nil {
					t.Fatal(err)
				}

				return oldFile, newDir, func() { os.RemoveAll(tmpDir) }
			},
			expectErr: false,
		},
		{
			name: "rename non-existing file",
			setup: func() (string, string, func()) {
				tmpDir, err := os.MkdirTemp("", "rename_test_*")
				if err != nil {
					t.Fatal(err)
				}

				oldFile := filepath.Join(tmpDir, "nonexistent.txt")
				newFile := filepath.Join(tmpDir, "new.txt")

				return oldFile, newFile, func() { os.RemoveAll(tmpDir) }
			},
			expectErr: true,
		},
		{
			name: "rename to invalid path",
			setup: func() (string, string, func()) {
				tmpDir, err := os.MkdirTemp("", "rename_test_*")
				if err != nil {
					t.Fatal(err)
				}

				oldFile := filepath.Join(tmpDir, "old.txt")
				err = os.WriteFile(oldFile, []byte("test"), 0644)
				if err != nil {
					t.Fatal(err)
				}

				// 无效路径（包含null字符）
				invalidPath := string([]byte{0})

				return oldFile, invalidPath, func() { os.RemoveAll(tmpDir) }
			},
			expectErr: true,
		},
		{
			name: "rename directory to non-existing path",
			setup: func() (string, string, func()) {
				tmpDir, err := os.MkdirTemp("", "rename_test_*")
				if err != nil {
					t.Fatal(err)
				}

				oldDir := filepath.Join(tmpDir, "old_dir")
				newDir := filepath.Join(tmpDir, "new_dir")

				err = os.Mkdir(oldDir, 0755)
				if err != nil {
					t.Fatal(err)
				}

				// 在目录中创建文件
				subFile := filepath.Join(oldDir, "file.txt")
				err = os.WriteFile(subFile, []byte("content"), 0644)
				if err != nil {
					t.Fatal(err)
				}

				return oldDir, newDir, func() { os.RemoveAll(tmpDir) }
			},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldpath, newpath, cleanup := tt.setup()
			defer cleanup()

			err := RenameForce(oldpath, newpath)

			if tt.expectErr && err == nil {
				t.Errorf("RenameForce(%s, %s) expected error, got nil", oldpath, newpath)
			}
			if !tt.expectErr && err != nil {
				t.Errorf("RenameForce(%s, %s) unexpected error: %v", oldpath, newpath, err)
			}

			// 如果操作成功，验证结果
			if !tt.expectErr && err == nil {
				// 检查新路径是否存在
				if !Exist(newpath) {
					t.Errorf("After RenameForce, newpath %s should exist", newpath)
				}
				// 检查旧路径是否不存在
				if Exist(oldpath) {
					t.Errorf("After RenameForce, oldpath %s should not exist", oldpath)
				}
			}
		})
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		name      string
		setup     func() (src, dst string, cleanup func())
		expectErr bool
		verify    func(t *testing.T, src, dst string)
	}{
		{
			name: "copy file to new location",
			setup: func() (string, string, func()) {
				tmpDir, err := os.MkdirTemp("", "copy_test_*")
				if err != nil {
					t.Fatal(err)
				}

				src := filepath.Join(tmpDir, "source.txt")
				dst := filepath.Join(tmpDir, "dest.txt")

				content := "This is test content for copying"
				err = os.WriteFile(src, []byte(content), 0644)
				if err != nil {
					t.Fatal(err)
				}

				return src, dst, func() { os.RemoveAll(tmpDir) }
			},
			expectErr: false,
			verify: func(t *testing.T, src, dst string) {
				// 验证目标文件存在
				if !Exist(dst) {
					t.Errorf("Destination file %s should exist after copy", dst)
					return
				}

				// 验证内容相同
				srcContent, err := os.ReadFile(src)
				if err != nil {
					t.Errorf("Failed to read source file: %v", err)
					return
				}

				dstContent, err := os.ReadFile(dst)
				if err != nil {
					t.Errorf("Failed to read destination file: %v", err)
					return
				}

				if string(srcContent) != string(dstContent) {
					t.Errorf("File content mismatch: src=%s, dst=%s", string(srcContent), string(dstContent))
				}

				// 验证文件权限
				srcInfo, err := os.Stat(src)
				if err != nil {
					t.Errorf("Failed to stat source file: %v", err)
					return
				}

				dstInfo, err := os.Stat(dst)
				if err != nil {
					t.Errorf("Failed to stat destination file: %v", err)
					return
				}

				if srcInfo.Mode() != dstInfo.Mode() {
					t.Errorf("File mode mismatch: src=%v, dst=%v", srcInfo.Mode(), dstInfo.Mode())
				}
			},
		},
		{
			name: "copy empty file",
			setup: func() (string, string, func()) {
				tmpDir, err := os.MkdirTemp("", "copy_test_*")
				if err != nil {
					t.Fatal(err)
				}

				src := filepath.Join(tmpDir, "empty.txt")
				dst := filepath.Join(tmpDir, "empty_copy.txt")

				err = os.WriteFile(src, []byte(""), 0644)
				if err != nil {
					t.Fatal(err)
				}

				return src, dst, func() { os.RemoveAll(tmpDir) }
			},
			expectErr: false,
			verify: func(t *testing.T, src, dst string) {
				content, err := os.ReadFile(dst)
				if err != nil {
					t.Errorf("Failed to read destination file: %v", err)
					return
				}
				if len(content) != 0 {
					t.Errorf("Expected empty file, got %d bytes", len(content))
				}
			},
		},
		{
			name: "copy file with special permissions",
			setup: func() (string, string, func()) {
				tmpDir, err := os.MkdirTemp("", "copy_test_*")
				if err != nil {
					t.Fatal(err)
				}

				src := filepath.Join(tmpDir, "special.txt")
				dst := filepath.Join(tmpDir, "special_copy.txt")

				err = os.WriteFile(src, []byte("special content"), 0600) // 只有所有者可读写
				if err != nil {
					t.Fatal(err)
				}

				return src, dst, func() { os.RemoveAll(tmpDir) }
			},
			expectErr: false,
			verify: func(t *testing.T, src, dst string) {
				srcInfo, _ := os.Stat(src)
				dstInfo, _ := os.Stat(dst)

				if srcInfo.Mode() != dstInfo.Mode() {
					t.Errorf("Permission not preserved: src=%v, dst=%v", srcInfo.Mode(), dstInfo.Mode())
				}
			},
		},
		{
			name: "copy overwrite existing file",
			setup: func() (string, string, func()) {
				tmpDir, err := os.MkdirTemp("", "copy_test_*")
				if err != nil {
					t.Fatal(err)
				}

				src := filepath.Join(tmpDir, "source.txt")
				dst := filepath.Join(tmpDir, "existing.txt")

				err = os.WriteFile(src, []byte("new content"), 0644)
				if err != nil {
					t.Fatal(err)
				}

				err = os.WriteFile(dst, []byte("old content"), 0644)
				if err != nil {
					t.Fatal(err)
				}

				return src, dst, func() { os.RemoveAll(tmpDir) }
			},
			expectErr: false,
			verify: func(t *testing.T, src, dst string) {
				content, err := os.ReadFile(dst)
				if err != nil {
					t.Errorf("Failed to read destination: %v", err)
					return
				}
				if string(content) != "new content" {
					t.Errorf("Expected 'new content', got '%s'", string(content))
				}
			},
		},
		{
			name: "copy non-existing source file",
			setup: func() (string, string, func()) {
				tmpDir, err := os.MkdirTemp("", "copy_test_*")
				if err != nil {
					t.Fatal(err)
				}

				src := filepath.Join(tmpDir, "nonexistent.txt")
				dst := filepath.Join(tmpDir, "dest.txt")

				return src, dst, func() { os.RemoveAll(tmpDir) }
			},
			expectErr: true,
			verify:    func(t *testing.T, src, dst string) {},
		},
		{
			name: "copy to invalid destination",
			setup: func() (string, string, func()) {
				tmpDir, err := os.MkdirTemp("", "copy_test_*")
				if err != nil {
					t.Fatal(err)
				}

				src := filepath.Join(tmpDir, "source.txt")
				err = os.WriteFile(src, []byte("test"), 0644)
				if err != nil {
					t.Fatal(err)
				}

				// 无效目标路径
				dst := "/invalid/nonexistent/path/dest.txt"

				return src, dst, func() { os.RemoveAll(tmpDir) }
			},
			expectErr: true,
			verify:    func(t *testing.T, src, dst string) {},
		},
		{
			name: "copy directory (should fail)",
			setup: func() (string, string, func()) {
				tmpDir, err := os.MkdirTemp("", "copy_test_*")
				if err != nil {
					t.Fatal(err)
				}

				srcDir := filepath.Join(tmpDir, "source_dir")
				dst := filepath.Join(tmpDir, "dest.txt")

				err = os.Mkdir(srcDir, 0755)
				if err != nil {
					t.Fatal(err)
				}

				return srcDir, dst, func() { os.RemoveAll(tmpDir) }
			},
			expectErr: true, // Copy函数不支持目录复制
			verify:    func(t *testing.T, src, dst string) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src, dst, cleanup := tt.setup()
			defer cleanup()

			err := Copy(src, dst)

			if tt.expectErr && err == nil {
				t.Errorf("Copy(%s, %s) expected error, got nil", src, dst)
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Copy(%s, %s) unexpected error: %v", src, dst, err)
			}

			// 如果操作成功，运行验证函数
			if !tt.expectErr && err == nil {
				tt.verify(t, src, dst)
			}
		})
	}
}

func TestCopy_LargeFile(t *testing.T) {
	// 测试大文件复制
	tmpDir, err := os.MkdirTemp("", "copy_large_test_*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	src := filepath.Join(tmpDir, "large.txt")
	dst := filepath.Join(tmpDir, "large_copy.txt")

	// 创建1MB的测试文件
	largeContent := make([]byte, 1024*1024)
	for i := range largeContent {
		largeContent[i] = byte(i % 256)
	}

	err = os.WriteFile(src, largeContent, 0644)
	if err != nil {
		t.Fatal(err)
	}

	err = Copy(src, dst)
	if err != nil {
		t.Fatalf("Failed to copy large file: %v", err)
	}

	// 验证文件大小
	srcInfo, err := os.Stat(src)
	if err != nil {
		t.Fatal(err)
	}

	dstInfo, err := os.Stat(dst)
	if err != nil {
		t.Fatal(err)
	}

	if srcInfo.Size() != dstInfo.Size() {
		t.Errorf("File size mismatch: src=%d, dst=%d", srcInfo.Size(), dstInfo.Size())
	}

	// 验证内容（抽样检查）
	dstContent, err := os.ReadFile(dst)
	if err != nil {
		t.Fatal(err)
	}

	if len(dstContent) != len(largeContent) {
		t.Errorf("Content length mismatch: expected %d, got %d", len(largeContent), len(dstContent))
	}

	// 检查前后几个字节
	for i := 0; i < 10; i++ {
		if dstContent[i] != largeContent[i] {
			t.Errorf("Content mismatch at position %d: expected %d, got %d", i, largeContent[i], dstContent[i])
		}
	}

	for i := len(largeContent) - 10; i < len(largeContent); i++ {
		if dstContent[i] != largeContent[i] {
			t.Errorf("Content mismatch at position %d: expected %d, got %d", i, largeContent[i], dstContent[i])
		}
	}
}
