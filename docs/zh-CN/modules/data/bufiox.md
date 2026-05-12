---
title: bufiox
description: 自定义数据扫描函数
---

# bufiox

`bufiox` 包提供用于数据分割的扫描函数，主要配合 `bufio.Scanner` 使用。

## 适用场景

- **自定义分隔符**：需要按指定字节序列分割数据（非标准换行符）
- **流式处理**：处理大文件或网络流，避免一次性加载
- **协议解析**：解析二进制协议或自定义格式文件

## 主要功能

### 按行分割

```go
import (
    "bufio"
    "strings"
    "github.com/lazygophers/utils/bufiox"
)

data := "line1\nline2\r\nline3"
scanner := bufio.NewScanner(strings.NewReader(data))
scanner.Split(bufiox.ScanLines)

for scanner.Scan() {
    fmt.Println(scanner.Text()) // 自动处理 CRLF/LF
}
```

### 自定义分隔符

```go
// 按自定义字节序列分割
data := "A::B::C"
scanner := bufio.NewScanner(strings.NewReader(data))
scanner.Split(bufiox.ScanBy([]byte("::")))

for scanner.Scan() {
    fmt.Println(scanner.Text()) // "A", "B", "C"
}
```

## 与标准库对比

| 功能 | 标准库 | bufiox |
|------|--------|--------|
| 按行分割 | `bufio.ScanLines` | `bufiox.ScanLines`（相同） |
| 自定义分隔符 | 需自己实现 | `bufiox.ScanBy` |

## 使用建议

1. **标准日志**：使用 `ScanLines`，自动处理 Windows/Unix 换行
2. **自定义协议**：使用 `ScanBy([]byte("DELIM"))` 按协议分隔符分割
3. **大文件处理**：配合 `bufio.Scanner` 避免内存溢出

## 示例

### 解析 CSV 文件

```go
file, _ := os.Open("data.csv")
scanner := bufio.NewScanner(file)
scanner.Split(bufiox.ScanBy([]byte(",")))

for scanner.Scan() {
    field := scanner.Text()
    // 处理每个字段
}
```

### 解析键值对

```go
data := "name=张三&age=25&city=北京"
scanner := bufio.NewScanner(strings.NewReader(data))
scanner.Split(bufiox.ScanBy([]byte("&")))

for scanner.Scan() {
    kv := strings.Split(scanner.Text(), "=")
    // 处理键值对
}
```
