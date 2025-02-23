package defaults

import (
	"reflect"
	"strconv"
)

func setDefault(vv reflect.Value) {
	for vv.Type().Kind() == reflect.Ptr {
		vv = vv.Elem()
	}

	if vv.Type().Kind() != reflect.Struct {
		panic("value must be a struct or ptr")
	}

	for i := 0; i < vv.NumField(); i++ {
		field := vv.Field(i)
		fieldType := vv.Type().Field(i)

		defaultStr := fieldType.Tag.Get("default")
		if defaultStr == "-" {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			if defaultStr != "" && field.String() == "" {
				field.SetString(defaultStr)
			}

		case reflect.Uint,
			reflect.Uint8,
			reflect.Uint16,
			reflect.Uint32,
			reflect.Uint64:
			if defaultStr != "" && field.Uint() == 0 {
				val, err := strconv.ParseUint(defaultStr, 10, 64)
				if err == nil {
					field.SetUint(val)
				}
			}

		case reflect.Int,
			reflect.Int8,
			reflect.Int16,
			reflect.Int32,
			reflect.Int64:
			if defaultStr != "" && field.Int() == 0 {
				val, err := strconv.ParseInt(defaultStr, 10, 64)
				if err == nil {
					field.SetInt(val)
				}
			}

		case reflect.Float32,
			reflect.Float64:
			if defaultStr != "" && field.Float() == 0 {
				val, err := strconv.ParseFloat(defaultStr, 64)
				if err == nil {
					field.SetFloat(val)
				}
			}

		case reflect.Bool:
			if defaultStr != "" && field.Bool() == false {
				val, err := strconv.ParseBool(defaultStr)
				if err == nil {
					field.SetBool(val)
				}
			}

		case reflect.Ptr:
			if !field.CanSet() {
				continue
			}

			if field.IsNil() {
				field.Set(reflect.New(field.Type().Elem()))

				field = field.Elem()
				field.Set(reflect.New(field.Type()).Elem())
			}

			setDefault(field)

		case reflect.Struct:
			setDefault(field)

		case reflect.Map:
		// do nothing

		case reflect.Interface:
			// do nothing

		case reflect.Slice:
			// do nothing

			//default:
			//	if !field.CanSet() {
			//		continue
			//	}

			//panic("unknown kind " + field.Kind().String())
		}
	}
}

func SetDefaults(value interface{}) {
	setDefault(reflect.ValueOf(value))
}
