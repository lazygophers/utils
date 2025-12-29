---
title: osx - OS Operations
---

# osx - OS Operations

## Overview

The osx module provides operating system utilities for file and directory operations with enhanced functionality.

## Functions

### Exists()

Check if path exists.

```go
func Exists(path string) bool
```

**Parameters:**
- `path` - Path to check

**Returns:**
- true if path exists
- false otherwise

**Example:**
```go
if osx.Exists("config.json") {
    log.Info("Config file exists")
} else {
    log.Warn("Config file not found")
}
```

---

### IsDir()

Check if path is a directory.

```go
func IsDir(path string) bool
```

**Parameters:**
- `path` - Path to check

**Returns:**
- true if path is a directory
- false otherwise

**Example:**
```go
if osx.IsDir("/tmp") {
    log.Info("/tmp is a directory")
}
```

---

### IsFile()

Check if path is a file.

```go
func IsFile(path string) bool
```

**Parameters:**
- `path` - Path to check

**Returns:**
- true if path is a file
- false otherwise

**Example:**
```go
if osx.IsFile("config.json") {
    log.Info("config.json is a file")
}
```

---

### Exist()

Check if path exists (alias for Exists).

```go
func Exist(path string) bool
```

**Parameters:**
- `path` - Path to check

**Returns:**
- true if path exists
- false otherwise

**Example:**
```go
if osx.Exist("data") {
    log.Info("data exists")
}
```

---

### FsHasFile()

Check if file exists in filesystem.

```go
func FsHasFile(fs fs.FS, path string) bool
```

**Parameters:**
- `fs` - Filesystem interface
- `path` - Path to check

**Returns:**
- true if file exists
- false otherwise

**Example:**
```go
if osx.FsHasFile(os.DirFS("/tmp"), "test.txt") {
    log.Info("test.txt exists in /tmp")
}
```

---

### RenameForce()

Rename file, removing destination if exists.

```go
func RenameForce(oldpath, newpath string) error
```

**Parameters:**
- `oldpath` - Source path
- `newpath` - Destination path

**Returns:**
- Error if operation fails

**Example:**
```go
if err := osx.RenameForce("old.txt", "new.txt"); err != nil {
    log.Errorf("Failed to rename: %v", err)
}
```

---

### Copy()

Copy file from source to destination.

```go
func Copy(src, dst string) error
```

**Parameters:**
- `src` - Source file path
- `dst` - Destination file path

**Returns:**
- Error if operation fails

**Example:**
```go
if err := osx.Copy("source.txt", "destination.txt"); err != nil {
    log.Errorf("Failed to copy: %v", err)
}
```

---

## Usage Patterns

### File Existence Check

```go
func checkFile(path string) bool {
    if !osx.Exists(path) {
        log.Warnf("File not found: %s", path)
        return false
    }
    
    if !osx.IsFile(path) {
        log.Warnf("Path is not a file: %s", path)
        return false
    }
    
    return true
}
```

### Directory Operations

```go
func ensureDirectory(path string) error {
    if osx.Exists(path) {
        if !osx.IsDir(path) {
            return fmt.Errorf("path exists but is not a directory: %s", path)
        }
        return nil
    }
    
    return os.MkdirAll(path, 0755)
}
```

### File Copy with Backup

```go
func copyWithBackup(src, dst string) error {
    if !osx.Exists(src) {
        return fmt.Errorf("source file not found: %s", src)
    }
    
    // Create backup if destination exists
    if osx.Exists(dst) {
        backupPath := dst + ".bak"
        if err := osx.Copy(dst, backupPath); err != nil {
            return fmt.Errorf("failed to create backup: %w", err)
        }
    }
    
    // Copy file
    return osx.Copy(src, dst)
}
```

### Safe Rename

```go
func safeRename(oldpath, newpath string) error {
    if !osx.Exists(oldpath) {
        return fmt.Errorf("source file not found: %s", oldpath)
    }
    
    if osx.Exists(newpath) {
        log.Warnf("Destination exists, will be overwritten: %s", newpath)
    }
    
    return osx.RenameForce(oldpath, newpath)
}
```

### Batch File Operations

```go
func copyFiles(srcDir, dstDir string, files []string) error {
    for _, file := range files {
        srcPath := filepath.Join(srcDir, file)
        dstPath := filepath.Join(dstDir, file)
        
        if err := osx.Copy(srcPath, dstPath); err != nil {
            return fmt.Errorf("failed to copy %s: %w", file, err)
        }
    }
    
    return nil
}
```

---

## Best Practices

### Error Handling

```go
// Good: Check file existence before operations
func readFile(path string) ([]byte, error) {
    if !osx.Exists(path) {
        return nil, fmt.Errorf("file not found: %s", path)
    }
    
    return os.ReadFile(path)
}

// Good: Handle copy errors
func safeCopy(src, dst string) error {
    if err := osx.Copy(src, dst); err != nil {
        log.Errorf("Copy failed: %v", err)
        return err
    }
    return nil
}
```

### File System Checks

```go
// Good: Verify file type before operations
func processFile(path string) error {
    if !osx.Exists(path) {
        return fmt.Errorf("file not found: %s", path)
    }
    
    if osx.IsDir(path) {
        return fmt.Errorf("expected file, got directory: %s", path)
    }
    
    // Process file
    return nil
}
```

### Atomic Operations

```go
// Good: Use atomic rename for file updates
func atomicWrite(path string, data []byte) error {
    tmpPath := path + ".tmp"
    
    // Write to temporary file
    if err := os.WriteFile(tmpPath, data, 0644); err != nil {
        return err
    }
    
    // Atomic rename
    return osx.RenameForce(tmpPath, path)
}
```

---

## Related Documentation

- [runtime](/en/modules/runtime) - Runtime information
- [config](/en/modules/config) - Configuration management
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
