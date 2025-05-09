package defaults

import (
	"reflect"
	"strconv"
)

func setDefault(vv reflect.Value, defaultStr string) {
	switch vv.Kind() {
	case reflect.String:
		if defaultStr != "" && vv.String() == "" {
			vv.SetString(defaultStr)
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if defaultStr != "" && vv.Uint() == 0 {
			val, err := strconv.ParseUint(defaultStr, 10, 64)
			if err != nil {
				panic("invalid default value for uint field: " + defaultStr)
			}
			vv.SetUint(val)
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if defaultStr != "" && vv.Int() == 0 {
			val, err := strconv.ParseInt(defaultStr, 10, 64)
			if err != nil {
				panic("invalid default value for int field: " + defaultStr)
			}
			vv.SetInt(val)
		}

	case reflect.Float32, reflect.Float64:
		if defaultStr != "" && vv.Float() == 0 {
			val, err := strconv.ParseFloat(defaultStr, 64)
			if err != nil {
				panic("invalid default value for float field: " + defaultStr)
			}
			vv.SetFloat(val)
		}

	case reflect.Bool:
		if defaultStr != "" && vv.Bool() == false {
			val, err := strconv.ParseBool(defaultStr)
			if err != nil {
				panic("invalid default value for bool field: " + defaultStr)
			}
			vv.SetBool(val)
		}

	case reflect.Ptr:
		if vv.IsNil() {
			vv.Set(reflect.New(vv.Type().Elem()))
		}
		// 修复：递归调用时传递 defaultStr 参数
		for vv.Kind() == reflect.Ptr {
			vv = vv.Elem()
			if vv.Kind() == reflect.Ptr && vv.IsNil() {
				vv.Set(reflect.New(vv.Type().Elem()))
			}
		}
		setDefault(vv, defaultStr)

	case reflect.Struct:
		for i := 0; i < vv.NumField(); i++ {
			field := vv.Field(i)
			fieldType := vv.Type().Field(i)

			if !field.CanSet() {
				continue
			}

			setDefault(field, fieldType.Tag.Get("default"))
		}

	case reflect.Interface:
		// 跳过 interface 类型，因为 interface 本身没有默认值
		return

	case reflect.Slice: // 新增对 slice 类型的支持
		if vv.IsNil() { // 如果 slice 未初始化，则初始化为空切片
			vv.Set(reflect.MakeSlice(vv.Type(), 0, 0))
		}
		return

	default:
		panic("unknown kind " + vv.Kind().String())
	}
}

func SetDefaults(value interface{}) {
	setDefault(reflect.ValueOf(value), "")
}