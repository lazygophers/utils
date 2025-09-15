# stringx - 高级字符串操作工具

`stringx` 包提供高性能的字符串操作工具，具有零拷贝优化、Unicode 支持和高级字符串处理功能。它扩展了 Go 标准 `strings` 包，为常见字符串操作提供附加功能。

## 功能特性

- **零拷贝转换**: 使用 unsafe 操作进行高效的字符串/字节切片转换
- **大小写转换**: 驼峰命名、蛇形命名、短横线命名和其他大小写转换
- **字符串生成**: 可自定义字符集的随机字符串生成
- **Unicode 支持**: 完整的 Unicode 文本处理支持
- **性能优化**: 内存高效操作，最小化分配
- **验证函数**: 字符串验证和检查工具
- **文本处理**: 高级文本操作和格式化函数

## 安装

```bash
go get github.com/lazygophers/utils/stringx
```

## 使用示例

### 零拷贝字符串/字节转换

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/stringx"
)

func main() {
    // 零拷贝字符串到字节转换
    str := "hello world"
    bytes := stringx.ToBytes(str)
    fmt.Printf("字符串: %s, 字节: %v\n", str, bytes)

    // 零拷贝字节到字符串转换
    data := []byte("hello world")
    result := stringx.ToString(data)
    fmt.Printf("字节: %v, 字符串: %s\n", data, result)
}
```

### 大小写转换

```go
// 驼峰命名转蛇形命名
camelCase := "getUserProfile"
snakeCase := stringx.Camel2Snake(camelCase)
fmt.Println(snakeCase) // get_user_profile

// 蛇形命名转驼峰命名
snakeStr := "user_profile_data"
camelStr := stringx.Snake2Camel(snakeStr)
fmt.Println(camelStr) // userProfileData

// 短横线命名转换
kebabCase := stringx.Camel2Kebab("getUserProfile")
fmt.Println(kebabCase) // get-user-profile

// 帕斯卡命名转换
pascalCase := stringx.Snake2Pascal("user_profile")
fmt.Println(pascalCase) // UserProfile
```

### 随机字符串生成

```go
// 使用默认字符集（字母数字）生成随机字符串
randomStr := stringx.RandString(10)
fmt.Println(randomStr) // 例如："Kj8mNq2Lp9"

