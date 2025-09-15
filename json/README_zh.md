# json - 增强 JSON 处理

`json` 包提供了增强的 JSON 处理工具，具有更好的错误处理、类型安全和便利函数。它扩展了 Go 标准 `encoding/json` 包，为常见的 JSON 操作提供了附加功能。

## 功能特性

- **增强错误处理**: 更好的错误消息和错误上下文
- **类型安全操作**: 类型安全 JSON 操作的泛型函数
- **便利函数**: 简化的 Marshal/Unmarshal 操作
- **格式化打印**: 可自定义缩进的格式化 JSON 输出
- **路径操作**: 基于 JSON 路径的值提取和修改
- **验证**: JSON 模式验证和结构检查
- **流式支持**: 高效的流式 JSON 处理

## 安装

```bash
go get github.com/lazygophers/utils/json
```

## 使用示例

### 基本 JSON 操作

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/json"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    user := User{ID: 1, Name: "Alice", Age: 30}

    // 编组为 JSON 字符串
    jsonStr, err := json.MarshalString(user)
    if err != nil {
        panic(err)
    }
    fmt.Println(jsonStr) // {"id":1,"name":"Alice","age":30}

    // 从 JSON 字符串解组
    var newUser User
    err = json.UnmarshalString(jsonStr, &newUser)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%+v\n", newUser)
}
```

### 格式化打印

```go
user := User{ID: 1, Name: "Alice", Age: 30}

// 使用默认缩进的格式化打印
pretty, err := json.MarshalPretty(user)
if err != nil {
    panic(err)
}
fmt.Println(string(pretty))
// 输出:
// {
//   "id": 1,
//   "name": "Alice",
//   "age": 30
// }

// 使用自定义缩进的格式化打印
pretty, err = json.MarshalIndent(user, "", "    ")
if err != nil {
    panic(err)
}
fmt.Println(string(pretty))
```

### 类型安全操作

```go
// 泛型编组函数
data := map[string]int{"apple": 5, "banana": 3}
result, err := json.Marshal[map[string]int](data)
if err != nil {
    panic(err)
}

// 泛型解组函数
var restored map[string]int
err = json.Unmarshal[map[string]int](result, &restored)
if err != nil {
    panic(err)
}
fmt.Println(restored) // map[apple:5 banana:3]
```

### JSON 路径操作

```go
jsonData := `{
    "user": {
        "id": 123,
        "profile": {
            "name": "Alice",
            "age": 30
        }
    },
    "settings": {
        "theme": "dark"
    }
}`

// 通过路径提取值
name, err := json.GetValueByPath(jsonData, "user.profile.name")
if err != nil {
    panic(err)
}
fmt.Println(name) // "Alice"

// 通过路径设置值
modified, err := json.SetValueByPath(jsonData, "user.profile.age", 31)
if err != nil {
    panic(err)
}
fmt.Println(modified)
```

### 验证和结构检查

```go
// 验证 JSON 结构
valid := json.IsValidJSON(`{"name": "Alice", "age": 30}`)
fmt.Println(valid) // true

// 检查 JSON 是否有必需字段
hasFields := json.HasFields(jsonData, []string{"user.id", "user.profile.name"})
fmt.Println(hasFields) // true

// 根据模式验证
schema := `{
    "type": "object",
    "properties": {
        "name": {"type": "string"},
        "age": {"type": "number"}
    },
    "required": ["name", "age"]
}`

valid, err = json.ValidateSchema(jsonData, schema)
if err != nil {
    panic(err)
}
fmt.Println(valid)
```

### 流式操作

```go
// 流式编码器
var buf bytes.Buffer
encoder := json.NewStreamEncoder(&buf)

users := []User{
    {ID: 1, Name: "Alice", Age: 30},
    {ID: 2, Name: "Bob", Age: 25},
}

for _, user := range users {
    err := encoder.Encode(user)
    if err != nil {
        panic(err)
    }
}

