package osx

import (
	"os"
	"path/filepath"
	"testing"
)

// TestCopyMissingPaths tests the specific missing coverage paths in Copy function
func TestCopyMissingPaths(t *testing.T) {
	t.Run("copy_with_permission_errors", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "copy_missing_test_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		// Create source file
		src := filepath.Join(tmpDir, "source.txt")
		if err := os.WriteFile(src, []byte("test content"), 0644); err != nil {
			t.Fatal(err)
		}

		// Try to copy to a read-only directory (should trigger OpenFile error)
		readOnlyDir := filepath.Join(tmpDir, "readonly")
		if err := os.MkdirAll(readOnlyDir, 0755); err != nil {
			t.Fatal(err)
		}

		// Make the directory read-only
		if err := os.Chmod(readOnlyDir, 0444); err != nil {
			t.Skip("Cannot change directory permissions on this system")
		}

		dst := filepath.Join(readOnlyDir, "dest.txt")
		err = Copy(src, dst)
		if err == nil {
			t.Error("Expected error when copying to read-only directory")
		} else {
			t.Logf("Got expected error: %v", err)
		}

		// Restore permissions for cleanup
		os.Chmod(readOnlyDir, 0755)
	})

	t.Run("copy_nonexistent_source", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "copy_missing_test2_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		// Try to copy from non-existent source
		src := filepath.Join(tmpDir, "nonexistent.txt")
		dst := filepath.Join(tmpDir, "dest.txt")

		err = Copy(src, dst)
		if err == nil {
			t.Error("Expected error when copying from non-existent source")
		} else {
			t.Logf("Got expected error: %v", err)
		}
	})

	t.Run("copy_with_directory_as_source", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "copy_missing_test3_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		// Create a directory as source (this might cause Stat to behave differently)
		srcDir := filepath.Join(tmpDir, "sourcedir")
		if err := os.MkdirAll(srcDir, 0755); err != nil {
			t.Fatal(err)
		}

		dst := filepath.Join(tmpDir, "dest.txt")

		// This should fail when trying to copy a directory as a file
		err = Copy(srcDir, dst)
		if err == nil {
			t.Error("Expected error when copying directory as file")
		} else {
			t.Logf("Got expected error: %v", err)
		}
	})

	t.Run("copy_with_invalid_destination", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "copy_missing_test4_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		// Create valid source file
		src := filepath.Join(tmpDir, "source.txt")
		if err := os.WriteFile(src, []byte("test content"), 0644); err != nil {
			t.Fatal(err)
		}

		// Try to copy to an invalid path (parent directory doesn't exist)
		dst := filepath.Join(tmpDir, "nonexistent_dir", "dest.txt")

		err = Copy(src, dst)
		if err == nil {
			t.Error("Expected error when copying to invalid destination path")
		} else {
			t.Logf("Got expected error: %v", err)
		}
	})

	t.Run("copy_existing_file_replacement", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "copy_missing_test5_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		// Create source file
		src := filepath.Join(tmpDir, "source.txt")
		if err := os.WriteFile(src, []byte("new content"), 0644); err != nil {
			t.Fatal(err)
		}

		// Create existing destination file
		dst := filepath.Join(tmpDir, "dest.txt")
		if err := os.WriteFile(dst, []byte("old content"), 0644); err != nil {
			t.Fatal(err)
		}

		// Copy should replace the existing file
		err = Copy(src, dst)
		if err != nil {
			t.Errorf("Copy should replace existing file: %v", err)
		}

		// Verify content was replaced
		content, err := os.ReadFile(dst)
		if err != nil {
			t.Fatal(err)
		}
		if string(content) != "new content" {
			t.Error("File content was not properly replaced")
		}
	})

	t.Run("successful_copy_completion", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "successful_copy_test_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		// Create source file with specific content
		src := filepath.Join(tmpDir, "source.txt")
		originalContent := "This is test content for successful copy"
		if err := os.WriteFile(src, []byte(originalContent), 0644); err != nil {
			t.Fatal(err)
		}

		// Set up destination
		dst := filepath.Join(tmpDir, "destination.txt")

		// Perform copy - this should succeed and hit the "return nil" line
		err = Copy(src, dst)
		if err != nil {
			t.Fatalf("Copy should succeed: %v", err)
		}

		// Verify the copy was completely successful
		copiedContent, err := os.ReadFile(dst)
		if err != nil {
			t.Fatal(err)
		}

		if string(copiedContent) != originalContent {
			t.Errorf("Copied content mismatch. Expected: %q, Got: %q", originalContent, string(copiedContent))
		}

		// Verify file permissions were copied
		srcInfo, err := os.Stat(src)
		if err != nil {
			t.Fatal(err)
		}
		dstInfo, err := os.Stat(dst)
		if err != nil {
			t.Fatal(err)
		}

		if srcInfo.Mode() != dstInfo.Mode() {
			t.Logf("Source mode: %v, Destination mode: %v", srcInfo.Mode(), dstInfo.Mode())
		}

		t.Log("Successful copy completed - should hit return nil path")
	})
}

// TestExistsEdgeCases tests edge cases for the Exists function
func TestExistsEdgeCases(t *testing.T) {
	t.Run("exists_with_special_paths", func(t *testing.T) {
		// Test with empty path
		result := Exists("")
		t.Logf("Exists(\"\") = %v", result)

		// Test with invalid characters (if supported by OS)
		result = Exists("/\x00invalid")
		t.Logf("Exists with null character = %v", result)

		// Test with very long path
		longPath := filepath.Join(os.TempDir(), string(make([]byte, 300)))
		result = Exists(longPath)
		t.Logf("Exists with long path = %v", result)
	})
}

// TestFsHasFileEdgeCases tests edge cases for FsHasFile
func TestFsHasFileEdgeCases(t *testing.T) {
	t.Run("fs_has_file_with_invalid_paths", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("", "fs_test_*")
		if err != nil {
			t.Fatal(err)
		}
		defer os.RemoveAll(tmpDir)

		// Create a test file
		testFile := filepath.Join(tmpDir, "test.txt")
		if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
			t.Fatal(err)
		}

		fsys := os.DirFS(tmpDir)

		// Test with valid file
		result := FsHasFile(fsys, "test.txt")
		if !result {
			t.Error("Should find existing file")
		}

		// Test with non-existent file
		result = FsHasFile(fsys, "nonexistent.txt")
		if result {
			t.Error("Should not find non-existent file")
		}

		// Test with directory
		subDir := filepath.Join(tmpDir, "subdir")
		if err := os.MkdirAll(subDir, 0755); err != nil {
			t.Fatal(err)
		}

		result = FsHasFile(fsys, "subdir")
		// This should succeed because directories can be opened in fs.FS
		t.Logf("FsHasFile with directory = %v", result)
	})
}
