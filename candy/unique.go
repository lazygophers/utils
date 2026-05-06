// Package candy 提供 Go 语法糖工具函数，简化常见的编程操作
package candy

import (
	"golang.org/x/exp/constraints"
)

// Unique 返回切片中的唯一元素，保持原始顺序
//
// 类型参数：
//   - T: 可排序的类型，支持所有可比较的内置类型
//
// 参数：
//   - ss: 输入切片，包含可排序类型的元素
//
// 返回值：
//   - []T: 去重后的切片，保持原始顺序
//
// 特点：
//   - 保持原始元素的顺序
//   - 使用 map 高效去重，时间复杂度 O(n)
//   - 空切片安全，返回空切片而非 nil
//   - 只支持可排序类型，确保类型安全
//
// 示例：
//
//	// 对整数切片去重
//	numbers := []int{1, 2, 2, 3, 4, 4, 5}
//	unique := Unique(numbers)
//	// unique 的值为 []int{1, 2, 3, 4, 5}
//
//	// 对字符串切片去重
//	names := []string{"Alice", "Bob", "Alice", "Charlie", "Bob"}
//	uniqueNames := Unique(names)
//	// uniqueNames 的值为 []string{"Alice", "Bob", "Charlie"}
//
//	// 对浮点数切片去重
//	floats := []float64{1.1, 2.2, 1.1, 3.3, 2.2}
//	uniqueFloats := Unique(floats)
//	// uniqueFloats 的值为 []float64{1.1, 2.2, 3.3}
func Unique[T constraints.Ordered](ss []T) (ret []T) {
	// 空切片处理
	if len(ss) == 0 {
		return make([]T, 0)
	}

	// 小数据集优化：线性搜索比 map 更快
	const smallThreshold = 32
	if len(ss) < smallThreshold {
		ret = make([]T, 0, len(ss))
		for _, s := range ss {
			found := false
			for _, v := range ret {
				if v == s {
					found = true
					break
				}
			}
			if !found {
				ret = append(ret, s)
			}
		}
		return
	}

	// 大数据集：使用预分配 map 优化性能
	// 估算唯一元素数量，避免过度分配
	estimatedSize := len(ss)
	if estimatedSize > 100 {
		// 假设至少有 30% 的重复率
		estimatedSize = estimatedSize * 70 / 100
	}

	ret = make([]T, 0, estimatedSize)
	m := make(map[T]struct{}, estimatedSize)
	for _, s := range ss {
		if _, ok := m[s]; !ok {
			m[s] = struct{}{}
			ret = append(ret, s)
		}
	}

	return
}

// UniqueUsing 返回切片中的唯一元素，保持原始顺序，使用提供的函数来生成用于比较的键
//
// 类型参数：
//   - T: 任意类型，支持所有可比较的类型
//
// 参数：
//   - ss: 输入切片，包含任意类型的元素
//   - f: 键生成函数，用于从元素中提取比较键
//
// 返回值：
//   - []T: 去重后的切片，保持原始顺序
//
// 特点：
//   - 保持原始元素的顺序
//   - 使用自定义键生成函数，支持复杂类型去重
//   - 空切片安全，返回空切片而非 nil
//   - 适用于结构体、自定义类型等复杂类型
//   - 时间复杂度 O(n)，空间复杂度 O(n)
//
// 示例：
//
//	// 对结构体切片去重（按 ID）
//	type User struct {
//	    ID   int
//	    Name string
//	}
//	users := []User{{1, "Alice"}, {2, "Bob"}, {1, "Alice2"}}
//	uniqueUsers := UniqueUsing(users, func(u User) any { return u.ID })
//	// uniqueUsers 的值为 []User{{1, "Alice"}, {2, "Bob"}}
//
//	// 对字符串切片按长度去重
//	words := []string{"apple", "banana", "orange", "kiwi"}
//	uniqueLengths := UniqueUsing(words, func(s string) any { return len(s) })
//	// uniqueLengths 的值为 []string{"apple", "banana", "orange"}
//
//	// 对切片按首字母去重
//	names := []string{"Alice", "Bob", "Anna", "Charlie", "Bob"}
//	uniqueFirstLetters := UniqueUsing(names, func(s string) any { return s[0] })
//	// uniqueFirstLetters 的值为 []string{"Alice", "Bob", "Charlie"}
func UniqueUsing[T any](ss []T, f func(T) any) (ret []T) {
	// 优化版本：
	// 1. 预分配结果切片容量，减少append时的内存重新分配
	// 2. 预分配map容量，减少map扩容开销
	// 3. 使用索引循环避免range的值拷贝开销
	// 4. 预计算长度避免重复调用
	n := len(ss)
	if n == 0 {
		return []T{}
	}

	ret = make([]T, 0, n)
	m := make(map[any]struct{}, n)

	for i := 0; i < n; i++ {
		key := f(ss[i])
		if _, ok := m[key]; !ok {
			m[key] = struct{}{}
			ret = append(ret, ss[i])
		}
	}

	return ret
}
