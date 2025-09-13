package osx

import (
	"os"
	"path/filepath"
	"syscall"
	"testing"
)

// 最后的尝试：通过文件描述符操作来模拟stat错误
func TestCopy_DescriptorStatError(t *testing.T) {
	t.Run("copy_with_closed_descriptor_race", func(t *testing.T) {
		// 这种方法虽然有风险，但是可能能够触发stat错误
		tmpDir, err := os.MkdirTemp("", "copy_descriptor_test_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		src := filepath.Join(tmpDir, "source.txt")
		dst := filepath.Join(tmpDir, "dest.txt")

		// 创建一个源文件
		err = os.WriteFile(src, []byte("descriptor test content"), 0644)
		if err != nil {
			t.Fatal(err)
		}

		// 尝试多次，看看能否触发竞态条件
		for attempt := 0; attempt < 100; attempt++ {
			err = Copy(src, dst)
			if err != nil {
				if attempt == 0 {
					t.Logf("Copy attempt %d failed: %v", attempt, err)
				}
			}
			// 清理目标文件以便下次尝试
			os.Remove(dst)
		}
	})

	t.Run("copy_with_invalid_file_operations", func(t *testing.T) {
		// 尝试创建一些可能导致stat失败的特殊文件
		tmpDir, err := os.MkdirTemp("", "copy_invalid_test_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		// 测试不同的文件类型和权限组合
		testFiles := []struct {
			name string
			mode os.FileMode
			data []byte
		}{
			{"socket_like", 0600, []byte("socket test")},
			{"device_like", 0660, []byte("device test")},
			{"special_perm", 0000, []byte("special permission test")},
			{"setuid", 04755, []byte("setuid test")},
			{"setgid", 02755, []byte("setgid test")},
			{"sticky", 01755, []byte("sticky test")},
		}

		for _, tf := range testFiles {
			src := filepath.Join(tmpDir, tf.name+"_src.txt")
			dst := filepath.Join(tmpDir, tf.name+"_dst.txt")

			// 创建文件
			err = os.WriteFile(src, tf.data, 0644)
			if err != nil {
				t.Logf("Cannot create file %s: %v", tf.name, err)
				continue
			}

			// 设置特殊权限
			err = os.Chmod(src, tf.mode)
			if err != nil {
				t.Logf("Cannot set mode for %s: %v", tf.name, err)
				os.Chmod(src, 0644) // 恢复权限以便清理
				continue
			}

			// 尝试复制
			err = Copy(src, dst)

			// 恢复权限以便清理
			os.Chmod(src, 0644)

			if err != nil {
				t.Logf("Copy %s failed: %v", tf.name, err)
			} else {
				t.Logf("Copy %s succeeded", tf.name)
			}
		}
	})
}

// 使用系统调用的方式尝试触发stat错误
func TestCopy_SystemCallStatError(t *testing.T) {
	t.Run("copy_with_manual_fd_manipulation", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "copy_syscall_test_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		src := filepath.Join(tmpDir, "syscall_source.txt")
		dst := filepath.Join(tmpDir, "syscall_dest.txt")

		// 创建源文件
		err = os.WriteFile(src, []byte("syscall test content"), 0644)
		if err != nil {
			t.Fatal(err)
		}

		// 手动打开文件
		fd, err := syscall.Open(src, syscall.O_RDONLY, 0)
		if err != nil {
			t.Fatal(err)
		}
		defer syscall.Close(fd)

		// 获取文件状态
		var stat syscall.Stat_t
		err = syscall.Fstat(fd, &stat)
		if err != nil {
			t.Logf("Manual fstat failed: %v", err)
		} else {
			t.Logf("Manual fstat succeeded, size: %d", stat.Size)
		}

		// 正常执行Copy操作
		err = Copy(src, dst)
		if err != nil {
			t.Logf("Copy after manual fd operations failed: %v", err)
		} else {
			t.Logf("Copy after manual fd operations succeeded")
		}
	})

	t.Run("copy_with_truncated_file", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "copy_truncate_test_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		src := filepath.Join(tmpDir, "truncate_source.txt")
		dst := filepath.Join(tmpDir, "truncate_dest.txt")

		// 创建一个较大的文件
		largeContent := make([]byte, 1024)
		for i := range largeContent {
			largeContent[i] = byte(i % 256)
		}

		err = os.WriteFile(src, largeContent, 0644)
		if err != nil {
			t.Fatal(err)
		}

		// 在另一个goroutine中truncate文件
		go func() {
			// 等待一点点时间
			for i := 0; i < 10; i++ {
				// 尝试截断文件
				if f, err := os.OpenFile(src, os.O_WRONLY, 0644); err == nil {
					f.Truncate(100) // 截断到100字节
					f.Close()
				}
			}
		}()

		// 执行复制
		err = Copy(src, dst)
		if err != nil {
			t.Logf("Copy with truncation failed: %v", err)
		} else {
			t.Logf("Copy with truncation succeeded")

			// 检查复制结果的大小
			if info, err := os.Stat(dst); err == nil {
				t.Logf("Copied file size: %d bytes", info.Size())
			}
		}
	})
}

// 尝试创建一个可能导致stat错误的自定义情况
func TestCopy_CustomStatError(t *testing.T) {
	t.Run("copy_with_file_replacement", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "copy_replacement_test_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		src := filepath.Join(tmpDir, "replacement_source.txt")
		dst := filepath.Join(tmpDir, "replacement_dest.txt")

		// 创建源文件
		originalContent := []byte("original content for replacement test")
		err = os.WriteFile(src, originalContent, 0644)
		if err != nil {
			t.Fatal(err)
		}

		// 在复制过程中替换文件
		go func() {
			// 等待复制开始
			for i := 0; i < 50; i++ {
				// 用不同内容替换源文件
				newContent := []byte("replacement content " + string(rune('0'+i%10)))
				os.WriteFile(src, newContent, 0644)
			}
		}()

		// 执行复制
		err = Copy(src, dst)
		if err != nil {
			t.Logf("Copy with file replacement failed: %v", err)
		} else {
			t.Logf("Copy with file replacement succeeded")

			// 检查最终复制的内容
			if content, err := os.ReadFile(dst); err == nil {
				maxLen := 50
				if len(content) < maxLen {
					maxLen = len(content)
				}
				t.Logf("Final copied content: %s", string(content[:maxLen]))
			}
		}
	})
}
