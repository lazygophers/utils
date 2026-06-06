package validator

import (
	"os"
	"reflect"
	"testing"
)

func TestValidateDir(t *testing.T) {
	// Create temp dir
	dir, _ := os.MkdirTemp("", "validator_test")
	defer os.RemoveAll(dir)

	fl := paramFL{field: reflect.ValueOf(dir), param: ""}
	if !validateDir(fl) {
		t.Error("existing dir should pass")
	}
	fl2 := paramFL{field: reflect.ValueOf("/nonexistent/path/xyz"), param: ""}
	if validateDir(fl2) {
		t.Error("nonexistent path should fail")
	}
	fl3 := paramFL{field: reflect.ValueOf(""), param: ""}
	if validateDir(fl3) {
		t.Error("empty string should fail")
	}
}

func TestValidateDirPath(t *testing.T) {
	fl := paramFL{field: reflect.ValueOf("/tmp/some/dir"), param: ""}
	if !validateDirPath(fl) {
		t.Error("valid dir path should pass")
	}
	fl2 := paramFL{field: reflect.ValueOf(""), param: ""}
	if validateDirPath(fl2) {
		t.Error("empty string should fail")
	}
	fl3 := paramFL{field: reflect.ValueOf("/tmp\x00bad"), param: ""}
	if validateDirPath(fl3) {
		t.Error("null byte path should fail")
	}
}

func TestValidateFile(t *testing.T) {
	f, _ := os.CreateTemp("", "validator_test_file")
	f.Close()
	defer os.Remove(f.Name())

	fl := paramFL{field: reflect.ValueOf(f.Name()), param: ""}
	if !validateFile(fl) {
		t.Error("existing file should pass")
	}
	fl2 := paramFL{field: reflect.ValueOf("/nonexistent/file.txt"), param: ""}
	if validateFile(fl2) {
		t.Error("nonexistent file should fail")
	}
	fl3 := paramFL{field: reflect.ValueOf(""), param: ""}
	if validateFile(fl3) {
		t.Error("empty string should fail")
	}
}

func TestValidateFilePath(t *testing.T) {
	fl := paramFL{field: reflect.ValueOf("/tmp/some/file.txt"), param: ""}
	if !validateFilePath(fl) {
		t.Error("valid file path should pass")
	}
	fl2 := paramFL{field: reflect.ValueOf(""), param: ""}
	if validateFilePath(fl2) {
		t.Error("empty string should fail")
	}
}

func TestValidateImage(t *testing.T) {
	for _, ext := range []string{".jpg", ".png", ".gif", ".svg", ".webp", ".avif"} {
		fl := paramFL{field: reflect.ValueOf("test" + ext), param: ""}
		if !validateImage(fl) {
			t.Errorf("image extension %s should pass", ext)
		}
	}
	fl := paramFL{field: reflect.ValueOf("test.txt"), param: ""}
	if validateImage(fl) {
		t.Error(".txt should fail image check")
	}
	fl2 := paramFL{field: reflect.ValueOf(""), param: ""}
	if validateImage(fl2) {
		t.Error("empty string should fail")
	}
}
