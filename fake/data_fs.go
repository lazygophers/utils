package fake

import (
	"embed"
	"io/fs"
	"path/filepath"
)

// MultiFS 组合多个文件系统
type MultiFS struct {
	filesystems map[string]embed.FS
}

func (m *MultiFS) Open(name string) (fs.File, error) {
	// 尝试从所有文件系统中打开文件
	for _, fsys := range m.filesystems {
		if file, err := fsys.Open(name); err == nil {
			return file, nil
		}
	}
	return nil, fs.ErrNotExist
}

func (m *MultiFS) ReadDir(name string) ([]fs.DirEntry, error) {
	var allEntries []fs.DirEntry
	seen := make(map[string]bool)

	for _, fsys := range m.filesystems {
		if entries, err := fsys.ReadDir(name); err == nil {
			for _, entry := range entries {
				if !seen[entry.Name()] {
					allEntries = append(allEntries, entry)
					seen[entry.Name()] = true
				}
			}
		}
	}

	if len(allEntries) == 0 {
		return nil, fs.ErrNotExist
	}
	return allEntries, nil
}

func (m *MultiFS) ReadFile(name string) ([]byte, error) {
	for _, fsys := range m.filesystems {
		if data, err := fsys.ReadFile(name); err == nil {
			return data, nil
		}
	}
	return nil, fs.ErrNotExist
}

// GetAvailableLanguages 获取可用的语言列表
func (m *MultiFS) GetAvailableLanguages() []string {
	var langs []string
	entries, err := fs.ReadDir(m, "data")
	if err != nil {
		return langs
	}

	for _, entry := range entries {
		if entry.IsDir() {
			langs = append(langs, entry.Name())
		}
	}
	return langs
}

// HasLanguage 检查是否包含指定语言
func (m *MultiFS) HasLanguage(lang string) bool {
	_, err := fs.ReadDir(m, filepath.Join("data", lang))
	return err == nil
}

var dataFS = &MultiFS{
	filesystems: make(map[string]embed.FS),
}

// DataFS 全局文件系统实例
var DataFS fs.FS = dataFS

// RegisterLanguageFS 注册语言文件系统
func RegisterLanguageFS(lang string, fsys embed.FS) {
	dataFS.filesystems[lang] = fsys
}
