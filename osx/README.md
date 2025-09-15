# OSX - OS-specific File and Process Operations

The `osx` module provides cross-platform utilities for file system operations and OS-specific functionality. Despite its name, this module works on all operating systems and provides essential file system utilities for path checking, file copying, and force renaming operations.

## Features

- **File Existence Checking**: Multiple ways to check if files/directories exist
- **Type Detection**: Distinguish between files and directories
- **Force Operations**: Rename files with automatic cleanup of destination conflicts
- **File System Integration**: Works with both regular file systems and `fs.FS` interfaces
- **Atomic Operations**: Safe file operations with proper error handling
- **Cross-Platform**: Works consistently across different operating systems

## Installation

```bash
go get github.com/lazygophers/utils
```

## Usage

### Basic File System Checks

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/osx"
)

func main() {
    // Check if file or directory exists
    if osx.Exists("./config.json") {
        fmt.Println("Config file exists")
    }

    // Alternative existence check (stricter)
    if osx.Exist("./config.json") {
        fmt.Println("Config file definitely exists")
    }

    // Check if path is a directory
    if osx.IsDir("./logs") {
        fmt.Println("Logs directory found")
    }

    // Check if path is a file
    if osx.IsFile("./app.log") {
        fmt.Println("App log file found")
    }
}
```

### File System Interface Operations

```go
package main

import (
    "embed"
    "fmt"
    "github.com/lazygophers/utils/osx"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
    // Check if file exists in embedded file system
    if osx.FsHasFile(staticFiles, "static/index.html") {
        fmt.Println("Index file found in embedded FS")
    }

    // Works with any fs.FS implementation
    if osx.FsHasFile(os.DirFS("./templates"), "header.html") {
        fmt.Println("Header template found")
    }
}
```

### File Operations

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/osx"
    "log"
)

func main() {
    // Copy file with permission preservation
    err := osx.Copy("source.txt", "destination.txt")
    if err != nil {
        log.Fatal("Copy failed:", err)
    }

    // Force rename - removes destination if it exists
    err = osx.RenameForce("old_name.txt", "new_name.txt")
    if err != nil {
        log.Fatal("Rename failed:", err)
    }

    fmt.Println("File operations completed successfully")
}
```

### Safe File Processing Pattern

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/osx"
    "log"
    "os"
)

func processFileIfExists(filename string) error {
    // Use Exist() for strict existence checking
    if !osx.Exist(filename) {
        return fmt.Errorf("file %s not found", filename)
    }

    if !osx.IsFile(filename) {
        return fmt.Errorf("%s is not a regular file", filename)
    }

    // File exists and is a regular file - safe to process
    data, err := os.ReadFile(filename)
    if err != nil {
        return err
    }

    fmt.Printf("Processed %d bytes from %s\n", len(data), filename)
    return nil
}

func main() {
    if err := processFileIfExists("config.json"); err != nil {
        log.Printf("Processing failed: %v", err)
    }
}
```

### Backup and Replace Pattern

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/osx"
    "log"
    "os"
    "time"
)

func safeFileUpdate(original, newContent string) error {
    backup := original + ".backup." + time.Now().Format("20060102150405")

    // Create backup if original exists
    if osx.Exist(original) {
        if err := osx.Copy(original, backup); err != nil {
            return fmt.Errorf("backup failed: %w", err)
        }
        defer func() {
            // Clean up backup on success
            os.Remove(backup)
        }()
    }

    // Write new content to temporary file
    temp := original + ".tmp"
    if err := os.WriteFile(temp, []byte(newContent), 0644); err != nil {
        return fmt.Errorf("write temp failed: %w", err)
    }

    // Atomic rename
    if err := osx.RenameForce(temp, original); err != nil {
        // Restore backup on failure
        if osx.Exist(backup) {
            osx.RenameForce(backup, original)
        }
        return fmt.Errorf("rename failed: %w", err)
    }

    fmt.Println("File updated successfully")
    return nil
}

func main() {
    newData := `{"version": "2.0", "updated": "` + time.Now().Format(time.RFC3339) + `"}`
    if err := safeFileUpdate("config.json", newData); err != nil {
        log.Fatal(err)
    }
}
```

## API Reference

### File Existence Functions

#### `Exists(path string) bool`
Checks if a file or directory exists at the given path. Uses `os.IsExist()` for error checking, which may return true for some error conditions.

**Parameters:**
- `path`: File or directory path to check

