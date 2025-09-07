// Package candy 提供 Go 语法糖工具函数，简化常见的编程操作
package candy

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
	// 空切片检查，返回空切片而非 nil
	if len(ss) == 0 {
		return []T{}
	}

	// 创建映射用于记录已出现的键值
	m := make(map[any]struct{})

	// 遍历输入切片
	for _, s := range ss {
		// 使用映射函数提取键值
		key := f(s)

		// 如果键值未出现过，则添加到结果切片
		if _, ok := m[key]; !ok {
			m[key] = struct{}{}
			ret = append(ret, s)
		}
	}

	return ret
}
