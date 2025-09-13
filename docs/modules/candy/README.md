# Candy 模块文档

## 📋 概述

Candy 模块是 LazyGophers Utils 的核心类型转换工具包，提供了丰富的"语法糖"和便捷函数，专注于安全、高效的类型转换操作。

## 🎯 核心功能

### 类型转换
- **布尔转换**: `ToBool()` - 智能布尔值转换
- **数字转换**: `ToInt()`, `ToFloat64()`, `ToUint()` 等完整数字类型转换
- **字符串转换**: `ToString()`, `ToBytes()` - 高性能字符串处理
- **切片转换**: `ToInt64Slice()`, `ToFloat64Slice()` 等批量转换
- **映射转换**: `ToMap()`, `ToMapStringAny()` 等复杂结构转换

### 集合操作
- **数组处理**: `All()`, `Any()`, `Contains()` - 数组状态检查
- **过滤操作**: `Filter()`, `FilterNot()` - 条件过滤
- **聚合操作**: `Sum()`, `Average()`, `Min()`, `Max()` - 数值统计
- **去重操作**: `Unique()`, `UniqueUsing()` - 数据去重
- **排序操作**: `Sort()`, `SortUsing()` - 灵活排序

### 数学运算
- **基础运算**: `Abs()`, `Pow()`, `Sqrt()`, `Cbrt()` - 数学函数
- **随机操作**: `Random()`, `Shuffle()` - 随机数生成和数组洗牌

### 实用工具
- **深度操作**: `DeepCopy()`, `DeepEqual()` - 深度拷贝和比较
- **数组分块**: `Chunk()`, `Drop()` - 数组分割和处理
- **数据提取**: `Pluck()` 系列 - 结构体字段提取

## 📖 详细文档

### 类型转换函数

#### ToBool()
```go
func ToBool(v interface{}) bool
```
**功能**: 将任意类型转换为布尔值

**转换规则**:
- **bool**: 直接返回原值
- **数字类型**: 0 为 false，其他为 true
- **浮点类型**: 0.0 或 NaN 为 false，其他为 true
- **字符串/字节**: "true", "1", "t", "y", "yes", "on" 为 true；"false", "0", "f", "n", "no", "off" 为 false
- **其他**: 根据具体类型判断

**示例**:
```go
fmt.Println(candy.ToBool(1))        // true
fmt.Println(candy.ToBool(0))        // false
fmt.Println(candy.ToBool("yes"))    // true
fmt.Println(candy.ToBool("false"))  // false
```

#### ToString()
```go
func ToString(v interface{}) string
```
**功能**: 将任意类型转换为字符串

**性能特点**: 
- 基础类型: O(1) 时间复杂度
- 复杂类型: 使用 JSON 序列化，O(n) 时间复杂度

**示例**:
```go
fmt.Println(candy.ToString(123))           // "123"
fmt.Println(candy.ToString([]int{1, 2, 3})) // "[1,2,3]"
```

#### ToInt64()
```go
func ToInt64(v interface{}) int64
```
**功能**: 将任意类型转换为 int64

**支持类型**:
- 所有整型和浮点型
- 字符串数字
- 布尔值 (true=1, false=0)

**示例**:
```go
fmt.Println(candy.ToInt64("123"))   // 123
fmt.Println(candy.ToInt64(3.14))    // 3
fmt.Println(candy.ToInt64(true))    // 1
```

### 集合操作函数

#### All()
```go
func All[T any](slice []T, predicate func(T) bool) bool
```
**功能**: 检查切片中所有元素是否都满足条件

**示例**:
```go
numbers := []int{2, 4, 6, 8}
result := candy.All(numbers, func(n int) bool {
    return n%2 == 0
})
fmt.Println(result) // true (所有数字都是偶数)
```

#### Filter()
```go
func Filter[T any](slice []T, predicate func(T) bool) []T
```
**功能**: 根据条件过滤切片元素

**示例**:
```go
numbers := []int{1, 2, 3, 4, 5, 6}
evens := candy.Filter(numbers, func(n int) bool {
    return n%2 == 0
})
fmt.Println(evens) // [2, 4, 6]
```

#### Sum()
```go
func Sum[T Number](slice []T) T
```
**功能**: 计算数字切片的总和

**约束**: T 必须是数字类型 (int, float 等)

