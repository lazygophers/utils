package candy

// 反射版切片工具：仅用于无法使用泛型/约束的场景。
import "reflect"

// RemoveSlice 从源切片中移除指定的元素
// src 是源切片，rm 是要移除的元素切片
// 返回移除指定元素后的新切片
func RemoveSlice(src interface{}, rm interface{}) interface{} {
	// 快速路径：int 类型
	if srcInt, ok := src.([]int); ok {
		if rmInt, ok := rm.([]int); ok {
			m := make(map[int]struct{}, len(rmInt))
			for _, v := range rmInt {
				m[v] = struct{}{}
			}
			result := make([]int, 0, len(srcInt)-len(rmInt)/2)
			for _, v := range srcInt {
				if _, ok := m[v]; !ok {
					result = append(result, v)
				}
			}
			return result
		}
	}

	// 快速路径：string 类型
	if srcStr, ok := src.([]string); ok {
		if rmStr, ok := rm.([]string); ok {
			m := make(map[string]struct{}, len(rmStr))
			for _, v := range rmStr {
				m[v] = struct{}{}
			}
			result := make([]string, 0, len(srcStr)-len(rmStr)/2)
			for _, v := range srcStr {
				if _, ok := m[v]; !ok {
					result = append(result, v)
				}
			}
			return result
		}
	}

	// 反射路径
	at := reflect.TypeOf(src)
	if at.Kind() != reflect.Slice {
		panic("src is not slice")
	}

	bt := reflect.TypeOf(rm)
	if bt.Kind() != reflect.Slice {
		panic("rm is not slice")
	}

	if at.Elem().Kind() != bt.Elem().Kind() {
		panic("src and rm are not same type")
	}

	m := map[interface{}]bool{}
	bv := reflect.ValueOf(rm)
	for i := 0; i < bv.Len(); i++ {
		m[bv.Index(i).Interface()] = true
	}

	av := reflect.ValueOf(src)
	c := reflect.MakeSlice(at, 0, av.Len()-bv.Len()/2) // 预分配容量
	for i := 0; i < av.Len(); i++ {
		if !m[av.Index(i).Interface()] {
			c = reflect.Append(c, av.Index(i))
		}
	}

	return c.Interface()
}

// DiffSlice 比较两个切片的差异
// 返回第一个切片中存在但第二个切片中不存在的元素，以及第二个切片中存在但第一个切片中不存在的元素
func DiffSlice(a interface{}, b interface{}) (interface{}, interface{}) {
	// 快速路径：int 类型
	if aInt, ok := a.([]int); ok {
		if bInt, ok := b.([]int); ok {
			m := make(map[int]struct{}, len(bInt))
			for _, v := range bInt {
				m[v] = struct{}{}
			}

			aOnly := make([]int, 0, len(aInt)/2)
			bOnly := make([]int, 0, len(bInt)/2)

			for _, v := range aInt {
				if _, exists := m[v]; !exists {
					aOnly = append(aOnly, v)
				} else {
					delete(m, v)
				}
			}

			for v := range m {
				bOnly = append(bOnly, v)
			}

			return aOnly, bOnly
		}
	}

	// 快速路径：string 类型
	if aStr, ok := a.([]string); ok {
		if bStr, ok := b.([]string); ok {
			m := make(map[string]struct{}, len(bStr))
			for _, v := range bStr {
				m[v] = struct{}{}
			}

			aOnly := make([]string, 0, len(aStr)/2)
			bOnly := make([]string, 0, len(bStr)/2)

			for _, v := range aStr {
				if _, exists := m[v]; !exists {
					aOnly = append(aOnly, v)
				} else {
					delete(m, v)
				}
			}

			for v := range m {
				bOnly = append(bOnly, v)
			}

			return aOnly, bOnly
		}
	}

	// 反射路径
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	if av.Kind() != reflect.Slice {
		panic("a is not slice")
	}
	if bv.Kind() != reflect.Slice {
		panic("b is not slice")
	}

	if av.Type().Elem().Kind() != bv.Type().Elem().Kind() {
		panic("a and b are not same type")
	}

	m := map[interface{}]reflect.Value{}
	for i := 0; i < bv.Len(); i++ {
		m[bv.Index(i).Interface()] = bv.Index(i)
	}

	c := reflect.MakeSlice(av.Type(), 0, av.Len()/2) // 预分配
	d := reflect.MakeSlice(bv.Type(), 0, bv.Len()/2)

	for i := 0; i < av.Len(); i++ {
		elem := av.Index(i)
		if !m[elem.Interface()].IsValid() {
			c = reflect.Append(c, elem)
		} else {
			delete(m, elem.Interface())
		}
	}

	for _, value := range m {
		d = reflect.Append(d, value)
	}

	return c.Interface(), d.Interface()
}