// 使用自定义字符集生成随机字符串
customStr := stringx.RandStringWithCharset(8, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
fmt.Println(customStr) // 例如："MXKQWERT"

// 生成随机字母数字字符串
alphaNum := stringx.RandAlphaNumeric(12)
fmt.Println(alphaNum) // 例如："Abc123Def456"

// 生成随机字母字符串
alpha := stringx.RandAlphabetic(6)
fmt.Println(alpha) // 例如："XyZabc"

// 生成随机数字字符串
numeric := stringx.RandNumeric(4)
fmt.Println(numeric) // 例如："1234"
```

### 字符串验证和检查

```go
// 检查字符串是否为字母数字
isAlphaNum := stringx.IsAlphaNumeric("Abc123")
fmt.Println(isAlphaNum) // true

// 检查字符串是否仅为字母
isAlpha := stringx.IsAlphabetic("HelloWorld")
fmt.Println(isAlpha) // true

// 检查字符串是否仅为数字
isNumeric := stringx.IsNumeric("12345")
fmt.Println(isNumeric) // true

// 检查字符串是否为 ASCII
isASCII := stringx.IsASCII("Hello World")
fmt.Println(isASCII) // true

// 检查字符串是否为空或空白
isEmpty := stringx.IsBlank("   ")
fmt.Println(isEmpty) // true
```

### 文本处理

```go
// 反转字符串
reversed := stringx.Reverse("hello")
fmt.Println(reversed) // "olleh"

// 首字母大写
capitalized := stringx.Capitalize("hello world")
fmt.Println(capitalized) // "Hello world"

// 标题大小写
title := stringx.ToTitle("hello world")
fmt.Println(title) // "Hello World"

// 居中对齐字符串
centered := stringx.Center("hello", 10, " ")
fmt.Println(centered) // "  hello   "

// 截断字符串并添加省略号
truncated := stringx.Truncate("This is a long string", 10)
fmt.Println(truncated) // "This is..."

// 从字符串中移除重复字符
noDupes := stringx.RemoveDuplicates("hello")
fmt.Println(noDupes) // "helo"
```

### 高级字符串操作

```go
// 分割字符串并去除空格
parts := stringx.SplitAndTrim("apple, banana, cherry", ",")
fmt.Println(parts) // ["apple", "banana", "cherry"]

// 连接非空字符串
joined := stringx.JoinNonEmpty([]string{"hello", "", "world", ""}, " ")
fmt.Println(joined) // "hello world"

// 从字符串中提取单词
words := stringx.ExtractWords("Hello, World! How are you?")
fmt.Println(words) // ["Hello", "World", "How", "are", "you"]

// 计算子字符串出现次数
count := stringx.CountOccurrences("hello world hello", "hello")
fmt.Println(count) // 2

// 替换多个子字符串
replacements := map[string]string{"hello": "hi", "world": "earth"}
replaced := stringx.ReplaceMultiple("hello world", replacements)
fmt.Println(replaced) // "hi earth"
```

## API 参考

### 零拷贝转换

- `ToString(b []byte) string` - 将字节切片转换为字符串（零拷贝）
- `ToBytes(s string) []byte` - 将字符串转换为字节切片（零拷贝）

### 大小写转换

- `Camel2Snake(s string) string` - 将 camelCase 转换为 snake_case
- `Snake2Camel(s string) string` - 将 snake_case 转换为 camelCase
- `Camel2Kebab(s string) string` - 将 camelCase 转换为 kebab-case
- `Snake2Pascal(s string) string` - 将 snake_case 转换为 PascalCase
- `Pascal2Snake(s string) string` - 将 PascalCase 转换为 snake_case
- `Kebab2Camel(s string) string` - 将 kebab-case 转换为 camelCase

### 随机字符串生成

- `RandString(length int) string` - 生成随机字母数字字符串
- `RandStringWithCharset(length int, charset string) string` - 使用自定义字符集生成
- `RandAlphaNumeric(length int) string` - 生成字母数字字符串
- `RandAlphabetic(length int) string` - 生成字母字符串
- `RandNumeric(length int) string` - 生成数字字符串
- `RandHex(length int) string` - 生成十六进制字符串
- `RandBase64(length int) string` - 生成 base64 字符串

### 字符串验证

- `IsAlphaNumeric(s string) bool` - 检查字符串是否为字母数字
- `IsAlphabetic(s string) bool` - 检查字符串是否为字母
- `IsNumeric(s string) bool` - 检查字符串是否为数字
- `IsASCII(s string) bool` - 检查字符串是否仅包含 ASCII 字符
- `IsBlank(s string) bool` - 检查字符串是否为空或空白
- `IsEmpty(s string) bool` - 检查字符串是否为空
- `IsUpper(s string) bool` - 检查字符串是否为大写
- `IsLower(s string) bool` - 检查字符串是否为小写

### 文本处理

- `Reverse(s string) string` - 反转字符串
- `Capitalize(s string) string` - 首字母大写
- `ToTitle(s string) string` - 转换为标题大小写
- `Center(s string, width int, fillChar string) string` - 居中对齐
- `PadLeft(s string, width int, padChar string) string` - 左填充
- `PadRight(s string, width int, padChar string) string` - 右填充
- `Truncate(s string, length int) string` - 截断并添加省略号
- `TruncateWords(s string, wordCount int) string` - 按单词数截断

### 高级操作

- `SplitAndTrim(s, sep string) []string` - 分割并去除空白
- `JoinNonEmpty(strs []string, sep string) string` - 连接非空字符串
- `ExtractWords(s string) []string` - 从文本中提取单词
- `CountOccurrences(s, substr string) int` - 计算子字符串出现次数
- `ReplaceMultiple(s string, replacements map[string]string) string` - 多重替换
- `RemoveDuplicates(s string) string` - 移除重复字符
- `Similarity(s1, s2 string) float64` - 计算字符串相似度

### Unicode 支持

- `ContainsUnicode(s string) bool` - 检查字符串是否包含 Unicode
- `UnicodeLength(s string) int` - 获取 Unicode 字符数
- `UnicodeSubstring(s string, start, length int) string` - Unicode 感知的子字符串
- `NormalizeSpaces(s string) string` - 规范化空白字符

## 性能特性

该包针对高性能进行了优化：

### 零拷贝操作
```go
// 这些操作不分配新内存
bytes := stringx.ToBytes("hello")    // O(1)，无分配
str := stringx.ToString([]byte{...}) // O(1)，无分配
```

### 优化的大小写转换
- 仅 ASCII 字符串使用优化的快速路径
- Unicode 字符串使用高效缓冲
- 预计算容量以最小化重新分配

### 内存高效的随机生成
- 重用字符集和缓冲区
- 重复调用的最小分配

## 最佳实践

1. **使用零拷贝函数**: 使用 `ToString()` 和 `ToBytes()` 进行高效转换
2. **选择适当的函数**: 使用特定验证函数而不是正则表达式
3. **批量操作预分配**: 处理大量字符串时考虑预分配
4. **正确处理 Unicode**: 对国际文本使用 Unicode 感知函数
5. **验证输入**: 在公共 API 中始终验证字符串输入

## 示例

### 配置键转换

```go
// 在不同格式间转换配置键
configKeys := []string{
    "database_host",
    "database_port",
    "api_timeout",
    "log_level",
}

// 转换为环境变量格式
for _, key := range configKeys {
    envKey := strings.ToUpper(key)
    fmt.Printf("%s -> %s\n", key, envKey)
}

// 转换为 camelCase 用于 JSON
for _, key := range configKeys {
    jsonKey := stringx.Snake2Camel(key)
    fmt.Printf("%s -> %s\n", key, jsonKey)
}
```

### 文本处理管道

```go
// 处理用户输入文本
userInput := "  Hello, WORLD!  How are YOU today?  "

// 清理和规范化
cleaned := strings.TrimSpace(userInput)
normalized := stringx.NormalizeSpaces(cleaned)
title := stringx.ToTitle(normalized)

fmt.Println(title) // "Hello, World! How Are You Today?"

// 提取和分析单词
words := stringx.ExtractWords(normalized)
fmt.Printf("单词数: %d\n", len(words))
```

### ID 和令牌生成

```go
// 生成各种类型的标识符
sessionID := stringx.RandHex(32)
apiKey := stringx.RandBase64(24)
userToken := stringx.RandAlphaNumeric(16)

fmt.Printf("会话 ID: %s\n", sessionID)
fmt.Printf("API 密钥: %s\n", apiKey)
fmt.Printf("用户令牌: %s\n", userToken)
```

## 相关包

- `candy` - 类型转换工具
- `validator` - 字符串验证工具
- `json` - JSON 字符串处理