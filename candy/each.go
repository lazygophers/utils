// Package candy 包含实用的语法糖工具函数
package candy

import (
	"reflect"
)

// Each 遍历切片、数组或映射的每个元素，并对每个元素执行指定的函数
//
// 参数:
//   - collection: 要遍历的集合，支持切片、数组、映射
//   - fn: 对每个元素执行的函数，接收元素的索引和值
//
// 示例:
//
//	// 遍历切片
//	Each([]int{1, 2, 3}, func(index int, value int) {
//	    fmt.Printf("索引 %d: 值 %d\n", index, value)
//	})
//
//	// 遍历映射
//	Each(map[string]int{"a": 1, "b": 2}, func(index int, value int) {
//	    fmt.Printf("索引 %d: 值 %d\n", index, value)
//	})
//
// 注意:
//   - 如果 collection 不是切片、数组或映射，函数会 panic
//   - 对于映射，index 参数是连续的整数，从 0 开始
//   - 函数使用反射机制，性能上不如直接的 for 循环
func Each(collection interface{}, fn func(index int, value interface{})) {
	val := reflect.ValueOf(collection)

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			fn(i, val.Index(i).Interface())
		}
	case reflect.Map:
		keys := val.MapKeys()
		for i, key := range keys {
			fn(i, val.MapIndex(key).Interface())
		}
	default:
		panic("Each: collection must be a slice, array, or map")
	}
}
