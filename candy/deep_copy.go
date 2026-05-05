package candy

import (
	"reflect"

	"github.com/lazygophers/log"
)

// deepCopyValue 是 DeepCopy 的内部实现核心。
// 它接收两个 reflect.Value (v1 为源, v2 为目标)，并递归地将内容从 v1 拷贝到 v2。
//
// 注意：此函数为 unexported，不应在包外直接调用。
func deepCopyValue(v1, v2 reflect.Value) {
	if !v1.IsValid() || !v2.IsValid() {
		return
	}

	// 解引用指针，直到获取到实际的值
	for v1.Kind() == reflect.Ptr {
		// 如果源指针是 nil，则无需继续
		if v1.IsNil() {
			return
		}
		v1 = v1.Elem()
	}
	for v2.Kind() == reflect.Ptr {
		// 如果目标指针是 nil，则为其分配新内存
		if v2.IsNil() {
			v2.Set(reflect.New(v2.Type().Elem()))
		}
		v2 = v2.Elem()
	}

	// 类型必须匹配才能拷贝
	if v1.Type() != v2.Type() {
		log.Panicf("源类型 %s 与目标类型 %s 不匹配", v1.Type(), v2.Type())
	}

	switch v1.Kind() {
	// 拷贝 Map
	case reflect.Map:
		if v1.IsNil() {
			v2.Set(reflect.Zero(v2.Type()))
			return
		}
		// 为目标 Map 创建实例
		v2.Set(reflect.MakeMap(v1.Type()))
		// 遍历 Map 并递归拷贝每一个键值对
		for _, k := range v1.MapKeys() {
			val1 := v1.MapIndex(k)
			val2 := reflect.New(val1.Type()).Elem()
			deepCopyValue(val1, val2)
			v2.SetMapIndex(k, val2)
		}

	// 拷贝 Slice
	case reflect.Slice:
		if v1.IsNil() {
			v2.Set(reflect.Zero(v2.Type()))
			return
		}
		// 为目标 Slice 创建实例
		v2.Set(reflect.MakeSlice(v1.Type(), v1.Len(), v1.Cap()))
		// 遍历 Slice 并递归拷贝每一个元素
		for i := 0; i < v1.Len(); i++ {
			deepCopyValue(v1.Index(i), v2.Index(i))
		}

	// 拷贝 Array
	case reflect.Array:
		// 遍历 Array 并递归拷贝每一个元素
		for i := 0; i < v1.Len(); i++ {
			deepCopyValue(v1.Index(i), v2.Index(i))
		}

	// 拷贝 Struct
	case reflect.Struct:
		// 遍历 Struct 并递归拷贝每一个字段
		for i := 0; i < v1.NumField(); i++ {
			srcField := v1.Field(i)
			dstField := v2.Field(i)

			// 跳过不可设置的字段（通常是未导出字段）
			if !dstField.CanSet() {
				// 如果无法访问，跳过该字段
				continue
			}

			deepCopyValue(srcField, dstField)
		}

	// 拷贝 Interface
	case reflect.Interface:
		if v1.IsNil() {
			return
		}
		// 获取接口的实际值
		srcElem := v1.Elem()
		// 创建一个新的目标值，类型与源相同
		dstElem := reflect.New(srcElem.Type()).Elem()
		// 递归拷贝
		deepCopyValue(srcElem, dstElem)
		// 将拷贝后的值设置给目标接口
		v2.Set(dstElem)

	// 拷贝基本类型
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v2.CanSet() {
			v2.SetInt(v1.Int())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if v2.CanSet() {
			v2.SetUint(v1.Uint())
		}
	case reflect.Float32, reflect.Float64:
		if v2.CanSet() {
			v2.SetFloat(v1.Float())
		}
	case reflect.Complex64, reflect.Complex128:
		if v2.CanSet() {
			v2.SetComplex(v1.Complex())
		}
	case reflect.String:
		if v2.CanSet() {
			v2.SetString(v1.String())
		}
	case reflect.Bool:
		if v2.CanSet() {
			v2.SetBool(v1.Bool())
		}

	default:
		// 对于未处理的类型，直接 panic
		log.Panicf("未处理的类型: %s", v1.Kind())
	}
}

// isBasicType 判断是否为基本类型，用于快速路径优化
func isBasicType(kind reflect.Kind) bool {
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.String, reflect.Bool:
		return true
	default:
		return false
	}
}

