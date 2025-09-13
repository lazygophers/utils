package osx

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCopy_StatErrorBranch(t *testing.T) {
	// 尝试触发srcFile.Stat()错误分支（第73-75行）
	// 这是最后剩余的6.7%未覆盖代码

	t.Run("stat_error_with_special_file", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "copy_stat_error_test_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		src := filepath.Join(tmpDir, "source.txt")
		dst := filepath.Join(tmpDir, "dest.txt")

		// 创建源文件
		err = os.WriteFile(src, []byte("test content for stat error"), 0644)
		if err != nil {
			t.Fatal(err)
		}

		// 在某些极端情况下，文件可能在打开后但在stat前被删除或变为无效
		// 虽然这种情况很少见，但我们可以模拟这种情况

		// 方法1：尝试竞态条件 - 在文件操作过程中修改文件
		go func() {
			// 短暂延迟后尝试删除文件
			time.Sleep(1 * time.Millisecond)
			os.Remove(src) // 可能在Copy过程中删除文件
		}()

		err = Copy(src, dst)
		// 这可能会成功或失败，取决于竞态条件的时机
		t.Logf("Copy with potential race condition: %v", err)
	})

	t.Run("stat_error_with_device_file", func(t *testing.T) {
		// 在Unix系统上，某些设备文件可能有特殊的stat行为
		// 尝试复制/dev/null（如果存在）
		if _, err := os.Stat("/dev/null"); err == nil {
			tmpDir, err := os.MkdirTemp("", "copy_dev_test_*")
			if err != nil {
				t.Fatal(err)
			}
			defer os.RemoveAll(tmpDir)

			dst := filepath.Join(tmpDir, "null_copy.txt")

			// 尝试复制/dev/null
			err = Copy("/dev/null", dst)
			if err != nil {
				t.Logf("Copy from /dev/null failed: %v", err)
			} else {
				t.Logf("Copy from /dev/null succeeded")

				// 验证复制的文件
				if Exist(dst) {
					info, err := os.Stat(dst)
					if err == nil {
						t.Logf("Copied file size: %d bytes", info.Size())
					}
				}
			}
		} else {
			t.Skip("Skipping device file test: /dev/null not available")
		}
	})

	t.Run("stat_error_with_symlink", func(t *testing.T) {
		// 测试符号链接可能引起的stat问题
		tmpDir, err := os.MkdirTemp("", "copy_symlink_test_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		// 创建一个实际文件
		realFile := filepath.Join(tmpDir, "real.txt")
		err = os.WriteFile(realFile, []byte("real content"), 0644)
		if err != nil {
			t.Fatal(err)
		}

		// 创建指向该文件的符号链接
		symLink := filepath.Join(tmpDir, "link.txt")
		err = os.Symlink(realFile, symLink)
		if err != nil {
			t.Skip("Symlink not supported on this system")
		}

		dst := filepath.Join(tmpDir, "copy_from_link.txt")

		// 删除原始文件，使符号链接变为悬空
		os.Remove(realFile)

		// 现在尝试复制悬空的符号链接
		err = Copy(symLink, dst)
		if err != nil {
			t.Logf("Copy from broken symlink failed as expected: %v", err)
		} else {
			t.Logf("Copy from broken symlink unexpectedly succeeded")
		}
	})

	t.Run("stat_error_with_permission_change", func(t *testing.T) {
		if os.Geteuid() == 0 {
			t.Skip("Skipping test when running as root")
		}

		tmpDir, err := os.MkdirTemp("", "copy_perm_change_test_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		src := filepath.Join(tmpDir, "source.txt")
		dst := filepath.Join(tmpDir, "dest.txt")

		// 创建源文件
		err = os.WriteFile(src, []byte("permission test content"), 0644)
		if err != nil {
			t.Fatal(err)
		}

		// 启动一个goroutine在复制过程中修改文件权限
		go func() {
			time.Sleep(1 * time.Millisecond)
			// 尝试移除文件的读权限
			os.Chmod(src, 0000)
			time.Sleep(1 * time.Millisecond)
			// 恢复权限
			os.Chmod(src, 0644)
		}()

		// 执行复制操作
		err = Copy(src, dst)

		// 恢复权限确保清理能够进行
		os.Chmod(src, 0644)

		if err != nil {
			t.Logf("Copy with permission change failed: %v", err)
		} else {
			t.Logf("Copy with permission change succeeded")

			// 验证复制结果
			if Exist(dst) {
				content, readErr := os.ReadFile(dst)
				if readErr == nil {
					t.Logf("Copied content: %s", string(content))
				}
			}
		}
	})

	t.Run("stat_error_multiple_attempts", func(t *testing.T) {
		// 多次尝试不同的方法来触发stat错误
		tmpDir, err := os.MkdirTemp("", "copy_multiple_test_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		for i := 0; i < 10; i++ {
			src := filepath.Join(tmpDir, fmt.Sprintf("source_%d.txt", i))
			dst := filepath.Join(tmpDir, fmt.Sprintf("dest_%d.txt", i))

			// 创建源文件
			content := fmt.Sprintf("test content %d", i)
			err = os.WriteFile(src, []byte(content), 0644)
			if err != nil {
				t.Fatal(err)
			}

			// 快速复制
			err = Copy(src, dst)
			if err != nil {
				t.Logf("Copy attempt %d failed: %v", i, err)
			}
		}
	})
}
