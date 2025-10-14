package candy

import (
	"cmp"
	"reflect"
)

// Spare 返回在 against 中但不在 ss 中的元素
// 该函数实现了集合差集操作，返回在 against 切片中存在但在 ss 切片中不存在的所有元素
// 注意：该函数与 Remove 函数功能相同，都是返回差集结果
//
// 参数：
//   - ss: 作为参考集合的切片
//   - against: 作为被比较集合的切片
//
// 返回：
//   - []T: 在 against 中但不在 ss 中的元素组成的切片
//
// 示例：
//
//	ss := []int{1, 2, 3}
//	against := []int{2, 3, 4, 5}
//	result := Spare(ss, against) // 返回 [4, 5]
func Spare[T cmp.Ordered](ss []T, against []T) (result []T) {
	// 使用 make 初始化，确保返回空切片而非 nil
	result = make([]T, 0)
	set := make(map[T]struct{}, len(ss))

	// 将 ss 中的所有元素添加到 map 中用于快速查找
	for _, s := range ss {
		set[s] = struct{}{}
	}

	// 遍历 against 切片，找出不在 ss 中的元素
	for _, s := range against {
		if _, ok := set[s]; !ok {
			result = append(result, s)
		}
	}
	return
}

// Remove 从第一个切片中移除第二个切片中存在的元素
// 返回第一个切片中但不在第二个切片中的元素
//
// 参数:
//   - ss: 源切片
//   - toRemove: 要移除的元素列表
//
// 返回值:
//   - result: ss 中移除了 toRemove 中存在元素后的结果
//
// 示例:
//
//	ss := []int{1, 2, 3, 4, 5}
//	toRemove := []int{2, 4, 6}
//	result := Remove(ss, toRemove)
//	// result = [1, 3, 5]
func Remove[T cmp.Ordered](ss []T, toRemove []T) (result []T) {
	// 使用 make 初始化，确保返回空切片而非 nil
	result = make([]T, 0)
	removeSet := make(map[T]struct{}, len(toRemove))

	// 构建要移除元素的集合
	for _, item := range toRemove {
		removeSet[item] = struct{}{}
	}

	// 遍历原切片，只保留不在移除集合中的元素
	for _, item := range ss {
		if _, shouldRemove := removeSet[item]; !shouldRemove {
			result = append(result, item)
		}
	}
	return
}

// RemoveIndex 移除指定索引的元素
// 该函数从切片中移除指定索引位置的元素，并返回新的切片
// 如果索引无效（超出范围或为负数），则返回空切片
func RemoveIndex[T any](ss []T, index int) []T {
	// 边界检查：如果切片为空或索引无效，返回空切片
	if len(ss) == 0 || index < 0 || index >= len(ss) {
		return make([]T, 0)
	}

	// 处理移除第一个元素的特殊情况
	if index == 0 {
		return ss[1:]
	}

	// 处理移除最后一个元素的特殊情况
	if index == len(ss)-1 {
		return ss[:len(ss)-1]
	}

	// 一般情况：使用 append 将索引前后的元素拼接起来
	return append(ss[:index], ss[index+1:]...)
}

// RemoveSlice 从源切片中移除指定的元素
// src 是源切片，rm 是要移除的元素切片
// 返回移除指定元素后的新切片
func RemoveSlice(src interface{}, rm interface{}) interface{} {
	at := reflect.TypeOf(src)
	if at.Kind() != reflect.Slice {
		panic("a is not slice")
	}

	bt := reflect.TypeOf(rm)
	if bt.Kind() != reflect.Slice {
		panic("b is not slice")
	}

	atm := at.Elem()
	btm := bt.Elem()

	if atm.Kind() != btm.Kind() {
		panic("a and b are not same type")
	}

	m := map[interface{}]bool{}

	bv := reflect.ValueOf(rm)
	for i := 0; i < bv.Len(); i++ {
		m[bv.Index(i).Interface()] = true
	}

	c := reflect.MakeSlice(at, 0, 0)
	av := reflect.ValueOf(src)
	for i := 0; i < av.Len(); i++ {
		if !m[av.Index(i).Interface()] {
			c = reflect.Append(c, av.Index(i))
			delete(m, av.Index(i).Interface())
		}
	}

	return c.Interface()
}

