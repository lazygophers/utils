# candy - 类型转换工具

`candy` 包为 Go 提供了全面的类型转换工具。它提供了不同数据类型之间的安全、灵活转换，对数字、字符串和布尔值转换提供了广泛支持。

## 功能特性

- **类型转换**: 在不同原始类型之间安全转换
- **集合操作**: 使用函数式编程模式处理切片和数组
- **布尔转换**: 支持多种格式的灵活字符串到布尔值转换
- **数字转换**: 带溢出保护的安全数字转换
- **字符串工具**: 字符串操作和转换函数
- **数组操作**: 处理数组和切片的函数

## 安装

```bash
go get github.com/lazygophers/utils/candy
```

## 使用示例

### 布尔值转换

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/candy"
)

func main() {
    // 各种布尔值转换
    fmt.Println(candy.ToBool("true"))   // true
    fmt.Println(candy.ToBool("yes"))    // true
    fmt.Println(candy.ToBool("1"))      // true
    fmt.Println(candy.ToBool("on"))     // true
    fmt.Println(candy.ToBool("false"))  // false
    fmt.Println(candy.ToBool("no"))     // false
    fmt.Println(candy.ToBool("0"))      // false
    fmt.Println(candy.ToBool("off"))    // false
    fmt.Println(candy.ToBool(42))       // true (非零)
    fmt.Println(candy.ToBool(0))        // false (零)
    fmt.Println(candy.ToBool("hello"))  // true (非空字符串)
    fmt.Println(candy.ToBool(""))       // false (空字符串)
}
```

### 切片操作

```go
// 检查是否有任何元素满足条件
numbers := []int{1, 2, 3, 4, 5}
hasEven := candy.Any(numbers, func(n int) bool { return n%2 == 0 })
fmt.Println(hasEven) // true

// 检查是否所有元素都满足条件
allPositive := candy.All(numbers, func(n int) bool { return n > 0 })
fmt.Println(allPositive) // true

// 计算平均值
average := candy.Average(numbers)
fmt.Println(average) // 3.0

// 获取绝对值
negatives := []int{-1, -2, 3, -4}
absolutes := candy.Abs(negatives)
fmt.Println(absolutes) // [1, 2, 3, 4]
```

### 数组转换

```go
// 将切片转换为字符串数组
items := []interface{}{"apple", 42, true}
stringArray := candy.ToArrayString(items)
fmt.Println(stringArray) // ["apple", "42", "true"]
```

## API 参考

### 布尔值转换

- `ToBool(val interface{}) bool` - 将任何值转换为布尔值

布尔值转换规则：
- **数字**: 0 = false，非零 = true
- **字符串**: "true", "1", "t", "y", "yes", "on" = true；"false", "0", "f", "n", "no", "off" = false
- **其他字符串**: 非空 = true，空 = false
- **nil**: false
- **其他类型**: false

### 切片操作

- `Any[T any](ss []T, f func(T) bool) bool` - 检查是否有任何元素满足条件
- `All[T any](ss []T, f func(T) bool) bool` - 检查是否所有元素都满足条件
- `Average[T numeric](ss []T) float64` - 计算数字切片的平均值
- `Sum[T numeric](ss []T) T` - 计算数字切片的总和
- `Min[T comparable](ss []T) T` - 找到最小值
- `Max[T comparable](ss []T) T` - 找到最大值

### 数组工具

- `Abs[T numeric](ss []T) []T` - 获取所有元素的绝对值
- `ToArrayString(val interface{}) []string` - 将任何切片转换为字符串数组
- `Bottom[T any](ss []T, n int) []T` - 获取底部 N 个元素
- `Top[T any](ss []T, n int) []T` - 获取顶部 N 个元素
- `Unique[T comparable](ss []T) []T` - 移除重复元素
- `Reverse[T any](ss []T) []T` - 反转切片顺序

### 类型转换

- `ToString(val interface{}) string` - 将任何值转换为字符串
- `ToInt(val interface{}) (int, error)` - 转换为 int 并处理错误
- `ToInt64(val interface{}) (int64, error)` - 转换为 int64 并处理错误
- `ToFloat64(val interface{}) (float64, error)` - 转换为 float64 并处理错误
- `ToBytes(val interface{}) []byte` - 转换为字节切片

### 数学操作

- `Abs[T numeric](val T) T` - 获取绝对值
- `Max[T comparable](a, b T) T` - 获取两个值的最大值
- `Min[T comparable](a, b T) T` - 获取两个值的最小值
- `Clamp[T comparable](val, min, max T) T` - 将值限制在最小值和最大值之间

## 类型约束

该包使用 Go 泛型和以下类型约束：

```go
type numeric interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
    ~float32 | ~float64
}

type comparable interface {
    comparable
}
```

## 性能

该包针对性能进行了优化：
- 尽可能实现零分配操作
- 类型特定实现以避免反射
- 使用预分配容量的高效切片操作
- 泛型实现提供类型安全且无运行时开销

## 错误处理

大多数转换函数使用无 panic 模式，返回合理的默认值：
- 布尔转换对无法识别的类型默认为 `false`
- 数字转换对无效输入返回零值
- 字符串转换总是成功，返回字符串表示

对于需要错误信息的情况，使用返回错误的变体：
- `ToIntE()`, `ToInt64E()`, `ToFloat64E()` 等

## 示例

### 数据处理管道

```go
// 处理混合数据切片
data := []interface{}{1, "2", 3.5, "4", true}

// 将所有数据转换为数字，过滤正数，获取平均值
numbers := candy.Map(data, func(v interface{}) float64 {
    if f, err := candy.ToFloat64(v); err == nil {
        return f
    }
    return 0
})

positives := candy.Filter(numbers, func(n float64) bool {
    return n > 0
})

average := candy.Average(positives)
fmt.Printf("正数的平均值: %.2f\n", average)
```

### 配置解析

```go
// 解析配置值
config := map[string]interface{}{
    "debug":     "true",
    "port":      "8080",
    "timeout":   "30s",
    "enabled":   1,
}

debug := candy.ToBool(config["debug"])       // true
port := candy.ToString(config["port"])       // "8080"
enabled := candy.ToBool(config["enabled"])   // true
```

## 相关包

- `anyx` - 用于处理 interface{} 的泛型类型工具
- `stringx` - 高级字符串操作工具
- `validator` - 结构体验证工具