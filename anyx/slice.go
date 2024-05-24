package anyx

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"reflect"
)

func pluck(list interface{}, fieldName string, deferVal interface{}) interface{} {
	v := reflect.ValueOf(list)
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		if v.Len() == 0 {
			return deferVal
		}

		ev := v.Type().Elem()
		evs := ev
		for evs.Kind() == reflect.Ptr {
			evs = evs.Elem()
		}

		switch evs.Kind() {
		case reflect.Struct:
			// 如果是 struct，则取出 fieldName 的值
			field, ok := evs.FieldByName(fieldName)
			if !ok {
				panic(fmt.Sprintf("field %s not found", fieldName))
			}

			result := reflect.MakeSlice(reflect.SliceOf(field.Type), v.Len(), v.Len())

			for i := 0; i < v.Len(); i++ {
				ev := v.Index(i)
				for ev.Kind() == reflect.Ptr {
					ev = ev.Elem()
				}
				if ev.Kind() != reflect.Struct {
					panic("element is not a struct")
				}
				if !ev.IsValid() {
					continue
				}
				result.Index(i).Set(ev.FieldByIndex(field.Index))
			}

			return result.Interface()
		case reflect.Slice, reflect.Array:
			var ev reflect.Value
			var c int
			for i := 0; i < v.Len(); i++ {
				ev = v.Index(i)
				for i := 0; i < ev.Len(); i++ {
					c += ev.Index(i).Len()
				}
			}

			result := reflect.MakeSlice(ev.Type(), c, c)
			var idx int
			for i := 0; i < v.Len(); i++ {
				ev := v.Index(i)
				for i := 0; i < ev.Len(); i++ {
					result.Index(idx).Set(ev.Index(i))
					idx++
				}
			}

			return result.Interface()
		default:
			panic("list element type is not supported")
		}

	default:
		panic("list must be an array or slice")
	}
}

func PluckInt(list interface{}, fieldName string) []int {
	return pluck(list, fieldName, []int{}).([]int)
}

func PluckInt32(list interface{}, fieldName string) []int32 {
	return pluck(list, fieldName, []int32{}).([]int32)
}

func PluckUint32(list interface{}, fileName string) []uint32 {
	return pluck(list, fileName, []uint32{}).([]uint32)
}

func PluckInt64(list interface{}, fieldName string) []int64 {
	return pluck(list, fieldName, []int64{}).([]int64)
}

func PluckUint64(list interface{}, fieldName string) []uint64 {
	return pluck(list, fieldName, []uint64{}).([]uint64)
}

func PluckString(list interface{}, fieldName string) []string {
	return pluck(list, fieldName, []string{}).([]string)
}

func PluckStringSlice(list interface{}, fieldName string) [][]string {
	return pluck(list, fieldName, [][]string{}).([][]string)
}

// DiffSlice 传入两个slice
// 如果 a 或者 b 不为 slice 会 panic
// 如果 a 与 b 的元素类型不一致，也会 panic
// 返回的第一个参数为 a 比 b 多的，类型为 a 的类型
// 返回的第二个参数为 b 比 a 多的，类型为 b 的类型
func DiffSlice(a interface{}, b interface{}) (interface{}, interface{}) {
	at := reflect.TypeOf(a)
	if at.Kind() != reflect.Slice {
		panic("a is not slice")
	}

	bt := reflect.TypeOf(b)
	if bt.Kind() != reflect.Slice {
		panic("b is not slice")
	}

	atm := at.Elem()
	btm := bt.Elem()

	if atm.Kind() != btm.Kind() {
		panic("a and b are not same type")
	}

	m := map[interface{}]reflect.Value{}

	bv := reflect.ValueOf(b)
	for i := 0; i < bv.Len(); i++ {
		m[bv.Index(i).Interface()] = bv.Index(i)
	}

	c := reflect.MakeSlice(at, 0, 0)
	d := reflect.MakeSlice(bt, 0, 0)
	av := reflect.ValueOf(a)
	for i := 0; i < av.Len(); i++ {
		if !m[av.Index(i).Interface()].IsValid() {
			c = reflect.Append(c, av.Index(i))
		} else {
			delete(m, av.Index(i).Interface())
		}
	}

	for _, value := range m {
		d = reflect.Append(d, value)
	}

	return c.Interface(), d.Interface()
}

