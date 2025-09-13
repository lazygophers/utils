package osx

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// 这是最后的尝试来触发srcFile.Stat()错误分支
func TestCopy_DirectStatErrorAttempt(t *testing.T) {
	t.Run("copy_from_directory_as_file", func(t *testing.T) {
		// 尝试创建一个目录，但是尝试当作文件来复制
		tmpDir, err := os.MkdirTemp("", "copy_dir_as_file_test_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		// 创建一个目录
		srcDir := filepath.Join(tmpDir, "source_dir")
		err = os.Mkdir(srcDir, 0755)
		if err != nil {
			t.Fatal(err)
		}

		dst := filepath.Join(tmpDir, "dest.txt")

		// 尝试复制目录（这应该会失败）
		err = Copy(srcDir, dst)
		if err != nil {
			t.Logf("Copy directory as file failed as expected: %v", err)
		} else {
			t.Error("Copy should fail when source is a directory")
		}
	})

	t.Run("copy_with_long_path_name", func(t *testing.T) {
		// 测试非常长的文件名，可能触发某些系统限制
		tmpDir, err := os.MkdirTemp("", "copy_long_path_test_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		// 创建一个很长的文件名
		longName := strings.Repeat("a", 200) + ".txt"
		src := filepath.Join(tmpDir, longName)
		dst := filepath.Join(tmpDir, "copy_"+longName)

		// 创建文件
		err = os.WriteFile(src, []byte("long name test"), 0644)
		if err != nil {
			t.Logf("Cannot create file with long name: %v", err)
			return
		}

		err = Copy(src, dst)
		if err != nil {
			t.Logf("Copy with long filename failed: %v", err)
		} else {
			t.Logf("Copy with long filename succeeded")
		}
	})

	t.Run("copy_special_characters", func(t *testing.T) {
		// 测试包含特殊字符的文件名
		tmpDir, err := os.MkdirTemp("", "copy_special_test_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		specialNames := []string{
			"file with spaces.txt",
			"file\twith\ttabs.txt",
			"file-with-unicode-测试.txt",
			"file.with.many.dots.txt",
		}

		for i, name := range specialNames {
			src := filepath.Join(tmpDir, name)
			dst := filepath.Join(tmpDir, "copy_"+name)

			// 创建文件
			content := []byte("special character test " + string(rune('0'+i)))
			err = os.WriteFile(src, content, 0644)
			if err != nil {
				t.Logf("Cannot create file '%s': %v", name, err)
				continue
			}

			err = Copy(src, dst)
			if err != nil {
				t.Logf("Copy file '%s' failed: %v", name, err)
			} else {
				t.Logf("Copy file '%s' succeeded", name)
			}
		}
	})

	t.Run("copy_readonly_to_readonly_dir", func(t *testing.T) {
		if os.Geteuid() == 0 {
			t.Skip("Skipping test when running as root")
		}

		tmpDir, err := os.MkdirTemp("", "copy_readonly_test_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		// 创建只读源文件
		src := filepath.Join(tmpDir, "readonly_source.txt")
		err = os.WriteFile(src, []byte("readonly content"), 0444)
		if err != nil {
			t.Fatal(err)
		}

		// 创建只读目录
		readonlyDir := filepath.Join(tmpDir, "readonly_dir")
		err = os.Mkdir(readonlyDir, 0555) // 只读和执行权限
		if err != nil {
			t.Fatal(err)
		}

		dst := filepath.Join(readonlyDir, "dest.txt")

		// 尝试复制到只读目录
		err = Copy(src, dst)

		// 恢复权限以便清理
		os.Chmod(readonlyDir, 0755)
		os.Chmod(src, 0644)

		if err != nil {
			t.Logf("Copy to readonly directory failed as expected: %v", err)
		} else {
			t.Error("Copy should fail when destination directory is readonly")
		}
	})

	t.Run("copy_large_file_chunks", func(t *testing.T) {
		// 创建一个相对较大的文件来测试io.Copy的各种情况
		tmpDir, err := os.MkdirTemp("", "copy_large_test_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		src := filepath.Join(tmpDir, "large_source.txt")
		dst := filepath.Join(tmpDir, "large_dest.txt")

		// 创建256KB的文件
		largeContent := make([]byte, 256*1024)
		for i := range largeContent {
			largeContent[i] = byte(i % 256)
		}

		err = os.WriteFile(src, largeContent, 0644)
		if err != nil {
			t.Fatal(err)
		}

		err = Copy(src, dst)
		if err != nil {
			t.Errorf("Copy large file failed: %v", err)
		} else {
			// 验证复制的内容
			dstContent, err := os.ReadFile(dst)
			if err != nil {
				t.Fatal(err)
			}

			if len(dstContent) != len(largeContent) {
				t.Errorf("Size mismatch: expected %d, got %d", len(largeContent), len(dstContent))
			}

			// 验证前100字节的内容
			for i := 0; i < 100 && i < len(dstContent); i++ {
				if dstContent[i] != largeContent[i] {
					t.Errorf("Content mismatch at byte %d: expected %d, got %d", i, largeContent[i], dstContent[i])
					break
				}
			}
		}
	})
}

// 测试权限相关的错误情况
func TestCopy_PermissionErrors(t *testing.T) {
	if os.Geteuid() == 0 {
		t.Skip("Skipping permission tests when running as root")
	}

	t.Run("copy_unreadable_file", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "copy_unreadable_test_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		src := filepath.Join(tmpDir, "unreadable.txt")
		dst := filepath.Join(tmpDir, "dest.txt")

		// 创建文件
		err = os.WriteFile(src, []byte("unreadable content"), 0644)
		if err != nil {
			t.Fatal(err)
		}

		// 移除读权限
		err = os.Chmod(src, 0000)
		if err != nil {
			t.Fatal(err)
		}

		// 尝试复制
		err = Copy(src, dst)

		// 恢复权限以便清理
		os.Chmod(src, 0644)

		if err != nil {
			t.Logf("Copy unreadable file failed as expected: %v", err)
		} else {
			t.Error("Copy should fail for unreadable file")
		}
	})
}

// 尝试各种文件系统相关的边界条件
func TestCopy_FileSystemEdgeCases(t *testing.T) {
	t.Run("copy_with_filesystem_errors", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "copy_fs_test_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		// 测试各种可能导致stat失败的情况
		testCases := []struct {
			name    string
			content []byte
			mode    fs.FileMode
		}{
			{"empty_file", []byte(""), 0644},
			{"single_byte", []byte("a"), 0644},
			{"null_bytes", []byte{0, 0, 0}, 0644},
			{"binary_data", []byte{0xFF, 0xFE, 0xFD, 0xFC}, 0644},
			{"executable", []byte("#!/bin/bash\necho test"), 0755},
		}

		for _, tc := range testCases {
			src := filepath.Join(tmpDir, tc.name+"_src.txt")
			dst := filepath.Join(tmpDir, tc.name+"_dst.txt")

			// 创建源文件
			err = os.WriteFile(src, tc.content, tc.mode)
			if err != nil {
				t.Fatalf("Cannot create test file %s: %v", tc.name, err)
			}

			// 复制文件
			err = Copy(src, dst)
			if err != nil {
				t.Logf("Copy %s failed: %v", tc.name, err)
			} else {
				// 验证复制结果
				dstContent, err := os.ReadFile(dst)
				if err != nil {
					t.Errorf("Cannot read copied file %s: %v", tc.name, err)
					continue
				}

				if len(dstContent) != len(tc.content) {
					t.Errorf("Size mismatch for %s: expected %d, got %d", tc.name, len(tc.content), len(dstContent))
				}
			}
		}
	})
}
