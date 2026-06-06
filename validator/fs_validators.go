package validator

import (
	"os"
	"path/filepath"
	"strings"
)

// FSValidators 返回所有文件系统验证器注册表
func FSValidators() map[string]ValidatorFunc {
	return map[string]ValidatorFunc{
		"dir":      validateDir,
		"dirpath":  validateDirPath,
		"file":     validateFile,
		"filepath": validateFilePath,
		"image":    validateImage,
	}
}

func validateDir(fl FieldLevel) bool {
	path := fl.Field().String()
	if path == "" {
		return false
	}
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func validateDirPath(fl FieldLevel) bool {
	path := fl.Field().String()
	if path == "" {
		return false
	}
	// 目录路径只需格式合法，不必存在
	cleaned := filepath.Clean(path)
	// 简单检查：不允许包含 null 字节
	return !strings.Contains(cleaned, "\x00")
}

func validateFile(fl FieldLevel) bool {
	path := fl.Field().String()
	if path == "" {
		return false
	}
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func validateFilePath(fl FieldLevel) bool {
	path := fl.Field().String()
	if path == "" {
		return false
	}
	dir := filepath.Dir(path)
	cleaned := filepath.Clean(dir)
	return !strings.Contains(cleaned, "\x00") && cleaned != ""
}

var imageExts = map[string]bool{
	".jpg": true, ".jpeg": true, ".png": true, ".gif": true,
	".bmp": true, ".svg": true, ".webp": true, ".tiff": true,
	".tif": true, ".ico": true, ".avif": true,
}

func validateImage(fl FieldLevel) bool {
	path := fl.Field().String()
	if path == "" {
		return false
	}
	ext := strings.ToLower(filepath.Ext(path))
	return imageExts[ext]
}
