# OSX - 操作系统特定的文件和进程操作

`osx` 模块提供跨平台的文件系统操作和操作系统特定功能的实用工具。尽管名称为 osx，但该模块在所有操作系统上都能正常工作，并提供路径检查、文件复制和强制重命名操作等基本文件系统实用工具。

## 功能特性

- **文件存在性检查**: 多种方式检查文件/目录是否存在
- **类型检测**: 区分文件和目录
- **强制操作**: 重命名文件时自动清理目标冲突
- **文件系统集成**: 兼容常规文件系统和 `fs.FS` 接口
- **原子操作**: 安全的文件操作，具有适当的错误处理
- **跨平台**: 在不同操作系统间工作一致

## 安装

```bash
go get github.com/lazygophers/utils
```

## 使用方法

### 基本文件系统检查

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/osx"
)

func main() {
    // 检查文件或目录是否存在
    if osx.Exists("./config.json") {
        fmt.Println("配置文件存在")
    }

    // 替代的存在性检查（更严格）
    if osx.Exist("./config.json") {
        fmt.Println("配置文件确实存在")
    }

    // 检查路径是否为目录
    if osx.IsDir("./logs") {
        fmt.Println("找到日志目录")
    }

    // 检查路径是否为文件
    if osx.IsFile("./app.log") {
        fmt.Println("找到应用日志文件")
    }
}
```

### 文件系统接口操作

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
    // 检查嵌入式文件系统中是否存在文件
    if osx.FsHasFile(staticFiles, "static/index.html") {
        fmt.Println("在嵌入式文件系统中找到索引文件")
    }

    // 兼容任何 fs.FS 实现
    if osx.FsHasFile(os.DirFS("./templates"), "header.html") {
        fmt.Println("找到头部模板")
    }
}
```

### 文件操作

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/osx"
    "log"
)

func main() {
    // 复制文件并保持权限
    err := osx.Copy("source.txt", "destination.txt")
    if err != nil {
        log.Fatal("复制失败:", err)
    }

    // 强制重命名 - 如果目标存在则移除
    err = osx.RenameForce("old_name.txt", "new_name.txt")
    if err != nil {
        log.Fatal("重命名失败:", err)
    }

    fmt.Println("文件操作成功完成")
}
```

### 安全文件处理模式

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/osx"
    "log"
    "os"
)

func processFileIfExists(filename string) error {
    // 使用 Exist() 进行严格的存在性检查
    if !osx.Exist(filename) {
        return fmt.Errorf("文件 %s 不存在", filename)
    }

    if !osx.IsFile(filename) {
        return fmt.Errorf("%s 不是常规文件", filename)
    }

    // 文件存在且是常规文件 - 安全处理
    data, err := os.ReadFile(filename)
    if err != nil {
        return err
    }

    fmt.Printf("从 %s 处理了 %d 字节\n", filename, len(data))
    return nil
}

func main() {
    if err := processFileIfExists("config.json"); err != nil {
        log.Printf("处理失败: %v", err)
    }
}
```

### 备份和替换模式

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

    // 如果原文件存在则创建备份
    if osx.Exist(original) {
        if err := osx.Copy(original, backup); err != nil {
            return fmt.Errorf("备份失败: %w", err)
        }
        defer func() {
            // 成功时清理备份
            os.Remove(backup)
        }()
    }

    // 将新内容写入临时文件
    temp := original + ".tmp"
    if err := os.WriteFile(temp, []byte(newContent), 0644); err != nil {
        return fmt.Errorf("写入临时文件失败: %w", err)
    }

    // 原子重命名
    if err := osx.RenameForce(temp, original); err != nil {
        // 失败时恢复备份
        if osx.Exist(backup) {
            osx.RenameForce(backup, original)
        }
        return fmt.Errorf("重命名失败: %w", err)
    }

    fmt.Println("文件更新成功")
    return nil
}

