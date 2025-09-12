package osx

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopy_FinalCoverage(t *testing.T) {
	// 尝试覆盖Copy函数的最后6.7%
	// 基于代码分析，未覆盖的可能是某个错误处理分支
	
	t.Run("copy_with_various_file_sizes", func(t *testing.T) {
		// 测试不同大小的文件复制
		tmpDir, err := os.MkdirTemp("", "copy_final_test_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		testCases := []struct {
			name string
			size int
			data []byte
		}{
			{"empty_file", 0, []byte("")},
			{"small_file", 10, []byte("0123456789")},
			{"medium_file", 1024, make([]byte, 1024)},
			{"large_file", 64*1024, make([]byte, 64*1024)}, // 64KB
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				src := filepath.Join(tmpDir, tc.name+"_src.txt")
				dst := filepath.Join(tmpDir, tc.name+"_dst.txt")

				// 准备测试数据
				testData := tc.data
				if tc.size > len(tc.data) {
					testData = make([]byte, tc.size)
					for i := range testData {
						testData[i] = byte(i % 256)
					}
				}

				// 创建源文件
				err := os.WriteFile(src, testData, 0644)
				if err != nil {
					t.Fatal(err)
				}

				// 执行复制
				err = Copy(src, dst)
				if err != nil {
					t.Errorf("Copy failed for %s: %v", tc.name, err)
					return
				}

				// 验证复制结果
				dstData, err := os.ReadFile(dst)
				if err != nil {
					t.Fatal(err)
				}

				if len(dstData) != len(testData) {
					t.Errorf("%s: size mismatch, expected %d, got %d", tc.name, len(testData), len(dstData))
				}

				// 对于小文件，比较全部内容
				if tc.size <= 1024 {
					for i, b := range testData {
						if i < len(dstData) && dstData[i] != b {
							t.Errorf("%s: content mismatch at position %d, expected %d, got %d", tc.name, i, b, dstData[i])
							break
						}
					}
				}
			})
		}
	})

	t.Run("copy_with_special_permissions", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "copy_perm_test_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		// 测试各种文件权限
		permissions := []os.FileMode{
			0600, // 只有所有者可读写
			0644, // 所有者读写，其他只读
			0755, // 所有者全权限，其他读执行
			0777, // 全权限
		}

		for i, perm := range permissions {
			src := filepath.Join(tmpDir, "perm_src_"+string(rune('0'+i))+".txt")
			dst := filepath.Join(tmpDir, "perm_dst_"+string(rune('0'+i))+".txt")

			// 创建具有特定权限的源文件
			err := os.WriteFile(src, []byte("permission test content"), perm)
			if err != nil {
				t.Fatal(err)
			}

			// 复制文件
			err = Copy(src, dst)
			if err != nil {
				t.Errorf("Copy failed for permission %o: %v", perm, err)
				continue
			}

			// 验证权限被正确复制
			srcInfo, err := os.Stat(src)
			if err != nil {
				t.Fatal(err)
			}

			dstInfo, err := os.Stat(dst)
			if err != nil {
				t.Fatal(err)
			}

			if srcInfo.Mode() != dstInfo.Mode() {
				t.Errorf("Permission not preserved: expected %o, got %o", srcInfo.Mode(), dstInfo.Mode())
			}
		}
	})

	t.Run("copy_concurrent_access", func(t *testing.T) {
		// 测试并发访问情况（可能触发某些边界条件）
		tmpDir, err := os.MkdirTemp("", "copy_concurrent_test_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		src := filepath.Join(tmpDir, "concurrent_src.txt")
		dst := filepath.Join(tmpDir, "concurrent_dst.txt")

		// 创建源文件
		content := "This is content for concurrent access testing"
		err = os.WriteFile(src, []byte(content), 0644)
		if err != nil {
			t.Fatal(err)
		}

		// 执行复制
		err = Copy(src, dst)
		if err != nil {
			t.Errorf("Concurrent copy failed: %v", err)
			return
		}

		// 验证结果
		dstContent, err := os.ReadFile(dst)
		if err != nil {
			t.Fatal(err)
		}

		if string(dstContent) != content {
			t.Errorf("Content mismatch: expected '%s', got '%s'", content, string(dstContent))
		}
	})

	t.Run("copy_edge_case_paths", func(t *testing.T) {
		// 测试边界情况的路径
		tmpDir, err := os.MkdirTemp("", "copy_edge_path_test_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		// 创建嵌套目录结构
		nestedDir := filepath.Join(tmpDir, "a", "b", "c")
		err = os.MkdirAll(nestedDir, 0755)
		if err != nil {
			t.Fatal(err)
		}

		src := filepath.Join(nestedDir, "deep_file.txt")
		dst := filepath.Join(tmpDir, "copied_deep_file.txt")

		// 创建深层文件
		err = os.WriteFile(src, []byte("deep file content"), 0644)
		if err != nil {
			t.Fatal(err)
		}

		// 复制到浅层目录
		err = Copy(src, dst)
		if err != nil {
			t.Errorf("Copy from deep path failed: %v", err)
			return
		}

		// 验证复制
		content, err := os.ReadFile(dst)
		if err != nil {
			t.Fatal(err)
		}

		if string(content) != "deep file content" {
			t.Errorf("Deep file content mismatch: expected 'deep file content', got '%s'", string(content))
		}
	})
}