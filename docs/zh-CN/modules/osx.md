---
title: osx - 操作系统操作
---

# osx - 操作系统操作

## 概述

osx 模块为文件和目录操作提供操作系统工具，具有增强的功能。

## 函数

### Exists()

检查路径是否存在。

```go
func Exists(path string) bool
```

**参数：**
- `path` - 要检查的路径

**返回值：**
- 如果路径存在，返回 true
- 否则返回 false

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

检查路径是否为目录。

```go
func IsDir(path string) bool
```

**参数：**
- `path` - 要检查的路径

**返回值：**
- 如果路径是目录，返回 true
- 否则返回 false

**示例：**
```go
if osx.IsDir("/tmp") {
    log.Info("/tmp 是目录")
}
```

---

### IsFile()

检查路径是否为文件。

```go
func IsFile(path string) bool
```

**参数：**
- `path` - 要检查的路径

**返回值：**
- 如果路径是文件，返回 true
- 否则返回 false

**示例：**
```go
if osx.IsFile("config.json") {
    log.Info("config.json 是文件")
}
```

---

### Exist()

检查路径是否存在（Exists 的别名）。

```go
func Exist(path string) bool
```

**参数：**
- `path` - 要检查的路径

**返回值：**
- 如果路径存在，返回 true
- 否则返回 false

**示例：**
```go
if osx.Exist("data") {
    log.Info("data 存在")
}
```

---

### FsHasFile()

检查文件是否存在于文件系统中。

```go
func FsHasFile(fs fs.FS, path string) bool
```

**参数：**
- `fs` - 文件系统接口
- `path` - 要检查的路径

**返回值：**
- 如果文件存在，返回 true
- 否则返回 false

**示例：**
```go
if osx.FsHasFile(os.DirFS("/tmp"), "test.txt") {
    log.Info("test.txt 存在于 /tmp")
}
```

---

### RenameForce()

重命名文件，如果目标存在则删除。

```go
func RenameForce(oldpath, newpath string) error
```

**参数：**
- `oldpath` - 源路径
- `newpath` - 目标路径

**返回值：**
- 如果操作失败，返回错误

**示例：**
```go
if err := osx.RenameForce("old.txt", "new.txt"); err != nil {
    log.Errorf("重命名失败: %v", err)
}
```

---

### Copy()

从源复制文件到目标。

```go
func Copy(src, dst string) error
```

**参数：**
- `src` - 源文件路径
- `dst` - 目标文件路径

**返回值：**
- 如果操作失败，返回错误

**示例：**
```go
if err := osx.Copy("source.txt", "destination.txt"); err != nil {
    log.Errorf("复制失败: %v", err)
}
```

---

## 使用模式

### 文件存在性检查

```go
func checkFile(path string) bool {
    if !osx.Exists(path) {
        log.Warnf("文件未找到: %s", path)
        return false
    }
    
    if !osx.IsFile(path) {
        log.Warnf("路径不是文件: %s", path)
        return false
    }
    
    return true
}
```

### 目录操作

```go
func ensureDirectory(path string) error {
    if osx.Exists(path) {
        if !osx.IsDir(path) {
            return fmt.Errorf("路径存在但不是目录: %s", path)
        }
        return nil
    }
    
    return os.MkdirAll(path, 0755)
}
```

### 带备份的文件复制

```go
func copyWithBackup(src, dst string) error {
    if !osx.Exists(src) {
        return fmt.Errorf("源文件未找到: %s", src)
    }
    
    // 如果目标存在，创建备份
    if osx.Exists(dst) {
        backupPath := dst + ".bak"
        if err := osx.Copy(dst, backupPath); err != nil {
            return fmt.Errorf("创建备份失败: %w", err)
        }
    }
    
    // 复制文件
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
        log.Warnf("目标存在，将被覆盖: %s", newpath)
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
            return fmt.Errorf("复制 %s 失败: %w", file, err)
        }
    }
    
    return nil
}
```

---

## 最佳实践

### 错误处理

```go
// 好：操作前检查文件存在性
func readFile(path string) ([]byte, error) {
    if !osx.Exists(path) {
        return nil, fmt.Errorf("文件未找到: %s", path)
    }
    
    return os.ReadFile(path)
}

// 好：处理复制错误
func safeCopy(src, dst string) error {
    if err := osx.Copy(src, dst); err != nil {
        log.Errorf("复制失败: %v", err)
        return err
    }
    return nil
}
```

### 文件系统检查

```go
// 好：操作前验证文件类型
func processFile(path string) error {
    if !osx.Exists(path) {
        return fmt.Errorf("文件未找到: %s", path)
    }
    
    if osx.IsDir(path) {
        return fmt.Errorf("期望文件，得到目录: %s", path)
    }
    
    // 处理文件
    return nil
}
```

### 原子操作

```go
// 好：使用原子重命名进行文件更新
func atomicWrite(path string, data []byte) error {
    tmpPath := path + ".tmp"
    
    // 写入临时文件
    if err := os.WriteFile(tmpPath, data, 0644); err != nil {
        return err
    }
    
    // 原子重命名
    return osx.RenameForce(tmpPath, path)
}
```

---

## 相关文档

- [runtime](/zh-CN/modules/runtime) - 运行时信息
- [config](/zh-CN/modules/config) - 配置管理
- [API 文档](/zh-CN/api/overview)
- [模块概览](/zh-CN/modules/overview)