func main() {
    newData := `{"version": "2.0", "updated": "` + time.Now().Format(time.RFC3339) + `"}`
    if err := safeFileUpdate("config.json", newData); err != nil {
        log.Fatal(err)
    }
}
```

## API 参考

### 文件存在性函数

#### `Exists(path string) bool`
检查给定路径是否存在文件或目录。使用 `os.IsExist()` 进行错误检查，在某些错误条件下可能返回 true。

**参数:**
- `path`: 要检查的文件或目录路径

**返回值:**
- `bool`: 如果路径存在或 `os.IsExist()` 对错误返回 true，则为 true

#### `Exist(path string) bool`
更严格的存在性检查，只有当路径确实存在时才返回 true（`os.Stat()` 无错误）。

**参数:**
- `path`: 要检查的文件或目录路径

**返回值:**
- `bool`: 只有当路径确实存在时才为 true

### 类型检测函数

#### `IsDir(path string) bool`
检查给定路径是否存在且为目录。

**参数:**
- `path`: 要检查的路径

**返回值:**
- `bool`: 如果路径存在且为目录则为 true

#### `IsFile(path string) bool`
检查给定路径是否存在且为常规文件（不是目录）。

**参数:**
- `path`: 要检查的路径

**返回值:**
- `bool`: 如果路径存在且为常规文件则为 true

### 文件系统接口函数

#### `FsHasFile(fs fs.FS, path string) bool`
检查文件系统接口中是否存在文件（如 embed.FS、os.DirFS 等）。

**参数:**
- `fs`: 实现 `fs.FS` 接口的任何类型
- `path`: 文件系统内的文件路径

**返回值:**
- `bool`: 如果文件在文件系统中存在则为 true

### 文件操作

#### `Copy(src, dst string) error`
从源复制文件到目标，保持文件权限。

**参数:**
- `src`: 源文件路径
- `dst`: 目标文件路径

**返回值:**
- `error`: 如果复制操作失败则返回错误

**特性:**
- 保持原文件权限
- 使用高效的 `io.Copy` 进行数据传输
- 正确关闭文件句柄

#### `RenameForce(oldpath, newpath string) error`
重命名文件或目录，如果目标已存在则强制移除。

**参数:**
- `oldpath`: 文件/目录的当前路径
- `newpath`: 文件/目录的新路径

**返回值:**
- `error`: 如果操作失败则返回错误

**特性:**
- 如果目标存在则自动移除
- 尽可能进行原子操作
- 支持文件和目录

## 最佳实践

### 1. 选择正确的存在性检查
```go
// 使用 Exist() 进行严格检查
if osx.Exist(filename) {
    // 文件确实存在
}

// 如果需要处理边界情况，使用 Exists()
if osx.Exists(filename) {
    // 文件可能存在或可能有权限问题
}
```

### 2. 始终检查文件类型
```go
if osx.Exist(path) {
    if osx.IsDir(path) {
        // 处理目录
    } else if osx.IsFile(path) {
        // 处理常规文件
    } else {
        // 处理特殊文件（符号链接、设备等）
    }
}
```

### 3. 正确处理错误
```go
if err := osx.Copy(src, dst); err != nil {
    log.Printf("复制失败: %v", err)
    // 适当处理错误
}
```

### 4. 使用原子操作
```go
// 不要这样做 - 不是原子的
os.Remove(target)
os.Rename(source, target)

// 这样做 - 尽可能原子
osx.RenameForce(source, target)
```

## 性能注意事项

- **文件系统调用**: 所有函数都进行系统调用；如需要可缓存结果
- **错误处理**: `Exist()` 比 `Exists()` 更快，因为错误检查更少
- **复制操作**: 使用 `io.Copy` 进行高效的大文件复制
- **批量操作**: 尽可能组合多个文件检查

## 线程安全

OSX 模块中的所有函数都是线程安全的，因为它们不维护内部状态。但是要注意：

- 文件系统状态可能在调用之间发生变化
- 对并发文件操作使用适当的锁定
- 如果操作相同路径，应同步 `RenameForce` 操作

## 错误处理

该模块提供清晰的错误消息并使用 Go 的标准错误处理模式：

```go
err := osx.Copy("source.txt", "dest.txt")
if err != nil {
    switch {
    case os.IsNotExist(err):
        fmt.Println("源文件未找到")
    case os.IsPermission(err):
        fmt.Println("权限被拒绝")
    default:
        fmt.Printf("复制失败: %v\n", err)
    }
}
```

## 常见用例

1. **配置文件管理**: 安全地检查和更新配置文件
2. **日志文件操作**: 管理日志轮转和清理
3. **临时文件处理**: 创建和管理临时文件
4. **构建系统**: 构建管道中的文件复制和组织
5. **嵌入式文件系统**: 处理嵌入式资源和模板
6. **数据处理**: 安全的文件处理与存在性检查

## 相关包

- [`os`](https://pkg.go.dev/os): Go 标准操作系统接口
- [`io/fs`](https://pkg.go.dev/io/fs): 文件系统接口定义
- [`path/filepath`](https://pkg.go.dev/path/filepath): 文件路径操作
- [`embed`](https://pkg.go.dev/embed): 嵌入式文件系统