---
title: bufiox
description: 自訂数據掃描函數
---

# bufiox

`bufiox` 包提供用於數據分割的掃描函數，主要配合 `bufio.Scanner` 使用。

## 適用場景

- **自訂分隔符**：需要按定位元組序列分割數據（非標準換行符）
- **串流處理**：處理大文件或網絡串流，避免一次性加載
- **協議解析**：解析二進制協議或自訂格式文件

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
    fmt.Println(scanner.Text()) // 自動處理 CRLF/LF
}
```

### 自訂分隔符

```go
// 按自訂位元組序列分割
data := "A::B::C"
scanner := bufio.NewScanner(strings.NewReader(data))
scanner.Split(bufiox.ScanBy([]byte("::")))

for scanner.Scan() {
    fmt.Println(scanner.Text()) // "A", "B", "C"
}
```

## 與標準庫對比

| 功能 | 標準庫 | bufiox |
|------|--------|--------|
| 按行分割 | `bufio.ScanLines` | `bufiox.ScanLines`（相同） |
| 自訂分隔符 | 需自己實現 | `bufiox.ScanBy` |

## 使用建議

1. **標準日誌**：使用 `ScanLines`，自動處理 Windows/Unix 換行
2. **自訂協議**：使用 `ScanBy([]byte("DELIM"))` 按協議分隔符分割
3. **大文件處理**：配合 `bufio.Scanner` 避免記憶體溢出

## 範例

### 解析 CSV 文件

```go
file, _ := os.Open("data.csv")
scanner := bufio.NewScanner(file)
scanner.Split(bufiox.ScanBy([]byte(",")))

for scanner.Scan() {
    field := scanner.Text()
    // 處理每個字段
}
```

### 解析鍵值對

```go
data := "name=張三&age=25&city=台北"
scanner := bufio.NewScanner(strings.NewReader(data))
scanner.Split(bufiox.ScanBy([]byte("&")))

for scanner.Scan() {
    kv := strings.Split(scanner.Text(), "=")
    // 處理鍵值對
}
```