// Diff 计算两个有序切片之间的差异
//
// 参数:
//   - ss: 第一个切片
//   - against: 第二个切片
//
// 返回值:
//   - added: 在 against 中存在但不在 ss 中的元素
//   - removed: 在 ss 中存在但不在 against 中的元素
//
// 示例:
//
//	ss := []int{1, 2, 3}
//	against := []int{2, 3, 4}
//	added, removed := Diff(ss, against)
//	// added = [4]
//	// removed = [1]
func Diff[T cmp.Ordered](ss []T, against []T) (added, removed []T) {
	removed = Remove(ss, against)
	added = Remove(against, ss)

	return
}

// DiffSlice 比较两个切片的差异
// 返回第一个切片中存在但第二个切片中不存在的元素，以及第二个切片中存在但第一个切片中不存在的元素
func DiffSlice(a interface{}, b interface{}) (interface{}, interface{}) {
	at := reflect.TypeOf(a)
	if at.Kind() != reflect.Slice {
		panic("a is not slice")
	}

	bt := reflect.TypeOf(b)
	if bt.Kind() != reflect.Slice {
		panic("b is not slice")
	}

	atm := at.Elem()
	btm := bt.Elem()

	if atm.Kind() != btm.Kind() {
		panic("a and b are not same type")
	}

	m := map[interface{}]reflect.Value{}

	bv := reflect.ValueOf(b)
	for i := 0; i < bv.Len(); i++ {
		m[bv.Index(i).Interface()] = bv.Index(i)
	}

	c := reflect.MakeSlice(at, 0, 0)
	d := reflect.MakeSlice(bt, 0, 0)
	av := reflect.ValueOf(a)
	for i := 0; i < av.Len(); i++ {
		if !m[av.Index(i).Interface()].IsValid() {
			c = reflect.Append(c, av.Index(i))
		} else {
			delete(m, av.Index(i).Interface())
		}
	}

	for _, value := range m {
		d = reflect.Append(d, value)
	}

	return c.Interface(), d.Interface()
}

// First 返回切片中的第一个元素
// 如果切片为空，返回类型的零值
//
// 泛型参数 T 可以是任意类型
//
// 示例：
//
//	nums := []int{1, 2, 3}
//	first := First(nums) // 返回 1
//
//	empty := []string{}
//	first := First(empty) // 返回 "" (string 的零值)
func First[T any](ss []T) (ret T) {
	if len(ss) == 0 {
		return
	}

	return ss[0]
}

// FirstOr 返回切片的第一个元素，如果切片为空则返回指定的默认值
//
// 该函数使用泛型支持任意类型的切片，提供了安全的空切片处理机制。
// 在访问切片第一个元素之前，会先检查切片长度，避免 panic。
//
// 参数:
//   - ss: 任意类型的切片
//   - or: 当切片为空时返回的默认值
//
// 返回:
//   - 切片的第一个元素，如果切片为空则返回默认值
//
// 示例:
//
//	numbers := []int{1, 2, 3}
//	first := FirstOr(numbers, 0)     // 返回 1
//
//	empty := []int{}
//	defaultVal := FirstOr(empty, 0) // 返回 0
func FirstOr[T any](ss []T, or T) (ret T) {
	if len(ss) == 0 {
		return or
	}

	return ss[0]
}

// Last 返回切片中的最后一个元素
// 如果切片为空，返回类型的零值
// 该函数使用泛型实现，支持任意类型的切片
//
// 参数:
//   - ss: 任意类型的切片
//
// 返回:
//   - T: 切片中的最后一个元素，如果切片为空则返回类型零值
//
// 示例:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	last := Last(numbers) // 返回 5
//
//	empty := []string{}
//	result := Last(empty) // 返回 ""
func Last[T any](ss []T) (ret T) {
	if len(ss) == 0 {
		return
	}

	return ss[len(ss)-1]
}

