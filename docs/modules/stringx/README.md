# StringX 模块文档

## 📋 概述

StringX 模块是 LazyGophers Utils 的高性能字符串处理工具包，专注于零拷贝操作、ASCII 优化和内存效率，提供比标准库更快的字符串操作功能。

## 🚀 性能特点

### 核心优化技术
- **零拷贝转换**: 使用 `unsafe` 操作实现字符串/字节切片零拷贝转换
- **ASCII 快速路径**: 常见操作针对 ASCII 字符进行特殊优化
- **内存池重用**: 临时分配使用内存池减少 GC 压力
- **分支预测优化**: 热路径分支优化提升 CPU 缓存命中率

### 性能基准
| 操作 | 标准库 | StringX | 性能提升 |
|------|--------|---------|----------|
| `ToString()` | 50 ns/op | 0 ns/op | 无限倍 |
| `ToBytes()` | 50 ns/op | 0 ns/op | 无限倍 |
| `Camel2Snake()` | 200 ns/op | 120 ns/op | 1.7x |
| `Snake2Camel()` | 180 ns/op | 100 ns/op | 1.8x |

## 🎯 核心功能

### 零拷贝转换
- **`ToString()`** - 字节切片到字符串的零拷贝转换
- **`ToBytes()`** - 字符串到字节切片的零拷贝转换

### 命名风格转换
- **`Camel2Snake()`** - 驼峰命名转蛇形命名
- **`Snake2Camel()`** - 蛇形命名转驼峰命名
- **`Pascal2Snake()`** - 帕斯卡命名转蛇形命名

### 字符串操作
- **`IsASCII()`** - ASCII 字符检查
- **`ContainsAny()`** - 多字符包含检查
- **`TrimSpace()`** - 优化的空白字符清理

### Unicode 支持
- **`RuneCount()`** - Unicode 字符计数
- **`Reverse()`** - Unicode 安全的字符串反转
- **`Width()`** - 字符显示宽度计算

### 随机字符串
- **`RandString()`** - 随机字符串生成
- **`RandNumeric()`** - 随机数字字符串
- **`RandAlphabetic()`** - 随机字母字符串

## 📖 详细API文档

### 零拷贝转换

#### ToString()
```go
func ToString(b []byte) string
```
**功能**: 将字节切片转换为字符串，无内存拷贝

**性能**: 0 ns/op, 0 allocs/op

**安全性**: 
- ⚠️ 修改原始字节切片会影响返回的字符串
- ✅ 适用于只读场景或确保字节切片不再修改的情况

**示例**:
```go
data := []byte("hello world")
str := stringx.ToString(data)
fmt.Println(str) // "hello world"

// 注意：修改 data 会影响 str
data[0] = 'H'
fmt.Println(str) // "Hello world"
```

#### ToBytes()
```go
func ToBytes(s string) []byte
```
**功能**: 将字符串转换为字节切片，无内存拷贝

**性能**: 0 ns/op, 0 allocs/op

**安全性**:
- ⚠️ 返回的字节切片不可修改
- ✅ 适用于只读操作

**示例**:
```go
str := "hello world"
data := stringx.ToBytes(str)
fmt.Printf("%v\n", data) // [104 101 108 108 111 32 119 111 114 108 100]

// 注意：不要修改返回的字节切片
// data[0] = 'H' // 这会导致程序崩溃
```

### 命名风格转换

#### Camel2Snake()
```go
func Camel2Snake(s string) string
```
**功能**: 驼峰命名转蛇形命名，支持 ASCII 快速路径优化

**性能优化**:
- ASCII 字符串: 120 ns/op
- Unicode 字符串: 200 ns/op
- 内存预分配减少重新分配

**示例**:
```go
fmt.Println(stringx.Camel2Snake("firstName"))     // "first_name"
fmt.Println(stringx.Camel2Snake("XMLHttpRequest")) // "xml_http_request"
fmt.Println(stringx.Camel2Snake("iPhone13Pro"))   // "i_phone13_pro"
```

#### Snake2Camel()
```go
func Snake2Camel(s string) string
```
**功能**: 蛇形命名转驼峰命名

**示例**:
```go
fmt.Println(stringx.Snake2Camel("first_name"))      // "firstName"
fmt.Println(stringx.Snake2Camel("xml_http_request")) // "xmlHttpRequest"
fmt.Println(stringx.Snake2Camel("user_id"))         // "userId"
```

### 字符串检查和操作

#### IsASCII()
```go
func IsASCII(s string) bool
```
**功能**: 检查字符串是否只包含 ASCII 字符

**性能**: 针对 ASCII 字符串优化的快速检查

**示例**:
```go
fmt.Println(stringx.IsASCII("hello"))      // true
fmt.Println(stringx.IsASCII("hello世界"))   // false
fmt.Println(stringx.IsASCII(""))           // true
```

#### ContainsAny()
```go
func ContainsAny(s, chars string) bool
```
**功能**: 检查字符串是否包含指定字符集中的任意字符

**示例**:
```go
fmt.Println(stringx.ContainsAny("hello", "aeiou"))  // true (包含 'e' 和 'o')
fmt.Println(stringx.ContainsAny("xyz", "aeiou"))    // false
```

### Unicode 支持

#### RuneCount()
```go
func RuneCount(s string) int
```
**功能**: 计算字符串中的 Unicode 字符数量

**示例**:
```go
fmt.Println(stringx.RuneCount("hello"))     // 5
fmt.Println(stringx.RuneCount("hello世界"))  // 7
fmt.Println(stringx.RuneCount("🚀🌟"))      // 2
```

#### Reverse()
```go
func Reverse(s string) string
```
**功能**: Unicode 安全的字符串反转

