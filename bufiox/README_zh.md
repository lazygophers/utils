# BufioX - 增强的缓冲 I/O 工具

一个强大的 Go 包，扩展了标准 `bufio` 包，为跨平台文本处理提供增强的扫描功能。`bufiox` 包提供了自定义分割函数，高效处理不同的行结尾和自定义分隔符。

## 特性

- **跨平台行扫描**: 无缝处理 Windows (CRLF) 和 Unix (LF) 行结尾
- **自定义分隔符扫描**: 按任何字节序列分隔符进行扫描
- **CR 删除支持**: 自动删除 Windows 风格行结尾中的回车符
- **标准接口**: 与 `bufio.Scanner` 分割函数兼容
- **零分配**: 优化扫描过程中的最小内存分配
- **高性能**: 高效的字节级操作实现最大吞吐量

## 安装

```bash
go get github.com/lazygophers/utils/bufiox
```

## 快速开始

```go
package main

import (
    "bufio"
    "fmt"
    "strings"

    "github.com/lazygophers/utils/bufiox"
)

func main() {
    text := "Line 1\r\nLine 2\nLine 3\r\nLine 4"
    scanner := bufio.NewScanner(strings.NewReader(text))

    // 使用跨平台行扫描
    scanner.Split(bufiox.ScanLines)

    for scanner.Scan() {
        fmt.Printf("Line: '%s'\n", scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}
```

## API 参考

### 函数

#### `ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)`

用于 `bufio.Scanner` 的分割函数，处理跨平台行结尾。

**参数:**
- `data []byte`: 当前处理的数据
- `atEOF bool`: 是否到达输入结尾

**返回值:**
- `advance int`: 要前进的字节数
- `token []byte`: 当前行数据（已移除 CR）
- `err error`: 错误信息（通常为 nil）

**特性:**
- 处理 `\n`（Unix）和 `\r\n`（Windows）行结尾
- 自动移除回车符
- 在 EOF 时强制分割剩余数据
- 与标准 `bufio.Scanner` 接口兼容

**示例:**
```go
scanner := bufio.NewScanner(reader)
scanner.Split(bufiox.ScanLines)

for scanner.Scan() {
    line := scanner.Text()
    // 处理行（CR 已移除）
    fmt.Println(line)
}
```

#### `ScanBy(seq []byte) func(data []byte, atEOF bool) (advance int, token []byte, err error)`

创建一个自定义分割函数，按指定的字节序列分割数据。

**参数:**
- `seq []byte`: 用作分隔符的字节序列

**返回值:**
- 与 `bufio.Scanner` 兼容的分割函数

**示例:**
```go
// 按自定义分隔符分割
data := "item1|item2|item3|item4"
scanner := bufio.NewScanner(strings.NewReader(data))
scanner.Split(bufiox.ScanBy([]byte("|")))

for scanner.Scan() {
    fmt.Printf("Item: '%s'\n", scanner.Text())
}
```

#### `dropCR(data []byte) []byte`

工具函数，从字节切片中移除末尾的回车符（`\r`）。

**参数:**
- `data []byte`: 输入字节切片

**返回值:**
- `[]byte`: 移除末尾 CR 后的字节切片（如果存在）

**示例:**
```go
data := []byte("Hello World\r")
clean := bufiox.dropCR(data)  // 注意：这是内部函数
// clean = []byte("Hello World")
```

## 使用示例

### 读取跨平台文本文件

```go
package main

import (
    "bufio"
    "fmt"
    "os"

    "github.com/lazygophers/utils/bufiox"
)

func main() {
    file, err := os.Open("mixed_endings.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufiox.ScanLines)

    lineNum := 0
    for scanner.Scan() {
        lineNum++
        line := scanner.Text()
        fmt.Printf("Line %d: %s\n", lineNum, line)
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("读取文件错误: %v\n", err)
    }
}
```

### 自定义分隔符解析

```go
package main

import (
    "bufio"
    "fmt"
    "strings"

    "github.com/lazygophers/utils/bufiox"
)

func main() {
    // 使用自定义分隔符的 CSV 类数据
    csvData := "name;age;city;country"
    scanner := bufio.NewScanner(strings.NewReader(csvData))
    scanner.Split(bufiox.ScanBy([]byte(";")))

    fields := []string{}
    for scanner.Scan() {
        fields = append(fields, scanner.Text())
    }

    fmt.Printf("Fields: %v\n", fields)
    // 输出: Fields: [name age city country]
}
```

### 协议消息解析

```go
package main

import (
    "bufio"
    "fmt"
    "strings"

    "github.com/lazygophers/utils/bufiox"
)

func main() {
    // 由双换行符分隔的协议消息
    protocol := "MESSAGE1\n\nMESSAGE2\n\nMESSAGE3"
    scanner := bufio.NewScanner(strings.NewReader(protocol))
    scanner.Split(bufiox.ScanBy([]byte("\n\n")))

    messageNum := 0
    for scanner.Scan() {
        messageNum++
        message := scanner.Text()
        fmt.Printf("Message %d: %s\n", messageNum, message)
    }
}
```