// LastOr 返回切片中的最后一个元素
// 如果切片为空，返回指定的默认值
// 该函数使用泛型实现，支持任意类型的切片
//
// 参数:
//   - ss: 任意类型的切片
//   - or: 当切片为空时返回的默认值
//
// 返回:
//   - T: 切片中的最后一个元素，如果切片为空则返回指定的默认值
//
// 示例:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	last := LastOr(numbers, 0) // 返回 5
//
//	empty := []string{}
//	result := LastOr(empty, "default") // 返回 "default"
func LastOr[T any](ss []T, or T) (ret T) {
	if len(ss) == 0 {
		return or
	}

	return ss[len(ss)-1]
}

// Index 返回元素 sub 在切片 ss 中的索引位置
// 如果未找到，返回 -1
// 这是一个泛型函数，支持所有可排序的类型
func Index[T cmp.Ordered](ss []T, sub T) int {
	if len(ss) == 0 {
		return -1
	}

	for i, s := range ss {
		if s == sub {
			return i
		}
	}

	return -1
}

// Same 返回在 against 和 ss 中都存在的元素
// 用于查找两个有序集合的交集
func Same[T cmp.Ordered](against []T, ss []T) (result []T) {
	// 使用 make 初始化，确保返回空切片而非 nil
	result = make([]T, 0)
	set := make(map[T]struct{}, len(ss))

	for _, s := range ss {
		set[s] = struct{}{}
	}

	for _, s := range against {
		if _, ok := set[s]; ok {
			result = append(result, s)
		}
	}
	return
}

// SliceEqual 判断两个切片是否相等，不考虑元素顺序
// 使用 map 来统计元素出现次数，确保每个元素在两个切片中出现次数相同
// 处理了 nil 切片的特殊情况：nil 和空切片视为相等
func SliceEqual[T any](a, b []T) bool {
	// 处理 nil 切片的情况：nil 和空切片视为相等
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	// 使用 map 来跟踪每个元素的出现次数
	am := make(map[any]int, len(a))
	for _, v := range a {
		am[v]++
	}

	for _, v := range b {
		if count, ok := am[v]; !ok || count == 0 {
			return false
		}
		am[v]--
	}

	// 检查所有元素的计数是否都为0
	for _, count := range am {
		if count != 0 {
			return false
		}
	}

	return true
}

// Drop 丢弃切片前 n 个元素
//
// 参数:
//   - ss: 输入切片，支持任意类型
//   - n: 要丢弃的元素个数
//
// 返回值:
//   - []T: 丢弃前 n 个元素后的新切片
//
// 特点:
//   - 支持任意类型的切片
//   - 如果 n 为负数，当作 0 处理
//   - 如果 n 大于切片长度，返回空切片
//   - 不修改原切片，返回新切片
//
// 示例:
//
//	result := Drop([]int{1, 2, 3, 4, 5}, 2)      // 返回 [3, 4, 5]
//	result := Drop([]string{"a", "b", "c"}, 0)  // 返回 ["a", "b", "c"]
//	result := Drop([]int{1, 2, 3}, 5)      // 返回 []
//	result := Drop([]int{1, 2, 3}, -1)     // 返回 [1, 2, 3]
func Drop[T any](ss []T, n int) []T {
	if n < 0 {
		n = 0
	}

	if n > len(ss) {
		n = len(ss)
	}

	return ss[n:]
}