// 流式解码器
decoder := json.NewStreamDecoder(&buf)
for decoder.More() {
    var user User
    err := decoder.Decode(&user)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%+v\n", user)
}
```

## API 参考

### 基本函数

- `Marshal(v interface{}) ([]byte, error)` - 将值编组为 JSON 字节
- `MarshalString(v interface{}) (string, error)` - 将值编组为 JSON 字符串
- `Unmarshal(data []byte, v interface{}) error` - 解组 JSON 字节
- `UnmarshalString(s string, v interface{}) error` - 解组 JSON 字符串

### 格式化打印

- `MarshalPretty(v interface{}) ([]byte, error)` - 使用默认格式化编组
- `MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)` - 使用自定义缩进编组
- `PrettyPrint(data []byte) ([]byte, error)` - 格式化现有 JSON 数据

### 泛型函数

- `Marshal[T any](v T) ([]byte, error)` - 类型安全编组
- `Unmarshal[T any](data []byte, v *T) error` - 类型安全解组
- `MarshalString[T any](v T) (string, error)` - 类型安全编组为字符串

### 路径操作

- `GetValueByPath(jsonStr, path string) (interface{}, error)` - 通过 JSON 路径获取值
- `SetValueByPath(jsonStr, path string, value interface{}) (string, error)` - 通过路径设置值
- `DeleteByPath(jsonStr, path string) (string, error)` - 通过路径删除值
- `PathExists(jsonStr, path string) bool` - 检查路径是否存在

### 验证

- `IsValidJSON(s string) bool` - 检查字符串是否为有效 JSON
- `ValidateSchema(jsonStr, schema string) (bool, error)` - 根据 JSON 模式验证
- `HasFields(jsonStr string, fields []string) bool` - 检查必需字段
- `GetType(jsonStr, path string) (string, error)` - 获取路径上值的类型

### 流式处理

- `NewStreamEncoder(w io.Writer) *StreamEncoder` - 创建流式编码器
- `NewStreamDecoder(r io.Reader) *StreamDecoder` - 创建流式解码器

### 工具函数

- `Merge(json1, json2 string) (string, error)` - 合并两个 JSON 对象
- `Filter(jsonStr string, keys []string) (string, error)` - 按键过滤 JSON
- `Transform(jsonStr string, transformer func(key, value interface{}) interface{}) (string, error)` - 转换 JSON 值
- `Compare(json1, json2 string) (bool, error)` - 比较两个 JSON 结构

## 错误处理

该包提供增强的错误信息：

```go
type JSONError struct {
    Op   string // 失败的操作
    Path string // 发生错误的 JSON 路径
    Err  error  // 底层错误
}

func (e *JSONError) Error() string {
    return fmt.Sprintf("json: %s 在路径 %s: %v", e.Op, e.Path, e.Err)
}
```

### 错误类型

- `IsSyntaxError(err error) bool` - 检查 JSON 语法错误
- `IsTypeError(err error) bool` - 检查类型转换错误
- `IsPathError(err error) bool` - 检查路径相关错误

## 性能优化

- **零拷贝操作**: 尽可能减少内存分配
- **流式支持**: 高效处理大型 JSON 文件
- **路径缓存**: 缓存经常使用的 JSON 路径
- **池重用**: 重用缓冲区和解码器以获得更好的性能

## 最佳实践

1. **使用类型安全函数**: 优先使用泛型函数以获得编译时安全性
2. **处理错误**: 始终检查和适当处理 JSON 错误
3. **验证输入**: 在处理之前验证 JSON 结构
4. **使用流式处理**: 对大型 JSON 数据集使用流式处理
5. **缓存路径**: 缓存经常访问的 JSON 路径

## 示例

### 配置处理

```go
configJSON := `{
    "server": {
        "host": "localhost",
        "port": 8080,
        "ssl": {
            "enabled": true,
            "cert": "/path/to/cert.pem"
        }
    },
    "database": {
        "url": "postgres://localhost/mydb",
        "pool_size": 10
    }
}`

// 提取特定配置值
host, _ := json.GetValueByPath(configJSON, "server.host")
port, _ := json.GetValueByPath(configJSON, "server.port")
sslEnabled, _ := json.GetValueByPath(configJSON, "server.ssl.enabled")

fmt.Printf("服务器: %s:%d (SSL: %v)\n", host, port, sslEnabled)
```

### API 响应处理

```go
type APIResponse struct {
    Status  string      `json:"status"`
    Data    interface{} `json:"data"`
    Message string      `json:"message,omitempty"`
}

// 创建和编组响应
response := APIResponse{
    Status: "success",
    Data:   map[string]interface{}{"user_id": 123, "name": "Alice"},
}

jsonResponse, err := json.MarshalPretty(response)
if err != nil {
    panic(err)
}

fmt.Println(string(jsonResponse))
```

### 数据转换

```go
// 转换 JSON 数据
originalJSON := `{"prices": [10, 20, 30], "currency": "USD"}`

transformed, err := json.Transform(originalJSON, func(key, value interface{}) interface{} {
    if key == "prices" && reflect.TypeOf(value).Kind() == reflect.Slice {
        // 将价格从 USD 转换为 EUR（示例汇率）
        prices := value.([]interface{})
        for i, p := range prices {
            if price, ok := p.(float64); ok {
                prices[i] = price * 0.85 // 示例转换率
            }
        }
        return prices
    }
    if key == "currency" {
        return "EUR"
    }
    return value
})

if err != nil {
    panic(err)
}
fmt.Println(transformed)
```

## 相关包

- `candy` - 类型转换工具
- `validator` - 结构体验证工具
- `stringx` - 字符串操作工具