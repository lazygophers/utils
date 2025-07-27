package anyx

import (
	"reflect"

	"github.com/lazygophers/log"
)

// deepValueEqual 深度比较两个 reflect.Value 是否相等。
//
// 此函数为内部实现，递归地比较不同类型的值。
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
		return v1.Interface() == v2.Interface()
	}
}

// DeepEqual 是一个泛型函数，用于深度比较两个任意类型的值 x 和 y 是否相等。
//
// 它利用反射（reflection）来递归地检查所有字段和元素。
func DeepEqual[M any](x, y M) bool {
	v1 := reflect.ValueOf(x)
	v2 := reflect.ValueOf(y)
	return deepValueEqual(v1, v2)
}

// deepCopyValue 是实际执行深拷贝逻辑的内部函数。
//
// 它处理不同类型的 `reflect.Value`，并将其内容从 v1 拷贝到 v2。
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

	// 确保解引用后值仍然有效
	if v1.Kind() == reflect.Invalid || v2.Kind() == reflect.Invalid {
		return
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
			deepCopyValue(v1.Field(i), v2.Field(i))
		}

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

	case reflect.Invalid:
		// 无效类型，不处理

	default:
		// 对于未处理的类型，直接 panic
		log.Panicf("未处理的类型: %s", v1.Kind())
	}
}

// DeepCopy 是一个泛型函数，用于将 src 的内容深拷贝到 dst。
//
// **注意**：dst 必须是一个指向有效内存的指针，否则会引发 panic。
//
// 此函数通过反射递归地复制所有字段和元素，确保拷贝后的对象是完全独立的。
func DeepCopy[M any](src, dst M) {
	v1 := reflect.ValueOf(src)
	v2 := reflect.ValueOf(dst)
	deepCopyValue(v1, v2)
}
