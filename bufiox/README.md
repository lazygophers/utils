# bufiox

## 项目简介
高效处理缓冲区的Go语言工具库，支持跨平台换行符处理和流式操作

## 核心函数
```go
func ScanBy(r io.Reader, delimiter byte) ([]byte, error)
func ScanLines(r io.Reader) ([]string, error)
```

## 使用示例
```go
// 基础使用
lines, _ := ScanLines(os.Stdin)
for _, line := range lines {
    fmt.Println(line)
}

// 自定义分隔符
data, _ := ScanBy(bytes.NewBuffer([]byte("a\nb\nc")), '\n')
```

## 跨平台支持
- 自动识别并处理 Windows(CRLF) / Linux(LF) / Mac(CR) 换行符
- 通过 `bytes.Equal(line, []byte{'\r', '\n'})` 可手动检测特定换行符

This directory contains buffered I/O utility functions for the lazygophers project. Current implementation includes:  
- `scan.go`: Implements custom buffered scanning logic for efficient data processing.