// RemoveSlice 传入两个slice
// 如果 src 或者 rm 不为 slice 会 panic
// 如果 src 与 rm 的元素类型不一致，也会 panic
// 返回的第一个参数为 src 中不在 rm 中的元素，数据类型与 src 一致
func RemoveSlice(src interface{}, rm interface{}) interface{} {
	at := reflect.TypeOf(src)
	if at.Kind() != reflect.Slice {
		panic("a is not slice")
	}

	bt := reflect.TypeOf(src)
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

func KeyBy(list interface{}, fieldName string) interface{} {
	lv := reflect.ValueOf(list)

	switch lv.Kind() {
	case reflect.Slice, reflect.Array:
	default:
		panic("list required slice or array type")
	}

	ev := lv.Type().Elem()
	evs := ev
	for evs.Kind() == reflect.Ptr {
		evs = evs.Elem()
	}

	if evs.Kind() != reflect.Struct {
		panic("list element is not struct")
	}

	field, ok := evs.FieldByName(fieldName)
	if !ok {
		panic(fmt.Sprintf("field %s not found", fieldName))
	}

	m := reflect.MakeMapWithSize(reflect.MapOf(field.Type, ev), lv.Len())
	for i := 0; i < lv.Len(); i++ {
		elem := lv.Index(i)
		elemStruct := elem
		for elemStruct.Kind() == reflect.Ptr {
			elemStruct = elemStruct.Elem()
		}

		// 如果是nil的，意味着key和value同时不存在，所以跳过不处理
		if !elemStruct.IsValid() {
			continue
		}

		if elemStruct.Kind() != reflect.Struct {
			panic("element not struct")
		}

		m.SetMapIndex(elemStruct.FieldByIndex(field.Index), elem)
	}

	return m.Interface()
}

func KeyByUint64[M any](list []*M, fieldName string) map[uint64]*M {
	if len(list) == 0 {
		return map[uint64]*M{}
	}

	lv := reflect.ValueOf(list)

	ev := lv.Type().Elem()
	evs := ev
	for evs.Kind() == reflect.Ptr {
		evs = evs.Elem()
	}

	field, ok := evs.FieldByName(fieldName)
	if !ok {
		panic(fmt.Sprintf("field %s not found", fieldName))
	}

	m := make(map[uint64]*M, lv.Len())
	for i := 0; i < lv.Len(); i++ {
		elem := lv.Index(i)
		elemStruct := elem
		for elemStruct.Kind() == reflect.Ptr {
			elemStruct = elemStruct.Elem()
		}

		// 如果是nil的，意味着key和value同时不存在，所以跳过不处理
		if !elemStruct.IsValid() {
			continue
		}

		if elemStruct.Kind() != reflect.Struct {
			panic("element not struct")
		}

		m[elemStruct.FieldByIndex(field.Index).Uint()] = elem.Interface().(*M)
		//m.SetMapIndex(elemStruct.FieldByIndex(field.Index), elem)
	}

	return m
}

func KeyByInt64[M any](list []*M, fieldName string) map[int64]*M {
	if len(list) == 0 {
		return map[int64]*M{}
	}

	lv := reflect.ValueOf(list)

	ev := lv.Type().Elem()
	evs := ev
	for evs.Kind() == reflect.Ptr {
		evs = evs.Elem()
	}

	field, ok := evs.FieldByName(fieldName)
	if !ok {
		panic(fmt.Sprintf("field %s not found", fieldName))
	}

	m := make(map[int64]*M, lv.Len())
	for i := 0; i < lv.Len(); i++ {
		elem := lv.Index(i)
		elemStruct := elem
		for elemStruct.Kind() == reflect.Ptr {
			elemStruct = elemStruct.Elem()
		}

		// 如果是nil的，意味着key和value同时不存在，所以跳过不处理
		if !elemStruct.IsValid() {
			continue
		}

		if elemStruct.Kind() != reflect.Struct {
			panic("element not struct")
		}

		m[elemStruct.FieldByIndex(field.Index).Int()] = elem.Interface().(*M)
		//m.SetMapIndex(elemStruct.FieldByIndex(field.Index), elem)
	}

	return m
}

func Slice2Map[M constraints.Ordered](list []M) map[M]bool {
	m := make(map[M]bool, len(list))

	for _, v := range list {
		m[v] = true
	}

	return m
}
