# anyx - 泛型类型工具

`anyx` 包提供了处理 `interface{}` (any) 类型和泛型 map 操作的工具。它为从 map 中提取键和值以及转换数据结构提供了类型安全的操作。

## 功能特性

- **值类型检测**: 检测任何值的类型类别（数字、字符串、布尔值、未知）
- **Map 键提取**: 使用类型特定函数从 map 中提取键
- **Map 值提取**: 使用类型特定函数从 map 中提取值
- **类型转换**: 安全地在不同数据类型之间转换
- **切片转 Map**: 将切片转换为 map 以实现快速查找
- **KeyBy 操作**: 从结构体切片创建索引 map

## 安装

```bash
go get github.com/lazygophers/utils/anyx
```

## 使用示例

### 值类型检测

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/anyx"
)

func main() {
    fmt.Println(anyx.CheckValueType(42))       // ValueNumber
    fmt.Println(anyx.CheckValueType("hello"))  // ValueString
    fmt.Println(anyx.CheckValueType(true))     // ValueBool
}
```

### Map 键提取

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
keys := anyx.MapKeysString(m)
fmt.Println(keys) // [a b c]

numMap := map[int]string{1: "one", 2: "two"}
numKeys := anyx.MapKeysInt(numMap)
fmt.Println(numKeys) // [1 2]
```

### Map 值提取

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}
values := anyx.MapValues(m)
fmt.Println(values) // [1 2 3]

// 对于 interface{} map
anyMap := map[string]interface{}{"x": 1, "y": "hello"}
anyValues := anyx.MapValuesAny(anyMap)
fmt.Println(anyValues) // [1 hello]
```

### KeyBy 操作

```go
type User struct {
    ID   uint64
    Name string
}

users := []*User{
    {ID: 1, Name: "Alice"},
    {ID: 2, Name: "Bob"},
}

// 创建按 ID 索引的 map
userMap := anyx.KeyByUint64(users, "ID")
fmt.Println(userMap[1].Name) // Alice

// 创建按名称索引的 map
nameMap := anyx.KeyByString(users, "Name")
fmt.Println(nameMap["Bob"].ID) // 2
```

### 切片转 Map 转换

```go
slice := []string{"apple", "banana", "cherry"}
lookup := anyx.Slice2Map(slice)
fmt.Println(lookup["apple"])  // true
fmt.Println(lookup["grape"])  // false
```

## API 参考

### 值类型函数

- `CheckValueType(val interface{}) ValueType` - 确定值的类型类别

### Map 键提取函数

- `MapKeysString(m interface{}) []string` - 提取字符串键
- `MapKeysInt(m interface{}) []int` - 提取 int 键
- `MapKeysUint64(m interface{}) []uint64` - 提取 uint64 键
- `MapKeysFloat64(m interface{}) []float64` - 提取 float64 键
- `MapKeysAny(m interface{}) []interface{}` - 提取任何类型的键

### Map 值提取函数

- `MapValues[K, V any](m map[K]V) []V` - 提取值（泛型）
- `MapValuesAny(m interface{}) []interface{}` - 提取任何类型的值
- `MapValuesString(m interface{}) []string` - 提取字符串值
- `MapValuesInt(m interface{}) []int` - 提取 int 值

### KeyBy 函数

- `KeyBy(list interface{}, fieldName string) interface{}` - 通用 KeyBy 操作
- `KeyByUint64[M any](list []*M, fieldName string) map[uint64]*M` - 使用 uint64 键的 KeyBy
- `KeyByString[M any](list []*M, fieldName string) map[string]*M` - 使用字符串键的 KeyBy
- `KeyByInt64[M any](list []*M, fieldName string) map[int64]*M` - 使用 int64 键的 KeyBy

### 工具函数

- `MergeMap[K, V any](source, target map[K]V) map[K]V` - 合并两个 map
- `Slice2Map[M comparable](list []M) map[M]bool` - 将切片转换为查找 map

## 性能

该包针对性能进行了优化：
- 使用已知容量预分配切片
- 最小化内存分配
- 使用类型特定函数以避免不必要的反射开销

## 注意事项

- 使用 `interface{}` 的函数使用反射，如果类型不匹配预期可能会 panic
- 泛型版本（使用类型参数）因更好的类型安全性和性能而更受推荐
- 所有键提取函数都能优雅地处理 nil map
- KeyBy 函数期望结构体元素，如果找不到字段将会 panic

## 相关包

- `candy` - 类型转换工具
- `stringx` - 字符串操作工具
- `randx` - 随机值生成