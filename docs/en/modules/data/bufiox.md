---
title: bufiox
description: Custom data scanning functions
---

# bufiox

The `bufiox` package provides scanning functions for data splitting, primarily for use with `bufio.Scanner`.

## Use Cases

- **Custom delimiters**: Split data by specified byte sequences (non-standard newlines)
- **Stream processing**: Process large files or network streams without loading entirely
- **Protocol parsing**: Parse binary protocols or custom format files

## Key Features

### Split by lines

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
    fmt.Println(scanner.Text()) // Automatically handles CRLF/LF
}
```

### Custom delimiter

```go
// Split by custom byte sequence
data := "A::B::C"
scanner := bufio.NewScanner(strings.NewReader(data))
scanner.Split(bufiox.ScanBy([]byte("::")))

for scanner.Scan() {
    fmt.Println(scanner.Text()) // "A", "B", "C"
}
```

## Comparison with Standard Library

| Feature | Standard Library | bufiox |
|---------|------------------|--------|
| Split by lines | `bufio.ScanLines` | `bufiox.ScanLines` (same) |
| Custom delimiter | Must implement yourself | `bufiox.ScanBy` |

## Usage Recommendations

1. **Standard logs**: Use `ScanLines`, automatically handles Windows/Unix newlines
2. **Custom protocols**: Use `ScanBy([]byte("DELIM"))` to split by protocol delimiters
3. **Large file processing**: Use with `bufio.Scanner` to avoid memory overflow

## Examples

### Parse CSV files

```go
file, _ := os.Open("data.csv")
scanner := bufio.NewScanner(file)
scanner.Split(bufiox.ScanBy([]byte(",")))

for scanner.Scan() {
    field := scanner.Text()
    // Process each field
}
```

### Parse key-value pairs

```go
data := "name=John&age=25&city=NYC"
scanner := bufio.NewScanner(strings.NewReader(data))
scanner.Split(bufiox.ScanBy([]byte("&")))

for scanner.Scan() {
    kv := strings.Split(scanner.Text(), "=")
    // Process key-value pairs
}
```
