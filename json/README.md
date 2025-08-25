# json

JSON 处理工具集，提供高性能的 JSON 序列化、反序列化以及文件操作功能。本模块通过条件编译在不同平台自动选择最优的 JSON 库实现。

## 功能特性

### 🚀 核心功能

- **智能引擎选择**：在 Linux AMD64 和 Darwin 平台使用高性能的 sonic 引擎，其他平台使用标准库
- **序列化/反序列化**：支持基本的数据结构转换
- **字符串操作**：提供便捷的字符串形式 JSON 处理
- **流式处理**：支持编码器和解码器模式

### 📁 文件操作

- **文件读取**：直接从文件反序列化到结构体
- **文件写入**：将结构体序列化并保存到文件
- **错误处理**：提供普通和 Must（panic）两种模式

### 🛠️ 辅助功能

- **格式化输出**：支持 JSON 缩进格式化
- **强制操作**：提供 Must 系列函数，出错时 panic

## 安装

```bash
go get github.com/lazygophers/utils/json
```

## 快速开始

### 基本序列化和反序列化

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/json"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func main() {
    // 序列化
    user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
    data, err := json.Marshal(user)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(data))
    // 输出: {"id":1,"name":"Alice","email":"alice@example.com"}
    
    // 反序列化
    var newUser User
    err = json.Unmarshal(data, &newUser)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%+v\n", newUser)
    // 输出: {ID:1 Name:Alice Email:alice@example.com}
}
```

### 字符串操作

```go
// 序列化为字符串
str, err := json.MarshalString(user)
if err != nil {
    panic(err)
}
fmt.Println(str)

// 从字符串反序列化
var u User
err = json.UnmarshalString(str, &u)
```

### 文件操作

```go
// 写入 JSON 文件
err := json.MarshalToFile("user.json", user)
if err != nil {
    panic(err)
}

// 从文件读取
var fileUser User
err = json.UnmarshalFromFile("user.json", &fileUser)
if err != nil {
    panic(err)
}
```

### Must 系列函数（出错时 panic）

```go
// MustMarshal - 出错时 panic
data := json.MustMarshal(user)

// MustMarshalString - 出错时 panic
str := json.MustMarshalString(user)

// MustMarshalToFile - 出错时 panic
json.MustMarshalToFile("user.json", user)

// MustUnmarshalFromFile - 出错时 panic
var mustUser User
json.MustUnmarshalFromFile("user.json", &mustUser)
```

### 流式处理

```go
import (
    "os"
    "github.com/lazygophers/utils/json"
)

// 编码器
file, _ := os.Create("stream.json")
encoder := json.NewEncoder(file)
encoder.Encode(user1)
encoder.Encode(user2)
defer file.Close()

// 解码器
file, _ = os.Open("stream.json")
decoder := json.NewDecoder(file)
var users []User
for decoder.More() {
    var u User
    if err := decoder.Decode(&u); err == nil {
        users = append(users, u)
    }
}
```

### JSON 格式化

```go
import (
    "bytes"
    "github.com/lazygophers/utils/json"
)

// 格式化 JSON 输出
var buf bytes.Buffer
err := json.Indent(&buf, data, "", "  ")  // 使用两个空格缩进
if err != nil {
    panic(err)
}
fmt.Println(buf.String())
```

## 性能优化

本模块在不同平台上自动选择最优的 JSON 处理引擎：

- **Linux AMD64 / Darwin**：使用 [sonic](https://github.com/bytedance/sonic) 引擎，性能提升 2-3 倍
- **其他平台**：使用标准库 `encoding/json`，保证兼容性

### 性能对比

| 操作 | 标准库 | sonic | 提升 |
|------|--------|-------|------|
| Marshal | 100% | ~300% | 3x |
| Unmarshal | 100% | ~200% | 2x |

## API 参考

### 序列化函数

- `func Marshal(v any) ([]byte, error)` - 序列化为字节切片
- `func MarshalString(v any) (string, error)` - 序列化为字符串
- `func MustMarshal(v any) []byte` - 序列化为字节切片（出错时 panic）
- `func MustMarshalString(v any) string` - 序列化为字符串（出错时 panic）

### 反序列化函数

- `func Unmarshal(data []byte, v any) error` - 从字节切片反序列化
- `func UnmarshalString(data string, v any) error` - 从字符串反序列化

### 文件操作函数

- `func MarshalToFile(filename string, v any) error` - 序列化到文件
- `func UnmarshalFromFile(filename string, v any) error` - 从文件反序列化
- `func MustMarshalToFile(filename string, v any)` - 序列化到文件（出错时 panic）
- `func MustUnmarshalFromFile(filename string, v any)` - 从文件反序列化（出错时 panic）

### 编码器/解码器

- `func NewEncoder(w io.Writer) *json.Encoder` - 创建编码器
- `func NewDecoder(r io.Reader) *json.Decoder` - 创建解码器

### 辅助函数

- `func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error` - 格式化 JSON

## 注意事项

1. **平台兼容性**：sonic 引擎仅在 Linux AMD64 和 Darwin 平台启用
2. **错误处理**：Must 系列函数会 panic，请确保在可恢复的环境中使用
3. **文件权限**：文件操作函数需要相应的读写权限
4. **内存管理**：大文件处理时注意内存使用，建议使用流式处理

## 许可证

本项目采用 AGPL-3.0 许可证，详见 [LICENSE](../../LICENSE) 文件。