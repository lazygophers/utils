# anyx
通用类型转换与数据操作工具库

## 功能概述

anyx 提供了丰富的类型转换和数据操作功能，包括：

- **基础类型转换**：支持 bool、int、float、string 等各种类型间的安全转换
- **复合数据结构操作**：提供 slice、map、pointer 等复杂数据结构的操作工具
- **反射工具**：基于反射的通用数据处理函数
- **JSON 支持**：内置 JSON 序列化/反序列化能力

## 核心功能

### 类型转换函数

#### 布尔值转换
- [`ToBool(val)`](bool.go:47) - 将任意类型转换为布尔值
  - 支持：bool、所有整型、浮点型、string、[]byte
  - 字符串支持："true"/"false"、"1"/"0"、"t"/"f"、"y"/"n"、"yes"/"no"、"on"/"off"

#### 整型转换
- [`ToInt(val)`](int.go:8) - 转换为 int
- [`ToInt8(val)`](int.go:56) - 转换为 int8
- [`ToInt16(val)`](int.go:104) - 转换为 int16
- [`ToInt32(val)`](int.go:152) - 转换为 int32
- [`ToInt64(val)`](int.go:200) - 转换为 int64（支持 time.Duration）
- [`ToInt64Slice(val)`](int.go:250) - 将切片转换为 []int64

#### 字符串转换
- [`ToString(val)`](string.go:16) - 将任意类型转换为字符串
- [`ToBytes(val)`](string.go:76) - 将任意类型转换为 []byte
- [`ToStringSlice(val, seqs...)`](string.go:158) - 将任意类型转换为字符串切片
- [`ToArrayString(val)`](string.go:144) - 将数组/切片转换为字符串切片

#### 浮点型转换
- [`ToFloat(val)`](float.go:8) - 转换为 float64
- [`ToFloat32(val)`](float.go:49) - 转换为 float32
- [`ToFloat64(val)`](float.go:93) - 转换为 float64
- [`ToFloat64Slice(val)`](float.go:142) - 将切片转换为 []float64

#### 无符号整型转换
- [`ToUint(val)`](uint.go:8) - 转换为 uint
- [`ToUint8(val)`](uint.go:56) - 转换为 uint8
- [`ToUint16(val)`](uint.go:104) - 转换为 uint16
- [`ToUint32(val)`](uint.go:152) - 转换为 uint32
- [`ToUint64(val)`](uint.go:200) - 转换为 uint64

### 数据结构操作

#### 切片操作
- [`PluckInt(list, fieldName)`](slice.go:76) - 从结构体切片中提取 int 字段
- [`PluckInt32(list, fieldName)`](slice.go:80) - 从结构体切片中提取 int32 字段
- [`PluckInt64(list, fieldName)`](slice.go:88) - 从结构体切片中提取 int64 字段
- [`PluckUint32(list, fieldName)`](slice.go:84) - 从结构体切片中提取 uint32 字段
- [`PluckUint64(list, fieldName)`](slice.go:92) - 从结构体切片中提取 uint64 字段
- [`PluckString(list, fieldName)`](slice.go:96) - 从结构体切片中提取 string 字段
- [`PluckStringSlice(list, fieldName)`](slice.go:100) - 从结构体切片中提取 []string 字段
- [`DiffSlice(a, b)`](slice.go:104) - 比较两个切片的差异，返回 (a中有b无的元素, b中有a无的元素)
- [`RemoveSlice(src, rm)`](slice.go:147) - 从源切片中移除指定元素

