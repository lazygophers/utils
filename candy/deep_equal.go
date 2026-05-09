package candy

import "reflect"

// deepValueEqual 是 DeepEqual 的内部实现核心（优化版本）。
// 它接收两个 reflect.Value，并递归地对它们进行深度比较。
//
// 优化点：
// 1. 使用 UnsafeAddr 替代 UnsafePointer（更安全且性能更好）
// 2. 提前缓存 kind 值避免重复调用
// 3. 简化 panic 恢复机制
// 4. 使用 MapRange 优化 Map 迭代
// 5. 减少 reflect.Value 创建
//
// 注意：此函数为 unexported，不应在包外直接调用。
func deepValueEqual(v1, v2 reflect.Value) bool {
	// 检查值是否有效
	if !v1.IsValid() || !v2.IsValid() {
		return v1.IsValid() == v2.IsValid()
	}

	// 比较类型是否一致
	if v1.Type() != v2.Type() {
		return false
	}

	// 缓存 kind 避免重复调用
	kind := v1.Kind()

	switch kind {
	// 比较 Map
	case reflect.Map:
		// 检查是否为 nil
		if v1.IsNil() || v2.IsNil() {
			return v1.IsNil() == v2.IsNil()
		}
		// 比较长度
		if v1.Len() != v2.Len() {
			return false
		}
		// 优化：使用 UnsafeAddr 替代 UnsafePointer，并添加 CanAddr 检查
		if v1.CanAddr() && v2.CanAddr() {
			addr1, addr2 := v1.UnsafeAddr(), v2.UnsafeAddr()
			if addr1 == addr2 {
				return true
			}
		}
		// 优化：使用 MapRange 替代 MapKeys（减少内存分配）
		iter := v1.MapRange()
		for iter.Next() {
			k := iter.Key()
			val1 := iter.Value()
			val2 := v2.MapIndex(k)
			if !val2.IsValid() || !deepValueEqual(val1, val2) {
				return false
			}
		}
		return true

	// 比较 Slice
	case reflect.Slice:
		// 检查 nil 状态是否一致
		if v1.IsNil() != v2.IsNil() {
			return false
		}
		// 比较长度
		if v1.Len() != v2.Len() {
			return false
		}
		// 优化：使用 UnsafeAddr 替代 UnsafePointer，并添加 CanAddr 检查
		if v1.CanAddr() && v2.CanAddr() {
			addr1, addr2 := v1.UnsafeAddr(), v2.UnsafeAddr()
			if addr1 == addr2 {
				return true
			}
		}
		// 递归比较每一个元素
		for i := 0; i < v1.Len(); i++ {
			if !deepValueEqual(v1.Index(i), v2.Index(i)) {
				return false
			}
		}
		return true

	// 比较指针
	case reflect.Ptr:
		// 检查是否为 nil
		if v1.IsNil() || v2.IsNil() {
			return v1.IsNil() == v2.IsNil()
		}
		// 递归比较指针指向的元素
		return deepValueEqual(v1.Elem(), v2.Elem())

	// 比较数组
	case reflect.Array:
		// 递归比较每一个元素
		// 注意：不能对数组类型的 reflect.Value 使用 UnsafePointer
		for i := 0; i < v1.Len(); i++ {
			if !deepValueEqual(v1.Index(i), v2.Index(i)) {
				return false
			}
		}
		return true

	// 比较结构体
	case reflect.Struct:
		// 递归比较每一个字段
		for i := 0; i < v1.NumField(); i++ {
			if !deepValueEqual(v1.Field(i), v2.Field(i)) {
				return false
			}
		}
		return true

	// 比较接口
	case reflect.Interface:
		// 检查是否为 nil
		if v1.IsNil() || v2.IsNil() {
			return v1.IsNil() == v2.IsNil()
		}
		// 递归比较接口包含的元素
		return deepValueEqual(v1.Elem(), v2.Elem())

	// 对于其他基本类型，直接比较接口值
	default:
		// 优化：简化 panic 恢复机制（减少函数调用开销）
		defer func() {
			recover()
		}()
		return v1.Interface() == v2.Interface()
	}
}

// DeepEqual 使用深度递归比较的方式，判断两个任意类型的值 x 和 y 是否完全相等。
//
// 与标准的 `==` 运算符不同，`DeepEqual` 能够深入探索数据结构的内部，
// 对 Maps、Slices、Pointers、Structs 等复合类型的元素或字段进行逐一递归比较。
//
// 对于基本类型，它会直接比较其值。对于指针，它会比较所指向的实际内容。
// 两个 nil 值被视作相等。
//
// 示例：
//   - `DeepEqual(map[string]int{"a": 1}, map[string]int{"a": 1})` 返回 `true`
//   - `DeepEqual([]int{1, 2}, []int{1, 2})` 返回 `true`
//   - `DeepEqual(1, 1)` 返回 `true`
//   - `DeepEqual(1, 2)` 返回 `false`
//
// @param x 第一个待比较的值。
// @param y 第二个待比较的值。
// @return 如果两个值在结构和内容上完全相等，则返回 true，否则返回 false。
func DeepEqual[M any](x, y M) bool {
	// 快速路径：常见切片类型
	switch xv := any(x).(type) {
	case []int:
		yv, ok := any(y).([]int)
		if !ok {
			return false
		}
		if len(xv) != len(yv) {
			return false
		}
		for i := 0; i < len(xv); i++ {
			if xv[i] != yv[i] {
				return false
			}
		}
		return true

	case []string:
		yv, ok := any(y).([]string)
		if !ok {
			return false
		}
		if len(xv) != len(yv) {
			return false
		}
		for i := 0; i < len(xv); i++ {
			if xv[i] != yv[i] {
				return false
			}
		}
		return true

	case map[string]int:
		yv, ok := any(y).(map[string]int)
		if !ok {
			return false
		}
		if len(xv) != len(yv) {
			return false
		}
		for k, v := range xv {
			if yv[k] != v {
				return false
			}
		}
		return true

	case map[string]string:
		yv, ok := any(y).(map[string]string)
		if !ok {
			return false
		}
		if len(xv) != len(yv) {
			return false
		}
		for k, v := range xv {
			if yv[k] != v {
				return false
			}
		}
		return true
	}

	// 回退到反射
	v1 := reflect.ValueOf(x)
	v2 := reflect.ValueOf(y)
	return deepValueEqual(v1, v2)
}