// DeepCopy 通过深度递归的方式，将源对象 `src` 的内容完全复制到目标对象 `dst`。
//
// 此函数会创建一个源对象的完整、独立的副本。修改副本不会对原始对象产生任何影响。
// 它能够处理 Maps、Slices、Pointers、Structs 等各种复杂类型。
//
// **重要提示**:
// 参数 `dst` **必须**是一个指向目标对象的指针，且该指针必须已经被初始化（例如，通过 `new` 或 `&`）。
// 如果 `dst` 是一个 nil 指针或者不是指针类型，函数将在运行时引发 `panic`，因为无法向无效的内存地址写入数据。
//
// 示例：
//
//	var src = map[string]int{"a": 1}
//	var dst map[string]int
//	DeepCopy(src, &dst) // 正确用法
//
// @param src 源对象，待拷贝的数据。
// @param dst 目标对象的指针，用于接收拷贝后的数据。
func DeepCopy(src, dst any) {
	// 优化：对常见类型使用类型特化，避免反射开销
	switch s := src.(type) {
	case []int:
		if d, ok := dst.(*[]int); ok {
			if s == nil {
				*d = nil
			} else {
				*d = make([]int, len(s))
				copy(*d, s)
			}
			return
		}
	case []int8:
		if d, ok := dst.(*[]int8); ok {
			if s == nil {
				*d = nil
			} else {
				*d = make([]int8, len(s))
				copy(*d, s)
			}
			return
		}
	case []int16:
		if d, ok := dst.(*[]int16); ok {
			if s == nil {
				*d = nil
			} else {
				*d = make([]int16, len(s))
				copy(*d, s)
			}
			return
		}
	case []int32:
		if d, ok := dst.(*[]int32); ok {
			if s == nil {
				*d = nil
			} else {
				*d = make([]int32, len(s))
				copy(*d, s)
			}
			return
		}
	case []int64:
		if d, ok := dst.(*[]int64); ok {
			if s == nil {
				*d = nil
			} else {
				*d = make([]int64, len(s))
				copy(*d, s)
			}
			return
		}
	case []uint:
		if d, ok := dst.(*[]uint); ok {
			if s == nil {
				*d = nil
			} else {
				*d = make([]uint, len(s))
				copy(*d, s)
			}
			return
		}
	case []uint8:
		if d, ok := dst.(*[]uint8); ok {
			if s == nil {
				*d = nil
			} else {
				*d = make([]uint8, len(s))
				copy(*d, s)
			}
			return
		}
	case []uint16:
		if d, ok := dst.(*[]uint16); ok {
			if s == nil {
				*d = nil
			} else {
				*d = make([]uint16, len(s))
				copy(*d, s)
			}
			return
		}
	case []uint32:
		if d, ok := dst.(*[]uint32); ok {
			if s == nil {
				*d = nil
			} else {
				*d = make([]uint32, len(s))
				copy(*d, s)
			}
			return
		}
	case []uint64:
		if d, ok := dst.(*[]uint64); ok {
			if s == nil {
				*d = nil
			} else {
				*d = make([]uint64, len(s))
				copy(*d, s)
			}
			return
		}
	case []float32:
		if d, ok := dst.(*[]float32); ok {
			if s == nil {
				*d = nil
			} else {
				*d = make([]float32, len(s))
				copy(*d, s)
			}
			return
		}
	case []float64:
		if d, ok := dst.(*[]float64); ok {
			if s == nil {
				*d = nil
			} else {
				*d = make([]float64, len(s))
				copy(*d, s)
			}
			return
		}
	case []string:
		if d, ok := dst.(*[]string); ok {
			if s == nil {
				*d = nil
			} else {
				*d = make([]string, len(s))
				copy(*d, s)
			}
			return
		}
	case []bool:
		if d, ok := dst.(*[]bool); ok {
			if s == nil {
				*d = nil
			} else {
				*d = make([]bool, len(s))
				copy(*d, s)
			}
			return
		}
	}

	// 对于其他类型，使用优化后的反射实现
	v1 := reflect.ValueOf(src)
	v2 := reflect.ValueOf(dst)
	deepCopyValueOptimized(v1, v2)
}