// Filter 函数用于过滤切片中的元素，返回满足条件的所有元素
//
// 参数：
//   - ss: 要过滤的切片，可以是任意类型的切片
//   - f: 过滤函数，接收一个元素并返回布尔值，true 表示保留该元素
//
// 返回值：
//   - []T: 包含所有满足条件的元素的新切片
//
// 示例：
//
//	numbers := []int{1, 2, 3, 4, 5}
//	evenNumbers := Filter(numbers, func(n int) bool {
//	    return n%2 == 0
//	})
//	// evenNumbers = [2, 4]
//
// 注意事项：
//   - 使用 make 初始化返回切片，确保返回空切片而非 nil
//   - 该函数不会修改原始切片
//   - 时间复杂度为 O(n)，其中 n 为切片长度
func Filter[T any](ss []T, f func(T) bool) []T {
	if len(ss) == 0 {
		return []T{}
	}

	// 使用原始长度的1/4作为初始容量预估，减少重新分配
	ret := make([]T, 0, len(ss)/4+1)
	for _, s := range ss {
		if f(s) {
			ret = append(ret, s)
		}
	}

	return ret
}

// FilterNot 对切片进行反向过滤，保留不满足指定条件的元素
//
// 参数:
//   - ss: 输入切片，支持任意类型
//   - f: 谓词函数，接收一个元素并返回布尔值
//
// 返回值:
//   - []T: 包含所有不满足谓词函数条件的新切片
//
// 特点:
//   - 支持任意类型的切片
//   - 是 Filter 函数的补集，用于语义上更清晰的反向过滤操作
//   - 不修改原切片，返回新切片
//   - 空切片输入返回空切片
//
// 示例:
//
//	// 过滤偶数，保留奇数
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	result := FilterNot(numbers, func(n int) bool {
//	    return n % 2 == 0
//	})
//	// result = [1, 3, 5]
//
//	// 过滤空字符串，保留非空字符串
//	strings := []string{"hello", "", "world", ""}
//	result := FilterNot(strings, func(s string) bool {
//	    return s == ""
//	})
//	// result = ["hello", "world"]
//
//	// 过滤负数，保留非负数
//	nums := []int{-1, 0, 1, -2, 2}
//	result := FilterNot(nums, func(n int) bool {
//	    return n < 0
//	})
//	// result = [0, 1, 2]
func FilterNot[T any](ss []T, f func(T) bool) []T {
	// 使用 make 初始化，确保返回空切片而非 nil
	us := make([]T, 0)
	for _, s := range ss {
		if !f(s) {
			us = append(us, s)
		}
	}
	return us
}

// Contains 检查切片中是否包含指定元素
func Contains[T cmp.Ordered](ss []T, s T) bool {
	return ContainsUsing(ss, func(v T) bool {
		return s == v
	})
}

// ContainsUsing 使用自定义函数检查切片中是否包含满足条件的元素
//
// 泛型约束：T 可以是任意类型
// 参数：
//   - ss: 要搜索的切片
//   - f: 判断函数，如果函数返回 true，则表示找到匹配的元素
//
// 返回值：
//   - bool: 如果找到满足条件的元素返回 true，否则返回 false
//
// 示例：
//
//	numbers := []int{1, 2, 3, 4, 5}
//	found := ContainsUsing(numbers, func(v int) bool {
//	    return v > 3
//	}) // 返回 true
func ContainsUsing[T any](ss []T, f func(v T) bool) bool {
	// 遍历切片中的每个元素
	for _, v := range ss {
		// 使用自定义函数判断当前元素是否满足条件
		if f(v) {
			return true
		}
	}

	// 遍历完所有元素都没有找到满足条件的，返回 false
	return false
}

// Bottom 返回切片的最后 n 个元素
func Bottom[T any](ss []T, n int) (ret []T) {
	if n <= 0 {
		return []T{}
	}
	if n > len(ss) {
		n = len(ss)
	}

	ret = make([]T, n)
	copy(ret, ss[len(ss)-n:])
	return ret
}

// Chunk 将切片分割成指定大小的子切片
func Chunk[T any](ss []T, size int) (ret [][]T) {
	if len(ss) == 0 || size <= 0 {
		return [][]T{}
	}

	for i := 0; i < len(ss); i += size {
		end := i + size
		if end > len(ss) {
			end = len(ss)
		}

		ret = append(ret, ss[i:end])
	}

	return
}
