---
title: osx - 操作系統操作
---

# osx - 操作系統操作

## 概述

osx 模組為文件和目錄操作提供操作系統工具，具有增強的功能。

## 函數

### Exists()

檢查路徑是否存在。

```go
func Exists(path string) bool
```

**參數：**
- `path` - 要檢查的路徑

**返回值：**
- 如果路徑存在，返回 true
- 否則返回 false

**示例：**
```go
if osx.Exists("config.json") {
    log.Info("配置文件存在")
} else {
    log.Warn("配置文件未找到")
}
```

---

### IsDir()

檢查路徑是否為目錄。

```go
func IsDir(path string) bool
```

**參數：**
- `path` - 要檢查的路徑

**返回值：**
- 如果路徑是目錄，返回 true
- 否則返回 false

**示例：**
```go
if osx.IsDir("/tmp") {
    log.Info("/tmp 是目錄")
}
```

---

### IsFile()

檢查路徑是否為文件。

```go
func IsFile(path string) bool
```

**參數：**
- `path` - 要檢查的路徑

**返回值：**
- 如果路徑是文件，返回 true
- 否則返回 false

**示例：**
```go
if osx.IsFile("config.json") {
    log.Info("config.json 是文件")
}
```

---

### Exist()

檢查路徑是否存在（Exists 的別名）。

```go
func Exist(path string) bool
```

**參數：**
- `path` - 要檢查的路徑

**返回值：**
- 如果路徑存在，返回 true
- 否則返回 false

**示例：**
```go
if osx.Exist("data") {
    log.Info("data 存在")
}
```

---

### FsHasFile()

檢查文件是否存在於文件系統中。

```go
func FsHasFile(fs fs.FS, path string) bool
```

**參數：**
- `fs` - 文件系統接口
- `path` - 要檢查的路徑

**返回值：**
- 如果文件存在，返回 true
- 否則返回 false

**示例：**
```go
if osx.FsHasFile(os.DirFS("/tmp"), "test.txt") {
    log.Info("test.txt 存在於 /tmp")
}
```

---

### RenameForce()

重命名文件，如果目標存在則刪除。

```go
func RenameForce(oldpath, newpath string) error
```

**參數：**
- `oldpath` - 源路徑
- `newpath` - 目標路徑

**返回值：**
- 如果操作失敗，返回錯誤

**示例：**
```go
if err := osx.RenameForce("old.txt", "new.txt"); err != nil {
    log.Errorf("重命名失敗: %v", err)
}
```

---

### Copy()

從源複製文件到目標。

```go
func Copy(src, dst string) error
```

**參數：**
- `src` - 源文件路徑
- `dst` - 目標文件路徑

**返回值：**
- 如果操作失敗，返回錯誤

**示例：**
```go
if err := osx.Copy("source.txt", "destination.txt"); err != nil {
    log.Errorf("複製失敗: %v", err)
}
```

---

## 使用模式

### 文件存在性檢查

```go
func checkFile(path string) bool {
    if !osx.Exists(path) {
        log.Warnf("文件未找到: %s", path)
        return false
    }
    
    if !osx.IsFile(path) {
        log.Warnf("路徑不是文件: %s", path)
        return false
    }
    
    return true
}
```

### 目錄操作

```go
func ensureDirectory(path string) error {
    if osx.Exists(path) {
        if !osx.IsDir(path) {
            return fmt.Errorf("路徑存在但不是目錄: %s", path)
        }
        return nil
    }
    
    return os.MkdirAll(path, 0755)
}
```

### 帶備份的文件複製

```go
func copyWithBackup(src, dst string) error {
    if !osx.Exists(src) {
        return fmt.Errorf("源文件未找到: %s", src)
    }
    
    // 如果目標存在，創建備份
    if osx.Exists(dst) {
        backupPath := dst + ".bak"
        if err := osx.Copy(dst, backupPath); err != nil {
            return fmt.Errorf("創建備份失敗: %w", err)
        }
    }
    
    // 複製文件
    return osx.Copy(src, dst)
}
```

### 安全重命名

```go
func safeRename(oldpath, newpath string) error {
    if !osx.Exists(oldpath) {
        return fmt.Errorf("源文件未找到: %s", oldpath)
    }
    
    if osx.Exists(newpath) {
        log.Warnf("目標存在，將被覆蓋: %s", newpath)
    }
    
    return osx.RenameForce(oldpath, newpath)
}
```

### 批量文件操作

```go
func copyFiles(srcDir, dstDir string, files []string) error {
    for _, file := range files {
        srcPath := filepath.Join(srcDir, file)
        dstPath := filepath.Join(dstDir, file)
        
        if err := osx.Copy(srcPath, dstPath); err != nil {
            return fmt.Errorf("複製 %s 失敗: %w", file, err)
        }
    }
    
    return nil
}
```

---

## 最佳實踐

### 錯誤處理

```go
// 好：操作前檢查文件存在性
func readFile(path string) ([]byte, error) {
    if !osx.Exists(path) {
        return nil, fmt.Errorf("文件未找到: %s", path)
    }
    
    return os.ReadFile(path)
}

// 好：處理複製錯誤
func safeCopy(src, dst string) error {
    if err := osx.Copy(src, dst); err != nil {
        log.Errorf("複製失敗: %v", err)
        return err
    }
    return nil
}
```

### 文件系統檢查

```go
// 好：操作前驗證文件類型
func processFile(path string) error {
    if !osx.Exists(path) {
        return fmt.Errorf("文件未找到: %s", path)
    }
    
    if osx.IsDir(path) {
        return fmt.Errorf("期望文件，得到目錄: %s", path)
    }
    
    // 處理文件
    return nil
}
```

### 原子操作

```go
// 好：使用原子重命名進行文件更新
func atomicWrite(path string, data []byte) error {
    tmpPath := path + ".tmp"
    
    // 寫入臨時文件
    if err := os.WriteFile(tmpPath, data, 0644); err != nil {
        return err
    }
    
    // 原子重命名
    return osx.RenameForce(tmpPath, path)
}
```

---

## 相關文檔

- [runtime](/zh-TW/modules/runtime) - 運行時資訊
- [config](/zh-TW/modules/config) - 配置管理
- [API 文檔](/zh-TW/api/overview)
- [模組概覽](/zh-TW/modules/overview)