// deepCopyValueOptimized 是优化后的内部实现，相比原版本增加了：
// 1. Map预分配容量
// 2. 基本类型切片的快速复制路径
// 3. 减少不必要的反射调用
func deepCopyValueOptimized(v1, v2 reflect.Value) {
	if !v1.IsValid() || !v2.IsValid() {
		return
	}

	// 解引用指针，直到获取到实际的值
	for v1.Kind() == reflect.Ptr {
		// 如果源指针是 nil，则无需继续
		if v1.IsNil() {
			return
		}
		v1 = v1.Elem()
	}
	for v2.Kind() == reflect.Ptr {
		// 如果目标指针是 nil，则为其分配新内存
		if v2.IsNil() {
			v2.Set(reflect.New(v2.Type().Elem()))
		}
		v2 = v2.Elem()
	}

	// 类型必须匹配才能拷贝
	if v1.Type() != v2.Type() {
		log.Panicf("源类型 %s 与目标类型 %s 不匹配", v1.Type(), v2.Type())
	}

	switch v1.Kind() {
	// 拷贝 Map
	case reflect.Map:
		if v1.IsNil() {
			v2.Set(reflect.Zero(v2.Type()))
			return
		}
		// 优化：预分配Map容量
		v2.Set(reflect.MakeMapWithSize(v1.Type(), v1.Len()))
		// 遍历 Map 并递归拷贝每一个键值对
		for _, k := range v1.MapKeys() {
			val1 := v1.MapIndex(k)
			val2 := reflect.New(val1.Type()).Elem()
			deepCopyValueOptimized(val1, val2)
			v2.SetMapIndex(k, val2)
		}

	// 拷贝 Slice
	case reflect.Slice:
		if v1.IsNil() {
			v2.Set(reflect.Zero(v2.Type()))
			return
		}
		// 为目标 Slice 创建实例
		v2.Set(reflect.MakeSlice(v1.Type(), v1.Len(), v1.Cap()))

		// 优化：基本类型元素使用快速路径
		elemKind := v1.Type().Elem().Kind()
		if isBasicType(elemKind) {
			for i := 0; i < v1.Len(); i++ {
				v2.Index(i).Set(v1.Index(i))
			}
			return
		}

		// 复杂类型元素需要递归拷贝
		for i := 0; i < v1.Len(); i++ {
			deepCopyValueOptimized(v1.Index(i), v2.Index(i))
		}

	// 拷贝 Array
	case reflect.Array:
		// 遍历 Array 并递归拷贝每一个元素
		for i := 0; i < v1.Len(); i++ {
			deepCopyValueOptimized(v1.Index(i), v2.Index(i))
		}

	// 拷贝 Struct
	case reflect.Struct:
		// 遍历 Struct 并递归拷贝每一个字段
		for i := 0; i < v1.NumField(); i++ {
			srcField := v1.Field(i)
			dstField := v2.Field(i)

			// 跳过不可设置的字段（通常是未导出字段）
			if !dstField.CanSet() {
				// 如果无法访问，跳过该字段
				continue
			}

			deepCopyValueOptimized(srcField, dstField)
		}

	// 拷贝 Interface
	case reflect.Interface:
		if v1.IsNil() {
			return
		}
		// 获取接口的实际值
		srcElem := v1.Elem()
		// 创建一个新的目标值，类型与源相同
		dstElem := reflect.New(srcElem.Type()).Elem()
		// 递归拷贝
		deepCopyValueOptimized(srcElem, dstElem)
		// 将拷贝后的值设置给目标接口
		v2.Set(dstElem)

	// 拷贝基本类型
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v2.CanSet() {
			v2.SetInt(v1.Int())
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if v2.CanSet() {
			v2.SetUint(v1.Uint())
		}
	case reflect.Float32, reflect.Float64:
		if v2.CanSet() {
			v2.SetFloat(v1.Float())
		}
	case reflect.Complex64, reflect.Complex128:
		if v2.CanSet() {
			v2.SetComplex(v1.Complex())
		}
	case reflect.String:
		if v2.CanSet() {
			v2.SetString(v1.String())
		}
	case reflect.Bool:
		if v2.CanSet() {
			v2.SetBool(v1.Bool())
		}

	default:
		// 对于未处理的类型，直接 panic
		log.Panicf("未处理的类型: %s", v1.Kind())
	}
}
