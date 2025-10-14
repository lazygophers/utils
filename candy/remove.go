package candy

import (
	"cmp"
	"reflect"
)

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