#### 映射操作
- [`CheckValueType(val)`](map.go:21) - 检查值的类型（数字、字符串、布尔）
- [`MapKeysString(m)`](map.go:36) - 提取 map[string]V 的所有键
- [`MapKeysUint32(m)`](map.go:58) - 提取 map[uint32]V 的所有键
- [`MapKeysUint64(m)`](map.go:80) - 提取 map[uint64]V 的所有键
- [`MapKeysInt32(m)`](map.go:99) - 提取 map[int32]V 的所有键
- [`MapKeysInt64(m)`](map.go:118) - 提取 map[int64]V 的所有键
- [`MapValues[K, V](m)`](map.go:137) - 提取 map 的所有值（泛型函数）
- [`MergeMap[K, V](source, target)`](map.go:145) - 合并两个 map
- [`KeyBy(list, fieldName)`](map.go:157) - 将结构体切片转换为以指定字段为键的 map
- [`KeyByUint64[M](list, fieldName)`](map.go:203) - 将结构体指针切片转换为 map[uint64]*M
- [`KeyByInt64[M](list, fieldName)`](map.go:243) - 将结构体指针切片转换为 map[int64]*M
- [`KeyByString[M](list, fieldName)`](map.go:283) - 将结构体指针切片转换为 map[string]*M
- [`Slice2Map[M](list)`](map.go:323) - 将切片转换为 map[M]bool
- [`ToMap(v)`](map.go:350) - 将任意类型转换为 map[string]interface{}
- [`ToMapStringAny(v)`](map.go:333) - 将任意类型转换为 map[string]interface{}
- [`ToMapStringString(v)`](map.go:371) - 将任意类型转换为 map[string]string
- [`ToMapStringInt64(v)`](map.go:388) - 将任意类型转换为 map[string]int64
- [`ToMapInt64String(v)`](map.go:405) - 将任意类型转换为 map[int64]string
- [`ToMapInt32String(v)`](map.go:422) - 将任意类型转换为 map[int32]string
- [`ToMapStringArrayString(v)`](map.go:439) - 将任意类型转换为 map[string][]string

#### 深度复制
- [`DeepCopy(src, dst)`](deep.go:12) - 深度复制结构体
- [`DeepClone(src)`](deep.go:67) - 深度克隆任意类型

#### 指针操作
- [`ToPointer[T](v)`](ptr.go:13) - 将值转换为指针（泛型函数）
- [`ToPointerSlice[T](v)`](ptr.go:28) - 将切片转换为指针切片（泛型函数）
- [`ValueOfPointer[T](p)`](ptr.go:43) - 获取指针指向的值（泛型函数）
- [`ValueOfPointerSlice[T](p)`](ptr.go:58) - 获取指针切片的值切片（泛型函数）

## 注释规范

### GoDoc风格要求
**所有公共API必须遵循GoDoc注释规范**：
- 包注释：使用`// Package anyx`开头描述包功能
- 函数注释：包含功能描述、参数说明和返回值约定
- 示例注释：使用`// Example`标签提供完整用例
- 重要：注释应使用中文描述，但保留英文标签

### interface{}类型处理
当使用`interface{}`参数时必须：
1. 在注释中明确支持的类型列表
2. 使用类型断言前进行类型检查
3. 推荐使用类型开关处理多态
4. 错误路径必须返回具体错误类型

### 反射操作性能
反射操作需注意：
- 避免在热点代码中频繁使用
- 重复使用的反射结果应通过sync.Pool缓存
- 复杂结构遍历应使用深度优先算法
- 性能敏感场景建议预分配内存池

### 错误处理约定
统一错误处理规范：
- 所有函数返回`error`类型
- 错误信息使用`log.Errorf`记录
- 通用错误使用`fmt.Errorf`包装
- 业务错误使用自定义错误类型
- 必须处理所有非nil错误返回
- 测试用例需覆盖错误分支

## 使用示例
```
```go
// 基础类型转换
s := anyx.ToString(42) // "42"
i := anyx.ToInt("123") // 123

// 深度复制对象
a := struct{ X int }{42}
b := struct{ X int }{}
anyx.DeepCopy(&a, &b)

// 嵌套路径访问
m := anyx.NewMap(map[string]interface{}{
    "user": map[string]interface{}{
        "id": 123,
    },
})
id := m.GetInt("user.id") // 123
```
