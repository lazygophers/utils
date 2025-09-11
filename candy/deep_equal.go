package candy

import "reflect"

// deepValueEqual 是 DeepEqual 的内部实现核心。
// 它接收两个 reflect.Value，并递归地对它们进行深度比较。
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

	// 根据值的类型进行分类比较
	switch v1.Kind() {
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
		// 如果指针相同，则内容必然相同
		if v1.UnsafePointer() == v2.UnsafePointer() {
			return true
		}
		// 递归比较每一个键值对
		for _, k := range v1.MapKeys() {
			val1 := v1.MapIndex(k)
			val2 := v2.MapIndex(k)
			if !val1.IsValid() || !val2.IsValid() || !deepValueEqual(val1, val2) {
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
		// 如果指针相同，则内容必然相同
		if v1.UnsafePointer() == v2.UnsafePointer() {
			return true
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
		// 对于不可比较的类型（如函数、映射、切片），需要捕获panic
		var result bool
		var panicked bool
		
		func() {
			defer func() {
				if recover() != nil {
					panicked = true
				}
			}()
			
			// 尝试直接比较，如果类型不可比较会触发panic被上面捕获
			result = v1.Interface() == v2.Interface()
		}()
		
		// 如果发生了panic，说明类型不可比较，返回false
		if panicked {
			return false
		}
		
		return result
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
	v1 := reflect.ValueOf(x)
	v2 := reflect.ValueOf(y)
	return deepValueEqual(v1, v2)
}