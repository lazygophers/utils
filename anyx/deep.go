package anyx

import (
	"reflect"

	"github.com/lazygophers/log"
)

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

	for v1.Kind() == reflect.Ptr {
		v1 = v1.Elem()
	}

	for v2.Kind() == reflect.Ptr {
		v2 = v2.Elem()
	}

	if v1.Kind() == reflect.Invalid {
		return
	}

	if v2.Kind() == reflect.Invalid {
		return
	}

	if v1.Kind() != v2.Kind() {
		log.Panicf("source kind is %s different than destination kind %s", v1.Kind(), v2.Kind())
	}

	switch v1.Kind() {
	case reflect.Map:
		if v1.IsNil() {
			v2.SetZero()
			return
		}

		for _, k := range v1.MapKeys() {
			val1 := v1.MapIndex(k)
			val2 := v2.MapIndex(k)

			if !val1.IsValid() || !val2.IsValid() {
				continue
			}

			if val2.IsNil() || val2.IsZero() {
				v2.SetMapIndex(k, reflect.New(val2.Type().Elem()).Elem())
				val2 = v2.MapIndex(k)
			}

			deepCopyValue(val1, val2)
		}

	case reflect.Slice:
		if v1.IsNil() {
			return
		}

		for i := 0; i < v1.Len(); i++ {
			val1 := v1.Index(i)
			for v2.Len() <= i {
				v2 = reflect.Append(v2, reflect.New(v2.Type().Elem()).Elem())
			}
			val2 := v2.Index(i)

			if !val1.IsValid() || !val2.IsValid() {
				continue
			}

			val2.Set(reflect.New(val2.Type().Elem()))
			deepCopyValue(val1, val2.Elem())
		}

	case reflect.Array:
		for i := 0; i < v1.Len(); i++ {
			val1 := v1.Index(i)
			val2 := v2.Index(i)

			if !val1.IsValid() || !val2.IsValid() {
				continue
			}

			if val2.Kind() == reflect.Array {
				deepCopyValue(val1, val2)
			} else {
				val2.Set(val1)
			}
		}

	case reflect.Struct:
		for i := 0; i < v1.NumField(); i++ {
			val1 := v1.Field(i)
			val2 := v2.Field(i)

			if !val1.IsValid() || !val2.IsValid() {
				continue
			}

			deepCopyValue(val1, val2)
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v2.IsZero() && v2.CanSet() {
			v2.SetInt(v1.Int())
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if v2.IsZero() && v2.CanSet() {
			v2.SetUint(v1.Uint())
		}

	case reflect.Float32, reflect.Float64:
		if v2.IsZero() && v2.CanSet() {
			v2.SetFloat(v1.Float())
		}

	case reflect.Complex64, reflect.Complex128:
		if v2.IsZero() && v2.CanSet() {
			v2.SetComplex(v1.Complex())
		}

	case reflect.String:
		if v2.IsZero() && v2.CanSet() {
			v2.SetString(v1.String())
		}

	case reflect.Bool:
		if !v2.Bool() {
			v2.SetBool(v1.Bool())
		}

	case reflect.Invalid:

	default:
		log.Panicf("unhandled type %s", v1.Kind())
		v2.Set(v1)
	}
}

func DeepCopy[M any](src, dst M) {
	v1 := reflect.ValueOf(src)
	v2 := reflect.ValueOf(dst)

	reflect.Copy(v1, v2)
}