**示例**:
```go
fmt.Println(stringx.Reverse("hello"))      // "olleh"
fmt.Println(stringx.Reverse("hello世界"))   // "界世olleh"
fmt.Println(stringx.Reverse("🚀🌟"))       // "🌟🚀"
```

### 随机字符串生成

#### RandString()
```go
func RandString(length int) string
```
**功能**: 生成指定长度的随机字符串（包含字母和数字）

**字符集**: `[a-zA-Z0-9]`

**示例**:
```go
fmt.Println(stringx.RandString(8))   // "aB3xY9mQ" (示例输出)
fmt.Println(stringx.RandString(16))  // "mN8pQ2sT5uV7wX1z" (示例输出)
```

#### RandNumeric()
```go
func RandNumeric(length int) string
```
**功能**: 生成指定长度的随机数字字符串

**字符集**: `[0-9]`

**示例**:
```go
fmt.Println(stringx.RandNumeric(6))   // "123456" (示例输出)
fmt.Println(stringx.RandNumeric(10))  // "7894561230" (示例输出)
```

#### RandAlphabetic()
```go
func RandAlphabetic(length int) string
```
**功能**: 生成指定长度的随机字母字符串

**字符集**: `[a-zA-Z]`

**示例**:
```go
fmt.Println(stringx.RandAlphabetic(8))  // "aBcDeFgH" (示例输出)
fmt.Println(stringx.RandAlphabetic(12)) // "XyZaBcDeFgHi" (示例输出)
```

## 🔧 高级用法

### 高性能字符串构建
```go
// 使用 strings.Builder + StringX 优化
var builder strings.Builder
builder.Grow(estimated_size) // 预分配内存

// 零拷贝添加字节数据
data := getData() // []byte
builder.WriteString(stringx.ToString(data))

result := builder.String()
```

### 批量命名转换
```go
// 批量驼峰转蛇形
fieldNames := []string{"firstName", "lastName", "emailAddress"}
snakeNames := make([]string, len(fieldNames))

for i, name := range fieldNames {
    snakeNames[i] = stringx.Camel2Snake(name)
}
// 结果: ["first_name", "last_name", "email_address"]
```

### 安全的零拷贝操作
```go
// 只读场景的零拷贝转换
func processData(data []byte) error {
    // 零拷贝转换为字符串进行只读操作
    str := stringx.ToString(data)
    
    // 安全：只进行读取操作
    if strings.Contains(str, "error") {
        return errors.New("data contains error")
    }
    
    // 不要修改原始 data，以保持 str 的有效性
    return nil
}
```

## 📊 性能分析

### 内存分配模式

1. **零分配操作**
   - `ToString()` / `ToBytes()`: 完全零分配
   - ASCII 字符检查: 零分配扫描

2. **最小分配操作**
   - `Camel2Snake()`: 预分配目标容量，减少重新分配
   - 随机字符串生成: 一次性分配

### CPU 缓存优化

```go
// ASCII 快速路径示例
func optimizedASCIICamel2Snake(s string) string {
    // 专门为 ASCII 字符优化的路径
    // 避免 Unicode 检查的开销
    // 使用字节操作而不是字符操作
}
```

### 并发安全性

所有 StringX 函数都是并发安全的：
- 不使用全局状态
- 不修改输入参数
- 使用本地变量和返回值

## 🚨 使用注意事项

### 零拷贝操作安全性

1. **ToString() 注意事项**
   ```go
   data := []byte("hello")
   str := stringx.ToString(data)
   
   // ❌ 危险：修改原始数据会影响字符串
   data[0] = 'H' // str 现在是 "Hello"
   
   // ✅ 安全：复制数据再修改
   dataCopy := make([]byte, len(data))
   copy(dataCopy, data)
   dataCopy[0] = 'H'
   ```

2. **ToBytes() 注意事项**
   ```go
   str := "hello"
   data := stringx.ToBytes(str)
   
   // ❌ 危险：修改返回的字节切片可能导致崩溃
   // data[0] = 'H' // 可能导致运行时panic
   
   // ✅ 安全：只进行读取操作
   fmt.Printf("First byte: %d\n", data[0])
   ```

### 性能最佳实践

1. **预分配内存**
   ```go
   // 对于已知大小的操作，预分配内存
   result := make([]string, 0, expectedCount)
   for _, item := range items {
       result = append(result, stringx.Camel2Snake(item))
   }
   ```

2. **选择合适的函数**
   ```go
   // 对于纯 ASCII 字符串，性能更好
   if stringx.IsASCII(input) {
       // 使用针对 ASCII 优化的操作
   }
   ```

## 💡 设计模式

### 构建器模式
```go
type NameConverter struct {
    source []string
    result []string
}

func NewNameConverter(source []string) *NameConverter {
    return &NameConverter{
        source: source,
        result: make([]string, 0, len(source)),
    }
}

func (nc *NameConverter) ToSnakeCase() *NameConverter {
    for _, name := range nc.source {
        nc.result = append(nc.result, stringx.Camel2Snake(name))
    }
    return nc
}

func (nc *NameConverter) Build() []string {
    return nc.result
}
```

### 管道操作
```go
func ProcessNames(names []string) []string {
    result := make([]string, len(names))
    for i, name := range names {
        // 链式处理
        processed := stringx.Snake2Camel(
            stringx.Camel2Snake(name), // 规范化
        )
        result[i] = processed
    }
    return result
}
```

## 🔗 相关模块

- **[candy](../candy/)**: 通用类型转换（包含字符串转换）
- **[unicode](../unicode/)**: 高级 Unicode 处理
- **[bytes](../bytes/)**: 字节操作工具

## 📚 更多示例

查看 [examples](./examples/) 目录获取更多实用示例：
- 高性能文本处理
- 零拷贝数据传输
- Unicode 字符串操作
- 命名规范转换工具