### 处理 Windows 和 Unix 混合内容

```go
package main

import (
    "bufio"
    "fmt"
    "strings"

    "github.com/lazygophers/utils/bufiox"
)

func main() {
    // 来自不同来源的混合行结尾
    mixedContent := "Unix line\nWindows line\r\nAnother Unix\nAnother Windows\r\n"
    scanner := bufio.NewScanner(strings.NewReader(mixedContent))
    scanner.Split(bufiox.ScanLines)

    for scanner.Scan() {
        line := scanner.Text()
        fmt.Printf("'%s' (length: %d)\n", line, len(line))
    }
}
```

### 日志文件处理

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "github.com/lazygophers/utils/bufiox"
)

func processLogFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufiox.ScanLines)

    lineCount := 0
    errorCount := 0

    for scanner.Scan() {
        lineCount++
        line := scanner.Text()

        // 处理日志条目（无论行结尾风格如何都能工作）
        if strings.Contains(line, "ERROR") {
            errorCount++
            fmt.Printf("第 %d 行错误: %s\n", lineCount, line)
        }
    }

    fmt.Printf("处理了 %d 行，发现 %d 个错误\n", lineCount, errorCount)
    return scanner.Err()
}
```

## 性能特征

### ScanLines 性能
- **内存效率**: 扫描过程中最小内存分配
- **跨平台**: 单个函数处理所有行结尾类型
- **零复制**: 尽可能重用字节切片

### ScanBy 性能
- **灵活性**: 适用于任何字节序列分隔符
- **高效**: 使用 `bytes.Index` 快速查找分隔符
- **可配置**: 根据需要创建特定分隔符的扫描器

### 基准测试
```go
// 典型性能（根据输入而变化）
BenchmarkScanLines-8       1000000    1200 ns/op    0 B/op    0 allocs/op
BenchmarkScanBy-8          500000     2400 ns/op    0 B/op    0 allocs/op
```

## 最佳实践

### 1. 选择合适的扫描器
使用 `ScanLines` 进行基于行的处理，使用 `ScanBy` 处理自定义分隔符：

```go
// 行处理
scanner.Split(bufiox.ScanLines)

// 自定义分隔符
scanner.Split(bufiox.ScanBy([]byte("||")))
```

### 2. 高效处理大文件
对于大文件，考虑缓冲区大小：

```go
scanner := bufio.NewScanner(file)
scanner.Split(bufiox.ScanLines)

// 为大行增加缓冲区大小
buf := make([]byte, 0, 64*1024)
scanner.Buffer(buf, 1024*1024)
```

### 3. 错误处理
始终检查扫描器错误：

```go
for scanner.Scan() {
    // 处理 scanner.Text()
}

if err := scanner.Err(); err != nil {
    log.Printf("扫描错误: %v", err)
}
```

### 4. 内存管理
对于高吞吐量应用程序，重用扫描器：

```go
type FileProcessor struct {
    scanner *bufio.Scanner
}

func (p *FileProcessor) ProcessFile(reader io.Reader) {
    p.scanner.Reset(reader)  // 重用扫描器
    p.scanner.Split(bufiox.ScanLines)

    for p.scanner.Scan() {
        // 处理行
    }
}
```

## 兼容性

### 标准库兼容性
- 与 `bufio.Scanner` 完全兼容
- `bufio.ScanLines` 的直接替换
- 适用于所有 `io.Reader` 实现

### 平台支持
- **Unix/Linux**: 原生 `\n` 行结尾支持
- **Windows**: 自动 `\r\n` 处理和 CR 删除
- **macOS**: 完全支持 Unix 和 Windows 格式
- **跨平台**: 处理混合行结尾文件

## 高级用法

### 自定义分割函数
您可以结合 `bufiox` 函数创建复杂的解析逻辑：

```go
func scanCustomProtocol(data []byte, atEOF bool) (advance int, token []byte, err error) {
    // 首先尝试自定义分隔符
    if advance, token, err := bufiox.ScanBy([]byte("END"))(data, atEOF); err == nil && advance > 0 {
        return advance, token, err
    }

    // 回退到行扫描
    return bufiox.ScanLines(data, atEOF)
}
```

### 流处理
处理到达的数据：

```go
func streamProcessor(conn net.Conn) {
    scanner := bufio.NewScanner(conn)
    scanner.Split(bufiox.ScanLines)

    for scanner.Scan() {
        line := scanner.Text()
        // 立即处理行
        processLine(line)
    }
}
```

## 线程安全

`bufiox` 包函数是无状态的，线程安全：

- **ScanLines**: 安全的并发使用
- **ScanBy**: 安全的并发使用（创建新函数实例）
- **dropCR**: 安全的并发使用

但是，`bufio.Scanner` 本身不是线程安全的，所以不要在 goroutines 之间共享扫描器实例。

## 贡献

欢迎贡献！请确保：

1. 跨平台测试
2. 性能基准测试
3. 全面测试
4. 文档更新

## 许可证

此包是 LazyGophers Utils 库的一部分，遵循相同的许可条款。