package anyx

import "reflect"

func deepValueEqual(v1, v2 reflect.Value) bool {
	if !v1.IsValid() || !v2.IsValid() {
		return v1.IsValid() == v2.IsValid()
	}

	if v1.Type() != v2.Type() {
		return false
	}

	switch v1.Kind() {
	case reflect.Map:
		if v1.IsNil() || v2.IsNil() {
			return v1.IsNil() == v2.IsNil()
		}

		if v1.Len() != v2.Len() {
			return false
		}

		if v1.UnsafePointer() == v2.UnsafePointer() {
			return true
		}

		for _, k := range v1.MapKeys() {
			val1 := v1.MapIndex(k)
			val2 := v2.MapIndex(k)
			if !val1.IsValid() || !val2.IsValid() || !deepValueEqual(val1, val2) {
				return false
			}
		}

		return true

	case reflect.Slice:
		if v1.IsNil() != v2.IsNil() {
			return false
		}

		if v1.Len() != v2.Len() {
			return false
		}

		if v1.UnsafePointer() == v2.UnsafePointer() {
			return true
		}

		for i := 0; i < v1.Len(); i++ {
			if !deepValueEqual(v1.Index(i), v2.Index(i)) {
				return false
			}
		}

		return true

	case reflect.Ptr:
		if v1.IsNil() || v2.IsNil() {
			return v1.IsNil() == v2.IsNil()
		}

		return deepValueEqual(v1.Elem(), v2.Elem())

	case reflect.Array:
		if v1.UnsafePointer() == v2.UnsafePointer() {
			return true
		}

		for i := 0; i < v1.Len(); i++ {
			if !deepValueEqual(v1.Index(i), v2.Index(i)) {
				return false
			}
		}

		return true

	case reflect.Struct:
		for i := 0; i < v1.NumField(); i++ {
			if !deepValueEqual(v1.Field(i), v2.Field(i)) {
				return false
			}
		}

		return true

	case reflect.Interface:
		if v1.IsNil() || v2.IsNil() {
			return v1.IsNil() == v2.IsNil()
		}

		return deepValueEqual(v1.Elem(), v2.Elem())

	default:
		return v1.Interface() == v2.Interface()
	}
}

func DeepEqual[M any](x, y M) bool {
	v1 := reflect.ValueOf(x)
	v2 := reflect.ValueOf(y)

	return deepValueEqual(v1, v2)
}

func deepCopyValue(v1, v2 reflect.Value) {
	if !v1.IsValid() || !v2.IsValid() {
		return
	}

	if v1.Type() != v2.Type() {
		return
	}

	switch v1.Kind() {
	case reflect.Map:
		if v1.IsNil() || v2.IsNil() {
			v2.Set(v1)
			return
		}

		if v1.UnsafePointer() == v2.UnsafePointer() {
			return
		}

		v2.Set(reflect.MakeMap(v1.Type()))
		for _, k := range v1.MapKeys() {
			val1 := v1.MapIndex(k)
			val2 := reflect.New(v2.Type().Elem()).Elem()
			deepCopyValue(val1, val2)
			v2.SetMapIndex(k, val2)
		}

	case reflect.Slice:
		if v1.IsNil() || v2.IsNil() {
			v2.Set(v1)
			return
		}

		if v1.UnsafePointer() == v2.UnsafePointer() {
			return
		}

		v2.Set(reflect.MakeSlice(v1.Type(), v1.Len(), v1.Len()))
		for i := 0; i < v1.Len(); i++ {
			val1 := v1.Index(i)
			val2 := reflect.New(v2.Type().Elem()).Elem()
			deepCopyValue(val1, val2)
			v2.Index(i).Set(val2)
		}

	case reflect.Ptr:
		if v1.IsNil() || v2.IsNil() {
			v2.Set(v1)
			return
		}

		if v1.UnsafePointer() == v2.UnsafePointer() {
			return
		}

		val2 := reflect.New(v2.Type().Elem()).Elem()
		deepCopyValue(v1.Elem(), val2)
		v2.Elem().Set(val2)

	case reflect.Array:
		if v1.UnsafePointer() == v2.UnsafePointer() {
			return
		}

		for i := 0; i < v1.Len(); i++ {
			val1 := v1.Index(i)
			val2 := reflect.New(v2.Type().Elem()).Elem()
			deepCopyValue(val1, val2)
			v2.Index(i).Set(val2)
		}

	case reflect.Struct:
		for i := 0; i < v1.NumField(); i++ {
			val1 := v1.Field(i)
			val2 := reflect.New(v2.Field(i).Type()).Elem()
			deepCopyValue(val1, val2)
			v2.Field(i).Set(val2)
		}

	case reflect.Interface:
		if v1.IsNil() || v2.IsNil() {
			v2.Set(v1)
			return
		}

		val2 := reflect.New(v2.Elem().Type()).Elem()
		deepCopyValue(v1.Elem(), val2)
		v2.Elem().Set(val2)

	default:
		v2.Set(v1)
	}
}

func DeepCopy[M any](src, dst M) {
	v1 := reflect.ValueOf(src)
	v2 := reflect.ValueOf(dst)

	deepCopyValue(v1, v2)
}