**Returns:**
- `bool`: True if the path exists or if `os.IsExist()` returns true for the error

#### `Exist(path string) bool`
Stricter existence check that only returns true if the path definitely exists (no error from `os.Stat()`).

**Parameters:**
- `path`: File or directory path to check

**Returns:**
- `bool`: True only if the path definitely exists

### Type Detection Functions

#### `IsDir(path string) bool`
Checks if the given path exists and is a directory.

**Parameters:**
- `path`: Path to check

**Returns:**
- `bool`: True if path exists and is a directory

#### `IsFile(path string) bool`
Checks if the given path exists and is a regular file (not a directory).

**Parameters:**
- `path`: Path to check

**Returns:**
- `bool`: True if path exists and is a regular file

### File System Interface Functions

#### `FsHasFile(fs fs.FS, path string) bool`
Checks if a file exists within a file system interface (such as embed.FS, os.DirFS, etc.).

**Parameters:**
- `fs`: Any type implementing `fs.FS` interface
- `path`: File path within the file system

**Returns:**
- `bool`: True if the file exists in the file system

### File Operations

#### `Copy(src, dst string) error`
Copies a file from source to destination, preserving file permissions.

**Parameters:**
- `src`: Source file path
- `dst`: Destination file path

**Returns:**
- `error`: Error if copy operation fails

**Features:**
- Preserves original file permissions
- Uses efficient `io.Copy` for data transfer
- Properly closes file handles

#### `RenameForce(oldpath, newpath string) error`
Renames a file or directory, forcibly removing the destination if it already exists.

**Parameters:**
- `oldpath`: Current path of the file/directory
- `newpath`: New path for the file/directory

**Returns:**
- `error`: Error if the operation fails

**Features:**
- Automatically removes destination if it exists
- Atomic operation where possible
- Works with both files and directories

## Best Practices

### 1. Choose the Right Existence Check
```go
// Use Exist() for strict checking
if osx.Exist(filename) {
    // File definitely exists
}

// Use Exists() if you need to handle edge cases
if osx.Exists(filename) {
    // File might exist or there might be permission issues
}
```

### 2. Always Check File Type
```go
if osx.Exist(path) {
    if osx.IsDir(path) {
        // Handle directory
    } else if osx.IsFile(path) {
        // Handle regular file
    } else {
        // Handle special files (symlinks, devices, etc.)
    }
}
```

### 3. Handle Errors Properly
```go
if err := osx.Copy(src, dst); err != nil {
    log.Printf("Copy failed: %v", err)
    // Handle error appropriately
}
```

### 4. Use Atomic Operations
```go
// Don't do this - not atomic
os.Remove(target)
os.Rename(source, target)

// Do this - atomic when possible
osx.RenameForce(source, target)
```

## Performance Considerations

- **File System Calls**: All functions make system calls; cache results if needed
- **Error Handling**: `Exist()` is faster than `Exists()` as it does less error checking
- **Copy Operations**: Uses `io.Copy` for efficient large file copying
- **Batch Operations**: Group multiple file checks when possible

## Thread Safety

All functions in the OSX module are thread-safe as they don't maintain internal state. However, be aware that:

- File system state can change between calls
- Use appropriate locking for concurrent file operations
- `RenameForce` operations should be synchronized if operating on the same paths

## Error Handling

The module provides clear error messages and uses Go's standard error handling patterns:

```go
err := osx.Copy("source.txt", "dest.txt")
if err != nil {
    switch {
    case os.IsNotExist(err):
        fmt.Println("Source file not found")
    case os.IsPermission(err):
        fmt.Println("Permission denied")
    default:
        fmt.Printf("Copy failed: %v\n", err)
    }
}
```

## Common Use Cases

1. **Configuration File Management**: Check and update config files safely
2. **Log File Operations**: Manage log rotation and cleanup
3. **Temporary File Handling**: Create and manage temporary files
4. **Build Systems**: File copying and organization in build pipelines
5. **Embedded File Systems**: Work with embedded assets and templates
6. **Data Processing**: Safe file processing with existence checking

## Related Packages

- [`os`](https://pkg.go.dev/os): Standard Go OS interface
- [`io/fs`](https://pkg.go.dev/io/fs): File system interface definitions
- [`path/filepath`](https://pkg.go.dev/path/filepath): File path manipulation
- [`embed`](https://pkg.go.dev/embed): Embedded file systems