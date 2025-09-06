package candy

// UniqueUsing 使用自定义函数提取键值对切片进行去重
// 该函数通过提供的映射函数 f 从每个元素中提取键值，然后基于这些键值进行去重
// 保留原始切片中第一次出现的元素顺序
//
// 参数:
//   ss: 输入切片，支持任意类型
//   f: 映射函数，用于从元素中提取比较键值
//
// 返回:
//   去重后的切片，保持原始顺序
//
// 示例:
//   // 对整型切片去重
//   result := UniqueUsing([]int{1, 2, 3, 2, 1}, func(n int) any { return n })
//   // result 为 []int{1, 2, 3}
//
//   // 按结构体字段去重
//   type Person struct { Name string; Age int }
//   people := []Person{{"Alice", 25}, {"Bob", 30}, {"Alice", 35}}
//   result := UniqueUsing(people, func(p Person) any { return p.Name })
//   // result 保留第一个 Alice
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