**示例**:
```go
numbers := []int{1, 2, 3, 4, 5}
total := candy.Sum(numbers)
fmt.Println(total) // 15
```

### 数学运算函数

#### Abs()
```go
func Abs[T Number](x T) T
```
**功能**: 计算数字的绝对值

**示例**:
```go
fmt.Println(candy.Abs(-5))    // 5
fmt.Println(candy.Abs(3.14))  // 3.14
```

#### Random()
```go
func Random[T any](slice []T) T
```
**功能**: 从切片中随机选择一个元素

**示例**:
```go
colors := []string{"red", "green", "blue"}
randomColor := candy.Random(colors)
fmt.Println(randomColor) // "red", "green", 或 "blue" 中的一个
```

### 实用工具函数

#### DeepCopy()
```go
func DeepCopy[T any](src T) T
```
**功能**: 深度拷贝任意类型的数据

**特点**: 
- 使用 JSON 序列化/反序列化实现
- 支持嵌套结构体、切片、映射

**示例**:
```go
original := map[string][]int{
    "numbers": {1, 2, 3},
}
copied := candy.DeepCopy(original)
copied["numbers"][0] = 999
fmt.Println(original["numbers"][0]) // 1 (原始数据未改变)
```

#### Chunk()
```go
func Chunk[T any](slice []T, size int) [][]T
```
**功能**: 将切片分割成指定大小的子切片

**示例**:
```go
numbers := []int{1, 2, 3, 4, 5, 6, 7}
chunks := candy.Chunk(numbers, 3)
fmt.Println(chunks) // [[1, 2, 3], [4, 5, 6], [7]]
```

## 🔧 高级用法

### 管道式操作
```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

result := candy.Filter(numbers, func(n int) bool {
    return n%2 == 0  // 过滤偶数
})

result = candy.Map(result, func(n int) int {
    return n * n     // 平方
})

sum := candy.Sum(result)
fmt.Println(sum) // 220 (2²+4²+6²+8²+10²)
```

### 复杂类型转换
```go
// 结构体转映射
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

user := User{Name: "Alice", Age: 30}
userMap := candy.ToMapStringAny(user)
fmt.Println(userMap) // map[name:Alice age:30]
```

## 📊 性能特点

### 性能优化
- **零分配转换**: 基础类型转换实现零内存分配
- **泛型优化**: 使用 Go 1.18+ 泛型，消除反射开销
- **类型安全**: 编译时类型检查，避免运行时错误

### 基准测试结果
| 操作 | 时间复杂度 | 内存分配 | 适用场景 |
|------|------------|----------|----------|
| `ToBool()` | O(1) | 0 allocs | 高频转换 |
| `ToString()` | O(1) - O(n) | 最小化 | 通用转换 |
| `Filter()` | O(n) | 1 alloc | 数据过滤 |
| `Sort()` | O(n log n) | 0 allocs | 原地排序 |

## 🚨 使用注意事项

### 类型转换
1. **精度丢失**: 浮点数转整数会丢失小数部分
2. **溢出风险**: 大数值转换为小类型可能溢出
3. **nil 处理**: nil 指针转换为对应类型的零值

### 性能考虑
1. **复杂类型**: 结构体转换使用 JSON，性能相对较低
2. **大数据**: 对于大切片操作，考虑使用并发版本
3. **内存使用**: 深拷贝操作会复制所有数据

## 💡 最佳实践

### 1. 类型安全
```go
// 推荐：使用类型断言检查
if val, ok := v.(string); ok {
    result := candy.ToBool(val)
}

// 或者使用 candy 的安全转换
result := candy.ToBool(v) // 内部处理所有类型
```

### 2. 性能优化
```go
// 对于已知类型，直接使用
if str, ok := v.(string); ok {
    // 直接处理字符串，避免通用转换
    return str != ""
}

// 对于未知类型，使用 candy
return candy.ToBool(v)
```

### 3. 错误处理
```go
// 检查转换结果的合理性
if result := candy.ToInt64(userInput); result < 0 {
    // 处理异常情况
}
```

## 🔗 相关模块

- **[stringx](../stringx/)**: 字符串专用操作
- **[anyx](../anyx/)**: Any 类型处理
- **[json](../json/)**: JSON 序列化增强

## 📚 更多示例

查看 [examples](./examples/) 目录获取更多实用示例和用法。