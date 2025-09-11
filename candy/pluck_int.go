package candy

import (
	"fmt"
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

// PluckInt 从结构体切片中提取指定字段的 int 值
func PluckInt(list interface{}, fieldName string) []int {
	return pluck(list, fieldName, []int{}).([]int)
}