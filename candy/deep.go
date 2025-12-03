package candy

import (
	"reflect"
	"unsafe"

	"github.com/lazygophers/log"
)

// Comparable 定义可比较的类型约束
type Comparable interface {
	comparable
}

// Copyable 定义可复制的基本类型约束
type Copyable interface {
	~bool | ~string |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~complex64 | ~complex128
}

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
			srcField := v1.Field(i)
			dstField := v2.Field(i)

			// 跳过不可设置的字段（通常是未导出字段）
			if !dstField.CanSet() {
				// 对于未导出字段，尝试使用 unsafe 包强制设置
				if srcField.CanInterface() {
					dstField = reflect.NewAt(dstField.Type(), unsafe.Pointer(dstField.UnsafeAddr())).Elem()
				} else {
					// 如果无法访问，跳过该字段
					continue
				}
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

	case reflect.Invalid:
		// 无效类型，不处理

	default:
		// 对于未处理的类型，直接 panic
		log.Panicf("未处理的类型: %s", v1.Kind())
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
	v1 := reflect.ValueOf(src)
	v2 := reflect.ValueOf(dst)
	deepCopyValue(v1, v2)
}

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

// TypedSliceCopy 类型安全的切片复制
func TypedSliceCopy[T any](src []T) []T {
	if src == nil {
		return nil
	}

	dst := make([]T, len(src))

	// 检查是否是基本类型
	if len(src) > 0 {
		switch any(src[0]).(type) {
		case bool, string,
			int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64,
			float32, float64,
			complex64, complex128:
			// 对于基本类型，使用快速复制
			copy(dst, src)
			return dst
		}
	}

	// 对于复杂类型，逐个深度复制
	for i := range src {
		DeepCopy(src[i], &dst[i])
	}
	return dst
}

// TypedMapCopy 类型安全的 map 复制
func TypedMapCopy[K comparable, V any](src map[K]V) map[K]V {
	if src == nil {
		return nil
	}

	dst := make(map[K]V, len(src))

	// 检查值类型是否是基本类型
	var sampleValue V
	switch any(sampleValue).(type) {
	case bool, string,
		int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64,
		complex64, complex128:
		// 对于基本类型，直接复制
		for k, v := range src {
			dst[k] = v
		}
		return dst
	}

	// 对于复杂类型，深度复制值
	for k, v := range src {
		var dstV V
		DeepCopy(v, &dstV)
		dst[k] = dstV
	}
	return dst
}

// GenericSliceEqual 类型安全的切片比较
func GenericSliceEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	// 检查是否指向相同内存
	if len(a) > 0 && len(b) > 0 {
		aPtr := (*reflect.SliceHeader)(unsafe.Pointer(&a)).Data
		bPtr := (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data
		if aPtr == bPtr {
			return true
		}
	}

	// 逐个比较元素
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// MapEqual 类型安全的 map 比较
func MapEqual[K, V comparable](a, b map[K]V) bool {
	if len(a) != len(b) {
		return false
	}

	for k, av := range a {
		if bv, ok := b[k]; !ok || av != bv {
			return false
		}
	}
	return true
}

// PointerEqual 安全的指针比较
func PointerEqual[T comparable](a, b *T) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// StructEqual 结构体字段比较（需要手动实现）
// 这个函数展示了如何为特定结构体类型实现高性能比较
func StructEqual[T any](a, b T, comparer func(T, T) bool) bool {
	return comparer(a, b)
}

// Clone 通用克隆函数，根据类型选择最佳策略
func Clone[T any](src T) T {
	var dst T
	DeepCopy(src, &dst)
	return dst
}

// CloneSlice 切片克隆的便捷函数
func CloneSlice[T any](src []T) []T {
	return TypedSliceCopy(src)
}

// CloneMap map 克隆的便捷函数
func CloneMap[K comparable, V any](src map[K]V) map[K]V {
	return TypedMapCopy(src)
}

// Equal 通用相等性检查，根据类型选择最佳策略
func Equal[T comparable](a, b T) bool {
	return a == b
}

// EqualSlice 切片相等性检查的便捷函数
func EqualSlice[T comparable](a, b []T) bool {
	return SliceEqual(a, b)
}

// EqualMap map 相等性检查的便捷函数
func EqualMap[K, V comparable](a, b map[K]V) bool {
	return MapEqual(a, b)
}
